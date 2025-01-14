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
	"time"
)

// ParameterTimelineEntry Details about a single change.
type ParameterTimelineEntry struct {
	HistoryDate time.Time `json:"history_date"`
	HistoryType HistoryTypeEnum `json:"history_type"`
	// The unique identifier of a user.
	HistoryUser NullableString `json:"history_user,omitempty"`
	// The affected environment(s).
	HistoryEnvironments []ParameterTimelineEntryEnvironment `json:"history_environments"`
	// The component of the parameter that changed.
	HistoryModel HistoryModelEnum `json:"history_model"`
	// The affected parameter.
	HistoryParameter ParameterTimelineEntryParameter `json:"history_parameter"`
}

// NewParameterTimelineEntry instantiates a new ParameterTimelineEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewParameterTimelineEntry(historyDate time.Time, historyType HistoryTypeEnum, historyEnvironments []ParameterTimelineEntryEnvironment, historyModel HistoryModelEnum, historyParameter ParameterTimelineEntryParameter) *ParameterTimelineEntry {
	this := ParameterTimelineEntry{}
	this.HistoryDate = historyDate
	this.HistoryType = historyType
	this.HistoryEnvironments = historyEnvironments
	this.HistoryModel = historyModel
	this.HistoryParameter = historyParameter
	return &this
}

// NewParameterTimelineEntryWithDefaults instantiates a new ParameterTimelineEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewParameterTimelineEntryWithDefaults() *ParameterTimelineEntry {
	this := ParameterTimelineEntry{}
	return &this
}

// GetHistoryDate returns the HistoryDate field value
func (o *ParameterTimelineEntry) GetHistoryDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.HistoryDate
}

// GetHistoryDateOk returns a tuple with the HistoryDate field value
// and a boolean to check if the value has been set.
func (o *ParameterTimelineEntry) GetHistoryDateOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.HistoryDate, true
}

// SetHistoryDate sets field value
func (o *ParameterTimelineEntry) SetHistoryDate(v time.Time) {
	o.HistoryDate = v
}

// GetHistoryType returns the HistoryType field value
func (o *ParameterTimelineEntry) GetHistoryType() HistoryTypeEnum {
	if o == nil {
		var ret HistoryTypeEnum
		return ret
	}

	return o.HistoryType
}

// GetHistoryTypeOk returns a tuple with the HistoryType field value
// and a boolean to check if the value has been set.
func (o *ParameterTimelineEntry) GetHistoryTypeOk() (*HistoryTypeEnum, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.HistoryType, true
}

// SetHistoryType sets field value
func (o *ParameterTimelineEntry) SetHistoryType(v HistoryTypeEnum) {
	o.HistoryType = v
}

// GetHistoryUser returns the HistoryUser field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ParameterTimelineEntry) GetHistoryUser() string {
	if o == nil || o.HistoryUser.Get() == nil {
		var ret string
		return ret
	}
	return *o.HistoryUser.Get()
}

// GetHistoryUserOk returns a tuple with the HistoryUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ParameterTimelineEntry) GetHistoryUserOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.HistoryUser.Get(), o.HistoryUser.IsSet()
}

// HasHistoryUser returns a boolean if a field has been set.
func (o *ParameterTimelineEntry) HasHistoryUser() bool {
	if o != nil && o.HistoryUser.IsSet() {
		return true
	}

	return false
}

// SetHistoryUser gets a reference to the given NullableString and assigns it to the HistoryUser field.
func (o *ParameterTimelineEntry) SetHistoryUser(v string) {
	o.HistoryUser.Set(&v)
}
// SetHistoryUserNil sets the value for HistoryUser to be an explicit nil
func (o *ParameterTimelineEntry) SetHistoryUserNil() {
	o.HistoryUser.Set(nil)
}

// UnsetHistoryUser ensures that no value is present for HistoryUser, not even an explicit nil
func (o *ParameterTimelineEntry) UnsetHistoryUser() {
	o.HistoryUser.Unset()
}

// GetHistoryEnvironments returns the HistoryEnvironments field value
func (o *ParameterTimelineEntry) GetHistoryEnvironments() []ParameterTimelineEntryEnvironment {
	if o == nil {
		var ret []ParameterTimelineEntryEnvironment
		return ret
	}

	return o.HistoryEnvironments
}

// GetHistoryEnvironmentsOk returns a tuple with the HistoryEnvironments field value
// and a boolean to check if the value has been set.
func (o *ParameterTimelineEntry) GetHistoryEnvironmentsOk() (*[]ParameterTimelineEntryEnvironment, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.HistoryEnvironments, true
}

// SetHistoryEnvironments sets field value
func (o *ParameterTimelineEntry) SetHistoryEnvironments(v []ParameterTimelineEntryEnvironment) {
	o.HistoryEnvironments = v
}

// GetHistoryModel returns the HistoryModel field value
func (o *ParameterTimelineEntry) GetHistoryModel() HistoryModelEnum {
	if o == nil {
		var ret HistoryModelEnum
		return ret
	}

	return o.HistoryModel
}

// GetHistoryModelOk returns a tuple with the HistoryModel field value
// and a boolean to check if the value has been set.
func (o *ParameterTimelineEntry) GetHistoryModelOk() (*HistoryModelEnum, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.HistoryModel, true
}

// SetHistoryModel sets field value
func (o *ParameterTimelineEntry) SetHistoryModel(v HistoryModelEnum) {
	o.HistoryModel = v
}

// GetHistoryParameter returns the HistoryParameter field value
func (o *ParameterTimelineEntry) GetHistoryParameter() ParameterTimelineEntryParameter {
	if o == nil {
		var ret ParameterTimelineEntryParameter
		return ret
	}

	return o.HistoryParameter
}

// GetHistoryParameterOk returns a tuple with the HistoryParameter field value
// and a boolean to check if the value has been set.
func (o *ParameterTimelineEntry) GetHistoryParameterOk() (*ParameterTimelineEntryParameter, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.HistoryParameter, true
}

// SetHistoryParameter sets field value
func (o *ParameterTimelineEntry) SetHistoryParameter(v ParameterTimelineEntryParameter) {
	o.HistoryParameter = v
}

func (o ParameterTimelineEntry) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["history_date"] = o.HistoryDate
	}
	if true {
		toSerialize["history_type"] = o.HistoryType
	}
	if o.HistoryUser.IsSet() {
		toSerialize["history_user"] = o.HistoryUser.Get()
	}
	if true {
		toSerialize["history_environments"] = o.HistoryEnvironments
	}
	if true {
		toSerialize["history_model"] = o.HistoryModel
	}
	if true {
		toSerialize["history_parameter"] = o.HistoryParameter
	}
	return json.Marshal(toSerialize)
}

type NullableParameterTimelineEntry struct {
	value *ParameterTimelineEntry
	isSet bool
}

func (v NullableParameterTimelineEntry) Get() *ParameterTimelineEntry {
	return v.value
}

func (v *NullableParameterTimelineEntry) Set(val *ParameterTimelineEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableParameterTimelineEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableParameterTimelineEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableParameterTimelineEntry(val *ParameterTimelineEntry) *NullableParameterTimelineEntry {
	return &NullableParameterTimelineEntry{value: val, isSet: true}
}

func (v NullableParameterTimelineEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableParameterTimelineEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


