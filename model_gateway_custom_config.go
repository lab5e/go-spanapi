/*
The Span API

API for device, collection, output and firmware management

API version: 4.6.1 squirming-codi
Contact: dev@lab5e.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package spanapi

import (
	"encoding/json"
)

// checks if the GatewayCustomConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GatewayCustomConfig{}

// GatewayCustomConfig struct for GatewayCustomConfig
type GatewayCustomConfig struct {
	Params *map[string]string `json:"params,omitempty"`
}

// NewGatewayCustomConfig instantiates a new GatewayCustomConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayCustomConfig() *GatewayCustomConfig {
	this := GatewayCustomConfig{}
	return &this
}

// NewGatewayCustomConfigWithDefaults instantiates a new GatewayCustomConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayCustomConfigWithDefaults() *GatewayCustomConfig {
	this := GatewayCustomConfig{}
	return &this
}

// GetParams returns the Params field value if set, zero value otherwise.
func (o *GatewayCustomConfig) GetParams() map[string]string {
	if o == nil || IsNil(o.Params) {
		var ret map[string]string
		return ret
	}
	return *o.Params
}

// GetParamsOk returns a tuple with the Params field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCustomConfig) GetParamsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Params) {
		return nil, false
	}
	return o.Params, true
}

// HasParams returns a boolean if a field has been set.
func (o *GatewayCustomConfig) HasParams() bool {
	if o != nil && !IsNil(o.Params) {
		return true
	}

	return false
}

// SetParams gets a reference to the given map[string]string and assigns it to the Params field.
func (o *GatewayCustomConfig) SetParams(v map[string]string) {
	o.Params = &v
}

func (o GatewayCustomConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GatewayCustomConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Params) {
		toSerialize["params"] = o.Params
	}
	return toSerialize, nil
}

type NullableGatewayCustomConfig struct {
	value *GatewayCustomConfig
	isSet bool
}

func (v NullableGatewayCustomConfig) Get() *GatewayCustomConfig {
	return v.value
}

func (v *NullableGatewayCustomConfig) Set(val *GatewayCustomConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayCustomConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayCustomConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayCustomConfig(val *GatewayCustomConfig) *NullableGatewayCustomConfig {
	return &NullableGatewayCustomConfig{value: val, isSet: true}
}

func (v NullableGatewayCustomConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayCustomConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


