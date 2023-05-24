# ProjectAppParamsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** |  | 
**Value** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewProjectAppParamsDto

`func NewProjectAppParamsDto(key string, ) *ProjectAppParamsDto`

NewProjectAppParamsDto instantiates a new ProjectAppParamsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectAppParamsDtoWithDefaults

`func NewProjectAppParamsDtoWithDefaults() *ProjectAppParamsDto`

NewProjectAppParamsDtoWithDefaults instantiates a new ProjectAppParamsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *ProjectAppParamsDto) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ProjectAppParamsDto) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ProjectAppParamsDto) SetKey(v string)`

SetKey sets Key field to given value.


### GetValue

`func (o *ProjectAppParamsDto) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ProjectAppParamsDto) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ProjectAppParamsDto) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ProjectAppParamsDto) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *ProjectAppParamsDto) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *ProjectAppParamsDto) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


