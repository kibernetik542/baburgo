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

// checks if the OpenStackProjectListQuery type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OpenStackProjectListQuery{}

// OpenStackProjectListQuery struct for OpenStackProjectListQuery
type OpenStackProjectListQuery struct {
	OpenStackUser NullableString `json:"openStackUser,omitempty"`
	OpenStackPassword NullableString `json:"openStackPassword,omitempty"`
	OpenStackUrl NullableString `json:"openStackUrl,omitempty"`
	OpenStackDomain NullableString `json:"openStackDomain,omitempty"`
	ApplicationCredEnabled *bool `json:"applicationCredEnabled,omitempty"`
	IsAdmin *bool `json:"isAdmin,omitempty"`
}

// NewOpenStackProjectListQuery instantiates a new OpenStackProjectListQuery object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpenStackProjectListQuery() *OpenStackProjectListQuery {
	this := OpenStackProjectListQuery{}
	return &this
}

// NewOpenStackProjectListQueryWithDefaults instantiates a new OpenStackProjectListQuery object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpenStackProjectListQueryWithDefaults() *OpenStackProjectListQuery {
	this := OpenStackProjectListQuery{}
	return &this
}

// GetOpenStackUser returns the OpenStackUser field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpenStackProjectListQuery) GetOpenStackUser() string {
	if o == nil || IsNil(o.OpenStackUser.Get()) {
		var ret string
		return ret
	}
	return *o.OpenStackUser.Get()
}

// GetOpenStackUserOk returns a tuple with the OpenStackUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpenStackProjectListQuery) GetOpenStackUserOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OpenStackUser.Get(), o.OpenStackUser.IsSet()
}

// HasOpenStackUser returns a boolean if a field has been set.
func (o *OpenStackProjectListQuery) HasOpenStackUser() bool {
	if o != nil && o.OpenStackUser.IsSet() {
		return true
	}

	return false
}

// SetOpenStackUser gets a reference to the given NullableString and assigns it to the OpenStackUser field.
func (o *OpenStackProjectListQuery) SetOpenStackUser(v string) {
	o.OpenStackUser.Set(&v)
}
// SetOpenStackUserNil sets the value for OpenStackUser to be an explicit nil
func (o *OpenStackProjectListQuery) SetOpenStackUserNil() {
	o.OpenStackUser.Set(nil)
}

// UnsetOpenStackUser ensures that no value is present for OpenStackUser, not even an explicit nil
func (o *OpenStackProjectListQuery) UnsetOpenStackUser() {
	o.OpenStackUser.Unset()
}

// GetOpenStackPassword returns the OpenStackPassword field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpenStackProjectListQuery) GetOpenStackPassword() string {
	if o == nil || IsNil(o.OpenStackPassword.Get()) {
		var ret string
		return ret
	}
	return *o.OpenStackPassword.Get()
}

// GetOpenStackPasswordOk returns a tuple with the OpenStackPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpenStackProjectListQuery) GetOpenStackPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OpenStackPassword.Get(), o.OpenStackPassword.IsSet()
}

// HasOpenStackPassword returns a boolean if a field has been set.
func (o *OpenStackProjectListQuery) HasOpenStackPassword() bool {
	if o != nil && o.OpenStackPassword.IsSet() {
		return true
	}

	return false
}

// SetOpenStackPassword gets a reference to the given NullableString and assigns it to the OpenStackPassword field.
func (o *OpenStackProjectListQuery) SetOpenStackPassword(v string) {
	o.OpenStackPassword.Set(&v)
}
// SetOpenStackPasswordNil sets the value for OpenStackPassword to be an explicit nil
func (o *OpenStackProjectListQuery) SetOpenStackPasswordNil() {
	o.OpenStackPassword.Set(nil)
}

// UnsetOpenStackPassword ensures that no value is present for OpenStackPassword, not even an explicit nil
func (o *OpenStackProjectListQuery) UnsetOpenStackPassword() {
	o.OpenStackPassword.Unset()
}

// GetOpenStackUrl returns the OpenStackUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpenStackProjectListQuery) GetOpenStackUrl() string {
	if o == nil || IsNil(o.OpenStackUrl.Get()) {
		var ret string
		return ret
	}
	return *o.OpenStackUrl.Get()
}

// GetOpenStackUrlOk returns a tuple with the OpenStackUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpenStackProjectListQuery) GetOpenStackUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OpenStackUrl.Get(), o.OpenStackUrl.IsSet()
}

// HasOpenStackUrl returns a boolean if a field has been set.
func (o *OpenStackProjectListQuery) HasOpenStackUrl() bool {
	if o != nil && o.OpenStackUrl.IsSet() {
		return true
	}

	return false
}

// SetOpenStackUrl gets a reference to the given NullableString and assigns it to the OpenStackUrl field.
func (o *OpenStackProjectListQuery) SetOpenStackUrl(v string) {
	o.OpenStackUrl.Set(&v)
}
// SetOpenStackUrlNil sets the value for OpenStackUrl to be an explicit nil
func (o *OpenStackProjectListQuery) SetOpenStackUrlNil() {
	o.OpenStackUrl.Set(nil)
}

// UnsetOpenStackUrl ensures that no value is present for OpenStackUrl, not even an explicit nil
func (o *OpenStackProjectListQuery) UnsetOpenStackUrl() {
	o.OpenStackUrl.Unset()
}

