# \ConfigurationsAPI

All URIs are relative to *https://api.aptible.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateConfigurationForApp**](ConfigurationsAPI.md#CreateConfigurationForApp) | **Post** /apps/{app_id}/configurations | create configuration
[**CreateConfigurationForDatabase**](ConfigurationsAPI.md#CreateConfigurationForDatabase) | **Post** /databases/{database_id}/configurations | create configuration
[**DeleteConfiguration**](ConfigurationsAPI.md#DeleteConfiguration) | **Delete** /configurations/{id} | delete configuration
[**GetConfiguration**](ConfigurationsAPI.md#GetConfiguration) | **Get** /configurations/{id} | show configuration
[**ListConfigurationsForApp**](ConfigurationsAPI.md#ListConfigurationsForApp) | **Get** /apps/{app_id}/configurations | list configurations
[**ListConfigurationsForDatabase**](ConfigurationsAPI.md#ListConfigurationsForDatabase) | **Get** /databases/{database_id}/configurations | list configurations



## CreateConfigurationForApp

> Configuration CreateConfigurationForApp(ctx, appId).CreateConfigurationForAppRequest(createConfigurationForAppRequest).Execute()

create configuration

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
	appId := int32(56) // int32 | app_id
	createConfigurationForAppRequest := *openapiclient.NewCreateConfigurationForAppRequest(map[string]interface{}(123)) // CreateConfigurationForAppRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationsAPI.CreateConfigurationForApp(context.Background(), appId).CreateConfigurationForAppRequest(createConfigurationForAppRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationsAPI.CreateConfigurationForApp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateConfigurationForApp`: Configuration
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationsAPI.CreateConfigurationForApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** | app_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateConfigurationForAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createConfigurationForAppRequest** | [**CreateConfigurationForAppRequest**](CreateConfigurationForAppRequest.md) |  | 

### Return type

[**Configuration**](Configuration.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateConfigurationForDatabase

> Configuration CreateConfigurationForDatabase(ctx, databaseId).CreateConfigurationForAppRequest(createConfigurationForAppRequest).Execute()

create configuration

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
	createConfigurationForAppRequest := *openapiclient.NewCreateConfigurationForAppRequest(map[string]interface{}(123)) // CreateConfigurationForAppRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationsAPI.CreateConfigurationForDatabase(context.Background(), databaseId).CreateConfigurationForAppRequest(createConfigurationForAppRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationsAPI.CreateConfigurationForDatabase``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateConfigurationForDatabase`: Configuration
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationsAPI.CreateConfigurationForDatabase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**databaseId** | **int32** | database_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateConfigurationForDatabaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createConfigurationForAppRequest** | [**CreateConfigurationForAppRequest**](CreateConfigurationForAppRequest.md) |  | 

### Return type

[**Configuration**](Configuration.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteConfiguration

> DeleteConfiguration(ctx, id).Execute()

delete configuration

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
	r, err := apiClient.ConfigurationsAPI.DeleteConfiguration(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationsAPI.DeleteConfiguration``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteConfigurationRequest struct via the builder pattern


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


## GetConfiguration

> Configuration GetConfiguration(ctx, id).Execute()

show configuration

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
	resp, r, err := apiClient.ConfigurationsAPI.GetConfiguration(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationsAPI.GetConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConfiguration`: Configuration
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationsAPI.GetConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Configuration**](Configuration.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListConfigurationsForApp

> ListConfigurationsForApp200Response ListConfigurationsForApp(ctx, appId).Page(page).Execute()

list configurations

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
	appId := int32(56) // int32 | app_id
	page := int32(56) // int32 | current page of results for pagination (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationsAPI.ListConfigurationsForApp(context.Background(), appId).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationsAPI.ListConfigurationsForApp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListConfigurationsForApp`: ListConfigurationsForApp200Response
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationsAPI.ListConfigurationsForApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** | app_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiListConfigurationsForAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | current page of results for pagination | 

### Return type

[**ListConfigurationsForApp200Response**](ListConfigurationsForApp200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListConfigurationsForDatabase

> ListConfigurationsForApp200Response ListConfigurationsForDatabase(ctx, databaseId).Page(page).Execute()

list configurations

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
	resp, r, err := apiClient.ConfigurationsAPI.ListConfigurationsForDatabase(context.Background(), databaseId).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationsAPI.ListConfigurationsForDatabase``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListConfigurationsForDatabase`: ListConfigurationsForApp200Response
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationsAPI.ListConfigurationsForDatabase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**databaseId** | **int32** | database_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiListConfigurationsForDatabaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | current page of results for pagination | 

### Return type

[**ListConfigurationsForApp200Response**](ListConfigurationsForApp200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

