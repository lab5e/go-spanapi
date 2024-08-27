# UpdateCollectionBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TeamId** | Pointer to **string** | The team ID that owns the collection. This field is required. When you create new collections the default is to use your private team ID. | [optional] 
**Firmware** | Pointer to [**CollectionFirmware**](CollectionFirmware.md) |  | [optional] 
**Tags** | Pointer to **map[string]string** | Tags for the collection. Tags are metadata fields that you can assign to the collection. | [optional] 
**Enabled** | Pointer to **bool** | Enabled flag for the collection. A collection may be disabled or enabled to save time. | [optional] 

## Methods

### NewUpdateCollectionBody

`func NewUpdateCollectionBody() *UpdateCollectionBody`

NewUpdateCollectionBody instantiates a new UpdateCollectionBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateCollectionBodyWithDefaults

`func NewUpdateCollectionBodyWithDefaults() *UpdateCollectionBody`

NewUpdateCollectionBodyWithDefaults instantiates a new UpdateCollectionBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTeamId

`func (o *UpdateCollectionBody) GetTeamId() string`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *UpdateCollectionBody) GetTeamIdOk() (*string, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *UpdateCollectionBody) SetTeamId(v string)`

SetTeamId sets TeamId field to given value.

### HasTeamId

`func (o *UpdateCollectionBody) HasTeamId() bool`

HasTeamId returns a boolean if a field has been set.

### GetFirmware

`func (o *UpdateCollectionBody) GetFirmware() CollectionFirmware`

GetFirmware returns the Firmware field if non-nil, zero value otherwise.

### GetFirmwareOk

`func (o *UpdateCollectionBody) GetFirmwareOk() (*CollectionFirmware, bool)`

GetFirmwareOk returns a tuple with the Firmware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmware

`func (o *UpdateCollectionBody) SetFirmware(v CollectionFirmware)`

SetFirmware sets Firmware field to given value.

### HasFirmware

`func (o *UpdateCollectionBody) HasFirmware() bool`

HasFirmware returns a boolean if a field has been set.

### GetTags

`func (o *UpdateCollectionBody) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpdateCollectionBody) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UpdateCollectionBody) SetTags(v map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *UpdateCollectionBody) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetEnabled

`func (o *UpdateCollectionBody) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateCollectionBody) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateCollectionBody) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdateCollectionBody) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


