//go:build !ignore_autogenerated

/*
Copyright 2024 Kubotal SAS.

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
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeviceSelector) DeepCopyInto(out *DeviceSelector) {
	*out = *in
	if in.Paths != nil {
		in, out := &in.Paths, &out.Paths
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.OptionalPaths != nil {
		in, out := &in.OptionalPaths, &out.OptionalPaths
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ForceWipeDevicesAndDestroyAllData != nil {
		in, out := &in.ForceWipeDevicesAndDestroyAllData, &out.ForceWipeDevicesAndDestroyAllData
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeviceSelector.
func (in *DeviceSelector) DeepCopy() *DeviceSelector {
	if in == nil {
		return nil
	}
	out := new(DeviceSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExcludedDevice) DeepCopyInto(out *ExcludedDevice) {
	*out = *in
	if in.Reasons != nil {
		in, out := &in.Reasons, &out.Reasons
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExcludedDevice.
func (in *ExcludedDevice) DeepCopy() *ExcludedDevice {
	if in == nil {
		return nil
	}
	out := new(ExcludedDevice)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LVMVolumeGroup) DeepCopyInto(out *LVMVolumeGroup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LVMVolumeGroup.
func (in *LVMVolumeGroup) DeepCopy() *LVMVolumeGroup {
	if in == nil {
		return nil
	}
	out := new(LVMVolumeGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LVMVolumeGroup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LVMVolumeGroupList) DeepCopyInto(out *LVMVolumeGroupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]LVMVolumeGroup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LVMVolumeGroupList.
func (in *LVMVolumeGroupList) DeepCopy() *LVMVolumeGroupList {
	if in == nil {
		return nil
	}
	out := new(LVMVolumeGroupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LVMVolumeGroupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LVMVolumeGroupNodeStatus) DeepCopyInto(out *LVMVolumeGroupNodeStatus) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LVMVolumeGroupNodeStatus.
func (in *LVMVolumeGroupNodeStatus) DeepCopy() *LVMVolumeGroupNodeStatus {
	if in == nil {
		return nil
	}
	out := new(LVMVolumeGroupNodeStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LVMVolumeGroupNodeStatus) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LVMVolumeGroupNodeStatusList) DeepCopyInto(out *LVMVolumeGroupNodeStatusList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]LVMVolumeGroupNodeStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LVMVolumeGroupNodeStatusList.
func (in *LVMVolumeGroupNodeStatusList) DeepCopy() *LVMVolumeGroupNodeStatusList {
	if in == nil {
		return nil
	}
	out := new(LVMVolumeGroupNodeStatusList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LVMVolumeGroupNodeStatusList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LVMVolumeGroupNodeStatusSpec) DeepCopyInto(out *LVMVolumeGroupNodeStatusSpec) {
	*out = *in
	if in.LVMVGStatus != nil {
		in, out := &in.LVMVGStatus, &out.LVMVGStatus
		*out = make([]VGStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LVMVolumeGroupNodeStatusSpec.
func (in *LVMVolumeGroupNodeStatusSpec) DeepCopy() *LVMVolumeGroupNodeStatusSpec {
	if in == nil {
		return nil
	}
	out := new(LVMVolumeGroupNodeStatusSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LVMVolumeGroupNodeStatusStatus) DeepCopyInto(out *LVMVolumeGroupNodeStatusStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LVMVolumeGroupNodeStatusStatus.
func (in *LVMVolumeGroupNodeStatusStatus) DeepCopy() *LVMVolumeGroupNodeStatusStatus {
	if in == nil {
		return nil
	}
	out := new(LVMVolumeGroupNodeStatusStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LVMVolumeGroupSpec) DeepCopyInto(out *LVMVolumeGroupSpec) {
	*out = *in
	if in.Selectors != nil {
		in, out := &in.Selectors, &out.Selectors
		*out = make([]Selector, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ThinPoolConfig != nil {
		in, out := &in.ThinPoolConfig, &out.ThinPoolConfig
		*out = new(ThinPoolConfig)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LVMVolumeGroupSpec.
func (in *LVMVolumeGroupSpec) DeepCopy() *LVMVolumeGroupSpec {
	if in == nil {
		return nil
	}
	out := new(LVMVolumeGroupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LVMVolumeGroupStatus) DeepCopyInto(out *LVMVolumeGroupStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LVMVolumeGroupStatus.
func (in *LVMVolumeGroupStatus) DeepCopy() *LVMVolumeGroupStatus {
	if in == nil {
		return nil
	}
	out := new(LVMVolumeGroupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Selector) DeepCopyInto(out *Selector) {
	*out = *in
	if in.DeviceSelector != nil {
		in, out := &in.DeviceSelector, &out.DeviceSelector
		*out = new(DeviceSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = new(v1.NodeSelector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Selector.
func (in *Selector) DeepCopy() *Selector {
	if in == nil {
		return nil
	}
	out := new(Selector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ThinPoolConfig) DeepCopyInto(out *ThinPoolConfig) {
	*out = *in
	if in.ChunkSize != nil {
		in, out := &in.ChunkSize, &out.ChunkSize
		x := (*in).DeepCopy()
		*out = &x
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ThinPoolConfig.
func (in *ThinPoolConfig) DeepCopy() *ThinPoolConfig {
	if in == nil {
		return nil
	}
	out := new(ThinPoolConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VGStatus) DeepCopyInto(out *VGStatus) {
	*out = *in
	if in.Devices != nil {
		in, out := &in.Devices, &out.Devices
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Excluded != nil {
		in, out := &in.Excluded, &out.Excluded
		*out = make([]ExcludedDevice, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VGStatus.
func (in *VGStatus) DeepCopy() *VGStatus {
	if in == nil {
		return nil
	}
	out := new(VGStatus)
	in.DeepCopyInto(out)
	return out
}
