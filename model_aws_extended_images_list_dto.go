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

// checks if the AwsExtendedImagesListDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AwsExtendedImagesListDto{}

// AwsExtendedImagesListDto struct for AwsExtendedImagesListDto
type AwsExtendedImagesListDto struct {
	Id NullableString `json:"id,omitempty"`
	Name NullableString `json:"name,omitempty"`
	Description NullableString `json:"description,omitempty"`
	PlatformDetails NullableString `json:"platformDetails,omitempty"`
	OwnerId NullableString `json:"ownerId,omitempty"`
	OwnerAlias NullableString `json:"ownerAlias,omitempty"`
	Logo NullableString `json:"logo,omitempty"`
}

// NewAwsExtendedImagesListDto instantiates a new AwsExtendedImagesListDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAwsExtendedImagesListDto() *AwsExtendedImagesListDto {
	this := AwsExtendedImagesListDto{}
	return &this
}

// NewAwsExtendedImagesListDtoWithDefaults instantiates a new AwsExtendedImagesListDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAwsExtendedImagesListDtoWithDefaults() *AwsExtendedImagesListDto {
	this := AwsExtendedImagesListDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AwsExtendedImagesListDto) GetId() string {
	if o == nil || IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AwsExtendedImagesListDto) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *AwsExtendedImagesListDto) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *AwsExtendedImagesListDto) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *AwsExtendedImagesListDto) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *AwsExtendedImagesListDto) UnsetId() {
	o.Id.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AwsExtendedImagesListDto) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AwsExtendedImagesListDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *AwsExtendedImagesListDto) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *AwsExtendedImagesListDto) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *AwsExtendedImagesListDto) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *AwsExtendedImagesListDto) UnsetName() {
	o.Name.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AwsExtendedImagesListDto) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AwsExtendedImagesListDto) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *AwsExtendedImagesListDto) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *AwsExtendedImagesListDto) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *AwsExtendedImagesListDto) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *AwsExtendedImagesListDto) UnsetDescription() {
	o.Description.Unset()
}

// GetPlatformDetails returns the PlatformDetails field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AwsExtendedImagesListDto) GetPlatformDetails() string {
	if o == nil || IsNil(o.PlatformDetails.Get()) {
		var ret string
		return ret
	}
	return *o.PlatformDetails.Get()
}

// GetPlatformDetailsOk returns a tuple with the PlatformDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AwsExtendedImagesListDto) GetPlatformDetailsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PlatformDetails.Get(), o.PlatformDetails.IsSet()
}

// HasPlatformDetails returns a boolean if a field has been set.
func (o *AwsExtendedImagesListDto) HasPlatformDetails() bool {
	if o != nil && o.PlatformDetails.IsSet() {
		return true
	}

	return false
}

// SetPlatformDetails gets a reference to the given NullableString and assigns it to the PlatformDetails field.
func (o *AwsExtendedImagesListDto) SetPlatformDetails(v string) {
	o.PlatformDetails.Set(&v)
}
// SetPlatformDetailsNil sets the value for PlatformDetails to be an explicit nil
func (o *AwsExtendedImagesListDto) SetPlatformDetailsNil() {
	o.PlatformDetails.Set(nil)
}

// UnsetPlatformDetails ensures that no value is present for PlatformDetails, not even an explicit nil
func (o *AwsExtendedImagesListDto) UnsetPlatformDetails() {
	o.PlatformDetails.Unset()
}

// GetOwnerId returns the OwnerId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AwsExtendedImagesListDto) GetOwnerId() string {
	if o == nil || IsNil(o.OwnerId.Get()) {
		var ret string
		return ret
	}
	return *o.OwnerId.Get()
}

// GetOwnerIdOk returns a tuple with the OwnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AwsExtendedImagesListDto) GetOwnerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OwnerId.Get(), o.OwnerId.IsSet()
}

// HasOwnerId returns a boolean if a field has been set.
func (o *AwsExtendedImagesListDto) HasOwnerId() bool {
	if o != nil && o.OwnerId.IsSet() {
		return true
	}

	return false
}

// SetOwnerId gets a reference to the given NullableString and assigns it to the OwnerId field.
func (o *AwsExtendedImagesListDto) SetOwnerId(v string) {
	o.OwnerId.Set(&v)
}
// SetOwnerIdNil sets the value for OwnerId to be an explicit nil
func (o *AwsExtendedImagesListDto) SetOwnerIdNil() {
	o.OwnerId.Set(nil)
}

