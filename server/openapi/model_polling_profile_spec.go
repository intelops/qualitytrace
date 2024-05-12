/*
 * TraceTest
 *
 * OpenAPI definition for TraceTest endpoint and resources
 *
 * API version: 0.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// PollingProfileSpec - Represents the attributes of a Polling Profile.
type PollingProfileSpec struct {

	// ID of this Polling Profile.
	Id string `json:"id"`

	// Name given for this profile.
	Name string `json:"name"`

	// Is default polling profile
	Default bool `json:"default,omitempty"`

	// Name of the strategy that will be used on this profile.
	Strategy string `json:"strategy"`

	Periodic PollingProfileSpecPeriodic `json:"periodic,omitempty"`
}

// AssertPollingProfileSpecRequired checks if the required fields are not zero-ed
func AssertPollingProfileSpecRequired(obj PollingProfileSpec) error {
	elements := map[string]interface{}{
		"id":       obj.Id,
		"name":     obj.Name,
		"strategy": obj.Strategy,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertPollingProfileSpecPeriodicRequired(obj.Periodic); err != nil {
		return err
	}
	return nil
}

// AssertRecursePollingProfileSpecRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of PollingProfileSpec (e.g. [][]PollingProfileSpec), otherwise ErrTypeAssertionError is thrown.
func AssertRecursePollingProfileSpecRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aPollingProfileSpec, ok := obj.(PollingProfileSpec)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertPollingProfileSpecRequired(aPollingProfileSpec)
	})
}
