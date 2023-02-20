# CreateDeviceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tags** | Pointer to **map[string]string** | Tags are metadata for the device that you can set. These are just strings. | [optional] 
**Firmware** | Pointer to [**FirmwareMetadata**](FirmwareMetadata.md) |  | [optional] 
**Config** | Pointer to [**DeviceConfig**](DeviceConfig.md) |  | [optional] 
**Metadata** | Pointer to [**DeviceMetadata**](DeviceMetadata.md) |  | [optional] 
**Imsi** | Pointer to **string** | Deprecated: The IMSI is replaced by CellularIoTMetadata | [optional] 
**Imei** | Pointer to **string** | The IMEI number is the unique ID for your hardware as seen by the network. Obviously you might have a completely different view on things. This field is deprecated.  Deprecated: The IMEI is replaced by CellularIoTMetadata | [optional] 
**Network** | Pointer to [**NetworkMetadata**](NetworkMetadata.md) |  | [optional] 

## Methods

### NewCreateDeviceRequest

`func NewCreateDeviceRequest() *CreateDeviceRequest`

NewCreateDeviceRequest instantiates a new CreateDeviceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDeviceRequestWithDefaults

`func NewCreateDeviceRequestWithDefaults() *CreateDeviceRequest`

NewCreateDeviceRequestWithDefaults instantiates a new CreateDeviceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTags

`func (o *CreateDeviceRequest) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateDeviceRequest) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateDeviceRequest) SetTags(v map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateDeviceRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetFirmware

`func (o *CreateDeviceRequest) GetFirmware() FirmwareMetadata`

GetFirmware returns the Firmware field if non-nil, zero value otherwise.

### GetFirmwareOk

`func (o *CreateDeviceRequest) GetFirmwareOk() (*FirmwareMetadata, bool)`

GetFirmwareOk returns a tuple with the Firmware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmware

`func (o *CreateDeviceRequest) SetFirmware(v FirmwareMetadata)`

SetFirmware sets Firmware field to given value.

### HasFirmware

`func (o *CreateDeviceRequest) HasFirmware() bool`

HasFirmware returns a boolean if a field has been set.

### GetConfig

`func (o *CreateDeviceRequest) GetConfig() DeviceConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *CreateDeviceRequest) GetConfigOk() (*DeviceConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *CreateDeviceRequest) SetConfig(v DeviceConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *CreateDeviceRequest) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetMetadata

`func (o *CreateDeviceRequest) GetMetadata() DeviceMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreateDeviceRequest) GetMetadataOk() (*DeviceMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreateDeviceRequest) SetMetadata(v DeviceMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreateDeviceRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetImsi

`func (o *CreateDeviceRequest) GetImsi() string`

GetImsi returns the Imsi field if non-nil, zero value otherwise.

### GetImsiOk

`func (o *CreateDeviceRequest) GetImsiOk() (*string, bool)`

GetImsiOk returns a tuple with the Imsi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImsi

`func (o *CreateDeviceRequest) SetImsi(v string)`

SetImsi sets Imsi field to given value.

### HasImsi

`func (o *CreateDeviceRequest) HasImsi() bool`

HasImsi returns a boolean if a field has been set.

### GetImei

`func (o *CreateDeviceRequest) GetImei() string`

GetImei returns the Imei field if non-nil, zero value otherwise.

### GetImeiOk

`func (o *CreateDeviceRequest) GetImeiOk() (*string, bool)`

GetImeiOk returns a tuple with the Imei field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImei

`func (o *CreateDeviceRequest) SetImei(v string)`

SetImei sets Imei field to given value.

### HasImei

`func (o *CreateDeviceRequest) HasImei() bool`

HasImei returns a boolean if a field has been set.

### GetNetwork

`func (o *CreateDeviceRequest) GetNetwork() NetworkMetadata`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *CreateDeviceRequest) GetNetworkOk() (*NetworkMetadata, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *CreateDeviceRequest) SetNetwork(v NetworkMetadata)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *CreateDeviceRequest) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


