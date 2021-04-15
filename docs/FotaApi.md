# \FotaApi

All URIs are relative to *https://api.lab5e.com/span*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SpanClearFirmwareError**](FotaApi.md#SpanClearFirmwareError) | **Delete** /collections/{collectionId}/devices/{deviceId}/fwerror | Clear FOTA error
[**SpanCreateFirmware**](FotaApi.md#SpanCreateFirmware) | **Post** /collections/{collectionId}/firmware | Create firmware
[**SpanDeleteFirmware**](FotaApi.md#SpanDeleteFirmware) | **Delete** /collections/{collectionId}/firmware/{imageId} | Delete firmware
[**SpanFirmwareUsage**](FotaApi.md#SpanFirmwareUsage) | **Patch** /collections/{collectionId}/firmware/{imageId}/usage | Firmware usage
[**SpanListFirmware**](FotaApi.md#SpanListFirmware) | **Get** /collections/{collectionId}/firmware | List firmware
[**SpanRetrieveFirmware**](FotaApi.md#SpanRetrieveFirmware) | **Get** /collections/{collectionId}/firmware/{imageId} | Retrieve firmware
[**SpanUpdateFirmware**](FotaApi.md#SpanUpdateFirmware) | **Patch** /collections/{collectionId}/firmware/{imageId} | Update firmware



## SpanClearFirmwareError

> ClearFirmwareErrorResponse SpanClearFirmwareError(ctx, collectionId, deviceId).Execute()

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
    resp, r, err := api_client.FotaApi.SpanClearFirmwareError(context.Background(), collectionId, deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FotaApi.SpanClearFirmwareError``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SpanClearFirmwareError`: ClearFirmwareErrorResponse
    fmt.Fprintf(os.Stdout, "Response from `FotaApi.SpanClearFirmwareError`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** |  | 
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSpanClearFirmwareErrorRequest struct via the builder pattern


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


## SpanCreateFirmware

> Firmware SpanCreateFirmware(ctx, collectionId).Body(body).Execute()

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
    resp, r, err := api_client.FotaApi.SpanCreateFirmware(context.Background(), collectionId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FotaApi.SpanCreateFirmware``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SpanCreateFirmware`: Firmware
    fmt.Fprintf(os.Stdout, "Response from `FotaApi.SpanCreateFirmware`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSpanCreateFirmwareRequest struct via the builder pattern


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


## SpanDeleteFirmware

> Firmware SpanDeleteFirmware(ctx, collectionId, imageId).Execute()

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
    resp, r, err := api_client.FotaApi.SpanDeleteFirmware(context.Background(), collectionId, imageId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FotaApi.SpanDeleteFirmware``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SpanDeleteFirmware`: Firmware
    fmt.Fprintf(os.Stdout, "Response from `FotaApi.SpanDeleteFirmware`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** |  | 
**imageId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSpanDeleteFirmwareRequest struct via the builder pattern


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


## SpanFirmwareUsage

> FirmwareUsageResponse SpanFirmwareUsage(ctx, collectionId, imageId).Execute()

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
    resp, r, err := api_client.FotaApi.SpanFirmwareUsage(context.Background(), collectionId, imageId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FotaApi.SpanFirmwareUsage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SpanFirmwareUsage`: FirmwareUsageResponse
    fmt.Fprintf(os.Stdout, "Response from `FotaApi.SpanFirmwareUsage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** |  | 
**imageId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSpanFirmwareUsageRequest struct via the builder pattern


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


## SpanListFirmware

> ListFirmwareResponse SpanListFirmware(ctx, collectionId).Execute()

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
    resp, r, err := api_client.FotaApi.SpanListFirmware(context.Background(), collectionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FotaApi.SpanListFirmware``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SpanListFirmware`: ListFirmwareResponse
    fmt.Fprintf(os.Stdout, "Response from `FotaApi.SpanListFirmware`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSpanListFirmwareRequest struct via the builder pattern


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


## SpanRetrieveFirmware

> Firmware SpanRetrieveFirmware(ctx, collectionId, imageId).Execute()

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
    resp, r, err := api_client.FotaApi.SpanRetrieveFirmware(context.Background(), collectionId, imageId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FotaApi.SpanRetrieveFirmware``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SpanRetrieveFirmware`: Firmware
    fmt.Fprintf(os.Stdout, "Response from `FotaApi.SpanRetrieveFirmware`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** |  | 
**imageId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSpanRetrieveFirmwareRequest struct via the builder pattern


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


## SpanUpdateFirmware

> Firmware SpanUpdateFirmware(ctx, collectionId, imageId).Body(body).Execute()

Update firmware

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
    resp, r, err := api_client.FotaApi.SpanUpdateFirmware(context.Background(), collectionId, imageId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FotaApi.SpanUpdateFirmware``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SpanUpdateFirmware`: Firmware
    fmt.Fprintf(os.Stdout, "Response from `FotaApi.SpanUpdateFirmware`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** | Collection ID  Collection ID for the collection owning the firmware image. | 
**imageId** | **string** | Firmware image ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSpanUpdateFirmwareRequest struct via the builder pattern


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

