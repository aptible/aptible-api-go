# \AppExternalAwsRdsConnectionsAPI

All URIs are relative to *https://api.aptible.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAppExternalAwsRdsConnection**](AppExternalAwsRdsConnectionsAPI.md#DeleteAppExternalAwsRdsConnection) | **Delete** /app_external_aws_rds_connections/{id} | delete app_external_aws_rds_connection
[**GetAppExternalAwsRdsConnection**](AppExternalAwsRdsConnectionsAPI.md#GetAppExternalAwsRdsConnection) | **Get** /app_external_aws_rds_connections/{id} | show app_external_aws_rds_connection
[**PatchAppExternalAwsRdsConnection**](AppExternalAwsRdsConnectionsAPI.md#PatchAppExternalAwsRdsConnection) | **Patch** /app_external_aws_rds_connections/{id} | update app_external_aws_rds_connection
[**UpdateAppExternalAwsRdsConnection**](AppExternalAwsRdsConnectionsAPI.md#UpdateAppExternalAwsRdsConnection) | **Put** /app_external_aws_rds_connections/{id} | update app_external_aws_rds_connection



## DeleteAppExternalAwsRdsConnection

> DeleteAppExternalAwsRdsConnection(ctx, id).Execute()

delete app_external_aws_rds_connection

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
	r, err := apiClient.AppExternalAwsRdsConnectionsAPI.DeleteAppExternalAwsRdsConnection(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppExternalAwsRdsConnectionsAPI.DeleteAppExternalAwsRdsConnection``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteAppExternalAwsRdsConnectionRequest struct via the builder pattern


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


## GetAppExternalAwsRdsConnection

> AppExternalAwsRdsConnection GetAppExternalAwsRdsConnection(ctx, id).Execute()

show app_external_aws_rds_connection

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
	resp, r, err := apiClient.AppExternalAwsRdsConnectionsAPI.GetAppExternalAwsRdsConnection(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppExternalAwsRdsConnectionsAPI.GetAppExternalAwsRdsConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAppExternalAwsRdsConnection`: AppExternalAwsRdsConnection
	fmt.Fprintf(os.Stdout, "Response from `AppExternalAwsRdsConnectionsAPI.GetAppExternalAwsRdsConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAppExternalAwsRdsConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AppExternalAwsRdsConnection**](AppExternalAwsRdsConnection.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchAppExternalAwsRdsConnection

> PatchAppExternalAwsRdsConnection(ctx, id).UpdateAppExternalAwsRdsConnectionRequest(updateAppExternalAwsRdsConnectionRequest).Execute()

update app_external_aws_rds_connection

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
	updateAppExternalAwsRdsConnectionRequest := *openapiclient.NewUpdateAppExternalAwsRdsConnectionRequest() // UpdateAppExternalAwsRdsConnectionRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AppExternalAwsRdsConnectionsAPI.PatchAppExternalAwsRdsConnection(context.Background(), id).UpdateAppExternalAwsRdsConnectionRequest(updateAppExternalAwsRdsConnectionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppExternalAwsRdsConnectionsAPI.PatchAppExternalAwsRdsConnection``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPatchAppExternalAwsRdsConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateAppExternalAwsRdsConnectionRequest** | [**UpdateAppExternalAwsRdsConnectionRequest**](UpdateAppExternalAwsRdsConnectionRequest.md) |  | 

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


## UpdateAppExternalAwsRdsConnection

> UpdateAppExternalAwsRdsConnection(ctx, id).UpdateAppExternalAwsRdsConnectionRequest(updateAppExternalAwsRdsConnectionRequest).Execute()

update app_external_aws_rds_connection

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
	updateAppExternalAwsRdsConnectionRequest := *openapiclient.NewUpdateAppExternalAwsRdsConnectionRequest() // UpdateAppExternalAwsRdsConnectionRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AppExternalAwsRdsConnectionsAPI.UpdateAppExternalAwsRdsConnection(context.Background(), id).UpdateAppExternalAwsRdsConnectionRequest(updateAppExternalAwsRdsConnectionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppExternalAwsRdsConnectionsAPI.UpdateAppExternalAwsRdsConnection``: %v\n", err)
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

Other parameters are passed through a pointer to a apiUpdateAppExternalAwsRdsConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateAppExternalAwsRdsConnectionRequest** | [**UpdateAppExternalAwsRdsConnectionRequest**](UpdateAppExternalAwsRdsConnectionRequest.md) |  | 

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

