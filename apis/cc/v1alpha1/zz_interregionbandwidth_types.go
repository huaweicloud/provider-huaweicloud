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

type InterRegionBandwidthInitParameters struct {

	// Inter-region bandwidth size.
	// Inter-region bandwidth.
	Bandwidth *float64 `json:"bandwidth,omitempty" tf:"bandwidth,omitempty"`

	// Bandwidth package ID.
	// Bandwidth package ID.
	// +crossplane:generate:reference:type=github.com/huaweicloud/provider-huaweicloud/apis/cc/v1alpha1.BandwidthPackage
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	BandwidthPackageID *string `json:"bandwidthPackageId,omitempty" tf:"bandwidth_package_id,omitempty"`

	// Reference to a BandwidthPackage in cc to populate bandwidthPackageId.
	// +kubebuilder:validation:Optional
	BandwidthPackageIDRef *v1.Reference `json:"bandwidthPackageIdRef,omitempty" tf:"-"`

	// Selector for a BandwidthPackage in cc to populate bandwidthPackageId.
	// +kubebuilder:validation:Optional
	BandwidthPackageIDSelector *v1.Selector `json:"bandwidthPackageIdSelector,omitempty" tf:"-"`

	// Cloud connection ID.
	// Cloud connection ID.
	CloudConnectionID *string `json:"cloudConnectionId,omitempty" tf:"cloud_connection_id,omitempty"`

	// Two regions to which bandwidth is allocated.
	// Two regions to which bandwidth is allocated.
	InterRegionIds []*string `json:"interRegionIds,omitempty" tf:"inter_region_ids,omitempty"`

	// Specifies the region in which to create the resource.
	// If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

type InterRegionBandwidthObservation struct {

	// Inter-region bandwidth size.
	// Inter-region bandwidth.
	Bandwidth *float64 `json:"bandwidth,omitempty" tf:"bandwidth,omitempty"`

	// Bandwidth package ID.
	// Bandwidth package ID.
	BandwidthPackageID *string `json:"bandwidthPackageId,omitempty" tf:"bandwidth_package_id,omitempty"`

	// Cloud connection ID.
	// Cloud connection ID.
	CloudConnectionID *string `json:"cloudConnectionId,omitempty" tf:"cloud_connection_id,omitempty"`

	// The resource ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Two regions to which bandwidth is allocated.
	// Two regions to which bandwidth is allocated.
	InterRegionIds []*string `json:"interRegionIds,omitempty" tf:"inter_region_ids,omitempty"`

	// Details about regions of the inter-region bandwidth.
	// The inter_regions structure is documented below.
	InterRegions []InterRegionsObservation `json:"interRegions,omitempty" tf:"inter_regions,omitempty"`

	// Specifies the region in which to create the resource.
	// If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

type InterRegionBandwidthParameters struct {

	// Inter-region bandwidth size.
	// Inter-region bandwidth.
	// +kubebuilder:validation:Optional
	Bandwidth *float64 `json:"bandwidth,omitempty" tf:"bandwidth,omitempty"`

	// Bandwidth package ID.
	// Bandwidth package ID.
	// +crossplane:generate:reference:type=github.com/huaweicloud/provider-huaweicloud/apis/cc/v1alpha1.BandwidthPackage
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	BandwidthPackageID *string `json:"bandwidthPackageId,omitempty" tf:"bandwidth_package_id,omitempty"`

	// Reference to a BandwidthPackage in cc to populate bandwidthPackageId.
	// +kubebuilder:validation:Optional
	BandwidthPackageIDRef *v1.Reference `json:"bandwidthPackageIdRef,omitempty" tf:"-"`

	// Selector for a BandwidthPackage in cc to populate bandwidthPackageId.
	// +kubebuilder:validation:Optional
	BandwidthPackageIDSelector *v1.Selector `json:"bandwidthPackageIdSelector,omitempty" tf:"-"`

	// Cloud connection ID.
	// Cloud connection ID.
	// +kubebuilder:validation:Optional
	CloudConnectionID *string `json:"cloudConnectionId,omitempty" tf:"cloud_connection_id,omitempty"`

	// Two regions to which bandwidth is allocated.
	// Two regions to which bandwidth is allocated.
	// +kubebuilder:validation:Optional
	InterRegionIds []*string `json:"interRegionIds,omitempty" tf:"inter_region_ids,omitempty"`

	// Specifies the region in which to create the resource.
	// If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

type InterRegionsInitParameters struct {
}

type InterRegionsObservation struct {

	// Inter-region bandwidth ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// ID of the local region where the inter-region bandwidth is used.
	LocalRegionID *string `json:"localRegionId,omitempty" tf:"local_region_id,omitempty"`

	// Project ID of a region where the inter-region bandwidth is used.
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// ID of the remote region where the inter-region bandwidth is used.
	RemoteRegionID *string `json:"remoteRegionId,omitempty" tf:"remote_region_id,omitempty"`
}

type InterRegionsParameters struct {
}

// InterRegionBandwidthSpec defines the desired state of InterRegionBandwidth
type InterRegionBandwidthSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     InterRegionBandwidthParameters `json:"forProvider"`
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
	InitProvider InterRegionBandwidthInitParameters `json:"initProvider,omitempty"`
}

// InterRegionBandwidthStatus defines the observed state of InterRegionBandwidth.
type InterRegionBandwidthStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        InterRegionBandwidthObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// InterRegionBandwidth is the Schema for the InterRegionBandwidths API. ""
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,huaweicloud}
type InterRegionBandwidth struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.bandwidth) || (has(self.initProvider) && has(self.initProvider.bandwidth))",message="spec.forProvider.bandwidth is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.cloudConnectionId) || (has(self.initProvider) && has(self.initProvider.cloudConnectionId))",message="spec.forProvider.cloudConnectionId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.interRegionIds) || (has(self.initProvider) && has(self.initProvider.interRegionIds))",message="spec.forProvider.interRegionIds is a required parameter"
	Spec   InterRegionBandwidthSpec   `json:"spec"`
	Status InterRegionBandwidthStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// InterRegionBandwidthList contains a list of InterRegionBandwidths
type InterRegionBandwidthList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []InterRegionBandwidth `json:"items"`
}

// Repository type metadata.
var (
	InterRegionBandwidth_Kind             = "InterRegionBandwidth"
	InterRegionBandwidth_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: InterRegionBandwidth_Kind}.String()
	InterRegionBandwidth_KindAPIVersion   = InterRegionBandwidth_Kind + "." + CRDGroupVersion.String()
	InterRegionBandwidth_GroupVersionKind = CRDGroupVersion.WithKind(InterRegionBandwidth_Kind)
)

func init() {
	SchemeBuilder.Register(&InterRegionBandwidth{}, &InterRegionBandwidthList{})
}
