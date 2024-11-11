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

type GroupPermsInitParameters struct {

	// Indicates the name of a topic or consumer group.
	// Indicates the name of a topic or consumer group.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Indicates the permissions of the topic or consumer group.
	// Value options: PUB|SUB, PUB, SUB, DENY.
	// Indicates the permissions of the topic or consumer group.
	// Value options: **PUB|SUB**, **PUB**, **SUB**, **DENY**.
	Perm *string `json:"perm,omitempty" tf:"perm,omitempty"`
}

type GroupPermsObservation struct {

	// Indicates the name of a topic or consumer group.
	// Indicates the name of a topic or consumer group.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Indicates the permissions of the topic or consumer group.
	// Value options: PUB|SUB, PUB, SUB, DENY.
	// Indicates the permissions of the topic or consumer group.
	// Value options: **PUB|SUB**, **PUB**, **SUB**, **DENY**.
	Perm *string `json:"perm,omitempty" tf:"perm,omitempty"`
}

type GroupPermsParameters struct {

	// Indicates the name of a topic or consumer group.
	// Indicates the name of a topic or consumer group.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Indicates the permissions of the topic or consumer group.
	// Value options: PUB|SUB, PUB, SUB, DENY.
	// Indicates the permissions of the topic or consumer group.
	// Value options: **PUB|SUB**, **PUB**, **SUB**, **DENY**.
	// +kubebuilder:validation:Optional
	Perm *string `json:"perm,omitempty" tf:"perm,omitempty"`
}

type RocketmqUserInitParameters struct {

	// Specifies the name of the user, which starts with a letter, consists of 7
	// to 64 characters and can contain only letters, digits, hyphens (-), and underscores (_).
	// Changing this parameter will create a new resource.
	// Specifies the access key of the user.
	AccessKey *string `json:"accessKey,omitempty" tf:"access_key,omitempty"`

	// Specifies whether the user is an administrator.
	// Specifies whether the user is an administrator.
	Admin *bool `json:"admin,omitempty" tf:"admin,omitempty"`

	// Specifies the default consumer group permissions.
	// Value options: PUB|SUB, PUB, SUB, DENY.
	// Specifies the default consumer group permissions.
	// Value options: **PUB|SUB**, **PUB**, **SUB**, **DENY**.
	DefaultGroupPerm *string `json:"defaultGroupPerm,omitempty" tf:"default_group_perm,omitempty"`

	// Specifies the default topic permissions.
	// Value options: PUB|SUB, PUB, SUB, DENY.
	// Specifies the default topic permissions.
	// Value options: **PUB|SUB**, **PUB**, **SUB**, **DENY**.
	DefaultTopicPerm *string `json:"defaultTopicPerm,omitempty" tf:"default_topic_perm,omitempty"`

	// Specifies the special consumer group permissions.
	// The permission structure is documented below.
	// Specifies the special consumer group permissions.
	GroupPerms []GroupPermsInitParameters `json:"groupPerms,omitempty" tf:"group_perms,omitempty"`

	// Specifies the ID of the rocketMQ instance.
	// Changing this parameter will create a new resource.
	// Specifies the ID of the rocketMQ instance.
	// +crossplane:generate:reference:type=github.com/huaweicloud/provider-huaweicloud/apis/dms/v1alpha1.RocketmqInstance
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// Reference to a RocketmqInstance in dms to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDRef *v1.Reference `json:"instanceIdRef,omitempty" tf:"-"`

	// Selector for a RocketmqInstance in dms to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDSelector *v1.Selector `json:"instanceIdSelector,omitempty" tf:"-"`

	// Specifies the region in which to create the resource.
	// If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Specifies the password of the user. Use 8 to 32 characters. Contain at
	// least three of the following character types:
	// Specifies the secret key of the user.
	SecretKey *string `json:"secretKey,omitempty" tf:"secret_key,omitempty"`

	// Specifies the special topic permissions.
	// The permission structure is documented below.
	// Specifies the special topic permissions.
	TopicPerms []TopicPermsInitParameters `json:"topicPerms,omitempty" tf:"topic_perms,omitempty"`

	// Specifies the IP address whitelist.
	// Specifies the IP address whitelist.
	WhiteRemoteAddress *string `json:"whiteRemoteAddress,omitempty" tf:"white_remote_address,omitempty"`
}

