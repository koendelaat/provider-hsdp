//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CdrServiceAccountObservation) DeepCopyInto(out *CdrServiceAccountObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CdrServiceAccountObservation.
func (in *CdrServiceAccountObservation) DeepCopy() *CdrServiceAccountObservation {
	if in == nil {
		return nil
	}
	out := new(CdrServiceAccountObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CdrServiceAccountParameters) DeepCopyInto(out *CdrServiceAccountParameters) {
	*out = *in
	if in.PrivateKey != nil {
		in, out := &in.PrivateKey, &out.PrivateKey
		*out = new(string)
		**out = **in
	}
	if in.ServiceID != nil {
		in, out := &in.ServiceID, &out.ServiceID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CdrServiceAccountParameters.
func (in *CdrServiceAccountParameters) DeepCopy() *CdrServiceAccountParameters {
	if in == nil {
		return nil
	}
	out := new(CdrServiceAccountParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FHIRStoreObservation) DeepCopyInto(out *FHIRStoreObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FHIRStoreObservation.
func (in *FHIRStoreObservation) DeepCopy() *FHIRStoreObservation {
	if in == nil {
		return nil
	}
	out := new(FHIRStoreObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FHIRStoreParameters) DeepCopyInto(out *FHIRStoreParameters) {
	*out = *in
	if in.MpiEndpoint != nil {
		in, out := &in.MpiEndpoint, &out.MpiEndpoint
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FHIRStoreParameters.
func (in *FHIRStoreParameters) DeepCopy() *FHIRStoreParameters {
	if in == nil {
		return nil
	}
	out := new(FHIRStoreParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Notification) DeepCopyInto(out *Notification) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Notification.
func (in *Notification) DeepCopy() *Notification {
	if in == nil {
		return nil
	}
	out := new(Notification)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Notification) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NotificationList) DeepCopyInto(out *NotificationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Notification, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NotificationList.
func (in *NotificationList) DeepCopy() *NotificationList {
	if in == nil {
		return nil
	}
	out := new(NotificationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NotificationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NotificationObservation) DeepCopyInto(out *NotificationObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NotificationObservation.
func (in *NotificationObservation) DeepCopy() *NotificationObservation {
	if in == nil {
		return nil
	}
	out := new(NotificationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NotificationParameters) DeepCopyInto(out *NotificationParameters) {
	*out = *in
	if in.ConfigURL != nil {
		in, out := &in.ConfigURL, &out.ConfigURL
		*out = new(string)
		**out = **in
	}
	if in.DefaultOrganizationID != nil {
		in, out := &in.DefaultOrganizationID, &out.DefaultOrganizationID
		*out = new(string)
		**out = **in
	}
	if in.DefaultOrganizationIDSelector != nil {
		in, out := &in.DefaultOrganizationIDSelector, &out.DefaultOrganizationIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.EndpointURL != nil {
		in, out := &in.EndpointURL, &out.EndpointURL
		*out = new(string)
		**out = **in
	}
	if in.OrganizationID != nil {
		in, out := &in.OrganizationID, &out.OrganizationID
		*out = new(string)
		**out = **in
	}
	if in.OrganizationRef != nil {
		in, out := &in.OrganizationRef, &out.OrganizationRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NotificationParameters.
func (in *NotificationParameters) DeepCopy() *NotificationParameters {
	if in == nil {
		return nil
	}
	out := new(NotificationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NotificationSpec) DeepCopyInto(out *NotificationSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NotificationSpec.
func (in *NotificationSpec) DeepCopy() *NotificationSpec {
	if in == nil {
		return nil
	}
	out := new(NotificationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NotificationStatus) DeepCopyInto(out *NotificationStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NotificationStatus.
func (in *NotificationStatus) DeepCopy() *NotificationStatus {
	if in == nil {
		return nil
	}
	out := new(NotificationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectStore) DeepCopyInto(out *ObjectStore) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectStore.
func (in *ObjectStore) DeepCopy() *ObjectStore {
	if in == nil {
		return nil
	}
	out := new(ObjectStore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ObjectStore) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectStoreList) DeepCopyInto(out *ObjectStoreList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ObjectStore, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectStoreList.
func (in *ObjectStoreList) DeepCopy() *ObjectStoreList {
	if in == nil {
		return nil
	}
	out := new(ObjectStoreList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ObjectStoreList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectStoreObservation) DeepCopyInto(out *ObjectStoreObservation) {
	*out = *in
	if in.AccessType != nil {
		in, out := &in.AccessType, &out.AccessType
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectStoreObservation.
func (in *ObjectStoreObservation) DeepCopy() *ObjectStoreObservation {
	if in == nil {
		return nil
	}
	out := new(ObjectStoreObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectStoreParameters) DeepCopyInto(out *ObjectStoreParameters) {
	*out = *in
	if in.ConfigURL != nil {
		in, out := &in.ConfigURL, &out.ConfigURL
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.ForceDelete != nil {
		in, out := &in.ForceDelete, &out.ForceDelete
		*out = new(bool)
		**out = **in
	}
	if in.OrganizationID != nil {
		in, out := &in.OrganizationID, &out.OrganizationID
		*out = new(string)
		**out = **in
	}
	if in.OrganizationIDSelector != nil {
		in, out := &in.OrganizationIDSelector, &out.OrganizationIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.OrganizationRef != nil {
		in, out := &in.OrganizationRef, &out.OrganizationRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.S3CredsAccess != nil {
		in, out := &in.S3CredsAccess, &out.S3CredsAccess
		*out = make([]S3CredsAccessParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.StaticAccess != nil {
		in, out := &in.StaticAccess, &out.StaticAccess
		*out = make([]StaticAccessParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectStoreParameters.
func (in *ObjectStoreParameters) DeepCopy() *ObjectStoreParameters {
	if in == nil {
		return nil
	}
	out := new(ObjectStoreParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectStoreSpec) DeepCopyInto(out *ObjectStoreSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectStoreSpec.
func (in *ObjectStoreSpec) DeepCopy() *ObjectStoreSpec {
	if in == nil {
		return nil
	}
	out := new(ObjectStoreSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectStoreStatus) DeepCopyInto(out *ObjectStoreStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectStoreStatus.
func (in *ObjectStoreStatus) DeepCopy() *ObjectStoreStatus {
	if in == nil {
		return nil
	}
	out := new(ObjectStoreStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Repository) DeepCopyInto(out *Repository) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Repository.
func (in *Repository) DeepCopy() *Repository {
	if in == nil {
		return nil
	}
	out := new(Repository)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Repository) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RepositoryList) DeepCopyInto(out *RepositoryList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Repository, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RepositoryList.
func (in *RepositoryList) DeepCopy() *RepositoryList {
	if in == nil {
		return nil
	}
	out := new(RepositoryList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RepositoryList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RepositoryNotificationObservation) DeepCopyInto(out *RepositoryNotificationObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RepositoryNotificationObservation.
func (in *RepositoryNotificationObservation) DeepCopy() *RepositoryNotificationObservation {
	if in == nil {
		return nil
	}
	out := new(RepositoryNotificationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RepositoryNotificationParameters) DeepCopyInto(out *RepositoryNotificationParameters) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.OrganizationID != nil {
		in, out := &in.OrganizationID, &out.OrganizationID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RepositoryNotificationParameters.
func (in *RepositoryNotificationParameters) DeepCopy() *RepositoryNotificationParameters {
	if in == nil {
		return nil
	}
	out := new(RepositoryNotificationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RepositoryObservation) DeepCopyInto(out *RepositoryObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RepositoryObservation.
func (in *RepositoryObservation) DeepCopy() *RepositoryObservation {
	if in == nil {
		return nil
	}
	out := new(RepositoryObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RepositoryParameters) DeepCopyInto(out *RepositoryParameters) {
	*out = *in
	if in.ConfigURL != nil {
		in, out := &in.ConfigURL, &out.ConfigURL
		*out = new(string)
		**out = **in
	}
	if in.Notification != nil {
		in, out := &in.Notification, &out.Notification
		*out = make([]RepositoryNotificationParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ObjectStoreID != nil {
		in, out := &in.ObjectStoreID, &out.ObjectStoreID
		*out = new(string)
		**out = **in
	}
	if in.OrganizationID != nil {
		in, out := &in.OrganizationID, &out.OrganizationID
		*out = new(string)
		**out = **in
	}
	if in.OrganizationIDSelector != nil {
		in, out := &in.OrganizationIDSelector, &out.OrganizationIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.OrganizationRef != nil {
		in, out := &in.OrganizationRef, &out.OrganizationRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.RepositoryOrganizationID != nil {
		in, out := &in.RepositoryOrganizationID, &out.RepositoryOrganizationID
		*out = new(string)
		**out = **in
	}
	if in.StoreAsComposite != nil {
		in, out := &in.StoreAsComposite, &out.StoreAsComposite
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RepositoryParameters.
func (in *RepositoryParameters) DeepCopy() *RepositoryParameters {
	if in == nil {
		return nil
	}
	out := new(RepositoryParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RepositorySpec) DeepCopyInto(out *RepositorySpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RepositorySpec.
func (in *RepositorySpec) DeepCopy() *RepositorySpec {
	if in == nil {
		return nil
	}
	out := new(RepositorySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RepositoryStatus) DeepCopyInto(out *RepositoryStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RepositoryStatus.
func (in *RepositoryStatus) DeepCopy() *RepositoryStatus {
	if in == nil {
		return nil
	}
	out := new(RepositoryStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3CredsAccessObservation) DeepCopyInto(out *S3CredsAccessObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3CredsAccessObservation.
func (in *S3CredsAccessObservation) DeepCopy() *S3CredsAccessObservation {
	if in == nil {
		return nil
	}
	out := new(S3CredsAccessObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3CredsAccessParameters) DeepCopyInto(out *S3CredsAccessParameters) {
	*out = *in
	if in.BucketName != nil {
		in, out := &in.BucketName, &out.BucketName
		*out = new(string)
		**out = **in
	}
	if in.Endpoint != nil {
		in, out := &in.Endpoint, &out.Endpoint
		*out = new(string)
		**out = **in
	}
	if in.FolderPath != nil {
		in, out := &in.FolderPath, &out.FolderPath
		*out = new(string)
		**out = **in
	}
	out.ProductKeySecretRef = in.ProductKeySecretRef
	if in.ServiceAccount != nil {
		in, out := &in.ServiceAccount, &out.ServiceAccount
		*out = make([]ServiceAccountParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3CredsAccessParameters.
func (in *S3CredsAccessParameters) DeepCopy() *S3CredsAccessParameters {
	if in == nil {
		return nil
	}
	out := new(S3CredsAccessParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceAccountObservation) DeepCopyInto(out *ServiceAccountObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceAccountObservation.
func (in *ServiceAccountObservation) DeepCopy() *ServiceAccountObservation {
	if in == nil {
		return nil
	}
	out := new(ServiceAccountObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceAccountParameters) DeepCopyInto(out *ServiceAccountParameters) {
	*out = *in
	if in.AccessTokenEndpoint != nil {
		in, out := &in.AccessTokenEndpoint, &out.AccessTokenEndpoint
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	out.PrivateKeySecretRef = in.PrivateKeySecretRef
	if in.ServiceID != nil {
		in, out := &in.ServiceID, &out.ServiceID
		*out = new(string)
		**out = **in
	}
	if in.TokenEndpoint != nil {
		in, out := &in.TokenEndpoint, &out.TokenEndpoint
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceAccountParameters.
func (in *ServiceAccountParameters) DeepCopy() *ServiceAccountParameters {
	if in == nil {
		return nil
	}
	out := new(ServiceAccountParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StaticAccessObservation) DeepCopyInto(out *StaticAccessObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StaticAccessObservation.
func (in *StaticAccessObservation) DeepCopy() *StaticAccessObservation {
	if in == nil {
		return nil
	}
	out := new(StaticAccessObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StaticAccessParameters) DeepCopyInto(out *StaticAccessParameters) {
	*out = *in
	out.AccessKeySecretRef = in.AccessKeySecretRef
	if in.BucketName != nil {
		in, out := &in.BucketName, &out.BucketName
		*out = new(string)
		**out = **in
	}
	if in.Endpoint != nil {
		in, out := &in.Endpoint, &out.Endpoint
		*out = new(string)
		**out = **in
	}
	out.SecretKeySecretRef = in.SecretKeySecretRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StaticAccessParameters.
func (in *StaticAccessParameters) DeepCopy() *StaticAccessParameters {
	if in == nil {
		return nil
	}
	out := new(StaticAccessParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StoreConfig) DeepCopyInto(out *StoreConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StoreConfig.
func (in *StoreConfig) DeepCopy() *StoreConfig {
	if in == nil {
		return nil
	}
	out := new(StoreConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StoreConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StoreConfigList) DeepCopyInto(out *StoreConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]StoreConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StoreConfigList.
func (in *StoreConfigList) DeepCopy() *StoreConfigList {
	if in == nil {
		return nil
	}
	out := new(StoreConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StoreConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StoreConfigObservation) DeepCopyInto(out *StoreConfigObservation) {
	*out = *in
	if in.CdrServiceAccount != nil {
		in, out := &in.CdrServiceAccount, &out.CdrServiceAccount
		*out = make([]CdrServiceAccountObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DataManagementURL != nil {
		in, out := &in.DataManagementURL, &out.DataManagementURL
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.QidoURL != nil {
		in, out := &in.QidoURL, &out.QidoURL
		*out = new(string)
		**out = **in
	}
	if in.StowURL != nil {
		in, out := &in.StowURL, &out.StowURL
		*out = new(string)
		**out = **in
	}
	if in.WadoURL != nil {
		in, out := &in.WadoURL, &out.WadoURL
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StoreConfigObservation.
func (in *StoreConfigObservation) DeepCopy() *StoreConfigObservation {
	if in == nil {
		return nil
	}
	out := new(StoreConfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StoreConfigParameters) DeepCopyInto(out *StoreConfigParameters) {
	*out = *in
	if in.CdrServiceAccount != nil {
		in, out := &in.CdrServiceAccount, &out.CdrServiceAccount
		*out = make([]CdrServiceAccountParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ConfigURL != nil {
		in, out := &in.ConfigURL, &out.ConfigURL
		*out = new(string)
		**out = **in
	}
	if in.FHIRStore != nil {
		in, out := &in.FHIRStore, &out.FHIRStore
		*out = make([]FHIRStoreParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.OrganizationID != nil {
		in, out := &in.OrganizationID, &out.OrganizationID
		*out = new(string)
		**out = **in
	}
	if in.OrganizationIDSelector != nil {
		in, out := &in.OrganizationIDSelector, &out.OrganizationIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.OrganizationRef != nil {
		in, out := &in.OrganizationRef, &out.OrganizationRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StoreConfigParameters.
func (in *StoreConfigParameters) DeepCopy() *StoreConfigParameters {
	if in == nil {
		return nil
	}
	out := new(StoreConfigParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StoreConfigSpec) DeepCopyInto(out *StoreConfigSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StoreConfigSpec.
func (in *StoreConfigSpec) DeepCopy() *StoreConfigSpec {
	if in == nil {
		return nil
	}
	out := new(StoreConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StoreConfigStatus) DeepCopyInto(out *StoreConfigStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StoreConfigStatus.
func (in *StoreConfigStatus) DeepCopy() *StoreConfigStatus {
	if in == nil {
		return nil
	}
	out := new(StoreConfigStatus)
	in.DeepCopyInto(out)
	return out
}
