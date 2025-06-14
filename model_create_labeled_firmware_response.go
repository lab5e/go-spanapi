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

// checks if the CreateLabeledFirmwareResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateLabeledFirmwareResponse{}

// CreateLabeledFirmwareResponse Response object when creating a labeled firmware image
type CreateLabeledFirmwareResponse struct {
	Image *LabeledFirmware `json:"image,omitempty"`
}

// NewCreateLabeledFirmwareResponse instantiates a new CreateLabeledFirmwareResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateLabeledFirmwareResponse() *CreateLabeledFirmwareResponse {
	this := CreateLabeledFirmwareResponse{}
	return &this
}

// NewCreateLabeledFirmwareResponseWithDefaults instantiates a new CreateLabeledFirmwareResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateLabeledFirmwareResponseWithDefaults() *CreateLabeledFirmwareResponse {
	this := CreateLabeledFirmwareResponse{}
	return &this
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *CreateLabeledFirmwareResponse) GetImage() LabeledFirmware {
	if o == nil || IsNil(o.Image) {
		var ret LabeledFirmware
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLabeledFirmwareResponse) GetImageOk() (*LabeledFirmware, bool) {
	if o == nil || IsNil(o.Image) {
		return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *CreateLabeledFirmwareResponse) HasImage() bool {
	if o != nil && !IsNil(o.Image) {
		return true
	}

	return false
}

// SetImage gets a reference to the given LabeledFirmware and assigns it to the Image field.
func (o *CreateLabeledFirmwareResponse) SetImage(v LabeledFirmware) {
	o.Image = &v
}

func (o CreateLabeledFirmwareResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateLabeledFirmwareResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Image) {
		toSerialize["image"] = o.Image
	}
	return toSerialize, nil
}

type NullableCreateLabeledFirmwareResponse struct {
	value *CreateLabeledFirmwareResponse
	isSet bool
}

func (v NullableCreateLabeledFirmwareResponse) Get() *CreateLabeledFirmwareResponse {
	return v.value
}

func (v *NullableCreateLabeledFirmwareResponse) Set(val *CreateLabeledFirmwareResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateLabeledFirmwareResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateLabeledFirmwareResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateLabeledFirmwareResponse(val *CreateLabeledFirmwareResponse) *NullableCreateLabeledFirmwareResponse {
	return &NullableCreateLabeledFirmwareResponse{value: val, isSet: true}
}

func (v NullableCreateLabeledFirmwareResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateLabeledFirmwareResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


