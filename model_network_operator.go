/*
The Span API

API for device, collection, output and firmware management

API version: 4.2.1 petulant-kyan
Contact: dev@lab5e.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package spanapi

import (
	"encoding/json"
)

// NetworkOperator Operator holds information on the network operator. There might be several operators involved; one operator is running the network your devices are connected to and the SIM card in your device belongs to a different operator.  Deprecated: Replaced by CellularIoTMetadata
type NetworkOperator struct {
	// The Mobile Country Code for the operator.
	Mcc *int32 `json:"mcc,omitempty"`
	Mnc *int32 `json:"mnc,omitempty"`
	Country *string `json:"country,omitempty"`
	Network *string `json:"network,omitempty"`
	CountryCode *string `json:"countryCode,omitempty"`
}

// NewNetworkOperator instantiates a new NetworkOperator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkOperator() *NetworkOperator {
	this := NetworkOperator{}
	return &this
}

// NewNetworkOperatorWithDefaults instantiates a new NetworkOperator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkOperatorWithDefaults() *NetworkOperator {
	this := NetworkOperator{}
	return &this
}

// GetMcc returns the Mcc field value if set, zero value otherwise.
func (o *NetworkOperator) GetMcc() int32 {
	if o == nil || o.Mcc == nil {
		var ret int32
		return ret
	}
	return *o.Mcc
}

// GetMccOk returns a tuple with the Mcc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkOperator) GetMccOk() (*int32, bool) {
	if o == nil || o.Mcc == nil {
		return nil, false
	}
	return o.Mcc, true
}

// HasMcc returns a boolean if a field has been set.
func (o *NetworkOperator) HasMcc() bool {
	if o != nil && o.Mcc != nil {
		return true
	}

	return false
}

// SetMcc gets a reference to the given int32 and assigns it to the Mcc field.
func (o *NetworkOperator) SetMcc(v int32) {
	o.Mcc = &v
}

// GetMnc returns the Mnc field value if set, zero value otherwise.
func (o *NetworkOperator) GetMnc() int32 {
	if o == nil || o.Mnc == nil {
		var ret int32
		return ret
	}
	return *o.Mnc
}

// GetMncOk returns a tuple with the Mnc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkOperator) GetMncOk() (*int32, bool) {
	if o == nil || o.Mnc == nil {
		return nil, false
	}
	return o.Mnc, true
}

// HasMnc returns a boolean if a field has been set.
func (o *NetworkOperator) HasMnc() bool {
	if o != nil && o.Mnc != nil {
		return true
	}

	return false
}

// SetMnc gets a reference to the given int32 and assigns it to the Mnc field.
func (o *NetworkOperator) SetMnc(v int32) {
	o.Mnc = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *NetworkOperator) GetCountry() string {
	if o == nil || o.Country == nil {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkOperator) GetCountryOk() (*string, bool) {
	if o == nil || o.Country == nil {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *NetworkOperator) HasCountry() bool {
	if o != nil && o.Country != nil {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *NetworkOperator) SetCountry(v string) {
	o.Country = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *NetworkOperator) GetNetwork() string {
	if o == nil || o.Network == nil {
		var ret string
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkOperator) GetNetworkOk() (*string, bool) {
	if o == nil || o.Network == nil {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *NetworkOperator) HasNetwork() bool {
	if o != nil && o.Network != nil {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given string and assigns it to the Network field.
func (o *NetworkOperator) SetNetwork(v string) {
	o.Network = &v
}

// GetCountryCode returns the CountryCode field value if set, zero value otherwise.
func (o *NetworkOperator) GetCountryCode() string {
	if o == nil || o.CountryCode == nil {
		var ret string
		return ret
	}
	return *o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkOperator) GetCountryCodeOk() (*string, bool) {
	if o == nil || o.CountryCode == nil {
		return nil, false
	}
	return o.CountryCode, true
}

// HasCountryCode returns a boolean if a field has been set.
func (o *NetworkOperator) HasCountryCode() bool {
	if o != nil && o.CountryCode != nil {
		return true
	}

	return false
}

// SetCountryCode gets a reference to the given string and assigns it to the CountryCode field.
func (o *NetworkOperator) SetCountryCode(v string) {
	o.CountryCode = &v
}

func (o NetworkOperator) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Mcc != nil {
		toSerialize["mcc"] = o.Mcc
	}
	if o.Mnc != nil {
		toSerialize["mnc"] = o.Mnc
	}
	if o.Country != nil {
		toSerialize["country"] = o.Country
	}
	if o.Network != nil {
		toSerialize["network"] = o.Network
	}
	if o.CountryCode != nil {
		toSerialize["countryCode"] = o.CountryCode
	}
	return json.Marshal(toSerialize)
}

type NullableNetworkOperator struct {
	value *NetworkOperator
	isSet bool
}

func (v NullableNetworkOperator) Get() *NetworkOperator {
	return v.value
}

func (v *NullableNetworkOperator) Set(val *NetworkOperator) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkOperator) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkOperator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkOperator(val *NetworkOperator) *NullableNetworkOperator {
	return &NullableNetworkOperator{value: val, isSet: true}
}

func (v NullableNetworkOperator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkOperator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


