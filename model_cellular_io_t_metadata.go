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

// checks if the CellularIoTMetadata type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CellularIoTMetadata{}

// CellularIoTMetadata This is the metadata for a Cellular IoT device connected via an APN.
type CellularIoTMetadata struct {
	GatewayId *string `json:"gatewayId,omitempty"`
	// Allocated IP address.
	AllocatedIp *string `json:"allocatedIp,omitempty"`
	AllocatedAt *string `json:"allocatedAt,omitempty"`
	CellId *string `json:"cellId,omitempty"`
	// the provider in use.  The Mobile Country Code for the operator.
	Mcc *int32 `json:"mcc,omitempty"`
	Mnc *int32 `json:"mnc,omitempty"`
	Country *string `json:"country,omitempty"`
	Network *string `json:"network,omitempty"`
	CountryCode *string `json:"countryCode,omitempty"`
	LastUpdate *string `json:"lastUpdate,omitempty"`
}

// NewCellularIoTMetadata instantiates a new CellularIoTMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCellularIoTMetadata() *CellularIoTMetadata {
	this := CellularIoTMetadata{}
	return &this
}

// NewCellularIoTMetadataWithDefaults instantiates a new CellularIoTMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCellularIoTMetadataWithDefaults() *CellularIoTMetadata {
	this := CellularIoTMetadata{}
	return &this
}

// GetGatewayId returns the GatewayId field value if set, zero value otherwise.
func (o *CellularIoTMetadata) GetGatewayId() string {
	if o == nil || IsNil(o.GatewayId) {
		var ret string
		return ret
	}
	return *o.GatewayId
}

// GetGatewayIdOk returns a tuple with the GatewayId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CellularIoTMetadata) GetGatewayIdOk() (*string, bool) {
	if o == nil || IsNil(o.GatewayId) {
		return nil, false
	}
	return o.GatewayId, true
}

// HasGatewayId returns a boolean if a field has been set.
func (o *CellularIoTMetadata) HasGatewayId() bool {
	if o != nil && !IsNil(o.GatewayId) {
		return true
	}

	return false
}

// SetGatewayId gets a reference to the given string and assigns it to the GatewayId field.
func (o *CellularIoTMetadata) SetGatewayId(v string) {
	o.GatewayId = &v
}

// GetAllocatedIp returns the AllocatedIp field value if set, zero value otherwise.
func (o *CellularIoTMetadata) GetAllocatedIp() string {
	if o == nil || IsNil(o.AllocatedIp) {
		var ret string
		return ret
	}
	return *o.AllocatedIp
}

// GetAllocatedIpOk returns a tuple with the AllocatedIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CellularIoTMetadata) GetAllocatedIpOk() (*string, bool) {
	if o == nil || IsNil(o.AllocatedIp) {
		return nil, false
	}
	return o.AllocatedIp, true
}

// HasAllocatedIp returns a boolean if a field has been set.
func (o *CellularIoTMetadata) HasAllocatedIp() bool {
	if o != nil && !IsNil(o.AllocatedIp) {
		return true
	}

	return false
}

// SetAllocatedIp gets a reference to the given string and assigns it to the AllocatedIp field.
func (o *CellularIoTMetadata) SetAllocatedIp(v string) {
	o.AllocatedIp = &v
}

// GetAllocatedAt returns the AllocatedAt field value if set, zero value otherwise.
func (o *CellularIoTMetadata) GetAllocatedAt() string {
	if o == nil || IsNil(o.AllocatedAt) {
		var ret string
		return ret
	}
	return *o.AllocatedAt
}

// GetAllocatedAtOk returns a tuple with the AllocatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CellularIoTMetadata) GetAllocatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.AllocatedAt) {
		return nil, false
	}
	return o.AllocatedAt, true
}

// HasAllocatedAt returns a boolean if a field has been set.
func (o *CellularIoTMetadata) HasAllocatedAt() bool {
	if o != nil && !IsNil(o.AllocatedAt) {
		return true
	}

	return false
}

// SetAllocatedAt gets a reference to the given string and assigns it to the AllocatedAt field.
func (o *CellularIoTMetadata) SetAllocatedAt(v string) {
	o.AllocatedAt = &v
}

// GetCellId returns the CellId field value if set, zero value otherwise.
func (o *CellularIoTMetadata) GetCellId() string {
	if o == nil || IsNil(o.CellId) {
		var ret string
		return ret
	}
	return *o.CellId
}

// GetCellIdOk returns a tuple with the CellId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CellularIoTMetadata) GetCellIdOk() (*string, bool) {
	if o == nil || IsNil(o.CellId) {
		return nil, false
	}
	return o.CellId, true
}

// HasCellId returns a boolean if a field has been set.
func (o *CellularIoTMetadata) HasCellId() bool {
	if o != nil && !IsNil(o.CellId) {
		return true
	}

	return false
}

// SetCellId gets a reference to the given string and assigns it to the CellId field.
func (o *CellularIoTMetadata) SetCellId(v string) {
	o.CellId = &v
}

// GetMcc returns the Mcc field value if set, zero value otherwise.
func (o *CellularIoTMetadata) GetMcc() int32 {
	if o == nil || IsNil(o.Mcc) {
		var ret int32
		return ret
	}
	return *o.Mcc
}

// GetMccOk returns a tuple with the Mcc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CellularIoTMetadata) GetMccOk() (*int32, bool) {
	if o == nil || IsNil(o.Mcc) {
		return nil, false
	}
	return o.Mcc, true
}

