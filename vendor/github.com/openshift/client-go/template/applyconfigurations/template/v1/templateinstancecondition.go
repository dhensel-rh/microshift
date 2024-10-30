// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/openshift/api/template/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// TemplateInstanceConditionApplyConfiguration represents a declarative configuration of the TemplateInstanceCondition type for use
// with apply.
type TemplateInstanceConditionApplyConfiguration struct {
	Type               *v1.TemplateInstanceConditionType `json:"type,omitempty"`
	Status             *corev1.ConditionStatus           `json:"status,omitempty"`
	LastTransitionTime *metav1.Time                      `json:"lastTransitionTime,omitempty"`
	Reason             *string                           `json:"reason,omitempty"`
	Message            *string                           `json:"message,omitempty"`
}

// TemplateInstanceConditionApplyConfiguration constructs a declarative configuration of the TemplateInstanceCondition type for use with
// apply.
func TemplateInstanceCondition() *TemplateInstanceConditionApplyConfiguration {
	return &TemplateInstanceConditionApplyConfiguration{}
}

// WithType sets the Type field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Type field is set to the value of the last call.
func (b *TemplateInstanceConditionApplyConfiguration) WithType(value v1.TemplateInstanceConditionType) *TemplateInstanceConditionApplyConfiguration {
	b.Type = &value
	return b
}

// WithStatus sets the Status field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Status field is set to the value of the last call.
func (b *TemplateInstanceConditionApplyConfiguration) WithStatus(value corev1.ConditionStatus) *TemplateInstanceConditionApplyConfiguration {
	b.Status = &value
	return b
}

// WithLastTransitionTime sets the LastTransitionTime field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LastTransitionTime field is set to the value of the last call.
func (b *TemplateInstanceConditionApplyConfiguration) WithLastTransitionTime(value metav1.Time) *TemplateInstanceConditionApplyConfiguration {
	b.LastTransitionTime = &value
	return b
}

// WithReason sets the Reason field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Reason field is set to the value of the last call.
func (b *TemplateInstanceConditionApplyConfiguration) WithReason(value string) *TemplateInstanceConditionApplyConfiguration {
	b.Reason = &value
	return b
}

// WithMessage sets the Message field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Message field is set to the value of the last call.
func (b *TemplateInstanceConditionApplyConfiguration) WithMessage(value string) *TemplateInstanceConditionApplyConfiguration {
	b.Message = &value
	return b
}
