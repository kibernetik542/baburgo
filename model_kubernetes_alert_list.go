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

// checks if the KubernetesAlertList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KubernetesAlertList{}

// KubernetesAlertList struct for KubernetesAlertList
type KubernetesAlertList struct {
	Data []KubernetesAlertDto `json:"data,omitempty"`
	TotalCount *int32 `json:"totalCount,omitempty"`
}

// NewKubernetesAlertList instantiates a new KubernetesAlertList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesAlertList() *KubernetesAlertList {
	this := KubernetesAlertList{}
	return &this
}

// NewKubernetesAlertListWithDefaults instantiates a new KubernetesAlertList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesAlertListWithDefaults() *KubernetesAlertList {
	this := KubernetesAlertList{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesAlertList) GetData() []KubernetesAlertDto {
	if o == nil {
		var ret []KubernetesAlertDto
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesAlertList) GetDataOk() ([]KubernetesAlertDto, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *KubernetesAlertList) HasData() bool {
	if o != nil && IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []KubernetesAlertDto and assigns it to the Data field.
func (o *KubernetesAlertList) SetData(v []KubernetesAlertDto) {
	o.Data = v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *KubernetesAlertList) GetTotalCount() int32 {
	if o == nil || IsNil(o.TotalCount) {
		var ret int32
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesAlertList) GetTotalCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalCount) {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *KubernetesAlertList) HasTotalCount() bool {
	if o != nil && !IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int32 and assigns it to the TotalCount field.
func (o *KubernetesAlertList) SetTotalCount(v int32) {
	o.TotalCount = &v
}

func (o KubernetesAlertList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KubernetesAlertList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.TotalCount) {
		toSerialize["totalCount"] = o.TotalCount
	}
	return toSerialize, nil
}

type NullableKubernetesAlertList struct {
	value *KubernetesAlertList
	isSet bool
}

func (v NullableKubernetesAlertList) Get() *KubernetesAlertList {
	return v.value
}

func (v *NullableKubernetesAlertList) Set(val *KubernetesAlertList) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesAlertList) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesAlertList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesAlertList(val *KubernetesAlertList) *NullableKubernetesAlertList {
	return &NullableKubernetesAlertList{value: val, isSet: true}
}

func (v NullableKubernetesAlertList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesAlertList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


