# OutputLogEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Time** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Repeated** | Pointer to **int32** |  | [optional] 

## Methods

### NewOutputLogEntry

`func NewOutputLogEntry() *OutputLogEntry`

NewOutputLogEntry instantiates a new OutputLogEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutputLogEntryWithDefaults

`func NewOutputLogEntryWithDefaults() *OutputLogEntry`

NewOutputLogEntryWithDefaults instantiates a new OutputLogEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTime

`func (o *OutputLogEntry) GetTime() string`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *OutputLogEntry) GetTimeOk() (*string, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *OutputLogEntry) SetTime(v string)`

SetTime sets Time field to given value.

### HasTime

`func (o *OutputLogEntry) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetMessage

`func (o *OutputLogEntry) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *OutputLogEntry) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *OutputLogEntry) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *OutputLogEntry) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetRepeated

`func (o *OutputLogEntry) GetRepeated() int32`

GetRepeated returns the Repeated field if non-nil, zero value otherwise.

### GetRepeatedOk

`func (o *OutputLogEntry) GetRepeatedOk() (*int32, bool)`

GetRepeatedOk returns a tuple with the Repeated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepeated

`func (o *OutputLogEntry) SetRepeated(v int32)`

SetRepeated sets Repeated field to given value.

### HasRepeated

`func (o *OutputLogEntry) HasRepeated() bool`

HasRepeated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


