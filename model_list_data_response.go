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

// checks if the ListDataResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListDataResponse{}

// ListDataResponse List of device payloads
type ListDataResponse struct {
	Data []OutputDataMessage `json:"data,omitempty"`
}

// NewListDataResponse instantiates a new ListDataResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListDataResponse() *ListDataResponse {
	this := ListDataResponse{}
	return &this
}

// NewListDataResponseWithDefaults instantiates a new ListDataResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListDataResponseWithDefaults() *ListDataResponse {
	this := ListDataResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ListDataResponse) GetData() []OutputDataMessage {
	if o == nil || IsNil(o.Data) {
		var ret []OutputDataMessage
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListDataResponse) GetDataOk() ([]OutputDataMessage, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ListDataResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []OutputDataMessage and assigns it to the Data field.
func (o *ListDataResponse) SetData(v []OutputDataMessage) {
	o.Data = v
}

func (o ListDataResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListDataResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableListDataResponse struct {
	value *ListDataResponse
	isSet bool
}

func (v NullableListDataResponse) Get() *ListDataResponse {
	return v.value
}

func (v *NullableListDataResponse) Set(val *ListDataResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListDataResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListDataResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListDataResponse(val *ListDataResponse) *NullableListDataResponse {
	return &NullableListDataResponse{value: val, isSet: true}
}

func (v NullableListDataResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListDataResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


