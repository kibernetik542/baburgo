/*
Taikun - WebApi

This Api will be responsible for overall data distribution and authorization.

API version: v1
Contact: noreply@itera.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the CreateProjectGroupCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateProjectGroupCommand{}

// CreateProjectGroupCommand struct for CreateProjectGroupCommand
type CreateProjectGroupCommand struct {
	Name string `json:"name"`
	OrganizationId NullableInt32 `json:"organizationId,omitempty"`
	ProjectIds []int32 `json:"projectIds,omitempty"`
}

// NewCreateProjectGroupCommand instantiates a new CreateProjectGroupCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateProjectGroupCommand(name string) *CreateProjectGroupCommand {
	this := CreateProjectGroupCommand{}
	this.Name = name
	return &this
}

// NewCreateProjectGroupCommandWithDefaults instantiates a new CreateProjectGroupCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateProjectGroupCommandWithDefaults() *CreateProjectGroupCommand {
	this := CreateProjectGroupCommand{}
	return &this
}

// GetName returns the Name field value
func (o *CreateProjectGroupCommand) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateProjectGroupCommand) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateProjectGroupCommand) SetName(v string) {
	o.Name = v
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateProjectGroupCommand) GetOrganizationId() int32 {
	if o == nil || IsNil(o.OrganizationId.Get()) {
		var ret int32
		return ret
	}
	return *o.OrganizationId.Get()
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateProjectGroupCommand) GetOrganizationIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.OrganizationId.Get(), o.OrganizationId.IsSet()
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *CreateProjectGroupCommand) HasOrganizationId() bool {
	if o != nil && o.OrganizationId.IsSet() {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given NullableInt32 and assigns it to the OrganizationId field.
func (o *CreateProjectGroupCommand) SetOrganizationId(v int32) {
	o.OrganizationId.Set(&v)
}
// SetOrganizationIdNil sets the value for OrganizationId to be an explicit nil
func (o *CreateProjectGroupCommand) SetOrganizationIdNil() {
	o.OrganizationId.Set(nil)
}

// UnsetOrganizationId ensures that no value is present for OrganizationId, not even an explicit nil
func (o *CreateProjectGroupCommand) UnsetOrganizationId() {
	o.OrganizationId.Unset()
}

// GetProjectIds returns the ProjectIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateProjectGroupCommand) GetProjectIds() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}
	return o.ProjectIds
}

// GetProjectIdsOk returns a tuple with the ProjectIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateProjectGroupCommand) GetProjectIdsOk() ([]int32, bool) {
	if o == nil || IsNil(o.ProjectIds) {
		return nil, false
	}
	return o.ProjectIds, true
}

// HasProjectIds returns a boolean if a field has been set.
func (o *CreateProjectGroupCommand) HasProjectIds() bool {
	if o != nil && IsNil(o.ProjectIds) {
		return true
	}

	return false
}

// SetProjectIds gets a reference to the given []int32 and assigns it to the ProjectIds field.
func (o *CreateProjectGroupCommand) SetProjectIds(v []int32) {
	o.ProjectIds = v
}

func (o CreateProjectGroupCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateProjectGroupCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if o.OrganizationId.IsSet() {
		toSerialize["organizationId"] = o.OrganizationId.Get()
	}
	if o.ProjectIds != nil {
		toSerialize["projectIds"] = o.ProjectIds
	}
	return toSerialize, nil
}

type NullableCreateProjectGroupCommand struct {
	value *CreateProjectGroupCommand
	isSet bool
}

func (v NullableCreateProjectGroupCommand) Get() *CreateProjectGroupCommand {
	return v.value
}

func (v *NullableCreateProjectGroupCommand) Set(val *CreateProjectGroupCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateProjectGroupCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateProjectGroupCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateProjectGroupCommand(val *CreateProjectGroupCommand) *NullableCreateProjectGroupCommand {
	return &NullableCreateProjectGroupCommand{value: val, isSet: true}
}

func (v NullableCreateProjectGroupCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateProjectGroupCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


