# UpdateGatewayRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**CollectionId** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**GatewayType**](GatewayType.md) |  | [optional] [default to GATEWAYTYPE_UNKNOWN]
**Config** | Pointer to [**GatewayConfig**](GatewayConfig.md) |  | [optional] 
**Tags** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewUpdateGatewayRequest

`func NewUpdateGatewayRequest() *UpdateGatewayRequest`

NewUpdateGatewayRequest instantiates a new UpdateGatewayRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateGatewayRequestWithDefaults

`func NewUpdateGatewayRequestWithDefaults() *UpdateGatewayRequest`

NewUpdateGatewayRequestWithDefaults instantiates a new UpdateGatewayRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateGatewayRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateGatewayRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateGatewayRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateGatewayRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCollectionId

`func (o *UpdateGatewayRequest) GetCollectionId() string`

GetCollectionId returns the CollectionId field if non-nil, zero value otherwise.

### GetCollectionIdOk

`func (o *UpdateGatewayRequest) GetCollectionIdOk() (*string, bool)`

GetCollectionIdOk returns a tuple with the CollectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionId

`func (o *UpdateGatewayRequest) SetCollectionId(v string)`

SetCollectionId sets CollectionId field to given value.

### HasCollectionId

`func (o *UpdateGatewayRequest) HasCollectionId() bool`

HasCollectionId returns a boolean if a field has been set.

### GetType

`func (o *UpdateGatewayRequest) GetType() GatewayType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateGatewayRequest) GetTypeOk() (*GatewayType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateGatewayRequest) SetType(v GatewayType)`

SetType sets Type field to given value.

### HasType

`func (o *UpdateGatewayRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetConfig

`func (o *UpdateGatewayRequest) GetConfig() GatewayConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *UpdateGatewayRequest) GetConfigOk() (*GatewayConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *UpdateGatewayRequest) SetConfig(v GatewayConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *UpdateGatewayRequest) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetTags

`func (o *UpdateGatewayRequest) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpdateGatewayRequest) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UpdateGatewayRequest) SetTags(v map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *UpdateGatewayRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


