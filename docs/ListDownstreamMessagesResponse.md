# ListDownstreamMessagesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Messages** | Pointer to [**[]MessageDownstream**](MessageDownstream.md) |  | [optional] 

## Methods

### NewListDownstreamMessagesResponse

`func NewListDownstreamMessagesResponse() *ListDownstreamMessagesResponse`

NewListDownstreamMessagesResponse instantiates a new ListDownstreamMessagesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListDownstreamMessagesResponseWithDefaults

`func NewListDownstreamMessagesResponseWithDefaults() *ListDownstreamMessagesResponse`

NewListDownstreamMessagesResponseWithDefaults instantiates a new ListDownstreamMessagesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessages

`func (o *ListDownstreamMessagesResponse) GetMessages() []MessageDownstream`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *ListDownstreamMessagesResponse) GetMessagesOk() (*[]MessageDownstream, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *ListDownstreamMessagesResponse) SetMessages(v []MessageDownstream)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *ListDownstreamMessagesResponse) HasMessages() bool`

HasMessages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


