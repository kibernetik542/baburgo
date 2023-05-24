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

// checks if the AzureDashboardCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AzureDashboardCommand{}

// AzureDashboardCommand struct for AzureDashboardCommand
type AzureDashboardCommand struct {
	CloudId *int32 `json:"cloudId,omitempty"`
	FilterBy NullableString `json:"filterBy,omitempty"`
}

// NewAzureDashboardCommand instantiates a new AzureDashboardCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureDashboardCommand() *AzureDashboardCommand {
	this := AzureDashboardCommand{}
	return &this
}

// NewAzureDashboardCommandWithDefaults instantiates a new AzureDashboardCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureDashboardCommandWithDefaults() *AzureDashboardCommand {
	this := AzureDashboardCommand{}
	return &this
}

// GetCloudId returns the CloudId field value if set, zero value otherwise.
func (o *AzureDashboardCommand) GetCloudId() int32 {
	if o == nil || IsNil(o.CloudId) {
		var ret int32
		return ret
	}
	return *o.CloudId
}

// GetCloudIdOk returns a tuple with the CloudId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureDashboardCommand) GetCloudIdOk() (*int32, bool) {
	if o == nil || IsNil(o.CloudId) {
		return nil, false
	}
	return o.CloudId, true
}

// HasCloudId returns a boolean if a field has been set.
func (o *AzureDashboardCommand) HasCloudId() bool {
	if o != nil && !IsNil(o.CloudId) {
		return true
	}

	return false
}

// SetCloudId gets a reference to the given int32 and assigns it to the CloudId field.
func (o *AzureDashboardCommand) SetCloudId(v int32) {
	o.CloudId = &v
}

// GetFilterBy returns the FilterBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AzureDashboardCommand) GetFilterBy() string {
	if o == nil || IsNil(o.FilterBy.Get()) {
		var ret string
		return ret
	}
	return *o.FilterBy.Get()
}

// GetFilterByOk returns a tuple with the FilterBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AzureDashboardCommand) GetFilterByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FilterBy.Get(), o.FilterBy.IsSet()
}

// HasFilterBy returns a boolean if a field has been set.
func (o *AzureDashboardCommand) HasFilterBy() bool {
	if o != nil && o.FilterBy.IsSet() {
		return true
	}

	return false
}

// SetFilterBy gets a reference to the given NullableString and assigns it to the FilterBy field.
func (o *AzureDashboardCommand) SetFilterBy(v string) {
	o.FilterBy.Set(&v)
}
// SetFilterByNil sets the value for FilterBy to be an explicit nil
func (o *AzureDashboardCommand) SetFilterByNil() {
	o.FilterBy.Set(nil)
}

// UnsetFilterBy ensures that no value is present for FilterBy, not even an explicit nil
func (o *AzureDashboardCommand) UnsetFilterBy() {
	o.FilterBy.Unset()
}

func (o AzureDashboardCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AzureDashboardCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CloudId) {
		toSerialize["cloudId"] = o.CloudId
	}
	if o.FilterBy.IsSet() {
		toSerialize["filterBy"] = o.FilterBy.Get()
	}
	return toSerialize, nil
}

type NullableAzureDashboardCommand struct {
	value *AzureDashboardCommand
	isSet bool
}

func (v NullableAzureDashboardCommand) Get() *AzureDashboardCommand {
	return v.value
}

func (v *NullableAzureDashboardCommand) Set(val *AzureDashboardCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureDashboardCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureDashboardCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureDashboardCommand(val *AzureDashboardCommand) *NullableAzureDashboardCommand {
	return &NullableAzureDashboardCommand{value: val, isSet: true}
}

func (v NullableAzureDashboardCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureDashboardCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


