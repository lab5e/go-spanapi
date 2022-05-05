# MessageUpstream

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MessageId** | Pointer to **string** |  | [optional] 
**CollectionId** | Pointer to **string** |  | [optional] 
**DeviceId** | Pointer to **string** |  | [optional] 
**GatewayId** | Pointer to **string** |  | [optional] 
**Transport** | Pointer to [**MessageTransport**](MessageTransport.md) |  | [optional] [default to MESSAGETRANSPORT_UNSPECIFIED]
**Received** | Pointer to **string** |  | [optional] 
**Payload** | Pointer to **string** |  | [optional] 

## Methods

### NewMessageUpstream

`func NewMessageUpstream() *MessageUpstream`

NewMessageUpstream instantiates a new MessageUpstream object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageUpstreamWithDefaults

`func NewMessageUpstreamWithDefaults() *MessageUpstream`

NewMessageUpstreamWithDefaults instantiates a new MessageUpstream object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessageId

`func (o *MessageUpstream) GetMessageId() string`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *MessageUpstream) GetMessageIdOk() (*string, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *MessageUpstream) SetMessageId(v string)`

SetMessageId sets MessageId field to given value.

### HasMessageId

`func (o *MessageUpstream) HasMessageId() bool`

HasMessageId returns a boolean if a field has been set.

### GetCollectionId

`func (o *MessageUpstream) GetCollectionId() string`

GetCollectionId returns the CollectionId field if non-nil, zero value otherwise.

### GetCollectionIdOk

`func (o *MessageUpstream) GetCollectionIdOk() (*string, bool)`

GetCollectionIdOk returns a tuple with the CollectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionId

`func (o *MessageUpstream) SetCollectionId(v string)`

SetCollectionId sets CollectionId field to given value.

### HasCollectionId

`func (o *MessageUpstream) HasCollectionId() bool`

HasCollectionId returns a boolean if a field has been set.

### GetDeviceId

`func (o *MessageUpstream) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *MessageUpstream) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *MessageUpstream) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *MessageUpstream) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetGatewayId

`func (o *MessageUpstream) GetGatewayId() string`

GetGatewayId returns the GatewayId field if non-nil, zero value otherwise.

### GetGatewayIdOk

`func (o *MessageUpstream) GetGatewayIdOk() (*string, bool)`

GetGatewayIdOk returns a tuple with the GatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayId

`func (o *MessageUpstream) SetGatewayId(v string)`

SetGatewayId sets GatewayId field to given value.

### HasGatewayId

`func (o *MessageUpstream) HasGatewayId() bool`

HasGatewayId returns a boolean if a field has been set.

### GetTransport

`func (o *MessageUpstream) GetTransport() MessageTransport`

GetTransport returns the Transport field if non-nil, zero value otherwise.

### GetTransportOk

`func (o *MessageUpstream) GetTransportOk() (*MessageTransport, bool)`

GetTransportOk returns a tuple with the Transport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransport

`func (o *MessageUpstream) SetTransport(v MessageTransport)`

SetTransport sets Transport field to given value.

### HasTransport

`func (o *MessageUpstream) HasTransport() bool`

HasTransport returns a boolean if a field has been set.

### GetReceived

`func (o *MessageUpstream) GetReceived() string`

GetReceived returns the Received field if non-nil, zero value otherwise.

### GetReceivedOk

`func (o *MessageUpstream) GetReceivedOk() (*string, bool)`

GetReceivedOk returns a tuple with the Received field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceived

`func (o *MessageUpstream) SetReceived(v string)`

SetReceived sets Received field to given value.

### HasReceived

`func (o *MessageUpstream) HasReceived() bool`

HasReceived returns a boolean if a field has been set.

### GetPayload

`func (o *MessageUpstream) GetPayload() string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *MessageUpstream) GetPayloadOk() (*string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *MessageUpstream) SetPayload(v string)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *MessageUpstream) HasPayload() bool`

HasPayload returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


