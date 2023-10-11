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

// checks if the GatewayStats type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GatewayStats{}

// GatewayStats This is statistics for a single gateway
type GatewayStats struct {
	MessagesUpstream *int32 `json:"messagesUpstream,omitempty"`
	MessagesDownstream *int32 `json:"messagesDownstream,omitempty"`
	BytesUpstreamMb *float32 `json:"bytesUpstreamMb,omitempty"`
	BytesDownstreamMb *float32 `json:"bytesDownstreamMb,omitempty"`
}

// NewGatewayStats instantiates a new GatewayStats object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayStats() *GatewayStats {
	this := GatewayStats{}
	return &this
}

// NewGatewayStatsWithDefaults instantiates a new GatewayStats object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayStatsWithDefaults() *GatewayStats {
	this := GatewayStats{}
	return &this
}

// GetMessagesUpstream returns the MessagesUpstream field value if set, zero value otherwise.
func (o *GatewayStats) GetMessagesUpstream() int32 {
	if o == nil || IsNil(o.MessagesUpstream) {
		var ret int32
		return ret
	}
	return *o.MessagesUpstream
}

// GetMessagesUpstreamOk returns a tuple with the MessagesUpstream field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayStats) GetMessagesUpstreamOk() (*int32, bool) {
	if o == nil || IsNil(o.MessagesUpstream) {
		return nil, false
	}
	return o.MessagesUpstream, true
}

// HasMessagesUpstream returns a boolean if a field has been set.
func (o *GatewayStats) HasMessagesUpstream() bool {
	if o != nil && !IsNil(o.MessagesUpstream) {
		return true
	}

	return false
}

// SetMessagesUpstream gets a reference to the given int32 and assigns it to the MessagesUpstream field.
func (o *GatewayStats) SetMessagesUpstream(v int32) {
	o.MessagesUpstream = &v
}

// GetMessagesDownstream returns the MessagesDownstream field value if set, zero value otherwise.
func (o *GatewayStats) GetMessagesDownstream() int32 {
	if o == nil || IsNil(o.MessagesDownstream) {
		var ret int32
		return ret
	}
	return *o.MessagesDownstream
}

// GetMessagesDownstreamOk returns a tuple with the MessagesDownstream field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayStats) GetMessagesDownstreamOk() (*int32, bool) {
	if o == nil || IsNil(o.MessagesDownstream) {
		return nil, false
	}
	return o.MessagesDownstream, true
}

// HasMessagesDownstream returns a boolean if a field has been set.
func (o *GatewayStats) HasMessagesDownstream() bool {
	if o != nil && !IsNil(o.MessagesDownstream) {
		return true
	}

	return false
}

// SetMessagesDownstream gets a reference to the given int32 and assigns it to the MessagesDownstream field.
func (o *GatewayStats) SetMessagesDownstream(v int32) {
	o.MessagesDownstream = &v
}

// GetBytesUpstreamMb returns the BytesUpstreamMb field value if set, zero value otherwise.
func (o *GatewayStats) GetBytesUpstreamMb() float32 {
	if o == nil || IsNil(o.BytesUpstreamMb) {
		var ret float32
		return ret
	}
	return *o.BytesUpstreamMb
}

// GetBytesUpstreamMbOk returns a tuple with the BytesUpstreamMb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayStats) GetBytesUpstreamMbOk() (*float32, bool) {
	if o == nil || IsNil(o.BytesUpstreamMb) {
		return nil, false
	}
	return o.BytesUpstreamMb, true
}

// HasBytesUpstreamMb returns a boolean if a field has been set.
func (o *GatewayStats) HasBytesUpstreamMb() bool {
	if o != nil && !IsNil(o.BytesUpstreamMb) {
		return true
	}

	return false
}

// SetBytesUpstreamMb gets a reference to the given float32 and assigns it to the BytesUpstreamMb field.
func (o *GatewayStats) SetBytesUpstreamMb(v float32) {
	o.BytesUpstreamMb = &v
}

// GetBytesDownstreamMb returns the BytesDownstreamMb field value if set, zero value otherwise.
func (o *GatewayStats) GetBytesDownstreamMb() float32 {
	if o == nil || IsNil(o.BytesDownstreamMb) {
		var ret float32
		return ret
	}
	return *o.BytesDownstreamMb
}

// GetBytesDownstreamMbOk returns a tuple with the BytesDownstreamMb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayStats) GetBytesDownstreamMbOk() (*float32, bool) {
	if o == nil || IsNil(o.BytesDownstreamMb) {
		return nil, false
	}
	return o.BytesDownstreamMb, true
}

// HasBytesDownstreamMb returns a boolean if a field has been set.
func (o *GatewayStats) HasBytesDownstreamMb() bool {
	if o != nil && !IsNil(o.BytesDownstreamMb) {
		return true
	}

	return false
}

// SetBytesDownstreamMb gets a reference to the given float32 and assigns it to the BytesDownstreamMb field.
func (o *GatewayStats) SetBytesDownstreamMb(v float32) {
	o.BytesDownstreamMb = &v
}

func (o GatewayStats) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GatewayStats) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MessagesUpstream) {
		toSerialize["messagesUpstream"] = o.MessagesUpstream
	}
	if !IsNil(o.MessagesDownstream) {
		toSerialize["messagesDownstream"] = o.MessagesDownstream
	}
	if !IsNil(o.BytesUpstreamMb) {
		toSerialize["bytesUpstreamMb"] = o.BytesUpstreamMb
	}
	if !IsNil(o.BytesDownstreamMb) {
		toSerialize["bytesDownstreamMb"] = o.BytesDownstreamMb
	}
	return toSerialize, nil
}

type NullableGatewayStats struct {
	value *GatewayStats
	isSet bool
}

func (v NullableGatewayStats) Get() *GatewayStats {
	return v.value
}

func (v *NullableGatewayStats) Set(val *GatewayStats) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayStats) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayStats) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayStats(val *GatewayStats) *NullableGatewayStats {
	return &NullableGatewayStats{value: val, isSet: true}
}

func (v NullableGatewayStats) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayStats) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


