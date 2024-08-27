# Firmware

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImageId** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** | Make sure that the firmware image reports this version. If the version reported by this image is different the FOTA process won&#39;t be reported as successful since the device will report a version different from this. | [optional] 
**Filename** | Pointer to **string** | This is just for internal house keeping, making it potentially easier to identify the firmware image. | [optional] 
**Sha256** | Pointer to **string** | To ensure each image is unique the SHA-256 hash for all images in a collection must be unqique | [optional] 
**Length** | Pointer to **int32** |  | [optional] 
**CollectionId** | Pointer to **string** | Collection ID for the collection owning the firmware image. | [optional] 
**Created** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **map[string]string** | Tags for firmware image. | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewFirmware

`func NewFirmware() *Firmware`

NewFirmware instantiates a new Firmware object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirmwareWithDefaults

`func NewFirmwareWithDefaults() *Firmware`

NewFirmwareWithDefaults instantiates a new Firmware object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImageId

`func (o *Firmware) GetImageId() string`

GetImageId returns the ImageId field if non-nil, zero value otherwise.

### GetImageIdOk

`func (o *Firmware) GetImageIdOk() (*string, bool)`

GetImageIdOk returns a tuple with the ImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageId

`func (o *Firmware) SetImageId(v string)`

SetImageId sets ImageId field to given value.

### HasImageId

`func (o *Firmware) HasImageId() bool`

HasImageId returns a boolean if a field has been set.

### GetVersion

`func (o *Firmware) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Firmware) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Firmware) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Firmware) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetFilename

`func (o *Firmware) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *Firmware) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *Firmware) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *Firmware) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### GetSha256

`func (o *Firmware) GetSha256() string`

GetSha256 returns the Sha256 field if non-nil, zero value otherwise.

### GetSha256Ok

`func (o *Firmware) GetSha256Ok() (*string, bool)`

GetSha256Ok returns a tuple with the Sha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256

`func (o *Firmware) SetSha256(v string)`

SetSha256 sets Sha256 field to given value.

### HasSha256

`func (o *Firmware) HasSha256() bool`

HasSha256 returns a boolean if a field has been set.

### GetLength

`func (o *Firmware) GetLength() int32`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *Firmware) GetLengthOk() (*int32, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *Firmware) SetLength(v int32)`

SetLength sets Length field to given value.

### HasLength

`func (o *Firmware) HasLength() bool`

HasLength returns a boolean if a field has been set.

### GetCollectionId

`func (o *Firmware) GetCollectionId() string`

GetCollectionId returns the CollectionId field if non-nil, zero value otherwise.

### GetCollectionIdOk

`func (o *Firmware) GetCollectionIdOk() (*string, bool)`

GetCollectionIdOk returns a tuple with the CollectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionId

`func (o *Firmware) SetCollectionId(v string)`

SetCollectionId sets CollectionId field to given value.

### HasCollectionId

`func (o *Firmware) HasCollectionId() bool`

HasCollectionId returns a boolean if a field has been set.

### GetCreated

`func (o *Firmware) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Firmware) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Firmware) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *Firmware) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetTags

`func (o *Firmware) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Firmware) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Firmware) SetTags(v map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Firmware) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetEnabled

`func (o *Firmware) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *Firmware) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *Firmware) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *Firmware) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


