# Gateway

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GatewayId** | Pointer to **string** |  | [optional] 
**CollectionId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**BuiltIn** | Pointer to **bool** |  | [optional] 
**Type** | Pointer to [**GatewayType**](GatewayType.md) |  | [optional] [default to GATEWAYTYPE_UNKNOWN]
**Config** | Pointer to [**GatewayConfig**](GatewayConfig.md) |  | [optional] 
**Tags** | Pointer to **map[string]string** |  | [optional] 
**Status** | Pointer to [**GatewayStatus**](GatewayStatus.md) |  | [optional] [default to GATEWAYSTATUS_UNKNOWN]
**Enabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewGateway

`func NewGateway() *Gateway`

NewGateway instantiates a new Gateway object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayWithDefaults

`func NewGatewayWithDefaults() *Gateway`

NewGatewayWithDefaults instantiates a new Gateway object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGatewayId

`func (o *Gateway) GetGatewayId() string`

GetGatewayId returns the GatewayId field if non-nil, zero value otherwise.

### GetGatewayIdOk

`func (o *Gateway) GetGatewayIdOk() (*string, bool)`

GetGatewayIdOk returns a tuple with the GatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayId

`func (o *Gateway) SetGatewayId(v string)`

SetGatewayId sets GatewayId field to given value.

### HasGatewayId

`func (o *Gateway) HasGatewayId() bool`

HasGatewayId returns a boolean if a field has been set.

### GetCollectionId

`func (o *Gateway) GetCollectionId() string`

GetCollectionId returns the CollectionId field if non-nil, zero value otherwise.

### GetCollectionIdOk

`func (o *Gateway) GetCollectionIdOk() (*string, bool)`

GetCollectionIdOk returns a tuple with the CollectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionId

`func (o *Gateway) SetCollectionId(v string)`

SetCollectionId sets CollectionId field to given value.

### HasCollectionId

`func (o *Gateway) HasCollectionId() bool`

HasCollectionId returns a boolean if a field has been set.

### GetName

`func (o *Gateway) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Gateway) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Gateway) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Gateway) HasName() bool`

HasName returns a boolean if a field has been set.

### GetBuiltIn

`func (o *Gateway) GetBuiltIn() bool`

GetBuiltIn returns the BuiltIn field if non-nil, zero value otherwise.

### GetBuiltInOk

`func (o *Gateway) GetBuiltInOk() (*bool, bool)`

GetBuiltInOk returns a tuple with the BuiltIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuiltIn

`func (o *Gateway) SetBuiltIn(v bool)`

SetBuiltIn sets BuiltIn field to given value.

### HasBuiltIn

`func (o *Gateway) HasBuiltIn() bool`

HasBuiltIn returns a boolean if a field has been set.

### GetType

`func (o *Gateway) GetType() GatewayType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Gateway) GetTypeOk() (*GatewayType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Gateway) SetType(v GatewayType)`

SetType sets Type field to given value.

### HasType

`func (o *Gateway) HasType() bool`

HasType returns a boolean if a field has been set.

### GetConfig

`func (o *Gateway) GetConfig() GatewayConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *Gateway) GetConfigOk() (*GatewayConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *Gateway) SetConfig(v GatewayConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *Gateway) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetTags

`func (o *Gateway) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Gateway) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Gateway) SetTags(v map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Gateway) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetStatus

`func (o *Gateway) GetStatus() GatewayStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Gateway) GetStatusOk() (*GatewayStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Gateway) SetStatus(v GatewayStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Gateway) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetEnabled

`func (o *Gateway) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *Gateway) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *Gateway) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *Gateway) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


