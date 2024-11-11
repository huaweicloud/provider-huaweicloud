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

type ChannelDetailsInitParameters struct {
}

type ChannelDetailsObservation struct {

	// Indicates the connection details.
	ConnectionName *string `json:"connectionName,omitempty" tf:"connection_name,omitempty"`

	// Specifies the queue name.
	// Changing this creates a new resource.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Indicates the channel quantity.
	Number *float64 `json:"number,omitempty" tf:"number,omitempty"`

	// Indicates the IP address of the connected consumer.
	PeerHost *string `json:"peerHost,omitempty" tf:"peer_host,omitempty"`

	// Indicates the port of the process of the connected consumer.
	PeerPort *float64 `json:"peerPort,omitempty" tf:"peer_port,omitempty"`

	// Indicates the consumer username. If ACL is enabled, the real username will be returned, otherwise null will
	// be returned.
	User *string `json:"user,omitempty" tf:"user,omitempty"`
}

type ChannelDetailsParameters struct {
}

type ConsumerDetailsInitParameters struct {
}

type ConsumerDetailsObservation struct {

	// Indicates whether manual acknowledgement is enabled on the consumer client.
	AckRequired *bool `json:"ackRequired,omitempty" tf:"ack_required,omitempty"`

	// Indicates the consumer connections.
	// The channel_details structure is documented below.
	ChannelDetails []ChannelDetailsObservation `json:"channelDetails,omitempty" tf:"channel_details,omitempty"`

	// Indicates the consumer tag.
	ConsumerTag *string `json:"consumerTag,omitempty" tf:"consumer_tag,omitempty"`

	// Indicates the consumer client preset value.
	PrefetchCount *float64 `json:"prefetchCount,omitempty" tf:"prefetch_count,omitempty"`
}

type ConsumerDetailsParameters struct {
}

type QueueBindingsInitParameters struct {
}

type QueueBindingsObservation struct {

	// Indicates the binding target name.
	Destination *string `json:"destination,omitempty" tf:"destination,omitempty"`

	// Indicates the binding target type.
	DestinationType *string `json:"destinationType,omitempty" tf:"destination_type,omitempty"`

	// Indicates the URL-translated routing key.
	PropertiesKey *string `json:"propertiesKey,omitempty" tf:"properties_key,omitempty"`

	// Indicates the binding key-value.
	RoutingKey *string `json:"routingKey,omitempty" tf:"routing_key,omitempty"`

	// Indicates the exchange name.
	Source *string `json:"source,omitempty" tf:"source,omitempty"`
}

type QueueBindingsParameters struct {
}

type RabbitmqQueueInitParameters struct {

	// Specifies whether to enable auto delete.
	// Changing this creates a new resource.
	AutoDelete *bool `json:"autoDelete,omitempty" tf:"auto_delete,omitempty"`

	// Specifies the name of the dead letter exchange.
	// It's required when dead_letter_routing_key is specified.
	// Changing this creates a new resource.
	DeadLetterExchange *string `json:"deadLetterExchange,omitempty" tf:"dead_letter_exchange,omitempty"`

	// Specifies the routing key of the dead letter exchange.
	// Changing this creates a new resource.
	DeadLetterRoutingKey *string `json:"deadLetterRoutingKey,omitempty" tf:"dead_letter_routing_key,omitempty"`

	// Specifies whether to enable durable. Defaults to false.
	// Changing this creates a new resource.
	Durable *bool `json:"durable,omitempty" tf:"durable,omitempty"`

	// Specifies the DMS RabbitMQ instance ID.
	// Changing this creates a new resource.
	// +crossplane:generate:reference:type=github.com/huaweicloud/provider-huaweicloud/apis/dms/v1alpha1.RabbitmqInstance
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// Reference to a RabbitmqInstance in dms to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDRef *v1.Reference `json:"instanceIdRef,omitempty" tf:"-"`

	// Selector for a RabbitmqInstance in dms to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDSelector *v1.Selector `json:"instanceIdSelector,omitempty" tf:"-"`

	// Specifies the lazy mode. Valid value is lazy.
	// Changing this creates a new resource.
	LazyMode *string `json:"lazyMode,omitempty" tf:"lazy_mode,omitempty"`

	// Specifies how long a message in this queue can be retained.
	// Changing this creates a new resource.
	MessageTTL *float64 `json:"messageTtl,omitempty" tf:"message_ttl,omitempty"`

	// Specifies the queue name.
	// Changing this creates a new resource.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies the region in which to create the resource.
	// If omitted, the provider-level region will be used.
	// Changing this creates a new resource.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Specifies the vhost name.
	// Changing this creates a new resource.
	Vhost *string `json:"vhost,omitempty" tf:"vhost,omitempty"`
}

