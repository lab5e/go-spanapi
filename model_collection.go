/*
The Span API

API for device, collection, output and firmware management

API version: 4.2.4 curable-andres
Contact: dev@lab5e.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package spanapi

import (
	"encoding/json"
)

// Collection This is a collection
type Collection struct {
	// The ID of the collection. This is assigned by the backend.
	CollectionId *string `json:"collectionId,omitempty"`
	// The team ID that owns the collection. This field is required. When you create new collections the default is to use your private team ID.
	TeamId   *string             `json:"teamId,omitempty"`
	Firmware *CollectionFirmware `json:"firmware,omitempty"`
	// Tags for the collection. Tags are metadata fields that you can assign to the collection.
	Tags *map[string]string `json:"tags,omitempty"`
}

// NewCollection instantiates a new Collection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollection() *Collection {
	this := Collection{}
	return &this
}

// NewCollectionWithDefaults instantiates a new Collection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionWithDefaults() *Collection {
	this := Collection{}
	return &this
}

// GetCollectionId returns the CollectionId field value if set, zero value otherwise.
func (o *Collection) GetCollectionId() string {
	if o == nil || o.CollectionId == nil {
		var ret string
		return ret
	}
	return *o.CollectionId
}

// GetCollectionIdOk returns a tuple with the CollectionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Collection) GetCollectionIdOk() (*string, bool) {
	if o == nil || o.CollectionId == nil {
		return nil, false
	}
	return o.CollectionId, true
}

// HasCollectionId returns a boolean if a field has been set.
func (o *Collection) HasCollectionId() bool {
	if o != nil && o.CollectionId != nil {
		return true
	}

	return false
}

// SetCollectionId gets a reference to the given string and assigns it to the CollectionId field.
func (o *Collection) SetCollectionId(v string) {
	o.CollectionId = &v
}

// GetTeamId returns the TeamId field value if set, zero value otherwise.
func (o *Collection) GetTeamId() string {
	if o == nil || o.TeamId == nil {
		var ret string
		return ret
	}
	return *o.TeamId
}

// GetTeamIdOk returns a tuple with the TeamId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Collection) GetTeamIdOk() (*string, bool) {
	if o == nil || o.TeamId == nil {
		return nil, false
	}
	return o.TeamId, true
}

// HasTeamId returns a boolean if a field has been set.
func (o *Collection) HasTeamId() bool {
	if o != nil && o.TeamId != nil {
		return true
	}

	return false
}

// SetTeamId gets a reference to the given string and assigns it to the TeamId field.
func (o *Collection) SetTeamId(v string) {
	o.TeamId = &v
}

// GetFirmware returns the Firmware field value if set, zero value otherwise.
func (o *Collection) GetFirmware() CollectionFirmware {
	if o == nil || o.Firmware == nil {
		var ret CollectionFirmware
		return ret
	}
	return *o.Firmware
}

// GetFirmwareOk returns a tuple with the Firmware field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Collection) GetFirmwareOk() (*CollectionFirmware, bool) {
	if o == nil || o.Firmware == nil {
		return nil, false
	}
	return o.Firmware, true
}

// HasFirmware returns a boolean if a field has been set.
func (o *Collection) HasFirmware() bool {
	if o != nil && o.Firmware != nil {
		return true
	}

	return false
}

// SetFirmware gets a reference to the given CollectionFirmware and assigns it to the Firmware field.
func (o *Collection) SetFirmware(v CollectionFirmware) {
	o.Firmware = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *Collection) GetTags() map[string]string {
	if o == nil || o.Tags == nil {
		var ret map[string]string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Collection) GetTagsOk() (*map[string]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *Collection) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given map[string]string and assigns it to the Tags field.
func (o *Collection) SetTags(v map[string]string) {
	o.Tags = &v
}

func (o Collection) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CollectionId != nil {
		toSerialize["collectionId"] = o.CollectionId
	}
	if o.TeamId != nil {
		toSerialize["teamId"] = o.TeamId
	}
	if o.Firmware != nil {
		toSerialize["firmware"] = o.Firmware
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	return json.Marshal(toSerialize)
}

type NullableCollection struct {
	value *Collection
	isSet bool
}

func (v NullableCollection) Get() *Collection {
	return v.value
}

func (v *NullableCollection) Set(val *Collection) {
	v.value = val
	v.isSet = true
}

func (v NullableCollection) IsSet() bool {
	return v.isSet
}

func (v *NullableCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollection(val *Collection) *NullableCollection {
	return &NullableCollection{value: val, isSet: true}
}

func (v NullableCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
