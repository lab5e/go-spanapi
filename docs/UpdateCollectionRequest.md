# UpdateCollectionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TeamId** | Pointer to **string** | The team ID that owns the collection. This field is required. When you create new collections the default is to use your private team ID. | [optional] 
**Firmware** | Pointer to [**CollectionFirmware**](CollectionFirmware.md) |  | [optional] 
**Tags** | Pointer to **map[string]string** | Tags for the collection. Tags are metadata fields that you can assign to the collection. | [optional] 

## Methods

### NewUpdateCollectionRequest

`func NewUpdateCollectionRequest() *UpdateCollectionRequest`

NewUpdateCollectionRequest instantiates a new UpdateCollectionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateCollectionRequestWithDefaults

`func NewUpdateCollectionRequestWithDefaults() *UpdateCollectionRequest`

NewUpdateCollectionRequestWithDefaults instantiates a new UpdateCollectionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTeamId

`func (o *UpdateCollectionRequest) GetTeamId() string`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *UpdateCollectionRequest) GetTeamIdOk() (*string, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *UpdateCollectionRequest) SetTeamId(v string)`

SetTeamId sets TeamId field to given value.

### HasTeamId

`func (o *UpdateCollectionRequest) HasTeamId() bool`

HasTeamId returns a boolean if a field has been set.

### GetFirmware

`func (o *UpdateCollectionRequest) GetFirmware() CollectionFirmware`

GetFirmware returns the Firmware field if non-nil, zero value otherwise.

### GetFirmwareOk

`func (o *UpdateCollectionRequest) GetFirmwareOk() (*CollectionFirmware, bool)`

GetFirmwareOk returns a tuple with the Firmware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmware

`func (o *UpdateCollectionRequest) SetFirmware(v CollectionFirmware)`

SetFirmware sets Firmware field to given value.

### HasFirmware

`func (o *UpdateCollectionRequest) HasFirmware() bool`

HasFirmware returns a boolean if a field has been set.

### GetTags

`func (o *UpdateCollectionRequest) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpdateCollectionRequest) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UpdateCollectionRequest) SetTags(v map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *UpdateCollectionRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


