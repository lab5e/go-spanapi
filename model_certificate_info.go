/*
The Span API

API for device, collection, output and firmware management

API version: 4.4.0 lean-joline
Contact: dev@lab5e.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package spanapi

import (
	"encoding/json"
)

// CertificateInfo Certificate information
type CertificateInfo struct {
	CertificateSerial *string `json:"certificateSerial,omitempty"`
	Expires           *string `json:"expires,omitempty"`
}

// NewCertificateInfo instantiates a new CertificateInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCertificateInfo() *CertificateInfo {
	this := CertificateInfo{}
	return &this
}

// NewCertificateInfoWithDefaults instantiates a new CertificateInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCertificateInfoWithDefaults() *CertificateInfo {
	this := CertificateInfo{}
	return &this
}

// GetCertificateSerial returns the CertificateSerial field value if set, zero value otherwise.
func (o *CertificateInfo) GetCertificateSerial() string {
	if o == nil || o.CertificateSerial == nil {
		var ret string
		return ret
	}
	return *o.CertificateSerial
}

// GetCertificateSerialOk returns a tuple with the CertificateSerial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateInfo) GetCertificateSerialOk() (*string, bool) {
	if o == nil || o.CertificateSerial == nil {
		return nil, false
	}
	return o.CertificateSerial, true
}

// HasCertificateSerial returns a boolean if a field has been set.
func (o *CertificateInfo) HasCertificateSerial() bool {
	if o != nil && o.CertificateSerial != nil {
		return true
	}

	return false
}

// SetCertificateSerial gets a reference to the given string and assigns it to the CertificateSerial field.
func (o *CertificateInfo) SetCertificateSerial(v string) {
	o.CertificateSerial = &v
}

// GetExpires returns the Expires field value if set, zero value otherwise.
func (o *CertificateInfo) GetExpires() string {
	if o == nil || o.Expires == nil {
		var ret string
		return ret
	}
	return *o.Expires
}

// GetExpiresOk returns a tuple with the Expires field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateInfo) GetExpiresOk() (*string, bool) {
	if o == nil || o.Expires == nil {
		return nil, false
	}
	return o.Expires, true
}

// HasExpires returns a boolean if a field has been set.
func (o *CertificateInfo) HasExpires() bool {
	if o != nil && o.Expires != nil {
		return true
	}

	return false
}

// SetExpires gets a reference to the given string and assigns it to the Expires field.
func (o *CertificateInfo) SetExpires(v string) {
	o.Expires = &v
}

func (o CertificateInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CertificateSerial != nil {
		toSerialize["certificateSerial"] = o.CertificateSerial
	}
	if o.Expires != nil {
		toSerialize["expires"] = o.Expires
	}
	return json.Marshal(toSerialize)
}

type NullableCertificateInfo struct {
	value *CertificateInfo
	isSet bool
}

func (v NullableCertificateInfo) Get() *CertificateInfo {
	return v.value
}

func (v *NullableCertificateInfo) Set(val *CertificateInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableCertificateInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableCertificateInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCertificateInfo(val *CertificateInfo) *NullableCertificateInfo {
	return &NullableCertificateInfo{value: val, isSet: true}
}

func (v NullableCertificateInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCertificateInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
