# ListBlobResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Blobs** | Pointer to [**[]Blob**](Blob.md) |  | [optional] 

## Methods

### NewListBlobResponse

`func NewListBlobResponse() *ListBlobResponse`

NewListBlobResponse instantiates a new ListBlobResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListBlobResponseWithDefaults

`func NewListBlobResponseWithDefaults() *ListBlobResponse`

NewListBlobResponseWithDefaults instantiates a new ListBlobResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlobs

`func (o *ListBlobResponse) GetBlobs() []Blob`

GetBlobs returns the Blobs field if non-nil, zero value otherwise.

### GetBlobsOk

`func (o *ListBlobResponse) GetBlobsOk() (*[]Blob, bool)`

GetBlobsOk returns a tuple with the Blobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlobs

`func (o *ListBlobResponse) SetBlobs(v []Blob)`

SetBlobs sets Blobs field to given value.

### HasBlobs

`func (o *ListBlobResponse) HasBlobs() bool`

HasBlobs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


