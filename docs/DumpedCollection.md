# DumpedCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Collection** | Pointer to [**Collection**](Collection.md) |  | [optional] 
**Devices** | Pointer to [**[]DumpedDevice**](DumpedDevice.md) |  | [optional] 
**Outputs** | Pointer to [**[]Output**](Output.md) |  | [optional] 

## Methods

### NewDumpedCollection

`func NewDumpedCollection() *DumpedCollection`

NewDumpedCollection instantiates a new DumpedCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDumpedCollectionWithDefaults

`func NewDumpedCollectionWithDefaults() *DumpedCollection`

NewDumpedCollectionWithDefaults instantiates a new DumpedCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollection

`func (o *DumpedCollection) GetCollection() Collection`

GetCollection returns the Collection field if non-nil, zero value otherwise.

### GetCollectionOk

`func (o *DumpedCollection) GetCollectionOk() (*Collection, bool)`

GetCollectionOk returns a tuple with the Collection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollection

`func (o *DumpedCollection) SetCollection(v Collection)`

SetCollection sets Collection field to given value.

### HasCollection

`func (o *DumpedCollection) HasCollection() bool`

HasCollection returns a boolean if a field has been set.

### GetDevices

`func (o *DumpedCollection) GetDevices() []DumpedDevice`

GetDevices returns the Devices field if non-nil, zero value otherwise.

### GetDevicesOk

`func (o *DumpedCollection) GetDevicesOk() (*[]DumpedDevice, bool)`

GetDevicesOk returns a tuple with the Devices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevices

`func (o *DumpedCollection) SetDevices(v []DumpedDevice)`

SetDevices sets Devices field to given value.

### HasDevices

`func (o *DumpedCollection) HasDevices() bool`

HasDevices returns a boolean if a field has been set.

### GetOutputs

`func (o *DumpedCollection) GetOutputs() []Output`

GetOutputs returns the Outputs field if non-nil, zero value otherwise.

### GetOutputsOk

`func (o *DumpedCollection) GetOutputsOk() (*[]Output, bool)`

GetOutputsOk returns a tuple with the Outputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputs

`func (o *DumpedCollection) SetOutputs(v []Output)`

SetOutputs sets Outputs field to given value.

### HasOutputs

`func (o *DumpedCollection) HasOutputs() bool`

HasOutputs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


