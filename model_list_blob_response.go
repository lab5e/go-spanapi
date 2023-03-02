/*
The Span API

API for device, collection, output and firmware management

API version: 4.4.2 nonviolent-adelbert
Contact: dev@lab5e.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package spanapi

import (
	"encoding/json"
)

// ListBlobResponse Response object when listing blobs for a collection
type ListBlobResponse struct {
	Blobs []Blob `json:"blobs,omitempty"`
}

// NewListBlobResponse instantiates a new ListBlobResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListBlobResponse() *ListBlobResponse {
	this := ListBlobResponse{}
	return &this
}

// NewListBlobResponseWithDefaults instantiates a new ListBlobResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListBlobResponseWithDefaults() *ListBlobResponse {
	this := ListBlobResponse{}
	return &this
}

// GetBlobs returns the Blobs field value if set, zero value otherwise.
func (o *ListBlobResponse) GetBlobs() []Blob {
	if o == nil || o.Blobs == nil {
		var ret []Blob
		return ret
	}
	return o.Blobs
}

// GetBlobsOk returns a tuple with the Blobs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListBlobResponse) GetBlobsOk() ([]Blob, bool) {
	if o == nil || o.Blobs == nil {
		return nil, false
	}
	return o.Blobs, true
}

// HasBlobs returns a boolean if a field has been set.
func (o *ListBlobResponse) HasBlobs() bool {
	if o != nil && o.Blobs != nil {
		return true
	}

	return false
}

// SetBlobs gets a reference to the given []Blob and assigns it to the Blobs field.
func (o *ListBlobResponse) SetBlobs(v []Blob) {
	o.Blobs = v
}

func (o ListBlobResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Blobs != nil {
		toSerialize["blobs"] = o.Blobs
	}
	return json.Marshal(toSerialize)
}

type NullableListBlobResponse struct {
	value *ListBlobResponse
	isSet bool
}

func (v NullableListBlobResponse) Get() *ListBlobResponse {
	return v.value
}

func (v *NullableListBlobResponse) Set(val *ListBlobResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListBlobResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListBlobResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListBlobResponse(val *ListBlobResponse) *NullableListBlobResponse {
	return &NullableListBlobResponse{value: val, isSet: true}
}

func (v NullableListBlobResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListBlobResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


