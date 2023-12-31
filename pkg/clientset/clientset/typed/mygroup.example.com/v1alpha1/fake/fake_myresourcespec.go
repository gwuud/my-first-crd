// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/gwuud/my-first-crd/pkg/apis/mygroup.example.com/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeMyResourceSpecs implements MyResourceSpecInterface
type FakeMyResourceSpecs struct {
	Fake *FakeMygroupV1alpha1
	ns   string
}

var myresourcespecsResource = v1alpha1.SchemeGroupVersion.WithResource("myresourcespecs")

var myresourcespecsKind = v1alpha1.SchemeGroupVersion.WithKind("MyResourceSpec")

// Get takes name of the myResourceSpec, and returns the corresponding myResourceSpec object, and an error if there is any.
func (c *FakeMyResourceSpecs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.MyResourceSpec, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(myresourcespecsResource, c.ns, name), &v1alpha1.MyResourceSpec{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MyResourceSpec), err
}

// List takes label and field selectors, and returns the list of MyResourceSpecs that match those selectors.
func (c *FakeMyResourceSpecs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.MyResourceSpecList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(myresourcespecsResource, myresourcespecsKind, c.ns, opts), &v1alpha1.MyResourceSpecList{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MyResourceSpecList), err
}

// Watch returns a watch.Interface that watches the requested myResourceSpecs.
func (c *FakeMyResourceSpecs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(myresourcespecsResource, c.ns, opts))

}

// Create takes the representation of a myResourceSpec and creates it.  Returns the server's representation of the myResourceSpec, and an error, if there is any.
func (c *FakeMyResourceSpecs) Create(ctx context.Context, myResourceSpec *v1alpha1.MyResourceSpec, opts v1.CreateOptions) (result *v1alpha1.MyResourceSpec, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(myresourcespecsResource, c.ns, myResourceSpec), &v1alpha1.MyResourceSpec{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MyResourceSpec), err
}

// Update takes the representation of a myResourceSpec and updates it. Returns the server's representation of the myResourceSpec, and an error, if there is any.
func (c *FakeMyResourceSpecs) Update(ctx context.Context, myResourceSpec *v1alpha1.MyResourceSpec, opts v1.UpdateOptions) (result *v1alpha1.MyResourceSpec, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(myresourcespecsResource, c.ns, myResourceSpec), &v1alpha1.MyResourceSpec{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MyResourceSpec), err
}

// Delete takes name of the myResourceSpec and deletes it. Returns an error if one occurs.
func (c *FakeMyResourceSpecs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(myresourcespecsResource, c.ns, name, opts), &v1alpha1.MyResourceSpec{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeMyResourceSpecs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(myresourcespecsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.MyResourceSpecList{})
	return err
}

// Patch applies the patch and returns the patched myResourceSpec.
func (c *FakeMyResourceSpecs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.MyResourceSpec, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(myresourcespecsResource, c.ns, name, pt, data, subresources...), &v1alpha1.MyResourceSpec{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MyResourceSpec), err
}
