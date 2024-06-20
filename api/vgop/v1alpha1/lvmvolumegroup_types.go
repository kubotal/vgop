/*
Copyright © 2023 Red Hat, Inc.
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

package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// LVMVolumeGroupSpec defines the desired state of LVMVolumeGroup
type LVMVolumeGroupSpec struct {
	// DeviceSelector is a set of rules that should match for a device to be included in this VG
	// +optional
	DeviceSelector *DeviceSelector `json:"deviceSelector,omitempty"`

	// NodeSelector chooses nodes
	// +optional
	NodeSelector *corev1.NodeSelector `json:"nodeSelector,omitempty"`

	// ThinPoolConfig contains configurations for the thin-pool
	// +optional
	ThinPoolConfig *ThinPoolConfig `json:"thinPoolConfig,omitempty"`

	// Default is a flag to indicate whether the device-class is the default
	// +optional
	Default bool `json:"default,omitempty"`
}

// LVMVolumeGroupStatus defines the observed state of LVMVolumeGroup
type LVMVolumeGroupStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// LVMVolumeGroup is the Schema for the lvmvolumegroups API
type LVMVolumeGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   LVMVolumeGroupSpec   `json:"spec,omitempty"`
	Status LVMVolumeGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LVMVolumeGroupList contains a list of LVMVolumeGroup
type LVMVolumeGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LVMVolumeGroup `json:"items"`
}

func init() {
	SchemeBuilder.Register(&LVMVolumeGroup{}, &LVMVolumeGroupList{})
}

// DeviceSelector specifies the list of criteria that have to match before a device is assigned
type DeviceSelector struct {
	// MinSize is the minimum size of the device which needs to be included. Defaults to `1Gi` if empty
	// +optional
	// MinSize *resource.Quantity `json:"minSize,omitempty"`

	// Paths specify the device paths.
	// +optional
	Paths []string `json:"paths,omitempty"`

	// OptionalPaths specify the optional device paths.
	// +optional
	OptionalPaths []string `json:"optionalPaths,omitempty"`

	// ForceWipeDevicesAndDestroyAllData is a flag to force wipe the selected devices.
	// This wipes the file signatures on the devices. Use this feature with caution.
	// Force wipe the devices only when you know that they do not contain any important data.
	// +optional
	ForceWipeDevicesAndDestroyAllData *bool `json:"forceWipeDevicesAndDestroyAllData,omitempty"`
}

type ThinPoolConfig struct {
	// Name specifies a name for the thin pool.
	// +kubebuilder:validation:Required
	// +required
	Name string `json:"name"`

	// SizePercent specifies the percentage of space in the LVM volume group for creating the thin pool.
	// +kubebuilder:default=90
	// +kubebuilder:validation:Minimum=10
	// +kubebuilder:validation:Maximum=90
	SizePercent int `json:"sizePercent,omitempty"`

	// OverProvisionRatio specifies a factor by which you can provision additional storage based on the available storage in the thin pool. To prevent over-provisioning through validation, set this field to 1.
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Required
	// +required
	OverprovisionRatio int `json:"overprovisionRatio"`

	// ChunkSizeCalculationPolicy specifies the policy to calculate the chunk size for the underlying volume.
	// When set to Host, the chunk size is calculated based on the lvm2 host setting on the node.
	// When set to Static, the chunk size is calculated based on the static size attribute provided within ChunkSize.
	// +kubebuilder:default=Static
	// +kubebuilder:validation:Enum=Host;Static
	// +required
	ChunkSizeCalculationPolicy ChunkSizeCalculationPolicy `json:"chunkSizeCalculationPolicy,omitempty"`

	// ChunkSize specifies the statically calculated chunk size for the thin pool.
	// Thus, It is only used when the ChunkSizeCalculationPolicy is set to Static.
	// No ChunkSize with a ChunkSizeCalculationPolicy set to Static will result in a default chunk size of 128Ki.
	// It can be between 64Ki and 1Gi due to the underlying limitations of lvm2.
	// +optional
	ChunkSize *resource.Quantity `json:"chunkSize,omitempty"`
}

// ChunkSizeCalculationPolicy specifies the policy to calculate the chunk size for the underlying volume.
// for more information, see man lvm.
type ChunkSizeCalculationPolicy string

const (
// ChunkSizeCalculationPolicyHost calculates the chunk size based on the lvm2 host setting on the node.
// ChunkSizeCalculationPolicyHost ChunkSizeCalculationPolicy = "Host"
// ChunkSizeCalculationPolicyStatic calculates the chunk size based on a static size attribute.
// ChunkSizeCalculationPolicyStatic ChunkSizeCalculationPolicy = "Static"
)

var (
	ChunkSizeDefault = resource.MustParse("128Ki")
	// ChunkSizeMaximum = resource.MustParse("1Gi")
	// ChunkSizeMinimum = resource.MustParse("64Ki")
)
