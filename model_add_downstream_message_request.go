/*
The Span API

API for device, collection, output and firmware management

API version: 4.9.6 authoritarian-betty
Contact: dev@lab5e.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package spanapi

import (
	"encoding/json"
)

// checks if the AddDownstreamMessageRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddDownstreamMessageRequest{}

// AddDownstreamMessageRequest This is the request object to send messages out to the devices
type AddDownstreamMessageRequest struct {
	Payload *string `json:"payload,omitempty"`
}

// NewAddDownstreamMessageRequest instantiates a new AddDownstreamMessageRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddDownstreamMessageRequest() *AddDownstreamMessageRequest {
	this := AddDownstreamMessageRequest{}
	return &this
}

// NewAddDownstreamMessageRequestWithDefaults instantiates a new AddDownstreamMessageRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddDownstreamMessageRequestWithDefaults() *AddDownstreamMessageRequest {
	this := AddDownstreamMessageRequest{}
	return &this
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *AddDownstreamMessageRequest) GetPayload() string {
	if o == nil || IsNil(o.Payload) {
		var ret string
		return ret
	}
	return *o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddDownstreamMessageRequest) GetPayloadOk() (*string, bool) {
	if o == nil || IsNil(o.Payload) {
		return nil, false
	}
	return o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *AddDownstreamMessageRequest) HasPayload() bool {
	if o != nil && !IsNil(o.Payload) {
		return true
	}

	return false
}

// SetPayload gets a reference to the given string and assigns it to the Payload field.
func (o *AddDownstreamMessageRequest) SetPayload(v string) {
	o.Payload = &v
}

func (o AddDownstreamMessageRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddDownstreamMessageRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Payload) {
		toSerialize["payload"] = o.Payload
	}
	return toSerialize, nil
}

type NullableAddDownstreamMessageRequest struct {
	value *AddDownstreamMessageRequest
	isSet bool
}

func (v NullableAddDownstreamMessageRequest) Get() *AddDownstreamMessageRequest {
	return v.value
}

func (v *NullableAddDownstreamMessageRequest) Set(val *AddDownstreamMessageRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddDownstreamMessageRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddDownstreamMessageRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddDownstreamMessageRequest(val *AddDownstreamMessageRequest) *NullableAddDownstreamMessageRequest {
	return &NullableAddDownstreamMessageRequest{value: val, isSet: true}
}

func (v NullableAddDownstreamMessageRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddDownstreamMessageRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


