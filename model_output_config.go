/*
The Span API

API for device, collection, output and firmware management

API version: 4.3.0 grouchy-aloysius
Contact: dev@lab5e.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package spanapi

import (
	"encoding/json"
)

// OutputConfig Configuration for outputs.
type OutputConfig struct {
	// URL for the webhook.
	Url               *string `json:"url,omitempty"`
	BasicAuthUser     *string `json:"basicAuthUser,omitempty"`
	BasicAuthPass     *string `json:"basicAuthPass,omitempty"`
	CustomHeaderName  *string `json:"customHeaderName,omitempty"`
	CustomHeaderValue *string `json:"customHeaderValue,omitempty"`
	Host              *string `json:"host,omitempty"`
	Port              *int32  `json:"port,omitempty"`
	Key               *string `json:"key,omitempty"`
	EventName         *string `json:"eventName,omitempty"`
	AsIsPayload       *bool   `json:"asIsPayload,omitempty"`
	Endpoint          *string `json:"endpoint,omitempty"`
	// MQTT configuration: Disable certificate checks. Default is off.
	DisableCertCheck *bool   `json:"disableCertCheck,omitempty"`
	Username         *string `json:"username,omitempty"`
	Password         *string `json:"password,omitempty"`
	ClientId         *string `json:"clientId,omitempty"`
	TopicName        *string `json:"topicName,omitempty"`
	TopicTemplate    *string `json:"topicTemplate,omitempty"`
	PayloadFormat    *string `json:"payloadFormat,omitempty"`
	PayloadTemplate  *string `json:"payloadTemplate,omitempty"`
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
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputConfig) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *OutputConfig) HasUrl() bool {
	if o != nil && o.Url != nil {
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
	if o == nil || o.BasicAuthUser == nil {
		var ret string
		return ret
	}
	return *o.BasicAuthUser
}

// GetBasicAuthUserOk returns a tuple with the BasicAuthUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputConfig) GetBasicAuthUserOk() (*string, bool) {
	if o == nil || o.BasicAuthUser == nil {
		return nil, false
	}
	return o.BasicAuthUser, true
}

// HasBasicAuthUser returns a boolean if a field has been set.
func (o *OutputConfig) HasBasicAuthUser() bool {
	if o != nil && o.BasicAuthUser != nil {
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
	if o == nil || o.BasicAuthPass == nil {
		var ret string
		return ret
	}
	return *o.BasicAuthPass
}

// GetBasicAuthPassOk returns a tuple with the BasicAuthPass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputConfig) GetBasicAuthPassOk() (*string, bool) {
	if o == nil || o.BasicAuthPass == nil {
		return nil, false
	}
	return o.BasicAuthPass, true
}

// HasBasicAuthPass returns a boolean if a field has been set.
func (o *OutputConfig) HasBasicAuthPass() bool {
	if o != nil && o.BasicAuthPass != nil {
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
	if o == nil || o.CustomHeaderName == nil {
		var ret string
		return ret
	}
	return *o.CustomHeaderName
}

// GetCustomHeaderNameOk returns a tuple with the CustomHeaderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputConfig) GetCustomHeaderNameOk() (*string, bool) {
	if o == nil || o.CustomHeaderName == nil {
		return nil, false
	}
	return o.CustomHeaderName, true
}

// HasCustomHeaderName returns a boolean if a field has been set.
func (o *OutputConfig) HasCustomHeaderName() bool {
	if o != nil && o.CustomHeaderName != nil {
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
	if o == nil || o.CustomHeaderValue == nil {
		var ret string
		return ret
	}
	return *o.CustomHeaderValue
}

// GetCustomHeaderValueOk returns a tuple with the CustomHeaderValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputConfig) GetCustomHeaderValueOk() (*string, bool) {
	if o == nil || o.CustomHeaderValue == nil {
		return nil, false
	}
	return o.CustomHeaderValue, true
}

// HasCustomHeaderValue returns a boolean if a field has been set.
func (o *OutputConfig) HasCustomHeaderValue() bool {
	if o != nil && o.CustomHeaderValue != nil {
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
	if o == nil || o.Host == nil {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputConfig) GetHostOk() (*string, bool) {
	if o == nil || o.Host == nil {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *OutputConfig) HasHost() bool {
	if o != nil && o.Host != nil {
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
	if o == nil || o.Port == nil {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputConfig) GetPortOk() (*int32, bool) {
	if o == nil || o.Port == nil {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *OutputConfig) HasPort() bool {
	if o != nil && o.Port != nil {
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
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputConfig) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *OutputConfig) HasKey() bool {
	if o != nil && o.Key != nil {
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
	if o == nil || o.EventName == nil {
		var ret string
		return ret
	}
	return *o.EventName
}

// GetEventNameOk returns a tuple with the EventName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputConfig) GetEventNameOk() (*string, bool) {
	if o == nil || o.EventName == nil {
		return nil, false
	}
	return o.EventName, true
}

// HasEventName returns a boolean if a field has been set.
func (o *OutputConfig) HasEventName() bool {
	if o != nil && o.EventName != nil {
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
	if o == nil || o.AsIsPayload == nil {
		var ret bool
		return ret
	}
	return *o.AsIsPayload
}

// GetAsIsPayloadOk returns a tuple with the AsIsPayload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputConfig) GetAsIsPayloadOk() (*bool, bool) {
	if o == nil || o.AsIsPayload == nil {
		return nil, false
	}
	return o.AsIsPayload, true
}

// HasAsIsPayload returns a boolean if a field has been set.
func (o *OutputConfig) HasAsIsPayload() bool {
	if o != nil && o.AsIsPayload != nil {
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
	if o == nil || o.Endpoint == nil {
		var ret string
		return ret
	}
	return *o.Endpoint
}

// GetEndpointOk returns a tuple with the Endpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputConfig) GetEndpointOk() (*string, bool) {
	if o == nil || o.Endpoint == nil {
		return nil, false
	}
	return o.Endpoint, true
}

// HasEndpoint returns a boolean if a field has been set.
func (o *OutputConfig) HasEndpoint() bool {
	if o != nil && o.Endpoint != nil {
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
	if o == nil || o.DisableCertCheck == nil {
		var ret bool
		return ret
	}
	return *o.DisableCertCheck
}

// GetDisableCertCheckOk returns a tuple with the DisableCertCheck field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputConfig) GetDisableCertCheckOk() (*bool, bool) {
	if o == nil || o.DisableCertCheck == nil {
		return nil, false
	}
	return o.DisableCertCheck, true
}

// HasDisableCertCheck returns a boolean if a field has been set.
func (o *OutputConfig) HasDisableCertCheck() bool {
	if o != nil && o.DisableCertCheck != nil {
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
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputConfig) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *OutputConfig) HasUsername() bool {
	if o != nil && o.Username != nil {
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
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputConfig) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *OutputConfig) HasPassword() bool {
	if o != nil && o.Password != nil {
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
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputConfig) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *OutputConfig) HasClientId() bool {
	if o != nil && o.ClientId != nil {
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
	if o == nil || o.TopicName == nil {
		var ret string
		return ret
	}
	return *o.TopicName
}

// GetTopicNameOk returns a tuple with the TopicName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputConfig) GetTopicNameOk() (*string, bool) {
	if o == nil || o.TopicName == nil {
		return nil, false
	}
	return o.TopicName, true
}

// HasTopicName returns a boolean if a field has been set.
func (o *OutputConfig) HasTopicName() bool {
	if o != nil && o.TopicName != nil {
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
	if o == nil || o.TopicTemplate == nil {
		var ret string
		return ret
	}
	return *o.TopicTemplate
}

// GetTopicTemplateOk returns a tuple with the TopicTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputConfig) GetTopicTemplateOk() (*string, bool) {
	if o == nil || o.TopicTemplate == nil {
		return nil, false
	}
	return o.TopicTemplate, true
}

// HasTopicTemplate returns a boolean if a field has been set.
func (o *OutputConfig) HasTopicTemplate() bool {
	if o != nil && o.TopicTemplate != nil {
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
	if o == nil || o.PayloadFormat == nil {
		var ret string
		return ret
	}
	return *o.PayloadFormat
}

// GetPayloadFormatOk returns a tuple with the PayloadFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputConfig) GetPayloadFormatOk() (*string, bool) {
	if o == nil || o.PayloadFormat == nil {
		return nil, false
	}
	return o.PayloadFormat, true
}

// HasPayloadFormat returns a boolean if a field has been set.
func (o *OutputConfig) HasPayloadFormat() bool {
	if o != nil && o.PayloadFormat != nil {
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
	if o == nil || o.PayloadTemplate == nil {
		var ret string
		return ret
	}
	return *o.PayloadTemplate
}

// GetPayloadTemplateOk returns a tuple with the PayloadTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputConfig) GetPayloadTemplateOk() (*string, bool) {
	if o == nil || o.PayloadTemplate == nil {
		return nil, false
	}
	return o.PayloadTemplate, true
}

// HasPayloadTemplate returns a boolean if a field has been set.
func (o *OutputConfig) HasPayloadTemplate() bool {
	if o != nil && o.PayloadTemplate != nil {
		return true
	}

	return false
}

// SetPayloadTemplate gets a reference to the given string and assigns it to the PayloadTemplate field.
func (o *OutputConfig) SetPayloadTemplate(v string) {
	o.PayloadTemplate = &v
}

func (o OutputConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	if o.BasicAuthUser != nil {
		toSerialize["basicAuthUser"] = o.BasicAuthUser
	}
	if o.BasicAuthPass != nil {
		toSerialize["basicAuthPass"] = o.BasicAuthPass
	}
	if o.CustomHeaderName != nil {
		toSerialize["customHeaderName"] = o.CustomHeaderName
	}
	if o.CustomHeaderValue != nil {
		toSerialize["customHeaderValue"] = o.CustomHeaderValue
	}
	if o.Host != nil {
		toSerialize["host"] = o.Host
	}
	if o.Port != nil {
		toSerialize["port"] = o.Port
	}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	if o.EventName != nil {
		toSerialize["eventName"] = o.EventName
	}
	if o.AsIsPayload != nil {
		toSerialize["asIsPayload"] = o.AsIsPayload
	}
	if o.Endpoint != nil {
		toSerialize["endpoint"] = o.Endpoint
	}
	if o.DisableCertCheck != nil {
		toSerialize["disableCertCheck"] = o.DisableCertCheck
	}
	if o.Username != nil {
		toSerialize["username"] = o.Username
	}
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}
	if o.ClientId != nil {
		toSerialize["clientId"] = o.ClientId
	}
	if o.TopicName != nil {
		toSerialize["topicName"] = o.TopicName
	}
	if o.TopicTemplate != nil {
		toSerialize["topicTemplate"] = o.TopicTemplate
	}
	if o.PayloadFormat != nil {
		toSerialize["payloadFormat"] = o.PayloadFormat
	}
	if o.PayloadTemplate != nil {
		toSerialize["payloadTemplate"] = o.PayloadTemplate
	}
	return json.Marshal(toSerialize)
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
