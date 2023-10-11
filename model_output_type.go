/*
The Span API

API for device, collection, output and firmware management

API version: 4.7.0 actionable-aryanna
Contact: dev@lab5e.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package spanapi

import (
	"encoding/json"
	"fmt"
)

// OutputType Output types available   - undefined: The undefined output type is an invalid type
type OutputType string

// List of OutputType
const (
	OUTPUTTYPE_UNDEFINED OutputType = "undefined"
	OUTPUTTYPE_WEBHOOK OutputType = "webhook"
	OUTPUTTYPE_UDPOUT OutputType = "udpout"
	OUTPUTTYPE_MQTTCLIENT OutputType = "mqttclient"
	OUTPUTTYPE_IFTTT OutputType = "ifttt"
	OUTPUTTYPE_MQTTBROKER OutputType = "mqttbroker"
)

// All allowed values of OutputType enum
var AllowedOutputTypeEnumValues = []OutputType{
	"undefined",
	"webhook",
	"udpout",
	"mqttclient",
	"ifttt",
	"mqttbroker",
}

func (v *OutputType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OutputType(value)
	for _, existing := range AllowedOutputTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OutputType", value)
}

// NewOutputTypeFromValue returns a pointer to a valid OutputType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOutputTypeFromValue(v string) (*OutputType, error) {
	ev := OutputType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OutputType: valid values are %v", v, AllowedOutputTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OutputType) IsValid() bool {
	for _, existing := range AllowedOutputTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OutputType value
func (v OutputType) Ptr() *OutputType {
	return &v
}

type NullableOutputType struct {
	value *OutputType
	isSet bool
}

func (v NullableOutputType) Get() *OutputType {
	return v.value
}

func (v *NullableOutputType) Set(val *OutputType) {
	v.value = val
	v.isSet = true
}

func (v NullableOutputType) IsSet() bool {
	return v.isSet
}

func (v *NullableOutputType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOutputType(val *OutputType) *NullableOutputType {
	return &NullableOutputType{value: val, isSet: true}
}

func (v NullableOutputType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOutputType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

