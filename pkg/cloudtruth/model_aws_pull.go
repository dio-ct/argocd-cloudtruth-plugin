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

// AwsPull Pull actions can be configured to get configuration and secrets from integrations on demand.
type AwsPull struct {
	Url string `json:"url"`
	// Unique identifier for the action.
	Id string `json:"id"`
	// The action name.
	Name string `json:"name"`
	// The optional description for the action.
	Description *string `json:"description,omitempty"`
	// The most recent task run for this action.
	LatestTask NullableAwsPullTask `json:"latest_task"`
	CreatedAt time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	// When set to dry-run mode an action will report the changes that it would have made in task steps, however those changes are not actually performed.
	DryRun *bool `json:"dry_run,omitempty"`
	// The AWS region this pull uses.  This region must be enabled in the integration.
	Region AwsRegionEnum `json:"region"`
	// The AWS service this pull uses.  This service must be enabled in the integration.
	Service AwsServiceEnum `json:"service"`
	// Defines a path through the integration to the location where values will be pulled.  The following mustache-style substitutions must be used in the resource locator string:    - ``{{ environment }}`` to identify the environment name   - ``{{ parameter }}`` to identify the parameter name   - ``{{ project }}`` to identify the project name
	Resource string `json:"resource"`
}

// NewAwsPull instantiates a new AwsPull object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAwsPull(url string, id string, name string, latestTask NullableAwsPullTask, createdAt time.Time, modifiedAt time.Time, region AwsRegionEnum, service AwsServiceEnum, resource string) *AwsPull {
	this := AwsPull{}
	this.Url = url
	this.Id = id
	this.Name = name
	this.LatestTask = latestTask
	this.CreatedAt = createdAt
	this.ModifiedAt = modifiedAt
	this.Region = region
	this.Service = service
	this.Resource = resource
	return &this
}

// NewAwsPullWithDefaults instantiates a new AwsPull object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAwsPullWithDefaults() *AwsPull {
	this := AwsPull{}
	return &this
}

// GetUrl returns the Url field value
func (o *AwsPull) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *AwsPull) GetUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *AwsPull) SetUrl(v string) {
	o.Url = v
}

// GetId returns the Id field value
func (o *AwsPull) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AwsPull) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AwsPull) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *AwsPull) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AwsPull) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AwsPull) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AwsPull) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsPull) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AwsPull) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AwsPull) SetDescription(v string) {
	o.Description = &v
}

// GetLatestTask returns the LatestTask field value
// If the value is explicit nil, the zero value for AwsPullTask will be returned
func (o *AwsPull) GetLatestTask() AwsPullTask {
	if o == nil || o.LatestTask.Get() == nil {
		var ret AwsPullTask
		return ret
	}

	return *o.LatestTask.Get()
}

// GetLatestTaskOk returns a tuple with the LatestTask field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AwsPull) GetLatestTaskOk() (*AwsPullTask, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LatestTask.Get(), o.LatestTask.IsSet()
}

// SetLatestTask sets field value
func (o *AwsPull) SetLatestTask(v AwsPullTask) {
	o.LatestTask.Set(&v)
}

// GetCreatedAt returns the CreatedAt field value
func (o *AwsPull) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *AwsPull) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *AwsPull) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetModifiedAt returns the ModifiedAt field value
func (o *AwsPull) GetModifiedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value
// and a boolean to check if the value has been set.
func (o *AwsPull) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ModifiedAt, true
}

// SetModifiedAt sets field value
func (o *AwsPull) SetModifiedAt(v time.Time) {
	o.ModifiedAt = v
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *AwsPull) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsPull) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *AwsPull) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *AwsPull) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetRegion returns the Region field value
func (o *AwsPull) GetRegion() AwsRegionEnum {
	if o == nil {
		var ret AwsRegionEnum
		return ret
	}

	return o.Region
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
func (o *AwsPull) GetRegionOk() (*AwsRegionEnum, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Region, true
}

// SetRegion sets field value
func (o *AwsPull) SetRegion(v AwsRegionEnum) {
	o.Region = v
}

// GetService returns the Service field value
func (o *AwsPull) GetService() AwsServiceEnum {
	if o == nil {
		var ret AwsServiceEnum
		return ret
	}

	return o.Service
}

// GetServiceOk returns a tuple with the Service field value
// and a boolean to check if the value has been set.
func (o *AwsPull) GetServiceOk() (*AwsServiceEnum, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Service, true
}

// SetService sets field value
func (o *AwsPull) SetService(v AwsServiceEnum) {
	o.Service = v
}

// GetResource returns the Resource field value
func (o *AwsPull) GetResource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Resource
}

// GetResourceOk returns a tuple with the Resource field value
// and a boolean to check if the value has been set.
func (o *AwsPull) GetResourceOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Resource, true
}

// SetResource sets field value
func (o *AwsPull) SetResource(v string) {
	o.Resource = v
}

func (o AwsPull) MarshalJSON() ([]byte, error) {
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
	if true {
		toSerialize["latest_task"] = o.LatestTask.Get()
	}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if true {
		toSerialize["modified_at"] = o.ModifiedAt
	}
	if o.DryRun != nil {
		toSerialize["dry_run"] = o.DryRun
	}
	if true {
		toSerialize["region"] = o.Region
	}
	if true {
		toSerialize["service"] = o.Service
	}
	if true {
		toSerialize["resource"] = o.Resource
	}
	return json.Marshal(toSerialize)
}

type NullableAwsPull struct {
	value *AwsPull
	isSet bool
}

func (v NullableAwsPull) Get() *AwsPull {
	return v.value
}

func (v *NullableAwsPull) Set(val *AwsPull) {
	v.value = val
	v.isSet = true
}

func (v NullableAwsPull) IsSet() bool {
	return v.isSet
}

func (v *NullableAwsPull) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAwsPull(val *AwsPull) *NullableAwsPull {
	return &NullableAwsPull{value: val, isSet: true}
}

func (v NullableAwsPull) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAwsPull) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

