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

// VerifyCertificateResponse Response when verifying a certificate. The valid flag is set to true when the certificate is valid.  Any errors will be added to the errors array.
type VerifyCertificateResponse struct {
	Valid  *bool    `json:"valid,omitempty"`
	Errors []string `json:"errors,omitempty"`
}

// NewVerifyCertificateResponse instantiates a new VerifyCertificateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVerifyCertificateResponse() *VerifyCertificateResponse {
	this := VerifyCertificateResponse{}
	return &this
}

// NewVerifyCertificateResponseWithDefaults instantiates a new VerifyCertificateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVerifyCertificateResponseWithDefaults() *VerifyCertificateResponse {
	this := VerifyCertificateResponse{}
	return &this
}

// GetValid returns the Valid field value if set, zero value otherwise.
func (o *VerifyCertificateResponse) GetValid() bool {
	if o == nil || o.Valid == nil {
		var ret bool
		return ret
	}
	return *o.Valid
}

// GetValidOk returns a tuple with the Valid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyCertificateResponse) GetValidOk() (*bool, bool) {
	if o == nil || o.Valid == nil {
		return nil, false
	}
	return o.Valid, true
}

// HasValid returns a boolean if a field has been set.
func (o *VerifyCertificateResponse) HasValid() bool {
	if o != nil && o.Valid != nil {
		return true
	}

	return false
}

// SetValid gets a reference to the given bool and assigns it to the Valid field.
func (o *VerifyCertificateResponse) SetValid(v bool) {
	o.Valid = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *VerifyCertificateResponse) GetErrors() []string {
	if o == nil || o.Errors == nil {
		var ret []string
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyCertificateResponse) GetErrorsOk() ([]string, bool) {
	if o == nil || o.Errors == nil {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *VerifyCertificateResponse) HasErrors() bool {
	if o != nil && o.Errors != nil {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []string and assigns it to the Errors field.
func (o *VerifyCertificateResponse) SetErrors(v []string) {
	o.Errors = v
}

func (o VerifyCertificateResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Valid != nil {
		toSerialize["valid"] = o.Valid
	}
	if o.Errors != nil {
		toSerialize["errors"] = o.Errors
	}
	return json.Marshal(toSerialize)
}

type NullableVerifyCertificateResponse struct {
	value *VerifyCertificateResponse
	isSet bool
}

func (v NullableVerifyCertificateResponse) Get() *VerifyCertificateResponse {
	return v.value
}

func (v *NullableVerifyCertificateResponse) Set(val *VerifyCertificateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableVerifyCertificateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableVerifyCertificateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVerifyCertificateResponse(val *VerifyCertificateResponse) *NullableVerifyCertificateResponse {
	return &NullableVerifyCertificateResponse{value: val, isSet: true}
}

func (v NullableVerifyCertificateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVerifyCertificateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
