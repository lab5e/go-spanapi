/*
The Span API

API for device, collection, output and firmware management

API version: 4.9.5 spattered-kelvin
Contact: dev@lab5e.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package spanapi

import (
	"encoding/json"
)

// checks if the Any type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Any{}

// Any struct for Any
type Any struct {
	Type *string `json:"@type,omitempty"`
}

// NewAny instantiates a new Any object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAny() *Any {
	this := Any{}
	return &this
}

// NewAnyWithDefaults instantiates a new Any object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnyWithDefaults() *Any {
	this := Any{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Any) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Any) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Any) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Any) SetType(v string) {
	o.Type = &v
}

func (o Any) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Any) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["@type"] = o.Type
	}
	return toSerialize, nil
}

type NullableAny struct {
	value *Any
	isSet bool
}

func (v NullableAny) Get() *Any {
	return v.value
}

func (v *NullableAny) Set(val *Any) {
	v.value = val
	v.isSet = true
}

func (v NullableAny) IsSet() bool {
	return v.isSet
}

func (v *NullableAny) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAny(val *Any) *NullableAny {
	return &NullableAny{value: val, isSet: true}
}

func (v NullableAny) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAny) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


