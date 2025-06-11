# \FotaApi

All URIs are relative to *https://api.lab5e.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssignTargetImage**](FotaApi.md#AssignTargetImage) | **Patch** /span/collections/{collectionId}/firmware/labeled/{imageId}/assign | BETA: Assign a target image
[**ClearFirmwareError**](FotaApi.md#ClearFirmwareError) | **Delete** /span/collections/{collectionId}/devices/{deviceId}/fwerror | Clear FOTA error
[**CreateFirmware**](FotaApi.md#CreateFirmware) | **Post** /span/collections/{collectionId}/firmware | Create firmware
[**CreateLabeledFirmware**](FotaApi.md#CreateLabeledFirmware) | **Post** /span/collections/{collectionId}/firmware/labeled | BETA: Create a labeled firmware image
[**DeleteFirmware**](FotaApi.md#DeleteFirmware) | **Delete** /span/collections/{collectionId}/firmware/{imageId} | Delete firmware
[**DeleteLabeledImage**](FotaApi.md#DeleteLabeledImage) | **Delete** /span/collections/{collectionId}/firmware/labeled/{imageId} | BETA: Remove a tagged firmware image
[**FirmwareUsage**](FotaApi.md#FirmwareUsage) | **Get** /span/collections/{collectionId}/firmware/{imageId}/usage | Firmware usage
[**GetImageState**](FotaApi.md#GetImageState) | **Patch** /span/collections/{collectionId}/devices/{deviceId}/fotastate | BETA: Get state for a single device
[**ListFirmware**](FotaApi.md#ListFirmware) | **Get** /span/collections/{collectionId}/firmware | List firmware
[**ListImageStates**](FotaApi.md#ListImageStates) | **Get** /span/collections/{collectionId}/firmware/labeled/states | BETA: List image assignments plus states
[**ListImageVersionHistory**](FotaApi.md#ListImageVersionHistory) | **Patch** /span/collections/{collectionId}/devices/{deviceId}/fotalog | BETA: List version history for a single device
[**ListLabeledFirmware**](FotaApi.md#ListLabeledFirmware) | **Get** /span/collections/{collectionId}/firmware/labeled | BETA: List the labeled firmware images for a collection
[**RetrieveFirmware**](FotaApi.md#RetrieveFirmware) | **Get** /span/collections/{collectionId}/firmware/{imageId} | Retrieve firmware
[**RetrieveFirmwareStats**](FotaApi.md#RetrieveFirmwareStats) | **Get** /span/collections/{collectionId}/firmware/{imageId}/stats | Retrieve firmware statistics
[**UpdateFirmware**](FotaApi.md#UpdateFirmware) | **Patch** /span/collections/{existingCollectionId}/firmware/{imageId} | Update firmware



## AssignTargetImage

> AssignTargetImageResponse AssignTargetImage(ctx, collectionId, imageId).Label(label).DeviceId(deviceId).Execute()

BETA: Assign a target image



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
    label := "label_example" // string |  (optional)
    deviceId := "deviceId_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FotaApi.AssignTargetImage(context.Background(), collectionId, imageId).Label(label).DeviceId(deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FotaApi.AssignTargetImage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AssignTargetImage`: AssignTargetImageResponse
    fmt.Fprintf(os.Stdout, "Response from `FotaApi.AssignTargetImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** |  | 
**imageId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssignTargetImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **label** | **string** |  | 
 **deviceId** | **string** |  | 

### Return type

[**AssignTargetImageResponse**](AssignTargetImageResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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
    body := *openapiclient.NewCreateFirmwareBody() // CreateFirmwareBody | 

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

 **body** | [**CreateFirmwareBody**](CreateFirmwareBody.md) |  | 

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


## CreateLabeledFirmware

> CreateLabeledFirmwareResponse CreateLabeledFirmware(ctx, collectionId).Version(version).Label(label).ImageRef(imageRef).Execute()

BETA: Create a labeled firmware image



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
    version := "version_example" // string |  (optional)
    label := "label_example" // string |  (optional)
    imageRef := "imageRef_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FotaApi.CreateLabeledFirmware(context.Background(), collectionId).Version(version).Label(label).ImageRef(imageRef).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FotaApi.CreateLabeledFirmware``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateLabeledFirmware`: CreateLabeledFirmwareResponse
    fmt.Fprintf(os.Stdout, "Response from `FotaApi.CreateLabeledFirmware`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateLabeledFirmwareRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | **string** |  | 
 **label** | **string** |  | 
 **imageRef** | **string** |  | 

### Return type

[**CreateLabeledFirmwareResponse**](CreateLabeledFirmwareResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: Not defined
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


## DeleteLabeledImage

> map[string]interface{} DeleteLabeledImage(ctx, collectionId, imageId).Execute()

BETA: Remove a tagged firmware image



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
    resp, r, err := apiClient.FotaApi.DeleteLabeledImage(context.Background(), collectionId, imageId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FotaApi.DeleteLabeledImage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteLabeledImage`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `FotaApi.DeleteLabeledImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** |  | 
**imageId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLabeledImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]interface{}**

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


## GetImageState

> GetImageStateResponse GetImageState(ctx, collectionId, deviceId).Execute()

BETA: Get state for a single device

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
    resp, r, err := apiClient.FotaApi.GetImageState(context.Background(), collectionId, deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FotaApi.GetImageState``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetImageState`: GetImageStateResponse
    fmt.Fprintf(os.Stdout, "Response from `FotaApi.GetImageState`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** |  | 
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetImageStateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetImageStateResponse**](GetImageStateResponse.md)

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


## ListImageStates

> ListImageStatesResponse ListImageStates(ctx, collectionId).Execute()

BETA: List image assignments plus states



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
    resp, r, err := apiClient.FotaApi.ListImageStates(context.Background(), collectionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FotaApi.ListImageStates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListImageStates`: ListImageStatesResponse
    fmt.Fprintf(os.Stdout, "Response from `FotaApi.ListImageStates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListImageStatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListImageStatesResponse**](ListImageStatesResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListImageVersionHistory

> ListImageVersionHistoryResponse ListImageVersionHistory(ctx, collectionId, deviceId).Execute()

BETA: List version history for a single device

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
    resp, r, err := apiClient.FotaApi.ListImageVersionHistory(context.Background(), collectionId, deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FotaApi.ListImageVersionHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListImageVersionHistory`: ListImageVersionHistoryResponse
    fmt.Fprintf(os.Stdout, "Response from `FotaApi.ListImageVersionHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** |  | 
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListImageVersionHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ListImageVersionHistoryResponse**](ListImageVersionHistoryResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLabeledFirmware

> ListLabeledFirmwareResponse ListLabeledFirmware(ctx, collectionId).Execute()

BETA: List the labeled firmware images for a collection

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
    resp, r, err := apiClient.FotaApi.ListLabeledFirmware(context.Background(), collectionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FotaApi.ListLabeledFirmware``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLabeledFirmware`: ListLabeledFirmwareResponse
    fmt.Fprintf(os.Stdout, "Response from `FotaApi.ListLabeledFirmware`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListLabeledFirmwareRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListLabeledFirmwareResponse**](ListLabeledFirmwareResponse.md)

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
    body := *openapiclient.NewUpdateFirmwareBody() // UpdateFirmwareBody | 

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


 **body** | [**UpdateFirmwareBody**](UpdateFirmwareBody.md) |  | 

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

