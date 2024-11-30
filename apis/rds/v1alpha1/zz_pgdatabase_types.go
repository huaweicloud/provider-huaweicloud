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

type PgDatabaseInitParameters struct {

	// Specifies the database character set.
	// For details, see documentation.
	// Defaults to UTF8.
	// Specifies the database character set.
	CharacterSet *string `json:"characterSet,omitempty" tf:"character_set,omitempty"`

	// Specifies the database description. The value contains 0 to 512 characters.
	// Specifies the database description.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies the ID of the RDS PostgreSQL instance.
	// Specifies the ID of the RDS PostgreSQL instance.
	// +crossplane:generate:reference:type=github.com/huaweicloud/provider-huaweicloud/apis/rds/v1alpha1.Instance
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// Reference to a Instance in rds to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDRef *v1.Reference `json:"instanceIdRef,omitempty" tf:"-"`

	// Selector for a Instance in rds to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDSelector *v1.Selector `json:"instanceIdSelector,omitempty" tf:"-"`

	// Specifies whether to revoke the PUBLIC CREATE permission of
	// the public schema.
	// Specifies whether to revoke the PUBLIC CREATE permission of the public schema.
	IsRevokePublicPrivilege *bool `json:"isRevokePublicPrivilege,omitempty" tf:"is_revoke_public_privilege,omitempty"`

	// Specifies the database collocation.
	// For details, see documentation.
	// Defaults to en_US.UTF-8.
	// Specifies the database collocation.
	LcCollate *string `json:"lcCollate,omitempty" tf:"lc_collate,omitempty"`

	// Specifies the database classification.
	// For details, see documentation.
	// Defaults to: en_US.UTF-8.
	// Specifies the database classification.
	LcCtype *string `json:"lcCtype,omitempty" tf:"lc_ctype,omitempty"`

	// Specifies the database name. The value contains 1 to 63 characters, including
	// letters, digits, and underscores (_). It cannot start with pg or a digit, and must be different from RDS for
	// PostgreSQL template library names. RDS for PostgreSQL template libraries include postgres, template0, and
	// template1.
	// Specifies the database name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies the database user. The value must be an existing username and must be different
	// from system usernames. Defaults to root.
	// Specifies the database user.
	Owner *string `json:"owner,omitempty" tf:"owner,omitempty"`

	// Specifies the region in which to create the resource.
	// If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Specifies the name of the database template. Value options: template0,
	// template1. Defaults to template1.
	// Specifies the name of the database template.
	Template *string `json:"template,omitempty" tf:"template,omitempty"`
}

type PgDatabaseObservation struct {

	// Specifies the database character set.
	// For details, see documentation.
	// Defaults to UTF8.
	// Specifies the database character set.
	CharacterSet *string `json:"characterSet,omitempty" tf:"character_set,omitempty"`

	// Specifies the database description. The value contains 0 to 512 characters.
	// Specifies the database description.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The resource ID of database which is formatted <instance_id>/<name>.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the ID of the RDS PostgreSQL instance.
	// Specifies the ID of the RDS PostgreSQL instance.
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// Specifies whether to revoke the PUBLIC CREATE permission of
	// the public schema.
	// Specifies whether to revoke the PUBLIC CREATE permission of the public schema.
	IsRevokePublicPrivilege *bool `json:"isRevokePublicPrivilege,omitempty" tf:"is_revoke_public_privilege,omitempty"`

	// Specifies the database collocation.
	// For details, see documentation.
	// Defaults to en_US.UTF-8.
	// Specifies the database collocation.
	LcCollate *string `json:"lcCollate,omitempty" tf:"lc_collate,omitempty"`

	// Specifies the database classification.
	// For details, see documentation.
	// Defaults to: en_US.UTF-8.
	// Specifies the database classification.
	LcCtype *string `json:"lcCtype,omitempty" tf:"lc_ctype,omitempty"`

	// Specifies the database name. The value contains 1 to 63 characters, including
	// letters, digits, and underscores (_). It cannot start with pg or a digit, and must be different from RDS for
	// PostgreSQL template library names. RDS for PostgreSQL template libraries include postgres, template0, and
	// template1.
	// Specifies the database name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies the database user. The value must be an existing username and must be different
	// from system usernames. Defaults to root.
	// Specifies the database user.
	Owner *string `json:"owner,omitempty" tf:"owner,omitempty"`

	// Specifies the region in which to create the resource.
	// If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Indicates the database size, in bytes.
	// Indicates the database size, in bytes.
	Size *float64 `json:"size,omitempty" tf:"size,omitempty"`

	// Specifies the name of the database template. Value options: template0,
	// template1. Defaults to template1.
	// Specifies the name of the database template.
	Template *string `json:"template,omitempty" tf:"template,omitempty"`
}

