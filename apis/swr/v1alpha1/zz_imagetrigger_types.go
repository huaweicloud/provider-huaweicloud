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

type ImageTriggerInitParameters struct {

	// Specifies the ID of the cluster.
	// It is required when type is set to cce.
	// Specifies the ID of the cluster.
	// +crossplane:generate:reference:type=github.com/huaweicloud/provider-huaweicloud/apis/cce/v1alpha1.Cluster
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// Reference to a Cluster in cce to populate clusterId.
	// +kubebuilder:validation:Optional
	ClusterIDRef *v1.Reference `json:"clusterIdRef,omitempty" tf:"-"`

	// Selector for a Cluster in cce to populate clusterId.
	// +kubebuilder:validation:Optional
	ClusterIDSelector *v1.Selector `json:"clusterIdSelector,omitempty" tf:"-"`

	// Specifies the name of the cluster.
	// Specifies the name of the cluster.
	ClusterName *string `json:"clusterName,omitempty" tf:"cluster_name,omitempty"`

	// Specifies the trigger condition type.
	// Value options all, tag, regular.
	// Specifies the trigger condition type.
	ConditionType *string `json:"conditionType,omitempty" tf:"condition_type,omitempty"`

	// Specifies the trigger condition value. Value options:
	// Specifies the trigger condition value.
	ConditionValue *string `json:"conditionValue,omitempty" tf:"condition_value,omitempty"`

	// Specifies the name of the container to be updated.
	// By default, all containers using this image are updated.
	// Specifies the name of the container to be updated.
	Container *string `json:"container,omitempty" tf:"container,omitempty"`

	// Specifies whether to enable the trigger.
	// Value options true, false. Default to true
	// Specifies whether to enable the trigger.
	Enabled *string `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Specifies the trigger name.
	// Specifies the trigger name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies the namespace where the application is located.
	// Specifies the namespace where the application is located.
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// Specifies the name of the organization.
	// Specifies the name of the organization.
	// +crossplane:generate:reference:type=github.com/huaweicloud/provider-huaweicloud/apis/swr/v1alpha1.Organization
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("name",true)
	Organization *string `json:"organization,omitempty" tf:"organization,omitempty"`

	// Reference to a Organization in swr to populate organization.
	// +kubebuilder:validation:Optional
	OrganizationRef *v1.Reference `json:"organizationRef,omitempty" tf:"-"`

	// Selector for a Organization in swr to populate organization.
	// +kubebuilder:validation:Optional
	OrganizationSelector *v1.Selector `json:"organizationSelector,omitempty" tf:"-"`

	// Specifies the region in which to create the resource.
	// If omitted, the provider-level region will be used.
	// Changing this parameter will create a new resource.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Specifies the name of the repository.
	// Specifies the name of the repository.
	// +crossplane:generate:reference:type=github.com/huaweicloud/provider-huaweicloud/apis/swr/v1alpha1.Repository
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("name",true)
	Repository *string `json:"repository,omitempty" tf:"repository,omitempty"`

	// Reference to a Repository in swr to populate repository.
	// +kubebuilder:validation:Optional
	RepositoryRef *v1.Reference `json:"repositoryRef,omitempty" tf:"-"`

	// Selector for a Repository in swr to populate repository.
	// +kubebuilder:validation:Optional
	RepositorySelector *v1.Selector `json:"repositorySelector,omitempty" tf:"-"`

	// Specifies the trigger type.
	// Value options cce, cci. Default to cce.
	// Specifies the trigger type.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// Specifies the name of the application.
	// Specifies the name of the application.
	WorkloadName *string `json:"workloadName,omitempty" tf:"workload_name,omitempty"`

	// Specifies the type of the application.
	// Value options: deployments, statefulsets.
	// Specifies the type of the application.
	WorkloadType *string `json:"workloadType,omitempty" tf:"workload_type,omitempty"`
}

type ImageTriggerObservation struct {

	// Specifies the ID of the cluster.
	// It is required when type is set to cce.
	// Specifies the ID of the cluster.
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// Specifies the name of the cluster.
	// Specifies the name of the cluster.
	ClusterName *string `json:"clusterName,omitempty" tf:"cluster_name,omitempty"`

	// Specifies the trigger condition type.
	// Value options all, tag, regular.
	// Specifies the trigger condition type.
	ConditionType *string `json:"conditionType,omitempty" tf:"condition_type,omitempty"`

	// Specifies the trigger condition value. Value options:
	// Specifies the trigger condition value.
	ConditionValue *string `json:"conditionValue,omitempty" tf:"condition_value,omitempty"`

	// Specifies the name of the container to be updated.
	// By default, all containers using this image are updated.
	// Specifies the name of the container to be updated.
	Container *string `json:"container,omitempty" tf:"container,omitempty"`

	// Indicates the creation time.
	// Indicates the creation time.
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// Indicates the creator name of the trigger.
	// Indicates the creator name of the trigger.
	CreatorName *string `json:"creatorName,omitempty" tf:"creator_name,omitempty"`

	// Specifies whether to enable the trigger.
	// Value options true, false. Default to true
	// Specifies whether to enable the trigger.
	Enabled *string `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The resource ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the trigger name.
	// Specifies the trigger name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies the namespace where the application is located.
	// Specifies the namespace where the application is located.
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// Specifies the name of the organization.
	// Specifies the name of the organization.
	Organization *string `json:"organization,omitempty" tf:"organization,omitempty"`

	// Specifies the region in which to create the resource.
	// If omitted, the provider-level region will be used.
	// Changing this parameter will create a new resource.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Specifies the name of the repository.
	// Specifies the name of the repository.
	Repository *string `json:"repository,omitempty" tf:"repository,omitempty"`

	// Specifies the trigger type.
	// Value options cce, cci. Default to cce.
	// Specifies the trigger type.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// Specifies the name of the application.
	// Specifies the name of the application.
	WorkloadName *string `json:"workloadName,omitempty" tf:"workload_name,omitempty"`

	// Specifies the type of the application.
	// Value options: deployments, statefulsets.
	// Specifies the type of the application.
	WorkloadType *string `json:"workloadType,omitempty" tf:"workload_type,omitempty"`
}

