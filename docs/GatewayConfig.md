# GatewayConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ciot** | Pointer to [**GatewayCIoTConfig**](GatewayCIoTConfig.md) |  | [optional] 
**Inet** | Pointer to [**GatewayInetConfig**](GatewayInetConfig.md) |  | [optional] 
**User** | Pointer to [**GatewayCustomConfig**](GatewayCustomConfig.md) |  | [optional] 

## Methods

### NewGatewayConfig

`func NewGatewayConfig() *GatewayConfig`

NewGatewayConfig instantiates a new GatewayConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayConfigWithDefaults

`func NewGatewayConfigWithDefaults() *GatewayConfig`

NewGatewayConfigWithDefaults instantiates a new GatewayConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCiot

`func (o *GatewayConfig) GetCiot() GatewayCIoTConfig`

GetCiot returns the Ciot field if non-nil, zero value otherwise.

### GetCiotOk

`func (o *GatewayConfig) GetCiotOk() (*GatewayCIoTConfig, bool)`

GetCiotOk returns a tuple with the Ciot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiot

`func (o *GatewayConfig) SetCiot(v GatewayCIoTConfig)`

SetCiot sets Ciot field to given value.

### HasCiot

`func (o *GatewayConfig) HasCiot() bool`

HasCiot returns a boolean if a field has been set.

### GetInet

`func (o *GatewayConfig) GetInet() GatewayInetConfig`

GetInet returns the Inet field if non-nil, zero value otherwise.

### GetInetOk

`func (o *GatewayConfig) GetInetOk() (*GatewayInetConfig, bool)`

GetInetOk returns a tuple with the Inet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInet

`func (o *GatewayConfig) SetInet(v GatewayInetConfig)`

SetInet sets Inet field to given value.

### HasInet

`func (o *GatewayConfig) HasInet() bool`

HasInet returns a boolean if a field has been set.

### GetUser

`func (o *GatewayConfig) GetUser() GatewayCustomConfig`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *GatewayConfig) GetUserOk() (*GatewayCustomConfig, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *GatewayConfig) SetUser(v GatewayCustomConfig)`

SetUser sets User field to given value.

### HasUser

`func (o *GatewayConfig) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


