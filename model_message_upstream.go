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
)

// checks if the MessageUpstream type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MessageUpstream{}

// MessageUpstream This is the messages sent from the device to the backend service
type MessageUpstream struct {
	MessageId *string `json:"messageId,omitempty"`
	CollectionId *string `json:"collectionId,omitempty"`
	DeviceId *string `json:"deviceId,omitempty"`
	GatewayId *string `json:"gatewayId,omitempty"`
	Transport *MessageTransport `json:"transport,omitempty"`
	Received *string `json:"received,omitempty"`
	Payload *string `json:"payload,omitempty"`
}

// NewMessageUpstream instantiates a new MessageUpstream object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessageUpstream() *MessageUpstream {
	this := MessageUpstream{}
	var transport MessageTransport = MESSAGETRANSPORT_UNSPECIFIED
	this.Transport = &transport
	return &this
}

// NewMessageUpstreamWithDefaults instantiates a new MessageUpstream object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessageUpstreamWithDefaults() *MessageUpstream {
	this := MessageUpstream{}
	var transport MessageTransport = MESSAGETRANSPORT_UNSPECIFIED
	this.Transport = &transport
	return &this
}

// GetMessageId returns the MessageId field value if set, zero value otherwise.
func (o *MessageUpstream) GetMessageId() string {
	if o == nil || IsNil(o.MessageId) {
		var ret string
		return ret
	}
	return *o.MessageId
}

// GetMessageIdOk returns a tuple with the MessageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageUpstream) GetMessageIdOk() (*string, bool) {
	if o == nil || IsNil(o.MessageId) {
		return nil, false
	}
	return o.MessageId, true
}

// HasMessageId returns a boolean if a field has been set.
func (o *MessageUpstream) HasMessageId() bool {
	if o != nil && !IsNil(o.MessageId) {
		return true
	}

	return false
}

// SetMessageId gets a reference to the given string and assigns it to the MessageId field.
func (o *MessageUpstream) SetMessageId(v string) {
	o.MessageId = &v
}

// GetCollectionId returns the CollectionId field value if set, zero value otherwise.
func (o *MessageUpstream) GetCollectionId() string {
	if o == nil || IsNil(o.CollectionId) {
		var ret string
		return ret
	}
	return *o.CollectionId
}

// GetCollectionIdOk returns a tuple with the CollectionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageUpstream) GetCollectionIdOk() (*string, bool) {
	if o == nil || IsNil(o.CollectionId) {
		return nil, false
	}
	return o.CollectionId, true
}

// HasCollectionId returns a boolean if a field has been set.
func (o *MessageUpstream) HasCollectionId() bool {
	if o != nil && !IsNil(o.CollectionId) {
		return true
	}

	return false
}

// SetCollectionId gets a reference to the given string and assigns it to the CollectionId field.
func (o *MessageUpstream) SetCollectionId(v string) {
	o.CollectionId = &v
}

// GetDeviceId returns the DeviceId field value if set, zero value otherwise.
func (o *MessageUpstream) GetDeviceId() string {
	if o == nil || IsNil(o.DeviceId) {
		var ret string
		return ret
	}
	return *o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageUpstream) GetDeviceIdOk() (*string, bool) {
	if o == nil || IsNil(o.DeviceId) {
		return nil, false
	}
	return o.DeviceId, true
}

// HasDeviceId returns a boolean if a field has been set.
func (o *MessageUpstream) HasDeviceId() bool {
	if o != nil && !IsNil(o.DeviceId) {
		return true
	}

	return false
}

// SetDeviceId gets a reference to the given string and assigns it to the DeviceId field.
func (o *MessageUpstream) SetDeviceId(v string) {
	o.DeviceId = &v
}

// GetGatewayId returns the GatewayId field value if set, zero value otherwise.
func (o *MessageUpstream) GetGatewayId() string {
	if o == nil || IsNil(o.GatewayId) {
		var ret string
		return ret
	}
	return *o.GatewayId
}

