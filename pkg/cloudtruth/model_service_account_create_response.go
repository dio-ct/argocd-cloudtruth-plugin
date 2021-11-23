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

// ServiceAccountCreateResponse struct for ServiceAccountCreateResponse
type ServiceAccountCreateResponse struct {
	Url string `json:"url"`
	Id string `json:"id"`
	User User `json:"user"`
	// An optional description of the process or system using the service account.
	Description *string `json:"description,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	// The most recent date and time the service account was used.
	LastUsedAt time.Time `json:"last_used_at"`
	// The API Key to use as a Bearer token for the service account.
	Apikey string `json:"apikey"`
}

// NewServiceAccountCreateResponse instantiates a new ServiceAccountCreateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceAccountCreateResponse(url string, id string, user User, createdAt time.Time, modifiedAt time.Time, lastUsedAt time.Time, apikey string) *ServiceAccountCreateResponse {
	this := ServiceAccountCreateResponse{}
	this.Url = url
	this.Id = id
	this.User = user
	this.CreatedAt = createdAt
	this.ModifiedAt = modifiedAt
	this.LastUsedAt = lastUsedAt
	this.Apikey = apikey
	return &this
}

// NewServiceAccountCreateResponseWithDefaults instantiates a new ServiceAccountCreateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceAccountCreateResponseWithDefaults() *ServiceAccountCreateResponse {
	this := ServiceAccountCreateResponse{}
	return &this
}

// GetUrl returns the Url field value
func (o *ServiceAccountCreateResponse) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *ServiceAccountCreateResponse) GetUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *ServiceAccountCreateResponse) SetUrl(v string) {
	o.Url = v
}

// GetId returns the Id field value
func (o *ServiceAccountCreateResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ServiceAccountCreateResponse) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ServiceAccountCreateResponse) SetId(v string) {
	o.Id = v
}

// GetUser returns the User field value
func (o *ServiceAccountCreateResponse) GetUser() User {
	if o == nil {
		var ret User
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *ServiceAccountCreateResponse) GetUserOk() (*User, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *ServiceAccountCreateResponse) SetUser(v User) {
	o.User = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ServiceAccountCreateResponse) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccountCreateResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ServiceAccountCreateResponse) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ServiceAccountCreateResponse) SetDescription(v string) {
	o.Description = &v
}

// GetCreatedAt returns the CreatedAt field value
func (o *ServiceAccountCreateResponse) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ServiceAccountCreateResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *ServiceAccountCreateResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetModifiedAt returns the ModifiedAt field value
func (o *ServiceAccountCreateResponse) GetModifiedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value
// and a boolean to check if the value has been set.
func (o *ServiceAccountCreateResponse) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ModifiedAt, true
}

// SetModifiedAt sets field value
func (o *ServiceAccountCreateResponse) SetModifiedAt(v time.Time) {
	o.ModifiedAt = v
}

// GetLastUsedAt returns the LastUsedAt field value
func (o *ServiceAccountCreateResponse) GetLastUsedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastUsedAt
}

// GetLastUsedAtOk returns a tuple with the LastUsedAt field value
// and a boolean to check if the value has been set.
func (o *ServiceAccountCreateResponse) GetLastUsedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LastUsedAt, true
}

// SetLastUsedAt sets field value
func (o *ServiceAccountCreateResponse) SetLastUsedAt(v time.Time) {
	o.LastUsedAt = v
}

// GetApikey returns the Apikey field value
func (o *ServiceAccountCreateResponse) GetApikey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Apikey
}

// GetApikeyOk returns a tuple with the Apikey field value
// and a boolean to check if the value has been set.
func (o *ServiceAccountCreateResponse) GetApikeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Apikey, true
}

// SetApikey sets field value
func (o *ServiceAccountCreateResponse) SetApikey(v string) {
	o.Apikey = v
}

func (o ServiceAccountCreateResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["url"] = o.Url
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["user"] = o.User
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if true {
		toSerialize["modified_at"] = o.ModifiedAt
	}
	if true {
		toSerialize["last_used_at"] = o.LastUsedAt
	}
	if true {
		toSerialize["apikey"] = o.Apikey
	}
	return json.Marshal(toSerialize)
}

type NullableServiceAccountCreateResponse struct {
	value *ServiceAccountCreateResponse
	isSet bool
}

func (v NullableServiceAccountCreateResponse) Get() *ServiceAccountCreateResponse {
	return v.value
}

func (v *NullableServiceAccountCreateResponse) Set(val *ServiceAccountCreateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceAccountCreateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceAccountCreateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceAccountCreateResponse(val *ServiceAccountCreateResponse) *NullableServiceAccountCreateResponse {
	return &NullableServiceAccountCreateResponse{value: val, isSet: true}
}

func (v NullableServiceAccountCreateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceAccountCreateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

