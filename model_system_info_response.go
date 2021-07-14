/*
 * The Span API
 *
 * API for device, collection, output and firmware management
 *
 * API version: 4.1.16 spooky-devante
 * Contact: dev@lab5e.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package spanapi

import (
	"encoding/json"
)

// SystemInfoResponse struct for SystemInfoResponse
type SystemInfoResponse struct {
	Version          *string    `json:"version,omitempty"`
	BuildDate        *string    `json:"buildDate,omitempty"`
	ReleaseName      *string    `json:"releaseName,omitempty"`
	DefaultFieldMask *FieldMask `json:"defaultFieldMask,omitempty"`
	ForcedFieldMask  *FieldMask `json:"forcedFieldMask,omitempty"`
}

// NewSystemInfoResponse instantiates a new SystemInfoResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSystemInfoResponse() *SystemInfoResponse {
	this := SystemInfoResponse{}
	return &this
}

// NewSystemInfoResponseWithDefaults instantiates a new SystemInfoResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSystemInfoResponseWithDefaults() *SystemInfoResponse {
	this := SystemInfoResponse{}
	return &this
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *SystemInfoResponse) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInfoResponse) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *SystemInfoResponse) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *SystemInfoResponse) SetVersion(v string) {
	o.Version = &v
}

// GetBuildDate returns the BuildDate field value if set, zero value otherwise.
func (o *SystemInfoResponse) GetBuildDate() string {
	if o == nil || o.BuildDate == nil {
		var ret string
		return ret
	}
	return *o.BuildDate
}

// GetBuildDateOk returns a tuple with the BuildDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInfoResponse) GetBuildDateOk() (*string, bool) {
	if o == nil || o.BuildDate == nil {
		return nil, false
	}
	return o.BuildDate, true
}

// HasBuildDate returns a boolean if a field has been set.
func (o *SystemInfoResponse) HasBuildDate() bool {
	if o != nil && o.BuildDate != nil {
		return true
	}

	return false
}

// SetBuildDate gets a reference to the given string and assigns it to the BuildDate field.
func (o *SystemInfoResponse) SetBuildDate(v string) {
	o.BuildDate = &v
}

// GetReleaseName returns the ReleaseName field value if set, zero value otherwise.
func (o *SystemInfoResponse) GetReleaseName() string {
	if o == nil || o.ReleaseName == nil {
		var ret string
		return ret
	}
	return *o.ReleaseName
}

// GetReleaseNameOk returns a tuple with the ReleaseName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInfoResponse) GetReleaseNameOk() (*string, bool) {
	if o == nil || o.ReleaseName == nil {
		return nil, false
	}
	return o.ReleaseName, true
}

// HasReleaseName returns a boolean if a field has been set.
func (o *SystemInfoResponse) HasReleaseName() bool {
	if o != nil && o.ReleaseName != nil {
		return true
	}

	return false
}

// SetReleaseName gets a reference to the given string and assigns it to the ReleaseName field.
func (o *SystemInfoResponse) SetReleaseName(v string) {
	o.ReleaseName = &v
}

// GetDefaultFieldMask returns the DefaultFieldMask field value if set, zero value otherwise.
func (o *SystemInfoResponse) GetDefaultFieldMask() FieldMask {
	if o == nil || o.DefaultFieldMask == nil {
		var ret FieldMask
		return ret
	}
	return *o.DefaultFieldMask
}

// GetDefaultFieldMaskOk returns a tuple with the DefaultFieldMask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInfoResponse) GetDefaultFieldMaskOk() (*FieldMask, bool) {
	if o == nil || o.DefaultFieldMask == nil {
		return nil, false
	}
	return o.DefaultFieldMask, true
}

// HasDefaultFieldMask returns a boolean if a field has been set.
func (o *SystemInfoResponse) HasDefaultFieldMask() bool {
	if o != nil && o.DefaultFieldMask != nil {
		return true
	}

	return false
}

// SetDefaultFieldMask gets a reference to the given FieldMask and assigns it to the DefaultFieldMask field.
func (o *SystemInfoResponse) SetDefaultFieldMask(v FieldMask) {
	o.DefaultFieldMask = &v
}

// GetForcedFieldMask returns the ForcedFieldMask field value if set, zero value otherwise.
func (o *SystemInfoResponse) GetForcedFieldMask() FieldMask {
	if o == nil || o.ForcedFieldMask == nil {
		var ret FieldMask
		return ret
	}
	return *o.ForcedFieldMask
}

// GetForcedFieldMaskOk returns a tuple with the ForcedFieldMask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInfoResponse) GetForcedFieldMaskOk() (*FieldMask, bool) {
	if o == nil || o.ForcedFieldMask == nil {
		return nil, false
	}
	return o.ForcedFieldMask, true
}

// HasForcedFieldMask returns a boolean if a field has been set.
func (o *SystemInfoResponse) HasForcedFieldMask() bool {
	if o != nil && o.ForcedFieldMask != nil {
		return true
	}

	return false
}

// SetForcedFieldMask gets a reference to the given FieldMask and assigns it to the ForcedFieldMask field.
func (o *SystemInfoResponse) SetForcedFieldMask(v FieldMask) {
	o.ForcedFieldMask = &v
}

func (o SystemInfoResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.BuildDate != nil {
		toSerialize["buildDate"] = o.BuildDate
	}
	if o.ReleaseName != nil {
		toSerialize["releaseName"] = o.ReleaseName
	}
	if o.DefaultFieldMask != nil {
		toSerialize["defaultFieldMask"] = o.DefaultFieldMask
	}
	if o.ForcedFieldMask != nil {
		toSerialize["forcedFieldMask"] = o.ForcedFieldMask
	}
	return json.Marshal(toSerialize)
}

type NullableSystemInfoResponse struct {
	value *SystemInfoResponse
	isSet bool
}

func (v NullableSystemInfoResponse) Get() *SystemInfoResponse {
	return v.value
}

func (v *NullableSystemInfoResponse) Set(val *SystemInfoResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSystemInfoResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSystemInfoResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSystemInfoResponse(val *SystemInfoResponse) *NullableSystemInfoResponse {
	return &NullableSystemInfoResponse{value: val, isSet: true}
}

func (v NullableSystemInfoResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSystemInfoResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