type PgDatabaseParameters struct {

	// Specifies the database character set.
	// For details, see documentation.
	// Defaults to UTF8.
	// Specifies the database character set.
	// +kubebuilder:validation:Optional
	CharacterSet *string `json:"characterSet,omitempty" tf:"character_set,omitempty"`

	// Specifies the database description. The value contains 0 to 512 characters.
	// Specifies the database description.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies the ID of the RDS PostgreSQL instance.
	// Specifies the ID of the RDS PostgreSQL instance.
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

	// Specifies whether to revoke the PUBLIC CREATE permission of
	// the public schema.
	// Specifies whether to revoke the PUBLIC CREATE permission of the public schema.
	// +kubebuilder:validation:Optional
	IsRevokePublicPrivilege *bool `json:"isRevokePublicPrivilege,omitempty" tf:"is_revoke_public_privilege,omitempty"`

	// Specifies the database collocation.
	// For details, see documentation.
	// Defaults to en_US.UTF-8.
	// Specifies the database collocation.
	// +kubebuilder:validation:Optional
	LcCollate *string `json:"lcCollate,omitempty" tf:"lc_collate,omitempty"`

	// Specifies the database classification.
	// For details, see documentation.
	// Defaults to: en_US.UTF-8.
	// Specifies the database classification.
	// +kubebuilder:validation:Optional
	LcCtype *string `json:"lcCtype,omitempty" tf:"lc_ctype,omitempty"`

	// Specifies the database name. The value contains 1 to 63 characters, including
	// letters, digits, and underscores (_). It cannot start with pg or a digit, and must be different from RDS for
	// PostgreSQL template library names. RDS for PostgreSQL template libraries include postgres, template0, and
	// template1.
	// Specifies the database name.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies the database user. The value must be an existing username and must be different
	// from system usernames. Defaults to root.
	// Specifies the database user.
	// +kubebuilder:validation:Optional
	Owner *string `json:"owner,omitempty" tf:"owner,omitempty"`

	// Specifies the region in which to create the resource.
	// If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Specifies the name of the database template. Value options: template0,
	// template1. Defaults to template1.
	// Specifies the name of the database template.
	// +kubebuilder:validation:Optional
	Template *string `json:"template,omitempty" tf:"template,omitempty"`
}

// PgDatabaseSpec defines the desired state of PgDatabase
type PgDatabaseSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PgDatabaseParameters `json:"forProvider"`
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
	InitProvider PgDatabaseInitParameters `json:"initProvider,omitempty"`
}

// PgDatabaseStatus defines the observed state of PgDatabase.
type PgDatabaseStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PgDatabaseObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// PgDatabase is the Schema for the PgDatabases API. ""
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,huaweicloud}
type PgDatabase struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   PgDatabaseSpec   `json:"spec"`
	Status PgDatabaseStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PgDatabaseList contains a list of PgDatabases
type PgDatabaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PgDatabase `json:"items"`
}

// Repository type metadata.
var (
	PgDatabase_Kind             = "PgDatabase"
	PgDatabase_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PgDatabase_Kind}.String()
	PgDatabase_KindAPIVersion   = PgDatabase_Kind + "." + CRDGroupVersion.String()
	PgDatabase_GroupVersionKind = CRDGroupVersion.WithKind(PgDatabase_Kind)
)

func init() {
	SchemeBuilder.Register(&PgDatabase{}, &PgDatabaseList{})
}