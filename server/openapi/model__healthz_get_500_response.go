/*
 * TraceTest
 *
 * OpenAPI definition for TraceTest endpoint and resources
 *
 * API version: 0.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type HealthzGet500Response struct {
	Status string `json:"status,omitempty"`
}

// AssertHealthzGet500ResponseRequired checks if the required fields are not zero-ed
func AssertHealthzGet500ResponseRequired(obj HealthzGet500Response) error {
	return nil
}

// AssertRecurseHealthzGet500ResponseRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of HealthzGet500Response (e.g. [][]HealthzGet500Response), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseHealthzGet500ResponseRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aHealthzGet500Response, ok := obj.(HealthzGet500Response)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertHealthzGet500ResponseRequired(aHealthzGet500Response)
	})
}
