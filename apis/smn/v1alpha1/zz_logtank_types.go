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

type LogtankInitParameters struct {

	// The lts log group ID.
	LogGroupID *string `json:"logGroupId,omitempty" tf:"log_group_id,omitempty"`

	// The lts log stream ID.
	LogStreamID *string `json:"logStreamId,omitempty" tf:"log_stream_id,omitempty"`

	// The region in which to create the SMN logtank resource. If omitted, the
	// provider-level region will be used. Changing this parameter will create a new resource.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Resource identifier of a topic, which is unique.
	// Changing this parameter will create a new resource.
	// +crossplane:generate:reference:type=github.com/huaweicloud/provider-huaweicloud/apis/smn/v1alpha1.Topic
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	TopicUrn *string `json:"topicUrn,omitempty" tf:"topic_urn,omitempty"`

	// Reference to a Topic in smn to populate topicUrn.
	// +kubebuilder:validation:Optional
	TopicUrnRef *v1.Reference `json:"topicUrnRef,omitempty" tf:"-"`

	// Selector for a Topic in smn to populate topicUrn.
	// +kubebuilder:validation:Optional
	TopicUrnSelector *v1.Selector `json:"topicUrnSelector,omitempty" tf:"-"`
}

type LogtankObservation struct {

	// Time when the logtank was created.
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// The resource ID. The value is the topic URN.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The lts log group ID.
	LogGroupID *string `json:"logGroupId,omitempty" tf:"log_group_id,omitempty"`

	// The lts log stream ID.
	LogStreamID *string `json:"logStreamId,omitempty" tf:"log_stream_id,omitempty"`

	// The ID of the logtank.
	LogtankID *string `json:"logtankId,omitempty" tf:"logtank_id,omitempty"`

	// The region in which to create the SMN logtank resource. If omitted, the
	// provider-level region will be used. Changing this parameter will create a new resource.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Resource identifier of a topic, which is unique.
	// Changing this parameter will create a new resource.
	TopicUrn *string `json:"topicUrn,omitempty" tf:"topic_urn,omitempty"`

	// Time when the logtank was updated.
	UpdatedAt *string `json:"updatedAt,omitempty" tf:"updated_at,omitempty"`
}

type LogtankParameters struct {

	// The lts log group ID.
	// +kubebuilder:validation:Optional
	LogGroupID *string `json:"logGroupId,omitempty" tf:"log_group_id,omitempty"`

	// The lts log stream ID.
	// +kubebuilder:validation:Optional
	LogStreamID *string `json:"logStreamId,omitempty" tf:"log_stream_id,omitempty"`

	// The region in which to create the SMN logtank resource. If omitted, the
	// provider-level region will be used. Changing this parameter will create a new resource.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Resource identifier of a topic, which is unique.
	// Changing this parameter will create a new resource.
	// +crossplane:generate:reference:type=github.com/huaweicloud/provider-huaweicloud/apis/smn/v1alpha1.Topic
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	TopicUrn *string `json:"topicUrn,omitempty" tf:"topic_urn,omitempty"`

	// Reference to a Topic in smn to populate topicUrn.
	// +kubebuilder:validation:Optional
	TopicUrnRef *v1.Reference `json:"topicUrnRef,omitempty" tf:"-"`

	// Selector for a Topic in smn to populate topicUrn.
	// +kubebuilder:validation:Optional
	TopicUrnSelector *v1.Selector `json:"topicUrnSelector,omitempty" tf:"-"`
}

// LogtankSpec defines the desired state of Logtank
type LogtankSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LogtankParameters `json:"forProvider"`
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
	InitProvider LogtankInitParameters `json:"initProvider,omitempty"`
}

// LogtankStatus defines the observed state of Logtank.
type LogtankStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LogtankObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Logtank is the Schema for the Logtanks API. ""
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,huaweicloud}
type Logtank struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.logGroupId) || (has(self.initProvider) && has(self.initProvider.logGroupId))",message="spec.forProvider.logGroupId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.logStreamId) || (has(self.initProvider) && has(self.initProvider.logStreamId))",message="spec.forProvider.logStreamId is a required parameter"
	Spec   LogtankSpec   `json:"spec"`
	Status LogtankStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LogtankList contains a list of Logtanks
type LogtankList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Logtank `json:"items"`
}

// Repository type metadata.
var (
	Logtank_Kind             = "Logtank"
	Logtank_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Logtank_Kind}.String()
	Logtank_KindAPIVersion   = Logtank_Kind + "." + CRDGroupVersion.String()
	Logtank_GroupVersionKind = CRDGroupVersion.WithKind(Logtank_Kind)
)

func init() {
	SchemeBuilder.Register(&Logtank{}, &LogtankList{})
}
