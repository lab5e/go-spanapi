/*
The Span API

API for device, collection, output and firmware management

API version: 4.9.5 spattered-kelvin
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
	MessagesUpstream *string `json:"messagesUpstream,omitempty"`
	MessagesDownstream *string `json:"messagesDownstream,omitempty"`
	BytesUpstream *string `json:"bytesUpstream,omitempty"`
	BytesDownstream *string `json:"bytesDownstream,omitempty"`
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
func (o *GatewayStats) GetMessagesUpstream() string {
	if o == nil || IsNil(o.MessagesUpstream) {
		var ret string
		return ret
	}
	return *o.MessagesUpstream
}

// GetMessagesUpstreamOk returns a tuple with the MessagesUpstream field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayStats) GetMessagesUpstreamOk() (*string, bool) {
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

// SetMessagesUpstream gets a reference to the given string and assigns it to the MessagesUpstream field.
func (o *GatewayStats) SetMessagesUpstream(v string) {
	o.MessagesUpstream = &v
}

// GetMessagesDownstream returns the MessagesDownstream field value if set, zero value otherwise.
func (o *GatewayStats) GetMessagesDownstream() string {
	if o == nil || IsNil(o.MessagesDownstream) {
		var ret string
		return ret
	}
	return *o.MessagesDownstream
}

// GetMessagesDownstreamOk returns a tuple with the MessagesDownstream field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayStats) GetMessagesDownstreamOk() (*string, bool) {
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

// SetMessagesDownstream gets a reference to the given string and assigns it to the MessagesDownstream field.
func (o *GatewayStats) SetMessagesDownstream(v string) {
	o.MessagesDownstream = &v
}

// GetBytesUpstream returns the BytesUpstream field value if set, zero value otherwise.
func (o *GatewayStats) GetBytesUpstream() string {
	if o == nil || IsNil(o.BytesUpstream) {
		var ret string
		return ret
	}
	return *o.BytesUpstream
}

// GetBytesUpstreamOk returns a tuple with the BytesUpstream field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayStats) GetBytesUpstreamOk() (*string, bool) {
	if o == nil || IsNil(o.BytesUpstream) {
		return nil, false
	}
	return o.BytesUpstream, true
}

// HasBytesUpstream returns a boolean if a field has been set.
func (o *GatewayStats) HasBytesUpstream() bool {
	if o != nil && !IsNil(o.BytesUpstream) {
		return true
	}

	return false
}

// SetBytesUpstream gets a reference to the given string and assigns it to the BytesUpstream field.
func (o *GatewayStats) SetBytesUpstream(v string) {
	o.BytesUpstream = &v
}

// GetBytesDownstream returns the BytesDownstream field value if set, zero value otherwise.
func (o *GatewayStats) GetBytesDownstream() string {
	if o == nil || IsNil(o.BytesDownstream) {
		var ret string
		return ret
	}
	return *o.BytesDownstream
}

// GetBytesDownstreamOk returns a tuple with the BytesDownstream field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayStats) GetBytesDownstreamOk() (*string, bool) {
	if o == nil || IsNil(o.BytesDownstream) {
		return nil, false
	}
	return o.BytesDownstream, true
}

// HasBytesDownstream returns a boolean if a field has been set.
func (o *GatewayStats) HasBytesDownstream() bool {
	if o != nil && !IsNil(o.BytesDownstream) {
		return true
	}

	return false
}

// SetBytesDownstream gets a reference to the given string and assigns it to the BytesDownstream field.
func (o *GatewayStats) SetBytesDownstream(v string) {
	o.BytesDownstream = &v
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
	if !IsNil(o.BytesUpstream) {
		toSerialize["bytesUpstream"] = o.BytesUpstream
	}
	if !IsNil(o.BytesDownstream) {
		toSerialize["bytesDownstream"] = o.BytesDownstream
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


