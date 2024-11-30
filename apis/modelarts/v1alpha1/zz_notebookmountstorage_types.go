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

type NotebookMountStorageInitParameters struct {

	// Specifies the local mount directory.
	// Only the subdirectory of /data/ can be mounted. The format is : /data/dir1/.
	// Changing this parameter will create a new resource.
	LocalMountDirectory *string `json:"localMountDirectory,omitempty" tf:"local_mount_directory,omitempty"`

	// Specifies ID of notebook which the storage be mounted to.
	// Changing this parameter will create a new resource.
	// +crossplane:generate:reference:type=github.com/huaweicloud/provider-huaweicloud/apis/modelarts/v1alpha1.Notebook
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	NotebookID *string `json:"notebookId,omitempty" tf:"notebook_id,omitempty"`

	// Reference to a Notebook in modelarts to populate notebookId.
	// +kubebuilder:validation:Optional
	NotebookIDRef *v1.Reference `json:"notebookIdRef,omitempty" tf:"-"`

	// Selector for a Notebook in modelarts to populate notebookId.
	// +kubebuilder:validation:Optional
	NotebookIDSelector *v1.Selector `json:"notebookIdSelector,omitempty" tf:"-"`

	// The region in which to create the resource. If omitted, the
	// provider-level region will be used. Changing this parameter will create a new resource.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Specifies the path of Parallel File System (PFS) or its folders in OBS.
	// The format is : obs://obs-bucket/folder/. Changing this parameter will create a new resource.
	StoragePath *string `json:"storagePath,omitempty" tf:"storage_path,omitempty"`
}

type NotebookMountStorageObservation struct {

	// The resource ID in format of notebook_id/mount_id. It is composed of the ID of notebook and mount ID,
	// separated by a slash.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the local mount directory.
	// Only the subdirectory of /data/ can be mounted. The format is : /data/dir1/.
	// Changing this parameter will create a new resource.
	LocalMountDirectory *string `json:"localMountDirectory,omitempty" tf:"local_mount_directory,omitempty"`

	// The mount ID.
	MountID *string `json:"mountId,omitempty" tf:"mount_id,omitempty"`

	// Specifies ID of notebook which the storage be mounted to.
	// Changing this parameter will create a new resource.
	NotebookID *string `json:"notebookId,omitempty" tf:"notebook_id,omitempty"`

	// The region in which to create the resource. If omitted, the
	// provider-level region will be used. Changing this parameter will create a new resource.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// mount status. Valid values include: MOUNTING, MOUNT_FAILED, MOUNTED, UNMOUNTING,
	// UNMOUNT_FAILED, UNMOUNTED.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// Specifies the path of Parallel File System (PFS) or its folders in OBS.
	// The format is : obs://obs-bucket/folder/. Changing this parameter will create a new resource.
	StoragePath *string `json:"storagePath,omitempty" tf:"storage_path,omitempty"`

	// The type of storage system.  The value is OBSFS.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type NotebookMountStorageParameters struct {

	// Specifies the local mount directory.
	// Only the subdirectory of /data/ can be mounted. The format is : /data/dir1/.
	// Changing this parameter will create a new resource.
	// +kubebuilder:validation:Optional
	LocalMountDirectory *string `json:"localMountDirectory,omitempty" tf:"local_mount_directory,omitempty"`

	// Specifies ID of notebook which the storage be mounted to.
	// Changing this parameter will create a new resource.
	// +crossplane:generate:reference:type=github.com/huaweicloud/provider-huaweicloud/apis/modelarts/v1alpha1.Notebook
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	NotebookID *string `json:"notebookId,omitempty" tf:"notebook_id,omitempty"`

	// Reference to a Notebook in modelarts to populate notebookId.
	// +kubebuilder:validation:Optional
	NotebookIDRef *v1.Reference `json:"notebookIdRef,omitempty" tf:"-"`

	// Selector for a Notebook in modelarts to populate notebookId.
	// +kubebuilder:validation:Optional
	NotebookIDSelector *v1.Selector `json:"notebookIdSelector,omitempty" tf:"-"`

	// The region in which to create the resource. If omitted, the
	// provider-level region will be used. Changing this parameter will create a new resource.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Specifies the path of Parallel File System (PFS) or its folders in OBS.
	// The format is : obs://obs-bucket/folder/. Changing this parameter will create a new resource.
	// +kubebuilder:validation:Optional
	StoragePath *string `json:"storagePath,omitempty" tf:"storage_path,omitempty"`
}

// NotebookMountStorageSpec defines the desired state of NotebookMountStorage
type NotebookMountStorageSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NotebookMountStorageParameters `json:"forProvider"`
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
	InitProvider NotebookMountStorageInitParameters `json:"initProvider,omitempty"`
}

// NotebookMountStorageStatus defines the observed state of NotebookMountStorage.
type NotebookMountStorageStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NotebookMountStorageObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// NotebookMountStorage is the Schema for the NotebookMountStorages API. ""
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,huaweicloud}
type NotebookMountStorage struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.localMountDirectory) || (has(self.initProvider) && has(self.initProvider.localMountDirectory))",message="spec.forProvider.localMountDirectory is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.storagePath) || (has(self.initProvider) && has(self.initProvider.storagePath))",message="spec.forProvider.storagePath is a required parameter"
	Spec   NotebookMountStorageSpec   `json:"spec"`
	Status NotebookMountStorageStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NotebookMountStorageList contains a list of NotebookMountStorages
type NotebookMountStorageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NotebookMountStorage `json:"items"`
}

// Repository type metadata.
var (
	NotebookMountStorage_Kind             = "NotebookMountStorage"
	NotebookMountStorage_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: NotebookMountStorage_Kind}.String()
	NotebookMountStorage_KindAPIVersion   = NotebookMountStorage_Kind + "." + CRDGroupVersion.String()
	NotebookMountStorage_GroupVersionKind = CRDGroupVersion.WithKind(NotebookMountStorage_Kind)
)

func init() {
	SchemeBuilder.Register(&NotebookMountStorage{}, &NotebookMountStorageList{})
}