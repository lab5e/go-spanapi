# OutputDataMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**OutputMessageType**](OutputMessageType.md) |  | [optional] [default to UNKNOWN]
**Device** | Pointer to [**Device**](Device.md) |  | [optional] 
**Payload** | Pointer to **string** |  | [optional] 
**Received** | Pointer to **string** | Received time for message. Value is ms since epoch. | [optional] 
**Transport** | Pointer to **string** |  | [optional] 
**UdpMetaData** | Pointer to [**UDPMetadata**](UDPMetadata.md) |  | [optional] 
**CoapMetaData** | Pointer to [**CoAPMetadata**](CoAPMetadata.md) |  | [optional] 
**MessageId** | Pointer to **string** |  | [optional] 

## Methods

### NewOutputDataMessage

`func NewOutputDataMessage() *OutputDataMessage`

NewOutputDataMessage instantiates a new OutputDataMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutputDataMessageWithDefaults

`func NewOutputDataMessageWithDefaults() *OutputDataMessage`

NewOutputDataMessageWithDefaults instantiates a new OutputDataMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *OutputDataMessage) GetType() OutputMessageType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OutputDataMessage) GetTypeOk() (*OutputMessageType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OutputDataMessage) SetType(v OutputMessageType)`

SetType sets Type field to given value.

### HasType

`func (o *OutputDataMessage) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDevice

`func (o *OutputDataMessage) GetDevice() Device`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *OutputDataMessage) GetDeviceOk() (*Device, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *OutputDataMessage) SetDevice(v Device)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *OutputDataMessage) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetPayload

`func (o *OutputDataMessage) GetPayload() string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *OutputDataMessage) GetPayloadOk() (*string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *OutputDataMessage) SetPayload(v string)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *OutputDataMessage) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### GetReceived

`func (o *OutputDataMessage) GetReceived() string`

GetReceived returns the Received field if non-nil, zero value otherwise.

### GetReceivedOk

`func (o *OutputDataMessage) GetReceivedOk() (*string, bool)`

GetReceivedOk returns a tuple with the Received field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceived

`func (o *OutputDataMessage) SetReceived(v string)`

SetReceived sets Received field to given value.

### HasReceived

`func (o *OutputDataMessage) HasReceived() bool`

HasReceived returns a boolean if a field has been set.

### GetTransport

`func (o *OutputDataMessage) GetTransport() string`

GetTransport returns the Transport field if non-nil, zero value otherwise.

### GetTransportOk

`func (o *OutputDataMessage) GetTransportOk() (*string, bool)`

GetTransportOk returns a tuple with the Transport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransport

`func (o *OutputDataMessage) SetTransport(v string)`

SetTransport sets Transport field to given value.

### HasTransport

`func (o *OutputDataMessage) HasTransport() bool`

HasTransport returns a boolean if a field has been set.

### GetUdpMetaData

`func (o *OutputDataMessage) GetUdpMetaData() UDPMetadata`

GetUdpMetaData returns the UdpMetaData field if non-nil, zero value otherwise.

### GetUdpMetaDataOk

`func (o *OutputDataMessage) GetUdpMetaDataOk() (*UDPMetadata, bool)`

GetUdpMetaDataOk returns a tuple with the UdpMetaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdpMetaData

`func (o *OutputDataMessage) SetUdpMetaData(v UDPMetadata)`

SetUdpMetaData sets UdpMetaData field to given value.

### HasUdpMetaData

`func (o *OutputDataMessage) HasUdpMetaData() bool`

HasUdpMetaData returns a boolean if a field has been set.

### GetCoapMetaData

`func (o *OutputDataMessage) GetCoapMetaData() CoAPMetadata`

GetCoapMetaData returns the CoapMetaData field if non-nil, zero value otherwise.

### GetCoapMetaDataOk

`func (o *OutputDataMessage) GetCoapMetaDataOk() (*CoAPMetadata, bool)`

GetCoapMetaDataOk returns a tuple with the CoapMetaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoapMetaData

`func (o *OutputDataMessage) SetCoapMetaData(v CoAPMetadata)`

SetCoapMetaData sets CoapMetaData field to given value.

### HasCoapMetaData

`func (o *OutputDataMessage) HasCoapMetaData() bool`

HasCoapMetaData returns a boolean if a field has been set.

### GetMessageId

`func (o *OutputDataMessage) GetMessageId() string`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *OutputDataMessage) GetMessageIdOk() (*string, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *OutputDataMessage) SetMessageId(v string)`

SetMessageId sets MessageId field to given value.

### HasMessageId

`func (o *OutputDataMessage) HasMessageId() bool`

HasMessageId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


