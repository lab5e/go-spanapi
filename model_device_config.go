/*
The Span API

API for device, collection, output and firmware management

API version: 4.3.0 grouchy-aloysius
Contact: dev@lab5e.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package spanapi

import (
	"encoding/json"
)

// DeviceConfig This is the configuration for the device via the various gateways.
type DeviceConfig struct {
	Ciot *CellularIoTConfig `json:"ciot,omitempty"`
}

// NewDeviceConfig instantiates a new DeviceConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceConfig() *DeviceConfig {
	this := DeviceConfig{}
	return &this
}

// NewDeviceConfigWithDefaults instantiates a new DeviceConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceConfigWithDefaults() *DeviceConfig {
	this := DeviceConfig{}
	return &this
}

// GetCiot returns the Ciot field value if set, zero value otherwise.
func (o *DeviceConfig) GetCiot() CellularIoTConfig {
	if o == nil || o.Ciot == nil {
		var ret CellularIoTConfig
		return ret
	}
	return *o.Ciot
}

// GetCiotOk returns a tuple with the Ciot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceConfig) GetCiotOk() (*CellularIoTConfig, bool) {
	if o == nil || o.Ciot == nil {
		return nil, false
	}
	return o.Ciot, true
}

// HasCiot returns a boolean if a field has been set.
func (o *DeviceConfig) HasCiot() bool {
	if o != nil && o.Ciot != nil {
		return true
	}

	return false
}

// SetCiot gets a reference to the given CellularIoTConfig and assigns it to the Ciot field.
func (o *DeviceConfig) SetCiot(v CellularIoTConfig) {
	o.Ciot = &v
}

func (o DeviceConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Ciot != nil {
		toSerialize["ciot"] = o.Ciot
	}
	return json.Marshal(toSerialize)
}

type NullableDeviceConfig struct {
	value *DeviceConfig
	isSet bool
}

func (v NullableDeviceConfig) Get() *DeviceConfig {
	return v.value
}

func (v *NullableDeviceConfig) Set(val *DeviceConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceConfig(val *DeviceConfig) *NullableDeviceConfig {
	return &NullableDeviceConfig{value: val, isSet: true}
}

func (v NullableDeviceConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


