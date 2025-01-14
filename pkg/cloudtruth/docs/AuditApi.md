# \AuditApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuditList**](AuditApi.md#AuditList) | **Get** /api/v1/audit/ | 
[**AuditRetrieve**](AuditApi.md#AuditRetrieve) | **Get** /api/v1/audit/{id}/ | 
[**AuditSummaryRetrieve**](AuditApi.md#AuditSummaryRetrieve) | **Get** /api/v1/audit/summary/ | 



## AuditList

> PaginatedAuditTrailList AuditList(ctx).Action(action).Earliest(earliest).Latest(latest).ObjectId(objectId).ObjectType(objectType).Ordering(ordering).Page(page).PageSize(pageSize).UserId(userId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    action := "action_example" // string | The action that was taken. (optional)
    earliest := time.Now() // time.Time |  (optional)
    latest := time.Now() // time.Time |  (optional)
    objectId := "objectId_example" // string |  (optional)
    objectType := "objectType_example" // string |  (optional)
    ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)
    userId := "userId_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuditApi.AuditList(context.Background()).Action(action).Earliest(earliest).Latest(latest).ObjectId(objectId).ObjectType(objectType).Ordering(ordering).Page(page).PageSize(pageSize).UserId(userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuditApi.AuditList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuditList`: PaginatedAuditTrailList
    fmt.Fprintf(os.Stdout, "Response from `AuditApi.AuditList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuditListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **action** | **string** | The action that was taken. | 
 **earliest** | **time.Time** |  | 
 **latest** | **time.Time** |  | 
 **objectId** | **string** |  | 
 **objectType** | **string** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **userId** | **string** |  | 

### Return type

[**PaginatedAuditTrailList**](PaginatedAuditTrailList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuditRetrieve

> AuditTrail AuditRetrieve(ctx, id).Execute()





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
    id := TODO // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuditApi.AuditRetrieve(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuditApi.AuditRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuditRetrieve`: AuditTrail
    fmt.Fprintf(os.Stdout, "Response from `AuditApi.AuditRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuditRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AuditTrail**](AuditTrail.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuditSummaryRetrieve

> AuditTrailSummary AuditSummaryRetrieve(ctx).Execute()





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
    resp, r, err := api_client.AuditApi.AuditSummaryRetrieve(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuditApi.AuditSummaryRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuditSummaryRetrieve`: AuditTrailSummary
    fmt.Fprintf(os.Stdout, "Response from `AuditApi.AuditSummaryRetrieve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAuditSummaryRetrieveRequest struct via the builder pattern


### Return type

[**AuditTrailSummary**](AuditTrailSummary.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

