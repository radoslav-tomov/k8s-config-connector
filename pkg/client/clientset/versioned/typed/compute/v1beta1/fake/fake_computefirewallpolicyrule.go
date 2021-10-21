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

	v1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/compute/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeComputeFirewallPolicyRules implements ComputeFirewallPolicyRuleInterface
type FakeComputeFirewallPolicyRules struct {
	Fake *FakeComputeV1beta1
	ns   string
}

var computefirewallpolicyrulesResource = schema.GroupVersionResource{Group: "compute.cnrm.cloud.google.com", Version: "v1beta1", Resource: "computefirewallpolicyrules"}

var computefirewallpolicyrulesKind = schema.GroupVersionKind{Group: "compute.cnrm.cloud.google.com", Version: "v1beta1", Kind: "ComputeFirewallPolicyRule"}

// Get takes name of the computeFirewallPolicyRule, and returns the corresponding computeFirewallPolicyRule object, and an error if there is any.
func (c *FakeComputeFirewallPolicyRules) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.ComputeFirewallPolicyRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(computefirewallpolicyrulesResource, c.ns, name), &v1beta1.ComputeFirewallPolicyRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ComputeFirewallPolicyRule), err
}

// List takes label and field selectors, and returns the list of ComputeFirewallPolicyRules that match those selectors.
func (c *FakeComputeFirewallPolicyRules) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.ComputeFirewallPolicyRuleList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(computefirewallpolicyrulesResource, computefirewallpolicyrulesKind, c.ns, opts), &v1beta1.ComputeFirewallPolicyRuleList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.ComputeFirewallPolicyRuleList{ListMeta: obj.(*v1beta1.ComputeFirewallPolicyRuleList).ListMeta}
	for _, item := range obj.(*v1beta1.ComputeFirewallPolicyRuleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested computeFirewallPolicyRules.
func (c *FakeComputeFirewallPolicyRules) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(computefirewallpolicyrulesResource, c.ns, opts))

}

// Create takes the representation of a computeFirewallPolicyRule and creates it.  Returns the server's representation of the computeFirewallPolicyRule, and an error, if there is any.
func (c *FakeComputeFirewallPolicyRules) Create(ctx context.Context, computeFirewallPolicyRule *v1beta1.ComputeFirewallPolicyRule, opts v1.CreateOptions) (result *v1beta1.ComputeFirewallPolicyRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(computefirewallpolicyrulesResource, c.ns, computeFirewallPolicyRule), &v1beta1.ComputeFirewallPolicyRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ComputeFirewallPolicyRule), err
}

// Update takes the representation of a computeFirewallPolicyRule and updates it. Returns the server's representation of the computeFirewallPolicyRule, and an error, if there is any.
func (c *FakeComputeFirewallPolicyRules) Update(ctx context.Context, computeFirewallPolicyRule *v1beta1.ComputeFirewallPolicyRule, opts v1.UpdateOptions) (result *v1beta1.ComputeFirewallPolicyRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(computefirewallpolicyrulesResource, c.ns, computeFirewallPolicyRule), &v1beta1.ComputeFirewallPolicyRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ComputeFirewallPolicyRule), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeComputeFirewallPolicyRules) UpdateStatus(ctx context.Context, computeFirewallPolicyRule *v1beta1.ComputeFirewallPolicyRule, opts v1.UpdateOptions) (*v1beta1.ComputeFirewallPolicyRule, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(computefirewallpolicyrulesResource, "status", c.ns, computeFirewallPolicyRule), &v1beta1.ComputeFirewallPolicyRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ComputeFirewallPolicyRule), err
}

// Delete takes name of the computeFirewallPolicyRule and deletes it. Returns an error if one occurs.
func (c *FakeComputeFirewallPolicyRules) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(computefirewallpolicyrulesResource, c.ns, name), &v1beta1.ComputeFirewallPolicyRule{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeComputeFirewallPolicyRules) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(computefirewallpolicyrulesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.ComputeFirewallPolicyRuleList{})
	return err
}

// Patch applies the patch and returns the patched computeFirewallPolicyRule.
func (c *FakeComputeFirewallPolicyRules) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.ComputeFirewallPolicyRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(computefirewallpolicyrulesResource, c.ns, name, pt, data, subresources...), &v1beta1.ComputeFirewallPolicyRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ComputeFirewallPolicyRule), err
}
