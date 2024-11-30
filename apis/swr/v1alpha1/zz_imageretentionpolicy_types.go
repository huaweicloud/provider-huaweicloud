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

type ImageRetentionPolicyInitParameters struct {

	// Specifies the number of retention.
	// Specifies the number of retention.
	Number *float64 `json:"number,omitempty" tf:"number,omitempty"`

	// Specifies the name of the organization.
	// Specifies the name of the organization.
	Organization *string `json:"organization,omitempty" tf:"organization,omitempty"`

	// Specifies the region in which to create the resource.
	// If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Specifies the name of the repository.
	// Specifies the name of the repository.
	Repository *string `json:"repository,omitempty" tf:"repository,omitempty"`

	// Specifies the image tags that are not counted in the retention policy
	// The TagSelector structure is documented below.
	// Specifies the image tags that are not counted in the retention policy
	TagSelectors []TagSelectorsInitParameters `json:"tagSelectors,omitempty" tf:"tag_selectors,omitempty"`

	// Specifies the retention policy type.
	// Value options: date_rule, tag_rule.
	// Specifies the retention policy type.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ImageRetentionPolicyObservation struct {

	// The resource ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the number of retention.
	// Specifies the number of retention.
	Number *float64 `json:"number,omitempty" tf:"number,omitempty"`

	// Specifies the name of the organization.
	// Specifies the name of the organization.
	Organization *string `json:"organization,omitempty" tf:"organization,omitempty"`

	// Specifies the region in which to create the resource.
	// If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Specifies the name of the repository.
	// Specifies the name of the repository.
	Repository *string `json:"repository,omitempty" tf:"repository,omitempty"`

	// Specifies the image tags that are not counted in the retention policy
	// The TagSelector structure is documented below.
	// Specifies the image tags that are not counted in the retention policy
	TagSelectors []TagSelectorsObservation `json:"tagSelectors,omitempty" tf:"tag_selectors,omitempty"`

	// Specifies the retention policy type.
	// Value options: date_rule, tag_rule.
	// Specifies the retention policy type.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ImageRetentionPolicyParameters struct {

	// Specifies the number of retention.
	// Specifies the number of retention.
	// +kubebuilder:validation:Optional
	Number *float64 `json:"number,omitempty" tf:"number,omitempty"`

	// Specifies the name of the organization.
	// Specifies the name of the organization.
	// +kubebuilder:validation:Optional
	Organization *string `json:"organization,omitempty" tf:"organization,omitempty"`

	// Specifies the region in which to create the resource.
	// If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Specifies the name of the repository.
	// Specifies the name of the repository.
	// +kubebuilder:validation:Optional
	Repository *string `json:"repository,omitempty" tf:"repository,omitempty"`

	// Specifies the image tags that are not counted in the retention policy
	// The TagSelector structure is documented below.
	// Specifies the image tags that are not counted in the retention policy
	// +kubebuilder:validation:Optional
	TagSelectors []TagSelectorsParameters `json:"tagSelectors,omitempty" tf:"tag_selectors,omitempty"`

	// Specifies the retention policy type.
	// Value options: date_rule, tag_rule.
	// Specifies the retention policy type.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type TagSelectorsInitParameters struct {

	// Specifies the Matching rule. Value options: label, regexp.
	// Specifies the Matching rule.
	Kind *string `json:"kind,omitempty" tf:"kind,omitempty"`

	// Specifies the Matching pattern.
	// Specifies the Matching pattern.
	Pattern *string `json:"pattern,omitempty" tf:"pattern,omitempty"`
}

type TagSelectorsObservation struct {

	// Specifies the Matching rule. Value options: label, regexp.
	// Specifies the Matching rule.
	Kind *string `json:"kind,omitempty" tf:"kind,omitempty"`

	// Specifies the Matching pattern.
	// Specifies the Matching pattern.
	Pattern *string `json:"pattern,omitempty" tf:"pattern,omitempty"`
}

type TagSelectorsParameters struct {

	// Specifies the Matching rule. Value options: label, regexp.
	// Specifies the Matching rule.
	// +kubebuilder:validation:Optional
	Kind *string `json:"kind,omitempty" tf:"kind,omitempty"`

	// Specifies the Matching pattern.
	// Specifies the Matching pattern.
	// +kubebuilder:validation:Optional
	Pattern *string `json:"pattern,omitempty" tf:"pattern,omitempty"`
}

// ImageRetentionPolicySpec defines the desired state of ImageRetentionPolicy
type ImageRetentionPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ImageRetentionPolicyParameters `json:"forProvider"`
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
	InitProvider ImageRetentionPolicyInitParameters `json:"initProvider,omitempty"`
}

// ImageRetentionPolicyStatus defines the observed state of ImageRetentionPolicy.
type ImageRetentionPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ImageRetentionPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ImageRetentionPolicy is the Schema for the ImageRetentionPolicys API. ""
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,huaweicloud}
type ImageRetentionPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.number) || (has(self.initProvider) && has(self.initProvider.number))",message="spec.forProvider.number is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.organization) || (has(self.initProvider) && has(self.initProvider.organization))",message="spec.forProvider.organization is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.repository) || (has(self.initProvider) && has(self.initProvider.repository))",message="spec.forProvider.repository is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.type) || (has(self.initProvider) && has(self.initProvider.type))",message="spec.forProvider.type is a required parameter"
	Spec   ImageRetentionPolicySpec   `json:"spec"`
	Status ImageRetentionPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ImageRetentionPolicyList contains a list of ImageRetentionPolicys
type ImageRetentionPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ImageRetentionPolicy `json:"items"`
}

// Repository type metadata.
var (
	ImageRetentionPolicy_Kind             = "ImageRetentionPolicy"
	ImageRetentionPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ImageRetentionPolicy_Kind}.String()
	ImageRetentionPolicy_KindAPIVersion   = ImageRetentionPolicy_Kind + "." + CRDGroupVersion.String()
	ImageRetentionPolicy_GroupVersionKind = CRDGroupVersion.WithKind(ImageRetentionPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&ImageRetentionPolicy{}, &ImageRetentionPolicyList{})
}
