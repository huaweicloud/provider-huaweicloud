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

type GlobalEipInitParameters struct {

	// Specifies the access site name.
	// Changing this creates a new resource.
	AccessSite *string `json:"accessSite,omitempty" tf:"access_site,omitempty"`

	// Specifies the description of the global EIP.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies the enterprise project id to which the global EIP
	// belongs. Changing this creates a new resource.
	EnterpriseProjectID *string `json:"enterpriseProjectId,omitempty" tf:"enterprise_project_id,omitempty"`

	// Specifies the global EIP pool name.
	// Changing this creates a new resource.
	GeipPoolName *string `json:"geipPoolName,omitempty" tf:"geip_pool_name,omitempty"`

	// Specifies the internet bandwidth id which the global EIP use.
	// Changing this creates a new resource.
	InternetBandwidthID *string `json:"internetBandwidthId,omitempty" tf:"internet_bandwidth_id,omitempty"`

	// Specifies the name of the global EIP.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies the tags of the global EIP.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type GlobalEipObservation struct {

	// Specifies the access site name.
	// Changing this creates a new resource.
	AccessSite *string `json:"accessSite,omitempty" tf:"access_site,omitempty"`

	// The ID of the associate instance.
	AssociateInstanceID *string `json:"associateInstanceId,omitempty" tf:"associate_instance_id,omitempty"`

	// The region of the associate instance.
	AssociateInstanceRegion *string `json:"associateInstanceRegion,omitempty" tf:"associate_instance_region,omitempty"`

	// The type of the associate instance.
	AssociateInstanceType *string `json:"associateInstanceType,omitempty" tf:"associate_instance_type,omitempty"`

	// The create time of the global EIP.
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// Specifies the description of the global EIP.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies the enterprise project id to which the global EIP
	// belongs. Changing this creates a new resource.
	EnterpriseProjectID *string `json:"enterpriseProjectId,omitempty" tf:"enterprise_project_id,omitempty"`

	// The global EIP is frozen or not.
	Frozen *bool `json:"frozen,omitempty" tf:"frozen,omitempty"`

	// The frozen info of the global EIP.
	FrozenInfo *string `json:"frozenInfo,omitempty" tf:"frozen_info,omitempty"`

	// Specifies the global EIP pool name.
	// Changing this creates a new resource.
	GeipPoolName *string `json:"geipPoolName,omitempty" tf:"geip_pool_name,omitempty"`

	// The ID of the global connection bandwidth.
	GlobalConnectionBandwidthID *string `json:"globalConnectionBandwidthId,omitempty" tf:"global_connection_bandwidth_id,omitempty"`

	// The type of the global connection bandwidth.
	GlobalConnectionBandwidthType *string `json:"globalConnectionBandwidthType,omitempty" tf:"global_connection_bandwidth_type,omitempty"`

	// The resource ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ip address of the global EIP.
	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`

	// The ip version of the global EIP.
	IPVersion *float64 `json:"ipVersion,omitempty" tf:"ip_version,omitempty"`

	// Specifies the internet bandwidth id which the global EIP use.
	// Changing this creates a new resource.
	InternetBandwidthID *string `json:"internetBandwidthId,omitempty" tf:"internet_bandwidth_id,omitempty"`

	// The the internet service provider of the global EIP.
	Isp *string `json:"isp,omitempty" tf:"isp,omitempty"`

	// Specifies the name of the global EIP.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The global EIP is polluted or not.
	Polluted *bool `json:"polluted,omitempty" tf:"polluted,omitempty"`

	// The status of the global EIP.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// Specifies the tags of the global EIP.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The update time of the global EIP.
	UpdatedAt *string `json:"updatedAt,omitempty" tf:"updated_at,omitempty"`
}

type GlobalEipParameters struct {

	// Specifies the access site name.
	// Changing this creates a new resource.
	// +kubebuilder:validation:Optional
	AccessSite *string `json:"accessSite,omitempty" tf:"access_site,omitempty"`

	// Specifies the description of the global EIP.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies the enterprise project id to which the global EIP
	// belongs. Changing this creates a new resource.
	// +kubebuilder:validation:Optional
	EnterpriseProjectID *string `json:"enterpriseProjectId,omitempty" tf:"enterprise_project_id,omitempty"`

	// Specifies the global EIP pool name.
	// Changing this creates a new resource.
	// +kubebuilder:validation:Optional
	GeipPoolName *string `json:"geipPoolName,omitempty" tf:"geip_pool_name,omitempty"`

	// Specifies the internet bandwidth id which the global EIP use.
	// Changing this creates a new resource.
	// +kubebuilder:validation:Optional
	InternetBandwidthID *string `json:"internetBandwidthId,omitempty" tf:"internet_bandwidth_id,omitempty"`

	// Specifies the name of the global EIP.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies the tags of the global EIP.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// GlobalEipSpec defines the desired state of GlobalEip
type GlobalEipSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     GlobalEipParameters `json:"forProvider"`
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
	InitProvider GlobalEipInitParameters `json:"initProvider,omitempty"`
}

// GlobalEipStatus defines the observed state of GlobalEip.
type GlobalEipStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        GlobalEipObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// GlobalEip is the Schema for the GlobalEips API. ""
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,huaweicloud}
type GlobalEip struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.accessSite) || (has(self.initProvider) && has(self.initProvider.accessSite))",message="spec.forProvider.accessSite is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.geipPoolName) || (has(self.initProvider) && has(self.initProvider.geipPoolName))",message="spec.forProvider.geipPoolName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.internetBandwidthId) || (has(self.initProvider) && has(self.initProvider.internetBandwidthId))",message="spec.forProvider.internetBandwidthId is a required parameter"
	Spec   GlobalEipSpec   `json:"spec"`
	Status GlobalEipStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GlobalEipList contains a list of GlobalEips
type GlobalEipList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GlobalEip `json:"items"`
}

// Repository type metadata.
var (
	GlobalEip_Kind             = "GlobalEip"
	GlobalEip_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: GlobalEip_Kind}.String()
	GlobalEip_KindAPIVersion   = GlobalEip_Kind + "." + CRDGroupVersion.String()
	GlobalEip_GroupVersionKind = CRDGroupVersion.WithKind(GlobalEip_Kind)
)

func init() {
	SchemeBuilder.Register(&GlobalEip{}, &GlobalEipList{})
}