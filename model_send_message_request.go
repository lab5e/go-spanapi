/*
 * The Span API
 *
 * API for device, collection, output and firmware management
 *
 * API version: 4.1.18 disgruntled-jerald
 * Contact: dev@lab5e.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package spanapi

import (
	"encoding/json"
)

// SendMessageRequest struct for SendMessageRequest
type SendMessageRequest struct {
	CollectionId *string `json:"collectionId,omitempty"`
	DeviceId *string `json:"deviceId,omitempty"`
	Port *int32 `json:"port,omitempty"`
	Payload *string `json:"payload,omitempty"`
	// Valid transports are \"udp\", \"coap\", \"coap-pull\", \"udp-pull\", \"coap-push\", \"udp-push\". \"udp\" is equivalent to \"udp-push\" and \"coap\" is equivalent to \"coap-push\". Push messages are sent unsolicited to the device wheil pull messages are sent whenever the device wither sends data upstream (for UDP) or does a CoAP request to the CoAP service in span.
	Transport *string `json:"transport,omitempty"`
	CoapPath *string `json:"coapPath,omitempty"`
}

// NewSendMessageRequest instantiates a new SendMessageRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSendMessageRequest() *SendMessageRequest {
	this := SendMessageRequest{}
	return &this
}

// NewSendMessageRequestWithDefaults instantiates a new SendMessageRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSendMessageRequestWithDefaults() *SendMessageRequest {
	this := SendMessageRequest{}
	return &this
}

// GetCollectionId returns the CollectionId field value if set, zero value otherwise.
func (o *SendMessageRequest) GetCollectionId() string {
	if o == nil || o.CollectionId == nil {
		var ret string
		return ret
	}
	return *o.CollectionId
}

// GetCollectionIdOk returns a tuple with the CollectionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendMessageRequest) GetCollectionIdOk() (*string, bool) {
	if o == nil || o.CollectionId == nil {
		return nil, false
	}
	return o.CollectionId, true
}

// HasCollectionId returns a boolean if a field has been set.
func (o *SendMessageRequest) HasCollectionId() bool {
	if o != nil && o.CollectionId != nil {
		return true
	}

	return false
}

// SetCollectionId gets a reference to the given string and assigns it to the CollectionId field.
func (o *SendMessageRequest) SetCollectionId(v string) {
	o.CollectionId = &v
}

// GetDeviceId returns the DeviceId field value if set, zero value otherwise.
func (o *SendMessageRequest) GetDeviceId() string {
	if o == nil || o.DeviceId == nil {
		var ret string
		return ret
	}
	return *o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendMessageRequest) GetDeviceIdOk() (*string, bool) {
	if o == nil || o.DeviceId == nil {
		return nil, false
	}
	return o.DeviceId, true
}

// HasDeviceId returns a boolean if a field has been set.
func (o *SendMessageRequest) HasDeviceId() bool {
	if o != nil && o.DeviceId != nil {
		return true
	}

	return false
}

// SetDeviceId gets a reference to the given string and assigns it to the DeviceId field.
func (o *SendMessageRequest) SetDeviceId(v string) {
	o.DeviceId = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *SendMessageRequest) GetPort() int32 {
	if o == nil || o.Port == nil {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendMessageRequest) GetPortOk() (*int32, bool) {
	if o == nil || o.Port == nil {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *SendMessageRequest) HasPort() bool {
	if o != nil && o.Port != nil {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *SendMessageRequest) SetPort(v int32) {
	o.Port = &v
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *SendMessageRequest) GetPayload() string {
	if o == nil || o.Payload == nil {
		var ret string
		return ret
	}
	return *o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendMessageRequest) GetPayloadOk() (*string, bool) {
	if o == nil || o.Payload == nil {
		return nil, false
	}
	return o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *SendMessageRequest) HasPayload() bool {
	if o != nil && o.Payload != nil {
		return true
	}

	return false
}

// SetPayload gets a reference to the given string and assigns it to the Payload field.
func (o *SendMessageRequest) SetPayload(v string) {
	o.Payload = &v
}

// GetTransport returns the Transport field value if set, zero value otherwise.
func (o *SendMessageRequest) GetTransport() string {
	if o == nil || o.Transport == nil {
		var ret string
		return ret
	}
	return *o.Transport
}

// GetTransportOk returns a tuple with the Transport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendMessageRequest) GetTransportOk() (*string, bool) {
	if o == nil || o.Transport == nil {
		return nil, false
	}
	return o.Transport, true
}

// HasTransport returns a boolean if a field has been set.
func (o *SendMessageRequest) HasTransport() bool {
	if o != nil && o.Transport != nil {
		return true
	}

	return false
}

// SetTransport gets a reference to the given string and assigns it to the Transport field.
func (o *SendMessageRequest) SetTransport(v string) {
	o.Transport = &v
}

// GetCoapPath returns the CoapPath field value if set, zero value otherwise.
func (o *SendMessageRequest) GetCoapPath() string {
	if o == nil || o.CoapPath == nil {
		var ret string
		return ret
	}
	return *o.CoapPath
}

// GetCoapPathOk returns a tuple with the CoapPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendMessageRequest) GetCoapPathOk() (*string, bool) {
	if o == nil || o.CoapPath == nil {
		return nil, false
	}
	return o.CoapPath, true
}

// HasCoapPath returns a boolean if a field has been set.
func (o *SendMessageRequest) HasCoapPath() bool {
	if o != nil && o.CoapPath != nil {
		return true
	}

	return false
}

// SetCoapPath gets a reference to the given string and assigns it to the CoapPath field.
func (o *SendMessageRequest) SetCoapPath(v string) {
	o.CoapPath = &v
}

func (o SendMessageRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CollectionId != nil {
		toSerialize["collectionId"] = o.CollectionId
	}
	if o.DeviceId != nil {
		toSerialize["deviceId"] = o.DeviceId
	}
	if o.Port != nil {
		toSerialize["port"] = o.Port
	}
	if o.Payload != nil {
		toSerialize["payload"] = o.Payload
	}
	if o.Transport != nil {
		toSerialize["transport"] = o.Transport
	}
	if o.CoapPath != nil {
		toSerialize["coapPath"] = o.CoapPath
	}
	return json.Marshal(toSerialize)
}

type NullableSendMessageRequest struct {
	value *SendMessageRequest
	isSet bool
}

func (v NullableSendMessageRequest) Get() *SendMessageRequest {
	return v.value
}

func (v *NullableSendMessageRequest) Set(val *SendMessageRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSendMessageRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSendMessageRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSendMessageRequest(val *SendMessageRequest) *NullableSendMessageRequest {
	return &NullableSendMessageRequest{value: val, isSet: true}
}

func (v NullableSendMessageRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSendMessageRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


