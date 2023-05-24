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

// checks if the TanzuFlavorList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TanzuFlavorList{}

// TanzuFlavorList struct for TanzuFlavorList
type TanzuFlavorList struct {
	Data []TanzuFlavorsListDto `json:"data,omitempty"`
	TotalCount *int32 `json:"totalCount,omitempty"`
}

// NewTanzuFlavorList instantiates a new TanzuFlavorList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTanzuFlavorList() *TanzuFlavorList {
	this := TanzuFlavorList{}
	return &this
}

// NewTanzuFlavorListWithDefaults instantiates a new TanzuFlavorList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTanzuFlavorListWithDefaults() *TanzuFlavorList {
	this := TanzuFlavorList{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TanzuFlavorList) GetData() []TanzuFlavorsListDto {
	if o == nil {
		var ret []TanzuFlavorsListDto
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TanzuFlavorList) GetDataOk() ([]TanzuFlavorsListDto, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *TanzuFlavorList) HasData() bool {
	if o != nil && IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []TanzuFlavorsListDto and assigns it to the Data field.
func (o *TanzuFlavorList) SetData(v []TanzuFlavorsListDto) {
	o.Data = v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *TanzuFlavorList) GetTotalCount() int32 {
	if o == nil || IsNil(o.TotalCount) {
		var ret int32
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TanzuFlavorList) GetTotalCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalCount) {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *TanzuFlavorList) HasTotalCount() bool {
	if o != nil && !IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int32 and assigns it to the TotalCount field.
func (o *TanzuFlavorList) SetTotalCount(v int32) {
	o.TotalCount = &v
}

func (o TanzuFlavorList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TanzuFlavorList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.TotalCount) {
		toSerialize["totalCount"] = o.TotalCount
	}
	return toSerialize, nil
}

type NullableTanzuFlavorList struct {
	value *TanzuFlavorList
	isSet bool
}

func (v NullableTanzuFlavorList) Get() *TanzuFlavorList {
	return v.value
}

func (v *NullableTanzuFlavorList) Set(val *TanzuFlavorList) {
	v.value = val
	v.isSet = true
}

func (v NullableTanzuFlavorList) IsSet() bool {
	return v.isSet
}

func (v *NullableTanzuFlavorList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTanzuFlavorList(val *TanzuFlavorList) *NullableTanzuFlavorList {
	return &NullableTanzuFlavorList{value: val, isSet: true}
}

func (v NullableTanzuFlavorList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTanzuFlavorList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


