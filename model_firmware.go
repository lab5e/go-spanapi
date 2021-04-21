/*
 * The Span API
 *
 * API for device, collection, output and firmware management
 *
 * API version: 4.1.9 receding-glennis
 * Contact: dev@lab5e.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package spanapi

import (
	"encoding/json"
)

// Firmware Firmware images aren't served back out through the API, only the metadata.
type Firmware struct {
	ImageId *string `json:"imageId,omitempty"`
	// Make sure that the firmware image reports this version. If the version reported by this image is different the FOTA process won't be reported as successful since the device will report a version different from this.
	Version *string `json:"version,omitempty"`
	// This is just for internal house keeping, making it potentially easier to identify the firmware image.
	Filename *string `json:"filename,omitempty"`
	// To ensure each image is unique the SHA-256 hash for all images in a collection must be unqique
	Sha256 *string `json:"sha256,omitempty"`
	Length *int32 `json:"length,omitempty"`
	// Collection ID for the collection owning the firmware image.
	CollectionId *string `json:"collectionId,omitempty"`
	Created *string `json:"created,omitempty"`
	// Tags for firmware image.
	Tags *map[string]string `json:"tags,omitempty"`
}

// NewFirmware instantiates a new Firmware object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFirmware() *Firmware {
	this := Firmware{}
	return &this
}

// NewFirmwareWithDefaults instantiates a new Firmware object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFirmwareWithDefaults() *Firmware {
	this := Firmware{}
	return &this
}

// GetImageId returns the ImageId field value if set, zero value otherwise.
func (o *Firmware) GetImageId() string {
	if o == nil || o.ImageId == nil {
		var ret string
		return ret
	}
	return *o.ImageId
}

// GetImageIdOk returns a tuple with the ImageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Firmware) GetImageIdOk() (*string, bool) {
	if o == nil || o.ImageId == nil {
		return nil, false
	}
	return o.ImageId, true
}

// HasImageId returns a boolean if a field has been set.
func (o *Firmware) HasImageId() bool {
	if o != nil && o.ImageId != nil {
		return true
	}

	return false
}

// SetImageId gets a reference to the given string and assigns it to the ImageId field.
func (o *Firmware) SetImageId(v string) {
	o.ImageId = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *Firmware) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Firmware) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *Firmware) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *Firmware) SetVersion(v string) {
	o.Version = &v
}

// GetFilename returns the Filename field value if set, zero value otherwise.
func (o *Firmware) GetFilename() string {
	if o == nil || o.Filename == nil {
		var ret string
		return ret
	}
	return *o.Filename
}

// GetFilenameOk returns a tuple with the Filename field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Firmware) GetFilenameOk() (*string, bool) {
	if o == nil || o.Filename == nil {
		return nil, false
	}
	return o.Filename, true
}

// HasFilename returns a boolean if a field has been set.
func (o *Firmware) HasFilename() bool {
	if o != nil && o.Filename != nil {
		return true
	}

	return false
}

// SetFilename gets a reference to the given string and assigns it to the Filename field.
func (o *Firmware) SetFilename(v string) {
	o.Filename = &v
}

// GetSha256 returns the Sha256 field value if set, zero value otherwise.
func (o *Firmware) GetSha256() string {
	if o == nil || o.Sha256 == nil {
		var ret string
		return ret
	}
	return *o.Sha256
}

// GetSha256Ok returns a tuple with the Sha256 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Firmware) GetSha256Ok() (*string, bool) {
	if o == nil || o.Sha256 == nil {
		return nil, false
	}
	return o.Sha256, true
}

// HasSha256 returns a boolean if a field has been set.
func (o *Firmware) HasSha256() bool {
	if o != nil && o.Sha256 != nil {
		return true
	}

	return false
}

// SetSha256 gets a reference to the given string and assigns it to the Sha256 field.
func (o *Firmware) SetSha256(v string) {
	o.Sha256 = &v
}

// GetLength returns the Length field value if set, zero value otherwise.
func (o *Firmware) GetLength() int32 {
	if o == nil || o.Length == nil {
		var ret int32
		return ret
	}
	return *o.Length
}

// GetLengthOk returns a tuple with the Length field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Firmware) GetLengthOk() (*int32, bool) {
	if o == nil || o.Length == nil {
		return nil, false
	}
	return o.Length, true
}

// HasLength returns a boolean if a field has been set.
func (o *Firmware) HasLength() bool {
	if o != nil && o.Length != nil {
		return true
	}

	return false
}

// SetLength gets a reference to the given int32 and assigns it to the Length field.
func (o *Firmware) SetLength(v int32) {
	o.Length = &v
}

// GetCollectionId returns the CollectionId field value if set, zero value otherwise.
func (o *Firmware) GetCollectionId() string {
	if o == nil || o.CollectionId == nil {
		var ret string
		return ret
	}
	return *o.CollectionId
}

// GetCollectionIdOk returns a tuple with the CollectionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Firmware) GetCollectionIdOk() (*string, bool) {
	if o == nil || o.CollectionId == nil {
		return nil, false
	}
	return o.CollectionId, true
}

// HasCollectionId returns a boolean if a field has been set.
func (o *Firmware) HasCollectionId() bool {
	if o != nil && o.CollectionId != nil {
		return true
	}

	return false
}

// SetCollectionId gets a reference to the given string and assigns it to the CollectionId field.
func (o *Firmware) SetCollectionId(v string) {
	o.CollectionId = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *Firmware) GetCreated() string {
	if o == nil || o.Created == nil {
		var ret string
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Firmware) GetCreatedOk() (*string, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *Firmware) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given string and assigns it to the Created field.
func (o *Firmware) SetCreated(v string) {
	o.Created = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *Firmware) GetTags() map[string]string {
	if o == nil || o.Tags == nil {
		var ret map[string]string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Firmware) GetTagsOk() (*map[string]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *Firmware) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given map[string]string and assigns it to the Tags field.
func (o *Firmware) SetTags(v map[string]string) {
	o.Tags = &v
}

func (o Firmware) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ImageId != nil {
		toSerialize["imageId"] = o.ImageId
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.Filename != nil {
		toSerialize["filename"] = o.Filename
	}
	if o.Sha256 != nil {
		toSerialize["sha256"] = o.Sha256
	}
	if o.Length != nil {
		toSerialize["length"] = o.Length
	}
	if o.CollectionId != nil {
		toSerialize["collectionId"] = o.CollectionId
	}
	if o.Created != nil {
		toSerialize["created"] = o.Created
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	return json.Marshal(toSerialize)
}

type NullableFirmware struct {
	value *Firmware
	isSet bool
}

func (v NullableFirmware) Get() *Firmware {
	return v.value
}

func (v *NullableFirmware) Set(val *Firmware) {
	v.value = val
	v.isSet = true
}

func (v NullableFirmware) IsSet() bool {
	return v.isSet
}

func (v *NullableFirmware) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFirmware(val *Firmware) *NullableFirmware {
	return &NullableFirmware{value: val, isSet: true}
}

func (v NullableFirmware) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFirmware) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


