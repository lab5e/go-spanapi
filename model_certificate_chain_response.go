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

// checks if the CertificateChainResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CertificateChainResponse{}

// CertificateChainResponse Response when retrieving a certificate chain
type CertificateChainResponse struct {
	Chain *string `json:"chain,omitempty"`
}

// NewCertificateChainResponse instantiates a new CertificateChainResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCertificateChainResponse() *CertificateChainResponse {
	this := CertificateChainResponse{}
	return &this
}

// NewCertificateChainResponseWithDefaults instantiates a new CertificateChainResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCertificateChainResponseWithDefaults() *CertificateChainResponse {
	this := CertificateChainResponse{}
	return &this
}

// GetChain returns the Chain field value if set, zero value otherwise.
func (o *CertificateChainResponse) GetChain() string {
	if o == nil || IsNil(o.Chain) {
		var ret string
		return ret
	}
	return *o.Chain
}

// GetChainOk returns a tuple with the Chain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateChainResponse) GetChainOk() (*string, bool) {
	if o == nil || IsNil(o.Chain) {
		return nil, false
	}
	return o.Chain, true
}

// HasChain returns a boolean if a field has been set.
func (o *CertificateChainResponse) HasChain() bool {
	if o != nil && !IsNil(o.Chain) {
		return true
	}

	return false
}

// SetChain gets a reference to the given string and assigns it to the Chain field.
func (o *CertificateChainResponse) SetChain(v string) {
	o.Chain = &v
}

func (o CertificateChainResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CertificateChainResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Chain) {
		toSerialize["chain"] = o.Chain
	}
	return toSerialize, nil
}

type NullableCertificateChainResponse struct {
	value *CertificateChainResponse
	isSet bool
}

func (v NullableCertificateChainResponse) Get() *CertificateChainResponse {
	return v.value
}

func (v *NullableCertificateChainResponse) Set(val *CertificateChainResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCertificateChainResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCertificateChainResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCertificateChainResponse(val *CertificateChainResponse) *NullableCertificateChainResponse {
	return &NullableCertificateChainResponse{value: val, isSet: true}
}

func (v NullableCertificateChainResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCertificateChainResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


