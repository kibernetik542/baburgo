# UpdateOrganizationCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | **string** |  | 
**FullName** | **string** |  | 
**Phone** | Pointer to **NullableString** |  | [optional] 
**Email** | Pointer to **NullableString** |  | [optional] 
**Address** | Pointer to **NullableString** |  | [optional] 
**Country** | Pointer to **NullableString** |  | [optional] 
**City** | Pointer to **NullableString** |  | [optional] 
**VatNumber** | Pointer to **NullableString** |  | [optional] 
**IsLocked** | Pointer to **bool** |  | [optional] 
**BillingEmail** | Pointer to **NullableString** |  | [optional] 
**DiscountRate** | Pointer to **NullableFloat64** |  | [optional] 
**IsEligibleUpdateSubscription** | Pointer to **bool** |  | [optional] 

## Methods

### NewUpdateOrganizationCommand

`func NewUpdateOrganizationCommand(name string, fullName string, ) *UpdateOrganizationCommand`

NewUpdateOrganizationCommand instantiates a new UpdateOrganizationCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateOrganizationCommandWithDefaults

`func NewUpdateOrganizationCommandWithDefaults() *UpdateOrganizationCommand`

NewUpdateOrganizationCommandWithDefaults instantiates a new UpdateOrganizationCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateOrganizationCommand) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateOrganizationCommand) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateOrganizationCommand) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateOrganizationCommand) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *UpdateOrganizationCommand) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateOrganizationCommand) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateOrganizationCommand) SetName(v string)`

SetName sets Name field to given value.


### GetFullName

`func (o *UpdateOrganizationCommand) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *UpdateOrganizationCommand) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *UpdateOrganizationCommand) SetFullName(v string)`

SetFullName sets FullName field to given value.


### GetPhone

`func (o *UpdateOrganizationCommand) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *UpdateOrganizationCommand) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *UpdateOrganizationCommand) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *UpdateOrganizationCommand) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### SetPhoneNil

`func (o *UpdateOrganizationCommand) SetPhoneNil(b bool)`

 SetPhoneNil sets the value for Phone to be an explicit nil

### UnsetPhone
`func (o *UpdateOrganizationCommand) UnsetPhone()`

UnsetPhone ensures that no value is present for Phone, not even an explicit nil
### GetEmail

`func (o *UpdateOrganizationCommand) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UpdateOrganizationCommand) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UpdateOrganizationCommand) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UpdateOrganizationCommand) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *UpdateOrganizationCommand) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *UpdateOrganizationCommand) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetAddress

`func (o *UpdateOrganizationCommand) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *UpdateOrganizationCommand) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *UpdateOrganizationCommand) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *UpdateOrganizationCommand) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### SetAddressNil

`func (o *UpdateOrganizationCommand) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *UpdateOrganizationCommand) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil
### GetCountry

`func (o *UpdateOrganizationCommand) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *UpdateOrganizationCommand) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *UpdateOrganizationCommand) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *UpdateOrganizationCommand) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### SetCountryNil

