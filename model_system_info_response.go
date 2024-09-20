/*
The Span API

API for device, collection, output and firmware management

API version: 5.0.2 subversive-jamila
Contact: dev@lab5e.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package spanapi

import (
	"encoding/json"
)

// checks if the SystemInfoResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SystemInfoResponse{}

// SystemInfoResponse Response object for system information. This contains system-level information.
type SystemInfoResponse struct {
	// This is the system version
	Version *string `json:"version,omitempty"`
	// The build time for this version.
	BuildDate *string `json:"buildDate,omitempty"`
	// Human-readable code name for this release. This can be easier to remember than the version number.
	ReleaseName *string `json:"releaseName,omitempty"`
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
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInfoResponse) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *SystemInfoResponse) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
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
	if o == nil || IsNil(o.BuildDate) {
		var ret string
		return ret
	}
	return *o.BuildDate
}

// GetBuildDateOk returns a tuple with the BuildDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInfoResponse) GetBuildDateOk() (*string, bool) {
	if o == nil || IsNil(o.BuildDate) {
		return nil, false
	}
	return o.BuildDate, true
}

// HasBuildDate returns a boolean if a field has been set.
func (o *SystemInfoResponse) HasBuildDate() bool {
	if o != nil && !IsNil(o.BuildDate) {
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
	if o == nil || IsNil(o.ReleaseName) {
		var ret string
		return ret
	}
	return *o.ReleaseName
}

// GetReleaseNameOk returns a tuple with the ReleaseName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInfoResponse) GetReleaseNameOk() (*string, bool) {
	if o == nil || IsNil(o.ReleaseName) {
		return nil, false
	}
	return o.ReleaseName, true
}

// HasReleaseName returns a boolean if a field has been set.
func (o *SystemInfoResponse) HasReleaseName() bool {
	if o != nil && !IsNil(o.ReleaseName) {
		return true
	}

	return false
}

// SetReleaseName gets a reference to the given string and assigns it to the ReleaseName field.
func (o *SystemInfoResponse) SetReleaseName(v string) {
	o.ReleaseName = &v
}

func (o SystemInfoResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SystemInfoResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !IsNil(o.BuildDate) {
		toSerialize["buildDate"] = o.BuildDate
	}
	if !IsNil(o.ReleaseName) {
		toSerialize["releaseName"] = o.ReleaseName
	}
	return toSerialize, nil
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


