# CellularIoTMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GatewayId** | Pointer to **string** |  | [optional] 
**AllocatedIp** | Pointer to **string** | Allocated IP address. | [optional] 
**AllocatedAt** | Pointer to **string** |  | [optional] 
**CellId** | Pointer to **string** |  | [optional] 
**Mcc** | Pointer to **int32** | the provider in use.  The Mobile Country Code for the operator. | [optional] 
**Mnc** | Pointer to **int32** |  | [optional] 
**Country** | Pointer to **string** |  | [optional] 
**Network** | Pointer to **string** |  | [optional] 
**CountryCode** | Pointer to **string** |  | [optional] 
**LastUpdate** | Pointer to **string** |  | [optional] 
**LastImsi** | Pointer to **string** |  | [optional] 
**LastImei** | Pointer to **string** |  | [optional] 

## Methods

### NewCellularIoTMetadata

`func NewCellularIoTMetadata() *CellularIoTMetadata`

NewCellularIoTMetadata instantiates a new CellularIoTMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCellularIoTMetadataWithDefaults

`func NewCellularIoTMetadataWithDefaults() *CellularIoTMetadata`

NewCellularIoTMetadataWithDefaults instantiates a new CellularIoTMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGatewayId

`func (o *CellularIoTMetadata) GetGatewayId() string`

GetGatewayId returns the GatewayId field if non-nil, zero value otherwise.

### GetGatewayIdOk

`func (o *CellularIoTMetadata) GetGatewayIdOk() (*string, bool)`

GetGatewayIdOk returns a tuple with the GatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayId

`func (o *CellularIoTMetadata) SetGatewayId(v string)`

SetGatewayId sets GatewayId field to given value.

### HasGatewayId

`func (o *CellularIoTMetadata) HasGatewayId() bool`

HasGatewayId returns a boolean if a field has been set.

### GetAllocatedIp

`func (o *CellularIoTMetadata) GetAllocatedIp() string`

GetAllocatedIp returns the AllocatedIp field if non-nil, zero value otherwise.

### GetAllocatedIpOk

`func (o *CellularIoTMetadata) GetAllocatedIpOk() (*string, bool)`

GetAllocatedIpOk returns a tuple with the AllocatedIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocatedIp

`func (o *CellularIoTMetadata) SetAllocatedIp(v string)`

SetAllocatedIp sets AllocatedIp field to given value.

### HasAllocatedIp

`func (o *CellularIoTMetadata) HasAllocatedIp() bool`

HasAllocatedIp returns a boolean if a field has been set.

### GetAllocatedAt

`func (o *CellularIoTMetadata) GetAllocatedAt() string`

GetAllocatedAt returns the AllocatedAt field if non-nil, zero value otherwise.

### GetAllocatedAtOk

`func (o *CellularIoTMetadata) GetAllocatedAtOk() (*string, bool)`

GetAllocatedAtOk returns a tuple with the AllocatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocatedAt

`func (o *CellularIoTMetadata) SetAllocatedAt(v string)`

SetAllocatedAt sets AllocatedAt field to given value.

### HasAllocatedAt

`func (o *CellularIoTMetadata) HasAllocatedAt() bool`

HasAllocatedAt returns a boolean if a field has been set.

### GetCellId

`func (o *CellularIoTMetadata) GetCellId() string`

GetCellId returns the CellId field if non-nil, zero value otherwise.

### GetCellIdOk

`func (o *CellularIoTMetadata) GetCellIdOk() (*string, bool)`

GetCellIdOk returns a tuple with the CellId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCellId

`func (o *CellularIoTMetadata) SetCellId(v string)`

SetCellId sets CellId field to given value.

### HasCellId

`func (o *CellularIoTMetadata) HasCellId() bool`

HasCellId returns a boolean if a field has been set.

### GetMcc

`func (o *CellularIoTMetadata) GetMcc() int32`

GetMcc returns the Mcc field if non-nil, zero value otherwise.

### GetMccOk

`func (o *CellularIoTMetadata) GetMccOk() (*int32, bool)`

GetMccOk returns a tuple with the Mcc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcc

`func (o *CellularIoTMetadata) SetMcc(v int32)`

SetMcc sets Mcc field to given value.

### HasMcc

`func (o *CellularIoTMetadata) HasMcc() bool`

HasMcc returns a boolean if a field has been set.

### GetMnc

`func (o *CellularIoTMetadata) GetMnc() int32`

GetMnc returns the Mnc field if non-nil, zero value otherwise.

### GetMncOk

`func (o *CellularIoTMetadata) GetMncOk() (*int32, bool)`

GetMncOk returns a tuple with the Mnc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMnc

`func (o *CellularIoTMetadata) SetMnc(v int32)`

SetMnc sets Mnc field to given value.

### HasMnc

`func (o *CellularIoTMetadata) HasMnc() bool`

HasMnc returns a boolean if a field has been set.

### GetCountry

`func (o *CellularIoTMetadata) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *CellularIoTMetadata) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *CellularIoTMetadata) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *CellularIoTMetadata) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetNetwork

`func (o *CellularIoTMetadata) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *CellularIoTMetadata) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *CellularIoTMetadata) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *CellularIoTMetadata) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetCountryCode

`func (o *CellularIoTMetadata) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *CellularIoTMetadata) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *CellularIoTMetadata) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *CellularIoTMetadata) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetLastUpdate

`func (o *CellularIoTMetadata) GetLastUpdate() string`

GetLastUpdate returns the LastUpdate field if non-nil, zero value otherwise.

### GetLastUpdateOk

`func (o *CellularIoTMetadata) GetLastUpdateOk() (*string, bool)`

GetLastUpdateOk returns a tuple with the LastUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdate

`func (o *CellularIoTMetadata) SetLastUpdate(v string)`

SetLastUpdate sets LastUpdate field to given value.

### HasLastUpdate

`func (o *CellularIoTMetadata) HasLastUpdate() bool`

HasLastUpdate returns a boolean if a field has been set.

### GetLastImsi

`func (o *CellularIoTMetadata) GetLastImsi() string`

GetLastImsi returns the LastImsi field if non-nil, zero value otherwise.

### GetLastImsiOk

`func (o *CellularIoTMetadata) GetLastImsiOk() (*string, bool)`

GetLastImsiOk returns a tuple with the LastImsi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastImsi

`func (o *CellularIoTMetadata) SetLastImsi(v string)`

SetLastImsi sets LastImsi field to given value.

### HasLastImsi

`func (o *CellularIoTMetadata) HasLastImsi() bool`

HasLastImsi returns a boolean if a field has been set.

### GetLastImei

`func (o *CellularIoTMetadata) GetLastImei() string`

GetLastImei returns the LastImei field if non-nil, zero value otherwise.

### GetLastImeiOk

`func (o *CellularIoTMetadata) GetLastImeiOk() (*string, bool)`

GetLastImeiOk returns a tuple with the LastImei field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastImei

`func (o *CellularIoTMetadata) SetLastImei(v string)`

SetLastImei sets LastImei field to given value.

### HasLastImei

`func (o *CellularIoTMetadata) HasLastImei() bool`

HasLastImei returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


