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

// AuditTrail struct for AuditTrail
type AuditTrail struct {
	Url string `json:"url"`
	// A unique identifier for the audit record.
	Id string `json:"id"`
	// The action that was taken.
	Action string `json:"action"`
	// The id of the object associated with the action.
	ObjectId string `json:"object_id"`
	// The name of the object associated with the action, if applicable.
	ObjectName string `json:"object_name"`
	// The type of object associated with the action.
	ObjectType ObjectTypeEnum `json:"object_type"`
	// The timestamp of the activity that was audited.
	Timestamp time.Time `json:"timestamp"`
	// Details of the user associated with this change.
	User User `json:"user"`
}

// NewAuditTrail instantiates a new AuditTrail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuditTrail(url string, id string, action string, objectId string, objectName string, objectType ObjectTypeEnum, timestamp time.Time, user User) *AuditTrail {
	this := AuditTrail{}
	this.Url = url
	this.Id = id
	this.Action = action
	this.ObjectId = objectId
	this.ObjectName = objectName
	this.ObjectType = objectType
	this.Timestamp = timestamp
	this.User = user
	return &this
}

// NewAuditTrailWithDefaults instantiates a new AuditTrail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuditTrailWithDefaults() *AuditTrail {
	this := AuditTrail{}
	return &this
}

// GetUrl returns the Url field value
func (o *AuditTrail) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *AuditTrail) GetUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *AuditTrail) SetUrl(v string) {
	o.Url = v
}

// GetId returns the Id field value
func (o *AuditTrail) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AuditTrail) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AuditTrail) SetId(v string) {
	o.Id = v
}

// GetAction returns the Action field value
func (o *AuditTrail) GetAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *AuditTrail) GetActionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *AuditTrail) SetAction(v string) {
	o.Action = v
}

// GetObjectId returns the ObjectId field value
func (o *AuditTrail) GetObjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectId
}

// GetObjectIdOk returns a tuple with the ObjectId field value
// and a boolean to check if the value has been set.
func (o *AuditTrail) GetObjectIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectId, true
}

// SetObjectId sets field value
func (o *AuditTrail) SetObjectId(v string) {
	o.ObjectId = v
}

// GetObjectName returns the ObjectName field value
func (o *AuditTrail) GetObjectName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectName
}

// GetObjectNameOk returns a tuple with the ObjectName field value
// and a boolean to check if the value has been set.
func (o *AuditTrail) GetObjectNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectName, true
}

// SetObjectName sets field value
func (o *AuditTrail) SetObjectName(v string) {
	o.ObjectName = v
}

// GetObjectType returns the ObjectType field value
func (o *AuditTrail) GetObjectType() ObjectTypeEnum {
	if o == nil {
		var ret ObjectTypeEnum
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *AuditTrail) GetObjectTypeOk() (*ObjectTypeEnum, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *AuditTrail) SetObjectType(v ObjectTypeEnum) {
	o.ObjectType = v
}

// GetTimestamp returns the Timestamp field value
func (o *AuditTrail) GetTimestamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *AuditTrail) GetTimestampOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *AuditTrail) SetTimestamp(v time.Time) {
	o.Timestamp = v
}

// GetUser returns the User field value
func (o *AuditTrail) GetUser() User {
	if o == nil {
		var ret User
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *AuditTrail) GetUserOk() (*User, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *AuditTrail) SetUser(v User) {
	o.User = v
}

func (o AuditTrail) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["url"] = o.Url
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["action"] = o.Action
	}
	if true {
		toSerialize["object_id"] = o.ObjectId
	}
	if true {
		toSerialize["object_name"] = o.ObjectName
	}
	if true {
		toSerialize["object_type"] = o.ObjectType
	}
	if true {
		toSerialize["timestamp"] = o.Timestamp
	}
	if true {
		toSerialize["user"] = o.User
	}
	return json.Marshal(toSerialize)
}

type NullableAuditTrail struct {
	value *AuditTrail
	isSet bool
}

func (v NullableAuditTrail) Get() *AuditTrail {
	return v.value
}

func (v *NullableAuditTrail) Set(val *AuditTrail) {
	v.value = val
	v.isSet = true
}

func (v NullableAuditTrail) IsSet() bool {
	return v.isSet
}

func (v *NullableAuditTrail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuditTrail(val *AuditTrail) *NullableAuditTrail {
	return &NullableAuditTrail{value: val, isSet: true}
}

func (v NullableAuditTrail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuditTrail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


