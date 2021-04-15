# DataDumpResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Collections** | Pointer to [**[]DumpedCollection**](DumpedCollection.md) |  | [optional] 

## Methods

### NewDataDumpResponse

`func NewDataDumpResponse() *DataDumpResponse`

NewDataDumpResponse instantiates a new DataDumpResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataDumpResponseWithDefaults

`func NewDataDumpResponseWithDefaults() *DataDumpResponse`

NewDataDumpResponseWithDefaults instantiates a new DataDumpResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollections

`func (o *DataDumpResponse) GetCollections() []DumpedCollection`

GetCollections returns the Collections field if non-nil, zero value otherwise.

### GetCollectionsOk

`func (o *DataDumpResponse) GetCollectionsOk() (*[]DumpedCollection, bool)`

GetCollectionsOk returns a tuple with the Collections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollections

`func (o *DataDumpResponse) SetCollections(v []DumpedCollection)`

SetCollections sets Collections field to given value.

### HasCollections

`func (o *DataDumpResponse) HasCollections() bool`

HasCollections returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


