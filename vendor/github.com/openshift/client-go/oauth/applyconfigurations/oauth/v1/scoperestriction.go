// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// ScopeRestrictionApplyConfiguration represents an declarative configuration of the ScopeRestriction type for use
// with apply.
type ScopeRestrictionApplyConfiguration struct {
	ExactValues []string                                       `json:"literals,omitempty"`
	ClusterRole *ClusterRoleScopeRestrictionApplyConfiguration `json:"clusterRole,omitempty"`
}

// ScopeRestrictionApplyConfiguration constructs an declarative configuration of the ScopeRestriction type for use with
// apply.
func ScopeRestriction() *ScopeRestrictionApplyConfiguration {
	return &ScopeRestrictionApplyConfiguration{}
}

// WithExactValues adds the given value to the ExactValues field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the ExactValues field.
func (b *ScopeRestrictionApplyConfiguration) WithExactValues(values ...string) *ScopeRestrictionApplyConfiguration {
	for i := range values {
		b.ExactValues = append(b.ExactValues, values[i])
	}
	return b
}

// WithClusterRole sets the ClusterRole field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ClusterRole field is set to the value of the last call.
func (b *ScopeRestrictionApplyConfiguration) WithClusterRole(value *ClusterRoleScopeRestrictionApplyConfiguration) *ScopeRestrictionApplyConfiguration {
	b.ClusterRole = value
	return b
}