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

// checks if the HTTPAuth type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HTTPAuth{}

// HTTPAuth struct for HTTPAuth
type HTTPAuth struct {
	Type   *string         `json:"type,omitempty"`
	ApiKey *HTTPAuthApiKey `json:"apiKey,omitempty"`
	Basic  *HTTPAuthBasic  `json:"basic,omitempty"`
	Bearer *HTTPAuthBearer `json:"bearer,omitempty"`
}

// NewHTTPAuth instantiates a new HTTPAuth object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHTTPAuth() *HTTPAuth {
	this := HTTPAuth{}
	return &this
}

// NewHTTPAuthWithDefaults instantiates a new HTTPAuth object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHTTPAuthWithDefaults() *HTTPAuth {
	this := HTTPAuth{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *HTTPAuth) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HTTPAuth) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *HTTPAuth) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *HTTPAuth) SetType(v string) {
	o.Type = &v
}

// GetApiKey returns the ApiKey field value if set, zero value otherwise.
func (o *HTTPAuth) GetApiKey() HTTPAuthApiKey {
	if o == nil || isNil(o.ApiKey) {
		var ret HTTPAuthApiKey
		return ret
	}
	return *o.ApiKey
}

// GetApiKeyOk returns a tuple with the ApiKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HTTPAuth) GetApiKeyOk() (*HTTPAuthApiKey, bool) {
	if o == nil || isNil(o.ApiKey) {
		return nil, false
	}
	return o.ApiKey, true
}

// HasApiKey returns a boolean if a field has been set.
func (o *HTTPAuth) HasApiKey() bool {
	if o != nil && !isNil(o.ApiKey) {
		return true
	}

	return false
}

// SetApiKey gets a reference to the given HTTPAuthApiKey and assigns it to the ApiKey field.
func (o *HTTPAuth) SetApiKey(v HTTPAuthApiKey) {
	o.ApiKey = &v
}

// GetBasic returns the Basic field value if set, zero value otherwise.
func (o *HTTPAuth) GetBasic() HTTPAuthBasic {
	if o == nil || isNil(o.Basic) {
		var ret HTTPAuthBasic
		return ret
	}
	return *o.Basic
}

// GetBasicOk returns a tuple with the Basic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HTTPAuth) GetBasicOk() (*HTTPAuthBasic, bool) {
	if o == nil || isNil(o.Basic) {
		return nil, false
	}
	return o.Basic, true
}

// HasBasic returns a boolean if a field has been set.
func (o *HTTPAuth) HasBasic() bool {
	if o != nil && !isNil(o.Basic) {
		return true
	}

	return false
}

// SetBasic gets a reference to the given HTTPAuthBasic and assigns it to the Basic field.
func (o *HTTPAuth) SetBasic(v HTTPAuthBasic) {
	o.Basic = &v
}

// GetBearer returns the Bearer field value if set, zero value otherwise.
func (o *HTTPAuth) GetBearer() HTTPAuthBearer {
	if o == nil || isNil(o.Bearer) {
		var ret HTTPAuthBearer
		return ret
	}
	return *o.Bearer
}

// GetBearerOk returns a tuple with the Bearer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HTTPAuth) GetBearerOk() (*HTTPAuthBearer, bool) {
	if o == nil || isNil(o.Bearer) {
		return nil, false
	}
	return o.Bearer, true
}

// HasBearer returns a boolean if a field has been set.
func (o *HTTPAuth) HasBearer() bool {
	if o != nil && !isNil(o.Bearer) {
		return true
	}

	return false
}

// SetBearer gets a reference to the given HTTPAuthBearer and assigns it to the Bearer field.
func (o *HTTPAuth) SetBearer(v HTTPAuthBearer) {
	o.Bearer = &v
}

func (o HTTPAuth) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HTTPAuth) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.ApiKey) {
		toSerialize["apiKey"] = o.ApiKey
	}
	if !isNil(o.Basic) {
		toSerialize["basic"] = o.Basic
	}
	if !isNil(o.Bearer) {
		toSerialize["bearer"] = o.Bearer
	}
	return toSerialize, nil
}

type NullableHTTPAuth struct {
	value *HTTPAuth
	isSet bool
}

func (v NullableHTTPAuth) Get() *HTTPAuth {
	return v.value
}

func (v *NullableHTTPAuth) Set(val *HTTPAuth) {
	v.value = val
	v.isSet = true
}

func (v NullableHTTPAuth) IsSet() bool {
	return v.isSet
}

func (v *NullableHTTPAuth) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHTTPAuth(val *HTTPAuth) *NullableHTTPAuth {
	return &NullableHTTPAuth{value: val, isSet: true}
}

func (v NullableHTTPAuth) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHTTPAuth) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}