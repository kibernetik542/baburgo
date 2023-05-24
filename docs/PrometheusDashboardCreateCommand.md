# PrometheusDashboardCreateCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | Pointer to **int32** |  | [optional] 
**Name** | **string** |  | 
**Expression** | **string** |  | 
**Description** | **string** |  | 
**CategoryName** | **string** |  | 

## Methods

### NewPrometheusDashboardCreateCommand

`func NewPrometheusDashboardCreateCommand(name string, expression string, description string, categoryName string, ) *PrometheusDashboardCreateCommand`

NewPrometheusDashboardCreateCommand instantiates a new PrometheusDashboardCreateCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrometheusDashboardCreateCommandWithDefaults

`func NewPrometheusDashboardCreateCommandWithDefaults() *PrometheusDashboardCreateCommand`

NewPrometheusDashboardCreateCommandWithDefaults instantiates a new PrometheusDashboardCreateCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *PrometheusDashboardCreateCommand) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *PrometheusDashboardCreateCommand) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *PrometheusDashboardCreateCommand) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *PrometheusDashboardCreateCommand) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetName

`func (o *PrometheusDashboardCreateCommand) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PrometheusDashboardCreateCommand) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PrometheusDashboardCreateCommand) SetName(v string)`

SetName sets Name field to given value.


### GetExpression

`func (o *PrometheusDashboardCreateCommand) GetExpression() string`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *PrometheusDashboardCreateCommand) GetExpressionOk() (*string, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *PrometheusDashboardCreateCommand) SetExpression(v string)`

SetExpression sets Expression field to given value.


### GetDescription

`func (o *PrometheusDashboardCreateCommand) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PrometheusDashboardCreateCommand) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PrometheusDashboardCreateCommand) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetCategoryName

`func (o *PrometheusDashboardCreateCommand) GetCategoryName() string`

GetCategoryName returns the CategoryName field if non-nil, zero value otherwise.

### GetCategoryNameOk

`func (o *PrometheusDashboardCreateCommand) GetCategoryNameOk() (*string, bool)`

GetCategoryNameOk returns a tuple with the CategoryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryName

`func (o *PrometheusDashboardCreateCommand) SetCategoryName(v string)`

SetCategoryName sets CategoryName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


