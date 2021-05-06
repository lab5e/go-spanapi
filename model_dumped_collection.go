/*
 * The Span API
 *
 * API for device, collection, output and firmware management
 *
 * API version: 4.1.14 oversensitive-deante
 * Contact: dev@lab5e.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package spanapi

import (
	"encoding/json"
)

// DumpedCollection struct for DumpedCollection
type DumpedCollection struct {
	Collection *Collection     `json:"collection,omitempty"`
	Devices    *[]DumpedDevice `json:"devices,omitempty"`
	Outputs    *[]Output       `json:"outputs,omitempty"`
}

// NewDumpedCollection instantiates a new DumpedCollection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDumpedCollection() *DumpedCollection {
	this := DumpedCollection{}
	return &this
}

// NewDumpedCollectionWithDefaults instantiates a new DumpedCollection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDumpedCollectionWithDefaults() *DumpedCollection {
	this := DumpedCollection{}
	return &this
}

// GetCollection returns the Collection field value if set, zero value otherwise.
func (o *DumpedCollection) GetCollection() Collection {
	if o == nil || o.Collection == nil {
		var ret Collection
		return ret
	}
	return *o.Collection
}

// GetCollectionOk returns a tuple with the Collection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DumpedCollection) GetCollectionOk() (*Collection, bool) {
	if o == nil || o.Collection == nil {
		return nil, false
	}
	return o.Collection, true
}

// HasCollection returns a boolean if a field has been set.
func (o *DumpedCollection) HasCollection() bool {
	if o != nil && o.Collection != nil {
		return true
	}

	return false
}

// SetCollection gets a reference to the given Collection and assigns it to the Collection field.
func (o *DumpedCollection) SetCollection(v Collection) {
	o.Collection = &v
}

// GetDevices returns the Devices field value if set, zero value otherwise.
func (o *DumpedCollection) GetDevices() []DumpedDevice {
	if o == nil || o.Devices == nil {
		var ret []DumpedDevice
		return ret
	}
	return *o.Devices
}

// GetDevicesOk returns a tuple with the Devices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DumpedCollection) GetDevicesOk() (*[]DumpedDevice, bool) {
	if o == nil || o.Devices == nil {
		return nil, false
	}
	return o.Devices, true
}

// HasDevices returns a boolean if a field has been set.
func (o *DumpedCollection) HasDevices() bool {
	if o != nil && o.Devices != nil {
		return true
	}

	return false
}

// SetDevices gets a reference to the given []DumpedDevice and assigns it to the Devices field.
func (o *DumpedCollection) SetDevices(v []DumpedDevice) {
	o.Devices = &v
}

// GetOutputs returns the Outputs field value if set, zero value otherwise.
func (o *DumpedCollection) GetOutputs() []Output {
	if o == nil || o.Outputs == nil {
		var ret []Output
		return ret
	}
	return *o.Outputs
}

// GetOutputsOk returns a tuple with the Outputs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DumpedCollection) GetOutputsOk() (*[]Output, bool) {
	if o == nil || o.Outputs == nil {
		return nil, false
	}
	return o.Outputs, true
}

// HasOutputs returns a boolean if a field has been set.
func (o *DumpedCollection) HasOutputs() bool {
	if o != nil && o.Outputs != nil {
		return true
	}

	return false
}

// SetOutputs gets a reference to the given []Output and assigns it to the Outputs field.
func (o *DumpedCollection) SetOutputs(v []Output) {
	o.Outputs = &v
}

func (o DumpedCollection) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Collection != nil {
		toSerialize["collection"] = o.Collection
	}
	if o.Devices != nil {
		toSerialize["devices"] = o.Devices
	}
	if o.Outputs != nil {
		toSerialize["outputs"] = o.Outputs
	}
	return json.Marshal(toSerialize)
}

type NullableDumpedCollection struct {
	value *DumpedCollection
	isSet bool
}

func (v NullableDumpedCollection) Get() *DumpedCollection {
	return v.value
}

func (v *NullableDumpedCollection) Set(val *DumpedCollection) {
	v.value = val
	v.isSet = true
}

func (v NullableDumpedCollection) IsSet() bool {
	return v.isSet
}

func (v *NullableDumpedCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDumpedCollection(val *DumpedCollection) *NullableDumpedCollection {
	return &NullableDumpedCollection{value: val, isSet: true}
}

func (v NullableDumpedCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDumpedCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
