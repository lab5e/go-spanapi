/*
The Span API

API for device, collection, output and firmware management

API version: 5.0.1 humorous-jaron
Contact: dev@lab5e.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package spanapi

import (
	"encoding/json"
)

// checks if the CreateFirmwareBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateFirmwareBody{}

// CreateFirmwareBody Create a new firmware image
type CreateFirmwareBody struct {
	Image *string `json:"image,omitempty"`
	Version *string `json:"version,omitempty"`
	Filename *string `json:"filename,omitempty"`
	Tags *map[string]string `json:"tags,omitempty"`
}

// NewCreateFirmwareBody instantiates a new CreateFirmwareBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateFirmwareBody() *CreateFirmwareBody {
	this := CreateFirmwareBody{}
	return &this
}

// NewCreateFirmwareBodyWithDefaults instantiates a new CreateFirmwareBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateFirmwareBodyWithDefaults() *CreateFirmwareBody {
	this := CreateFirmwareBody{}
	return &this
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *CreateFirmwareBody) GetImage() string {
	if o == nil || IsNil(o.Image) {
		var ret string
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateFirmwareBody) GetImageOk() (*string, bool) {
	if o == nil || IsNil(o.Image) {
		return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *CreateFirmwareBody) HasImage() bool {
	if o != nil && !IsNil(o.Image) {
		return true
	}

	return false
}

// SetImage gets a reference to the given string and assigns it to the Image field.
func (o *CreateFirmwareBody) SetImage(v string) {
	o.Image = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *CreateFirmwareBody) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateFirmwareBody) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *CreateFirmwareBody) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *CreateFirmwareBody) SetVersion(v string) {
	o.Version = &v
}

// GetFilename returns the Filename field value if set, zero value otherwise.
func (o *CreateFirmwareBody) GetFilename() string {
	if o == nil || IsNil(o.Filename) {
		var ret string
		return ret
	}
	return *o.Filename
}

// GetFilenameOk returns a tuple with the Filename field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateFirmwareBody) GetFilenameOk() (*string, bool) {
	if o == nil || IsNil(o.Filename) {
		return nil, false
	}
	return o.Filename, true
}

// HasFilename returns a boolean if a field has been set.
func (o *CreateFirmwareBody) HasFilename() bool {
	if o != nil && !IsNil(o.Filename) {
		return true
	}

	return false
}

// SetFilename gets a reference to the given string and assigns it to the Filename field.
func (o *CreateFirmwareBody) SetFilename(v string) {
	o.Filename = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *CreateFirmwareBody) GetTags() map[string]string {
	if o == nil || IsNil(o.Tags) {
		var ret map[string]string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateFirmwareBody) GetTagsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *CreateFirmwareBody) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given map[string]string and assigns it to the Tags field.
func (o *CreateFirmwareBody) SetTags(v map[string]string) {
	o.Tags = &v
}

func (o CreateFirmwareBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateFirmwareBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Image) {
		toSerialize["image"] = o.Image
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !IsNil(o.Filename) {
		toSerialize["filename"] = o.Filename
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	return toSerialize, nil
}

type NullableCreateFirmwareBody struct {
	value *CreateFirmwareBody
	isSet bool
}

func (v NullableCreateFirmwareBody) Get() *CreateFirmwareBody {
	return v.value
}

func (v *NullableCreateFirmwareBody) Set(val *CreateFirmwareBody) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateFirmwareBody) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateFirmwareBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateFirmwareBody(val *CreateFirmwareBody) *NullableCreateFirmwareBody {
	return &NullableCreateFirmwareBody{value: val, isSet: true}
}

func (v NullableCreateFirmwareBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateFirmwareBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


