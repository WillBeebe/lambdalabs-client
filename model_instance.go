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

// checks if the Instance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Instance{}

// Instance Virtual machine (VM) in Lambda Cloud
type Instance struct {
	// Unique identifier (ID) of an instance
	Id string `json:"id"`
	// User-provided name for the instance
	Name NullableString `json:"name,omitempty"`
	// IPv4 address of the instance
	Ip NullableString `json:"ip,omitempty"`
	// The current status of the instance
	Status string `json:"status"`
	// Names of the SSH keys allowed to access the instance
	SshKeyNames []string `json:"ssh_key_names"`
	// Names of the file systems, if any, attached to the instance
	FileSystemNames []string `json:"file_system_names"`
	Region *Region `json:"region,omitempty"`
	InstanceType *InstanceType `json:"instance_type,omitempty"`
	// Hostname assigned to this instance, which resolves to the instance's IP.
	Hostname NullableString `json:"hostname,omitempty"`
	// Secret token used to log into the jupyter lab server hosted on the instance.
	JupyterToken NullableString `json:"jupyter_token,omitempty"`
	// URL that opens a jupyter lab notebook on the instance.
	JupyterUrl NullableString `json:"jupyter_url,omitempty"`
}

type _Instance Instance

// NewInstance instantiates a new Instance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstance(id string, status string, sshKeyNames []string, fileSystemNames []string) *Instance {
	this := Instance{}
	this.Id = id
	this.Status = status
	this.SshKeyNames = sshKeyNames
	this.FileSystemNames = fileSystemNames
	return &this
}

// NewInstanceWithDefaults instantiates a new Instance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstanceWithDefaults() *Instance {
	this := Instance{}
	return &this
}

// GetId returns the Id field value
func (o *Instance) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Instance) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Instance) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Instance) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Instance) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *Instance) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *Instance) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *Instance) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *Instance) UnsetName() {
	o.Name.Unset()
}

// GetIp returns the Ip field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Instance) GetIp() string {
	if o == nil || IsNil(o.Ip.Get()) {
		var ret string
		return ret
	}
	return *o.Ip.Get()
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Instance) GetIpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ip.Get(), o.Ip.IsSet()
}

// HasIp returns a boolean if a field has been set.
func (o *Instance) HasIp() bool {
	if o != nil && o.Ip.IsSet() {
		return true
	}

	return false
}

// SetIp gets a reference to the given NullableString and assigns it to the Ip field.
func (o *Instance) SetIp(v string) {
	o.Ip.Set(&v)
}
// SetIpNil sets the value for Ip to be an explicit nil
func (o *Instance) SetIpNil() {
	o.Ip.Set(nil)
}

// UnsetIp ensures that no value is present for Ip, not even an explicit nil
func (o *Instance) UnsetIp() {
	o.Ip.Unset()
}

// GetStatus returns the Status field value
func (o *Instance) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Instance) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Instance) SetStatus(v string) {
	o.Status = v
}

// GetSshKeyNames returns the SshKeyNames field value
func (o *Instance) GetSshKeyNames() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.SshKeyNames
}

// GetSshKeyNamesOk returns a tuple with the SshKeyNames field value
// and a boolean to check if the value has been set.
func (o *Instance) GetSshKeyNamesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SshKeyNames, true
}

// SetSshKeyNames sets field value
func (o *Instance) SetSshKeyNames(v []string) {
	o.SshKeyNames = v
}

// GetFileSystemNames returns the FileSystemNames field value
func (o *Instance) GetFileSystemNames() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.FileSystemNames
}

// GetFileSystemNamesOk returns a tuple with the FileSystemNames field value
// and a boolean to check if the value has been set.
func (o *Instance) GetFileSystemNamesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FileSystemNames, true
}

// SetFileSystemNames sets field value
func (o *Instance) SetFileSystemNames(v []string) {
	o.FileSystemNames = v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *Instance) GetRegion() Region {
	if o == nil || IsNil(o.Region) {
		var ret Region
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Instance) GetRegionOk() (*Region, bool) {
	if o == nil || IsNil(o.Region) {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *Instance) HasRegion() bool {
	if o != nil && !IsNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given Region and assigns it to the Region field.
func (o *Instance) SetRegion(v Region) {
	o.Region = &v
}

// GetInstanceType returns the InstanceType field value if set, zero value otherwise.
func (o *Instance) GetInstanceType() InstanceType {
	if o == nil || IsNil(o.InstanceType) {
		var ret InstanceType
		return ret
	}
	return *o.InstanceType
}

// GetInstanceTypeOk returns a tuple with the InstanceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Instance) GetInstanceTypeOk() (*InstanceType, bool) {
	if o == nil || IsNil(o.InstanceType) {
		return nil, false
	}
	return o.InstanceType, true
}

// HasInstanceType returns a boolean if a field has been set.
func (o *Instance) HasInstanceType() bool {
	if o != nil && !IsNil(o.InstanceType) {
		return true
	}

	return false
}

// SetInstanceType gets a reference to the given InstanceType and assigns it to the InstanceType field.
func (o *Instance) SetInstanceType(v InstanceType) {
	o.InstanceType = &v
}

// GetHostname returns the Hostname field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Instance) GetHostname() string {
	if o == nil || IsNil(o.Hostname.Get()) {
		var ret string
		return ret
	}
	return *o.Hostname.Get()
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Instance) GetHostnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Hostname.Get(), o.Hostname.IsSet()
}

