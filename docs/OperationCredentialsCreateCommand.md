# OperationCredentialsCreateCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**PrometheusUsername** | **string** |  | 
**PrometheusPassword** | **string** |  | 
**PrometheusUrl** | **string** |  | 
**OrganizationId** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewOperationCredentialsCreateCommand

`func NewOperationCredentialsCreateCommand(name string, prometheusUsername string, prometheusPassword string, prometheusUrl string, ) *OperationCredentialsCreateCommand`

NewOperationCredentialsCreateCommand instantiates a new OperationCredentialsCreateCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOperationCredentialsCreateCommandWithDefaults

`func NewOperationCredentialsCreateCommandWithDefaults() *OperationCredentialsCreateCommand`

NewOperationCredentialsCreateCommandWithDefaults instantiates a new OperationCredentialsCreateCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *OperationCredentialsCreateCommand) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OperationCredentialsCreateCommand) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OperationCredentialsCreateCommand) SetName(v string)`

SetName sets Name field to given value.


### GetPrometheusUsername

`func (o *OperationCredentialsCreateCommand) GetPrometheusUsername() string`

GetPrometheusUsername returns the PrometheusUsername field if non-nil, zero value otherwise.

### GetPrometheusUsernameOk

`func (o *OperationCredentialsCreateCommand) GetPrometheusUsernameOk() (*string, bool)`

GetPrometheusUsernameOk returns a tuple with the PrometheusUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrometheusUsername

`func (o *OperationCredentialsCreateCommand) SetPrometheusUsername(v string)`

SetPrometheusUsername sets PrometheusUsername field to given value.


### GetPrometheusPassword

`func (o *OperationCredentialsCreateCommand) GetPrometheusPassword() string`

GetPrometheusPassword returns the PrometheusPassword field if non-nil, zero value otherwise.

### GetPrometheusPasswordOk

`func (o *OperationCredentialsCreateCommand) GetPrometheusPasswordOk() (*string, bool)`

GetPrometheusPasswordOk returns a tuple with the PrometheusPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrometheusPassword

`func (o *OperationCredentialsCreateCommand) SetPrometheusPassword(v string)`

SetPrometheusPassword sets PrometheusPassword field to given value.


### GetPrometheusUrl

`func (o *OperationCredentialsCreateCommand) GetPrometheusUrl() string`

GetPrometheusUrl returns the PrometheusUrl field if non-nil, zero value otherwise.

### GetPrometheusUrlOk

`func (o *OperationCredentialsCreateCommand) GetPrometheusUrlOk() (*string, bool)`

GetPrometheusUrlOk returns a tuple with the PrometheusUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrometheusUrl

`func (o *OperationCredentialsCreateCommand) SetPrometheusUrl(v string)`

SetPrometheusUrl sets PrometheusUrl field to given value.


### GetOrganizationId

`func (o *OperationCredentialsCreateCommand) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *OperationCredentialsCreateCommand) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *OperationCredentialsCreateCommand) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *OperationCredentialsCreateCommand) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### SetOrganizationIdNil

`func (o *OperationCredentialsCreateCommand) SetOrganizationIdNil(b bool)`

 SetOrganizationIdNil sets the value for OrganizationId to be an explicit nil

### UnsetOrganizationId
`func (o *OperationCredentialsCreateCommand) UnsetOrganizationId()`

UnsetOrganizationId ensures that no value is present for OrganizationId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


