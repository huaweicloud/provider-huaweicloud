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

type GlobalInternetBandwidthInitParameters struct {

	// Specifies the access site name.
	// Changing this creates a new resource.
	AccessSite *string `json:"accessSite,omitempty" tf:"access_site,omitempty"`

	// Specifies the charge mode, for example, 95peak_guar.
	ChargeMode *string `json:"chargeMode,omitempty" tf:"charge_mode,omitempty"`

	// Specifies the description of the global internet bandwidth.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies the enterprise project id to which the global
	// internet bandwidth belongs. Changing this creates a new resource.
	EnterpriseProjectID *string `json:"enterpriseProjectId,omitempty" tf:"enterprise_project_id,omitempty"`

	// Specifies the ingress size of the global internet bandwidth.
	// It's not used for charge mode 95peak_guar.
	IngressSize *float64 `json:"ingressSize,omitempty" tf:"ingress_size,omitempty"`

	// Specifies the internet service provider of the global internet bandwidth.
	// Changing this creates a new resource.
	Isp *string `json:"isp,omitempty" tf:"isp,omitempty"`

	// Specifies the name of the global internet bandwidth.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies the size of the global internet bandwidth.
	// The value ranges from 300 Mbit/s to 5000 Mbit/s in normal.
	Size *float64 `json:"size,omitempty" tf:"size,omitempty"`

	// Specifies the tags of the global internet bandwidth.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Specifies the type of the global internet bandwidth.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type GlobalInternetBandwidthObservation struct {

	// Specifies the access site name.
	// Changing this creates a new resource.
	AccessSite *string `json:"accessSite,omitempty" tf:"access_site,omitempty"`

	// Specifies the charge mode, for example, 95peak_guar.
	ChargeMode *string `json:"chargeMode,omitempty" tf:"charge_mode,omitempty"`

	// The create time of the global internet bandwidth.
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// Specifies the description of the global internet bandwidth.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies the enterprise project id to which the global
	// internet bandwidth belongs. Changing this creates a new resource.
	EnterpriseProjectID *string `json:"enterpriseProjectId,omitempty" tf:"enterprise_project_id,omitempty"`

	// The frozen info of the global internet bandwidth.
	FrozenInfo *string `json:"frozenInfo,omitempty" tf:"frozen_info,omitempty"`

	// The resource ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the ingress size of the global internet bandwidth.
	// It's not used for charge mode 95peak_guar.
	IngressSize *float64 `json:"ingressSize,omitempty" tf:"ingress_size,omitempty"`

	// Specifies the internet service provider of the global internet bandwidth.
	// Changing this creates a new resource.
	Isp *string `json:"isp,omitempty" tf:"isp,omitempty"`

	// Specifies the name of the global internet bandwidth.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The enhanced 95% guaranteed rate of the global internet bandwidth.
	Ratio95Peak *float64 `json:"ratio95Peak,omitempty" tf:"ratio_95peak,omitempty"`

	// Specifies the size of the global internet bandwidth.
	// The value ranges from 300 Mbit/s to 5000 Mbit/s in normal.
	Size *float64 `json:"size,omitempty" tf:"size,omitempty"`

	// The status of the global internet bandwidth.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// Specifies the tags of the global internet bandwidth.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Specifies the type of the global internet bandwidth.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// The update time of the global internet bandwidth.
	UpdatedAt *string `json:"updatedAt,omitempty" tf:"updated_at,omitempty"`
}

type GlobalInternetBandwidthParameters struct {

	// Specifies the access site name.
	// Changing this creates a new resource.
	// +kubebuilder:validation:Optional
	AccessSite *string `json:"accessSite,omitempty" tf:"access_site,omitempty"`

	// Specifies the charge mode, for example, 95peak_guar.
	// +kubebuilder:validation:Optional
	ChargeMode *string `json:"chargeMode,omitempty" tf:"charge_mode,omitempty"`

	// Specifies the description of the global internet bandwidth.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies the enterprise project id to which the global
	// internet bandwidth belongs. Changing this creates a new resource.
	// +kubebuilder:validation:Optional
	EnterpriseProjectID *string `json:"enterpriseProjectId,omitempty" tf:"enterprise_project_id,omitempty"`

	// Specifies the ingress size of the global internet bandwidth.
	// It's not used for charge mode 95peak_guar.
	// +kubebuilder:validation:Optional
	IngressSize *float64 `json:"ingressSize,omitempty" tf:"ingress_size,omitempty"`

	// Specifies the internet service provider of the global internet bandwidth.
	// Changing this creates a new resource.
	// +kubebuilder:validation:Optional
	Isp *string `json:"isp,omitempty" tf:"isp,omitempty"`

	// Specifies the name of the global internet bandwidth.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies the size of the global internet bandwidth.
	// The value ranges from 300 Mbit/s to 5000 Mbit/s in normal.
	// +kubebuilder:validation:Optional
	Size *float64 `json:"size,omitempty" tf:"size,omitempty"`

	// Specifies the tags of the global internet bandwidth.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Specifies the type of the global internet bandwidth.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

// GlobalInternetBandwidthSpec defines the desired state of GlobalInternetBandwidth
type GlobalInternetBandwidthSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     GlobalInternetBandwidthParameters `json:"forProvider"`
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
	InitProvider GlobalInternetBandwidthInitParameters `json:"initProvider,omitempty"`
}

// GlobalInternetBandwidthStatus defines the observed state of GlobalInternetBandwidth.
type GlobalInternetBandwidthStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        GlobalInternetBandwidthObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// GlobalInternetBandwidth is the Schema for the GlobalInternetBandwidths API. ""
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,huaweicloud}
type GlobalInternetBandwidth struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.accessSite) || (has(self.initProvider) && has(self.initProvider.accessSite))",message="spec.forProvider.accessSite is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.chargeMode) || (has(self.initProvider) && has(self.initProvider.chargeMode))",message="spec.forProvider.chargeMode is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.isp) || (has(self.initProvider) && has(self.initProvider.isp))",message="spec.forProvider.isp is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.size) || (has(self.initProvider) && has(self.initProvider.size))",message="spec.forProvider.size is a required parameter"
	Spec   GlobalInternetBandwidthSpec   `json:"spec"`
	Status GlobalInternetBandwidthStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GlobalInternetBandwidthList contains a list of GlobalInternetBandwidths
type GlobalInternetBandwidthList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GlobalInternetBandwidth `json:"items"`
}

// Repository type metadata.
var (
	GlobalInternetBandwidth_Kind             = "GlobalInternetBandwidth"
	GlobalInternetBandwidth_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: GlobalInternetBandwidth_Kind}.String()
	GlobalInternetBandwidth_KindAPIVersion   = GlobalInternetBandwidth_Kind + "." + CRDGroupVersion.String()
	GlobalInternetBandwidth_GroupVersionKind = CRDGroupVersion.WithKind(GlobalInternetBandwidth_Kind)
)

func init() {
	SchemeBuilder.Register(&GlobalInternetBandwidth{}, &GlobalInternetBandwidthList{})
}
