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
func (in *CapoolAdditionalExtensions) DeepCopyInto(out *CapoolAdditionalExtensions) {
	*out = *in
	if in.ObjectIdPath != nil {
		in, out := &in.ObjectIdPath, &out.ObjectIdPath
		*out = make([]int, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CapoolAdditionalExtensions.
func (in *CapoolAdditionalExtensions) DeepCopy() *CapoolAdditionalExtensions {
	if in == nil {
		return nil
	}
	out := new(CapoolAdditionalExtensions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CapoolAllowedIssuanceModes) DeepCopyInto(out *CapoolAllowedIssuanceModes) {
	*out = *in
	if in.AllowConfigBasedIssuance != nil {
		in, out := &in.AllowConfigBasedIssuance, &out.AllowConfigBasedIssuance
		*out = new(bool)
		**out = **in
	}
	if in.AllowCsrBasedIssuance != nil {
		in, out := &in.AllowCsrBasedIssuance, &out.AllowCsrBasedIssuance
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CapoolAllowedIssuanceModes.
func (in *CapoolAllowedIssuanceModes) DeepCopy() *CapoolAllowedIssuanceModes {
	if in == nil {
		return nil
	}
	out := new(CapoolAllowedIssuanceModes)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CapoolAllowedKeyTypes) DeepCopyInto(out *CapoolAllowedKeyTypes) {
	*out = *in
	if in.EllipticCurve != nil {
		in, out := &in.EllipticCurve, &out.EllipticCurve
		*out = new(CapoolEllipticCurve)
		(*in).DeepCopyInto(*out)
	}
	if in.Rsa != nil {
		in, out := &in.Rsa, &out.Rsa
		*out = new(CapoolRsa)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CapoolAllowedKeyTypes.
func (in *CapoolAllowedKeyTypes) DeepCopy() *CapoolAllowedKeyTypes {
	if in == nil {
		return nil
	}
	out := new(CapoolAllowedKeyTypes)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CapoolBaseKeyUsage) DeepCopyInto(out *CapoolBaseKeyUsage) {
	*out = *in
	if in.CertSign != nil {
		in, out := &in.CertSign, &out.CertSign
		*out = new(bool)
		**out = **in
	}
	if in.ContentCommitment != nil {
		in, out := &in.ContentCommitment, &out.ContentCommitment
		*out = new(bool)
		**out = **in
	}
	if in.CrlSign != nil {
		in, out := &in.CrlSign, &out.CrlSign
		*out = new(bool)
		**out = **in
	}
	if in.DataEncipherment != nil {
		in, out := &in.DataEncipherment, &out.DataEncipherment
		*out = new(bool)
		**out = **in
	}
	if in.DecipherOnly != nil {
		in, out := &in.DecipherOnly, &out.DecipherOnly
		*out = new(bool)
		**out = **in
	}
	if in.DigitalSignature != nil {
		in, out := &in.DigitalSignature, &out.DigitalSignature
		*out = new(bool)
		**out = **in
	}
	if in.EncipherOnly != nil {
		in, out := &in.EncipherOnly, &out.EncipherOnly
		*out = new(bool)
		**out = **in
	}
	if in.KeyAgreement != nil {
		in, out := &in.KeyAgreement, &out.KeyAgreement
		*out = new(bool)
		**out = **in
	}
	if in.KeyEncipherment != nil {
		in, out := &in.KeyEncipherment, &out.KeyEncipherment
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CapoolBaseKeyUsage.
func (in *CapoolBaseKeyUsage) DeepCopy() *CapoolBaseKeyUsage {
	if in == nil {
		return nil
	}
	out := new(CapoolBaseKeyUsage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CapoolBaselineValues) DeepCopyInto(out *CapoolBaselineValues) {
	*out = *in
	if in.AdditionalExtensions != nil {
		in, out := &in.AdditionalExtensions, &out.AdditionalExtensions
		*out = make([]CapoolAdditionalExtensions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.AiaOcspServers != nil {
		in, out := &in.AiaOcspServers, &out.AiaOcspServers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.CaOptions != nil {
		in, out := &in.CaOptions, &out.CaOptions
		*out = new(CapoolCaOptions)
		(*in).DeepCopyInto(*out)
	}
	if in.KeyUsage != nil {
		in, out := &in.KeyUsage, &out.KeyUsage
		*out = new(CapoolKeyUsage)
		(*in).DeepCopyInto(*out)
	}
	if in.PolicyIds != nil {
		in, out := &in.PolicyIds, &out.PolicyIds
		*out = make([]CapoolPolicyIds, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CapoolBaselineValues.
func (in *CapoolBaselineValues) DeepCopy() *CapoolBaselineValues {
	if in == nil {
		return nil
	}
	out := new(CapoolBaselineValues)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CapoolCaOptions) DeepCopyInto(out *CapoolCaOptions) {
	*out = *in
	if in.IsCa != nil {
		in, out := &in.IsCa, &out.IsCa
		*out = new(bool)
		**out = **in
	}
	if in.MaxIssuerPathLength != nil {
		in, out := &in.MaxIssuerPathLength, &out.MaxIssuerPathLength
		*out = new(int)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CapoolCaOptions.
func (in *CapoolCaOptions) DeepCopy() *CapoolCaOptions {
	if in == nil {
		return nil
	}
	out := new(CapoolCaOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CapoolCelExpression) DeepCopyInto(out *CapoolCelExpression) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Expression != nil {
		in, out := &in.Expression, &out.Expression
		*out = new(string)
		**out = **in
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.Title != nil {
		in, out := &in.Title, &out.Title
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CapoolCelExpression.
func (in *CapoolCelExpression) DeepCopy() *CapoolCelExpression {
	if in == nil {
		return nil
	}
	out := new(CapoolCelExpression)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CapoolEllipticCurve) DeepCopyInto(out *CapoolEllipticCurve) {
	*out = *in
	if in.SignatureAlgorithm != nil {
		in, out := &in.SignatureAlgorithm, &out.SignatureAlgorithm
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CapoolEllipticCurve.
func (in *CapoolEllipticCurve) DeepCopy() *CapoolEllipticCurve {
	if in == nil {
		return nil
	}
	out := new(CapoolEllipticCurve)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CapoolExtendedKeyUsage) DeepCopyInto(out *CapoolExtendedKeyUsage) {
	*out = *in
	if in.ClientAuth != nil {
		in, out := &in.ClientAuth, &out.ClientAuth
		*out = new(bool)
		**out = **in
	}
	if in.CodeSigning != nil {
		in, out := &in.CodeSigning, &out.CodeSigning
		*out = new(bool)
		**out = **in
	}
	if in.EmailProtection != nil {
		in, out := &in.EmailProtection, &out.EmailProtection
		*out = new(bool)
		**out = **in
	}
	if in.OcspSigning != nil {
		in, out := &in.OcspSigning, &out.OcspSigning
		*out = new(bool)
		**out = **in
	}
	if in.ServerAuth != nil {
		in, out := &in.ServerAuth, &out.ServerAuth
		*out = new(bool)
		**out = **in
	}
	if in.TimeStamping != nil {
		in, out := &in.TimeStamping, &out.TimeStamping
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CapoolExtendedKeyUsage.
func (in *CapoolExtendedKeyUsage) DeepCopy() *CapoolExtendedKeyUsage {
	if in == nil {
		return nil
	}
	out := new(CapoolExtendedKeyUsage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CapoolIdentityConstraints) DeepCopyInto(out *CapoolIdentityConstraints) {
	*out = *in
	if in.CelExpression != nil {
		in, out := &in.CelExpression, &out.CelExpression
		*out = new(CapoolCelExpression)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CapoolIdentityConstraints.
func (in *CapoolIdentityConstraints) DeepCopy() *CapoolIdentityConstraints {
	if in == nil {
		return nil
	}
	out := new(CapoolIdentityConstraints)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CapoolIssuancePolicy) DeepCopyInto(out *CapoolIssuancePolicy) {
	*out = *in
	if in.AllowedIssuanceModes != nil {
		in, out := &in.AllowedIssuanceModes, &out.AllowedIssuanceModes
		*out = new(CapoolAllowedIssuanceModes)
		(*in).DeepCopyInto(*out)
	}
	if in.AllowedKeyTypes != nil {
		in, out := &in.AllowedKeyTypes, &out.AllowedKeyTypes
		*out = make([]CapoolAllowedKeyTypes, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.BaselineValues != nil {
		in, out := &in.BaselineValues, &out.BaselineValues
		*out = new(CapoolBaselineValues)
		(*in).DeepCopyInto(*out)
	}
	if in.IdentityConstraints != nil {
		in, out := &in.IdentityConstraints, &out.IdentityConstraints
		*out = new(CapoolIdentityConstraints)
		(*in).DeepCopyInto(*out)
	}
	if in.MaximumLifetime != nil {
		in, out := &in.MaximumLifetime, &out.MaximumLifetime
		*out = new(string)
		**out = **in
	}
	if in.PassthroughExtensions != nil {
		in, out := &in.PassthroughExtensions, &out.PassthroughExtensions
		*out = new(CapoolPassthroughExtensions)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CapoolIssuancePolicy.
func (in *CapoolIssuancePolicy) DeepCopy() *CapoolIssuancePolicy {
	if in == nil {
		return nil
	}
	out := new(CapoolIssuancePolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CapoolKeyUsage) DeepCopyInto(out *CapoolKeyUsage) {
	*out = *in
	if in.BaseKeyUsage != nil {
		in, out := &in.BaseKeyUsage, &out.BaseKeyUsage
		*out = new(CapoolBaseKeyUsage)
		(*in).DeepCopyInto(*out)
	}
	if in.ExtendedKeyUsage != nil {
		in, out := &in.ExtendedKeyUsage, &out.ExtendedKeyUsage
		*out = new(CapoolExtendedKeyUsage)
		(*in).DeepCopyInto(*out)
	}
	if in.UnknownExtendedKeyUsages != nil {
		in, out := &in.UnknownExtendedKeyUsages, &out.UnknownExtendedKeyUsages
		*out = make([]CapoolUnknownExtendedKeyUsages, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CapoolKeyUsage.
func (in *CapoolKeyUsage) DeepCopy() *CapoolKeyUsage {
	if in == nil {
		return nil
	}
	out := new(CapoolKeyUsage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CapoolObjectId) DeepCopyInto(out *CapoolObjectId) {
	*out = *in
	if in.ObjectIdPath != nil {
		in, out := &in.ObjectIdPath, &out.ObjectIdPath
		*out = make([]int, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CapoolObjectId.
func (in *CapoolObjectId) DeepCopy() *CapoolObjectId {
	if in == nil {
		return nil
	}
	out := new(CapoolObjectId)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CapoolPassthroughExtensions) DeepCopyInto(out *CapoolPassthroughExtensions) {
	*out = *in
	if in.AdditionalExtensions != nil {
		in, out := &in.AdditionalExtensions, &out.AdditionalExtensions
		*out = make([]CapoolAdditionalExtensions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.KnownExtensions != nil {
		in, out := &in.KnownExtensions, &out.KnownExtensions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CapoolPassthroughExtensions.
func (in *CapoolPassthroughExtensions) DeepCopy() *CapoolPassthroughExtensions {
	if in == nil {
		return nil
	}
	out := new(CapoolPassthroughExtensions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CapoolPolicyIds) DeepCopyInto(out *CapoolPolicyIds) {
	*out = *in
	if in.ObjectIdPath != nil {
		in, out := &in.ObjectIdPath, &out.ObjectIdPath
		*out = make([]int, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CapoolPolicyIds.
func (in *CapoolPolicyIds) DeepCopy() *CapoolPolicyIds {
	if in == nil {
		return nil
	}
	out := new(CapoolPolicyIds)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CapoolPublishingOptions) DeepCopyInto(out *CapoolPublishingOptions) {
	*out = *in
	if in.PublishCaCert != nil {
		in, out := &in.PublishCaCert, &out.PublishCaCert
		*out = new(bool)
		**out = **in
	}
	if in.PublishCrl != nil {
		in, out := &in.PublishCrl, &out.PublishCrl
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CapoolPublishingOptions.
func (in *CapoolPublishingOptions) DeepCopy() *CapoolPublishingOptions {
	if in == nil {
		return nil
	}
	out := new(CapoolPublishingOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CapoolRsa) DeepCopyInto(out *CapoolRsa) {
	*out = *in
	if in.MaxModulusSize != nil {
		in, out := &in.MaxModulusSize, &out.MaxModulusSize
		*out = new(int)
		**out = **in
	}
	if in.MinModulusSize != nil {
		in, out := &in.MinModulusSize, &out.MinModulusSize
		*out = new(int)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CapoolRsa.
func (in *CapoolRsa) DeepCopy() *CapoolRsa {
	if in == nil {
		return nil
	}
	out := new(CapoolRsa)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CapoolUnknownExtendedKeyUsages) DeepCopyInto(out *CapoolUnknownExtendedKeyUsages) {
	*out = *in
	if in.ObjectIdPath != nil {
		in, out := &in.ObjectIdPath, &out.ObjectIdPath
		*out = make([]int, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CapoolUnknownExtendedKeyUsages.
func (in *CapoolUnknownExtendedKeyUsages) DeepCopy() *CapoolUnknownExtendedKeyUsages {
	if in == nil {
		return nil
	}
	out := new(CapoolUnknownExtendedKeyUsages)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificatetemplateAdditionalExtensions) DeepCopyInto(out *CertificatetemplateAdditionalExtensions) {
	*out = *in
	if in.Critical != nil {
		in, out := &in.Critical, &out.Critical
		*out = new(bool)
		**out = **in
	}
	in.ObjectId.DeepCopyInto(&out.ObjectId)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificatetemplateAdditionalExtensions.
func (in *CertificatetemplateAdditionalExtensions) DeepCopy() *CertificatetemplateAdditionalExtensions {
	if in == nil {
		return nil
	}
	out := new(CertificatetemplateAdditionalExtensions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificatetemplateBaseKeyUsage) DeepCopyInto(out *CertificatetemplateBaseKeyUsage) {
	*out = *in
	if in.CertSign != nil {
		in, out := &in.CertSign, &out.CertSign
		*out = new(bool)
		**out = **in
	}
	if in.ContentCommitment != nil {
		in, out := &in.ContentCommitment, &out.ContentCommitment
		*out = new(bool)
		**out = **in
	}
	if in.CrlSign != nil {
		in, out := &in.CrlSign, &out.CrlSign
		*out = new(bool)
		**out = **in
	}
	if in.DataEncipherment != nil {
		in, out := &in.DataEncipherment, &out.DataEncipherment
		*out = new(bool)
		**out = **in
	}
	if in.DecipherOnly != nil {
		in, out := &in.DecipherOnly, &out.DecipherOnly
		*out = new(bool)
		**out = **in
	}
	if in.DigitalSignature != nil {
		in, out := &in.DigitalSignature, &out.DigitalSignature
		*out = new(bool)
		**out = **in
	}
	if in.EncipherOnly != nil {
		in, out := &in.EncipherOnly, &out.EncipherOnly
		*out = new(bool)
		**out = **in
	}
	if in.KeyAgreement != nil {
		in, out := &in.KeyAgreement, &out.KeyAgreement
		*out = new(bool)
		**out = **in
	}
	if in.KeyEncipherment != nil {
		in, out := &in.KeyEncipherment, &out.KeyEncipherment
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificatetemplateBaseKeyUsage.
func (in *CertificatetemplateBaseKeyUsage) DeepCopy() *CertificatetemplateBaseKeyUsage {
	if in == nil {
		return nil
	}
	out := new(CertificatetemplateBaseKeyUsage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificatetemplateCaOptions) DeepCopyInto(out *CertificatetemplateCaOptions) {
	*out = *in
	if in.IsCa != nil {
		in, out := &in.IsCa, &out.IsCa
		*out = new(bool)
		**out = **in
	}
	if in.MaxIssuerPathLength != nil {
		in, out := &in.MaxIssuerPathLength, &out.MaxIssuerPathLength
		*out = new(int)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificatetemplateCaOptions.
func (in *CertificatetemplateCaOptions) DeepCopy() *CertificatetemplateCaOptions {
	if in == nil {
		return nil
	}
	out := new(CertificatetemplateCaOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificatetemplateCelExpression) DeepCopyInto(out *CertificatetemplateCelExpression) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Expression != nil {
		in, out := &in.Expression, &out.Expression
		*out = new(string)
		**out = **in
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.Title != nil {
		in, out := &in.Title, &out.Title
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificatetemplateCelExpression.
func (in *CertificatetemplateCelExpression) DeepCopy() *CertificatetemplateCelExpression {
	if in == nil {
		return nil
	}
	out := new(CertificatetemplateCelExpression)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificatetemplateExtendedKeyUsage) DeepCopyInto(out *CertificatetemplateExtendedKeyUsage) {
	*out = *in
	if in.ClientAuth != nil {
		in, out := &in.ClientAuth, &out.ClientAuth
		*out = new(bool)
		**out = **in
	}
	if in.CodeSigning != nil {
		in, out := &in.CodeSigning, &out.CodeSigning
		*out = new(bool)
		**out = **in
	}
	if in.EmailProtection != nil {
		in, out := &in.EmailProtection, &out.EmailProtection
		*out = new(bool)
		**out = **in
	}
	if in.OcspSigning != nil {
		in, out := &in.OcspSigning, &out.OcspSigning
		*out = new(bool)
		**out = **in
	}
	if in.ServerAuth != nil {
		in, out := &in.ServerAuth, &out.ServerAuth
		*out = new(bool)
		**out = **in
	}
	if in.TimeStamping != nil {
		in, out := &in.TimeStamping, &out.TimeStamping
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificatetemplateExtendedKeyUsage.
func (in *CertificatetemplateExtendedKeyUsage) DeepCopy() *CertificatetemplateExtendedKeyUsage {
	if in == nil {
		return nil
	}
	out := new(CertificatetemplateExtendedKeyUsage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificatetemplateIdentityConstraints) DeepCopyInto(out *CertificatetemplateIdentityConstraints) {
	*out = *in
	if in.CelExpression != nil {
		in, out := &in.CelExpression, &out.CelExpression
		*out = new(CertificatetemplateCelExpression)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificatetemplateIdentityConstraints.
func (in *CertificatetemplateIdentityConstraints) DeepCopy() *CertificatetemplateIdentityConstraints {
	if in == nil {
		return nil
	}
	out := new(CertificatetemplateIdentityConstraints)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificatetemplateKeyUsage) DeepCopyInto(out *CertificatetemplateKeyUsage) {
	*out = *in
	if in.BaseKeyUsage != nil {
		in, out := &in.BaseKeyUsage, &out.BaseKeyUsage
		*out = new(CertificatetemplateBaseKeyUsage)
		(*in).DeepCopyInto(*out)
	}
	if in.ExtendedKeyUsage != nil {
		in, out := &in.ExtendedKeyUsage, &out.ExtendedKeyUsage
		*out = new(CertificatetemplateExtendedKeyUsage)
		(*in).DeepCopyInto(*out)
	}
	if in.UnknownExtendedKeyUsages != nil {
		in, out := &in.UnknownExtendedKeyUsages, &out.UnknownExtendedKeyUsages
		*out = make([]CertificatetemplateUnknownExtendedKeyUsages, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificatetemplateKeyUsage.
func (in *CertificatetemplateKeyUsage) DeepCopy() *CertificatetemplateKeyUsage {
	if in == nil {
		return nil
	}
	out := new(CertificatetemplateKeyUsage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificatetemplateObjectId) DeepCopyInto(out *CertificatetemplateObjectId) {
	*out = *in
	if in.ObjectIdPath != nil {
		in, out := &in.ObjectIdPath, &out.ObjectIdPath
		*out = make([]int, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificatetemplateObjectId.
func (in *CertificatetemplateObjectId) DeepCopy() *CertificatetemplateObjectId {
	if in == nil {
		return nil
	}
	out := new(CertificatetemplateObjectId)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificatetemplatePassthroughExtensions) DeepCopyInto(out *CertificatetemplatePassthroughExtensions) {
	*out = *in
	if in.AdditionalExtensions != nil {
		in, out := &in.AdditionalExtensions, &out.AdditionalExtensions
		*out = make([]CertificatetemplateAdditionalExtensions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.KnownExtensions != nil {
		in, out := &in.KnownExtensions, &out.KnownExtensions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificatetemplatePassthroughExtensions.
func (in *CertificatetemplatePassthroughExtensions) DeepCopy() *CertificatetemplatePassthroughExtensions {
	if in == nil {
		return nil
	}
	out := new(CertificatetemplatePassthroughExtensions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificatetemplatePolicyIds) DeepCopyInto(out *CertificatetemplatePolicyIds) {
	*out = *in
	if in.ObjectIdPath != nil {
		in, out := &in.ObjectIdPath, &out.ObjectIdPath
		*out = make([]int, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificatetemplatePolicyIds.
func (in *CertificatetemplatePolicyIds) DeepCopy() *CertificatetemplatePolicyIds {
	if in == nil {
		return nil
	}
	out := new(CertificatetemplatePolicyIds)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificatetemplatePredefinedValues) DeepCopyInto(out *CertificatetemplatePredefinedValues) {
	*out = *in
	if in.AdditionalExtensions != nil {
		in, out := &in.AdditionalExtensions, &out.AdditionalExtensions
		*out = make([]CertificatetemplateAdditionalExtensions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.AiaOcspServers != nil {
		in, out := &in.AiaOcspServers, &out.AiaOcspServers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.CaOptions != nil {
		in, out := &in.CaOptions, &out.CaOptions
		*out = new(CertificatetemplateCaOptions)
		(*in).DeepCopyInto(*out)
	}
	if in.KeyUsage != nil {
		in, out := &in.KeyUsage, &out.KeyUsage
		*out = new(CertificatetemplateKeyUsage)
		(*in).DeepCopyInto(*out)
	}
	if in.PolicyIds != nil {
		in, out := &in.PolicyIds, &out.PolicyIds
		*out = make([]CertificatetemplatePolicyIds, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificatetemplatePredefinedValues.
func (in *CertificatetemplatePredefinedValues) DeepCopy() *CertificatetemplatePredefinedValues {
	if in == nil {
		return nil
	}
	out := new(CertificatetemplatePredefinedValues)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificatetemplateUnknownExtendedKeyUsages) DeepCopyInto(out *CertificatetemplateUnknownExtendedKeyUsages) {
	*out = *in
	if in.ObjectIdPath != nil {
		in, out := &in.ObjectIdPath, &out.ObjectIdPath
		*out = make([]int, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificatetemplateUnknownExtendedKeyUsages.
func (in *CertificatetemplateUnknownExtendedKeyUsages) DeepCopy() *CertificatetemplateUnknownExtendedKeyUsages {
	if in == nil {
		return nil
	}
	out := new(CertificatetemplateUnknownExtendedKeyUsages)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrivateCACAPool) DeepCopyInto(out *PrivateCACAPool) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrivateCACAPool.
func (in *PrivateCACAPool) DeepCopy() *PrivateCACAPool {
	if in == nil {
		return nil
	}
	out := new(PrivateCACAPool)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PrivateCACAPool) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrivateCACAPoolList) DeepCopyInto(out *PrivateCACAPoolList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PrivateCACAPool, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrivateCACAPoolList.
func (in *PrivateCACAPoolList) DeepCopy() *PrivateCACAPoolList {
	if in == nil {
		return nil
	}
	out := new(PrivateCACAPoolList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PrivateCACAPoolList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrivateCACAPoolSpec) DeepCopyInto(out *PrivateCACAPoolSpec) {
	*out = *in
	if in.IssuancePolicy != nil {
		in, out := &in.IssuancePolicy, &out.IssuancePolicy
		*out = new(CapoolIssuancePolicy)
		(*in).DeepCopyInto(*out)
	}
	out.ProjectRef = in.ProjectRef
	if in.PublishingOptions != nil {
		in, out := &in.PublishingOptions, &out.PublishingOptions
		*out = new(CapoolPublishingOptions)
		(*in).DeepCopyInto(*out)
	}
	if in.ResourceID != nil {
		in, out := &in.ResourceID, &out.ResourceID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrivateCACAPoolSpec.
func (in *PrivateCACAPoolSpec) DeepCopy() *PrivateCACAPoolSpec {
	if in == nil {
		return nil
	}
	out := new(PrivateCACAPoolSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrivateCACAPoolStatus) DeepCopyInto(out *PrivateCACAPoolStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1alpha1.Condition, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrivateCACAPoolStatus.
func (in *PrivateCACAPoolStatus) DeepCopy() *PrivateCACAPoolStatus {
	if in == nil {
		return nil
	}
	out := new(PrivateCACAPoolStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrivateCACertificateTemplate) DeepCopyInto(out *PrivateCACertificateTemplate) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrivateCACertificateTemplate.
func (in *PrivateCACertificateTemplate) DeepCopy() *PrivateCACertificateTemplate {
	if in == nil {
		return nil
	}
	out := new(PrivateCACertificateTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PrivateCACertificateTemplate) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrivateCACertificateTemplateList) DeepCopyInto(out *PrivateCACertificateTemplateList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PrivateCACertificateTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrivateCACertificateTemplateList.
func (in *PrivateCACertificateTemplateList) DeepCopy() *PrivateCACertificateTemplateList {
	if in == nil {
		return nil
	}
	out := new(PrivateCACertificateTemplateList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PrivateCACertificateTemplateList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrivateCACertificateTemplateSpec) DeepCopyInto(out *PrivateCACertificateTemplateSpec) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.IdentityConstraints != nil {
		in, out := &in.IdentityConstraints, &out.IdentityConstraints
		*out = new(CertificatetemplateIdentityConstraints)
		(*in).DeepCopyInto(*out)
	}
	if in.PassthroughExtensions != nil {
		in, out := &in.PassthroughExtensions, &out.PassthroughExtensions
		*out = new(CertificatetemplatePassthroughExtensions)
		(*in).DeepCopyInto(*out)
	}
	if in.PredefinedValues != nil {
		in, out := &in.PredefinedValues, &out.PredefinedValues
		*out = new(CertificatetemplatePredefinedValues)
		(*in).DeepCopyInto(*out)
	}
	out.ProjectRef = in.ProjectRef
	if in.ResourceID != nil {
		in, out := &in.ResourceID, &out.ResourceID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrivateCACertificateTemplateSpec.
func (in *PrivateCACertificateTemplateSpec) DeepCopy() *PrivateCACertificateTemplateSpec {
	if in == nil {
		return nil
	}
	out := new(PrivateCACertificateTemplateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrivateCACertificateTemplateStatus) DeepCopyInto(out *PrivateCACertificateTemplateStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1alpha1.Condition, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrivateCACertificateTemplateStatus.
func (in *PrivateCACertificateTemplateStatus) DeepCopy() *PrivateCACertificateTemplateStatus {
	if in == nil {
		return nil
	}
	out := new(PrivateCACertificateTemplateStatus)
	in.DeepCopyInto(out)
	return out
}
