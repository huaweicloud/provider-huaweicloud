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

type ExtendLogLinkInitParameters struct {

	// Specifies the name of the file to be downloaded.
	// Specifies the name of the file to be downloaded.
	FileName *string `json:"fileName,omitempty" tf:"file_name,omitempty"`

	// Specifies the ID of the RDS instance.
	// Specifies the ID of the RDS instance.
	// +crossplane:generate:reference:type=github.com/huaweicloud/provider-huaweicloud/apis/rds/v1alpha1.Instance
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// Reference to a Instance in rds to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDRef *v1.Reference `json:"instanceIdRef,omitempty" tf:"-"`

	// Selector for a Instance in rds to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDSelector *v1.Selector `json:"instanceIdSelector,omitempty" tf:"-"`

	// Specifies the region in which to create the resource.
	// If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

type ExtendLogLinkObservation struct {

	// Indicates the creation time.
	// Indicates the creation time.
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// Indicates the download link.
	// Indicates the download link.
	FileLink *string `json:"fileLink,omitempty" tf:"file_link,omitempty"`

	// Specifies the name of the file to be downloaded.
	// Specifies the name of the file to be downloaded.
	FileName *string `json:"fileName,omitempty" tf:"file_name,omitempty"`

	// Indicates the file size in KB.
	// Indicates the file size in KB.
	FileSize *string `json:"fileSize,omitempty" tf:"file_size,omitempty"`

	// The resource ID. The format is  <instance_id>/<file_name>.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the ID of the RDS instance.
	// Specifies the ID of the RDS instance.
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// Specifies the region in which to create the resource.
	// If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Indicates the last update time.
	// Indicates the last update time.
	UpdatedAt *string `json:"updatedAt,omitempty" tf:"updated_at,omitempty"`
}

type ExtendLogLinkParameters struct {

	// Specifies the name of the file to be downloaded.
	// Specifies the name of the file to be downloaded.
	// +kubebuilder:validation:Optional
	FileName *string `json:"fileName,omitempty" tf:"file_name,omitempty"`

	// Specifies the ID of the RDS instance.
	// Specifies the ID of the RDS instance.
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

	// Specifies the region in which to create the resource.
	// If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

// ExtendLogLinkSpec defines the desired state of ExtendLogLink
type ExtendLogLinkSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ExtendLogLinkParameters `json:"forProvider"`
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
	InitProvider ExtendLogLinkInitParameters `json:"initProvider,omitempty"`
}

// ExtendLogLinkStatus defines the observed state of ExtendLogLink.
type ExtendLogLinkStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ExtendLogLinkObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ExtendLogLink is the Schema for the ExtendLogLinks API. Manages an RDS extend log link resource within HuaweiCloud.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,huaweicloud}
type ExtendLogLink struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.fileName) || (has(self.initProvider) && has(self.initProvider.fileName))",message="spec.forProvider.fileName is a required parameter"
	Spec   ExtendLogLinkSpec   `json:"spec"`
	Status ExtendLogLinkStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ExtendLogLinkList contains a list of ExtendLogLinks
type ExtendLogLinkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ExtendLogLink `json:"items"`
}

// Repository type metadata.
var (
	ExtendLogLink_Kind             = "ExtendLogLink"
	ExtendLogLink_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ExtendLogLink_Kind}.String()
	ExtendLogLink_KindAPIVersion   = ExtendLogLink_Kind + "." + CRDGroupVersion.String()
	ExtendLogLink_GroupVersionKind = CRDGroupVersion.WithKind(ExtendLogLink_Kind)
)

func init() {
	SchemeBuilder.Register(&ExtendLogLink{}, &ExtendLogLinkList{})
}
