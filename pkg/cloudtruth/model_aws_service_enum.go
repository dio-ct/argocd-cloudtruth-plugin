/*
CloudTruth Management API

CloudTruth centralizes your configuration parameters and secrets making them easier to manage and use as a team.

API version: v1
Contact: support@cloudtruth.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudtruth

import (
	"encoding/json"
	"fmt"
)

// AwsServiceEnum the model 'AwsServiceEnum'
type AwsServiceEnum string

// List of AwsServiceEnum
const (
	AWSSERVICEENUM_S3 AwsServiceEnum = "s3"
	AWSSERVICEENUM_SECRETSMANAGER AwsServiceEnum = "secretsmanager"
	AWSSERVICEENUM_SSM AwsServiceEnum = "ssm"
)

// All allowed values of AwsServiceEnum enum
var AllowedAwsServiceEnumEnumValues = []AwsServiceEnum{
	"s3",
	"secretsmanager",
	"ssm",
}

func (v *AwsServiceEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AwsServiceEnum(value)
	for _, existing := range AllowedAwsServiceEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AwsServiceEnum", value)
}

// NewAwsServiceEnumFromValue returns a pointer to a valid AwsServiceEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAwsServiceEnumFromValue(v string) (*AwsServiceEnum, error) {
	ev := AwsServiceEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AwsServiceEnum: valid values are %v", v, AllowedAwsServiceEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AwsServiceEnum) IsValid() bool {
	for _, existing := range AllowedAwsServiceEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AwsServiceEnum value
func (v AwsServiceEnum) Ptr() *AwsServiceEnum {
	return &v
}

type NullableAwsServiceEnum struct {
	value *AwsServiceEnum
	isSet bool
}

func (v NullableAwsServiceEnum) Get() *AwsServiceEnum {
	return v.value
}

func (v *NullableAwsServiceEnum) Set(val *AwsServiceEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableAwsServiceEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableAwsServiceEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAwsServiceEnum(val *AwsServiceEnum) *NullableAwsServiceEnum {
	return &NullableAwsServiceEnum{value: val, isSet: true}
}

func (v NullableAwsServiceEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAwsServiceEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

