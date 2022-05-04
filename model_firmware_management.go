/*
The Span API

API for device, collection, output and firmware management

API version: 4.2.0 alternative-adolphus
Contact: dev@lab5e.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package spanapi

import (
	"encoding/json"
	"fmt"
)

// FirmwareManagement The firmware management settings for a collection can either be \"disabled\", ie there is no firmware management for this collection, \"collection\"; devices are managed through the settings on the collection or \"device\" where each device is configured individual.
type FirmwareManagement string

// List of FirmwareManagement
const (
	UNSPECIFIED FirmwareManagement = "unspecified"
	DISABLED FirmwareManagement = "disabled"
	COLLECTION FirmwareManagement = "collection"
	DEVICE FirmwareManagement = "device"
)

// All allowed values of FirmwareManagement enum
var AllowedFirmwareManagementEnumValues = []FirmwareManagement{
	"unspecified",
	"disabled",
	"collection",
	"device",
}

func (v *FirmwareManagement) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FirmwareManagement(value)
	for _, existing := range AllowedFirmwareManagementEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FirmwareManagement", value)
}

// NewFirmwareManagementFromValue returns a pointer to a valid FirmwareManagement
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFirmwareManagementFromValue(v string) (*FirmwareManagement, error) {
	ev := FirmwareManagement(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FirmwareManagement: valid values are %v", v, AllowedFirmwareManagementEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FirmwareManagement) IsValid() bool {
	for _, existing := range AllowedFirmwareManagementEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FirmwareManagement value
func (v FirmwareManagement) Ptr() *FirmwareManagement {
	return &v
}

type NullableFirmwareManagement struct {
	value *FirmwareManagement
	isSet bool
}

func (v NullableFirmwareManagement) Get() *FirmwareManagement {
	return v.value
}

func (v *NullableFirmwareManagement) Set(val *FirmwareManagement) {
	v.value = val
	v.isSet = true
}

func (v NullableFirmwareManagement) IsSet() bool {
	return v.isSet
}

func (v *NullableFirmwareManagement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFirmwareManagement(val *FirmwareManagement) *NullableFirmwareManagement {
	return &NullableFirmwareManagement{value: val, isSet: true}
}

func (v NullableFirmwareManagement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFirmwareManagement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
