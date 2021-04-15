/*
 * The Span API
 *
 * API for device, collection, output and firmware management
 *
 * API version: 4.1.7 only-deshaun
 * Contact: dev@lab5e.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package spanapi

import (
	"encoding/json"
)

// DataDumpResponse struct for DataDumpResponse
type DataDumpResponse struct {
	Collections *[]DumpedCollection `json:"collections,omitempty"`
}

// NewDataDumpResponse instantiates a new DataDumpResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataDumpResponse() *DataDumpResponse {
	this := DataDumpResponse{}
	return &this
}

// NewDataDumpResponseWithDefaults instantiates a new DataDumpResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataDumpResponseWithDefaults() *DataDumpResponse {
	this := DataDumpResponse{}
	return &this
}

// GetCollections returns the Collections field value if set, zero value otherwise.
func (o *DataDumpResponse) GetCollections() []DumpedCollection {
	if o == nil || o.Collections == nil {
		var ret []DumpedCollection
		return ret
	}
	return *o.Collections
}

// GetCollectionsOk returns a tuple with the Collections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataDumpResponse) GetCollectionsOk() (*[]DumpedCollection, bool) {
	if o == nil || o.Collections == nil {
		return nil, false
	}
	return o.Collections, true
}

// HasCollections returns a boolean if a field has been set.
func (o *DataDumpResponse) HasCollections() bool {
	if o != nil && o.Collections != nil {
		return true
	}

	return false
}

// SetCollections gets a reference to the given []DumpedCollection and assigns it to the Collections field.
func (o *DataDumpResponse) SetCollections(v []DumpedCollection) {
	o.Collections = &v
}

func (o DataDumpResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Collections != nil {
		toSerialize["collections"] = o.Collections
	}
	return json.Marshal(toSerialize)
}

type NullableDataDumpResponse struct {
	value *DataDumpResponse
	isSet bool
}

func (v NullableDataDumpResponse) Get() *DataDumpResponse {
	return v.value
}

func (v *NullableDataDumpResponse) Set(val *DataDumpResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDataDumpResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDataDumpResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataDumpResponse(val *DataDumpResponse) *NullableDataDumpResponse {
	return &NullableDataDumpResponse{value: val, isSet: true}
}

func (v NullableDataDumpResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataDumpResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


