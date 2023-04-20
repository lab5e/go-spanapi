# \GatewaysApi

All URIs are relative to *https://api.lab5e.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateGateway**](GatewaysApi.md#CreateGateway) | **Post** /span/collections/{collectionId}/gateways | Create gateway
[**DeleteGateway**](GatewaysApi.md#DeleteGateway) | **Delete** /span/collections/{collectionId}/gateways/{gatewayId} | Delete gateway
[**GatewayCertificates**](GatewaysApi.md#GatewayCertificates) | **Get** /span/collections/{collectionId}/gateways/{gatewayId}/certs | Get issued certificate(s) for gateway
[**ListGateways**](GatewaysApi.md#ListGateways) | **Get** /span/collections/{collectionId}/gateways | List gateways
[**RetrieveGateway**](GatewaysApi.md#RetrieveGateway) | **Get** /span/collections/{collectionId}/gateways/{gatewayId} | Retrieve gateway
[**UpdateGateway**](GatewaysApi.md#UpdateGateway) | **Patch** /span/collections/{existingCollectionId}/gateways/{gatewayId} | Update gateway



## CreateGateway

> Gateway CreateGateway(ctx, collectionId).Body(body).Execute()

Create gateway



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
    body := *openapiclient.NewCreateGatewayRequest() // CreateGatewayRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GatewaysApi.CreateGateway(context.Background(), collectionId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GatewaysApi.CreateGateway``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateGateway`: Gateway
    fmt.Fprintf(os.Stdout, "Response from `GatewaysApi.CreateGateway`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateGatewayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**CreateGatewayRequest**](CreateGatewayRequest.md) |  | 

### Return type

[**Gateway**](Gateway.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGateway

> Gateway DeleteGateway(ctx, collectionId, gatewayId).Execute()

Delete gateway



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
    gatewayId := "gatewayId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GatewaysApi.DeleteGateway(context.Background(), collectionId, gatewayId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GatewaysApi.DeleteGateway``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteGateway`: Gateway
    fmt.Fprintf(os.Stdout, "Response from `GatewaysApi.DeleteGateway`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** |  | 
**gatewayId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGatewayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Gateway**](Gateway.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GatewayCertificates

> GatewayCertificateResponse GatewayCertificates(ctx, collectionId, gatewayId).Execute()

Get issued certificate(s) for gateway

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
    gatewayId := "gatewayId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GatewaysApi.GatewayCertificates(context.Background(), collectionId, gatewayId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GatewaysApi.GatewayCertificates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GatewayCertificates`: GatewayCertificateResponse
    fmt.Fprintf(os.Stdout, "Response from `GatewaysApi.GatewayCertificates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** |  | 
**gatewayId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGatewayCertificatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GatewayCertificateResponse**](GatewayCertificateResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGateways

> ListGatewayResponse ListGateways(ctx, collectionId).Execute()

List gateways



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
    resp, r, err := apiClient.GatewaysApi.ListGateways(context.Background(), collectionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GatewaysApi.ListGateways``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListGateways`: ListGatewayResponse
    fmt.Fprintf(os.Stdout, "Response from `GatewaysApi.ListGateways`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListGatewaysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListGatewayResponse**](ListGatewayResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveGateway

> Gateway RetrieveGateway(ctx, collectionId, gatewayId).Execute()

Retrieve gateway



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
    gatewayId := "gatewayId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GatewaysApi.RetrieveGateway(context.Background(), collectionId, gatewayId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GatewaysApi.RetrieveGateway``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveGateway`: Gateway
    fmt.Fprintf(os.Stdout, "Response from `GatewaysApi.RetrieveGateway`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** |  | 
**gatewayId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveGatewayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Gateway**](Gateway.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGateway

> Gateway UpdateGateway(ctx, existingCollectionId, gatewayId).Body(body).Execute()

Update gateway



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
    gatewayId := "gatewayId_example" // string | 
    body := *openapiclient.NewUpdateGatewayRequest() // UpdateGatewayRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GatewaysApi.UpdateGateway(context.Background(), existingCollectionId, gatewayId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GatewaysApi.UpdateGateway``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateGateway`: Gateway
    fmt.Fprintf(os.Stdout, "Response from `GatewaysApi.UpdateGateway`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**existingCollectionId** | **string** |  | 
**gatewayId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGatewayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**UpdateGatewayRequest**](UpdateGatewayRequest.md) |  | 

### Return type

[**Gateway**](Gateway.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

