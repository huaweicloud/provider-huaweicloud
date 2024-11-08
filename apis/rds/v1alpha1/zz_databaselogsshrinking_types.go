// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type DatabaseLogsShrinkingInitParameters struct {

	// Specifies the name of the database.
	// Specifies the name of the database.
	DBName *string `json:"dbName,omitempty" tf:"db_name,omitempty"`

	// Specifies the ID of instance.
	// Specifies the ID of instance.
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// The region in which to create the rds instance resource. If omitted, the
	// provider-level region will be used. Changing this creates a new resource.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

type DatabaseLogsShrinkingObservation struct {

	// Specifies the name of the database.
	// Specifies the name of the database.
	DBName *string `json:"dbName,omitempty" tf:"db_name,omitempty"`

	// The resource ID. The value is the instance ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the ID of instance.
	// Specifies the ID of instance.
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// The region in which to create the rds instance resource. If omitted, the
	// provider-level region will be used. Changing this creates a new resource.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

type DatabaseLogsShrinkingParameters struct {

	// Specifies the name of the database.
	// Specifies the name of the database.
	// +kubebuilder:validation:Optional
	DBName *string `json:"dbName,omitempty" tf:"db_name,omitempty"`

	// Specifies the ID of instance.
	// Specifies the ID of instance.
	// +kubebuilder:validation:Optional
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// The region in which to create the rds instance resource. If omitted, the
	// provider-level region will be used. Changing this creates a new resource.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

// DatabaseLogsShrinkingSpec defines the desired state of DatabaseLogsShrinking
type DatabaseLogsShrinkingSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DatabaseLogsShrinkingParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider DatabaseLogsShrinkingInitParameters `json:"initProvider,omitempty"`
}

// DatabaseLogsShrinkingStatus defines the observed state of DatabaseLogsShrinking.
type DatabaseLogsShrinkingStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DatabaseLogsShrinkingObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// DatabaseLogsShrinking is the Schema for the DatabaseLogsShrinkings API. Manages an RDS database logs shrinking resource within HuaweiCloud.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,huaweicloud}
type DatabaseLogsShrinking struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.dbName) || (has(self.initProvider) && has(self.initProvider.dbName))",message="spec.forProvider.dbName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.instanceId) || (has(self.initProvider) && has(self.initProvider.instanceId))",message="spec.forProvider.instanceId is a required parameter"
	Spec   DatabaseLogsShrinkingSpec   `json:"spec"`
	Status DatabaseLogsShrinkingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DatabaseLogsShrinkingList contains a list of DatabaseLogsShrinkings
type DatabaseLogsShrinkingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DatabaseLogsShrinking `json:"items"`
}

// Repository type metadata.
var (
	DatabaseLogsShrinking_Kind             = "DatabaseLogsShrinking"
	DatabaseLogsShrinking_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DatabaseLogsShrinking_Kind}.String()
	DatabaseLogsShrinking_KindAPIVersion   = DatabaseLogsShrinking_Kind + "." + CRDGroupVersion.String()
	DatabaseLogsShrinking_GroupVersionKind = CRDGroupVersion.WithKind(DatabaseLogsShrinking_Kind)
)

func init() {
	SchemeBuilder.Register(&DatabaseLogsShrinking{}, &DatabaseLogsShrinkingList{})
}
