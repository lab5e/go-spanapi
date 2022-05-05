/*
The Span API

API for device, collection, output and firmware management

API version: 4.2.2 smarmy-derik
Contact: dev@lab5e.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package spanapi

import (
	"encoding/json"
)

// Output Output resource. The configuration depends on the kind of output type. There are five outputs: Webhooks, UDP forwarding, IFTTT events, MQTT client and MQTT broker. The MQTT broker output is just used to configure the built-in MQTT broker in Span.
type Output struct {
	OutputId *string `json:"outputId,omitempty"`
	CollectionId *string `json:"collectionId,omitempty"`
	Type *OutputType `json:"type,omitempty"`
	Config *OutputConfig `json:"config,omitempty"`
	Enabled *bool `json:"enabled,omitempty"`
	Tags *map[string]string `json:"tags,omitempty"`
}

// NewOutput instantiates a new Output object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOutput() *Output {
	this := Output{}
	var type_ OutputType = OUTPUTTYPE_UNDEFINED
	this.Type = &type_
	return &this
}

// NewOutputWithDefaults instantiates a new Output object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOutputWithDefaults() *Output {
	this := Output{}
	var type_ OutputType = OUTPUTTYPE_UNDEFINED
	this.Type = &type_
	return &this
}

// GetOutputId returns the OutputId field value if set, zero value otherwise.
func (o *Output) GetOutputId() string {
	if o == nil || o.OutputId == nil {
		var ret string
		return ret
	}
	return *o.OutputId
}

// GetOutputIdOk returns a tuple with the OutputId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Output) GetOutputIdOk() (*string, bool) {
	if o == nil || o.OutputId == nil {
		return nil, false
	}
	return o.OutputId, true
}

// HasOutputId returns a boolean if a field has been set.
func (o *Output) HasOutputId() bool {
	if o != nil && o.OutputId != nil {
		return true
	}

	return false
}

// SetOutputId gets a reference to the given string and assigns it to the OutputId field.
func (o *Output) SetOutputId(v string) {
	o.OutputId = &v
}

// GetCollectionId returns the CollectionId field value if set, zero value otherwise.
func (o *Output) GetCollectionId() string {
	if o == nil || o.CollectionId == nil {
		var ret string
		return ret
	}
	return *o.CollectionId
}

// GetCollectionIdOk returns a tuple with the CollectionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Output) GetCollectionIdOk() (*string, bool) {
	if o == nil || o.CollectionId == nil {
		return nil, false
	}
	return o.CollectionId, true
}

// HasCollectionId returns a boolean if a field has been set.
func (o *Output) HasCollectionId() bool {
	if o != nil && o.CollectionId != nil {
		return true
	}

	return false
}

// SetCollectionId gets a reference to the given string and assigns it to the CollectionId field.
func (o *Output) SetCollectionId(v string) {
	o.CollectionId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Output) GetType() OutputType {
	if o == nil || o.Type == nil {
		var ret OutputType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Output) GetTypeOk() (*OutputType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Output) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given OutputType and assigns it to the Type field.
func (o *Output) SetType(v OutputType) {
	o.Type = &v
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *Output) GetConfig() OutputConfig {
	if o == nil || o.Config == nil {
		var ret OutputConfig
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Output) GetConfigOk() (*OutputConfig, bool) {
	if o == nil || o.Config == nil {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *Output) HasConfig() bool {
	if o != nil && o.Config != nil {
		return true
	}

	return false
}

// SetConfig gets a reference to the given OutputConfig and assigns it to the Config field.
func (o *Output) SetConfig(v OutputConfig) {
	o.Config = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *Output) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Output) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *Output) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *Output) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *Output) GetTags() map[string]string {
	if o == nil || o.Tags == nil {
		var ret map[string]string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Output) GetTagsOk() (*map[string]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *Output) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given map[string]string and assigns it to the Tags field.
func (o *Output) SetTags(v map[string]string) {
	o.Tags = &v
}

func (o Output) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.OutputId != nil {
		toSerialize["outputId"] = o.OutputId
	}
	if o.CollectionId != nil {
		toSerialize["collectionId"] = o.CollectionId
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Config != nil {
		toSerialize["config"] = o.Config
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	return json.Marshal(toSerialize)
}

type NullableOutput struct {
	value *Output
	isSet bool
}

func (v NullableOutput) Get() *Output {
	return v.value
}

func (v *NullableOutput) Set(val *Output) {
	v.value = val
	v.isSet = true
}

func (v NullableOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOutput(val *Output) *NullableOutput {
	return &NullableOutput{value: val, isSet: true}
}

func (v NullableOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


