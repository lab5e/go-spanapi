/*
The Span API

API for device, collection, output and firmware management

API version: 4.7.0 actionable-aryanna
Contact: dev@lab5e.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package spanapi

import (
	"encoding/json"
)

// checks if the CreateOutputRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateOutputRequest{}

// CreateOutputRequest Request type when creating new outputs
type CreateOutputRequest struct {
	Type *OutputType `json:"type,omitempty"`
	Config *OutputConfig `json:"config,omitempty"`
	Enabled *bool `json:"enabled,omitempty"`
	Tags *map[string]string `json:"tags,omitempty"`
}

// NewCreateOutputRequest instantiates a new CreateOutputRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOutputRequest() *CreateOutputRequest {
	this := CreateOutputRequest{}
	var type_ OutputType = OUTPUTTYPE_UNDEFINED
	this.Type = &type_
	return &this
}

// NewCreateOutputRequestWithDefaults instantiates a new CreateOutputRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOutputRequestWithDefaults() *CreateOutputRequest {
	this := CreateOutputRequest{}
	var type_ OutputType = OUTPUTTYPE_UNDEFINED
	this.Type = &type_
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CreateOutputRequest) GetType() OutputType {
	if o == nil || IsNil(o.Type) {
		var ret OutputType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOutputRequest) GetTypeOk() (*OutputType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CreateOutputRequest) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given OutputType and assigns it to the Type field.
func (o *CreateOutputRequest) SetType(v OutputType) {
	o.Type = &v
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *CreateOutputRequest) GetConfig() OutputConfig {
	if o == nil || IsNil(o.Config) {
		var ret OutputConfig
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOutputRequest) GetConfigOk() (*OutputConfig, bool) {
	if o == nil || IsNil(o.Config) {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *CreateOutputRequest) HasConfig() bool {
	if o != nil && !IsNil(o.Config) {
		return true
	}

	return false
}

// SetConfig gets a reference to the given OutputConfig and assigns it to the Config field.
func (o *CreateOutputRequest) SetConfig(v OutputConfig) {
	o.Config = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *CreateOutputRequest) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOutputRequest) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *CreateOutputRequest) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *CreateOutputRequest) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *CreateOutputRequest) GetTags() map[string]string {
	if o == nil || IsNil(o.Tags) {
		var ret map[string]string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOutputRequest) GetTagsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *CreateOutputRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given map[string]string and assigns it to the Tags field.
func (o *CreateOutputRequest) SetTags(v map[string]string) {
	o.Tags = &v
}

func (o CreateOutputRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateOutputRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Config) {
		toSerialize["config"] = o.Config
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	return toSerialize, nil
}

type NullableCreateOutputRequest struct {
	value *CreateOutputRequest
	isSet bool
}

func (v NullableCreateOutputRequest) Get() *CreateOutputRequest {
	return v.value
}

func (v *NullableCreateOutputRequest) Set(val *CreateOutputRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOutputRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOutputRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOutputRequest(val *CreateOutputRequest) *NullableCreateOutputRequest {
	return &NullableCreateOutputRequest{value: val, isSet: true}
}

func (v NullableCreateOutputRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOutputRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


