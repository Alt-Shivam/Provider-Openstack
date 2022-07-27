/*
Copyright 2022 The Crossplane Authors.

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
	"reflect"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

// FlavourParameters are the configurable fields of a Flavour.
type FlavourParameters struct {
	ConfigurableField string `json:"configurableField"`
}

// FlavourObservation are the observable fields of a Flavour.
type FlavourObservation struct {
	ObservableField string `json:"observableField,omitempty"`
}

// A FlavourSpec defines the desired state of a Flavour.
type FlavourSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       FlavourParameters `json:"forProvider"`
}

// A FlavourStatus represents the observed state of a Flavour.
type FlavourStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          FlavourObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// A Flavour is an example API type.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,openstack}
type Flavour struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   FlavourSpec   `json:"spec"`
	Status FlavourStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FlavourList contains a list of Flavour
type FlavourList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Flavour `json:"items"`
}

// Flavour type metadata.
var (
	FlavourKind             = reflect.TypeOf(Flavour{}).Name()
	FlavourGroupKind        = schema.GroupKind{Group: Group, Kind: FlavourKind}.String()
	FlavourKindAPIVersion   = FlavourKind + "." + SchemeGroupVersion.String()
	FlavourGroupVersionKind = SchemeGroupVersion.WithKind(FlavourKind)
)

func init() {
	SchemeBuilder.Register(&Flavour{}, &FlavourList{})
}
