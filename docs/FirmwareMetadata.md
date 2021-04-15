# FirmwareMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentFirmwareId** | Pointer to **string** |  | [optional] 
**TargetFirmwareId** | Pointer to **string** |  | [optional] 
**FirmwareVersion** | Pointer to **string** | Last reported firmware version. | [optional] 
**SerialNumber** | Pointer to **string** |  | [optional] 
**ModelNumber** | Pointer to **string** |  | [optional] 
**Manufacturer** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** | State of the firmware. | [optional] 
**StateMessage** | Pointer to **string** |  | [optional] 

## Methods

### NewFirmwareMetadata

`func NewFirmwareMetadata() *FirmwareMetadata`

NewFirmwareMetadata instantiates a new FirmwareMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirmwareMetadataWithDefaults

`func NewFirmwareMetadataWithDefaults() *FirmwareMetadata`

NewFirmwareMetadataWithDefaults instantiates a new FirmwareMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentFirmwareId

`func (o *FirmwareMetadata) GetCurrentFirmwareId() string`

GetCurrentFirmwareId returns the CurrentFirmwareId field if non-nil, zero value otherwise.

### GetCurrentFirmwareIdOk

`func (o *FirmwareMetadata) GetCurrentFirmwareIdOk() (*string, bool)`

GetCurrentFirmwareIdOk returns a tuple with the CurrentFirmwareId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentFirmwareId

`func (o *FirmwareMetadata) SetCurrentFirmwareId(v string)`

SetCurrentFirmwareId sets CurrentFirmwareId field to given value.

### HasCurrentFirmwareId

`func (o *FirmwareMetadata) HasCurrentFirmwareId() bool`

HasCurrentFirmwareId returns a boolean if a field has been set.

### GetTargetFirmwareId

`func (o *FirmwareMetadata) GetTargetFirmwareId() string`

GetTargetFirmwareId returns the TargetFirmwareId field if non-nil, zero value otherwise.

### GetTargetFirmwareIdOk

`func (o *FirmwareMetadata) GetTargetFirmwareIdOk() (*string, bool)`

GetTargetFirmwareIdOk returns a tuple with the TargetFirmwareId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetFirmwareId

`func (o *FirmwareMetadata) SetTargetFirmwareId(v string)`

SetTargetFirmwareId sets TargetFirmwareId field to given value.

### HasTargetFirmwareId

`func (o *FirmwareMetadata) HasTargetFirmwareId() bool`

HasTargetFirmwareId returns a boolean if a field has been set.

### GetFirmwareVersion

`func (o *FirmwareMetadata) GetFirmwareVersion() string`

GetFirmwareVersion returns the FirmwareVersion field if non-nil, zero value otherwise.

### GetFirmwareVersionOk

`func (o *FirmwareMetadata) GetFirmwareVersionOk() (*string, bool)`

GetFirmwareVersionOk returns a tuple with the FirmwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmwareVersion

`func (o *FirmwareMetadata) SetFirmwareVersion(v string)`

SetFirmwareVersion sets FirmwareVersion field to given value.

### HasFirmwareVersion

`func (o *FirmwareMetadata) HasFirmwareVersion() bool`

HasFirmwareVersion returns a boolean if a field has been set.

### GetSerialNumber

`func (o *FirmwareMetadata) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *FirmwareMetadata) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *FirmwareMetadata) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *FirmwareMetadata) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetModelNumber

`func (o *FirmwareMetadata) GetModelNumber() string`

GetModelNumber returns the ModelNumber field if non-nil, zero value otherwise.

### GetModelNumberOk

`func (o *FirmwareMetadata) GetModelNumberOk() (*string, bool)`

GetModelNumberOk returns a tuple with the ModelNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelNumber

`func (o *FirmwareMetadata) SetModelNumber(v string)`

SetModelNumber sets ModelNumber field to given value.

### HasModelNumber

`func (o *FirmwareMetadata) HasModelNumber() bool`

HasModelNumber returns a boolean if a field has been set.

### GetManufacturer

`func (o *FirmwareMetadata) GetManufacturer() string`

GetManufacturer returns the Manufacturer field if non-nil, zero value otherwise.

### GetManufacturerOk

`func (o *FirmwareMetadata) GetManufacturerOk() (*string, bool)`

GetManufacturerOk returns a tuple with the Manufacturer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturer

`func (o *FirmwareMetadata) SetManufacturer(v string)`

SetManufacturer sets Manufacturer field to given value.

### HasManufacturer

`func (o *FirmwareMetadata) HasManufacturer() bool`

HasManufacturer returns a boolean if a field has been set.

### GetState

`func (o *FirmwareMetadata) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *FirmwareMetadata) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *FirmwareMetadata) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *FirmwareMetadata) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStateMessage

`func (o *FirmwareMetadata) GetStateMessage() string`

GetStateMessage returns the StateMessage field if non-nil, zero value otherwise.

### GetStateMessageOk

`func (o *FirmwareMetadata) GetStateMessageOk() (*string, bool)`

GetStateMessageOk returns a tuple with the StateMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateMessage

`func (o *FirmwareMetadata) SetStateMessage(v string)`

SetStateMessage sets StateMessage field to given value.

### HasStateMessage

`func (o *FirmwareMetadata) HasStateMessage() bool`

HasStateMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


