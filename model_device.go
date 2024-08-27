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

// checks if the Device type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Device{}

// Device This a device
type Device struct {
	// The device ID is assigned by the backend.
	DeviceId *string `json:"deviceId,omitempty"`
	CollectionId *string `json:"collectionId,omitempty"`
	// Tags are metadata for the device that you can set. These are just strings.
	Tags *map[string]string `json:"tags,omitempty"`
	Firmware *FirmwareMetadata `json:"firmware,omitempty"`
	Config *DeviceConfig `json:"config,omitempty"`
	Metadata *DeviceMetadata `json:"metadata,omitempty"`
	LastGatewayId *string `json:"lastGatewayId,omitempty"`
	LastTransport *MessageTransport `json:"lastTransport,omitempty"`
	LastReceived *string `json:"lastReceived,omitempty"`
	LastPayload *string `json:"lastPayload,omitempty"`
	Enabled *bool `json:"enabled,omitempty"`
}

// NewDevice instantiates a new Device object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDevice() *Device {
	this := Device{}
	var lastTransport MessageTransport = MESSAGETRANSPORT_UNSPECIFIED
	this.LastTransport = &lastTransport
	return &this
}

// NewDeviceWithDefaults instantiates a new Device object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceWithDefaults() *Device {
	this := Device{}
	var lastTransport MessageTransport = MESSAGETRANSPORT_UNSPECIFIED
	this.LastTransport = &lastTransport
	return &this
}

// GetDeviceId returns the DeviceId field value if set, zero value otherwise.
func (o *Device) GetDeviceId() string {
	if o == nil || IsNil(o.DeviceId) {
		var ret string
		return ret
	}
	return *o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Device) GetDeviceIdOk() (*string, bool) {
	if o == nil || IsNil(o.DeviceId) {
		return nil, false
	}
	return o.DeviceId, true
}

// HasDeviceId returns a boolean if a field has been set.
func (o *Device) HasDeviceId() bool {
	if o != nil && !IsNil(o.DeviceId) {
		return true
	}

	return false
}

// SetDeviceId gets a reference to the given string and assigns it to the DeviceId field.
func (o *Device) SetDeviceId(v string) {
	o.DeviceId = &v
}

// GetCollectionId returns the CollectionId field value if set, zero value otherwise.
func (o *Device) GetCollectionId() string {
	if o == nil || IsNil(o.CollectionId) {
		var ret string
		return ret
	}
	return *o.CollectionId
}

// GetCollectionIdOk returns a tuple with the CollectionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Device) GetCollectionIdOk() (*string, bool) {
	if o == nil || IsNil(o.CollectionId) {
		return nil, false
	}
	return o.CollectionId, true
}

// HasCollectionId returns a boolean if a field has been set.
func (o *Device) HasCollectionId() bool {
	if o != nil && !IsNil(o.CollectionId) {
		return true
	}

	return false
}

// SetCollectionId gets a reference to the given string and assigns it to the CollectionId field.
func (o *Device) SetCollectionId(v string) {
	o.CollectionId = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *Device) GetTags() map[string]string {
	if o == nil || IsNil(o.Tags) {
		var ret map[string]string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Device) GetTagsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *Device) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given map[string]string and assigns it to the Tags field.
func (o *Device) SetTags(v map[string]string) {
	o.Tags = &v
}

// GetFirmware returns the Firmware field value if set, zero value otherwise.
func (o *Device) GetFirmware() FirmwareMetadata {
	if o == nil || IsNil(o.Firmware) {
		var ret FirmwareMetadata
		return ret
	}
	return *o.Firmware
}

// GetFirmwareOk returns a tuple with the Firmware field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Device) GetFirmwareOk() (*FirmwareMetadata, bool) {
	if o == nil || IsNil(o.Firmware) {
		return nil, false
	}
	return o.Firmware, true
}

// HasFirmware returns a boolean if a field has been set.
func (o *Device) HasFirmware() bool {
	if o != nil && !IsNil(o.Firmware) {
		return true
	}

	return false
}

