# \BackupsAPI

All URIs are relative to *https://api.aptible.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBackup**](BackupsAPI.md#GetBackup) | **Get** /backups/{id} | show backup
[**ListBackups**](BackupsAPI.md#ListBackups) | **Get** /backups | list backups
[**ListBackupsForAccount**](BackupsAPI.md#ListBackupsForAccount) | **Get** /accounts/{account_id}/backups | list backups
[**ListBackupsForDatabase**](BackupsAPI.md#ListBackupsForDatabase) | **Get** /databases/{database_id}/backups | list backups
[**ListCopiesForBackup**](BackupsAPI.md#ListCopiesForBackup) | **Get** /backups/{backup_id}/copies | list backups



## GetBackup

> Backup GetBackup(ctx, id).Execute()

show backup

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
	resp, r, err := apiClient.BackupsAPI.GetBackup(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupsAPI.GetBackup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBackup`: Backup
	fmt.Fprintf(os.Stdout, "Response from `BackupsAPI.GetBackup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Backup**](Backup.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBackups

> ListBackups200Response ListBackups(ctx).Page(page).PerPage(perPage).WithDeleted(withDeleted).Execute()

list backups

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
	perPage := int32(56) // int32 | number of results to return per page (optional)
	withDeleted := true // bool | wether deleted records should be returned or not (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupsAPI.ListBackups(context.Background()).Page(page).PerPage(perPage).WithDeleted(withDeleted).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupsAPI.ListBackups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListBackups`: ListBackups200Response
	fmt.Fprintf(os.Stdout, "Response from `BackupsAPI.ListBackups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListBackupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | current page of results for pagination | 
 **perPage** | **int32** | number of results to return per page | 
 **withDeleted** | **bool** | wether deleted records should be returned or not | 

### Return type

[**ListBackups200Response**](ListBackups200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBackupsForAccount

> ListBackups200Response ListBackupsForAccount(ctx, accountId).Page(page).PerPage(perPage).WithDeleted(withDeleted).Execute()

list backups

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
	perPage := int32(56) // int32 | number of results to return per page (optional)
	withDeleted := true // bool | wether deleted records should be returned or not (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupsAPI.ListBackupsForAccount(context.Background(), accountId).Page(page).PerPage(perPage).WithDeleted(withDeleted).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupsAPI.ListBackupsForAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListBackupsForAccount`: ListBackups200Response
	fmt.Fprintf(os.Stdout, "Response from `BackupsAPI.ListBackupsForAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **int32** | account_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiListBackupsForAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | current page of results for pagination | 
 **perPage** | **int32** | number of results to return per page | 
 **withDeleted** | **bool** | wether deleted records should be returned or not | 

### Return type

[**ListBackups200Response**](ListBackups200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBackupsForDatabase

> ListBackups200Response ListBackupsForDatabase(ctx, databaseId).Page(page).PerPage(perPage).WithDeleted(withDeleted).Execute()

list backups

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
	perPage := int32(56) // int32 | number of results to return per page (optional)
	withDeleted := true // bool | wether deleted records should be returned or not (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupsAPI.ListBackupsForDatabase(context.Background(), databaseId).Page(page).PerPage(perPage).WithDeleted(withDeleted).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupsAPI.ListBackupsForDatabase``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListBackupsForDatabase`: ListBackups200Response
	fmt.Fprintf(os.Stdout, "Response from `BackupsAPI.ListBackupsForDatabase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**databaseId** | **int32** | database_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiListBackupsForDatabaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | current page of results for pagination | 
 **perPage** | **int32** | number of results to return per page | 
 **withDeleted** | **bool** | wether deleted records should be returned or not | 

### Return type

[**ListBackups200Response**](ListBackups200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCopiesForBackup

> ListBackups200Response ListCopiesForBackup(ctx, backupId).Page(page).PerPage(perPage).WithDeleted(withDeleted).Execute()

list backups

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
	backupId := int32(56) // int32 | backup_id
	page := int32(56) // int32 | current page of results for pagination (optional)
	perPage := int32(56) // int32 | number of results to return per page (optional)
	withDeleted := true // bool | wether deleted records should be returned or not (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupsAPI.ListCopiesForBackup(context.Background(), backupId).Page(page).PerPage(perPage).WithDeleted(withDeleted).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupsAPI.ListCopiesForBackup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCopiesForBackup`: ListBackups200Response
	fmt.Fprintf(os.Stdout, "Response from `BackupsAPI.ListCopiesForBackup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**backupId** | **int32** | backup_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCopiesForBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | current page of results for pagination | 
 **perPage** | **int32** | number of results to return per page | 
 **withDeleted** | **bool** | wether deleted records should be returned or not | 

### Return type

[**ListBackups200Response**](ListBackups200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

