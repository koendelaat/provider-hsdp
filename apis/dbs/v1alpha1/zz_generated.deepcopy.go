//go:build !ignore_autogenerated

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
func (in *SqsSubscriber) DeepCopyInto(out *SqsSubscriber) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SqsSubscriber.
func (in *SqsSubscriber) DeepCopy() *SqsSubscriber {
	if in == nil {
		return nil
	}
	out := new(SqsSubscriber)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SqsSubscriber) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SqsSubscriberInitParameters) DeepCopyInto(out *SqsSubscriberInitParameters) {
	*out = *in
	if in.DeliveryDelaySeconds != nil {
		in, out := &in.DeliveryDelaySeconds, &out.DeliveryDelaySeconds
		*out = new(int64)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.MessageRetentionPeriodSeconds != nil {
		in, out := &in.MessageRetentionPeriodSeconds, &out.MessageRetentionPeriodSeconds
		*out = new(int64)
		**out = **in
	}
	if in.NameInfix != nil {
		in, out := &in.NameInfix, &out.NameInfix
		*out = new(string)
		**out = **in
	}
	if in.QueueType != nil {
		in, out := &in.QueueType, &out.QueueType
		*out = new(string)
		**out = **in
	}
	if in.ReceiveWaitTimeSeconds != nil {
		in, out := &in.ReceiveWaitTimeSeconds, &out.ReceiveWaitTimeSeconds
		*out = new(int64)
		**out = **in
	}
	if in.ServerSideEncryption != nil {
		in, out := &in.ServerSideEncryption, &out.ServerSideEncryption
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SqsSubscriberInitParameters.
func (in *SqsSubscriberInitParameters) DeepCopy() *SqsSubscriberInitParameters {
	if in == nil {
		return nil
	}
	out := new(SqsSubscriberInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SqsSubscriberList) DeepCopyInto(out *SqsSubscriberList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SqsSubscriber, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SqsSubscriberList.
func (in *SqsSubscriberList) DeepCopy() *SqsSubscriberList {
	if in == nil {
		return nil
	}
	out := new(SqsSubscriberList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SqsSubscriberList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SqsSubscriberObservation) DeepCopyInto(out *SqsSubscriberObservation) {
	*out = *in
	if in.DeliveryDelaySeconds != nil {
		in, out := &in.DeliveryDelaySeconds, &out.DeliveryDelaySeconds
		*out = new(int64)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.MessageRetentionPeriodSeconds != nil {
		in, out := &in.MessageRetentionPeriodSeconds, &out.MessageRetentionPeriodSeconds
		*out = new(int64)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.NameInfix != nil {
		in, out := &in.NameInfix, &out.NameInfix
		*out = new(string)
		**out = **in
	}
	if in.QueueName != nil {
		in, out := &in.QueueName, &out.QueueName
		*out = new(string)
		**out = **in
	}
	if in.QueueType != nil {
		in, out := &in.QueueType, &out.QueueType
		*out = new(string)
		**out = **in
	}
	if in.ReceiveWaitTimeSeconds != nil {
		in, out := &in.ReceiveWaitTimeSeconds, &out.ReceiveWaitTimeSeconds
		*out = new(int64)
		**out = **in
	}
	if in.ServerSideEncryption != nil {
		in, out := &in.ServerSideEncryption, &out.ServerSideEncryption
		*out = new(bool)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SqsSubscriberObservation.
func (in *SqsSubscriberObservation) DeepCopy() *SqsSubscriberObservation {
	if in == nil {
		return nil
	}
	out := new(SqsSubscriberObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SqsSubscriberParameters) DeepCopyInto(out *SqsSubscriberParameters) {
	*out = *in
	if in.DeliveryDelaySeconds != nil {
		in, out := &in.DeliveryDelaySeconds, &out.DeliveryDelaySeconds
		*out = new(int64)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.MessageRetentionPeriodSeconds != nil {
		in, out := &in.MessageRetentionPeriodSeconds, &out.MessageRetentionPeriodSeconds
		*out = new(int64)
		**out = **in
	}
	if in.NameInfix != nil {
		in, out := &in.NameInfix, &out.NameInfix
		*out = new(string)
		**out = **in
	}
	if in.QueueType != nil {
		in, out := &in.QueueType, &out.QueueType
		*out = new(string)
		**out = **in
	}
	if in.ReceiveWaitTimeSeconds != nil {
		in, out := &in.ReceiveWaitTimeSeconds, &out.ReceiveWaitTimeSeconds
		*out = new(int64)
		**out = **in
	}
	if in.ServerSideEncryption != nil {
		in, out := &in.ServerSideEncryption, &out.ServerSideEncryption
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SqsSubscriberParameters.
func (in *SqsSubscriberParameters) DeepCopy() *SqsSubscriberParameters {
	if in == nil {
		return nil
	}
	out := new(SqsSubscriberParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SqsSubscriberSpec) DeepCopyInto(out *SqsSubscriberSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SqsSubscriberSpec.
func (in *SqsSubscriberSpec) DeepCopy() *SqsSubscriberSpec {
	if in == nil {
		return nil
	}
	out := new(SqsSubscriberSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SqsSubscriberStatus) DeepCopyInto(out *SqsSubscriberStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SqsSubscriberStatus.
func (in *SqsSubscriberStatus) DeepCopy() *SqsSubscriberStatus {
	if in == nil {
		return nil
	}
	out := new(SqsSubscriberStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Subscription) DeepCopyInto(out *Subscription) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Subscription.
func (in *Subscription) DeepCopy() *Subscription {
	if in == nil {
		return nil
	}
	out := new(Subscription)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Subscription) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubscriptionInitParameters) DeepCopyInto(out *SubscriptionInitParameters) {
	*out = *in
	if in.DataType != nil {
		in, out := &in.DataType, &out.DataType
		*out = new(string)
		**out = **in
	}
	if in.DataTypeRefRef != nil {
		in, out := &in.DataTypeRefRef, &out.DataTypeRefRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.DataTypeSelector != nil {
		in, out := &in.DataTypeSelector, &out.DataTypeSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.DeliverDataOnly != nil {
		in, out := &in.DeliverDataOnly, &out.DeliverDataOnly
		*out = new(bool)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.KinesisStreamPartitionKey != nil {
		in, out := &in.KinesisStreamPartitionKey, &out.KinesisStreamPartitionKey
		*out = new(string)
		**out = **in
	}
	if in.NameInfix != nil {
		in, out := &in.NameInfix, &out.NameInfix
		*out = new(string)
		**out = **in
	}
	if in.SubscriberID != nil {
		in, out := &in.SubscriberID, &out.SubscriberID
		*out = new(string)
		**out = **in
	}
	if in.SubscriberRef != nil {
		in, out := &in.SubscriberRef, &out.SubscriberRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.SubscriberSelector != nil {
		in, out := &in.SubscriberSelector, &out.SubscriberSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubscriptionInitParameters.
func (in *SubscriptionInitParameters) DeepCopy() *SubscriptionInitParameters {
	if in == nil {
		return nil
	}
	out := new(SubscriptionInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubscriptionList) DeepCopyInto(out *SubscriptionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Subscription, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubscriptionList.
func (in *SubscriptionList) DeepCopy() *SubscriptionList {
	if in == nil {
		return nil
	}
	out := new(SubscriptionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SubscriptionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubscriptionObservation) DeepCopyInto(out *SubscriptionObservation) {
	*out = *in
	if in.DataType != nil {
		in, out := &in.DataType, &out.DataType
		*out = new(string)
		**out = **in
	}
	if in.DeliverDataOnly != nil {
		in, out := &in.DeliverDataOnly, &out.DeliverDataOnly
		*out = new(bool)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.KinesisStreamPartitionKey != nil {
		in, out := &in.KinesisStreamPartitionKey, &out.KinesisStreamPartitionKey
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.NameInfix != nil {
		in, out := &in.NameInfix, &out.NameInfix
		*out = new(string)
		**out = **in
	}
	if in.RuleName != nil {
		in, out := &in.RuleName, &out.RuleName
		*out = new(string)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
	if in.SubscriberID != nil {
		in, out := &in.SubscriberID, &out.SubscriberID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubscriptionObservation.
func (in *SubscriptionObservation) DeepCopy() *SubscriptionObservation {
	if in == nil {
		return nil
	}
	out := new(SubscriptionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubscriptionParameters) DeepCopyInto(out *SubscriptionParameters) {
	*out = *in
	if in.DataType != nil {
		in, out := &in.DataType, &out.DataType
		*out = new(string)
		**out = **in
	}
	if in.DataTypeRefRef != nil {
		in, out := &in.DataTypeRefRef, &out.DataTypeRefRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.DataTypeSelector != nil {
		in, out := &in.DataTypeSelector, &out.DataTypeSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.DeliverDataOnly != nil {
		in, out := &in.DeliverDataOnly, &out.DeliverDataOnly
		*out = new(bool)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.KinesisStreamPartitionKey != nil {
		in, out := &in.KinesisStreamPartitionKey, &out.KinesisStreamPartitionKey
		*out = new(string)
		**out = **in
	}
	if in.NameInfix != nil {
		in, out := &in.NameInfix, &out.NameInfix
		*out = new(string)
		**out = **in
	}
	if in.SubscriberID != nil {
		in, out := &in.SubscriberID, &out.SubscriberID
		*out = new(string)
		**out = **in
	}
	if in.SubscriberRef != nil {
		in, out := &in.SubscriberRef, &out.SubscriberRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.SubscriberSelector != nil {
		in, out := &in.SubscriberSelector, &out.SubscriberSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubscriptionParameters.
func (in *SubscriptionParameters) DeepCopy() *SubscriptionParameters {
	if in == nil {
		return nil
	}
	out := new(SubscriptionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubscriptionSpec) DeepCopyInto(out *SubscriptionSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubscriptionSpec.
func (in *SubscriptionSpec) DeepCopy() *SubscriptionSpec {
	if in == nil {
		return nil
	}
	out := new(SubscriptionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubscriptionStatus) DeepCopyInto(out *SubscriptionStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubscriptionStatus.
func (in *SubscriptionStatus) DeepCopy() *SubscriptionStatus {
	if in == nil {
		return nil
	}
	out := new(SubscriptionStatus)
	in.DeepCopyInto(out)
	return out
}
