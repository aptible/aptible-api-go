# \ContainersAPI

All URIs are relative to *https://api.aptible.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetContainer**](ContainersAPI.md#GetContainer) | **Get** /containers/{id} | show container
[**ListContainersForRelease**](ContainersAPI.md#ListContainersForRelease) | **Get** /releases/{release_id}/containers | list containers
[**LogDrainsLogDrainIdContainersGet**](ContainersAPI.md#LogDrainsLogDrainIdContainersGet) | **Get** /log_drains/{log_drain_id}/containers | list containers
[**MetricDrainsMetricDrainIdContainersGet**](ContainersAPI.md#MetricDrainsMetricDrainIdContainersGet) | **Get** /metric_drains/{metric_drain_id}/containers | list containers



## GetContainer

> Container GetContainer(ctx, id).Execute()

show container

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
	resp, r, err := apiClient.ContainersAPI.GetContainer(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContainersAPI.GetContainer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContainer`: Container
	fmt.Fprintf(os.Stdout, "Response from `ContainersAPI.GetContainer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContainerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Container**](Container.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListContainersForRelease

> ListContainersForRelease200Response ListContainersForRelease(ctx, releaseId).Page(page).Execute()

list containers

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
	releaseId := int32(56) // int32 | release_id
	page := int32(56) // int32 | current page of results for pagination (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContainersAPI.ListContainersForRelease(context.Background(), releaseId).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContainersAPI.ListContainersForRelease``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListContainersForRelease`: ListContainersForRelease200Response
	fmt.Fprintf(os.Stdout, "Response from `ContainersAPI.ListContainersForRelease`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**releaseId** | **int32** | release_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiListContainersForReleaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | current page of results for pagination | 

### Return type

[**ListContainersForRelease200Response**](ListContainersForRelease200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LogDrainsLogDrainIdContainersGet

> LogDrainsLogDrainIdContainersGet200Response LogDrainsLogDrainIdContainersGet(ctx, logDrainId).Page(page).Execute()

list containers

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
	resp, r, err := apiClient.ContainersAPI.LogDrainsLogDrainIdContainersGet(context.Background(), logDrainId).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContainersAPI.LogDrainsLogDrainIdContainersGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LogDrainsLogDrainIdContainersGet`: LogDrainsLogDrainIdContainersGet200Response
	fmt.Fprintf(os.Stdout, "Response from `ContainersAPI.LogDrainsLogDrainIdContainersGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**logDrainId** | **int32** | log_drain_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiLogDrainsLogDrainIdContainersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | current page of results for pagination | 

### Return type

[**LogDrainsLogDrainIdContainersGet200Response**](LogDrainsLogDrainIdContainersGet200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MetricDrainsMetricDrainIdContainersGet

> MetricDrainsMetricDrainIdContainersGet200Response MetricDrainsMetricDrainIdContainersGet(ctx, metricDrainId).Page(page).Execute()

list containers

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
	metricDrainId := int32(56) // int32 | metric_drain_id
	page := int32(56) // int32 | current page of results for pagination (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContainersAPI.MetricDrainsMetricDrainIdContainersGet(context.Background(), metricDrainId).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContainersAPI.MetricDrainsMetricDrainIdContainersGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MetricDrainsMetricDrainIdContainersGet`: MetricDrainsMetricDrainIdContainersGet200Response
	fmt.Fprintf(os.Stdout, "Response from `ContainersAPI.MetricDrainsMetricDrainIdContainersGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**metricDrainId** | **int32** | metric_drain_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiMetricDrainsMetricDrainIdContainersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | current page of results for pagination | 

### Return type

[**MetricDrainsMetricDrainIdContainersGet200Response**](MetricDrainsMetricDrainIdContainersGet200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

