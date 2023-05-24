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

// checks if the CreateNtpServerCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateNtpServerCommand{}

// CreateNtpServerCommand struct for CreateNtpServerCommand
type CreateNtpServerCommand struct {
	Address string `json:"address"`
	AccessProfileId int32 `json:"accessProfileId"`
}

// NewCreateNtpServerCommand instantiates a new CreateNtpServerCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNtpServerCommand(address string, accessProfileId int32) *CreateNtpServerCommand {
	this := CreateNtpServerCommand{}
	this.Address = address
	this.AccessProfileId = accessProfileId
	return &this
}

// NewCreateNtpServerCommandWithDefaults instantiates a new CreateNtpServerCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNtpServerCommandWithDefaults() *CreateNtpServerCommand {
	this := CreateNtpServerCommand{}
	return &this
}

// GetAddress returns the Address field value
func (o *CreateNtpServerCommand) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *CreateNtpServerCommand) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *CreateNtpServerCommand) SetAddress(v string) {
	o.Address = v
}

// GetAccessProfileId returns the AccessProfileId field value
func (o *CreateNtpServerCommand) GetAccessProfileId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AccessProfileId
}

// GetAccessProfileIdOk returns a tuple with the AccessProfileId field value
// and a boolean to check if the value has been set.
func (o *CreateNtpServerCommand) GetAccessProfileIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessProfileId, true
}

// SetAccessProfileId sets field value
func (o *CreateNtpServerCommand) SetAccessProfileId(v int32) {
	o.AccessProfileId = v
}

func (o CreateNtpServerCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateNtpServerCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["address"] = o.Address
	toSerialize["accessProfileId"] = o.AccessProfileId
	return toSerialize, nil
}

type NullableCreateNtpServerCommand struct {
	value *CreateNtpServerCommand
	isSet bool
}

func (v NullableCreateNtpServerCommand) Get() *CreateNtpServerCommand {
	return v.value
}

func (v *NullableCreateNtpServerCommand) Set(val *CreateNtpServerCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNtpServerCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNtpServerCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNtpServerCommand(val *CreateNtpServerCommand) *NullableCreateNtpServerCommand {
	return &NullableCreateNtpServerCommand{value: val, isSet: true}
}

func (v NullableCreateNtpServerCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNtpServerCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


