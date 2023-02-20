# InlineObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**GatewayType**](GatewayType.md) |  | [optional] [default to GATEWAYTYPE_UNKNOWN]
**Config** | Pointer to [**GatewayConfig**](GatewayConfig.md) |  | [optional] 
**Tags** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewInlineObject

`func NewInlineObject() *InlineObject`

NewInlineObject instantiates a new InlineObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObjectWithDefaults

`func NewInlineObjectWithDefaults() *InlineObject`

NewInlineObjectWithDefaults instantiates a new InlineObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineObject) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineObject) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineObject) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineObject) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *InlineObject) GetType() GatewayType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InlineObject) GetTypeOk() (*GatewayType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InlineObject) SetType(v GatewayType)`

SetType sets Type field to given value.

### HasType

`func (o *InlineObject) HasType() bool`

HasType returns a boolean if a field has been set.

### GetConfig

`func (o *InlineObject) GetConfig() GatewayConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *InlineObject) GetConfigOk() (*GatewayConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *InlineObject) SetConfig(v GatewayConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *InlineObject) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetTags

`func (o *InlineObject) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *InlineObject) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *InlineObject) SetTags(v map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *InlineObject) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


