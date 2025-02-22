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

type SnatRuleInitParameters struct {

	// Specifies the CIDR block connected by SNAT rule (DC side).
	// This parameter and subnet_id are alternative. Changing this will create a new resource.
	// The CIDR block connected by SNAT rule (DC side).
	Cidr *string `json:"cidr,omitempty" tf:"cidr,omitempty"`

	// Specifies the description of the SNAT rule.
	// The value is a string of no more than 255 characters, and angle brackets (<>) are not allowed.
	// The description of the SNAT rule.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies the IDs of floating IPs connected by SNAT rule.
	// Multiple floating IPs are separated using commas (,). The number of floating IP IDs cannot exceed 20.
	// The IDs of floating IPs connected by SNAT rule.
	// +crossplane:generate:reference:type=github.com/huaweicloud/provider-huaweicloud/apis/eip/v1alpha1.VpcEip
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	FloatingIPID *string `json:"floatingIpId,omitempty" tf:"floating_ip_id,omitempty"`

	// Reference to a VpcEip in eip to populate floatingIpId.
	// +kubebuilder:validation:Optional
	FloatingIPIDRef *v1.Reference `json:"floatingIpIdRef,omitempty" tf:"-"`

	// Selector for a VpcEip in eip to populate floatingIpId.
	// +kubebuilder:validation:Optional
	FloatingIPIDSelector *v1.Selector `json:"floatingIpIdSelector,omitempty" tf:"-"`

	// Specifies the IDs of global EIPs connected by SNAT rule.
	// Multiple global EIPs are separated using commas (,). The number of global EIP IDs cannot exceed 20.
	// The IDs (separated by commas) of global EIPs connected by SNAT rule.
	// +crossplane:generate:reference:type=github.com/huaweicloud/provider-huaweicloud/apis/eip/v1alpha1.GlobalEip
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	GlobalEIPID *string `json:"globalEipId,omitempty" tf:"global_eip_id,omitempty"`

	// Reference to a GlobalEip in eip to populate globalEipId.
	// +kubebuilder:validation:Optional
	GlobalEIPIDRef *v1.Reference `json:"globalEipIdRef,omitempty" tf:"-"`

	// Selector for a GlobalEip in eip to populate globalEipId.
	// +kubebuilder:validation:Optional
	GlobalEIPIDSelector *v1.Selector `json:"globalEipIdSelector,omitempty" tf:"-"`

	// Specifies the ID of the gateway to which the SNAT rule belongs.
	// Changing this will create a new resource.
	// schema: Required; The ID of the gateway to which the SNAT rule belongs.
	// +crossplane:generate:reference:type=github.com/huaweicloud/provider-huaweicloud/apis/nat/v1alpha1.Gateway
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	NATGatewayID *string `json:"natGatewayId,omitempty" tf:"nat_gateway_id,omitempty"`

	// Reference to a Gateway in nat to populate natGatewayId.
	// +kubebuilder:validation:Optional
	NATGatewayIDRef *v1.Reference `json:"natGatewayIdRef,omitempty" tf:"-"`

	// Selector for a Gateway in nat to populate natGatewayId.
	// +kubebuilder:validation:Optional
	NATGatewayIDSelector *v1.Selector `json:"natGatewayIdSelector,omitempty" tf:"-"`

	// Specifies a resource ID in UUID format.
	NetworkID *string `json:"networkId,omitempty" tf:"network_id,omitempty"`

	// Specifies the region where the SNAT rule is located.
	// If omitted, the provider-level region will be used. Changing this will create a new resource.
	// The region where the SNAT rule is located.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Specifies the resource scenario.
	// The valid values are 0 (VPC scenario) and 1 (Direct Connect scenario), and the default value is 0.
	// Only cidr can be specified over a Direct Connect connection. Changing this will create a new resource.
	// The resource type of the SNAT rule.
	SourceType *float64 `json:"sourceType,omitempty" tf:"source_type,omitempty"`

	// Specifies the network IDs of subnet connected by SNAT rule (VPC side).
	// This parameter and cidr are alternative. Changing this will create a new resource.
	// The network IDs of subnet connected by SNAT rule (VPC side).
	// +crossplane:generate:reference:type=github.com/huaweicloud/provider-huaweicloud/apis/vpc/v1alpha1.Subnet
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Reference to a Subnet in vpc to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// Selector for a Subnet in vpc to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`
}

