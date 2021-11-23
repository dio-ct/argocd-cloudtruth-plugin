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

// GitHubIntegrationCreate struct for GitHubIntegrationCreate
type GitHubIntegrationCreate struct {
	// An optional description for the integration.
	Description *string `json:"description,omitempty"`
	// Allow actions to write to the integration.
	Writable *bool `json:"writable,omitempty"`
	GhInstallationId int32 `json:"gh_installation_id"`
}

// NewGitHubIntegrationCreate instantiates a new GitHubIntegrationCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGitHubIntegrationCreate(ghInstallationId int32) *GitHubIntegrationCreate {
	this := GitHubIntegrationCreate{}
	this.GhInstallationId = ghInstallationId
	return &this
}

// NewGitHubIntegrationCreateWithDefaults instantiates a new GitHubIntegrationCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGitHubIntegrationCreateWithDefaults() *GitHubIntegrationCreate {
	this := GitHubIntegrationCreate{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *GitHubIntegrationCreate) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitHubIntegrationCreate) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *GitHubIntegrationCreate) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *GitHubIntegrationCreate) SetDescription(v string) {
	o.Description = &v
}

// GetWritable returns the Writable field value if set, zero value otherwise.
func (o *GitHubIntegrationCreate) GetWritable() bool {
	if o == nil || o.Writable == nil {
		var ret bool
		return ret
	}
	return *o.Writable
}

// GetWritableOk returns a tuple with the Writable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitHubIntegrationCreate) GetWritableOk() (*bool, bool) {
	if o == nil || o.Writable == nil {
		return nil, false
	}
	return o.Writable, true
}

// HasWritable returns a boolean if a field has been set.
func (o *GitHubIntegrationCreate) HasWritable() bool {
	if o != nil && o.Writable != nil {
		return true
	}

	return false
}

// SetWritable gets a reference to the given bool and assigns it to the Writable field.
func (o *GitHubIntegrationCreate) SetWritable(v bool) {
	o.Writable = &v
}

// GetGhInstallationId returns the GhInstallationId field value
func (o *GitHubIntegrationCreate) GetGhInstallationId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.GhInstallationId
}

// GetGhInstallationIdOk returns a tuple with the GhInstallationId field value
// and a boolean to check if the value has been set.
func (o *GitHubIntegrationCreate) GetGhInstallationIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.GhInstallationId, true
}

// SetGhInstallationId sets field value
func (o *GitHubIntegrationCreate) SetGhInstallationId(v int32) {
	o.GhInstallationId = v
}

func (o GitHubIntegrationCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Writable != nil {
		toSerialize["writable"] = o.Writable
	}
	if true {
		toSerialize["gh_installation_id"] = o.GhInstallationId
	}
	return json.Marshal(toSerialize)
}

type NullableGitHubIntegrationCreate struct {
	value *GitHubIntegrationCreate
	isSet bool
}

func (v NullableGitHubIntegrationCreate) Get() *GitHubIntegrationCreate {
	return v.value
}

func (v *NullableGitHubIntegrationCreate) Set(val *GitHubIntegrationCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableGitHubIntegrationCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableGitHubIntegrationCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGitHubIntegrationCreate(val *GitHubIntegrationCreate) *NullableGitHubIntegrationCreate {
	return &NullableGitHubIntegrationCreate{value: val, isSet: true}
}

func (v NullableGitHubIntegrationCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGitHubIntegrationCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


