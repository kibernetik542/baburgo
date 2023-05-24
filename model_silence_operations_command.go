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

// checks if the SilenceOperationsCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SilenceOperationsCommand{}

// SilenceOperationsCommand struct for SilenceOperationsCommand
type SilenceOperationsCommand struct {
	Id int32 `json:"id"`
	Mode string `json:"mode"`
	Reason NullableString `json:"reason,omitempty"`
}

// NewSilenceOperationsCommand instantiates a new SilenceOperationsCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSilenceOperationsCommand(id int32, mode string) *SilenceOperationsCommand {
	this := SilenceOperationsCommand{}
	this.Id = id
	this.Mode = mode
	return &this
}

// NewSilenceOperationsCommandWithDefaults instantiates a new SilenceOperationsCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSilenceOperationsCommandWithDefaults() *SilenceOperationsCommand {
	this := SilenceOperationsCommand{}
	return &this
}

// GetId returns the Id field value
func (o *SilenceOperationsCommand) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SilenceOperationsCommand) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SilenceOperationsCommand) SetId(v int32) {
	o.Id = v
}

// GetMode returns the Mode field value
func (o *SilenceOperationsCommand) GetMode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Mode
}

// GetModeOk returns a tuple with the Mode field value
// and a boolean to check if the value has been set.
func (o *SilenceOperationsCommand) GetModeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mode, true
}

// SetMode sets field value
func (o *SilenceOperationsCommand) SetMode(v string) {
	o.Mode = v
}

// GetReason returns the Reason field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SilenceOperationsCommand) GetReason() string {
	if o == nil || IsNil(o.Reason.Get()) {
		var ret string
		return ret
	}
	return *o.Reason.Get()
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SilenceOperationsCommand) GetReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Reason.Get(), o.Reason.IsSet()
}

// HasReason returns a boolean if a field has been set.
func (o *SilenceOperationsCommand) HasReason() bool {
	if o != nil && o.Reason.IsSet() {
		return true
	}

	return false
}

// SetReason gets a reference to the given NullableString and assigns it to the Reason field.
func (o *SilenceOperationsCommand) SetReason(v string) {
	o.Reason.Set(&v)
}
// SetReasonNil sets the value for Reason to be an explicit nil
func (o *SilenceOperationsCommand) SetReasonNil() {
	o.Reason.Set(nil)
}

// UnsetReason ensures that no value is present for Reason, not even an explicit nil
func (o *SilenceOperationsCommand) UnsetReason() {
	o.Reason.Unset()
}

func (o SilenceOperationsCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SilenceOperationsCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["mode"] = o.Mode
	if o.Reason.IsSet() {
		toSerialize["reason"] = o.Reason.Get()
	}
	return toSerialize, nil
}

type NullableSilenceOperationsCommand struct {
	value *SilenceOperationsCommand
	isSet bool
}

func (v NullableSilenceOperationsCommand) Get() *SilenceOperationsCommand {
	return v.value
}

func (v *NullableSilenceOperationsCommand) Set(val *SilenceOperationsCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableSilenceOperationsCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableSilenceOperationsCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSilenceOperationsCommand(val *SilenceOperationsCommand) *NullableSilenceOperationsCommand {
	return &NullableSilenceOperationsCommand{value: val, isSet: true}
}

func (v NullableSilenceOperationsCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSilenceOperationsCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


