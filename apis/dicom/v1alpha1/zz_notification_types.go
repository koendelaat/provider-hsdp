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

type NotificationObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type NotificationParameters struct {

	// +kubebuilder:validation:Required
	ConfigURL *string `json:"configUrl" tf:"config_url,omitempty"`

	// +crossplane:generate:reference:type=github.com/philips-software/provider-hsdp/apis/iam/v1alpha1.Organization
	// +crossplane:generate:reference:refFieldName=OrganizationRef
	// +kubebuilder:validation:Optional
	DefaultOrganizationID *string `json:"defaultOrganizationId,omitempty" tf:"default_organization_id,omitempty"`

	// Selector for a Organization in iam to populate defaultOrganizationId.
	// +kubebuilder:validation:Optional
	DefaultOrganizationIDSelector *v1.Selector `json:"defaultOrganizationIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// +kubebuilder:validation:Required
	EndpointURL *string `json:"endpointUrl" tf:"endpoint_url,omitempty"`

	// +kubebuilder:validation:Optional
	OrganizationID *string `json:"organizationId,omitempty" tf:"organization_id,omitempty"`

	// Reference to a Organization in iam to populate defaultOrganizationId.
	// +kubebuilder:validation:Optional
	OrganizationRef *v1.Reference `json:"organizationRef,omitempty" tf:"-"`
}

// NotificationSpec defines the desired state of Notification
type NotificationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NotificationParameters `json:"forProvider"`
}

// NotificationStatus defines the observed state of Notification.
type NotificationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NotificationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Notification is the Schema for the Notifications API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,hsdp}
type Notification struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NotificationSpec   `json:"spec"`
	Status            NotificationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NotificationList contains a list of Notifications
type NotificationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Notification `json:"items"`
}

// Repository type metadata.
var (
	Notification_Kind             = "Notification"
	Notification_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Notification_Kind}.String()
	Notification_KindAPIVersion   = Notification_Kind + "." + CRDGroupVersion.String()
	Notification_GroupVersionKind = CRDGroupVersion.WithKind(Notification_Kind)
)

func init() {
	SchemeBuilder.Register(&Notification{}, &NotificationList{})
}