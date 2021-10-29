/*
/*
@Time : 2019/12/23 3:08 下午
@Author : tianpeng.du
@File : types
@Software: GoLand
*/

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"
	"time"

	v1 "ip/pkg/apis/ip/v1"
	scheme "ip/pkg/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// IpsGetter has a method to return a IpInterface.
// A group's client should implement this interface.
type IpsGetter interface {
	Ips(namespace string) IpInterface
}

// IpInterface has methods to work with Ip resources.
type IpInterface interface {
	Create(ctx context.Context, ip *v1.Ip, opts metav1.CreateOptions) (*v1.Ip, error)
	Update(ctx context.Context, ip *v1.Ip, opts metav1.UpdateOptions) (*v1.Ip, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.Ip, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.IpList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Ip, err error)
	IpExpansion
}

// ips implements IpInterface
type ips struct {
	client rest.Interface
	ns     string
}

// newIps returns a Ips
func newIps(c *RocduV1Client, namespace string) *ips {
	return &ips{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the ip, and returns the corresponding ip object, and an error if there is any.
func (c *ips) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.Ip, err error) {
	result = &v1.Ip{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("ips").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Ips that match those selectors.
func (c *ips) List(ctx context.Context, opts metav1.ListOptions) (result *v1.IpList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.IpList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("ips").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested ips.
func (c *ips) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("ips").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a ip and creates it.  Returns the server's representation of the ip, and an error, if there is any.
func (c *ips) Create(ctx context.Context, ip *v1.Ip, opts metav1.CreateOptions) (result *v1.Ip, err error) {
	result = &v1.Ip{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("ips").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(ip).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a ip and updates it. Returns the server's representation of the ip, and an error, if there is any.
func (c *ips) Update(ctx context.Context, ip *v1.Ip, opts metav1.UpdateOptions) (result *v1.Ip, err error) {
	result = &v1.Ip{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("ips").
		Name(ip.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(ip).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the ip and deletes it. Returns an error if one occurs.
func (c *ips) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("ips").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *ips) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("ips").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched ip.
func (c *ips) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Ip, err error) {
	result = &v1.Ip{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("ips").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
