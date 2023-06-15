/*
The Span API

API for device, collection, output and firmware management

API version: 4.6.0 truthful-holli
Contact: dev@lab5e.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package spanapi

import (
	"encoding/json"
)

// checks if the FirmwareStats type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FirmwareStats{}

// FirmwareStats Statistics for a single firmware image
type FirmwareStats struct {
	FirmwareImageSize *int32 `json:"firmwareImageSize,omitempty"`
}

// NewFirmwareStats instantiates a new FirmwareStats object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFirmwareStats() *FirmwareStats {
	this := FirmwareStats{}
	return &this
}

// NewFirmwareStatsWithDefaults instantiates a new FirmwareStats object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFirmwareStatsWithDefaults() *FirmwareStats {
	this := FirmwareStats{}
	return &this
}

// GetFirmwareImageSize returns the FirmwareImageSize field value if set, zero value otherwise.
func (o *FirmwareStats) GetFirmwareImageSize() int32 {
	if o == nil || IsNil(o.FirmwareImageSize) {
		var ret int32
		return ret
	}
	return *o.FirmwareImageSize
}

// GetFirmwareImageSizeOk returns a tuple with the FirmwareImageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareStats) GetFirmwareImageSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.FirmwareImageSize) {
		return nil, false
	}
	return o.FirmwareImageSize, true
}

// HasFirmwareImageSize returns a boolean if a field has been set.
func (o *FirmwareStats) HasFirmwareImageSize() bool {
	if o != nil && !IsNil(o.FirmwareImageSize) {
		return true
	}

	return false
}

// SetFirmwareImageSize gets a reference to the given int32 and assigns it to the FirmwareImageSize field.
func (o *FirmwareStats) SetFirmwareImageSize(v int32) {
	o.FirmwareImageSize = &v
}

func (o FirmwareStats) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FirmwareStats) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FirmwareImageSize) {
		toSerialize["firmwareImageSize"] = o.FirmwareImageSize
	}
	return toSerialize, nil
}

type NullableFirmwareStats struct {
	value *FirmwareStats
	isSet bool
}

func (v NullableFirmwareStats) Get() *FirmwareStats {
	return v.value
}

func (v *NullableFirmwareStats) Set(val *FirmwareStats) {
	v.value = val
	v.isSet = true
}

func (v NullableFirmwareStats) IsSet() bool {
	return v.isSet
}

func (v *NullableFirmwareStats) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFirmwareStats(val *FirmwareStats) *NullableFirmwareStats {
	return &NullableFirmwareStats{value: val, isSet: true}
}

func (v NullableFirmwareStats) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFirmwareStats) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


