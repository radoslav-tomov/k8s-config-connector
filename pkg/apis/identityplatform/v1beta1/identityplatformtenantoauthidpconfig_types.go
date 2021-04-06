// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Config Connector and manual
//     changes will be clobbered when the file is regenerated.
//
// ----------------------------------------------------------------------------

// *** DISCLAIMER ***
// Config Connector's go-client for CRDs is currently in ALPHA, which means
// that future versions of the go-client may include breaking changes.
// Please try it out and give us feedback!

package v1beta1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type TenantoauthidpconfigClientSecret struct {
	/* Value of the field. Cannot be used if 'valueFrom' is specified. */
	Value string `json:"value,omitempty"`
	/* Source for the field's value. Cannot be used if 'value' is specified. */
	ValueFrom TenantoauthidpconfigValueFrom `json:"valueFrom,omitempty"`
}

type TenantoauthidpconfigResponseType struct {
	/* If true, authorization code is returned from IdP's authorization endpoint. */
	Code bool `json:"code,omitempty"`
	/* If true, ID token is returned from IdP's authorization endpoint. */
	IdToken bool `json:"idToken,omitempty"`
	/* If true, access token is returned from IdP's authorization endpoint. */
	Token bool `json:"token,omitempty"`
}

type TenantoauthidpconfigValueFrom struct {
	/* Reference to a value with the given key in the given Secret in the resource's namespace. */
	SecretKeyRef v1alpha1.ResourceRef `json:"secretKeyRef,omitempty"`
}

type IdentityPlatformTenantOAuthIDPConfigSpec struct {
	/* The client id of an OAuth client. */
	ClientId string `json:"clientId,omitempty"`
	/* The client secret of the OAuth client, to enable OIDC code flow. */
	ClientSecret TenantoauthidpconfigClientSecret `json:"clientSecret,omitempty"`
	/* The config's display name set by developers. */
	DisplayName string `json:"displayName,omitempty"`
	/* True if allows the user to sign in with the provider. */
	Enabled bool `json:"enabled,omitempty"`
	/* For OIDC Idps, the issuer identifier. */
	Issuer string `json:"issuer,omitempty"`
	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	ResourceID string `json:"resourceID,omitempty"`
	/* The multiple response type to request for in the OAuth authorization flow. This can possibly be a combination of set bits (e.g. {id\_token, token}). */
	ResponseType TenantoauthidpconfigResponseType `json:"responseType,omitempty"`
	/*  */
	TenantRef v1alpha1.ResourceRef `json:"tenantRef,omitempty"`
}

type IdentityPlatformTenantOAuthIDPConfigStatus struct {
	/* Conditions represent the latest available observations of the
	   IdentityPlatformTenantOAuthIDPConfig's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// IdentityPlatformTenantOAuthIDPConfig is the Schema for the identityplatform API
// +k8s:openapi-gen=true
type IdentityPlatformTenantOAuthIDPConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   IdentityPlatformTenantOAuthIDPConfigSpec   `json:"spec,omitempty"`
	Status IdentityPlatformTenantOAuthIDPConfigStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// IdentityPlatformTenantOAuthIDPConfigList contains a list of IdentityPlatformTenantOAuthIDPConfig
type IdentityPlatformTenantOAuthIDPConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IdentityPlatformTenantOAuthIDPConfig `json:"items"`
}

func init() {
	SchemeBuilder.Register(&IdentityPlatformTenantOAuthIDPConfig{}, &IdentityPlatformTenantOAuthIDPConfigList{})
}