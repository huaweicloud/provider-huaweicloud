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

type RabbitmqExchangeAssociateInitParameters struct {

	// Specifies the name of a target exchange or queue.
	// Changing this creates a new resource.
	Destination *string `json:"destination,omitempty" tf:"destination,omitempty"`

	// Specifies the type of the binding target.
	// The options are Exchange and Queue. Changing this creates a new resource.
	DestinationType *string `json:"destinationType,omitempty" tf:"destination_type,omitempty"`

	// Specifies the exchange name. Changing this creates a new resource.
	Exchange *string `json:"exchange,omitempty" tf:"exchange,omitempty"`

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

	// The region in which to create the resource.
	// If omitted, the provider-level region will be used. Changing this creates a new resource.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Specifies the binding key-value. Changing this creates a new resource.
	RoutingKey *string `json:"routingKey,omitempty" tf:"routing_key,omitempty"`

	// Specifies the vhost name. Changing this creates a new resource.
	Vhost *string `json:"vhost,omitempty" tf:"vhost,omitempty"`
}

type RabbitmqExchangeAssociateObservation struct {

	// Specifies the name of a target exchange or queue.
	// Changing this creates a new resource.
	Destination *string `json:"destination,omitempty" tf:"destination,omitempty"`

	// Specifies the type of the binding target.
	// The options are Exchange and Queue. Changing this creates a new resource.
	DestinationType *string `json:"destinationType,omitempty" tf:"destination_type,omitempty"`

	// Specifies the exchange name. Changing this creates a new resource.
	Exchange *string `json:"exchange,omitempty" tf:"exchange,omitempty"`

	// The resource ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the DMS RabbitMQ instance ID.
	// Changing this creates a new resource.
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// The URL-translated routing key.
	PropertiesKey *string `json:"propertiesKey,omitempty" tf:"properties_key,omitempty"`

	// The region in which to create the resource.
	// If omitted, the provider-level region will be used. Changing this creates a new resource.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Specifies the binding key-value. Changing this creates a new resource.
	RoutingKey *string `json:"routingKey,omitempty" tf:"routing_key,omitempty"`

	// Specifies the vhost name. Changing this creates a new resource.
	Vhost *string `json:"vhost,omitempty" tf:"vhost,omitempty"`
}

type RabbitmqExchangeAssociateParameters struct {

	// Specifies the name of a target exchange or queue.
	// Changing this creates a new resource.
	// +kubebuilder:validation:Optional
	Destination *string `json:"destination,omitempty" tf:"destination,omitempty"`

	// Specifies the type of the binding target.
	// The options are Exchange and Queue. Changing this creates a new resource.
	// +kubebuilder:validation:Optional
	DestinationType *string `json:"destinationType,omitempty" tf:"destination_type,omitempty"`

	// Specifies the exchange name. Changing this creates a new resource.
	// +kubebuilder:validation:Optional
	Exchange *string `json:"exchange,omitempty" tf:"exchange,omitempty"`

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

	// The region in which to create the resource.
	// If omitted, the provider-level region will be used. Changing this creates a new resource.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Specifies the binding key-value. Changing this creates a new resource.
	// +kubebuilder:validation:Optional
	RoutingKey *string `json:"routingKey,omitempty" tf:"routing_key,omitempty"`

	// Specifies the vhost name. Changing this creates a new resource.
	// +kubebuilder:validation:Optional
	Vhost *string `json:"vhost,omitempty" tf:"vhost,omitempty"`
}

// RabbitmqExchangeAssociateSpec defines the desired state of RabbitmqExchangeAssociate
type RabbitmqExchangeAssociateSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RabbitmqExchangeAssociateParameters `json:"forProvider"`
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
	InitProvider RabbitmqExchangeAssociateInitParameters `json:"initProvider,omitempty"`
}

// RabbitmqExchangeAssociateStatus defines the observed state of RabbitmqExchangeAssociate.
type RabbitmqExchangeAssociateStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RabbitmqExchangeAssociateObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// RabbitmqExchangeAssociate is the Schema for the RabbitmqExchangeAssociates API. Manages a DMS RabbitMQ exchange association resource within HuaweiCloud.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,huaweicloud}
type RabbitmqExchangeAssociate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.destination) || (has(self.initProvider) && has(self.initProvider.destination))",message="spec.forProvider.destination is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.destinationType) || (has(self.initProvider) && has(self.initProvider.destinationType))",message="spec.forProvider.destinationType is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.exchange) || (has(self.initProvider) && has(self.initProvider.exchange))",message="spec.forProvider.exchange is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.vhost) || (has(self.initProvider) && has(self.initProvider.vhost))",message="spec.forProvider.vhost is a required parameter"
	Spec   RabbitmqExchangeAssociateSpec   `json:"spec"`
	Status RabbitmqExchangeAssociateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RabbitmqExchangeAssociateList contains a list of RabbitmqExchangeAssociates
type RabbitmqExchangeAssociateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RabbitmqExchangeAssociate `json:"items"`
}

// Repository type metadata.
var (
	RabbitmqExchangeAssociate_Kind             = "RabbitmqExchangeAssociate"
	RabbitmqExchangeAssociate_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RabbitmqExchangeAssociate_Kind}.String()
	RabbitmqExchangeAssociate_KindAPIVersion   = RabbitmqExchangeAssociate_Kind + "." + CRDGroupVersion.String()
	RabbitmqExchangeAssociate_GroupVersionKind = CRDGroupVersion.WithKind(RabbitmqExchangeAssociate_Kind)
)

func init() {
	SchemeBuilder.Register(&RabbitmqExchangeAssociate{}, &RabbitmqExchangeAssociateList{})
}