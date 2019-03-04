// Copyright 2019 Istio Authors
//
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha2 "github.com/f4tq/istio-api-mod/pkg/kube/apis/config/v1alpha2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeAttributeManifests implements AttributeManifestInterface
type FakeAttributeManifests struct {
	Fake *FakeConfigV1alpha2
	ns   string
}

var attributemanifestsResource = schema.GroupVersionResource{Group: "config", Version: "v1alpha2", Resource: "attributemanifests"}

var attributemanifestsKind = schema.GroupVersionKind{Group: "config", Version: "v1alpha2", Kind: "AttributeManifest"}

// Get takes name of the attributeManifest, and returns the corresponding attributeManifest object, and an error if there is any.
func (c *FakeAttributeManifests) Get(name string, options v1.GetOptions) (result *v1alpha2.AttributeManifest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(attributemanifestsResource, c.ns, name), &v1alpha2.AttributeManifest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.AttributeManifest), err
}

// List takes label and field selectors, and returns the list of AttributeManifests that match those selectors.
func (c *FakeAttributeManifests) List(opts v1.ListOptions) (result *v1alpha2.AttributeManifestList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(attributemanifestsResource, attributemanifestsKind, c.ns, opts), &v1alpha2.AttributeManifestList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha2.AttributeManifestList{ListMeta: obj.(*v1alpha2.AttributeManifestList).ListMeta}
	for _, item := range obj.(*v1alpha2.AttributeManifestList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested attributeManifests.
func (c *FakeAttributeManifests) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(attributemanifestsResource, c.ns, opts))

}

// Create takes the representation of a attributeManifest and creates it.  Returns the server's representation of the attributeManifest, and an error, if there is any.
func (c *FakeAttributeManifests) Create(attributeManifest *v1alpha2.AttributeManifest) (result *v1alpha2.AttributeManifest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(attributemanifestsResource, c.ns, attributeManifest), &v1alpha2.AttributeManifest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.AttributeManifest), err
}

// Update takes the representation of a attributeManifest and updates it. Returns the server's representation of the attributeManifest, and an error, if there is any.
func (c *FakeAttributeManifests) Update(attributeManifest *v1alpha2.AttributeManifest) (result *v1alpha2.AttributeManifest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(attributemanifestsResource, c.ns, attributeManifest), &v1alpha2.AttributeManifest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.AttributeManifest), err
}

// Delete takes name of the attributeManifest and deletes it. Returns an error if one occurs.
func (c *FakeAttributeManifests) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(attributemanifestsResource, c.ns, name), &v1alpha2.AttributeManifest{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAttributeManifests) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(attributemanifestsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha2.AttributeManifestList{})
	return err
}

// Patch applies the patch and returns the patched attributeManifest.
func (c *FakeAttributeManifests) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha2.AttributeManifest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(attributemanifestsResource, c.ns, name, data, subresources...), &v1alpha2.AttributeManifest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.AttributeManifest), err
}
