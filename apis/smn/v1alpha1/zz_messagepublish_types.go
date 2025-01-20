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

type MessageAttributesInitParameters struct {

	// Specifies the property name.
	// Specifies the property name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies the property type.
	// The value can be STRING, STRING_ARRAY or PROTOCOL.
	// Specifies the property type.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// Specifies the property value.
	// This parameter is valid only when the type set to STRING. The attribute value can only contain Chinese
	// and English, numbers, and underscores, and the length is 1 to 32 characters.
	// Specifies the property value.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`

	// Specifies the property values.
	// This parameter is valid when the type set to STRING_ARRAY or PROTOCOL.
	// Specifies the property values.
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type MessageAttributesObservation struct {

	// Specifies the property name.
	// Specifies the property name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies the property type.
	// The value can be STRING, STRING_ARRAY or PROTOCOL.
	// Specifies the property type.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// Specifies the property value.
	// This parameter is valid only when the type set to STRING. The attribute value can only contain Chinese
	// and English, numbers, and underscores, and the length is 1 to 32 characters.
	// Specifies the property value.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`

	// Specifies the property values.
	// This parameter is valid when the type set to STRING_ARRAY or PROTOCOL.
	// Specifies the property values.
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type MessageAttributesParameters struct {

	// Specifies the property name.
	// Specifies the property name.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// Specifies the property type.
	// The value can be STRING, STRING_ARRAY or PROTOCOL.
	// Specifies the property type.
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`

	// Specifies the property value.
	// This parameter is valid only when the type set to STRING. The attribute value can only contain Chinese
	// and English, numbers, and underscores, and the length is 1 to 32 characters.
	// Specifies the property value.
	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`

	// Specifies the property values.
	// This parameter is valid when the type set to STRING_ARRAY or PROTOCOL.
	// Specifies the property values.
	// +kubebuilder:validation:Optional
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type MessagePublishInitParameters struct {
	EnableForceNew *string `json:"enableForceNew,omitempty" tf:"enable_force_new,omitempty"`

	// Specifies the message content.
	// Specifies the message content.
	Message *string `json:"message,omitempty" tf:"message,omitempty"`

	// Specifies the message filter policies of a subscriber.
	// The message_attributes structure is documented below.
	// Specifies the message filter policies of a subscriber.
	MessageAttributes []MessageAttributesInitParameters `json:"messageAttributes,omitempty" tf:"message_attributes,omitempty"`

	// Specifies the message structure.
	// Specifies the message structure.
	MessageStructure *string `json:"messageStructure,omitempty" tf:"message_structure,omitempty"`

	// Specifies the message template name.
	// Specifies the message template name.
	MessageTemplateName *string `json:"messageTemplateName,omitempty" tf:"message_template_name,omitempty"`

	// Specifies the region in which to create the resource.
	// If omitted, the provider-level region will be used.
	// Changing this creates a new resource.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Specifies the message title.
	// Specifies the message title.
	Subject *string `json:"subject,omitempty" tf:"subject,omitempty"`

	// Specifies a dictionary consisting of tag and parameters to replace the tag.
	// The value corresponding to the label in the message template. Message publishing using message template mode must
	// carry this parameter. The key in the dictionary is the parameter name in the message template, which should not
	// exceed 21 characters. The value in the dictionary is the value after replacing the parameters in the message
	// template, which does not exceed 1KB.
	// Specifies a dictionary consisting of tag and parameters to replace the tag.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Specifies the maximum retention time of the message within the SMN system.
	// After this retention time, the system will no longer send the message. The unit is second, and the default value
	// of the variable is 3600 second. The value is a positive integer and less than or equal to 3600*24.
	// Specifies the maximum retention time of the message within the SMN system.
	TimeToLive *string `json:"timeToLive,omitempty" tf:"time_to_live,omitempty"`

	// Specifies the resource identifier of a topic.
	// Specifies the resource identifier of a topic.
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

type MessagePublishObservation struct {
	EnableForceNew *string `json:"enableForceNew,omitempty" tf:"enable_force_new,omitempty"`

	// The resource ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the message content.
	// Specifies the message content.
	Message *string `json:"message,omitempty" tf:"message,omitempty"`

	// Specifies the message filter policies of a subscriber.
	// The message_attributes structure is documented below.
	// Specifies the message filter policies of a subscriber.
	MessageAttributes []MessageAttributesObservation `json:"messageAttributes,omitempty" tf:"message_attributes,omitempty"`

	// Specifies the message structure.
	// Specifies the message structure.
	MessageStructure *string `json:"messageStructure,omitempty" tf:"message_structure,omitempty"`

	// Specifies the message template name.
	// Specifies the message template name.
	MessageTemplateName *string `json:"messageTemplateName,omitempty" tf:"message_template_name,omitempty"`

	// Specifies the region in which to create the resource.
	// If omitted, the provider-level region will be used.
	// Changing this creates a new resource.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Specifies the message title.
	// Specifies the message title.
	Subject *string `json:"subject,omitempty" tf:"subject,omitempty"`

	// Specifies a dictionary consisting of tag and parameters to replace the tag.
	// The value corresponding to the label in the message template. Message publishing using message template mode must
	// carry this parameter. The key in the dictionary is the parameter name in the message template, which should not
	// exceed 21 characters. The value in the dictionary is the value after replacing the parameters in the message
	// template, which does not exceed 1KB.
	// Specifies a dictionary consisting of tag and parameters to replace the tag.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Specifies the maximum retention time of the message within the SMN system.
	// After this retention time, the system will no longer send the message. The unit is second, and the default value
	// of the variable is 3600 second. The value is a positive integer and less than or equal to 3600*24.
	// Specifies the maximum retention time of the message within the SMN system.
	TimeToLive *string `json:"timeToLive,omitempty" tf:"time_to_live,omitempty"`

	// Specifies the resource identifier of a topic.
	// Specifies the resource identifier of a topic.
	TopicUrn *string `json:"topicUrn,omitempty" tf:"topic_urn,omitempty"`
}

type MessagePublishParameters struct {

	// +kubebuilder:validation:Optional
	EnableForceNew *string `json:"enableForceNew,omitempty" tf:"enable_force_new,omitempty"`

	// Specifies the message content.
	// Specifies the message content.
	// +kubebuilder:validation:Optional
	Message *string `json:"message,omitempty" tf:"message,omitempty"`

	// Specifies the message filter policies of a subscriber.
	// The message_attributes structure is documented below.
	// Specifies the message filter policies of a subscriber.
	// +kubebuilder:validation:Optional
	MessageAttributes []MessageAttributesParameters `json:"messageAttributes,omitempty" tf:"message_attributes,omitempty"`

	// Specifies the message structure.
	// Specifies the message structure.
	// +kubebuilder:validation:Optional
	MessageStructure *string `json:"messageStructure,omitempty" tf:"message_structure,omitempty"`

	// Specifies the message template name.
	// Specifies the message template name.
	// +kubebuilder:validation:Optional
	MessageTemplateName *string `json:"messageTemplateName,omitempty" tf:"message_template_name,omitempty"`

	// Specifies the region in which to create the resource.
	// If omitted, the provider-level region will be used.
	// Changing this creates a new resource.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Specifies the message title.
	// Specifies the message title.
	// +kubebuilder:validation:Optional
	Subject *string `json:"subject,omitempty" tf:"subject,omitempty"`

	// Specifies a dictionary consisting of tag and parameters to replace the tag.
	// The value corresponding to the label in the message template. Message publishing using message template mode must
	// carry this parameter. The key in the dictionary is the parameter name in the message template, which should not
	// exceed 21 characters. The value in the dictionary is the value after replacing the parameters in the message
	// template, which does not exceed 1KB.
	// Specifies a dictionary consisting of tag and parameters to replace the tag.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Specifies the maximum retention time of the message within the SMN system.
	// After this retention time, the system will no longer send the message. The unit is second, and the default value
	// of the variable is 3600 second. The value is a positive integer and less than or equal to 3600*24.
	// Specifies the maximum retention time of the message within the SMN system.
	// +kubebuilder:validation:Optional
	TimeToLive *string `json:"timeToLive,omitempty" tf:"time_to_live,omitempty"`

	// Specifies the resource identifier of a topic.
	// Specifies the resource identifier of a topic.
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

// MessagePublishSpec defines the desired state of MessagePublish
type MessagePublishSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MessagePublishParameters `json:"forProvider"`
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
	InitProvider MessagePublishInitParameters `json:"initProvider,omitempty"`
}

// MessagePublishStatus defines the observed state of MessagePublish.
type MessagePublishStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MessagePublishObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// MessagePublish is the Schema for the MessagePublishs API. Manages a SMN message publishment resource within HuaweiCloud.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,huaweicloud}
type MessagePublish struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MessagePublishSpec   `json:"spec"`
	Status            MessagePublishStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MessagePublishList contains a list of MessagePublishs
type MessagePublishList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MessagePublish `json:"items"`
}

// Repository type metadata.
var (
	MessagePublish_Kind             = "MessagePublish"
	MessagePublish_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: MessagePublish_Kind}.String()
	MessagePublish_KindAPIVersion   = MessagePublish_Kind + "." + CRDGroupVersion.String()
	MessagePublish_GroupVersionKind = CRDGroupVersion.WithKind(MessagePublish_Kind)
)

func init() {
	SchemeBuilder.Register(&MessagePublish{}, &MessagePublishList{})
}