// UnsetOwnerId ensures that no value is present for OwnerId, not even an explicit nil
func (o *AwsExtendedImagesListDto) UnsetOwnerId() {
	o.OwnerId.Unset()
}

// GetOwnerAlias returns the OwnerAlias field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AwsExtendedImagesListDto) GetOwnerAlias() string {
	if o == nil || IsNil(o.OwnerAlias.Get()) {
		var ret string
		return ret
	}
	return *o.OwnerAlias.Get()
}

// GetOwnerAliasOk returns a tuple with the OwnerAlias field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AwsExtendedImagesListDto) GetOwnerAliasOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OwnerAlias.Get(), o.OwnerAlias.IsSet()
}

// HasOwnerAlias returns a boolean if a field has been set.
func (o *AwsExtendedImagesListDto) HasOwnerAlias() bool {
	if o != nil && o.OwnerAlias.IsSet() {
		return true
	}

	return false
}

// SetOwnerAlias gets a reference to the given NullableString and assigns it to the OwnerAlias field.
func (o *AwsExtendedImagesListDto) SetOwnerAlias(v string) {
	o.OwnerAlias.Set(&v)
}
// SetOwnerAliasNil sets the value for OwnerAlias to be an explicit nil
func (o *AwsExtendedImagesListDto) SetOwnerAliasNil() {
	o.OwnerAlias.Set(nil)
}

// UnsetOwnerAlias ensures that no value is present for OwnerAlias, not even an explicit nil
func (o *AwsExtendedImagesListDto) UnsetOwnerAlias() {
	o.OwnerAlias.Unset()
}

// GetLogo returns the Logo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AwsExtendedImagesListDto) GetLogo() string {
	if o == nil || IsNil(o.Logo.Get()) {
		var ret string
		return ret
	}
	return *o.Logo.Get()
}

// GetLogoOk returns a tuple with the Logo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AwsExtendedImagesListDto) GetLogoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Logo.Get(), o.Logo.IsSet()
}

// HasLogo returns a boolean if a field has been set.
func (o *AwsExtendedImagesListDto) HasLogo() bool {
	if o != nil && o.Logo.IsSet() {
		return true
	}

	return false
}

// SetLogo gets a reference to the given NullableString and assigns it to the Logo field.
func (o *AwsExtendedImagesListDto) SetLogo(v string) {
	o.Logo.Set(&v)
}
// SetLogoNil sets the value for Logo to be an explicit nil
func (o *AwsExtendedImagesListDto) SetLogoNil() {
	o.Logo.Set(nil)
}

// UnsetLogo ensures that no value is present for Logo, not even an explicit nil
func (o *AwsExtendedImagesListDto) UnsetLogo() {
	o.Logo.Unset()
}

func (o AwsExtendedImagesListDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AwsExtendedImagesListDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.PlatformDetails.IsSet() {
		toSerialize["platformDetails"] = o.PlatformDetails.Get()
	}
	if o.OwnerId.IsSet() {
		toSerialize["ownerId"] = o.OwnerId.Get()
	}
	if o.OwnerAlias.IsSet() {
		toSerialize["ownerAlias"] = o.OwnerAlias.Get()
	}
	if o.Logo.IsSet() {
		toSerialize["logo"] = o.Logo.Get()
	}
	return toSerialize, nil
}

type NullableAwsExtendedImagesListDto struct {
	value *AwsExtendedImagesListDto
	isSet bool
}

func (v NullableAwsExtendedImagesListDto) Get() *AwsExtendedImagesListDto {
	return v.value
}

func (v *NullableAwsExtendedImagesListDto) Set(val *AwsExtendedImagesListDto) {
	v.value = val
	v.isSet = true
}

func (v NullableAwsExtendedImagesListDto) IsSet() bool {
	return v.isSet
}

func (v *NullableAwsExtendedImagesListDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAwsExtendedImagesListDto(val *AwsExtendedImagesListDto) *NullableAwsExtendedImagesListDto {
	return &NullableAwsExtendedImagesListDto{value: val, isSet: true}
}

func (v NullableAwsExtendedImagesListDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAwsExtendedImagesListDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


