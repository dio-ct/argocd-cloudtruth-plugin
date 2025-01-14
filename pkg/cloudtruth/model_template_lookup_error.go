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
)

// TemplateLookupError Indicates errors occurred while retrieving values to substitute into the template.
type TemplateLookupError struct {
	Detail []TemplateLookupErrorEntry `json:"detail"`
}

// NewTemplateLookupError instantiates a new TemplateLookupError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateLookupError(detail []TemplateLookupErrorEntry) *TemplateLookupError {
	this := TemplateLookupError{}
	this.Detail = detail
	return &this
}

// NewTemplateLookupErrorWithDefaults instantiates a new TemplateLookupError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateLookupErrorWithDefaults() *TemplateLookupError {
	this := TemplateLookupError{}
	return &this
}

// GetDetail returns the Detail field value
func (o *TemplateLookupError) GetDetail() []TemplateLookupErrorEntry {
	if o == nil {
		var ret []TemplateLookupErrorEntry
		return ret
	}

	return o.Detail
}

// GetDetailOk returns a tuple with the Detail field value
// and a boolean to check if the value has been set.
func (o *TemplateLookupError) GetDetailOk() (*[]TemplateLookupErrorEntry, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Detail, true
}

// SetDetail sets field value
func (o *TemplateLookupError) SetDetail(v []TemplateLookupErrorEntry) {
	o.Detail = v
}

func (o TemplateLookupError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["detail"] = o.Detail
	}
	return json.Marshal(toSerialize)
}

type NullableTemplateLookupError struct {
	value *TemplateLookupError
	isSet bool
}

func (v NullableTemplateLookupError) Get() *TemplateLookupError {
	return v.value
}

func (v *NullableTemplateLookupError) Set(val *TemplateLookupError) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateLookupError) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateLookupError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateLookupError(val *TemplateLookupError) *NullableTemplateLookupError {
	return &NullableTemplateLookupError{value: val, isSet: true}
}

func (v NullableTemplateLookupError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateLookupError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


