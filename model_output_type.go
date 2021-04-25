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

// OutputType the model 'OutputType'
type OutputType string

// List of OutputType
const (
	UNDEFINED OutputType = "undefined"
	WEBHOOK OutputType = "webhook"
	UDP OutputType = "udp"
	MQTT OutputType = "mqtt"
	IFTTT OutputType = "ifttt"
)

func (v *OutputType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OutputType(value)
	for _, existing := range []OutputType{ "undefined", "webhook", "udp", "mqtt", "ifttt",   } {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OutputType", value)
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

