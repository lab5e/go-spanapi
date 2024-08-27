# UpdateOutputBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CollectionId** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**OutputType**](OutputType.md) |  | [optional] [default to OUTPUTTYPE_UNDEFINED]
**Config** | Pointer to [**OutputConfig**](OutputConfig.md) |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Tags** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewUpdateOutputBody

`func NewUpdateOutputBody() *UpdateOutputBody`

NewUpdateOutputBody instantiates a new UpdateOutputBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateOutputBodyWithDefaults

`func NewUpdateOutputBodyWithDefaults() *UpdateOutputBody`

NewUpdateOutputBodyWithDefaults instantiates a new UpdateOutputBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollectionId

`func (o *UpdateOutputBody) GetCollectionId() string`

GetCollectionId returns the CollectionId field if non-nil, zero value otherwise.

### GetCollectionIdOk

`func (o *UpdateOutputBody) GetCollectionIdOk() (*string, bool)`

GetCollectionIdOk returns a tuple with the CollectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionId

`func (o *UpdateOutputBody) SetCollectionId(v string)`

SetCollectionId sets CollectionId field to given value.

### HasCollectionId

`func (o *UpdateOutputBody) HasCollectionId() bool`

HasCollectionId returns a boolean if a field has been set.

### GetType

`func (o *UpdateOutputBody) GetType() OutputType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateOutputBody) GetTypeOk() (*OutputType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateOutputBody) SetType(v OutputType)`

SetType sets Type field to given value.

### HasType

`func (o *UpdateOutputBody) HasType() bool`

HasType returns a boolean if a field has been set.

### GetConfig

`func (o *UpdateOutputBody) GetConfig() OutputConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *UpdateOutputBody) GetConfigOk() (*OutputConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *UpdateOutputBody) SetConfig(v OutputConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *UpdateOutputBody) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetEnabled

`func (o *UpdateOutputBody) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateOutputBody) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateOutputBody) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdateOutputBody) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetTags

`func (o *UpdateOutputBody) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpdateOutputBody) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UpdateOutputBody) SetTags(v map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *UpdateOutputBody) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


