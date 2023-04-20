/*
The Span API

API for device, collection, output and firmware management

API version: 4.4.2 larger-lashanda
Contact: dev@lab5e.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package spanapi

import (
	"encoding/json"
)

// checks if the Collection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Collection{}

// Collection This is a collection
type Collection struct {
	// The ID of the collection. This is assigned by the backend.
	CollectionId *string `json:"collectionId,omitempty"`
	// The team ID that owns the collection. This field is required. When you create new collections the default is to use your private team ID.
	TeamId *string `json:"teamId,omitempty"`
	Firmware *CollectionFirmware `json:"firmware,omitempty"`
	// Tags for the collection. Tags are metadata fields that you can assign to the collection.
	Tags *map[string]string `json:"tags,omitempty"`
	UpstreamTimestamps []string `json:"upstreamTimestamps,omitempty"`
	DownstreamTimestamps []string `json:"downstreamTimestamps,omitempty"`
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
	if o == nil || IsNil(o.CollectionId) {
		var ret string
		return ret
	}
	return *o.CollectionId
}

// GetCollectionIdOk returns a tuple with the CollectionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Collection) GetCollectionIdOk() (*string, bool) {
	if o == nil || IsNil(o.CollectionId) {
		return nil, false
	}
	return o.CollectionId, true
}

// HasCollectionId returns a boolean if a field has been set.
func (o *Collection) HasCollectionId() bool {
	if o != nil && !IsNil(o.CollectionId) {
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
	if o == nil || IsNil(o.TeamId) {
		var ret string
		return ret
	}
	return *o.TeamId
}

// GetTeamIdOk returns a tuple with the TeamId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Collection) GetTeamIdOk() (*string, bool) {
	if o == nil || IsNil(o.TeamId) {
		return nil, false
	}
	return o.TeamId, true
}

// HasTeamId returns a boolean if a field has been set.
func (o *Collection) HasTeamId() bool {
	if o != nil && !IsNil(o.TeamId) {
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
	if o == nil || IsNil(o.Firmware) {
		var ret CollectionFirmware
		return ret
	}
	return *o.Firmware
}

// GetFirmwareOk returns a tuple with the Firmware field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Collection) GetFirmwareOk() (*CollectionFirmware, bool) {
	if o == nil || IsNil(o.Firmware) {
		return nil, false
	}
	return o.Firmware, true
}

// HasFirmware returns a boolean if a field has been set.
func (o *Collection) HasFirmware() bool {
	if o != nil && !IsNil(o.Firmware) {
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
	if o == nil || IsNil(o.Tags) {
		var ret map[string]string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Collection) GetTagsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *Collection) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given map[string]string and assigns it to the Tags field.
func (o *Collection) SetTags(v map[string]string) {
	o.Tags = &v
}

// GetUpstreamTimestamps returns the UpstreamTimestamps field value if set, zero value otherwise.
func (o *Collection) GetUpstreamTimestamps() []string {
	if o == nil || IsNil(o.UpstreamTimestamps) {
		var ret []string
		return ret
	}
	return o.UpstreamTimestamps
}

// GetUpstreamTimestampsOk returns a tuple with the UpstreamTimestamps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Collection) GetUpstreamTimestampsOk() ([]string, bool) {
	if o == nil || IsNil(o.UpstreamTimestamps) {
		return nil, false
	}
	return o.UpstreamTimestamps, true
}

// HasUpstreamTimestamps returns a boolean if a field has been set.
func (o *Collection) HasUpstreamTimestamps() bool {
	if o != nil && !IsNil(o.UpstreamTimestamps) {
		return true
	}

	return false
}

// SetUpstreamTimestamps gets a reference to the given []string and assigns it to the UpstreamTimestamps field.
func (o *Collection) SetUpstreamTimestamps(v []string) {
	o.UpstreamTimestamps = v
}

// GetDownstreamTimestamps returns the DownstreamTimestamps field value if set, zero value otherwise.
func (o *Collection) GetDownstreamTimestamps() []string {
	if o == nil || IsNil(o.DownstreamTimestamps) {
		var ret []string
		return ret
	}
	return o.DownstreamTimestamps
}

// GetDownstreamTimestampsOk returns a tuple with the DownstreamTimestamps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Collection) GetDownstreamTimestampsOk() ([]string, bool) {
	if o == nil || IsNil(o.DownstreamTimestamps) {
		return nil, false
	}
	return o.DownstreamTimestamps, true
}

// HasDownstreamTimestamps returns a boolean if a field has been set.
func (o *Collection) HasDownstreamTimestamps() bool {
	if o != nil && !IsNil(o.DownstreamTimestamps) {
		return true
	}

	return false
}

// SetDownstreamTimestamps gets a reference to the given []string and assigns it to the DownstreamTimestamps field.
func (o *Collection) SetDownstreamTimestamps(v []string) {
	o.DownstreamTimestamps = v
}

func (o Collection) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Collection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CollectionId) {
		toSerialize["collectionId"] = o.CollectionId
	}
	if !IsNil(o.TeamId) {
		toSerialize["teamId"] = o.TeamId
	}
	if !IsNil(o.Firmware) {
		toSerialize["firmware"] = o.Firmware
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.UpstreamTimestamps) {
		toSerialize["upstreamTimestamps"] = o.UpstreamTimestamps
	}
	if !IsNil(o.DownstreamTimestamps) {
		toSerialize["downstreamTimestamps"] = o.DownstreamTimestamps
	}
	return toSerialize, nil
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


