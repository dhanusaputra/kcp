/*
Copyright The KCP Authors.

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

package v1alpha1

import (
	v1alpha1 "github.com/kcp-dev/kcp/sdk/apis/scheduling/v1alpha1"
	conditionsv1alpha1 "github.com/kcp-dev/kcp/sdk/apis/third_party/conditions/apis/conditions/v1alpha1"
)

// PlacementStatusApplyConfiguration represents an declarative configuration of the PlacementStatus type for use
// with apply.
type PlacementStatusApplyConfiguration struct {
	Phase            *v1alpha1.PlacementPhase             `json:"phase,omitempty"`
	SelectedLocation *LocationReferenceApplyConfiguration `json:"selectedLocation,omitempty"`
	Conditions       *conditionsv1alpha1.Conditions       `json:"conditions,omitempty"`
}

// PlacementStatusApplyConfiguration constructs an declarative configuration of the PlacementStatus type for use with
// apply.
func PlacementStatus() *PlacementStatusApplyConfiguration {
	return &PlacementStatusApplyConfiguration{}
}

// WithPhase sets the Phase field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Phase field is set to the value of the last call.
func (b *PlacementStatusApplyConfiguration) WithPhase(value v1alpha1.PlacementPhase) *PlacementStatusApplyConfiguration {
	b.Phase = &value
	return b
}

// WithSelectedLocation sets the SelectedLocation field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SelectedLocation field is set to the value of the last call.
func (b *PlacementStatusApplyConfiguration) WithSelectedLocation(value *LocationReferenceApplyConfiguration) *PlacementStatusApplyConfiguration {
	b.SelectedLocation = value
	return b
}

// WithConditions sets the Conditions field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Conditions field is set to the value of the last call.
func (b *PlacementStatusApplyConfiguration) WithConditions(value conditionsv1alpha1.Conditions) *PlacementStatusApplyConfiguration {
	b.Conditions = &value
	return b
}
