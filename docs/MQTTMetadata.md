# MQTTMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Topic** | Pointer to **string** |  | [optional] 

## Methods

### NewMQTTMetadata

`func NewMQTTMetadata() *MQTTMetadata`

NewMQTTMetadata instantiates a new MQTTMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMQTTMetadataWithDefaults

`func NewMQTTMetadataWithDefaults() *MQTTMetadata`

NewMQTTMetadataWithDefaults instantiates a new MQTTMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTopic

`func (o *MQTTMetadata) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *MQTTMetadata) GetTopicOk() (*string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *MQTTMetadata) SetTopic(v string)`

SetTopic sets Topic field to given value.

### HasTopic

`func (o *MQTTMetadata) HasTopic() bool`

HasTopic returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


