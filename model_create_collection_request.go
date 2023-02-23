/*
The Span API

API for device, collection, output and firmware management

API version: 4.4.1 busy-janay
Contact: dev@lab5e.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package spanapi

import (
	"encoding/json"
)

// CreateCollectionRequest Request object when creating a collection. The collect ID is assigned by the service.
type CreateCollectionRequest struct {
	// The team ID that owns the collection. This field is required. When you create new collections the default is to use your private team ID.
	TeamId   *string             `json:"teamId,omitempty"`
	Firmware *CollectionFirmware `json:"firmware,omitempty"`
	// Tags for the collection. Tags are metadata fields that you can assign to the collection.
	Tags *map[string]string `json:"tags,omitempty"`
}

// NewCreateCollectionRequest instantiates a new CreateCollectionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateCollectionRequest() *CreateCollectionRequest {
	this := CreateCollectionRequest{}
	return &this
}

// NewCreateCollectionRequestWithDefaults instantiates a new CreateCollectionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateCollectionRequestWithDefaults() *CreateCollectionRequest {
	this := CreateCollectionRequest{}
	return &this
}

// GetTeamId returns the TeamId field value if set, zero value otherwise.
func (o *CreateCollectionRequest) GetTeamId() string {
	if o == nil || o.TeamId == nil {
		var ret string
		return ret
	}
	return *o.TeamId
}

// GetTeamIdOk returns a tuple with the TeamId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCollectionRequest) GetTeamIdOk() (*string, bool) {
	if o == nil || o.TeamId == nil {
		return nil, false
	}
	return o.TeamId, true
}

// HasTeamId returns a boolean if a field has been set.
func (o *CreateCollectionRequest) HasTeamId() bool {
	if o != nil && o.TeamId != nil {
		return true
	}

	return false
}

// SetTeamId gets a reference to the given string and assigns it to the TeamId field.
func (o *CreateCollectionRequest) SetTeamId(v string) {
	o.TeamId = &v
}

// GetFirmware returns the Firmware field value if set, zero value otherwise.
func (o *CreateCollectionRequest) GetFirmware() CollectionFirmware {
	if o == nil || o.Firmware == nil {
		var ret CollectionFirmware
		return ret
	}
	return *o.Firmware
}

// GetFirmwareOk returns a tuple with the Firmware field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCollectionRequest) GetFirmwareOk() (*CollectionFirmware, bool) {
	if o == nil || o.Firmware == nil {
		return nil, false
	}
	return o.Firmware, true
}

// HasFirmware returns a boolean if a field has been set.
func (o *CreateCollectionRequest) HasFirmware() bool {
	if o != nil && o.Firmware != nil {
		return true
	}

	return false
}

// SetFirmware gets a reference to the given CollectionFirmware and assigns it to the Firmware field.
func (o *CreateCollectionRequest) SetFirmware(v CollectionFirmware) {
	o.Firmware = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *CreateCollectionRequest) GetTags() map[string]string {
	if o == nil || o.Tags == nil {
		var ret map[string]string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCollectionRequest) GetTagsOk() (*map[string]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *CreateCollectionRequest) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given map[string]string and assigns it to the Tags field.
func (o *CreateCollectionRequest) SetTags(v map[string]string) {
	o.Tags = &v
}

func (o CreateCollectionRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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

type NullableCreateCollectionRequest struct {
	value *CreateCollectionRequest
	isSet bool
}

func (v NullableCreateCollectionRequest) Get() *CreateCollectionRequest {
	return v.value
}

func (v *NullableCreateCollectionRequest) Set(val *CreateCollectionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateCollectionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateCollectionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateCollectionRequest(val *CreateCollectionRequest) *NullableCreateCollectionRequest {
	return &NullableCreateCollectionRequest{value: val, isSet: true}
}

func (v NullableCreateCollectionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateCollectionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