type RabbitmqQueueObservation struct {

	// Specifies whether to enable auto delete.
	// Changing this creates a new resource.
	AutoDelete *bool `json:"autoDelete,omitempty" tf:"auto_delete,omitempty"`

	// Indicates the details of subscribed consumers.
	// The consumer_details structure is documented below.
	ConsumerDetails []ConsumerDetailsObservation `json:"consumerDetails,omitempty" tf:"consumer_details,omitempty"`

	// Indicates the connected consumers.
	Consumers *float64 `json:"consumers,omitempty" tf:"consumers,omitempty"`

	// Specifies the name of the dead letter exchange.
	// It's required when dead_letter_routing_key is specified.
	// Changing this creates a new resource.
	DeadLetterExchange *string `json:"deadLetterExchange,omitempty" tf:"dead_letter_exchange,omitempty"`

	// Specifies the routing key of the dead letter exchange.
	// Changing this creates a new resource.
	DeadLetterRoutingKey *string `json:"deadLetterRoutingKey,omitempty" tf:"dead_letter_routing_key,omitempty"`

	// Specifies whether to enable durable. Defaults to false.
	// Changing this creates a new resource.
	Durable *bool `json:"durable,omitempty" tf:"durable,omitempty"`

	// The resource ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the DMS RabbitMQ instance ID.
	// Changing this creates a new resource.
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// Specifies the lazy mode. Valid value is lazy.
	// Changing this creates a new resource.
	LazyMode *string `json:"lazyMode,omitempty" tf:"lazy_mode,omitempty"`

	// Specifies how long a message in this queue can be retained.
	// Changing this creates a new resource.
	MessageTTL *float64 `json:"messageTtl,omitempty" tf:"message_ttl,omitempty"`

	// Indicates the accumulated messages.
	Messages *float64 `json:"messages,omitempty" tf:"messages,omitempty"`

	// Specifies the queue name.
	// Changing this creates a new resource.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Indicates the policy.
	Policy *string `json:"policy,omitempty" tf:"policy,omitempty"`

	// Indicates the bindings to this queue.
	// The queue_bindings structure is documented below.
	QueueBindings []QueueBindingsObservation `json:"queueBindings,omitempty" tf:"queue_bindings,omitempty"`

	// Specifies the region in which to create the resource.
	// If omitted, the provider-level region will be used.
	// Changing this creates a new resource.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Specifies the vhost name.
	// Changing this creates a new resource.
	Vhost *string `json:"vhost,omitempty" tf:"vhost,omitempty"`
}

type RabbitmqQueueParameters struct {

	// Specifies whether to enable auto delete.
	// Changing this creates a new resource.
	// +kubebuilder:validation:Optional
	AutoDelete *bool `json:"autoDelete,omitempty" tf:"auto_delete,omitempty"`

	// Specifies the name of the dead letter exchange.
	// It's required when dead_letter_routing_key is specified.
	// Changing this creates a new resource.
	// +kubebuilder:validation:Optional
	DeadLetterExchange *string `json:"deadLetterExchange,omitempty" tf:"dead_letter_exchange,omitempty"`

	// Specifies the routing key of the dead letter exchange.
	// Changing this creates a new resource.
	// +kubebuilder:validation:Optional
	DeadLetterRoutingKey *string `json:"deadLetterRoutingKey,omitempty" tf:"dead_letter_routing_key,omitempty"`

	// Specifies whether to enable durable. Defaults to false.
	// Changing this creates a new resource.
	// +kubebuilder:validation:Optional
	Durable *bool `json:"durable,omitempty" tf:"durable,omitempty"`

	// Specifies the DMS RabbitMQ instance ID.
	// Changing this creates a new resource.
	// +crossplane:generate:reference:type=github.com/huaweicloud/provider-huaweicloud/apis/dms/v1alpha1.RabbitmqInstance
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// Reference to a RabbitmqInstance in dms to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDRef *v1.Reference `json:"instanceIdRef,omitempty" tf:"-"`

	// Selector for a RabbitmqInstance in dms to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDSelector *v1.Selector `json:"instanceIdSelector,omitempty" tf:"-"`

	// Specifies the lazy mode. Valid value is lazy.
	// Changing this creates a new resource.
	// +kubebuilder:validation:Optional
	LazyMode *string `json:"lazyMode,omitempty" tf:"lazy_mode,omitempty"`

	// Specifies how long a message in this queue can be retained.
	// Changing this creates a new resource.
	// +kubebuilder:validation:Optional
	MessageTTL *float64 `json:"messageTtl,omitempty" tf:"message_ttl,omitempty"`

	// Specifies the queue name.
	// Changing this creates a new resource.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies the region in which to create the resource.
	// If omitted, the provider-level region will be used.
	// Changing this creates a new resource.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Specifies the vhost name.
	// Changing this creates a new resource.
	// +kubebuilder:validation:Optional
	Vhost *string `json:"vhost,omitempty" tf:"vhost,omitempty"`
}

// RabbitmqQueueSpec defines the desired state of RabbitmqQueue
type RabbitmqQueueSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RabbitmqQueueParameters `json:"forProvider"`
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
	InitProvider RabbitmqQueueInitParameters `json:"initProvider,omitempty"`
}

// RabbitmqQueueStatus defines the observed state of RabbitmqQueue.
type RabbitmqQueueStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RabbitmqQueueObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// RabbitmqQueue is the Schema for the RabbitmqQueues API. Manages a DMS RabbitMQ queue resource within HuaweiCloud.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,huaweicloud}
type RabbitmqQueue struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.autoDelete) || (has(self.initProvider) && has(self.initProvider.autoDelete))",message="spec.forProvider.autoDelete is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.vhost) || (has(self.initProvider) && has(self.initProvider.vhost))",message="spec.forProvider.vhost is a required parameter"
	Spec   RabbitmqQueueSpec   `json:"spec"`
	Status RabbitmqQueueStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RabbitmqQueueList contains a list of RabbitmqQueues
type RabbitmqQueueList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RabbitmqQueue `json:"items"`
}

// Repository type metadata.
var (
	RabbitmqQueue_Kind             = "RabbitmqQueue"
	RabbitmqQueue_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RabbitmqQueue_Kind}.String()
	RabbitmqQueue_KindAPIVersion   = RabbitmqQueue_Kind + "." + CRDGroupVersion.String()
	RabbitmqQueue_GroupVersionKind = CRDGroupVersion.WithKind(RabbitmqQueue_Kind)
)

func init() {
	SchemeBuilder.Register(&RabbitmqQueue{}, &RabbitmqQueueList{})
}
