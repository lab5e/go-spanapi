# FieldMask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Imsi** | Pointer to **bool** |  | [optional] 
**Imei** | Pointer to **bool** |  | [optional] 
**Msisdn** | Pointer to **bool** |  | [optional] 
**Location** | Pointer to **bool** |  | [optional] 

## Methods

### NewFieldMask

`func NewFieldMask() *FieldMask`

NewFieldMask instantiates a new FieldMask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFieldMaskWithDefaults

`func NewFieldMaskWithDefaults() *FieldMask`

NewFieldMaskWithDefaults instantiates a new FieldMask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImsi

`func (o *FieldMask) GetImsi() bool`

GetImsi returns the Imsi field if non-nil, zero value otherwise.

### GetImsiOk

`func (o *FieldMask) GetImsiOk() (*bool, bool)`

GetImsiOk returns a tuple with the Imsi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImsi

`func (o *FieldMask) SetImsi(v bool)`

SetImsi sets Imsi field to given value.

### HasImsi

`func (o *FieldMask) HasImsi() bool`

HasImsi returns a boolean if a field has been set.

### GetImei

`func (o *FieldMask) GetImei() bool`

GetImei returns the Imei field if non-nil, zero value otherwise.

### GetImeiOk

`func (o *FieldMask) GetImeiOk() (*bool, bool)`

GetImeiOk returns a tuple with the Imei field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImei

`func (o *FieldMask) SetImei(v bool)`

SetImei sets Imei field to given value.

### HasImei

`func (o *FieldMask) HasImei() bool`

HasImei returns a boolean if a field has been set.

### GetMsisdn

`func (o *FieldMask) GetMsisdn() bool`

GetMsisdn returns the Msisdn field if non-nil, zero value otherwise.

### GetMsisdnOk

`func (o *FieldMask) GetMsisdnOk() (*bool, bool)`

GetMsisdnOk returns a tuple with the Msisdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsisdn

`func (o *FieldMask) SetMsisdn(v bool)`

SetMsisdn sets Msisdn field to given value.

### HasMsisdn

`func (o *FieldMask) HasMsisdn() bool`

HasMsisdn returns a boolean if a field has been set.

### GetLocation

`func (o *FieldMask) GetLocation() bool`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *FieldMask) GetLocationOk() (*bool, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *FieldMask) SetLocation(v bool)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *FieldMask) HasLocation() bool`

HasLocation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


