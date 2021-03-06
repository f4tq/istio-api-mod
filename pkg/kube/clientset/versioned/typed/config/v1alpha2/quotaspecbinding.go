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

package v1alpha2

import (
	v1alpha2 "github.com/f4tq/istio-api-mod/pkg/kube/apis/config/v1alpha2"
	scheme "github.com/f4tq/istio-api-mod/pkg/kube/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// QuotaSpecBindingsGetter has a method to return a QuotaSpecBindingInterface.
// A group's client should implement this interface.
type QuotaSpecBindingsGetter interface {
	QuotaSpecBindings(namespace string) QuotaSpecBindingInterface
}

// QuotaSpecBindingInterface has methods to work with QuotaSpecBinding resources.
type QuotaSpecBindingInterface interface {
	Create(*v1alpha2.QuotaSpecBinding) (*v1alpha2.QuotaSpecBinding, error)
	Update(*v1alpha2.QuotaSpecBinding) (*v1alpha2.QuotaSpecBinding, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha2.QuotaSpecBinding, error)
	List(opts v1.ListOptions) (*v1alpha2.QuotaSpecBindingList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha2.QuotaSpecBinding, err error)
	QuotaSpecBindingExpansion
}

// quotaSpecBindings implements QuotaSpecBindingInterface
type quotaSpecBindings struct {
	client rest.Interface
	ns     string
}

// newQuotaSpecBindings returns a QuotaSpecBindings
func newQuotaSpecBindings(c *ConfigV1alpha2Client, namespace string) *quotaSpecBindings {
	return &quotaSpecBindings{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the quotaSpecBinding, and returns the corresponding quotaSpecBinding object, and an error if there is any.
func (c *quotaSpecBindings) Get(name string, options v1.GetOptions) (result *v1alpha2.QuotaSpecBinding, err error) {
	result = &v1alpha2.QuotaSpecBinding{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("quotaspecbindings").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of QuotaSpecBindings that match those selectors.
func (c *quotaSpecBindings) List(opts v1.ListOptions) (result *v1alpha2.QuotaSpecBindingList, err error) {
	result = &v1alpha2.QuotaSpecBindingList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("quotaspecbindings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested quotaSpecBindings.
func (c *quotaSpecBindings) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("quotaspecbindings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a quotaSpecBinding and creates it.  Returns the server's representation of the quotaSpecBinding, and an error, if there is any.
func (c *quotaSpecBindings) Create(quotaSpecBinding *v1alpha2.QuotaSpecBinding) (result *v1alpha2.QuotaSpecBinding, err error) {
	result = &v1alpha2.QuotaSpecBinding{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("quotaspecbindings").
		Body(quotaSpecBinding).
		Do().
		Into(result)
	return
}

// Update takes the representation of a quotaSpecBinding and updates it. Returns the server's representation of the quotaSpecBinding, and an error, if there is any.
func (c *quotaSpecBindings) Update(quotaSpecBinding *v1alpha2.QuotaSpecBinding) (result *v1alpha2.QuotaSpecBinding, err error) {
	result = &v1alpha2.QuotaSpecBinding{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("quotaspecbindings").
		Name(quotaSpecBinding.Name).
		Body(quotaSpecBinding).
		Do().
		Into(result)
	return
}

// Delete takes name of the quotaSpecBinding and deletes it. Returns an error if one occurs.
func (c *quotaSpecBindings) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("quotaspecbindings").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *quotaSpecBindings) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("quotaspecbindings").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched quotaSpecBinding.
func (c *quotaSpecBindings) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha2.QuotaSpecBinding, err error) {
	result = &v1alpha2.QuotaSpecBinding{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("quotaspecbindings").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
