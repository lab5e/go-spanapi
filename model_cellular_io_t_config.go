/*
The Span API

API for device, collection, output and firmware management

API version: 5.0.2 subversive-jamila
Contact: dev@lab5e.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package spanapi

import (
	"encoding/json"
)

// checks if the CellularIoTConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CellularIoTConfig{}

// CellularIoTConfig This is the cellular IOT config
type CellularIoTConfig struct {
	Imsi *string `json:"imsi,omitempty"`
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
	if o == nil || IsNil(o.Imsi) {
		var ret string
		return ret
	}
	return *o.Imsi
}

// GetImsiOk returns a tuple with the Imsi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CellularIoTConfig) GetImsiOk() (*string, bool) {
	if o == nil || IsNil(o.Imsi) {
		return nil, false
	}
	return o.Imsi, true
}

// HasImsi returns a boolean if a field has been set.
func (o *CellularIoTConfig) HasImsi() bool {
	if o != nil && !IsNil(o.Imsi) {
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
	if o == nil || IsNil(o.Imei) {
		var ret string
		return ret
	}
	return *o.Imei
}

// GetImeiOk returns a tuple with the Imei field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CellularIoTConfig) GetImeiOk() (*string, bool) {
	if o == nil || IsNil(o.Imei) {
		return nil, false
	}
	return o.Imei, true
}

// HasImei returns a boolean if a field has been set.
func (o *CellularIoTConfig) HasImei() bool {
	if o != nil && !IsNil(o.Imei) {
		return true
	}

	return false
}

// SetImei gets a reference to the given string and assigns it to the Imei field.
func (o *CellularIoTConfig) SetImei(v string) {
	o.Imei = &v
}

func (o CellularIoTConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CellularIoTConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Imsi) {
		toSerialize["imsi"] = o.Imsi
	}
	if !IsNil(o.Imei) {
		toSerialize["imei"] = o.Imei
	}
	return toSerialize, nil
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


