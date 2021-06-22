# NetworkOperator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mcc** | Pointer to **int32** | The Mobil Country Code for the operator. | [optional] 
**Mnc** | Pointer to **int32** |  | [optional] 
**Country** | Pointer to **string** |  | [optional] 
**Network** | Pointer to **string** |  | [optional] 

## Methods

### NewNetworkOperator

`func NewNetworkOperator() *NetworkOperator`

NewNetworkOperator instantiates a new NetworkOperator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkOperatorWithDefaults

`func NewNetworkOperatorWithDefaults() *NetworkOperator`

NewNetworkOperatorWithDefaults instantiates a new NetworkOperator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMcc

`func (o *NetworkOperator) GetMcc() int32`

GetMcc returns the Mcc field if non-nil, zero value otherwise.

### GetMccOk

`func (o *NetworkOperator) GetMccOk() (*int32, bool)`

GetMccOk returns a tuple with the Mcc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcc

`func (o *NetworkOperator) SetMcc(v int32)`

SetMcc sets Mcc field to given value.

### HasMcc

`func (o *NetworkOperator) HasMcc() bool`

HasMcc returns a boolean if a field has been set.

### GetMnc

`func (o *NetworkOperator) GetMnc() int32`

GetMnc returns the Mnc field if non-nil, zero value otherwise.

### GetMncOk

`func (o *NetworkOperator) GetMncOk() (*int32, bool)`

GetMncOk returns a tuple with the Mnc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMnc

`func (o *NetworkOperator) SetMnc(v int32)`

SetMnc sets Mnc field to given value.

### HasMnc

`func (o *NetworkOperator) HasMnc() bool`

HasMnc returns a boolean if a field has been set.

### GetCountry

`func (o *NetworkOperator) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *NetworkOperator) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *NetworkOperator) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *NetworkOperator) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetNetwork

`func (o *NetworkOperator) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *NetworkOperator) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *NetworkOperator) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *NetworkOperator) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


