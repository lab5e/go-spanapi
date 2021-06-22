# Device

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceId** | Pointer to **string** | The device ID is assigned by the backend. | [optional] 
**CollectionId** | Pointer to **string** |  | [optional] 
**Imsi** | Pointer to **string** | The IMSI is the unique ID for the (e|nu|whatever)SIM card on your device. This is the primary identifier for your device on the network. | [optional] 
**Imei** | Pointer to **string** | The IMEI number is the unique ID for your hardware as seen by the network. Obviously you might have a completely different view on things. | [optional] 
**Tags** | Pointer to **map[string]string** | Tags are metadata for the device that you can set. These are just strings. | [optional] 
**Network** | Pointer to [**NetworkMetadata**](NetworkMetadata.md) |  | [optional] 
**Firmware** | Pointer to [**FirmwareMetadata**](FirmwareMetadata.md) |  | [optional] 
**Metadata** | Pointer to [**DeviceMetadata**](DeviceMetadata.md) |  | [optional] 

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

### GetImsi

`func (o *Device) GetImsi() string`

GetImsi returns the Imsi field if non-nil, zero value otherwise.

### GetImsiOk

`func (o *Device) GetImsiOk() (*string, bool)`

GetImsiOk returns a tuple with the Imsi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImsi

`func (o *Device) SetImsi(v string)`

SetImsi sets Imsi field to given value.

### HasImsi

`func (o *Device) HasImsi() bool`

HasImsi returns a boolean if a field has been set.

### GetImei

`func (o *Device) GetImei() string`

GetImei returns the Imei field if non-nil, zero value otherwise.

### GetImeiOk

`func (o *Device) GetImeiOk() (*string, bool)`

GetImeiOk returns a tuple with the Imei field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImei

`func (o *Device) SetImei(v string)`

SetImei sets Imei field to given value.

### HasImei

`func (o *Device) HasImei() bool`

HasImei returns a boolean if a field has been set.

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

### GetNetwork

`func (o *Device) GetNetwork() NetworkMetadata`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *Device) GetNetworkOk() (*NetworkMetadata, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *Device) SetNetwork(v NetworkMetadata)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *Device) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


