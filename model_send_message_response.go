/*
 * The Span API
 *
 * API for device, collection, output and firmware management
 *
 * API version: 4.1.7 prized-adelbert
 * Contact: dev@lab5e.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package spanapi

import (
	"encoding/json"
)

// SendMessageResponse struct for SendMessageResponse
type SendMessageResponse struct {
	CollectionId *string `json:"collectionId,omitempty"`
	DeviceId *string `json:"deviceId,omitempty"`
	BytesSent *int32 `json:"bytesSent,omitempty"`
}

// NewSendMessageResponse instantiates a new SendMessageResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSendMessageResponse() *SendMessageResponse {
	this := SendMessageResponse{}
	return &this
}

// NewSendMessageResponseWithDefaults instantiates a new SendMessageResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSendMessageResponseWithDefaults() *SendMessageResponse {
	this := SendMessageResponse{}
	return &this
}

// GetCollectionId returns the CollectionId field value if set, zero value otherwise.
func (o *SendMessageResponse) GetCollectionId() string {
	if o == nil || o.CollectionId == nil {
		var ret string
		return ret
	}
	return *o.CollectionId
}

// GetCollectionIdOk returns a tuple with the CollectionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendMessageResponse) GetCollectionIdOk() (*string, bool) {
	if o == nil || o.CollectionId == nil {
		return nil, false
	}
	return o.CollectionId, true
}

// HasCollectionId returns a boolean if a field has been set.
func (o *SendMessageResponse) HasCollectionId() bool {
	if o != nil && o.CollectionId != nil {
		return true
	}

	return false
}

// SetCollectionId gets a reference to the given string and assigns it to the CollectionId field.
func (o *SendMessageResponse) SetCollectionId(v string) {
	o.CollectionId = &v
}

// GetDeviceId returns the DeviceId field value if set, zero value otherwise.
func (o *SendMessageResponse) GetDeviceId() string {
	if o == nil || o.DeviceId == nil {
		var ret string
		return ret
	}
	return *o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendMessageResponse) GetDeviceIdOk() (*string, bool) {
	if o == nil || o.DeviceId == nil {
		return nil, false
	}
	return o.DeviceId, true
}

// HasDeviceId returns a boolean if a field has been set.
func (o *SendMessageResponse) HasDeviceId() bool {
	if o != nil && o.DeviceId != nil {
		return true
	}

	return false
}

// SetDeviceId gets a reference to the given string and assigns it to the DeviceId field.
func (o *SendMessageResponse) SetDeviceId(v string) {
	o.DeviceId = &v
}

// GetBytesSent returns the BytesSent field value if set, zero value otherwise.
func (o *SendMessageResponse) GetBytesSent() int32 {
	if o == nil || o.BytesSent == nil {
		var ret int32
		return ret
	}
	return *o.BytesSent
}

// GetBytesSentOk returns a tuple with the BytesSent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendMessageResponse) GetBytesSentOk() (*int32, bool) {
	if o == nil || o.BytesSent == nil {
		return nil, false
	}
	return o.BytesSent, true
}

// HasBytesSent returns a boolean if a field has been set.
func (o *SendMessageResponse) HasBytesSent() bool {
	if o != nil && o.BytesSent != nil {
		return true
	}

	return false
}

// SetBytesSent gets a reference to the given int32 and assigns it to the BytesSent field.
func (o *SendMessageResponse) SetBytesSent(v int32) {
	o.BytesSent = &v
}

func (o SendMessageResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CollectionId != nil {
		toSerialize["collectionId"] = o.CollectionId
	}
	if o.DeviceId != nil {
		toSerialize["deviceId"] = o.DeviceId
	}
	if o.BytesSent != nil {
		toSerialize["bytesSent"] = o.BytesSent
	}
	return json.Marshal(toSerialize)
}

type NullableSendMessageResponse struct {
	value *SendMessageResponse
	isSet bool
}

func (v NullableSendMessageResponse) Get() *SendMessageResponse {
	return v.value
}

func (v *NullableSendMessageResponse) Set(val *SendMessageResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSendMessageResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSendMessageResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSendMessageResponse(val *SendMessageResponse) *NullableSendMessageResponse {
	return &NullableSendMessageResponse{value: val, isSet: true}
}

func (v NullableSendMessageResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSendMessageResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


