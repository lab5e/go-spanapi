# OutputStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ForwardErrors** | Pointer to **int32** |  | [optional] 
**MessagesForwarded** | Pointer to **int32** |  | [optional] 
**BytesForwardedMb** | Pointer to **float32** |  | [optional] 

## Methods

### NewOutputStats

`func NewOutputStats() *OutputStats`

NewOutputStats instantiates a new OutputStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutputStatsWithDefaults

`func NewOutputStatsWithDefaults() *OutputStats`

NewOutputStatsWithDefaults instantiates a new OutputStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetForwardErrors

`func (o *OutputStats) GetForwardErrors() int32`

GetForwardErrors returns the ForwardErrors field if non-nil, zero value otherwise.

### GetForwardErrorsOk

`func (o *OutputStats) GetForwardErrorsOk() (*int32, bool)`

GetForwardErrorsOk returns a tuple with the ForwardErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardErrors

`func (o *OutputStats) SetForwardErrors(v int32)`

SetForwardErrors sets ForwardErrors field to given value.

### HasForwardErrors

`func (o *OutputStats) HasForwardErrors() bool`

HasForwardErrors returns a boolean if a field has been set.

### GetMessagesForwarded

`func (o *OutputStats) GetMessagesForwarded() int32`

GetMessagesForwarded returns the MessagesForwarded field if non-nil, zero value otherwise.

### GetMessagesForwardedOk

`func (o *OutputStats) GetMessagesForwardedOk() (*int32, bool)`

GetMessagesForwardedOk returns a tuple with the MessagesForwarded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagesForwarded

`func (o *OutputStats) SetMessagesForwarded(v int32)`

SetMessagesForwarded sets MessagesForwarded field to given value.

### HasMessagesForwarded

`func (o *OutputStats) HasMessagesForwarded() bool`

HasMessagesForwarded returns a boolean if a field has been set.

### GetBytesForwardedMb

`func (o *OutputStats) GetBytesForwardedMb() float32`

GetBytesForwardedMb returns the BytesForwardedMb field if non-nil, zero value otherwise.

### GetBytesForwardedMbOk

`func (o *OutputStats) GetBytesForwardedMbOk() (*float32, bool)`

GetBytesForwardedMbOk returns a tuple with the BytesForwardedMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytesForwardedMb

`func (o *OutputStats) SetBytesForwardedMb(v float32)`

SetBytesForwardedMb sets BytesForwardedMb field to given value.

### HasBytesForwardedMb

`func (o *OutputStats) HasBytesForwardedMb() bool`

HasBytesForwardedMb returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


