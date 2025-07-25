//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2025 The CloudPilot AI Authors.

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	status "github.com/awslabs/operatorpkg/status"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Alias) DeepCopyInto(out *Alias) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Alias.
func (in *Alias) DeepCopy() *Alias {
	if in == nil {
		return nil
	}
	out := new(Alias)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Disk) DeepCopyInto(out *Disk) {
	*out = *in
	if in.Categories != nil {
		in, out := &in.Categories, &out.Categories
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Disk.
func (in *Disk) DeepCopy() *Disk {
	if in == nil {
		return nil
	}
	out := new(Disk)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GCENodeClass) DeepCopyInto(out *GCENodeClass) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GCENodeClass.
func (in *GCENodeClass) DeepCopy() *GCENodeClass {
	if in == nil {
		return nil
	}
	out := new(GCENodeClass)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GCENodeClass) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GCENodeClassList) DeepCopyInto(out *GCENodeClassList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]GCENodeClass, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GCENodeClassList.
func (in *GCENodeClassList) DeepCopy() *GCENodeClassList {
	if in == nil {
		return nil
	}
	out := new(GCENodeClassList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GCENodeClassList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GCENodeClassSpec) DeepCopyInto(out *GCENodeClassSpec) {
	*out = *in
	if in.Disks != nil {
		in, out := &in.Disks, &out.Disks
		*out = new(Disk)
		(*in).DeepCopyInto(*out)
	}
	if in.ImageSelectorTerms != nil {
		in, out := &in.ImageSelectorTerms, &out.ImageSelectorTerms
		*out = make([]ImageSelectorTerm, len(*in))
		copy(*out, *in)
	}
	if in.ImageFamily != nil {
		in, out := &in.ImageFamily, &out.ImageFamily
		*out = new(string)
		**out = **in
	}
	if in.KubeletConfiguration != nil {
		in, out := &in.KubeletConfiguration, &out.KubeletConfiguration
		*out = new(KubeletConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Metadata != nil {
		in, out := &in.Metadata, &out.Metadata
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GCENodeClassSpec.
func (in *GCENodeClassSpec) DeepCopy() *GCENodeClassSpec {
	if in == nil {
		return nil
	}
	out := new(GCENodeClassSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GCENodeClassStatus) DeepCopyInto(out *GCENodeClassStatus) {
	*out = *in
	if in.Images != nil {
		in, out := &in.Images, &out.Images
		*out = make([]Image, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]status.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GCENodeClassStatus.
func (in *GCENodeClassStatus) DeepCopy() *GCENodeClassStatus {
	if in == nil {
		return nil
	}
	out := new(GCENodeClassStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Image) DeepCopyInto(out *Image) {
	*out = *in
	if in.Requirements != nil {
		in, out := &in.Requirements, &out.Requirements
		*out = make([]v1.NodeSelectorRequirement, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Image.
func (in *Image) DeepCopy() *Image {
	if in == nil {
		return nil
	}
	out := new(Image)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageSelectorTerm) DeepCopyInto(out *ImageSelectorTerm) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageSelectorTerm.
func (in *ImageSelectorTerm) DeepCopy() *ImageSelectorTerm {
	if in == nil {
		return nil
	}
	out := new(ImageSelectorTerm)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeletConfiguration) DeepCopyInto(out *KubeletConfiguration) {
	*out = *in
	if in.ClusterDNS != nil {
		in, out := &in.ClusterDNS, &out.ClusterDNS
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.MaxPods != nil {
		in, out := &in.MaxPods, &out.MaxPods
		*out = new(int32)
		**out = **in
	}
	if in.PodsPerCore != nil {
		in, out := &in.PodsPerCore, &out.PodsPerCore
		*out = new(int32)
		**out = **in
	}
	if in.SystemReserved != nil {
		in, out := &in.SystemReserved, &out.SystemReserved
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.KubeReserved != nil {
		in, out := &in.KubeReserved, &out.KubeReserved
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.EvictionHard != nil {
		in, out := &in.EvictionHard, &out.EvictionHard
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.EvictionSoft != nil {
		in, out := &in.EvictionSoft, &out.EvictionSoft
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.EvictionSoftGracePeriod != nil {
		in, out := &in.EvictionSoftGracePeriod, &out.EvictionSoftGracePeriod
		*out = make(map[string]metav1.Duration, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.EvictionMaxPodGracePeriod != nil {
		in, out := &in.EvictionMaxPodGracePeriod, &out.EvictionMaxPodGracePeriod
		*out = new(int32)
		**out = **in
	}
	if in.ImageGCHighThresholdPercent != nil {
		in, out := &in.ImageGCHighThresholdPercent, &out.ImageGCHighThresholdPercent
		*out = new(int32)
		**out = **in
	}
	if in.ImageGCLowThresholdPercent != nil {
		in, out := &in.ImageGCLowThresholdPercent, &out.ImageGCLowThresholdPercent
		*out = new(int32)
		**out = **in
	}
	if in.CPUCFSQuota != nil {
		in, out := &in.CPUCFSQuota, &out.CPUCFSQuota
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeletConfiguration.
func (in *KubeletConfiguration) DeepCopy() *KubeletConfiguration {
	if in == nil {
		return nil
	}
	out := new(KubeletConfiguration)
	in.DeepCopyInto(out)
	return out
}
