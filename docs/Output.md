# Output

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OutputId** | Pointer to **string** |  | [optional] 
**CollectionId** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**OutputType**](OutputType.md) |  | [optional] [default to UNDEFINED]
**Config** | Pointer to [**OutputConfig**](OutputConfig.md) |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Tags** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewOutput

`func NewOutput() *Output`

NewOutput instantiates a new Output object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutputWithDefaults

`func NewOutputWithDefaults() *Output`

NewOutputWithDefaults instantiates a new Output object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOutputId

`func (o *Output) GetOutputId() string`

GetOutputId returns the OutputId field if non-nil, zero value otherwise.

### GetOutputIdOk

`func (o *Output) GetOutputIdOk() (*string, bool)`

GetOutputIdOk returns a tuple with the OutputId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputId

`func (o *Output) SetOutputId(v string)`

SetOutputId sets OutputId field to given value.

### HasOutputId

`func (o *Output) HasOutputId() bool`

HasOutputId returns a boolean if a field has been set.

### GetCollectionId

`func (o *Output) GetCollectionId() string`

GetCollectionId returns the CollectionId field if non-nil, zero value otherwise.

### GetCollectionIdOk

`func (o *Output) GetCollectionIdOk() (*string, bool)`

GetCollectionIdOk returns a tuple with the CollectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionId

`func (o *Output) SetCollectionId(v string)`

SetCollectionId sets CollectionId field to given value.

### HasCollectionId

`func (o *Output) HasCollectionId() bool`

HasCollectionId returns a boolean if a field has been set.

### GetType

`func (o *Output) GetType() OutputType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Output) GetTypeOk() (*OutputType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Output) SetType(v OutputType)`

SetType sets Type field to given value.

### HasType

`func (o *Output) HasType() bool`

HasType returns a boolean if a field has been set.

### GetConfig

`func (o *Output) GetConfig() OutputConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *Output) GetConfigOk() (*OutputConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *Output) SetConfig(v OutputConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *Output) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetEnabled

`func (o *Output) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *Output) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *Output) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *Output) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetTags

`func (o *Output) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Output) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Output) SetTags(v map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Output) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


