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

type RoleSharingPolicyObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	RoleName *string `json:"roleName,omitempty" tf:"role_name,omitempty"`

	SourceOrganizationID *string `json:"sourceOrganizationId,omitempty" tf:"source_organization_id,omitempty"`
}

type RoleSharingPolicyParameters struct {

	// Reference to a Organization to populate targetOrganizationId.
	// +kubebuilder:validation:Optional
	OrganizationRef *v1.Reference `json:"organizationRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Purpose *string `json:"purpose,omitempty" tf:"purpose,omitempty"`

	// +crossplane:generate:reference:type=Role
	// +crossplane:generate:reference:refFieldName=RoleRef
	// +kubebuilder:validation:Optional
	RoleID *string `json:"roleId,omitempty" tf:"role_id,omitempty"`

	// Selector for a Role to populate roleId.
	// +kubebuilder:validation:Optional
	RoleIDSelector *v1.Selector `json:"roleIdSelector,omitempty" tf:"-"`

	// Reference to a Role to populate roleId.
	// +kubebuilder:validation:Optional
	RoleRef *v1.Reference `json:"roleRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	SharingPolicy *string `json:"sharingPolicy" tf:"sharing_policy,omitempty"`

	// +crossplane:generate:reference:type=Organization
	// +crossplane:generate:reference:refFieldName=OrganizationRef
	// +kubebuilder:validation:Optional
	TargetOrganizationID *string `json:"targetOrganizationId,omitempty" tf:"target_organization_id,omitempty"`

	// Selector for a Organization to populate targetOrganizationId.
	// +kubebuilder:validation:Optional
	TargetOrganizationIDSelector *v1.Selector `json:"targetOrganizationIdSelector,omitempty" tf:"-"`
}

// RoleSharingPolicySpec defines the desired state of RoleSharingPolicy
type RoleSharingPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RoleSharingPolicyParameters `json:"forProvider"`
}

// RoleSharingPolicyStatus defines the observed state of RoleSharingPolicy.
type RoleSharingPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RoleSharingPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RoleSharingPolicy is the Schema for the RoleSharingPolicys API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,hsdp}
type RoleSharingPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RoleSharingPolicySpec   `json:"spec"`
	Status            RoleSharingPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RoleSharingPolicyList contains a list of RoleSharingPolicys
type RoleSharingPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RoleSharingPolicy `json:"items"`
}

// Repository type metadata.
var (
	RoleSharingPolicy_Kind             = "RoleSharingPolicy"
	RoleSharingPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RoleSharingPolicy_Kind}.String()
	RoleSharingPolicy_KindAPIVersion   = RoleSharingPolicy_Kind + "." + CRDGroupVersion.String()
	RoleSharingPolicy_GroupVersionKind = CRDGroupVersion.WithKind(RoleSharingPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&RoleSharingPolicy{}, &RoleSharingPolicyList{})
}
