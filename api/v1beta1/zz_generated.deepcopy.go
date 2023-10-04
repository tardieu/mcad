//go:build !ignore_autogenerated
// +build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppWrapper) DeepCopyInto(out *AppWrapper) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppWrapper.
func (in *AppWrapper) DeepCopy() *AppWrapper {
	if in == nil {
		return nil
	}
	out := new(AppWrapper)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AppWrapper) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppWrapperDispatcherStatus) DeepCopyInto(out *AppWrapperDispatcherStatus) {
	*out = *in
	in.LastDispatchingTime.DeepCopyInto(&out.LastDispatchingTime)
	in.LastRequeuingTime.DeepCopyInto(&out.LastRequeuingTime)
	if in.Transitions != nil {
		in, out := &in.Transitions, &out.Transitions
		*out = make([]AppWrapperTransition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppWrapperDispatcherStatus.
func (in *AppWrapperDispatcherStatus) DeepCopy() *AppWrapperDispatcherStatus {
	if in == nil {
		return nil
	}
	out := new(AppWrapperDispatcherStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppWrapperList) DeepCopyInto(out *AppWrapperList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AppWrapper, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppWrapperList.
func (in *AppWrapperList) DeepCopy() *AppWrapperList {
	if in == nil {
		return nil
	}
	out := new(AppWrapperList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AppWrapperList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppWrapperResources) DeepCopyInto(out *AppWrapperResources) {
	*out = *in
	if in.GenericItems != nil {
		in, out := &in.GenericItems, &out.GenericItems
		*out = make([]GenericItem, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppWrapperResources.
func (in *AppWrapperResources) DeepCopy() *AppWrapperResources {
	if in == nil {
		return nil
	}
	out := new(AppWrapperResources)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppWrapperRunnerStatus) DeepCopyInto(out *AppWrapperRunnerStatus) {
	*out = *in
	in.LastRunningTime.DeepCopyInto(&out.LastRunningTime)
	in.LastRequeuingTime.DeepCopyInto(&out.LastRequeuingTime)
	if in.Transitions != nil {
		in, out := &in.Transitions, &out.Transitions
		*out = make([]AppWrapperTransition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppWrapperRunnerStatus.
func (in *AppWrapperRunnerStatus) DeepCopy() *AppWrapperRunnerStatus {
	if in == nil {
		return nil
	}
	out := new(AppWrapperRunnerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppWrapperSpec) DeepCopyInto(out *AppWrapperSpec) {
	*out = *in
	out.Scheduling = in.Scheduling
	in.Resources.DeepCopyInto(&out.Resources)
	in.DispatcherStatus.DeepCopyInto(&out.DispatcherStatus)
	out.Sustainable = in.Sustainable
	if in.DispatchingGates != nil {
		in, out := &in.DispatchingGates, &out.DispatchingGates
		*out = make([]DispatchingGate, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppWrapperSpec.
func (in *AppWrapperSpec) DeepCopy() *AppWrapperSpec {
	if in == nil {
		return nil
	}
	out := new(AppWrapperSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppWrapperStatus) DeepCopyInto(out *AppWrapperStatus) {
	*out = *in
	in.RunnerStatus.DeepCopyInto(&out.RunnerStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppWrapperStatus.
func (in *AppWrapperStatus) DeepCopy() *AppWrapperStatus {
	if in == nil {
		return nil
	}
	out := new(AppWrapperStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppWrapperTransition) DeepCopyInto(out *AppWrapperTransition) {
	*out = *in
	in.Time.DeepCopyInto(&out.Time)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppWrapperTransition.
func (in *AppWrapperTransition) DeepCopy() *AppWrapperTransition {
	if in == nil {
		return nil
	}
	out := new(AppWrapperTransition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterDecision) DeepCopyInto(out *ClusterDecision) {
	*out = *in
	out.TargetCluster = in.TargetCluster
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterDecision.
func (in *ClusterDecision) DeepCopy() *ClusterDecision {
	if in == nil {
		return nil
	}
	out := new(ClusterDecision)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterInfo) DeepCopyInto(out *ClusterInfo) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterInfo.
func (in *ClusterInfo) DeepCopy() *ClusterInfo {
	if in == nil {
		return nil
	}
	out := new(ClusterInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterInfo) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterInfoList) DeepCopyInto(out *ClusterInfoList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterInfo, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterInfoList.
func (in *ClusterInfoList) DeepCopy() *ClusterInfoList {
	if in == nil {
		return nil
	}
	out := new(ClusterInfoList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterInfoList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterInfoSpec) DeepCopyInto(out *ClusterInfoSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterInfoSpec.
func (in *ClusterInfoSpec) DeepCopy() *ClusterInfoSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterInfoSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterInfoStatus) DeepCopyInto(out *ClusterInfoStatus) {
	*out = *in
	if in.Capacity != nil {
		in, out := &in.Capacity, &out.Capacity
		*out = make(v1.ResourceList, len(*in))
		for key, val := range *in {
			(*out)[key] = val.DeepCopy()
		}
	}
	in.Time.DeepCopyInto(&out.Time)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterInfoStatus.
func (in *ClusterInfoStatus) DeepCopy() *ClusterInfoStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterInfoStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterReference) DeepCopyInto(out *ClusterReference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterReference.
func (in *ClusterReference) DeepCopy() *ClusterReference {
	if in == nil {
		return nil
	}
	out := new(ClusterReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSchedulingSpec) DeepCopyInto(out *ClusterSchedulingSpec) {
	*out = *in
	out.PolicyResult = in.PolicyResult
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSchedulingSpec.
func (in *ClusterSchedulingSpec) DeepCopy() *ClusterSchedulingSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterSchedulingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomPodResource) DeepCopyInto(out *CustomPodResource) {
	*out = *in
	if in.Requests != nil {
		in, out := &in.Requests, &out.Requests
		*out = make(v1.ResourceList, len(*in))
		for key, val := range *in {
			(*out)[key] = val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomPodResource.
func (in *CustomPodResource) DeepCopy() *CustomPodResource {
	if in == nil {
		return nil
	}
	out := new(CustomPodResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GenericItem) DeepCopyInto(out *GenericItem) {
	*out = *in
	if in.CustomPodResources != nil {
		in, out := &in.CustomPodResources, &out.CustomPodResources
		*out = make([]CustomPodResource, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.GenericTemplate.DeepCopyInto(&out.GenericTemplate)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GenericItem.
func (in *GenericItem) DeepCopy() *GenericItem {
	if in == nil {
		return nil
	}
	out := new(GenericItem)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RequeuingSpec) DeepCopyInto(out *RequeuingSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RequeuingSpec.
func (in *RequeuingSpec) DeepCopy() *RequeuingSpec {
	if in == nil {
		return nil
	}
	out := new(RequeuingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SchedulingSpec) DeepCopyInto(out *SchedulingSpec) {
	*out = *in
	out.Requeuing = in.Requeuing
	out.ClusterScheduling = in.ClusterScheduling
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SchedulingSpec.
func (in *SchedulingSpec) DeepCopy() *SchedulingSpec {
	if in == nil {
		return nil
	}
	out := new(SchedulingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SustainableSpec) DeepCopyInto(out *SustainableSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SustainableSpec.
func (in *SustainableSpec) DeepCopy() *SustainableSpec {
	if in == nil {
		return nil
	}
	out := new(SustainableSpec)
	in.DeepCopyInto(out)
	return out
}