// HasMcc returns a boolean if a field has been set.
func (o *CellularIoTMetadata) HasMcc() bool {
	if o != nil && !IsNil(o.Mcc) {
		return true
	}

	return false
}

// SetMcc gets a reference to the given int32 and assigns it to the Mcc field.
func (o *CellularIoTMetadata) SetMcc(v int32) {
	o.Mcc = &v
}

// GetMnc returns the Mnc field value if set, zero value otherwise.
func (o *CellularIoTMetadata) GetMnc() int32 {
	if o == nil || IsNil(o.Mnc) {
		var ret int32
		return ret
	}
	return *o.Mnc
}

// GetMncOk returns a tuple with the Mnc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CellularIoTMetadata) GetMncOk() (*int32, bool) {
	if o == nil || IsNil(o.Mnc) {
		return nil, false
	}
	return o.Mnc, true
}

// HasMnc returns a boolean if a field has been set.
func (o *CellularIoTMetadata) HasMnc() bool {
	if o != nil && !IsNil(o.Mnc) {
		return true
	}

	return false
}

// SetMnc gets a reference to the given int32 and assigns it to the Mnc field.
func (o *CellularIoTMetadata) SetMnc(v int32) {
	o.Mnc = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *CellularIoTMetadata) GetCountry() string {
	if o == nil || IsNil(o.Country) {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CellularIoTMetadata) GetCountryOk() (*string, bool) {
	if o == nil || IsNil(o.Country) {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *CellularIoTMetadata) HasCountry() bool {
	if o != nil && !IsNil(o.Country) {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *CellularIoTMetadata) SetCountry(v string) {
	o.Country = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *CellularIoTMetadata) GetNetwork() string {
	if o == nil || IsNil(o.Network) {
		var ret string
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CellularIoTMetadata) GetNetworkOk() (*string, bool) {
	if o == nil || IsNil(o.Network) {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *CellularIoTMetadata) HasNetwork() bool {
	if o != nil && !IsNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given string and assigns it to the Network field.
func (o *CellularIoTMetadata) SetNetwork(v string) {
	o.Network = &v
}

// GetCountryCode returns the CountryCode field value if set, zero value otherwise.
func (o *CellularIoTMetadata) GetCountryCode() string {
	if o == nil || IsNil(o.CountryCode) {
		var ret string
		return ret
	}
	return *o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CellularIoTMetadata) GetCountryCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CountryCode) {
		return nil, false
	}
	return o.CountryCode, true
}

// HasCountryCode returns a boolean if a field has been set.
func (o *CellularIoTMetadata) HasCountryCode() bool {
	if o != nil && !IsNil(o.CountryCode) {
		return true
	}

	return false
}

// SetCountryCode gets a reference to the given string and assigns it to the CountryCode field.
func (o *CellularIoTMetadata) SetCountryCode(v string) {
	o.CountryCode = &v
}

// GetLastUpdate returns the LastUpdate field value if set, zero value otherwise.
func (o *CellularIoTMetadata) GetLastUpdate() string {
	if o == nil || IsNil(o.LastUpdate) {
		var ret string
		return ret
	}
	return *o.LastUpdate
}

// GetLastUpdateOk returns a tuple with the LastUpdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CellularIoTMetadata) GetLastUpdateOk() (*string, bool) {
	if o == nil || IsNil(o.LastUpdate) {
		return nil, false
	}
	return o.LastUpdate, true
}

// HasLastUpdate returns a boolean if a field has been set.
func (o *CellularIoTMetadata) HasLastUpdate() bool {
	if o != nil && !IsNil(o.LastUpdate) {
		return true
	}

	return false
}

// SetLastUpdate gets a reference to the given string and assigns it to the LastUpdate field.
func (o *CellularIoTMetadata) SetLastUpdate(v string) {
	o.LastUpdate = &v
}

func (o CellularIoTMetadata) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CellularIoTMetadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GatewayId) {
		toSerialize["gatewayId"] = o.GatewayId
	}
	if !IsNil(o.AllocatedIp) {
		toSerialize["allocatedIp"] = o.AllocatedIp
	}
	if !IsNil(o.AllocatedAt) {
		toSerialize["allocatedAt"] = o.AllocatedAt
	}
	if !IsNil(o.CellId) {
		toSerialize["cellId"] = o.CellId
	}
	if !IsNil(o.Mcc) {
		toSerialize["mcc"] = o.Mcc
	}
	if !IsNil(o.Mnc) {
		toSerialize["mnc"] = o.Mnc
	}
	if !IsNil(o.Country) {
		toSerialize["country"] = o.Country
	}
	if !IsNil(o.Network) {
		toSerialize["network"] = o.Network
	}
	if !IsNil(o.CountryCode) {
		toSerialize["countryCode"] = o.CountryCode
	}
	if !IsNil(o.LastUpdate) {
		toSerialize["lastUpdate"] = o.LastUpdate
	}
	return toSerialize, nil
}

type NullableCellularIoTMetadata struct {
	value *CellularIoTMetadata
	isSet bool
}

func (v NullableCellularIoTMetadata) Get() *CellularIoTMetadata {
	return v.value
}

func (v *NullableCellularIoTMetadata) Set(val *CellularIoTMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableCellularIoTMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableCellularIoTMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCellularIoTMetadata(val *CellularIoTMetadata) *NullableCellularIoTMetadata {
	return &NullableCellularIoTMetadata{value: val, isSet: true}
}

func (v NullableCellularIoTMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCellularIoTMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


