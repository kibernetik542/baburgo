# \AzureCloudCredentialApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AzureCreate**](AzureCloudCredentialApi.md#AzureCreate) | **Post** /api/v1/azure/create | Add Azure credentials
[**AzureDashboard**](AzureCloudCredentialApi.md#AzureDashboard) | **Post** /api/v1/azure/quota/list | Fetch Azure quota list
[**AzureList**](AzureCloudCredentialApi.md#AzureList) | **Get** /api/v1/azure/list | Retrieve list of azure cloud credentials
[**AzureLocations**](AzureCloudCredentialApi.md#AzureLocations) | **Post** /api/v1/azure/locations | Fetch Azure location list
[**AzureOffers**](AzureCloudCredentialApi.md#AzureOffers) | **Get** /api/v1/azure/offers/{cloudId}/{publisher} | 
[**AzurePublishers**](AzureCloudCredentialApi.md#AzurePublishers) | **Get** /api/v1/azure/publishers/{cloudId} | List Azure publishers list
[**AzureSkus**](AzureCloudCredentialApi.md#AzureSkus) | **Get** /api/v1/azure/skus/{cloudId}/{publisher}/{offer} | 
[**AzureSubscriptions**](AzureCloudCredentialApi.md#AzureSubscriptions) | **Post** /api/v1/azure/subscriptions | Azure subscriptions list
[**AzureUpdate**](AzureCloudCredentialApi.md#AzureUpdate) | **Post** /api/v1/azure/update | Update Azure credentials
[**AzureZones**](AzureCloudCredentialApi.md#AzureZones) | **Post** /api/v1/azure/zones | Fetch Azure zone list



## AzureCreate

> ApiResponse AzureCreate(ctx).CreateAzureCloudCommand(createAzureCloudCommand).Execute()

Add Azure credentials

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
    createAzureCloudCommand := *openapiclient.NewCreateAzureCloudCommand("Name_example", "AzureSubscriptionId_example", "AzureClientId_example", "AzureClientSecret_example", "AzureTenantId_example", "AzureLocation_example") // CreateAzureCloudCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AzureCloudCredentialApi.AzureCreate(context.Background()).CreateAzureCloudCommand(createAzureCloudCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AzureCloudCredentialApi.AzureCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AzureCreate`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `AzureCloudCredentialApi.AzureCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAzureCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createAzureCloudCommand** | [**CreateAzureCloudCommand**](CreateAzureCloudCommand.md) |  | 

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


## AzureDashboard

> []AzureQuotaListRecordDto AzureDashboard(ctx).AzureDashboardCommand(azureDashboardCommand).Execute()

Fetch Azure quota list

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
    azureDashboardCommand := *openapiclient.NewAzureDashboardCommand() // AzureDashboardCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AzureCloudCredentialApi.AzureDashboard(context.Background()).AzureDashboardCommand(azureDashboardCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AzureCloudCredentialApi.AzureDashboard``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AzureDashboard`: []AzureQuotaListRecordDto
    fmt.Fprintf(os.Stdout, "Response from `AzureCloudCredentialApi.AzureDashboard`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAzureDashboardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **azureDashboardCommand** | [**AzureDashboardCommand**](AzureDashboardCommand.md) |  | 

### Return type

[**[]AzureQuotaListRecordDto**](AzureQuotaListRecordDto.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AzureList

> AzureCredentialList AzureList(ctx).Limit(limit).Offset(offset).OrganizationId(organizationId).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).Id(id).Execute()

Retrieve list of azure cloud credentials

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
    organizationId := int32(56) // int32 |  (optional)
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)
    search := "search_example" // string |  (optional)
    searchId := "searchId_example" // string |  (optional)
    id := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AzureCloudCredentialApi.AzureList(context.Background()).Limit(limit).Offset(offset).OrganizationId(organizationId).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).Id(id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AzureCloudCredentialApi.AzureList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AzureList`: AzureCredentialList
    fmt.Fprintf(os.Stdout, "Response from `AzureCloudCredentialApi.AzureList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAzureListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **organizationId** | **int32** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 
 **searchId** | **string** |  | 
 **id** | **int32** |  | 

### Return type

[**AzureCredentialList**](AzureCredentialList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AzureLocations

> []string AzureLocations(ctx).AzureLocationsCommand(azureLocationsCommand).Execute()

Fetch Azure location list

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
    azureLocationsCommand := *openapiclient.NewAzureLocationsCommand() // AzureLocationsCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AzureCloudCredentialApi.AzureLocations(context.Background()).AzureLocationsCommand(azureLocationsCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AzureCloudCredentialApi.AzureLocations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AzureLocations`: []string
    fmt.Fprintf(os.Stdout, "Response from `AzureCloudCredentialApi.AzureLocations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAzureLocationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **azureLocationsCommand** | [**AzureLocationsCommand**](AzureLocationsCommand.md) |  | 

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


## AzureOffers

> AzureOffersList AzureOffers(ctx, cloudId, publisher).Offset(offset).Limit(limit).SortBy(sortBy).SortDirection(sortDirection).Search(search).Execute()



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
    cloudId := int32(56) // int32 | 
    publisher := "publisher_example" // string | 
    offset := int32(56) // int32 |  (optional)
    limit := int32(56) // int32 |  (optional)
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)
    search := "search_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AzureCloudCredentialApi.AzureOffers(context.Background(), cloudId, publisher).Offset(offset).Limit(limit).SortBy(sortBy).SortDirection(sortDirection).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AzureCloudCredentialApi.AzureOffers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AzureOffers`: AzureOffersList
    fmt.Fprintf(os.Stdout, "Response from `AzureCloudCredentialApi.AzureOffers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudId** | **int32** |  | 
**publisher** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAzureOffersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **offset** | **int32** |  | 
 **limit** | **int32** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 

### Return type

[**AzureOffersList**](AzureOffersList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AzurePublishers

> AzurePublishersList AzurePublishers(ctx, cloudId).Offset(offset).Limit(limit).SortBy(sortBy).SortDirection(sortDirection).Search(search).Execute()

List Azure publishers list

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
    cloudId := int32(56) // int32 | 
    offset := int32(56) // int32 |  (optional)
    limit := int32(56) // int32 |  (optional)
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)
    search := "search_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AzureCloudCredentialApi.AzurePublishers(context.Background(), cloudId).Offset(offset).Limit(limit).SortBy(sortBy).SortDirection(sortDirection).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AzureCloudCredentialApi.AzurePublishers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AzurePublishers`: AzurePublishersList
    fmt.Fprintf(os.Stdout, "Response from `AzureCloudCredentialApi.AzurePublishers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAzurePublishersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **int32** |  | 
 **limit** | **int32** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 

### Return type

[**AzurePublishersList**](AzurePublishersList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AzureSkus

> AzureSkusList AzureSkus(ctx, cloudId, publisher, offer).Offset(offset).Limit(limit).SortBy(sortBy).SortDirection(sortDirection).Search(search).Execute()



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
    cloudId := int32(56) // int32 | 
    publisher := "publisher_example" // string | 
    offer := "offer_example" // string | 
    offset := int32(56) // int32 |  (optional)
    limit := int32(56) // int32 |  (optional)
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)
    search := "search_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AzureCloudCredentialApi.AzureSkus(context.Background(), cloudId, publisher, offer).Offset(offset).Limit(limit).SortBy(sortBy).SortDirection(sortDirection).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AzureCloudCredentialApi.AzureSkus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AzureSkus`: AzureSkusList
    fmt.Fprintf(os.Stdout, "Response from `AzureCloudCredentialApi.AzureSkus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudId** | **int32** |  | 
**publisher** | **string** |  | 
**offer** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAzureSkusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **offset** | **int32** |  | 
 **limit** | **int32** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 

### Return type

[**AzureSkusList**](AzureSkusList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AzureSubscriptions

> []CommonStringBasedDropdownDto AzureSubscriptions(ctx).AzureSubscriptionListCommand(azureSubscriptionListCommand).Execute()

Azure subscriptions list

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
    azureSubscriptionListCommand := *openapiclient.NewAzureSubscriptionListCommand("ClientId_example", "ClientSecret_example", "TenantId_example") // AzureSubscriptionListCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AzureCloudCredentialApi.AzureSubscriptions(context.Background()).AzureSubscriptionListCommand(azureSubscriptionListCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AzureCloudCredentialApi.AzureSubscriptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AzureSubscriptions`: []CommonStringBasedDropdownDto
    fmt.Fprintf(os.Stdout, "Response from `AzureCloudCredentialApi.AzureSubscriptions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAzureSubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **azureSubscriptionListCommand** | [**AzureSubscriptionListCommand**](AzureSubscriptionListCommand.md) |  | 

### Return type

[**[]CommonStringBasedDropdownDto**](CommonStringBasedDropdownDto.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AzureUpdate

> AzureUpdate(ctx).UpdateAzureCommand(updateAzureCommand).Execute()

Update Azure credentials

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
    updateAzureCommand := *openapiclient.NewUpdateAzureCommand(int32(123), "Name_example", "AzureClientSecret_example", "AzureClientId_example") // UpdateAzureCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AzureCloudCredentialApi.AzureUpdate(context.Background()).UpdateAzureCommand(updateAzureCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AzureCloudCredentialApi.AzureUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAzureUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateAzureCommand** | [**UpdateAzureCommand**](UpdateAzureCommand.md) |  | 

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


## AzureZones

> AzResult AzureZones(ctx).AzureZonesCommand(azureZonesCommand).Execute()

Fetch Azure zone list

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
    azureZonesCommand := *openapiclient.NewAzureZonesCommand() // AzureZonesCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AzureCloudCredentialApi.AzureZones(context.Background()).AzureZonesCommand(azureZonesCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AzureCloudCredentialApi.AzureZones``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AzureZones`: AzResult
    fmt.Fprintf(os.Stdout, "Response from `AzureCloudCredentialApi.AzureZones`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAzureZonesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **azureZonesCommand** | [**AzureZonesCommand**](AzureZonesCommand.md) |  | 

### Return type

[**AzResult**](AzResult.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

