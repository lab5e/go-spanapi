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

// checks if the CreateGatewayRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateGatewayRequest{}

// CreateGatewayRequest struct for CreateGatewayRequest
type CreateGatewayRequest struct {
	Name *string `json:"name,omitempty"`
	Type *GatewayType `json:"type,omitempty"`
	Config *GatewayConfig `json:"config,omitempty"`
	Tags *map[string]string `json:"tags,omitempty"`
}

// NewCreateGatewayRequest instantiates a new CreateGatewayRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateGatewayRequest() *CreateGatewayRequest {
	this := CreateGatewayRequest{}
	var type_ GatewayType = GATEWAYTYPE_UNKNOWN
	this.Type = &type_
	return &this
}

// NewCreateGatewayRequestWithDefaults instantiates a new CreateGatewayRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateGatewayRequestWithDefaults() *CreateGatewayRequest {
	this := CreateGatewayRequest{}
	var type_ GatewayType = GATEWAYTYPE_UNKNOWN
	this.Type = &type_
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CreateGatewayRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGatewayRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CreateGatewayRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CreateGatewayRequest) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CreateGatewayRequest) GetType() GatewayType {
	if o == nil || IsNil(o.Type) {
		var ret GatewayType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGatewayRequest) GetTypeOk() (*GatewayType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CreateGatewayRequest) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given GatewayType and assigns it to the Type field.
func (o *CreateGatewayRequest) SetType(v GatewayType) {
	o.Type = &v
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *CreateGatewayRequest) GetConfig() GatewayConfig {
	if o == nil || IsNil(o.Config) {
		var ret GatewayConfig
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGatewayRequest) GetConfigOk() (*GatewayConfig, bool) {
	if o == nil || IsNil(o.Config) {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *CreateGatewayRequest) HasConfig() bool {
	if o != nil && !IsNil(o.Config) {
		return true
	}

	return false
}

// SetConfig gets a reference to the given GatewayConfig and assigns it to the Config field.
func (o *CreateGatewayRequest) SetConfig(v GatewayConfig) {
	o.Config = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *CreateGatewayRequest) GetTags() map[string]string {
	if o == nil || IsNil(o.Tags) {
		var ret map[string]string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGatewayRequest) GetTagsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *CreateGatewayRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given map[string]string and assigns it to the Tags field.
func (o *CreateGatewayRequest) SetTags(v map[string]string) {
	o.Tags = &v
}

func (o CreateGatewayRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateGatewayRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Config) {
		toSerialize["config"] = o.Config
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	return toSerialize, nil
}

type NullableCreateGatewayRequest struct {
	value *CreateGatewayRequest
	isSet bool
}

func (v NullableCreateGatewayRequest) Get() *CreateGatewayRequest {
	return v.value
}

func (v *NullableCreateGatewayRequest) Set(val *CreateGatewayRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateGatewayRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateGatewayRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateGatewayRequest(val *CreateGatewayRequest) *NullableCreateGatewayRequest {
	return &NullableCreateGatewayRequest{value: val, isSet: true}
}

func (v NullableCreateGatewayRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateGatewayRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


