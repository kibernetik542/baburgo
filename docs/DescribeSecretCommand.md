# DescribeSecretCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | **int32** |  | 
**Name** | **string** |  | 
**Namespace** | **string** |  | 

## Methods

### NewDescribeSecretCommand

`func NewDescribeSecretCommand(projectId int32, name string, namespace string, ) *DescribeSecretCommand`

NewDescribeSecretCommand instantiates a new DescribeSecretCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeSecretCommandWithDefaults

`func NewDescribeSecretCommandWithDefaults() *DescribeSecretCommand`

NewDescribeSecretCommandWithDefaults instantiates a new DescribeSecretCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *DescribeSecretCommand) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *DescribeSecretCommand) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *DescribeSecretCommand) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.


### GetName

`func (o *DescribeSecretCommand) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DescribeSecretCommand) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DescribeSecretCommand) SetName(v string)`

SetName sets Name field to given value.


### GetNamespace

`func (o *DescribeSecretCommand) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *DescribeSecretCommand) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *DescribeSecretCommand) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


