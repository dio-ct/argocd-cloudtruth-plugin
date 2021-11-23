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

// PaginatedAwsPushTaskList struct for PaginatedAwsPushTaskList
type PaginatedAwsPushTaskList struct {
	Count *int32 `json:"count,omitempty"`
	Next NullableString `json:"next,omitempty"`
	Previous NullableString `json:"previous,omitempty"`
	Results *[]AwsPushTask `json:"results,omitempty"`
}

// NewPaginatedAwsPushTaskList instantiates a new PaginatedAwsPushTaskList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedAwsPushTaskList() *PaginatedAwsPushTaskList {
	this := PaginatedAwsPushTaskList{}
	return &this
}

// NewPaginatedAwsPushTaskListWithDefaults instantiates a new PaginatedAwsPushTaskList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedAwsPushTaskListWithDefaults() *PaginatedAwsPushTaskList {
	this := PaginatedAwsPushTaskList{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *PaginatedAwsPushTaskList) GetCount() int32 {
	if o == nil || o.Count == nil {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedAwsPushTaskList) GetCountOk() (*int32, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *PaginatedAwsPushTaskList) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *PaginatedAwsPushTaskList) SetCount(v int32) {
	o.Count = &v
}

// GetNext returns the Next field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaginatedAwsPushTaskList) GetNext() string {
	if o == nil || o.Next.Get() == nil {
		var ret string
		return ret
	}
	return *o.Next.Get()
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaginatedAwsPushTaskList) GetNextOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Next.Get(), o.Next.IsSet()
}

// HasNext returns a boolean if a field has been set.
func (o *PaginatedAwsPushTaskList) HasNext() bool {
	if o != nil && o.Next.IsSet() {
		return true
	}

	return false
}

// SetNext gets a reference to the given NullableString and assigns it to the Next field.
func (o *PaginatedAwsPushTaskList) SetNext(v string) {
	o.Next.Set(&v)
}
// SetNextNil sets the value for Next to be an explicit nil
func (o *PaginatedAwsPushTaskList) SetNextNil() {
	o.Next.Set(nil)
}

// UnsetNext ensures that no value is present for Next, not even an explicit nil
func (o *PaginatedAwsPushTaskList) UnsetNext() {
	o.Next.Unset()
}

// GetPrevious returns the Previous field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaginatedAwsPushTaskList) GetPrevious() string {
	if o == nil || o.Previous.Get() == nil {
		var ret string
		return ret
	}
	return *o.Previous.Get()
}

// GetPreviousOk returns a tuple with the Previous field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaginatedAwsPushTaskList) GetPreviousOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Previous.Get(), o.Previous.IsSet()
}

// HasPrevious returns a boolean if a field has been set.
func (o *PaginatedAwsPushTaskList) HasPrevious() bool {
	if o != nil && o.Previous.IsSet() {
		return true
	}

	return false
}

// SetPrevious gets a reference to the given NullableString and assigns it to the Previous field.
func (o *PaginatedAwsPushTaskList) SetPrevious(v string) {
	o.Previous.Set(&v)
}
// SetPreviousNil sets the value for Previous to be an explicit nil
func (o *PaginatedAwsPushTaskList) SetPreviousNil() {
	o.Previous.Set(nil)
}

// UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
func (o *PaginatedAwsPushTaskList) UnsetPrevious() {
	o.Previous.Unset()
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *PaginatedAwsPushTaskList) GetResults() []AwsPushTask {
	if o == nil || o.Results == nil {
		var ret []AwsPushTask
		return ret
	}
	return *o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedAwsPushTaskList) GetResultsOk() (*[]AwsPushTask, bool) {
	if o == nil || o.Results == nil {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *PaginatedAwsPushTaskList) HasResults() bool {
	if o != nil && o.Results != nil {
		return true
	}

	return false
}

// SetResults gets a reference to the given []AwsPushTask and assigns it to the Results field.
func (o *PaginatedAwsPushTaskList) SetResults(v []AwsPushTask) {
	o.Results = &v
}

func (o PaginatedAwsPushTaskList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Count != nil {
		toSerialize["count"] = o.Count
	}
	if o.Next.IsSet() {
		toSerialize["next"] = o.Next.Get()
	}
	if o.Previous.IsSet() {
		toSerialize["previous"] = o.Previous.Get()
	}
	if o.Results != nil {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedAwsPushTaskList struct {
	value *PaginatedAwsPushTaskList
	isSet bool
}

func (v NullablePaginatedAwsPushTaskList) Get() *PaginatedAwsPushTaskList {
	return v.value
}

func (v *NullablePaginatedAwsPushTaskList) Set(val *PaginatedAwsPushTaskList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedAwsPushTaskList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedAwsPushTaskList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedAwsPushTaskList(val *PaginatedAwsPushTaskList) *NullablePaginatedAwsPushTaskList {
	return &NullablePaginatedAwsPushTaskList{value: val, isSet: true}
}

func (v NullablePaginatedAwsPushTaskList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedAwsPushTaskList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

