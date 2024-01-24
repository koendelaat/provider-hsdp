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

type RepositoryInitParameters struct {

	// The base config URL of the DICOM Store instance
	ConfigURL *string `json:"configUrl,omitempty" tf:"config_url,omitempty"`

	// (Block, Optional)
	Notification []RepositoryNotificationInitParameters `json:"notification,omitempty" tf:"notification,omitempty"`

	// the Object store ID
	ObjectStoreID *string `json:"objectStoreId,omitempty" tf:"object_store_id,omitempty"`

	// The organization ID
	// +crossplane:generate:reference:type=github.com/philips-software/provider-hsdp/apis/iam/v1alpha1.Organization
	// +crossplane:generate:reference:refFieldName=OrganizationRef
	OrganizationID *string `json:"organizationId,omitempty" tf:"organization_id,omitempty"`

	// Selector for a Organization in iam to populate organizationId.
	// +kubebuilder:validation:Optional
	OrganizationIDSelector *v1.Selector `json:"organizationIdSelector,omitempty" tf:"-"`

	// Reference to a Organization in iam to populate organizationId.
	// +kubebuilder:validation:Optional
	OrganizationRef *v1.Reference `json:"organizationRef,omitempty" tf:"-"`

	// The organization ID attached to this repository.
	// When not specified, the root organization is used.
	RepositoryOrganizationID *string `json:"repositoryOrganizationId,omitempty" tf:"repository_organization_id,omitempty"`

	// Configure this repository as store as composite.
	StoreAsComposite *bool `json:"storeAsComposite,omitempty" tf:"store_as_composite,omitempty"`
}

type RepositoryNotificationInitParameters struct {

	// Enable notifications or not. Default: true
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The organization ID
	OrganizationID *string `json:"organizationId,omitempty" tf:"organization_id,omitempty"`
}

type RepositoryNotificationObservation struct {

	// Enable notifications or not. Default: true
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The organization ID
	OrganizationID *string `json:"organizationId,omitempty" tf:"organization_id,omitempty"`
}

type RepositoryNotificationParameters struct {

	// Enable notifications or not. Default: true
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The organization ID
	// +kubebuilder:validation:Optional
	OrganizationID *string `json:"organizationId" tf:"organization_id,omitempty"`
}

type RepositoryObservation struct {

	// The base config URL of the DICOM Store instance
	ConfigURL *string `json:"configUrl,omitempty" tf:"config_url,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (Block, Optional)
	Notification []RepositoryNotificationObservation `json:"notification,omitempty" tf:"notification,omitempty"`

	// the Object store ID
	ObjectStoreID *string `json:"objectStoreId,omitempty" tf:"object_store_id,omitempty"`

	// The organization ID
	OrganizationID *string `json:"organizationId,omitempty" tf:"organization_id,omitempty"`

	// The organization ID attached to this repository.
	// When not specified, the root organization is used.
	RepositoryOrganizationID *string `json:"repositoryOrganizationId,omitempty" tf:"repository_organization_id,omitempty"`

	// Configure this repository as store as composite.
	StoreAsComposite *bool `json:"storeAsComposite,omitempty" tf:"store_as_composite,omitempty"`
}

type RepositoryParameters struct {

	// The base config URL of the DICOM Store instance
	// +kubebuilder:validation:Optional
	ConfigURL *string `json:"configUrl,omitempty" tf:"config_url,omitempty"`

	// (Block, Optional)
	// +kubebuilder:validation:Optional
	Notification []RepositoryNotificationParameters `json:"notification,omitempty" tf:"notification,omitempty"`

	// the Object store ID
	// +kubebuilder:validation:Optional
	ObjectStoreID *string `json:"objectStoreId,omitempty" tf:"object_store_id,omitempty"`

	// The organization ID
	// +crossplane:generate:reference:type=github.com/philips-software/provider-hsdp/apis/iam/v1alpha1.Organization
	// +crossplane:generate:reference:refFieldName=OrganizationRef
	// +kubebuilder:validation:Optional
	OrganizationID *string `json:"organizationId,omitempty" tf:"organization_id,omitempty"`

	// Selector for a Organization in iam to populate organizationId.
	// +kubebuilder:validation:Optional
	OrganizationIDSelector *v1.Selector `json:"organizationIdSelector,omitempty" tf:"-"`

	// Reference to a Organization in iam to populate organizationId.
	// +kubebuilder:validation:Optional
	OrganizationRef *v1.Reference `json:"organizationRef,omitempty" tf:"-"`

	// The organization ID attached to this repository.
	// When not specified, the root organization is used.
	// +kubebuilder:validation:Optional
	RepositoryOrganizationID *string `json:"repositoryOrganizationId,omitempty" tf:"repository_organization_id,omitempty"`

	// Configure this repository as store as composite.
	// +kubebuilder:validation:Optional
	StoreAsComposite *bool `json:"storeAsComposite,omitempty" tf:"store_as_composite,omitempty"`
}

// RepositorySpec defines the desired state of Repository
type RepositorySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RepositoryParameters `json:"forProvider"`
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
	InitProvider RepositoryInitParameters `json:"initProvider,omitempty"`
}

// RepositoryStatus defines the observed state of Repository.
type RepositoryStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RepositoryObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Repository is the Schema for the Repositorys API. Manages HSDP DICOM Repository resources
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,hsdp}
type Repository struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.configUrl) || (has(self.initProvider) && has(self.initProvider.configUrl))",message="spec.forProvider.configUrl is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.objectStoreId) || (has(self.initProvider) && has(self.initProvider.objectStoreId))",message="spec.forProvider.objectStoreId is a required parameter"
	Spec   RepositorySpec   `json:"spec"`
	Status RepositoryStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RepositoryList contains a list of Repositorys
type RepositoryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Repository `json:"items"`
}

// Repository type metadata.
var (
	Repository_Kind             = "Repository"
	Repository_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Repository_Kind}.String()
	Repository_KindAPIVersion   = Repository_Kind + "." + CRDGroupVersion.String()
	Repository_GroupVersionKind = CRDGroupVersion.WithKind(Repository_Kind)
)

func init() {
	SchemeBuilder.Register(&Repository{}, &RepositoryList{})
}
