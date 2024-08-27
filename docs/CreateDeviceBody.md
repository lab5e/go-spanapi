# CreateDeviceBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tags** | Pointer to **map[string]string** | Tags are metadata for the device that you can set. These are just strings. | [optional] 
**Firmware** | Pointer to [**FirmwareMetadata**](FirmwareMetadata.md) |  | [optional] 
**Config** | Pointer to [**DeviceConfig**](DeviceConfig.md) |  | [optional] 
**Metadata** | Pointer to [**DeviceMetadata**](DeviceMetadata.md) |  | [optional] 

## Methods

### NewCreateDeviceBody

`func NewCreateDeviceBody() *CreateDeviceBody`

NewCreateDeviceBody instantiates a new CreateDeviceBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDeviceBodyWithDefaults

`func NewCreateDeviceBodyWithDefaults() *CreateDeviceBody`

NewCreateDeviceBodyWithDefaults instantiates a new CreateDeviceBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTags

`func (o *CreateDeviceBody) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateDeviceBody) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateDeviceBody) SetTags(v map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateDeviceBody) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetFirmware

`func (o *CreateDeviceBody) GetFirmware() FirmwareMetadata`

GetFirmware returns the Firmware field if non-nil, zero value otherwise.

### GetFirmwareOk

`func (o *CreateDeviceBody) GetFirmwareOk() (*FirmwareMetadata, bool)`

GetFirmwareOk returns a tuple with the Firmware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmware

`func (o *CreateDeviceBody) SetFirmware(v FirmwareMetadata)`

SetFirmware sets Firmware field to given value.

### HasFirmware

`func (o *CreateDeviceBody) HasFirmware() bool`

HasFirmware returns a boolean if a field has been set.

### GetConfig

`func (o *CreateDeviceBody) GetConfig() DeviceConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *CreateDeviceBody) GetConfigOk() (*DeviceConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *CreateDeviceBody) SetConfig(v DeviceConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *CreateDeviceBody) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetMetadata

`func (o *CreateDeviceBody) GetMetadata() DeviceMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreateDeviceBody) GetMetadataOk() (*DeviceMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreateDeviceBody) SetMetadata(v DeviceMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreateDeviceBody) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


