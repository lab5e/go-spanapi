/*
 * The Span API
 *
 * API for device, collection, output and firmware management
 *
 * API version: 4.1.12 infinite-dana
 * Contact: dev@lab5e.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package spanapi

import (
	"encoding/json"
	"fmt"
)

// OutputDataMessageOutputMessageType the model 'OutputDataMessageOutputMessageType'
type OutputDataMessageOutputMessageType string

// List of OutputDataMessageOutputMessageType
const (
	UNKNOWN   OutputDataMessageOutputMessageType = "unknown"
	KEEPALIVE OutputDataMessageOutputMessageType = "keepalive"
	DATA      OutputDataMessageOutputMessageType = "data"
)

func (v *OutputDataMessageOutputMessageType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OutputDataMessageOutputMessageType(value)
	for _, existing := range []OutputDataMessageOutputMessageType{"unknown", "keepalive", "data"} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OutputDataMessageOutputMessageType", value)
}

// Ptr returns reference to OutputDataMessageOutputMessageType value
func (v OutputDataMessageOutputMessageType) Ptr() *OutputDataMessageOutputMessageType {
	return &v
}

type NullableOutputDataMessageOutputMessageType struct {
	value *OutputDataMessageOutputMessageType
	isSet bool
}

func (v NullableOutputDataMessageOutputMessageType) Get() *OutputDataMessageOutputMessageType {
	return v.value
}

func (v *NullableOutputDataMessageOutputMessageType) Set(val *OutputDataMessageOutputMessageType) {
	v.value = val
	v.isSet = true
}

func (v NullableOutputDataMessageOutputMessageType) IsSet() bool {
	return v.isSet
}

func (v *NullableOutputDataMessageOutputMessageType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOutputDataMessageOutputMessageType(val *OutputDataMessageOutputMessageType) *NullableOutputDataMessageOutputMessageType {
	return &NullableOutputDataMessageOutputMessageType{value: val, isSet: true}
}

func (v NullableOutputDataMessageOutputMessageType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOutputDataMessageOutputMessageType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
