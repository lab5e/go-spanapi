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

// FieldMask struct for FieldMask
type FieldMask struct {
	Imsi *bool `json:"imsi,omitempty"`
	Imei *bool `json:"imei,omitempty"`
	Msisdn *bool `json:"msisdn,omitempty"`
	Location *bool `json:"location,omitempty"`
}

// NewFieldMask instantiates a new FieldMask object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFieldMask() *FieldMask {
	this := FieldMask{}
	return &this
}

// NewFieldMaskWithDefaults instantiates a new FieldMask object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFieldMaskWithDefaults() *FieldMask {
	this := FieldMask{}
	return &this
}

// GetImsi returns the Imsi field value if set, zero value otherwise.
func (o *FieldMask) GetImsi() bool {
	if o == nil || o.Imsi == nil {
		var ret bool
		return ret
	}
	return *o.Imsi
}

// GetImsiOk returns a tuple with the Imsi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldMask) GetImsiOk() (*bool, bool) {
	if o == nil || o.Imsi == nil {
		return nil, false
	}
	return o.Imsi, true
}

// HasImsi returns a boolean if a field has been set.
func (o *FieldMask) HasImsi() bool {
	if o != nil && o.Imsi != nil {
		return true
	}

	return false
}

// SetImsi gets a reference to the given bool and assigns it to the Imsi field.
func (o *FieldMask) SetImsi(v bool) {
	o.Imsi = &v
}

// GetImei returns the Imei field value if set, zero value otherwise.
func (o *FieldMask) GetImei() bool {
	if o == nil || o.Imei == nil {
		var ret bool
		return ret
	}
	return *o.Imei
}

// GetImeiOk returns a tuple with the Imei field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldMask) GetImeiOk() (*bool, bool) {
	if o == nil || o.Imei == nil {
		return nil, false
	}
	return o.Imei, true
}

// HasImei returns a boolean if a field has been set.
func (o *FieldMask) HasImei() bool {
	if o != nil && o.Imei != nil {
		return true
	}

	return false
}

// SetImei gets a reference to the given bool and assigns it to the Imei field.
func (o *FieldMask) SetImei(v bool) {
	o.Imei = &v
}

// GetMsisdn returns the Msisdn field value if set, zero value otherwise.
func (o *FieldMask) GetMsisdn() bool {
	if o == nil || o.Msisdn == nil {
		var ret bool
		return ret
	}
	return *o.Msisdn
}

// GetMsisdnOk returns a tuple with the Msisdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldMask) GetMsisdnOk() (*bool, bool) {
	if o == nil || o.Msisdn == nil {
		return nil, false
	}
	return o.Msisdn, true
}

// HasMsisdn returns a boolean if a field has been set.
func (o *FieldMask) HasMsisdn() bool {
	if o != nil && o.Msisdn != nil {
		return true
	}

	return false
}

// SetMsisdn gets a reference to the given bool and assigns it to the Msisdn field.
func (o *FieldMask) SetMsisdn(v bool) {
	o.Msisdn = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *FieldMask) GetLocation() bool {
	if o == nil || o.Location == nil {
		var ret bool
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldMask) GetLocationOk() (*bool, bool) {
	if o == nil || o.Location == nil {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *FieldMask) HasLocation() bool {
	if o != nil && o.Location != nil {
		return true
	}

	return false
}

// SetLocation gets a reference to the given bool and assigns it to the Location field.
func (o *FieldMask) SetLocation(v bool) {
	o.Location = &v
}

func (o FieldMask) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Imsi != nil {
		toSerialize["imsi"] = o.Imsi
	}
	if o.Imei != nil {
		toSerialize["imei"] = o.Imei
	}
	if o.Msisdn != nil {
		toSerialize["msisdn"] = o.Msisdn
	}
	if o.Location != nil {
		toSerialize["location"] = o.Location
	}
	return json.Marshal(toSerialize)
}

type NullableFieldMask struct {
	value *FieldMask
	isSet bool
}

func (v NullableFieldMask) Get() *FieldMask {
	return v.value
}

func (v *NullableFieldMask) Set(val *FieldMask) {
	v.value = val
	v.isSet = true
}

func (v NullableFieldMask) IsSet() bool {
	return v.isSet
}

func (v *NullableFieldMask) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFieldMask(val *FieldMask) *NullableFieldMask {
	return &NullableFieldMask{value: val, isSet: true}
}

func (v NullableFieldMask) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFieldMask) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


