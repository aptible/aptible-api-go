# \DatabasesAPI

All URIs are relative to *https://api.aptible.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDatabase**](DatabasesAPI.md#CreateDatabase) | **Post** /accounts/{account_id}/databases | create database
[**DeleteDatabase**](DatabasesAPI.md#DeleteDatabase) | **Delete** /databases/{id} | delete database
[**GetDatabase**](DatabasesAPI.md#GetDatabase) | **Get** /databases/{id} | show database
[**ListDatabases**](DatabasesAPI.md#ListDatabases) | **Get** /databases | list databases
[**ListDatabasesForAccount**](DatabasesAPI.md#ListDatabasesForAccount) | **Get** /accounts/{account_id}/databases | list databases
[**ListReplicasForDatabase**](DatabasesAPI.md#ListReplicasForDatabase) | **Get** /databases/{database_id}/dependents | list databases
[**PatchDatabase**](DatabasesAPI.md#PatchDatabase) | **Patch** /databases/{id} | update database
[**UpdateDatabase**](DatabasesAPI.md#UpdateDatabase) | **Put** /databases/{id} | update database



## CreateDatabase

> Database CreateDatabase(ctx, accountId).CreateDatabaseRequest(createDatabaseRequest).Execute()

create database

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
	createDatabaseRequest := *openapiclient.NewCreateDatabaseRequest("Handle_example", "Type_example") // CreateDatabaseRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DatabasesAPI.CreateDatabase(context.Background(), accountId).CreateDatabaseRequest(createDatabaseRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.CreateDatabase``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDatabase`: Database
	fmt.Fprintf(os.Stdout, "Response from `DatabasesAPI.CreateDatabase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **int32** | account_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDatabaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createDatabaseRequest** | [**CreateDatabaseRequest**](CreateDatabaseRequest.md) |  | 

### Return type

[**Database**](Database.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDatabase

> DeleteDatabase(ctx, id).Execute()

delete database

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
	r, err := apiClient.DatabasesAPI.DeleteDatabase(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.DeleteDatabase``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteDatabaseRequest struct via the builder pattern


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


## GetDatabase

> Database GetDatabase(ctx, id).Execute()

show database

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
	resp, r, err := apiClient.DatabasesAPI.GetDatabase(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.GetDatabase``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDatabase`: Database
	fmt.Fprintf(os.Stdout, "Response from `DatabasesAPI.GetDatabase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDatabaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Database**](Database.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDatabases

> ListDatabasesForAccount200Response ListDatabases(ctx).Page(page).Execute()

list databases

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
	resp, r, err := apiClient.DatabasesAPI.ListDatabases(context.Background()).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.ListDatabases``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDatabases`: ListDatabasesForAccount200Response
	fmt.Fprintf(os.Stdout, "Response from `DatabasesAPI.ListDatabases`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDatabasesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | current page of results for pagination | 

### Return type

[**ListDatabasesForAccount200Response**](ListDatabasesForAccount200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDatabasesForAccount

> ListDatabasesForAccount200Response ListDatabasesForAccount(ctx, accountId).Page(page).Execute()

list databases

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
	resp, r, err := apiClient.DatabasesAPI.ListDatabasesForAccount(context.Background(), accountId).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.ListDatabasesForAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDatabasesForAccount`: ListDatabasesForAccount200Response
	fmt.Fprintf(os.Stdout, "Response from `DatabasesAPI.ListDatabasesForAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **int32** | account_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiListDatabasesForAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | current page of results for pagination | 

### Return type

[**ListDatabasesForAccount200Response**](ListDatabasesForAccount200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListReplicasForDatabase

> ListDatabasesForAccount200Response ListReplicasForDatabase(ctx, databaseId).Page(page).Execute()

list databases

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
	databaseId := int32(56) // int32 | database_id
	page := int32(56) // int32 | current page of results for pagination (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DatabasesAPI.ListReplicasForDatabase(context.Background(), databaseId).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.ListReplicasForDatabase``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListReplicasForDatabase`: ListDatabasesForAccount200Response
	fmt.Fprintf(os.Stdout, "Response from `DatabasesAPI.ListReplicasForDatabase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**databaseId** | **int32** | database_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiListReplicasForDatabaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | current page of results for pagination | 

### Return type

[**ListDatabasesForAccount200Response**](ListDatabasesForAccount200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchDatabase

> PatchDatabase(ctx, id).UpdateDatabaseRequest(updateDatabaseRequest).Execute()

update database

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
	updateDatabaseRequest := *openapiclient.NewUpdateDatabaseRequest() // UpdateDatabaseRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DatabasesAPI.PatchDatabase(context.Background(), id).UpdateDatabaseRequest(updateDatabaseRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.PatchDatabase``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPatchDatabaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateDatabaseRequest** | [**UpdateDatabaseRequest**](UpdateDatabaseRequest.md) |  | 

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


## UpdateDatabase

> UpdateDatabase(ctx, id).UpdateDatabaseRequest(updateDatabaseRequest).Execute()

update database

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
	updateDatabaseRequest := *openapiclient.NewUpdateDatabaseRequest() // UpdateDatabaseRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DatabasesAPI.UpdateDatabase(context.Background(), id).UpdateDatabaseRequest(updateDatabaseRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.UpdateDatabase``: %v\n", err)
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

Other parameters are passed through a pointer to a apiUpdateDatabaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateDatabaseRequest** | [**UpdateDatabaseRequest**](UpdateDatabaseRequest.md) |  | 

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

