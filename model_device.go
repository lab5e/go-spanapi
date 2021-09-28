/*
 * The Span API
 *
 * API for device, collection, output and firmware management
 *
 * API version: 4.1.17 enhanced-allie
 * Contact: dev@lab5e.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package spanapi

import (
	"encoding/json"
)

// Device struct for Device
type Device struct {
	// The device ID is assigned by the backend.
	DeviceId *string `json:"deviceId,omitempty"`
	CollectionId *string `json:"collectionId,omitempty"`
	// The IMSI is the unique ID for the (e|nu|whatever)SIM card on your device. This is the primary identifier for your device on the network.
	Imsi *string `json:"imsi,omitempty"`
	// The IMEI number is the unique ID for your hardware as seen by the network. Obviously you might have a completely different view on things.
	Imei *string `json:"imei,omitempty"`
	// Tags are metadata for the device that you can set. These are just strings.
	Tags *map[string]string `json:"tags,omitempty"`
	Network *NetworkMetadata `json:"network,omitempty"`
	Firmware *FirmwareMetadata `json:"firmware,omitempty"`
	Metadata *DeviceMetadata `json:"metadata,omitempty"`
}

// NewDevice instantiates a new Device object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDevice() *Device {
	this := Device{}
	return &this
}

// NewDeviceWithDefaults instantiates a new Device object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceWithDefaults() *Device {
	this := Device{}
	return &this
}

// GetDeviceId returns the DeviceId field value if set, zero value otherwise.
func (o *Device) GetDeviceId() string {
	if o == nil || o.DeviceId == nil {
		var ret string
		return ret
	}
	return *o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Device) GetDeviceIdOk() (*string, bool) {
	if o == nil || o.DeviceId == nil {
		return nil, false
	}
	return o.DeviceId, true
}

// HasDeviceId returns a boolean if a field has been set.
func (o *Device) HasDeviceId() bool {
	if o != nil && o.DeviceId != nil {
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
	if o == nil || o.CollectionId == nil {
		var ret string
		return ret
	}
	return *o.CollectionId
}

// GetCollectionIdOk returns a tuple with the CollectionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Device) GetCollectionIdOk() (*string, bool) {
	if o == nil || o.CollectionId == nil {
		return nil, false
	}
	return o.CollectionId, true
}

// HasCollectionId returns a boolean if a field has been set.
func (o *Device) HasCollectionId() bool {
	if o != nil && o.CollectionId != nil {
		return true
	}

	return false
}

// SetCollectionId gets a reference to the given string and assigns it to the CollectionId field.
func (o *Device) SetCollectionId(v string) {
	o.CollectionId = &v
}

// GetImsi returns the Imsi field value if set, zero value otherwise.
func (o *Device) GetImsi() string {
	if o == nil || o.Imsi == nil {
		var ret string
		return ret
	}
	return *o.Imsi
}

// GetImsiOk returns a tuple with the Imsi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Device) GetImsiOk() (*string, bool) {
	if o == nil || o.Imsi == nil {
		return nil, false
	}
	return o.Imsi, true
}

// HasImsi returns a boolean if a field has been set.
func (o *Device) HasImsi() bool {
	if o != nil && o.Imsi != nil {
		return true
	}

	return false
}

// SetImsi gets a reference to the given string and assigns it to the Imsi field.
func (o *Device) SetImsi(v string) {
	o.Imsi = &v
}

// GetImei returns the Imei field value if set, zero value otherwise.
func (o *Device) GetImei() string {
	if o == nil || o.Imei == nil {
		var ret string
		return ret
	}
	return *o.Imei
}

// GetImeiOk returns a tuple with the Imei field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Device) GetImeiOk() (*string, bool) {
	if o == nil || o.Imei == nil {
		return nil, false
	}
	return o.Imei, true
}

// HasImei returns a boolean if a field has been set.
func (o *Device) HasImei() bool {
	if o != nil && o.Imei != nil {
		return true
	}

	return false
}

// SetImei gets a reference to the given string and assigns it to the Imei field.
func (o *Device) SetImei(v string) {
	o.Imei = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *Device) GetTags() map[string]string {
	if o == nil || o.Tags == nil {
		var ret map[string]string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Device) GetTagsOk() (*map[string]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *Device) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given map[string]string and assigns it to the Tags field.
func (o *Device) SetTags(v map[string]string) {
	o.Tags = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *Device) GetNetwork() NetworkMetadata {
	if o == nil || o.Network == nil {
		var ret NetworkMetadata
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Device) GetNetworkOk() (*NetworkMetadata, bool) {
	if o == nil || o.Network == nil {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *Device) HasNetwork() bool {
	if o != nil && o.Network != nil {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given NetworkMetadata and assigns it to the Network field.
func (o *Device) SetNetwork(v NetworkMetadata) {
	o.Network = &v
}

// GetFirmware returns the Firmware field value if set, zero value otherwise.
func (o *Device) GetFirmware() FirmwareMetadata {
	if o == nil || o.Firmware == nil {
		var ret FirmwareMetadata
		return ret
	}
	return *o.Firmware
}

// GetFirmwareOk returns a tuple with the Firmware field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Device) GetFirmwareOk() (*FirmwareMetadata, bool) {
	if o == nil || o.Firmware == nil {
		return nil, false
	}
	return o.Firmware, true
}

// HasFirmware returns a boolean if a field has been set.
func (o *Device) HasFirmware() bool {
	if o != nil && o.Firmware != nil {
		return true
	}

	return false
}

// SetFirmware gets a reference to the given FirmwareMetadata and assigns it to the Firmware field.
func (o *Device) SetFirmware(v FirmwareMetadata) {
	o.Firmware = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *Device) GetMetadata() DeviceMetadata {
	if o == nil || o.Metadata == nil {
		var ret DeviceMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Device) GetMetadataOk() (*DeviceMetadata, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *Device) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given DeviceMetadata and assigns it to the Metadata field.
func (o *Device) SetMetadata(v DeviceMetadata) {
	o.Metadata = &v
}

func (o Device) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DeviceId != nil {
		toSerialize["deviceId"] = o.DeviceId
	}
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
	if o.Network != nil {
		toSerialize["network"] = o.Network
	}
	if o.Firmware != nil {
		toSerialize["firmware"] = o.Firmware
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	return json.Marshal(toSerialize)
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


