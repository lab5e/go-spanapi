# OutputLogResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Logs** | Pointer to [**[]OutputLogEntry**](OutputLogEntry.md) |  | [optional] 

## Methods

### NewOutputLogResponse

`func NewOutputLogResponse() *OutputLogResponse`

NewOutputLogResponse instantiates a new OutputLogResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutputLogResponseWithDefaults

`func NewOutputLogResponseWithDefaults() *OutputLogResponse`

NewOutputLogResponseWithDefaults instantiates a new OutputLogResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogs

`func (o *OutputLogResponse) GetLogs() []OutputLogEntry`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *OutputLogResponse) GetLogsOk() (*[]OutputLogEntry, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *OutputLogResponse) SetLogs(v []OutputLogEntry)`

SetLogs sets Logs field to given value.

### HasLogs

`func (o *OutputLogResponse) HasLogs() bool`

HasLogs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


