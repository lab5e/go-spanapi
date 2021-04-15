# FirmwareUsageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImageId** | Pointer to **string** |  | [optional] 
**Targeted** | Pointer to **[]string** |  | [optional] 
**Current** | Pointer to **[]string** |  | [optional] 

## Methods

### NewFirmwareUsageResponse

`func NewFirmwareUsageResponse() *FirmwareUsageResponse`

NewFirmwareUsageResponse instantiates a new FirmwareUsageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirmwareUsageResponseWithDefaults

`func NewFirmwareUsageResponseWithDefaults() *FirmwareUsageResponse`

NewFirmwareUsageResponseWithDefaults instantiates a new FirmwareUsageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImageId

`func (o *FirmwareUsageResponse) GetImageId() string`

GetImageId returns the ImageId field if non-nil, zero value otherwise.

### GetImageIdOk

`func (o *FirmwareUsageResponse) GetImageIdOk() (*string, bool)`

GetImageIdOk returns a tuple with the ImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageId

`func (o *FirmwareUsageResponse) SetImageId(v string)`

SetImageId sets ImageId field to given value.

### HasImageId

`func (o *FirmwareUsageResponse) HasImageId() bool`

HasImageId returns a boolean if a field has been set.

### GetTargeted

`func (o *FirmwareUsageResponse) GetTargeted() []string`

GetTargeted returns the Targeted field if non-nil, zero value otherwise.

### GetTargetedOk

`func (o *FirmwareUsageResponse) GetTargetedOk() (*[]string, bool)`

GetTargetedOk returns a tuple with the Targeted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargeted

`func (o *FirmwareUsageResponse) SetTargeted(v []string)`

SetTargeted sets Targeted field to given value.

### HasTargeted

`func (o *FirmwareUsageResponse) HasTargeted() bool`

HasTargeted returns a boolean if a field has been set.

### GetCurrent

`func (o *FirmwareUsageResponse) GetCurrent() []string`

GetCurrent returns the Current field if non-nil, zero value otherwise.

### GetCurrentOk

`func (o *FirmwareUsageResponse) GetCurrentOk() (*[]string, bool)`

GetCurrentOk returns a tuple with the Current field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrent

`func (o *FirmwareUsageResponse) SetCurrent(v []string)`

SetCurrent sets Current field to given value.

### HasCurrent

`func (o *FirmwareUsageResponse) HasCurrent() bool`

HasCurrent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


