/*
 * The Span API
 *
 * API for device, collection, output and firmware management
 *
 * API version: 4.1.7 prized-adelbert
 * Contact: dev@lab5e.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package spanapi

import (
	"encoding/json"
)

// DataDumpRequest struct for DataDumpRequest
type DataDumpRequest struct {
	Comment *string `json:"comment,omitempty"`
}

// NewDataDumpRequest instantiates a new DataDumpRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataDumpRequest() *DataDumpRequest {
	this := DataDumpRequest{}
	return &this
}

// NewDataDumpRequestWithDefaults instantiates a new DataDumpRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataDumpRequestWithDefaults() *DataDumpRequest {
	this := DataDumpRequest{}
	return &this
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *DataDumpRequest) GetComment() string {
	if o == nil || o.Comment == nil {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataDumpRequest) GetCommentOk() (*string, bool) {
	if o == nil || o.Comment == nil {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *DataDumpRequest) HasComment() bool {
	if o != nil && o.Comment != nil {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *DataDumpRequest) SetComment(v string) {
	o.Comment = &v
}

func (o DataDumpRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Comment != nil {
		toSerialize["comment"] = o.Comment
	}
	return json.Marshal(toSerialize)
}

type NullableDataDumpRequest struct {
	value *DataDumpRequest
	isSet bool
}

func (v NullableDataDumpRequest) Get() *DataDumpRequest {
	return v.value
}

func (v *NullableDataDumpRequest) Set(val *DataDumpRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDataDumpRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDataDumpRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataDumpRequest(val *DataDumpRequest) *NullableDataDumpRequest {
	return &NullableDataDumpRequest{value: val, isSet: true}
}

func (v NullableDataDumpRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataDumpRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


