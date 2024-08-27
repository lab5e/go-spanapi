# CreateOutputBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**OutputType**](OutputType.md) |  | [optional] [default to OUTPUTTYPE_UNDEFINED]
**Config** | Pointer to [**OutputConfig**](OutputConfig.md) |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Tags** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewCreateOutputBody

`func NewCreateOutputBody() *CreateOutputBody`

NewCreateOutputBody instantiates a new CreateOutputBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOutputBodyWithDefaults

`func NewCreateOutputBodyWithDefaults() *CreateOutputBody`

NewCreateOutputBodyWithDefaults instantiates a new CreateOutputBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CreateOutputBody) GetType() OutputType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateOutputBody) GetTypeOk() (*OutputType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateOutputBody) SetType(v OutputType)`

SetType sets Type field to given value.

### HasType

`func (o *CreateOutputBody) HasType() bool`

HasType returns a boolean if a field has been set.

### GetConfig

`func (o *CreateOutputBody) GetConfig() OutputConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *CreateOutputBody) GetConfigOk() (*OutputConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *CreateOutputBody) SetConfig(v OutputConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *CreateOutputBody) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetEnabled

`func (o *CreateOutputBody) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CreateOutputBody) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CreateOutputBody) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CreateOutputBody) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetTags

`func (o *CreateOutputBody) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateOutputBody) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateOutputBody) SetTags(v map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateOutputBody) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


