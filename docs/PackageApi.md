# \PackageApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PackageDetails**](PackageApi.md#PackageDetails) | **Get** /api/v1/package/{repoName}/{packageName} | 
[**PackageList**](PackageApi.md#PackageList) | **Get** /api/v1/package/list | Retrieve all available packages
[**PackageValue**](PackageApi.md#PackageValue) | **Post** /api/v1/package/value | Get yaml based value
[**PackageValueAutocomplete**](PackageApi.md#PackageValueAutocomplete) | **Post** /api/v1/package/value-autocomplete | Get autocomplete dictionary
[**PackageVersions**](PackageApi.md#PackageVersions) | **Post** /api/v1/package/versions | Get available versions



## PackageDetails

> AvailablePackageDetailsDto PackageDetails(ctx, repoName, packageName).Execute()



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
    repoName := "repoName_example" // string | 
    packageName := "packageName_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PackageApi.PackageDetails(context.Background(), repoName, packageName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PackageApi.PackageDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PackageDetails`: AvailablePackageDetailsDto
    fmt.Fprintf(os.Stdout, "Response from `PackageApi.PackageDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repoName** | **string** |  | 
**packageName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPackageDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AvailablePackageDetailsDto**](AvailablePackageDetailsDto.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackageList

> AvailablePackagesList PackageList(ctx).Offset(offset).Limit(limit).SortBy(sortBy).SortDirection(sortDirection).Search(search).Id(id).CatalogId(catalogId).Execute()

Retrieve all available packages

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
    offset := int32(56) // int32 |  (optional)
    limit := int32(56) // int32 |  (optional)
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)
    search := "search_example" // string |  (optional)
    id := "id_example" // string |  (optional)
    catalogId := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PackageApi.PackageList(context.Background()).Offset(offset).Limit(limit).SortBy(sortBy).SortDirection(sortDirection).Search(search).Id(id).CatalogId(catalogId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PackageApi.PackageList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PackageList`: AvailablePackagesList
    fmt.Fprintf(os.Stdout, "Response from `PackageApi.PackageList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPackageListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int32** |  | 
 **limit** | **int32** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 
 **id** | **string** |  | 
 **catalogId** | **int32** |  | 

### Return type

[**AvailablePackagesList**](AvailablePackagesList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackageValue

> string PackageValue(ctx).GetCatalogAppValueCommand(getCatalogAppValueCommand).Execute()

Get yaml based value

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
    getCatalogAppValueCommand := *openapiclient.NewGetCatalogAppValueCommand("PackageId_example", "Version_example") // GetCatalogAppValueCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PackageApi.PackageValue(context.Background()).GetCatalogAppValueCommand(getCatalogAppValueCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PackageApi.PackageValue``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PackageValue`: string
    fmt.Fprintf(os.Stdout, "Response from `PackageApi.PackageValue`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPackageValueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getCatalogAppValueCommand** | [**GetCatalogAppValueCommand**](GetCatalogAppValueCommand.md) |  | 

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


## PackageValueAutocomplete

> []PackageAutocompleteDto PackageValueAutocomplete(ctx).GetCatalogAppValueAutocompleteCommand(getCatalogAppValueAutocompleteCommand).Execute()

Get autocomplete dictionary

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
    getCatalogAppValueAutocompleteCommand := *openapiclient.NewGetCatalogAppValueAutocompleteCommand("PackageId_example", "Version_example") // GetCatalogAppValueAutocompleteCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PackageApi.PackageValueAutocomplete(context.Background()).GetCatalogAppValueAutocompleteCommand(getCatalogAppValueAutocompleteCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PackageApi.PackageValueAutocomplete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PackageValueAutocomplete`: []PackageAutocompleteDto
    fmt.Fprintf(os.Stdout, "Response from `PackageApi.PackageValueAutocomplete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPackageValueAutocompleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getCatalogAppValueAutocompleteCommand** | [**GetCatalogAppValueAutocompleteCommand**](GetCatalogAppValueAutocompleteCommand.md) |  | 

### Return type

[**[]PackageAutocompleteDto**](PackageAutocompleteDto.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackageVersions

> []string PackageVersions(ctx).ListCatalogAppAvailableVersionsCommand(listCatalogAppAvailableVersionsCommand).Execute()

Get available versions

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
    listCatalogAppAvailableVersionsCommand := *openapiclient.NewListCatalogAppAvailableVersionsCommand() // ListCatalogAppAvailableVersionsCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PackageApi.PackageVersions(context.Background()).ListCatalogAppAvailableVersionsCommand(listCatalogAppAvailableVersionsCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PackageApi.PackageVersions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PackageVersions`: []string
    fmt.Fprintf(os.Stdout, "Response from `PackageApi.PackageVersions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPackageVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **listCatalogAppAvailableVersionsCommand** | [**ListCatalogAppAvailableVersionsCommand**](ListCatalogAppAvailableVersionsCommand.md) |  | 

### Return type

**[]string**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

