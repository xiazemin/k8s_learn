/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "code-generator/api/samplecontroller/v1alpha1"
	scheme "code-generator/generated/clientset/versioned/scheme"
	"context"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// FoosGetter has a method to return a FooInterface.
// A group's client should implement this interface.
type FoosGetter interface {
	Foos(namespace string) FooInterface
}

// FooInterface has methods to work with Foo resources.
type FooInterface interface {
	Create(ctx context.Context, foo *v1alpha1.Foo, opts v1.CreateOptions) (*v1alpha1.Foo, error)
	Update(ctx context.Context, foo *v1alpha1.Foo, opts v1.UpdateOptions) (*v1alpha1.Foo, error)
	UpdateStatus(ctx context.Context, foo *v1alpha1.Foo, opts v1.UpdateOptions) (*v1alpha1.Foo, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.Foo, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.FooList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Foo, err error)
	FooExpansion
}

// foos implements FooInterface
type foos struct {
	client rest.Interface
	ns     string
}

// newFoos returns a Foos
func newFoos(c *SamplecontrollerV1alpha1Client, namespace string) *foos {
	return &foos{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the foo, and returns the corresponding foo object, and an error if there is any.
func (c *foos) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Foo, err error) {
	result = &v1alpha1.Foo{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("foos").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Foos that match those selectors.
func (c *foos) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.FooList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.FooList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("foos").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested foos.
func (c *foos) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("foos").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a foo and creates it.  Returns the server's representation of the foo, and an error, if there is any.
func (c *foos) Create(ctx context.Context, foo *v1alpha1.Foo, opts v1.CreateOptions) (result *v1alpha1.Foo, err error) {
	result = &v1alpha1.Foo{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("foos").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(foo).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a foo and updates it. Returns the server's representation of the foo, and an error, if there is any.
func (c *foos) Update(ctx context.Context, foo *v1alpha1.Foo, opts v1.UpdateOptions) (result *v1alpha1.Foo, err error) {
	result = &v1alpha1.Foo{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("foos").
		Name(foo.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(foo).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *foos) UpdateStatus(ctx context.Context, foo *v1alpha1.Foo, opts v1.UpdateOptions) (result *v1alpha1.Foo, err error) {
	result = &v1alpha1.Foo{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("foos").
		Name(foo.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(foo).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the foo and deletes it. Returns an error if one occurs.
func (c *foos) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("foos").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *foos) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("foos").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched foo.
func (c *foos) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Foo, err error) {
	result = &v1alpha1.Foo{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("foos").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
