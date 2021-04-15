# \CollectionsApi

All URIs are relative to *https://api.lab5e.com/span*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SpanBroadcastMessage**](CollectionsApi.md#SpanBroadcastMessage) | **Post** /collections/{collectionId}/to | Broadcast message
[**SpanCreateCollection**](CollectionsApi.md#SpanCreateCollection) | **Post** /collections | Create collection
[**SpanDeleteCollection**](CollectionsApi.md#SpanDeleteCollection) | **Delete** /collections/{collectionId} | Delete collection
[**SpanListCollectionData**](CollectionsApi.md#SpanListCollectionData) | **Get** /collections/{collectionId}/data | Get payloads
[**SpanListCollections**](CollectionsApi.md#SpanListCollections) | **Get** /collections | List collections
[**SpanRetrieveCollection**](CollectionsApi.md#SpanRetrieveCollection) | **Get** /collections/{collectionId} | Retrieve collection
[**SpanUpdateCollection**](CollectionsApi.md#SpanUpdateCollection) | **Patch** /collections/{collectionId} | Update collection



## SpanBroadcastMessage

> MultiSendMessageResponse SpanBroadcastMessage(ctx, collectionId).Body(body).Execute()

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
    resp, r, err := api_client.CollectionsApi.SpanBroadcastMessage(context.Background(), collectionId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CollectionsApi.SpanBroadcastMessage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SpanBroadcastMessage`: MultiSendMessageResponse
    fmt.Fprintf(os.Stdout, "Response from `CollectionsApi.SpanBroadcastMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSpanBroadcastMessageRequest struct via the builder pattern


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


## SpanCreateCollection

> Collection SpanCreateCollection(ctx).Body(body).Execute()

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
    resp, r, err := api_client.CollectionsApi.SpanCreateCollection(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CollectionsApi.SpanCreateCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SpanCreateCollection`: Collection
    fmt.Fprintf(os.Stdout, "Response from `CollectionsApi.SpanCreateCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSpanCreateCollectionRequest struct via the builder pattern


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


## SpanDeleteCollection

> Collection SpanDeleteCollection(ctx, collectionId).Execute()

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
    resp, r, err := api_client.CollectionsApi.SpanDeleteCollection(context.Background(), collectionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CollectionsApi.SpanDeleteCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SpanDeleteCollection`: Collection
    fmt.Fprintf(os.Stdout, "Response from `CollectionsApi.SpanDeleteCollection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** | The ID of the collection you want to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiSpanDeleteCollectionRequest struct via the builder pattern


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


## SpanListCollectionData

> ListDataResponse SpanListCollectionData(ctx, collectionId).Limit(limit).Start(start).End(end).Offset(offset).Execute()

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
    resp, r, err := api_client.CollectionsApi.SpanListCollectionData(context.Background(), collectionId).Limit(limit).Start(start).End(end).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CollectionsApi.SpanListCollectionData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SpanListCollectionData`: ListDataResponse
    fmt.Fprintf(os.Stdout, "Response from `CollectionsApi.SpanListCollectionData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** | The collection ID requested. This is included in the request path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSpanListCollectionDataRequest struct via the builder pattern


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


## SpanListCollections

> ListCollectionResponse SpanListCollections(ctx).Execute()

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
    resp, r, err := api_client.CollectionsApi.SpanListCollections(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CollectionsApi.SpanListCollections``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SpanListCollections`: ListCollectionResponse
    fmt.Fprintf(os.Stdout, "Response from `CollectionsApi.SpanListCollections`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSpanListCollectionsRequest struct via the builder pattern


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


## SpanRetrieveCollection

> Collection SpanRetrieveCollection(ctx, collectionId).Execute()

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
    resp, r, err := api_client.CollectionsApi.SpanRetrieveCollection(context.Background(), collectionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CollectionsApi.SpanRetrieveCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SpanRetrieveCollection`: Collection
    fmt.Fprintf(os.Stdout, "Response from `CollectionsApi.SpanRetrieveCollection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** | The collection ID of the collection you are requesting | 

### Other Parameters

Other parameters are passed through a pointer to a apiSpanRetrieveCollectionRequest struct via the builder pattern


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


## SpanUpdateCollection

> Collection SpanUpdateCollection(ctx, collectionId).Body(body).Execute()

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
    resp, r, err := api_client.CollectionsApi.SpanUpdateCollection(context.Background(), collectionId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CollectionsApi.SpanUpdateCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SpanUpdateCollection`: Collection
    fmt.Fprintf(os.Stdout, "Response from `CollectionsApi.SpanUpdateCollection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** | The ID of the collection. This is assigned by the backend. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSpanUpdateCollectionRequest struct via the builder pattern


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

