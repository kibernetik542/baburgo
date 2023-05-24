# DescribeJobCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | **int32** |  | 
**Name** | **string** |  | 
**Namespace** | **string** |  | 

## Methods

### NewDescribeJobCommand

`func NewDescribeJobCommand(projectId int32, name string, namespace string, ) *DescribeJobCommand`

NewDescribeJobCommand instantiates a new DescribeJobCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeJobCommandWithDefaults

`func NewDescribeJobCommandWithDefaults() *DescribeJobCommand`

NewDescribeJobCommandWithDefaults instantiates a new DescribeJobCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *DescribeJobCommand) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *DescribeJobCommand) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *DescribeJobCommand) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.


### GetName

`func (o *DescribeJobCommand) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DescribeJobCommand) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DescribeJobCommand) SetName(v string)`

SetName sets Name field to given value.


### GetNamespace

`func (o *DescribeJobCommand) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *DescribeJobCommand) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *DescribeJobCommand) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


