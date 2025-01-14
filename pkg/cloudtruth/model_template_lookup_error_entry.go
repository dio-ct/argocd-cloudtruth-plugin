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

// TemplateLookupErrorEntry struct for TemplateLookupErrorEntry
type TemplateLookupErrorEntry struct {
	// The parameter id.
	ParameterId string `json:"parameter_id"`
	// The parameter name.
	ParameterName string `json:"parameter_name"`
	// The error code.
	ErrorCode string `json:"error_code"`
	// Details about the error.
	ErrorDetail string `json:"error_detail"`
}

// NewTemplateLookupErrorEntry instantiates a new TemplateLookupErrorEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateLookupErrorEntry(parameterId string, parameterName string, errorCode string, errorDetail string) *TemplateLookupErrorEntry {
	this := TemplateLookupErrorEntry{}
	this.ParameterId = parameterId
	this.ParameterName = parameterName
	this.ErrorCode = errorCode
	this.ErrorDetail = errorDetail
	return &this
}

// NewTemplateLookupErrorEntryWithDefaults instantiates a new TemplateLookupErrorEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateLookupErrorEntryWithDefaults() *TemplateLookupErrorEntry {
	this := TemplateLookupErrorEntry{}
	return &this
}

// GetParameterId returns the ParameterId field value
func (o *TemplateLookupErrorEntry) GetParameterId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ParameterId
}

// GetParameterIdOk returns a tuple with the ParameterId field value
// and a boolean to check if the value has been set.
func (o *TemplateLookupErrorEntry) GetParameterIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ParameterId, true
}

// SetParameterId sets field value
func (o *TemplateLookupErrorEntry) SetParameterId(v string) {
	o.ParameterId = v
}

// GetParameterName returns the ParameterName field value
func (o *TemplateLookupErrorEntry) GetParameterName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ParameterName
}

// GetParameterNameOk returns a tuple with the ParameterName field value
// and a boolean to check if the value has been set.
func (o *TemplateLookupErrorEntry) GetParameterNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ParameterName, true
}

// SetParameterName sets field value
func (o *TemplateLookupErrorEntry) SetParameterName(v string) {
	o.ParameterName = v
}

// GetErrorCode returns the ErrorCode field value
func (o *TemplateLookupErrorEntry) GetErrorCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value
// and a boolean to check if the value has been set.
func (o *TemplateLookupErrorEntry) GetErrorCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ErrorCode, true
}

// SetErrorCode sets field value
func (o *TemplateLookupErrorEntry) SetErrorCode(v string) {
	o.ErrorCode = v
}

// GetErrorDetail returns the ErrorDetail field value
func (o *TemplateLookupErrorEntry) GetErrorDetail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ErrorDetail
}

// GetErrorDetailOk returns a tuple with the ErrorDetail field value
// and a boolean to check if the value has been set.
func (o *TemplateLookupErrorEntry) GetErrorDetailOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ErrorDetail, true
}

// SetErrorDetail sets field value
func (o *TemplateLookupErrorEntry) SetErrorDetail(v string) {
	o.ErrorDetail = v
}

func (o TemplateLookupErrorEntry) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["parameter_id"] = o.ParameterId
	}
	if true {
		toSerialize["parameter_name"] = o.ParameterName
	}
	if true {
		toSerialize["error_code"] = o.ErrorCode
	}
	if true {
		toSerialize["error_detail"] = o.ErrorDetail
	}
	return json.Marshal(toSerialize)
}

type NullableTemplateLookupErrorEntry struct {
	value *TemplateLookupErrorEntry
	isSet bool
}

func (v NullableTemplateLookupErrorEntry) Get() *TemplateLookupErrorEntry {
	return v.value
}

func (v *NullableTemplateLookupErrorEntry) Set(val *TemplateLookupErrorEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateLookupErrorEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateLookupErrorEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateLookupErrorEntry(val *TemplateLookupErrorEntry) *NullableTemplateLookupErrorEntry {
	return &NullableTemplateLookupErrorEntry{value: val, isSet: true}
}

func (v NullableTemplateLookupErrorEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateLookupErrorEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