type SnatRuleObservation struct {

	// Specifies the CIDR block connected by SNAT rule (DC side).
	// This parameter and subnet_id are alternative. Changing this will create a new resource.
	// The CIDR block connected by SNAT rule (DC side).
	Cidr *string `json:"cidr,omitempty" tf:"cidr,omitempty"`

	// Specifies the description of the SNAT rule.
	// The value is a string of no more than 255 characters, and angle brackets (<>) are not allowed.
	// The description of the SNAT rule.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The actual floating IP address.
	// The floating IP addresses (separated by commas) connected by SNAT rule.
	FloatingIPAddress *string `json:"floatingIpAddress,omitempty" tf:"floating_ip_address,omitempty"`

	// Specifies the IDs of floating IPs connected by SNAT rule.
	// Multiple floating IPs are separated using commas (,). The number of floating IP IDs cannot exceed 20.
	// The IDs of floating IPs connected by SNAT rule.
	FloatingIPID *string `json:"floatingIpId,omitempty" tf:"floating_ip_id,omitempty"`

	// The global EIP addresses (separated by commas) connected by SNAT rule.
	// The global EIP addresses (separated by commas) connected by SNAT rule.
	GlobalEIPAddress *string `json:"globalEipAddress,omitempty" tf:"global_eip_address,omitempty"`

	// Specifies the IDs of global EIPs connected by SNAT rule.
	// Multiple global EIPs are separated using commas (,). The number of global EIP IDs cannot exceed 20.
	// The IDs (separated by commas) of global EIPs connected by SNAT rule.
	GlobalEIPID *string `json:"globalEipId,omitempty" tf:"global_eip_id,omitempty"`

	// Specifies a resource ID in UUID format.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the ID of the gateway to which the SNAT rule belongs.
	// Changing this will create a new resource.
	// schema: Required; The ID of the gateway to which the SNAT rule belongs.
	NATGatewayID *string `json:"natGatewayId,omitempty" tf:"nat_gateway_id,omitempty"`

	// Specifies a resource ID in UUID format.
	NetworkID *string `json:"networkId,omitempty" tf:"network_id,omitempty"`

	// Specifies the region where the SNAT rule is located.
	// If omitted, the provider-level region will be used. Changing this will create a new resource.
	// The region where the SNAT rule is located.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Specifies the resource scenario.
	// The valid values are 0 (VPC scenario) and 1 (Direct Connect scenario), and the default value is 0.
	// Only cidr can be specified over a Direct Connect connection. Changing this will create a new resource.
	// The resource type of the SNAT rule.
	SourceType *float64 `json:"sourceType,omitempty" tf:"source_type,omitempty"`

	// The status of the SNAT rule.
	// The status of the SNAT rule.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// Specifies the network IDs of subnet connected by SNAT rule (VPC side).
	// This parameter and cidr are alternative. Changing this will create a new resource.
	// The network IDs of subnet connected by SNAT rule (VPC side).
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`
}

type SnatRuleParameters struct {

	// Specifies the CIDR block connected by SNAT rule (DC side).
	// This parameter and subnet_id are alternative. Changing this will create a new resource.
	// The CIDR block connected by SNAT rule (DC side).
	// +kubebuilder:validation:Optional
	Cidr *string `json:"cidr,omitempty" tf:"cidr,omitempty"`

	// Specifies the description of the SNAT rule.
	// The value is a string of no more than 255 characters, and angle brackets (<>) are not allowed.
	// The description of the SNAT rule.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies the IDs of floating IPs connected by SNAT rule.
	// Multiple floating IPs are separated using commas (,). The number of floating IP IDs cannot exceed 20.
	// The IDs of floating IPs connected by SNAT rule.
	// +crossplane:generate:reference:type=github.com/huaweicloud/provider-huaweicloud/apis/eip/v1alpha1.VpcEip
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	FloatingIPID *string `json:"floatingIpId,omitempty" tf:"floating_ip_id,omitempty"`

	// Reference to a VpcEip in eip to populate floatingIpId.
	// +kubebuilder:validation:Optional
	FloatingIPIDRef *v1.Reference `json:"floatingIpIdRef,omitempty" tf:"-"`

	// Selector for a VpcEip in eip to populate floatingIpId.
	// +kubebuilder:validation:Optional
	FloatingIPIDSelector *v1.Selector `json:"floatingIpIdSelector,omitempty" tf:"-"`

	// Specifies the IDs of global EIPs connected by SNAT rule.
	// Multiple global EIPs are separated using commas (,). The number of global EIP IDs cannot exceed 20.
	// The IDs (separated by commas) of global EIPs connected by SNAT rule.
	// +crossplane:generate:reference:type=github.com/huaweicloud/provider-huaweicloud/apis/eip/v1alpha1.GlobalEip
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	GlobalEIPID *string `json:"globalEipId,omitempty" tf:"global_eip_id,omitempty"`

	// Reference to a GlobalEip in eip to populate globalEipId.
	// +kubebuilder:validation:Optional
	GlobalEIPIDRef *v1.Reference `json:"globalEipIdRef,omitempty" tf:"-"`

	// Selector for a GlobalEip in eip to populate globalEipId.
	// +kubebuilder:validation:Optional
	GlobalEIPIDSelector *v1.Selector `json:"globalEipIdSelector,omitempty" tf:"-"`

	// Specifies the ID of the gateway to which the SNAT rule belongs.
	// Changing this will create a new resource.
	// schema: Required; The ID of the gateway to which the SNAT rule belongs.
	// +crossplane:generate:reference:type=github.com/huaweicloud/provider-huaweicloud/apis/nat/v1alpha1.Gateway
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	NATGatewayID *string `json:"natGatewayId,omitempty" tf:"nat_gateway_id,omitempty"`

	// Reference to a Gateway in nat to populate natGatewayId.
	// +kubebuilder:validation:Optional
	NATGatewayIDRef *v1.Reference `json:"natGatewayIdRef,omitempty" tf:"-"`

	// Selector for a Gateway in nat to populate natGatewayId.
	// +kubebuilder:validation:Optional
	NATGatewayIDSelector *v1.Selector `json:"natGatewayIdSelector,omitempty" tf:"-"`

	// Specifies a resource ID in UUID format.
	// +kubebuilder:validation:Optional
	NetworkID *string `json:"networkId,omitempty" tf:"network_id,omitempty"`

	// Specifies the region where the SNAT rule is located.
	// If omitted, the provider-level region will be used. Changing this will create a new resource.
	// The region where the SNAT rule is located.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Specifies the resource scenario.
	// The valid values are 0 (VPC scenario) and 1 (Direct Connect scenario), and the default value is 0.
	// Only cidr can be specified over a Direct Connect connection. Changing this will create a new resource.
	// The resource type of the SNAT rule.
	// +kubebuilder:validation:Optional
	SourceType *float64 `json:"sourceType,omitempty" tf:"source_type,omitempty"`

	// Specifies the network IDs of subnet connected by SNAT rule (VPC side).
	// This parameter and cidr are alternative. Changing this will create a new resource.
	// The network IDs of subnet connected by SNAT rule (VPC side).
	// +crossplane:generate:reference:type=github.com/huaweicloud/provider-huaweicloud/apis/vpc/v1alpha1.Subnet
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Reference to a Subnet in vpc to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// Selector for a Subnet in vpc to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`
}

// SnatRuleSpec defines the desired state of SnatRule
type SnatRuleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SnatRuleParameters `json:"forProvider"`
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
	InitProvider SnatRuleInitParameters `json:"initProvider,omitempty"`
}

// SnatRuleStatus defines the observed state of SnatRule.
type SnatRuleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SnatRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// SnatRule is the Schema for the SnatRules API. ""
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,huaweicloud}
type SnatRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SnatRuleSpec   `json:"spec"`
	Status            SnatRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SnatRuleList contains a list of SnatRules
type SnatRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SnatRule `json:"items"`
}

// Repository type metadata.
var (
	SnatRule_Kind             = "SnatRule"
	SnatRule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SnatRule_Kind}.String()
	SnatRule_KindAPIVersion   = SnatRule_Kind + "." + CRDGroupVersion.String()
	SnatRule_GroupVersionKind = CRDGroupVersion.WithKind(SnatRule_Kind)
)

func init() {
	SchemeBuilder.Register(&SnatRule{}, &SnatRuleList{})
}