// GetOpenStackDomain returns the OpenStackDomain field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpenStackProjectListQuery) GetOpenStackDomain() string {
	if o == nil || IsNil(o.OpenStackDomain.Get()) {
		var ret string
		return ret
	}
	return *o.OpenStackDomain.Get()
}

// GetOpenStackDomainOk returns a tuple with the OpenStackDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpenStackProjectListQuery) GetOpenStackDomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OpenStackDomain.Get(), o.OpenStackDomain.IsSet()
}

// HasOpenStackDomain returns a boolean if a field has been set.
func (o *OpenStackProjectListQuery) HasOpenStackDomain() bool {
	if o != nil && o.OpenStackDomain.IsSet() {
		return true
	}

	return false
}

// SetOpenStackDomain gets a reference to the given NullableString and assigns it to the OpenStackDomain field.
func (o *OpenStackProjectListQuery) SetOpenStackDomain(v string) {
	o.OpenStackDomain.Set(&v)
}
// SetOpenStackDomainNil sets the value for OpenStackDomain to be an explicit nil
func (o *OpenStackProjectListQuery) SetOpenStackDomainNil() {
	o.OpenStackDomain.Set(nil)
}

// UnsetOpenStackDomain ensures that no value is present for OpenStackDomain, not even an explicit nil
func (o *OpenStackProjectListQuery) UnsetOpenStackDomain() {
	o.OpenStackDomain.Unset()
}

// GetApplicationCredEnabled returns the ApplicationCredEnabled field value if set, zero value otherwise.
func (o *OpenStackProjectListQuery) GetApplicationCredEnabled() bool {
	if o == nil || IsNil(o.ApplicationCredEnabled) {
		var ret bool
		return ret
	}
	return *o.ApplicationCredEnabled
}

// GetApplicationCredEnabledOk returns a tuple with the ApplicationCredEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenStackProjectListQuery) GetApplicationCredEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.ApplicationCredEnabled) {
		return nil, false
	}
	return o.ApplicationCredEnabled, true
}

// HasApplicationCredEnabled returns a boolean if a field has been set.
func (o *OpenStackProjectListQuery) HasApplicationCredEnabled() bool {
	if o != nil && !IsNil(o.ApplicationCredEnabled) {
		return true
	}

	return false
}

// SetApplicationCredEnabled gets a reference to the given bool and assigns it to the ApplicationCredEnabled field.
func (o *OpenStackProjectListQuery) SetApplicationCredEnabled(v bool) {
	o.ApplicationCredEnabled = &v
}

// GetIsAdmin returns the IsAdmin field value if set, zero value otherwise.
func (o *OpenStackProjectListQuery) GetIsAdmin() bool {
	if o == nil || IsNil(o.IsAdmin) {
		var ret bool
		return ret
	}
	return *o.IsAdmin
}

// GetIsAdminOk returns a tuple with the IsAdmin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenStackProjectListQuery) GetIsAdminOk() (*bool, bool) {
	if o == nil || IsNil(o.IsAdmin) {
		return nil, false
	}
	return o.IsAdmin, true
}

// HasIsAdmin returns a boolean if a field has been set.
func (o *OpenStackProjectListQuery) HasIsAdmin() bool {
	if o != nil && !IsNil(o.IsAdmin) {
		return true
	}

	return false
}

// SetIsAdmin gets a reference to the given bool and assigns it to the IsAdmin field.
func (o *OpenStackProjectListQuery) SetIsAdmin(v bool) {
	o.IsAdmin = &v
}

func (o OpenStackProjectListQuery) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OpenStackProjectListQuery) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.OpenStackUser.IsSet() {
		toSerialize["openStackUser"] = o.OpenStackUser.Get()
	}
	if o.OpenStackPassword.IsSet() {
		toSerialize["openStackPassword"] = o.OpenStackPassword.Get()
	}
	if o.OpenStackUrl.IsSet() {
		toSerialize["openStackUrl"] = o.OpenStackUrl.Get()
	}
	if o.OpenStackDomain.IsSet() {
		toSerialize["openStackDomain"] = o.OpenStackDomain.Get()
	}
	if !IsNil(o.ApplicationCredEnabled) {
		toSerialize["applicationCredEnabled"] = o.ApplicationCredEnabled
	}
	if !IsNil(o.IsAdmin) {
		toSerialize["isAdmin"] = o.IsAdmin
	}
	return toSerialize, nil
}

type NullableOpenStackProjectListQuery struct {
	value *OpenStackProjectListQuery
	isSet bool
}

func (v NullableOpenStackProjectListQuery) Get() *OpenStackProjectListQuery {
	return v.value
}

func (v *NullableOpenStackProjectListQuery) Set(val *OpenStackProjectListQuery) {
	v.value = val
	v.isSet = true
}

func (v NullableOpenStackProjectListQuery) IsSet() bool {
	return v.isSet
}

func (v *NullableOpenStackProjectListQuery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpenStackProjectListQuery(val *OpenStackProjectListQuery) *NullableOpenStackProjectListQuery {
	return &NullableOpenStackProjectListQuery{value: val, isSet: true}
}

func (v NullableOpenStackProjectListQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpenStackProjectListQuery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


