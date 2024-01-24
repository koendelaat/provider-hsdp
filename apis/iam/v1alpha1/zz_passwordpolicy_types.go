// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ChallengePolicyInitParameters struct {

	// (Mandatory) A Multi-valued String attribute that contains one or more default question a user may use when setting their challenge questions.
	// +listType=set
	DefaultQuestions []*string `json:"defaultQuestions,omitempty" tf:"default_questions,omitempty"`

	// (Mandatory) An Integer indicates the maximum number of failed reset password attempts using challenges.
	MaxIncorrectAttempts *int64 `json:"maxIncorrectAttempts,omitempty" tf:"max_incorrect_attempts,omitempty"`

	// (Mandatory) An Integer indicating the minimum number of challenge answers a user MUST answer when attempting to reset their password.
	MinAnswerCount *int64 `json:"minAnswerCount,omitempty" tf:"min_answer_count,omitempty"`

	// (Mandatory) An Integer indicating the minimum number of challenge questions a user MUST answer when setting challenge question answers.
	MinQuestionCount *int64 `json:"minQuestionCount,omitempty" tf:"min_question_count,omitempty"`
}

type ChallengePolicyObservation struct {

	// (Mandatory) A Multi-valued String attribute that contains one or more default question a user may use when setting their challenge questions.
	// +listType=set
	DefaultQuestions []*string `json:"defaultQuestions,omitempty" tf:"default_questions,omitempty"`

	// (Mandatory) An Integer indicates the maximum number of failed reset password attempts using challenges.
	MaxIncorrectAttempts *int64 `json:"maxIncorrectAttempts,omitempty" tf:"max_incorrect_attempts,omitempty"`

	// (Mandatory) An Integer indicating the minimum number of challenge answers a user MUST answer when attempting to reset their password.
	MinAnswerCount *int64 `json:"minAnswerCount,omitempty" tf:"min_answer_count,omitempty"`

	// (Mandatory) An Integer indicating the minimum number of challenge questions a user MUST answer when setting challenge question answers.
	MinQuestionCount *int64 `json:"minQuestionCount,omitempty" tf:"min_question_count,omitempty"`
}

type ChallengePolicyParameters struct {

	// (Mandatory) A Multi-valued String attribute that contains one or more default question a user may use when setting their challenge questions.
	// +kubebuilder:validation:Optional
	// +listType=set
	DefaultQuestions []*string `json:"defaultQuestions,omitempty" tf:"default_questions,omitempty"`

	// (Mandatory) An Integer indicates the maximum number of failed reset password attempts using challenges.
	// +kubebuilder:validation:Optional
	MaxIncorrectAttempts *int64 `json:"maxIncorrectAttempts,omitempty" tf:"max_incorrect_attempts,omitempty"`

	// (Mandatory) An Integer indicating the minimum number of challenge answers a user MUST answer when attempting to reset their password.
	// +kubebuilder:validation:Optional
	MinAnswerCount *int64 `json:"minAnswerCount,omitempty" tf:"min_answer_count,omitempty"`

	// (Mandatory) An Integer indicating the minimum number of challenge questions a user MUST answer when setting challenge question answers.
	// +kubebuilder:validation:Optional
	MinQuestionCount *int64 `json:"minQuestionCount,omitempty" tf:"min_question_count,omitempty"`
}

type ComplexityInitParameters struct {

	// The maximum number of characters password can contain.
	MaxLength *int64 `json:"maxLength,omitempty" tf:"max_length,omitempty"`

	// The minimum number of characters password can contain. Default 8
	MinLength *int64 `json:"minLength,omitempty" tf:"min_length,omitempty"`

	// The minimum number of lower characters password can contain.
	MinLowercase *int64 `json:"minLowercase,omitempty" tf:"min_lowercase,omitempty"`

	MinNumerics *int64 `json:"minNumerics,omitempty" tf:"min_numerics,omitempty"`

	// The minimum number of special characters password can contain.
	MinSpecialChars *int64 `json:"minSpecialChars,omitempty" tf:"min_special_chars,omitempty"`

	// The minimum number of uppercase characters password can contain.
	MinUppercase *int64 `json:"minUppercase,omitempty" tf:"min_uppercase,omitempty"`
}

