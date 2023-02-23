/*
The Span API

API for device, collection, output and firmware management

API version: 4.4.1 busy-janay
Contact: dev@lab5e.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package spanapi

import (
	"encoding/json"
)

// ListOutputResponse List outputs
type ListOutputResponse struct {
	CollectionId *string  `json:"collectionId,omitempty"`
	Outputs      []Output `json:"outputs,omitempty"`
}

// NewListOutputResponse instantiates a new ListOutputResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListOutputResponse() *ListOutputResponse {
	this := ListOutputResponse{}
	return &this
}

// NewListOutputResponseWithDefaults instantiates a new ListOutputResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListOutputResponseWithDefaults() *ListOutputResponse {
	this := ListOutputResponse{}
	return &this
}

// GetCollectionId returns the CollectionId field value if set, zero value otherwise.
func (o *ListOutputResponse) GetCollectionId() string {
	if o == nil || o.CollectionId == nil {
		var ret string
		return ret
	}
	return *o.CollectionId
}

// GetCollectionIdOk returns a tuple with the CollectionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListOutputResponse) GetCollectionIdOk() (*string, bool) {
	if o == nil || o.CollectionId == nil {
		return nil, false
	}
	return o.CollectionId, true
}

// HasCollectionId returns a boolean if a field has been set.
func (o *ListOutputResponse) HasCollectionId() bool {
	if o != nil && o.CollectionId != nil {
		return true
	}

	return false
}

// SetCollectionId gets a reference to the given string and assigns it to the CollectionId field.
func (o *ListOutputResponse) SetCollectionId(v string) {
	o.CollectionId = &v
}

// GetOutputs returns the Outputs field value if set, zero value otherwise.
func (o *ListOutputResponse) GetOutputs() []Output {
	if o == nil || o.Outputs == nil {
		var ret []Output
		return ret
	}
	return o.Outputs
}

// GetOutputsOk returns a tuple with the Outputs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListOutputResponse) GetOutputsOk() ([]Output, bool) {
	if o == nil || o.Outputs == nil {
		return nil, false
	}
	return o.Outputs, true
}

// HasOutputs returns a boolean if a field has been set.
func (o *ListOutputResponse) HasOutputs() bool {
	if o != nil && o.Outputs != nil {
		return true
	}

	return false
}

// SetOutputs gets a reference to the given []Output and assigns it to the Outputs field.
func (o *ListOutputResponse) SetOutputs(v []Output) {
	o.Outputs = v
}

func (o ListOutputResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CollectionId != nil {
		toSerialize["collectionId"] = o.CollectionId
	}
	if o.Outputs != nil {
		toSerialize["outputs"] = o.Outputs
	}
	return json.Marshal(toSerialize)
}

type NullableListOutputResponse struct {
	value *ListOutputResponse
	isSet bool
}

func (v NullableListOutputResponse) Get() *ListOutputResponse {
	return v.value
}

func (v *NullableListOutputResponse) Set(val *ListOutputResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListOutputResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListOutputResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListOutputResponse(val *ListOutputResponse) *NullableListOutputResponse {
	return &NullableListOutputResponse{value: val, isSet: true}
}

func (v NullableListOutputResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListOutputResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
