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

// OutputLogResponse struct for OutputLogResponse
type OutputLogResponse struct {
	Logs *[]OutputLogEntry `json:"logs,omitempty"`
}

// NewOutputLogResponse instantiates a new OutputLogResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOutputLogResponse() *OutputLogResponse {
	this := OutputLogResponse{}
	return &this
}

// NewOutputLogResponseWithDefaults instantiates a new OutputLogResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOutputLogResponseWithDefaults() *OutputLogResponse {
	this := OutputLogResponse{}
	return &this
}

// GetLogs returns the Logs field value if set, zero value otherwise.
func (o *OutputLogResponse) GetLogs() []OutputLogEntry {
	if o == nil || o.Logs == nil {
		var ret []OutputLogEntry
		return ret
	}
	return *o.Logs
}

// GetLogsOk returns a tuple with the Logs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputLogResponse) GetLogsOk() (*[]OutputLogEntry, bool) {
	if o == nil || o.Logs == nil {
		return nil, false
	}
	return o.Logs, true
}

// HasLogs returns a boolean if a field has been set.
func (o *OutputLogResponse) HasLogs() bool {
	if o != nil && o.Logs != nil {
		return true
	}

	return false
}

// SetLogs gets a reference to the given []OutputLogEntry and assigns it to the Logs field.
func (o *OutputLogResponse) SetLogs(v []OutputLogEntry) {
	o.Logs = &v
}

func (o OutputLogResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Logs != nil {
		toSerialize["logs"] = o.Logs
	}
	return json.Marshal(toSerialize)
}

type NullableOutputLogResponse struct {
	value *OutputLogResponse
	isSet bool
}

func (v NullableOutputLogResponse) Get() *OutputLogResponse {
	return v.value
}

func (v *NullableOutputLogResponse) Set(val *OutputLogResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOutputLogResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOutputLogResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOutputLogResponse(val *OutputLogResponse) *NullableOutputLogResponse {
	return &NullableOutputLogResponse{value: val, isSet: true}
}

func (v NullableOutputLogResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOutputLogResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


