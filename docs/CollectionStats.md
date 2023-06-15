# CollectionStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceCount** | Pointer to **int32** |  | [optional] 
**OutputCount** | Pointer to **int32** |  | [optional] 
**FirmwareCount** | Pointer to **int32** |  | [optional] 
**BlobCount** | Pointer to **int32** |  | [optional] 
**GatewayCount** | Pointer to **int32** |  | [optional] 
**Devices** | Pointer to [**DeviceStats**](DeviceStats.md) |  | [optional] 
**Outputs** | Pointer to [**OutputStats**](OutputStats.md) |  | [optional] 
**Firmware** | Pointer to [**FirmwareStats**](FirmwareStats.md) |  | [optional] 
**Blobs** | Pointer to [**BlobStats**](BlobStats.md) |  | [optional] 
**Gateways** | Pointer to [**GatewayStats**](GatewayStats.md) |  | [optional] 

## Methods

### NewCollectionStats

`func NewCollectionStats() *CollectionStats`

NewCollectionStats instantiates a new CollectionStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionStatsWithDefaults

`func NewCollectionStatsWithDefaults() *CollectionStats`

NewCollectionStatsWithDefaults instantiates a new CollectionStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceCount

`func (o *CollectionStats) GetDeviceCount() int32`

GetDeviceCount returns the DeviceCount field if non-nil, zero value otherwise.

### GetDeviceCountOk

`func (o *CollectionStats) GetDeviceCountOk() (*int32, bool)`

GetDeviceCountOk returns a tuple with the DeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceCount

`func (o *CollectionStats) SetDeviceCount(v int32)`

SetDeviceCount sets DeviceCount field to given value.

### HasDeviceCount

`func (o *CollectionStats) HasDeviceCount() bool`

HasDeviceCount returns a boolean if a field has been set.

### GetOutputCount

`func (o *CollectionStats) GetOutputCount() int32`

GetOutputCount returns the OutputCount field if non-nil, zero value otherwise.

### GetOutputCountOk

`func (o *CollectionStats) GetOutputCountOk() (*int32, bool)`

GetOutputCountOk returns a tuple with the OutputCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputCount

`func (o *CollectionStats) SetOutputCount(v int32)`

SetOutputCount sets OutputCount field to given value.

### HasOutputCount

`func (o *CollectionStats) HasOutputCount() bool`

HasOutputCount returns a boolean if a field has been set.

### GetFirmwareCount

`func (o *CollectionStats) GetFirmwareCount() int32`

GetFirmwareCount returns the FirmwareCount field if non-nil, zero value otherwise.

### GetFirmwareCountOk

`func (o *CollectionStats) GetFirmwareCountOk() (*int32, bool)`

GetFirmwareCountOk returns a tuple with the FirmwareCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmwareCount

`func (o *CollectionStats) SetFirmwareCount(v int32)`

SetFirmwareCount sets FirmwareCount field to given value.

### HasFirmwareCount

`func (o *CollectionStats) HasFirmwareCount() bool`

HasFirmwareCount returns a boolean if a field has been set.

### GetBlobCount

`func (o *CollectionStats) GetBlobCount() int32`

GetBlobCount returns the BlobCount field if non-nil, zero value otherwise.

### GetBlobCountOk

`func (o *CollectionStats) GetBlobCountOk() (*int32, bool)`

GetBlobCountOk returns a tuple with the BlobCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlobCount

`func (o *CollectionStats) SetBlobCount(v int32)`

SetBlobCount sets BlobCount field to given value.

### HasBlobCount

`func (o *CollectionStats) HasBlobCount() bool`

HasBlobCount returns a boolean if a field has been set.

### GetGatewayCount

`func (o *CollectionStats) GetGatewayCount() int32`

GetGatewayCount returns the GatewayCount field if non-nil, zero value otherwise.

### GetGatewayCountOk

`func (o *CollectionStats) GetGatewayCountOk() (*int32, bool)`

GetGatewayCountOk returns a tuple with the GatewayCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayCount

`func (o *CollectionStats) SetGatewayCount(v int32)`

SetGatewayCount sets GatewayCount field to given value.

### HasGatewayCount

`func (o *CollectionStats) HasGatewayCount() bool`

HasGatewayCount returns a boolean if a field has been set.

### GetDevices

`func (o *CollectionStats) GetDevices() DeviceStats`

GetDevices returns the Devices field if non-nil, zero value otherwise.

### GetDevicesOk

`func (o *CollectionStats) GetDevicesOk() (*DeviceStats, bool)`

GetDevicesOk returns a tuple with the Devices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevices

`func (o *CollectionStats) SetDevices(v DeviceStats)`

SetDevices sets Devices field to given value.

### HasDevices

`func (o *CollectionStats) HasDevices() bool`

HasDevices returns a boolean if a field has been set.

### GetOutputs

`func (o *CollectionStats) GetOutputs() OutputStats`

GetOutputs returns the Outputs field if non-nil, zero value otherwise.

### GetOutputsOk

`func (o *CollectionStats) GetOutputsOk() (*OutputStats, bool)`

GetOutputsOk returns a tuple with the Outputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputs

`func (o *CollectionStats) SetOutputs(v OutputStats)`

SetOutputs sets Outputs field to given value.

### HasOutputs

`func (o *CollectionStats) HasOutputs() bool`

HasOutputs returns a boolean if a field has been set.

### GetFirmware

`func (o *CollectionStats) GetFirmware() FirmwareStats`

GetFirmware returns the Firmware field if non-nil, zero value otherwise.

### GetFirmwareOk

`func (o *CollectionStats) GetFirmwareOk() (*FirmwareStats, bool)`

GetFirmwareOk returns a tuple with the Firmware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmware

`func (o *CollectionStats) SetFirmware(v FirmwareStats)`

SetFirmware sets Firmware field to given value.

### HasFirmware

`func (o *CollectionStats) HasFirmware() bool`

HasFirmware returns a boolean if a field has been set.

### GetBlobs

`func (o *CollectionStats) GetBlobs() BlobStats`

GetBlobs returns the Blobs field if non-nil, zero value otherwise.

### GetBlobsOk

`func (o *CollectionStats) GetBlobsOk() (*BlobStats, bool)`

GetBlobsOk returns a tuple with the Blobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlobs

`func (o *CollectionStats) SetBlobs(v BlobStats)`

SetBlobs sets Blobs field to given value.

### HasBlobs

`func (o *CollectionStats) HasBlobs() bool`

HasBlobs returns a boolean if a field has been set.

### GetGateways

`func (o *CollectionStats) GetGateways() GatewayStats`

GetGateways returns the Gateways field if non-nil, zero value otherwise.

### GetGatewaysOk

`func (o *CollectionStats) GetGatewaysOk() (*GatewayStats, bool)`

GetGatewaysOk returns a tuple with the Gateways field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateways

`func (o *CollectionStats) SetGateways(v GatewayStats)`

SetGateways sets Gateways field to given value.

### HasGateways

`func (o *CollectionStats) HasGateways() bool`

HasGateways returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


