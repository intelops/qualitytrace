/*
 * TraceTest
 *
 * OpenAPI definition for TraceTest endpoint and resources
 *
 * API version: 0.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type TestSummary struct {
	Runs int32 `json:"runs,omitempty"`

	LastRun TestSummaryLastRun `json:"lastRun,omitempty"`
}

// AssertTestSummaryRequired checks if the required fields are not zero-ed
func AssertTestSummaryRequired(obj TestSummary) error {
	if err := AssertTestSummaryLastRunRequired(obj.LastRun); err != nil {
		return err
	}
	return nil
}

// AssertRecurseTestSummaryRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of TestSummary (e.g. [][]TestSummary), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseTestSummaryRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aTestSummary, ok := obj.(TestSummary)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertTestSummaryRequired(aTestSummary)
	})
}
