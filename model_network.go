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

// Network Network is the network devices runs on. They communicate via gateway to Span.
type Network struct {
	NetworkId *string `json:"networkId,omitempty"`
	Tags *map[string]string `json:"tags,omitempty"`
}

// NewNetwork instantiates a new Network object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetwork() *Network {
	this := Network{}
	return &this
}

// NewNetworkWithDefaults instantiates a new Network object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkWithDefaults() *Network {
	this := Network{}
	return &this
}

// GetNetworkId returns the NetworkId field value if set, zero value otherwise.
func (o *Network) GetNetworkId() string {
	if o == nil || o.NetworkId == nil {
		var ret string
		return ret
	}
	return *o.NetworkId
}

// GetNetworkIdOk returns a tuple with the NetworkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Network) GetNetworkIdOk() (*string, bool) {
	if o == nil || o.NetworkId == nil {
		return nil, false
	}
	return o.NetworkId, true
}

// HasNetworkId returns a boolean if a field has been set.
func (o *Network) HasNetworkId() bool {
	if o != nil && o.NetworkId != nil {
		return true
	}

	return false
}

// SetNetworkId gets a reference to the given string and assigns it to the NetworkId field.
func (o *Network) SetNetworkId(v string) {
	o.NetworkId = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *Network) GetTags() map[string]string {
	if o == nil || o.Tags == nil {
		var ret map[string]string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Network) GetTagsOk() (*map[string]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *Network) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given map[string]string and assigns it to the Tags field.
func (o *Network) SetTags(v map[string]string) {
	o.Tags = &v
}

func (o Network) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.NetworkId != nil {
		toSerialize["networkId"] = o.NetworkId
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	return json.Marshal(toSerialize)
}

type NullableNetwork struct {
	value *Network
	isSet bool
}

func (v NullableNetwork) Get() *Network {
	return v.value
}

func (v *NullableNetwork) Set(val *Network) {
	v.value = val
	v.isSet = true
}

func (v NullableNetwork) IsSet() bool {
	return v.isSet
}

func (v *NullableNetwork) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetwork(val *Network) *NullableNetwork {
	return &NullableNetwork{value: val, isSet: true}
}

func (v NullableNetwork) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetwork) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


