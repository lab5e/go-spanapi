# DeviceMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ciot** | Pointer to [**CellularIoTMetadata**](CellularIoTMetadata.md) |  | [optional] 
**Inet** | Pointer to [**InetMetadata**](InetMetadata.md) |  | [optional] 
**Gateway** | Pointer to [**GatewayDeviceMetadata**](GatewayDeviceMetadata.md) |  | [optional] 

## Methods

### NewDeviceMetadata

`func NewDeviceMetadata() *DeviceMetadata`

NewDeviceMetadata instantiates a new DeviceMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceMetadataWithDefaults

`func NewDeviceMetadataWithDefaults() *DeviceMetadata`

NewDeviceMetadataWithDefaults instantiates a new DeviceMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCiot

`func (o *DeviceMetadata) GetCiot() CellularIoTMetadata`

GetCiot returns the Ciot field if non-nil, zero value otherwise.

### GetCiotOk

`func (o *DeviceMetadata) GetCiotOk() (*CellularIoTMetadata, bool)`

GetCiotOk returns a tuple with the Ciot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiot

`func (o *DeviceMetadata) SetCiot(v CellularIoTMetadata)`

SetCiot sets Ciot field to given value.

### HasCiot

`func (o *DeviceMetadata) HasCiot() bool`

HasCiot returns a boolean if a field has been set.

### GetInet

`func (o *DeviceMetadata) GetInet() InetMetadata`

GetInet returns the Inet field if non-nil, zero value otherwise.

### GetInetOk

`func (o *DeviceMetadata) GetInetOk() (*InetMetadata, bool)`

GetInetOk returns a tuple with the Inet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInet

`func (o *DeviceMetadata) SetInet(v InetMetadata)`

SetInet sets Inet field to given value.

### HasInet

`func (o *DeviceMetadata) HasInet() bool`

HasInet returns a boolean if a field has been set.

### GetGateway

`func (o *DeviceMetadata) GetGateway() GatewayDeviceMetadata`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *DeviceMetadata) GetGatewayOk() (*GatewayDeviceMetadata, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *DeviceMetadata) SetGateway(v GatewayDeviceMetadata)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *DeviceMetadata) HasGateway() bool`

HasGateway returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


