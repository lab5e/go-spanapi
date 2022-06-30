/*
The Span API

API for device, collection, output and firmware management

API version: 4.2.4 curable-andres
Contact: dev@lab5e.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package spanapi

import (
	"encoding/json"
)

// CreateFirmwareRequest Create a new firmware image
type CreateFirmwareRequest struct {
	Image *string `json:"image,omitempty"`
	Version *string `json:"version,omitempty"`
	Filename *string `json:"filename,omitempty"`
	Tags *map[string]string `json:"tags,omitempty"`
}

// NewCreateFirmwareRequest instantiates a new CreateFirmwareRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateFirmwareRequest() *CreateFirmwareRequest {
	this := CreateFirmwareRequest{}
	return &this
}

// NewCreateFirmwareRequestWithDefaults instantiates a new CreateFirmwareRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateFirmwareRequestWithDefaults() *CreateFirmwareRequest {
	this := CreateFirmwareRequest{}
	return &this
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *CreateFirmwareRequest) GetImage() string {
	if o == nil || o.Image == nil {
		var ret string
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateFirmwareRequest) GetImageOk() (*string, bool) {
	if o == nil || o.Image == nil {
		return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *CreateFirmwareRequest) HasImage() bool {
	if o != nil && o.Image != nil {
		return true
	}

	return false
}

// SetImage gets a reference to the given string and assigns it to the Image field.
func (o *CreateFirmwareRequest) SetImage(v string) {
	o.Image = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *CreateFirmwareRequest) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateFirmwareRequest) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *CreateFirmwareRequest) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *CreateFirmwareRequest) SetVersion(v string) {
	o.Version = &v
}

// GetFilename returns the Filename field value if set, zero value otherwise.
func (o *CreateFirmwareRequest) GetFilename() string {
	if o == nil || o.Filename == nil {
		var ret string
		return ret
	}
	return *o.Filename
}

// GetFilenameOk returns a tuple with the Filename field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateFirmwareRequest) GetFilenameOk() (*string, bool) {
	if o == nil || o.Filename == nil {
		return nil, false
	}
	return o.Filename, true
}

// HasFilename returns a boolean if a field has been set.
func (o *CreateFirmwareRequest) HasFilename() bool {
	if o != nil && o.Filename != nil {
		return true
	}

	return false
}

// SetFilename gets a reference to the given string and assigns it to the Filename field.
func (o *CreateFirmwareRequest) SetFilename(v string) {
	o.Filename = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *CreateFirmwareRequest) GetTags() map[string]string {
	if o == nil || o.Tags == nil {
		var ret map[string]string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateFirmwareRequest) GetTagsOk() (*map[string]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *CreateFirmwareRequest) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given map[string]string and assigns it to the Tags field.
func (o *CreateFirmwareRequest) SetTags(v map[string]string) {
	o.Tags = &v
}

func (o CreateFirmwareRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Image != nil {
		toSerialize["image"] = o.Image
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.Filename != nil {
		toSerialize["filename"] = o.Filename
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	return json.Marshal(toSerialize)
}

type NullableCreateFirmwareRequest struct {
	value *CreateFirmwareRequest
	isSet bool
}

func (v NullableCreateFirmwareRequest) Get() *CreateFirmwareRequest {
	return v.value
}

func (v *NullableCreateFirmwareRequest) Set(val *CreateFirmwareRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateFirmwareRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateFirmwareRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateFirmwareRequest(val *CreateFirmwareRequest) *NullableCreateFirmwareRequest {
	return &NullableCreateFirmwareRequest{value: val, isSet: true}
}

func (v NullableCreateFirmwareRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateFirmwareRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


