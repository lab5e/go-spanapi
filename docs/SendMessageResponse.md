# SendMessageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CollectionId** | Pointer to **string** |  | [optional] 
**DeviceId** | Pointer to **string** |  | [optional] 
**BytesSent** | Pointer to **int32** |  | [optional] 

## Methods

### NewSendMessageResponse

`func NewSendMessageResponse() *SendMessageResponse`

NewSendMessageResponse instantiates a new SendMessageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendMessageResponseWithDefaults

`func NewSendMessageResponseWithDefaults() *SendMessageResponse`

NewSendMessageResponseWithDefaults instantiates a new SendMessageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollectionId

`func (o *SendMessageResponse) GetCollectionId() string`

GetCollectionId returns the CollectionId field if non-nil, zero value otherwise.

### GetCollectionIdOk

`func (o *SendMessageResponse) GetCollectionIdOk() (*string, bool)`

GetCollectionIdOk returns a tuple with the CollectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionId

`func (o *SendMessageResponse) SetCollectionId(v string)`

SetCollectionId sets CollectionId field to given value.

### HasCollectionId

`func (o *SendMessageResponse) HasCollectionId() bool`

HasCollectionId returns a boolean if a field has been set.

### GetDeviceId

`func (o *SendMessageResponse) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *SendMessageResponse) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *SendMessageResponse) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *SendMessageResponse) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetBytesSent

`func (o *SendMessageResponse) GetBytesSent() int32`

GetBytesSent returns the BytesSent field if non-nil, zero value otherwise.

### GetBytesSentOk

`func (o *SendMessageResponse) GetBytesSentOk() (*int32, bool)`

GetBytesSentOk returns a tuple with the BytesSent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytesSent

`func (o *SendMessageResponse) SetBytesSent(v int32)`

SetBytesSent sets BytesSent field to given value.

### HasBytesSent

`func (o *SendMessageResponse) HasBytesSent() bool`

HasBytesSent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


