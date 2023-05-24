# CreateAzureCloudCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**AzureSubscriptionId** | **string** |  | 
**AzureClientId** | **string** |  | 
**AzureClientSecret** | **string** |  | 
**AzureTenantId** | **string** |  | 
**AzureLocation** | **string** |  | 
**AzCount** | Pointer to **int32** |  | [optional] 
**OrganizationId** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewCreateAzureCloudCommand

`func NewCreateAzureCloudCommand(name string, azureSubscriptionId string, azureClientId string, azureClientSecret string, azureTenantId string, azureLocation string, ) *CreateAzureCloudCommand`

NewCreateAzureCloudCommand instantiates a new CreateAzureCloudCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAzureCloudCommandWithDefaults

`func NewCreateAzureCloudCommandWithDefaults() *CreateAzureCloudCommand`

NewCreateAzureCloudCommandWithDefaults instantiates a new CreateAzureCloudCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateAzureCloudCommand) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateAzureCloudCommand) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateAzureCloudCommand) SetName(v string)`

SetName sets Name field to given value.


### GetAzureSubscriptionId

`func (o *CreateAzureCloudCommand) GetAzureSubscriptionId() string`

GetAzureSubscriptionId returns the AzureSubscriptionId field if non-nil, zero value otherwise.

### GetAzureSubscriptionIdOk

`func (o *CreateAzureCloudCommand) GetAzureSubscriptionIdOk() (*string, bool)`

GetAzureSubscriptionIdOk returns a tuple with the AzureSubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureSubscriptionId

`func (o *CreateAzureCloudCommand) SetAzureSubscriptionId(v string)`

SetAzureSubscriptionId sets AzureSubscriptionId field to given value.


### GetAzureClientId

`func (o *CreateAzureCloudCommand) GetAzureClientId() string`

GetAzureClientId returns the AzureClientId field if non-nil, zero value otherwise.

### GetAzureClientIdOk

`func (o *CreateAzureCloudCommand) GetAzureClientIdOk() (*string, bool)`

GetAzureClientIdOk returns a tuple with the AzureClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureClientId

`func (o *CreateAzureCloudCommand) SetAzureClientId(v string)`

SetAzureClientId sets AzureClientId field to given value.


### GetAzureClientSecret

`func (o *CreateAzureCloudCommand) GetAzureClientSecret() string`

GetAzureClientSecret returns the AzureClientSecret field if non-nil, zero value otherwise.

### GetAzureClientSecretOk

`func (o *CreateAzureCloudCommand) GetAzureClientSecretOk() (*string, bool)`

GetAzureClientSecretOk returns a tuple with the AzureClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureClientSecret

`func (o *CreateAzureCloudCommand) SetAzureClientSecret(v string)`

SetAzureClientSecret sets AzureClientSecret field to given value.


### GetAzureTenantId

`func (o *CreateAzureCloudCommand) GetAzureTenantId() string`

GetAzureTenantId returns the AzureTenantId field if non-nil, zero value otherwise.

### GetAzureTenantIdOk

`func (o *CreateAzureCloudCommand) GetAzureTenantIdOk() (*string, bool)`

GetAzureTenantIdOk returns a tuple with the AzureTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureTenantId

`func (o *CreateAzureCloudCommand) SetAzureTenantId(v string)`

SetAzureTenantId sets AzureTenantId field to given value.


### GetAzureLocation

`func (o *CreateAzureCloudCommand) GetAzureLocation() string`

GetAzureLocation returns the AzureLocation field if non-nil, zero value otherwise.

### GetAzureLocationOk

`func (o *CreateAzureCloudCommand) GetAzureLocationOk() (*string, bool)`

GetAzureLocationOk returns a tuple with the AzureLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureLocation

`func (o *CreateAzureCloudCommand) SetAzureLocation(v string)`

SetAzureLocation sets AzureLocation field to given value.


### GetAzCount

`func (o *CreateAzureCloudCommand) GetAzCount() int32`

GetAzCount returns the AzCount field if non-nil, zero value otherwise.

### GetAzCountOk

`func (o *CreateAzureCloudCommand) GetAzCountOk() (*int32, bool)`

GetAzCountOk returns a tuple with the AzCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzCount

`func (o *CreateAzureCloudCommand) SetAzCount(v int32)`

SetAzCount sets AzCount field to given value.

### HasAzCount

`func (o *CreateAzureCloudCommand) HasAzCount() bool`

HasAzCount returns a boolean if a field has been set.

### GetOrganizationId

`func (o *CreateAzureCloudCommand) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *CreateAzureCloudCommand) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *CreateAzureCloudCommand) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *CreateAzureCloudCommand) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### SetOrganizationIdNil

`func (o *CreateAzureCloudCommand) SetOrganizationIdNil(b bool)`

 SetOrganizationIdNil sets the value for OrganizationId to be an explicit nil

### UnsetOrganizationId
`func (o *CreateAzureCloudCommand) UnsetOrganizationId()`

UnsetOrganizationId ensures that no value is present for OrganizationId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


