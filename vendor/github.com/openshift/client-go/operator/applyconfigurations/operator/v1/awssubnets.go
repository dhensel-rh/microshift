// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/openshift/api/operator/v1"
)

// AWSSubnetsApplyConfiguration represents a declarative configuration of the AWSSubnets type for use
// with apply.
type AWSSubnetsApplyConfiguration struct {
	IDs   []v1.AWSSubnetID   `json:"ids,omitempty"`
	Names []v1.AWSSubnetName `json:"names,omitempty"`
}

// AWSSubnetsApplyConfiguration constructs a declarative configuration of the AWSSubnets type for use with
// apply.
func AWSSubnets() *AWSSubnetsApplyConfiguration {
	return &AWSSubnetsApplyConfiguration{}
}

// WithIDs adds the given value to the IDs field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the IDs field.
func (b *AWSSubnetsApplyConfiguration) WithIDs(values ...v1.AWSSubnetID) *AWSSubnetsApplyConfiguration {
	for i := range values {
		b.IDs = append(b.IDs, values[i])
	}
	return b
}

// WithNames adds the given value to the Names field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Names field.
func (b *AWSSubnetsApplyConfiguration) WithNames(values ...v1.AWSSubnetName) *AWSSubnetsApplyConfiguration {
	for i := range values {
		b.Names = append(b.Names, values[i])
	}
	return b
}
