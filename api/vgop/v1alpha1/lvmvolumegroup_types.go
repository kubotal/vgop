/*
Copyright 2024.

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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// LVMVolumeGroupSpec defines the desired state of LVMVolumeGroup
type LVMVolumeGroupSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of LVMVolumeGroup. Edit lvmvolumegroup_types.go to remove/update
	Foo string `json:"foo,omitempty"`
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
