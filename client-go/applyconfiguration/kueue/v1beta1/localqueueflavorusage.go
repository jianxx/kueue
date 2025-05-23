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
// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1beta1

import (
	kueuev1beta1 "sigs.k8s.io/kueue/apis/kueue/v1beta1"
)

// LocalQueueFlavorUsageApplyConfiguration represents a declarative configuration of the LocalQueueFlavorUsage type for use
// with apply.
type LocalQueueFlavorUsageApplyConfiguration struct {
	Name      *kueuev1beta1.ResourceFlavorReference       `json:"name,omitempty"`
	Resources []LocalQueueResourceUsageApplyConfiguration `json:"resources,omitempty"`
}

// LocalQueueFlavorUsageApplyConfiguration constructs a declarative configuration of the LocalQueueFlavorUsage type for use with
// apply.
func LocalQueueFlavorUsage() *LocalQueueFlavorUsageApplyConfiguration {
	return &LocalQueueFlavorUsageApplyConfiguration{}
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *LocalQueueFlavorUsageApplyConfiguration) WithName(value kueuev1beta1.ResourceFlavorReference) *LocalQueueFlavorUsageApplyConfiguration {
	b.Name = &value
	return b
}

// WithResources adds the given value to the Resources field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Resources field.
func (b *LocalQueueFlavorUsageApplyConfiguration) WithResources(values ...*LocalQueueResourceUsageApplyConfiguration) *LocalQueueFlavorUsageApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithResources")
		}
		b.Resources = append(b.Resources, *values[i])
	}
	return b
}
