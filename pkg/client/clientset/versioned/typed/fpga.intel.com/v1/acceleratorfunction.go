// Copyright 2020 Intel Corporation. All Rights Reserved.
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

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"
	"time"

	v1 "github.com/intel/intel-device-plugins-for-kubernetes/pkg/apis/fpga.intel.com/v1"
	scheme "github.com/intel/intel-device-plugins-for-kubernetes/pkg/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// AcceleratorFunctionsGetter has a method to return a AcceleratorFunctionInterface.
// A group's client should implement this interface.
type AcceleratorFunctionsGetter interface {
	AcceleratorFunctions(namespace string) AcceleratorFunctionInterface
}

// AcceleratorFunctionInterface has methods to work with AcceleratorFunction resources.
type AcceleratorFunctionInterface interface {
	Create(ctx context.Context, acceleratorFunction *v1.AcceleratorFunction, opts metav1.CreateOptions) (*v1.AcceleratorFunction, error)
	Update(ctx context.Context, acceleratorFunction *v1.AcceleratorFunction, opts metav1.UpdateOptions) (*v1.AcceleratorFunction, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.AcceleratorFunction, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.AcceleratorFunctionList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.AcceleratorFunction, err error)
	AcceleratorFunctionExpansion
}

// acceleratorFunctions implements AcceleratorFunctionInterface
type acceleratorFunctions struct {
	client rest.Interface
	ns     string
}

// newAcceleratorFunctions returns a AcceleratorFunctions
func newAcceleratorFunctions(c *FpgaV1Client, namespace string) *acceleratorFunctions {
	return &acceleratorFunctions{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the acceleratorFunction, and returns the corresponding acceleratorFunction object, and an error if there is any.
func (c *acceleratorFunctions) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.AcceleratorFunction, err error) {
	result = &v1.AcceleratorFunction{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("acceleratorfunctions").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AcceleratorFunctions that match those selectors.
func (c *acceleratorFunctions) List(ctx context.Context, opts metav1.ListOptions) (result *v1.AcceleratorFunctionList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.AcceleratorFunctionList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("acceleratorfunctions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested acceleratorFunctions.
func (c *acceleratorFunctions) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("acceleratorfunctions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a acceleratorFunction and creates it.  Returns the server's representation of the acceleratorFunction, and an error, if there is any.
func (c *acceleratorFunctions) Create(ctx context.Context, acceleratorFunction *v1.AcceleratorFunction, opts metav1.CreateOptions) (result *v1.AcceleratorFunction, err error) {
	result = &v1.AcceleratorFunction{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("acceleratorfunctions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(acceleratorFunction).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a acceleratorFunction and updates it. Returns the server's representation of the acceleratorFunction, and an error, if there is any.
func (c *acceleratorFunctions) Update(ctx context.Context, acceleratorFunction *v1.AcceleratorFunction, opts metav1.UpdateOptions) (result *v1.AcceleratorFunction, err error) {
	result = &v1.AcceleratorFunction{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("acceleratorfunctions").
		Name(acceleratorFunction.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(acceleratorFunction).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the acceleratorFunction and deletes it. Returns an error if one occurs.
func (c *acceleratorFunctions) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("acceleratorfunctions").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *acceleratorFunctions) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("acceleratorfunctions").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched acceleratorFunction.
func (c *acceleratorFunctions) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.AcceleratorFunction, err error) {
	result = &v1.AcceleratorFunction{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("acceleratorfunctions").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
