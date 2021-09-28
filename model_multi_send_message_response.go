/*
 * The Span API
 *
 * API for device, collection, output and firmware management
 *
 * API version: 4.1.17 enhanced-allie
 * Contact: dev@lab5e.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package spanapi

import (
	"encoding/json"
)

// MultiSendMessageResponse Broadcast message result. The errors array contains the list of errors ocurred when sending a message.
type MultiSendMessageResponse struct {
	Errors *[]MessageSendResult `json:"errors,omitempty"`
	Sent *int32 `json:"sent,omitempty"`
	Failed *int32 `json:"failed,omitempty"`
}

// NewMultiSendMessageResponse instantiates a new MultiSendMessageResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMultiSendMessageResponse() *MultiSendMessageResponse {
	this := MultiSendMessageResponse{}
	return &this
}

// NewMultiSendMessageResponseWithDefaults instantiates a new MultiSendMessageResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMultiSendMessageResponseWithDefaults() *MultiSendMessageResponse {
	this := MultiSendMessageResponse{}
	return &this
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *MultiSendMessageResponse) GetErrors() []MessageSendResult {
	if o == nil || o.Errors == nil {
		var ret []MessageSendResult
		return ret
	}
	return *o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultiSendMessageResponse) GetErrorsOk() (*[]MessageSendResult, bool) {
	if o == nil || o.Errors == nil {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *MultiSendMessageResponse) HasErrors() bool {
	if o != nil && o.Errors != nil {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []MessageSendResult and assigns it to the Errors field.
func (o *MultiSendMessageResponse) SetErrors(v []MessageSendResult) {
	o.Errors = &v
}

// GetSent returns the Sent field value if set, zero value otherwise.
func (o *MultiSendMessageResponse) GetSent() int32 {
	if o == nil || o.Sent == nil {
		var ret int32
		return ret
	}
	return *o.Sent
}

// GetSentOk returns a tuple with the Sent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultiSendMessageResponse) GetSentOk() (*int32, bool) {
	if o == nil || o.Sent == nil {
		return nil, false
	}
	return o.Sent, true
}

// HasSent returns a boolean if a field has been set.
func (o *MultiSendMessageResponse) HasSent() bool {
	if o != nil && o.Sent != nil {
		return true
	}

	return false
}

// SetSent gets a reference to the given int32 and assigns it to the Sent field.
func (o *MultiSendMessageResponse) SetSent(v int32) {
	o.Sent = &v
}

// GetFailed returns the Failed field value if set, zero value otherwise.
func (o *MultiSendMessageResponse) GetFailed() int32 {
	if o == nil || o.Failed == nil {
		var ret int32
		return ret
	}
	return *o.Failed
}

// GetFailedOk returns a tuple with the Failed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultiSendMessageResponse) GetFailedOk() (*int32, bool) {
	if o == nil || o.Failed == nil {
		return nil, false
	}
	return o.Failed, true
}

// HasFailed returns a boolean if a field has been set.
func (o *MultiSendMessageResponse) HasFailed() bool {
	if o != nil && o.Failed != nil {
		return true
	}

	return false
}

// SetFailed gets a reference to the given int32 and assigns it to the Failed field.
func (o *MultiSendMessageResponse) SetFailed(v int32) {
	o.Failed = &v
}

func (o MultiSendMessageResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Errors != nil {
		toSerialize["errors"] = o.Errors
	}
	if o.Sent != nil {
		toSerialize["sent"] = o.Sent
	}
	if o.Failed != nil {
		toSerialize["failed"] = o.Failed
	}
	return json.Marshal(toSerialize)
}

type NullableMultiSendMessageResponse struct {
	value *MultiSendMessageResponse
	isSet bool
}

func (v NullableMultiSendMessageResponse) Get() *MultiSendMessageResponse {
	return v.value
}

func (v *NullableMultiSendMessageResponse) Set(val *MultiSendMessageResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMultiSendMessageResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMultiSendMessageResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMultiSendMessageResponse(val *MultiSendMessageResponse) *NullableMultiSendMessageResponse {
	return &NullableMultiSendMessageResponse{value: val, isSet: true}
}

func (v NullableMultiSendMessageResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMultiSendMessageResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


