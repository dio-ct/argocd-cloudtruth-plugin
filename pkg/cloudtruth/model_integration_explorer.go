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

// IntegrationExplorer Describes the content available at a given location.
type IntegrationExplorer struct {
	Fqn string `json:"fqn"`
	NodeType NodeTypeEnum `json:"node_type"`
	Secret *bool `json:"secret,omitempty"`
	Name *string `json:"name,omitempty"`
	ContentType NullableString `json:"content_type,omitempty"`
	ContentData NullableString `json:"content_data,omitempty"`
	ContentSize NullableInt32 `json:"content_size,omitempty"`
	ContentKeys []string `json:"content_keys,omitempty"`
}

// NewIntegrationExplorer instantiates a new IntegrationExplorer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIntegrationExplorer(fqn string, nodeType NodeTypeEnum) *IntegrationExplorer {
	this := IntegrationExplorer{}
	this.Fqn = fqn
	this.NodeType = nodeType
	return &this
}

// NewIntegrationExplorerWithDefaults instantiates a new IntegrationExplorer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIntegrationExplorerWithDefaults() *IntegrationExplorer {
	this := IntegrationExplorer{}
	return &this
}

// GetFqn returns the Fqn field value
func (o *IntegrationExplorer) GetFqn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Fqn
}

// GetFqnOk returns a tuple with the Fqn field value
// and a boolean to check if the value has been set.
func (o *IntegrationExplorer) GetFqnOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Fqn, true
}

// SetFqn sets field value
func (o *IntegrationExplorer) SetFqn(v string) {
	o.Fqn = v
}

// GetNodeType returns the NodeType field value
func (o *IntegrationExplorer) GetNodeType() NodeTypeEnum {
	if o == nil {
		var ret NodeTypeEnum
		return ret
	}

	return o.NodeType
}

// GetNodeTypeOk returns a tuple with the NodeType field value
// and a boolean to check if the value has been set.
func (o *IntegrationExplorer) GetNodeTypeOk() (*NodeTypeEnum, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.NodeType, true
}

// SetNodeType sets field value
func (o *IntegrationExplorer) SetNodeType(v NodeTypeEnum) {
	o.NodeType = v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *IntegrationExplorer) GetSecret() bool {
	if o == nil || o.Secret == nil {
		var ret bool
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationExplorer) GetSecretOk() (*bool, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *IntegrationExplorer) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given bool and assigns it to the Secret field.
func (o *IntegrationExplorer) SetSecret(v bool) {
	o.Secret = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *IntegrationExplorer) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationExplorer) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *IntegrationExplorer) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *IntegrationExplorer) SetName(v string) {
	o.Name = &v
}

// GetContentType returns the ContentType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IntegrationExplorer) GetContentType() string {
	if o == nil || o.ContentType.Get() == nil {
		var ret string
		return ret
	}
	return *o.ContentType.Get()
}

// GetContentTypeOk returns a tuple with the ContentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IntegrationExplorer) GetContentTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ContentType.Get(), o.ContentType.IsSet()
}

// HasContentType returns a boolean if a field has been set.
func (o *IntegrationExplorer) HasContentType() bool {
	if o != nil && o.ContentType.IsSet() {
		return true
	}

	return false
}

// SetContentType gets a reference to the given NullableString and assigns it to the ContentType field.
func (o *IntegrationExplorer) SetContentType(v string) {
	o.ContentType.Set(&v)
}
// SetContentTypeNil sets the value for ContentType to be an explicit nil
func (o *IntegrationExplorer) SetContentTypeNil() {
	o.ContentType.Set(nil)
}

// UnsetContentType ensures that no value is present for ContentType, not even an explicit nil
func (o *IntegrationExplorer) UnsetContentType() {
	o.ContentType.Unset()
}

// GetContentData returns the ContentData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IntegrationExplorer) GetContentData() string {
	if o == nil || o.ContentData.Get() == nil {
		var ret string
		return ret
	}
	return *o.ContentData.Get()
}

// GetContentDataOk returns a tuple with the ContentData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IntegrationExplorer) GetContentDataOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ContentData.Get(), o.ContentData.IsSet()
}

