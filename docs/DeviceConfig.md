# DeviceConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ciot** | Pointer to [**CellularIoTConfig**](CellularIoTConfig.md) |  | [optional] 

## Methods

### NewDeviceConfig

`func NewDeviceConfig() *DeviceConfig`

NewDeviceConfig instantiates a new DeviceConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceConfigWithDefaults

`func NewDeviceConfigWithDefaults() *DeviceConfig`

NewDeviceConfigWithDefaults instantiates a new DeviceConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCiot

`func (o *DeviceConfig) GetCiot() CellularIoTConfig`

GetCiot returns the Ciot field if non-nil, zero value otherwise.

### GetCiotOk

`func (o *DeviceConfig) GetCiotOk() (*CellularIoTConfig, bool)`

GetCiotOk returns a tuple with the Ciot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiot

`func (o *DeviceConfig) SetCiot(v CellularIoTConfig)`

SetCiot sets Ciot field to given value.

### HasCiot

`func (o *DeviceConfig) HasCiot() bool`

HasCiot returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


