/*
 * TraceTest
 *
 * OpenAPI definition for TraceTest endpoint and resources
 *
 * API version: 0.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type HttpHeader struct {
	Key string `json:"key,omitempty"`

	Value string `json:"value,omitempty"`
}

// AssertHttpHeaderRequired checks if the required fields are not zero-ed
func AssertHttpHeaderRequired(obj HttpHeader) error {
	return nil
}

// AssertRecurseHttpHeaderRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of HttpHeader (e.g. [][]HttpHeader), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseHttpHeaderRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aHttpHeader, ok := obj.(HttpHeader)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertHttpHeaderRequired(aHttpHeader)
	})
}
