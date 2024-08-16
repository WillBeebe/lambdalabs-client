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

// checks if the InstanceTypes200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InstanceTypes200Response{}

// InstanceTypes200Response struct for InstanceTypes200Response
type InstanceTypes200Response struct {
	// Dict of instance_type_name to instance_type and region availability.
	Data map[string]InstanceTypes200ResponseDataValue `json:"data"`
}

type _InstanceTypes200Response InstanceTypes200Response

// NewInstanceTypes200Response instantiates a new InstanceTypes200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstanceTypes200Response(data map[string]InstanceTypes200ResponseDataValue) *InstanceTypes200Response {
	this := InstanceTypes200Response{}
	this.Data = data
	return &this
}

// NewInstanceTypes200ResponseWithDefaults instantiates a new InstanceTypes200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstanceTypes200ResponseWithDefaults() *InstanceTypes200Response {
	this := InstanceTypes200Response{}
	return &this
}

// GetData returns the Data field value
func (o *InstanceTypes200Response) GetData() map[string]InstanceTypes200ResponseDataValue {
	if o == nil {
		var ret map[string]InstanceTypes200ResponseDataValue
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *InstanceTypes200Response) GetDataOk() (*map[string]InstanceTypes200ResponseDataValue, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *InstanceTypes200Response) SetData(v map[string]InstanceTypes200ResponseDataValue) {
	o.Data = v
}

func (o InstanceTypes200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InstanceTypes200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *InstanceTypes200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
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

	varInstanceTypes200Response := _InstanceTypes200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varInstanceTypes200Response)

	if err != nil {
		return err
	}

	*o = InstanceTypes200Response(varInstanceTypes200Response)

	return err
}

type NullableInstanceTypes200Response struct {
	value *InstanceTypes200Response
	isSet bool
}

func (v NullableInstanceTypes200Response) Get() *InstanceTypes200Response {
	return v.value
}

func (v *NullableInstanceTypes200Response) Set(val *InstanceTypes200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableInstanceTypes200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableInstanceTypes200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstanceTypes200Response(val *InstanceTypes200Response) *NullableInstanceTypes200Response {
	return &NullableInstanceTypes200Response{value: val, isSet: true}
}

func (v NullableInstanceTypes200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstanceTypes200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


