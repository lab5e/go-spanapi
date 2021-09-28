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

// NetworkMetadata Network metadata for devices.
type NetworkMetadata struct {
	// Allocated IP address.
	AllocatedIp *string `json:"allocatedIp,omitempty"`
	AllocatedAt *string `json:"allocatedAt,omitempty"`
	// Cell ID for device. This might not be set, depending on the provider in use.
	CellId *string `json:"cellId,omitempty"`
}

// NewNetworkMetadata instantiates a new NetworkMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkMetadata() *NetworkMetadata {
	this := NetworkMetadata{}
	return &this
}

// NewNetworkMetadataWithDefaults instantiates a new NetworkMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkMetadataWithDefaults() *NetworkMetadata {
	this := NetworkMetadata{}
	return &this
}

// GetAllocatedIp returns the AllocatedIp field value if set, zero value otherwise.
func (o *NetworkMetadata) GetAllocatedIp() string {
	if o == nil || o.AllocatedIp == nil {
		var ret string
		return ret
	}
	return *o.AllocatedIp
}

// GetAllocatedIpOk returns a tuple with the AllocatedIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkMetadata) GetAllocatedIpOk() (*string, bool) {
	if o == nil || o.AllocatedIp == nil {
		return nil, false
	}
	return o.AllocatedIp, true
}

// HasAllocatedIp returns a boolean if a field has been set.
func (o *NetworkMetadata) HasAllocatedIp() bool {
	if o != nil && o.AllocatedIp != nil {
		return true
	}

	return false
}

// SetAllocatedIp gets a reference to the given string and assigns it to the AllocatedIp field.
func (o *NetworkMetadata) SetAllocatedIp(v string) {
	o.AllocatedIp = &v
}

// GetAllocatedAt returns the AllocatedAt field value if set, zero value otherwise.
func (o *NetworkMetadata) GetAllocatedAt() string {
	if o == nil || o.AllocatedAt == nil {
		var ret string
		return ret
	}
	return *o.AllocatedAt
}

// GetAllocatedAtOk returns a tuple with the AllocatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkMetadata) GetAllocatedAtOk() (*string, bool) {
	if o == nil || o.AllocatedAt == nil {
		return nil, false
	}
	return o.AllocatedAt, true
}

// HasAllocatedAt returns a boolean if a field has been set.
func (o *NetworkMetadata) HasAllocatedAt() bool {
	if o != nil && o.AllocatedAt != nil {
		return true
	}

	return false
}

// SetAllocatedAt gets a reference to the given string and assigns it to the AllocatedAt field.
func (o *NetworkMetadata) SetAllocatedAt(v string) {
	o.AllocatedAt = &v
}

// GetCellId returns the CellId field value if set, zero value otherwise.
func (o *NetworkMetadata) GetCellId() string {
	if o == nil || o.CellId == nil {
		var ret string
		return ret
	}
	return *o.CellId
}

// GetCellIdOk returns a tuple with the CellId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkMetadata) GetCellIdOk() (*string, bool) {
	if o == nil || o.CellId == nil {
		return nil, false
	}
	return o.CellId, true
}

// HasCellId returns a boolean if a field has been set.
func (o *NetworkMetadata) HasCellId() bool {
	if o != nil && o.CellId != nil {
		return true
	}

	return false
}

// SetCellId gets a reference to the given string and assigns it to the CellId field.
func (o *NetworkMetadata) SetCellId(v string) {
	o.CellId = &v
}

func (o NetworkMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AllocatedIp != nil {
		toSerialize["allocatedIp"] = o.AllocatedIp
	}
	if o.AllocatedAt != nil {
		toSerialize["allocatedAt"] = o.AllocatedAt
	}
	if o.CellId != nil {
		toSerialize["cellId"] = o.CellId
	}
	return json.Marshal(toSerialize)
}

type NullableNetworkMetadata struct {
	value *NetworkMetadata
	isSet bool
}

func (v NullableNetworkMetadata) Get() *NetworkMetadata {
	return v.value
}

func (v *NullableNetworkMetadata) Set(val *NetworkMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkMetadata(val *NetworkMetadata) *NullableNetworkMetadata {
	return &NullableNetworkMetadata{value: val, isSet: true}
}

func (v NullableNetworkMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
