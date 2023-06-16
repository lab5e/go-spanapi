/*
The Span API

API for device, collection, output and firmware management

API version: 4.6.1 squirming-codi
Contact: dev@lab5e.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package spanapi

import (
	"encoding/json"
	"fmt"
)

// GatewayType the model 'GatewayType'
type GatewayType string

// List of GatewayType
const (
	GATEWAYTYPE_UNKNOWN GatewayType = "unknown"
	GATEWAYTYPE_CIOT GatewayType = "ciot"
	GATEWAYTYPE_INET GatewayType = "inet"
	GATEWAYTYPE_LORA GatewayType = "lora"
	GATEWAYTYPE_OPENTHREAD GatewayType = "openthread"
	GATEWAYTYPE_ZIGBEE GatewayType = "zigbee"
	GATEWAYTYPE_MATTER GatewayType = "matter"
	GATEWAYTYPE_CUSTOM GatewayType = "custom"
)

// All allowed values of GatewayType enum
var AllowedGatewayTypeEnumValues = []GatewayType{
	"unknown",
	"ciot",
	"inet",
	"lora",
	"openthread",
	"zigbee",
	"matter",
	"custom",
}

func (v *GatewayType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GatewayType(value)
	for _, existing := range AllowedGatewayTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GatewayType", value)
}

// NewGatewayTypeFromValue returns a pointer to a valid GatewayType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGatewayTypeFromValue(v string) (*GatewayType, error) {
	ev := GatewayType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GatewayType: valid values are %v", v, AllowedGatewayTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GatewayType) IsValid() bool {
	for _, existing := range AllowedGatewayTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GatewayType value
func (v GatewayType) Ptr() *GatewayType {
	return &v
}

type NullableGatewayType struct {
	value *GatewayType
	isSet bool
}

func (v NullableGatewayType) Get() *GatewayType {
	return v.value
}

func (v *NullableGatewayType) Set(val *GatewayType) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayType) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayType(val *GatewayType) *NullableGatewayType {
	return &NullableGatewayType{value: val, isSet: true}
}

func (v NullableGatewayType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

