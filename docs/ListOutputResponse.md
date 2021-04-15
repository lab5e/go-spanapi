# ListOutputResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CollectionId** | Pointer to **string** |  | [optional] 
**Outputs** | Pointer to [**[]Output**](Output.md) |  | [optional] 

## Methods

### NewListOutputResponse

`func NewListOutputResponse() *ListOutputResponse`

NewListOutputResponse instantiates a new ListOutputResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListOutputResponseWithDefaults

`func NewListOutputResponseWithDefaults() *ListOutputResponse`

NewListOutputResponseWithDefaults instantiates a new ListOutputResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollectionId

`func (o *ListOutputResponse) GetCollectionId() string`

GetCollectionId returns the CollectionId field if non-nil, zero value otherwise.

### GetCollectionIdOk

`func (o *ListOutputResponse) GetCollectionIdOk() (*string, bool)`

GetCollectionIdOk returns a tuple with the CollectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionId

`func (o *ListOutputResponse) SetCollectionId(v string)`

SetCollectionId sets CollectionId field to given value.

### HasCollectionId

`func (o *ListOutputResponse) HasCollectionId() bool`

HasCollectionId returns a boolean if a field has been set.

### GetOutputs

`func (o *ListOutputResponse) GetOutputs() []Output`

GetOutputs returns the Outputs field if non-nil, zero value otherwise.

### GetOutputsOk

`func (o *ListOutputResponse) GetOutputsOk() (*[]Output, bool)`

GetOutputsOk returns a tuple with the Outputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputs

`func (o *ListOutputResponse) SetOutputs(v []Output)`

SetOutputs sets Outputs field to given value.

### HasOutputs

`func (o *ListOutputResponse) HasOutputs() bool`

HasOutputs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


