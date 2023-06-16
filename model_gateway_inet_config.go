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

// checks if the GatewayInetConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GatewayInetConfig{}

// GatewayInetConfig struct for GatewayInetConfig
type GatewayInetConfig struct {
	DtlsEndpoint *string `json:"dtlsEndpoint,omitempty"`
	CoapEndpoint *string `json:"coapEndpoint,omitempty"`
	MqttEndpoint *string `json:"mqttEndpoint,omitempty"`
}

// NewGatewayInetConfig instantiates a new GatewayInetConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayInetConfig() *GatewayInetConfig {
	this := GatewayInetConfig{}
	return &this
}

// NewGatewayInetConfigWithDefaults instantiates a new GatewayInetConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayInetConfigWithDefaults() *GatewayInetConfig {
	this := GatewayInetConfig{}
	return &this
}

// GetDtlsEndpoint returns the DtlsEndpoint field value if set, zero value otherwise.
func (o *GatewayInetConfig) GetDtlsEndpoint() string {
	if o == nil || IsNil(o.DtlsEndpoint) {
		var ret string
		return ret
	}
	return *o.DtlsEndpoint
}

// GetDtlsEndpointOk returns a tuple with the DtlsEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayInetConfig) GetDtlsEndpointOk() (*string, bool) {
	if o == nil || IsNil(o.DtlsEndpoint) {
		return nil, false
	}
	return o.DtlsEndpoint, true
}

// HasDtlsEndpoint returns a boolean if a field has been set.
func (o *GatewayInetConfig) HasDtlsEndpoint() bool {
	if o != nil && !IsNil(o.DtlsEndpoint) {
		return true
	}

	return false
}

// SetDtlsEndpoint gets a reference to the given string and assigns it to the DtlsEndpoint field.
func (o *GatewayInetConfig) SetDtlsEndpoint(v string) {
	o.DtlsEndpoint = &v
}

// GetCoapEndpoint returns the CoapEndpoint field value if set, zero value otherwise.
func (o *GatewayInetConfig) GetCoapEndpoint() string {
	if o == nil || IsNil(o.CoapEndpoint) {
		var ret string
		return ret
	}
	return *o.CoapEndpoint
}

// GetCoapEndpointOk returns a tuple with the CoapEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayInetConfig) GetCoapEndpointOk() (*string, bool) {
	if o == nil || IsNil(o.CoapEndpoint) {
		return nil, false
	}
	return o.CoapEndpoint, true
}

// HasCoapEndpoint returns a boolean if a field has been set.
func (o *GatewayInetConfig) HasCoapEndpoint() bool {
	if o != nil && !IsNil(o.CoapEndpoint) {
		return true
	}

	return false
}

// SetCoapEndpoint gets a reference to the given string and assigns it to the CoapEndpoint field.
func (o *GatewayInetConfig) SetCoapEndpoint(v string) {
	o.CoapEndpoint = &v
}

// GetMqttEndpoint returns the MqttEndpoint field value if set, zero value otherwise.
func (o *GatewayInetConfig) GetMqttEndpoint() string {
	if o == nil || IsNil(o.MqttEndpoint) {
		var ret string
		return ret
	}
	return *o.MqttEndpoint
}

// GetMqttEndpointOk returns a tuple with the MqttEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayInetConfig) GetMqttEndpointOk() (*string, bool) {
	if o == nil || IsNil(o.MqttEndpoint) {
		return nil, false
	}
	return o.MqttEndpoint, true
}

// HasMqttEndpoint returns a boolean if a field has been set.
func (o *GatewayInetConfig) HasMqttEndpoint() bool {
	if o != nil && !IsNil(o.MqttEndpoint) {
		return true
	}

	return false
}

// SetMqttEndpoint gets a reference to the given string and assigns it to the MqttEndpoint field.
func (o *GatewayInetConfig) SetMqttEndpoint(v string) {
	o.MqttEndpoint = &v
}

func (o GatewayInetConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GatewayInetConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DtlsEndpoint) {
		toSerialize["dtlsEndpoint"] = o.DtlsEndpoint
	}
	if !IsNil(o.CoapEndpoint) {
		toSerialize["coapEndpoint"] = o.CoapEndpoint
	}
	if !IsNil(o.MqttEndpoint) {
		toSerialize["mqttEndpoint"] = o.MqttEndpoint
	}
	return toSerialize, nil
}

type NullableGatewayInetConfig struct {
	value *GatewayInetConfig
	isSet bool
}

func (v NullableGatewayInetConfig) Get() *GatewayInetConfig {
	return v.value
}

func (v *NullableGatewayInetConfig) Set(val *GatewayInetConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayInetConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayInetConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayInetConfig(val *GatewayInetConfig) *NullableGatewayInetConfig {
	return &NullableGatewayInetConfig{value: val, isSet: true}
}

func (v NullableGatewayInetConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayInetConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


