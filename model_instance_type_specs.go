/*
Lambda Cloud API

API for interacting with the Lambda GPU Cloud

API version: 1.5.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the InstanceTypeSpecs type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InstanceTypeSpecs{}

// InstanceTypeSpecs struct for InstanceTypeSpecs
type InstanceTypeSpecs struct {
	// Number of virtual CPUs
	Vcpus int32 `json:"vcpus"`
	// Amount of RAM, in gibibytes (GiB)
	MemoryGib int32 `json:"memory_gib"`
	// Amount of storage, in gibibytes (GiB).
	StorageGib int32 `json:"storage_gib"`
	// Number of GPUs
	Gpus int32 `json:"gpus"`
}

type _InstanceTypeSpecs InstanceTypeSpecs

// NewInstanceTypeSpecs instantiates a new InstanceTypeSpecs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstanceTypeSpecs(vcpus int32, memoryGib int32, storageGib int32, gpus int32) *InstanceTypeSpecs {
	this := InstanceTypeSpecs{}
	this.Vcpus = vcpus
	this.MemoryGib = memoryGib
	this.StorageGib = storageGib
	this.Gpus = gpus
	return &this
}

// NewInstanceTypeSpecsWithDefaults instantiates a new InstanceTypeSpecs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstanceTypeSpecsWithDefaults() *InstanceTypeSpecs {
	this := InstanceTypeSpecs{}
	return &this
}

// GetVcpus returns the Vcpus field value
func (o *InstanceTypeSpecs) GetVcpus() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Vcpus
}

// GetVcpusOk returns a tuple with the Vcpus field value
// and a boolean to check if the value has been set.
func (o *InstanceTypeSpecs) GetVcpusOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Vcpus, true
}

// SetVcpus sets field value
func (o *InstanceTypeSpecs) SetVcpus(v int32) {
	o.Vcpus = v
}

// GetMemoryGib returns the MemoryGib field value
func (o *InstanceTypeSpecs) GetMemoryGib() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MemoryGib
}

// GetMemoryGibOk returns a tuple with the MemoryGib field value
// and a boolean to check if the value has been set.
func (o *InstanceTypeSpecs) GetMemoryGibOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MemoryGib, true
}

// SetMemoryGib sets field value
func (o *InstanceTypeSpecs) SetMemoryGib(v int32) {
	o.MemoryGib = v
}

// GetStorageGib returns the StorageGib field value
func (o *InstanceTypeSpecs) GetStorageGib() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.StorageGib
}

// GetStorageGibOk returns a tuple with the StorageGib field value
// and a boolean to check if the value has been set.
func (o *InstanceTypeSpecs) GetStorageGibOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StorageGib, true
}

// SetStorageGib sets field value
func (o *InstanceTypeSpecs) SetStorageGib(v int32) {
	o.StorageGib = v
}

// GetGpus returns the Gpus field value
func (o *InstanceTypeSpecs) GetGpus() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Gpus
}

// GetGpusOk returns a tuple with the Gpus field value
// and a boolean to check if the value has been set.
func (o *InstanceTypeSpecs) GetGpusOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Gpus, true
}

// SetGpus sets field value
func (o *InstanceTypeSpecs) SetGpus(v int32) {
	o.Gpus = v
}

func (o InstanceTypeSpecs) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InstanceTypeSpecs) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["vcpus"] = o.Vcpus
	toSerialize["memory_gib"] = o.MemoryGib
	toSerialize["storage_gib"] = o.StorageGib
	toSerialize["gpus"] = o.Gpus
	return toSerialize, nil
}

func (o *InstanceTypeSpecs) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"vcpus",
		"memory_gib",
		"storage_gib",
		"gpus",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varInstanceTypeSpecs := _InstanceTypeSpecs{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varInstanceTypeSpecs)

	if err != nil {
		return err
	}

	*o = InstanceTypeSpecs(varInstanceTypeSpecs)

	return err
}

type NullableInstanceTypeSpecs struct {
	value *InstanceTypeSpecs
	isSet bool
}

func (v NullableInstanceTypeSpecs) Get() *InstanceTypeSpecs {
	return v.value
}

func (v *NullableInstanceTypeSpecs) Set(val *InstanceTypeSpecs) {
	v.value = val
	v.isSet = true
}

func (v NullableInstanceTypeSpecs) IsSet() bool {
	return v.isSet
}

func (v *NullableInstanceTypeSpecs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstanceTypeSpecs(val *InstanceTypeSpecs) *NullableInstanceTypeSpecs {
	return &NullableInstanceTypeSpecs{value: val, isSet: true}
}

func (v NullableInstanceTypeSpecs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstanceTypeSpecs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


