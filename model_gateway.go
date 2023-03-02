/*
The Span API

API for device, collection, output and firmware management

API version: 4.4.2 nonviolent-adelbert
Contact: dev@lab5e.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package spanapi

import (
	"encoding/json"
)

// Gateway A gateway is a connection between devices and Span
type Gateway struct {
	GatewayId *string `json:"gatewayId,omitempty"`
	CollectionId *string `json:"collectionId,omitempty"`
	Name *string `json:"name,omitempty"`
	BuiltIn *bool `json:"builtIn,omitempty"`
	Type *GatewayType `json:"type,omitempty"`
	Config *GatewayConfig `json:"config,omitempty"`
	Tags *map[string]string `json:"tags,omitempty"`
	Status *GatewayStatus `json:"status,omitempty"`
}

// NewGateway instantiates a new Gateway object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGateway() *Gateway {
	this := Gateway{}
	var type_ GatewayType = GATEWAYTYPE_UNKNOWN
	this.Type = &type_
	var status GatewayStatus = GATEWAYSTATUS_UNKNOWN
	this.Status = &status
	return &this
}

// NewGatewayWithDefaults instantiates a new Gateway object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayWithDefaults() *Gateway {
	this := Gateway{}
	var type_ GatewayType = GATEWAYTYPE_UNKNOWN
	this.Type = &type_
	var status GatewayStatus = GATEWAYSTATUS_UNKNOWN
	this.Status = &status
	return &this
}

// GetGatewayId returns the GatewayId field value if set, zero value otherwise.
func (o *Gateway) GetGatewayId() string {
	if o == nil || o.GatewayId == nil {
		var ret string
		return ret
	}
	return *o.GatewayId
}

// GetGatewayIdOk returns a tuple with the GatewayId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Gateway) GetGatewayIdOk() (*string, bool) {
	if o == nil || o.GatewayId == nil {
		return nil, false
	}
	return o.GatewayId, true
}

// HasGatewayId returns a boolean if a field has been set.
func (o *Gateway) HasGatewayId() bool {
	if o != nil && o.GatewayId != nil {
		return true
	}

	return false
}

// SetGatewayId gets a reference to the given string and assigns it to the GatewayId field.
func (o *Gateway) SetGatewayId(v string) {
	o.GatewayId = &v
}

// GetCollectionId returns the CollectionId field value if set, zero value otherwise.
func (o *Gateway) GetCollectionId() string {
	if o == nil || o.CollectionId == nil {
		var ret string
		return ret
	}
	return *o.CollectionId
}

// GetCollectionIdOk returns a tuple with the CollectionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Gateway) GetCollectionIdOk() (*string, bool) {
	if o == nil || o.CollectionId == nil {
		return nil, false
	}
	return o.CollectionId, true
}

// HasCollectionId returns a boolean if a field has been set.
func (o *Gateway) HasCollectionId() bool {
	if o != nil && o.CollectionId != nil {
		return true
	}

	return false
}

// SetCollectionId gets a reference to the given string and assigns it to the CollectionId field.
func (o *Gateway) SetCollectionId(v string) {
	o.CollectionId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Gateway) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Gateway) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Gateway) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Gateway) SetName(v string) {
	o.Name = &v
}

// GetBuiltIn returns the BuiltIn field value if set, zero value otherwise.
func (o *Gateway) GetBuiltIn() bool {
	if o == nil || o.BuiltIn == nil {
		var ret bool
		return ret
	}
	return *o.BuiltIn
}

// GetBuiltInOk returns a tuple with the BuiltIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Gateway) GetBuiltInOk() (*bool, bool) {
	if o == nil || o.BuiltIn == nil {
		return nil, false
	}
	return o.BuiltIn, true
}

// HasBuiltIn returns a boolean if a field has been set.
func (o *Gateway) HasBuiltIn() bool {
	if o != nil && o.BuiltIn != nil {
		return true
	}

	return false
}

// SetBuiltIn gets a reference to the given bool and assigns it to the BuiltIn field.
func (o *Gateway) SetBuiltIn(v bool) {
	o.BuiltIn = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Gateway) GetType() GatewayType {
	if o == nil || o.Type == nil {
		var ret GatewayType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Gateway) GetTypeOk() (*GatewayType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Gateway) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given GatewayType and assigns it to the Type field.
func (o *Gateway) SetType(v GatewayType) {
	o.Type = &v
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *Gateway) GetConfig() GatewayConfig {
	if o == nil || o.Config == nil {
		var ret GatewayConfig
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Gateway) GetConfigOk() (*GatewayConfig, bool) {
	if o == nil || o.Config == nil {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *Gateway) HasConfig() bool {
	if o != nil && o.Config != nil {
		return true
	}

	return false
}

// SetConfig gets a reference to the given GatewayConfig and assigns it to the Config field.
func (o *Gateway) SetConfig(v GatewayConfig) {
	o.Config = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *Gateway) GetTags() map[string]string {
	if o == nil || o.Tags == nil {
		var ret map[string]string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Gateway) GetTagsOk() (*map[string]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *Gateway) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given map[string]string and assigns it to the Tags field.
func (o *Gateway) SetTags(v map[string]string) {
	o.Tags = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Gateway) GetStatus() GatewayStatus {
	if o == nil || o.Status == nil {
		var ret GatewayStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Gateway) GetStatusOk() (*GatewayStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Gateway) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given GatewayStatus and assigns it to the Status field.
func (o *Gateway) SetStatus(v GatewayStatus) {
	o.Status = &v
}

func (o Gateway) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.GatewayId != nil {
		toSerialize["gatewayId"] = o.GatewayId
	}
	if o.CollectionId != nil {
		toSerialize["collectionId"] = o.CollectionId
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.BuiltIn != nil {
		toSerialize["builtIn"] = o.BuiltIn
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Config != nil {
		toSerialize["config"] = o.Config
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableGateway struct {
	value *Gateway
	isSet bool
}

func (v NullableGateway) Get() *Gateway {
	return v.value
}

func (v *NullableGateway) Set(val *Gateway) {
	v.value = val
	v.isSet = true
}

func (v NullableGateway) IsSet() bool {
	return v.isSet
}

func (v *NullableGateway) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGateway(val *Gateway) *NullableGateway {
	return &NullableGateway{value: val, isSet: true}
}

func (v NullableGateway) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGateway) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


