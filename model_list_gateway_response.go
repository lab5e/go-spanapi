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

// checks if the ListGatewayResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListGatewayResponse{}

// ListGatewayResponse Response when listing gateways
type ListGatewayResponse struct {
	Gateways []Gateway `json:"gateways,omitempty"`
}

// NewListGatewayResponse instantiates a new ListGatewayResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListGatewayResponse() *ListGatewayResponse {
	this := ListGatewayResponse{}
	return &this
}

// NewListGatewayResponseWithDefaults instantiates a new ListGatewayResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListGatewayResponseWithDefaults() *ListGatewayResponse {
	this := ListGatewayResponse{}
	return &this
}

// GetGateways returns the Gateways field value if set, zero value otherwise.
func (o *ListGatewayResponse) GetGateways() []Gateway {
	if o == nil || IsNil(o.Gateways) {
		var ret []Gateway
		return ret
	}
	return o.Gateways
}

// GetGatewaysOk returns a tuple with the Gateways field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListGatewayResponse) GetGatewaysOk() ([]Gateway, bool) {
	if o == nil || IsNil(o.Gateways) {
		return nil, false
	}
	return o.Gateways, true
}

// HasGateways returns a boolean if a field has been set.
func (o *ListGatewayResponse) HasGateways() bool {
	if o != nil && !IsNil(o.Gateways) {
		return true
	}

	return false
}

// SetGateways gets a reference to the given []Gateway and assigns it to the Gateways field.
func (o *ListGatewayResponse) SetGateways(v []Gateway) {
	o.Gateways = v
}

func (o ListGatewayResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListGatewayResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Gateways) {
		toSerialize["gateways"] = o.Gateways
	}
	return toSerialize, nil
}

type NullableListGatewayResponse struct {
	value *ListGatewayResponse
	isSet bool
}

func (v NullableListGatewayResponse) Get() *ListGatewayResponse {
	return v.value
}

func (v *NullableListGatewayResponse) Set(val *ListGatewayResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListGatewayResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListGatewayResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListGatewayResponse(val *ListGatewayResponse) *NullableListGatewayResponse {
	return &NullableListGatewayResponse{value: val, isSet: true}
}

func (v NullableListGatewayResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListGatewayResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


