# DeviceConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ciot** | Pointer to [**CellularIoTConfig**](CellularIoTConfig.md) |  | [optional] 
**Inet** | Pointer to **map[string]interface{}** | This is the configuration for an internet-connected device. There are currently no configuration options for internet devices; the device is identified via the clientcertificate.  This is empty since there&#39;s no configuration required for the internet  gateway | [optional] 
**Gateway** | Pointer to [**map[string]GatewayDeviceConfig**](GatewayDeviceConfig.md) |  | [optional] 

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

### GetInet

`func (o *DeviceConfig) GetInet() map[string]interface{}`

GetInet returns the Inet field if non-nil, zero value otherwise.

### GetInetOk

`func (o *DeviceConfig) GetInetOk() (*map[string]interface{}, bool)`

GetInetOk returns a tuple with the Inet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInet

`func (o *DeviceConfig) SetInet(v map[string]interface{})`

SetInet sets Inet field to given value.

### HasInet

`func (o *DeviceConfig) HasInet() bool`

HasInet returns a boolean if a field has been set.

### GetGateway

`func (o *DeviceConfig) GetGateway() map[string]GatewayDeviceConfig`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *DeviceConfig) GetGatewayOk() (*map[string]GatewayDeviceConfig, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *DeviceConfig) SetGateway(v map[string]GatewayDeviceConfig)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *DeviceConfig) HasGateway() bool`

HasGateway returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


