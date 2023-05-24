# BackupCredentialsSearchCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | Pointer to **NullableInt32** |  | [optional] 
**Offset** | Pointer to **NullableInt32** |  | [optional] 
**SearchTerm** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewBackupCredentialsSearchCommand

`func NewBackupCredentialsSearchCommand() *BackupCredentialsSearchCommand`

NewBackupCredentialsSearchCommand instantiates a new BackupCredentialsSearchCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupCredentialsSearchCommandWithDefaults

`func NewBackupCredentialsSearchCommandWithDefaults() *BackupCredentialsSearchCommand`

NewBackupCredentialsSearchCommandWithDefaults instantiates a new BackupCredentialsSearchCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *BackupCredentialsSearchCommand) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *BackupCredentialsSearchCommand) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *BackupCredentialsSearchCommand) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *BackupCredentialsSearchCommand) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### SetLimitNil

`func (o *BackupCredentialsSearchCommand) SetLimitNil(b bool)`

 SetLimitNil sets the value for Limit to be an explicit nil

### UnsetLimit
`func (o *BackupCredentialsSearchCommand) UnsetLimit()`

UnsetLimit ensures that no value is present for Limit, not even an explicit nil
### GetOffset

`func (o *BackupCredentialsSearchCommand) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *BackupCredentialsSearchCommand) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *BackupCredentialsSearchCommand) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *BackupCredentialsSearchCommand) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### SetOffsetNil

`func (o *BackupCredentialsSearchCommand) SetOffsetNil(b bool)`

 SetOffsetNil sets the value for Offset to be an explicit nil

### UnsetOffset
`func (o *BackupCredentialsSearchCommand) UnsetOffset()`

UnsetOffset ensures that no value is present for Offset, not even an explicit nil
### GetSearchTerm

`func (o *BackupCredentialsSearchCommand) GetSearchTerm() string`

GetSearchTerm returns the SearchTerm field if non-nil, zero value otherwise.

### GetSearchTermOk

`func (o *BackupCredentialsSearchCommand) GetSearchTermOk() (*string, bool)`

GetSearchTermOk returns a tuple with the SearchTerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchTerm

`func (o *BackupCredentialsSearchCommand) SetSearchTerm(v string)`

SetSearchTerm sets SearchTerm field to given value.

### HasSearchTerm

`func (o *BackupCredentialsSearchCommand) HasSearchTerm() bool`

HasSearchTerm returns a boolean if a field has been set.

### SetSearchTermNil

`func (o *BackupCredentialsSearchCommand) SetSearchTermNil(b bool)`

 SetSearchTermNil sets the value for SearchTerm to be an explicit nil

### UnsetSearchTerm
`func (o *BackupCredentialsSearchCommand) UnsetSearchTerm()`

UnsetSearchTerm ensures that no value is present for SearchTerm, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


