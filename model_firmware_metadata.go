/*
 * The Span API
 *
 * API for device, collection, output and firmware management
 *
 * API version: 4.1.18 disgruntled-jerald
 * Contact: dev@lab5e.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package spanapi

import (
	"encoding/json"
)

// FirmwareMetadata Metadata about firmware on devices.
type FirmwareMetadata struct {
	CurrentFirmwareId *string `json:"currentFirmwareId,omitempty"`
	TargetFirmwareId *string `json:"targetFirmwareId,omitempty"`
	// Last reported firmware version.
	FirmwareVersion *string `json:"firmwareVersion,omitempty"`
	SerialNumber *string `json:"serialNumber,omitempty"`
	ModelNumber *string `json:"modelNumber,omitempty"`
	Manufacturer *string `json:"manufacturer,omitempty"`
	// State of the firmware.
	State *string `json:"state,omitempty"`
	StateMessage *string `json:"stateMessage,omitempty"`
}

// NewFirmwareMetadata instantiates a new FirmwareMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFirmwareMetadata() *FirmwareMetadata {
	this := FirmwareMetadata{}
	return &this
}

// NewFirmwareMetadataWithDefaults instantiates a new FirmwareMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFirmwareMetadataWithDefaults() *FirmwareMetadata {
	this := FirmwareMetadata{}
	return &this
}

// GetCurrentFirmwareId returns the CurrentFirmwareId field value if set, zero value otherwise.
func (o *FirmwareMetadata) GetCurrentFirmwareId() string {
	if o == nil || o.CurrentFirmwareId == nil {
		var ret string
		return ret
	}
	return *o.CurrentFirmwareId
}

// GetCurrentFirmwareIdOk returns a tuple with the CurrentFirmwareId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareMetadata) GetCurrentFirmwareIdOk() (*string, bool) {
	if o == nil || o.CurrentFirmwareId == nil {
		return nil, false
	}
	return o.CurrentFirmwareId, true
}

// HasCurrentFirmwareId returns a boolean if a field has been set.
func (o *FirmwareMetadata) HasCurrentFirmwareId() bool {
	if o != nil && o.CurrentFirmwareId != nil {
		return true
	}

	return false
}

// SetCurrentFirmwareId gets a reference to the given string and assigns it to the CurrentFirmwareId field.
func (o *FirmwareMetadata) SetCurrentFirmwareId(v string) {
	o.CurrentFirmwareId = &v
}

// GetTargetFirmwareId returns the TargetFirmwareId field value if set, zero value otherwise.
func (o *FirmwareMetadata) GetTargetFirmwareId() string {
	if o == nil || o.TargetFirmwareId == nil {
		var ret string
		return ret
	}
	return *o.TargetFirmwareId
}

// GetTargetFirmwareIdOk returns a tuple with the TargetFirmwareId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareMetadata) GetTargetFirmwareIdOk() (*string, bool) {
	if o == nil || o.TargetFirmwareId == nil {
		return nil, false
	}
	return o.TargetFirmwareId, true
}

// HasTargetFirmwareId returns a boolean if a field has been set.
func (o *FirmwareMetadata) HasTargetFirmwareId() bool {
	if o != nil && o.TargetFirmwareId != nil {
		return true
	}

	return false
}

// SetTargetFirmwareId gets a reference to the given string and assigns it to the TargetFirmwareId field.
func (o *FirmwareMetadata) SetTargetFirmwareId(v string) {
	o.TargetFirmwareId = &v
}

// GetFirmwareVersion returns the FirmwareVersion field value if set, zero value otherwise.
func (o *FirmwareMetadata) GetFirmwareVersion() string {
	if o == nil || o.FirmwareVersion == nil {
		var ret string
		return ret
	}
	return *o.FirmwareVersion
}

// GetFirmwareVersionOk returns a tuple with the FirmwareVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareMetadata) GetFirmwareVersionOk() (*string, bool) {
	if o == nil || o.FirmwareVersion == nil {
		return nil, false
	}
	return o.FirmwareVersion, true
}

// HasFirmwareVersion returns a boolean if a field has been set.
func (o *FirmwareMetadata) HasFirmwareVersion() bool {
	if o != nil && o.FirmwareVersion != nil {
		return true
	}

	return false
}

// SetFirmwareVersion gets a reference to the given string and assigns it to the FirmwareVersion field.
func (o *FirmwareMetadata) SetFirmwareVersion(v string) {
	o.FirmwareVersion = &v
}

// GetSerialNumber returns the SerialNumber field value if set, zero value otherwise.
func (o *FirmwareMetadata) GetSerialNumber() string {
	if o == nil || o.SerialNumber == nil {
		var ret string
		return ret
	}
	return *o.SerialNumber
}

// GetSerialNumberOk returns a tuple with the SerialNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareMetadata) GetSerialNumberOk() (*string, bool) {
	if o == nil || o.SerialNumber == nil {
		return nil, false
	}
	return o.SerialNumber, true
}

// HasSerialNumber returns a boolean if a field has been set.
func (o *FirmwareMetadata) HasSerialNumber() bool {
	if o != nil && o.SerialNumber != nil {
		return true
	}

	return false
}

// SetSerialNumber gets a reference to the given string and assigns it to the SerialNumber field.
func (o *FirmwareMetadata) SetSerialNumber(v string) {
	o.SerialNumber = &v
}

// GetModelNumber returns the ModelNumber field value if set, zero value otherwise.
func (o *FirmwareMetadata) GetModelNumber() string {
	if o == nil || o.ModelNumber == nil {
		var ret string
		return ret
	}
	return *o.ModelNumber
}

// GetModelNumberOk returns a tuple with the ModelNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareMetadata) GetModelNumberOk() (*string, bool) {
	if o == nil || o.ModelNumber == nil {
		return nil, false
	}
	return o.ModelNumber, true
}

// HasModelNumber returns a boolean if a field has been set.
func (o *FirmwareMetadata) HasModelNumber() bool {
	if o != nil && o.ModelNumber != nil {
		return true
	}

	return false
}

// SetModelNumber gets a reference to the given string and assigns it to the ModelNumber field.
func (o *FirmwareMetadata) SetModelNumber(v string) {
	o.ModelNumber = &v
}

// GetManufacturer returns the Manufacturer field value if set, zero value otherwise.
func (o *FirmwareMetadata) GetManufacturer() string {
	if o == nil || o.Manufacturer == nil {
		var ret string
		return ret
	}
	return *o.Manufacturer
}

// GetManufacturerOk returns a tuple with the Manufacturer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareMetadata) GetManufacturerOk() (*string, bool) {
	if o == nil || o.Manufacturer == nil {
		return nil, false
	}
	return o.Manufacturer, true
}

// HasManufacturer returns a boolean if a field has been set.
func (o *FirmwareMetadata) HasManufacturer() bool {
	if o != nil && o.Manufacturer != nil {
		return true
	}

	return false
}

// SetManufacturer gets a reference to the given string and assigns it to the Manufacturer field.
func (o *FirmwareMetadata) SetManufacturer(v string) {
	o.Manufacturer = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *FirmwareMetadata) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareMetadata) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *FirmwareMetadata) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *FirmwareMetadata) SetState(v string) {
	o.State = &v
}

// GetStateMessage returns the StateMessage field value if set, zero value otherwise.
func (o *FirmwareMetadata) GetStateMessage() string {
	if o == nil || o.StateMessage == nil {
		var ret string
		return ret
	}
	return *o.StateMessage
}

// GetStateMessageOk returns a tuple with the StateMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareMetadata) GetStateMessageOk() (*string, bool) {
	if o == nil || o.StateMessage == nil {
		return nil, false
	}
	return o.StateMessage, true
}

// HasStateMessage returns a boolean if a field has been set.
func (o *FirmwareMetadata) HasStateMessage() bool {
	if o != nil && o.StateMessage != nil {
		return true
	}

	return false
}

// SetStateMessage gets a reference to the given string and assigns it to the StateMessage field.
func (o *FirmwareMetadata) SetStateMessage(v string) {
	o.StateMessage = &v
}

func (o FirmwareMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CurrentFirmwareId != nil {
		toSerialize["currentFirmwareId"] = o.CurrentFirmwareId
	}
	if o.TargetFirmwareId != nil {
		toSerialize["targetFirmwareId"] = o.TargetFirmwareId
	}
	if o.FirmwareVersion != nil {
		toSerialize["firmwareVersion"] = o.FirmwareVersion
	}
	if o.SerialNumber != nil {
		toSerialize["serialNumber"] = o.SerialNumber
	}
	if o.ModelNumber != nil {
		toSerialize["modelNumber"] = o.ModelNumber
	}
	if o.Manufacturer != nil {
		toSerialize["manufacturer"] = o.Manufacturer
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.StateMessage != nil {
		toSerialize["stateMessage"] = o.StateMessage
	}
	return json.Marshal(toSerialize)
}

type NullableFirmwareMetadata struct {
	value *FirmwareMetadata
	isSet bool
}

func (v NullableFirmwareMetadata) Get() *FirmwareMetadata {
	return v.value
}

func (v *NullableFirmwareMetadata) Set(val *FirmwareMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableFirmwareMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableFirmwareMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFirmwareMetadata(val *FirmwareMetadata) *NullableFirmwareMetadata {
	return &NullableFirmwareMetadata{value: val, isSet: true}
}

func (v NullableFirmwareMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFirmwareMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


