/*
The Span API

API for device, collection, output and firmware management

API version: 5.0.3 pitch-dark-elza
Contact: dev@lab5e.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package spanapi

import (
	"encoding/json"
)

// checks if the FirmwareUsageResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FirmwareUsageResponse{}

// FirmwareUsageResponse Firmware usage report
type FirmwareUsageResponse struct {
	ImageId *string `json:"imageId,omitempty"`
	Targeted []string `json:"targeted,omitempty"`
	Current []string `json:"current,omitempty"`
}

// NewFirmwareUsageResponse instantiates a new FirmwareUsageResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFirmwareUsageResponse() *FirmwareUsageResponse {
	this := FirmwareUsageResponse{}
	return &this
}

// NewFirmwareUsageResponseWithDefaults instantiates a new FirmwareUsageResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFirmwareUsageResponseWithDefaults() *FirmwareUsageResponse {
	this := FirmwareUsageResponse{}
	return &this
}

// GetImageId returns the ImageId field value if set, zero value otherwise.
func (o *FirmwareUsageResponse) GetImageId() string {
	if o == nil || IsNil(o.ImageId) {
		var ret string
		return ret
	}
	return *o.ImageId
}

// GetImageIdOk returns a tuple with the ImageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareUsageResponse) GetImageIdOk() (*string, bool) {
	if o == nil || IsNil(o.ImageId) {
		return nil, false
	}
	return o.ImageId, true
}

// HasImageId returns a boolean if a field has been set.
func (o *FirmwareUsageResponse) HasImageId() bool {
	if o != nil && !IsNil(o.ImageId) {
		return true
	}

	return false
}

// SetImageId gets a reference to the given string and assigns it to the ImageId field.
func (o *FirmwareUsageResponse) SetImageId(v string) {
	o.ImageId = &v
}

// GetTargeted returns the Targeted field value if set, zero value otherwise.
func (o *FirmwareUsageResponse) GetTargeted() []string {
	if o == nil || IsNil(o.Targeted) {
		var ret []string
		return ret
	}
	return o.Targeted
}

// GetTargetedOk returns a tuple with the Targeted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareUsageResponse) GetTargetedOk() ([]string, bool) {
	if o == nil || IsNil(o.Targeted) {
		return nil, false
	}
	return o.Targeted, true
}

// HasTargeted returns a boolean if a field has been set.
func (o *FirmwareUsageResponse) HasTargeted() bool {
	if o != nil && !IsNil(o.Targeted) {
		return true
	}

	return false
}

// SetTargeted gets a reference to the given []string and assigns it to the Targeted field.
func (o *FirmwareUsageResponse) SetTargeted(v []string) {
	o.Targeted = v
}

// GetCurrent returns the Current field value if set, zero value otherwise.
func (o *FirmwareUsageResponse) GetCurrent() []string {
	if o == nil || IsNil(o.Current) {
		var ret []string
		return ret
	}
	return o.Current
}

// GetCurrentOk returns a tuple with the Current field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareUsageResponse) GetCurrentOk() ([]string, bool) {
	if o == nil || IsNil(o.Current) {
		return nil, false
	}
	return o.Current, true
}

// HasCurrent returns a boolean if a field has been set.
func (o *FirmwareUsageResponse) HasCurrent() bool {
	if o != nil && !IsNil(o.Current) {
		return true
	}

	return false
}

// SetCurrent gets a reference to the given []string and assigns it to the Current field.
func (o *FirmwareUsageResponse) SetCurrent(v []string) {
	o.Current = v
}

func (o FirmwareUsageResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FirmwareUsageResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ImageId) {
		toSerialize["imageId"] = o.ImageId
	}
	if !IsNil(o.Targeted) {
		toSerialize["targeted"] = o.Targeted
	}
	if !IsNil(o.Current) {
		toSerialize["current"] = o.Current
	}
	return toSerialize, nil
}

type NullableFirmwareUsageResponse struct {
	value *FirmwareUsageResponse
	isSet bool
}

func (v NullableFirmwareUsageResponse) Get() *FirmwareUsageResponse {
	return v.value
}

func (v *NullableFirmwareUsageResponse) Set(val *FirmwareUsageResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableFirmwareUsageResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableFirmwareUsageResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFirmwareUsageResponse(val *FirmwareUsageResponse) *NullableFirmwareUsageResponse {
	return &NullableFirmwareUsageResponse{value: val, isSet: true}
}

func (v NullableFirmwareUsageResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFirmwareUsageResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


