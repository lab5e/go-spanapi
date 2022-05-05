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

// SignCertificateRequest Request certificate signing
type SignCertificateRequest struct {
	GatewayId *string `json:"gatewayId,omitempty"`
	DeviceId  *string `json:"deviceId,omitempty"`
	Csr       *string `json:"csr,omitempty"`
}

// NewSignCertificateRequest instantiates a new SignCertificateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignCertificateRequest() *SignCertificateRequest {
	this := SignCertificateRequest{}
	return &this
}

// NewSignCertificateRequestWithDefaults instantiates a new SignCertificateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignCertificateRequestWithDefaults() *SignCertificateRequest {
	this := SignCertificateRequest{}
	return &this
}

// GetGatewayId returns the GatewayId field value if set, zero value otherwise.
func (o *SignCertificateRequest) GetGatewayId() string {
	if o == nil || o.GatewayId == nil {
		var ret string
		return ret
	}
	return *o.GatewayId
}

// GetGatewayIdOk returns a tuple with the GatewayId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignCertificateRequest) GetGatewayIdOk() (*string, bool) {
	if o == nil || o.GatewayId == nil {
		return nil, false
	}
	return o.GatewayId, true
}

// HasGatewayId returns a boolean if a field has been set.
func (o *SignCertificateRequest) HasGatewayId() bool {
	if o != nil && o.GatewayId != nil {
		return true
	}

	return false
}

// SetGatewayId gets a reference to the given string and assigns it to the GatewayId field.
func (o *SignCertificateRequest) SetGatewayId(v string) {
	o.GatewayId = &v
}

// GetDeviceId returns the DeviceId field value if set, zero value otherwise.
func (o *SignCertificateRequest) GetDeviceId() string {
	if o == nil || o.DeviceId == nil {
		var ret string
		return ret
	}
	return *o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignCertificateRequest) GetDeviceIdOk() (*string, bool) {
	if o == nil || o.DeviceId == nil {
		return nil, false
	}
	return o.DeviceId, true
}

// HasDeviceId returns a boolean if a field has been set.
func (o *SignCertificateRequest) HasDeviceId() bool {
	if o != nil && o.DeviceId != nil {
		return true
	}

	return false
}

// SetDeviceId gets a reference to the given string and assigns it to the DeviceId field.
func (o *SignCertificateRequest) SetDeviceId(v string) {
	o.DeviceId = &v
}

// GetCsr returns the Csr field value if set, zero value otherwise.
func (o *SignCertificateRequest) GetCsr() string {
	if o == nil || o.Csr == nil {
		var ret string
		return ret
	}
	return *o.Csr
}

// GetCsrOk returns a tuple with the Csr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignCertificateRequest) GetCsrOk() (*string, bool) {
	if o == nil || o.Csr == nil {
		return nil, false
	}
	return o.Csr, true
}

// HasCsr returns a boolean if a field has been set.
func (o *SignCertificateRequest) HasCsr() bool {
	if o != nil && o.Csr != nil {
		return true
	}

	return false
}

// SetCsr gets a reference to the given string and assigns it to the Csr field.
func (o *SignCertificateRequest) SetCsr(v string) {
	o.Csr = &v
}

func (o SignCertificateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.GatewayId != nil {
		toSerialize["gatewayId"] = o.GatewayId
	}
	if o.DeviceId != nil {
		toSerialize["deviceId"] = o.DeviceId
	}
	if o.Csr != nil {
		toSerialize["csr"] = o.Csr
	}
	return json.Marshal(toSerialize)
}

type NullableSignCertificateRequest struct {
	value *SignCertificateRequest
	isSet bool
}

func (v NullableSignCertificateRequest) Get() *SignCertificateRequest {
	return v.value
}

func (v *NullableSignCertificateRequest) Set(val *SignCertificateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSignCertificateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSignCertificateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignCertificateRequest(val *SignCertificateRequest) *NullableSignCertificateRequest {
	return &NullableSignCertificateRequest{value: val, isSet: true}
}

func (v NullableSignCertificateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignCertificateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
