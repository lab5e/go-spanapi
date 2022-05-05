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

// ListUpstreamMessagesResponse Response object when listing upstream messages
type ListUpstreamMessagesResponse struct {
	Messages []MessageUpstream `json:"messages,omitempty"`
}

// NewListUpstreamMessagesResponse instantiates a new ListUpstreamMessagesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListUpstreamMessagesResponse() *ListUpstreamMessagesResponse {
	this := ListUpstreamMessagesResponse{}
	return &this
}

// NewListUpstreamMessagesResponseWithDefaults instantiates a new ListUpstreamMessagesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListUpstreamMessagesResponseWithDefaults() *ListUpstreamMessagesResponse {
	this := ListUpstreamMessagesResponse{}
	return &this
}

// GetMessages returns the Messages field value if set, zero value otherwise.
func (o *ListUpstreamMessagesResponse) GetMessages() []MessageUpstream {
	if o == nil || o.Messages == nil {
		var ret []MessageUpstream
		return ret
	}
	return o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListUpstreamMessagesResponse) GetMessagesOk() ([]MessageUpstream, bool) {
	if o == nil || o.Messages == nil {
		return nil, false
	}
	return o.Messages, true
}

// HasMessages returns a boolean if a field has been set.
func (o *ListUpstreamMessagesResponse) HasMessages() bool {
	if o != nil && o.Messages != nil {
		return true
	}

	return false
}

// SetMessages gets a reference to the given []MessageUpstream and assigns it to the Messages field.
func (o *ListUpstreamMessagesResponse) SetMessages(v []MessageUpstream) {
	o.Messages = v
}

func (o ListUpstreamMessagesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Messages != nil {
		toSerialize["messages"] = o.Messages
	}
	return json.Marshal(toSerialize)
}

type NullableListUpstreamMessagesResponse struct {
	value *ListUpstreamMessagesResponse
	isSet bool
}

func (v NullableListUpstreamMessagesResponse) Get() *ListUpstreamMessagesResponse {
	return v.value
}

func (v *NullableListUpstreamMessagesResponse) Set(val *ListUpstreamMessagesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListUpstreamMessagesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListUpstreamMessagesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListUpstreamMessagesResponse(val *ListUpstreamMessagesResponse) *NullableListUpstreamMessagesResponse {
	return &NullableListUpstreamMessagesResponse{value: val, isSet: true}
}

func (v NullableListUpstreamMessagesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListUpstreamMessagesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


