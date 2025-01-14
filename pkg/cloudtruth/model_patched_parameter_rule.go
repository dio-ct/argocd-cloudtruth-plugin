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

// PatchedParameterRule A type of `ModelSerializer` that uses hyperlinked relationships with compound keys instead of primary key relationships.  Specifically:  * A 'url' field is included instead of the 'id' field. * Relationships to other instances are hyperlinks, instead of primary keys.  NOTE: this only works with DRF 3.1.0 and above.
type PatchedParameterRule struct {
	Url *string `json:"url,omitempty"`
	Id *string `json:"id,omitempty"`
	// The parameter this rule is for.
	Parameter *string `json:"parameter,omitempty"`
	Type *ParameterRuleTypeEnum `json:"type,omitempty"`
	Constraint *string `json:"constraint,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	ModifiedAt *time.Time `json:"modified_at,omitempty"`
}

// NewPatchedParameterRule instantiates a new PatchedParameterRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedParameterRule() *PatchedParameterRule {
	this := PatchedParameterRule{}
	return &this
}

// NewPatchedParameterRuleWithDefaults instantiates a new PatchedParameterRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedParameterRuleWithDefaults() *PatchedParameterRule {
	this := PatchedParameterRule{}
	return &this
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *PatchedParameterRule) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedParameterRule) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *PatchedParameterRule) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *PatchedParameterRule) SetUrl(v string) {
	o.Url = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PatchedParameterRule) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedParameterRule) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PatchedParameterRule) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PatchedParameterRule) SetId(v string) {
	o.Id = &v
}

// GetParameter returns the Parameter field value if set, zero value otherwise.
func (o *PatchedParameterRule) GetParameter() string {
	if o == nil || o.Parameter == nil {
		var ret string
		return ret
	}
	return *o.Parameter
}

// GetParameterOk returns a tuple with the Parameter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedParameterRule) GetParameterOk() (*string, bool) {
	if o == nil || o.Parameter == nil {
		return nil, false
	}
	return o.Parameter, true
}

// HasParameter returns a boolean if a field has been set.
func (o *PatchedParameterRule) HasParameter() bool {
	if o != nil && o.Parameter != nil {
		return true
	}

	return false
}

// SetParameter gets a reference to the given string and assigns it to the Parameter field.
func (o *PatchedParameterRule) SetParameter(v string) {
	o.Parameter = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PatchedParameterRule) GetType() ParameterRuleTypeEnum {
	if o == nil || o.Type == nil {
		var ret ParameterRuleTypeEnum
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedParameterRule) GetTypeOk() (*ParameterRuleTypeEnum, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PatchedParameterRule) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given ParameterRuleTypeEnum and assigns it to the Type field.
func (o *PatchedParameterRule) SetType(v ParameterRuleTypeEnum) {
	o.Type = &v
}

// GetConstraint returns the Constraint field value if set, zero value otherwise.
func (o *PatchedParameterRule) GetConstraint() string {
	if o == nil || o.Constraint == nil {
		var ret string
		return ret
	}
	return *o.Constraint
}

// GetConstraintOk returns a tuple with the Constraint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedParameterRule) GetConstraintOk() (*string, bool) {
	if o == nil || o.Constraint == nil {
		return nil, false
	}
	return o.Constraint, true
}

// HasConstraint returns a boolean if a field has been set.
func (o *PatchedParameterRule) HasConstraint() bool {
	if o != nil && o.Constraint != nil {
		return true
	}

	return false
}

// SetConstraint gets a reference to the given string and assigns it to the Constraint field.
func (o *PatchedParameterRule) SetConstraint(v string) {
	o.Constraint = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *PatchedParameterRule) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedParameterRule) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *PatchedParameterRule) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *PatchedParameterRule) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise.
func (o *PatchedParameterRule) GetModifiedAt() time.Time {
	if o == nil || o.ModifiedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedParameterRule) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil || o.ModifiedAt == nil {
		return nil, false
	}
	return o.ModifiedAt, true
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *PatchedParameterRule) HasModifiedAt() bool {
	if o != nil && o.ModifiedAt != nil {
		return true
	}

	return false
}

// SetModifiedAt gets a reference to the given time.Time and assigns it to the ModifiedAt field.
func (o *PatchedParameterRule) SetModifiedAt(v time.Time) {
	o.ModifiedAt = &v
}

func (o PatchedParameterRule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Parameter != nil {
		toSerialize["parameter"] = o.Parameter
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Constraint != nil {
		toSerialize["constraint"] = o.Constraint
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.ModifiedAt != nil {
		toSerialize["modified_at"] = o.ModifiedAt
	}
	return json.Marshal(toSerialize)
}

type NullablePatchedParameterRule struct {
	value *PatchedParameterRule
	isSet bool
}

func (v NullablePatchedParameterRule) Get() *PatchedParameterRule {
	return v.value
}

func (v *NullablePatchedParameterRule) Set(val *PatchedParameterRule) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedParameterRule) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedParameterRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedParameterRule(val *PatchedParameterRule) *NullablePatchedParameterRule {
	return &NullablePatchedParameterRule{value: val, isSet: true}
}

func (v NullablePatchedParameterRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedParameterRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


