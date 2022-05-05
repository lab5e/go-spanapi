/*
The Span API

API for device, collection, output and firmware management

API version: 4.2.1 petulant-kyan
Contact: dev@lab5e.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package spanapi

import (
	"encoding/json"
)

// ListDownstreamMessagesResponse Response object when listing downstream messages
type ListDownstreamMessagesResponse struct {
	Messages []MessageDownstream `json:"messages,omitempty"`
}

// NewListDownstreamMessagesResponse instantiates a new ListDownstreamMessagesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListDownstreamMessagesResponse() *ListDownstreamMessagesResponse {
	this := ListDownstreamMessagesResponse{}
	return &this
}

// NewListDownstreamMessagesResponseWithDefaults instantiates a new ListDownstreamMessagesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListDownstreamMessagesResponseWithDefaults() *ListDownstreamMessagesResponse {
	this := ListDownstreamMessagesResponse{}
	return &this
}

// GetMessages returns the Messages field value if set, zero value otherwise.
func (o *ListDownstreamMessagesResponse) GetMessages() []MessageDownstream {
	if o == nil || o.Messages == nil {
		var ret []MessageDownstream
		return ret
	}
	return o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListDownstreamMessagesResponse) GetMessagesOk() ([]MessageDownstream, bool) {
	if o == nil || o.Messages == nil {
		return nil, false
	}
	return o.Messages, true
}

// HasMessages returns a boolean if a field has been set.
func (o *ListDownstreamMessagesResponse) HasMessages() bool {
	if o != nil && o.Messages != nil {
		return true
	}

	return false
}

// SetMessages gets a reference to the given []MessageDownstream and assigns it to the Messages field.
func (o *ListDownstreamMessagesResponse) SetMessages(v []MessageDownstream) {
	o.Messages = v
}

func (o ListDownstreamMessagesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Messages != nil {
		toSerialize["messages"] = o.Messages
	}
	return json.Marshal(toSerialize)
}

type NullableListDownstreamMessagesResponse struct {
	value *ListDownstreamMessagesResponse
	isSet bool
}

func (v NullableListDownstreamMessagesResponse) Get() *ListDownstreamMessagesResponse {
	return v.value
}

func (v *NullableListDownstreamMessagesResponse) Set(val *ListDownstreamMessagesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListDownstreamMessagesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListDownstreamMessagesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListDownstreamMessagesResponse(val *ListDownstreamMessagesResponse) *NullableListDownstreamMessagesResponse {
	return &NullableListDownstreamMessagesResponse{value: val, isSet: true}
}

func (v NullableListDownstreamMessagesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListDownstreamMessagesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


