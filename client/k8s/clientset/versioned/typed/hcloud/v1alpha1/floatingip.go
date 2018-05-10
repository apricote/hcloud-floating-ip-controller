package v1alpha1

import (
	v1alpha1 "github.com/apricote/hcloud-floating-ip-operator/apis/hcloud/v1alpha1"
	scheme "github.com/apricote/hcloud-floating-ip-operator/client/k8s/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// FloatingIPsGetter has a method to return a FloatingIPInterface.
// A group's client should implement this interface.
type FloatingIPsGetter interface {
	FloatingIPs() FloatingIPInterface
}

// FloatingIPInterface has methods to work with FloatingIP resources.
type FloatingIPInterface interface {
	Create(*v1alpha1.FloatingIP) (*v1alpha1.FloatingIP, error)
	Update(*v1alpha1.FloatingIP) (*v1alpha1.FloatingIP, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.FloatingIP, error)
	List(opts v1.ListOptions) (*v1alpha1.FloatingIPList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.FloatingIP, err error)
	FloatingIPExpansion
}

// floatingIPs implements FloatingIPInterface
type floatingIPs struct {
	client rest.Interface
}

// newFloatingIPs returns a FloatingIPs
func newFloatingIPs(c *HcloudV1alpha1Client) *floatingIPs {
	return &floatingIPs{
		client: c.RESTClient(),
	}
}

// Get takes name of the floatingIP, and returns the corresponding floatingIP object, and an error if there is any.
func (c *floatingIPs) Get(name string, options v1.GetOptions) (result *v1alpha1.FloatingIP, err error) {
	result = &v1alpha1.FloatingIP{}
	err = c.client.Get().
		Resource("floatingips").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of FloatingIPs that match those selectors.
func (c *floatingIPs) List(opts v1.ListOptions) (result *v1alpha1.FloatingIPList, err error) {
	result = &v1alpha1.FloatingIPList{}
	err = c.client.Get().
		Resource("floatingips").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested floatingIPs.
func (c *floatingIPs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Resource("floatingips").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a floatingIP and creates it.  Returns the server's representation of the floatingIP, and an error, if there is any.
func (c *floatingIPs) Create(floatingIP *v1alpha1.FloatingIP) (result *v1alpha1.FloatingIP, err error) {
	result = &v1alpha1.FloatingIP{}
	err = c.client.Post().
		Resource("floatingips").
		Body(floatingIP).
		Do().
		Into(result)
	return
}

// Update takes the representation of a floatingIP and updates it. Returns the server's representation of the floatingIP, and an error, if there is any.
func (c *floatingIPs) Update(floatingIP *v1alpha1.FloatingIP) (result *v1alpha1.FloatingIP, err error) {
	result = &v1alpha1.FloatingIP{}
	err = c.client.Put().
		Resource("floatingips").
		Name(floatingIP.Name).
		Body(floatingIP).
		Do().
		Into(result)
	return
}

// Delete takes name of the floatingIP and deletes it. Returns an error if one occurs.
func (c *floatingIPs) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("floatingips").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *floatingIPs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Resource("floatingips").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched floatingIP.
func (c *floatingIPs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.FloatingIP, err error) {
	result = &v1alpha1.FloatingIP{}
	err = c.client.Patch(pt).
		Resource("floatingips").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
