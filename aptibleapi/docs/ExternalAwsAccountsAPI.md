# \ExternalAwsAccountsAPI

All URIs are relative to *https://api.aptible.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateExternalAwsAccount**](ExternalAwsAccountsAPI.md#CreateExternalAwsAccount) | **Post** /external_aws_accounts | create external_aws_account
[**CreateExternalAwsResource**](ExternalAwsAccountsAPI.md#CreateExternalAwsResource) | **Post** /external_aws_accounts/{external_aws_account_id}/external_aws_resources | create external_aws_resource
[**DeleteExternalAwsAccount**](ExternalAwsAccountsAPI.md#DeleteExternalAwsAccount) | **Delete** /external_aws_accounts/{id} | delete external_aws_account
[**GetExternalAwsAccount**](ExternalAwsAccountsAPI.md#GetExternalAwsAccount) | **Get** /external_aws_accounts/{id} | show external_aws_account
[**ListAppDatabaseConnectionsForExternalAwsAccount**](ExternalAwsAccountsAPI.md#ListAppDatabaseConnectionsForExternalAwsAccount) | **Get** /external_aws_accounts/{external_aws_account_id}/app_external_aws_rds_connections | list app_external_aws_rds_connections for external_aws_account
[**ListExternalAwsAccounts**](ExternalAwsAccountsAPI.md#ListExternalAwsAccounts) | **Get** /external_aws_accounts | list external_aws_accounts
[**ListExternalAwsResourcesForExternalAwsAccount**](ExternalAwsAccountsAPI.md#ListExternalAwsResourcesForExternalAwsAccount) | **Get** /external_aws_accounts/{external_aws_account_id}/external_aws_resources | list external_aws_resources
[**PatchExternalAwsAccount**](ExternalAwsAccountsAPI.md#PatchExternalAwsAccount) | **Patch** /external_aws_accounts/{id} | update external_aws_account
[**UpdateExternalAwsAccount**](ExternalAwsAccountsAPI.md#UpdateExternalAwsAccount) | **Put** /external_aws_accounts/{id} | update external_aws_account



## CreateExternalAwsAccount

> ExternalAwsAccount CreateExternalAwsAccount(ctx).CreateExternalAwsAccountRequest(createExternalAwsAccountRequest).Execute()

create external_aws_account

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
	createExternalAwsAccountRequest := *openapiclient.NewCreateExternalAwsAccountRequest("OrganizationId_example", "AwsAccountId_example", "AccountName_example", "AwsRegionPrimary_example") // CreateExternalAwsAccountRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExternalAwsAccountsAPI.CreateExternalAwsAccount(context.Background()).CreateExternalAwsAccountRequest(createExternalAwsAccountRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExternalAwsAccountsAPI.CreateExternalAwsAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateExternalAwsAccount`: ExternalAwsAccount
	fmt.Fprintf(os.Stdout, "Response from `ExternalAwsAccountsAPI.CreateExternalAwsAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateExternalAwsAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createExternalAwsAccountRequest** | [**CreateExternalAwsAccountRequest**](CreateExternalAwsAccountRequest.md) |  | 

### Return type

[**ExternalAwsAccount**](ExternalAwsAccount.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateExternalAwsResource

> ExternalAwsResource CreateExternalAwsResource(ctx, externalAwsAccountId).CreateExternalAwsResourceRequest(createExternalAwsResourceRequest).Execute()

create external_aws_resource

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
	externalAwsAccountId := int32(56) // int32 | external_aws_account_id
	createExternalAwsResourceRequest := *openapiclient.NewCreateExternalAwsResourceRequest("ResourceType_example", "ResourceArn_example", "ResourceId_example", "Region_example", "SyncStatus_example", map[string]interface{}(123), []openapiclient.CreateExternalAwsResourceRequestTagsInner{*openapiclient.NewCreateExternalAwsResourceRequestTagsInner("Key_example", "Value_example")}) // CreateExternalAwsResourceRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExternalAwsAccountsAPI.CreateExternalAwsResource(context.Background(), externalAwsAccountId).CreateExternalAwsResourceRequest(createExternalAwsResourceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExternalAwsAccountsAPI.CreateExternalAwsResource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateExternalAwsResource`: ExternalAwsResource
	fmt.Fprintf(os.Stdout, "Response from `ExternalAwsAccountsAPI.CreateExternalAwsResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalAwsAccountId** | **int32** | external_aws_account_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateExternalAwsResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createExternalAwsResourceRequest** | [**CreateExternalAwsResourceRequest**](CreateExternalAwsResourceRequest.md) |  | 

### Return type

[**ExternalAwsResource**](ExternalAwsResource.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteExternalAwsAccount

> DeleteExternalAwsAccount(ctx, id).Execute()

delete external_aws_account

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
	r, err := apiClient.ExternalAwsAccountsAPI.DeleteExternalAwsAccount(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExternalAwsAccountsAPI.DeleteExternalAwsAccount``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteExternalAwsAccountRequest struct via the builder pattern


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


## GetExternalAwsAccount

> ExternalAwsAccount GetExternalAwsAccount(ctx, id).Execute()

show external_aws_account

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
	resp, r, err := apiClient.ExternalAwsAccountsAPI.GetExternalAwsAccount(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExternalAwsAccountsAPI.GetExternalAwsAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExternalAwsAccount`: ExternalAwsAccount
	fmt.Fprintf(os.Stdout, "Response from `ExternalAwsAccountsAPI.GetExternalAwsAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExternalAwsAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ExternalAwsAccount**](ExternalAwsAccount.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAppDatabaseConnectionsForExternalAwsAccount

> ListAppDatabaseConnectionsForExternalAwsAccount200Response ListAppDatabaseConnectionsForExternalAwsAccount(ctx, externalAwsAccountId).Execute()

list app_external_aws_rds_connections for external_aws_account

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
	externalAwsAccountId := int32(56) // int32 | external_aws_account_id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExternalAwsAccountsAPI.ListAppDatabaseConnectionsForExternalAwsAccount(context.Background(), externalAwsAccountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExternalAwsAccountsAPI.ListAppDatabaseConnectionsForExternalAwsAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAppDatabaseConnectionsForExternalAwsAccount`: ListAppDatabaseConnectionsForExternalAwsAccount200Response
	fmt.Fprintf(os.Stdout, "Response from `ExternalAwsAccountsAPI.ListAppDatabaseConnectionsForExternalAwsAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalAwsAccountId** | **int32** | external_aws_account_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAppDatabaseConnectionsForExternalAwsAccountRequest struct via the builder pattern


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


## ListExternalAwsAccounts

> ListExternalAwsAccounts200Response ListExternalAwsAccounts(ctx).Page(page).Execute()

list external_aws_accounts

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
	resp, r, err := apiClient.ExternalAwsAccountsAPI.ListExternalAwsAccounts(context.Background()).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExternalAwsAccountsAPI.ListExternalAwsAccounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListExternalAwsAccounts`: ListExternalAwsAccounts200Response
	fmt.Fprintf(os.Stdout, "Response from `ExternalAwsAccountsAPI.ListExternalAwsAccounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListExternalAwsAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | current page of results for pagination | 

### Return type

[**ListExternalAwsAccounts200Response**](ListExternalAwsAccounts200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListExternalAwsResourcesForExternalAwsAccount

> ListExternalAwsResourcesForExternalAwsAccount200Response ListExternalAwsResourcesForExternalAwsAccount(ctx, externalAwsAccountId).Page(page).Execute()

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
	externalAwsAccountId := int32(56) // int32 | external_aws_account_id
	page := int32(56) // int32 | current page of results for pagination (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExternalAwsAccountsAPI.ListExternalAwsResourcesForExternalAwsAccount(context.Background(), externalAwsAccountId).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExternalAwsAccountsAPI.ListExternalAwsResourcesForExternalAwsAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListExternalAwsResourcesForExternalAwsAccount`: ListExternalAwsResourcesForExternalAwsAccount200Response
	fmt.Fprintf(os.Stdout, "Response from `ExternalAwsAccountsAPI.ListExternalAwsResourcesForExternalAwsAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalAwsAccountId** | **int32** | external_aws_account_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiListExternalAwsResourcesForExternalAwsAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | current page of results for pagination | 

### Return type

[**ListExternalAwsResourcesForExternalAwsAccount200Response**](ListExternalAwsResourcesForExternalAwsAccount200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchExternalAwsAccount

> PatchExternalAwsAccount(ctx, id).UpdateExternalAwsAccountRequest(updateExternalAwsAccountRequest).Execute()

update external_aws_account

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
	updateExternalAwsAccountRequest := *openapiclient.NewUpdateExternalAwsAccountRequest() // UpdateExternalAwsAccountRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ExternalAwsAccountsAPI.PatchExternalAwsAccount(context.Background(), id).UpdateExternalAwsAccountRequest(updateExternalAwsAccountRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExternalAwsAccountsAPI.PatchExternalAwsAccount``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPatchExternalAwsAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateExternalAwsAccountRequest** | [**UpdateExternalAwsAccountRequest**](UpdateExternalAwsAccountRequest.md) |  | 

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


## UpdateExternalAwsAccount

> UpdateExternalAwsAccount(ctx, id).UpdateExternalAwsAccountRequest(updateExternalAwsAccountRequest).Execute()

update external_aws_account

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
	updateExternalAwsAccountRequest := *openapiclient.NewUpdateExternalAwsAccountRequest() // UpdateExternalAwsAccountRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ExternalAwsAccountsAPI.UpdateExternalAwsAccount(context.Background(), id).UpdateExternalAwsAccountRequest(updateExternalAwsAccountRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExternalAwsAccountsAPI.UpdateExternalAwsAccount``: %v\n", err)
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

Other parameters are passed through a pointer to a apiUpdateExternalAwsAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateExternalAwsAccountRequest** | [**UpdateExternalAwsAccountRequest**](UpdateExternalAwsAccountRequest.md) |  | 

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