// SetFirmware gets a reference to the given FirmwareMetadata and assigns it to the Firmware field.
func (o *Device) SetFirmware(v FirmwareMetadata) {
	o.Firmware = &v
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *Device) GetConfig() DeviceConfig {
	if o == nil || IsNil(o.Config) {
		var ret DeviceConfig
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Device) GetConfigOk() (*DeviceConfig, bool) {
	if o == nil || IsNil(o.Config) {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *Device) HasConfig() bool {
	if o != nil && !IsNil(o.Config) {
		return true
	}

	return false
}

// SetConfig gets a reference to the given DeviceConfig and assigns it to the Config field.
func (o *Device) SetConfig(v DeviceConfig) {
	o.Config = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *Device) GetMetadata() DeviceMetadata {
	if o == nil || IsNil(o.Metadata) {
		var ret DeviceMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Device) GetMetadataOk() (*DeviceMetadata, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *Device) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given DeviceMetadata and assigns it to the Metadata field.
func (o *Device) SetMetadata(v DeviceMetadata) {
	o.Metadata = &v
}

// GetLastGatewayId returns the LastGatewayId field value if set, zero value otherwise.
func (o *Device) GetLastGatewayId() string {
	if o == nil || IsNil(o.LastGatewayId) {
		var ret string
		return ret
	}
	return *o.LastGatewayId
}

// GetLastGatewayIdOk returns a tuple with the LastGatewayId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Device) GetLastGatewayIdOk() (*string, bool) {
	if o == nil || IsNil(o.LastGatewayId) {
		return nil, false
	}
	return o.LastGatewayId, true
}

// HasLastGatewayId returns a boolean if a field has been set.
func (o *Device) HasLastGatewayId() bool {
	if o != nil && !IsNil(o.LastGatewayId) {
		return true
	}

	return false
}

// SetLastGatewayId gets a reference to the given string and assigns it to the LastGatewayId field.
func (o *Device) SetLastGatewayId(v string) {
	o.LastGatewayId = &v
}

// GetLastTransport returns the LastTransport field value if set, zero value otherwise.
func (o *Device) GetLastTransport() MessageTransport {
	if o == nil || IsNil(o.LastTransport) {
		var ret MessageTransport
		return ret
	}
	return *o.LastTransport
}

// GetLastTransportOk returns a tuple with the LastTransport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Device) GetLastTransportOk() (*MessageTransport, bool) {
	if o == nil || IsNil(o.LastTransport) {
		return nil, false
	}
	return o.LastTransport, true
}

// HasLastTransport returns a boolean if a field has been set.
func (o *Device) HasLastTransport() bool {
	if o != nil && !IsNil(o.LastTransport) {
		return true
	}

	return false
}

// SetLastTransport gets a reference to the given MessageTransport and assigns it to the LastTransport field.
func (o *Device) SetLastTransport(v MessageTransport) {
	o.LastTransport = &v
}

// GetLastReceived returns the LastReceived field value if set, zero value otherwise.
func (o *Device) GetLastReceived() string {
	if o == nil || IsNil(o.LastReceived) {
		var ret string
		return ret
	}
	return *o.LastReceived
}

// GetLastReceivedOk returns a tuple with the LastReceived field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Device) GetLastReceivedOk() (*string, bool) {
	if o == nil || IsNil(o.LastReceived) {
		return nil, false
	}
	return o.LastReceived, true
}

// HasLastReceived returns a boolean if a field has been set.
func (o *Device) HasLastReceived() bool {
	if o != nil && !IsNil(o.LastReceived) {
		return true
	}

	return false
}

// SetLastReceived gets a reference to the given string and assigns it to the LastReceived field.
func (o *Device) SetLastReceived(v string) {
	o.LastReceived = &v
}

// GetLastPayload returns the LastPayload field value if set, zero value otherwise.
func (o *Device) GetLastPayload() string {
	if o == nil || IsNil(o.LastPayload) {
		var ret string
		return ret
	}
	return *o.LastPayload
}

// GetLastPayloadOk returns a tuple with the LastPayload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Device) GetLastPayloadOk() (*string, bool) {
	if o == nil || IsNil(o.LastPayload) {
		return nil, false
	}
	return o.LastPayload, true
}

// HasLastPayload returns a boolean if a field has been set.
func (o *Device) HasLastPayload() bool {
	if o != nil && !IsNil(o.LastPayload) {
		return true
	}

	return false
}

// SetLastPayload gets a reference to the given string and assigns it to the LastPayload field.
func (o *Device) SetLastPayload(v string) {
	o.LastPayload = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *Device) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Device) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *Device) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *Device) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o Device) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Device) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DeviceId) {
		toSerialize["deviceId"] = o.DeviceId
	}
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
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.LastGatewayId) {
		toSerialize["lastGatewayId"] = o.LastGatewayId
	}
	if !IsNil(o.LastTransport) {
		toSerialize["lastTransport"] = o.LastTransport
	}
	if !IsNil(o.LastReceived) {
		toSerialize["lastReceived"] = o.LastReceived
	}
	if !IsNil(o.LastPayload) {
		toSerialize["lastPayload"] = o.LastPayload
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return toSerialize, nil
}

type NullableDevice struct {
	value *Device
	isSet bool
}

func (v NullableDevice) Get() *Device {
	return v.value
}

func (v *NullableDevice) Set(val *Device) {
	v.value = val
	v.isSet = true
}

func (v NullableDevice) IsSet() bool {
	return v.isSet
}

func (v *NullableDevice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDevice(val *Device) *NullableDevice {
	return &NullableDevice{value: val, isSet: true}
}

func (v NullableDevice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDevice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


