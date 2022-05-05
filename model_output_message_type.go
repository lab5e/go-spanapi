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

// OutputMessageType the model 'OutputMessageType'
type OutputMessageType string

// List of OutputMessageType
const (
	OUTPUTMESSAGETYPE_UNKNOWN OutputMessageType = "unknown"
	OUTPUTMESSAGETYPE_KEEPALIVE OutputMessageType = "keepalive"
	OUTPUTMESSAGETYPE_DATA OutputMessageType = "data"
)

// All allowed values of OutputMessageType enum
var AllowedOutputMessageTypeEnumValues = []OutputMessageType{
	"unknown",
	"keepalive",
	"data",
}

func (v *OutputMessageType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OutputMessageType(value)
	for _, existing := range AllowedOutputMessageTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OutputMessageType", value)
}

// NewOutputMessageTypeFromValue returns a pointer to a valid OutputMessageType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOutputMessageTypeFromValue(v string) (*OutputMessageType, error) {
	ev := OutputMessageType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OutputMessageType: valid values are %v", v, AllowedOutputMessageTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OutputMessageType) IsValid() bool {
	for _, existing := range AllowedOutputMessageTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OutputMessageType value
func (v OutputMessageType) Ptr() *OutputMessageType {
	return &v
}

type NullableOutputMessageType struct {
	value *OutputMessageType
	isSet bool
}

func (v NullableOutputMessageType) Get() *OutputMessageType {
	return v.value
}

func (v *NullableOutputMessageType) Set(val *OutputMessageType) {
	v.value = val
	v.isSet = true
}

func (v NullableOutputMessageType) IsSet() bool {
	return v.isSet
}

func (v *NullableOutputMessageType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOutputMessageType(val *OutputMessageType) *NullableOutputMessageType {
	return &NullableOutputMessageType{value: val, isSet: true}
}

func (v NullableOutputMessageType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOutputMessageType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

