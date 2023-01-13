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

// DeviceMetadata This is the metadata for devices.
type DeviceMetadata struct {
	SimOperator *NetworkOperator     `json:"simOperator,omitempty"`
	Ciot        *CellularIoTMetadata `json:"ciot,omitempty"`
	Inet        *InetMetadata        `json:"inet,omitempty"`
}

// NewDeviceMetadata instantiates a new DeviceMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceMetadata() *DeviceMetadata {
	this := DeviceMetadata{}
	return &this
}

// NewDeviceMetadataWithDefaults instantiates a new DeviceMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceMetadataWithDefaults() *DeviceMetadata {
	this := DeviceMetadata{}
	return &this
}

// GetSimOperator returns the SimOperator field value if set, zero value otherwise.
func (o *DeviceMetadata) GetSimOperator() NetworkOperator {
	if o == nil || o.SimOperator == nil {
		var ret NetworkOperator
		return ret
	}
	return *o.SimOperator
}

// GetSimOperatorOk returns a tuple with the SimOperator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceMetadata) GetSimOperatorOk() (*NetworkOperator, bool) {
	if o == nil || o.SimOperator == nil {
		return nil, false
	}
	return o.SimOperator, true
}

// HasSimOperator returns a boolean if a field has been set.
func (o *DeviceMetadata) HasSimOperator() bool {
	if o != nil && o.SimOperator != nil {
		return true
	}

	return false
}

// SetSimOperator gets a reference to the given NetworkOperator and assigns it to the SimOperator field.
func (o *DeviceMetadata) SetSimOperator(v NetworkOperator) {
	o.SimOperator = &v
}

// GetCiot returns the Ciot field value if set, zero value otherwise.
func (o *DeviceMetadata) GetCiot() CellularIoTMetadata {
	if o == nil || o.Ciot == nil {
		var ret CellularIoTMetadata
		return ret
	}
	return *o.Ciot
}

// GetCiotOk returns a tuple with the Ciot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceMetadata) GetCiotOk() (*CellularIoTMetadata, bool) {
	if o == nil || o.Ciot == nil {
		return nil, false
	}
	return o.Ciot, true
}

// HasCiot returns a boolean if a field has been set.
func (o *DeviceMetadata) HasCiot() bool {
	if o != nil && o.Ciot != nil {
		return true
	}

	return false
}

// SetCiot gets a reference to the given CellularIoTMetadata and assigns it to the Ciot field.
func (o *DeviceMetadata) SetCiot(v CellularIoTMetadata) {
	o.Ciot = &v
}

// GetInet returns the Inet field value if set, zero value otherwise.
func (o *DeviceMetadata) GetInet() InetMetadata {
	if o == nil || o.Inet == nil {
		var ret InetMetadata
		return ret
	}
	return *o.Inet
}

// GetInetOk returns a tuple with the Inet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceMetadata) GetInetOk() (*InetMetadata, bool) {
	if o == nil || o.Inet == nil {
		return nil, false
	}
	return o.Inet, true
}

// HasInet returns a boolean if a field has been set.
func (o *DeviceMetadata) HasInet() bool {
	if o != nil && o.Inet != nil {
		return true
	}

	return false
}

// SetInet gets a reference to the given InetMetadata and assigns it to the Inet field.
func (o *DeviceMetadata) SetInet(v InetMetadata) {
	o.Inet = &v
}

func (o DeviceMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SimOperator != nil {
		toSerialize["simOperator"] = o.SimOperator
	}
	if o.Ciot != nil {
		toSerialize["ciot"] = o.Ciot
	}
	if o.Inet != nil {
		toSerialize["inet"] = o.Inet
	}
	return json.Marshal(toSerialize)
}

type NullableDeviceMetadata struct {
	value *DeviceMetadata
	isSet bool
}

func (v NullableDeviceMetadata) Get() *DeviceMetadata {
	return v.value
}

func (v *NullableDeviceMetadata) Set(val *DeviceMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceMetadata(val *DeviceMetadata) *NullableDeviceMetadata {
	return &NullableDeviceMetadata{value: val, isSet: true}
}

func (v NullableDeviceMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
