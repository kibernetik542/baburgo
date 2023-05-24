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

// checks if the ImportBackupStorageLocationCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImportBackupStorageLocationCommand{}

// ImportBackupStorageLocationCommand struct for ImportBackupStorageLocationCommand
type ImportBackupStorageLocationCommand struct {
	TargetProjectId *int32 `json:"targetProjectId,omitempty"`
	SourceProjectId *int32 `json:"sourceProjectId,omitempty"`
}

// NewImportBackupStorageLocationCommand instantiates a new ImportBackupStorageLocationCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImportBackupStorageLocationCommand() *ImportBackupStorageLocationCommand {
	this := ImportBackupStorageLocationCommand{}
	return &this
}

// NewImportBackupStorageLocationCommandWithDefaults instantiates a new ImportBackupStorageLocationCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImportBackupStorageLocationCommandWithDefaults() *ImportBackupStorageLocationCommand {
	this := ImportBackupStorageLocationCommand{}
	return &this
}

// GetTargetProjectId returns the TargetProjectId field value if set, zero value otherwise.
func (o *ImportBackupStorageLocationCommand) GetTargetProjectId() int32 {
	if o == nil || IsNil(o.TargetProjectId) {
		var ret int32
		return ret
	}
	return *o.TargetProjectId
}

// GetTargetProjectIdOk returns a tuple with the TargetProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportBackupStorageLocationCommand) GetTargetProjectIdOk() (*int32, bool) {
	if o == nil || IsNil(o.TargetProjectId) {
		return nil, false
	}
	return o.TargetProjectId, true
}

// HasTargetProjectId returns a boolean if a field has been set.
func (o *ImportBackupStorageLocationCommand) HasTargetProjectId() bool {
	if o != nil && !IsNil(o.TargetProjectId) {
		return true
	}

	return false
}

// SetTargetProjectId gets a reference to the given int32 and assigns it to the TargetProjectId field.
func (o *ImportBackupStorageLocationCommand) SetTargetProjectId(v int32) {
	o.TargetProjectId = &v
}

// GetSourceProjectId returns the SourceProjectId field value if set, zero value otherwise.
func (o *ImportBackupStorageLocationCommand) GetSourceProjectId() int32 {
	if o == nil || IsNil(o.SourceProjectId) {
		var ret int32
		return ret
	}
	return *o.SourceProjectId
}

// GetSourceProjectIdOk returns a tuple with the SourceProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportBackupStorageLocationCommand) GetSourceProjectIdOk() (*int32, bool) {
	if o == nil || IsNil(o.SourceProjectId) {
		return nil, false
	}
	return o.SourceProjectId, true
}

// HasSourceProjectId returns a boolean if a field has been set.
func (o *ImportBackupStorageLocationCommand) HasSourceProjectId() bool {
	if o != nil && !IsNil(o.SourceProjectId) {
		return true
	}

	return false
}

// SetSourceProjectId gets a reference to the given int32 and assigns it to the SourceProjectId field.
func (o *ImportBackupStorageLocationCommand) SetSourceProjectId(v int32) {
	o.SourceProjectId = &v
}

func (o ImportBackupStorageLocationCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ImportBackupStorageLocationCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TargetProjectId) {
		toSerialize["targetProjectId"] = o.TargetProjectId
	}
	if !IsNil(o.SourceProjectId) {
		toSerialize["sourceProjectId"] = o.SourceProjectId
	}
	return toSerialize, nil
}

type NullableImportBackupStorageLocationCommand struct {
	value *ImportBackupStorageLocationCommand
	isSet bool
}

func (v NullableImportBackupStorageLocationCommand) Get() *ImportBackupStorageLocationCommand {
	return v.value
}

func (v *NullableImportBackupStorageLocationCommand) Set(val *ImportBackupStorageLocationCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableImportBackupStorageLocationCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableImportBackupStorageLocationCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImportBackupStorageLocationCommand(val *ImportBackupStorageLocationCommand) *NullableImportBackupStorageLocationCommand {
	return &NullableImportBackupStorageLocationCommand{value: val, isSet: true}
}

func (v NullableImportBackupStorageLocationCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImportBackupStorageLocationCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


