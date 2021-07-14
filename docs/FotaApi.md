# \FotaApi

All URIs are relative to *https://api.lab5e.com/span*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClearFirmwareError**](FotaApi.md#ClearFirmwareError) | **Delete** /collections/{collectionId}/devices/{deviceId}/fwerror | Clear FOTA error
[**CreateFirmware**](FotaApi.md#CreateFirmware) | **Post** /collections/{collectionId}/firmware | Create firmware
[**DeleteFirmware**](FotaApi.md#DeleteFirmware) | **Delete** /collections/{collectionId}/firmware/{imageId} | Delete firmware
[**FirmwareUsage**](FotaApi.md#FirmwareUsage) | **Get** /collections/{collectionId}/firmware/{imageId}/usage | Firmware usage
[**ListFirmware**](FotaApi.md#ListFirmware) | **Get** /collections/{collectionId}/firmware | List firmware
[**RetrieveFirmware**](FotaApi.md#RetrieveFirmware) | **Get** /collections/{collectionId}/firmware/{imageId} | Retrieve firmware
[**UpdateFirmware**](FotaApi.md#UpdateFirmware) | **Patch** /collections/{collectionId}/firmware/{imageId} | Update firmware. Only the version and tags fields can be updated. The other fields will be ignored..



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
    openapiclient "./openapi"
)

func main() {
    collectionId := "collectionId_example" // string | 
    deviceId := "deviceId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FotaApi.ClearFirmwareError(context.Background(), collectionId, deviceId).Execute()
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
    openapiclient "./openapi"
)

func main() {
    collectionId := "collectionId_example" // string | 
    body := *openapiclient.NewCreateFirmwareRequest() // CreateFirmwareRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FotaApi.CreateFirmware(context.Background(), collectionId).Body(body).Execute()
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
    openapiclient "./openapi"
)

func main() {
    collectionId := "collectionId_example" // string | 
    imageId := "imageId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FotaApi.DeleteFirmware(context.Background(), collectionId, imageId).Execute()
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
    openapiclient "./openapi"
)

func main() {
    collectionId := "collectionId_example" // string | 
    imageId := "imageId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FotaApi.FirmwareUsage(context.Background(), collectionId, imageId).Execute()
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
    openapiclient "./openapi"
)

func main() {
    collectionId := "collectionId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FotaApi.ListFirmware(context.Background(), collectionId).Execute()
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
    openapiclient "./openapi"
)

func main() {
    collectionId := "collectionId_example" // string | 
    imageId := "imageId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FotaApi.RetrieveFirmware(context.Background(), collectionId, imageId).Execute()
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


## UpdateFirmware

> Firmware UpdateFirmware(ctx, collectionId, imageId).Body(body).Execute()

Update firmware. Only the version and tags fields can be updated. The other fields will be ignored..

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    collectionId := "collectionId_example" // string | Collection ID  Collection ID for the collection owning the firmware image.
    imageId := "imageId_example" // string | Firmware image ID
    body := *openapiclient.NewFirmware() // Firmware | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FotaApi.UpdateFirmware(context.Background(), collectionId, imageId).Body(body).Execute()
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
**collectionId** | **string** | Collection ID  Collection ID for the collection owning the firmware image. | 
**imageId** | **string** | Firmware image ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFirmwareRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**Firmware**](Firmware.md) |  | 

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

