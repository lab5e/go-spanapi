# \CollectionsApi

All URIs are relative to *https://api.lab5e.com/span*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BroadcastMessage**](CollectionsApi.md#BroadcastMessage) | **Post** /collections/{collectionId}/to | Broadcast message
[**CreateCollection**](CollectionsApi.md#CreateCollection) | **Post** /collections | Create collection
[**DeleteCollection**](CollectionsApi.md#DeleteCollection) | **Delete** /collections/{collectionId} | Delete collection
[**ListCollectionData**](CollectionsApi.md#ListCollectionData) | **Get** /collections/{collectionId}/data | Get payloads
[**ListCollections**](CollectionsApi.md#ListCollections) | **Get** /collections | List collections
[**RetrieveCollection**](CollectionsApi.md#RetrieveCollection) | **Get** /collections/{collectionId} | Retrieve collection
[**UpdateCollection**](CollectionsApi.md#UpdateCollection) | **Patch** /collections/{collectionId} | Update collection



## BroadcastMessage

> MultiSendMessageResponse BroadcastMessage(ctx, collectionId).Body(body).Execute()

Broadcast message



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
    body := *openapiclient.NewBroadcastMessageRequest() // BroadcastMessageRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CollectionsApi.BroadcastMessage(context.Background(), collectionId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CollectionsApi.BroadcastMessage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BroadcastMessage`: MultiSendMessageResponse
    fmt.Fprintf(os.Stdout, "Response from `CollectionsApi.BroadcastMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBroadcastMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**BroadcastMessageRequest**](BroadcastMessageRequest.md) |  | 

### Return type

[**MultiSendMessageResponse**](MultiSendMessageResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCollection

> Collection CreateCollection(ctx).Body(body).Execute()

Create collection



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
    body := *openapiclient.NewCollection() // Collection | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CollectionsApi.CreateCollection(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CollectionsApi.CreateCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCollection`: Collection
    fmt.Fprintf(os.Stdout, "Response from `CollectionsApi.CreateCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**Collection**](Collection.md) |  | 

### Return type

[**Collection**](Collection.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCollection

> Collection DeleteCollection(ctx, collectionId).Execute()

Delete collection



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
    collectionId := "collectionId_example" // string | The ID of the collection you want to delete

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CollectionsApi.DeleteCollection(context.Background(), collectionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CollectionsApi.DeleteCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteCollection`: Collection
    fmt.Fprintf(os.Stdout, "Response from `CollectionsApi.DeleteCollection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** | The ID of the collection you want to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Collection**](Collection.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCollectionData

> ListDataResponse ListCollectionData(ctx, collectionId).Limit(limit).Start(start).End(end).Offset(offset).Execute()

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
    collectionId := "collectionId_example" // string | The collection ID requested. This is included in the request path.
    limit := int32(56) // int32 | Limit the number of payloads to return. The default is 512. (optional)
    start := "start_example" // string | Start of time range. The default is 24 hours ago. Value is in milliseconds since epoch. (optional)
    end := "end_example" // string | End of time range. The default is the current time stamp. Value is in milliseconds since epoch. (optional)
    offset := "offset_example" // string | The message offset based on the message ID. This parameter can't be combined with the start and end parameters. If no parameter is set the first N messages will be returned. If this parameter is set the next N messages (from newest to oldest) with message ID less than the offset will be returned. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CollectionsApi.ListCollectionData(context.Background(), collectionId).Limit(limit).Start(start).End(end).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CollectionsApi.ListCollectionData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCollectionData`: ListDataResponse
    fmt.Fprintf(os.Stdout, "Response from `CollectionsApi.ListCollectionData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** | The collection ID requested. This is included in the request path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCollectionDataRequest struct via the builder pattern


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


## ListCollections

> ListCollectionResponse ListCollections(ctx).Execute()

List collections



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CollectionsApi.ListCollections(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CollectionsApi.ListCollections``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCollections`: ListCollectionResponse
    fmt.Fprintf(os.Stdout, "Response from `CollectionsApi.ListCollections`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListCollectionsRequest struct via the builder pattern


### Return type

[**ListCollectionResponse**](ListCollectionResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveCollection

> Collection RetrieveCollection(ctx, collectionId).Execute()

Retrieve collection

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
    collectionId := "collectionId_example" // string | The collection ID of the collection you are requesting

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CollectionsApi.RetrieveCollection(context.Background(), collectionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CollectionsApi.RetrieveCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveCollection`: Collection
    fmt.Fprintf(os.Stdout, "Response from `CollectionsApi.RetrieveCollection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** | The collection ID of the collection you are requesting | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Collection**](Collection.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCollection

> Collection UpdateCollection(ctx, collectionId).Body(body).Execute()

Update collection



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
    collectionId := "collectionId_example" // string | The ID of the collection. This is assigned by the backend.
    body := *openapiclient.NewCollection() // Collection | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CollectionsApi.UpdateCollection(context.Background(), collectionId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CollectionsApi.UpdateCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCollection`: Collection
    fmt.Fprintf(os.Stdout, "Response from `CollectionsApi.UpdateCollection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** | The ID of the collection. This is assigned by the backend. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**Collection**](Collection.md) |  | 

### Return type

[**Collection**](Collection.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

