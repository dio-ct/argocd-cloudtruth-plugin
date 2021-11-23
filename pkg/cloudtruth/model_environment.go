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

// Environment struct for Environment
type Environment struct {
	Url string `json:"url"`
	// A unique identifier for the environment.
	Id string `json:"id"`
	// The environment name.
	Name string `json:"name"`
	// A description of the environment.  You may find it helpful to document how this environment is used to assist others when they need to maintain software that uses this content.
	Description *string `json:"description,omitempty"`
	// Environments can inherit from a single parent environment which provides values for parameters when specific environments do not have a value set.  Every organization has one default environment that cannot be removed.
	Parent NullableString `json:"parent,omitempty"`
	// This is the opposite of `parent`, see that field for more details.
	Children []string `json:"children"`
	CreatedAt time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
}

// NewEnvironment instantiates a new Environment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvironment(url string, id string, name string, children []string, createdAt time.Time, modifiedAt time.Time) *Environment {
	this := Environment{}
	this.Url = url
	this.Id = id
	this.Name = name
	this.Children = children
	this.CreatedAt = createdAt
	this.ModifiedAt = modifiedAt
	return &this
}

// NewEnvironmentWithDefaults instantiates a new Environment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvironmentWithDefaults() *Environment {
	this := Environment{}
	return &this
}

// GetUrl returns the Url field value
func (o *Environment) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *Environment) GetUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *Environment) SetUrl(v string) {
	o.Url = v
}

// GetId returns the Id field value
func (o *Environment) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Environment) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Environment) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *Environment) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Environment) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Environment) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Environment) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Environment) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Environment) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Environment) SetDescription(v string) {
	o.Description = &v
}

// GetParent returns the Parent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Environment) GetParent() string {
	if o == nil || o.Parent.Get() == nil {
		var ret string
		return ret
	}
	return *o.Parent.Get()
}

// GetParentOk returns a tuple with the Parent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Environment) GetParentOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Parent.Get(), o.Parent.IsSet()
}

// HasParent returns a boolean if a field has been set.
func (o *Environment) HasParent() bool {
	if o != nil && o.Parent.IsSet() {
		return true
	}

	return false
}

// SetParent gets a reference to the given NullableString and assigns it to the Parent field.
func (o *Environment) SetParent(v string) {
	o.Parent.Set(&v)
}
// SetParentNil sets the value for Parent to be an explicit nil
func (o *Environment) SetParentNil() {
	o.Parent.Set(nil)
}

// UnsetParent ensures that no value is present for Parent, not even an explicit nil
func (o *Environment) UnsetParent() {
	o.Parent.Unset()
}

// GetChildren returns the Children field value
func (o *Environment) GetChildren() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Children
}

// GetChildrenOk returns a tuple with the Children field value
// and a boolean to check if the value has been set.
func (o *Environment) GetChildrenOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Children, true
}

// SetChildren sets field value
func (o *Environment) SetChildren(v []string) {
	o.Children = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *Environment) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Environment) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *Environment) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetModifiedAt returns the ModifiedAt field value
func (o *Environment) GetModifiedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value
// and a boolean to check if the value has been set.
func (o *Environment) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ModifiedAt, true
}

// SetModifiedAt sets field value
func (o *Environment) SetModifiedAt(v time.Time) {
	o.ModifiedAt = v
}

func (o Environment) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["url"] = o.Url
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Parent.IsSet() {
		toSerialize["parent"] = o.Parent.Get()
	}
	if true {
		toSerialize["children"] = o.Children
	}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if true {
		toSerialize["modified_at"] = o.ModifiedAt
	}
	return json.Marshal(toSerialize)
}

type NullableEnvironment struct {
	value *Environment
	isSet bool
}

func (v NullableEnvironment) Get() *Environment {
	return v.value
}

func (v *NullableEnvironment) Set(val *Environment) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvironment) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvironment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvironment(val *Environment) *NullableEnvironment {
	return &NullableEnvironment{value: val, isSet: true}
}

func (v NullableEnvironment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvironment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

