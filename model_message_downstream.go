/*
The Span API

API for device, collection, output and firmware management

API version: 4.3.0 grouchy-aloysius
Contact: dev@lab5e.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package spanapi

import (
	"encoding/json"
)

// MessageDownstream Downstream messages are sent from the backend to the devices.
type MessageDownstream struct {
	MessageId    *string           `json:"messageId,omitempty"`
	CollectionId *string           `json:"collectionId,omitempty"`
	DeviceId     *string           `json:"deviceId,omitempty"`
	GatewayId    *string           `json:"gatewayId,omitempty"`
	CreatedTime  *string           `json:"createdTime,omitempty"`
	SentTime     *string           `json:"sentTime,omitempty"`
	Transport    *MessageTransport `json:"transport,omitempty"`
	State        *MessageState     `json:"state,omitempty"`
	Payload      *string           `json:"payload,omitempty"`
}

// NewMessageDownstream instantiates a new MessageDownstream object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessageDownstream() *MessageDownstream {
	this := MessageDownstream{}
	var transport MessageTransport = MESSAGETRANSPORT_UNSPECIFIED
	this.Transport = &transport
	var state MessageState = MESSAGESTATE_UNSPECIFIED
	this.State = &state
	return &this
}

// NewMessageDownstreamWithDefaults instantiates a new MessageDownstream object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessageDownstreamWithDefaults() *MessageDownstream {
	this := MessageDownstream{}
	var transport MessageTransport = MESSAGETRANSPORT_UNSPECIFIED
	this.Transport = &transport
	var state MessageState = MESSAGESTATE_UNSPECIFIED
	this.State = &state
	return &this
}

// GetMessageId returns the MessageId field value if set, zero value otherwise.
func (o *MessageDownstream) GetMessageId() string {
	if o == nil || o.MessageId == nil {
		var ret string
		return ret
	}
	return *o.MessageId
}

// GetMessageIdOk returns a tuple with the MessageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageDownstream) GetMessageIdOk() (*string, bool) {
	if o == nil || o.MessageId == nil {
		return nil, false
	}
	return o.MessageId, true
}

// HasMessageId returns a boolean if a field has been set.
func (o *MessageDownstream) HasMessageId() bool {
	if o != nil && o.MessageId != nil {
		return true
	}

	return false
}

// SetMessageId gets a reference to the given string and assigns it to the MessageId field.
func (o *MessageDownstream) SetMessageId(v string) {
	o.MessageId = &v
}

// GetCollectionId returns the CollectionId field value if set, zero value otherwise.
func (o *MessageDownstream) GetCollectionId() string {
	if o == nil || o.CollectionId == nil {
		var ret string
		return ret
	}
	return *o.CollectionId
}

// GetCollectionIdOk returns a tuple with the CollectionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageDownstream) GetCollectionIdOk() (*string, bool) {
	if o == nil || o.CollectionId == nil {
		return nil, false
	}
	return o.CollectionId, true
}

// HasCollectionId returns a boolean if a field has been set.
func (o *MessageDownstream) HasCollectionId() bool {
	if o != nil && o.CollectionId != nil {
		return true
	}

	return false
}

// SetCollectionId gets a reference to the given string and assigns it to the CollectionId field.
func (o *MessageDownstream) SetCollectionId(v string) {
	o.CollectionId = &v
}

// GetDeviceId returns the DeviceId field value if set, zero value otherwise.
func (o *MessageDownstream) GetDeviceId() string {
	if o == nil || o.DeviceId == nil {
		var ret string
		return ret
	}
	return *o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageDownstream) GetDeviceIdOk() (*string, bool) {
	if o == nil || o.DeviceId == nil {
		return nil, false
	}
	return o.DeviceId, true
}

// HasDeviceId returns a boolean if a field has been set.
func (o *MessageDownstream) HasDeviceId() bool {
	if o != nil && o.DeviceId != nil {
		return true
	}

	return false
}

// SetDeviceId gets a reference to the given string and assigns it to the DeviceId field.
func (o *MessageDownstream) SetDeviceId(v string) {
	o.DeviceId = &v
}

// GetGatewayId returns the GatewayId field value if set, zero value otherwise.
func (o *MessageDownstream) GetGatewayId() string {
	if o == nil || o.GatewayId == nil {
		var ret string
		return ret
	}
	return *o.GatewayId
}

// GetGatewayIdOk returns a tuple with the GatewayId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageDownstream) GetGatewayIdOk() (*string, bool) {
	if o == nil || o.GatewayId == nil {
		return nil, false
	}
	return o.GatewayId, true
}

// HasGatewayId returns a boolean if a field has been set.
func (o *MessageDownstream) HasGatewayId() bool {
	if o != nil && o.GatewayId != nil {
		return true
	}

	return false
}

// SetGatewayId gets a reference to the given string and assigns it to the GatewayId field.
func (o *MessageDownstream) SetGatewayId(v string) {
	o.GatewayId = &v
}

// GetCreatedTime returns the CreatedTime field value if set, zero value otherwise.
func (o *MessageDownstream) GetCreatedTime() string {
	if o == nil || o.CreatedTime == nil {
		var ret string
		return ret
	}
	return *o.CreatedTime
}

// GetCreatedTimeOk returns a tuple with the CreatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageDownstream) GetCreatedTimeOk() (*string, bool) {
	if o == nil || o.CreatedTime == nil {
		return nil, false
	}
	return o.CreatedTime, true
}

// HasCreatedTime returns a boolean if a field has been set.
func (o *MessageDownstream) HasCreatedTime() bool {
	if o != nil && o.CreatedTime != nil {
		return true
	}

	return false
}

// SetCreatedTime gets a reference to the given string and assigns it to the CreatedTime field.
func (o *MessageDownstream) SetCreatedTime(v string) {
	o.CreatedTime = &v
}

// GetSentTime returns the SentTime field value if set, zero value otherwise.
func (o *MessageDownstream) GetSentTime() string {
	if o == nil || o.SentTime == nil {
		var ret string
		return ret
	}
	return *o.SentTime
}

// GetSentTimeOk returns a tuple with the SentTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageDownstream) GetSentTimeOk() (*string, bool) {
	if o == nil || o.SentTime == nil {
		return nil, false
	}
	return o.SentTime, true
}

// HasSentTime returns a boolean if a field has been set.
func (o *MessageDownstream) HasSentTime() bool {
	if o != nil && o.SentTime != nil {
		return true
	}

	return false
}

// SetSentTime gets a reference to the given string and assigns it to the SentTime field.
func (o *MessageDownstream) SetSentTime(v string) {
	o.SentTime = &v
}

// GetTransport returns the Transport field value if set, zero value otherwise.
func (o *MessageDownstream) GetTransport() MessageTransport {
	if o == nil || o.Transport == nil {
		var ret MessageTransport
		return ret
	}
	return *o.Transport
}

// GetTransportOk returns a tuple with the Transport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageDownstream) GetTransportOk() (*MessageTransport, bool) {
	if o == nil || o.Transport == nil {
		return nil, false
	}
	return o.Transport, true
}

// HasTransport returns a boolean if a field has been set.
func (o *MessageDownstream) HasTransport() bool {
	if o != nil && o.Transport != nil {
		return true
	}

	return false
}

// SetTransport gets a reference to the given MessageTransport and assigns it to the Transport field.
func (o *MessageDownstream) SetTransport(v MessageTransport) {
	o.Transport = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *MessageDownstream) GetState() MessageState {
	if o == nil || o.State == nil {
		var ret MessageState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageDownstream) GetStateOk() (*MessageState, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *MessageDownstream) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given MessageState and assigns it to the State field.
func (o *MessageDownstream) SetState(v MessageState) {
	o.State = &v
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *MessageDownstream) GetPayload() string {
	if o == nil || o.Payload == nil {
		var ret string
		return ret
	}
	return *o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageDownstream) GetPayloadOk() (*string, bool) {
	if o == nil || o.Payload == nil {
		return nil, false
	}
	return o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *MessageDownstream) HasPayload() bool {
	if o != nil && o.Payload != nil {
		return true
	}

	return false
}

// SetPayload gets a reference to the given string and assigns it to the Payload field.
func (o *MessageDownstream) SetPayload(v string) {
	o.Payload = &v
}

func (o MessageDownstream) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.MessageId != nil {
		toSerialize["messageId"] = o.MessageId
	}
	if o.CollectionId != nil {
		toSerialize["collectionId"] = o.CollectionId
	}
	if o.DeviceId != nil {
		toSerialize["deviceId"] = o.DeviceId
	}
	if o.GatewayId != nil {
		toSerialize["gatewayId"] = o.GatewayId
	}
	if o.CreatedTime != nil {
		toSerialize["createdTime"] = o.CreatedTime
	}
	if o.SentTime != nil {
		toSerialize["sentTime"] = o.SentTime
	}
	if o.Transport != nil {
		toSerialize["transport"] = o.Transport
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.Payload != nil {
		toSerialize["payload"] = o.Payload
	}
	return json.Marshal(toSerialize)
}

type NullableMessageDownstream struct {
	value *MessageDownstream
	isSet bool
}

func (v NullableMessageDownstream) Get() *MessageDownstream {
	return v.value
}

func (v *NullableMessageDownstream) Set(val *MessageDownstream) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageDownstream) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageDownstream) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageDownstream(val *MessageDownstream) *NullableMessageDownstream {
	return &NullableMessageDownstream{value: val, isSet: true}
}

func (v NullableMessageDownstream) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageDownstream) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
