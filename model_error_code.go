/*
Lambda Cloud API

API for interacting with the Lambda GPU Cloud

API version: 1.5.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// ErrorCode Unique identifier for the type of error
type ErrorCode string

// List of errorCode
const (
	GLOBAL_UNKNOWN ErrorCode = "global/unknown"
	GLOBAL_INVALID_API_KEY ErrorCode = "global/invalid-api-key"
	GLOBAL_ACCOUNT_INACTIVE ErrorCode = "global/account-inactive"
	GLOBAL_INVALID_ADDRESS ErrorCode = "global/invalid-address"
	GLOBAL_INVALID_PARAMETERS ErrorCode = "global/invalid-parameters"
	GLOBAL_OBJECT_DOES_NOT_EXIST ErrorCode = "global/object-does-not-exist"
	GLOBAL_QUOTA_EXCEEDED ErrorCode = "global/quota-exceeded"
	INSTANCE_OPERATIONS_LAUNCH_INSUFFICIENT_CAPACITY ErrorCode = "instance-operations/launch/insufficient-capacity"
	INSTANCE_OPERATIONS_LAUNCH_FILE_SYSTEM_IN_WRONG_REGION ErrorCode = "instance-operations/launch/file-system-in-wrong-region"
	SSH_KEYS_KEY_IN_USE ErrorCode = "ssh-keys/key-in-use"
)

// All allowed values of ErrorCode enum
var AllowedErrorCodeEnumValues = []ErrorCode{
	"global/unknown",
	"global/invalid-api-key",
	"global/account-inactive",
	"global/invalid-address",
	"global/invalid-parameters",
	"global/object-does-not-exist",
	"global/quota-exceeded",
	"instance-operations/launch/insufficient-capacity",
	"instance-operations/launch/file-system-in-wrong-region",
	"ssh-keys/key-in-use",
}

func (v *ErrorCode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ErrorCode(value)
	for _, existing := range AllowedErrorCodeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ErrorCode", value)
}

// NewErrorCodeFromValue returns a pointer to a valid ErrorCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewErrorCodeFromValue(v string) (*ErrorCode, error) {
	ev := ErrorCode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ErrorCode: valid values are %v", v, AllowedErrorCodeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ErrorCode) IsValid() bool {
	for _, existing := range AllowedErrorCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to errorCode value
func (v ErrorCode) Ptr() *ErrorCode {
	return &v
}

type NullableErrorCode struct {
	value *ErrorCode
	isSet bool
}

func (v NullableErrorCode) Get() *ErrorCode {
	return v.value
}

func (v *NullableErrorCode) Set(val *ErrorCode) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorCode) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorCode(val *ErrorCode) *NullableErrorCode {
	return &NullableErrorCode{value: val, isSet: true}
}

func (v NullableErrorCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

