# \ExternalAwsResourcesAPI

All URIs are relative to *https://api.aptible.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteExternalAwsResource**](ExternalAwsResourcesAPI.md#DeleteExternalAwsResource) | **Delete** /external_aws_resources/{id} | delete external_aws_resource
[**GetExternalAwsResource**](ExternalAwsResourcesAPI.md#GetExternalAwsResource) | **Get** /external_aws_resources/{id} | show external_aws_resource
[**ListAppDatabaseConnectionsForExternalAwsResource**](ExternalAwsResourcesAPI.md#ListAppDatabaseConnectionsForExternalAwsResource) | **Get** /external_aws_resources/{external_aws_resource_id}/app_external_aws_rds_connections | list app database connections for external aws resource
[**ListExternalAwsDatabaseCredentialsForExternalAwsResource**](ExternalAwsResourcesAPI.md#ListExternalAwsDatabaseCredentialsForExternalAwsResource) | **Get** /external_aws_resources/{external_aws_resource_id}/external_aws_database_credentials | list external aws database credentials for resource
[**ListExternalAwsResources**](ExternalAwsResourcesAPI.md#ListExternalAwsResources) | **Get** /external_aws_resources | list external_aws_resources
[**PatchExternalAwsResource**](ExternalAwsResourcesAPI.md#PatchExternalAwsResource) | **Patch** /external_aws_resources/{id} | update external_aws_resource
[**UpdateExternalAwsResource**](ExternalAwsResourcesAPI.md#UpdateExternalAwsResource) | **Put** /external_aws_resources/{id} | update external_aws_resource



## DeleteExternalAwsResource

> DeleteExternalAwsResource(ctx, id).Execute()

delete external_aws_resource

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
	r, err := apiClient.ExternalAwsResourcesAPI.DeleteExternalAwsResource(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExternalAwsResourcesAPI.DeleteExternalAwsResource``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteExternalAwsResourceRequest struct via the builder pattern


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


## GetExternalAwsResource

> ExternalAwsResource GetExternalAwsResource(ctx, id).Execute()

show external_aws_resource

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
	resp, r, err := apiClient.ExternalAwsResourcesAPI.GetExternalAwsResource(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExternalAwsResourcesAPI.GetExternalAwsResource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExternalAwsResource`: ExternalAwsResource
	fmt.Fprintf(os.Stdout, "Response from `ExternalAwsResourcesAPI.GetExternalAwsResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExternalAwsResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ExternalAwsResource**](ExternalAwsResource.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAppDatabaseConnectionsForExternalAwsResource

> ListAppDatabaseConnectionsForExternalAwsAccount200Response ListAppDatabaseConnectionsForExternalAwsResource(ctx, externalAwsResourceId).Execute()

list app database connections for external aws resource

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
	externalAwsResourceId := int32(56) // int32 | external_aws_resource_id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExternalAwsResourcesAPI.ListAppDatabaseConnectionsForExternalAwsResource(context.Background(), externalAwsResourceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExternalAwsResourcesAPI.ListAppDatabaseConnectionsForExternalAwsResource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAppDatabaseConnectionsForExternalAwsResource`: ListAppDatabaseConnectionsForExternalAwsAccount200Response
	fmt.Fprintf(os.Stdout, "Response from `ExternalAwsResourcesAPI.ListAppDatabaseConnectionsForExternalAwsResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalAwsResourceId** | **int32** | external_aws_resource_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAppDatabaseConnectionsForExternalAwsResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListAppDatabaseConnectionsForExternalAwsAccount200Response**](ListAppDatabaseConnectionsForExternalAwsAccount200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListExternalAwsDatabaseCredentialsForExternalAwsResource

> ListExternalAwsDatabaseCredentialsForExternalAwsResource200Response ListExternalAwsDatabaseCredentialsForExternalAwsResource(ctx, externalAwsResourceId).Page(page).Execute()

list external aws database credentials for resource

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
	externalAwsResourceId := int32(56) // int32 | external_aws_resource_id
	page := int32(56) // int32 | current page of results for pagination (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExternalAwsResourcesAPI.ListExternalAwsDatabaseCredentialsForExternalAwsResource(context.Background(), externalAwsResourceId).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExternalAwsResourcesAPI.ListExternalAwsDatabaseCredentialsForExternalAwsResource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListExternalAwsDatabaseCredentialsForExternalAwsResource`: ListExternalAwsDatabaseCredentialsForExternalAwsResource200Response
	fmt.Fprintf(os.Stdout, "Response from `ExternalAwsResourcesAPI.ListExternalAwsDatabaseCredentialsForExternalAwsResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalAwsResourceId** | **int32** | external_aws_resource_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiListExternalAwsDatabaseCredentialsForExternalAwsResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | current page of results for pagination | 

### Return type

[**ListExternalAwsDatabaseCredentialsForExternalAwsResource200Response**](ListExternalAwsDatabaseCredentialsForExternalAwsResource200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListExternalAwsResources

> ListExternalAwsResources200Response ListExternalAwsResources(ctx).Page(page).Execute()

list external_aws_resources

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
	page := int32(56) // int32 | current page of results for pagination (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExternalAwsResourcesAPI.ListExternalAwsResources(context.Background()).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExternalAwsResourcesAPI.ListExternalAwsResources``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListExternalAwsResources`: ListExternalAwsResources200Response
	fmt.Fprintf(os.Stdout, "Response from `ExternalAwsResourcesAPI.ListExternalAwsResources`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListExternalAwsResourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | current page of results for pagination | 

### Return type

[**ListExternalAwsResources200Response**](ListExternalAwsResources200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchExternalAwsResource

> PatchExternalAwsResource(ctx, id).UpdateExternalAwsResourceRequest(updateExternalAwsResourceRequest).Execute()

update external_aws_resource

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
	updateExternalAwsResourceRequest := *openapiclient.NewUpdateExternalAwsResourceRequest() // UpdateExternalAwsResourceRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ExternalAwsResourcesAPI.PatchExternalAwsResource(context.Background(), id).UpdateExternalAwsResourceRequest(updateExternalAwsResourceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExternalAwsResourcesAPI.PatchExternalAwsResource``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPatchExternalAwsResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateExternalAwsResourceRequest** | [**UpdateExternalAwsResourceRequest**](UpdateExternalAwsResourceRequest.md) |  | 

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


## UpdateExternalAwsResource

> UpdateExternalAwsResource(ctx, id).UpdateExternalAwsResourceRequest(updateExternalAwsResourceRequest).Execute()

update external_aws_resource

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
	updateExternalAwsResourceRequest := *openapiclient.NewUpdateExternalAwsResourceRequest() // UpdateExternalAwsResourceRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ExternalAwsResourcesAPI.UpdateExternalAwsResource(context.Background(), id).UpdateExternalAwsResourceRequest(updateExternalAwsResourceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExternalAwsResourcesAPI.UpdateExternalAwsResource``: %v\n", err)
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

Other parameters are passed through a pointer to a apiUpdateExternalAwsResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateExternalAwsResourceRequest** | [**UpdateExternalAwsResourceRequest**](UpdateExternalAwsResourceRequest.md) |  | 

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

