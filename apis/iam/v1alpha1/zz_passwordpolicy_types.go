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

type PasswordPolicyInitParameters struct {

	// Specifies the maximum number of times that a character is allowed
	// to consecutively present in a password. The value ranges from 0 to 32 and defaults to 0 which indicates that
	// consecutive identical characters are allowed in a password. For example, value 2 indicates that two consecutive
	// identical characters are not allowed in a password.
	MaximumConsecutiveIdenticalChars *float64 `json:"maximumConsecutiveIdenticalChars,omitempty" tf:"maximum_consecutive_identical_chars,omitempty"`

	// Specifies the minimum period (minutes) after which users are allowed to make
	// a password change. The value ranges from 0 to 1,440 and defaults to 0.
	MinimumPasswordAge *float64 `json:"minimumPasswordAge,omitempty" tf:"minimum_password_age,omitempty"`

	// Specifies the minimum number of characters that a password must contain.
	// The value ranges from 6 to 32 and defaults to 8.
	MinimumPasswordLength *float64 `json:"minimumPasswordLength,omitempty" tf:"minimum_password_length,omitempty"`

	// Specifies the member of previously used passwords that are
	// not allowed. The value ranges from 0 to 10 and defaults to 1. For example, value 3 indicates that the user cannot
	// set the last three passwords that the user has previously used when setting a new password.
	NumberOfRecentPasswordsDisallowed *float64 `json:"numberOfRecentPasswordsDisallowed,omitempty" tf:"number_of_recent_passwords_disallowed,omitempty"`

	// Specifies the minimum number of character types that a password must contain.
	// The value ranges from 2 to 4 and defaults to 2 which indicates that a password must contain at least two of the following:
	// uppercase letters, lowercase letters, digits, and special characters.
	PasswordCharCombination *float64 `json:"passwordCharCombination,omitempty" tf:"password_char_combination,omitempty"`

	// Specifies whether the password can be the username or the username
	// spelled backwards. Defaults to true, which indicates that the username or the inversion of username is not allowed to
	// be used as a password.
	PasswordNotUsernameOrInvert *bool `json:"passwordNotUsernameOrInvert,omitempty" tf:"password_not_username_or_invert,omitempty"`

	// Specifies the password validity period (days).
	// The value ranges from 0 to 180 and defaults to 0 which indicates that this requirement does not apply.
	PasswordValidityPeriod *float64 `json:"passwordValidityPeriod,omitempty" tf:"password_validity_period,omitempty"`
}

type PasswordPolicyObservation struct {

	// The ID of account password policy, which is the same as the account ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the maximum number of times that a character is allowed
	// to consecutively present in a password. The value ranges from 0 to 32 and defaults to 0 which indicates that
	// consecutive identical characters are allowed in a password. For example, value 2 indicates that two consecutive
	// identical characters are not allowed in a password.
	MaximumConsecutiveIdenticalChars *float64 `json:"maximumConsecutiveIdenticalChars,omitempty" tf:"maximum_consecutive_identical_chars,omitempty"`

	// The maximum number of characters that a password can contain.
	MaximumPasswordLength *float64 `json:"maximumPasswordLength,omitempty" tf:"maximum_password_length,omitempty"`

	// Specifies the minimum period (minutes) after which users are allowed to make
	// a password change. The value ranges from 0 to 1,440 and defaults to 0.
	MinimumPasswordAge *float64 `json:"minimumPasswordAge,omitempty" tf:"minimum_password_age,omitempty"`

	// Specifies the minimum number of characters that a password must contain.
	// The value ranges from 6 to 32 and defaults to 8.
	MinimumPasswordLength *float64 `json:"minimumPasswordLength,omitempty" tf:"minimum_password_length,omitempty"`

	// Specifies the member of previously used passwords that are
	// not allowed. The value ranges from 0 to 10 and defaults to 1. For example, value 3 indicates that the user cannot
	// set the last three passwords that the user has previously used when setting a new password.
	NumberOfRecentPasswordsDisallowed *float64 `json:"numberOfRecentPasswordsDisallowed,omitempty" tf:"number_of_recent_passwords_disallowed,omitempty"`

	// Specifies the minimum number of character types that a password must contain.
	// The value ranges from 2 to 4 and defaults to 2 which indicates that a password must contain at least two of the following:
	// uppercase letters, lowercase letters, digits, and special characters.
	PasswordCharCombination *float64 `json:"passwordCharCombination,omitempty" tf:"password_char_combination,omitempty"`

	// Specifies whether the password can be the username or the username
	// spelled backwards. Defaults to true, which indicates that the username or the inversion of username is not allowed to
	// be used as a password.
	PasswordNotUsernameOrInvert *bool `json:"passwordNotUsernameOrInvert,omitempty" tf:"password_not_username_or_invert,omitempty"`

	// Specifies the password validity period (days).
	// The value ranges from 0 to 180 and defaults to 0 which indicates that this requirement does not apply.
	PasswordValidityPeriod *float64 `json:"passwordValidityPeriod,omitempty" tf:"password_validity_period,omitempty"`
}

