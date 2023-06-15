# \FotaApi

All URIs are relative to *https://api.lab5e.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClearFirmwareError**](FotaApi.md#ClearFirmwareError) | **Delete** /span/collections/{collectionId}/devices/{deviceId}/fwerror | Clear FOTA error
[**CreateFirmware**](FotaApi.md#CreateFirmware) | **Post** /span/collections/{collectionId}/firmware | Create firmware
[**DeleteFirmware**](FotaApi.md#DeleteFirmware) | **Delete** /span/collections/{collectionId}/firmware/{imageId} | Delete firmware
[**FirmwareUsage**](FotaApi.md#FirmwareUsage) | **Get** /span/collections/{collectionId}/firmware/{imageId}/usage | Firmware usage
[**ListFirmware**](FotaApi.md#ListFirmware) | **Get** /span/collections/{collectionId}/firmware | List firmware
[**RetrieveFirmware**](FotaApi.md#RetrieveFirmware) | **Get** /span/collections/{collectionId}/firmware/{imageId} | Retrieve firmware
[**RetrieveFirmwareStats**](FotaApi.md#RetrieveFirmwareStats) | **Get** /span/collections/{collectionId}/firmware/{imageId}/stats | Retrieve firmware statistics
[**UpdateFirmware**](FotaApi.md#UpdateFirmware) | **Patch** /span/collections/{existingCollectionId}/firmware/{imageId} | Update firmware



## ClearFirmwareError

> ClearFirmwareErrorResponse ClearFirmwareError(ctx, collectionId, deviceId).Execute()

Clear FOTA error

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/lab5e/go-spanapi"
)

func main() {
    collectionId := "collectionId_example" // string | 
    deviceId := "deviceId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FotaApi.ClearFirmwareError(context.Background(), collectionId, deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FotaApi.ClearFirmwareError``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ClearFirmwareError`: ClearFirmwareErrorResponse
    fmt.Fprintf(os.Stdout, "Response from `FotaApi.ClearFirmwareError`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** |  | 
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiClearFirmwareErrorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ClearFirmwareErrorResponse**](ClearFirmwareErrorResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateFirmware

> Firmware CreateFirmware(ctx, collectionId).Body(body).Execute()

Create firmware



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/lab5e/go-spanapi"
)

func main() {
    collectionId := "collectionId_example" // string | 
    body := *openapiclient.NewCreateFirmwareRequest() // CreateFirmwareRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FotaApi.CreateFirmware(context.Background(), collectionId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FotaApi.CreateFirmware``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateFirmware`: Firmware
    fmt.Fprintf(os.Stdout, "Response from `FotaApi.CreateFirmware`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateFirmwareRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**CreateFirmwareRequest**](CreateFirmwareRequest.md) |  | 

### Return type

[**Firmware**](Firmware.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFirmware

> Firmware DeleteFirmware(ctx, collectionId, imageId).Execute()

Delete firmware

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/lab5e/go-spanapi"
)

func main() {
    collectionId := "collectionId_example" // string | 
    imageId := "imageId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FotaApi.DeleteFirmware(context.Background(), collectionId, imageId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FotaApi.DeleteFirmware``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteFirmware`: Firmware
    fmt.Fprintf(os.Stdout, "Response from `FotaApi.DeleteFirmware`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** |  | 
**imageId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFirmwareRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Firmware**](Firmware.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FirmwareUsage

> FirmwareUsageResponse FirmwareUsage(ctx, collectionId, imageId).Execute()

Firmware usage

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/lab5e/go-spanapi"
)

func main() {
    collectionId := "collectionId_example" // string | 
    imageId := "imageId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FotaApi.FirmwareUsage(context.Background(), collectionId, imageId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FotaApi.FirmwareUsage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FirmwareUsage`: FirmwareUsageResponse
    fmt.Fprintf(os.Stdout, "Response from `FotaApi.FirmwareUsage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** |  | 
**imageId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFirmwareUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**FirmwareUsageResponse**](FirmwareUsageResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFirmware

> ListFirmwareResponse ListFirmware(ctx, collectionId).Execute()

List firmware

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/lab5e/go-spanapi"
)

func main() {
    collectionId := "collectionId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FotaApi.ListFirmware(context.Background(), collectionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FotaApi.ListFirmware``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListFirmware`: ListFirmwareResponse
    fmt.Fprintf(os.Stdout, "Response from `FotaApi.ListFirmware`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListFirmwareRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListFirmwareResponse**](ListFirmwareResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveFirmware

> Firmware RetrieveFirmware(ctx, collectionId, imageId).Execute()

Retrieve firmware

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/lab5e/go-spanapi"
)

func main() {
    collectionId := "collectionId_example" // string | 
    imageId := "imageId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FotaApi.RetrieveFirmware(context.Background(), collectionId, imageId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FotaApi.RetrieveFirmware``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveFirmware`: Firmware
    fmt.Fprintf(os.Stdout, "Response from `FotaApi.RetrieveFirmware`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** |  | 
**imageId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveFirmwareRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Firmware**](Firmware.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveFirmwareStats

> FirmwareStats RetrieveFirmwareStats(ctx, collectionId, imageId).Execute()

Retrieve firmware statistics

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/lab5e/go-spanapi"
)

func main() {
    collectionId := "collectionId_example" // string | 
    imageId := "imageId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FotaApi.RetrieveFirmwareStats(context.Background(), collectionId, imageId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FotaApi.RetrieveFirmwareStats``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveFirmwareStats`: FirmwareStats
    fmt.Fprintf(os.Stdout, "Response from `FotaApi.RetrieveFirmwareStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** |  | 
**imageId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveFirmwareStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**FirmwareStats**](FirmwareStats.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFirmware

> Firmware UpdateFirmware(ctx, existingCollectionId, imageId).Body(body).Execute()

Update firmware



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/lab5e/go-spanapi"
)

func main() {
    existingCollectionId := "existingCollectionId_example" // string | 
    imageId := "imageId_example" // string | 
    body := *openapiclient.NewUpdateFirmwareRequest() // UpdateFirmwareRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FotaApi.UpdateFirmware(context.Background(), existingCollectionId, imageId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FotaApi.UpdateFirmware``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateFirmware`: Firmware
    fmt.Fprintf(os.Stdout, "Response from `FotaApi.UpdateFirmware`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**existingCollectionId** | **string** |  | 
**imageId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFirmwareRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**UpdateFirmwareRequest**](UpdateFirmwareRequest.md) |  | 

### Return type

[**Firmware**](Firmware.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

