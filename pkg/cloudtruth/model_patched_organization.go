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

// PatchedOrganization struct for PatchedOrganization
type PatchedOrganization struct {
	Url *string `json:"url,omitempty"`
	// A unique identifier for the organization.
	Id *string `json:"id,omitempty"`
	// The organization name.
	Name *string `json:"name,omitempty"`
	// Indicates if this Organization is the one currently targeted by the Bearer token used by the client to authorize.
	Current *bool `json:"current,omitempty"`
	SubscriptionExpiresAt NullableTime `json:"subscription_expires_at,omitempty"`
	SubscriptionId NullableString `json:"subscription_id,omitempty"`
	SubscriptionPlanId NullableString `json:"subscription_plan_id,omitempty"`
	SubscriptionPlanName NullableString `json:"subscription_plan_name,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	ModifiedAt *time.Time `json:"modified_at,omitempty"`
}

// NewPatchedOrganization instantiates a new PatchedOrganization object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedOrganization() *PatchedOrganization {
	this := PatchedOrganization{}
	return &this
}

// NewPatchedOrganizationWithDefaults instantiates a new PatchedOrganization object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedOrganizationWithDefaults() *PatchedOrganization {
	this := PatchedOrganization{}
	return &this
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *PatchedOrganization) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedOrganization) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *PatchedOrganization) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *PatchedOrganization) SetUrl(v string) {
	o.Url = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PatchedOrganization) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedOrganization) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PatchedOrganization) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PatchedOrganization) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedOrganization) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedOrganization) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedOrganization) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedOrganization) SetName(v string) {
	o.Name = &v
}

// GetCurrent returns the Current field value if set, zero value otherwise.
func (o *PatchedOrganization) GetCurrent() bool {
	if o == nil || o.Current == nil {
		var ret bool
		return ret
	}
	return *o.Current
}

// GetCurrentOk returns a tuple with the Current field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedOrganization) GetCurrentOk() (*bool, bool) {
	if o == nil || o.Current == nil {
		return nil, false
	}
	return o.Current, true
}

// HasCurrent returns a boolean if a field has been set.
func (o *PatchedOrganization) HasCurrent() bool {
	if o != nil && o.Current != nil {
		return true
	}

	return false
}

// SetCurrent gets a reference to the given bool and assigns it to the Current field.
func (o *PatchedOrganization) SetCurrent(v bool) {
	o.Current = &v
}

// GetSubscriptionExpiresAt returns the SubscriptionExpiresAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedOrganization) GetSubscriptionExpiresAt() time.Time {
	if o == nil || o.SubscriptionExpiresAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.SubscriptionExpiresAt.Get()
}

// GetSubscriptionExpiresAtOk returns a tuple with the SubscriptionExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedOrganization) GetSubscriptionExpiresAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SubscriptionExpiresAt.Get(), o.SubscriptionExpiresAt.IsSet()
}

// HasSubscriptionExpiresAt returns a boolean if a field has been set.
func (o *PatchedOrganization) HasSubscriptionExpiresAt() bool {
	if o != nil && o.SubscriptionExpiresAt.IsSet() {
		return true
	}

	return false
}

// SetSubscriptionExpiresAt gets a reference to the given NullableTime and assigns it to the SubscriptionExpiresAt field.
func (o *PatchedOrganization) SetSubscriptionExpiresAt(v time.Time) {
	o.SubscriptionExpiresAt.Set(&v)
}
// SetSubscriptionExpiresAtNil sets the value for SubscriptionExpiresAt to be an explicit nil
func (o *PatchedOrganization) SetSubscriptionExpiresAtNil() {
	o.SubscriptionExpiresAt.Set(nil)
}

// UnsetSubscriptionExpiresAt ensures that no value is present for SubscriptionExpiresAt, not even an explicit nil
func (o *PatchedOrganization) UnsetSubscriptionExpiresAt() {
	o.SubscriptionExpiresAt.Unset()
}

// GetSubscriptionId returns the SubscriptionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedOrganization) GetSubscriptionId() string {
	if o == nil || o.SubscriptionId.Get() == nil {
		var ret string
		return ret
	}
	return *o.SubscriptionId.Get()
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedOrganization) GetSubscriptionIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SubscriptionId.Get(), o.SubscriptionId.IsSet()
}

// HasSubscriptionId returns a boolean if a field has been set.
func (o *PatchedOrganization) HasSubscriptionId() bool {
	if o != nil && o.SubscriptionId.IsSet() {
		return true
	}

	return false
}

// SetSubscriptionId gets a reference to the given NullableString and assigns it to the SubscriptionId field.
func (o *PatchedOrganization) SetSubscriptionId(v string) {
	o.SubscriptionId.Set(&v)
}
// SetSubscriptionIdNil sets the value for SubscriptionId to be an explicit nil
func (o *PatchedOrganization) SetSubscriptionIdNil() {
	o.SubscriptionId.Set(nil)
}

// UnsetSubscriptionId ensures that no value is present for SubscriptionId, not even an explicit nil
func (o *PatchedOrganization) UnsetSubscriptionId() {
	o.SubscriptionId.Unset()
}

// GetSubscriptionPlanId returns the SubscriptionPlanId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedOrganization) GetSubscriptionPlanId() string {
	if o == nil || o.SubscriptionPlanId.Get() == nil {
		var ret string
		return ret
	}
	return *o.SubscriptionPlanId.Get()
}

// GetSubscriptionPlanIdOk returns a tuple with the SubscriptionPlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedOrganization) GetSubscriptionPlanIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SubscriptionPlanId.Get(), o.SubscriptionPlanId.IsSet()
}

// HasSubscriptionPlanId returns a boolean if a field has been set.
func (o *PatchedOrganization) HasSubscriptionPlanId() bool {
	if o != nil && o.SubscriptionPlanId.IsSet() {
		return true
	}

	return false
}

// SetSubscriptionPlanId gets a reference to the given NullableString and assigns it to the SubscriptionPlanId field.
func (o *PatchedOrganization) SetSubscriptionPlanId(v string) {
	o.SubscriptionPlanId.Set(&v)
}
// SetSubscriptionPlanIdNil sets the value for SubscriptionPlanId to be an explicit nil
func (o *PatchedOrganization) SetSubscriptionPlanIdNil() {
	o.SubscriptionPlanId.Set(nil)
}

// UnsetSubscriptionPlanId ensures that no value is present for SubscriptionPlanId, not even an explicit nil
func (o *PatchedOrganization) UnsetSubscriptionPlanId() {
	o.SubscriptionPlanId.Unset()
}

// GetSubscriptionPlanName returns the SubscriptionPlanName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedOrganization) GetSubscriptionPlanName() string {
	if o == nil || o.SubscriptionPlanName.Get() == nil {
		var ret string
		return ret
	}
	return *o.SubscriptionPlanName.Get()
}

// GetSubscriptionPlanNameOk returns a tuple with the SubscriptionPlanName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedOrganization) GetSubscriptionPlanNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SubscriptionPlanName.Get(), o.SubscriptionPlanName.IsSet()
}

// HasSubscriptionPlanName returns a boolean if a field has been set.
func (o *PatchedOrganization) HasSubscriptionPlanName() bool {
	if o != nil && o.SubscriptionPlanName.IsSet() {
		return true
	}

	return false
}

// SetSubscriptionPlanName gets a reference to the given NullableString and assigns it to the SubscriptionPlanName field.
func (o *PatchedOrganization) SetSubscriptionPlanName(v string) {
	o.SubscriptionPlanName.Set(&v)
}
// SetSubscriptionPlanNameNil sets the value for SubscriptionPlanName to be an explicit nil
func (o *PatchedOrganization) SetSubscriptionPlanNameNil() {
	o.SubscriptionPlanName.Set(nil)
}

// UnsetSubscriptionPlanName ensures that no value is present for SubscriptionPlanName, not even an explicit nil
func (o *PatchedOrganization) UnsetSubscriptionPlanName() {
	o.SubscriptionPlanName.Unset()
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *PatchedOrganization) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedOrganization) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *PatchedOrganization) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *PatchedOrganization) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise.
func (o *PatchedOrganization) GetModifiedAt() time.Time {
	if o == nil || o.ModifiedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedOrganization) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil || o.ModifiedAt == nil {
		return nil, false
	}
	return o.ModifiedAt, true
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *PatchedOrganization) HasModifiedAt() bool {
	if o != nil && o.ModifiedAt != nil {
		return true
	}

	return false
}

// SetModifiedAt gets a reference to the given time.Time and assigns it to the ModifiedAt field.
func (o *PatchedOrganization) SetModifiedAt(v time.Time) {
	o.ModifiedAt = &v
}

func (o PatchedOrganization) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Current != nil {
		toSerialize["current"] = o.Current
	}
	if o.SubscriptionExpiresAt.IsSet() {
		toSerialize["subscription_expires_at"] = o.SubscriptionExpiresAt.Get()
	}
	if o.SubscriptionId.IsSet() {
		toSerialize["subscription_id"] = o.SubscriptionId.Get()
	}
	if o.SubscriptionPlanId.IsSet() {
		toSerialize["subscription_plan_id"] = o.SubscriptionPlanId.Get()
	}
	if o.SubscriptionPlanName.IsSet() {
		toSerialize["subscription_plan_name"] = o.SubscriptionPlanName.Get()
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.ModifiedAt != nil {
		toSerialize["modified_at"] = o.ModifiedAt
	}
	return json.Marshal(toSerialize)
}

type NullablePatchedOrganization struct {
	value *PatchedOrganization
	isSet bool
}

func (v NullablePatchedOrganization) Get() *PatchedOrganization {
	return v.value
}

func (v *NullablePatchedOrganization) Set(val *PatchedOrganization) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedOrganization) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedOrganization) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedOrganization(val *PatchedOrganization) *NullablePatchedOrganization {
	return &NullablePatchedOrganization{value: val, isSet: true}
}

func (v NullablePatchedOrganization) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedOrganization) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


