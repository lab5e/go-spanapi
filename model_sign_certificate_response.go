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

// SignCertificateResponse Response when signing a certificate
type SignCertificateResponse struct {
	Certificate *string `json:"certificate,omitempty"`
	Chain *string `json:"chain,omitempty"`
}

// NewSignCertificateResponse instantiates a new SignCertificateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignCertificateResponse() *SignCertificateResponse {
	this := SignCertificateResponse{}
	return &this
}

// NewSignCertificateResponseWithDefaults instantiates a new SignCertificateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignCertificateResponseWithDefaults() *SignCertificateResponse {
	this := SignCertificateResponse{}
	return &this
}

// GetCertificate returns the Certificate field value if set, zero value otherwise.
func (o *SignCertificateResponse) GetCertificate() string {
	if o == nil || o.Certificate == nil {
		var ret string
		return ret
	}
	return *o.Certificate
}

// GetCertificateOk returns a tuple with the Certificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignCertificateResponse) GetCertificateOk() (*string, bool) {
	if o == nil || o.Certificate == nil {
		return nil, false
	}
	return o.Certificate, true
}

// HasCertificate returns a boolean if a field has been set.
func (o *SignCertificateResponse) HasCertificate() bool {
	if o != nil && o.Certificate != nil {
		return true
	}

	return false
}

// SetCertificate gets a reference to the given string and assigns it to the Certificate field.
func (o *SignCertificateResponse) SetCertificate(v string) {
	o.Certificate = &v
}

// GetChain returns the Chain field value if set, zero value otherwise.
func (o *SignCertificateResponse) GetChain() string {
	if o == nil || o.Chain == nil {
		var ret string
		return ret
	}
	return *o.Chain
}

// GetChainOk returns a tuple with the Chain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignCertificateResponse) GetChainOk() (*string, bool) {
	if o == nil || o.Chain == nil {
		return nil, false
	}
	return o.Chain, true
}

// HasChain returns a boolean if a field has been set.
func (o *SignCertificateResponse) HasChain() bool {
	if o != nil && o.Chain != nil {
		return true
	}

	return false
}

// SetChain gets a reference to the given string and assigns it to the Chain field.
func (o *SignCertificateResponse) SetChain(v string) {
	o.Chain = &v
}

func (o SignCertificateResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Certificate != nil {
		toSerialize["certificate"] = o.Certificate
	}
	if o.Chain != nil {
		toSerialize["chain"] = o.Chain
	}
	return json.Marshal(toSerialize)
}

type NullableSignCertificateResponse struct {
	value *SignCertificateResponse
	isSet bool
}

func (v NullableSignCertificateResponse) Get() *SignCertificateResponse {
	return v.value
}

func (v *NullableSignCertificateResponse) Set(val *SignCertificateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSignCertificateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSignCertificateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignCertificateResponse(val *SignCertificateResponse) *NullableSignCertificateResponse {
	return &NullableSignCertificateResponse{value: val, isSet: true}
}

func (v NullableSignCertificateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignCertificateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


