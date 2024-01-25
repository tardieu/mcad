//go:build !ignore_autogenerated

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

package v1alpha1

import (
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BoxedJob) DeepCopyInto(out *BoxedJob) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BoxedJob.
func (in *BoxedJob) DeepCopy() *BoxedJob {
	if in == nil {
		return nil
	}
	out := new(BoxedJob)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BoxedJob) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BoxedJobComponent) DeepCopyInto(out *BoxedJobComponent) {
	*out = *in
	if in.Topology != nil {
		in, out := &in.Topology, &out.Topology
		*out = make([]BoxedJobPodSet, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BoxedJobComponent.
func (in *BoxedJobComponent) DeepCopy() *BoxedJobComponent {
	if in == nil {
		return nil
	}
	out := new(BoxedJobComponent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BoxedJobList) DeepCopyInto(out *BoxedJobList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BoxedJob, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BoxedJobList.
func (in *BoxedJobList) DeepCopy() *BoxedJobList {
	if in == nil {
		return nil
	}
	out := new(BoxedJobList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BoxedJobList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BoxedJobPodSet) DeepCopyInto(out *BoxedJobPodSet) {
	*out = *in
	if in.Requests != nil {
		in, out := &in.Requests, &out.Requests
		*out = make(v1.ResourceList, len(*in))
		for key, val := range *in {
			(*out)[key] = val.DeepCopy()
		}
	}
	if in.Limits != nil {
		in, out := &in.Limits, &out.Limits
		*out = make(v1.ResourceList, len(*in))
		for key, val := range *in {
			(*out)[key] = val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BoxedJobPodSet.
func (in *BoxedJobPodSet) DeepCopy() *BoxedJobPodSet {
	if in == nil {
		return nil
	}
	out := new(BoxedJobPodSet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BoxedJobSpec) DeepCopyInto(out *BoxedJobSpec) {
	*out = *in
	if in.Components != nil {
		in, out := &in.Components, &out.Components
		*out = make([]BoxedJobComponent, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BoxedJobSpec.
func (in *BoxedJobSpec) DeepCopy() *BoxedJobSpec {
	if in == nil {
		return nil
	}
	out := new(BoxedJobSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BoxedJobStatus) DeepCopyInto(out *BoxedJobStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BoxedJobStatus.
func (in *BoxedJobStatus) DeepCopy() *BoxedJobStatus {
	if in == nil {
		return nil
	}
	out := new(BoxedJobStatus)
	in.DeepCopyInto(out)
	return out
}
