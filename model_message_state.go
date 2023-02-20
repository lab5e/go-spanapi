/*
The Span API

API for device, collection, output and firmware management

API version: 4.4.0 lean-joline
Contact: dev@lab5e.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package spanapi

import (
	"encoding/json"
	"fmt"
)

// MessageState the model 'MessageState'
type MessageState string

// List of MessageState
const (
	MESSAGESTATE_UNSPECIFIED MessageState = "unspecified"
	MESSAGESTATE_PENDING     MessageState = "pending"
	MESSAGESTATE_SENT        MessageState = "sent"
	MESSAGESTATE_FAILED      MessageState = "failed"
)

// All allowed values of MessageState enum
var AllowedMessageStateEnumValues = []MessageState{
	"unspecified",
	"pending",
	"sent",
	"failed",
}

func (v *MessageState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MessageState(value)
	for _, existing := range AllowedMessageStateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MessageState", value)
}

// NewMessageStateFromValue returns a pointer to a valid MessageState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMessageStateFromValue(v string) (*MessageState, error) {
	ev := MessageState(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MessageState: valid values are %v", v, AllowedMessageStateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MessageState) IsValid() bool {
	for _, existing := range AllowedMessageStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MessageState value
func (v MessageState) Ptr() *MessageState {
	return &v
}

type NullableMessageState struct {
	value *MessageState
	isSet bool
}

func (v NullableMessageState) Get() *MessageState {
	return v.value
}

func (v *NullableMessageState) Set(val *MessageState) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageState) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageState(val *MessageState) *NullableMessageState {
	return &NullableMessageState{value: val, isSet: true}
}

func (v NullableMessageState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
