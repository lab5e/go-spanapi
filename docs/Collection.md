# Collection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CollectionId** | Pointer to **string** | The ID of the collection. This is assigned by the backend. | [optional] 
**TeamId** | Pointer to **string** | The team ID that owns the collection. This field is required. When you create new collections the default is to use your private team ID. | [optional] 
**Firmware** | Pointer to [**CollectionFirmware**](CollectionFirmware.md) |  | [optional] 
**Tags** | Pointer to **map[string]string** | Tags for the collection. Tags are metadata fields that you can assign to the collection. | [optional] 
**UpstreamTimestamps** | Pointer to **[]string** |  | [optional] 
**DownstreamTimestamps** | Pointer to **[]string** |  | [optional] 

## Methods

### NewCollection

`func NewCollection() *Collection`

NewCollection instantiates a new Collection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionWithDefaults

`func NewCollectionWithDefaults() *Collection`

NewCollectionWithDefaults instantiates a new Collection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollectionId

`func (o *Collection) GetCollectionId() string`

GetCollectionId returns the CollectionId field if non-nil, zero value otherwise.

### GetCollectionIdOk

`func (o *Collection) GetCollectionIdOk() (*string, bool)`

GetCollectionIdOk returns a tuple with the CollectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionId

`func (o *Collection) SetCollectionId(v string)`

SetCollectionId sets CollectionId field to given value.

### HasCollectionId

`func (o *Collection) HasCollectionId() bool`

HasCollectionId returns a boolean if a field has been set.

### GetTeamId

`func (o *Collection) GetTeamId() string`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *Collection) GetTeamIdOk() (*string, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *Collection) SetTeamId(v string)`

SetTeamId sets TeamId field to given value.

### HasTeamId

`func (o *Collection) HasTeamId() bool`

HasTeamId returns a boolean if a field has been set.

### GetFirmware

`func (o *Collection) GetFirmware() CollectionFirmware`

GetFirmware returns the Firmware field if non-nil, zero value otherwise.

### GetFirmwareOk

`func (o *Collection) GetFirmwareOk() (*CollectionFirmware, bool)`

GetFirmwareOk returns a tuple with the Firmware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmware

`func (o *Collection) SetFirmware(v CollectionFirmware)`

SetFirmware sets Firmware field to given value.

### HasFirmware

`func (o *Collection) HasFirmware() bool`

HasFirmware returns a boolean if a field has been set.

### GetTags

`func (o *Collection) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Collection) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Collection) SetTags(v map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Collection) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUpstreamTimestamps

`func (o *Collection) GetUpstreamTimestamps() []string`

GetUpstreamTimestamps returns the UpstreamTimestamps field if non-nil, zero value otherwise.

### GetUpstreamTimestampsOk

`func (o *Collection) GetUpstreamTimestampsOk() (*[]string, bool)`

GetUpstreamTimestampsOk returns a tuple with the UpstreamTimestamps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpstreamTimestamps

`func (o *Collection) SetUpstreamTimestamps(v []string)`

SetUpstreamTimestamps sets UpstreamTimestamps field to given value.

### HasUpstreamTimestamps

`func (o *Collection) HasUpstreamTimestamps() bool`

HasUpstreamTimestamps returns a boolean if a field has been set.

### GetDownstreamTimestamps

`func (o *Collection) GetDownstreamTimestamps() []string`

GetDownstreamTimestamps returns the DownstreamTimestamps field if non-nil, zero value otherwise.

### GetDownstreamTimestampsOk

`func (o *Collection) GetDownstreamTimestampsOk() (*[]string, bool)`

GetDownstreamTimestampsOk returns a tuple with the DownstreamTimestamps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownstreamTimestamps

`func (o *Collection) SetDownstreamTimestamps(v []string)`

SetDownstreamTimestamps sets DownstreamTimestamps field to given value.

### HasDownstreamTimestamps

`func (o *Collection) HasDownstreamTimestamps() bool`

HasDownstreamTimestamps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


