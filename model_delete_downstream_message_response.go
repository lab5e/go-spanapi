/*
The Span API

API for device, collection, output and firmware management

API version: 4.2.3 lower-elian
Contact: dev@lab5e.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package spanapi

import (
	"encoding/json"
)

// DeleteDownstreamMessageResponse Response object when deleting a downstream message
type DeleteDownstreamMessageResponse struct {
	MessageId *string `json:"messageId,omitempty"`
}

// NewDeleteDownstreamMessageResponse instantiates a new DeleteDownstreamMessageResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteDownstreamMessageResponse() *DeleteDownstreamMessageResponse {
	this := DeleteDownstreamMessageResponse{}
	return &this
}

// NewDeleteDownstreamMessageResponseWithDefaults instantiates a new DeleteDownstreamMessageResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteDownstreamMessageResponseWithDefaults() *DeleteDownstreamMessageResponse {
	this := DeleteDownstreamMessageResponse{}
	return &this
}

// GetMessageId returns the MessageId field value if set, zero value otherwise.
func (o *DeleteDownstreamMessageResponse) GetMessageId() string {
	if o == nil || o.MessageId == nil {
		var ret string
		return ret
	}
	return *o.MessageId
}

// GetMessageIdOk returns a tuple with the MessageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteDownstreamMessageResponse) GetMessageIdOk() (*string, bool) {
	if o == nil || o.MessageId == nil {
		return nil, false
	}
	return o.MessageId, true
}

// HasMessageId returns a boolean if a field has been set.
func (o *DeleteDownstreamMessageResponse) HasMessageId() bool {
	if o != nil && o.MessageId != nil {
		return true
	}

	return false
}

// SetMessageId gets a reference to the given string and assigns it to the MessageId field.
func (o *DeleteDownstreamMessageResponse) SetMessageId(v string) {
	o.MessageId = &v
}

func (o DeleteDownstreamMessageResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.MessageId != nil {
		toSerialize["messageId"] = o.MessageId
	}
	return json.Marshal(toSerialize)
}

type NullableDeleteDownstreamMessageResponse struct {
	value *DeleteDownstreamMessageResponse
	isSet bool
}

func (v NullableDeleteDownstreamMessageResponse) Get() *DeleteDownstreamMessageResponse {
	return v.value
}

func (v *NullableDeleteDownstreamMessageResponse) Set(val *DeleteDownstreamMessageResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteDownstreamMessageResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteDownstreamMessageResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteDownstreamMessageResponse(val *DeleteDownstreamMessageResponse) *NullableDeleteDownstreamMessageResponse {
	return &NullableDeleteDownstreamMessageResponse{value: val, isSet: true}
}

func (v NullableDeleteDownstreamMessageResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteDownstreamMessageResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
