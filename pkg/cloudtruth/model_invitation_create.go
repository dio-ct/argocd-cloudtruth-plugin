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

// InvitationCreate struct for InvitationCreate
type InvitationCreate struct {
	// The email address of the user to be invited.
	Email string `json:"email"`
	// The role that the user will have in the organization, should the user accept.
	Role RoleEnum `json:"role"`
}

// NewInvitationCreate instantiates a new InvitationCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvitationCreate(email string, role RoleEnum) *InvitationCreate {
	this := InvitationCreate{}
	this.Email = email
	this.Role = role
	return &this
}

// NewInvitationCreateWithDefaults instantiates a new InvitationCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvitationCreateWithDefaults() *InvitationCreate {
	this := InvitationCreate{}
	return &this
}

// GetEmail returns the Email field value
func (o *InvitationCreate) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *InvitationCreate) GetEmailOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *InvitationCreate) SetEmail(v string) {
	o.Email = v
}

// GetRole returns the Role field value
func (o *InvitationCreate) GetRole() RoleEnum {
	if o == nil {
		var ret RoleEnum
		return ret
	}

	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *InvitationCreate) GetRoleOk() (*RoleEnum, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value
func (o *InvitationCreate) SetRole(v RoleEnum) {
	o.Role = v
}

func (o InvitationCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["email"] = o.Email
	}
	if true {
		toSerialize["role"] = o.Role
	}
	return json.Marshal(toSerialize)
}

type NullableInvitationCreate struct {
	value *InvitationCreate
	isSet bool
}

func (v NullableInvitationCreate) Get() *InvitationCreate {
	return v.value
}

func (v *NullableInvitationCreate) Set(val *InvitationCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableInvitationCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableInvitationCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvitationCreate(val *InvitationCreate) *NullableInvitationCreate {
	return &NullableInvitationCreate{value: val, isSet: true}
}

func (v NullableInvitationCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvitationCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


