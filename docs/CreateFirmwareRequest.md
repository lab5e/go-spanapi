# CreateFirmwareRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CollectionId** | Pointer to **string** |  | [optional] 
**Image** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**Filename** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewCreateFirmwareRequest

`func NewCreateFirmwareRequest() *CreateFirmwareRequest`

NewCreateFirmwareRequest instantiates a new CreateFirmwareRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateFirmwareRequestWithDefaults

`func NewCreateFirmwareRequestWithDefaults() *CreateFirmwareRequest`

NewCreateFirmwareRequestWithDefaults instantiates a new CreateFirmwareRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollectionId

`func (o *CreateFirmwareRequest) GetCollectionId() string`

GetCollectionId returns the CollectionId field if non-nil, zero value otherwise.

### GetCollectionIdOk

`func (o *CreateFirmwareRequest) GetCollectionIdOk() (*string, bool)`

GetCollectionIdOk returns a tuple with the CollectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionId

`func (o *CreateFirmwareRequest) SetCollectionId(v string)`

SetCollectionId sets CollectionId field to given value.

### HasCollectionId

`func (o *CreateFirmwareRequest) HasCollectionId() bool`

HasCollectionId returns a boolean if a field has been set.

### GetImage

`func (o *CreateFirmwareRequest) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *CreateFirmwareRequest) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *CreateFirmwareRequest) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *CreateFirmwareRequest) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetVersion

`func (o *CreateFirmwareRequest) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CreateFirmwareRequest) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CreateFirmwareRequest) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *CreateFirmwareRequest) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetFilename

`func (o *CreateFirmwareRequest) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *CreateFirmwareRequest) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *CreateFirmwareRequest) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *CreateFirmwareRequest) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### GetTags

`func (o *CreateFirmwareRequest) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateFirmwareRequest) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateFirmwareRequest) SetTags(v map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateFirmwareRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


