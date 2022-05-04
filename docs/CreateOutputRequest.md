# CreateOutputRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**OutputType**](OutputType.md) |  | [optional] [default to UNDEFINED]
**Config** | Pointer to [**OutputConfig**](OutputConfig.md) |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Tags** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewCreateOutputRequest

`func NewCreateOutputRequest() *CreateOutputRequest`

NewCreateOutputRequest instantiates a new CreateOutputRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOutputRequestWithDefaults

`func NewCreateOutputRequestWithDefaults() *CreateOutputRequest`

NewCreateOutputRequestWithDefaults instantiates a new CreateOutputRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CreateOutputRequest) GetType() OutputType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateOutputRequest) GetTypeOk() (*OutputType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateOutputRequest) SetType(v OutputType)`

SetType sets Type field to given value.

### HasType

`func (o *CreateOutputRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetConfig

`func (o *CreateOutputRequest) GetConfig() OutputConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *CreateOutputRequest) GetConfigOk() (*OutputConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *CreateOutputRequest) SetConfig(v OutputConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *CreateOutputRequest) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetEnabled

`func (o *CreateOutputRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CreateOutputRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CreateOutputRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CreateOutputRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetTags

`func (o *CreateOutputRequest) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateOutputRequest) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateOutputRequest) SetTags(v map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateOutputRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


