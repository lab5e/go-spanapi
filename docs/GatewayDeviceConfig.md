# GatewayDeviceConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GatewayId** | Pointer to **string** | This is the ID of the gateway this configuration applies to. | [optional] 
**Params** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewGatewayDeviceConfig

`func NewGatewayDeviceConfig() *GatewayDeviceConfig`

NewGatewayDeviceConfig instantiates a new GatewayDeviceConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayDeviceConfigWithDefaults

`func NewGatewayDeviceConfigWithDefaults() *GatewayDeviceConfig`

NewGatewayDeviceConfigWithDefaults instantiates a new GatewayDeviceConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGatewayId

`func (o *GatewayDeviceConfig) GetGatewayId() string`

GetGatewayId returns the GatewayId field if non-nil, zero value otherwise.

### GetGatewayIdOk

`func (o *GatewayDeviceConfig) GetGatewayIdOk() (*string, bool)`

GetGatewayIdOk returns a tuple with the GatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayId

`func (o *GatewayDeviceConfig) SetGatewayId(v string)`

SetGatewayId sets GatewayId field to given value.

### HasGatewayId

`func (o *GatewayDeviceConfig) HasGatewayId() bool`

HasGatewayId returns a boolean if a field has been set.

### GetParams

`func (o *GatewayDeviceConfig) GetParams() map[string]string`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *GatewayDeviceConfig) GetParamsOk() (*map[string]string, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *GatewayDeviceConfig) SetParams(v map[string]string)`

SetParams sets Params field to given value.

### HasParams

`func (o *GatewayDeviceConfig) HasParams() bool`

HasParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


