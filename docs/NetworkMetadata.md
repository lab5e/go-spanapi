# NetworkMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllocatedIp** | Pointer to **string** | Allocated IP address. | [optional] 
**AllocatedAt** | Pointer to **string** |  | [optional] 
**CellId** | Pointer to **string** | Cell ID for device. This might not be set, depending on the provider in use. | [optional] 

## Methods

### NewNetworkMetadata

`func NewNetworkMetadata() *NetworkMetadata`

NewNetworkMetadata instantiates a new NetworkMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkMetadataWithDefaults

`func NewNetworkMetadataWithDefaults() *NetworkMetadata`

NewNetworkMetadataWithDefaults instantiates a new NetworkMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllocatedIp

`func (o *NetworkMetadata) GetAllocatedIp() string`

GetAllocatedIp returns the AllocatedIp field if non-nil, zero value otherwise.

### GetAllocatedIpOk

`func (o *NetworkMetadata) GetAllocatedIpOk() (*string, bool)`

GetAllocatedIpOk returns a tuple with the AllocatedIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocatedIp

`func (o *NetworkMetadata) SetAllocatedIp(v string)`

SetAllocatedIp sets AllocatedIp field to given value.

### HasAllocatedIp

`func (o *NetworkMetadata) HasAllocatedIp() bool`

HasAllocatedIp returns a boolean if a field has been set.

### GetAllocatedAt

`func (o *NetworkMetadata) GetAllocatedAt() string`

GetAllocatedAt returns the AllocatedAt field if non-nil, zero value otherwise.

### GetAllocatedAtOk

`func (o *NetworkMetadata) GetAllocatedAtOk() (*string, bool)`

GetAllocatedAtOk returns a tuple with the AllocatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocatedAt

`func (o *NetworkMetadata) SetAllocatedAt(v string)`

SetAllocatedAt sets AllocatedAt field to given value.

### HasAllocatedAt

`func (o *NetworkMetadata) HasAllocatedAt() bool`

HasAllocatedAt returns a boolean if a field has been set.

### GetCellId

`func (o *NetworkMetadata) GetCellId() string`

GetCellId returns the CellId field if non-nil, zero value otherwise.

### GetCellIdOk

`func (o *NetworkMetadata) GetCellIdOk() (*string, bool)`

GetCellIdOk returns a tuple with the CellId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCellId

`func (o *NetworkMetadata) SetCellId(v string)`

SetCellId sets CellId field to given value.

### HasCellId

`func (o *NetworkMetadata) HasCellId() bool`

HasCellId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


