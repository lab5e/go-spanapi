# \BlobsApi

All URIs are relative to *https://api.lab5e.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteBlob**](BlobsApi.md#DeleteBlob) | **Delete** /span/collections/{collectionId}/blobs/{blobId} | Remove a blob stored on a collection
[**ListBlobs**](BlobsApi.md#ListBlobs) | **Get** /span/collections/{collectionId}/blobs | List the blobs for a collection



## DeleteBlob

> map[string]interface{} DeleteBlob(ctx, collectionId, blobId).Execute()

Remove a blob stored on a collection



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
    blobId := "blobId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BlobsApi.DeleteBlob(context.Background(), collectionId, blobId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlobsApi.DeleteBlob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteBlob`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `BlobsApi.DeleteBlob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** |  | 
**blobId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBlobRequest struct via the builder pattern


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


## ListBlobs

> ListBlobResponse ListBlobs(ctx, collectionId).Limit(limit).Execute()

List the blobs for a collection



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
    limit := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BlobsApi.ListBlobs(context.Background(), collectionId).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlobsApi.ListBlobs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBlobs`: ListBlobResponse
    fmt.Fprintf(os.Stdout, "Response from `BlobsApi.ListBlobs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListBlobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | 

### Return type

[**ListBlobResponse**](ListBlobResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