// GetGatewayIdOk returns a tuple with the GatewayId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageUpstream) GetGatewayIdOk() (*string, bool) {
	if o == nil || IsNil(o.GatewayId) {
		return nil, false
	}
	return o.GatewayId, true
}

// HasGatewayId returns a boolean if a field has been set.
func (o *MessageUpstream) HasGatewayId() bool {
	if o != nil && !IsNil(o.GatewayId) {
		return true
	}

	return false
}

// SetGatewayId gets a reference to the given string and assigns it to the GatewayId field.
func (o *MessageUpstream) SetGatewayId(v string) {
	o.GatewayId = &v
}

// GetTransport returns the Transport field value if set, zero value otherwise.
func (o *MessageUpstream) GetTransport() MessageTransport {
	if o == nil || IsNil(o.Transport) {
		var ret MessageTransport
		return ret
	}
	return *o.Transport
}

// GetTransportOk returns a tuple with the Transport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageUpstream) GetTransportOk() (*MessageTransport, bool) {
	if o == nil || IsNil(o.Transport) {
		return nil, false
	}
	return o.Transport, true
}

// HasTransport returns a boolean if a field has been set.
func (o *MessageUpstream) HasTransport() bool {
	if o != nil && !IsNil(o.Transport) {
		return true
	}

	return false
}

// SetTransport gets a reference to the given MessageTransport and assigns it to the Transport field.
func (o *MessageUpstream) SetTransport(v MessageTransport) {
	o.Transport = &v
}

// GetReceived returns the Received field value if set, zero value otherwise.
func (o *MessageUpstream) GetReceived() string {
	if o == nil || IsNil(o.Received) {
		var ret string
		return ret
	}
	return *o.Received
}

// GetReceivedOk returns a tuple with the Received field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageUpstream) GetReceivedOk() (*string, bool) {
	if o == nil || IsNil(o.Received) {
		return nil, false
	}
	return o.Received, true
}

// HasReceived returns a boolean if a field has been set.
func (o *MessageUpstream) HasReceived() bool {
	if o != nil && !IsNil(o.Received) {
		return true
	}

	return false
}

// SetReceived gets a reference to the given string and assigns it to the Received field.
func (o *MessageUpstream) SetReceived(v string) {
	o.Received = &v
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *MessageUpstream) GetPayload() string {
	if o == nil || IsNil(o.Payload) {
		var ret string
		return ret
	}
	return *o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageUpstream) GetPayloadOk() (*string, bool) {
	if o == nil || IsNil(o.Payload) {
		return nil, false
	}
	return o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *MessageUpstream) HasPayload() bool {
	if o != nil && !IsNil(o.Payload) {
		return true
	}

	return false
}

// SetPayload gets a reference to the given string and assigns it to the Payload field.
func (o *MessageUpstream) SetPayload(v string) {
	o.Payload = &v
}

func (o MessageUpstream) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MessageUpstream) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MessageId) {
		toSerialize["messageId"] = o.MessageId
	}
	if !IsNil(o.CollectionId) {
		toSerialize["collectionId"] = o.CollectionId
	}
	if !IsNil(o.DeviceId) {
		toSerialize["deviceId"] = o.DeviceId
	}
	if !IsNil(o.GatewayId) {
		toSerialize["gatewayId"] = o.GatewayId
	}
	if !IsNil(o.Transport) {
		toSerialize["transport"] = o.Transport
	}
	if !IsNil(o.Received) {
		toSerialize["received"] = o.Received
	}
	if !IsNil(o.Payload) {
		toSerialize["payload"] = o.Payload
	}
	return toSerialize, nil
}

type NullableMessageUpstream struct {
	value *MessageUpstream
	isSet bool
}

func (v NullableMessageUpstream) Get() *MessageUpstream {
	return v.value
}

func (v *NullableMessageUpstream) Set(val *MessageUpstream) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageUpstream) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageUpstream) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageUpstream(val *MessageUpstream) *NullableMessageUpstream {
	return &NullableMessageUpstream{value: val, isSet: true}
}

func (v NullableMessageUpstream) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageUpstream) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


