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

type PgDatabasePrivilegeInitParameters struct {

	// Specifies the database name. Changing this creates a new resource.
	// Specifies the database name.
	DBName *string `json:"dbName,omitempty" tf:"db_name,omitempty"`

	// Specifies the RDS instance ID. Changing this will create a new resource.
	// Specifies the ID of the RDS PostgreSQL instance.
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// The region in which to create the RDS database privilege resource. If omitted,
	// the provider-level region will be used. Changing this creates a new resource.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Specifies the account that associated with the database.
	// The users structure is documented below.
	// Specifies the account that associated with the database
	Users []PgDatabasePrivilegeUsersInitParameters `json:"users,omitempty" tf:"users,omitempty"`
}

type PgDatabasePrivilegeObservation struct {

	// Specifies the database name. Changing this creates a new resource.
	// Specifies the database name.
	DBName *string `json:"dbName,omitempty" tf:"db_name,omitempty"`

	// The resource ID of database privilege which is formatted <instance_id>/<db_name>.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the RDS instance ID. Changing this will create a new resource.
	// Specifies the ID of the RDS PostgreSQL instance.
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// The region in which to create the RDS database privilege resource. If omitted,
	// the provider-level region will be used. Changing this creates a new resource.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Specifies the account that associated with the database.
	// The users structure is documented below.
	// Specifies the account that associated with the database
	Users []PgDatabasePrivilegeUsersObservation `json:"users,omitempty" tf:"users,omitempty"`
}

type PgDatabasePrivilegeParameters struct {

	// Specifies the database name. Changing this creates a new resource.
	// Specifies the database name.
	// +kubebuilder:validation:Optional
	DBName *string `json:"dbName,omitempty" tf:"db_name,omitempty"`

	// Specifies the RDS instance ID. Changing this will create a new resource.
	// Specifies the ID of the RDS PostgreSQL instance.
	// +kubebuilder:validation:Optional
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// The region in which to create the RDS database privilege resource. If omitted,
	// the provider-level region will be used. Changing this creates a new resource.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Specifies the account that associated with the database.
	// The users structure is documented below.
	// Specifies the account that associated with the database
	// +kubebuilder:validation:Optional
	Users []PgDatabasePrivilegeUsersParameters `json:"users,omitempty" tf:"users,omitempty"`
}

type PgDatabasePrivilegeUsersInitParameters struct {

	// Specifies the username of the database account.
	// Specifies the username of the database account.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies the read-only permission. The value can be:
	// Specifies the read-only permission.
	Readonly *bool `json:"readonly,omitempty" tf:"readonly,omitempty"`

	// Specifies the name of the schema.
	// Specifies the name of the schema.
	SchemaName *string `json:"schemaName,omitempty" tf:"schema_name,omitempty"`
}

type PgDatabasePrivilegeUsersObservation struct {

	// Specifies the username of the database account.
	// Specifies the username of the database account.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies the read-only permission. The value can be:
	// Specifies the read-only permission.
	Readonly *bool `json:"readonly,omitempty" tf:"readonly,omitempty"`

	// Specifies the name of the schema.
	// Specifies the name of the schema.
	SchemaName *string `json:"schemaName,omitempty" tf:"schema_name,omitempty"`
}

type PgDatabasePrivilegeUsersParameters struct {

	// Specifies the username of the database account.
	// Specifies the username of the database account.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// Specifies the read-only permission. The value can be:
	// Specifies the read-only permission.
	// +kubebuilder:validation:Optional
	Readonly *bool `json:"readonly" tf:"readonly,omitempty"`

	// Specifies the name of the schema.
	// Specifies the name of the schema.
	// +kubebuilder:validation:Optional
	SchemaName *string `json:"schemaName" tf:"schema_name,omitempty"`
}

// PgDatabasePrivilegeSpec defines the desired state of PgDatabasePrivilege
type PgDatabasePrivilegeSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PgDatabasePrivilegeParameters `json:"forProvider"`
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
	InitProvider PgDatabasePrivilegeInitParameters `json:"initProvider,omitempty"`
}

// PgDatabasePrivilegeStatus defines the observed state of PgDatabasePrivilege.
type PgDatabasePrivilegeStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PgDatabasePrivilegeObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// PgDatabasePrivilege is the Schema for the PgDatabasePrivileges API. Manages an RDS PostgreSQL database privilege resource within HuaweiCloud.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,huaweicloud}
type PgDatabasePrivilege struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.dbName) || (has(self.initProvider) && has(self.initProvider.dbName))",message="spec.forProvider.dbName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.instanceId) || (has(self.initProvider) && has(self.initProvider.instanceId))",message="spec.forProvider.instanceId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.users) || (has(self.initProvider) && has(self.initProvider.users))",message="spec.forProvider.users is a required parameter"
	Spec   PgDatabasePrivilegeSpec   `json:"spec"`
	Status PgDatabasePrivilegeStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PgDatabasePrivilegeList contains a list of PgDatabasePrivileges
type PgDatabasePrivilegeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PgDatabasePrivilege `json:"items"`
}

// Repository type metadata.
var (
	PgDatabasePrivilege_Kind             = "PgDatabasePrivilege"
	PgDatabasePrivilege_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PgDatabasePrivilege_Kind}.String()
	PgDatabasePrivilege_KindAPIVersion   = PgDatabasePrivilege_Kind + "." + CRDGroupVersion.String()
	PgDatabasePrivilege_GroupVersionKind = CRDGroupVersion.WithKind(PgDatabasePrivilege_Kind)
)

func init() {
	SchemeBuilder.Register(&PgDatabasePrivilege{}, &PgDatabasePrivilegeList{})
}
