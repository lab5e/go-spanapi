# SystemInfoResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | Pointer to **string** | This is the system version | [optional] 
**BuildDate** | Pointer to **string** | The build time for this version. | [optional] 
**ReleaseName** | Pointer to **string** | Human-readable code name for this release. This can be easier to remember than the version number. | [optional] 

## Methods

### NewSystemInfoResponse

`func NewSystemInfoResponse() *SystemInfoResponse`

NewSystemInfoResponse instantiates a new SystemInfoResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemInfoResponseWithDefaults

`func NewSystemInfoResponseWithDefaults() *SystemInfoResponse`

NewSystemInfoResponseWithDefaults instantiates a new SystemInfoResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *SystemInfoResponse) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SystemInfoResponse) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SystemInfoResponse) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *SystemInfoResponse) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetBuildDate

`func (o *SystemInfoResponse) GetBuildDate() string`

GetBuildDate returns the BuildDate field if non-nil, zero value otherwise.

### GetBuildDateOk

`func (o *SystemInfoResponse) GetBuildDateOk() (*string, bool)`

GetBuildDateOk returns a tuple with the BuildDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildDate

`func (o *SystemInfoResponse) SetBuildDate(v string)`

SetBuildDate sets BuildDate field to given value.

### HasBuildDate

`func (o *SystemInfoResponse) HasBuildDate() bool`

HasBuildDate returns a boolean if a field has been set.

### GetReleaseName

`func (o *SystemInfoResponse) GetReleaseName() string`

GetReleaseName returns the ReleaseName field if non-nil, zero value otherwise.

### GetReleaseNameOk

`func (o *SystemInfoResponse) GetReleaseNameOk() (*string, bool)`

GetReleaseNameOk returns a tuple with the ReleaseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseName

`func (o *SystemInfoResponse) SetReleaseName(v string)`

SetReleaseName sets ReleaseName field to given value.

### HasReleaseName

`func (o *SystemInfoResponse) HasReleaseName() bool`

HasReleaseName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


