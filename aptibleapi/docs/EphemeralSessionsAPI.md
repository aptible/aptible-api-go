# \EphemeralSessionsAPI

All URIs are relative to *https://api.aptible.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEphemeralSession**](EphemeralSessionsAPI.md#GetEphemeralSession) | **Get** /ephemeral_sessions/{id} | show ephemeral_session
[**ListEphemeralSessionsForApp**](EphemeralSessionsAPI.md#ListEphemeralSessionsForApp) | **Get** /apps/{app_id}/ephemeral_sessions | list ephemeral_sessions
[**ListEphemeralSessionsForOperation**](EphemeralSessionsAPI.md#ListEphemeralSessionsForOperation) | **Get** /operations/{operation_id}/ephemeral_sessions | list ephemeral_sessions



## GetEphemeralSession

> EphemeralSession GetEphemeralSession(ctx, id).Execute()

show ephemeral_session

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/aptible/aptible-api-go/aptibleapi"
)

func main() {
	id := int32(56) // int32 | id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EphemeralSessionsAPI.GetEphemeralSession(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EphemeralSessionsAPI.GetEphemeralSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEphemeralSession`: EphemeralSession
	fmt.Fprintf(os.Stdout, "Response from `EphemeralSessionsAPI.GetEphemeralSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEphemeralSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EphemeralSession**](EphemeralSession.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEphemeralSessionsForApp

> ListEphemeralSessionsForApp200Response ListEphemeralSessionsForApp(ctx, appId).Page(page).Execute()

list ephemeral_sessions

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/aptible/aptible-api-go/aptibleapi"
)

func main() {
	appId := int32(56) // int32 | app_id
	page := int32(56) // int32 | current page of results for pagination (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EphemeralSessionsAPI.ListEphemeralSessionsForApp(context.Background(), appId).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EphemeralSessionsAPI.ListEphemeralSessionsForApp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListEphemeralSessionsForApp`: ListEphemeralSessionsForApp200Response
	fmt.Fprintf(os.Stdout, "Response from `EphemeralSessionsAPI.ListEphemeralSessionsForApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** | app_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiListEphemeralSessionsForAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | current page of results for pagination | 

### Return type

[**ListEphemeralSessionsForApp200Response**](ListEphemeralSessionsForApp200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEphemeralSessionsForOperation

> ListEphemeralSessionsForApp200Response ListEphemeralSessionsForOperation(ctx, operationId).Page(page).Execute()

list ephemeral_sessions

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/aptible/aptible-api-go/aptibleapi"
)

func main() {
	operationId := int32(56) // int32 | operation_id
	page := int32(56) // int32 | current page of results for pagination (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EphemeralSessionsAPI.ListEphemeralSessionsForOperation(context.Background(), operationId).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EphemeralSessionsAPI.ListEphemeralSessionsForOperation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListEphemeralSessionsForOperation`: ListEphemeralSessionsForApp200Response
	fmt.Fprintf(os.Stdout, "Response from `EphemeralSessionsAPI.ListEphemeralSessionsForOperation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operationId** | **int32** | operation_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiListEphemeralSessionsForOperationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | current page of results for pagination | 

### Return type

[**ListEphemeralSessionsForApp200Response**](ListEphemeralSessionsForApp200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