// HasContentData returns a boolean if a field has been set.
func (o *IntegrationExplorer) HasContentData() bool {
	if o != nil && o.ContentData.IsSet() {
		return true
	}

	return false
}

// SetContentData gets a reference to the given NullableString and assigns it to the ContentData field.
func (o *IntegrationExplorer) SetContentData(v string) {
	o.ContentData.Set(&v)
}
// SetContentDataNil sets the value for ContentData to be an explicit nil
func (o *IntegrationExplorer) SetContentDataNil() {
	o.ContentData.Set(nil)
}

// UnsetContentData ensures that no value is present for ContentData, not even an explicit nil
func (o *IntegrationExplorer) UnsetContentData() {
	o.ContentData.Unset()
}

// GetContentSize returns the ContentSize field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IntegrationExplorer) GetContentSize() int32 {
	if o == nil || o.ContentSize.Get() == nil {
		var ret int32
		return ret
	}
	return *o.ContentSize.Get()
}

// GetContentSizeOk returns a tuple with the ContentSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IntegrationExplorer) GetContentSizeOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ContentSize.Get(), o.ContentSize.IsSet()
}

// HasContentSize returns a boolean if a field has been set.
func (o *IntegrationExplorer) HasContentSize() bool {
	if o != nil && o.ContentSize.IsSet() {
		return true
	}

	return false
}

// SetContentSize gets a reference to the given NullableInt32 and assigns it to the ContentSize field.
func (o *IntegrationExplorer) SetContentSize(v int32) {
	o.ContentSize.Set(&v)
}
// SetContentSizeNil sets the value for ContentSize to be an explicit nil
func (o *IntegrationExplorer) SetContentSizeNil() {
	o.ContentSize.Set(nil)
}

// UnsetContentSize ensures that no value is present for ContentSize, not even an explicit nil
func (o *IntegrationExplorer) UnsetContentSize() {
	o.ContentSize.Unset()
}

// GetContentKeys returns the ContentKeys field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IntegrationExplorer) GetContentKeys() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.ContentKeys
}

// GetContentKeysOk returns a tuple with the ContentKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IntegrationExplorer) GetContentKeysOk() (*[]string, bool) {
	if o == nil || o.ContentKeys == nil {
		return nil, false
	}
	return &o.ContentKeys, true
}

// HasContentKeys returns a boolean if a field has been set.
func (o *IntegrationExplorer) HasContentKeys() bool {
	if o != nil && o.ContentKeys != nil {
		return true
	}

	return false
}

// SetContentKeys gets a reference to the given []string and assigns it to the ContentKeys field.
func (o *IntegrationExplorer) SetContentKeys(v []string) {
	o.ContentKeys = v
}

func (o IntegrationExplorer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["fqn"] = o.Fqn
	}
	if true {
		toSerialize["node_type"] = o.NodeType
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.ContentType.IsSet() {
		toSerialize["content_type"] = o.ContentType.Get()
	}
	if o.ContentData.IsSet() {
		toSerialize["content_data"] = o.ContentData.Get()
	}
	if o.ContentSize.IsSet() {
		toSerialize["content_size"] = o.ContentSize.Get()
	}
	if o.ContentKeys != nil {
		toSerialize["content_keys"] = o.ContentKeys
	}
	return json.Marshal(toSerialize)
}

type NullableIntegrationExplorer struct {
	value *IntegrationExplorer
	isSet bool
}

func (v NullableIntegrationExplorer) Get() *IntegrationExplorer {
	return v.value
}

func (v *NullableIntegrationExplorer) Set(val *IntegrationExplorer) {
	v.value = val
	v.isSet = true
}

func (v NullableIntegrationExplorer) IsSet() bool {
	return v.isSet
}

func (v *NullableIntegrationExplorer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIntegrationExplorer(val *IntegrationExplorer) *NullableIntegrationExplorer {
	return &NullableIntegrationExplorer{value: val, isSet: true}
}

func (v NullableIntegrationExplorer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIntegrationExplorer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

