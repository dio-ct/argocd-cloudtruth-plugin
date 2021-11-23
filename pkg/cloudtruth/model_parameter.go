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

// Parameter A single parameter inside of a project.
type Parameter struct {
	Url string `json:"url"`
	// A unique identifier for the parameter.
	Id string `json:"id"`
	// The parameter name.
	Name string `json:"name"`
	// A description of the parameter.  You may find it helpful to document how this parameter is used to assist others when they need to maintain software that uses this content.
	Description *string `json:"description,omitempty"`
	// Indicates if this content is secret or not.  When a parameter is considered to be a secret, any internal values are stored in a dedicated vault for your organization for maximum security.  External values are inspected on-demand to ensure they align with the parameter's secret setting and if they do not, those external values are not allowed to be used.
	Secret *bool `json:"secret,omitempty"`
	Type *ParameterTypeEnum `json:"type,omitempty"`
	// Rules applied to this parameter.
	Rules []ParameterRule `json:"rules"`
	// The project that the parameter is within.
	Project string `json:"project"`
	// The project name that the parameter is within
	ProjectName string `json:"project_name"`
	// Templates that reference this Parameter.
	ReferencingTemplates []string `json:"referencing_templates"`
	// Dynamic values that reference this Parameter.
	ReferencingValues []string `json:"referencing_values"`
	//              Each parameter has an effective value in every environment based on             environment inheritance and which environments have had a value set.              Environments inherit from a single parent to form a tree, as a result             a single parameter may have different values present for each environment.             When a value is not explicitly set in an environment, the parent environment             is consulted to see if it has a value defined, and so on.              The dictionary of values has an environment url as the key, and the optional             Value record that it resolves to.  If the Value.environment matches the key,             then it is an explicit value set for that environment.  If they differ, the             value was obtained from a parent environment (directly or indirectly).  If the             value is None then no value has ever been set in any environment for this             parameter.              key: Environment url             value: optional Value record         
	Values map[string]Value `json:"values"`
	CreatedAt time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
}

// NewParameter instantiates a new Parameter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewParameter(url string, id string, name string, rules []ParameterRule, project string, projectName string, referencingTemplates []string, referencingValues []string, values map[string]Value, createdAt time.Time, modifiedAt time.Time) *Parameter {
	this := Parameter{}
	this.Url = url
	this.Id = id
	this.Name = name
	this.Rules = rules
	this.Project = project
	this.ProjectName = projectName
	this.ReferencingTemplates = referencingTemplates
	this.ReferencingValues = referencingValues
	this.Values = values
	this.CreatedAt = createdAt
	this.ModifiedAt = modifiedAt
	return &this
}

// NewParameterWithDefaults instantiates a new Parameter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewParameterWithDefaults() *Parameter {
	this := Parameter{}
	return &this
}

// GetUrl returns the Url field value
func (o *Parameter) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *Parameter) GetUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *Parameter) SetUrl(v string) {
	o.Url = v
}

// GetId returns the Id field value
func (o *Parameter) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Parameter) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Parameter) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *Parameter) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Parameter) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Parameter) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Parameter) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Parameter) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Parameter) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Parameter) SetDescription(v string) {
	o.Description = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *Parameter) GetSecret() bool {
	if o == nil || o.Secret == nil {
		var ret bool
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Parameter) GetSecretOk() (*bool, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *Parameter) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given bool and assigns it to the Secret field.
func (o *Parameter) SetSecret(v bool) {
	o.Secret = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Parameter) GetType() ParameterTypeEnum {
	if o == nil || o.Type == nil {
		var ret ParameterTypeEnum
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Parameter) GetTypeOk() (*ParameterTypeEnum, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Parameter) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given ParameterTypeEnum and assigns it to the Type field.
func (o *Parameter) SetType(v ParameterTypeEnum) {
	o.Type = &v
}

// GetRules returns the Rules field value
func (o *Parameter) GetRules() []ParameterRule {
	if o == nil {
		var ret []ParameterRule
		return ret
	}

	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value
// and a boolean to check if the value has been set.
func (o *Parameter) GetRulesOk() (*[]ParameterRule, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Rules, true
}

// SetRules sets field value
func (o *Parameter) SetRules(v []ParameterRule) {
	o.Rules = v
}

// GetProject returns the Project field value
func (o *Parameter) GetProject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Project
}

// GetProjectOk returns a tuple with the Project field value
// and a boolean to check if the value has been set.
func (o *Parameter) GetProjectOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Project, true
}

// SetProject sets field value
func (o *Parameter) SetProject(v string) {
	o.Project = v
}

// GetProjectName returns the ProjectName field value
func (o *Parameter) GetProjectName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectName
}

// GetProjectNameOk returns a tuple with the ProjectName field value
// and a boolean to check if the value has been set.
func (o *Parameter) GetProjectNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ProjectName, true
}

// SetProjectName sets field value
func (o *Parameter) SetProjectName(v string) {
	o.ProjectName = v
}

// GetReferencingTemplates returns the ReferencingTemplates field value
func (o *Parameter) GetReferencingTemplates() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ReferencingTemplates
}

// GetReferencingTemplatesOk returns a tuple with the ReferencingTemplates field value
// and a boolean to check if the value has been set.
func (o *Parameter) GetReferencingTemplatesOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ReferencingTemplates, true
}

// SetReferencingTemplates sets field value
func (o *Parameter) SetReferencingTemplates(v []string) {
	o.ReferencingTemplates = v
}

// GetReferencingValues returns the ReferencingValues field value
func (o *Parameter) GetReferencingValues() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ReferencingValues
}

// GetReferencingValuesOk returns a tuple with the ReferencingValues field value
// and a boolean to check if the value has been set.
func (o *Parameter) GetReferencingValuesOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ReferencingValues, true
}

// SetReferencingValues sets field value
func (o *Parameter) SetReferencingValues(v []string) {
	o.ReferencingValues = v
}

// GetValues returns the Values field value
func (o *Parameter) GetValues() map[string]Value {
	if o == nil {
		var ret map[string]Value
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *Parameter) GetValuesOk() (*map[string]Value, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Values, true
}

// SetValues sets field value
func (o *Parameter) SetValues(v map[string]Value) {
	o.Values = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *Parameter) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Parameter) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *Parameter) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetModifiedAt returns the ModifiedAt field value
func (o *Parameter) GetModifiedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value
// and a boolean to check if the value has been set.
func (o *Parameter) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ModifiedAt, true
}

// SetModifiedAt sets field value
func (o *Parameter) SetModifiedAt(v time.Time) {
	o.ModifiedAt = v
}

func (o Parameter) MarshalJSON() ([]byte, error) {
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
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["rules"] = o.Rules
	}
	if true {
		toSerialize["project"] = o.Project
	}
	if true {
		toSerialize["project_name"] = o.ProjectName
	}
	if true {
		toSerialize["referencing_templates"] = o.ReferencingTemplates
	}
	if true {
		toSerialize["referencing_values"] = o.ReferencingValues
	}
	if true {
		toSerialize["values"] = o.Values
	}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if true {
		toSerialize["modified_at"] = o.ModifiedAt
	}
	return json.Marshal(toSerialize)
}

type NullableParameter struct {
	value *Parameter
	isSet bool
}

func (v NullableParameter) Get() *Parameter {
	return v.value
}

func (v *NullableParameter) Set(val *Parameter) {
	v.value = val
	v.isSet = true
}

func (v NullableParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableParameter(val *Parameter) *NullableParameter {
	return &NullableParameter{value: val, isSet: true}
}

func (v NullableParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


