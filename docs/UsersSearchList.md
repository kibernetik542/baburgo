# UsersSearchList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]UsersSearchResponseData**](UsersSearchResponseData.md) |  | [optional] 
**TotalCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewUsersSearchList

`func NewUsersSearchList() *UsersSearchList`

NewUsersSearchList instantiates a new UsersSearchList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsersSearchListWithDefaults

`func NewUsersSearchListWithDefaults() *UsersSearchList`

NewUsersSearchListWithDefaults instantiates a new UsersSearchList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *UsersSearchList) GetData() []UsersSearchResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *UsersSearchList) GetDataOk() (*[]UsersSearchResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *UsersSearchList) SetData(v []UsersSearchResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *UsersSearchList) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *UsersSearchList) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *UsersSearchList) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetTotalCount

`func (o *UsersSearchList) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *UsersSearchList) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *UsersSearchList) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *UsersSearchList) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


