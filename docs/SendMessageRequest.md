# SendMessageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CollectionId** | Pointer to **string** |  | [optional] 
**DeviceId** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **int32** |  | [optional] 
**Payload** | Pointer to **string** |  | [optional] 
**Transport** | Pointer to **string** | Valid transports are \&quot;udp\&quot;, \&quot;coap\&quot;, \&quot;coap-pull\&quot;, \&quot;udp-pull\&quot;, \&quot;coap-push\&quot;, \&quot;udp-push\&quot;. \&quot;udp\&quot; is equivalent to \&quot;udp-push\&quot; and \&quot;coap\&quot; is equivalent to \&quot;coap-push\&quot;. Push messages are sent unsolicited to the device wheil pull messages are sent whenever the device wither sends data upstream (for UDP) or does a CoAP request to the CoAP service in span. | [optional] 
**CoapPath** | Pointer to **string** |  | [optional] 

## Methods

### NewSendMessageRequest

`func NewSendMessageRequest() *SendMessageRequest`

NewSendMessageRequest instantiates a new SendMessageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendMessageRequestWithDefaults

`func NewSendMessageRequestWithDefaults() *SendMessageRequest`

NewSendMessageRequestWithDefaults instantiates a new SendMessageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollectionId

`func (o *SendMessageRequest) GetCollectionId() string`

GetCollectionId returns the CollectionId field if non-nil, zero value otherwise.

### GetCollectionIdOk

`func (o *SendMessageRequest) GetCollectionIdOk() (*string, bool)`

GetCollectionIdOk returns a tuple with the CollectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionId

`func (o *SendMessageRequest) SetCollectionId(v string)`

SetCollectionId sets CollectionId field to given value.

### HasCollectionId

`func (o *SendMessageRequest) HasCollectionId() bool`

HasCollectionId returns a boolean if a field has been set.

### GetDeviceId

`func (o *SendMessageRequest) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *SendMessageRequest) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *SendMessageRequest) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *SendMessageRequest) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetPort

`func (o *SendMessageRequest) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *SendMessageRequest) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *SendMessageRequest) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *SendMessageRequest) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetPayload

`func (o *SendMessageRequest) GetPayload() string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *SendMessageRequest) GetPayloadOk() (*string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *SendMessageRequest) SetPayload(v string)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *SendMessageRequest) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### GetTransport

`func (o *SendMessageRequest) GetTransport() string`

GetTransport returns the Transport field if non-nil, zero value otherwise.

### GetTransportOk

`func (o *SendMessageRequest) GetTransportOk() (*string, bool)`

GetTransportOk returns a tuple with the Transport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransport

`func (o *SendMessageRequest) SetTransport(v string)`

SetTransport sets Transport field to given value.

### HasTransport

`func (o *SendMessageRequest) HasTransport() bool`

HasTransport returns a boolean if a field has been set.

### GetCoapPath

`func (o *SendMessageRequest) GetCoapPath() string`

GetCoapPath returns the CoapPath field if non-nil, zero value otherwise.

### GetCoapPathOk

`func (o *SendMessageRequest) GetCoapPathOk() (*string, bool)`

GetCoapPathOk returns a tuple with the CoapPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoapPath

`func (o *SendMessageRequest) SetCoapPath(v string)`

SetCoapPath sets CoapPath field to given value.

### HasCoapPath

`func (o *SendMessageRequest) HasCoapPath() bool`

HasCoapPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


