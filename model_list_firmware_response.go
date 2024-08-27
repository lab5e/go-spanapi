/*
The Span API

API for device, collection, output and firmware management

API version: 5.0.0 convulsive-launa
Contact: dev@lab5e.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package spanapi

import (
	"encoding/json"
)

// checks if the ListFirmwareResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListFirmwareResponse{}

// ListFirmwareResponse List firmware response
type ListFirmwareResponse struct {
	Images []Firmware `json:"images,omitempty"`
}

// NewListFirmwareResponse instantiates a new ListFirmwareResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListFirmwareResponse() *ListFirmwareResponse {
	this := ListFirmwareResponse{}
	return &this
}

// NewListFirmwareResponseWithDefaults instantiates a new ListFirmwareResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListFirmwareResponseWithDefaults() *ListFirmwareResponse {
	this := ListFirmwareResponse{}
	return &this
}

// GetImages returns the Images field value if set, zero value otherwise.
func (o *ListFirmwareResponse) GetImages() []Firmware {
	if o == nil || IsNil(o.Images) {
		var ret []Firmware
		return ret
	}
	return o.Images
}

// GetImagesOk returns a tuple with the Images field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListFirmwareResponse) GetImagesOk() ([]Firmware, bool) {
	if o == nil || IsNil(o.Images) {
		return nil, false
	}
	return o.Images, true
}

// HasImages returns a boolean if a field has been set.
func (o *ListFirmwareResponse) HasImages() bool {
	if o != nil && !IsNil(o.Images) {
		return true
	}

	return false
}

// SetImages gets a reference to the given []Firmware and assigns it to the Images field.
func (o *ListFirmwareResponse) SetImages(v []Firmware) {
	o.Images = v
}

func (o ListFirmwareResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListFirmwareResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Images) {
		toSerialize["images"] = o.Images
	}
	return toSerialize, nil
}

type NullableListFirmwareResponse struct {
	value *ListFirmwareResponse
	isSet bool
}

func (v NullableListFirmwareResponse) Get() *ListFirmwareResponse {
	return v.value
}

func (v *NullableListFirmwareResponse) Set(val *ListFirmwareResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListFirmwareResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListFirmwareResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListFirmwareResponse(val *ListFirmwareResponse) *NullableListFirmwareResponse {
	return &NullableListFirmwareResponse{value: val, isSet: true}
}

func (v NullableListFirmwareResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListFirmwareResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


