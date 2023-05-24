# PatchNodeLabelsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** |  | 
**Value** | Pointer to **NullableString** |  | [optional] 
**Mode** | **string** |  | 

## Methods

### NewPatchNodeLabelsDto

`func NewPatchNodeLabelsDto(key string, mode string, ) *PatchNodeLabelsDto`

NewPatchNodeLabelsDto instantiates a new PatchNodeLabelsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchNodeLabelsDtoWithDefaults

`func NewPatchNodeLabelsDtoWithDefaults() *PatchNodeLabelsDto`

NewPatchNodeLabelsDtoWithDefaults instantiates a new PatchNodeLabelsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *PatchNodeLabelsDto) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *PatchNodeLabelsDto) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *PatchNodeLabelsDto) SetKey(v string)`

SetKey sets Key field to given value.


### GetValue

`func (o *PatchNodeLabelsDto) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *PatchNodeLabelsDto) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *PatchNodeLabelsDto) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *PatchNodeLabelsDto) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *PatchNodeLabelsDto) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *PatchNodeLabelsDto) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetMode

`func (o *PatchNodeLabelsDto) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *PatchNodeLabelsDto) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *PatchNodeLabelsDto) SetMode(v string)`

SetMode sets Mode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


