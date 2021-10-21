//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// *** DISCLAIMER ***
// Config Connector's go-client for CRDs is currently in ALPHA, which means
// that future versions of the go-client may include breaking changes.
// Please try it out and give us feedback!

// Code generated by main. DO NOT EDIT.

package v1beta1

import (
	v1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuditconfigAuditLogConfigs) DeepCopyInto(out *AuditconfigAuditLogConfigs) {
	*out = *in
	if in.ExemptedMembers != nil {
		in, out := &in.ExemptedMembers, &out.ExemptedMembers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuditconfigAuditLogConfigs.
func (in *AuditconfigAuditLogConfigs) DeepCopy() *AuditconfigAuditLogConfigs {
	if in == nil {
		return nil
	}
	out := new(AuditconfigAuditLogConfigs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IAMAuditConfig) DeepCopyInto(out *IAMAuditConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IAMAuditConfig.
func (in *IAMAuditConfig) DeepCopy() *IAMAuditConfig {
	if in == nil {
		return nil
	}
	out := new(IAMAuditConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IAMAuditConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IAMAuditConfigList) DeepCopyInto(out *IAMAuditConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IAMAuditConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IAMAuditConfigList.
func (in *IAMAuditConfigList) DeepCopy() *IAMAuditConfigList {
	if in == nil {
		return nil
	}
	out := new(IAMAuditConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IAMAuditConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IAMAuditConfigSpec) DeepCopyInto(out *IAMAuditConfigSpec) {
	*out = *in
	if in.AuditLogConfigs != nil {
		in, out := &in.AuditLogConfigs, &out.AuditLogConfigs
		*out = make([]AuditconfigAuditLogConfigs, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	out.ResourceRef = in.ResourceRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IAMAuditConfigSpec.
func (in *IAMAuditConfigSpec) DeepCopy() *IAMAuditConfigSpec {
	if in == nil {
		return nil
	}
	out := new(IAMAuditConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IAMAuditConfigStatus) DeepCopyInto(out *IAMAuditConfigStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1alpha1.Condition, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IAMAuditConfigStatus.
func (in *IAMAuditConfigStatus) DeepCopy() *IAMAuditConfigStatus {
	if in == nil {
		return nil
	}
	out := new(IAMAuditConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IAMCustomRole) DeepCopyInto(out *IAMCustomRole) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IAMCustomRole.
func (in *IAMCustomRole) DeepCopy() *IAMCustomRole {
	if in == nil {
		return nil
	}
	out := new(IAMCustomRole)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IAMCustomRole) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IAMCustomRoleList) DeepCopyInto(out *IAMCustomRoleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IAMCustomRole, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IAMCustomRoleList.
func (in *IAMCustomRoleList) DeepCopy() *IAMCustomRoleList {
	if in == nil {
		return nil
	}
	out := new(IAMCustomRoleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IAMCustomRoleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IAMCustomRoleSpec) DeepCopyInto(out *IAMCustomRoleSpec) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Permissions != nil {
		in, out := &in.Permissions, &out.Permissions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ResourceID != nil {
		in, out := &in.ResourceID, &out.ResourceID
		*out = new(string)
		**out = **in
	}
	if in.Stage != nil {
		in, out := &in.Stage, &out.Stage
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IAMCustomRoleSpec.
func (in *IAMCustomRoleSpec) DeepCopy() *IAMCustomRoleSpec {
	if in == nil {
		return nil
	}
	out := new(IAMCustomRoleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IAMCustomRoleStatus) DeepCopyInto(out *IAMCustomRoleStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1alpha1.Condition, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IAMCustomRoleStatus.
func (in *IAMCustomRoleStatus) DeepCopy() *IAMCustomRoleStatus {
	if in == nil {
		return nil
	}
	out := new(IAMCustomRoleStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IAMPartialPolicy) DeepCopyInto(out *IAMPartialPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IAMPartialPolicy.
func (in *IAMPartialPolicy) DeepCopy() *IAMPartialPolicy {
	if in == nil {
		return nil
	}
	out := new(IAMPartialPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IAMPartialPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IAMPartialPolicyList) DeepCopyInto(out *IAMPartialPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IAMPartialPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IAMPartialPolicyList.
func (in *IAMPartialPolicyList) DeepCopy() *IAMPartialPolicyList {
	if in == nil {
		return nil
	}
	out := new(IAMPartialPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IAMPartialPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IAMPartialPolicySpec) DeepCopyInto(out *IAMPartialPolicySpec) {
	*out = *in
	if in.Bindings != nil {
		in, out := &in.Bindings, &out.Bindings
		*out = make([]PartialpolicyBindings, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	out.ResourceRef = in.ResourceRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IAMPartialPolicySpec.
func (in *IAMPartialPolicySpec) DeepCopy() *IAMPartialPolicySpec {
	if in == nil {
		return nil
	}
	out := new(IAMPartialPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IAMPartialPolicyStatus) DeepCopyInto(out *IAMPartialPolicyStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1alpha1.Condition, len(*in))
		copy(*out, *in)
	}
	if in.AllBindings != nil {
		in, out := &in.AllBindings, &out.AllBindings
		*out = make([]PartialpolicyAllBindingsStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.LastAppliedBindings != nil {
		in, out := &in.LastAppliedBindings, &out.LastAppliedBindings
		*out = make([]PartialpolicyLastAppliedBindingsStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IAMPartialPolicyStatus.
func (in *IAMPartialPolicyStatus) DeepCopy() *IAMPartialPolicyStatus {
	if in == nil {
		return nil
	}
	out := new(IAMPartialPolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IAMPolicy) DeepCopyInto(out *IAMPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IAMPolicy.
func (in *IAMPolicy) DeepCopy() *IAMPolicy {
	if in == nil {
		return nil
	}
	out := new(IAMPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IAMPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IAMPolicyList) DeepCopyInto(out *IAMPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IAMPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IAMPolicyList.
func (in *IAMPolicyList) DeepCopy() *IAMPolicyList {
	if in == nil {
		return nil
	}
	out := new(IAMPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IAMPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IAMPolicyMember) DeepCopyInto(out *IAMPolicyMember) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IAMPolicyMember.
func (in *IAMPolicyMember) DeepCopy() *IAMPolicyMember {
	if in == nil {
		return nil
	}
	out := new(IAMPolicyMember)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IAMPolicyMember) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IAMPolicyMemberList) DeepCopyInto(out *IAMPolicyMemberList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IAMPolicyMember, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IAMPolicyMemberList.
func (in *IAMPolicyMemberList) DeepCopy() *IAMPolicyMemberList {
	if in == nil {
		return nil
	}
	out := new(IAMPolicyMemberList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IAMPolicyMemberList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IAMPolicyMemberSpec) DeepCopyInto(out *IAMPolicyMemberSpec) {
	*out = *in
	if in.Condition != nil {
		in, out := &in.Condition, &out.Condition
		*out = new(PolicymemberCondition)
		(*in).DeepCopyInto(*out)
	}
	if in.Member != nil {
		in, out := &in.Member, &out.Member
		*out = new(string)
		**out = **in
	}
	if in.MemberFrom != nil {
		in, out := &in.MemberFrom, &out.MemberFrom
		*out = new(PolicymemberMemberFrom)
		(*in).DeepCopyInto(*out)
	}
	out.ResourceRef = in.ResourceRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IAMPolicyMemberSpec.
func (in *IAMPolicyMemberSpec) DeepCopy() *IAMPolicyMemberSpec {
	if in == nil {
		return nil
	}
	out := new(IAMPolicyMemberSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IAMPolicyMemberStatus) DeepCopyInto(out *IAMPolicyMemberStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1alpha1.Condition, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IAMPolicyMemberStatus.
func (in *IAMPolicyMemberStatus) DeepCopy() *IAMPolicyMemberStatus {
	if in == nil {
		return nil
	}
	out := new(IAMPolicyMemberStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IAMPolicySpec) DeepCopyInto(out *IAMPolicySpec) {
	*out = *in
	if in.AuditConfigs != nil {
		in, out := &in.AuditConfigs, &out.AuditConfigs
		*out = make([]PolicyAuditConfigs, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Bindings != nil {
		in, out := &in.Bindings, &out.Bindings
		*out = make([]PolicyBindings, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	out.ResourceRef = in.ResourceRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IAMPolicySpec.
func (in *IAMPolicySpec) DeepCopy() *IAMPolicySpec {
	if in == nil {
		return nil
	}
	out := new(IAMPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IAMPolicyStatus) DeepCopyInto(out *IAMPolicyStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1alpha1.Condition, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IAMPolicyStatus.
func (in *IAMPolicyStatus) DeepCopy() *IAMPolicyStatus {
	if in == nil {
		return nil
	}
	out := new(IAMPolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IAMServiceAccount) DeepCopyInto(out *IAMServiceAccount) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IAMServiceAccount.
func (in *IAMServiceAccount) DeepCopy() *IAMServiceAccount {
	if in == nil {
		return nil
	}
	out := new(IAMServiceAccount)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IAMServiceAccount) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IAMServiceAccountKey) DeepCopyInto(out *IAMServiceAccountKey) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IAMServiceAccountKey.
func (in *IAMServiceAccountKey) DeepCopy() *IAMServiceAccountKey {
	if in == nil {
		return nil
	}
	out := new(IAMServiceAccountKey)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IAMServiceAccountKey) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IAMServiceAccountKeyList) DeepCopyInto(out *IAMServiceAccountKeyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IAMServiceAccountKey, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IAMServiceAccountKeyList.
func (in *IAMServiceAccountKeyList) DeepCopy() *IAMServiceAccountKeyList {
	if in == nil {
		return nil
	}
	out := new(IAMServiceAccountKeyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IAMServiceAccountKeyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IAMServiceAccountKeySpec) DeepCopyInto(out *IAMServiceAccountKeySpec) {
	*out = *in
	if in.KeyAlgorithm != nil {
		in, out := &in.KeyAlgorithm, &out.KeyAlgorithm
		*out = new(string)
		**out = **in
	}
	if in.PrivateKeyType != nil {
		in, out := &in.PrivateKeyType, &out.PrivateKeyType
		*out = new(string)
		**out = **in
	}
	if in.PublicKeyData != nil {
		in, out := &in.PublicKeyData, &out.PublicKeyData
		*out = new(string)
		**out = **in
	}
	if in.PublicKeyType != nil {
		in, out := &in.PublicKeyType, &out.PublicKeyType
		*out = new(string)
		**out = **in
	}
	out.ServiceAccountRef = in.ServiceAccountRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IAMServiceAccountKeySpec.
func (in *IAMServiceAccountKeySpec) DeepCopy() *IAMServiceAccountKeySpec {
	if in == nil {
		return nil
	}
	out := new(IAMServiceAccountKeySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IAMServiceAccountKeyStatus) DeepCopyInto(out *IAMServiceAccountKeyStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1alpha1.Condition, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IAMServiceAccountKeyStatus.
func (in *IAMServiceAccountKeyStatus) DeepCopy() *IAMServiceAccountKeyStatus {
	if in == nil {
		return nil
	}
	out := new(IAMServiceAccountKeyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IAMServiceAccountList) DeepCopyInto(out *IAMServiceAccountList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IAMServiceAccount, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IAMServiceAccountList.
func (in *IAMServiceAccountList) DeepCopy() *IAMServiceAccountList {
	if in == nil {
		return nil
	}
	out := new(IAMServiceAccountList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IAMServiceAccountList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IAMServiceAccountSpec) DeepCopyInto(out *IAMServiceAccountSpec) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Disabled != nil {
		in, out := &in.Disabled, &out.Disabled
		*out = new(bool)
		**out = **in
	}
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.ResourceID != nil {
		in, out := &in.ResourceID, &out.ResourceID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IAMServiceAccountSpec.
func (in *IAMServiceAccountSpec) DeepCopy() *IAMServiceAccountSpec {
	if in == nil {
		return nil
	}
	out := new(IAMServiceAccountSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IAMServiceAccountStatus) DeepCopyInto(out *IAMServiceAccountStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1alpha1.Condition, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IAMServiceAccountStatus.
func (in *IAMServiceAccountStatus) DeepCopy() *IAMServiceAccountStatus {
	if in == nil {
		return nil
	}
	out := new(IAMServiceAccountStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PartialpolicyAllBindingsStatus) DeepCopyInto(out *PartialpolicyAllBindingsStatus) {
	*out = *in
	out.Condition = in.Condition
	if in.Members != nil {
		in, out := &in.Members, &out.Members
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PartialpolicyAllBindingsStatus.
func (in *PartialpolicyAllBindingsStatus) DeepCopy() *PartialpolicyAllBindingsStatus {
	if in == nil {
		return nil
	}
	out := new(PartialpolicyAllBindingsStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PartialpolicyBindings) DeepCopyInto(out *PartialpolicyBindings) {
	*out = *in
	if in.Condition != nil {
		in, out := &in.Condition, &out.Condition
		*out = new(PartialpolicyCondition)
		(*in).DeepCopyInto(*out)
	}
	if in.Members != nil {
		in, out := &in.Members, &out.Members
		*out = make([]PartialpolicyMembers, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PartialpolicyBindings.
func (in *PartialpolicyBindings) DeepCopy() *PartialpolicyBindings {
	if in == nil {
		return nil
	}
	out := new(PartialpolicyBindings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PartialpolicyCondition) DeepCopyInto(out *PartialpolicyCondition) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PartialpolicyCondition.
func (in *PartialpolicyCondition) DeepCopy() *PartialpolicyCondition {
	if in == nil {
		return nil
	}
	out := new(PartialpolicyCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PartialpolicyConditionStatus) DeepCopyInto(out *PartialpolicyConditionStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PartialpolicyConditionStatus.
func (in *PartialpolicyConditionStatus) DeepCopy() *PartialpolicyConditionStatus {
	if in == nil {
		return nil
	}
	out := new(PartialpolicyConditionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PartialpolicyLastAppliedBindingsStatus) DeepCopyInto(out *PartialpolicyLastAppliedBindingsStatus) {
	*out = *in
	out.Condition = in.Condition
	if in.Members != nil {
		in, out := &in.Members, &out.Members
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PartialpolicyLastAppliedBindingsStatus.
func (in *PartialpolicyLastAppliedBindingsStatus) DeepCopy() *PartialpolicyLastAppliedBindingsStatus {
	if in == nil {
		return nil
	}
	out := new(PartialpolicyLastAppliedBindingsStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PartialpolicyMembers) DeepCopyInto(out *PartialpolicyMembers) {
	*out = *in
	if in.Member != nil {
		in, out := &in.Member, &out.Member
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PartialpolicyMembers.
func (in *PartialpolicyMembers) DeepCopy() *PartialpolicyMembers {
	if in == nil {
		return nil
	}
	out := new(PartialpolicyMembers)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyAuditConfigs) DeepCopyInto(out *PolicyAuditConfigs) {
	*out = *in
	if in.AuditLogConfigs != nil {
		in, out := &in.AuditLogConfigs, &out.AuditLogConfigs
		*out = make([]PolicyAuditLogConfigs, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyAuditConfigs.
func (in *PolicyAuditConfigs) DeepCopy() *PolicyAuditConfigs {
	if in == nil {
		return nil
	}
	out := new(PolicyAuditConfigs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyAuditLogConfigs) DeepCopyInto(out *PolicyAuditLogConfigs) {
	*out = *in
	if in.ExemptedMembers != nil {
		in, out := &in.ExemptedMembers, &out.ExemptedMembers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyAuditLogConfigs.
func (in *PolicyAuditLogConfigs) DeepCopy() *PolicyAuditLogConfigs {
	if in == nil {
		return nil
	}
	out := new(PolicyAuditLogConfigs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyBindings) DeepCopyInto(out *PolicyBindings) {
	*out = *in
	if in.Condition != nil {
		in, out := &in.Condition, &out.Condition
		*out = new(PolicyCondition)
		(*in).DeepCopyInto(*out)
	}
	if in.Members != nil {
		in, out := &in.Members, &out.Members
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyBindings.
func (in *PolicyBindings) DeepCopy() *PolicyBindings {
	if in == nil {
		return nil
	}
	out := new(PolicyBindings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyCondition) DeepCopyInto(out *PolicyCondition) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyCondition.
func (in *PolicyCondition) DeepCopy() *PolicyCondition {
	if in == nil {
		return nil
	}
	out := new(PolicyCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicymemberCondition) DeepCopyInto(out *PolicymemberCondition) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicymemberCondition.
func (in *PolicymemberCondition) DeepCopy() *PolicymemberCondition {
	if in == nil {
		return nil
	}
	out := new(PolicymemberCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicymemberMemberFrom) DeepCopyInto(out *PolicymemberMemberFrom) {
	*out = *in
	if in.LogSinkRef != nil {
		in, out := &in.LogSinkRef, &out.LogSinkRef
		*out = new(v1alpha1.IAMResourceRef)
		**out = **in
	}
	if in.ServiceAccountRef != nil {
		in, out := &in.ServiceAccountRef, &out.ServiceAccountRef
		*out = new(v1alpha1.IAMResourceRef)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicymemberMemberFrom.
func (in *PolicymemberMemberFrom) DeepCopy() *PolicymemberMemberFrom {
	if in == nil {
		return nil
	}
	out := new(PolicymemberMemberFrom)
	in.DeepCopyInto(out)
	return out
}