type ComplexityObservation struct {

	// The maximum number of characters password can contain.
	MaxLength *int64 `json:"maxLength,omitempty" tf:"max_length,omitempty"`

	// The minimum number of characters password can contain. Default 8
	MinLength *int64 `json:"minLength,omitempty" tf:"min_length,omitempty"`

	// The minimum number of lower characters password can contain.
	MinLowercase *int64 `json:"minLowercase,omitempty" tf:"min_lowercase,omitempty"`

	MinNumerics *int64 `json:"minNumerics,omitempty" tf:"min_numerics,omitempty"`

	// The minimum number of special characters password can contain.
	MinSpecialChars *int64 `json:"minSpecialChars,omitempty" tf:"min_special_chars,omitempty"`

	// The minimum number of uppercase characters password can contain.
	MinUppercase *int64 `json:"minUppercase,omitempty" tf:"min_uppercase,omitempty"`
}

type ComplexityParameters struct {

	// The maximum number of characters password can contain.
	// +kubebuilder:validation:Optional
	MaxLength *int64 `json:"maxLength,omitempty" tf:"max_length,omitempty"`

	// The minimum number of characters password can contain. Default 8
	// +kubebuilder:validation:Optional
	MinLength *int64 `json:"minLength,omitempty" tf:"min_length,omitempty"`

	// The minimum number of lower characters password can contain.
	// +kubebuilder:validation:Optional
	MinLowercase *int64 `json:"minLowercase,omitempty" tf:"min_lowercase,omitempty"`

	// +kubebuilder:validation:Optional
	MinNumerics *int64 `json:"minNumerics,omitempty" tf:"min_numerics,omitempty"`

	// The minimum number of special characters password can contain.
	// +kubebuilder:validation:Optional
	MinSpecialChars *int64 `json:"minSpecialChars,omitempty" tf:"min_special_chars,omitempty"`

	// The minimum number of uppercase characters password can contain.
	// +kubebuilder:validation:Optional
	MinUppercase *int64 `json:"minUppercase,omitempty" tf:"min_uppercase,omitempty"`
}

type PasswordPolicyInitParameters struct {

	// (Mandatory, if challenges_enabled = true) Specify KBA settings
	ChallengePolicy []ChallengePolicyInitParameters `json:"challengePolicy,omitempty" tf:"challenge_policy,omitempty"`

	// A boolean value indicating if challenges are enabled at organization level. If the value is set to true, challenge_policy attribute is mandatory.
	ChallengesEnabled *bool `json:"challengesEnabled,omitempty" tf:"challenges_enabled,omitempty"`

	// Different criteria that decides the strength of user password for an organization. Block
	Complexity []ComplexityInitParameters `json:"complexity,omitempty" tf:"complexity,omitempty"`

	// number - The number of days after which the user's password expires.
	ExpiryPeriodInDays *int64 `json:"expiryPeriodInDays,omitempty" tf:"expiry_period_in_days,omitempty"`

	// The number of previous passwords that cannot be used as new password.
	HistoryCount *int64 `json:"historyCount,omitempty" tf:"history_count,omitempty"`

	// The UUID of the IAM Org to apply this policy to
	// +crossplane:generate:reference:type=Organization
	// +crossplane:generate:reference:refFieldName=OrganizationRef
	ManagingOrganization *string `json:"managingOrganization,omitempty" tf:"managing_organization,omitempty"`

	// Selector for a Organization to populate managingOrganization.
	// +kubebuilder:validation:Optional
	ManagingOrganizationSelector *v1.Selector `json:"managingOrganizationSelector,omitempty" tf:"-"`

	// Reference to a Organization to populate managingOrganization.
	// +kubebuilder:validation:Optional
	OrganizationRef *v1.Reference `json:"organizationRef,omitempty" tf:"-"`
}

