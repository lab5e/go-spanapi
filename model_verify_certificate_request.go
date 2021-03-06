/*
The Span API

API for device, collection, output and firmware management

API version: 4.2.4 curable-andres
Contact: dev@lab5e.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package spanapi

import (
	"encoding/json"
)

// VerifyCertificateRequest Verify a certificate
type VerifyCertificateRequest struct {
	GatewayId   *string `json:"gatewayId,omitempty"`
	DeviceId    *string `json:"deviceId,omitempty"`
	Certificate *string `json:"certificate,omitempty"`
}

// NewVerifyCertificateRequest instantiates a new VerifyCertificateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVerifyCertificateRequest() *VerifyCertificateRequest {
	this := VerifyCertificateRequest{}
	return &this
}

// NewVerifyCertificateRequestWithDefaults instantiates a new VerifyCertificateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVerifyCertificateRequestWithDefaults() *VerifyCertificateRequest {
	this := VerifyCertificateRequest{}
	return &this
}

// GetGatewayId returns the GatewayId field value if set, zero value otherwise.
func (o *VerifyCertificateRequest) GetGatewayId() string {
	if o == nil || o.GatewayId == nil {
		var ret string
		return ret
	}
	return *o.GatewayId
}

// GetGatewayIdOk returns a tuple with the GatewayId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyCertificateRequest) GetGatewayIdOk() (*string, bool) {
	if o == nil || o.GatewayId == nil {
		return nil, false
	}
	return o.GatewayId, true
}

// HasGatewayId returns a boolean if a field has been set.
func (o *VerifyCertificateRequest) HasGatewayId() bool {
	if o != nil && o.GatewayId != nil {
		return true
	}

	return false
}

// SetGatewayId gets a reference to the given string and assigns it to the GatewayId field.
func (o *VerifyCertificateRequest) SetGatewayId(v string) {
	o.GatewayId = &v
}

// GetDeviceId returns the DeviceId field value if set, zero value otherwise.
func (o *VerifyCertificateRequest) GetDeviceId() string {
	if o == nil || o.DeviceId == nil {
		var ret string
		return ret
	}
	return *o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyCertificateRequest) GetDeviceIdOk() (*string, bool) {
	if o == nil || o.DeviceId == nil {
		return nil, false
	}
	return o.DeviceId, true
}

// HasDeviceId returns a boolean if a field has been set.
func (o *VerifyCertificateRequest) HasDeviceId() bool {
	if o != nil && o.DeviceId != nil {
		return true
	}

	return false
}

// SetDeviceId gets a reference to the given string and assigns it to the DeviceId field.
func (o *VerifyCertificateRequest) SetDeviceId(v string) {
	o.DeviceId = &v
}

// GetCertificate returns the Certificate field value if set, zero value otherwise.
func (o *VerifyCertificateRequest) GetCertificate() string {
	if o == nil || o.Certificate == nil {
		var ret string
		return ret
	}
	return *o.Certificate
}

// GetCertificateOk returns a tuple with the Certificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyCertificateRequest) GetCertificateOk() (*string, bool) {
	if o == nil || o.Certificate == nil {
		return nil, false
	}
	return o.Certificate, true
}

// HasCertificate returns a boolean if a field has been set.
func (o *VerifyCertificateRequest) HasCertificate() bool {
	if o != nil && o.Certificate != nil {
		return true
	}

	return false
}

// SetCertificate gets a reference to the given string and assigns it to the Certificate field.
func (o *VerifyCertificateRequest) SetCertificate(v string) {
	o.Certificate = &v
}

func (o VerifyCertificateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.GatewayId != nil {
		toSerialize["gatewayId"] = o.GatewayId
	}
	if o.DeviceId != nil {
		toSerialize["deviceId"] = o.DeviceId
	}
	if o.Certificate != nil {
		toSerialize["certificate"] = o.Certificate
	}
	return json.Marshal(toSerialize)
}

type NullableVerifyCertificateRequest struct {
	value *VerifyCertificateRequest
	isSet bool
}

func (v NullableVerifyCertificateRequest) Get() *VerifyCertificateRequest {
	return v.value
}

func (v *NullableVerifyCertificateRequest) Set(val *VerifyCertificateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableVerifyCertificateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableVerifyCertificateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVerifyCertificateRequest(val *VerifyCertificateRequest) *NullableVerifyCertificateRequest {
	return &NullableVerifyCertificateRequest{value: val, isSet: true}
}

func (v NullableVerifyCertificateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVerifyCertificateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
