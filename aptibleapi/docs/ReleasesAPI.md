# \ReleasesAPI

All URIs are relative to *https://api.aptible.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetRelease**](ReleasesAPI.md#GetRelease) | **Get** /releases/{id} | show release
[**ListReleasesForService**](ReleasesAPI.md#ListReleasesForService) | **Get** /services/{service_id}/releases | list releases



## GetRelease

> Release GetRelease(ctx, id).Execute()

show release

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
	resp, r, err := apiClient.ReleasesAPI.GetRelease(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReleasesAPI.GetRelease``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRelease`: Release
	fmt.Fprintf(os.Stdout, "Response from `ReleasesAPI.GetRelease`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReleaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Release**](Release.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListReleasesForService

> ListReleasesForService200Response ListReleasesForService(ctx, serviceId).Page(page).Execute()

list releases

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
	serviceId := int32(56) // int32 | service_id
	page := int32(56) // int32 | current page of results for pagination (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReleasesAPI.ListReleasesForService(context.Background(), serviceId).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReleasesAPI.ListReleasesForService``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListReleasesForService`: ListReleasesForService200Response
	fmt.Fprintf(os.Stdout, "Response from `ReleasesAPI.ListReleasesForService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **int32** | service_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiListReleasesForServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | current page of results for pagination | 

### Return type

[**ListReleasesForService200Response**](ListReleasesForService200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