type PasswordPolicyObservation struct {

	// (Mandatory, if challenges_enabled = true) Specify KBA settings
	ChallengePolicy []ChallengePolicyObservation `json:"challengePolicy,omitempty" tf:"challenge_policy,omitempty"`

	// A boolean value indicating if challenges are enabled at organization level. If the value is set to true, challenge_policy attribute is mandatory.
	ChallengesEnabled *bool `json:"challengesEnabled,omitempty" tf:"challenges_enabled,omitempty"`

	// Different criteria that decides the strength of user password for an organization. Block
	Complexity []ComplexityObservation `json:"complexity,omitempty" tf:"complexity,omitempty"`

	// number - The number of days after which the user's password expires.
	ExpiryPeriodInDays *int64 `json:"expiryPeriodInDays,omitempty" tf:"expiry_period_in_days,omitempty"`

	// The number of previous passwords that cannot be used as new password.
	HistoryCount *int64 `json:"historyCount,omitempty" tf:"history_count,omitempty"`

	// The GUID of the password policy
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The UUID of the IAM Org to apply this policy to
	ManagingOrganization *string `json:"managingOrganization,omitempty" tf:"managing_organization,omitempty"`

	Policy *string `json:"Policy,omitempty" tf:"_policy,omitempty"`
}

type PasswordPolicyParameters struct {

	// (Mandatory, if challenges_enabled = true) Specify KBA settings
	// +kubebuilder:validation:Optional
	ChallengePolicy []ChallengePolicyParameters `json:"challengePolicy,omitempty" tf:"challenge_policy,omitempty"`

	// A boolean value indicating if challenges are enabled at organization level. If the value is set to true, challenge_policy attribute is mandatory.
	// +kubebuilder:validation:Optional
	ChallengesEnabled *bool `json:"challengesEnabled,omitempty" tf:"challenges_enabled,omitempty"`

	// Different criteria that decides the strength of user password for an organization. Block
	// +kubebuilder:validation:Optional
	Complexity []ComplexityParameters `json:"complexity,omitempty" tf:"complexity,omitempty"`

	// number - The number of days after which the user's password expires.
	// +kubebuilder:validation:Optional
	ExpiryPeriodInDays *int64 `json:"expiryPeriodInDays,omitempty" tf:"expiry_period_in_days,omitempty"`

	// The number of previous passwords that cannot be used as new password.
	// +kubebuilder:validation:Optional
	HistoryCount *int64 `json:"historyCount,omitempty" tf:"history_count,omitempty"`

	// The UUID of the IAM Org to apply this policy to
	// +crossplane:generate:reference:type=Organization
	// +crossplane:generate:reference:refFieldName=OrganizationRef
	// +kubebuilder:validation:Optional
	ManagingOrganization *string `json:"managingOrganization,omitempty" tf:"managing_organization,omitempty"`

	// Selector for a Organization to populate managingOrganization.
	// +kubebuilder:validation:Optional
	ManagingOrganizationSelector *v1.Selector `json:"managingOrganizationSelector,omitempty" tf:"-"`

	// Reference to a Organization to populate managingOrganization.
	// +kubebuilder:validation:Optional
	OrganizationRef *v1.Reference `json:"organizationRef,omitempty" tf:"-"`
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

// PasswordPolicy is the Schema for the PasswordPolicys API. Manages HSDP IAM Password policy resources
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,hsdp}
type PasswordPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.complexity) || (has(self.initProvider) && has(self.initProvider.complexity))",message="spec.forProvider.complexity is a required parameter"
	Spec   PasswordPolicySpec   `json:"spec"`
	Status PasswordPolicyStatus `json:"status,omitempty"`
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
