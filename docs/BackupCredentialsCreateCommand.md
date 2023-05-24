# BackupCredentialsCreateCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**S3Name** | **string** |  | 
**S3AccessKeyId** | **string** |  | 
**S3SecretKey** | **string** |  | 
**S3Endpoint** | **string** |  | 
**S3Region** | Pointer to **NullableString** |  | [optional] 
**OrganizationId** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewBackupCredentialsCreateCommand

`func NewBackupCredentialsCreateCommand(s3Name string, s3AccessKeyId string, s3SecretKey string, s3Endpoint string, ) *BackupCredentialsCreateCommand`

NewBackupCredentialsCreateCommand instantiates a new BackupCredentialsCreateCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupCredentialsCreateCommandWithDefaults

`func NewBackupCredentialsCreateCommandWithDefaults() *BackupCredentialsCreateCommand`

NewBackupCredentialsCreateCommandWithDefaults instantiates a new BackupCredentialsCreateCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetS3Name

`func (o *BackupCredentialsCreateCommand) GetS3Name() string`

GetS3Name returns the S3Name field if non-nil, zero value otherwise.

### GetS3NameOk

`func (o *BackupCredentialsCreateCommand) GetS3NameOk() (*string, bool)`

GetS3NameOk returns a tuple with the S3Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3Name

`func (o *BackupCredentialsCreateCommand) SetS3Name(v string)`

SetS3Name sets S3Name field to given value.


### GetS3AccessKeyId

`func (o *BackupCredentialsCreateCommand) GetS3AccessKeyId() string`

GetS3AccessKeyId returns the S3AccessKeyId field if non-nil, zero value otherwise.

### GetS3AccessKeyIdOk

`func (o *BackupCredentialsCreateCommand) GetS3AccessKeyIdOk() (*string, bool)`

GetS3AccessKeyIdOk returns a tuple with the S3AccessKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3AccessKeyId

`func (o *BackupCredentialsCreateCommand) SetS3AccessKeyId(v string)`

SetS3AccessKeyId sets S3AccessKeyId field to given value.


### GetS3SecretKey

`func (o *BackupCredentialsCreateCommand) GetS3SecretKey() string`

GetS3SecretKey returns the S3SecretKey field if non-nil, zero value otherwise.

### GetS3SecretKeyOk

`func (o *BackupCredentialsCreateCommand) GetS3SecretKeyOk() (*string, bool)`

GetS3SecretKeyOk returns a tuple with the S3SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3SecretKey

`func (o *BackupCredentialsCreateCommand) SetS3SecretKey(v string)`

SetS3SecretKey sets S3SecretKey field to given value.


### GetS3Endpoint

`func (o *BackupCredentialsCreateCommand) GetS3Endpoint() string`

GetS3Endpoint returns the S3Endpoint field if non-nil, zero value otherwise.

### GetS3EndpointOk

`func (o *BackupCredentialsCreateCommand) GetS3EndpointOk() (*string, bool)`

GetS3EndpointOk returns a tuple with the S3Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3Endpoint

`func (o *BackupCredentialsCreateCommand) SetS3Endpoint(v string)`

SetS3Endpoint sets S3Endpoint field to given value.


### GetS3Region

`func (o *BackupCredentialsCreateCommand) GetS3Region() string`

GetS3Region returns the S3Region field if non-nil, zero value otherwise.

### GetS3RegionOk

`func (o *BackupCredentialsCreateCommand) GetS3RegionOk() (*string, bool)`

GetS3RegionOk returns a tuple with the S3Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3Region

`func (o *BackupCredentialsCreateCommand) SetS3Region(v string)`

SetS3Region sets S3Region field to given value.

### HasS3Region

`func (o *BackupCredentialsCreateCommand) HasS3Region() bool`

HasS3Region returns a boolean if a field has been set.

### SetS3RegionNil

`func (o *BackupCredentialsCreateCommand) SetS3RegionNil(b bool)`

 SetS3RegionNil sets the value for S3Region to be an explicit nil

### UnsetS3Region
`func (o *BackupCredentialsCreateCommand) UnsetS3Region()`

UnsetS3Region ensures that no value is present for S3Region, not even an explicit nil
### GetOrganizationId

`func (o *BackupCredentialsCreateCommand) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *BackupCredentialsCreateCommand) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *BackupCredentialsCreateCommand) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *BackupCredentialsCreateCommand) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### SetOrganizationIdNil

`func (o *BackupCredentialsCreateCommand) SetOrganizationIdNil(b bool)`

 SetOrganizationIdNil sets the value for OrganizationId to be an explicit nil

### UnsetOrganizationId
`func (o *BackupCredentialsCreateCommand) UnsetOrganizationId()`

UnsetOrganizationId ensures that no value is present for OrganizationId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


