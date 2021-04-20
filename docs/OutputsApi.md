# \OutputsApi

All URIs are relative to *https://api.lab5e.com/span*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOutput**](OutputsApi.md#CreateOutput) | **Post** /collections/{collectionId}/outputs | Create output
[**DeleteOutput**](OutputsApi.md#DeleteOutput) | **Delete** /collections/{collectionId}/outputs/{outputId} | Delete output
[**ListOutputs**](OutputsApi.md#ListOutputs) | **Get** /collections/{collectionId}/outputs | List outputs
[**Logs**](OutputsApi.md#Logs) | **Get** /collections/{collectionId}/outputs/{outputId}/logs | Output logs
[**RetrieveOutput**](OutputsApi.md#RetrieveOutput) | **Get** /collections/{collectionId}/outputs/{outputId} | Retrieve output
[**Status**](OutputsApi.md#Status) | **Get** /collections/{collectionId}/outputs/{outputId}/status | Output status
[**UpdateOutput**](OutputsApi.md#UpdateOutput) | **Patch** /collections/{collectionId}/outputs/{outputId} | Update output



## CreateOutput

> Output CreateOutput(ctx, collectionId).Body(body).Execute()

Create output

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
    body := *openapiclient.NewOutput() // Output | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OutputsApi.CreateOutput(context.Background(), collectionId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OutputsApi.CreateOutput``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOutput`: Output
    fmt.Fprintf(os.Stdout, "Response from `OutputsApi.CreateOutput`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOutputRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**Output**](Output.md) |  | 

### Return type

[**Output**](Output.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOutput

> Output DeleteOutput(ctx, collectionId, outputId).Execute()

Delete output

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
    outputId := "outputId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OutputsApi.DeleteOutput(context.Background(), collectionId, outputId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OutputsApi.DeleteOutput``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteOutput`: Output
    fmt.Fprintf(os.Stdout, "Response from `OutputsApi.DeleteOutput`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** |  | 
**outputId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOutputRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Output**](Output.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOutputs

> ListOutputResponse ListOutputs(ctx, collectionId).Execute()

List outputs

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
    resp, r, err := api_client.OutputsApi.ListOutputs(context.Background(), collectionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OutputsApi.ListOutputs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListOutputs`: ListOutputResponse
    fmt.Fprintf(os.Stdout, "Response from `OutputsApi.ListOutputs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOutputsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListOutputResponse**](ListOutputResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Logs

> OutputLogResponse Logs(ctx, collectionId, outputId).Execute()

Output logs

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
    outputId := "outputId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OutputsApi.Logs(context.Background(), collectionId, outputId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OutputsApi.Logs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Logs`: OutputLogResponse
    fmt.Fprintf(os.Stdout, "Response from `OutputsApi.Logs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** |  | 
**outputId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OutputLogResponse**](OutputLogResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveOutput

> Output RetrieveOutput(ctx, collectionId, outputId).Execute()

Retrieve output

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
    outputId := "outputId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OutputsApi.RetrieveOutput(context.Background(), collectionId, outputId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OutputsApi.RetrieveOutput``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveOutput`: Output
    fmt.Fprintf(os.Stdout, "Response from `OutputsApi.RetrieveOutput`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** |  | 
**outputId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveOutputRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Output**](Output.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Status

> OutputStatusResponse Status(ctx, collectionId, outputId).Execute()

Output status

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
    outputId := "outputId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OutputsApi.Status(context.Background(), collectionId, outputId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OutputsApi.Status``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Status`: OutputStatusResponse
    fmt.Fprintf(os.Stdout, "Response from `OutputsApi.Status`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** |  | 
**outputId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OutputStatusResponse**](OutputStatusResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOutput

> Output UpdateOutput(ctx, collectionId, outputId).Body(body).Execute()

Update output



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
    outputId := "outputId_example" // string | 
    body := *openapiclient.NewOutput() // Output | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OutputsApi.UpdateOutput(context.Background(), collectionId, outputId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OutputsApi.UpdateOutput``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOutput`: Output
    fmt.Fprintf(os.Stdout, "Response from `OutputsApi.UpdateOutput`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** |  | 
**outputId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOutputRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**Output**](Output.md) |  | 

### Return type

[**Output**](Output.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

