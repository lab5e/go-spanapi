/*
The Span API

API for device, collection, output and firmware management

API version: 4.2.2 smarmy-derik
Contact: dev@lab5e.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package spanapi

import (
	"encoding/json"
)

// CreateCertificateRequest Request object to create a new certificate.
type CreateCertificateRequest struct {
	GatewayId *string `json:"gatewayId,omitempty"`
	DeviceId  *string `json:"deviceId,omitempty"`
}

// NewCreateCertificateRequest instantiates a new CreateCertificateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateCertificateRequest() *CreateCertificateRequest {
	this := CreateCertificateRequest{}
	return &this
}

// NewCreateCertificateRequestWithDefaults instantiates a new CreateCertificateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateCertificateRequestWithDefaults() *CreateCertificateRequest {
	this := CreateCertificateRequest{}
	return &this
}

// GetGatewayId returns the GatewayId field value if set, zero value otherwise.
func (o *CreateCertificateRequest) GetGatewayId() string {
	if o == nil || o.GatewayId == nil {
		var ret string
		return ret
	}
	return *o.GatewayId
}

// GetGatewayIdOk returns a tuple with the GatewayId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCertificateRequest) GetGatewayIdOk() (*string, bool) {
	if o == nil || o.GatewayId == nil {
		return nil, false
	}
	return o.GatewayId, true
}

// HasGatewayId returns a boolean if a field has been set.
func (o *CreateCertificateRequest) HasGatewayId() bool {
	if o != nil && o.GatewayId != nil {
		return true
	}

	return false
}

// SetGatewayId gets a reference to the given string and assigns it to the GatewayId field.
func (o *CreateCertificateRequest) SetGatewayId(v string) {
	o.GatewayId = &v
}

// GetDeviceId returns the DeviceId field value if set, zero value otherwise.
func (o *CreateCertificateRequest) GetDeviceId() string {
	if o == nil || o.DeviceId == nil {
		var ret string
		return ret
	}
	return *o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCertificateRequest) GetDeviceIdOk() (*string, bool) {
	if o == nil || o.DeviceId == nil {
		return nil, false
	}
	return o.DeviceId, true
}

// HasDeviceId returns a boolean if a field has been set.
func (o *CreateCertificateRequest) HasDeviceId() bool {
	if o != nil && o.DeviceId != nil {
		return true
	}

	return false
}

// SetDeviceId gets a reference to the given string and assigns it to the DeviceId field.
func (o *CreateCertificateRequest) SetDeviceId(v string) {
	o.DeviceId = &v
}

func (o CreateCertificateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.GatewayId != nil {
		toSerialize["gatewayId"] = o.GatewayId
	}
	if o.DeviceId != nil {
		toSerialize["deviceId"] = o.DeviceId
	}
	return json.Marshal(toSerialize)
}

type NullableCreateCertificateRequest struct {
	value *CreateCertificateRequest
	isSet bool
}

func (v NullableCreateCertificateRequest) Get() *CreateCertificateRequest {
	return v.value
}

func (v *NullableCreateCertificateRequest) Set(val *CreateCertificateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateCertificateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateCertificateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateCertificateRequest(val *CreateCertificateRequest) *NullableCreateCertificateRequest {
	return &NullableCreateCertificateRequest{value: val, isSet: true}
}

func (v NullableCreateCertificateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateCertificateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
