/*
The Span API

API for device, collection, output and firmware management

API version: 4.6.1 squirming-codi
Contact: dev@lab5e.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package spanapi

import (
	"encoding/json"
)

// checks if the CollectionStats type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CollectionStats{}

// CollectionStats This is statistics for an collection.
type CollectionStats struct {
	DeviceCount *int32 `json:"deviceCount,omitempty"`
	OutputCount *int32 `json:"outputCount,omitempty"`
	FirmwareCount *int32 `json:"firmwareCount,omitempty"`
	BlobCount *int32 `json:"blobCount,omitempty"`
	GatewayCount *int32 `json:"gatewayCount,omitempty"`
	Devices *DeviceStats `json:"devices,omitempty"`
	Outputs *OutputStats `json:"outputs,omitempty"`
	Firmware *FirmwareStats `json:"firmware,omitempty"`
	Blobs *BlobStats `json:"blobs,omitempty"`
	Gateways *GatewayStats `json:"gateways,omitempty"`
}

// NewCollectionStats instantiates a new CollectionStats object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionStats() *CollectionStats {
	this := CollectionStats{}
	return &this
}

// NewCollectionStatsWithDefaults instantiates a new CollectionStats object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionStatsWithDefaults() *CollectionStats {
	this := CollectionStats{}
	return &this
}

// GetDeviceCount returns the DeviceCount field value if set, zero value otherwise.
func (o *CollectionStats) GetDeviceCount() int32 {
	if o == nil || IsNil(o.DeviceCount) {
		var ret int32
		return ret
	}
	return *o.DeviceCount
}

// GetDeviceCountOk returns a tuple with the DeviceCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionStats) GetDeviceCountOk() (*int32, bool) {
	if o == nil || IsNil(o.DeviceCount) {
		return nil, false
	}
	return o.DeviceCount, true
}

// HasDeviceCount returns a boolean if a field has been set.
func (o *CollectionStats) HasDeviceCount() bool {
	if o != nil && !IsNil(o.DeviceCount) {
		return true
	}

	return false
}

// SetDeviceCount gets a reference to the given int32 and assigns it to the DeviceCount field.
func (o *CollectionStats) SetDeviceCount(v int32) {
	o.DeviceCount = &v
}

// GetOutputCount returns the OutputCount field value if set, zero value otherwise.
func (o *CollectionStats) GetOutputCount() int32 {
	if o == nil || IsNil(o.OutputCount) {
		var ret int32
		return ret
	}
	return *o.OutputCount
}

// GetOutputCountOk returns a tuple with the OutputCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionStats) GetOutputCountOk() (*int32, bool) {
	if o == nil || IsNil(o.OutputCount) {
		return nil, false
	}
	return o.OutputCount, true
}

// HasOutputCount returns a boolean if a field has been set.
func (o *CollectionStats) HasOutputCount() bool {
	if o != nil && !IsNil(o.OutputCount) {
		return true
	}

	return false
}

// SetOutputCount gets a reference to the given int32 and assigns it to the OutputCount field.
func (o *CollectionStats) SetOutputCount(v int32) {
	o.OutputCount = &v
}

// GetFirmwareCount returns the FirmwareCount field value if set, zero value otherwise.
func (o *CollectionStats) GetFirmwareCount() int32 {
	if o == nil || IsNil(o.FirmwareCount) {
		var ret int32
		return ret
	}
	return *o.FirmwareCount
}

// GetFirmwareCountOk returns a tuple with the FirmwareCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionStats) GetFirmwareCountOk() (*int32, bool) {
	if o == nil || IsNil(o.FirmwareCount) {
		return nil, false
	}
	return o.FirmwareCount, true
}

// HasFirmwareCount returns a boolean if a field has been set.
func (o *CollectionStats) HasFirmwareCount() bool {
	if o != nil && !IsNil(o.FirmwareCount) {
		return true
	}

	return false
}

// SetFirmwareCount gets a reference to the given int32 and assigns it to the FirmwareCount field.
func (o *CollectionStats) SetFirmwareCount(v int32) {
	o.FirmwareCount = &v
}

// GetBlobCount returns the BlobCount field value if set, zero value otherwise.
func (o *CollectionStats) GetBlobCount() int32 {
	if o == nil || IsNil(o.BlobCount) {
		var ret int32
		return ret
	}
	return *o.BlobCount
}

// GetBlobCountOk returns a tuple with the BlobCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionStats) GetBlobCountOk() (*int32, bool) {
	if o == nil || IsNil(o.BlobCount) {
		return nil, false
	}
	return o.BlobCount, true
}

// HasBlobCount returns a boolean if a field has been set.
func (o *CollectionStats) HasBlobCount() bool {
	if o != nil && !IsNil(o.BlobCount) {
		return true
	}

	return false
}

// SetBlobCount gets a reference to the given int32 and assigns it to the BlobCount field.
func (o *CollectionStats) SetBlobCount(v int32) {
	o.BlobCount = &v
}

// GetGatewayCount returns the GatewayCount field value if set, zero value otherwise.
func (o *CollectionStats) GetGatewayCount() int32 {
	if o == nil || IsNil(o.GatewayCount) {
		var ret int32
		return ret
	}
	return *o.GatewayCount
}

// GetGatewayCountOk returns a tuple with the GatewayCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionStats) GetGatewayCountOk() (*int32, bool) {
	if o == nil || IsNil(o.GatewayCount) {
		return nil, false
	}
	return o.GatewayCount, true
}

// HasGatewayCount returns a boolean if a field has been set.
func (o *CollectionStats) HasGatewayCount() bool {
	if o != nil && !IsNil(o.GatewayCount) {
		return true
	}

	return false
}

// SetGatewayCount gets a reference to the given int32 and assigns it to the GatewayCount field.
func (o *CollectionStats) SetGatewayCount(v int32) {
	o.GatewayCount = &v
}

// GetDevices returns the Devices field value if set, zero value otherwise.
func (o *CollectionStats) GetDevices() DeviceStats {
	if o == nil || IsNil(o.Devices) {
		var ret DeviceStats
		return ret
	}
	return *o.Devices
}

// GetDevicesOk returns a tuple with the Devices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionStats) GetDevicesOk() (*DeviceStats, bool) {
	if o == nil || IsNil(o.Devices) {
		return nil, false
	}
	return o.Devices, true
}

// HasDevices returns a boolean if a field has been set.
func (o *CollectionStats) HasDevices() bool {
	if o != nil && !IsNil(o.Devices) {
		return true
	}

	return false
}

// SetDevices gets a reference to the given DeviceStats and assigns it to the Devices field.
func (o *CollectionStats) SetDevices(v DeviceStats) {
	o.Devices = &v
}

// GetOutputs returns the Outputs field value if set, zero value otherwise.
func (o *CollectionStats) GetOutputs() OutputStats {
	if o == nil || IsNil(o.Outputs) {
		var ret OutputStats
		return ret
	}
	return *o.Outputs
}

// GetOutputsOk returns a tuple with the Outputs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionStats) GetOutputsOk() (*OutputStats, bool) {
	if o == nil || IsNil(o.Outputs) {
		return nil, false
	}
	return o.Outputs, true
}

// HasOutputs returns a boolean if a field has been set.
func (o *CollectionStats) HasOutputs() bool {
	if o != nil && !IsNil(o.Outputs) {
		return true
	}

	return false
}

// SetOutputs gets a reference to the given OutputStats and assigns it to the Outputs field.
func (o *CollectionStats) SetOutputs(v OutputStats) {
	o.Outputs = &v
}

// GetFirmware returns the Firmware field value if set, zero value otherwise.
func (o *CollectionStats) GetFirmware() FirmwareStats {
	if o == nil || IsNil(o.Firmware) {
		var ret FirmwareStats
		return ret
	}
	return *o.Firmware
}

// GetFirmwareOk returns a tuple with the Firmware field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionStats) GetFirmwareOk() (*FirmwareStats, bool) {
	if o == nil || IsNil(o.Firmware) {
		return nil, false
	}
	return o.Firmware, true
}

// HasFirmware returns a boolean if a field has been set.
func (o *CollectionStats) HasFirmware() bool {
	if o != nil && !IsNil(o.Firmware) {
		return true
	}

	return false
}

// SetFirmware gets a reference to the given FirmwareStats and assigns it to the Firmware field.
func (o *CollectionStats) SetFirmware(v FirmwareStats) {
	o.Firmware = &v
}

// GetBlobs returns the Blobs field value if set, zero value otherwise.
func (o *CollectionStats) GetBlobs() BlobStats {
	if o == nil || IsNil(o.Blobs) {
		var ret BlobStats
		return ret
	}
	return *o.Blobs
}

// GetBlobsOk returns a tuple with the Blobs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionStats) GetBlobsOk() (*BlobStats, bool) {
	if o == nil || IsNil(o.Blobs) {
		return nil, false
	}
	return o.Blobs, true
}

// HasBlobs returns a boolean if a field has been set.
func (o *CollectionStats) HasBlobs() bool {
	if o != nil && !IsNil(o.Blobs) {
		return true
	}

	return false
}

// SetBlobs gets a reference to the given BlobStats and assigns it to the Blobs field.
func (o *CollectionStats) SetBlobs(v BlobStats) {
	o.Blobs = &v
}

// GetGateways returns the Gateways field value if set, zero value otherwise.
func (o *CollectionStats) GetGateways() GatewayStats {
	if o == nil || IsNil(o.Gateways) {
		var ret GatewayStats
		return ret
	}
	return *o.Gateways
}

// GetGatewaysOk returns a tuple with the Gateways field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionStats) GetGatewaysOk() (*GatewayStats, bool) {
	if o == nil || IsNil(o.Gateways) {
		return nil, false
	}
	return o.Gateways, true
}

// HasGateways returns a boolean if a field has been set.
func (o *CollectionStats) HasGateways() bool {
	if o != nil && !IsNil(o.Gateways) {
		return true
	}

	return false
}

// SetGateways gets a reference to the given GatewayStats and assigns it to the Gateways field.
func (o *CollectionStats) SetGateways(v GatewayStats) {
	o.Gateways = &v
}

func (o CollectionStats) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CollectionStats) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DeviceCount) {
		toSerialize["deviceCount"] = o.DeviceCount
	}
	if !IsNil(o.OutputCount) {
		toSerialize["outputCount"] = o.OutputCount
	}
	if !IsNil(o.FirmwareCount) {
		toSerialize["firmwareCount"] = o.FirmwareCount
	}
	if !IsNil(o.BlobCount) {
		toSerialize["blobCount"] = o.BlobCount
	}
	if !IsNil(o.GatewayCount) {
		toSerialize["gatewayCount"] = o.GatewayCount
	}
	if !IsNil(o.Devices) {
		toSerialize["devices"] = o.Devices
	}
	if !IsNil(o.Outputs) {
		toSerialize["outputs"] = o.Outputs
	}
	if !IsNil(o.Firmware) {
		toSerialize["firmware"] = o.Firmware
	}
	if !IsNil(o.Blobs) {
		toSerialize["blobs"] = o.Blobs
	}
	if !IsNil(o.Gateways) {
		toSerialize["gateways"] = o.Gateways
	}
	return toSerialize, nil
}

type NullableCollectionStats struct {
	value *CollectionStats
	isSet bool
}

func (v NullableCollectionStats) Get() *CollectionStats {
	return v.value
}

func (v *NullableCollectionStats) Set(val *CollectionStats) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionStats) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionStats) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionStats(val *CollectionStats) *NullableCollectionStats {
	return &NullableCollectionStats{value: val, isSet: true}
}

func (v NullableCollectionStats) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionStats) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


