# \AuthManagementApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthForgotPassword**](AuthManagementApi.md#AuthForgotPassword) | **Post** /api/v1/auth/forgotpassword | Generate reset password token if you forgot password
[**AuthLogin**](AuthManagementApi.md#AuthLogin) | **Post** /api/v1/auth/login | Login to API
[**AuthRefresh**](AuthManagementApi.md#AuthRefresh) | **Post** /api/v1/auth/refresh | Refresh bearer token that generated automatically by API
[**AuthResetPassword**](AuthManagementApi.md#AuthResetPassword) | **Post** /api/v1/auth/resetpassword | Reset password
[**AuthTrial**](AuthManagementApi.md#AuthTrial) | **Post** /api/v1/auth/trial | Try free



## AuthForgotPassword

> AuthForgotPassword(ctx).ForgotPasswordCommand(forgotPasswordCommand).Execute()

Generate reset password token if you forgot password

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
    forgotPasswordCommand := *openapiclient.NewForgotPasswordCommand() // ForgotPasswordCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AuthManagementApi.AuthForgotPassword(context.Background()).ForgotPasswordCommand(forgotPasswordCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthManagementApi.AuthForgotPassword``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthForgotPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **forgotPasswordCommand** | [**ForgotPasswordCommand**](ForgotPasswordCommand.md) |  | 

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


## AuthLogin

> GetToken AuthLogin(ctx).LoginCommand(loginCommand).Execute()

Login to API

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
    loginCommand := *openapiclient.NewLoginCommand() // LoginCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthManagementApi.AuthLogin(context.Background()).LoginCommand(loginCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthManagementApi.AuthLogin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthLogin`: GetToken
    fmt.Fprintf(os.Stdout, "Response from `AuthManagementApi.AuthLogin`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthLoginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **loginCommand** | [**LoginCommand**](LoginCommand.md) |  | 

### Return type

[**GetToken**](GetToken.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthRefresh

> GetToken AuthRefresh(ctx).RefreshTokenCommand(refreshTokenCommand).Execute()

Refresh bearer token that generated automatically by API

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
    refreshTokenCommand := *openapiclient.NewRefreshTokenCommand() // RefreshTokenCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthManagementApi.AuthRefresh(context.Background()).RefreshTokenCommand(refreshTokenCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthManagementApi.AuthRefresh``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthRefresh`: GetToken
    fmt.Fprintf(os.Stdout, "Response from `AuthManagementApi.AuthRefresh`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthRefreshRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **refreshTokenCommand** | [**RefreshTokenCommand**](RefreshTokenCommand.md) |  | 

### Return type

[**GetToken**](GetToken.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthResetPassword

> AuthResetPassword(ctx).ResetPasswordCommand(resetPasswordCommand).Execute()

Reset password

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
    resetPasswordCommand := *openapiclient.NewResetPasswordCommand("Token_example", "NewPassword_example") // ResetPasswordCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AuthManagementApi.AuthResetPassword(context.Background()).ResetPasswordCommand(resetPasswordCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthManagementApi.AuthResetPassword``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthResetPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resetPasswordCommand** | [**ResetPasswordCommand**](ResetPasswordCommand.md) |  | 

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


## AuthTrial

> AuthTrial(ctx).TryForFreeCommand(tryForFreeCommand).Execute()

Try free

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
    tryForFreeCommand := *openapiclient.NewTryForFreeCommand("OrganizationName_example", "Username_example", "Email_example") // TryForFreeCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AuthManagementApi.AuthTrial(context.Background()).TryForFreeCommand(tryForFreeCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthManagementApi.AuthTrial``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthTrialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tryForFreeCommand** | [**TryForFreeCommand**](TryForFreeCommand.md) |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

