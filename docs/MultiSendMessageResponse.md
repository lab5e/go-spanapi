# MultiSendMessageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | Pointer to [**[]MessageSendResult**](MessageSendResult.md) |  | [optional] 
**Sent** | Pointer to **int32** |  | [optional] 
**Failed** | Pointer to **int32** |  | [optional] 

## Methods

### NewMultiSendMessageResponse

`func NewMultiSendMessageResponse() *MultiSendMessageResponse`

NewMultiSendMessageResponse instantiates a new MultiSendMessageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultiSendMessageResponseWithDefaults

`func NewMultiSendMessageResponseWithDefaults() *MultiSendMessageResponse`

NewMultiSendMessageResponseWithDefaults instantiates a new MultiSendMessageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrors

`func (o *MultiSendMessageResponse) GetErrors() []MessageSendResult`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *MultiSendMessageResponse) GetErrorsOk() (*[]MessageSendResult, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *MultiSendMessageResponse) SetErrors(v []MessageSendResult)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *MultiSendMessageResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetSent

`func (o *MultiSendMessageResponse) GetSent() int32`

GetSent returns the Sent field if non-nil, zero value otherwise.

### GetSentOk

`func (o *MultiSendMessageResponse) GetSentOk() (*int32, bool)`

GetSentOk returns a tuple with the Sent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSent

`func (o *MultiSendMessageResponse) SetSent(v int32)`

SetSent sets Sent field to given value.

### HasSent

`func (o *MultiSendMessageResponse) HasSent() bool`

HasSent returns a boolean if a field has been set.

### GetFailed

`func (o *MultiSendMessageResponse) GetFailed() int32`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *MultiSendMessageResponse) GetFailedOk() (*int32, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *MultiSendMessageResponse) SetFailed(v int32)`

SetFailed sets Failed field to given value.

### HasFailed

`func (o *MultiSendMessageResponse) HasFailed() bool`

HasFailed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


