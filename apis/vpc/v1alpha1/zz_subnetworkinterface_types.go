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

type SubNetworkInterfaceInitParameters struct {

	// Specifies the description of the supplementary network interface.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies the private IPv4 address of the supplementary network interface.
	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`

	// Specifies the IPv6 address is it enabled of the supplementary network interface.
	IPv6Enable *bool `json:"ipv6Enable,omitempty" tf:"ipv6_enable,omitempty"`

	// Specifies the IPv6 address of the supplementary network interface.
	IPv6IPAddress *string `json:"ipv6IpAddress,omitempty" tf:"ipv6_ip_address,omitempty"`

	// Specifies the ID of the elastic network interface to which the
	// supplementary network interface belongs.
	// Changing this creates a new resource.
	ParentID *string `json:"parentId,omitempty" tf:"parent_id,omitempty"`

	// Specifies the region in which to create the resource.
	// If omitted, the provider-level region will be used.
	// Changing this creates a new resource.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Specifies the list of the security groups IDs to which the supplementary
	// network interface belongs.
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" tf:"security_group_ids,omitempty"`

	// Specifies the ID of the subnet to which the supplementary network
	// interface belongs.
	// Changing this creates a new resource.
	// +crossplane:generate:reference:type=github.com/huaweicloud/provider-huaweicloud/apis/vpc/v1alpha1.Subnet
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Reference to a Subnet in vpc to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// Selector for a Subnet in vpc to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// Specifies the vlan ID of the supplementary network interface.
	// The valid value is range from 1 t0 4094.
	VlanID *string `json:"vlanId,omitempty" tf:"vlan_id,omitempty"`
}

type SubNetworkInterfaceObservation struct {

	// The create time of the supplementary network interface.
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// Specifies the description of the supplementary network interface.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The resource ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the private IPv4 address of the supplementary network interface.
	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`

	// Specifies the IPv6 address is it enabled of the supplementary network interface.
	IPv6Enable *bool `json:"ipv6Enable,omitempty" tf:"ipv6_enable,omitempty"`

	// Specifies the IPv6 address of the supplementary network interface.
	IPv6IPAddress *string `json:"ipv6IpAddress,omitempty" tf:"ipv6_ip_address,omitempty"`

	// The MAC address of the supplementary network interface.
	MacAddress *string `json:"macAddress,omitempty" tf:"mac_address,omitempty"`

	// The ID of the ECS to which the supplementary network interface belongs.
	ParentDeviceID *string `json:"parentDeviceId,omitempty" tf:"parent_device_id,omitempty"`

	// Specifies the ID of the elastic network interface to which the
	// supplementary network interface belongs.
	// Changing this creates a new resource.
	ParentID *string `json:"parentId,omitempty" tf:"parent_id,omitempty"`

	// Specifies the region in which to create the resource.
	// If omitted, the provider-level region will be used.
	// Changing this creates a new resource.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Specifies the list of the security groups IDs to which the supplementary
	// network interface belongs.
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" tf:"security_group_ids,omitempty"`

	// The status of the supplementary network interface.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// Specifies the ID of the subnet to which the supplementary network
	// interface belongs.
	// Changing this creates a new resource.
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// The ID of the VPC to which the supplementary network interface belongs.
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// Specifies the vlan ID of the supplementary network interface.
	// The valid value is range from 1 t0 4094.
	VlanID *string `json:"vlanId,omitempty" tf:"vlan_id,omitempty"`
}

type SubNetworkInterfaceParameters struct {

	// Specifies the description of the supplementary network interface.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies the private IPv4 address of the supplementary network interface.
	// +kubebuilder:validation:Optional
	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`

	// Specifies the IPv6 address is it enabled of the supplementary network interface.
	// +kubebuilder:validation:Optional
	IPv6Enable *bool `json:"ipv6Enable,omitempty" tf:"ipv6_enable,omitempty"`

	// Specifies the IPv6 address of the supplementary network interface.
	// +kubebuilder:validation:Optional
	IPv6IPAddress *string `json:"ipv6IpAddress,omitempty" tf:"ipv6_ip_address,omitempty"`

	// Specifies the ID of the elastic network interface to which the
	// supplementary network interface belongs.
	// Changing this creates a new resource.
	// +kubebuilder:validation:Optional
	ParentID *string `json:"parentId,omitempty" tf:"parent_id,omitempty"`

	// Specifies the region in which to create the resource.
	// If omitted, the provider-level region will be used.
	// Changing this creates a new resource.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Specifies the list of the security groups IDs to which the supplementary
	// network interface belongs.
	// +kubebuilder:validation:Optional
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" tf:"security_group_ids,omitempty"`

	// Specifies the ID of the subnet to which the supplementary network
	// interface belongs.
	// Changing this creates a new resource.
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

	// Specifies the vlan ID of the supplementary network interface.
	// The valid value is range from 1 t0 4094.
	// +kubebuilder:validation:Optional
	VlanID *string `json:"vlanId,omitempty" tf:"vlan_id,omitempty"`
}

// SubNetworkInterfaceSpec defines the desired state of SubNetworkInterface
type SubNetworkInterfaceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SubNetworkInterfaceParameters `json:"forProvider"`
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
	InitProvider SubNetworkInterfaceInitParameters `json:"initProvider,omitempty"`
}

// SubNetworkInterfaceStatus defines the observed state of SubNetworkInterface.
type SubNetworkInterfaceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SubNetworkInterfaceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// SubNetworkInterface is the Schema for the SubNetworkInterfaces API. ""
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,huaweicloud}
type SubNetworkInterface struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.parentId) || (has(self.initProvider) && has(self.initProvider.parentId))",message="spec.forProvider.parentId is a required parameter"
	Spec   SubNetworkInterfaceSpec   `json:"spec"`
	Status SubNetworkInterfaceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SubNetworkInterfaceList contains a list of SubNetworkInterfaces
type SubNetworkInterfaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SubNetworkInterface `json:"items"`
}

// Repository type metadata.
var (
	SubNetworkInterface_Kind             = "SubNetworkInterface"
	SubNetworkInterface_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SubNetworkInterface_Kind}.String()
	SubNetworkInterface_KindAPIVersion   = SubNetworkInterface_Kind + "." + CRDGroupVersion.String()
	SubNetworkInterface_GroupVersionKind = CRDGroupVersion.WithKind(SubNetworkInterface_Kind)
)

func init() {
	SchemeBuilder.Register(&SubNetworkInterface{}, &SubNetworkInterfaceList{})
}
