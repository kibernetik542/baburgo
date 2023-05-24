# \BillingApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BillingCreate**](BillingApi.md#BillingCreate) | **Post** /api/v1/billing/create | Add billing summary
[**BillingExportCsv**](BillingApi.md#BillingExportCsv) | **Get** /api/v1/billing/export | Export Csv
[**BillingGroupedList**](BillingApi.md#BillingGroupedList) | **Get** /api/v1/billing/grouped | Retrieve a grouped list of billing summaries
[**BillingList**](BillingApi.md#BillingList) | **Get** /api/v1/billing | Retrieve billing info



## BillingCreate

> BillingCreate(ctx).CreateBillingSummaryCommand(createBillingSummaryCommand).Execute()

Add billing summary

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
    createBillingSummaryCommand := *openapiclient.NewCreateBillingSummaryCommand(int32(123), int32(123)) // CreateBillingSummaryCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.BillingApi.BillingCreate(context.Background()).CreateBillingSummaryCommand(createBillingSummaryCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingApi.BillingCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBillingCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createBillingSummaryCommand** | [**CreateBillingSummaryCommand**](CreateBillingSummaryCommand.md) |  | 

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


## BillingExportCsv

> BillingExportCsv(ctx).IsEmailEnabled(isEmailEnabled).OrganizationId(organizationId).StartDate(startDate).EndDate(endDate).IsDeleted(isDeleted).Execute()

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
    organizationId := int32(56) // int32 |  (optional)
    startDate := "startDate_example" // string |  (optional)
    endDate := "endDate_example" // string |  (optional)
    isDeleted := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.BillingApi.BillingExportCsv(context.Background()).IsEmailEnabled(isEmailEnabled).OrganizationId(organizationId).StartDate(startDate).EndDate(endDate).IsDeleted(isDeleted).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingApi.BillingExportCsv``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBillingExportCsvRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **isEmailEnabled** | **bool** |  | 
 **organizationId** | **int32** |  | 
 **startDate** | **string** |  | 
 **endDate** | **string** |  | 
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


## BillingGroupedList

> []GroupedBillingInfo BillingGroupedList(ctx).OrganizationId(organizationId).PeriodDuration(periodDuration).IsDeleted(isDeleted).Execute()

Retrieve a grouped list of billing summaries

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
    organizationId := int32(56) // int32 |  (optional)
    periodDuration := "periodDuration_example" // string |  (optional)
    isDeleted := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BillingApi.BillingGroupedList(context.Background()).OrganizationId(organizationId).PeriodDuration(periodDuration).IsDeleted(isDeleted).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingApi.BillingGroupedList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BillingGroupedList`: []GroupedBillingInfo
    fmt.Fprintf(os.Stdout, "Response from `BillingApi.BillingGroupedList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBillingGroupedListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organizationId** | **int32** |  | 
 **periodDuration** | **string** |  | 
 **isDeleted** | **bool** |  | 

### Return type

[**[]GroupedBillingInfo**](GroupedBillingInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BillingList

> BillingInfo BillingList(ctx).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).StartDate(startDate).EndDate(endDate).OrganizationId(organizationId).IsDeleted(isDeleted).ProjectId(projectId).Execute()

Retrieve billing info

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
    limit := int32(56) // int32 |  (optional)
    offset := int32(56) // int32 |  (optional)
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)
    startDate := "startDate_example" // string |  (optional)
    endDate := "endDate_example" // string |  (optional)
    organizationId := int32(56) // int32 |  (optional)
    isDeleted := true // bool |  (optional)
    projectId := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BillingApi.BillingList(context.Background()).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).StartDate(startDate).EndDate(endDate).OrganizationId(organizationId).IsDeleted(isDeleted).ProjectId(projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingApi.BillingList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BillingList`: BillingInfo
    fmt.Fprintf(os.Stdout, "Response from `BillingApi.BillingList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBillingListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **startDate** | **string** |  | 
 **endDate** | **string** |  | 
 **organizationId** | **int32** |  | 
 **isDeleted** | **bool** |  | 
 **projectId** | **int32** |  | 

### Return type

[**BillingInfo**](BillingInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

