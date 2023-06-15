/*
The Span API

API for device, collection, output and firmware management

API version: 4.6.0 truthful-holli
Contact: dev@lab5e.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package spanapi

import (
	"encoding/json"
)

// checks if the OutputStats type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OutputStats{}

// OutputStats Statistics for a single data router
type OutputStats struct {
	ForwardErrors *int32 `json:"forwardErrors,omitempty"`
	MessagesForwarded *int32 `json:"messagesForwarded,omitempty"`
	BytesForwardedMb *float32 `json:"bytesForwardedMb,omitempty"`
}

// NewOutputStats instantiates a new OutputStats object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOutputStats() *OutputStats {
	this := OutputStats{}
	return &this
}

// NewOutputStatsWithDefaults instantiates a new OutputStats object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOutputStatsWithDefaults() *OutputStats {
	this := OutputStats{}
	return &this
}

// GetForwardErrors returns the ForwardErrors field value if set, zero value otherwise.
func (o *OutputStats) GetForwardErrors() int32 {
	if o == nil || IsNil(o.ForwardErrors) {
		var ret int32
		return ret
	}
	return *o.ForwardErrors
}

// GetForwardErrorsOk returns a tuple with the ForwardErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputStats) GetForwardErrorsOk() (*int32, bool) {
	if o == nil || IsNil(o.ForwardErrors) {
		return nil, false
	}
	return o.ForwardErrors, true
}

// HasForwardErrors returns a boolean if a field has been set.
func (o *OutputStats) HasForwardErrors() bool {
	if o != nil && !IsNil(o.ForwardErrors) {
		return true
	}

	return false
}

// SetForwardErrors gets a reference to the given int32 and assigns it to the ForwardErrors field.
func (o *OutputStats) SetForwardErrors(v int32) {
	o.ForwardErrors = &v
}

// GetMessagesForwarded returns the MessagesForwarded field value if set, zero value otherwise.
func (o *OutputStats) GetMessagesForwarded() int32 {
	if o == nil || IsNil(o.MessagesForwarded) {
		var ret int32
		return ret
	}
	return *o.MessagesForwarded
}

// GetMessagesForwardedOk returns a tuple with the MessagesForwarded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputStats) GetMessagesForwardedOk() (*int32, bool) {
	if o == nil || IsNil(o.MessagesForwarded) {
		return nil, false
	}
	return o.MessagesForwarded, true
}

// HasMessagesForwarded returns a boolean if a field has been set.
func (o *OutputStats) HasMessagesForwarded() bool {
	if o != nil && !IsNil(o.MessagesForwarded) {
		return true
	}

	return false
}

// SetMessagesForwarded gets a reference to the given int32 and assigns it to the MessagesForwarded field.
func (o *OutputStats) SetMessagesForwarded(v int32) {
	o.MessagesForwarded = &v
}

// GetBytesForwardedMb returns the BytesForwardedMb field value if set, zero value otherwise.
func (o *OutputStats) GetBytesForwardedMb() float32 {
	if o == nil || IsNil(o.BytesForwardedMb) {
		var ret float32
		return ret
	}
	return *o.BytesForwardedMb
}

// GetBytesForwardedMbOk returns a tuple with the BytesForwardedMb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputStats) GetBytesForwardedMbOk() (*float32, bool) {
	if o == nil || IsNil(o.BytesForwardedMb) {
		return nil, false
	}
	return o.BytesForwardedMb, true
}

// HasBytesForwardedMb returns a boolean if a field has been set.
func (o *OutputStats) HasBytesForwardedMb() bool {
	if o != nil && !IsNil(o.BytesForwardedMb) {
		return true
	}

	return false
}

// SetBytesForwardedMb gets a reference to the given float32 and assigns it to the BytesForwardedMb field.
func (o *OutputStats) SetBytesForwardedMb(v float32) {
	o.BytesForwardedMb = &v
}

func (o OutputStats) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OutputStats) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ForwardErrors) {
		toSerialize["forwardErrors"] = o.ForwardErrors
	}
	if !IsNil(o.MessagesForwarded) {
		toSerialize["messagesForwarded"] = o.MessagesForwarded
	}
	if !IsNil(o.BytesForwardedMb) {
		toSerialize["bytesForwardedMb"] = o.BytesForwardedMb
	}
	return toSerialize, nil
}

type NullableOutputStats struct {
	value *OutputStats
	isSet bool
}

func (v NullableOutputStats) Get() *OutputStats {
	return v.value
}

func (v *NullableOutputStats) Set(val *OutputStats) {
	v.value = val
	v.isSet = true
}

func (v NullableOutputStats) IsSet() bool {
	return v.isSet
}

func (v *NullableOutputStats) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOutputStats(val *OutputStats) *NullableOutputStats {
	return &NullableOutputStats{value: val, isSet: true}
}

func (v NullableOutputStats) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOutputStats) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


