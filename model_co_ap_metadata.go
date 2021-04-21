/*
 * The Span API
 *
 * API for device, collection, output and firmware management
 *
 * API version: 4.1.9 receding-glennis
 * Contact: dev@lab5e.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package spanapi

import (
	"encoding/json"
)

// CoAPMetadata struct for CoAPMetadata
type CoAPMetadata struct {
	Code *string `json:"code,omitempty"`
	Path *string `json:"path,omitempty"`
}

// NewCoAPMetadata instantiates a new CoAPMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoAPMetadata() *CoAPMetadata {
	this := CoAPMetadata{}
	return &this
}

// NewCoAPMetadataWithDefaults instantiates a new CoAPMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoAPMetadataWithDefaults() *CoAPMetadata {
	this := CoAPMetadata{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *CoAPMetadata) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoAPMetadata) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *CoAPMetadata) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *CoAPMetadata) SetCode(v string) {
	o.Code = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *CoAPMetadata) GetPath() string {
	if o == nil || o.Path == nil {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoAPMetadata) GetPathOk() (*string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *CoAPMetadata) HasPath() bool {
	if o != nil && o.Path != nil {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *CoAPMetadata) SetPath(v string) {
	o.Path = &v
}

func (o CoAPMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.Path != nil {
		toSerialize["path"] = o.Path
	}
	return json.Marshal(toSerialize)
}

type NullableCoAPMetadata struct {
	value *CoAPMetadata
	isSet bool
}

func (v NullableCoAPMetadata) Get() *CoAPMetadata {
	return v.value
}

func (v *NullableCoAPMetadata) Set(val *CoAPMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableCoAPMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableCoAPMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoAPMetadata(val *CoAPMetadata) *NullableCoAPMetadata {
	return &NullableCoAPMetadata{value: val, isSet: true}
}

func (v NullableCoAPMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoAPMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
