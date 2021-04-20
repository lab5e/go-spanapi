# \DevicesApi

All URIs are relative to *https://api.lab5e.com/span*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDevice**](DevicesApi.md#CreateDevice) | **Post** /collections/{collectionId}/devices | Create device
[**DeleteDevice**](DevicesApi.md#DeleteDevice) | **Delete** /collections/{collectionId}/devices/{deviceId} | Remove device
[**ListDeviceData**](DevicesApi.md#ListDeviceData) | **Get** /collections/{collectionId}/devices/{deviceId}/data | Get payloads
[**ListDevices**](DevicesApi.md#ListDevices) | **Get** /collections/{collectionId}/devices | List devices
[**RetrieveDevice**](DevicesApi.md#RetrieveDevice) | **Get** /collections/{collectionId}/devices/{deviceId} | Retrieve device
[**SendMessage**](DevicesApi.md#SendMessage) | **Post** /collections/{collectionId}/devices/{deviceId}/to | Send message
[**UpdateDevice**](DevicesApi.md#UpdateDevice) | **Patch** /collections/{existingCollectionId}/devices/{deviceId} | Update device



## CreateDevice

> Device CreateDevice(ctx, collectionId).Body(body).Execute()

Create device



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
    collectionId := "collectionId_example" // string | This is the containing collection
    body := *openapiclient.NewDevice() // Device | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DevicesApi.CreateDevice(context.Background(), collectionId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.CreateDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDevice`: Device
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.CreateDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** | This is the containing collection | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**Device**](Device.md) |  | 

### Return type

[**Device**](Device.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDevice

> Device DeleteDevice(ctx, collectionId, deviceId).Execute()

Remove device

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
    resp, r, err := api_client.DevicesApi.DeleteDevice(context.Background(), collectionId, deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.DeleteDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteDevice`: Device
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.DeleteDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** |  | 
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Device**](Device.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDeviceData

> ListDataResponse ListDeviceData(ctx, collectionId, deviceId).Limit(limit).Start(start).End(end).Offset(offset).Execute()

Get payloads



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
    collectionId := "collectionId_example" // string | The collection ID. This is included in the request path.
    deviceId := "deviceId_example" // string | The device ID. This is included in the request path.
    limit := int32(56) // int32 | Limit the number of payloads to return. The default is 512. (optional)
    start := "start_example" // string | Start of time range. The default is 24 hours ago. Value is in milliseconds since epoch. (optional)
    end := "end_example" // string | End of time range. The default is the current time stamp. Value is in milliseconds since epoch. (optional)
    offset := "offset_example" // string | The message offset based on the message ID. This parameter can't be combined with the start and end parameters. If no parameter is set the first N messages will be returned. If this parameter is set the next N messages (from newest to oldest) with message ID less than the offset will be returned. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DevicesApi.ListDeviceData(context.Background(), collectionId, deviceId).Limit(limit).Start(start).End(end).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.ListDeviceData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDeviceData`: ListDataResponse
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.ListDeviceData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** | The collection ID. This is included in the request path. | 
**deviceId** | **string** | The device ID. This is included in the request path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListDeviceDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **int32** | Limit the number of payloads to return. The default is 512. | 
 **start** | **string** | Start of time range. The default is 24 hours ago. Value is in milliseconds since epoch. | 
 **end** | **string** | End of time range. The default is the current time stamp. Value is in milliseconds since epoch. | 
 **offset** | **string** | The message offset based on the message ID. This parameter can&#39;t be combined with the start and end parameters. If no parameter is set the first N messages will be returned. If this parameter is set the next N messages (from newest to oldest) with message ID less than the offset will be returned. | 

### Return type

[**ListDataResponse**](ListDataResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDevices

> ListDevicesResponse ListDevices(ctx, collectionId).Execute()

List devices

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
    resp, r, err := api_client.DevicesApi.ListDevices(context.Background(), collectionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.ListDevices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDevices`: ListDevicesResponse
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.ListDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListDevicesResponse**](ListDevicesResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveDevice

> Device RetrieveDevice(ctx, collectionId, deviceId).Execute()

Retrieve device



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
    resp, r, err := api_client.DevicesApi.RetrieveDevice(context.Background(), collectionId, deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.RetrieveDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveDevice`: Device
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.RetrieveDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** |  | 
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Device**](Device.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendMessage

> SendMessageResponse SendMessage(ctx, collectionId, deviceId).Body(body).Execute()

Send message



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
    body := *openapiclient.NewSendMessageRequest() // SendMessageRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DevicesApi.SendMessage(context.Background(), collectionId, deviceId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.SendMessage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SendMessage`: SendMessageResponse
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.SendMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** |  | 
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSendMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**SendMessageRequest**](SendMessageRequest.md) |  | 

### Return type

[**SendMessageResponse**](SendMessageResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDevice

> Device UpdateDevice(ctx, existingCollectionId, deviceId).Body(body).Execute()

Update device

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
    existingCollectionId := "existingCollectionId_example" // string | 
    deviceId := "deviceId_example" // string | 
    body := *openapiclient.NewUpdateDeviceRequest() // UpdateDeviceRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DevicesApi.UpdateDevice(context.Background(), existingCollectionId, deviceId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.UpdateDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDevice`: Device
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.UpdateDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**existingCollectionId** | **string** |  | 
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**UpdateDeviceRequest**](UpdateDeviceRequest.md) |  | 

### Return type

[**Device**](Device.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

