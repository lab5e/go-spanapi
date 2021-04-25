/*
 * The Span API
 *
 * API for device, collection, output and firmware management
 *
 * API version: 4.1.10 hulking-betty
 * Contact: dev@lab5e.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package spanapi

import (
	"encoding/json"
	"fmt"
)

// CollectionFirmwareFirmwareManagement The firmware management settings for a collection can either be \"disabled\", ie there is no firmware management for this collection, \"collection\"; devices are managed through the settings on the collection or \"device\" where each device is configured individual.
type CollectionFirmwareFirmwareManagement string

// List of CollectionFirmwareFirmwareManagement
const (
	UNSPECIFIED CollectionFirmwareFirmwareManagement = "unspecified"
	DISABLED CollectionFirmwareFirmwareManagement = "disabled"
	COLLECTION CollectionFirmwareFirmwareManagement = "collection"
	DEVICE CollectionFirmwareFirmwareManagement = "device"
)

func (v *CollectionFirmwareFirmwareManagement) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CollectionFirmwareFirmwareManagement(value)
	for _, existing := range []CollectionFirmwareFirmwareManagement{ "unspecified", "disabled", "collection", "device",   } {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CollectionFirmwareFirmwareManagement", value)
}

// Ptr returns reference to CollectionFirmwareFirmwareManagement value
func (v CollectionFirmwareFirmwareManagement) Ptr() *CollectionFirmwareFirmwareManagement {
	return &v
}

type NullableCollectionFirmwareFirmwareManagement struct {
	value *CollectionFirmwareFirmwareManagement
	isSet bool
}

func (v NullableCollectionFirmwareFirmwareManagement) Get() *CollectionFirmwareFirmwareManagement {
	return v.value
}

func (v *NullableCollectionFirmwareFirmwareManagement) Set(val *CollectionFirmwareFirmwareManagement) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionFirmwareFirmwareManagement) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionFirmwareFirmwareManagement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionFirmwareFirmwareManagement(val *CollectionFirmwareFirmwareManagement) *NullableCollectionFirmwareFirmwareManagement {
	return &NullableCollectionFirmwareFirmwareManagement{value: val, isSet: true}
}

func (v NullableCollectionFirmwareFirmwareManagement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionFirmwareFirmwareManagement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

