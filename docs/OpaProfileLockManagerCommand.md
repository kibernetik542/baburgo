# OpaProfileLockManagerCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Mode** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewOpaProfileLockManagerCommand

`func NewOpaProfileLockManagerCommand() *OpaProfileLockManagerCommand`

NewOpaProfileLockManagerCommand instantiates a new OpaProfileLockManagerCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpaProfileLockManagerCommandWithDefaults

`func NewOpaProfileLockManagerCommandWithDefaults() *OpaProfileLockManagerCommand`

NewOpaProfileLockManagerCommandWithDefaults instantiates a new OpaProfileLockManagerCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OpaProfileLockManagerCommand) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OpaProfileLockManagerCommand) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OpaProfileLockManagerCommand) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *OpaProfileLockManagerCommand) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMode

`func (o *OpaProfileLockManagerCommand) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *OpaProfileLockManagerCommand) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *OpaProfileLockManagerCommand) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *OpaProfileLockManagerCommand) HasMode() bool`

HasMode returns a boolean if a field has been set.

### SetModeNil

`func (o *OpaProfileLockManagerCommand) SetModeNil(b bool)`

 SetModeNil sets the value for Mode to be an explicit nil

### UnsetMode
`func (o *OpaProfileLockManagerCommand) UnsetMode()`

UnsetMode ensures that no value is present for Mode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


