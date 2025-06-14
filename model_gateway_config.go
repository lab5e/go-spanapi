/*
The Span API

API for device, collection, output and firmware management

API version: 5.0.5 contented-jamila
Contact: dev@lab5e.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package spanapi

import (
	"encoding/json"
)

// checks if the GatewayConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GatewayConfig{}

// GatewayConfig struct for GatewayConfig
type GatewayConfig struct {
	Ciot *GatewayCIoTConfig `json:"ciot,omitempty"`
	Inet *GatewayInetConfig `json:"inet,omitempty"`
	User *GatewayCustomConfig `json:"user,omitempty"`
}

// NewGatewayConfig instantiates a new GatewayConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayConfig() *GatewayConfig {
	this := GatewayConfig{}
	return &this
}

// NewGatewayConfigWithDefaults instantiates a new GatewayConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayConfigWithDefaults() *GatewayConfig {
	this := GatewayConfig{}
	return &this
}

// GetCiot returns the Ciot field value if set, zero value otherwise.
func (o *GatewayConfig) GetCiot() GatewayCIoTConfig {
	if o == nil || IsNil(o.Ciot) {
		var ret GatewayCIoTConfig
		return ret
	}
	return *o.Ciot
}

// GetCiotOk returns a tuple with the Ciot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayConfig) GetCiotOk() (*GatewayCIoTConfig, bool) {
	if o == nil || IsNil(o.Ciot) {
		return nil, false
	}
	return o.Ciot, true
}

// HasCiot returns a boolean if a field has been set.
func (o *GatewayConfig) HasCiot() bool {
	if o != nil && !IsNil(o.Ciot) {
		return true
	}

	return false
}

// SetCiot gets a reference to the given GatewayCIoTConfig and assigns it to the Ciot field.
func (o *GatewayConfig) SetCiot(v GatewayCIoTConfig) {
	o.Ciot = &v
}

// GetInet returns the Inet field value if set, zero value otherwise.
func (o *GatewayConfig) GetInet() GatewayInetConfig {
	if o == nil || IsNil(o.Inet) {
		var ret GatewayInetConfig
		return ret
	}
	return *o.Inet
}

// GetInetOk returns a tuple with the Inet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayConfig) GetInetOk() (*GatewayInetConfig, bool) {
	if o == nil || IsNil(o.Inet) {
		return nil, false
	}
	return o.Inet, true
}

// HasInet returns a boolean if a field has been set.
func (o *GatewayConfig) HasInet() bool {
	if o != nil && !IsNil(o.Inet) {
		return true
	}

	return false
}

// SetInet gets a reference to the given GatewayInetConfig and assigns it to the Inet field.
func (o *GatewayConfig) SetInet(v GatewayInetConfig) {
	o.Inet = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *GatewayConfig) GetUser() GatewayCustomConfig {
	if o == nil || IsNil(o.User) {
		var ret GatewayCustomConfig
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayConfig) GetUserOk() (*GatewayCustomConfig, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *GatewayConfig) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given GatewayCustomConfig and assigns it to the User field.
func (o *GatewayConfig) SetUser(v GatewayCustomConfig) {
	o.User = &v
}

func (o GatewayConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GatewayConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ciot) {
		toSerialize["ciot"] = o.Ciot
	}
	if !IsNil(o.Inet) {
		toSerialize["inet"] = o.Inet
	}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	return toSerialize, nil
}

type NullableGatewayConfig struct {
	value *GatewayConfig
	isSet bool
}

func (v NullableGatewayConfig) Get() *GatewayConfig {
	return v.value
}

func (v *NullableGatewayConfig) Set(val *GatewayConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayConfig(val *GatewayConfig) *NullableGatewayConfig {
	return &NullableGatewayConfig{value: val, isSet: true}
}

func (v NullableGatewayConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


