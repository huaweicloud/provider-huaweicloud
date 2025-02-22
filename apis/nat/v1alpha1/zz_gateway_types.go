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

type GatewayInitParameters struct {

	// Specifies whether auto-renew is enabled. This parameter is only valid when
	// charging_mode is set to prePaid. Valid values are true and false. Defaults to false.
	AutoRenew *string `json:"autoRenew,omitempty" tf:"auto_renew,omitempty"`

	// Specifies the charging mode of the NAT gateway.
	// The valid values are as follows:
	ChargingMode *string `json:"chargingMode,omitempty" tf:"charging_mode,omitempty"`

	// Specifies the description of the NAT gateway, which contain maximum of 512
	// characters, and angle brackets (<) and (>) are not allowed.
	// The description of the NAT gateway.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies the enterprise project ID of the NAT gateway.
	// Changing this will create a new resource.
	// The enterprise project ID of the NAT gateway.
	EnterpriseProjectID *string `json:"enterpriseProjectId,omitempty" tf:"enterprise_project_id,omitempty"`

	// Specifies the NAT gateway name.
	// The valid length is limited from 1 to 64, only letters, digits, hyphens (-) and underscores (_) are allowed.
	// The NAT gateway name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies the private IP address of the NAT gateway.
	// The IP address must be one of the IP addresses of the VPC subnet associated with the NAT gateway.
	// If not spacified, it will be automatically allocated.
	// Changing this will creates a new resource.
	// The private IP address of the NAT gateway.
	NgportIPAddress *string `json:"ngportIpAddress,omitempty" tf:"ngport_ip_address,omitempty"`

	// Specifies the charging period of the NAT gateway.
	// If period_unit is set to month, the value ranges from 1 to 9.
	// If period_unit is set to year, the value ranges from 1 to 3.
	// This parameter is mandatory if charging_mode is set to prePaid.
	// Changing this will create a new resource.
	Period *float64 `json:"period,omitempty" tf:"period,omitempty"`

	// Specifies the charging period unit of the NAT gateway.
	// Valid values are month and year. This parameter is mandatory if charging_mode is set to prePaid.
	// Changing this will create a new resource.
	PeriodUnit *string `json:"periodUnit,omitempty" tf:"period_unit,omitempty"`

	// Specifies the region where the NAT gateway is located.
	// If omitted, the provider-level region will be used. Changing this will create a new resource.
	// The region where the NAT gateway is located.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Specifies the specification of the NAT gateway. The valid values are as follows:
	// The specification of the NAT gateway.
	Spec *string `json:"spec,omitempty" tf:"spec,omitempty"`

	// Specifies the subnet ID of the downstream interface (the next hop of the
	// DVR) of the NAT gateway.
	// Changing this will create a new resource.
	// The network ID of the downstream interface (the next hop of the DVR) of the NAT gateway.
	// +crossplane:generate:reference:type=github.com/huaweicloud/provider-huaweicloud/apis/vpc/v1alpha1.Subnet
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Reference to a Subnet in vpc to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// Selector for a Subnet in vpc to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// Specifies the key/value pairs to associate with the NAT gateway.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Specifies the ID of the VPC to which the NAT gateway belongs.
	// Changing this will create a new resource.
	// The ID of the VPC to which the NAT gateway belongs.
	// +crossplane:generate:reference:type=github.com/huaweicloud/provider-huaweicloud/apis/vpc/v1alpha1.VPC
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// Reference to a VPC in vpc to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDRef *v1.Reference `json:"vpcIdRef,omitempty" tf:"-"`

	// Selector for a VPC in vpc to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDSelector *v1.Selector `json:"vpcIdSelector,omitempty" tf:"-"`
}

