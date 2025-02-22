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

type BackupInitParameters struct {

	// List of self-built Microsoft SQL Server databases that are partially
	// backed up.
	// (Only Microsoft SQL Server supports partial backups.).
	// List of self-built Microsoft SQL Server databases that are partially backed up.
	Databases []DatabasesInitParameters `json:"databases,omitempty" tf:"databases,omitempty"`

	// The description about the backup.
	// It contains a maximum of 256 characters and cannot contain the following special characters: >!<"&'=.
	// The description about the backup.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Instance ID.
	// Instance ID.
	// +crossplane:generate:reference:type=github.com/huaweicloud/provider-huaweicloud/apis/rds/v1alpha1.Instance
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// Reference to a Instance in rds to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDRef *v1.Reference `json:"instanceIdRef,omitempty" tf:"-"`

	// Selector for a Instance in rds to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDSelector *v1.Selector `json:"instanceIdSelector,omitempty" tf:"-"`

	// Backup name.
	// It must be 4 to 64 characters long, start with a letter, and contain only letters (case-sensitive),
	// digits, hyphens (-), and underscores (_).
	// Backup name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies the region in which to create the resource.
	// If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

type BackupObservation struct {

	// Whether a DDM instance has been associated.
	// Whether a DDM instance has been associated.
	AssociatedWithDdm *bool `json:"associatedWithDdm,omitempty" tf:"associated_with_ddm,omitempty"`

	// Backup start time in the "yyyy-mm-ddThh:mm:ssZ" format.
	// Backup start time in the "yyyy-mm-ddThh:mm:ssZ" format.
	BeginTime *string `json:"beginTime,omitempty" tf:"begin_time,omitempty"`

	// List of self-built Microsoft SQL Server databases that are partially
	// backed up.
	// (Only Microsoft SQL Server supports partial backups.).
	// List of self-built Microsoft SQL Server databases that are partially backed up.
	Databases []DatabasesObservation `json:"databases,omitempty" tf:"databases,omitempty"`

	// The description about the backup.
	// It contains a maximum of 256 characters and cannot contain the following special characters: >!<"&'=.
	// The description about the backup.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Backup end time in the "yyyy-mm-ddThh:mm:ssZ" format.
	// Backup end time in the "yyyy-mm-ddThh:mm:ssZ" format.
	EndTime *string `json:"endTime,omitempty" tf:"end_time,omitempty"`

	// The resource ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Instance ID.
	// Instance ID.
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// Backup name.
	// It must be 4 to 64 characters long, start with a letter, and contain only letters (case-sensitive),
	// digits, hyphens (-), and underscores (_).
	// Backup name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies the region in which to create the resource.
	// If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Backup size in KB.
	// Backup size in KB.
	Size *float64 `json:"size,omitempty" tf:"size,omitempty"`

	// Backup status.
	// The options are as follows:
	// Backup status.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type BackupParameters struct {

	// List of self-built Microsoft SQL Server databases that are partially
	// backed up.
	// (Only Microsoft SQL Server supports partial backups.).
	// List of self-built Microsoft SQL Server databases that are partially backed up.
	// +kubebuilder:validation:Optional
	Databases []DatabasesParameters `json:"databases,omitempty" tf:"databases,omitempty"`

	// The description about the backup.
	// It contains a maximum of 256 characters and cannot contain the following special characters: >!<"&'=.
	// The description about the backup.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Instance ID.
	// Instance ID.
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

	// Backup name.
	// It must be 4 to 64 characters long, start with a letter, and contain only letters (case-sensitive),
	// digits, hyphens (-), and underscores (_).
	// Backup name.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies the region in which to create the resource.
	// If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

type DatabasesInitParameters struct {

	// Backup name.
	// It must be 4 to 64 characters long, start with a letter, and contain only letters (case-sensitive),
	// digits, hyphens (-), and underscores (_).
	// Database to be backed up for Microsoft SQL Server.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type DatabasesObservation struct {

	// Backup name.
	// It must be 4 to 64 characters long, start with a letter, and contain only letters (case-sensitive),
	// digits, hyphens (-), and underscores (_).
	// Database to be backed up for Microsoft SQL Server.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type DatabasesParameters struct {

	// Backup name.
	// It must be 4 to 64 characters long, start with a letter, and contain only letters (case-sensitive),
	// digits, hyphens (-), and underscores (_).
	// Database to be backed up for Microsoft SQL Server.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`
}

// BackupSpec defines the desired state of Backup
type BackupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BackupParameters `json:"forProvider"`
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
	InitProvider BackupInitParameters `json:"initProvider,omitempty"`
}

// BackupStatus defines the observed state of Backup.
type BackupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BackupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Backup is the Schema for the Backups API. ""
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,huaweicloud}
type Backup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   BackupSpec   `json:"spec"`
	Status BackupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BackupList contains a list of Backups
type BackupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Backup `json:"items"`
}

// Repository type metadata.
var (
	Backup_Kind             = "Backup"
	Backup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Backup_Kind}.String()
	Backup_KindAPIVersion   = Backup_Kind + "." + CRDGroupVersion.String()
	Backup_GroupVersionKind = CRDGroupVersion.WithKind(Backup_Kind)
)

func init() {
	SchemeBuilder.Register(&Backup{}, &BackupList{})
}
