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

// checks if the WhiteListDomainCreateCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WhiteListDomainCreateCommand{}

// WhiteListDomainCreateCommand struct for WhiteListDomainCreateCommand
type WhiteListDomainCreateCommand struct {
	WhiteListDomains []WhiteListDomainCreateDto `json:"whiteListDomains,omitempty"`
	PartnerId NullableInt32 `json:"partnerId,omitempty"`
}

// NewWhiteListDomainCreateCommand instantiates a new WhiteListDomainCreateCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWhiteListDomainCreateCommand() *WhiteListDomainCreateCommand {
	this := WhiteListDomainCreateCommand{}
	return &this
}

// NewWhiteListDomainCreateCommandWithDefaults instantiates a new WhiteListDomainCreateCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWhiteListDomainCreateCommandWithDefaults() *WhiteListDomainCreateCommand {
	this := WhiteListDomainCreateCommand{}
	return &this
}

// GetWhiteListDomains returns the WhiteListDomains field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WhiteListDomainCreateCommand) GetWhiteListDomains() []WhiteListDomainCreateDto {
	if o == nil {
		var ret []WhiteListDomainCreateDto
		return ret
	}
	return o.WhiteListDomains
}

// GetWhiteListDomainsOk returns a tuple with the WhiteListDomains field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WhiteListDomainCreateCommand) GetWhiteListDomainsOk() ([]WhiteListDomainCreateDto, bool) {
	if o == nil || IsNil(o.WhiteListDomains) {
		return nil, false
	}
	return o.WhiteListDomains, true
}

// HasWhiteListDomains returns a boolean if a field has been set.
func (o *WhiteListDomainCreateCommand) HasWhiteListDomains() bool {
	if o != nil && IsNil(o.WhiteListDomains) {
		return true
	}

	return false
}

// SetWhiteListDomains gets a reference to the given []WhiteListDomainCreateDto and assigns it to the WhiteListDomains field.
func (o *WhiteListDomainCreateCommand) SetWhiteListDomains(v []WhiteListDomainCreateDto) {
	o.WhiteListDomains = v
}

// GetPartnerId returns the PartnerId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WhiteListDomainCreateCommand) GetPartnerId() int32 {
	if o == nil || IsNil(o.PartnerId.Get()) {
		var ret int32
		return ret
	}
	return *o.PartnerId.Get()
}

// GetPartnerIdOk returns a tuple with the PartnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WhiteListDomainCreateCommand) GetPartnerIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PartnerId.Get(), o.PartnerId.IsSet()
}

// HasPartnerId returns a boolean if a field has been set.
func (o *WhiteListDomainCreateCommand) HasPartnerId() bool {
	if o != nil && o.PartnerId.IsSet() {
		return true
	}

	return false
}

// SetPartnerId gets a reference to the given NullableInt32 and assigns it to the PartnerId field.
func (o *WhiteListDomainCreateCommand) SetPartnerId(v int32) {
	o.PartnerId.Set(&v)
}
// SetPartnerIdNil sets the value for PartnerId to be an explicit nil
func (o *WhiteListDomainCreateCommand) SetPartnerIdNil() {
	o.PartnerId.Set(nil)
}

// UnsetPartnerId ensures that no value is present for PartnerId, not even an explicit nil
func (o *WhiteListDomainCreateCommand) UnsetPartnerId() {
	o.PartnerId.Unset()
}

func (o WhiteListDomainCreateCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WhiteListDomainCreateCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.WhiteListDomains != nil {
		toSerialize["whiteListDomains"] = o.WhiteListDomains
	}
	if o.PartnerId.IsSet() {
		toSerialize["partnerId"] = o.PartnerId.Get()
	}
	return toSerialize, nil
}

type NullableWhiteListDomainCreateCommand struct {
	value *WhiteListDomainCreateCommand
	isSet bool
}

func (v NullableWhiteListDomainCreateCommand) Get() *WhiteListDomainCreateCommand {
	return v.value
}

func (v *NullableWhiteListDomainCreateCommand) Set(val *WhiteListDomainCreateCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableWhiteListDomainCreateCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableWhiteListDomainCreateCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWhiteListDomainCreateCommand(val *WhiteListDomainCreateCommand) *NullableWhiteListDomainCreateCommand {
	return &NullableWhiteListDomainCreateCommand{value: val, isSet: true}
}

func (v NullableWhiteListDomainCreateCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWhiteListDomainCreateCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


