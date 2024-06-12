/*
TraceTest

OpenAPI definition for TraceTest endpoint and resources

API version: 0.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the HealthzGet200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HealthzGet200Response{}

// HealthzGet200Response struct for HealthzGet200Response
type HealthzGet200Response struct {
	Status *string `json:"status,omitempty"`
}

// NewHealthzGet200Response instantiates a new HealthzGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHealthzGet200Response() *HealthzGet200Response {
	this := HealthzGet200Response{}
	return &this
}

// NewHealthzGet200ResponseWithDefaults instantiates a new HealthzGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHealthzGet200ResponseWithDefaults() *HealthzGet200Response {
	this := HealthzGet200Response{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *HealthzGet200Response) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HealthzGet200Response) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *HealthzGet200Response) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *HealthzGet200Response) SetStatus(v string) {
	o.Status = &v
}

func (o HealthzGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HealthzGet200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableHealthzGet200Response struct {
	value *HealthzGet200Response
	isSet bool
}

func (v NullableHealthzGet200Response) Get() *HealthzGet200Response {
	return v.value
}

func (v *NullableHealthzGet200Response) Set(val *HealthzGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableHealthzGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableHealthzGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHealthzGet200Response(val *HealthzGet200Response) *NullableHealthzGet200Response {
	return &NullableHealthzGet200Response{value: val, isSet: true}
}

func (v NullableHealthzGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHealthzGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
