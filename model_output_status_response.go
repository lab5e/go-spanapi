/*
The Span API

API for device, collection, output and firmware management

API version: 4.2.2 smarmy-derik
Contact: dev@lab5e.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package spanapi

import (
	"encoding/json"
)

// OutputStatusResponse Show status of output
type OutputStatusResponse struct {
	CollectionId *string `json:"collectionId,omitempty"`
	OutputId     *string `json:"outputId,omitempty"`
	Enabled      *bool   `json:"enabled,omitempty"`
	ErrorCount   *int32  `json:"errorCount,omitempty"`
	Forwarded    *int32  `json:"forwarded,omitempty"`
	Received     *int32  `json:"received,omitempty"`
	Retransmits  *int32  `json:"retransmits,omitempty"`
}

// NewOutputStatusResponse instantiates a new OutputStatusResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOutputStatusResponse() *OutputStatusResponse {
	this := OutputStatusResponse{}
	return &this
}

// NewOutputStatusResponseWithDefaults instantiates a new OutputStatusResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOutputStatusResponseWithDefaults() *OutputStatusResponse {
	this := OutputStatusResponse{}
	return &this
}

// GetCollectionId returns the CollectionId field value if set, zero value otherwise.
func (o *OutputStatusResponse) GetCollectionId() string {
	if o == nil || o.CollectionId == nil {
		var ret string
		return ret
	}
	return *o.CollectionId
}

// GetCollectionIdOk returns a tuple with the CollectionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputStatusResponse) GetCollectionIdOk() (*string, bool) {
	if o == nil || o.CollectionId == nil {
		return nil, false
	}
	return o.CollectionId, true
}

// HasCollectionId returns a boolean if a field has been set.
func (o *OutputStatusResponse) HasCollectionId() bool {
	if o != nil && o.CollectionId != nil {
		return true
	}

	return false
}

// SetCollectionId gets a reference to the given string and assigns it to the CollectionId field.
func (o *OutputStatusResponse) SetCollectionId(v string) {
	o.CollectionId = &v
}

// GetOutputId returns the OutputId field value if set, zero value otherwise.
func (o *OutputStatusResponse) GetOutputId() string {
	if o == nil || o.OutputId == nil {
		var ret string
		return ret
	}
	return *o.OutputId
}

// GetOutputIdOk returns a tuple with the OutputId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputStatusResponse) GetOutputIdOk() (*string, bool) {
	if o == nil || o.OutputId == nil {
		return nil, false
	}
	return o.OutputId, true
}

// HasOutputId returns a boolean if a field has been set.
func (o *OutputStatusResponse) HasOutputId() bool {
	if o != nil && o.OutputId != nil {
		return true
	}

	return false
}

// SetOutputId gets a reference to the given string and assigns it to the OutputId field.
func (o *OutputStatusResponse) SetOutputId(v string) {
	o.OutputId = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *OutputStatusResponse) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputStatusResponse) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *OutputStatusResponse) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *OutputStatusResponse) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetErrorCount returns the ErrorCount field value if set, zero value otherwise.
func (o *OutputStatusResponse) GetErrorCount() int32 {
	if o == nil || o.ErrorCount == nil {
		var ret int32
		return ret
	}
	return *o.ErrorCount
}

// GetErrorCountOk returns a tuple with the ErrorCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputStatusResponse) GetErrorCountOk() (*int32, bool) {
	if o == nil || o.ErrorCount == nil {
		return nil, false
	}
	return o.ErrorCount, true
}

// HasErrorCount returns a boolean if a field has been set.
func (o *OutputStatusResponse) HasErrorCount() bool {
	if o != nil && o.ErrorCount != nil {
		return true
	}

	return false
}

// SetErrorCount gets a reference to the given int32 and assigns it to the ErrorCount field.
func (o *OutputStatusResponse) SetErrorCount(v int32) {
	o.ErrorCount = &v
}

// GetForwarded returns the Forwarded field value if set, zero value otherwise.
func (o *OutputStatusResponse) GetForwarded() int32 {
	if o == nil || o.Forwarded == nil {
		var ret int32
		return ret
	}
	return *o.Forwarded
}

// GetForwardedOk returns a tuple with the Forwarded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputStatusResponse) GetForwardedOk() (*int32, bool) {
	if o == nil || o.Forwarded == nil {
		return nil, false
	}
	return o.Forwarded, true
}

// HasForwarded returns a boolean if a field has been set.
func (o *OutputStatusResponse) HasForwarded() bool {
	if o != nil && o.Forwarded != nil {
		return true
	}

	return false
}

// SetForwarded gets a reference to the given int32 and assigns it to the Forwarded field.
func (o *OutputStatusResponse) SetForwarded(v int32) {
	o.Forwarded = &v
}

// GetReceived returns the Received field value if set, zero value otherwise.
func (o *OutputStatusResponse) GetReceived() int32 {
	if o == nil || o.Received == nil {
		var ret int32
		return ret
	}
	return *o.Received
}

// GetReceivedOk returns a tuple with the Received field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputStatusResponse) GetReceivedOk() (*int32, bool) {
	if o == nil || o.Received == nil {
		return nil, false
	}
	return o.Received, true
}

// HasReceived returns a boolean if a field has been set.
func (o *OutputStatusResponse) HasReceived() bool {
	if o != nil && o.Received != nil {
		return true
	}

	return false
}

// SetReceived gets a reference to the given int32 and assigns it to the Received field.
func (o *OutputStatusResponse) SetReceived(v int32) {
	o.Received = &v
}

// GetRetransmits returns the Retransmits field value if set, zero value otherwise.
func (o *OutputStatusResponse) GetRetransmits() int32 {
	if o == nil || o.Retransmits == nil {
		var ret int32
		return ret
	}
	return *o.Retransmits
}

// GetRetransmitsOk returns a tuple with the Retransmits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputStatusResponse) GetRetransmitsOk() (*int32, bool) {
	if o == nil || o.Retransmits == nil {
		return nil, false
	}
	return o.Retransmits, true
}

// HasRetransmits returns a boolean if a field has been set.
func (o *OutputStatusResponse) HasRetransmits() bool {
	if o != nil && o.Retransmits != nil {
		return true
	}

	return false
}

// SetRetransmits gets a reference to the given int32 and assigns it to the Retransmits field.
func (o *OutputStatusResponse) SetRetransmits(v int32) {
	o.Retransmits = &v
}

func (o OutputStatusResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CollectionId != nil {
		toSerialize["collectionId"] = o.CollectionId
	}
	if o.OutputId != nil {
		toSerialize["outputId"] = o.OutputId
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.ErrorCount != nil {
		toSerialize["errorCount"] = o.ErrorCount
	}
	if o.Forwarded != nil {
		toSerialize["forwarded"] = o.Forwarded
	}
	if o.Received != nil {
		toSerialize["received"] = o.Received
	}
	if o.Retransmits != nil {
		toSerialize["retransmits"] = o.Retransmits
	}
	return json.Marshal(toSerialize)
}

type NullableOutputStatusResponse struct {
	value *OutputStatusResponse
	isSet bool
}

func (v NullableOutputStatusResponse) Get() *OutputStatusResponse {
	return v.value
}

func (v *NullableOutputStatusResponse) Set(val *OutputStatusResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOutputStatusResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOutputStatusResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOutputStatusResponse(val *OutputStatusResponse) *NullableOutputStatusResponse {
	return &NullableOutputStatusResponse{value: val, isSet: true}
}

func (v NullableOutputStatusResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOutputStatusResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
