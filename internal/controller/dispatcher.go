/*
Copyright 2023 IBM Corporation.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controller

import (
	"context"
	"errors"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"

	mcadv1beta1 "github.com/tardieu/mcad/api/v1beta1"
)

type Dispatcher struct {
	AppWrapperReconciler
	Events chan event.GenericEvent
}

func (r *Dispatcher) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	// req == "*/*", attempt to select and dispatch one appWrapper
	if req.Namespace == "*" && req.Name == "*" {
		return r.dispatch(ctx)
	}

	// get deep copy of AppWrapper object in reconciler cache
	appWrapper := &mcadv1beta1.AppWrapper{}
	if err := r.Get(ctx, req.NamespacedName, appWrapper); err != nil {
		// no such AppWrapper, nothing to reconcile, not an error
		return ctrl.Result{}, nil
	}

	// abort and requeue reconciliation if reconciler cache is stale
	if err := r.checkCachedPhase(appWrapper); err != nil {
		return ctrl.Result{}, err
	}

	// handle deletion
	if !appWrapper.DeletionTimestamp.IsZero() {
		if appWrapper.Status.Phase == mcadv1beta1.Empty {
			// remove finalizer
			if controllerutil.RemoveFinalizer(appWrapper, finalizer) {
				if err := r.Update(ctx, appWrapper); err != nil {
					return ctrl.Result{}, err
				}
			}
			r.deleteCachedPhase(appWrapper) // remove appWrapper from cache
			r.triggerDispatch()             // cluster may have more available capacity
		}
		return ctrl.Result{}, nil
	}

	// handle other phases
	switch appWrapper.Spec.DispatcherStatus.Phase {
	case mcadv1beta1.Succeeded, mcadv1beta1.Queued:
		r.triggerDispatch()

	case mcadv1beta1.Dispatching:
		if appWrapper.Status.Phase == mcadv1beta1.Dispatching {
			return r.update(ctx, appWrapper, mcadv1beta1.Running)
		}

	case mcadv1beta1.Requeuing:
		if appWrapper.Status.Phase == mcadv1beta1.Empty {
			return r.update(ctx, appWrapper, mcadv1beta1.Queued)
		}

	case mcadv1beta1.Running:
		if appWrapper.Status.Phase == mcadv1beta1.Succeeded {
			return r.update(ctx, appWrapper, mcadv1beta1.Succeeded)
		}
		if appWrapper.Status.Phase == mcadv1beta1.Failed {
			return r.update(ctx, appWrapper, mcadv1beta1.Failed)
		}
		if appWrapper.Status.Phase == mcadv1beta1.Errored {
			return r.requeueOrFail(ctx, appWrapper)
		}

	case mcadv1beta1.Empty:
		// add finalizer
		if controllerutil.AddFinalizer(appWrapper, finalizer) {
			if err := r.Update(ctx, appWrapper); err != nil {
				return ctrl.Result{}, err
			}
		}
		// set queued status only after adding finalizer
		return r.update(ctx, appWrapper, mcadv1beta1.Queued)
	}
	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *Dispatcher) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&mcadv1beta1.AppWrapper{}).
		Watches(&mcadv1beta1.ClusterInfo{}, handler.EnqueueRequestsFromMapFunc(r.clusterInfoMapFunc)).
		WatchesRawSource(&source.Channel{Source: r.Events}, &handler.EnqueueRequestForObject{}).
		Complete(r)
}

// Trigger dispatch on cluster capacity change
func (r *Dispatcher) clusterInfoMapFunc(ctx context.Context, obj client.Object) []reconcile.Request {
	r.triggerDispatch()
	return nil
}

// Update AppWrapper status
func (r *Dispatcher) update(ctx context.Context, appWrapper *mcadv1beta1.AppWrapper, phase mcadv1beta1.AppWrapperPhase) (ctrl.Result, error) {
	log := log.FromContext(ctx)
	now := metav1.Now()
	// log transition
	transition := mcadv1beta1.AppWrapperTransition{Time: now, Phase: phase}
	appWrapper.Spec.DispatcherStatus.Transitions = append(appWrapper.Spec.DispatcherStatus.Transitions, transition)
	appWrapper.Spec.DispatcherStatus.Phase = phase
	if err := r.Update(ctx, appWrapper); err != nil {
		return ctrl.Result{}, err // etcd update failed, abort and requeue reconciliation
	}
	log.Info(string(phase))
	// cache AppWrapper status
	r.addCachedPhase(appWrapper)
	return ctrl.Result{}, nil
}

// Set requeuing or failed status depending on retry count
func (r *Dispatcher) requeueOrFail(ctx context.Context, appWrapper *mcadv1beta1.AppWrapper) (ctrl.Result, error) {
	if appWrapper.Spec.DispatcherStatus.Requeued < appWrapper.Spec.Scheduling.Requeuing.MaxNumRequeuings {
		appWrapper.Spec.DispatcherStatus.Requeued += 1
		return r.update(ctx, appWrapper, mcadv1beta1.Requeuing)
	}
	return r.update(ctx, appWrapper, mcadv1beta1.Failed)
}

// Trigger dispatch
func (r *Dispatcher) triggerDispatch() {
	select {
	case r.Events <- event.GenericEvent{Object: &metav1.PartialObjectMetadata{ObjectMeta: metav1.ObjectMeta{Namespace: "*", Name: "*"}}}:
	default:
	}
}

// Attempt to select and dispatch one appWrapper
func (r *Dispatcher) dispatch(ctx context.Context) (ctrl.Result, error) {
	appWrapper, err := r.selectForDispatch(ctx)
	if err != nil {
		return ctrl.Result{}, err
	}
	if appWrapper == nil { // no appWrapper eligible for dispatch
		return ctrl.Result{RequeueAfter: dispatchDelay}, nil // retry to dispatch later
	}
	// abort and requeue reconciliation if reconciler cache is stale
	if err := r.checkCachedPhase(appWrapper); err != nil {
		return ctrl.Result{}, err
	}
	if appWrapper.Spec.DispatcherStatus.Phase != mcadv1beta1.Queued {
		// this check should be redundant but better be defensive
		return ctrl.Result{}, errors.New("not queued")
	}
	// set dispatching status
	if _, err := r.update(ctx, appWrapper, mcadv1beta1.Dispatching); err != nil {
		return ctrl.Result{}, err
	}
	return ctrl.Result{Requeue: true}, nil // requeue to continue to dispatch queued appWrappers
}
