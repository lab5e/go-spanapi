# \DevicesApi

All URIs are relative to *https://api.lab5e.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddDownstreamMessage**](DevicesApi.md#AddDownstreamMessage) | **Post** /span/collections/{collectionId}/devices/{deviceId}/outbox | Add message to oubox
[**CreateDevice**](DevicesApi.md#CreateDevice) | **Post** /span/collections/{collectionId}/devices | Create device
[**DeleteDevice**](DevicesApi.md#DeleteDevice) | **Delete** /span/collections/{collectionId}/devices/{deviceId} | Remove device.
[**DeleteDownstreamMessage**](DevicesApi.md#DeleteDownstreamMessage) | **Delete** /span/collections/{collectionId}/devices/{deviceId}/outbox/{messageId} | Delete outgoing message
[**DeviceCertificate**](DevicesApi.md#DeviceCertificate) | **Get** /span/collections/{collectionId}/devices/{deviceId}/certs | Get issued certificate(s) for device
[**ListDeviceData**](DevicesApi.md#ListDeviceData) | **Get** /span/collections/{collectionId}/devices/{deviceId}/data | Retrieve data from device
[**ListDevices**](DevicesApi.md#ListDevices) | **Get** /span/collections/{collectionId}/devices | List devices in collection.
[**ListDownstreamMessages**](DevicesApi.md#ListDownstreamMessages) | **Get** /span/collections/{collectionId}/devices/{deviceId}/outbox | List the messages in the outbox
[**ListUpstreamMessages**](DevicesApi.md#ListUpstreamMessages) | **Get** /span/collections/{collectionId}/devices/{deviceId}/inbox | List incoming messages
[**RetrieveDevice**](DevicesApi.md#RetrieveDevice) | **Get** /span/collections/{collectionId}/devices/{deviceId} | Retrieve device
[**UpdateDevice**](DevicesApi.md#UpdateDevice) | **Patch** /span/collections/{existingCollectionId}/devices/{deviceId} | Update device



## AddDownstreamMessage

> MessageDownstream AddDownstreamMessage(ctx, collectionId, deviceId).Body(body).Execute()

Add message to oubox



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
    body := *openapiclient.NewAddDownstreamMessageRequest() // AddDownstreamMessageRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.AddDownstreamMessage(context.Background(), collectionId, deviceId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.AddDownstreamMessage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddDownstreamMessage`: MessageDownstream
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.AddDownstreamMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** |  | 
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddDownstreamMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**AddDownstreamMessageRequest**](AddDownstreamMessageRequest.md) |  | 

### Return type

[**MessageDownstream**](MessageDownstream.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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
    openapiclient "github.com/lab5e/go-spanapi"
)

func main() {
    collectionId := "collectionId_example" // string | This is the containing collection
    body := *openapiclient.NewCreateDeviceRequest() // CreateDeviceRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.CreateDevice(context.Background(), collectionId).Body(body).Execute()
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

 **body** | [**CreateDeviceRequest**](CreateDeviceRequest.md) |  | 

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

Remove device.

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
    collectionId := "collectionId_example" // string | This is the containing collection
    deviceId := "deviceId_example" // string | The device ID is assigned by the backend.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.DeleteDevice(context.Background(), collectionId, deviceId).Execute()
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
**collectionId** | **string** | This is the containing collection | 
**deviceId** | **string** | The device ID is assigned by the backend. | 

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


## DeleteDownstreamMessage

> DeleteDownstreamMessageResponse DeleteDownstreamMessage(ctx, collectionId, deviceId, messageId).Execute()

Delete outgoing message



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
    messageId := "messageId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.DeleteDownstreamMessage(context.Background(), collectionId, deviceId, messageId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.DeleteDownstreamMessage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteDownstreamMessage`: DeleteDownstreamMessageResponse
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.DeleteDownstreamMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** |  | 
**deviceId** | **string** |  | 
**messageId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDownstreamMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**DeleteDownstreamMessageResponse**](DeleteDownstreamMessageResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceCertificate

> DeviceCertificateResponse DeviceCertificate(ctx, collectionId, deviceId).Execute()

Get issued certificate(s) for device

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
    resp, r, err := apiClient.DevicesApi.DeviceCertificate(context.Background(), collectionId, deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.DeviceCertificate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceCertificate`: DeviceCertificateResponse
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.DeviceCertificate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** |  | 
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceCertificateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DeviceCertificateResponse**](DeviceCertificateResponse.md)

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

Retrieve data from device



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
    collectionId := "collectionId_example" // string | The collection ID. This is included in the request path.
    deviceId := "deviceId_example" // string | The device ID. This is included in the request path.
    limit := int32(56) // int32 | Limit the number of payloads to return. The default is 512. (optional)
    start := "start_example" // string | Start of time range. The default is 24 hours ago. Value is in milliseconds since epoch. (optional)
    end := "end_example" // string | End of time range. The default is the current time stamp. Value is in milliseconds since epoch. (optional)
    offset := "offset_example" // string | The message offset based on the message ID. This parameter can't be combined with the start and end parameters. If no parameter is set the first N messages will be returned. If this parameter is set the next N messages (from newest to oldest) with message ID less than the offset will be returned. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.ListDeviceData(context.Background(), collectionId, deviceId).Limit(limit).Start(start).End(end).Offset(offset).Execute()
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

List devices in collection.

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
    resp, r, err := apiClient.DevicesApi.ListDevices(context.Background(), collectionId).Execute()
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


## ListDownstreamMessages

> ListDownstreamMessagesResponse ListDownstreamMessages(ctx, collectionId, deviceId).Limit(limit).Start(start).End(end).Offset(offset).Execute()

List the messages in the outbox



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
    limit := int32(56) // int32 |  (optional)
    start := "start_example" // string | Start of time range. The default is 24 hours ago. Value is in milliseconds since epoch. (optional)
    end := "end_example" // string | End of time range. The default is the current time stamp. Value is in milliseconds since epoch. (optional)
    offset := "offset_example" // string | The message offset based on the message ID. This parameter can't be combined with the start and end parameters. If no parameter is set the first N messages will be returned. If this parameter is set the next N messages (from newest to oldest) with message ID less than the offset will be returned. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.ListDownstreamMessages(context.Background(), collectionId, deviceId).Limit(limit).Start(start).End(end).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.ListDownstreamMessages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDownstreamMessages`: ListDownstreamMessagesResponse
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.ListDownstreamMessages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** |  | 
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListDownstreamMessagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **int32** |  | 
 **start** | **string** | Start of time range. The default is 24 hours ago. Value is in milliseconds since epoch. | 
 **end** | **string** | End of time range. The default is the current time stamp. Value is in milliseconds since epoch. | 
 **offset** | **string** | The message offset based on the message ID. This parameter can&#39;t be combined with the start and end parameters. If no parameter is set the first N messages will be returned. If this parameter is set the next N messages (from newest to oldest) with message ID less than the offset will be returned. | 

### Return type

[**ListDownstreamMessagesResponse**](ListDownstreamMessagesResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUpstreamMessages

> ListUpstreamMessagesResponse ListUpstreamMessages(ctx, collectionId, deviceId).Limit(limit).Start(start).End(end).Offset(offset).Execute()

List incoming messages



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
    limit := int32(56) // int32 |  (optional)
    start := "start_example" // string | Start of time range. The default is 24 hours ago. Value is in milliseconds since epoch. (optional)
    end := "end_example" // string | End of time range. The default is the current time stamp. Value is in milliseconds since epoch. (optional)
    offset := "offset_example" // string | The message offset based on the message ID. This parameter can't be combined with the start and end parameters. If no parameter is set the first N messages will be returned. If this parameter is set the next N messages (from newest to oldest) with message ID less than the offset will be returned. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.ListUpstreamMessages(context.Background(), collectionId, deviceId).Limit(limit).Start(start).End(end).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.ListUpstreamMessages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUpstreamMessages`: ListUpstreamMessagesResponse
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.ListUpstreamMessages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** |  | 
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListUpstreamMessagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **int32** |  | 
 **start** | **string** | Start of time range. The default is 24 hours ago. Value is in milliseconds since epoch. | 
 **end** | **string** | End of time range. The default is the current time stamp. Value is in milliseconds since epoch. | 
 **offset** | **string** | The message offset based on the message ID. This parameter can&#39;t be combined with the start and end parameters. If no parameter is set the first N messages will be returned. If this parameter is set the next N messages (from newest to oldest) with message ID less than the offset will be returned. | 

### Return type

[**ListUpstreamMessagesResponse**](ListUpstreamMessagesResponse.md)

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
    openapiclient "github.com/lab5e/go-spanapi"
)

func main() {
    collectionId := "collectionId_example" // string | This is the containing collection
    deviceId := "deviceId_example" // string | The device ID is assigned by the backend.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.RetrieveDevice(context.Background(), collectionId, deviceId).Execute()
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
**collectionId** | **string** | This is the containing collection | 
**deviceId** | **string** | The device ID is assigned by the backend. | 

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
    openapiclient "github.com/lab5e/go-spanapi"
)

func main() {
    existingCollectionId := "existingCollectionId_example" // string | 
    deviceId := "deviceId_example" // string | 
    body := *openapiclient.NewUpdateDeviceRequest() // UpdateDeviceRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.UpdateDevice(context.Background(), existingCollectionId, deviceId).Body(body).Execute()
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

