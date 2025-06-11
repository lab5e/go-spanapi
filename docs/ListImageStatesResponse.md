# ListImageStatesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**States** | Pointer to [**[]DeviceImageState**](DeviceImageState.md) |  | [optional] 

## Methods

### NewListImageStatesResponse

`func NewListImageStatesResponse() *ListImageStatesResponse`

NewListImageStatesResponse instantiates a new ListImageStatesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListImageStatesResponseWithDefaults

`func NewListImageStatesResponseWithDefaults() *ListImageStatesResponse`

NewListImageStatesResponseWithDefaults instantiates a new ListImageStatesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStates

`func (o *ListImageStatesResponse) GetStates() []DeviceImageState`

GetStates returns the States field if non-nil, zero value otherwise.

### GetStatesOk

`func (o *ListImageStatesResponse) GetStatesOk() (*[]DeviceImageState, bool)`

GetStatesOk returns a tuple with the States field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStates

`func (o *ListImageStatesResponse) SetStates(v []DeviceImageState)`

SetStates sets States field to given value.

### HasStates

`func (o *ListImageStatesResponse) HasStates() bool`

HasStates returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


