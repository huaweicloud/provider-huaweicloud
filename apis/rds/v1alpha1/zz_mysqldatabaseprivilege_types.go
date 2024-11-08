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

type MySQLDatabasePrivilegeInitParameters struct {

	// Specifies the database name. Changing this creates a new resource.
	// Specifies the database name.
	DBName *string `json:"dbName,omitempty" tf:"db_name,omitempty"`

	// Specifies the RDS instance ID. Changing this will create a new resource.
	// Specifies the ID of the RDS Mysql instance.
	// +crossplane:generate:reference:type=github.com/huaweicloud/provider-huaweicloud/apis/rds/v1alpha1.Instance
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// Reference to a Instance in rds to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDRef *v1.Reference `json:"instanceIdRef,omitempty" tf:"-"`

	// Selector for a Instance in rds to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDSelector *v1.Selector `json:"instanceIdSelector,omitempty" tf:"-"`

	// The region in which to create the RDS database privilege resource. If omitted,
	// the provider-level region will be used. Changing this creates a new resource.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Specifies the account that associated with the database. Structure is documented below.
	// Specifies the account that associated with the database.
	Users []UsersInitParameters `json:"users,omitempty" tf:"users,omitempty"`
}

type MySQLDatabasePrivilegeObservation struct {

	// Specifies the database name. Changing this creates a new resource.
	// Specifies the database name.
	DBName *string `json:"dbName,omitempty" tf:"db_name,omitempty"`

	// The resource ID of database privilege which is formatted <instance_id>/<db_name>.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the RDS instance ID. Changing this will create a new resource.
	// Specifies the ID of the RDS Mysql instance.
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// The region in which to create the RDS database privilege resource. If omitted,
	// the provider-level region will be used. Changing this creates a new resource.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Specifies the account that associated with the database. Structure is documented below.
	// Specifies the account that associated with the database.
	Users []UsersObservation `json:"users,omitempty" tf:"users,omitempty"`
}

type MySQLDatabasePrivilegeParameters struct {

	// Specifies the database name. Changing this creates a new resource.
	// Specifies the database name.
	// +kubebuilder:validation:Optional
	DBName *string `json:"dbName,omitempty" tf:"db_name,omitempty"`

	// Specifies the RDS instance ID. Changing this will create a new resource.
	// Specifies the ID of the RDS Mysql instance.
	// +crossplane:generate:reference:type=github.com/huaweicloud/provider-huaweicloud/apis/rds/v1alpha1.Instance
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// Reference to a Instance in rds to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDRef *v1.Reference `json:"instanceIdRef,omitempty" tf:"-"`

	// Selector for a Instance in rds to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDSelector *v1.Selector `json:"instanceIdSelector,omitempty" tf:"-"`

	// The region in which to create the RDS database privilege resource. If omitted,
	// the provider-level region will be used. Changing this creates a new resource.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Specifies the account that associated with the database. Structure is documented below.
	// Specifies the account that associated with the database.
	// +kubebuilder:validation:Optional
	Users []UsersParameters `json:"users,omitempty" tf:"users,omitempty"`
}

type UsersInitParameters struct {

	// Specifies the username of the database account.
	// Specifies the username of the database account.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies the read-only permission. The value can be:
	// Specifies the read-only permission.
	Readonly *bool `json:"readonly,omitempty" tf:"readonly,omitempty"`
}

type UsersObservation struct {

	// Specifies the username of the database account.
	// Specifies the username of the database account.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies the read-only permission. The value can be:
	// Specifies the read-only permission.
	Readonly *bool `json:"readonly,omitempty" tf:"readonly,omitempty"`
}

type UsersParameters struct {

	// Specifies the username of the database account.
	// Specifies the username of the database account.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// Specifies the read-only permission. The value can be:
	// Specifies the read-only permission.
	// +kubebuilder:validation:Optional
	Readonly *bool `json:"readonly,omitempty" tf:"readonly,omitempty"`
}

// MySQLDatabasePrivilegeSpec defines the desired state of MySQLDatabasePrivilege
type MySQLDatabasePrivilegeSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MySQLDatabasePrivilegeParameters `json:"forProvider"`
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
	InitProvider MySQLDatabasePrivilegeInitParameters `json:"initProvider,omitempty"`
}

// MySQLDatabasePrivilegeStatus defines the observed state of MySQLDatabasePrivilege.
type MySQLDatabasePrivilegeStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MySQLDatabasePrivilegeObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// MySQLDatabasePrivilege is the Schema for the MySQLDatabasePrivileges API. ""
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,huaweicloud}
type MySQLDatabasePrivilege struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.dbName) || (has(self.initProvider) && has(self.initProvider.dbName))",message="spec.forProvider.dbName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.users) || (has(self.initProvider) && has(self.initProvider.users))",message="spec.forProvider.users is a required parameter"
	Spec   MySQLDatabasePrivilegeSpec   `json:"spec"`
	Status MySQLDatabasePrivilegeStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MySQLDatabasePrivilegeList contains a list of MySQLDatabasePrivileges
type MySQLDatabasePrivilegeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MySQLDatabasePrivilege `json:"items"`
}

// Repository type metadata.
var (
	MySQLDatabasePrivilege_Kind             = "MySQLDatabasePrivilege"
	MySQLDatabasePrivilege_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: MySQLDatabasePrivilege_Kind}.String()
	MySQLDatabasePrivilege_KindAPIVersion   = MySQLDatabasePrivilege_Kind + "." + CRDGroupVersion.String()
	MySQLDatabasePrivilege_GroupVersionKind = CRDGroupVersion.WithKind(MySQLDatabasePrivilege_Kind)
)

func init() {
	SchemeBuilder.Register(&MySQLDatabasePrivilege{}, &MySQLDatabasePrivilegeList{})
}
