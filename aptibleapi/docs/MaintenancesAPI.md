# \MaintenancesAPI

All URIs are relative to *https://api.aptible.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMaintenanceApps**](MaintenancesAPI.md#GetMaintenanceApps) | **Get** /maintenances/apps | list apps with maintenance info
[**GetMaintenanceDatabases**](MaintenancesAPI.md#GetMaintenanceDatabases) | **Get** /maintenances/databases | list databases with maintenance info



## GetMaintenanceApps

> GetMaintenanceApps200Response GetMaintenanceApps(ctx).Page(page).PerPage(perPage).Execute()

list apps with maintenance info



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
	page := int32(56) // int32 | Current page of paginated results (optional)
	perPage := int32(56) // int32 | Number of results to return per page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MaintenancesAPI.GetMaintenanceApps(context.Background()).Page(page).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MaintenancesAPI.GetMaintenanceApps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMaintenanceApps`: GetMaintenanceApps200Response
	fmt.Fprintf(os.Stdout, "Response from `MaintenancesAPI.GetMaintenanceApps`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMaintenanceAppsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Current page of paginated results | 
 **perPage** | **int32** | Number of results to return per page | 

### Return type

[**GetMaintenanceApps200Response**](GetMaintenanceApps200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMaintenanceDatabases

> GetMaintenanceDatabases200Response GetMaintenanceDatabases(ctx).Page(page).PerPage(perPage).Execute()

list databases with maintenance info



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
	page := int32(56) // int32 | Current page of paginated results (optional)
	perPage := int32(56) // int32 | Number of results to return per page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MaintenancesAPI.GetMaintenanceDatabases(context.Background()).Page(page).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MaintenancesAPI.GetMaintenanceDatabases``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMaintenanceDatabases`: GetMaintenanceDatabases200Response
	fmt.Fprintf(os.Stdout, "Response from `MaintenancesAPI.GetMaintenanceDatabases`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMaintenanceDatabasesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Current page of paginated results | 
 **perPage** | **int32** | Number of results to return per page | 

### Return type

[**GetMaintenanceDatabases200Response**](GetMaintenanceDatabases200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

