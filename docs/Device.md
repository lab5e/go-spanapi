# Device

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceId** | Pointer to **string** | The device ID is assigned by the backend. | [optional] 
**CollectionId** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **map[string]string** | Tags are metadata for the device that you can set. These are just strings. | [optional] 
**Firmware** | Pointer to [**FirmwareMetadata**](FirmwareMetadata.md) |  | [optional] 
**Config** | Pointer to [**DeviceConfig**](DeviceConfig.md) |  | [optional] 
**Metadata** | Pointer to [**DeviceMetadata**](DeviceMetadata.md) |  | [optional] 
**LastGatewayId** | Pointer to **string** |  | [optional] 
**LastTransport** | Pointer to [**MessageTransport**](MessageTransport.md) |  | [optional] [default to MESSAGETRANSPORT_UNSPECIFIED]
**LastReceived** | Pointer to **string** |  | [optional] 
**LastPayload** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewDevice

`func NewDevice() *Device`

NewDevice instantiates a new Device object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceWithDefaults

`func NewDeviceWithDefaults() *Device`

NewDeviceWithDefaults instantiates a new Device object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceId

`func (o *Device) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *Device) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *Device) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *Device) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetCollectionId

`func (o *Device) GetCollectionId() string`

GetCollectionId returns the CollectionId field if non-nil, zero value otherwise.

### GetCollectionIdOk

`func (o *Device) GetCollectionIdOk() (*string, bool)`

GetCollectionIdOk returns a tuple with the CollectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionId

`func (o *Device) SetCollectionId(v string)`

SetCollectionId sets CollectionId field to given value.

### HasCollectionId

`func (o *Device) HasCollectionId() bool`

HasCollectionId returns a boolean if a field has been set.

### GetTags

`func (o *Device) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Device) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Device) SetTags(v map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Device) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetFirmware

`func (o *Device) GetFirmware() FirmwareMetadata`

GetFirmware returns the Firmware field if non-nil, zero value otherwise.

### GetFirmwareOk

`func (o *Device) GetFirmwareOk() (*FirmwareMetadata, bool)`

GetFirmwareOk returns a tuple with the Firmware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmware

`func (o *Device) SetFirmware(v FirmwareMetadata)`

SetFirmware sets Firmware field to given value.

### HasFirmware

`func (o *Device) HasFirmware() bool`

HasFirmware returns a boolean if a field has been set.

### GetConfig

`func (o *Device) GetConfig() DeviceConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *Device) GetConfigOk() (*DeviceConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *Device) SetConfig(v DeviceConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *Device) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetMetadata

`func (o *Device) GetMetadata() DeviceMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Device) GetMetadataOk() (*DeviceMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Device) SetMetadata(v DeviceMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Device) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetLastGatewayId

`func (o *Device) GetLastGatewayId() string`

GetLastGatewayId returns the LastGatewayId field if non-nil, zero value otherwise.

### GetLastGatewayIdOk

`func (o *Device) GetLastGatewayIdOk() (*string, bool)`

GetLastGatewayIdOk returns a tuple with the LastGatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastGatewayId

`func (o *Device) SetLastGatewayId(v string)`

SetLastGatewayId sets LastGatewayId field to given value.

### HasLastGatewayId

`func (o *Device) HasLastGatewayId() bool`

HasLastGatewayId returns a boolean if a field has been set.

### GetLastTransport

`func (o *Device) GetLastTransport() MessageTransport`

GetLastTransport returns the LastTransport field if non-nil, zero value otherwise.

### GetLastTransportOk

`func (o *Device) GetLastTransportOk() (*MessageTransport, bool)`

GetLastTransportOk returns a tuple with the LastTransport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTransport

`func (o *Device) SetLastTransport(v MessageTransport)`

SetLastTransport sets LastTransport field to given value.

### HasLastTransport

`func (o *Device) HasLastTransport() bool`

HasLastTransport returns a boolean if a field has been set.

### GetLastReceived

`func (o *Device) GetLastReceived() string`

GetLastReceived returns the LastReceived field if non-nil, zero value otherwise.

### GetLastReceivedOk

`func (o *Device) GetLastReceivedOk() (*string, bool)`

GetLastReceivedOk returns a tuple with the LastReceived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReceived

`func (o *Device) SetLastReceived(v string)`

SetLastReceived sets LastReceived field to given value.

### HasLastReceived

`func (o *Device) HasLastReceived() bool`

HasLastReceived returns a boolean if a field has been set.

### GetLastPayload

`func (o *Device) GetLastPayload() string`

GetLastPayload returns the LastPayload field if non-nil, zero value otherwise.

### GetLastPayloadOk

`func (o *Device) GetLastPayloadOk() (*string, bool)`

GetLastPayloadOk returns a tuple with the LastPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPayload

`func (o *Device) SetLastPayload(v string)`

SetLastPayload sets LastPayload field to given value.

### HasLastPayload

`func (o *Device) HasLastPayload() bool`

HasLastPayload returns a boolean if a field has been set.

### GetEnabled

`func (o *Device) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *Device) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *Device) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *Device) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


