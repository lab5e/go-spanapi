# CreateGatewayBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**GatewayType**](GatewayType.md) |  | [optional] [default to GATEWAYTYPE_UNKNOWN]
**Config** | Pointer to [**GatewayConfig**](GatewayConfig.md) |  | [optional] 
**Tags** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewCreateGatewayBody

`func NewCreateGatewayBody() *CreateGatewayBody`

NewCreateGatewayBody instantiates a new CreateGatewayBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateGatewayBodyWithDefaults

`func NewCreateGatewayBodyWithDefaults() *CreateGatewayBody`

NewCreateGatewayBodyWithDefaults instantiates a new CreateGatewayBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateGatewayBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateGatewayBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateGatewayBody) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateGatewayBody) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *CreateGatewayBody) GetType() GatewayType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateGatewayBody) GetTypeOk() (*GatewayType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateGatewayBody) SetType(v GatewayType)`

SetType sets Type field to given value.

### HasType

`func (o *CreateGatewayBody) HasType() bool`

HasType returns a boolean if a field has been set.

### GetConfig

`func (o *CreateGatewayBody) GetConfig() GatewayConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *CreateGatewayBody) GetConfigOk() (*GatewayConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *CreateGatewayBody) SetConfig(v GatewayConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *CreateGatewayBody) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetTags

`func (o *CreateGatewayBody) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateGatewayBody) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateGatewayBody) SetTags(v map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateGatewayBody) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


