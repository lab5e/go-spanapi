/*
The Span API

API for device, collection, output and firmware management

API version: 4.9.6 authoritarian-betty
Contact: dev@lab5e.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package spanapi

import (
	"encoding/json"
)

// checks if the GatewayCIoTConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GatewayCIoTConfig{}

// GatewayCIoTConfig struct for GatewayCIoTConfig
type GatewayCIoTConfig struct {
	Apn *string `json:"apn,omitempty"`
	UdpEndpoint *string `json:"udpEndpoint,omitempty"`
	CoapEndpoint *string `json:"coapEndpoint,omitempty"`
}

// NewGatewayCIoTConfig instantiates a new GatewayCIoTConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayCIoTConfig() *GatewayCIoTConfig {
	this := GatewayCIoTConfig{}
	return &this
}

// NewGatewayCIoTConfigWithDefaults instantiates a new GatewayCIoTConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayCIoTConfigWithDefaults() *GatewayCIoTConfig {
	this := GatewayCIoTConfig{}
	return &this
}

// GetApn returns the Apn field value if set, zero value otherwise.
func (o *GatewayCIoTConfig) GetApn() string {
	if o == nil || IsNil(o.Apn) {
		var ret string
		return ret
	}
	return *o.Apn
}

// GetApnOk returns a tuple with the Apn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCIoTConfig) GetApnOk() (*string, bool) {
	if o == nil || IsNil(o.Apn) {
		return nil, false
	}
	return o.Apn, true
}

// HasApn returns a boolean if a field has been set.
func (o *GatewayCIoTConfig) HasApn() bool {
	if o != nil && !IsNil(o.Apn) {
		return true
	}

	return false
}

// SetApn gets a reference to the given string and assigns it to the Apn field.
func (o *GatewayCIoTConfig) SetApn(v string) {
	o.Apn = &v
}

// GetUdpEndpoint returns the UdpEndpoint field value if set, zero value otherwise.
func (o *GatewayCIoTConfig) GetUdpEndpoint() string {
	if o == nil || IsNil(o.UdpEndpoint) {
		var ret string
		return ret
	}
	return *o.UdpEndpoint
}

// GetUdpEndpointOk returns a tuple with the UdpEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCIoTConfig) GetUdpEndpointOk() (*string, bool) {
	if o == nil || IsNil(o.UdpEndpoint) {
		return nil, false
	}
	return o.UdpEndpoint, true
}

// HasUdpEndpoint returns a boolean if a field has been set.
func (o *GatewayCIoTConfig) HasUdpEndpoint() bool {
	if o != nil && !IsNil(o.UdpEndpoint) {
		return true
	}

	return false
}

// SetUdpEndpoint gets a reference to the given string and assigns it to the UdpEndpoint field.
func (o *GatewayCIoTConfig) SetUdpEndpoint(v string) {
	o.UdpEndpoint = &v
}

// GetCoapEndpoint returns the CoapEndpoint field value if set, zero value otherwise.
func (o *GatewayCIoTConfig) GetCoapEndpoint() string {
	if o == nil || IsNil(o.CoapEndpoint) {
		var ret string
		return ret
	}
	return *o.CoapEndpoint
}

// GetCoapEndpointOk returns a tuple with the CoapEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCIoTConfig) GetCoapEndpointOk() (*string, bool) {
	if o == nil || IsNil(o.CoapEndpoint) {
		return nil, false
	}
	return o.CoapEndpoint, true
}

// HasCoapEndpoint returns a boolean if a field has been set.
func (o *GatewayCIoTConfig) HasCoapEndpoint() bool {
	if o != nil && !IsNil(o.CoapEndpoint) {
		return true
	}

	return false
}

// SetCoapEndpoint gets a reference to the given string and assigns it to the CoapEndpoint field.
func (o *GatewayCIoTConfig) SetCoapEndpoint(v string) {
	o.CoapEndpoint = &v
}

func (o GatewayCIoTConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GatewayCIoTConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Apn) {
		toSerialize["apn"] = o.Apn
	}
	if !IsNil(o.UdpEndpoint) {
		toSerialize["udpEndpoint"] = o.UdpEndpoint
	}
	if !IsNil(o.CoapEndpoint) {
		toSerialize["coapEndpoint"] = o.CoapEndpoint
	}
	return toSerialize, nil
}

type NullableGatewayCIoTConfig struct {
	value *GatewayCIoTConfig
	isSet bool
}

func (v NullableGatewayCIoTConfig) Get() *GatewayCIoTConfig {
	return v.value
}

func (v *NullableGatewayCIoTConfig) Set(val *GatewayCIoTConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayCIoTConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayCIoTConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayCIoTConfig(val *GatewayCIoTConfig) *NullableGatewayCIoTConfig {
	return &NullableGatewayCIoTConfig{value: val, isSet: true}
}

func (v NullableGatewayCIoTConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayCIoTConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


