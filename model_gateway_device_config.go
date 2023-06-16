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

// checks if the GatewayDeviceConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GatewayDeviceConfig{}

// GatewayDeviceConfig Configuration parameters for a device in a user-managed gateway. The configuration parameters depends on the type of gateway.
type GatewayDeviceConfig struct {
	// This is the ID of the gateway this configuration applies to.
	GatewayId *string `json:"gatewayId,omitempty"`
	Params *map[string]string `json:"params,omitempty"`
}

// NewGatewayDeviceConfig instantiates a new GatewayDeviceConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayDeviceConfig() *GatewayDeviceConfig {
	this := GatewayDeviceConfig{}
	return &this
}

// NewGatewayDeviceConfigWithDefaults instantiates a new GatewayDeviceConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayDeviceConfigWithDefaults() *GatewayDeviceConfig {
	this := GatewayDeviceConfig{}
	return &this
}

// GetGatewayId returns the GatewayId field value if set, zero value otherwise.
func (o *GatewayDeviceConfig) GetGatewayId() string {
	if o == nil || IsNil(o.GatewayId) {
		var ret string
		return ret
	}
	return *o.GatewayId
}

// GetGatewayIdOk returns a tuple with the GatewayId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayDeviceConfig) GetGatewayIdOk() (*string, bool) {
	if o == nil || IsNil(o.GatewayId) {
		return nil, false
	}
	return o.GatewayId, true
}

// HasGatewayId returns a boolean if a field has been set.
func (o *GatewayDeviceConfig) HasGatewayId() bool {
	if o != nil && !IsNil(o.GatewayId) {
		return true
	}

	return false
}

// SetGatewayId gets a reference to the given string and assigns it to the GatewayId field.
func (o *GatewayDeviceConfig) SetGatewayId(v string) {
	o.GatewayId = &v
}

// GetParams returns the Params field value if set, zero value otherwise.
func (o *GatewayDeviceConfig) GetParams() map[string]string {
	if o == nil || IsNil(o.Params) {
		var ret map[string]string
		return ret
	}
	return *o.Params
}

// GetParamsOk returns a tuple with the Params field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayDeviceConfig) GetParamsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Params) {
		return nil, false
	}
	return o.Params, true
}

// HasParams returns a boolean if a field has been set.
func (o *GatewayDeviceConfig) HasParams() bool {
	if o != nil && !IsNil(o.Params) {
		return true
	}

	return false
}

// SetParams gets a reference to the given map[string]string and assigns it to the Params field.
func (o *GatewayDeviceConfig) SetParams(v map[string]string) {
	o.Params = &v
}

func (o GatewayDeviceConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GatewayDeviceConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GatewayId) {
		toSerialize["gatewayId"] = o.GatewayId
	}
	if !IsNil(o.Params) {
		toSerialize["params"] = o.Params
	}
	return toSerialize, nil
}

type NullableGatewayDeviceConfig struct {
	value *GatewayDeviceConfig
	isSet bool
}

func (v NullableGatewayDeviceConfig) Get() *GatewayDeviceConfig {
	return v.value
}

func (v *NullableGatewayDeviceConfig) Set(val *GatewayDeviceConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayDeviceConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayDeviceConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayDeviceConfig(val *GatewayDeviceConfig) *NullableGatewayDeviceConfig {
	return &NullableGatewayDeviceConfig{value: val, isSet: true}
}

func (v NullableGatewayDeviceConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayDeviceConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


