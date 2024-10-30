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

package fake

import (
	"context"
	json "encoding/json"
	"fmt"

	v1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	networkingv1 "k8s.io/client-go/applyconfigurations/networking/v1"
	testing "k8s.io/client-go/testing"
)

// FakeIngressClasses implements IngressClassInterface
type FakeIngressClasses struct {
	Fake *FakeNetworkingV1
}

var ingressclassesResource = v1.SchemeGroupVersion.WithResource("ingressclasses")

var ingressclassesKind = v1.SchemeGroupVersion.WithKind("IngressClass")

// Get takes name of the ingressClass, and returns the corresponding ingressClass object, and an error if there is any.
func (c *FakeIngressClasses) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.IngressClass, err error) {
	emptyResult := &v1.IngressClass{}
	obj, err := c.Fake.
		Invokes(testing.NewRootGetActionWithOptions(ingressclassesResource, name, options), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.IngressClass), err
}

// List takes label and field selectors, and returns the list of IngressClasses that match those selectors.
func (c *FakeIngressClasses) List(ctx context.Context, opts metav1.ListOptions) (result *v1.IngressClassList, err error) {
	emptyResult := &v1.IngressClassList{}
	obj, err := c.Fake.
		Invokes(testing.NewRootListActionWithOptions(ingressclassesResource, ingressclassesKind, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.IngressClassList{ListMeta: obj.(*v1.IngressClassList).ListMeta}
	for _, item := range obj.(*v1.IngressClassList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested ingressClasses.
func (c *FakeIngressClasses) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchActionWithOptions(ingressclassesResource, opts))
}

// Create takes the representation of a ingressClass and creates it.  Returns the server's representation of the ingressClass, and an error, if there is any.
func (c *FakeIngressClasses) Create(ctx context.Context, ingressClass *v1.IngressClass, opts metav1.CreateOptions) (result *v1.IngressClass, err error) {
	emptyResult := &v1.IngressClass{}
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateActionWithOptions(ingressclassesResource, ingressClass, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.IngressClass), err
}

// Update takes the representation of a ingressClass and updates it. Returns the server's representation of the ingressClass, and an error, if there is any.
func (c *FakeIngressClasses) Update(ctx context.Context, ingressClass *v1.IngressClass, opts metav1.UpdateOptions) (result *v1.IngressClass, err error) {
	emptyResult := &v1.IngressClass{}
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateActionWithOptions(ingressclassesResource, ingressClass, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.IngressClass), err
}

// Delete takes name of the ingressClass and deletes it. Returns an error if one occurs.
func (c *FakeIngressClasses) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(ingressclassesResource, name, opts), &v1.IngressClass{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeIngressClasses) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewRootDeleteCollectionActionWithOptions(ingressclassesResource, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v1.IngressClassList{})
	return err
}

// Patch applies the patch and returns the patched ingressClass.
func (c *FakeIngressClasses) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.IngressClass, err error) {
	emptyResult := &v1.IngressClass{}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceActionWithOptions(ingressclassesResource, name, pt, data, opts, subresources...), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.IngressClass), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied ingressClass.
func (c *FakeIngressClasses) Apply(ctx context.Context, ingressClass *networkingv1.IngressClassApplyConfiguration, opts metav1.ApplyOptions) (result *v1.IngressClass, err error) {
	if ingressClass == nil {
		return nil, fmt.Errorf("ingressClass provided to Apply must not be nil")
	}
	data, err := json.Marshal(ingressClass)
	if err != nil {
		return nil, err
	}
	name := ingressClass.Name
	if name == nil {
		return nil, fmt.Errorf("ingressClass.Name must be provided to Apply")
	}
	emptyResult := &v1.IngressClass{}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceActionWithOptions(ingressclassesResource, *name, types.ApplyPatchType, data, opts.ToPatchOptions()), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.IngressClass), err
}
