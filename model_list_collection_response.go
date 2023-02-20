/*
The Span API

API for device, collection, output and firmware management

API version: 4.4.0 lean-joline
Contact: dev@lab5e.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package spanapi

import (
	"encoding/json"
)

// ListCollectionResponse Collection list. The list contains all the collections you have access to.
type ListCollectionResponse struct {
	Collections []Collection `json:"collections,omitempty"`
}

// NewListCollectionResponse instantiates a new ListCollectionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListCollectionResponse() *ListCollectionResponse {
	this := ListCollectionResponse{}
	return &this
}

// NewListCollectionResponseWithDefaults instantiates a new ListCollectionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListCollectionResponseWithDefaults() *ListCollectionResponse {
	this := ListCollectionResponse{}
	return &this
}

// GetCollections returns the Collections field value if set, zero value otherwise.
func (o *ListCollectionResponse) GetCollections() []Collection {
	if o == nil || o.Collections == nil {
		var ret []Collection
		return ret
	}
	return o.Collections
}

// GetCollectionsOk returns a tuple with the Collections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListCollectionResponse) GetCollectionsOk() ([]Collection, bool) {
	if o == nil || o.Collections == nil {
		return nil, false
	}
	return o.Collections, true
}

// HasCollections returns a boolean if a field has been set.
func (o *ListCollectionResponse) HasCollections() bool {
	if o != nil && o.Collections != nil {
		return true
	}

	return false
}

// SetCollections gets a reference to the given []Collection and assigns it to the Collections field.
func (o *ListCollectionResponse) SetCollections(v []Collection) {
	o.Collections = v
}

func (o ListCollectionResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Collections != nil {
		toSerialize["collections"] = o.Collections
	}
	return json.Marshal(toSerialize)
}

type NullableListCollectionResponse struct {
	value *ListCollectionResponse
	isSet bool
}

func (v NullableListCollectionResponse) Get() *ListCollectionResponse {
	return v.value
}

func (v *NullableListCollectionResponse) Set(val *ListCollectionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListCollectionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListCollectionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListCollectionResponse(val *ListCollectionResponse) *NullableListCollectionResponse {
	return &NullableListCollectionResponse{value: val, isSet: true}
}

func (v NullableListCollectionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListCollectionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
