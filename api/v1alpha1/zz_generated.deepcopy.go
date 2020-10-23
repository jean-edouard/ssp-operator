// +build !ignore_autogenerated

/*


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
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SSP) DeepCopyInto(out *SSP) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SSP.
func (in *SSP) DeepCopy() *SSP {
	if in == nil {
		return nil
	}
	out := new(SSP)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SSP) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SSPList) DeepCopyInto(out *SSPList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SSP, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SSPList.
func (in *SSPList) DeepCopy() *SSPList {
	if in == nil {
		return nil
	}
	out := new(SSPList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SSPList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SSPSpec) DeepCopyInto(out *SSPSpec) {
	*out = *in
	in.TemplateValidator.DeepCopyInto(&out.TemplateValidator)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SSPSpec.
func (in *SSPSpec) DeepCopy() *SSPSpec {
	if in == nil {
		return nil
	}
	out := new(SSPSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SSPStatus) DeepCopyInto(out *SSPStatus) {
	*out = *in
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SSPStatus.
func (in *SSPStatus) DeepCopy() *SSPStatus {
	if in == nil {
		return nil
	}
	out := new(SSPStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TemplateValidator) DeepCopyInto(out *TemplateValidator) {
	*out = *in
	in.Placement.DeepCopyInto(&out.Placement)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TemplateValidator.
func (in *TemplateValidator) DeepCopy() *TemplateValidator {
	if in == nil {
		return nil
	}
	out := new(TemplateValidator)
	in.DeepCopyInto(out)
	return out
}