type ImageTriggerParameters struct {

	// Specifies the ID of the cluster.
	// It is required when type is set to cce.
	// Specifies the ID of the cluster.
	// +crossplane:generate:reference:type=github.com/huaweicloud/provider-huaweicloud/apis/cce/v1alpha1.Cluster
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// Reference to a Cluster in cce to populate clusterId.
	// +kubebuilder:validation:Optional
	ClusterIDRef *v1.Reference `json:"clusterIdRef,omitempty" tf:"-"`

	// Selector for a Cluster in cce to populate clusterId.
	// +kubebuilder:validation:Optional
	ClusterIDSelector *v1.Selector `json:"clusterIdSelector,omitempty" tf:"-"`

	// Specifies the name of the cluster.
	// Specifies the name of the cluster.
	// +kubebuilder:validation:Optional
	ClusterName *string `json:"clusterName,omitempty" tf:"cluster_name,omitempty"`

	// Specifies the trigger condition type.
	// Value options all, tag, regular.
	// Specifies the trigger condition type.
	// +kubebuilder:validation:Optional
	ConditionType *string `json:"conditionType,omitempty" tf:"condition_type,omitempty"`

	// Specifies the trigger condition value. Value options:
	// Specifies the trigger condition value.
	// +kubebuilder:validation:Optional
	ConditionValue *string `json:"conditionValue,omitempty" tf:"condition_value,omitempty"`

	// Specifies the name of the container to be updated.
	// By default, all containers using this image are updated.
	// Specifies the name of the container to be updated.
	// +kubebuilder:validation:Optional
	Container *string `json:"container,omitempty" tf:"container,omitempty"`

	// Specifies whether to enable the trigger.
	// Value options true, false. Default to true
	// Specifies whether to enable the trigger.
	// +kubebuilder:validation:Optional
	Enabled *string `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Specifies the trigger name.
	// Specifies the trigger name.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies the namespace where the application is located.
	// Specifies the namespace where the application is located.
	// +kubebuilder:validation:Optional
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// Specifies the name of the organization.
	// Specifies the name of the organization.
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

	// Specifies the region in which to create the resource.
	// If omitted, the provider-level region will be used.
	// Changing this parameter will create a new resource.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Specifies the name of the repository.
	// Specifies the name of the repository.
	// +crossplane:generate:reference:type=github.com/huaweicloud/provider-huaweicloud/apis/swr/v1alpha1.Repository
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("name",true)
	// +kubebuilder:validation:Optional
	Repository *string `json:"repository,omitempty" tf:"repository,omitempty"`

	// Reference to a Repository in swr to populate repository.
	// +kubebuilder:validation:Optional
	RepositoryRef *v1.Reference `json:"repositoryRef,omitempty" tf:"-"`

	// Selector for a Repository in swr to populate repository.
	// +kubebuilder:validation:Optional
	RepositorySelector *v1.Selector `json:"repositorySelector,omitempty" tf:"-"`

	// Specifies the trigger type.
	// Value options cce, cci. Default to cce.
	// Specifies the trigger type.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// Specifies the name of the application.
	// Specifies the name of the application.
	// +kubebuilder:validation:Optional
	WorkloadName *string `json:"workloadName,omitempty" tf:"workload_name,omitempty"`

	// Specifies the type of the application.
	// Value options: deployments, statefulsets.
	// Specifies the type of the application.
	// +kubebuilder:validation:Optional
	WorkloadType *string `json:"workloadType,omitempty" tf:"workload_type,omitempty"`
}

// ImageTriggerSpec defines the desired state of ImageTrigger
type ImageTriggerSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ImageTriggerParameters `json:"forProvider"`
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
	InitProvider ImageTriggerInitParameters `json:"initProvider,omitempty"`
}

// ImageTriggerStatus defines the observed state of ImageTrigger.
type ImageTriggerStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ImageTriggerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ImageTrigger is the Schema for the ImageTriggers API. ""
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,huaweicloud}
type ImageTrigger struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.conditionType) || (has(self.initProvider) && has(self.initProvider.conditionType))",message="spec.forProvider.conditionType is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.conditionValue) || (has(self.initProvider) && has(self.initProvider.conditionValue))",message="spec.forProvider.conditionValue is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.__namespace__) || (has(self.initProvider) && has(self.initProvider.__namespace__))",message="spec.forProvider.namespace is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.workloadName) || (has(self.initProvider) && has(self.initProvider.workloadName))",message="spec.forProvider.workloadName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.workloadType) || (has(self.initProvider) && has(self.initProvider.workloadType))",message="spec.forProvider.workloadType is a required parameter"
	Spec   ImageTriggerSpec   `json:"spec"`
	Status ImageTriggerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ImageTriggerList contains a list of ImageTriggers
type ImageTriggerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ImageTrigger `json:"items"`
}

// Repository type metadata.
var (
	ImageTrigger_Kind             = "ImageTrigger"
	ImageTrigger_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ImageTrigger_Kind}.String()
	ImageTrigger_KindAPIVersion   = ImageTrigger_Kind + "." + CRDGroupVersion.String()
	ImageTrigger_GroupVersionKind = CRDGroupVersion.WithKind(ImageTrigger_Kind)
)

func init() {
	SchemeBuilder.Register(&ImageTrigger{}, &ImageTriggerList{})
}
