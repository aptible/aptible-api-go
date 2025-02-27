# \LogDrainsAPI

All URIs are relative to *https://api.aptible.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateLogDrain**](LogDrainsAPI.md#CreateLogDrain) | **Post** /accounts/{account_id}/log_drains | create log_drain
[**DeleteLogDrain**](LogDrainsAPI.md#DeleteLogDrain) | **Delete** /log_drains/{id} | delete log_drain
[**GetLogDrain**](LogDrainsAPI.md#GetLogDrain) | **Get** /log_drains/{id} | show log_drain
[**ListLogDrainsForAccount**](LogDrainsAPI.md#ListLogDrainsForAccount) | **Get** /accounts/{account_id}/log_drains | list log_drains
[**PatchLogDrain**](LogDrainsAPI.md#PatchLogDrain) | **Patch** /log_drains/{id} | update log_drain
[**UpdateLogDrain**](LogDrainsAPI.md#UpdateLogDrain) | **Put** /log_drains/{id} | update log_drain



## CreateLogDrain

> LogDrain CreateLogDrain(ctx, accountId).CreateLogDrainRequest(createLogDrainRequest).Execute()

create log_drain

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
	accountId := int32(56) // int32 | account_id
	createLogDrainRequest := *openapiclient.NewCreateLogDrainRequest("Handle_example", "DrainType_example") // CreateLogDrainRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogDrainsAPI.CreateLogDrain(context.Background(), accountId).CreateLogDrainRequest(createLogDrainRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogDrainsAPI.CreateLogDrain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateLogDrain`: LogDrain
	fmt.Fprintf(os.Stdout, "Response from `LogDrainsAPI.CreateLogDrain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **int32** | account_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateLogDrainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createLogDrainRequest** | [**CreateLogDrainRequest**](CreateLogDrainRequest.md) |  | 

### Return type

[**LogDrain**](LogDrain.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteLogDrain

> DeleteLogDrain(ctx, id).Execute()

delete log_drain

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
	r, err := apiClient.LogDrainsAPI.DeleteLogDrain(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogDrainsAPI.DeleteLogDrain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLogDrainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLogDrain

> LogDrain GetLogDrain(ctx, id).Execute()

show log_drain

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
	resp, r, err := apiClient.LogDrainsAPI.GetLogDrain(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogDrainsAPI.GetLogDrain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLogDrain`: LogDrain
	fmt.Fprintf(os.Stdout, "Response from `LogDrainsAPI.GetLogDrain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLogDrainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LogDrain**](LogDrain.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLogDrainsForAccount

> ListLogDrainsForAccount200Response ListLogDrainsForAccount(ctx, accountId).Page(page).Execute()

list log_drains

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
	accountId := int32(56) // int32 | account_id
	page := int32(56) // int32 | current page of results for pagination (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogDrainsAPI.ListLogDrainsForAccount(context.Background(), accountId).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogDrainsAPI.ListLogDrainsForAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListLogDrainsForAccount`: ListLogDrainsForAccount200Response
	fmt.Fprintf(os.Stdout, "Response from `LogDrainsAPI.ListLogDrainsForAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **int32** | account_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiListLogDrainsForAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | current page of results for pagination | 

### Return type

[**ListLogDrainsForAccount200Response**](ListLogDrainsForAccount200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchLogDrain

> PatchLogDrain(ctx, id).UpdateLogDrainRequest(updateLogDrainRequest).Execute()

update log_drain

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
	updateLogDrainRequest := *openapiclient.NewUpdateLogDrainRequest() // UpdateLogDrainRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LogDrainsAPI.PatchLogDrain(context.Background(), id).UpdateLogDrainRequest(updateLogDrainRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogDrainsAPI.PatchLogDrain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchLogDrainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateLogDrainRequest** | [**UpdateLogDrainRequest**](UpdateLogDrainRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLogDrain

> UpdateLogDrain(ctx, id).UpdateLogDrainRequest(updateLogDrainRequest).Execute()

update log_drain

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
	updateLogDrainRequest := *openapiclient.NewUpdateLogDrainRequest() // UpdateLogDrainRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LogDrainsAPI.UpdateLogDrain(context.Background(), id).UpdateLogDrainRequest(updateLogDrainRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogDrainsAPI.UpdateLogDrain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLogDrainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateLogDrainRequest** | [**UpdateLogDrainRequest**](UpdateLogDrainRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

