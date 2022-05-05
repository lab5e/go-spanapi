/*
The Span API

API for device, collection, output and firmware management

API version: 4.2.2 smarmy-derik
Contact: dev@lab5e.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package spanapi

import (
	"encoding/json"
)

// UpdateDeviceRequest Updating the device
type UpdateDeviceRequest struct {
	// The collection id for the device. This field is optional and can be omitted if the collection id isn't changed. When changing to a new collection you must be an owner of the other collection, ie an administrator of the team that owns the new collection.
	CollectionId *string `json:"collectionId,omitempty"`
	// Deprecated: this is replaced by the Config type The IMSI is the unique ID for the (e|nu|whatever)SIM card on your device. This is the primary identifier for your device on the network.
	Imsi *string `json:"imsi,omitempty"`
	// Deprecated: This is replaced by the Config type The IMEI number is the unique ID for your hardware as seen by the network. Obviously you might have a completely different view on things.
	Imei *string `json:"imei,omitempty"`
	// Tags are metadata for the device that you can set. These are just strings.
	Tags *map[string]string `json:"tags,omitempty"`
	Firmware *FirmwareMetadata `json:"firmware,omitempty"`
	Config *DeviceConfig `json:"config,omitempty"`
}

// NewUpdateDeviceRequest instantiates a new UpdateDeviceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateDeviceRequest() *UpdateDeviceRequest {
	this := UpdateDeviceRequest{}
	return &this
}

// NewUpdateDeviceRequestWithDefaults instantiates a new UpdateDeviceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateDeviceRequestWithDefaults() *UpdateDeviceRequest {
	this := UpdateDeviceRequest{}
	return &this
}

// GetCollectionId returns the CollectionId field value if set, zero value otherwise.
func (o *UpdateDeviceRequest) GetCollectionId() string {
	if o == nil || o.CollectionId == nil {
		var ret string
		return ret
	}
	return *o.CollectionId
}

// GetCollectionIdOk returns a tuple with the CollectionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceRequest) GetCollectionIdOk() (*string, bool) {
	if o == nil || o.CollectionId == nil {
		return nil, false
	}
	return o.CollectionId, true
}

// HasCollectionId returns a boolean if a field has been set.
func (o *UpdateDeviceRequest) HasCollectionId() bool {
	if o != nil && o.CollectionId != nil {
		return true
	}

	return false
}

// SetCollectionId gets a reference to the given string and assigns it to the CollectionId field.
func (o *UpdateDeviceRequest) SetCollectionId(v string) {
	o.CollectionId = &v
}

// GetImsi returns the Imsi field value if set, zero value otherwise.
func (o *UpdateDeviceRequest) GetImsi() string {
	if o == nil || o.Imsi == nil {
		var ret string
		return ret
	}
	return *o.Imsi
}

// GetImsiOk returns a tuple with the Imsi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceRequest) GetImsiOk() (*string, bool) {
	if o == nil || o.Imsi == nil {
		return nil, false
	}
	return o.Imsi, true
}

// HasImsi returns a boolean if a field has been set.
func (o *UpdateDeviceRequest) HasImsi() bool {
	if o != nil && o.Imsi != nil {
		return true
	}

	return false
}

// SetImsi gets a reference to the given string and assigns it to the Imsi field.
func (o *UpdateDeviceRequest) SetImsi(v string) {
	o.Imsi = &v
}

// GetImei returns the Imei field value if set, zero value otherwise.
func (o *UpdateDeviceRequest) GetImei() string {
	if o == nil || o.Imei == nil {
		var ret string
		return ret
	}
	return *o.Imei
}

// GetImeiOk returns a tuple with the Imei field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceRequest) GetImeiOk() (*string, bool) {
	if o == nil || o.Imei == nil {
		return nil, false
	}
	return o.Imei, true
}

// HasImei returns a boolean if a field has been set.
func (o *UpdateDeviceRequest) HasImei() bool {
	if o != nil && o.Imei != nil {
		return true
	}

	return false
}

// SetImei gets a reference to the given string and assigns it to the Imei field.
func (o *UpdateDeviceRequest) SetImei(v string) {
	o.Imei = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *UpdateDeviceRequest) GetTags() map[string]string {
	if o == nil || o.Tags == nil {
		var ret map[string]string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceRequest) GetTagsOk() (*map[string]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *UpdateDeviceRequest) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given map[string]string and assigns it to the Tags field.
func (o *UpdateDeviceRequest) SetTags(v map[string]string) {
	o.Tags = &v
}

// GetFirmware returns the Firmware field value if set, zero value otherwise.
func (o *UpdateDeviceRequest) GetFirmware() FirmwareMetadata {
	if o == nil || o.Firmware == nil {
		var ret FirmwareMetadata
		return ret
	}
	return *o.Firmware
}

// GetFirmwareOk returns a tuple with the Firmware field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceRequest) GetFirmwareOk() (*FirmwareMetadata, bool) {
	if o == nil || o.Firmware == nil {
		return nil, false
	}
	return o.Firmware, true
}

// HasFirmware returns a boolean if a field has been set.
func (o *UpdateDeviceRequest) HasFirmware() bool {
	if o != nil && o.Firmware != nil {
		return true
	}

	return false
}

// SetFirmware gets a reference to the given FirmwareMetadata and assigns it to the Firmware field.
func (o *UpdateDeviceRequest) SetFirmware(v FirmwareMetadata) {
	o.Firmware = &v
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *UpdateDeviceRequest) GetConfig() DeviceConfig {
	if o == nil || o.Config == nil {
		var ret DeviceConfig
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceRequest) GetConfigOk() (*DeviceConfig, bool) {
	if o == nil || o.Config == nil {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *UpdateDeviceRequest) HasConfig() bool {
	if o != nil && o.Config != nil {
		return true
	}

	return false
}

// SetConfig gets a reference to the given DeviceConfig and assigns it to the Config field.
func (o *UpdateDeviceRequest) SetConfig(v DeviceConfig) {
	o.Config = &v
}

func (o UpdateDeviceRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CollectionId != nil {
		toSerialize["collectionId"] = o.CollectionId
	}
	if o.Imsi != nil {
		toSerialize["imsi"] = o.Imsi
	}
	if o.Imei != nil {
		toSerialize["imei"] = o.Imei
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.Firmware != nil {
		toSerialize["firmware"] = o.Firmware
	}
	if o.Config != nil {
		toSerialize["config"] = o.Config
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateDeviceRequest struct {
	value *UpdateDeviceRequest
	isSet bool
}

func (v NullableUpdateDeviceRequest) Get() *UpdateDeviceRequest {
	return v.value
}

func (v *NullableUpdateDeviceRequest) Set(val *UpdateDeviceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateDeviceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateDeviceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateDeviceRequest(val *UpdateDeviceRequest) *NullableUpdateDeviceRequest {
	return &NullableUpdateDeviceRequest{value: val, isSet: true}
}

func (v NullableUpdateDeviceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateDeviceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


