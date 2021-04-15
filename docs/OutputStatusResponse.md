# OutputStatusResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CollectionId** | Pointer to **string** |  | [optional] 
**OutputId** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**ErrorCount** | Pointer to **int32** |  | [optional] 
**Forwarded** | Pointer to **int32** |  | [optional] 
**Received** | Pointer to **int32** |  | [optional] 
**Retransmits** | Pointer to **int32** |  | [optional] 

## Methods

### NewOutputStatusResponse

`func NewOutputStatusResponse() *OutputStatusResponse`

NewOutputStatusResponse instantiates a new OutputStatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutputStatusResponseWithDefaults

`func NewOutputStatusResponseWithDefaults() *OutputStatusResponse`

NewOutputStatusResponseWithDefaults instantiates a new OutputStatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollectionId

`func (o *OutputStatusResponse) GetCollectionId() string`

GetCollectionId returns the CollectionId field if non-nil, zero value otherwise.

### GetCollectionIdOk

`func (o *OutputStatusResponse) GetCollectionIdOk() (*string, bool)`

GetCollectionIdOk returns a tuple with the CollectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionId

`func (o *OutputStatusResponse) SetCollectionId(v string)`

SetCollectionId sets CollectionId field to given value.

### HasCollectionId

`func (o *OutputStatusResponse) HasCollectionId() bool`

HasCollectionId returns a boolean if a field has been set.

### GetOutputId

`func (o *OutputStatusResponse) GetOutputId() string`

GetOutputId returns the OutputId field if non-nil, zero value otherwise.

### GetOutputIdOk

`func (o *OutputStatusResponse) GetOutputIdOk() (*string, bool)`

GetOutputIdOk returns a tuple with the OutputId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputId

`func (o *OutputStatusResponse) SetOutputId(v string)`

SetOutputId sets OutputId field to given value.

### HasOutputId

`func (o *OutputStatusResponse) HasOutputId() bool`

HasOutputId returns a boolean if a field has been set.

### GetEnabled

`func (o *OutputStatusResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *OutputStatusResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *OutputStatusResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *OutputStatusResponse) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetErrorCount

`func (o *OutputStatusResponse) GetErrorCount() int32`

GetErrorCount returns the ErrorCount field if non-nil, zero value otherwise.

### GetErrorCountOk

`func (o *OutputStatusResponse) GetErrorCountOk() (*int32, bool)`

GetErrorCountOk returns a tuple with the ErrorCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCount

`func (o *OutputStatusResponse) SetErrorCount(v int32)`

SetErrorCount sets ErrorCount field to given value.

### HasErrorCount

`func (o *OutputStatusResponse) HasErrorCount() bool`

HasErrorCount returns a boolean if a field has been set.

### GetForwarded

`func (o *OutputStatusResponse) GetForwarded() int32`

GetForwarded returns the Forwarded field if non-nil, zero value otherwise.

### GetForwardedOk

`func (o *OutputStatusResponse) GetForwardedOk() (*int32, bool)`

GetForwardedOk returns a tuple with the Forwarded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwarded

`func (o *OutputStatusResponse) SetForwarded(v int32)`

SetForwarded sets Forwarded field to given value.

### HasForwarded

`func (o *OutputStatusResponse) HasForwarded() bool`

HasForwarded returns a boolean if a field has been set.

### GetReceived

`func (o *OutputStatusResponse) GetReceived() int32`

GetReceived returns the Received field if non-nil, zero value otherwise.

### GetReceivedOk

`func (o *OutputStatusResponse) GetReceivedOk() (*int32, bool)`

GetReceivedOk returns a tuple with the Received field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceived

`func (o *OutputStatusResponse) SetReceived(v int32)`

SetReceived sets Received field to given value.

### HasReceived

`func (o *OutputStatusResponse) HasReceived() bool`

HasReceived returns a boolean if a field has been set.

### GetRetransmits

`func (o *OutputStatusResponse) GetRetransmits() int32`

GetRetransmits returns the Retransmits field if non-nil, zero value otherwise.

### GetRetransmitsOk

`func (o *OutputStatusResponse) GetRetransmitsOk() (*int32, bool)`

GetRetransmitsOk returns a tuple with the Retransmits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetransmits

`func (o *OutputStatusResponse) SetRetransmits(v int32)`

SetRetransmits sets Retransmits field to given value.

### HasRetransmits

`func (o *OutputStatusResponse) HasRetransmits() bool`

HasRetransmits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


