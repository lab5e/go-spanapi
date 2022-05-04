# SignCertificateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Certificate** | Pointer to **string** |  | [optional] 
**Chain** | Pointer to **string** |  | [optional] 

## Methods

### NewSignCertificateResponse

`func NewSignCertificateResponse() *SignCertificateResponse`

NewSignCertificateResponse instantiates a new SignCertificateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignCertificateResponseWithDefaults

`func NewSignCertificateResponseWithDefaults() *SignCertificateResponse`

NewSignCertificateResponseWithDefaults instantiates a new SignCertificateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificate

`func (o *SignCertificateResponse) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *SignCertificateResponse) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *SignCertificateResponse) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *SignCertificateResponse) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### GetChain

`func (o *SignCertificateResponse) GetChain() string`

GetChain returns the Chain field if non-nil, zero value otherwise.

### GetChainOk

`func (o *SignCertificateResponse) GetChainOk() (*string, bool)`

GetChainOk returns a tuple with the Chain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChain

`func (o *SignCertificateResponse) SetChain(v string)`

SetChain sets Chain field to given value.

### HasChain

`func (o *SignCertificateResponse) HasChain() bool`

HasChain returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


