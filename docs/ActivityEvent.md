# ActivityEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Event** | Pointer to **string** |  | [optional] 
**Time** | Pointer to **string** |  | [optional] 
**CollectionId** | Pointer to **string** |  | [optional] 
**DeviceId** | Pointer to **string** |  | [optional] 
**GatewayId** | Pointer to **string** |  | [optional] 
**Data** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewActivityEvent

`func NewActivityEvent() *ActivityEvent`

NewActivityEvent instantiates a new ActivityEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivityEventWithDefaults

`func NewActivityEventWithDefaults() *ActivityEvent`

NewActivityEventWithDefaults instantiates a new ActivityEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvent

`func (o *ActivityEvent) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *ActivityEvent) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *ActivityEvent) SetEvent(v string)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *ActivityEvent) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### GetTime

`func (o *ActivityEvent) GetTime() string`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *ActivityEvent) GetTimeOk() (*string, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *ActivityEvent) SetTime(v string)`

SetTime sets Time field to given value.

### HasTime

`func (o *ActivityEvent) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetCollectionId

`func (o *ActivityEvent) GetCollectionId() string`

GetCollectionId returns the CollectionId field if non-nil, zero value otherwise.

### GetCollectionIdOk

`func (o *ActivityEvent) GetCollectionIdOk() (*string, bool)`

GetCollectionIdOk returns a tuple with the CollectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionId

`func (o *ActivityEvent) SetCollectionId(v string)`

SetCollectionId sets CollectionId field to given value.

### HasCollectionId

`func (o *ActivityEvent) HasCollectionId() bool`

HasCollectionId returns a boolean if a field has been set.

### GetDeviceId

`func (o *ActivityEvent) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *ActivityEvent) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *ActivityEvent) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *ActivityEvent) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetGatewayId

`func (o *ActivityEvent) GetGatewayId() string`

GetGatewayId returns the GatewayId field if non-nil, zero value otherwise.

### GetGatewayIdOk

`func (o *ActivityEvent) GetGatewayIdOk() (*string, bool)`

GetGatewayIdOk returns a tuple with the GatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayId

`func (o *ActivityEvent) SetGatewayId(v string)`

SetGatewayId sets GatewayId field to given value.

### HasGatewayId

`func (o *ActivityEvent) HasGatewayId() bool`

HasGatewayId returns a boolean if a field has been set.

### GetData

`func (o *ActivityEvent) GetData() map[string]string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ActivityEvent) GetDataOk() (*map[string]string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ActivityEvent) SetData(v map[string]string)`

SetData sets Data field to given value.

### HasData

`func (o *ActivityEvent) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


