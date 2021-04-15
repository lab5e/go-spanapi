# CollectionFirmware

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentFirmwareId** | Pointer to **string** | The current firmware is the firmware that the devices are currently using. | [optional] 
**TargetFirmwareId** | Pointer to **string** | The target firmware is set to the desired firmware image for the devices in this collection. If the management is set to \&quot;device\&quot; this will only be used if the target firmware isn&#39;t set on the device itself. | [optional] 
**Management** | Pointer to [**CollectionFirmwareFirmwareManagement**](CollectionFirmwareFirmwareManagement.md) |  | [optional] [default to UNSPECIFIED]

## Methods

### NewCollectionFirmware

`func NewCollectionFirmware() *CollectionFirmware`

NewCollectionFirmware instantiates a new CollectionFirmware object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionFirmwareWithDefaults

`func NewCollectionFirmwareWithDefaults() *CollectionFirmware`

NewCollectionFirmwareWithDefaults instantiates a new CollectionFirmware object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentFirmwareId

`func (o *CollectionFirmware) GetCurrentFirmwareId() string`

GetCurrentFirmwareId returns the CurrentFirmwareId field if non-nil, zero value otherwise.

### GetCurrentFirmwareIdOk

`func (o *CollectionFirmware) GetCurrentFirmwareIdOk() (*string, bool)`

GetCurrentFirmwareIdOk returns a tuple with the CurrentFirmwareId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentFirmwareId

`func (o *CollectionFirmware) SetCurrentFirmwareId(v string)`

SetCurrentFirmwareId sets CurrentFirmwareId field to given value.

### HasCurrentFirmwareId

`func (o *CollectionFirmware) HasCurrentFirmwareId() bool`

HasCurrentFirmwareId returns a boolean if a field has been set.

### GetTargetFirmwareId

`func (o *CollectionFirmware) GetTargetFirmwareId() string`

GetTargetFirmwareId returns the TargetFirmwareId field if non-nil, zero value otherwise.

### GetTargetFirmwareIdOk

`func (o *CollectionFirmware) GetTargetFirmwareIdOk() (*string, bool)`

GetTargetFirmwareIdOk returns a tuple with the TargetFirmwareId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetFirmwareId

`func (o *CollectionFirmware) SetTargetFirmwareId(v string)`

SetTargetFirmwareId sets TargetFirmwareId field to given value.

### HasTargetFirmwareId

`func (o *CollectionFirmware) HasTargetFirmwareId() bool`

HasTargetFirmwareId returns a boolean if a field has been set.

### GetManagement

`func (o *CollectionFirmware) GetManagement() CollectionFirmwareFirmwareManagement`

GetManagement returns the Management field if non-nil, zero value otherwise.

### GetManagementOk

`func (o *CollectionFirmware) GetManagementOk() (*CollectionFirmwareFirmwareManagement, bool)`

GetManagementOk returns a tuple with the Management field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagement

`func (o *CollectionFirmware) SetManagement(v CollectionFirmwareFirmwareManagement)`

SetManagement sets Management field to given value.

### HasManagement

`func (o *CollectionFirmware) HasManagement() bool`

HasManagement returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


