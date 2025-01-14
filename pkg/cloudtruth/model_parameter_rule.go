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

// ParameterRule A type of `ModelSerializer` that uses hyperlinked relationships with compound keys instead of primary key relationships.  Specifically:  * A 'url' field is included instead of the 'id' field. * Relationships to other instances are hyperlinks, instead of primary keys.  NOTE: this only works with DRF 3.1.0 and above.
type ParameterRule struct {
	Url string `json:"url"`
	Id string `json:"id"`
	// The parameter this rule is for.
	Parameter string `json:"parameter"`
	Type ParameterRuleTypeEnum `json:"type"`
	Constraint string `json:"constraint"`
	CreatedAt time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
}

// NewParameterRule instantiates a new ParameterRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewParameterRule(url string, id string, parameter string, type_ ParameterRuleTypeEnum, constraint string, createdAt time.Time, modifiedAt time.Time) *ParameterRule {
	this := ParameterRule{}
	this.Url = url
	this.Id = id
	this.Parameter = parameter
	this.Type = type_
	this.Constraint = constraint
	this.CreatedAt = createdAt
	this.ModifiedAt = modifiedAt
	return &this
}

// NewParameterRuleWithDefaults instantiates a new ParameterRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewParameterRuleWithDefaults() *ParameterRule {
	this := ParameterRule{}
	return &this
}

// GetUrl returns the Url field value
func (o *ParameterRule) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *ParameterRule) GetUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *ParameterRule) SetUrl(v string) {
	o.Url = v
}

// GetId returns the Id field value
func (o *ParameterRule) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ParameterRule) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ParameterRule) SetId(v string) {
	o.Id = v
}

// GetParameter returns the Parameter field value
func (o *ParameterRule) GetParameter() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Parameter
}

// GetParameterOk returns a tuple with the Parameter field value
// and a boolean to check if the value has been set.
func (o *ParameterRule) GetParameterOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Parameter, true
}

// SetParameter sets field value
func (o *ParameterRule) SetParameter(v string) {
	o.Parameter = v
}

// GetType returns the Type field value
func (o *ParameterRule) GetType() ParameterRuleTypeEnum {
	if o == nil {
		var ret ParameterRuleTypeEnum
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ParameterRule) GetTypeOk() (*ParameterRuleTypeEnum, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ParameterRule) SetType(v ParameterRuleTypeEnum) {
	o.Type = v
}

// GetConstraint returns the Constraint field value
func (o *ParameterRule) GetConstraint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Constraint
}

// GetConstraintOk returns a tuple with the Constraint field value
// and a boolean to check if the value has been set.
func (o *ParameterRule) GetConstraintOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Constraint, true
}

// SetConstraint sets field value
func (o *ParameterRule) SetConstraint(v string) {
	o.Constraint = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *ParameterRule) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ParameterRule) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *ParameterRule) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetModifiedAt returns the ModifiedAt field value
func (o *ParameterRule) GetModifiedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value
// and a boolean to check if the value has been set.
func (o *ParameterRule) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ModifiedAt, true
}

// SetModifiedAt sets field value
func (o *ParameterRule) SetModifiedAt(v time.Time) {
	o.ModifiedAt = v
}

func (o ParameterRule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["url"] = o.Url
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["parameter"] = o.Parameter
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["constraint"] = o.Constraint
	}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if true {
		toSerialize["modified_at"] = o.ModifiedAt
	}
	return json.Marshal(toSerialize)
}

type NullableParameterRule struct {
	value *ParameterRule
	isSet bool
}

func (v NullableParameterRule) Get() *ParameterRule {
	return v.value
}

func (v *NullableParameterRule) Set(val *ParameterRule) {
	v.value = val
	v.isSet = true
}

func (v NullableParameterRule) IsSet() bool {
	return v.isSet
}

func (v *NullableParameterRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableParameterRule(val *ParameterRule) *NullableParameterRule {
	return &NullableParameterRule{value: val, isSet: true}
}

func (v NullableParameterRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableParameterRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


