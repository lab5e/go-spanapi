/*
The Span API

API for device, collection, output and firmware management

API version: 4.2.3 lower-elian
Contact: dev@lab5e.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package spanapi

import (
	"encoding/json"
)

// InetMetadata Metadata for devices connected via the internet gateway. This metadata shows the configuration for the  last message transmission.
type InetMetadata struct {
	GatewayId         *string `json:"gatewayId,omitempty"`
	LastUpdate        *string `json:"lastUpdate,omitempty"`
	RemoteAddress     *string `json:"remoteAddress,omitempty"`
	CertificateSerial *string `json:"certificateSerial,omitempty"`
}

// NewInetMetadata instantiates a new InetMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInetMetadata() *InetMetadata {
	this := InetMetadata{}
	return &this
}

// NewInetMetadataWithDefaults instantiates a new InetMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInetMetadataWithDefaults() *InetMetadata {
	this := InetMetadata{}
	return &this
}

// GetGatewayId returns the GatewayId field value if set, zero value otherwise.
func (o *InetMetadata) GetGatewayId() string {
	if o == nil || o.GatewayId == nil {
		var ret string
		return ret
	}
	return *o.GatewayId
}

// GetGatewayIdOk returns a tuple with the GatewayId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InetMetadata) GetGatewayIdOk() (*string, bool) {
	if o == nil || o.GatewayId == nil {
		return nil, false
	}
	return o.GatewayId, true
}

// HasGatewayId returns a boolean if a field has been set.
func (o *InetMetadata) HasGatewayId() bool {
	if o != nil && o.GatewayId != nil {
		return true
	}

	return false
}

// SetGatewayId gets a reference to the given string and assigns it to the GatewayId field.
func (o *InetMetadata) SetGatewayId(v string) {
	o.GatewayId = &v
}

// GetLastUpdate returns the LastUpdate field value if set, zero value otherwise.
func (o *InetMetadata) GetLastUpdate() string {
	if o == nil || o.LastUpdate == nil {
		var ret string
		return ret
	}
	return *o.LastUpdate
}

// GetLastUpdateOk returns a tuple with the LastUpdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InetMetadata) GetLastUpdateOk() (*string, bool) {
	if o == nil || o.LastUpdate == nil {
		return nil, false
	}
	return o.LastUpdate, true
}

// HasLastUpdate returns a boolean if a field has been set.
func (o *InetMetadata) HasLastUpdate() bool {
	if o != nil && o.LastUpdate != nil {
		return true
	}

	return false
}

// SetLastUpdate gets a reference to the given string and assigns it to the LastUpdate field.
func (o *InetMetadata) SetLastUpdate(v string) {
	o.LastUpdate = &v
}

// GetRemoteAddress returns the RemoteAddress field value if set, zero value otherwise.
func (o *InetMetadata) GetRemoteAddress() string {
	if o == nil || o.RemoteAddress == nil {
		var ret string
		return ret
	}
	return *o.RemoteAddress
}

// GetRemoteAddressOk returns a tuple with the RemoteAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InetMetadata) GetRemoteAddressOk() (*string, bool) {
	if o == nil || o.RemoteAddress == nil {
		return nil, false
	}
	return o.RemoteAddress, true
}

// HasRemoteAddress returns a boolean if a field has been set.
func (o *InetMetadata) HasRemoteAddress() bool {
	if o != nil && o.RemoteAddress != nil {
		return true
	}

	return false
}

// SetRemoteAddress gets a reference to the given string and assigns it to the RemoteAddress field.
func (o *InetMetadata) SetRemoteAddress(v string) {
	o.RemoteAddress = &v
}

// GetCertificateSerial returns the CertificateSerial field value if set, zero value otherwise.
func (o *InetMetadata) GetCertificateSerial() string {
	if o == nil || o.CertificateSerial == nil {
		var ret string
		return ret
	}
	return *o.CertificateSerial
}

// GetCertificateSerialOk returns a tuple with the CertificateSerial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InetMetadata) GetCertificateSerialOk() (*string, bool) {
	if o == nil || o.CertificateSerial == nil {
		return nil, false
	}
	return o.CertificateSerial, true
}

// HasCertificateSerial returns a boolean if a field has been set.
func (o *InetMetadata) HasCertificateSerial() bool {
	if o != nil && o.CertificateSerial != nil {
		return true
	}

	return false
}

// SetCertificateSerial gets a reference to the given string and assigns it to the CertificateSerial field.
func (o *InetMetadata) SetCertificateSerial(v string) {
	o.CertificateSerial = &v
}

func (o InetMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.GatewayId != nil {
		toSerialize["gatewayId"] = o.GatewayId
	}
	if o.LastUpdate != nil {
		toSerialize["lastUpdate"] = o.LastUpdate
	}
	if o.RemoteAddress != nil {
		toSerialize["remoteAddress"] = o.RemoteAddress
	}
	if o.CertificateSerial != nil {
		toSerialize["certificateSerial"] = o.CertificateSerial
	}
	return json.Marshal(toSerialize)
}

type NullableInetMetadata struct {
	value *InetMetadata
	isSet bool
}

func (v NullableInetMetadata) Get() *InetMetadata {
	return v.value
}

func (v *NullableInetMetadata) Set(val *InetMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableInetMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableInetMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInetMetadata(val *InetMetadata) *NullableInetMetadata {
	return &NullableInetMetadata{value: val, isSet: true}
}

func (v NullableInetMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInetMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
