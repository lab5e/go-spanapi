/*
The Span API

API for device, collection, output and firmware management

API version: 5.0.0 convulsive-launa
Contact: dev@lab5e.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package spanapi

import (
	"encoding/json"
)

// checks if the ListDevicesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListDevicesResponse{}

// ListDevicesResponse List device response
type ListDevicesResponse struct {
	Devices []Device `json:"devices,omitempty"`
}

// NewListDevicesResponse instantiates a new ListDevicesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListDevicesResponse() *ListDevicesResponse {
	this := ListDevicesResponse{}
	return &this
}

// NewListDevicesResponseWithDefaults instantiates a new ListDevicesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListDevicesResponseWithDefaults() *ListDevicesResponse {
	this := ListDevicesResponse{}
	return &this
}

// GetDevices returns the Devices field value if set, zero value otherwise.
func (o *ListDevicesResponse) GetDevices() []Device {
	if o == nil || IsNil(o.Devices) {
		var ret []Device
		return ret
	}
	return o.Devices
}

// GetDevicesOk returns a tuple with the Devices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListDevicesResponse) GetDevicesOk() ([]Device, bool) {
	if o == nil || IsNil(o.Devices) {
		return nil, false
	}
	return o.Devices, true
}

// HasDevices returns a boolean if a field has been set.
func (o *ListDevicesResponse) HasDevices() bool {
	if o != nil && !IsNil(o.Devices) {
		return true
	}

	return false
}

// SetDevices gets a reference to the given []Device and assigns it to the Devices field.
func (o *ListDevicesResponse) SetDevices(v []Device) {
	o.Devices = v
}

func (o ListDevicesResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListDevicesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Devices) {
		toSerialize["devices"] = o.Devices
	}
	return toSerialize, nil
}

type NullableListDevicesResponse struct {
	value *ListDevicesResponse
	isSet bool
}

func (v NullableListDevicesResponse) Get() *ListDevicesResponse {
	return v.value
}

func (v *NullableListDevicesResponse) Set(val *ListDevicesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListDevicesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListDevicesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListDevicesResponse(val *ListDevicesResponse) *NullableListDevicesResponse {
	return &NullableListDevicesResponse{value: val, isSet: true}
}

func (v NullableListDevicesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListDevicesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


