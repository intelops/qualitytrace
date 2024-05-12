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

// checks if the ConfigurationResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfigurationResource{}

// ConfigurationResource Represents a configuration structured into the Resources format.
type ConfigurationResource struct {
	// Represents the type of this resource. It should always be set as 'Config'.
	Type *string                    `json:"type,omitempty"`
	Spec *ConfigurationResourceSpec `json:"spec,omitempty"`
}

// NewConfigurationResource instantiates a new ConfigurationResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigurationResource() *ConfigurationResource {
	this := ConfigurationResource{}
	return &this
}

// NewConfigurationResourceWithDefaults instantiates a new ConfigurationResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigurationResourceWithDefaults() *ConfigurationResource {
	this := ConfigurationResource{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ConfigurationResource) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationResource) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ConfigurationResource) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ConfigurationResource) SetType(v string) {
	o.Type = &v
}

// GetSpec returns the Spec field value if set, zero value otherwise.
func (o *ConfigurationResource) GetSpec() ConfigurationResourceSpec {
	if o == nil || isNil(o.Spec) {
		var ret ConfigurationResourceSpec
		return ret
	}
	return *o.Spec
}

// GetSpecOk returns a tuple with the Spec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationResource) GetSpecOk() (*ConfigurationResourceSpec, bool) {
	if o == nil || isNil(o.Spec) {
		return nil, false
	}
	return o.Spec, true
}

// HasSpec returns a boolean if a field has been set.
func (o *ConfigurationResource) HasSpec() bool {
	if o != nil && !isNil(o.Spec) {
		return true
	}

	return false
}

// SetSpec gets a reference to the given ConfigurationResourceSpec and assigns it to the Spec field.
func (o *ConfigurationResource) SetSpec(v ConfigurationResourceSpec) {
	o.Spec = &v
}

func (o ConfigurationResource) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConfigurationResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.Spec) {
		toSerialize["spec"] = o.Spec
	}
	return toSerialize, nil
}

type NullableConfigurationResource struct {
	value *ConfigurationResource
	isSet bool
}

func (v NullableConfigurationResource) Get() *ConfigurationResource {
	return v.value
}

func (v *NullableConfigurationResource) Set(val *ConfigurationResource) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigurationResource) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigurationResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigurationResource(val *ConfigurationResource) *NullableConfigurationResource {
	return &NullableConfigurationResource{value: val, isSet: true}
}

func (v NullableConfigurationResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigurationResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
