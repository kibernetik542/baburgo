# AwsFlavorListDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ram** | Pointer to **int64** |  | [optional] 
**Cpu** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **interface{}** |  | [optional] 
**LinuxPrice** | Pointer to **NullableString** |  | [optional] 
**WindowsPrice** | Pointer to **NullableString** |  | [optional] 
**WindowsSpotPrice** | Pointer to **NullableString** |  | [optional] 
**LinuxSpotPrice** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewAwsFlavorListDto

`func NewAwsFlavorListDto() *AwsFlavorListDto`

NewAwsFlavorListDto instantiates a new AwsFlavorListDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsFlavorListDtoWithDefaults

`func NewAwsFlavorListDtoWithDefaults() *AwsFlavorListDto`

NewAwsFlavorListDtoWithDefaults instantiates a new AwsFlavorListDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRam

`func (o *AwsFlavorListDto) GetRam() int64`

GetRam returns the Ram field if non-nil, zero value otherwise.

### GetRamOk

`func (o *AwsFlavorListDto) GetRamOk() (*int64, bool)`

GetRamOk returns a tuple with the Ram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRam

`func (o *AwsFlavorListDto) SetRam(v int64)`

SetRam sets Ram field to given value.

### HasRam

`func (o *AwsFlavorListDto) HasRam() bool`

HasRam returns a boolean if a field has been set.

### GetCpu

`func (o *AwsFlavorListDto) GetCpu() int32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *AwsFlavorListDto) GetCpuOk() (*int32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *AwsFlavorListDto) SetCpu(v int32)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *AwsFlavorListDto) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetName

`func (o *AwsFlavorListDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AwsFlavorListDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AwsFlavorListDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AwsFlavorListDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *AwsFlavorListDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AwsFlavorListDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *AwsFlavorListDto) GetDescription() interface{}`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AwsFlavorListDto) GetDescriptionOk() (*interface{}, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AwsFlavorListDto) SetDescription(v interface{})`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AwsFlavorListDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *AwsFlavorListDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *AwsFlavorListDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetLinuxPrice

`func (o *AwsFlavorListDto) GetLinuxPrice() string`

GetLinuxPrice returns the LinuxPrice field if non-nil, zero value otherwise.

### GetLinuxPriceOk

`func (o *AwsFlavorListDto) GetLinuxPriceOk() (*string, bool)`

GetLinuxPriceOk returns a tuple with the LinuxPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinuxPrice

`func (o *AwsFlavorListDto) SetLinuxPrice(v string)`

SetLinuxPrice sets LinuxPrice field to given value.

### HasLinuxPrice

`func (o *AwsFlavorListDto) HasLinuxPrice() bool`

HasLinuxPrice returns a boolean if a field has been set.

### SetLinuxPriceNil

`func (o *AwsFlavorListDto) SetLinuxPriceNil(b bool)`

 SetLinuxPriceNil sets the value for LinuxPrice to be an explicit nil

### UnsetLinuxPrice
`func (o *AwsFlavorListDto) UnsetLinuxPrice()`

UnsetLinuxPrice ensures that no value is present for LinuxPrice, not even an explicit nil
### GetWindowsPrice

`func (o *AwsFlavorListDto) GetWindowsPrice() string`

GetWindowsPrice returns the WindowsPrice field if non-nil, zero value otherwise.

### GetWindowsPriceOk

`func (o *AwsFlavorListDto) GetWindowsPriceOk() (*string, bool)`

GetWindowsPriceOk returns a tuple with the WindowsPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowsPrice

`func (o *AwsFlavorListDto) SetWindowsPrice(v string)`

SetWindowsPrice sets WindowsPrice field to given value.

### HasWindowsPrice

`func (o *AwsFlavorListDto) HasWindowsPrice() bool`

HasWindowsPrice returns a boolean if a field has been set.

### SetWindowsPriceNil

`func (o *AwsFlavorListDto) SetWindowsPriceNil(b bool)`

 SetWindowsPriceNil sets the value for WindowsPrice to be an explicit nil

### UnsetWindowsPrice
`func (o *AwsFlavorListDto) UnsetWindowsPrice()`

UnsetWindowsPrice ensures that no value is present for WindowsPrice, not even an explicit nil
### GetWindowsSpotPrice

`func (o *AwsFlavorListDto) GetWindowsSpotPrice() string`

GetWindowsSpotPrice returns the WindowsSpotPrice field if non-nil, zero value otherwise.

### GetWindowsSpotPriceOk

`func (o *AwsFlavorListDto) GetWindowsSpotPriceOk() (*string, bool)`

GetWindowsSpotPriceOk returns a tuple with the WindowsSpotPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowsSpotPrice

`func (o *AwsFlavorListDto) SetWindowsSpotPrice(v string)`

SetWindowsSpotPrice sets WindowsSpotPrice field to given value.

### HasWindowsSpotPrice

`func (o *AwsFlavorListDto) HasWindowsSpotPrice() bool`

HasWindowsSpotPrice returns a boolean if a field has been set.

### SetWindowsSpotPriceNil

`func (o *AwsFlavorListDto) SetWindowsSpotPriceNil(b bool)`

 SetWindowsSpotPriceNil sets the value for WindowsSpotPrice to be an explicit nil

### UnsetWindowsSpotPrice
`func (o *AwsFlavorListDto) UnsetWindowsSpotPrice()`

UnsetWindowsSpotPrice ensures that no value is present for WindowsSpotPrice, not even an explicit nil
### GetLinuxSpotPrice

`func (o *AwsFlavorListDto) GetLinuxSpotPrice() string`

GetLinuxSpotPrice returns the LinuxSpotPrice field if non-nil, zero value otherwise.

### GetLinuxSpotPriceOk

`func (o *AwsFlavorListDto) GetLinuxSpotPriceOk() (*string, bool)`

GetLinuxSpotPriceOk returns a tuple with the LinuxSpotPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinuxSpotPrice

`func (o *AwsFlavorListDto) SetLinuxSpotPrice(v string)`

SetLinuxSpotPrice sets LinuxSpotPrice field to given value.

### HasLinuxSpotPrice

`func (o *AwsFlavorListDto) HasLinuxSpotPrice() bool`

HasLinuxSpotPrice returns a boolean if a field has been set.

### SetLinuxSpotPriceNil

`func (o *AwsFlavorListDto) SetLinuxSpotPriceNil(b bool)`

 SetLinuxSpotPriceNil sets the value for LinuxSpotPrice to be an explicit nil

### UnsetLinuxSpotPrice
`func (o *AwsFlavorListDto) UnsetLinuxSpotPrice()`

UnsetLinuxSpotPrice ensures that no value is present for LinuxSpotPrice, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


