/*
 * The Span API
 *
 * API for device, collection, output and firmware management
 *
 * API version: 4.1.14 oversensitive-deante
 * Contact: dev@lab5e.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package spanapi

import (
	"encoding/json"
)

// CollectionFirmware struct for CollectionFirmware
type CollectionFirmware struct {
	// The current firmware is the firmware that the devices are currently using.
	CurrentFirmwareId *string `json:"currentFirmwareId,omitempty"`
	// The target firmware is set to the desired firmware image for the devices in this collection. If the management is set to \"device\" this will only be used if the target firmware isn't set on the device itself.
	TargetFirmwareId *string                               `json:"targetFirmwareId,omitempty"`
	Management       *CollectionFirmwareFirmwareManagement `json:"management,omitempty"`
}

// NewCollectionFirmware instantiates a new CollectionFirmware object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionFirmware() *CollectionFirmware {
	this := CollectionFirmware{}
	var management CollectionFirmwareFirmwareManagement = UNSPECIFIED
	this.Management = &management
	return &this
}

// NewCollectionFirmwareWithDefaults instantiates a new CollectionFirmware object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionFirmwareWithDefaults() *CollectionFirmware {
	this := CollectionFirmware{}
	var management CollectionFirmwareFirmwareManagement = UNSPECIFIED
	this.Management = &management
	return &this
}

// GetCurrentFirmwareId returns the CurrentFirmwareId field value if set, zero value otherwise.
func (o *CollectionFirmware) GetCurrentFirmwareId() string {
	if o == nil || o.CurrentFirmwareId == nil {
		var ret string
		return ret
	}
	return *o.CurrentFirmwareId
}

// GetCurrentFirmwareIdOk returns a tuple with the CurrentFirmwareId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionFirmware) GetCurrentFirmwareIdOk() (*string, bool) {
	if o == nil || o.CurrentFirmwareId == nil {
		return nil, false
	}
	return o.CurrentFirmwareId, true
}

// HasCurrentFirmwareId returns a boolean if a field has been set.
func (o *CollectionFirmware) HasCurrentFirmwareId() bool {
	if o != nil && o.CurrentFirmwareId != nil {
		return true
	}

	return false
}

// SetCurrentFirmwareId gets a reference to the given string and assigns it to the CurrentFirmwareId field.
func (o *CollectionFirmware) SetCurrentFirmwareId(v string) {
	o.CurrentFirmwareId = &v
}

// GetTargetFirmwareId returns the TargetFirmwareId field value if set, zero value otherwise.
func (o *CollectionFirmware) GetTargetFirmwareId() string {
	if o == nil || o.TargetFirmwareId == nil {
		var ret string
		return ret
	}
	return *o.TargetFirmwareId
}

// GetTargetFirmwareIdOk returns a tuple with the TargetFirmwareId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionFirmware) GetTargetFirmwareIdOk() (*string, bool) {
	if o == nil || o.TargetFirmwareId == nil {
		return nil, false
	}
	return o.TargetFirmwareId, true
}

// HasTargetFirmwareId returns a boolean if a field has been set.
func (o *CollectionFirmware) HasTargetFirmwareId() bool {
	if o != nil && o.TargetFirmwareId != nil {
		return true
	}

	return false
}

// SetTargetFirmwareId gets a reference to the given string and assigns it to the TargetFirmwareId field.
func (o *CollectionFirmware) SetTargetFirmwareId(v string) {
	o.TargetFirmwareId = &v
}

// GetManagement returns the Management field value if set, zero value otherwise.
func (o *CollectionFirmware) GetManagement() CollectionFirmwareFirmwareManagement {
	if o == nil || o.Management == nil {
		var ret CollectionFirmwareFirmwareManagement
		return ret
	}
	return *o.Management
}

// GetManagementOk returns a tuple with the Management field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionFirmware) GetManagementOk() (*CollectionFirmwareFirmwareManagement, bool) {
	if o == nil || o.Management == nil {
		return nil, false
	}
	return o.Management, true
}

// HasManagement returns a boolean if a field has been set.
func (o *CollectionFirmware) HasManagement() bool {
	if o != nil && o.Management != nil {
		return true
	}

	return false
}

// SetManagement gets a reference to the given CollectionFirmwareFirmwareManagement and assigns it to the Management field.
func (o *CollectionFirmware) SetManagement(v CollectionFirmwareFirmwareManagement) {
	o.Management = &v
}

func (o CollectionFirmware) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CurrentFirmwareId != nil {
		toSerialize["currentFirmwareId"] = o.CurrentFirmwareId
	}
	if o.TargetFirmwareId != nil {
		toSerialize["targetFirmwareId"] = o.TargetFirmwareId
	}
	if o.Management != nil {
		toSerialize["management"] = o.Management
	}
	return json.Marshal(toSerialize)
}

type NullableCollectionFirmware struct {
	value *CollectionFirmware
	isSet bool
}

func (v NullableCollectionFirmware) Get() *CollectionFirmware {
	return v.value
}

func (v *NullableCollectionFirmware) Set(val *CollectionFirmware) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionFirmware) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionFirmware) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionFirmware(val *CollectionFirmware) *NullableCollectionFirmware {
	return &NullableCollectionFirmware{value: val, isSet: true}
}

func (v NullableCollectionFirmware) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionFirmware) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
