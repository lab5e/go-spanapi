/*
 * The Span API
 *
 * API for device, collection, output and firmware management
 *
 * API version: 4.1.14 oversensitive-deante
 * Contact: dev@lab5e.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package spanapi

import (
	"encoding/json"
)

// MessageSendResult struct for MessageSendResult
type MessageSendResult struct {
	DeviceId *string `json:"deviceId,omitempty"`
	Message *string `json:"message,omitempty"`
}

// NewMessageSendResult instantiates a new MessageSendResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessageSendResult() *MessageSendResult {
	this := MessageSendResult{}
	return &this
}

// NewMessageSendResultWithDefaults instantiates a new MessageSendResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessageSendResultWithDefaults() *MessageSendResult {
	this := MessageSendResult{}
	return &this
}

// GetDeviceId returns the DeviceId field value if set, zero value otherwise.
func (o *MessageSendResult) GetDeviceId() string {
	if o == nil || o.DeviceId == nil {
		var ret string
		return ret
	}
	return *o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageSendResult) GetDeviceIdOk() (*string, bool) {
	if o == nil || o.DeviceId == nil {
		return nil, false
	}
	return o.DeviceId, true
}

// HasDeviceId returns a boolean if a field has been set.
func (o *MessageSendResult) HasDeviceId() bool {
	if o != nil && o.DeviceId != nil {
		return true
	}

	return false
}

// SetDeviceId gets a reference to the given string and assigns it to the DeviceId field.
func (o *MessageSendResult) SetDeviceId(v string) {
	o.DeviceId = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *MessageSendResult) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageSendResult) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *MessageSendResult) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *MessageSendResult) SetMessage(v string) {
	o.Message = &v
}

func (o MessageSendResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DeviceId != nil {
		toSerialize["deviceId"] = o.DeviceId
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	return json.Marshal(toSerialize)
}

type NullableMessageSendResult struct {
	value *MessageSendResult
	isSet bool
}

func (v NullableMessageSendResult) Get() *MessageSendResult {
	return v.value
}

func (v *NullableMessageSendResult) Set(val *MessageSendResult) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageSendResult) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageSendResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageSendResult(val *MessageSendResult) *NullableMessageSendResult {
	return &NullableMessageSendResult{value: val, isSet: true}
}

func (v NullableMessageSendResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageSendResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


