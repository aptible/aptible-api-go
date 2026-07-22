# \ContainersAPI

All URIs are relative to *https://api.aptible.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetContainer**](ContainersAPI.md#GetContainer) | **Get** /containers/{id} | show container
[**ListContainersForLogDrain**](ContainersAPI.md#ListContainersForLogDrain) | **Get** /log_drains/{log_drain_id}/containers | list containers
[**ListContainersForMetricDrain**](ContainersAPI.md#ListContainersForMetricDrain) | **Get** /metric_drains/{metric_drain_id}/containers | list containers
[**ListContainersForRelease**](ContainersAPI.md#ListContainersForRelease) | **Get** /releases/{release_id}/containers | list containers



## GetContainer

> Container GetContainer(ctx, id).NoEmbed(noEmbed).Prefer(prefer).Execute()

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
	noEmbed := true // bool | When true, omits embedded resources from the response. Also triggered when the Prefer: no_sensitive_extras=true header is present. (optional)
	prefer := "prefer_example" // string | When set to no_sensitive_extras=true, omits sensitive fields and embedded resources from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContainersAPI.GetContainer(context.Background(), id).NoEmbed(noEmbed).Prefer(prefer).Execute()
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

 **noEmbed** | **bool** | When true, omits embedded resources from the response. Also triggered when the Prefer: no_sensitive_extras&#x3D;true header is present. | 
 **prefer** | **string** | When set to no_sensitive_extras&#x3D;true, omits sensitive fields and embedded resources from the response. | 

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


## ListContainersForLogDrain

> ListContainersForRelease200Response ListContainersForLogDrain(ctx, logDrainId).Page(page).PerPage(perPage).NoEmbed(noEmbed).Prefer(prefer).Execute()

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
	page := int32(56) // int32 | Current page of paginated results (optional)
	perPage := int32(56) // int32 | Number of results to return per page (optional)
	noEmbed := true // bool | When true, omits embedded resources from the response. Also triggered when the Prefer: no_sensitive_extras=true header is present. (optional)
	prefer := "prefer_example" // string | When set to no_sensitive_extras=true, omits sensitive fields and embedded resources from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContainersAPI.ListContainersForLogDrain(context.Background(), logDrainId).Page(page).PerPage(perPage).NoEmbed(noEmbed).Prefer(prefer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContainersAPI.ListContainersForLogDrain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListContainersForLogDrain`: ListContainersForRelease200Response
	fmt.Fprintf(os.Stdout, "Response from `ContainersAPI.ListContainersForLogDrain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**logDrainId** | **int32** | log_drain_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiListContainersForLogDrainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Current page of paginated results | 
 **perPage** | **int32** | Number of results to return per page | 
 **noEmbed** | **bool** | When true, omits embedded resources from the response. Also triggered when the Prefer: no_sensitive_extras&#x3D;true header is present. | 
 **prefer** | **string** | When set to no_sensitive_extras&#x3D;true, omits sensitive fields and embedded resources from the response. | 

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


## ListContainersForMetricDrain

> ListContainersForRelease200Response ListContainersForMetricDrain(ctx, metricDrainId).Page(page).PerPage(perPage).NoEmbed(noEmbed).Prefer(prefer).Execute()

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
	page := int32(56) // int32 | Current page of paginated results (optional)
	perPage := int32(56) // int32 | Number of results to return per page (optional)
	noEmbed := true // bool | When true, omits embedded resources from the response. Also triggered when the Prefer: no_sensitive_extras=true header is present. (optional)
	prefer := "prefer_example" // string | When set to no_sensitive_extras=true, omits sensitive fields and embedded resources from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContainersAPI.ListContainersForMetricDrain(context.Background(), metricDrainId).Page(page).PerPage(perPage).NoEmbed(noEmbed).Prefer(prefer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContainersAPI.ListContainersForMetricDrain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListContainersForMetricDrain`: ListContainersForRelease200Response
	fmt.Fprintf(os.Stdout, "Response from `ContainersAPI.ListContainersForMetricDrain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**metricDrainId** | **int32** | metric_drain_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiListContainersForMetricDrainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Current page of paginated results | 
 **perPage** | **int32** | Number of results to return per page | 
 **noEmbed** | **bool** | When true, omits embedded resources from the response. Also triggered when the Prefer: no_sensitive_extras&#x3D;true header is present. | 
 **prefer** | **string** | When set to no_sensitive_extras&#x3D;true, omits sensitive fields and embedded resources from the response. | 

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


## ListContainersForRelease

> ListContainersForRelease200Response ListContainersForRelease(ctx, releaseId).Page(page).PerPage(perPage).NoEmbed(noEmbed).Prefer(prefer).Execute()

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
	page := int32(56) // int32 | Current page of paginated results (optional)
	perPage := int32(56) // int32 | Number of results to return per page (optional)
	noEmbed := true // bool | When true, omits embedded resources from the response. Also triggered when the Prefer: no_sensitive_extras=true header is present. (optional)
	prefer := "prefer_example" // string | When set to no_sensitive_extras=true, omits sensitive fields and embedded resources from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContainersAPI.ListContainersForRelease(context.Background(), releaseId).Page(page).PerPage(perPage).NoEmbed(noEmbed).Prefer(prefer).Execute()
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

 **page** | **int32** | Current page of paginated results | 
 **perPage** | **int32** | Number of results to return per page | 
 **noEmbed** | **bool** | When true, omits embedded resources from the response. Also triggered when the Prefer: no_sensitive_extras&#x3D;true header is present. | 
 **prefer** | **string** | When set to no_sensitive_extras&#x3D;true, omits sensitive fields and embedded resources from the response. | 

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

