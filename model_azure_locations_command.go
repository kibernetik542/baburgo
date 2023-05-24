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

// checks if the AzureLocationsCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AzureLocationsCommand{}

// AzureLocationsCommand struct for AzureLocationsCommand
type AzureLocationsCommand struct {
	AzureSubscriptionId NullableString `json:"azureSubscriptionId,omitempty"`
	AzureClientId NullableString `json:"azureClientId,omitempty"`
	AzureClientSecret NullableString `json:"azureClientSecret,omitempty"`
	AzureTenantId NullableString `json:"azureTenantId,omitempty"`
}

// NewAzureLocationsCommand instantiates a new AzureLocationsCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureLocationsCommand() *AzureLocationsCommand {
	this := AzureLocationsCommand{}
	return &this
}

// NewAzureLocationsCommandWithDefaults instantiates a new AzureLocationsCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureLocationsCommandWithDefaults() *AzureLocationsCommand {
	this := AzureLocationsCommand{}
	return &this
}

// GetAzureSubscriptionId returns the AzureSubscriptionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AzureLocationsCommand) GetAzureSubscriptionId() string {
	if o == nil || IsNil(o.AzureSubscriptionId.Get()) {
		var ret string
		return ret
	}
	return *o.AzureSubscriptionId.Get()
}

// GetAzureSubscriptionIdOk returns a tuple with the AzureSubscriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AzureLocationsCommand) GetAzureSubscriptionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AzureSubscriptionId.Get(), o.AzureSubscriptionId.IsSet()
}

// HasAzureSubscriptionId returns a boolean if a field has been set.
func (o *AzureLocationsCommand) HasAzureSubscriptionId() bool {
	if o != nil && o.AzureSubscriptionId.IsSet() {
		return true
	}

	return false
}

// SetAzureSubscriptionId gets a reference to the given NullableString and assigns it to the AzureSubscriptionId field.
func (o *AzureLocationsCommand) SetAzureSubscriptionId(v string) {
	o.AzureSubscriptionId.Set(&v)
}
// SetAzureSubscriptionIdNil sets the value for AzureSubscriptionId to be an explicit nil
func (o *AzureLocationsCommand) SetAzureSubscriptionIdNil() {
	o.AzureSubscriptionId.Set(nil)
}

// UnsetAzureSubscriptionId ensures that no value is present for AzureSubscriptionId, not even an explicit nil
func (o *AzureLocationsCommand) UnsetAzureSubscriptionId() {
	o.AzureSubscriptionId.Unset()
}

// GetAzureClientId returns the AzureClientId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AzureLocationsCommand) GetAzureClientId() string {
	if o == nil || IsNil(o.AzureClientId.Get()) {
		var ret string
		return ret
	}
	return *o.AzureClientId.Get()
}

// GetAzureClientIdOk returns a tuple with the AzureClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AzureLocationsCommand) GetAzureClientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AzureClientId.Get(), o.AzureClientId.IsSet()
}

// HasAzureClientId returns a boolean if a field has been set.
func (o *AzureLocationsCommand) HasAzureClientId() bool {
	if o != nil && o.AzureClientId.IsSet() {
		return true
	}

	return false
}

// SetAzureClientId gets a reference to the given NullableString and assigns it to the AzureClientId field.
func (o *AzureLocationsCommand) SetAzureClientId(v string) {
	o.AzureClientId.Set(&v)
}
// SetAzureClientIdNil sets the value for AzureClientId to be an explicit nil
func (o *AzureLocationsCommand) SetAzureClientIdNil() {
	o.AzureClientId.Set(nil)
}

// UnsetAzureClientId ensures that no value is present for AzureClientId, not even an explicit nil
func (o *AzureLocationsCommand) UnsetAzureClientId() {
	o.AzureClientId.Unset()
}

// GetAzureClientSecret returns the AzureClientSecret field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AzureLocationsCommand) GetAzureClientSecret() string {
	if o == nil || IsNil(o.AzureClientSecret.Get()) {
		var ret string
		return ret
	}
	return *o.AzureClientSecret.Get()
}

// GetAzureClientSecretOk returns a tuple with the AzureClientSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AzureLocationsCommand) GetAzureClientSecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AzureClientSecret.Get(), o.AzureClientSecret.IsSet()
}

// HasAzureClientSecret returns a boolean if a field has been set.
func (o *AzureLocationsCommand) HasAzureClientSecret() bool {
	if o != nil && o.AzureClientSecret.IsSet() {
		return true
	}

	return false
}

// SetAzureClientSecret gets a reference to the given NullableString and assigns it to the AzureClientSecret field.
func (o *AzureLocationsCommand) SetAzureClientSecret(v string) {
	o.AzureClientSecret.Set(&v)
}
// SetAzureClientSecretNil sets the value for AzureClientSecret to be an explicit nil
func (o *AzureLocationsCommand) SetAzureClientSecretNil() {
	o.AzureClientSecret.Set(nil)
}

// UnsetAzureClientSecret ensures that no value is present for AzureClientSecret, not even an explicit nil
func (o *AzureLocationsCommand) UnsetAzureClientSecret() {
	o.AzureClientSecret.Unset()
}

// GetAzureTenantId returns the AzureTenantId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AzureLocationsCommand) GetAzureTenantId() string {
	if o == nil || IsNil(o.AzureTenantId.Get()) {
		var ret string
		return ret
	}
	return *o.AzureTenantId.Get()
}

// GetAzureTenantIdOk returns a tuple with the AzureTenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AzureLocationsCommand) GetAzureTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AzureTenantId.Get(), o.AzureTenantId.IsSet()
}

// HasAzureTenantId returns a boolean if a field has been set.
func (o *AzureLocationsCommand) HasAzureTenantId() bool {
	if o != nil && o.AzureTenantId.IsSet() {
		return true
	}

	return false
}

// SetAzureTenantId gets a reference to the given NullableString and assigns it to the AzureTenantId field.
func (o *AzureLocationsCommand) SetAzureTenantId(v string) {
	o.AzureTenantId.Set(&v)
}
// SetAzureTenantIdNil sets the value for AzureTenantId to be an explicit nil
func (o *AzureLocationsCommand) SetAzureTenantIdNil() {
	o.AzureTenantId.Set(nil)
}

// UnsetAzureTenantId ensures that no value is present for AzureTenantId, not even an explicit nil
func (o *AzureLocationsCommand) UnsetAzureTenantId() {
	o.AzureTenantId.Unset()
}

func (o AzureLocationsCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AzureLocationsCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AzureSubscriptionId.IsSet() {
		toSerialize["azureSubscriptionId"] = o.AzureSubscriptionId.Get()
	}
	if o.AzureClientId.IsSet() {
		toSerialize["azureClientId"] = o.AzureClientId.Get()
	}
	if o.AzureClientSecret.IsSet() {
		toSerialize["azureClientSecret"] = o.AzureClientSecret.Get()
	}
	if o.AzureTenantId.IsSet() {
		toSerialize["azureTenantId"] = o.AzureTenantId.Get()
	}
	return toSerialize, nil
}

type NullableAzureLocationsCommand struct {
	value *AzureLocationsCommand
	isSet bool
}

func (v NullableAzureLocationsCommand) Get() *AzureLocationsCommand {
	return v.value
}

func (v *NullableAzureLocationsCommand) Set(val *AzureLocationsCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureLocationsCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureLocationsCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureLocationsCommand(val *AzureLocationsCommand) *NullableAzureLocationsCommand {
	return &NullableAzureLocationsCommand{value: val, isSet: true}
}

func (v NullableAzureLocationsCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureLocationsCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


