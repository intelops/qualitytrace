/*
 * TraceTest
 *
 * OpenAPI definition for TraceTest endpoint and resources
 *
 * API version: 0.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type GrpcClientSettings struct {
	Endpoint string `json:"endpoint,omitempty"`

	ReadBufferSize float32 `json:"readBufferSize,omitempty"`

	WriteBufferSize float32 `json:"writeBufferSize,omitempty"`

	WaitForReady bool `json:"waitForReady,omitempty"`

	Headers map[string]string `json:"headers,omitempty"`

	BalancerName string `json:"balancerName,omitempty"`

	Compression string `json:"compression,omitempty"`

	Tls Tls `json:"tls,omitempty"`

	Auth HttpAuth `json:"auth,omitempty"`
}

// AssertGrpcClientSettingsRequired checks if the required fields are not zero-ed
func AssertGrpcClientSettingsRequired(obj GrpcClientSettings) error {
	if err := AssertTlsRequired(obj.Tls); err != nil {
		return err
	}
	if err := AssertHttpAuthRequired(obj.Auth); err != nil {
		return err
	}
	return nil
}

// AssertRecurseGrpcClientSettingsRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of GrpcClientSettings (e.g. [][]GrpcClientSettings), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseGrpcClientSettingsRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aGrpcClientSettings, ok := obj.(GrpcClientSettings)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertGrpcClientSettingsRequired(aGrpcClientSettings)
	})
}
