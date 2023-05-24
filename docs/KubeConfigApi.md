# \KubeConfigApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**KubeconfigCreate**](KubeConfigApi.md#KubeconfigCreate) | **Post** /api/v1/kubeconfig | Create kube config
[**KubeconfigDelete**](KubeConfigApi.md#KubeconfigDelete) | **Post** /api/v1/kubeconfig/delete | Delete kube config
[**KubeconfigDeleteByProjectId**](KubeConfigApi.md#KubeconfigDeleteByProjectId) | **Post** /api/v1/kubeconfig/delete-by-project-id | Delete kube config by project id
[**KubeconfigDownload**](KubeConfigApi.md#KubeconfigDownload) | **Post** /api/v1/kubeconfig/download | Download kube config file for user by project Id
[**KubeconfigInteractiveShell**](KubeConfigApi.md#KubeconfigInteractiveShell) | **Post** /api/v1/kubeconfig/interactive-shell | Interactive shell for user kube config
[**KubeconfigList**](KubeConfigApi.md#KubeconfigList) | **Get** /api/v1/kubeconfig | Retrieve a list of kube configs for project



## KubeconfigCreate

> ApiResponse KubeconfigCreate(ctx).CreateKubeConfigCommand(createKubeConfigCommand).Execute()

Create kube config

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
    createKubeConfigCommand := *openapiclient.NewCreateKubeConfigCommand("Name_example", int32(123)) // CreateKubeConfigCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubeConfigApi.KubeconfigCreate(context.Background()).CreateKubeConfigCommand(createKubeConfigCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubeConfigApi.KubeconfigCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubeconfigCreate`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `KubeConfigApi.KubeconfigCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKubeconfigCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createKubeConfigCommand** | [**CreateKubeConfigCommand**](CreateKubeConfigCommand.md) |  | 

### Return type

[**ApiResponse**](ApiResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubeconfigDelete

> KubeconfigDelete(ctx).DeleteKubeConfigCommand(deleteKubeConfigCommand).Execute()

Delete kube config

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
    deleteKubeConfigCommand := *openapiclient.NewDeleteKubeConfigCommand() // DeleteKubeConfigCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.KubeConfigApi.KubeconfigDelete(context.Background()).DeleteKubeConfigCommand(deleteKubeConfigCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubeConfigApi.KubeconfigDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKubeconfigDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteKubeConfigCommand** | [**DeleteKubeConfigCommand**](DeleteKubeConfigCommand.md) |  | 

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


## KubeconfigDeleteByProjectId

> KubeconfigDeleteByProjectId(ctx).DeleteKubeConfigByProjectIdCommand(deleteKubeConfigByProjectIdCommand).Execute()

Delete kube config by project id

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
    deleteKubeConfigByProjectIdCommand := *openapiclient.NewDeleteKubeConfigByProjectIdCommand() // DeleteKubeConfigByProjectIdCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.KubeConfigApi.KubeconfigDeleteByProjectId(context.Background()).DeleteKubeConfigByProjectIdCommand(deleteKubeConfigByProjectIdCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubeConfigApi.KubeconfigDeleteByProjectId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKubeconfigDeleteByProjectIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteKubeConfigByProjectIdCommand** | [**DeleteKubeConfigByProjectIdCommand**](DeleteKubeConfigByProjectIdCommand.md) |  | 

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


## KubeconfigDownload

> string KubeconfigDownload(ctx).DownloadKubeConfigCommand(downloadKubeConfigCommand).Execute()

Download kube config file for user by project Id

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
    downloadKubeConfigCommand := *openapiclient.NewDownloadKubeConfigCommand() // DownloadKubeConfigCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubeConfigApi.KubeconfigDownload(context.Background()).DownloadKubeConfigCommand(downloadKubeConfigCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubeConfigApi.KubeconfigDownload``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubeconfigDownload`: string
    fmt.Fprintf(os.Stdout, "Response from `KubeConfigApi.KubeconfigDownload`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKubeconfigDownloadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **downloadKubeConfigCommand** | [**DownloadKubeConfigCommand**](DownloadKubeConfigCommand.md) |  | 

### Return type

**string**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubeconfigInteractiveShell

> string KubeconfigInteractiveShell(ctx).KubeConfigInteractiveShellCommand(kubeConfigInteractiveShellCommand).Execute()

Interactive shell for user kube config

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
    kubeConfigInteractiveShellCommand := *openapiclient.NewKubeConfigInteractiveShellCommand(int32(123), "Token_example", int32(123)) // KubeConfigInteractiveShellCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubeConfigApi.KubeconfigInteractiveShell(context.Background()).KubeConfigInteractiveShellCommand(kubeConfigInteractiveShellCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubeConfigApi.KubeconfigInteractiveShell``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubeconfigInteractiveShell`: string
    fmt.Fprintf(os.Stdout, "Response from `KubeConfigApi.KubeconfigInteractiveShell`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKubeconfigInteractiveShellRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kubeConfigInteractiveShellCommand** | [**KubeConfigInteractiveShellCommand**](KubeConfigInteractiveShellCommand.md) |  | 

### Return type

**string**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubeconfigList

> KubeConfigForUserList KubeconfigList(ctx).ProjectId(projectId).OrganizationId(organizationId).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).Id(id).Execute()

Retrieve a list of kube configs for project

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
    projectId := int32(56) // int32 | 
    organizationId := int32(56) // int32 |  (optional)
    limit := int32(56) // int32 |  (optional)
    offset := int32(56) // int32 |  (optional)
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)
    search := "search_example" // string |  (optional)
    id := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubeConfigApi.KubeconfigList(context.Background()).ProjectId(projectId).OrganizationId(organizationId).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).Id(id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubeConfigApi.KubeconfigList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubeconfigList`: KubeConfigForUserList
    fmt.Fprintf(os.Stdout, "Response from `KubeConfigApi.KubeconfigList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKubeconfigListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectId** | **int32** |  | 
 **organizationId** | **int32** |  | 
 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 
 **id** | **int32** |  | 

### Return type

[**KubeConfigForUserList**](KubeConfigForUserList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

