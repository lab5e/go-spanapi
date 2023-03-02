# UpdateDeviceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CollectionId** | Pointer to **string** | The collection id for the device. This field is optional and can be omitted if the collection id isn&#39;t changed. When changing to a new collection you must be an owner of the other collection, ie an administrator of the team that owns the new collection. | [optional] 
**Tags** | Pointer to **map[string]string** | Tags are metadata for the device that you can set. These are just strings. | [optional] 
**Firmware** | Pointer to [**FirmwareMetadata**](FirmwareMetadata.md) |  | [optional] 
**Config** | Pointer to [**DeviceConfig**](DeviceConfig.md) |  | [optional] 

## Methods

### NewUpdateDeviceRequest

`func NewUpdateDeviceRequest() *UpdateDeviceRequest`

NewUpdateDeviceRequest instantiates a new UpdateDeviceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDeviceRequestWithDefaults

`func NewUpdateDeviceRequestWithDefaults() *UpdateDeviceRequest`

NewUpdateDeviceRequestWithDefaults instantiates a new UpdateDeviceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollectionId

`func (o *UpdateDeviceRequest) GetCollectionId() string`

GetCollectionId returns the CollectionId field if non-nil, zero value otherwise.

### GetCollectionIdOk

`func (o *UpdateDeviceRequest) GetCollectionIdOk() (*string, bool)`

GetCollectionIdOk returns a tuple with the CollectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionId

`func (o *UpdateDeviceRequest) SetCollectionId(v string)`

SetCollectionId sets CollectionId field to given value.

### HasCollectionId

`func (o *UpdateDeviceRequest) HasCollectionId() bool`

HasCollectionId returns a boolean if a field has been set.

### GetTags

`func (o *UpdateDeviceRequest) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpdateDeviceRequest) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UpdateDeviceRequest) SetTags(v map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *UpdateDeviceRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetFirmware

`func (o *UpdateDeviceRequest) GetFirmware() FirmwareMetadata`

GetFirmware returns the Firmware field if non-nil, zero value otherwise.

### GetFirmwareOk

`func (o *UpdateDeviceRequest) GetFirmwareOk() (*FirmwareMetadata, bool)`

GetFirmwareOk returns a tuple with the Firmware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmware

`func (o *UpdateDeviceRequest) SetFirmware(v FirmwareMetadata)`

SetFirmware sets Firmware field to given value.

### HasFirmware

`func (o *UpdateDeviceRequest) HasFirmware() bool`

HasFirmware returns a boolean if a field has been set.

### GetConfig

`func (o *UpdateDeviceRequest) GetConfig() DeviceConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *UpdateDeviceRequest) GetConfigOk() (*DeviceConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *UpdateDeviceRequest) SetConfig(v DeviceConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *UpdateDeviceRequest) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


