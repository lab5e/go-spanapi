/*
The Span API

API for device, collection, output and firmware management

API version: 4.9.6 authoritarian-betty
Contact: dev@lab5e.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package spanapi

import (
	"encoding/json"
)

// checks if the MQTTMetadata type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MQTTMetadata{}

// MQTTMetadata MQTT metadata for messages received through one of the MQTT endpoints. This is an EXPERIMENTAL feature.
type MQTTMetadata struct {
	Topic *string `json:"topic,omitempty"`
}

// NewMQTTMetadata instantiates a new MQTTMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMQTTMetadata() *MQTTMetadata {
	this := MQTTMetadata{}
	return &this
}

// NewMQTTMetadataWithDefaults instantiates a new MQTTMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMQTTMetadataWithDefaults() *MQTTMetadata {
	this := MQTTMetadata{}
	return &this
}

// GetTopic returns the Topic field value if set, zero value otherwise.
func (o *MQTTMetadata) GetTopic() string {
	if o == nil || IsNil(o.Topic) {
		var ret string
		return ret
	}
	return *o.Topic
}

// GetTopicOk returns a tuple with the Topic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MQTTMetadata) GetTopicOk() (*string, bool) {
	if o == nil || IsNil(o.Topic) {
		return nil, false
	}
	return o.Topic, true
}

// HasTopic returns a boolean if a field has been set.
func (o *MQTTMetadata) HasTopic() bool {
	if o != nil && !IsNil(o.Topic) {
		return true
	}

	return false
}

// SetTopic gets a reference to the given string and assigns it to the Topic field.
func (o *MQTTMetadata) SetTopic(v string) {
	o.Topic = &v
}

func (o MQTTMetadata) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MQTTMetadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Topic) {
		toSerialize["topic"] = o.Topic
	}
	return toSerialize, nil
}

type NullableMQTTMetadata struct {
	value *MQTTMetadata
	isSet bool
}

func (v NullableMQTTMetadata) Get() *MQTTMetadata {
	return v.value
}

func (v *NullableMQTTMetadata) Set(val *MQTTMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableMQTTMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableMQTTMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMQTTMetadata(val *MQTTMetadata) *NullableMQTTMetadata {
	return &NullableMQTTMetadata{value: val, isSet: true}
}

func (v NullableMQTTMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMQTTMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


