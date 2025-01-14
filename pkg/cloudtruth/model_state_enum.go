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

// StateEnum the model 'StateEnum'
type StateEnum string

// List of StateEnum
const (
	STATEENUM_QUEUED StateEnum = "queued"
	STATEENUM_RUNNING StateEnum = "running"
	STATEENUM_SKIPPED StateEnum = "skipped"
	STATEENUM_SUCCESS StateEnum = "success"
	STATEENUM_FAILURE StateEnum = "failure"
)

// All allowed values of StateEnum enum
var AllowedStateEnumEnumValues = []StateEnum{
	"queued",
	"running",
	"skipped",
	"success",
	"failure",
}

func (v *StateEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := StateEnum(value)
	for _, existing := range AllowedStateEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid StateEnum", value)
}

// NewStateEnumFromValue returns a pointer to a valid StateEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewStateEnumFromValue(v string) (*StateEnum, error) {
	ev := StateEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for StateEnum: valid values are %v", v, AllowedStateEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v StateEnum) IsValid() bool {
	for _, existing := range AllowedStateEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to StateEnum value
func (v StateEnum) Ptr() *StateEnum {
	return &v
}

type NullableStateEnum struct {
	value *StateEnum
	isSet bool
}

func (v NullableStateEnum) Get() *StateEnum {
	return v.value
}

func (v *NullableStateEnum) Set(val *StateEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableStateEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableStateEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStateEnum(val *StateEnum) *NullableStateEnum {
	return &NullableStateEnum{value: val, isSet: true}
}

func (v NullableStateEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStateEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

