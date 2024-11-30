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

type BucketReplicationInitParameters struct {

	// Specifies the IAM agency applied to the cross-region replication.
	Agency *string `json:"agency,omitempty" tf:"agency,omitempty"`

	// Specifies the name of the source bucket.
	// +crossplane:generate:reference:type=github.com/huaweicloud/provider-huaweicloud/apis/obs/v1alpha1.Bucket
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("bucket",true)
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// Reference to a Bucket in obs to populate bucket.
	// +kubebuilder:validation:Optional
	BucketRef *v1.Reference `json:"bucketRef,omitempty" tf:"-"`

	// Selector for a Bucket in obs to populate bucket.
	// +kubebuilder:validation:Optional
	BucketSelector *v1.Selector `json:"bucketSelector,omitempty" tf:"-"`

	// Specifies the name of the destination bucket.
	// +crossplane:generate:reference:type=github.com/huaweicloud/provider-huaweicloud/apis/obs/v1alpha1.Bucket
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("bucket",true)
	DestinationBucket *string `json:"destinationBucket,omitempty" tf:"destination_bucket,omitempty"`

	// Reference to a Bucket in obs to populate destinationBucket.
	// +kubebuilder:validation:Optional
	DestinationBucketRef *v1.Reference `json:"destinationBucketRef,omitempty" tf:"-"`

	// Selector for a Bucket in obs to populate destinationBucket.
	// +kubebuilder:validation:Optional
	DestinationBucketSelector *v1.Selector `json:"destinationBucketSelector,omitempty" tf:"-"`

	// Specifies the region in which to create the resource.
	// If omitted, the provider-level region will be used.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Specifies the configurations of object cross-region replication management.
	// The rule_struct structure is documented below.
	Rule []RuleInitParameters `json:"rule,omitempty" tf:"rule,omitempty"`
}

type BucketReplicationObservation struct {

	// Specifies the IAM agency applied to the cross-region replication.
	Agency *string `json:"agency,omitempty" tf:"agency,omitempty"`

	// Specifies the name of the source bucket.
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// Specifies the name of the destination bucket.
	DestinationBucket *string `json:"destinationBucket,omitempty" tf:"destination_bucket,omitempty"`

	// The name of the bucket.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the region in which to create the resource.
	// If omitted, the provider-level region will be used.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Specifies the configurations of object cross-region replication management.
	// The rule_struct structure is documented below.
	Rule []RuleObservation `json:"rule,omitempty" tf:"rule,omitempty"`
}

type BucketReplicationParameters struct {

	// Specifies the IAM agency applied to the cross-region replication.
	// +kubebuilder:validation:Optional
	Agency *string `json:"agency,omitempty" tf:"agency,omitempty"`

	// Specifies the name of the source bucket.
	// +crossplane:generate:reference:type=github.com/huaweicloud/provider-huaweicloud/apis/obs/v1alpha1.Bucket
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("bucket",true)
	// +kubebuilder:validation:Optional
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// Reference to a Bucket in obs to populate bucket.
	// +kubebuilder:validation:Optional
	BucketRef *v1.Reference `json:"bucketRef,omitempty" tf:"-"`

	// Selector for a Bucket in obs to populate bucket.
	// +kubebuilder:validation:Optional
	BucketSelector *v1.Selector `json:"bucketSelector,omitempty" tf:"-"`

	// Specifies the name of the destination bucket.
	// +crossplane:generate:reference:type=github.com/huaweicloud/provider-huaweicloud/apis/obs/v1alpha1.Bucket
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("bucket",true)
	// +kubebuilder:validation:Optional
	DestinationBucket *string `json:"destinationBucket,omitempty" tf:"destination_bucket,omitempty"`

	// Reference to a Bucket in obs to populate destinationBucket.
	// +kubebuilder:validation:Optional
	DestinationBucketRef *v1.Reference `json:"destinationBucketRef,omitempty" tf:"-"`

	// Selector for a Bucket in obs to populate destinationBucket.
	// +kubebuilder:validation:Optional
	DestinationBucketSelector *v1.Selector `json:"destinationBucketSelector,omitempty" tf:"-"`

	// Specifies the region in which to create the resource.
	// If omitted, the provider-level region will be used.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Specifies the configurations of object cross-region replication management.
	// The rule_struct structure is documented below.
	// +kubebuilder:validation:Optional
	Rule []RuleParameters `json:"rule,omitempty" tf:"rule,omitempty"`
}

