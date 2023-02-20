/*
The Span API

API for device, collection, output and firmware management

API version: 4.4.0 lean-joline
Contact: dev@lab5e.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package spanapi

import (
	"encoding/json"
)

// CellularIoTConfig This is the cellular IOT config
type CellularIoTConfig struct {
	Imsi *string `json:"imsi,omitempty"`
	// on your device. This is the primary identifier for your device on the network.  The IMEI number is the unique ID for your hardware as
	Imei *string `json:"imei,omitempty"`
}

// NewCellularIoTConfig instantiates a new CellularIoTConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCellularIoTConfig() *CellularIoTConfig {
	this := CellularIoTConfig{}
	return &this
}

// NewCellularIoTConfigWithDefaults instantiates a new CellularIoTConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCellularIoTConfigWithDefaults() *CellularIoTConfig {
	this := CellularIoTConfig{}
	return &this
}

// GetImsi returns the Imsi field value if set, zero value otherwise.
func (o *CellularIoTConfig) GetImsi() string {
	if o == nil || o.Imsi == nil {
		var ret string
		return ret
	}
	return *o.Imsi
}

// GetImsiOk returns a tuple with the Imsi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CellularIoTConfig) GetImsiOk() (*string, bool) {
	if o == nil || o.Imsi == nil {
		return nil, false
	}
	return o.Imsi, true
}

// HasImsi returns a boolean if a field has been set.
func (o *CellularIoTConfig) HasImsi() bool {
	if o != nil && o.Imsi != nil {
		return true
	}

	return false
}

// SetImsi gets a reference to the given string and assigns it to the Imsi field.
func (o *CellularIoTConfig) SetImsi(v string) {
	o.Imsi = &v
}

// GetImei returns the Imei field value if set, zero value otherwise.
func (o *CellularIoTConfig) GetImei() string {
	if o == nil || o.Imei == nil {
		var ret string
		return ret
	}
	return *o.Imei
}

// GetImeiOk returns a tuple with the Imei field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CellularIoTConfig) GetImeiOk() (*string, bool) {
	if o == nil || o.Imei == nil {
		return nil, false
	}
	return o.Imei, true
}

// HasImei returns a boolean if a field has been set.
func (o *CellularIoTConfig) HasImei() bool {
	if o != nil && o.Imei != nil {
		return true
	}

	return false
}

// SetImei gets a reference to the given string and assigns it to the Imei field.
func (o *CellularIoTConfig) SetImei(v string) {
	o.Imei = &v
}

func (o CellularIoTConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Imsi != nil {
		toSerialize["imsi"] = o.Imsi
	}
	if o.Imei != nil {
		toSerialize["imei"] = o.Imei
	}
	return json.Marshal(toSerialize)
}

type NullableCellularIoTConfig struct {
	value *CellularIoTConfig
	isSet bool
}

func (v NullableCellularIoTConfig) Get() *CellularIoTConfig {
	return v.value
}

func (v *NullableCellularIoTConfig) Set(val *CellularIoTConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableCellularIoTConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableCellularIoTConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCellularIoTConfig(val *CellularIoTConfig) *NullableCellularIoTConfig {
	return &NullableCellularIoTConfig{value: val, isSet: true}
}

func (v NullableCellularIoTConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCellularIoTConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
