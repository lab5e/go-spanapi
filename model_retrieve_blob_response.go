/*
The Span API

API for device, collection, output and firmware management

API version: 4.4.1 busy-janay
Contact: dev@lab5e.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package spanapi

import (
	"encoding/json"
)

// RetrieveBlobResponse This is not available throught the API, just as a regular HTTP response
type RetrieveBlobResponse struct {
	ContentType *string `json:"contentType,omitempty"`
	Size *int32 `json:"size,omitempty"`
	Data *string `json:"data,omitempty"`
}

// NewRetrieveBlobResponse instantiates a new RetrieveBlobResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRetrieveBlobResponse() *RetrieveBlobResponse {
	this := RetrieveBlobResponse{}
	return &this
}

// NewRetrieveBlobResponseWithDefaults instantiates a new RetrieveBlobResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRetrieveBlobResponseWithDefaults() *RetrieveBlobResponse {
	this := RetrieveBlobResponse{}
	return &this
}

// GetContentType returns the ContentType field value if set, zero value otherwise.
func (o *RetrieveBlobResponse) GetContentType() string {
	if o == nil || o.ContentType == nil {
		var ret string
		return ret
	}
	return *o.ContentType
}

// GetContentTypeOk returns a tuple with the ContentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RetrieveBlobResponse) GetContentTypeOk() (*string, bool) {
	if o == nil || o.ContentType == nil {
		return nil, false
	}
	return o.ContentType, true
}

// HasContentType returns a boolean if a field has been set.
func (o *RetrieveBlobResponse) HasContentType() bool {
	if o != nil && o.ContentType != nil {
		return true
	}

	return false
}

// SetContentType gets a reference to the given string and assigns it to the ContentType field.
func (o *RetrieveBlobResponse) SetContentType(v string) {
	o.ContentType = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *RetrieveBlobResponse) GetSize() int32 {
	if o == nil || o.Size == nil {
		var ret int32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RetrieveBlobResponse) GetSizeOk() (*int32, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *RetrieveBlobResponse) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// SetSize gets a reference to the given int32 and assigns it to the Size field.
func (o *RetrieveBlobResponse) SetSize(v int32) {
	o.Size = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *RetrieveBlobResponse) GetData() string {
	if o == nil || o.Data == nil {
		var ret string
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RetrieveBlobResponse) GetDataOk() (*string, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *RetrieveBlobResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given string and assigns it to the Data field.
func (o *RetrieveBlobResponse) SetData(v string) {
	o.Data = &v
}

func (o RetrieveBlobResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ContentType != nil {
		toSerialize["contentType"] = o.ContentType
	}
	if o.Size != nil {
		toSerialize["size"] = o.Size
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableRetrieveBlobResponse struct {
	value *RetrieveBlobResponse
	isSet bool
}

func (v NullableRetrieveBlobResponse) Get() *RetrieveBlobResponse {
	return v.value
}

func (v *NullableRetrieveBlobResponse) Set(val *RetrieveBlobResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRetrieveBlobResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRetrieveBlobResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRetrieveBlobResponse(val *RetrieveBlobResponse) *NullableRetrieveBlobResponse {
	return &NullableRetrieveBlobResponse{value: val, isSet: true}
}

func (v NullableRetrieveBlobResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRetrieveBlobResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


