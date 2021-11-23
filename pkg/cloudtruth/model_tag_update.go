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

// TagUpdate Details for updating a tag.
type TagUpdate struct {
	// A unique identifier for the tag.
	Id string `json:"id"`
	// The tag name. Tag names may contain alphanumeric, hyphen, underscore, or period characters. Tag names are case sensitive. The name cannot be modified.
	Name string `json:"name"`
	// A description of the tag.  You may find it helpful to document how this tag is used to assist others when they need to maintain software that uses this content.
	Description *string `json:"description,omitempty"`
	// The point in time this tag represents.  If explicitly set to `null` then the current time will be used.
	Timestamp NullableTime `json:"timestamp,omitempty"`
}

// NewTagUpdate instantiates a new TagUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTagUpdate(id string, name string) *TagUpdate {
	this := TagUpdate{}
	this.Id = id
	this.Name = name
	return &this
}

// NewTagUpdateWithDefaults instantiates a new TagUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTagUpdateWithDefaults() *TagUpdate {
	this := TagUpdate{}
	return &this
}

// GetId returns the Id field value
func (o *TagUpdate) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TagUpdate) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TagUpdate) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *TagUpdate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TagUpdate) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TagUpdate) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *TagUpdate) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagUpdate) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *TagUpdate) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *TagUpdate) SetDescription(v string) {
	o.Description = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TagUpdate) GetTimestamp() time.Time {
	if o == nil || o.Timestamp.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.Timestamp.Get()
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TagUpdate) GetTimestampOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Timestamp.Get(), o.Timestamp.IsSet()
}

// HasTimestamp returns a boolean if a field has been set.
func (o *TagUpdate) HasTimestamp() bool {
	if o != nil && o.Timestamp.IsSet() {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given NullableTime and assigns it to the Timestamp field.
func (o *TagUpdate) SetTimestamp(v time.Time) {
	o.Timestamp.Set(&v)
}
// SetTimestampNil sets the value for Timestamp to be an explicit nil
func (o *TagUpdate) SetTimestampNil() {
	o.Timestamp.Set(nil)
}

// UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
func (o *TagUpdate) UnsetTimestamp() {
	o.Timestamp.Unset()
}

func (o TagUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Timestamp.IsSet() {
		toSerialize["timestamp"] = o.Timestamp.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableTagUpdate struct {
	value *TagUpdate
	isSet bool
}

func (v NullableTagUpdate) Get() *TagUpdate {
	return v.value
}

func (v *NullableTagUpdate) Set(val *TagUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableTagUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableTagUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTagUpdate(val *TagUpdate) *NullableTagUpdate {
	return &NullableTagUpdate{value: val, isSet: true}
}

func (v NullableTagUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTagUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


