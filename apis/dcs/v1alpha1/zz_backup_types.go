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

	// Specifies the format of the DCS instance backup.
	// Value options: aof, rdb. Default to rdb.
	// Specifies the format of the DCS instance backup.
	BackupFormat *string `json:"backupFormat,omitempty" tf:"backup_format,omitempty"`

	// Specifies the description of DCS instance backup.
	// Specifies the description of DCS instance backup.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies the ID of the DCS instance.
	// Specifies the ID of the DCS instance.
	// +crossplane:generate:reference:type=github.com/huaweicloud/provider-huaweicloud/apis/dcs/v1alpha1.Instance
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// Reference to a Instance in dcs to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDRef *v1.Reference `json:"instanceIdRef,omitempty" tf:"-"`

	// Selector for a Instance in dcs to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDSelector *v1.Selector `json:"instanceIdSelector,omitempty" tf:"-"`

	// Specifies the region in which to create the resource.
	// If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

type BackupObservation struct {

	// Specifies the format of the DCS instance backup.
	// Value options: aof, rdb. Default to rdb.
	// Specifies the format of the DCS instance backup.
	BackupFormat *string `json:"backupFormat,omitempty" tf:"backup_format,omitempty"`

	// Indicates the ID of the DCS instance backup.
	// Indicates the ID of the DCS instance backup.
	BackupID *string `json:"backupId,omitempty" tf:"backup_id,omitempty"`

	// Indicates the time when the backup task is created. The format is yyyy-mm-dd hh:mm:ss.
	// The value is in UTC format.
	// Indicates the time when the backup task is created.
	BeginTime *string `json:"beginTime,omitempty" tf:"begin_time,omitempty"`

	// Specifies the description of DCS instance backup.
	// Specifies the description of DCS instance backup.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Indicates the time at which DCS instance backup is completed. The format is yyyy-mm-dd hh:mm:ss.
	// The value is in UTC format.
	// Indicates the time at which DCS instance backup is completed.
	EndTime *string `json:"endTime,omitempty" tf:"end_time,omitempty"`

	// The resource ID in format of <instance_id>/<backup_id>.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the ID of the DCS instance.
	// Specifies the ID of the DCS instance.
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// Indicates whether restoration is supported. Value Options: TRUE, FALSE.
	// Indicates whether restoration is supported. Value Options: **TRUE**, **FALSE**.
	IsSupportRestore *string `json:"isSupportRestore,omitempty" tf:"is_support_restore,omitempty"`

	// Indicates the backup name.
	// Indicates the backup name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies the region in which to create the resource.
	// If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Indicates the size of the backup file (byte).
	// Indicates the size of the backup file (byte).
	Size *float64 `json:"size,omitempty" tf:"size,omitempty"`

	// Indicates the backup status. Valid value:
	// Indicates the backup status.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// Indicates the backup type. Valid value:
	// Indicates the backup type.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type BackupParameters struct {

	// Specifies the format of the DCS instance backup.
	// Value options: aof, rdb. Default to rdb.
	// Specifies the format of the DCS instance backup.
	// +kubebuilder:validation:Optional
	BackupFormat *string `json:"backupFormat,omitempty" tf:"backup_format,omitempty"`

	// Specifies the description of DCS instance backup.
	// Specifies the description of DCS instance backup.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies the ID of the DCS instance.
	// Specifies the ID of the DCS instance.
	// +crossplane:generate:reference:type=github.com/huaweicloud/provider-huaweicloud/apis/dcs/v1alpha1.Instance
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// Reference to a Instance in dcs to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDRef *v1.Reference `json:"instanceIdRef,omitempty" tf:"-"`

	// Selector for a Instance in dcs to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDSelector *v1.Selector `json:"instanceIdSelector,omitempty" tf:"-"`

	// Specifies the region in which to create the resource.
	// If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
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
	Spec              BackupSpec   `json:"spec"`
	Status            BackupStatus `json:"status,omitempty"`
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