type GatewayObservation struct {

	// Specifies whether auto-renew is enabled. This parameter is only valid when
	// charging_mode is set to prePaid. Valid values are true and false. Defaults to false.
	AutoRenew *string `json:"autoRenew,omitempty" tf:"auto_renew,omitempty"`

	// Specifies the charging mode of the NAT gateway.
	// The valid values are as follows:
	ChargingMode *string `json:"chargingMode,omitempty" tf:"charging_mode,omitempty"`

	// Specifies the description of the NAT gateway, which contain maximum of 512
	// characters, and angle brackets (<) and (>) are not allowed.
	// The description of the NAT gateway.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies the enterprise project ID of the NAT gateway.
	// Changing this will create a new resource.
	// The enterprise project ID of the NAT gateway.
	EnterpriseProjectID *string `json:"enterpriseProjectId,omitempty" tf:"enterprise_project_id,omitempty"`

	// The resource ID in UUID format.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the NAT gateway name.
	// The valid length is limited from 1 to 64, only letters, digits, hyphens (-) and underscores (_) are allowed.
	// The NAT gateway name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies the private IP address of the NAT gateway.
	// The IP address must be one of the IP addresses of the VPC subnet associated with the NAT gateway.
	// If not spacified, it will be automatically allocated.
	// Changing this will creates a new resource.
	// The private IP address of the NAT gateway.
	NgportIPAddress *string `json:"ngportIpAddress,omitempty" tf:"ngport_ip_address,omitempty"`

	// Specifies the charging period of the NAT gateway.
	// If period_unit is set to month, the value ranges from 1 to 9.
	// If period_unit is set to year, the value ranges from 1 to 3.
	// This parameter is mandatory if charging_mode is set to prePaid.
	// Changing this will create a new resource.
	Period *float64 `json:"period,omitempty" tf:"period,omitempty"`

	// Specifies the charging period unit of the NAT gateway.
	// Valid values are month and year. This parameter is mandatory if charging_mode is set to prePaid.
	// Changing this will create a new resource.
	PeriodUnit *string `json:"periodUnit,omitempty" tf:"period_unit,omitempty"`

	// Specifies the region where the NAT gateway is located.
	// If omitted, the provider-level region will be used. Changing this will create a new resource.
	// The region where the NAT gateway is located.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Specifies the specification of the NAT gateway. The valid values are as follows:
	// The specification of the NAT gateway.
	Spec *string `json:"spec,omitempty" tf:"spec,omitempty"`

	// The current status of the NAT gateway.
	// The current status of the NAT gateway.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// Specifies the subnet ID of the downstream interface (the next hop of the
	// DVR) of the NAT gateway.
	// Changing this will create a new resource.
	// The network ID of the downstream interface (the next hop of the DVR) of the NAT gateway.
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Specifies the key/value pairs to associate with the NAT gateway.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Specifies the ID of the VPC to which the NAT gateway belongs.
	// Changing this will create a new resource.
	// The ID of the VPC to which the NAT gateway belongs.
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`
}

type GatewayParameters struct {

	// Specifies whether auto-renew is enabled. This parameter is only valid when
	// charging_mode is set to prePaid. Valid values are true and false. Defaults to false.
	// +kubebuilder:validation:Optional
	AutoRenew *string `json:"autoRenew,omitempty" tf:"auto_renew,omitempty"`

	// Specifies the charging mode of the NAT gateway.
	// The valid values are as follows:
	// +kubebuilder:validation:Optional
	ChargingMode *string `json:"chargingMode,omitempty" tf:"charging_mode,omitempty"`

	// Specifies the description of the NAT gateway, which contain maximum of 512
	// characters, and angle brackets (<) and (>) are not allowed.
	// The description of the NAT gateway.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies the enterprise project ID of the NAT gateway.
	// Changing this will create a new resource.
	// The enterprise project ID of the NAT gateway.
	// +kubebuilder:validation:Optional
	EnterpriseProjectID *string `json:"enterpriseProjectId,omitempty" tf:"enterprise_project_id,omitempty"`

	// Specifies the NAT gateway name.
	// The valid length is limited from 1 to 64, only letters, digits, hyphens (-) and underscores (_) are allowed.
	// The NAT gateway name.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies the private IP address of the NAT gateway.
	// The IP address must be one of the IP addresses of the VPC subnet associated with the NAT gateway.
	// If not spacified, it will be automatically allocated.
	// Changing this will creates a new resource.
	// The private IP address of the NAT gateway.
	// +kubebuilder:validation:Optional
	NgportIPAddress *string `json:"ngportIpAddress,omitempty" tf:"ngport_ip_address,omitempty"`

	// Specifies the charging period of the NAT gateway.
	// If period_unit is set to month, the value ranges from 1 to 9.
	// If period_unit is set to year, the value ranges from 1 to 3.
	// This parameter is mandatory if charging_mode is set to prePaid.
	// Changing this will create a new resource.
	// +kubebuilder:validation:Optional
	Period *float64 `json:"period,omitempty" tf:"period,omitempty"`

	// Specifies the charging period unit of the NAT gateway.
	// Valid values are month and year. This parameter is mandatory if charging_mode is set to prePaid.
	// Changing this will create a new resource.
	// +kubebuilder:validation:Optional
	PeriodUnit *string `json:"periodUnit,omitempty" tf:"period_unit,omitempty"`

	// Specifies the region where the NAT gateway is located.
	// If omitted, the provider-level region will be used. Changing this will create a new resource.
	// The region where the NAT gateway is located.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Specifies the specification of the NAT gateway. The valid values are as follows:
	// The specification of the NAT gateway.
	// +kubebuilder:validation:Optional
	Spec *string `json:"spec,omitempty" tf:"spec,omitempty"`

	// Specifies the subnet ID of the downstream interface (the next hop of the
	// DVR) of the NAT gateway.
	// Changing this will create a new resource.
	// The network ID of the downstream interface (the next hop of the DVR) of the NAT gateway.
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

	// Specifies the key/value pairs to associate with the NAT gateway.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Specifies the ID of the VPC to which the NAT gateway belongs.
	// Changing this will create a new resource.
	// The ID of the VPC to which the NAT gateway belongs.
	// +crossplane:generate:reference:type=github.com/huaweicloud/provider-huaweicloud/apis/vpc/v1alpha1.VPC
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// Reference to a VPC in vpc to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDRef *v1.Reference `json:"vpcIdRef,omitempty" tf:"-"`

	// Selector for a VPC in vpc to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDSelector *v1.Selector `json:"vpcIdSelector,omitempty" tf:"-"`
}

// GatewaySpec defines the desired state of Gateway
type GatewaySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     GatewayParameters `json:"forProvider"`
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
	InitProvider GatewayInitParameters `json:"initProvider,omitempty"`
}

// GatewayStatus defines the observed state of Gateway.
type GatewayStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        GatewayObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Gateway is the Schema for the Gateways API. ""
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,huaweicloud}
type Gateway struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.spec) || (has(self.initProvider) && has(self.initProvider.spec))",message="spec.forProvider.spec is a required parameter"
	Spec   GatewaySpec   `json:"spec"`
	Status GatewayStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GatewayList contains a list of Gateways
type GatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Gateway `json:"items"`
}

// Repository type metadata.
var (
	Gateway_Kind             = "Gateway"
	Gateway_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Gateway_Kind}.String()
	Gateway_KindAPIVersion   = Gateway_Kind + "." + CRDGroupVersion.String()
	Gateway_GroupVersionKind = CRDGroupVersion.WithKind(Gateway_Kind)
)

func init() {
	SchemeBuilder.Register(&Gateway{}, &GatewayList{})
}
