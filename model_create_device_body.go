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

// checks if the CreateDeviceBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateDeviceBody{}

// CreateDeviceBody Request object to create new devices
type CreateDeviceBody struct {
	// Tags are metadata for the device that you can set. These are just strings.
	Tags *map[string]string `json:"tags,omitempty"`
	Firmware *FirmwareMetadata `json:"firmware,omitempty"`
	Config *DeviceConfig `json:"config,omitempty"`
	Metadata *DeviceMetadata `json:"metadata,omitempty"`
}

// NewCreateDeviceBody instantiates a new CreateDeviceBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateDeviceBody() *CreateDeviceBody {
	this := CreateDeviceBody{}
	return &this
}

// NewCreateDeviceBodyWithDefaults instantiates a new CreateDeviceBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateDeviceBodyWithDefaults() *CreateDeviceBody {
	this := CreateDeviceBody{}
	return &this
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *CreateDeviceBody) GetTags() map[string]string {
	if o == nil || IsNil(o.Tags) {
		var ret map[string]string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDeviceBody) GetTagsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *CreateDeviceBody) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given map[string]string and assigns it to the Tags field.
func (o *CreateDeviceBody) SetTags(v map[string]string) {
	o.Tags = &v
}

// GetFirmware returns the Firmware field value if set, zero value otherwise.
func (o *CreateDeviceBody) GetFirmware() FirmwareMetadata {
	if o == nil || IsNil(o.Firmware) {
		var ret FirmwareMetadata
		return ret
	}
	return *o.Firmware
}

// GetFirmwareOk returns a tuple with the Firmware field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDeviceBody) GetFirmwareOk() (*FirmwareMetadata, bool) {
	if o == nil || IsNil(o.Firmware) {
		return nil, false
	}
	return o.Firmware, true
}

// HasFirmware returns a boolean if a field has been set.
func (o *CreateDeviceBody) HasFirmware() bool {
	if o != nil && !IsNil(o.Firmware) {
		return true
	}

	return false
}

// SetFirmware gets a reference to the given FirmwareMetadata and assigns it to the Firmware field.
func (o *CreateDeviceBody) SetFirmware(v FirmwareMetadata) {
	o.Firmware = &v
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *CreateDeviceBody) GetConfig() DeviceConfig {
	if o == nil || IsNil(o.Config) {
		var ret DeviceConfig
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDeviceBody) GetConfigOk() (*DeviceConfig, bool) {
	if o == nil || IsNil(o.Config) {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *CreateDeviceBody) HasConfig() bool {
	if o != nil && !IsNil(o.Config) {
		return true
	}

	return false
}

// SetConfig gets a reference to the given DeviceConfig and assigns it to the Config field.
func (o *CreateDeviceBody) SetConfig(v DeviceConfig) {
	o.Config = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *CreateDeviceBody) GetMetadata() DeviceMetadata {
	if o == nil || IsNil(o.Metadata) {
		var ret DeviceMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDeviceBody) GetMetadataOk() (*DeviceMetadata, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *CreateDeviceBody) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given DeviceMetadata and assigns it to the Metadata field.
func (o *CreateDeviceBody) SetMetadata(v DeviceMetadata) {
	o.Metadata = &v
}

func (o CreateDeviceBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateDeviceBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.Firmware) {
		toSerialize["firmware"] = o.Firmware
	}
	if !IsNil(o.Config) {
		toSerialize["config"] = o.Config
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	return toSerialize, nil
}

type NullableCreateDeviceBody struct {
	value *CreateDeviceBody
	isSet bool
}

func (v NullableCreateDeviceBody) Get() *CreateDeviceBody {
	return v.value
}

func (v *NullableCreateDeviceBody) Set(val *CreateDeviceBody) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateDeviceBody) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateDeviceBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateDeviceBody(val *CreateDeviceBody) *NullableCreateDeviceBody {
	return &NullableCreateDeviceBody{value: val, isSet: true}
}

func (v NullableCreateDeviceBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateDeviceBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


