# ListCollectionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Collections** | Pointer to [**[]Collection**](Collection.md) |  | [optional] 

## Methods

### NewListCollectionResponse

`func NewListCollectionResponse() *ListCollectionResponse`

NewListCollectionResponse instantiates a new ListCollectionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCollectionResponseWithDefaults

`func NewListCollectionResponseWithDefaults() *ListCollectionResponse`

NewListCollectionResponseWithDefaults instantiates a new ListCollectionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollections

`func (o *ListCollectionResponse) GetCollections() []Collection`

GetCollections returns the Collections field if non-nil, zero value otherwise.

### GetCollectionsOk

`func (o *ListCollectionResponse) GetCollectionsOk() (*[]Collection, bool)`

GetCollectionsOk returns a tuple with the Collections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollections

`func (o *ListCollectionResponse) SetCollections(v []Collection)`

SetCollections sets Collections field to given value.

### HasCollections

`func (o *ListCollectionResponse) HasCollections() bool`

HasCollections returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


