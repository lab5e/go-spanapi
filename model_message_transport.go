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

// MessageTransport The message transport can be UDP or CoAP.
type MessageTransport string

// List of MessageTransport
const (
	MESSAGETRANSPORT_UNSPECIFIED MessageTransport = "unspecified"
	MESSAGETRANSPORT_UDP MessageTransport = "udp"
	MESSAGETRANSPORT_COAP MessageTransport = "coap"
	MESSAGETRANSPORT_MQTT MessageTransport = "mqtt"
	MESSAGETRANSPORT_GATEWAY MessageTransport = "gateway"
	MESSAGETRANSPORT_COAPS MessageTransport = "coaps"
	MESSAGETRANSPORT_DTLS MessageTransport = "dtls"
)

// All allowed values of MessageTransport enum
var AllowedMessageTransportEnumValues = []MessageTransport{
	"unspecified",
	"udp",
	"coap",
	"mqtt",
	"gateway",
	"coaps",
	"dtls",
}

func (v *MessageTransport) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MessageTransport(value)
	for _, existing := range AllowedMessageTransportEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MessageTransport", value)
}

// NewMessageTransportFromValue returns a pointer to a valid MessageTransport
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMessageTransportFromValue(v string) (*MessageTransport, error) {
	ev := MessageTransport(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MessageTransport: valid values are %v", v, AllowedMessageTransportEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MessageTransport) IsValid() bool {
	for _, existing := range AllowedMessageTransportEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MessageTransport value
func (v MessageTransport) Ptr() *MessageTransport {
	return &v
}

type NullableMessageTransport struct {
	value *MessageTransport
	isSet bool
}

func (v NullableMessageTransport) Get() *MessageTransport {
	return v.value
}

func (v *NullableMessageTransport) Set(val *MessageTransport) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageTransport) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageTransport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageTransport(val *MessageTransport) *NullableMessageTransport {
	return &NullableMessageTransport{value: val, isSet: true}
}

func (v NullableMessageTransport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageTransport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

