# MessageDownstream

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MessageId** | Pointer to **string** |  | [optional] 
**CollectionId** | Pointer to **string** |  | [optional] 
**DeviceId** | Pointer to **string** |  | [optional] 
**GatewayId** | Pointer to **string** |  | [optional] 
**CreatedTime** | Pointer to **string** |  | [optional] 
**SentTime** | Pointer to **string** |  | [optional] 
**Transport** | Pointer to [**MessageTransport**](MessageTransport.md) |  | [optional] [default to MESSAGETRANSPORT_UNSPECIFIED]
**State** | Pointer to [**MessageState**](MessageState.md) |  | [optional] [default to MESSAGESTATE_UNSPECIFIED]
**Payload** | Pointer to **string** |  | [optional] 

## Methods

### NewMessageDownstream

`func NewMessageDownstream() *MessageDownstream`

NewMessageDownstream instantiates a new MessageDownstream object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageDownstreamWithDefaults

`func NewMessageDownstreamWithDefaults() *MessageDownstream`

NewMessageDownstreamWithDefaults instantiates a new MessageDownstream object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessageId

`func (o *MessageDownstream) GetMessageId() string`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *MessageDownstream) GetMessageIdOk() (*string, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *MessageDownstream) SetMessageId(v string)`

SetMessageId sets MessageId field to given value.

### HasMessageId

`func (o *MessageDownstream) HasMessageId() bool`

HasMessageId returns a boolean if a field has been set.

### GetCollectionId

`func (o *MessageDownstream) GetCollectionId() string`

GetCollectionId returns the CollectionId field if non-nil, zero value otherwise.

### GetCollectionIdOk

`func (o *MessageDownstream) GetCollectionIdOk() (*string, bool)`

GetCollectionIdOk returns a tuple with the CollectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionId

`func (o *MessageDownstream) SetCollectionId(v string)`

SetCollectionId sets CollectionId field to given value.

### HasCollectionId

`func (o *MessageDownstream) HasCollectionId() bool`

HasCollectionId returns a boolean if a field has been set.

### GetDeviceId

`func (o *MessageDownstream) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *MessageDownstream) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *MessageDownstream) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *MessageDownstream) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetGatewayId

`func (o *MessageDownstream) GetGatewayId() string`

GetGatewayId returns the GatewayId field if non-nil, zero value otherwise.

### GetGatewayIdOk

`func (o *MessageDownstream) GetGatewayIdOk() (*string, bool)`

GetGatewayIdOk returns a tuple with the GatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayId

`func (o *MessageDownstream) SetGatewayId(v string)`

SetGatewayId sets GatewayId field to given value.

### HasGatewayId

`func (o *MessageDownstream) HasGatewayId() bool`

HasGatewayId returns a boolean if a field has been set.

### GetCreatedTime

`func (o *MessageDownstream) GetCreatedTime() string`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *MessageDownstream) GetCreatedTimeOk() (*string, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *MessageDownstream) SetCreatedTime(v string)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *MessageDownstream) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetSentTime

`func (o *MessageDownstream) GetSentTime() string`

GetSentTime returns the SentTime field if non-nil, zero value otherwise.

### GetSentTimeOk

`func (o *MessageDownstream) GetSentTimeOk() (*string, bool)`

GetSentTimeOk returns a tuple with the SentTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentTime

`func (o *MessageDownstream) SetSentTime(v string)`

SetSentTime sets SentTime field to given value.

### HasSentTime

`func (o *MessageDownstream) HasSentTime() bool`

HasSentTime returns a boolean if a field has been set.

### GetTransport

`func (o *MessageDownstream) GetTransport() MessageTransport`

GetTransport returns the Transport field if non-nil, zero value otherwise.

### GetTransportOk

`func (o *MessageDownstream) GetTransportOk() (*MessageTransport, bool)`

GetTransportOk returns a tuple with the Transport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransport

`func (o *MessageDownstream) SetTransport(v MessageTransport)`

SetTransport sets Transport field to given value.

### HasTransport

`func (o *MessageDownstream) HasTransport() bool`

HasTransport returns a boolean if a field has been set.

### GetState

`func (o *MessageDownstream) GetState() MessageState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *MessageDownstream) GetStateOk() (*MessageState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *MessageDownstream) SetState(v MessageState)`

SetState sets State field to given value.

### HasState

`func (o *MessageDownstream) HasState() bool`

HasState returns a boolean if a field has been set.

### GetPayload

`func (o *MessageDownstream) GetPayload() string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *MessageDownstream) GetPayloadOk() (*string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *MessageDownstream) SetPayload(v string)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *MessageDownstream) HasPayload() bool`

HasPayload returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


