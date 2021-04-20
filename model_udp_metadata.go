/*
 * The Span API
 *
 * API for device, collection, output and firmware management
 *
 * API version: 4.1.8 adopted-kali
 * Contact: dev@lab5e.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package spanapi

import (
	"encoding/json"
)

// UDPMetadata struct for UDPMetadata
type UDPMetadata struct {
	LocalPort  *int32 `json:"localPort,omitempty"`
	RemotePort *int32 `json:"remotePort,omitempty"`
}

// NewUDPMetadata instantiates a new UDPMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUDPMetadata() *UDPMetadata {
	this := UDPMetadata{}
	return &this
}

// NewUDPMetadataWithDefaults instantiates a new UDPMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUDPMetadataWithDefaults() *UDPMetadata {
	this := UDPMetadata{}
	return &this
}

// GetLocalPort returns the LocalPort field value if set, zero value otherwise.
func (o *UDPMetadata) GetLocalPort() int32 {
	if o == nil || o.LocalPort == nil {
		var ret int32
		return ret
	}
	return *o.LocalPort
}

// GetLocalPortOk returns a tuple with the LocalPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UDPMetadata) GetLocalPortOk() (*int32, bool) {
	if o == nil || o.LocalPort == nil {
		return nil, false
	}
	return o.LocalPort, true
}

// HasLocalPort returns a boolean if a field has been set.
func (o *UDPMetadata) HasLocalPort() bool {
	if o != nil && o.LocalPort != nil {
		return true
	}

	return false
}

// SetLocalPort gets a reference to the given int32 and assigns it to the LocalPort field.
func (o *UDPMetadata) SetLocalPort(v int32) {
	o.LocalPort = &v
}

// GetRemotePort returns the RemotePort field value if set, zero value otherwise.
func (o *UDPMetadata) GetRemotePort() int32 {
	if o == nil || o.RemotePort == nil {
		var ret int32
		return ret
	}
	return *o.RemotePort
}

// GetRemotePortOk returns a tuple with the RemotePort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UDPMetadata) GetRemotePortOk() (*int32, bool) {
	if o == nil || o.RemotePort == nil {
		return nil, false
	}
	return o.RemotePort, true
}

// HasRemotePort returns a boolean if a field has been set.
func (o *UDPMetadata) HasRemotePort() bool {
	if o != nil && o.RemotePort != nil {
		return true
	}

	return false
}

// SetRemotePort gets a reference to the given int32 and assigns it to the RemotePort field.
func (o *UDPMetadata) SetRemotePort(v int32) {
	o.RemotePort = &v
}

func (o UDPMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.LocalPort != nil {
		toSerialize["localPort"] = o.LocalPort
	}
	if o.RemotePort != nil {
		toSerialize["remotePort"] = o.RemotePort
	}
	return json.Marshal(toSerialize)
}

type NullableUDPMetadata struct {
	value *UDPMetadata
	isSet bool
}

func (v NullableUDPMetadata) Get() *UDPMetadata {
	return v.value
}

func (v *NullableUDPMetadata) Set(val *UDPMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableUDPMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableUDPMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUDPMetadata(val *UDPMetadata) *NullableUDPMetadata {
	return &NullableUDPMetadata{value: val, isSet: true}
}

func (v NullableUDPMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUDPMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
