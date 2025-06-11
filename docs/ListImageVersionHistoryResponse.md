# ListImageVersionHistoryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CollectionId** | Pointer to **string** |  | [optional] 
**Log** | Pointer to [**[]ImageVersionEntry**](ImageVersionEntry.md) |  | [optional] 

## Methods

### NewListImageVersionHistoryResponse

`func NewListImageVersionHistoryResponse() *ListImageVersionHistoryResponse`

NewListImageVersionHistoryResponse instantiates a new ListImageVersionHistoryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListImageVersionHistoryResponseWithDefaults

`func NewListImageVersionHistoryResponseWithDefaults() *ListImageVersionHistoryResponse`

NewListImageVersionHistoryResponseWithDefaults instantiates a new ListImageVersionHistoryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollectionId

`func (o *ListImageVersionHistoryResponse) GetCollectionId() string`

GetCollectionId returns the CollectionId field if non-nil, zero value otherwise.

### GetCollectionIdOk

`func (o *ListImageVersionHistoryResponse) GetCollectionIdOk() (*string, bool)`

GetCollectionIdOk returns a tuple with the CollectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionId

`func (o *ListImageVersionHistoryResponse) SetCollectionId(v string)`

SetCollectionId sets CollectionId field to given value.

### HasCollectionId

`func (o *ListImageVersionHistoryResponse) HasCollectionId() bool`

HasCollectionId returns a boolean if a field has been set.

### GetLog

`func (o *ListImageVersionHistoryResponse) GetLog() []ImageVersionEntry`

GetLog returns the Log field if non-nil, zero value otherwise.

### GetLogOk

`func (o *ListImageVersionHistoryResponse) GetLogOk() (*[]ImageVersionEntry, bool)`

GetLogOk returns a tuple with the Log field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLog

`func (o *ListImageVersionHistoryResponse) SetLog(v []ImageVersionEntry)`

SetLog sets Log field to given value.

### HasLog

`func (o *ListImageVersionHistoryResponse) HasLog() bool`

HasLog returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