type RuleInitParameters struct {

	// Specifies cross-region replication rule status. Defaults to true.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Specifies cross-region replication history rule status. Defaults to false.
	// If the value is true, historical objects meeting this rule are copied.
	HistoryEnabled *bool `json:"historyEnabled,omitempty" tf:"history_enabled,omitempty"`

	// Specifies the prefix of an object key name, applicable to one or more objects.
	// The maximum length of a prefix is 1024 characters.
	// Duplicated prefixes are not supported. If omitted, all objects in the bucket will be managed by the lifecycle rule.
	// To copy a folder, end the prefix with a slash (/), for example, imgs/.
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`

	// Specifies the storage class for replicated objects. Valid values are STANDARD,
	// WARM (Infrequent Access) and COLD (Archive).
	// If omitted, the storage class of object copies is the same as that of objects in the source bucket.
	StorageClass *string `json:"storageClass,omitempty" tf:"storage_class,omitempty"`
}

type RuleObservation struct {

	// Specifies cross-region replication rule status. Defaults to true.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Specifies cross-region replication history rule status. Defaults to false.
	// If the value is true, historical objects meeting this rule are copied.
	HistoryEnabled *bool `json:"historyEnabled,omitempty" tf:"history_enabled,omitempty"`

	// The name of the bucket.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the prefix of an object key name, applicable to one or more objects.
	// The maximum length of a prefix is 1024 characters.
	// Duplicated prefixes are not supported. If omitted, all objects in the bucket will be managed by the lifecycle rule.
	// To copy a folder, end the prefix with a slash (/), for example, imgs/.
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`

	// Specifies the storage class for replicated objects. Valid values are STANDARD,
	// WARM (Infrequent Access) and COLD (Archive).
	// If omitted, the storage class of object copies is the same as that of objects in the source bucket.
	StorageClass *string `json:"storageClass,omitempty" tf:"storage_class,omitempty"`
}

type RuleParameters struct {

	// Specifies cross-region replication rule status. Defaults to true.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Specifies cross-region replication history rule status. Defaults to false.
	// If the value is true, historical objects meeting this rule are copied.
	// +kubebuilder:validation:Optional
	HistoryEnabled *bool `json:"historyEnabled,omitempty" tf:"history_enabled,omitempty"`

	// Specifies the prefix of an object key name, applicable to one or more objects.
	// The maximum length of a prefix is 1024 characters.
	// Duplicated prefixes are not supported. If omitted, all objects in the bucket will be managed by the lifecycle rule.
	// To copy a folder, end the prefix with a slash (/), for example, imgs/.
	// +kubebuilder:validation:Optional
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`

	// Specifies the storage class for replicated objects. Valid values are STANDARD,
	// WARM (Infrequent Access) and COLD (Archive).
	// If omitted, the storage class of object copies is the same as that of objects in the source bucket.
	// +kubebuilder:validation:Optional
	StorageClass *string `json:"storageClass,omitempty" tf:"storage_class,omitempty"`
}

// BucketReplicationSpec defines the desired state of BucketReplication
type BucketReplicationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BucketReplicationParameters `json:"forProvider"`
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
	InitProvider BucketReplicationInitParameters `json:"initProvider,omitempty"`
}

// BucketReplicationStatus defines the observed state of BucketReplication.
type BucketReplicationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BucketReplicationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// BucketReplication is the Schema for the BucketReplications API. ""
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,huaweicloud}
type BucketReplication struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.agency) || (has(self.initProvider) && has(self.initProvider.agency))",message="spec.forProvider.agency is a required parameter"
	Spec   BucketReplicationSpec   `json:"spec"`
	Status BucketReplicationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BucketReplicationList contains a list of BucketReplications
type BucketReplicationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BucketReplication `json:"items"`
}

// Repository type metadata.
var (
	BucketReplication_Kind             = "BucketReplication"
	BucketReplication_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BucketReplication_Kind}.String()
	BucketReplication_KindAPIVersion   = BucketReplication_Kind + "." + CRDGroupVersion.String()
	BucketReplication_GroupVersionKind = CRDGroupVersion.WithKind(BucketReplication_Kind)
)

func init() {
	SchemeBuilder.Register(&BucketReplication{}, &BucketReplicationList{})
}
