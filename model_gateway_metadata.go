/*
The Span API

API for device, collection, output and firmware management

API version: 4.4.2 nonviolent-adelbert
Contact: dev@lab5e.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package spanapi

import (
	"encoding/json"
)

// GatewayMetadata Metadata for gateway transports. The actual contents varies from gateway to gateway
type GatewayMetadata struct {
	Metadata *map[string]string `json:"metadata,omitempty"`
}

// NewGatewayMetadata instantiates a new GatewayMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayMetadata() *GatewayMetadata {
	this := GatewayMetadata{}
	return &this
}

// NewGatewayMetadataWithDefaults instantiates a new GatewayMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayMetadataWithDefaults() *GatewayMetadata {
	this := GatewayMetadata{}
	return &this
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *GatewayMetadata) GetMetadata() map[string]string {
	if o == nil || o.Metadata == nil {
		var ret map[string]string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayMetadata) GetMetadataOk() (*map[string]string, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *GatewayMetadata) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *GatewayMetadata) SetMetadata(v map[string]string) {
	o.Metadata = &v
}

func (o GatewayMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	return json.Marshal(toSerialize)
}

type NullableGatewayMetadata struct {
	value *GatewayMetadata
	isSet bool
}

func (v NullableGatewayMetadata) Get() *GatewayMetadata {
	return v.value
}

func (v *NullableGatewayMetadata) Set(val *GatewayMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayMetadata(val *GatewayMetadata) *NullableGatewayMetadata {
	return &NullableGatewayMetadata{value: val, isSet: true}
}

func (v NullableGatewayMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


