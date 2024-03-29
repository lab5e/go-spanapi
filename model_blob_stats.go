/*
The Span API

API for device, collection, output and firmware management

API version: 4.8.0 indecipherable-ferman
Contact: dev@lab5e.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package spanapi

import (
	"encoding/json"
)

// checks if the BlobStats type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BlobStats{}

// BlobStats Statistics for a single blob
type BlobStats struct {
	BlobBytesMb *float32 `json:"blobBytesMb,omitempty"`
}

// NewBlobStats instantiates a new BlobStats object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlobStats() *BlobStats {
	this := BlobStats{}
	return &this
}

// NewBlobStatsWithDefaults instantiates a new BlobStats object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlobStatsWithDefaults() *BlobStats {
	this := BlobStats{}
	return &this
}

// GetBlobBytesMb returns the BlobBytesMb field value if set, zero value otherwise.
func (o *BlobStats) GetBlobBytesMb() float32 {
	if o == nil || IsNil(o.BlobBytesMb) {
		var ret float32
		return ret
	}
	return *o.BlobBytesMb
}

// GetBlobBytesMbOk returns a tuple with the BlobBytesMb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlobStats) GetBlobBytesMbOk() (*float32, bool) {
	if o == nil || IsNil(o.BlobBytesMb) {
		return nil, false
	}
	return o.BlobBytesMb, true
}

// HasBlobBytesMb returns a boolean if a field has been set.
func (o *BlobStats) HasBlobBytesMb() bool {
	if o != nil && !IsNil(o.BlobBytesMb) {
		return true
	}

	return false
}

// SetBlobBytesMb gets a reference to the given float32 and assigns it to the BlobBytesMb field.
func (o *BlobStats) SetBlobBytesMb(v float32) {
	o.BlobBytesMb = &v
}

func (o BlobStats) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BlobStats) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BlobBytesMb) {
		toSerialize["blobBytesMb"] = o.BlobBytesMb
	}
	return toSerialize, nil
}

type NullableBlobStats struct {
	value *BlobStats
	isSet bool
}

func (v NullableBlobStats) Get() *BlobStats {
	return v.value
}

func (v *NullableBlobStats) Set(val *BlobStats) {
	v.value = val
	v.isSet = true
}

func (v NullableBlobStats) IsSet() bool {
	return v.isSet
}

func (v *NullableBlobStats) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlobStats(val *BlobStats) *NullableBlobStats {
	return &NullableBlobStats{value: val, isSet: true}
}

func (v NullableBlobStats) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlobStats) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


