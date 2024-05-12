/*
 * TraceTest
 *
 * OpenAPI definition for TraceTest endpoint and resources
 *
 * API version: 0.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type TriggerResult struct {
	Type string `json:"type,omitempty"`

	TriggerResult TriggerResultTriggerResult `json:"triggerResult,omitempty"`
}

// AssertTriggerResultRequired checks if the required fields are not zero-ed
func AssertTriggerResultRequired(obj TriggerResult) error {
	if err := AssertTriggerResultTriggerResultRequired(obj.TriggerResult); err != nil {
		return err
	}
	return nil
}

// AssertRecurseTriggerResultRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of TriggerResult (e.g. [][]TriggerResult), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseTriggerResultRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aTriggerResult, ok := obj.(TriggerResult)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertTriggerResultRequired(aTriggerResult)
	})
}
