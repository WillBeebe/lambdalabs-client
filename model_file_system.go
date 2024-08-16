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

// checks if the FileSystem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FileSystem{}

// FileSystem Information about a shared file system
type FileSystem struct {
	// Unique identifier (ID) of a file system
	Id string `json:"id"`
	// Name of a file system
	Name string `json:"name"`
	// A date and time, formatted as an ISO 8601 time stamp
	Created string `json:"created"`
	CreatedBy User `json:"created_by"`
	// Absolute path indicating where on instances the file system will be mounted
	MountPoint string `json:"mount_point"`
	Region Region `json:"region"`
	// Whether the file system is currently in use by an instance. File systems that are in use cannot be deleted.
	IsInUse bool `json:"is_in_use"`
	// Approximate amount of storage used by the file system, in bytes. This value is an estimate that is updated every several hours.
	BytesUsed *int32 `json:"bytes_used,omitempty"`
}

type _FileSystem FileSystem

// NewFileSystem instantiates a new FileSystem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileSystem(id string, name string, created string, createdBy User, mountPoint string, region Region, isInUse bool) *FileSystem {
	this := FileSystem{}
	this.Id = id
	this.Name = name
	this.Created = created
	this.CreatedBy = createdBy
	this.MountPoint = mountPoint
	this.Region = region
	this.IsInUse = isInUse
	return &this
}

// NewFileSystemWithDefaults instantiates a new FileSystem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileSystemWithDefaults() *FileSystem {
	this := FileSystem{}
	return &this
}

// GetId returns the Id field value
func (o *FileSystem) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *FileSystem) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *FileSystem) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *FileSystem) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FileSystem) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *FileSystem) SetName(v string) {
	o.Name = v
}

// GetCreated returns the Created field value
func (o *FileSystem) GetCreated() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *FileSystem) GetCreatedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *FileSystem) SetCreated(v string) {
	o.Created = v
}

// GetCreatedBy returns the CreatedBy field value
func (o *FileSystem) GetCreatedBy() User {
	if o == nil {
		var ret User
		return ret
	}

	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *FileSystem) GetCreatedByOk() (*User, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value
func (o *FileSystem) SetCreatedBy(v User) {
	o.CreatedBy = v
}

// GetMountPoint returns the MountPoint field value
func (o *FileSystem) GetMountPoint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MountPoint
}

// GetMountPointOk returns a tuple with the MountPoint field value
// and a boolean to check if the value has been set.
func (o *FileSystem) GetMountPointOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MountPoint, true
}

// SetMountPoint sets field value
func (o *FileSystem) SetMountPoint(v string) {
	o.MountPoint = v
}

// GetRegion returns the Region field value
func (o *FileSystem) GetRegion() Region {
	if o == nil {
		var ret Region
		return ret
	}

	return o.Region
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
func (o *FileSystem) GetRegionOk() (*Region, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Region, true
}

// SetRegion sets field value
func (o *FileSystem) SetRegion(v Region) {
	o.Region = v
}

// GetIsInUse returns the IsInUse field value
func (o *FileSystem) GetIsInUse() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsInUse
}

// GetIsInUseOk returns a tuple with the IsInUse field value
// and a boolean to check if the value has been set.
func (o *FileSystem) GetIsInUseOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsInUse, true
}

// SetIsInUse sets field value
func (o *FileSystem) SetIsInUse(v bool) {
	o.IsInUse = v
}

// GetBytesUsed returns the BytesUsed field value if set, zero value otherwise.
func (o *FileSystem) GetBytesUsed() int32 {
	if o == nil || IsNil(o.BytesUsed) {
		var ret int32
		return ret
	}
	return *o.BytesUsed
}

// GetBytesUsedOk returns a tuple with the BytesUsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileSystem) GetBytesUsedOk() (*int32, bool) {
	if o == nil || IsNil(o.BytesUsed) {
		return nil, false
	}
	return o.BytesUsed, true
}

// HasBytesUsed returns a boolean if a field has been set.
func (o *FileSystem) HasBytesUsed() bool {
	if o != nil && !IsNil(o.BytesUsed) {
		return true
	}

	return false
}

// SetBytesUsed gets a reference to the given int32 and assigns it to the BytesUsed field.
func (o *FileSystem) SetBytesUsed(v int32) {
	o.BytesUsed = &v
}

func (o FileSystem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FileSystem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["created"] = o.Created
	toSerialize["created_by"] = o.CreatedBy
	toSerialize["mount_point"] = o.MountPoint
	toSerialize["region"] = o.Region
	toSerialize["is_in_use"] = o.IsInUse
	if !IsNil(o.BytesUsed) {
		toSerialize["bytes_used"] = o.BytesUsed
	}
	return toSerialize, nil
}

func (o *FileSystem) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"created",
		"created_by",
		"mount_point",
		"region",
		"is_in_use",
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

	varFileSystem := _FileSystem{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFileSystem)

	if err != nil {
		return err
	}

	*o = FileSystem(varFileSystem)

	return err
}

type NullableFileSystem struct {
	value *FileSystem
	isSet bool
}

func (v NullableFileSystem) Get() *FileSystem {
	return v.value
}

func (v *NullableFileSystem) Set(val *FileSystem) {
	v.value = val
	v.isSet = true
}

func (v NullableFileSystem) IsSet() bool {
	return v.isSet
}

func (v *NullableFileSystem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileSystem(val *FileSystem) *NullableFileSystem {
	return &NullableFileSystem{value: val, isSet: true}
}

func (v NullableFileSystem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileSystem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


