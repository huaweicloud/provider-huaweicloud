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

type AgencyInitParameters struct {

	// Specifies an array of one or more role names which stand for the permissions
	// to be granted to agency on all resources, including those in enterprise projects, region-specific projects,
	// and global services under your account.
	// +listType=set
	AllResourcesRoles []*string `json:"allResourcesRoles,omitempty" tf:"all_resources_roles,omitempty"`

	// Specifies the name of delegated user domain.
	// schema: Required
	DelegatedDomainName *string `json:"delegatedDomainName,omitempty" tf:"delegated_domain_name,omitempty"`

	// Specifies the name of agency. The name is a string of 1 to 64 characters.
	// Changing this will create a new agency.
	// schema: Internal
	DelegatedServiceName *string `json:"delegatedServiceName,omitempty" tf:"delegated_service_name,omitempty"`

	// Specifies the supplementary information about the agency. The value is a string of
	// 0 to 255 characters, excluding these characters: '@#$%^&*<>\'.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies an array of one or more role names which stand for the permissions to be
	// granted to agency on domain.
	// +listType=set
	DomainRoles []*string `json:"domainRoles,omitempty" tf:"domain_roles,omitempty"`

	// Specifies the validity period of an agency. The valid value are FOREVER, ONEDAY
	// or the specific days, for example, "20". The default value is FOREVER.
	Duration *string `json:"duration,omitempty" tf:"duration,omitempty"`

	// Specifies the name of agency. The name is a string of 1 to 64 characters.
	// Changing this will create a new agency.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies an array of one or more roles and projects which are used to grant
	// permissions to agency on project. The structure is documented below.
	ProjectRole []ProjectRoleInitParameters `json:"projectRole,omitempty" tf:"project_role,omitempty"`
}

type AgencyObservation struct {

	// Specifies an array of one or more role names which stand for the permissions
	// to be granted to agency on all resources, including those in enterprise projects, region-specific projects,
	// and global services under your account.
	// +listType=set
	AllResourcesRoles []*string `json:"allResourcesRoles,omitempty" tf:"all_resources_roles,omitempty"`

	// The time when the agency was created.
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	// Specifies the name of delegated user domain.
	// schema: Required
	DelegatedDomainName *string `json:"delegatedDomainName,omitempty" tf:"delegated_domain_name,omitempty"`

	// Specifies the name of agency. The name is a string of 1 to 64 characters.
	// Changing this will create a new agency.
	// schema: Internal
	DelegatedServiceName *string `json:"delegatedServiceName,omitempty" tf:"delegated_service_name,omitempty"`

	// Specifies the supplementary information about the agency. The value is a string of
	// 0 to 255 characters, excluding these characters: '@#$%^&*<>\'.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies an array of one or more role names which stand for the permissions to be
	// granted to agency on domain.
	// +listType=set
	DomainRoles []*string `json:"domainRoles,omitempty" tf:"domain_roles,omitempty"`

	// Specifies the validity period of an agency. The valid value are FOREVER, ONEDAY
	// or the specific days, for example, "20". The default value is FOREVER.
	Duration *string `json:"duration,omitempty" tf:"duration,omitempty"`

	// The expiration time of agency.
	ExpireTime *string `json:"expireTime,omitempty" tf:"expire_time,omitempty"`

	// The agency ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the name of agency. The name is a string of 1 to 64 characters.
	// Changing this will create a new agency.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies an array of one or more roles and projects which are used to grant
	// permissions to agency on project. The structure is documented below.
	ProjectRole []ProjectRoleObservation `json:"projectRole,omitempty" tf:"project_role,omitempty"`
}

type AgencyParameters struct {

	// Specifies an array of one or more role names which stand for the permissions
	// to be granted to agency on all resources, including those in enterprise projects, region-specific projects,
	// and global services under your account.
	// +kubebuilder:validation:Optional
	// +listType=set
	AllResourcesRoles []*string `json:"allResourcesRoles,omitempty" tf:"all_resources_roles,omitempty"`

	// Specifies the name of delegated user domain.
	// schema: Required
	// +kubebuilder:validation:Optional
	DelegatedDomainName *string `json:"delegatedDomainName,omitempty" tf:"delegated_domain_name,omitempty"`

	// Specifies the name of agency. The name is a string of 1 to 64 characters.
	// Changing this will create a new agency.
	// schema: Internal
	// +kubebuilder:validation:Optional
	DelegatedServiceName *string `json:"delegatedServiceName,omitempty" tf:"delegated_service_name,omitempty"`

	// Specifies the supplementary information about the agency. The value is a string of
	// 0 to 255 characters, excluding these characters: '@#$%^&*<>\'.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies an array of one or more role names which stand for the permissions to be
	// granted to agency on domain.
	// +kubebuilder:validation:Optional
	// +listType=set
	DomainRoles []*string `json:"domainRoles,omitempty" tf:"domain_roles,omitempty"`

	// Specifies the validity period of an agency. The valid value are FOREVER, ONEDAY
	// or the specific days, for example, "20". The default value is FOREVER.
	// +kubebuilder:validation:Optional
	Duration *string `json:"duration,omitempty" tf:"duration,omitempty"`

	// Specifies the name of agency. The name is a string of 1 to 64 characters.
	// Changing this will create a new agency.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies an array of one or more roles and projects which are used to grant
	// permissions to agency on project. The structure is documented below.
	// +kubebuilder:validation:Optional
	ProjectRole []ProjectRoleParameters `json:"projectRole,omitempty" tf:"project_role,omitempty"`
}

type ProjectRoleInitParameters struct {

	// Specifies the name of project.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Specifies an array of role names.
	// +listType=set
	Roles []*string `json:"roles,omitempty" tf:"roles,omitempty"`
}

type ProjectRoleObservation struct {

	// Specifies the name of project.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Specifies an array of role names.
	// +listType=set
	Roles []*string `json:"roles,omitempty" tf:"roles,omitempty"`
}

type ProjectRoleParameters struct {

	// Specifies the name of project.
	// +kubebuilder:validation:Optional
	Project *string `json:"project" tf:"project,omitempty"`

	// Specifies an array of role names.
	// +kubebuilder:validation:Optional
	// +listType=set
	Roles []*string `json:"roles" tf:"roles,omitempty"`
}

// AgencySpec defines the desired state of Agency
type AgencySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AgencyParameters `json:"forProvider"`
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
	InitProvider AgencyInitParameters `json:"initProvider,omitempty"`
}

// AgencyStatus defines the observed state of Agency.
type AgencyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AgencyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Agency is the Schema for the Agencys API. ""
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,huaweicloud}
type Agency struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   AgencySpec   `json:"spec"`
	Status AgencyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AgencyList contains a list of Agencys
type AgencyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Agency `json:"items"`
}

// Repository type metadata.
var (
	Agency_Kind             = "Agency"
	Agency_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Agency_Kind}.String()
	Agency_KindAPIVersion   = Agency_Kind + "." + CRDGroupVersion.String()
	Agency_GroupVersionKind = CRDGroupVersion.WithKind(Agency_Kind)
)

func init() {
	SchemeBuilder.Register(&Agency{}, &AgencyList{})
}
