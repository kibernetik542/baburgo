# \NotificationsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NotificationsCreate**](NotificationsApi.md#NotificationsCreate) | **Post** /api/v1/notifications/add | Create notification
[**NotificationsExportCsv**](NotificationsApi.md#NotificationsExportCsv) | **Get** /api/v1/notifications/download | Export Csv
[**NotificationsNotifyOwner**](NotificationsApi.md#NotificationsNotifyOwner) | **Post** /api/v1/notifications/notifyowner | Notify owner
[**NotificationsOperationMessages**](NotificationsApi.md#NotificationsOperationMessages) | **Post** /api/v1/notifications/operations | Get project operations



## NotificationsCreate

> NotificationsCreate(ctx).NotificationSendCommand(notificationSendCommand).Execute()

Create notification

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/kibernetik542/baburgo"
)

func main() {
    notificationSendCommand := *openapiclient.NewNotificationSendCommand() // NotificationSendCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.NotificationsApi.NotificationsCreate(context.Background()).NotificationSendCommand(notificationSendCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsApi.NotificationsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNotificationsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **notificationSendCommand** | [**NotificationSendCommand**](NotificationSendCommand.md) |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NotificationsExportCsv

> NotificationsExportCsv(ctx).IsEmailEnabled(isEmailEnabled).SortBy(sortBy).SortDirection(sortDirection).StartDate(startDate).EndDate(endDate).OrganizationId(organizationId).FilterBy(filterBy).ProjectId(projectId).UserId(userId).IsDeleted(isDeleted).Execute()

Export Csv

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/kibernetik542/baburgo"
)

func main() {
    isEmailEnabled := true // bool | 
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)
    startDate := "startDate_example" // string |  (optional)
    endDate := "endDate_example" // string |  (optional)
    organizationId := int32(56) // int32 |  (optional)
    filterBy := "filterBy_example" // string |  (optional)
    projectId := int32(56) // int32 |  (optional)
    userId := "userId_example" // string |  (optional)
    isDeleted := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.NotificationsApi.NotificationsExportCsv(context.Background()).IsEmailEnabled(isEmailEnabled).SortBy(sortBy).SortDirection(sortDirection).StartDate(startDate).EndDate(endDate).OrganizationId(organizationId).FilterBy(filterBy).ProjectId(projectId).UserId(userId).IsDeleted(isDeleted).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsApi.NotificationsExportCsv``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNotificationsExportCsvRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **isEmailEnabled** | **bool** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **startDate** | **string** |  | 
 **endDate** | **string** |  | 
 **organizationId** | **int32** |  | 
 **filterBy** | **string** |  | 
 **projectId** | **int32** |  | 
 **userId** | **string** |  | 
 **isDeleted** | **bool** |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NotificationsNotifyOwner

> NotificationsNotifyOwner(ctx).NotifyOwnersCommand(notifyOwnersCommand).Execute()

Notify owner

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/kibernetik542/baburgo"
)

func main() {
    notifyOwnersCommand := *openapiclient.NewNotifyOwnersCommand() // NotifyOwnersCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.NotificationsApi.NotificationsNotifyOwner(context.Background()).NotifyOwnersCommand(notifyOwnersCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsApi.NotificationsNotifyOwner``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNotificationsNotifyOwnerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **notifyOwnersCommand** | [**NotifyOwnersCommand**](NotifyOwnersCommand.md) |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NotificationsOperationMessages

> interface{} NotificationsOperationMessages(ctx).GetProjectOperationCommand(getProjectOperationCommand).Execute()

Get project operations

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/kibernetik542/baburgo"
)

func main() {
    getProjectOperationCommand := *openapiclient.NewGetProjectOperationCommand() // GetProjectOperationCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsApi.NotificationsOperationMessages(context.Background()).GetProjectOperationCommand(getProjectOperationCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsApi.NotificationsOperationMessages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NotificationsOperationMessages`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `NotificationsApi.NotificationsOperationMessages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNotificationsOperationMessagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getProjectOperationCommand** | [**GetProjectOperationCommand**](GetProjectOperationCommand.md) |  | 

### Return type

**interface{}**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

