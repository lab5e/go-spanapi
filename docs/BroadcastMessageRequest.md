# BroadcastMessageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CollectionId** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **int32** |  | [optional] 
**Payload** | Pointer to **string** |  | [optional] 
**Transport** | Pointer to **string** | Valid transports are \&quot;udp\&quot;, \&quot;coap\&quot;, \&quot;coap-pull\&quot;, \&quot;udp-pull\&quot;, \&quot;coap-push\&quot;, \&quot;udp-push\&quot;. \&quot;udp\&quot; is equivalent to \&quot;udp-push\&quot; and \&quot;coap\&quot; is equivalent to \&quot;coap-push\&quot;. Push messages are sent unsolicited to the device wheil pull messages are sent whenever the device wither sends data upstream (for UDP) or does a CoAP request to the CoAP service in span. | [optional] 
**CoapPath** | Pointer to **string** |  | [optional] 

## Methods

### NewBroadcastMessageRequest

`func NewBroadcastMessageRequest() *BroadcastMessageRequest`

NewBroadcastMessageRequest instantiates a new BroadcastMessageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBroadcastMessageRequestWithDefaults

`func NewBroadcastMessageRequestWithDefaults() *BroadcastMessageRequest`

NewBroadcastMessageRequestWithDefaults instantiates a new BroadcastMessageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollectionId

`func (o *BroadcastMessageRequest) GetCollectionId() string`

GetCollectionId returns the CollectionId field if non-nil, zero value otherwise.

### GetCollectionIdOk

`func (o *BroadcastMessageRequest) GetCollectionIdOk() (*string, bool)`

GetCollectionIdOk returns a tuple with the CollectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionId

`func (o *BroadcastMessageRequest) SetCollectionId(v string)`

SetCollectionId sets CollectionId field to given value.

### HasCollectionId

`func (o *BroadcastMessageRequest) HasCollectionId() bool`

HasCollectionId returns a boolean if a field has been set.

### GetPort

`func (o *BroadcastMessageRequest) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *BroadcastMessageRequest) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *BroadcastMessageRequest) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *BroadcastMessageRequest) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetPayload

`func (o *BroadcastMessageRequest) GetPayload() string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *BroadcastMessageRequest) GetPayloadOk() (*string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *BroadcastMessageRequest) SetPayload(v string)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *BroadcastMessageRequest) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### GetTransport

`func (o *BroadcastMessageRequest) GetTransport() string`

GetTransport returns the Transport field if non-nil, zero value otherwise.

### GetTransportOk

`func (o *BroadcastMessageRequest) GetTransportOk() (*string, bool)`

GetTransportOk returns a tuple with the Transport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransport

`func (o *BroadcastMessageRequest) SetTransport(v string)`

SetTransport sets Transport field to given value.

### HasTransport

`func (o *BroadcastMessageRequest) HasTransport() bool`

HasTransport returns a boolean if a field has been set.

### GetCoapPath

`func (o *BroadcastMessageRequest) GetCoapPath() string`

GetCoapPath returns the CoapPath field if non-nil, zero value otherwise.

### GetCoapPathOk

`func (o *BroadcastMessageRequest) GetCoapPathOk() (*string, bool)`

GetCoapPathOk returns a tuple with the CoapPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoapPath

`func (o *BroadcastMessageRequest) SetCoapPath(v string)`

SetCoapPath sets CoapPath field to given value.

### HasCoapPath

`func (o *BroadcastMessageRequest) HasCoapPath() bool`

HasCoapPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


