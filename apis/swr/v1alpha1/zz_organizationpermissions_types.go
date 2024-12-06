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

type OrganizationPermissionsInitParameters struct {

	// Specifies the name of the organization (namespace) to be accessed.
	// Changing this creates a new resource.
	// +crossplane:generate:reference:type=github.com/huaweicloud/provider-huaweicloud/apis/swr/v1alpha1.Organization
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("name",true)
	Organization *string `json:"organization,omitempty" tf:"organization,omitempty"`

	// Reference to a Organization in swr to populate organization.
	// +kubebuilder:validation:Optional
	OrganizationRef *v1.Reference `json:"organizationRef,omitempty" tf:"-"`

	// Selector for a Organization in swr to populate organization.
	// +kubebuilder:validation:Optional
	OrganizationSelector *v1.Selector `json:"organizationSelector,omitempty" tf:"-"`

	// Specifies the region in which to create the resource. If omitted, the
	// provider-level region will be used. Changing this creates a new resource.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Specifies the users to access to the organization (namespace).
	// Structure is documented below.
	Users []OrganizationPermissionsUsersInitParameters `json:"users,omitempty" tf:"users,omitempty"`
}

type OrganizationPermissionsObservation struct {

	// The creator user name of the organization.
	Creator *string `json:"creator,omitempty" tf:"creator,omitempty"`

	// ID of the permissions. The value is the name of the organization.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the name of the organization (namespace) to be accessed.
	// Changing this creates a new resource.
	Organization *string `json:"organization,omitempty" tf:"organization,omitempty"`

	// Specifies the region in which to create the resource. If omitted, the
	// provider-level region will be used. Changing this creates a new resource.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The permission information of current user.
	SelfPermission []OrganizationPermissionsSelfPermissionObservation `json:"selfPermission,omitempty" tf:"self_permission,omitempty"`

	// Specifies the users to access to the organization (namespace).
	// Structure is documented below.
	Users []OrganizationPermissionsUsersObservation `json:"users,omitempty" tf:"users,omitempty"`
}

type OrganizationPermissionsParameters struct {

	// Specifies the name of the organization (namespace) to be accessed.
	// Changing this creates a new resource.
	// +crossplane:generate:reference:type=github.com/huaweicloud/provider-huaweicloud/apis/swr/v1alpha1.Organization
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("name",true)
	// +kubebuilder:validation:Optional
	Organization *string `json:"organization,omitempty" tf:"organization,omitempty"`

	// Reference to a Organization in swr to populate organization.
	// +kubebuilder:validation:Optional
	OrganizationRef *v1.Reference `json:"organizationRef,omitempty" tf:"-"`

	// Selector for a Organization in swr to populate organization.
	// +kubebuilder:validation:Optional
	OrganizationSelector *v1.Selector `json:"organizationSelector,omitempty" tf:"-"`

	// Specifies the region in which to create the resource. If omitted, the
	// provider-level region will be used. Changing this creates a new resource.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Specifies the users to access to the organization (namespace).
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	Users []OrganizationPermissionsUsersParameters `json:"users,omitempty" tf:"users,omitempty"`
}

type OrganizationPermissionsSelfPermissionInitParameters struct {
}

type OrganizationPermissionsSelfPermissionObservation struct {

	// The permission of current user.
	Permission *string `json:"permission,omitempty" tf:"permission,omitempty"`

	// The ID of current user.
	UserID *string `json:"userId,omitempty" tf:"user_id,omitempty"`

	// The name of current user.
	UserName *string `json:"userName,omitempty" tf:"user_name,omitempty"`
}

type OrganizationPermissionsSelfPermissionParameters struct {
}

