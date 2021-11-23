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

// TemplateTimeline struct for TemplateTimeline
type TemplateTimeline struct {
	// The number of records in this response.
	Count int32 `json:"count"`
	// If present, additional history can be retrieved using this timestamp in the next call for the as_of query parameter value.
	NextAsOf *time.Time `json:"next_as_of,omitempty"`
	Results []TemplateTimelineEntry `json:"results"`
}

// NewTemplateTimeline instantiates a new TemplateTimeline object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateTimeline(count int32, results []TemplateTimelineEntry) *TemplateTimeline {
	this := TemplateTimeline{}
	this.Count = count
	this.Results = results
	return &this
}

// NewTemplateTimelineWithDefaults instantiates a new TemplateTimeline object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateTimelineWithDefaults() *TemplateTimeline {
	this := TemplateTimeline{}
	return &this
}

// GetCount returns the Count field value
func (o *TemplateTimeline) GetCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *TemplateTimeline) GetCountOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *TemplateTimeline) SetCount(v int32) {
	o.Count = v
}

// GetNextAsOf returns the NextAsOf field value if set, zero value otherwise.
func (o *TemplateTimeline) GetNextAsOf() time.Time {
	if o == nil || o.NextAsOf == nil {
		var ret time.Time
		return ret
	}
	return *o.NextAsOf
}

// GetNextAsOfOk returns a tuple with the NextAsOf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateTimeline) GetNextAsOfOk() (*time.Time, bool) {
	if o == nil || o.NextAsOf == nil {
		return nil, false
	}
	return o.NextAsOf, true
}

// HasNextAsOf returns a boolean if a field has been set.
func (o *TemplateTimeline) HasNextAsOf() bool {
	if o != nil && o.NextAsOf != nil {
		return true
	}

	return false
}

// SetNextAsOf gets a reference to the given time.Time and assigns it to the NextAsOf field.
func (o *TemplateTimeline) SetNextAsOf(v time.Time) {
	o.NextAsOf = &v
}

// GetResults returns the Results field value
func (o *TemplateTimeline) GetResults() []TemplateTimelineEntry {
	if o == nil {
		var ret []TemplateTimelineEntry
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *TemplateTimeline) GetResultsOk() (*[]TemplateTimelineEntry, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Results, true
}

// SetResults sets field value
func (o *TemplateTimeline) SetResults(v []TemplateTimelineEntry) {
	o.Results = v
}

func (o TemplateTimeline) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["count"] = o.Count
	}
	if o.NextAsOf != nil {
		toSerialize["next_as_of"] = o.NextAsOf
	}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullableTemplateTimeline struct {
	value *TemplateTimeline
	isSet bool
}

func (v NullableTemplateTimeline) Get() *TemplateTimeline {
	return v.value
}

func (v *NullableTemplateTimeline) Set(val *TemplateTimeline) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateTimeline) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateTimeline) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateTimeline(val *TemplateTimeline) *NullableTemplateTimeline {
	return &NullableTemplateTimeline{value: val, isSet: true}
}

func (v NullableTemplateTimeline) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateTimeline) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

