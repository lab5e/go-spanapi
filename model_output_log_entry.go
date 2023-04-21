/*
The Span API

API for device, collection, output and firmware management

API version: 4.5.0 overwrought-dorla
Contact: dev@lab5e.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package spanapi

import (
	"encoding/json"
)

// checks if the OutputLogEntry type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OutputLogEntry{}

// OutputLogEntry Log entries for outputs
type OutputLogEntry struct {
	Time *string `json:"time,omitempty"`
	Message *string `json:"message,omitempty"`
	Repeated *int32 `json:"repeated,omitempty"`
}

// NewOutputLogEntry instantiates a new OutputLogEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOutputLogEntry() *OutputLogEntry {
	this := OutputLogEntry{}
	return &this
}

// NewOutputLogEntryWithDefaults instantiates a new OutputLogEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOutputLogEntryWithDefaults() *OutputLogEntry {
	this := OutputLogEntry{}
	return &this
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *OutputLogEntry) GetTime() string {
	if o == nil || IsNil(o.Time) {
		var ret string
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputLogEntry) GetTimeOk() (*string, bool) {
	if o == nil || IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *OutputLogEntry) HasTime() bool {
	if o != nil && !IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given string and assigns it to the Time field.
func (o *OutputLogEntry) SetTime(v string) {
	o.Time = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *OutputLogEntry) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputLogEntry) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *OutputLogEntry) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *OutputLogEntry) SetMessage(v string) {
	o.Message = &v
}

// GetRepeated returns the Repeated field value if set, zero value otherwise.
func (o *OutputLogEntry) GetRepeated() int32 {
	if o == nil || IsNil(o.Repeated) {
		var ret int32
		return ret
	}
	return *o.Repeated
}

// GetRepeatedOk returns a tuple with the Repeated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputLogEntry) GetRepeatedOk() (*int32, bool) {
	if o == nil || IsNil(o.Repeated) {
		return nil, false
	}
	return o.Repeated, true
}

// HasRepeated returns a boolean if a field has been set.
func (o *OutputLogEntry) HasRepeated() bool {
	if o != nil && !IsNil(o.Repeated) {
		return true
	}

	return false
}

// SetRepeated gets a reference to the given int32 and assigns it to the Repeated field.
func (o *OutputLogEntry) SetRepeated(v int32) {
	o.Repeated = &v
}

func (o OutputLogEntry) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OutputLogEntry) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.Repeated) {
		toSerialize["repeated"] = o.Repeated
	}
	return toSerialize, nil
}

type NullableOutputLogEntry struct {
	value *OutputLogEntry
	isSet bool
}

func (v NullableOutputLogEntry) Get() *OutputLogEntry {
	return v.value
}

func (v *NullableOutputLogEntry) Set(val *OutputLogEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableOutputLogEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableOutputLogEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOutputLogEntry(val *OutputLogEntry) *NullableOutputLogEntry {
	return &NullableOutputLogEntry{value: val, isSet: true}
}

func (v NullableOutputLogEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOutputLogEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


