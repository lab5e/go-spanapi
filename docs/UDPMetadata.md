# UDPMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocalPort** | Pointer to **int32** |  | [optional] 
**RemotePort** | Pointer to **int32** |  | [optional] 

## Methods

### NewUDPMetadata

`func NewUDPMetadata() *UDPMetadata`

NewUDPMetadata instantiates a new UDPMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUDPMetadataWithDefaults

`func NewUDPMetadataWithDefaults() *UDPMetadata`

NewUDPMetadataWithDefaults instantiates a new UDPMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocalPort

`func (o *UDPMetadata) GetLocalPort() int32`

GetLocalPort returns the LocalPort field if non-nil, zero value otherwise.

### GetLocalPortOk

`func (o *UDPMetadata) GetLocalPortOk() (*int32, bool)`

GetLocalPortOk returns a tuple with the LocalPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalPort

`func (o *UDPMetadata) SetLocalPort(v int32)`

SetLocalPort sets LocalPort field to given value.

### HasLocalPort

`func (o *UDPMetadata) HasLocalPort() bool`

HasLocalPort returns a boolean if a field has been set.

### GetRemotePort

`func (o *UDPMetadata) GetRemotePort() int32`

GetRemotePort returns the RemotePort field if non-nil, zero value otherwise.

### GetRemotePortOk

`func (o *UDPMetadata) GetRemotePortOk() (*int32, bool)`

GetRemotePortOk returns a tuple with the RemotePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemotePort

`func (o *UDPMetadata) SetRemotePort(v int32)`

SetRemotePort sets RemotePort field to given value.

### HasRemotePort

`func (o *UDPMetadata) HasRemotePort() bool`

HasRemotePort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