// HasHostname returns a boolean if a field has been set.
func (o *Instance) HasHostname() bool {
	if o != nil && o.Hostname.IsSet() {
		return true
	}

	return false
}

// SetHostname gets a reference to the given NullableString and assigns it to the Hostname field.
func (o *Instance) SetHostname(v string) {
	o.Hostname.Set(&v)
}
// SetHostnameNil sets the value for Hostname to be an explicit nil
func (o *Instance) SetHostnameNil() {
	o.Hostname.Set(nil)
}

// UnsetHostname ensures that no value is present for Hostname, not even an explicit nil
func (o *Instance) UnsetHostname() {
	o.Hostname.Unset()
}

// GetJupyterToken returns the JupyterToken field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Instance) GetJupyterToken() string {
	if o == nil || IsNil(o.JupyterToken.Get()) {
		var ret string
		return ret
	}
	return *o.JupyterToken.Get()
}

// GetJupyterTokenOk returns a tuple with the JupyterToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Instance) GetJupyterTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.JupyterToken.Get(), o.JupyterToken.IsSet()
}

// HasJupyterToken returns a boolean if a field has been set.
func (o *Instance) HasJupyterToken() bool {
	if o != nil && o.JupyterToken.IsSet() {
		return true
	}

	return false
}

// SetJupyterToken gets a reference to the given NullableString and assigns it to the JupyterToken field.
func (o *Instance) SetJupyterToken(v string) {
	o.JupyterToken.Set(&v)
}
// SetJupyterTokenNil sets the value for JupyterToken to be an explicit nil
func (o *Instance) SetJupyterTokenNil() {
	o.JupyterToken.Set(nil)
}

// UnsetJupyterToken ensures that no value is present for JupyterToken, not even an explicit nil
func (o *Instance) UnsetJupyterToken() {
	o.JupyterToken.Unset()
}

// GetJupyterUrl returns the JupyterUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Instance) GetJupyterUrl() string {
	if o == nil || IsNil(o.JupyterUrl.Get()) {
		var ret string
		return ret
	}
	return *o.JupyterUrl.Get()
}

// GetJupyterUrlOk returns a tuple with the JupyterUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Instance) GetJupyterUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.JupyterUrl.Get(), o.JupyterUrl.IsSet()
}

// HasJupyterUrl returns a boolean if a field has been set.
func (o *Instance) HasJupyterUrl() bool {
	if o != nil && o.JupyterUrl.IsSet() {
		return true
	}

	return false
}

// SetJupyterUrl gets a reference to the given NullableString and assigns it to the JupyterUrl field.
func (o *Instance) SetJupyterUrl(v string) {
	o.JupyterUrl.Set(&v)
}
// SetJupyterUrlNil sets the value for JupyterUrl to be an explicit nil
func (o *Instance) SetJupyterUrlNil() {
	o.JupyterUrl.Set(nil)
}

// UnsetJupyterUrl ensures that no value is present for JupyterUrl, not even an explicit nil
func (o *Instance) UnsetJupyterUrl() {
	o.JupyterUrl.Unset()
}

func (o Instance) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Instance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Ip.IsSet() {
		toSerialize["ip"] = o.Ip.Get()
	}
	toSerialize["status"] = o.Status
	toSerialize["ssh_key_names"] = o.SshKeyNames
	toSerialize["file_system_names"] = o.FileSystemNames
	if !IsNil(o.Region) {
		toSerialize["region"] = o.Region
	}
	if !IsNil(o.InstanceType) {
		toSerialize["instance_type"] = o.InstanceType
	}
	if o.Hostname.IsSet() {
		toSerialize["hostname"] = o.Hostname.Get()
	}
	if o.JupyterToken.IsSet() {
		toSerialize["jupyter_token"] = o.JupyterToken.Get()
	}
	if o.JupyterUrl.IsSet() {
		toSerialize["jupyter_url"] = o.JupyterUrl.Get()
	}
	return toSerialize, nil
}

func (o *Instance) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"status",
		"ssh_key_names",
		"file_system_names",
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

	varInstance := _Instance{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varInstance)

	if err != nil {
		return err
	}

	*o = Instance(varInstance)

	return err
}

type NullableInstance struct {
	value *Instance
	isSet bool
}

func (v NullableInstance) Get() *Instance {
	return v.value
}

func (v *NullableInstance) Set(val *Instance) {
	v.value = val
	v.isSet = true
}

func (v NullableInstance) IsSet() bool {
	return v.isSet
}

func (v *NullableInstance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstance(val *Instance) *NullableInstance {
	return &NullableInstance{value: val, isSet: true}
}

func (v NullableInstance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


