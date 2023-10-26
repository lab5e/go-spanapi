/*
The Span API

API for device, collection, output and firmware management

API version: 4.8.0 indecipherable-ferman
Contact: dev@lab5e.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package spanapi

import (
	"encoding/json"
)

// checks if the OutputDataMessage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OutputDataMessage{}

// OutputDataMessage The output data message contains payload plus metadata for a payload received from a device.
type OutputDataMessage struct {
	Type *OutputMessageType `json:"type,omitempty"`
	Device *Device `json:"device,omitempty"`
	Payload *string `json:"payload,omitempty"`
	// Received time for message. Value is ms since epoch.
	Received *string `json:"received,omitempty"`
	Transport *string `json:"transport,omitempty"`
	UdpMetaData *UDPMetadata `json:"udpMetaData,omitempty"`
	CoapMetaData *CoAPMetadata `json:"coapMetaData,omitempty"`
	MessageId *string `json:"messageId,omitempty"`
	MqttMetaData *MQTTMetadata `json:"mqttMetaData,omitempty"`
	GatewayMetaData *GatewayMetadata `json:"gatewayMetaData,omitempty"`
	GatewayId *string `json:"gatewayId,omitempty"`
}

// NewOutputDataMessage instantiates a new OutputDataMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOutputDataMessage() *OutputDataMessage {
	this := OutputDataMessage{}
	var type_ OutputMessageType = OUTPUTMESSAGETYPE_UNKNOWN
	this.Type = &type_
	return &this
}

// NewOutputDataMessageWithDefaults instantiates a new OutputDataMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOutputDataMessageWithDefaults() *OutputDataMessage {
	this := OutputDataMessage{}
	var type_ OutputMessageType = OUTPUTMESSAGETYPE_UNKNOWN
	this.Type = &type_
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *OutputDataMessage) GetType() OutputMessageType {
	if o == nil || IsNil(o.Type) {
		var ret OutputMessageType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputDataMessage) GetTypeOk() (*OutputMessageType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *OutputDataMessage) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given OutputMessageType and assigns it to the Type field.
func (o *OutputDataMessage) SetType(v OutputMessageType) {
	o.Type = &v
}

// GetDevice returns the Device field value if set, zero value otherwise.
func (o *OutputDataMessage) GetDevice() Device {
	if o == nil || IsNil(o.Device) {
		var ret Device
		return ret
	}
	return *o.Device
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputDataMessage) GetDeviceOk() (*Device, bool) {
	if o == nil || IsNil(o.Device) {
		return nil, false
	}
	return o.Device, true
}

// HasDevice returns a boolean if a field has been set.
func (o *OutputDataMessage) HasDevice() bool {
	if o != nil && !IsNil(o.Device) {
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
	if o == nil || IsNil(o.Payload) {
		var ret string
		return ret
	}
	return *o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputDataMessage) GetPayloadOk() (*string, bool) {
	if o == nil || IsNil(o.Payload) {
		return nil, false
	}
	return o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *OutputDataMessage) HasPayload() bool {
	if o != nil && !IsNil(o.Payload) {
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
	if o == nil || IsNil(o.Received) {
		var ret string
		return ret
	}
	return *o.Received
}

// GetReceivedOk returns a tuple with the Received field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputDataMessage) GetReceivedOk() (*string, bool) {
	if o == nil || IsNil(o.Received) {
		return nil, false
	}
	return o.Received, true
}

// HasReceived returns a boolean if a field has been set.
func (o *OutputDataMessage) HasReceived() bool {
	if o != nil && !IsNil(o.Received) {
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
	if o == nil || IsNil(o.Transport) {
		var ret string
		return ret
	}
	return *o.Transport
}

// GetTransportOk returns a tuple with the Transport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputDataMessage) GetTransportOk() (*string, bool) {
	if o == nil || IsNil(o.Transport) {
		return nil, false
	}
	return o.Transport, true
}

// HasTransport returns a boolean if a field has been set.
func (o *OutputDataMessage) HasTransport() bool {
	if o != nil && !IsNil(o.Transport) {
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
	if o == nil || IsNil(o.UdpMetaData) {
		var ret UDPMetadata
		return ret
	}
	return *o.UdpMetaData
}

// GetUdpMetaDataOk returns a tuple with the UdpMetaData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputDataMessage) GetUdpMetaDataOk() (*UDPMetadata, bool) {
	if o == nil || IsNil(o.UdpMetaData) {
		return nil, false
	}
	return o.UdpMetaData, true
}

// HasUdpMetaData returns a boolean if a field has been set.
func (o *OutputDataMessage) HasUdpMetaData() bool {
	if o != nil && !IsNil(o.UdpMetaData) {
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
	if o == nil || IsNil(o.CoapMetaData) {
		var ret CoAPMetadata
		return ret
	}
	return *o.CoapMetaData
}

// GetCoapMetaDataOk returns a tuple with the CoapMetaData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputDataMessage) GetCoapMetaDataOk() (*CoAPMetadata, bool) {
	if o == nil || IsNil(o.CoapMetaData) {
		return nil, false
	}
	return o.CoapMetaData, true
}

// HasCoapMetaData returns a boolean if a field has been set.
func (o *OutputDataMessage) HasCoapMetaData() bool {
	if o != nil && !IsNil(o.CoapMetaData) {
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
	if o == nil || IsNil(o.MessageId) {
		var ret string
		return ret
	}
	return *o.MessageId
}

// GetMessageIdOk returns a tuple with the MessageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputDataMessage) GetMessageIdOk() (*string, bool) {
	if o == nil || IsNil(o.MessageId) {
		return nil, false
	}
	return o.MessageId, true
}

// HasMessageId returns a boolean if a field has been set.
func (o *OutputDataMessage) HasMessageId() bool {
	if o != nil && !IsNil(o.MessageId) {
		return true
	}

	return false
}

// SetMessageId gets a reference to the given string and assigns it to the MessageId field.
func (o *OutputDataMessage) SetMessageId(v string) {
	o.MessageId = &v
}

// GetMqttMetaData returns the MqttMetaData field value if set, zero value otherwise.
func (o *OutputDataMessage) GetMqttMetaData() MQTTMetadata {
	if o == nil || IsNil(o.MqttMetaData) {
		var ret MQTTMetadata
		return ret
	}
	return *o.MqttMetaData
}

// GetMqttMetaDataOk returns a tuple with the MqttMetaData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputDataMessage) GetMqttMetaDataOk() (*MQTTMetadata, bool) {
	if o == nil || IsNil(o.MqttMetaData) {
		return nil, false
	}
	return o.MqttMetaData, true
}

// HasMqttMetaData returns a boolean if a field has been set.
func (o *OutputDataMessage) HasMqttMetaData() bool {
	if o != nil && !IsNil(o.MqttMetaData) {
		return true
	}

	return false
}

// SetMqttMetaData gets a reference to the given MQTTMetadata and assigns it to the MqttMetaData field.
func (o *OutputDataMessage) SetMqttMetaData(v MQTTMetadata) {
	o.MqttMetaData = &v
}

// GetGatewayMetaData returns the GatewayMetaData field value if set, zero value otherwise.
func (o *OutputDataMessage) GetGatewayMetaData() GatewayMetadata {
	if o == nil || IsNil(o.GatewayMetaData) {
		var ret GatewayMetadata
		return ret
	}
	return *o.GatewayMetaData
}

// GetGatewayMetaDataOk returns a tuple with the GatewayMetaData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputDataMessage) GetGatewayMetaDataOk() (*GatewayMetadata, bool) {
	if o == nil || IsNil(o.GatewayMetaData) {
		return nil, false
	}
	return o.GatewayMetaData, true
}

// HasGatewayMetaData returns a boolean if a field has been set.
func (o *OutputDataMessage) HasGatewayMetaData() bool {
	if o != nil && !IsNil(o.GatewayMetaData) {
		return true
	}

	return false
}

// SetGatewayMetaData gets a reference to the given GatewayMetadata and assigns it to the GatewayMetaData field.
func (o *OutputDataMessage) SetGatewayMetaData(v GatewayMetadata) {
	o.GatewayMetaData = &v
}

// GetGatewayId returns the GatewayId field value if set, zero value otherwise.
func (o *OutputDataMessage) GetGatewayId() string {
	if o == nil || IsNil(o.GatewayId) {
		var ret string
		return ret
	}
	return *o.GatewayId
}

// GetGatewayIdOk returns a tuple with the GatewayId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputDataMessage) GetGatewayIdOk() (*string, bool) {
	if o == nil || IsNil(o.GatewayId) {
		return nil, false
	}
	return o.GatewayId, true
}

// HasGatewayId returns a boolean if a field has been set.
func (o *OutputDataMessage) HasGatewayId() bool {
	if o != nil && !IsNil(o.GatewayId) {
		return true
	}

	return false
}

// SetGatewayId gets a reference to the given string and assigns it to the GatewayId field.
func (o *OutputDataMessage) SetGatewayId(v string) {
	o.GatewayId = &v
}

func (o OutputDataMessage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OutputDataMessage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Device) {
		toSerialize["device"] = o.Device
	}
	if !IsNil(o.Payload) {
		toSerialize["payload"] = o.Payload
	}
	if !IsNil(o.Received) {
		toSerialize["received"] = o.Received
	}
	if !IsNil(o.Transport) {
		toSerialize["transport"] = o.Transport
	}
	if !IsNil(o.UdpMetaData) {
		toSerialize["udpMetaData"] = o.UdpMetaData
	}
	if !IsNil(o.CoapMetaData) {
		toSerialize["coapMetaData"] = o.CoapMetaData
	}
	if !IsNil(o.MessageId) {
		toSerialize["messageId"] = o.MessageId
	}
	if !IsNil(o.MqttMetaData) {
		toSerialize["mqttMetaData"] = o.MqttMetaData
	}
	if !IsNil(o.GatewayMetaData) {
		toSerialize["gatewayMetaData"] = o.GatewayMetaData
	}
	if !IsNil(o.GatewayId) {
		toSerialize["gatewayId"] = o.GatewayId
	}
	return toSerialize, nil
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


