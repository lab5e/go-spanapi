# UpdateDeviceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExistingCollectionId** | Pointer to **string** |  | [optional] 
**DeviceId** | Pointer to **string** |  | [optional] 
**CollectionId** | Pointer to **string** | The collection id for the device. This field is optional and can be omitted if the collection id isn&#39;t changed. When changing to a new collection you must be an owner of the other collection, ie an administrator of the team that owns the new collection. | [optional] 
**Imsi** | Pointer to **string** | The IMSI is the unique ID for the (e|nu|whatever)SIM card on your device. This is the primary identifier for your device on the network. | [optional] 
**Imei** | Pointer to **string** | The IMEI number is the unique ID for your hardware as seen by the network. Obviously you might have a completely different view on things. | [optional] 
**Tags** | Pointer to **map[string]string** | Tags are metadata for the device that you can set. These are just strings. | [optional] 
**Firmware** | Pointer to [**FirmwareMetadata**](FirmwareMetadata.md) |  | [optional] 

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

### GetExistingCollectionId

`func (o *UpdateDeviceRequest) GetExistingCollectionId() string`

GetExistingCollectionId returns the ExistingCollectionId field if non-nil, zero value otherwise.

### GetExistingCollectionIdOk

`func (o *UpdateDeviceRequest) GetExistingCollectionIdOk() (*string, bool)`

GetExistingCollectionIdOk returns a tuple with the ExistingCollectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistingCollectionId

`func (o *UpdateDeviceRequest) SetExistingCollectionId(v string)`

SetExistingCollectionId sets ExistingCollectionId field to given value.

### HasExistingCollectionId

`func (o *UpdateDeviceRequest) HasExistingCollectionId() bool`

HasExistingCollectionId returns a boolean if a field has been set.

### GetDeviceId

`func (o *UpdateDeviceRequest) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *UpdateDeviceRequest) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *UpdateDeviceRequest) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *UpdateDeviceRequest) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

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

### GetImsi

`func (o *UpdateDeviceRequest) GetImsi() string`

GetImsi returns the Imsi field if non-nil, zero value otherwise.

### GetImsiOk

`func (o *UpdateDeviceRequest) GetImsiOk() (*string, bool)`

GetImsiOk returns a tuple with the Imsi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImsi

`func (o *UpdateDeviceRequest) SetImsi(v string)`

SetImsi sets Imsi field to given value.

### HasImsi

`func (o *UpdateDeviceRequest) HasImsi() bool`

HasImsi returns a boolean if a field has been set.

### GetImei

`func (o *UpdateDeviceRequest) GetImei() string`

GetImei returns the Imei field if non-nil, zero value otherwise.

### GetImeiOk

`func (o *UpdateDeviceRequest) GetImeiOk() (*string, bool)`

GetImeiOk returns a tuple with the Imei field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImei

`func (o *UpdateDeviceRequest) SetImei(v string)`

SetImei sets Imei field to given value.

### HasImei

`func (o *UpdateDeviceRequest) HasImei() bool`

HasImei returns a boolean if a field has been set.

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


