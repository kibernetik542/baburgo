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

// checks if the GroupedBillingInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupedBillingInfo{}

// GroupedBillingInfo struct for GroupedBillingInfo
type GroupedBillingInfo struct {
	Data []GroupedBillings `json:"data,omitempty"`
	ProjectId *int32 `json:"projectId,omitempty"`
	ProjectName NullableString `json:"projectName,omitempty"`
}

// NewGroupedBillingInfo instantiates a new GroupedBillingInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupedBillingInfo() *GroupedBillingInfo {
	this := GroupedBillingInfo{}
	return &this
}

// NewGroupedBillingInfoWithDefaults instantiates a new GroupedBillingInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupedBillingInfoWithDefaults() *GroupedBillingInfo {
	this := GroupedBillingInfo{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupedBillingInfo) GetData() []GroupedBillings {
	if o == nil {
		var ret []GroupedBillings
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupedBillingInfo) GetDataOk() ([]GroupedBillings, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GroupedBillingInfo) HasData() bool {
	if o != nil && IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []GroupedBillings and assigns it to the Data field.
func (o *GroupedBillingInfo) SetData(v []GroupedBillings) {
	o.Data = v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *GroupedBillingInfo) GetProjectId() int32 {
	if o == nil || IsNil(o.ProjectId) {
		var ret int32
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupedBillingInfo) GetProjectIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *GroupedBillingInfo) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given int32 and assigns it to the ProjectId field.
func (o *GroupedBillingInfo) SetProjectId(v int32) {
	o.ProjectId = &v
}

// GetProjectName returns the ProjectName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupedBillingInfo) GetProjectName() string {
	if o == nil || IsNil(o.ProjectName.Get()) {
		var ret string
		return ret
	}
	return *o.ProjectName.Get()
}

// GetProjectNameOk returns a tuple with the ProjectName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupedBillingInfo) GetProjectNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProjectName.Get(), o.ProjectName.IsSet()
}

// HasProjectName returns a boolean if a field has been set.
func (o *GroupedBillingInfo) HasProjectName() bool {
	if o != nil && o.ProjectName.IsSet() {
		return true
	}

	return false
}

// SetProjectName gets a reference to the given NullableString and assigns it to the ProjectName field.
func (o *GroupedBillingInfo) SetProjectName(v string) {
	o.ProjectName.Set(&v)
}
// SetProjectNameNil sets the value for ProjectName to be an explicit nil
func (o *GroupedBillingInfo) SetProjectNameNil() {
	o.ProjectName.Set(nil)
}

// UnsetProjectName ensures that no value is present for ProjectName, not even an explicit nil
func (o *GroupedBillingInfo) UnsetProjectName() {
	o.ProjectName.Unset()
}

func (o GroupedBillingInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupedBillingInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.ProjectId) {
		toSerialize["projectId"] = o.ProjectId
	}
	if o.ProjectName.IsSet() {
		toSerialize["projectName"] = o.ProjectName.Get()
	}
	return toSerialize, nil
}

type NullableGroupedBillingInfo struct {
	value *GroupedBillingInfo
	isSet bool
}

func (v NullableGroupedBillingInfo) Get() *GroupedBillingInfo {
	return v.value
}

func (v *NullableGroupedBillingInfo) Set(val *GroupedBillingInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupedBillingInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupedBillingInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupedBillingInfo(val *GroupedBillingInfo) *NullableGroupedBillingInfo {
	return &NullableGroupedBillingInfo{value: val, isSet: true}
}

func (v NullableGroupedBillingInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupedBillingInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


