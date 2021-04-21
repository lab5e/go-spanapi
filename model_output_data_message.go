/*
 * The Span API
 *
 * API for device, collection, output and firmware management
 *
 * API version: 4.1.9 receding-glennis
 * Contact: dev@lab5e.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package spanapi

import (
	"encoding/json"
)

// OutputDataMessage The output data message contains payload plus metadata for a payload received from a device.
type OutputDataMessage struct {
	Type *OutputDataMessageOutputMessageType `json:"type,omitempty"`
	Device *Device `json:"device,omitempty"`
	Payload *string `json:"payload,omitempty"`
	// Received time for message. Value is ms since epoch.
	Received *string `json:"received,omitempty"`
	Transport *string `json:"transport,omitempty"`
	UdpMetaData *UDPMetadata `json:"udpMetaData,omitempty"`
	CoapMetaData *CoAPMetadata `json:"coapMetaData,omitempty"`
	MessageId *string `json:"messageId,omitempty"`
}

// NewOutputDataMessage instantiates a new OutputDataMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOutputDataMessage() *OutputDataMessage {
	this := OutputDataMessage{}
	var type_ OutputDataMessageOutputMessageType = UNKNOWN
	this.Type = &type_
	return &this
}

// NewOutputDataMessageWithDefaults instantiates a new OutputDataMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOutputDataMessageWithDefaults() *OutputDataMessage {
	this := OutputDataMessage{}
	var type_ OutputDataMessageOutputMessageType = UNKNOWN
	this.Type = &type_
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *OutputDataMessage) GetType() OutputDataMessageOutputMessageType {
	if o == nil || o.Type == nil {
		var ret OutputDataMessageOutputMessageType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputDataMessage) GetTypeOk() (*OutputDataMessageOutputMessageType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *OutputDataMessage) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given OutputDataMessageOutputMessageType and assigns it to the Type field.
func (o *OutputDataMessage) SetType(v OutputDataMessageOutputMessageType) {
	o.Type = &v
}

// GetDevice returns the Device field value if set, zero value otherwise.
func (o *OutputDataMessage) GetDevice() Device {
	if o == nil || o.Device == nil {
		var ret Device
		return ret
	}
	return *o.Device
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputDataMessage) GetDeviceOk() (*Device, bool) {
	if o == nil || o.Device == nil {
		return nil, false
	}
	return o.Device, true
}

// HasDevice returns a boolean if a field has been set.
func (o *OutputDataMessage) HasDevice() bool {
	if o != nil && o.Device != nil {
		return true
	}

	return false
}

// SetDevice gets a reference to the given Device and assigns it to the Device field.
func (o *OutputDataMessage) SetDevice(v Device) {
	o.Device = &v
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *OutputDataMessage) GetPayload() string {
	if o == nil || o.Payload == nil {
		var ret string
		return ret
	}
	return *o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputDataMessage) GetPayloadOk() (*string, bool) {
	if o == nil || o.Payload == nil {
		return nil, false
	}
	return o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *OutputDataMessage) HasPayload() bool {
	if o != nil && o.Payload != nil {
		return true
	}

	return false
}

// SetPayload gets a reference to the given string and assigns it to the Payload field.
func (o *OutputDataMessage) SetPayload(v string) {
	o.Payload = &v
}

// GetReceived returns the Received field value if set, zero value otherwise.
func (o *OutputDataMessage) GetReceived() string {
	if o == nil || o.Received == nil {
		var ret string
		return ret
	}
	return *o.Received
}

// GetReceivedOk returns a tuple with the Received field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputDataMessage) GetReceivedOk() (*string, bool) {
	if o == nil || o.Received == nil {
		return nil, false
	}
	return o.Received, true
}

// HasReceived returns a boolean if a field has been set.
func (o *OutputDataMessage) HasReceived() bool {
	if o != nil && o.Received != nil {
		return true
	}

	return false
}

// SetReceived gets a reference to the given string and assigns it to the Received field.
func (o *OutputDataMessage) SetReceived(v string) {
	o.Received = &v
}

// GetTransport returns the Transport field value if set, zero value otherwise.
func (o *OutputDataMessage) GetTransport() string {
	if o == nil || o.Transport == nil {
		var ret string
		return ret
	}
	return *o.Transport
}

// GetTransportOk returns a tuple with the Transport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputDataMessage) GetTransportOk() (*string, bool) {
	if o == nil || o.Transport == nil {
		return nil, false
	}
	return o.Transport, true
}

// HasTransport returns a boolean if a field has been set.
func (o *OutputDataMessage) HasTransport() bool {
	if o != nil && o.Transport != nil {
		return true
	}

	return false
}

// SetTransport gets a reference to the given string and assigns it to the Transport field.
func (o *OutputDataMessage) SetTransport(v string) {
	o.Transport = &v
}

// GetUdpMetaData returns the UdpMetaData field value if set, zero value otherwise.
func (o *OutputDataMessage) GetUdpMetaData() UDPMetadata {
	if o == nil || o.UdpMetaData == nil {
		var ret UDPMetadata
		return ret
	}
	return *o.UdpMetaData
}

// GetUdpMetaDataOk returns a tuple with the UdpMetaData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputDataMessage) GetUdpMetaDataOk() (*UDPMetadata, bool) {
	if o == nil || o.UdpMetaData == nil {
		return nil, false
	}
	return o.UdpMetaData, true
}

// HasUdpMetaData returns a boolean if a field has been set.
func (o *OutputDataMessage) HasUdpMetaData() bool {
	if o != nil && o.UdpMetaData != nil {
		return true
	}

	return false
}

// SetUdpMetaData gets a reference to the given UDPMetadata and assigns it to the UdpMetaData field.
func (o *OutputDataMessage) SetUdpMetaData(v UDPMetadata) {
	o.UdpMetaData = &v
}

// GetCoapMetaData returns the CoapMetaData field value if set, zero value otherwise.
func (o *OutputDataMessage) GetCoapMetaData() CoAPMetadata {
	if o == nil || o.CoapMetaData == nil {
		var ret CoAPMetadata
		return ret
	}
	return *o.CoapMetaData
}

// GetCoapMetaDataOk returns a tuple with the CoapMetaData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputDataMessage) GetCoapMetaDataOk() (*CoAPMetadata, bool) {
	if o == nil || o.CoapMetaData == nil {
		return nil, false
	}
	return o.CoapMetaData, true
}

// HasCoapMetaData returns a boolean if a field has been set.
func (o *OutputDataMessage) HasCoapMetaData() bool {
	if o != nil && o.CoapMetaData != nil {
		return true
	}

	return false
}

// SetCoapMetaData gets a reference to the given CoAPMetadata and assigns it to the CoapMetaData field.
func (o *OutputDataMessage) SetCoapMetaData(v CoAPMetadata) {
	o.CoapMetaData = &v
}

// GetMessageId returns the MessageId field value if set, zero value otherwise.
func (o *OutputDataMessage) GetMessageId() string {
	if o == nil || o.MessageId == nil {
		var ret string
		return ret
	}
	return *o.MessageId
}

// GetMessageIdOk returns a tuple with the MessageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputDataMessage) GetMessageIdOk() (*string, bool) {
	if o == nil || o.MessageId == nil {
		return nil, false
	}
	return o.MessageId, true
}

// HasMessageId returns a boolean if a field has been set.
func (o *OutputDataMessage) HasMessageId() bool {
	if o != nil && o.MessageId != nil {
		return true
	}

	return false
}

// SetMessageId gets a reference to the given string and assigns it to the MessageId field.
func (o *OutputDataMessage) SetMessageId(v string) {
	o.MessageId = &v
}

func (o OutputDataMessage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Device != nil {
		toSerialize["device"] = o.Device
	}
	if o.Payload != nil {
		toSerialize["payload"] = o.Payload
	}
	if o.Received != nil {
		toSerialize["received"] = o.Received
	}
	if o.Transport != nil {
		toSerialize["transport"] = o.Transport
	}
	if o.UdpMetaData != nil {
		toSerialize["udpMetaData"] = o.UdpMetaData
	}
	if o.CoapMetaData != nil {
		toSerialize["coapMetaData"] = o.CoapMetaData
	}
	if o.MessageId != nil {
		toSerialize["messageId"] = o.MessageId
	}
	return json.Marshal(toSerialize)
}

type NullableOutputDataMessage struct {
	value *OutputDataMessage
	isSet bool
}

func (v NullableOutputDataMessage) Get() *OutputDataMessage {
	return v.value
}

func (v *NullableOutputDataMessage) Set(val *OutputDataMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableOutputDataMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableOutputDataMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOutputDataMessage(val *OutputDataMessage) *NullableOutputDataMessage {
	return &NullableOutputDataMessage{value: val, isSet: true}
}

func (v NullableOutputDataMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOutputDataMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