type PasswordPolicyParameters struct {

	// Specifies the maximum number of times that a character is allowed
	// to consecutively present in a password. The value ranges from 0 to 32 and defaults to 0 which indicates that
	// consecutive identical characters are allowed in a password. For example, value 2 indicates that two consecutive
	// identical characters are not allowed in a password.
	// +kubebuilder:validation:Optional
	MaximumConsecutiveIdenticalChars *float64 `json:"maximumConsecutiveIdenticalChars,omitempty" tf:"maximum_consecutive_identical_chars,omitempty"`

	// Specifies the minimum period (minutes) after which users are allowed to make
	// a password change. The value ranges from 0 to 1,440 and defaults to 0.
	// +kubebuilder:validation:Optional
	MinimumPasswordAge *float64 `json:"minimumPasswordAge,omitempty" tf:"minimum_password_age,omitempty"`

	// Specifies the minimum number of characters that a password must contain.
	// The value ranges from 6 to 32 and defaults to 8.
	// +kubebuilder:validation:Optional
	MinimumPasswordLength *float64 `json:"minimumPasswordLength,omitempty" tf:"minimum_password_length,omitempty"`

	// Specifies the member of previously used passwords that are
	// not allowed. The value ranges from 0 to 10 and defaults to 1. For example, value 3 indicates that the user cannot
	// set the last three passwords that the user has previously used when setting a new password.
	// +kubebuilder:validation:Optional
	NumberOfRecentPasswordsDisallowed *float64 `json:"numberOfRecentPasswordsDisallowed,omitempty" tf:"number_of_recent_passwords_disallowed,omitempty"`

	// Specifies the minimum number of character types that a password must contain.
	// The value ranges from 2 to 4 and defaults to 2 which indicates that a password must contain at least two of the following:
	// uppercase letters, lowercase letters, digits, and special characters.
	// +kubebuilder:validation:Optional
	PasswordCharCombination *float64 `json:"passwordCharCombination,omitempty" tf:"password_char_combination,omitempty"`

	// Specifies whether the password can be the username or the username
	// spelled backwards. Defaults to true, which indicates that the username or the inversion of username is not allowed to
	// be used as a password.
	// +kubebuilder:validation:Optional
	PasswordNotUsernameOrInvert *bool `json:"passwordNotUsernameOrInvert,omitempty" tf:"password_not_username_or_invert,omitempty"`

	// Specifies the password validity period (days).
	// The value ranges from 0 to 180 and defaults to 0 which indicates that this requirement does not apply.
	// +kubebuilder:validation:Optional
	PasswordValidityPeriod *float64 `json:"passwordValidityPeriod,omitempty" tf:"password_validity_period,omitempty"`
}

// PasswordPolicySpec defines the desired state of PasswordPolicy
type PasswordPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PasswordPolicyParameters `json:"forProvider"`
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
	InitProvider PasswordPolicyInitParameters `json:"initProvider,omitempty"`
}

// PasswordPolicyStatus defines the observed state of PasswordPolicy.
type PasswordPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PasswordPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// PasswordPolicy is the Schema for the PasswordPolicys API. ""
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,huaweicloud}
type PasswordPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PasswordPolicySpec   `json:"spec"`
	Status            PasswordPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PasswordPolicyList contains a list of PasswordPolicys
type PasswordPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PasswordPolicy `json:"items"`
}

// Repository type metadata.
var (
	PasswordPolicy_Kind             = "PasswordPolicy"
	PasswordPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PasswordPolicy_Kind}.String()
	PasswordPolicy_KindAPIVersion   = PasswordPolicy_Kind + "." + CRDGroupVersion.String()
	PasswordPolicy_GroupVersionKind = CRDGroupVersion.WithKind(PasswordPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&PasswordPolicy{}, &PasswordPolicyList{})
}