type RocketmqUserObservation struct {

	// Specifies the name of the user, which starts with a letter, consists of 7
	// to 64 characters and can contain only letters, digits, hyphens (-), and underscores (_).
	// Changing this parameter will create a new resource.
	// Specifies the access key of the user.
	AccessKey *string `json:"accessKey,omitempty" tf:"access_key,omitempty"`

	// Specifies whether the user is an administrator.
	// Specifies whether the user is an administrator.
	Admin *bool `json:"admin,omitempty" tf:"admin,omitempty"`

	// Specifies the default consumer group permissions.
	// Value options: PUB|SUB, PUB, SUB, DENY.
	// Specifies the default consumer group permissions.
	// Value options: **PUB|SUB**, **PUB**, **SUB**, **DENY**.
	DefaultGroupPerm *string `json:"defaultGroupPerm,omitempty" tf:"default_group_perm,omitempty"`

	// Specifies the default topic permissions.
	// Value options: PUB|SUB, PUB, SUB, DENY.
	// Specifies the default topic permissions.
	// Value options: **PUB|SUB**, **PUB**, **SUB**, **DENY**.
	DefaultTopicPerm *string `json:"defaultTopicPerm,omitempty" tf:"default_topic_perm,omitempty"`

	// Specifies the special consumer group permissions.
	// The permission structure is documented below.
	// Specifies the special consumer group permissions.
	GroupPerms []GroupPermsObservation `json:"groupPerms,omitempty" tf:"group_perms,omitempty"`

	// The resource ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the ID of the rocketMQ instance.
	// Changing this parameter will create a new resource.
	// Specifies the ID of the rocketMQ instance.
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// Specifies the region in which to create the resource.
	// If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Specifies the password of the user. Use 8 to 32 characters. Contain at
	// least three of the following character types:
	// Specifies the secret key of the user.
	SecretKey *string `json:"secretKey,omitempty" tf:"secret_key,omitempty"`

	// Specifies the special topic permissions.
	// The permission structure is documented below.
	// Specifies the special topic permissions.
	TopicPerms []TopicPermsObservation `json:"topicPerms,omitempty" tf:"topic_perms,omitempty"`

	// Specifies the IP address whitelist.
	// Specifies the IP address whitelist.
	WhiteRemoteAddress *string `json:"whiteRemoteAddress,omitempty" tf:"white_remote_address,omitempty"`
}

type RocketmqUserParameters struct {

	// Specifies the name of the user, which starts with a letter, consists of 7
	// to 64 characters and can contain only letters, digits, hyphens (-), and underscores (_).
	// Changing this parameter will create a new resource.
	// Specifies the access key of the user.
	// +kubebuilder:validation:Optional
	AccessKey *string `json:"accessKey,omitempty" tf:"access_key,omitempty"`

	// Specifies whether the user is an administrator.
	// Specifies whether the user is an administrator.
	// +kubebuilder:validation:Optional
	Admin *bool `json:"admin,omitempty" tf:"admin,omitempty"`

	// Specifies the default consumer group permissions.
	// Value options: PUB|SUB, PUB, SUB, DENY.
	// Specifies the default consumer group permissions.
	// Value options: **PUB|SUB**, **PUB**, **SUB**, **DENY**.
	// +kubebuilder:validation:Optional
	DefaultGroupPerm *string `json:"defaultGroupPerm,omitempty" tf:"default_group_perm,omitempty"`

	// Specifies the default topic permissions.
	// Value options: PUB|SUB, PUB, SUB, DENY.
	// Specifies the default topic permissions.
	// Value options: **PUB|SUB**, **PUB**, **SUB**, **DENY**.
	// +kubebuilder:validation:Optional
	DefaultTopicPerm *string `json:"defaultTopicPerm,omitempty" tf:"default_topic_perm,omitempty"`

	// Specifies the special consumer group permissions.
	// The permission structure is documented below.
	// Specifies the special consumer group permissions.
	// +kubebuilder:validation:Optional
	GroupPerms []GroupPermsParameters `json:"groupPerms,omitempty" tf:"group_perms,omitempty"`

	// Specifies the ID of the rocketMQ instance.
	// Changing this parameter will create a new resource.
	// Specifies the ID of the rocketMQ instance.
	// +crossplane:generate:reference:type=github.com/huaweicloud/provider-huaweicloud/apis/dms/v1alpha1.RocketmqInstance
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// Reference to a RocketmqInstance in dms to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDRef *v1.Reference `json:"instanceIdRef,omitempty" tf:"-"`

	// Selector for a RocketmqInstance in dms to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDSelector *v1.Selector `json:"instanceIdSelector,omitempty" tf:"-"`

	// Specifies the region in which to create the resource.
	// If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Specifies the password of the user. Use 8 to 32 characters. Contain at
	// least three of the following character types:
	// Specifies the secret key of the user.
	// +kubebuilder:validation:Optional
	SecretKey *string `json:"secretKey,omitempty" tf:"secret_key,omitempty"`

	// Specifies the special topic permissions.
	// The permission structure is documented below.
	// Specifies the special topic permissions.
	// +kubebuilder:validation:Optional
	TopicPerms []TopicPermsParameters `json:"topicPerms,omitempty" tf:"topic_perms,omitempty"`

	// Specifies the IP address whitelist.
	// Specifies the IP address whitelist.
	// +kubebuilder:validation:Optional
	WhiteRemoteAddress *string `json:"whiteRemoteAddress,omitempty" tf:"white_remote_address,omitempty"`
}

