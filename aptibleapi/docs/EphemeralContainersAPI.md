# \EphemeralContainersAPI

All URIs are relative to *https://api.aptible.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEphemeralContainer**](EphemeralContainersAPI.md#GetEphemeralContainer) | **Get** /ephemeral_containers/{id} | show ephemeral_container
[**ListEphemeralContainersForEphemeralSession**](EphemeralContainersAPI.md#ListEphemeralContainersForEphemeralSession) | **Get** /ephemeral_sessions/{ephemeral_session_id}/ephemeral_containers | list ephemeral_containers
[**ListEphemeralContainersForLogDrain**](EphemeralContainersAPI.md#ListEphemeralContainersForLogDrain) | **Get** /log_drains/{log_drain_id}/ephemeral_containers | list ephemeral_containers



## GetEphemeralContainer

> EphemeralContainer GetEphemeralContainer(ctx, id).Execute()

show ephemeral_container

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
	resp, r, err := apiClient.EphemeralContainersAPI.GetEphemeralContainer(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EphemeralContainersAPI.GetEphemeralContainer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEphemeralContainer`: EphemeralContainer
	fmt.Fprintf(os.Stdout, "Response from `EphemeralContainersAPI.GetEphemeralContainer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEphemeralContainerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EphemeralContainer**](EphemeralContainer.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEphemeralContainersForEphemeralSession

> ListEphemeralContainersForEphemeralSession200Response ListEphemeralContainersForEphemeralSession(ctx, ephemeralSessionId).Page(page).Execute()

list ephemeral_containers

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
	ephemeralSessionId := int32(56) // int32 | ephemeral_session_id
	page := int32(56) // int32 | current page of results for pagination (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EphemeralContainersAPI.ListEphemeralContainersForEphemeralSession(context.Background(), ephemeralSessionId).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EphemeralContainersAPI.ListEphemeralContainersForEphemeralSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListEphemeralContainersForEphemeralSession`: ListEphemeralContainersForEphemeralSession200Response
	fmt.Fprintf(os.Stdout, "Response from `EphemeralContainersAPI.ListEphemeralContainersForEphemeralSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ephemeralSessionId** | **int32** | ephemeral_session_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiListEphemeralContainersForEphemeralSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | current page of results for pagination | 

### Return type

[**ListEphemeralContainersForEphemeralSession200Response**](ListEphemeralContainersForEphemeralSession200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEphemeralContainersForLogDrain

> ListEphemeralContainersForEphemeralSession200Response ListEphemeralContainersForLogDrain(ctx, logDrainId).Page(page).Execute()

list ephemeral_containers

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
	logDrainId := int32(56) // int32 | log_drain_id
	page := int32(56) // int32 | current page of results for pagination (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EphemeralContainersAPI.ListEphemeralContainersForLogDrain(context.Background(), logDrainId).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EphemeralContainersAPI.ListEphemeralContainersForLogDrain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListEphemeralContainersForLogDrain`: ListEphemeralContainersForEphemeralSession200Response
	fmt.Fprintf(os.Stdout, "Response from `EphemeralContainersAPI.ListEphemeralContainersForLogDrain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**logDrainId** | **int32** | log_drain_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiListEphemeralContainersForLogDrainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | current page of results for pagination | 

### Return type

[**ListEphemeralContainersForEphemeralSession200Response**](ListEphemeralContainersForEphemeralSession200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

