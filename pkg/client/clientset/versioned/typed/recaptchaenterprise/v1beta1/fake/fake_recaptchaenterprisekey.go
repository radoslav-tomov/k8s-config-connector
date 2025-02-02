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

package fake

import (
	"context"

	v1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/recaptchaenterprise/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeRecaptchaEnterpriseKeys implements RecaptchaEnterpriseKeyInterface
type FakeRecaptchaEnterpriseKeys struct {
	Fake *FakeRecaptchaenterpriseV1beta1
	ns   string
}

var recaptchaenterprisekeysResource = schema.GroupVersionResource{Group: "recaptchaenterprise.cnrm.cloud.google.com", Version: "v1beta1", Resource: "recaptchaenterprisekeys"}

var recaptchaenterprisekeysKind = schema.GroupVersionKind{Group: "recaptchaenterprise.cnrm.cloud.google.com", Version: "v1beta1", Kind: "RecaptchaEnterpriseKey"}

// Get takes name of the recaptchaEnterpriseKey, and returns the corresponding recaptchaEnterpriseKey object, and an error if there is any.
func (c *FakeRecaptchaEnterpriseKeys) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.RecaptchaEnterpriseKey, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(recaptchaenterprisekeysResource, c.ns, name), &v1beta1.RecaptchaEnterpriseKey{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.RecaptchaEnterpriseKey), err
}

// List takes label and field selectors, and returns the list of RecaptchaEnterpriseKeys that match those selectors.
func (c *FakeRecaptchaEnterpriseKeys) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.RecaptchaEnterpriseKeyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(recaptchaenterprisekeysResource, recaptchaenterprisekeysKind, c.ns, opts), &v1beta1.RecaptchaEnterpriseKeyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.RecaptchaEnterpriseKeyList{ListMeta: obj.(*v1beta1.RecaptchaEnterpriseKeyList).ListMeta}
	for _, item := range obj.(*v1beta1.RecaptchaEnterpriseKeyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested recaptchaEnterpriseKeys.
func (c *FakeRecaptchaEnterpriseKeys) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(recaptchaenterprisekeysResource, c.ns, opts))

}

// Create takes the representation of a recaptchaEnterpriseKey and creates it.  Returns the server's representation of the recaptchaEnterpriseKey, and an error, if there is any.
func (c *FakeRecaptchaEnterpriseKeys) Create(ctx context.Context, recaptchaEnterpriseKey *v1beta1.RecaptchaEnterpriseKey, opts v1.CreateOptions) (result *v1beta1.RecaptchaEnterpriseKey, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(recaptchaenterprisekeysResource, c.ns, recaptchaEnterpriseKey), &v1beta1.RecaptchaEnterpriseKey{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.RecaptchaEnterpriseKey), err
}

// Update takes the representation of a recaptchaEnterpriseKey and updates it. Returns the server's representation of the recaptchaEnterpriseKey, and an error, if there is any.
func (c *FakeRecaptchaEnterpriseKeys) Update(ctx context.Context, recaptchaEnterpriseKey *v1beta1.RecaptchaEnterpriseKey, opts v1.UpdateOptions) (result *v1beta1.RecaptchaEnterpriseKey, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(recaptchaenterprisekeysResource, c.ns, recaptchaEnterpriseKey), &v1beta1.RecaptchaEnterpriseKey{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.RecaptchaEnterpriseKey), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeRecaptchaEnterpriseKeys) UpdateStatus(ctx context.Context, recaptchaEnterpriseKey *v1beta1.RecaptchaEnterpriseKey, opts v1.UpdateOptions) (*v1beta1.RecaptchaEnterpriseKey, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(recaptchaenterprisekeysResource, "status", c.ns, recaptchaEnterpriseKey), &v1beta1.RecaptchaEnterpriseKey{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.RecaptchaEnterpriseKey), err
}

// Delete takes name of the recaptchaEnterpriseKey and deletes it. Returns an error if one occurs.
func (c *FakeRecaptchaEnterpriseKeys) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(recaptchaenterprisekeysResource, c.ns, name), &v1beta1.RecaptchaEnterpriseKey{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeRecaptchaEnterpriseKeys) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(recaptchaenterprisekeysResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.RecaptchaEnterpriseKeyList{})
	return err
}

// Patch applies the patch and returns the patched recaptchaEnterpriseKey.
func (c *FakeRecaptchaEnterpriseKeys) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.RecaptchaEnterpriseKey, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(recaptchaenterprisekeysResource, c.ns, name, pt, data, subresources...), &v1beta1.RecaptchaEnterpriseKey{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.RecaptchaEnterpriseKey), err
}
