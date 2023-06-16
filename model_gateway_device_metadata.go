/*
The Span API

API for device, collection, output and firmware management

API version: 4.6.1 squirming-codi
Contact: dev@lab5e.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package spanapi

import (
	"encoding/json"
)

// checks if the GatewayDeviceMetadata type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GatewayDeviceMetadata{}

// GatewayDeviceMetadata Metadata for devices connected via user-managed gateways. This metadata shows the configuration for the last message transmission
type GatewayDeviceMetadata struct {
	GatewayId *string `json:"gatewayId,omitempty"`
	LastUpdate *string `json:"lastUpdate,omitempty"`
	Params *map[string]string `json:"params,omitempty"`
}

// NewGatewayDeviceMetadata instantiates a new GatewayDeviceMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayDeviceMetadata() *GatewayDeviceMetadata {
	this := GatewayDeviceMetadata{}
	return &this
}

// NewGatewayDeviceMetadataWithDefaults instantiates a new GatewayDeviceMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayDeviceMetadataWithDefaults() *GatewayDeviceMetadata {
	this := GatewayDeviceMetadata{}
	return &this
}

// GetGatewayId returns the GatewayId field value if set, zero value otherwise.
func (o *GatewayDeviceMetadata) GetGatewayId() string {
	if o == nil || IsNil(o.GatewayId) {
		var ret string
		return ret
	}
	return *o.GatewayId
}

// GetGatewayIdOk returns a tuple with the GatewayId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayDeviceMetadata) GetGatewayIdOk() (*string, bool) {
	if o == nil || IsNil(o.GatewayId) {
		return nil, false
	}
	return o.GatewayId, true
}

// HasGatewayId returns a boolean if a field has been set.
func (o *GatewayDeviceMetadata) HasGatewayId() bool {
	if o != nil && !IsNil(o.GatewayId) {
		return true
	}

	return false
}

// SetGatewayId gets a reference to the given string and assigns it to the GatewayId field.
func (o *GatewayDeviceMetadata) SetGatewayId(v string) {
	o.GatewayId = &v
}

// GetLastUpdate returns the LastUpdate field value if set, zero value otherwise.
func (o *GatewayDeviceMetadata) GetLastUpdate() string {
	if o == nil || IsNil(o.LastUpdate) {
		var ret string
		return ret
	}
	return *o.LastUpdate
}

// GetLastUpdateOk returns a tuple with the LastUpdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayDeviceMetadata) GetLastUpdateOk() (*string, bool) {
	if o == nil || IsNil(o.LastUpdate) {
		return nil, false
	}
	return o.LastUpdate, true
}

// HasLastUpdate returns a boolean if a field has been set.
func (o *GatewayDeviceMetadata) HasLastUpdate() bool {
	if o != nil && !IsNil(o.LastUpdate) {
		return true
	}

	return false
}

// SetLastUpdate gets a reference to the given string and assigns it to the LastUpdate field.
func (o *GatewayDeviceMetadata) SetLastUpdate(v string) {
	o.LastUpdate = &v
}

// GetParams returns the Params field value if set, zero value otherwise.
func (o *GatewayDeviceMetadata) GetParams() map[string]string {
	if o == nil || IsNil(o.Params) {
		var ret map[string]string
		return ret
	}
	return *o.Params
}

// GetParamsOk returns a tuple with the Params field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayDeviceMetadata) GetParamsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Params) {
		return nil, false
	}
	return o.Params, true
}

// HasParams returns a boolean if a field has been set.
func (o *GatewayDeviceMetadata) HasParams() bool {
	if o != nil && !IsNil(o.Params) {
		return true
	}

	return false
}

// SetParams gets a reference to the given map[string]string and assigns it to the Params field.
func (o *GatewayDeviceMetadata) SetParams(v map[string]string) {
	o.Params = &v
}

func (o GatewayDeviceMetadata) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GatewayDeviceMetadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GatewayId) {
		toSerialize["gatewayId"] = o.GatewayId
	}
	if !IsNil(o.LastUpdate) {
		toSerialize["lastUpdate"] = o.LastUpdate
	}
	if !IsNil(o.Params) {
		toSerialize["params"] = o.Params
	}
	return toSerialize, nil
}

type NullableGatewayDeviceMetadata struct {
	value *GatewayDeviceMetadata
	isSet bool
}

func (v NullableGatewayDeviceMetadata) Get() *GatewayDeviceMetadata {
	return v.value
}

func (v *NullableGatewayDeviceMetadata) Set(val *GatewayDeviceMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayDeviceMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayDeviceMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayDeviceMetadata(val *GatewayDeviceMetadata) *NullableGatewayDeviceMetadata {
	return &NullableGatewayDeviceMetadata{value: val, isSet: true}
}

func (v NullableGatewayDeviceMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayDeviceMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


