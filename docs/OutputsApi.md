# \OutputsApi

All URIs are relative to *https://api.lab5e.com/span*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SpanCreateOutput**](OutputsApi.md#SpanCreateOutput) | **Post** /collections/{collectionId}/outputs | Create output
[**SpanDeleteOutput**](OutputsApi.md#SpanDeleteOutput) | **Delete** /collections/{collectionId}/outputs/{outputId} | Delete output
[**SpanListOutputs**](OutputsApi.md#SpanListOutputs) | **Get** /collections/{collectionId}/outputs | List outputs
[**SpanLogs**](OutputsApi.md#SpanLogs) | **Get** /collections/{collectionId}/outputs/{outputId}/logs | Output logs
[**SpanRetrieveOutput**](OutputsApi.md#SpanRetrieveOutput) | **Get** /collections/{collectionId}/outputs/{outputId} | Retrieve output
[**SpanStatus**](OutputsApi.md#SpanStatus) | **Get** /collections/{collectionId}/outputs/{outputId}/status | Output status
[**SpanUpdateOutput**](OutputsApi.md#SpanUpdateOutput) | **Patch** /collections/{collectionId}/outputs/{outputId} | Update output



## SpanCreateOutput

> Output SpanCreateOutput(ctx, collectionId).Body(body).Execute()

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
    resp, r, err := api_client.OutputsApi.SpanCreateOutput(context.Background(), collectionId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OutputsApi.SpanCreateOutput``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SpanCreateOutput`: Output
    fmt.Fprintf(os.Stdout, "Response from `OutputsApi.SpanCreateOutput`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSpanCreateOutputRequest struct via the builder pattern


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


## SpanDeleteOutput

> Output SpanDeleteOutput(ctx, collectionId, outputId).Execute()

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
    resp, r, err := api_client.OutputsApi.SpanDeleteOutput(context.Background(), collectionId, outputId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OutputsApi.SpanDeleteOutput``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SpanDeleteOutput`: Output
    fmt.Fprintf(os.Stdout, "Response from `OutputsApi.SpanDeleteOutput`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** |  | 
**outputId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSpanDeleteOutputRequest struct via the builder pattern


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


## SpanListOutputs

> ListOutputResponse SpanListOutputs(ctx, collectionId).Execute()

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
    resp, r, err := api_client.OutputsApi.SpanListOutputs(context.Background(), collectionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OutputsApi.SpanListOutputs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SpanListOutputs`: ListOutputResponse
    fmt.Fprintf(os.Stdout, "Response from `OutputsApi.SpanListOutputs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSpanListOutputsRequest struct via the builder pattern


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


## SpanLogs

> OutputLogResponse SpanLogs(ctx, collectionId, outputId).Execute()

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
    resp, r, err := api_client.OutputsApi.SpanLogs(context.Background(), collectionId, outputId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OutputsApi.SpanLogs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SpanLogs`: OutputLogResponse
    fmt.Fprintf(os.Stdout, "Response from `OutputsApi.SpanLogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** |  | 
**outputId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSpanLogsRequest struct via the builder pattern


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


## SpanRetrieveOutput

> Output SpanRetrieveOutput(ctx, collectionId, outputId).Execute()

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
    resp, r, err := api_client.OutputsApi.SpanRetrieveOutput(context.Background(), collectionId, outputId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OutputsApi.SpanRetrieveOutput``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SpanRetrieveOutput`: Output
    fmt.Fprintf(os.Stdout, "Response from `OutputsApi.SpanRetrieveOutput`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** |  | 
**outputId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSpanRetrieveOutputRequest struct via the builder pattern


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


## SpanStatus

> OutputStatusResponse SpanStatus(ctx, collectionId, outputId).Execute()

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
    resp, r, err := api_client.OutputsApi.SpanStatus(context.Background(), collectionId, outputId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OutputsApi.SpanStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SpanStatus`: OutputStatusResponse
    fmt.Fprintf(os.Stdout, "Response from `OutputsApi.SpanStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** |  | 
**outputId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSpanStatusRequest struct via the builder pattern


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


## SpanUpdateOutput

> Output SpanUpdateOutput(ctx, collectionId, outputId).Body(body).Execute()

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
    resp, r, err := api_client.OutputsApi.SpanUpdateOutput(context.Background(), collectionId, outputId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OutputsApi.SpanUpdateOutput``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SpanUpdateOutput`: Output
    fmt.Fprintf(os.Stdout, "Response from `OutputsApi.SpanUpdateOutput`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** |  | 
**outputId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSpanUpdateOutputRequest struct via the builder pattern


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

