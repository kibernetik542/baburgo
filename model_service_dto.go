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

// checks if the ServiceDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceDto{}

// ServiceDto struct for ServiceDto
type ServiceDto struct {
	MetadataName NullableString `json:"metadataName,omitempty"`
	Namespace NullableString `json:"namespace,omitempty"`
	Age NullableString `json:"age,omitempty"`
	Type NullableString `json:"type,omitempty"`
	Ip interface{} `json:"ip,omitempty"`
}

// NewServiceDto instantiates a new ServiceDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceDto() *ServiceDto {
	this := ServiceDto{}
	return &this
}

// NewServiceDtoWithDefaults instantiates a new ServiceDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceDtoWithDefaults() *ServiceDto {
	this := ServiceDto{}
	return &this
}

// GetMetadataName returns the MetadataName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServiceDto) GetMetadataName() string {
	if o == nil || IsNil(o.MetadataName.Get()) {
		var ret string
		return ret
	}
	return *o.MetadataName.Get()
}

// GetMetadataNameOk returns a tuple with the MetadataName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServiceDto) GetMetadataNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MetadataName.Get(), o.MetadataName.IsSet()
}

// HasMetadataName returns a boolean if a field has been set.
func (o *ServiceDto) HasMetadataName() bool {
	if o != nil && o.MetadataName.IsSet() {
		return true
	}

	return false
}

// SetMetadataName gets a reference to the given NullableString and assigns it to the MetadataName field.
func (o *ServiceDto) SetMetadataName(v string) {
	o.MetadataName.Set(&v)
}
// SetMetadataNameNil sets the value for MetadataName to be an explicit nil
func (o *ServiceDto) SetMetadataNameNil() {
	o.MetadataName.Set(nil)
}

// UnsetMetadataName ensures that no value is present for MetadataName, not even an explicit nil
func (o *ServiceDto) UnsetMetadataName() {
	o.MetadataName.Unset()
}

// GetNamespace returns the Namespace field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServiceDto) GetNamespace() string {
	if o == nil || IsNil(o.Namespace.Get()) {
		var ret string
		return ret
	}
	return *o.Namespace.Get()
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServiceDto) GetNamespaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Namespace.Get(), o.Namespace.IsSet()
}

// HasNamespace returns a boolean if a field has been set.
func (o *ServiceDto) HasNamespace() bool {
	if o != nil && o.Namespace.IsSet() {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given NullableString and assigns it to the Namespace field.
func (o *ServiceDto) SetNamespace(v string) {
	o.Namespace.Set(&v)
}
// SetNamespaceNil sets the value for Namespace to be an explicit nil
func (o *ServiceDto) SetNamespaceNil() {
	o.Namespace.Set(nil)
}

// UnsetNamespace ensures that no value is present for Namespace, not even an explicit nil
func (o *ServiceDto) UnsetNamespace() {
	o.Namespace.Unset()
}

// GetAge returns the Age field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServiceDto) GetAge() string {
	if o == nil || IsNil(o.Age.Get()) {
		var ret string
		return ret
	}
	return *o.Age.Get()
}

// GetAgeOk returns a tuple with the Age field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServiceDto) GetAgeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Age.Get(), o.Age.IsSet()
}

// HasAge returns a boolean if a field has been set.
func (o *ServiceDto) HasAge() bool {
	if o != nil && o.Age.IsSet() {
		return true
	}

	return false
}

// SetAge gets a reference to the given NullableString and assigns it to the Age field.
func (o *ServiceDto) SetAge(v string) {
	o.Age.Set(&v)
}
// SetAgeNil sets the value for Age to be an explicit nil
func (o *ServiceDto) SetAgeNil() {
	o.Age.Set(nil)
}

// UnsetAge ensures that no value is present for Age, not even an explicit nil
func (o *ServiceDto) UnsetAge() {
	o.Age.Unset()
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServiceDto) GetType() string {
	if o == nil || IsNil(o.Type.Get()) {
		var ret string
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServiceDto) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *ServiceDto) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableString and assigns it to the Type field.
func (o *ServiceDto) SetType(v string) {
	o.Type.Set(&v)
}
// SetTypeNil sets the value for Type to be an explicit nil
func (o *ServiceDto) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *ServiceDto) UnsetType() {
	o.Type.Unset()
}

// GetIp returns the Ip field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServiceDto) GetIp() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServiceDto) GetIpOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Ip) {
		return nil, false
	}
	return &o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *ServiceDto) HasIp() bool {
	if o != nil && IsNil(o.Ip) {
		return true
	}

	return false
}

// SetIp gets a reference to the given interface{} and assigns it to the Ip field.
func (o *ServiceDto) SetIp(v interface{}) {
	o.Ip = v
}

func (o ServiceDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.MetadataName.IsSet() {
		toSerialize["metadataName"] = o.MetadataName.Get()
	}
	if o.Namespace.IsSet() {
		toSerialize["namespace"] = o.Namespace.Get()
	}
	if o.Age.IsSet() {
		toSerialize["age"] = o.Age.Get()
	}
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
	}
	if o.Ip != nil {
		toSerialize["ip"] = o.Ip
	}
	return toSerialize, nil
}

type NullableServiceDto struct {
	value *ServiceDto
	isSet bool
}

func (v NullableServiceDto) Get() *ServiceDto {
	return v.value
}

func (v *NullableServiceDto) Set(val *ServiceDto) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceDto) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceDto(val *ServiceDto) *NullableServiceDto {
	return &NullableServiceDto{value: val, isSet: true}
}

func (v NullableServiceDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