`func (o *UpdateOrganizationCommand) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *UpdateOrganizationCommand) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil
### GetCity

`func (o *UpdateOrganizationCommand) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *UpdateOrganizationCommand) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *UpdateOrganizationCommand) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *UpdateOrganizationCommand) HasCity() bool`

HasCity returns a boolean if a field has been set.

### SetCityNil

`func (o *UpdateOrganizationCommand) SetCityNil(b bool)`

 SetCityNil sets the value for City to be an explicit nil

### UnsetCity
`func (o *UpdateOrganizationCommand) UnsetCity()`

UnsetCity ensures that no value is present for City, not even an explicit nil
### GetVatNumber

`func (o *UpdateOrganizationCommand) GetVatNumber() string`

GetVatNumber returns the VatNumber field if non-nil, zero value otherwise.

### GetVatNumberOk

`func (o *UpdateOrganizationCommand) GetVatNumberOk() (*string, bool)`

GetVatNumberOk returns a tuple with the VatNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatNumber

`func (o *UpdateOrganizationCommand) SetVatNumber(v string)`

SetVatNumber sets VatNumber field to given value.

### HasVatNumber

`func (o *UpdateOrganizationCommand) HasVatNumber() bool`

HasVatNumber returns a boolean if a field has been set.

### SetVatNumberNil

`func (o *UpdateOrganizationCommand) SetVatNumberNil(b bool)`

 SetVatNumberNil sets the value for VatNumber to be an explicit nil

### UnsetVatNumber
`func (o *UpdateOrganizationCommand) UnsetVatNumber()`

UnsetVatNumber ensures that no value is present for VatNumber, not even an explicit nil
### GetIsLocked

`func (o *UpdateOrganizationCommand) GetIsLocked() bool`

GetIsLocked returns the IsLocked field if non-nil, zero value otherwise.

### GetIsLockedOk

`func (o *UpdateOrganizationCommand) GetIsLockedOk() (*bool, bool)`

GetIsLockedOk returns a tuple with the IsLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocked

`func (o *UpdateOrganizationCommand) SetIsLocked(v bool)`

SetIsLocked sets IsLocked field to given value.

### HasIsLocked

`func (o *UpdateOrganizationCommand) HasIsLocked() bool`

HasIsLocked returns a boolean if a field has been set.

### GetBillingEmail

`func (o *UpdateOrganizationCommand) GetBillingEmail() string`

GetBillingEmail returns the BillingEmail field if non-nil, zero value otherwise.

### GetBillingEmailOk

`func (o *UpdateOrganizationCommand) GetBillingEmailOk() (*string, bool)`

GetBillingEmailOk returns a tuple with the BillingEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingEmail

`func (o *UpdateOrganizationCommand) SetBillingEmail(v string)`

SetBillingEmail sets BillingEmail field to given value.

### HasBillingEmail

`func (o *UpdateOrganizationCommand) HasBillingEmail() bool`

HasBillingEmail returns a boolean if a field has been set.

### SetBillingEmailNil

`func (o *UpdateOrganizationCommand) SetBillingEmailNil(b bool)`

 SetBillingEmailNil sets the value for BillingEmail to be an explicit nil

### UnsetBillingEmail
`func (o *UpdateOrganizationCommand) UnsetBillingEmail()`

UnsetBillingEmail ensures that no value is present for BillingEmail, not even an explicit nil
### GetDiscountRate

`func (o *UpdateOrganizationCommand) GetDiscountRate() float64`

GetDiscountRate returns the DiscountRate field if non-nil, zero value otherwise.

### GetDiscountRateOk

`func (o *UpdateOrganizationCommand) GetDiscountRateOk() (*float64, bool)`

GetDiscountRateOk returns a tuple with the DiscountRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountRate

`func (o *UpdateOrganizationCommand) SetDiscountRate(v float64)`

SetDiscountRate sets DiscountRate field to given value.

### HasDiscountRate

`func (o *UpdateOrganizationCommand) HasDiscountRate() bool`

HasDiscountRate returns a boolean if a field has been set.

### SetDiscountRateNil

`func (o *UpdateOrganizationCommand) SetDiscountRateNil(b bool)`

 SetDiscountRateNil sets the value for DiscountRate to be an explicit nil

### UnsetDiscountRate
`func (o *UpdateOrganizationCommand) UnsetDiscountRate()`

UnsetDiscountRate ensures that no value is present for DiscountRate, not even an explicit nil
### GetIsEligibleUpdateSubscription

`func (o *UpdateOrganizationCommand) GetIsEligibleUpdateSubscription() bool`

GetIsEligibleUpdateSubscription returns the IsEligibleUpdateSubscription field if non-nil, zero value otherwise.

### GetIsEligibleUpdateSubscriptionOk

`func (o *UpdateOrganizationCommand) GetIsEligibleUpdateSubscriptionOk() (*bool, bool)`

GetIsEligibleUpdateSubscriptionOk returns a tuple with the IsEligibleUpdateSubscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEligibleUpdateSubscription

`func (o *UpdateOrganizationCommand) SetIsEligibleUpdateSubscription(v bool)`

SetIsEligibleUpdateSubscription sets IsEligibleUpdateSubscription field to given value.

### HasIsEligibleUpdateSubscription

`func (o *UpdateOrganizationCommand) HasIsEligibleUpdateSubscription() bool`

HasIsEligibleUpdateSubscription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


