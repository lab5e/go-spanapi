# CreateFirmwareBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Image** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**Filename** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewCreateFirmwareBody

`func NewCreateFirmwareBody() *CreateFirmwareBody`

NewCreateFirmwareBody instantiates a new CreateFirmwareBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateFirmwareBodyWithDefaults

`func NewCreateFirmwareBodyWithDefaults() *CreateFirmwareBody`

NewCreateFirmwareBodyWithDefaults instantiates a new CreateFirmwareBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImage

`func (o *CreateFirmwareBody) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *CreateFirmwareBody) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *CreateFirmwareBody) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *CreateFirmwareBody) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetVersion

`func (o *CreateFirmwareBody) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CreateFirmwareBody) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CreateFirmwareBody) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *CreateFirmwareBody) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetFilename

`func (o *CreateFirmwareBody) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *CreateFirmwareBody) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *CreateFirmwareBody) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *CreateFirmwareBody) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### GetTags

`func (o *CreateFirmwareBody) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateFirmwareBody) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateFirmwareBody) SetTags(v map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateFirmwareBody) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


