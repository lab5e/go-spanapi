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

// checks if the GatewayCertificateResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GatewayCertificateResponse{}

// GatewayCertificateResponse struct for GatewayCertificateResponse
type GatewayCertificateResponse struct {
	Certificates []CertificateInfo `json:"certificates,omitempty"`
}

// NewGatewayCertificateResponse instantiates a new GatewayCertificateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayCertificateResponse() *GatewayCertificateResponse {
	this := GatewayCertificateResponse{}
	return &this
}

// NewGatewayCertificateResponseWithDefaults instantiates a new GatewayCertificateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayCertificateResponseWithDefaults() *GatewayCertificateResponse {
	this := GatewayCertificateResponse{}
	return &this
}

// GetCertificates returns the Certificates field value if set, zero value otherwise.
func (o *GatewayCertificateResponse) GetCertificates() []CertificateInfo {
	if o == nil || IsNil(o.Certificates) {
		var ret []CertificateInfo
		return ret
	}
	return o.Certificates
}

// GetCertificatesOk returns a tuple with the Certificates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCertificateResponse) GetCertificatesOk() ([]CertificateInfo, bool) {
	if o == nil || IsNil(o.Certificates) {
		return nil, false
	}
	return o.Certificates, true
}

// HasCertificates returns a boolean if a field has been set.
func (o *GatewayCertificateResponse) HasCertificates() bool {
	if o != nil && !IsNil(o.Certificates) {
		return true
	}

	return false
}

// SetCertificates gets a reference to the given []CertificateInfo and assigns it to the Certificates field.
func (o *GatewayCertificateResponse) SetCertificates(v []CertificateInfo) {
	o.Certificates = v
}

func (o GatewayCertificateResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GatewayCertificateResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Certificates) {
		toSerialize["certificates"] = o.Certificates
	}
	return toSerialize, nil
}

type NullableGatewayCertificateResponse struct {
	value *GatewayCertificateResponse
	isSet bool
}

func (v NullableGatewayCertificateResponse) Get() *GatewayCertificateResponse {
	return v.value
}

func (v *NullableGatewayCertificateResponse) Set(val *GatewayCertificateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayCertificateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayCertificateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayCertificateResponse(val *GatewayCertificateResponse) *NullableGatewayCertificateResponse {
	return &NullableGatewayCertificateResponse{value: val, isSet: true}
}

func (v NullableGatewayCertificateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayCertificateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


