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

// checks if the UpdateDeviceBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateDeviceBody{}

// UpdateDeviceBody Updating the device
type UpdateDeviceBody struct {
	// The collection id for the device. This field is optional and can be omitted if the collection id isn't changed. When changing to a new collection you must be an owner of the other collection, ie an administrator of the team that owns the new collection.
	CollectionId *string `json:"collectionId,omitempty"`
	// Tags are metadata for the device that you can set. These are just strings.
	Tags *map[string]string `json:"tags,omitempty"`
	Firmware *FirmwareMetadata `json:"firmware,omitempty"`
	Config *DeviceConfig `json:"config,omitempty"`
	Enabled *bool `json:"enabled,omitempty"`
}

// NewUpdateDeviceBody instantiates a new UpdateDeviceBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateDeviceBody() *UpdateDeviceBody {
	this := UpdateDeviceBody{}
	return &this
}

// NewUpdateDeviceBodyWithDefaults instantiates a new UpdateDeviceBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateDeviceBodyWithDefaults() *UpdateDeviceBody {
	this := UpdateDeviceBody{}
	return &this
}

// GetCollectionId returns the CollectionId field value if set, zero value otherwise.
func (o *UpdateDeviceBody) GetCollectionId() string {
	if o == nil || IsNil(o.CollectionId) {
		var ret string
		return ret
	}
	return *o.CollectionId
}

// GetCollectionIdOk returns a tuple with the CollectionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceBody) GetCollectionIdOk() (*string, bool) {
	if o == nil || IsNil(o.CollectionId) {
		return nil, false
	}
	return o.CollectionId, true
}

// HasCollectionId returns a boolean if a field has been set.
func (o *UpdateDeviceBody) HasCollectionId() bool {
	if o != nil && !IsNil(o.CollectionId) {
		return true
	}

	return false
}

// SetCollectionId gets a reference to the given string and assigns it to the CollectionId field.
func (o *UpdateDeviceBody) SetCollectionId(v string) {
	o.CollectionId = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *UpdateDeviceBody) GetTags() map[string]string {
	if o == nil || IsNil(o.Tags) {
		var ret map[string]string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceBody) GetTagsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *UpdateDeviceBody) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given map[string]string and assigns it to the Tags field.
func (o *UpdateDeviceBody) SetTags(v map[string]string) {
	o.Tags = &v
}

// GetFirmware returns the Firmware field value if set, zero value otherwise.
func (o *UpdateDeviceBody) GetFirmware() FirmwareMetadata {
	if o == nil || IsNil(o.Firmware) {
		var ret FirmwareMetadata
		return ret
	}
	return *o.Firmware
}

// GetFirmwareOk returns a tuple with the Firmware field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceBody) GetFirmwareOk() (*FirmwareMetadata, bool) {
	if o == nil || IsNil(o.Firmware) {
		return nil, false
	}
	return o.Firmware, true
}

// HasFirmware returns a boolean if a field has been set.
func (o *UpdateDeviceBody) HasFirmware() bool {
	if o != nil && !IsNil(o.Firmware) {
		return true
	}

	return false
}

// SetFirmware gets a reference to the given FirmwareMetadata and assigns it to the Firmware field.
func (o *UpdateDeviceBody) SetFirmware(v FirmwareMetadata) {
	o.Firmware = &v
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *UpdateDeviceBody) GetConfig() DeviceConfig {
	if o == nil || IsNil(o.Config) {
		var ret DeviceConfig
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceBody) GetConfigOk() (*DeviceConfig, bool) {
	if o == nil || IsNil(o.Config) {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *UpdateDeviceBody) HasConfig() bool {
	if o != nil && !IsNil(o.Config) {
		return true
	}

	return false
}

// SetConfig gets a reference to the given DeviceConfig and assigns it to the Config field.
func (o *UpdateDeviceBody) SetConfig(v DeviceConfig) {
	o.Config = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *UpdateDeviceBody) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceBody) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *UpdateDeviceBody) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *UpdateDeviceBody) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o UpdateDeviceBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateDeviceBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CollectionId) {
		toSerialize["collectionId"] = o.CollectionId
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.Firmware) {
		toSerialize["firmware"] = o.Firmware
	}
	if !IsNil(o.Config) {
		toSerialize["config"] = o.Config
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return toSerialize, nil
}

type NullableUpdateDeviceBody struct {
	value *UpdateDeviceBody
	isSet bool
}

func (v NullableUpdateDeviceBody) Get() *UpdateDeviceBody {
	return v.value
}

func (v *NullableUpdateDeviceBody) Set(val *UpdateDeviceBody) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateDeviceBody) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateDeviceBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateDeviceBody(val *UpdateDeviceBody) *NullableUpdateDeviceBody {
	return &NullableUpdateDeviceBody{value: val, isSet: true}
}

func (v NullableUpdateDeviceBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateDeviceBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


