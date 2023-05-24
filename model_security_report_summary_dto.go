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

// checks if the SecurityReportSummaryDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SecurityReportSummaryDto{}

// SecurityReportSummaryDto struct for SecurityReportSummaryDto
type SecurityReportSummaryDto struct {
	Low *int64 `json:"low,omitempty"`
	High *int64 `json:"high,omitempty"`
	Medium *int64 `json:"medium,omitempty"`
	Unknown *int64 `json:"unknown,omitempty"`
	Critical *int64 `json:"critical,omitempty"`
}

// NewSecurityReportSummaryDto instantiates a new SecurityReportSummaryDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityReportSummaryDto() *SecurityReportSummaryDto {
	this := SecurityReportSummaryDto{}
	return &this
}

// NewSecurityReportSummaryDtoWithDefaults instantiates a new SecurityReportSummaryDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityReportSummaryDtoWithDefaults() *SecurityReportSummaryDto {
	this := SecurityReportSummaryDto{}
	return &this
}

// GetLow returns the Low field value if set, zero value otherwise.
func (o *SecurityReportSummaryDto) GetLow() int64 {
	if o == nil || IsNil(o.Low) {
		var ret int64
		return ret
	}
	return *o.Low
}

// GetLowOk returns a tuple with the Low field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityReportSummaryDto) GetLowOk() (*int64, bool) {
	if o == nil || IsNil(o.Low) {
		return nil, false
	}
	return o.Low, true
}

// HasLow returns a boolean if a field has been set.
func (o *SecurityReportSummaryDto) HasLow() bool {
	if o != nil && !IsNil(o.Low) {
		return true
	}

	return false
}

// SetLow gets a reference to the given int64 and assigns it to the Low field.
func (o *SecurityReportSummaryDto) SetLow(v int64) {
	o.Low = &v
}

// GetHigh returns the High field value if set, zero value otherwise.
func (o *SecurityReportSummaryDto) GetHigh() int64 {
	if o == nil || IsNil(o.High) {
		var ret int64
		return ret
	}
	return *o.High
}

// GetHighOk returns a tuple with the High field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityReportSummaryDto) GetHighOk() (*int64, bool) {
	if o == nil || IsNil(o.High) {
		return nil, false
	}
	return o.High, true
}

// HasHigh returns a boolean if a field has been set.
func (o *SecurityReportSummaryDto) HasHigh() bool {
	if o != nil && !IsNil(o.High) {
		return true
	}

	return false
}

// SetHigh gets a reference to the given int64 and assigns it to the High field.
func (o *SecurityReportSummaryDto) SetHigh(v int64) {
	o.High = &v
}

// GetMedium returns the Medium field value if set, zero value otherwise.
func (o *SecurityReportSummaryDto) GetMedium() int64 {
	if o == nil || IsNil(o.Medium) {
		var ret int64
		return ret
	}
	return *o.Medium
}

// GetMediumOk returns a tuple with the Medium field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityReportSummaryDto) GetMediumOk() (*int64, bool) {
	if o == nil || IsNil(o.Medium) {
		return nil, false
	}
	return o.Medium, true
}

// HasMedium returns a boolean if a field has been set.
func (o *SecurityReportSummaryDto) HasMedium() bool {
	if o != nil && !IsNil(o.Medium) {
		return true
	}

	return false
}

// SetMedium gets a reference to the given int64 and assigns it to the Medium field.
func (o *SecurityReportSummaryDto) SetMedium(v int64) {
	o.Medium = &v
}

// GetUnknown returns the Unknown field value if set, zero value otherwise.
func (o *SecurityReportSummaryDto) GetUnknown() int64 {
	if o == nil || IsNil(o.Unknown) {
		var ret int64
		return ret
	}
	return *o.Unknown
}

// GetUnknownOk returns a tuple with the Unknown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityReportSummaryDto) GetUnknownOk() (*int64, bool) {
	if o == nil || IsNil(o.Unknown) {
		return nil, false
	}
	return o.Unknown, true
}

// HasUnknown returns a boolean if a field has been set.
func (o *SecurityReportSummaryDto) HasUnknown() bool {
	if o != nil && !IsNil(o.Unknown) {
		return true
	}

	return false
}

// SetUnknown gets a reference to the given int64 and assigns it to the Unknown field.
func (o *SecurityReportSummaryDto) SetUnknown(v int64) {
	o.Unknown = &v
}

// GetCritical returns the Critical field value if set, zero value otherwise.
func (o *SecurityReportSummaryDto) GetCritical() int64 {
	if o == nil || IsNil(o.Critical) {
		var ret int64
		return ret
	}
	return *o.Critical
}

// GetCriticalOk returns a tuple with the Critical field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityReportSummaryDto) GetCriticalOk() (*int64, bool) {
	if o == nil || IsNil(o.Critical) {
		return nil, false
	}
	return o.Critical, true
}

// HasCritical returns a boolean if a field has been set.
func (o *SecurityReportSummaryDto) HasCritical() bool {
	if o != nil && !IsNil(o.Critical) {
		return true
	}

	return false
}

// SetCritical gets a reference to the given int64 and assigns it to the Critical field.
func (o *SecurityReportSummaryDto) SetCritical(v int64) {
	o.Critical = &v
}

func (o SecurityReportSummaryDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SecurityReportSummaryDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Low) {
		toSerialize["low"] = o.Low
	}
	if !IsNil(o.High) {
		toSerialize["high"] = o.High
	}
	if !IsNil(o.Medium) {
		toSerialize["medium"] = o.Medium
	}
	if !IsNil(o.Unknown) {
		toSerialize["unknown"] = o.Unknown
	}
	if !IsNil(o.Critical) {
		toSerialize["critical"] = o.Critical
	}
	return toSerialize, nil
}

type NullableSecurityReportSummaryDto struct {
	value *SecurityReportSummaryDto
	isSet bool
}

func (v NullableSecurityReportSummaryDto) Get() *SecurityReportSummaryDto {
	return v.value
}

func (v *NullableSecurityReportSummaryDto) Set(val *SecurityReportSummaryDto) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityReportSummaryDto) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityReportSummaryDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityReportSummaryDto(val *SecurityReportSummaryDto) *NullableSecurityReportSummaryDto {
	return &NullableSecurityReportSummaryDto{value: val, isSet: true}
}

func (v NullableSecurityReportSummaryDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityReportSummaryDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


