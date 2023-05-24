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

// checks if the AdminOrganizationsList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdminOrganizationsList{}

// AdminOrganizationsList struct for AdminOrganizationsList
type AdminOrganizationsList struct {
	Data []AdminOrganizationsListDto `json:"data,omitempty"`
	TotalCount *int32 `json:"totalCount,omitempty"`
}

// NewAdminOrganizationsList instantiates a new AdminOrganizationsList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdminOrganizationsList() *AdminOrganizationsList {
	this := AdminOrganizationsList{}
	return &this
}

// NewAdminOrganizationsListWithDefaults instantiates a new AdminOrganizationsList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdminOrganizationsListWithDefaults() *AdminOrganizationsList {
	this := AdminOrganizationsList{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AdminOrganizationsList) GetData() []AdminOrganizationsListDto {
	if o == nil {
		var ret []AdminOrganizationsListDto
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AdminOrganizationsList) GetDataOk() ([]AdminOrganizationsListDto, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AdminOrganizationsList) HasData() bool {
	if o != nil && IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []AdminOrganizationsListDto and assigns it to the Data field.
func (o *AdminOrganizationsList) SetData(v []AdminOrganizationsListDto) {
	o.Data = v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *AdminOrganizationsList) GetTotalCount() int32 {
	if o == nil || IsNil(o.TotalCount) {
		var ret int32
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdminOrganizationsList) GetTotalCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalCount) {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *AdminOrganizationsList) HasTotalCount() bool {
	if o != nil && !IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int32 and assigns it to the TotalCount field.
func (o *AdminOrganizationsList) SetTotalCount(v int32) {
	o.TotalCount = &v
}

func (o AdminOrganizationsList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdminOrganizationsList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.TotalCount) {
		toSerialize["totalCount"] = o.TotalCount
	}
	return toSerialize, nil
}

type NullableAdminOrganizationsList struct {
	value *AdminOrganizationsList
	isSet bool
}

func (v NullableAdminOrganizationsList) Get() *AdminOrganizationsList {
	return v.value
}

func (v *NullableAdminOrganizationsList) Set(val *AdminOrganizationsList) {
	v.value = val
	v.isSet = true
}

func (v NullableAdminOrganizationsList) IsSet() bool {
	return v.isSet
}

func (v *NullableAdminOrganizationsList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdminOrganizationsList(val *AdminOrganizationsList) *NullableAdminOrganizationsList {
	return &NullableAdminOrganizationsList{value: val, isSet: true}
}

func (v NullableAdminOrganizationsList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdminOrganizationsList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


