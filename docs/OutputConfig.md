# OutputConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** |  | [optional] 
**BasicAuthUser** | Pointer to **string** |  | [optional] 
**BasicAuthPass** | Pointer to **string** |  | [optional] 
**CustomHeaderName** | Pointer to **string** |  | [optional] 
**CustomHeaderValue** | Pointer to **string** |  | [optional] 
**Host** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **int32** |  | [optional] 
**Key** | Pointer to **string** |  | [optional] 
**EventName** | Pointer to **string** |  | [optional] 
**AsIsPayload** | Pointer to **bool** |  | [optional] 
**Endpoint** | Pointer to **string** |  | [optional] 
**DisableCertCheck** | Pointer to **bool** | MQTT configuration: Disable certificate checks. Default is off. | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**ClientId** | Pointer to **string** |  | [optional] 
**TopicName** | Pointer to **string** |  | [optional] 

## Methods

### NewOutputConfig

`func NewOutputConfig() *OutputConfig`

NewOutputConfig instantiates a new OutputConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutputConfigWithDefaults

`func NewOutputConfigWithDefaults() *OutputConfig`

NewOutputConfigWithDefaults instantiates a new OutputConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *OutputConfig) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *OutputConfig) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *OutputConfig) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *OutputConfig) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetBasicAuthUser

`func (o *OutputConfig) GetBasicAuthUser() string`

GetBasicAuthUser returns the BasicAuthUser field if non-nil, zero value otherwise.

### GetBasicAuthUserOk

`func (o *OutputConfig) GetBasicAuthUserOk() (*string, bool)`

GetBasicAuthUserOk returns a tuple with the BasicAuthUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicAuthUser

`func (o *OutputConfig) SetBasicAuthUser(v string)`

SetBasicAuthUser sets BasicAuthUser field to given value.

### HasBasicAuthUser

`func (o *OutputConfig) HasBasicAuthUser() bool`

HasBasicAuthUser returns a boolean if a field has been set.

### GetBasicAuthPass

`func (o *OutputConfig) GetBasicAuthPass() string`

GetBasicAuthPass returns the BasicAuthPass field if non-nil, zero value otherwise.

### GetBasicAuthPassOk

`func (o *OutputConfig) GetBasicAuthPassOk() (*string, bool)`

GetBasicAuthPassOk returns a tuple with the BasicAuthPass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicAuthPass

`func (o *OutputConfig) SetBasicAuthPass(v string)`

SetBasicAuthPass sets BasicAuthPass field to given value.

### HasBasicAuthPass

`func (o *OutputConfig) HasBasicAuthPass() bool`

HasBasicAuthPass returns a boolean if a field has been set.

### GetCustomHeaderName

`func (o *OutputConfig) GetCustomHeaderName() string`

GetCustomHeaderName returns the CustomHeaderName field if non-nil, zero value otherwise.

### GetCustomHeaderNameOk

`func (o *OutputConfig) GetCustomHeaderNameOk() (*string, bool)`

GetCustomHeaderNameOk returns a tuple with the CustomHeaderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomHeaderName

`func (o *OutputConfig) SetCustomHeaderName(v string)`

SetCustomHeaderName sets CustomHeaderName field to given value.

### HasCustomHeaderName

`func (o *OutputConfig) HasCustomHeaderName() bool`

HasCustomHeaderName returns a boolean if a field has been set.

### GetCustomHeaderValue

`func (o *OutputConfig) GetCustomHeaderValue() string`

GetCustomHeaderValue returns the CustomHeaderValue field if non-nil, zero value otherwise.

### GetCustomHeaderValueOk

`func (o *OutputConfig) GetCustomHeaderValueOk() (*string, bool)`

GetCustomHeaderValueOk returns a tuple with the CustomHeaderValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomHeaderValue

`func (o *OutputConfig) SetCustomHeaderValue(v string)`

SetCustomHeaderValue sets CustomHeaderValue field to given value.

### HasCustomHeaderValue

`func (o *OutputConfig) HasCustomHeaderValue() bool`

HasCustomHeaderValue returns a boolean if a field has been set.

### GetHost

`func (o *OutputConfig) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *OutputConfig) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *OutputConfig) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *OutputConfig) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetPort

`func (o *OutputConfig) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *OutputConfig) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *OutputConfig) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *OutputConfig) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetKey

`func (o *OutputConfig) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *OutputConfig) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *OutputConfig) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *OutputConfig) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetEventName

`func (o *OutputConfig) GetEventName() string`

GetEventName returns the EventName field if non-nil, zero value otherwise.

### GetEventNameOk

`func (o *OutputConfig) GetEventNameOk() (*string, bool)`

GetEventNameOk returns a tuple with the EventName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventName

`func (o *OutputConfig) SetEventName(v string)`

SetEventName sets EventName field to given value.

### HasEventName

`func (o *OutputConfig) HasEventName() bool`

HasEventName returns a boolean if a field has been set.

### GetAsIsPayload

`func (o *OutputConfig) GetAsIsPayload() bool`

GetAsIsPayload returns the AsIsPayload field if non-nil, zero value otherwise.

### GetAsIsPayloadOk

`func (o *OutputConfig) GetAsIsPayloadOk() (*bool, bool)`

GetAsIsPayloadOk returns a tuple with the AsIsPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsIsPayload

`func (o *OutputConfig) SetAsIsPayload(v bool)`

SetAsIsPayload sets AsIsPayload field to given value.

### HasAsIsPayload

`func (o *OutputConfig) HasAsIsPayload() bool`

HasAsIsPayload returns a boolean if a field has been set.

### GetEndpoint

`func (o *OutputConfig) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *OutputConfig) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *OutputConfig) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *OutputConfig) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetDisableCertCheck

`func (o *OutputConfig) GetDisableCertCheck() bool`

GetDisableCertCheck returns the DisableCertCheck field if non-nil, zero value otherwise.

### GetDisableCertCheckOk

`func (o *OutputConfig) GetDisableCertCheckOk() (*bool, bool)`

GetDisableCertCheckOk returns a tuple with the DisableCertCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableCertCheck

`func (o *OutputConfig) SetDisableCertCheck(v bool)`

SetDisableCertCheck sets DisableCertCheck field to given value.

### HasDisableCertCheck

`func (o *OutputConfig) HasDisableCertCheck() bool`

HasDisableCertCheck returns a boolean if a field has been set.

### GetUsername

`func (o *OutputConfig) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *OutputConfig) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *OutputConfig) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *OutputConfig) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *OutputConfig) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *OutputConfig) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *OutputConfig) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *OutputConfig) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetClientId

`func (o *OutputConfig) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *OutputConfig) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *OutputConfig) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *OutputConfig) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetTopicName

`func (o *OutputConfig) GetTopicName() string`

GetTopicName returns the TopicName field if non-nil, zero value otherwise.

### GetTopicNameOk

`func (o *OutputConfig) GetTopicNameOk() (*string, bool)`

GetTopicNameOk returns a tuple with the TopicName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopicName

`func (o *OutputConfig) SetTopicName(v string)`

SetTopicName sets TopicName field to given value.

### HasTopicName

`func (o *OutputConfig) HasTopicName() bool`

HasTopicName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


