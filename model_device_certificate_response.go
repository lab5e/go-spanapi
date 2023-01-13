/*
The Span API

API for device, collection, output and firmware management

API version: 4.3.0 grouchy-aloysius
Contact: dev@lab5e.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package spanapi

import (
	"encoding/json"
)

// DeviceCertificateResponse Response object for certificate info resource
type DeviceCertificateResponse struct {
	Certificates []CertificateInfo `json:"certificates,omitempty"`
}

// NewDeviceCertificateResponse instantiates a new DeviceCertificateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceCertificateResponse() *DeviceCertificateResponse {
	this := DeviceCertificateResponse{}
	return &this
}

// NewDeviceCertificateResponseWithDefaults instantiates a new DeviceCertificateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceCertificateResponseWithDefaults() *DeviceCertificateResponse {
	this := DeviceCertificateResponse{}
	return &this
}

// GetCertificates returns the Certificates field value if set, zero value otherwise.
func (o *DeviceCertificateResponse) GetCertificates() []CertificateInfo {
	if o == nil || o.Certificates == nil {
		var ret []CertificateInfo
		return ret
	}
	return o.Certificates
}

// GetCertificatesOk returns a tuple with the Certificates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceCertificateResponse) GetCertificatesOk() ([]CertificateInfo, bool) {
	if o == nil || o.Certificates == nil {
		return nil, false
	}
	return o.Certificates, true
}

// HasCertificates returns a boolean if a field has been set.
func (o *DeviceCertificateResponse) HasCertificates() bool {
	if o != nil && o.Certificates != nil {
		return true
	}

	return false
}

// SetCertificates gets a reference to the given []CertificateInfo and assigns it to the Certificates field.
func (o *DeviceCertificateResponse) SetCertificates(v []CertificateInfo) {
	o.Certificates = v
}

func (o DeviceCertificateResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Certificates != nil {
		toSerialize["certificates"] = o.Certificates
	}
	return json.Marshal(toSerialize)
}

type NullableDeviceCertificateResponse struct {
	value *DeviceCertificateResponse
	isSet bool
}

func (v NullableDeviceCertificateResponse) Get() *DeviceCertificateResponse {
	return v.value
}

func (v *NullableDeviceCertificateResponse) Set(val *DeviceCertificateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceCertificateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceCertificateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceCertificateResponse(val *DeviceCertificateResponse) *NullableDeviceCertificateResponse {
	return &NullableDeviceCertificateResponse{value: val, isSet: true}
}

func (v NullableDeviceCertificateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceCertificateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