type OrganizationPermissionsUsersInitParameters struct {

	// Specifies the permission of the existing HuaweiCloud user.
	// The values can be Manage, Write and Read.
	Permission *string `json:"permission,omitempty" tf:"permission,omitempty"`

	// Specifies the ID of the existing HuaweiCloud user.
	// +crossplane:generate:reference:type=github.com/huaweicloud/provider-huaweicloud/apis/iam/v1alpha1.User
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	UserID *string `json:"userId,omitempty" tf:"user_id,omitempty"`

	// Reference to a User in iam to populate userId.
	// +kubebuilder:validation:Optional
	UserIDRef *v1.Reference `json:"userIdRef,omitempty" tf:"-"`

	// Selector for a User in iam to populate userId.
	// +kubebuilder:validation:Optional
	UserIDSelector *v1.Selector `json:"userIdSelector,omitempty" tf:"-"`

	// Specifies the name of the existing HuaweiCloud user.
	UserName *string `json:"userName,omitempty" tf:"user_name,omitempty"`
}

type OrganizationPermissionsUsersObservation struct {

	// Specifies the permission of the existing HuaweiCloud user.
	// The values can be Manage, Write and Read.
	Permission *string `json:"permission,omitempty" tf:"permission,omitempty"`

	// Specifies the ID of the existing HuaweiCloud user.
	UserID *string `json:"userId,omitempty" tf:"user_id,omitempty"`

	// Specifies the name of the existing HuaweiCloud user.
	UserName *string `json:"userName,omitempty" tf:"user_name,omitempty"`
}

type OrganizationPermissionsUsersParameters struct {

	// Specifies the permission of the existing HuaweiCloud user.
	// The values can be Manage, Write and Read.
	// +kubebuilder:validation:Optional
	Permission *string `json:"permission" tf:"permission,omitempty"`

	// Specifies the ID of the existing HuaweiCloud user.
	// +crossplane:generate:reference:type=github.com/huaweicloud/provider-huaweicloud/apis/iam/v1alpha1.User
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	UserID *string `json:"userId,omitempty" tf:"user_id,omitempty"`

	// Reference to a User in iam to populate userId.
	// +kubebuilder:validation:Optional
	UserIDRef *v1.Reference `json:"userIdRef,omitempty" tf:"-"`

	// Selector for a User in iam to populate userId.
	// +kubebuilder:validation:Optional
	UserIDSelector *v1.Selector `json:"userIdSelector,omitempty" tf:"-"`

	// Specifies the name of the existing HuaweiCloud user.
	// +kubebuilder:validation:Optional
	UserName *string `json:"userName,omitempty" tf:"user_name,omitempty"`
}

// OrganizationPermissionsSpec defines the desired state of OrganizationPermissions
type OrganizationPermissionsSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     OrganizationPermissionsParameters `json:"forProvider"`
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
	InitProvider OrganizationPermissionsInitParameters `json:"initProvider,omitempty"`
}

// OrganizationPermissionsStatus defines the observed state of OrganizationPermissions.
type OrganizationPermissionsStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        OrganizationPermissionsObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// OrganizationPermissions is the Schema for the OrganizationPermissionss API. ""
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,huaweicloud}
type OrganizationPermissions struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.users) || (has(self.initProvider) && has(self.initProvider.users))",message="spec.forProvider.users is a required parameter"
	Spec   OrganizationPermissionsSpec   `json:"spec"`
	Status OrganizationPermissionsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OrganizationPermissionsList contains a list of OrganizationPermissionss
type OrganizationPermissionsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OrganizationPermissions `json:"items"`
}

// Repository type metadata.
var (
	OrganizationPermissions_Kind             = "OrganizationPermissions"
	OrganizationPermissions_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: OrganizationPermissions_Kind}.String()
	OrganizationPermissions_KindAPIVersion   = OrganizationPermissions_Kind + "." + CRDGroupVersion.String()
	OrganizationPermissions_GroupVersionKind = CRDGroupVersion.WithKind(OrganizationPermissions_Kind)
)

func init() {
	SchemeBuilder.Register(&OrganizationPermissions{}, &OrganizationPermissionsList{})
}
