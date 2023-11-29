/*
Copyright 2023 IONOS Cloud.

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
	clusterv1 "sigs.k8s.io/cluster-api/api/v1beta1"
)

// ProxmoxClusterTemplateSpec defines the desired state of ProxmoxClusterTemplate
type ProxmoxClusterTemplateSpec struct {
	Template ProxmoxClusterTemplateResource `json:"template"`
}

// ProxmoxClusterTemplateResource defines the spec and metadata for ProxmoxClusterTemplate supported by capi.
type ProxmoxClusterTemplateResource struct {
	// Standard object's metadata.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	// +optional
	ObjectMeta clusterv1.ObjectMeta `json:"metadata,omitempty"`
	Spec       ProxmoxClusterSpec   `json:"spec"`
}

// ProxmoxClusterTemplateStatus defines the observed state of ProxmoxClusterTemplate
type ProxmoxClusterTemplateStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// ProxmoxClusterTemplate is the Schema for the proxmoxclustertemplates API
type ProxmoxClusterTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ProxmoxClusterTemplateSpec   `json:"spec,omitempty"`
	Status ProxmoxClusterTemplateStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ProxmoxClusterTemplateList contains a list of ProxmoxClusterTemplate
type ProxmoxClusterTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProxmoxClusterTemplate `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ProxmoxClusterTemplate{}, &ProxmoxClusterTemplateList{})
}