# CreateBackupPolicyCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**IncludeNamespaces** | Pointer to **[]string** |  | [optional] 
**CronPeriod** | **string** |  | 
**RetentionPeriod** | **string** |  | 
**ProjectId** | **int32** |  | 

## Methods

### NewCreateBackupPolicyCommand

`func NewCreateBackupPolicyCommand(name string, cronPeriod string, retentionPeriod string, projectId int32, ) *CreateBackupPolicyCommand`

NewCreateBackupPolicyCommand instantiates a new CreateBackupPolicyCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateBackupPolicyCommandWithDefaults

`func NewCreateBackupPolicyCommandWithDefaults() *CreateBackupPolicyCommand`

NewCreateBackupPolicyCommandWithDefaults instantiates a new CreateBackupPolicyCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateBackupPolicyCommand) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateBackupPolicyCommand) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateBackupPolicyCommand) SetName(v string)`

SetName sets Name field to given value.


### GetIncludeNamespaces

`func (o *CreateBackupPolicyCommand) GetIncludeNamespaces() []string`

GetIncludeNamespaces returns the IncludeNamespaces field if non-nil, zero value otherwise.

### GetIncludeNamespacesOk

`func (o *CreateBackupPolicyCommand) GetIncludeNamespacesOk() (*[]string, bool)`

GetIncludeNamespacesOk returns a tuple with the IncludeNamespaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeNamespaces

`func (o *CreateBackupPolicyCommand) SetIncludeNamespaces(v []string)`

SetIncludeNamespaces sets IncludeNamespaces field to given value.

### HasIncludeNamespaces

`func (o *CreateBackupPolicyCommand) HasIncludeNamespaces() bool`

HasIncludeNamespaces returns a boolean if a field has been set.

### SetIncludeNamespacesNil

`func (o *CreateBackupPolicyCommand) SetIncludeNamespacesNil(b bool)`

 SetIncludeNamespacesNil sets the value for IncludeNamespaces to be an explicit nil

### UnsetIncludeNamespaces
`func (o *CreateBackupPolicyCommand) UnsetIncludeNamespaces()`

UnsetIncludeNamespaces ensures that no value is present for IncludeNamespaces, not even an explicit nil
### GetCronPeriod

`func (o *CreateBackupPolicyCommand) GetCronPeriod() string`

GetCronPeriod returns the CronPeriod field if non-nil, zero value otherwise.

### GetCronPeriodOk

`func (o *CreateBackupPolicyCommand) GetCronPeriodOk() (*string, bool)`

GetCronPeriodOk returns a tuple with the CronPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCronPeriod

`func (o *CreateBackupPolicyCommand) SetCronPeriod(v string)`

SetCronPeriod sets CronPeriod field to given value.


### GetRetentionPeriod

`func (o *CreateBackupPolicyCommand) GetRetentionPeriod() string`

GetRetentionPeriod returns the RetentionPeriod field if non-nil, zero value otherwise.

### GetRetentionPeriodOk

`func (o *CreateBackupPolicyCommand) GetRetentionPeriodOk() (*string, bool)`

GetRetentionPeriodOk returns a tuple with the RetentionPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionPeriod

`func (o *CreateBackupPolicyCommand) SetRetentionPeriod(v string)`

SetRetentionPeriod sets RetentionPeriod field to given value.


### GetProjectId

`func (o *CreateBackupPolicyCommand) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *CreateBackupPolicyCommand) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *CreateBackupPolicyCommand) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


