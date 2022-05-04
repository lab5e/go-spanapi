# CertificateInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertificateSerial** | Pointer to **string** |  | [optional] 
**Expires** | Pointer to **string** |  | [optional] 

## Methods

### NewCertificateInfo

`func NewCertificateInfo() *CertificateInfo`

NewCertificateInfo instantiates a new CertificateInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateInfoWithDefaults

`func NewCertificateInfoWithDefaults() *CertificateInfo`

NewCertificateInfoWithDefaults instantiates a new CertificateInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificateSerial

`func (o *CertificateInfo) GetCertificateSerial() string`

GetCertificateSerial returns the CertificateSerial field if non-nil, zero value otherwise.

### GetCertificateSerialOk

`func (o *CertificateInfo) GetCertificateSerialOk() (*string, bool)`

GetCertificateSerialOk returns a tuple with the CertificateSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateSerial

`func (o *CertificateInfo) SetCertificateSerial(v string)`

SetCertificateSerial sets CertificateSerial field to given value.

### HasCertificateSerial

`func (o *CertificateInfo) HasCertificateSerial() bool`

HasCertificateSerial returns a boolean if a field has been set.

### GetExpires

`func (o *CertificateInfo) GetExpires() string`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *CertificateInfo) GetExpiresOk() (*string, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *CertificateInfo) SetExpires(v string)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *CertificateInfo) HasExpires() bool`

HasExpires returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


