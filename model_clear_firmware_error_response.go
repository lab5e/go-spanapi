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

// ClearFirmwareErrorResponse struct for ClearFirmwareErrorResponse
type ClearFirmwareErrorResponse struct {
	Result *string `json:"result,omitempty"`
}

// NewClearFirmwareErrorResponse instantiates a new ClearFirmwareErrorResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClearFirmwareErrorResponse() *ClearFirmwareErrorResponse {
	this := ClearFirmwareErrorResponse{}
	return &this
}

// NewClearFirmwareErrorResponseWithDefaults instantiates a new ClearFirmwareErrorResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClearFirmwareErrorResponseWithDefaults() *ClearFirmwareErrorResponse {
	this := ClearFirmwareErrorResponse{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *ClearFirmwareErrorResponse) GetResult() string {
	if o == nil || o.Result == nil {
		var ret string
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClearFirmwareErrorResponse) GetResultOk() (*string, bool) {
	if o == nil || o.Result == nil {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *ClearFirmwareErrorResponse) HasResult() bool {
	if o != nil && o.Result != nil {
		return true
	}

	return false
}

// SetResult gets a reference to the given string and assigns it to the Result field.
func (o *ClearFirmwareErrorResponse) SetResult(v string) {
	o.Result = &v
}

func (o ClearFirmwareErrorResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Result != nil {
		toSerialize["result"] = o.Result
	}
	return json.Marshal(toSerialize)
}

type NullableClearFirmwareErrorResponse struct {
	value *ClearFirmwareErrorResponse
	isSet bool
}

func (v NullableClearFirmwareErrorResponse) Get() *ClearFirmwareErrorResponse {
	return v.value
}

func (v *NullableClearFirmwareErrorResponse) Set(val *ClearFirmwareErrorResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableClearFirmwareErrorResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableClearFirmwareErrorResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClearFirmwareErrorResponse(val *ClearFirmwareErrorResponse) *NullableClearFirmwareErrorResponse {
	return &NullableClearFirmwareErrorResponse{value: val, isSet: true}
}

func (v NullableClearFirmwareErrorResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClearFirmwareErrorResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