type TopicPermsInitParameters struct {

	// Indicates the name of a topic or consumer group.
	// Indicates the name of a topic or consumer group.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Indicates the permissions of the topic or consumer group.
	// Value options: PUB|SUB, PUB, SUB, DENY.
	// Indicates the permissions of the topic or consumer group.
	// Value options: **PUB|SUB**, **PUB**, **SUB**, **DENY**.
	Perm *string `json:"perm,omitempty" tf:"perm,omitempty"`
}

type TopicPermsObservation struct {

	// Indicates the name of a topic or consumer group.
	// Indicates the name of a topic or consumer group.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Indicates the permissions of the topic or consumer group.
	// Value options: PUB|SUB, PUB, SUB, DENY.
	// Indicates the permissions of the topic or consumer group.
	// Value options: **PUB|SUB**, **PUB**, **SUB**, **DENY**.
	Perm *string `json:"perm,omitempty" tf:"perm,omitempty"`
}

type TopicPermsParameters struct {

	// Indicates the name of a topic or consumer group.
	// Indicates the name of a topic or consumer group.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Indicates the permissions of the topic or consumer group.
	// Value options: PUB|SUB, PUB, SUB, DENY.
	// Indicates the permissions of the topic or consumer group.
	// Value options: **PUB|SUB**, **PUB**, **SUB**, **DENY**.
	// +kubebuilder:validation:Optional
	Perm *string `json:"perm,omitempty" tf:"perm,omitempty"`
}

// RocketmqUserSpec defines the desired state of RocketmqUser
type RocketmqUserSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RocketmqUserParameters `json:"forProvider"`
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
	InitProvider RocketmqUserInitParameters `json:"initProvider,omitempty"`
}

// RocketmqUserStatus defines the observed state of RocketmqUser.
type RocketmqUserStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RocketmqUserObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// RocketmqUser is the Schema for the RocketmqUsers API. ""
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,huaweicloud}
type RocketmqUser struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.accessKey) || (has(self.initProvider) && has(self.initProvider.accessKey))",message="spec.forProvider.accessKey is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.secretKey) || (has(self.initProvider) && has(self.initProvider.secretKey))",message="spec.forProvider.secretKey is a required parameter"
	Spec   RocketmqUserSpec   `json:"spec"`
	Status RocketmqUserStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RocketmqUserList contains a list of RocketmqUsers
type RocketmqUserList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RocketmqUser `json:"items"`
}

// Repository type metadata.
var (
	RocketmqUser_Kind             = "RocketmqUser"
	RocketmqUser_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RocketmqUser_Kind}.String()
	RocketmqUser_KindAPIVersion   = RocketmqUser_Kind + "." + CRDGroupVersion.String()
	RocketmqUser_GroupVersionKind = CRDGroupVersion.WithKind(RocketmqUser_Kind)
)

func init() {
	SchemeBuilder.Register(&RocketmqUser{}, &RocketmqUserList{})
}
