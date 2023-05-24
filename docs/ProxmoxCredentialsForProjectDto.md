# ProxmoxCredentialsForProjectDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | Pointer to **NullableString** |  | [optional] 
**Url** | Pointer to **NullableString** |  | [optional] 
**Password** | Pointer to **NullableString** |  | [optional] 
**Storage** | Pointer to **NullableString** |  | [optional] 
**VmTemplateName** | Pointer to **NullableString** |  | [optional] 
**ProxmoxNetworks** | Pointer to [**[]ProxmoxNetworkListDto**](ProxmoxNetworkListDto.md) |  | [optional] 

## Methods

### NewProxmoxCredentialsForProjectDto

`func NewProxmoxCredentialsForProjectDto() *ProxmoxCredentialsForProjectDto`

NewProxmoxCredentialsForProjectDto instantiates a new ProxmoxCredentialsForProjectDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProxmoxCredentialsForProjectDtoWithDefaults

`func NewProxmoxCredentialsForProjectDtoWithDefaults() *ProxmoxCredentialsForProjectDto`

NewProxmoxCredentialsForProjectDtoWithDefaults instantiates a new ProxmoxCredentialsForProjectDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *ProxmoxCredentialsForProjectDto) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ProxmoxCredentialsForProjectDto) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ProxmoxCredentialsForProjectDto) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ProxmoxCredentialsForProjectDto) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *ProxmoxCredentialsForProjectDto) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *ProxmoxCredentialsForProjectDto) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetUrl

`func (o *ProxmoxCredentialsForProjectDto) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ProxmoxCredentialsForProjectDto) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ProxmoxCredentialsForProjectDto) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ProxmoxCredentialsForProjectDto) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *ProxmoxCredentialsForProjectDto) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *ProxmoxCredentialsForProjectDto) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetPassword

`func (o *ProxmoxCredentialsForProjectDto) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ProxmoxCredentialsForProjectDto) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ProxmoxCredentialsForProjectDto) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *ProxmoxCredentialsForProjectDto) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *ProxmoxCredentialsForProjectDto) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *ProxmoxCredentialsForProjectDto) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetStorage

`func (o *ProxmoxCredentialsForProjectDto) GetStorage() string`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *ProxmoxCredentialsForProjectDto) GetStorageOk() (*string, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *ProxmoxCredentialsForProjectDto) SetStorage(v string)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *ProxmoxCredentialsForProjectDto) HasStorage() bool`

HasStorage returns a boolean if a field has been set.

### SetStorageNil

`func (o *ProxmoxCredentialsForProjectDto) SetStorageNil(b bool)`

 SetStorageNil sets the value for Storage to be an explicit nil

### UnsetStorage
`func (o *ProxmoxCredentialsForProjectDto) UnsetStorage()`

UnsetStorage ensures that no value is present for Storage, not even an explicit nil
### GetVmTemplateName

`func (o *ProxmoxCredentialsForProjectDto) GetVmTemplateName() string`

GetVmTemplateName returns the VmTemplateName field if non-nil, zero value otherwise.

### GetVmTemplateNameOk

`func (o *ProxmoxCredentialsForProjectDto) GetVmTemplateNameOk() (*string, bool)`

GetVmTemplateNameOk returns a tuple with the VmTemplateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmTemplateName

`func (o *ProxmoxCredentialsForProjectDto) SetVmTemplateName(v string)`

SetVmTemplateName sets VmTemplateName field to given value.

### HasVmTemplateName

`func (o *ProxmoxCredentialsForProjectDto) HasVmTemplateName() bool`

HasVmTemplateName returns a boolean if a field has been set.

### SetVmTemplateNameNil

`func (o *ProxmoxCredentialsForProjectDto) SetVmTemplateNameNil(b bool)`

 SetVmTemplateNameNil sets the value for VmTemplateName to be an explicit nil

### UnsetVmTemplateName
`func (o *ProxmoxCredentialsForProjectDto) UnsetVmTemplateName()`

UnsetVmTemplateName ensures that no value is present for VmTemplateName, not even an explicit nil
### GetProxmoxNetworks

`func (o *ProxmoxCredentialsForProjectDto) GetProxmoxNetworks() []ProxmoxNetworkListDto`

GetProxmoxNetworks returns the ProxmoxNetworks field if non-nil, zero value otherwise.

### GetProxmoxNetworksOk

`func (o *ProxmoxCredentialsForProjectDto) GetProxmoxNetworksOk() (*[]ProxmoxNetworkListDto, bool)`

GetProxmoxNetworksOk returns a tuple with the ProxmoxNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxmoxNetworks

`func (o *ProxmoxCredentialsForProjectDto) SetProxmoxNetworks(v []ProxmoxNetworkListDto)`

SetProxmoxNetworks sets ProxmoxNetworks field to given value.

### HasProxmoxNetworks

`func (o *ProxmoxCredentialsForProjectDto) HasProxmoxNetworks() bool`

HasProxmoxNetworks returns a boolean if a field has been set.

### SetProxmoxNetworksNil

`func (o *ProxmoxCredentialsForProjectDto) SetProxmoxNetworksNil(b bool)`

 SetProxmoxNetworksNil sets the value for ProxmoxNetworks to be an explicit nil

### UnsetProxmoxNetworks
`func (o *ProxmoxCredentialsForProjectDto) UnsetProxmoxNetworks()`

UnsetProxmoxNetworks ensures that no value is present for ProxmoxNetworks, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


