/*
The Span API

API for device, collection, output and firmware management

API version: 5.0.3 pitch-dark-elza
Contact: dev@lab5e.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package spanapi

import (
	"encoding/json"
)

// checks if the Blob type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Blob{}

// Blob This is a blob (binary large object) that the devices might upload to the service. This is messages that are typically too large to handle like regular status and sensor values, typically media files. The content type might be derived from the first few bytes of the blob and could possibly be incorrect.  Download the blob by accessing the blob URL field. This will work like a regular HTTP request for your client. Authentication is required as always.
type Blob struct {
	BlobId *string `json:"blobId,omitempty"`
	BlobPath *string `json:"blobPath,omitempty"`
	ContentType *string `json:"contentType,omitempty"`
	Size *string `json:"size,omitempty"`
	Created *string `json:"created,omitempty"`
	CollectionId *string `json:"collectionId,omitempty"`
	DeviceId *string `json:"deviceId,omitempty"`
	GatewayId *string `json:"gatewayId,omitempty"`
	Properties *map[string]string `json:"properties,omitempty"`
}

// NewBlob instantiates a new Blob object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlob() *Blob {
	this := Blob{}
	return &this
}

// NewBlobWithDefaults instantiates a new Blob object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlobWithDefaults() *Blob {
	this := Blob{}
	return &this
}

// GetBlobId returns the BlobId field value if set, zero value otherwise.
func (o *Blob) GetBlobId() string {
	if o == nil || IsNil(o.BlobId) {
		var ret string
		return ret
	}
	return *o.BlobId
}

// GetBlobIdOk returns a tuple with the BlobId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Blob) GetBlobIdOk() (*string, bool) {
	if o == nil || IsNil(o.BlobId) {
		return nil, false
	}
	return o.BlobId, true
}

// HasBlobId returns a boolean if a field has been set.
func (o *Blob) HasBlobId() bool {
	if o != nil && !IsNil(o.BlobId) {
		return true
	}

	return false
}

// SetBlobId gets a reference to the given string and assigns it to the BlobId field.
func (o *Blob) SetBlobId(v string) {
	o.BlobId = &v
}

// GetBlobPath returns the BlobPath field value if set, zero value otherwise.
func (o *Blob) GetBlobPath() string {
	if o == nil || IsNil(o.BlobPath) {
		var ret string
		return ret
	}
	return *o.BlobPath
}

// GetBlobPathOk returns a tuple with the BlobPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Blob) GetBlobPathOk() (*string, bool) {
	if o == nil || IsNil(o.BlobPath) {
		return nil, false
	}
	return o.BlobPath, true
}

// HasBlobPath returns a boolean if a field has been set.
func (o *Blob) HasBlobPath() bool {
	if o != nil && !IsNil(o.BlobPath) {
		return true
	}

	return false
}

// SetBlobPath gets a reference to the given string and assigns it to the BlobPath field.
func (o *Blob) SetBlobPath(v string) {
	o.BlobPath = &v
}

// GetContentType returns the ContentType field value if set, zero value otherwise.
func (o *Blob) GetContentType() string {
	if o == nil || IsNil(o.ContentType) {
		var ret string
		return ret
	}
	return *o.ContentType
}

// GetContentTypeOk returns a tuple with the ContentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Blob) GetContentTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ContentType) {
		return nil, false
	}
	return o.ContentType, true
}

// HasContentType returns a boolean if a field has been set.
func (o *Blob) HasContentType() bool {
	if o != nil && !IsNil(o.ContentType) {
		return true
	}

	return false
}

// SetContentType gets a reference to the given string and assigns it to the ContentType field.
func (o *Blob) SetContentType(v string) {
	o.ContentType = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *Blob) GetSize() string {
	if o == nil || IsNil(o.Size) {
		var ret string
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Blob) GetSizeOk() (*string, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *Blob) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given string and assigns it to the Size field.
func (o *Blob) SetSize(v string) {
	o.Size = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *Blob) GetCreated() string {
	if o == nil || IsNil(o.Created) {
		var ret string
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Blob) GetCreatedOk() (*string, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *Blob) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given string and assigns it to the Created field.
func (o *Blob) SetCreated(v string) {
	o.Created = &v
}

// GetCollectionId returns the CollectionId field value if set, zero value otherwise.
func (o *Blob) GetCollectionId() string {
	if o == nil || IsNil(o.CollectionId) {
		var ret string
		return ret
	}
	return *o.CollectionId
}

// GetCollectionIdOk returns a tuple with the CollectionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Blob) GetCollectionIdOk() (*string, bool) {
	if o == nil || IsNil(o.CollectionId) {
		return nil, false
	}
	return o.CollectionId, true
}

// HasCollectionId returns a boolean if a field has been set.
func (o *Blob) HasCollectionId() bool {
	if o != nil && !IsNil(o.CollectionId) {
		return true
	}

	return false
}

// SetCollectionId gets a reference to the given string and assigns it to the CollectionId field.
func (o *Blob) SetCollectionId(v string) {
	o.CollectionId = &v
}

// GetDeviceId returns the DeviceId field value if set, zero value otherwise.
func (o *Blob) GetDeviceId() string {
	if o == nil || IsNil(o.DeviceId) {
		var ret string
		return ret
	}
	return *o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Blob) GetDeviceIdOk() (*string, bool) {
	if o == nil || IsNil(o.DeviceId) {
		return nil, false
	}
	return o.DeviceId, true
}

// HasDeviceId returns a boolean if a field has been set.
func (o *Blob) HasDeviceId() bool {
	if o != nil && !IsNil(o.DeviceId) {
		return true
	}

	return false
}

// SetDeviceId gets a reference to the given string and assigns it to the DeviceId field.
func (o *Blob) SetDeviceId(v string) {
	o.DeviceId = &v
}

// GetGatewayId returns the GatewayId field value if set, zero value otherwise.
func (o *Blob) GetGatewayId() string {
	if o == nil || IsNil(o.GatewayId) {
		var ret string
		return ret
	}
	return *o.GatewayId
}

// GetGatewayIdOk returns a tuple with the GatewayId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Blob) GetGatewayIdOk() (*string, bool) {
	if o == nil || IsNil(o.GatewayId) {
		return nil, false
	}
	return o.GatewayId, true
}

// HasGatewayId returns a boolean if a field has been set.
func (o *Blob) HasGatewayId() bool {
	if o != nil && !IsNil(o.GatewayId) {
		return true
	}

	return false
}

// SetGatewayId gets a reference to the given string and assigns it to the GatewayId field.
func (o *Blob) SetGatewayId(v string) {
	o.GatewayId = &v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *Blob) GetProperties() map[string]string {
	if o == nil || IsNil(o.Properties) {
		var ret map[string]string
		return ret
	}
	return *o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Blob) GetPropertiesOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Properties) {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *Blob) HasProperties() bool {
	if o != nil && !IsNil(o.Properties) {
		return true
	}

	return false
}

// SetProperties gets a reference to the given map[string]string and assigns it to the Properties field.
func (o *Blob) SetProperties(v map[string]string) {
	o.Properties = &v
}

func (o Blob) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Blob) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BlobId) {
		toSerialize["blobId"] = o.BlobId
	}
	if !IsNil(o.BlobPath) {
		toSerialize["blobPath"] = o.BlobPath
	}
	if !IsNil(o.ContentType) {
		toSerialize["contentType"] = o.ContentType
	}
	if !IsNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !IsNil(o.CollectionId) {
		toSerialize["collectionId"] = o.CollectionId
	}
	if !IsNil(o.DeviceId) {
		toSerialize["deviceId"] = o.DeviceId
	}
	if !IsNil(o.GatewayId) {
		toSerialize["gatewayId"] = o.GatewayId
	}
	if !IsNil(o.Properties) {
		toSerialize["properties"] = o.Properties
	}
	return toSerialize, nil
}

type NullableBlob struct {
	value *Blob
	isSet bool
}

func (v NullableBlob) Get() *Blob {
	return v.value
}

func (v *NullableBlob) Set(val *Blob) {
	v.value = val
	v.isSet = true
}

func (v NullableBlob) IsSet() bool {
	return v.isSet
}

func (v *NullableBlob) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlob(val *Blob) *NullableBlob {
	return &NullableBlob{value: val, isSet: true}
}

func (v NullableBlob) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlob) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


