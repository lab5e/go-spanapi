/*
The Span API

API for device, collection, output and firmware management

API version: 4.7.0 actionable-aryanna
Contact: dev@lab5e.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package spanapi

import (
	"encoding/json"
)

// checks if the OutputConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OutputConfig{}

// OutputConfig Configuration for outputs.
type OutputConfig struct {
	// URL for the webhook.
	Url *string `json:"url,omitempty"`
	BasicAuthUser *string `json:"basicAuthUser,omitempty"`
	BasicAuthPass *string `json:"basicAuthPass,omitempty"`
	CustomHeaderName *string `json:"customHeaderName,omitempty"`
	CustomHeaderValue *string `json:"customHeaderValue,omitempty"`
	Host *string `json:"host,omitempty"`
	Port *int32 `json:"port,omitempty"`
	Key *string `json:"key,omitempty"`
	EventName *string `json:"eventName,omitempty"`
	AsIsPayload *bool `json:"asIsPayload,omitempty"`
	Endpoint *string `json:"endpoint,omitempty"`
	// MQTT configuration: Disable certificate checks. Default is off.
	DisableCertCheck *bool `json:"disableCertCheck,omitempty"`
	Username *string `json:"username,omitempty"`
	Password *string `json:"password,omitempty"`
	ClientId *string `json:"clientId,omitempty"`
	TopicName *string `json:"topicName,omitempty"`
	TopicTemplate *string `json:"topicTemplate,omitempty"`
	PayloadFormat *string `json:"payloadFormat,omitempty"`
	PayloadTemplate *string `json:"payloadTemplate,omitempty"`
}

// NewOutputConfig instantiates a new OutputConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOutputConfig() *OutputConfig {
	this := OutputConfig{}
	return &this
}

// NewOutputConfigWithDefaults instantiates a new OutputConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOutputConfigWithDefaults() *OutputConfig {
	this := OutputConfig{}
	return &this
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *OutputConfig) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputConfig) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *OutputConfig) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *OutputConfig) SetUrl(v string) {
	o.Url = &v
}

// GetBasicAuthUser returns the BasicAuthUser field value if set, zero value otherwise.
func (o *OutputConfig) GetBasicAuthUser() string {
	if o == nil || IsNil(o.BasicAuthUser) {
		var ret string
		return ret
	}
	return *o.BasicAuthUser
}

// GetBasicAuthUserOk returns a tuple with the BasicAuthUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputConfig) GetBasicAuthUserOk() (*string, bool) {
	if o == nil || IsNil(o.BasicAuthUser) {
		return nil, false
	}
	return o.BasicAuthUser, true
}

// HasBasicAuthUser returns a boolean if a field has been set.
func (o *OutputConfig) HasBasicAuthUser() bool {
	if o != nil && !IsNil(o.BasicAuthUser) {
		return true
	}

	return false
}

// SetBasicAuthUser gets a reference to the given string and assigns it to the BasicAuthUser field.
func (o *OutputConfig) SetBasicAuthUser(v string) {
	o.BasicAuthUser = &v
}

// GetBasicAuthPass returns the BasicAuthPass field value if set, zero value otherwise.
func (o *OutputConfig) GetBasicAuthPass() string {
	if o == nil || IsNil(o.BasicAuthPass) {
		var ret string
		return ret
	}
	return *o.BasicAuthPass
}

// GetBasicAuthPassOk returns a tuple with the BasicAuthPass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputConfig) GetBasicAuthPassOk() (*string, bool) {
	if o == nil || IsNil(o.BasicAuthPass) {
		return nil, false
	}
	return o.BasicAuthPass, true
}

// HasBasicAuthPass returns a boolean if a field has been set.
func (o *OutputConfig) HasBasicAuthPass() bool {
	if o != nil && !IsNil(o.BasicAuthPass) {
		return true
	}

	return false
}

// SetBasicAuthPass gets a reference to the given string and assigns it to the BasicAuthPass field.
func (o *OutputConfig) SetBasicAuthPass(v string) {
	o.BasicAuthPass = &v
}

// GetCustomHeaderName returns the CustomHeaderName field value if set, zero value otherwise.
func (o *OutputConfig) GetCustomHeaderName() string {
	if o == nil || IsNil(o.CustomHeaderName) {
		var ret string
		return ret
	}
	return *o.CustomHeaderName
}

// GetCustomHeaderNameOk returns a tuple with the CustomHeaderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputConfig) GetCustomHeaderNameOk() (*string, bool) {
	if o == nil || IsNil(o.CustomHeaderName) {
		return nil, false
	}
	return o.CustomHeaderName, true
}

// HasCustomHeaderName returns a boolean if a field has been set.
func (o *OutputConfig) HasCustomHeaderName() bool {
	if o != nil && !IsNil(o.CustomHeaderName) {
		return true
	}

	return false
}

// SetCustomHeaderName gets a reference to the given string and assigns it to the CustomHeaderName field.
func (o *OutputConfig) SetCustomHeaderName(v string) {
	o.CustomHeaderName = &v
}

// GetCustomHeaderValue returns the CustomHeaderValue field value if set, zero value otherwise.
func (o *OutputConfig) GetCustomHeaderValue() string {
	if o == nil || IsNil(o.CustomHeaderValue) {
		var ret string
		return ret
	}
	return *o.CustomHeaderValue
}

// GetCustomHeaderValueOk returns a tuple with the CustomHeaderValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputConfig) GetCustomHeaderValueOk() (*string, bool) {
	if o == nil || IsNil(o.CustomHeaderValue) {
		return nil, false
	}
	return o.CustomHeaderValue, true
}

// HasCustomHeaderValue returns a boolean if a field has been set.
func (o *OutputConfig) HasCustomHeaderValue() bool {
	if o != nil && !IsNil(o.CustomHeaderValue) {
		return true
	}

	return false
}

// SetCustomHeaderValue gets a reference to the given string and assigns it to the CustomHeaderValue field.
func (o *OutputConfig) SetCustomHeaderValue(v string) {
	o.CustomHeaderValue = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *OutputConfig) GetHost() string {
	if o == nil || IsNil(o.Host) {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputConfig) GetHostOk() (*string, bool) {
	if o == nil || IsNil(o.Host) {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *OutputConfig) HasHost() bool {
	if o != nil && !IsNil(o.Host) {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *OutputConfig) SetHost(v string) {
	o.Host = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *OutputConfig) GetPort() int32 {
	if o == nil || IsNil(o.Port) {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputConfig) GetPortOk() (*int32, bool) {
	if o == nil || IsNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *OutputConfig) HasPort() bool {
	if o != nil && !IsNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *OutputConfig) SetPort(v int32) {
	o.Port = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *OutputConfig) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputConfig) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *OutputConfig) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *OutputConfig) SetKey(v string) {
	o.Key = &v
}

// GetEventName returns the EventName field value if set, zero value otherwise.
func (o *OutputConfig) GetEventName() string {
	if o == nil || IsNil(o.EventName) {
		var ret string
		return ret
	}
	return *o.EventName
}

// GetEventNameOk returns a tuple with the EventName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputConfig) GetEventNameOk() (*string, bool) {
	if o == nil || IsNil(o.EventName) {
		return nil, false
	}
	return o.EventName, true
}

// HasEventName returns a boolean if a field has been set.
func (o *OutputConfig) HasEventName() bool {
	if o != nil && !IsNil(o.EventName) {
		return true
	}

	return false
}

// SetEventName gets a reference to the given string and assigns it to the EventName field.
func (o *OutputConfig) SetEventName(v string) {
	o.EventName = &v
}

// GetAsIsPayload returns the AsIsPayload field value if set, zero value otherwise.
func (o *OutputConfig) GetAsIsPayload() bool {
	if o == nil || IsNil(o.AsIsPayload) {
		var ret bool
		return ret
	}
	return *o.AsIsPayload
}

// GetAsIsPayloadOk returns a tuple with the AsIsPayload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputConfig) GetAsIsPayloadOk() (*bool, bool) {
	if o == nil || IsNil(o.AsIsPayload) {
		return nil, false
	}
	return o.AsIsPayload, true
}

// HasAsIsPayload returns a boolean if a field has been set.
func (o *OutputConfig) HasAsIsPayload() bool {
	if o != nil && !IsNil(o.AsIsPayload) {
		return true
	}

	return false
}

// SetAsIsPayload gets a reference to the given bool and assigns it to the AsIsPayload field.
func (o *OutputConfig) SetAsIsPayload(v bool) {
	o.AsIsPayload = &v
}

// GetEndpoint returns the Endpoint field value if set, zero value otherwise.
func (o *OutputConfig) GetEndpoint() string {
	if o == nil || IsNil(o.Endpoint) {
		var ret string
		return ret
	}
	return *o.Endpoint
}

// GetEndpointOk returns a tuple with the Endpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputConfig) GetEndpointOk() (*string, bool) {
	if o == nil || IsNil(o.Endpoint) {
		return nil, false
	}
	return o.Endpoint, true
}

// HasEndpoint returns a boolean if a field has been set.
func (o *OutputConfig) HasEndpoint() bool {
	if o != nil && !IsNil(o.Endpoint) {
		return true
	}

	return false
}

// SetEndpoint gets a reference to the given string and assigns it to the Endpoint field.
func (o *OutputConfig) SetEndpoint(v string) {
	o.Endpoint = &v
}

// GetDisableCertCheck returns the DisableCertCheck field value if set, zero value otherwise.
func (o *OutputConfig) GetDisableCertCheck() bool {
	if o == nil || IsNil(o.DisableCertCheck) {
		var ret bool
		return ret
	}
	return *o.DisableCertCheck
}

// GetDisableCertCheckOk returns a tuple with the DisableCertCheck field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputConfig) GetDisableCertCheckOk() (*bool, bool) {
	if o == nil || IsNil(o.DisableCertCheck) {
		return nil, false
	}
	return o.DisableCertCheck, true
}

// HasDisableCertCheck returns a boolean if a field has been set.
func (o *OutputConfig) HasDisableCertCheck() bool {
	if o != nil && !IsNil(o.DisableCertCheck) {
		return true
	}

	return false
}

// SetDisableCertCheck gets a reference to the given bool and assigns it to the DisableCertCheck field.
func (o *OutputConfig) SetDisableCertCheck(v bool) {
	o.DisableCertCheck = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *OutputConfig) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputConfig) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *OutputConfig) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *OutputConfig) SetUsername(v string) {
	o.Username = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *OutputConfig) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputConfig) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *OutputConfig) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *OutputConfig) SetPassword(v string) {
	o.Password = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *OutputConfig) GetClientId() string {
	if o == nil || IsNil(o.ClientId) {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputConfig) GetClientIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClientId) {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *OutputConfig) HasClientId() bool {
	if o != nil && !IsNil(o.ClientId) {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *OutputConfig) SetClientId(v string) {
	o.ClientId = &v
}

// GetTopicName returns the TopicName field value if set, zero value otherwise.
func (o *OutputConfig) GetTopicName() string {
	if o == nil || IsNil(o.TopicName) {
		var ret string
		return ret
	}
	return *o.TopicName
}

// GetTopicNameOk returns a tuple with the TopicName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputConfig) GetTopicNameOk() (*string, bool) {
	if o == nil || IsNil(o.TopicName) {
		return nil, false
	}
	return o.TopicName, true
}

// HasTopicName returns a boolean if a field has been set.
func (o *OutputConfig) HasTopicName() bool {
	if o != nil && !IsNil(o.TopicName) {
		return true
	}

	return false
}

// SetTopicName gets a reference to the given string and assigns it to the TopicName field.
func (o *OutputConfig) SetTopicName(v string) {
	o.TopicName = &v
}

// GetTopicTemplate returns the TopicTemplate field value if set, zero value otherwise.
func (o *OutputConfig) GetTopicTemplate() string {
	if o == nil || IsNil(o.TopicTemplate) {
		var ret string
		return ret
	}
	return *o.TopicTemplate
}

// GetTopicTemplateOk returns a tuple with the TopicTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputConfig) GetTopicTemplateOk() (*string, bool) {
	if o == nil || IsNil(o.TopicTemplate) {
		return nil, false
	}
	return o.TopicTemplate, true
}

// HasTopicTemplate returns a boolean if a field has been set.
func (o *OutputConfig) HasTopicTemplate() bool {
	if o != nil && !IsNil(o.TopicTemplate) {
		return true
	}

	return false
}

// SetTopicTemplate gets a reference to the given string and assigns it to the TopicTemplate field.
func (o *OutputConfig) SetTopicTemplate(v string) {
	o.TopicTemplate = &v
}

// GetPayloadFormat returns the PayloadFormat field value if set, zero value otherwise.
func (o *OutputConfig) GetPayloadFormat() string {
	if o == nil || IsNil(o.PayloadFormat) {
		var ret string
		return ret
	}
	return *o.PayloadFormat
}

// GetPayloadFormatOk returns a tuple with the PayloadFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputConfig) GetPayloadFormatOk() (*string, bool) {
	if o == nil || IsNil(o.PayloadFormat) {
		return nil, false
	}
	return o.PayloadFormat, true
}

// HasPayloadFormat returns a boolean if a field has been set.
func (o *OutputConfig) HasPayloadFormat() bool {
	if o != nil && !IsNil(o.PayloadFormat) {
		return true
	}

	return false
}

// SetPayloadFormat gets a reference to the given string and assigns it to the PayloadFormat field.
func (o *OutputConfig) SetPayloadFormat(v string) {
	o.PayloadFormat = &v
}

// GetPayloadTemplate returns the PayloadTemplate field value if set, zero value otherwise.
func (o *OutputConfig) GetPayloadTemplate() string {
	if o == nil || IsNil(o.PayloadTemplate) {
		var ret string
		return ret
	}
	return *o.PayloadTemplate
}

// GetPayloadTemplateOk returns a tuple with the PayloadTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputConfig) GetPayloadTemplateOk() (*string, bool) {
	if o == nil || IsNil(o.PayloadTemplate) {
		return nil, false
	}
	return o.PayloadTemplate, true
}

// HasPayloadTemplate returns a boolean if a field has been set.
func (o *OutputConfig) HasPayloadTemplate() bool {
	if o != nil && !IsNil(o.PayloadTemplate) {
		return true
	}

	return false
}

// SetPayloadTemplate gets a reference to the given string and assigns it to the PayloadTemplate field.
func (o *OutputConfig) SetPayloadTemplate(v string) {
	o.PayloadTemplate = &v
}

func (o OutputConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OutputConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.BasicAuthUser) {
		toSerialize["basicAuthUser"] = o.BasicAuthUser
	}
	if !IsNil(o.BasicAuthPass) {
		toSerialize["basicAuthPass"] = o.BasicAuthPass
	}
	if !IsNil(o.CustomHeaderName) {
		toSerialize["customHeaderName"] = o.CustomHeaderName
	}
	if !IsNil(o.CustomHeaderValue) {
		toSerialize["customHeaderValue"] = o.CustomHeaderValue
	}
	if !IsNil(o.Host) {
		toSerialize["host"] = o.Host
	}
	if !IsNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !IsNil(o.EventName) {
		toSerialize["eventName"] = o.EventName
	}
	if !IsNil(o.AsIsPayload) {
		toSerialize["asIsPayload"] = o.AsIsPayload
	}
	if !IsNil(o.Endpoint) {
		toSerialize["endpoint"] = o.Endpoint
	}
	if !IsNil(o.DisableCertCheck) {
		toSerialize["disableCertCheck"] = o.DisableCertCheck
	}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !IsNil(o.ClientId) {
		toSerialize["clientId"] = o.ClientId
	}
	if !IsNil(o.TopicName) {
		toSerialize["topicName"] = o.TopicName
	}
	if !IsNil(o.TopicTemplate) {
		toSerialize["topicTemplate"] = o.TopicTemplate
	}
	if !IsNil(o.PayloadFormat) {
		toSerialize["payloadFormat"] = o.PayloadFormat
	}
	if !IsNil(o.PayloadTemplate) {
		toSerialize["payloadTemplate"] = o.PayloadTemplate
	}
	return toSerialize, nil
}

type NullableOutputConfig struct {
	value *OutputConfig
	isSet bool
}

func (v NullableOutputConfig) Get() *OutputConfig {
	return v.value
}

func (v *NullableOutputConfig) Set(val *OutputConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableOutputConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableOutputConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOutputConfig(val *OutputConfig) *NullableOutputConfig {
	return &NullableOutputConfig{value: val, isSet: true}
}

func (v NullableOutputConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOutputConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


