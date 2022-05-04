/*
The Span API

API for device, collection, output and firmware management

API version: 4.2.0 alternative-adolphus
Contact: dev@lab5e.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package spanapi

import (
	"encoding/json"
)

// CreateCertificateResponse Response when creating a new certificate
type CreateCertificateResponse struct {
	Certificate *string `json:"certificate,omitempty"`
	PrivateKey *string `json:"privateKey,omitempty"`
	Chain *string `json:"chain,omitempty"`
}

// NewCreateCertificateResponse instantiates a new CreateCertificateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateCertificateResponse() *CreateCertificateResponse {
	this := CreateCertificateResponse{}
	return &this
}

// NewCreateCertificateResponseWithDefaults instantiates a new CreateCertificateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateCertificateResponseWithDefaults() *CreateCertificateResponse {
	this := CreateCertificateResponse{}
	return &this
}

// GetCertificate returns the Certificate field value if set, zero value otherwise.
func (o *CreateCertificateResponse) GetCertificate() string {
	if o == nil || o.Certificate == nil {
		var ret string
		return ret
	}
	return *o.Certificate
}

// GetCertificateOk returns a tuple with the Certificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCertificateResponse) GetCertificateOk() (*string, bool) {
	if o == nil || o.Certificate == nil {
		return nil, false
	}
	return o.Certificate, true
}

// HasCertificate returns a boolean if a field has been set.
func (o *CreateCertificateResponse) HasCertificate() bool {
	if o != nil && o.Certificate != nil {
		return true
	}

	return false
}

// SetCertificate gets a reference to the given string and assigns it to the Certificate field.
func (o *CreateCertificateResponse) SetCertificate(v string) {
	o.Certificate = &v
}

// GetPrivateKey returns the PrivateKey field value if set, zero value otherwise.
func (o *CreateCertificateResponse) GetPrivateKey() string {
	if o == nil || o.PrivateKey == nil {
		var ret string
		return ret
	}
	return *o.PrivateKey
}

// GetPrivateKeyOk returns a tuple with the PrivateKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCertificateResponse) GetPrivateKeyOk() (*string, bool) {
	if o == nil || o.PrivateKey == nil {
		return nil, false
	}
	return o.PrivateKey, true
}

// HasPrivateKey returns a boolean if a field has been set.
func (o *CreateCertificateResponse) HasPrivateKey() bool {
	if o != nil && o.PrivateKey != nil {
		return true
	}

	return false
}

// SetPrivateKey gets a reference to the given string and assigns it to the PrivateKey field.
func (o *CreateCertificateResponse) SetPrivateKey(v string) {
	o.PrivateKey = &v
}

// GetChain returns the Chain field value if set, zero value otherwise.
func (o *CreateCertificateResponse) GetChain() string {
	if o == nil || o.Chain == nil {
		var ret string
		return ret
	}
	return *o.Chain
}

// GetChainOk returns a tuple with the Chain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCertificateResponse) GetChainOk() (*string, bool) {
	if o == nil || o.Chain == nil {
		return nil, false
	}
	return o.Chain, true
}

// HasChain returns a boolean if a field has been set.
func (o *CreateCertificateResponse) HasChain() bool {
	if o != nil && o.Chain != nil {
		return true
	}

	return false
}

// SetChain gets a reference to the given string and assigns it to the Chain field.
func (o *CreateCertificateResponse) SetChain(v string) {
	o.Chain = &v
}

func (o CreateCertificateResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Certificate != nil {
		toSerialize["certificate"] = o.Certificate
	}
	if o.PrivateKey != nil {
		toSerialize["privateKey"] = o.PrivateKey
	}
	if o.Chain != nil {
		toSerialize["chain"] = o.Chain
	}
	return json.Marshal(toSerialize)
}

type NullableCreateCertificateResponse struct {
	value *CreateCertificateResponse
	isSet bool
}

func (v NullableCreateCertificateResponse) Get() *CreateCertificateResponse {
	return v.value
}

func (v *NullableCreateCertificateResponse) Set(val *CreateCertificateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateCertificateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateCertificateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateCertificateResponse(val *CreateCertificateResponse) *NullableCreateCertificateResponse {
	return &NullableCreateCertificateResponse{value: val, isSet: true}
}

func (v NullableCreateCertificateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateCertificateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

