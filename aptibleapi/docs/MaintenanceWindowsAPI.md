# \MaintenanceWindowsAPI

All URIs are relative to *https://api.aptible.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMaintenanceWindowForApp**](MaintenanceWindowsAPI.md#CreateMaintenanceWindowForApp) | **Post** /apps/{app_id}/maintenance_window | create maintenance_window
[**CreateMaintenanceWindowForDatabase**](MaintenanceWindowsAPI.md#CreateMaintenanceWindowForDatabase) | **Post** /databases/{database_id}/maintenance_window | create maintenance_window
[**DeleteMaintenanceWindowForApp**](MaintenanceWindowsAPI.md#DeleteMaintenanceWindowForApp) | **Delete** /apps/{app_id}/maintenance_window | delete maintenance_window
[**DeleteMaintenanceWindowForDatabase**](MaintenanceWindowsAPI.md#DeleteMaintenanceWindowForDatabase) | **Delete** /databases/{database_id}/maintenance_window | delete maintenance_window
[**GetMaintenanceWindow**](MaintenanceWindowsAPI.md#GetMaintenanceWindow) | **Get** /maintenance_windows/{id} | show maintenance_window
[**GetMaintenanceWindowForApp**](MaintenanceWindowsAPI.md#GetMaintenanceWindowForApp) | **Get** /apps/{app_id}/maintenance_window | show maintenance_window
[**GetMaintenanceWindowForDatabase**](MaintenanceWindowsAPI.md#GetMaintenanceWindowForDatabase) | **Get** /databases/{database_id}/maintenance_window | show maintenance_window
[**PatchMaintenanceWindowForApp**](MaintenanceWindowsAPI.md#PatchMaintenanceWindowForApp) | **Patch** /apps/{app_id}/maintenance_window | update maintenance_window
[**PatchMaintenanceWindowForDatabase**](MaintenanceWindowsAPI.md#PatchMaintenanceWindowForDatabase) | **Patch** /databases/{database_id}/maintenance_window | update maintenance_window
[**UpdateMaintenanceWindowForApp**](MaintenanceWindowsAPI.md#UpdateMaintenanceWindowForApp) | **Put** /apps/{app_id}/maintenance_window | update maintenance_window
[**UpdateMaintenanceWindowForDatabase**](MaintenanceWindowsAPI.md#UpdateMaintenanceWindowForDatabase) | **Put** /databases/{database_id}/maintenance_window | update maintenance_window



## CreateMaintenanceWindowForApp

> MaintenanceWindow CreateMaintenanceWindowForApp(ctx, appId).NoEmbed(noEmbed).Prefer(prefer).CreateMaintenanceWindowForAppRequest(createMaintenanceWindowForAppRequest).Execute()

create maintenance_window



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
	noEmbed := true // bool | When true, omits embedded resources from the response. Also triggered when the Prefer: no_sensitive_extras=true header is present. (optional)
	prefer := "prefer_example" // string | When set to no_sensitive_extras=true, omits sensitive fields and embedded resources from the response. (optional)
	createMaintenanceWindowForAppRequest := *openapiclient.NewCreateMaintenanceWindowForAppRequest("StartTimeUtc_example", int32(123)) // CreateMaintenanceWindowForAppRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MaintenanceWindowsAPI.CreateMaintenanceWindowForApp(context.Background(), appId).NoEmbed(noEmbed).Prefer(prefer).CreateMaintenanceWindowForAppRequest(createMaintenanceWindowForAppRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MaintenanceWindowsAPI.CreateMaintenanceWindowForApp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateMaintenanceWindowForApp`: MaintenanceWindow
	fmt.Fprintf(os.Stdout, "Response from `MaintenanceWindowsAPI.CreateMaintenanceWindowForApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** | app_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMaintenanceWindowForAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **noEmbed** | **bool** | When true, omits embedded resources from the response. Also triggered when the Prefer: no_sensitive_extras&#x3D;true header is present. | 
 **prefer** | **string** | When set to no_sensitive_extras&#x3D;true, omits sensitive fields and embedded resources from the response. | 
 **createMaintenanceWindowForAppRequest** | [**CreateMaintenanceWindowForAppRequest**](CreateMaintenanceWindowForAppRequest.md) |  | 

### Return type

[**MaintenanceWindow**](MaintenanceWindow.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMaintenanceWindowForDatabase

> MaintenanceWindow CreateMaintenanceWindowForDatabase(ctx, databaseId).NoEmbed(noEmbed).Prefer(prefer).CreateMaintenanceWindowForAppRequest(createMaintenanceWindowForAppRequest).Execute()

create maintenance_window



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
	noEmbed := true // bool | When true, omits embedded resources from the response. Also triggered when the Prefer: no_sensitive_extras=true header is present. (optional)
	prefer := "prefer_example" // string | When set to no_sensitive_extras=true, omits sensitive fields and embedded resources from the response. (optional)
	createMaintenanceWindowForAppRequest := *openapiclient.NewCreateMaintenanceWindowForAppRequest("StartTimeUtc_example", int32(123)) // CreateMaintenanceWindowForAppRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MaintenanceWindowsAPI.CreateMaintenanceWindowForDatabase(context.Background(), databaseId).NoEmbed(noEmbed).Prefer(prefer).CreateMaintenanceWindowForAppRequest(createMaintenanceWindowForAppRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MaintenanceWindowsAPI.CreateMaintenanceWindowForDatabase``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateMaintenanceWindowForDatabase`: MaintenanceWindow
	fmt.Fprintf(os.Stdout, "Response from `MaintenanceWindowsAPI.CreateMaintenanceWindowForDatabase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**databaseId** | **int32** | database_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMaintenanceWindowForDatabaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **noEmbed** | **bool** | When true, omits embedded resources from the response. Also triggered when the Prefer: no_sensitive_extras&#x3D;true header is present. | 
 **prefer** | **string** | When set to no_sensitive_extras&#x3D;true, omits sensitive fields and embedded resources from the response. | 
 **createMaintenanceWindowForAppRequest** | [**CreateMaintenanceWindowForAppRequest**](CreateMaintenanceWindowForAppRequest.md) |  | 

### Return type

[**MaintenanceWindow**](MaintenanceWindow.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMaintenanceWindowForApp

> DeleteMaintenanceWindowForApp(ctx, appId).Execute()

delete maintenance_window



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MaintenanceWindowsAPI.DeleteMaintenanceWindowForApp(context.Background(), appId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MaintenanceWindowsAPI.DeleteMaintenanceWindowForApp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** | app_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMaintenanceWindowForAppRequest struct via the builder pattern


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


## DeleteMaintenanceWindowForDatabase

> DeleteMaintenanceWindowForDatabase(ctx, databaseId).Execute()

delete maintenance_window



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MaintenanceWindowsAPI.DeleteMaintenanceWindowForDatabase(context.Background(), databaseId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MaintenanceWindowsAPI.DeleteMaintenanceWindowForDatabase``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**databaseId** | **int32** | database_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMaintenanceWindowForDatabaseRequest struct via the builder pattern


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


## GetMaintenanceWindow

> MaintenanceWindow GetMaintenanceWindow(ctx, id).NoEmbed(noEmbed).Prefer(prefer).Execute()

show maintenance_window



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
	resp, r, err := apiClient.MaintenanceWindowsAPI.GetMaintenanceWindow(context.Background(), id).NoEmbed(noEmbed).Prefer(prefer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MaintenanceWindowsAPI.GetMaintenanceWindow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMaintenanceWindow`: MaintenanceWindow
	fmt.Fprintf(os.Stdout, "Response from `MaintenanceWindowsAPI.GetMaintenanceWindow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMaintenanceWindowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **noEmbed** | **bool** | When true, omits embedded resources from the response. Also triggered when the Prefer: no_sensitive_extras&#x3D;true header is present. | 
 **prefer** | **string** | When set to no_sensitive_extras&#x3D;true, omits sensitive fields and embedded resources from the response. | 

### Return type

[**MaintenanceWindow**](MaintenanceWindow.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMaintenanceWindowForApp

> MaintenanceWindow GetMaintenanceWindowForApp(ctx, appId).NoEmbed(noEmbed).Prefer(prefer).Execute()

show maintenance_window



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
	noEmbed := true // bool | When true, omits embedded resources from the response. Also triggered when the Prefer: no_sensitive_extras=true header is present. (optional)
	prefer := "prefer_example" // string | When set to no_sensitive_extras=true, omits sensitive fields and embedded resources from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MaintenanceWindowsAPI.GetMaintenanceWindowForApp(context.Background(), appId).NoEmbed(noEmbed).Prefer(prefer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MaintenanceWindowsAPI.GetMaintenanceWindowForApp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMaintenanceWindowForApp`: MaintenanceWindow
	fmt.Fprintf(os.Stdout, "Response from `MaintenanceWindowsAPI.GetMaintenanceWindowForApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** | app_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMaintenanceWindowForAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **noEmbed** | **bool** | When true, omits embedded resources from the response. Also triggered when the Prefer: no_sensitive_extras&#x3D;true header is present. | 
 **prefer** | **string** | When set to no_sensitive_extras&#x3D;true, omits sensitive fields and embedded resources from the response. | 

### Return type

[**MaintenanceWindow**](MaintenanceWindow.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMaintenanceWindowForDatabase

> MaintenanceWindow GetMaintenanceWindowForDatabase(ctx, databaseId).NoEmbed(noEmbed).Prefer(prefer).Execute()

show maintenance_window



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
	noEmbed := true // bool | When true, omits embedded resources from the response. Also triggered when the Prefer: no_sensitive_extras=true header is present. (optional)
	prefer := "prefer_example" // string | When set to no_sensitive_extras=true, omits sensitive fields and embedded resources from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MaintenanceWindowsAPI.GetMaintenanceWindowForDatabase(context.Background(), databaseId).NoEmbed(noEmbed).Prefer(prefer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MaintenanceWindowsAPI.GetMaintenanceWindowForDatabase``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMaintenanceWindowForDatabase`: MaintenanceWindow
	fmt.Fprintf(os.Stdout, "Response from `MaintenanceWindowsAPI.GetMaintenanceWindowForDatabase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**databaseId** | **int32** | database_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMaintenanceWindowForDatabaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **noEmbed** | **bool** | When true, omits embedded resources from the response. Also triggered when the Prefer: no_sensitive_extras&#x3D;true header is present. | 
 **prefer** | **string** | When set to no_sensitive_extras&#x3D;true, omits sensitive fields and embedded resources from the response. | 

### Return type

[**MaintenanceWindow**](MaintenanceWindow.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchMaintenanceWindowForApp

> MaintenanceWindow PatchMaintenanceWindowForApp(ctx, appId).NoEmbed(noEmbed).Prefer(prefer).UpdateMaintenanceWindowForAppRequest(updateMaintenanceWindowForAppRequest).Execute()

update maintenance_window



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
	noEmbed := true // bool | When true, omits embedded resources from the response. Also triggered when the Prefer: no_sensitive_extras=true header is present. (optional)
	prefer := "prefer_example" // string | When set to no_sensitive_extras=true, omits sensitive fields and embedded resources from the response. (optional)
	updateMaintenanceWindowForAppRequest := *openapiclient.NewUpdateMaintenanceWindowForAppRequest() // UpdateMaintenanceWindowForAppRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MaintenanceWindowsAPI.PatchMaintenanceWindowForApp(context.Background(), appId).NoEmbed(noEmbed).Prefer(prefer).UpdateMaintenanceWindowForAppRequest(updateMaintenanceWindowForAppRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MaintenanceWindowsAPI.PatchMaintenanceWindowForApp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchMaintenanceWindowForApp`: MaintenanceWindow
	fmt.Fprintf(os.Stdout, "Response from `MaintenanceWindowsAPI.PatchMaintenanceWindowForApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** | app_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchMaintenanceWindowForAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **noEmbed** | **bool** | When true, omits embedded resources from the response. Also triggered when the Prefer: no_sensitive_extras&#x3D;true header is present. | 
 **prefer** | **string** | When set to no_sensitive_extras&#x3D;true, omits sensitive fields and embedded resources from the response. | 
 **updateMaintenanceWindowForAppRequest** | [**UpdateMaintenanceWindowForAppRequest**](UpdateMaintenanceWindowForAppRequest.md) |  | 

### Return type

[**MaintenanceWindow**](MaintenanceWindow.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchMaintenanceWindowForDatabase

> MaintenanceWindow PatchMaintenanceWindowForDatabase(ctx, databaseId).NoEmbed(noEmbed).Prefer(prefer).UpdateMaintenanceWindowForAppRequest(updateMaintenanceWindowForAppRequest).Execute()

update maintenance_window



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
	noEmbed := true // bool | When true, omits embedded resources from the response. Also triggered when the Prefer: no_sensitive_extras=true header is present. (optional)
	prefer := "prefer_example" // string | When set to no_sensitive_extras=true, omits sensitive fields and embedded resources from the response. (optional)
	updateMaintenanceWindowForAppRequest := *openapiclient.NewUpdateMaintenanceWindowForAppRequest() // UpdateMaintenanceWindowForAppRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MaintenanceWindowsAPI.PatchMaintenanceWindowForDatabase(context.Background(), databaseId).NoEmbed(noEmbed).Prefer(prefer).UpdateMaintenanceWindowForAppRequest(updateMaintenanceWindowForAppRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MaintenanceWindowsAPI.PatchMaintenanceWindowForDatabase``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchMaintenanceWindowForDatabase`: MaintenanceWindow
	fmt.Fprintf(os.Stdout, "Response from `MaintenanceWindowsAPI.PatchMaintenanceWindowForDatabase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**databaseId** | **int32** | database_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchMaintenanceWindowForDatabaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **noEmbed** | **bool** | When true, omits embedded resources from the response. Also triggered when the Prefer: no_sensitive_extras&#x3D;true header is present. | 
 **prefer** | **string** | When set to no_sensitive_extras&#x3D;true, omits sensitive fields and embedded resources from the response. | 
 **updateMaintenanceWindowForAppRequest** | [**UpdateMaintenanceWindowForAppRequest**](UpdateMaintenanceWindowForAppRequest.md) |  | 

### Return type

[**MaintenanceWindow**](MaintenanceWindow.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMaintenanceWindowForApp

> MaintenanceWindow UpdateMaintenanceWindowForApp(ctx, appId).NoEmbed(noEmbed).Prefer(prefer).UpdateMaintenanceWindowForAppRequest(updateMaintenanceWindowForAppRequest).Execute()

update maintenance_window



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
	noEmbed := true // bool | When true, omits embedded resources from the response. Also triggered when the Prefer: no_sensitive_extras=true header is present. (optional)
	prefer := "prefer_example" // string | When set to no_sensitive_extras=true, omits sensitive fields and embedded resources from the response. (optional)
	updateMaintenanceWindowForAppRequest := *openapiclient.NewUpdateMaintenanceWindowForAppRequest() // UpdateMaintenanceWindowForAppRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MaintenanceWindowsAPI.UpdateMaintenanceWindowForApp(context.Background(), appId).NoEmbed(noEmbed).Prefer(prefer).UpdateMaintenanceWindowForAppRequest(updateMaintenanceWindowForAppRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MaintenanceWindowsAPI.UpdateMaintenanceWindowForApp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateMaintenanceWindowForApp`: MaintenanceWindow
	fmt.Fprintf(os.Stdout, "Response from `MaintenanceWindowsAPI.UpdateMaintenanceWindowForApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** | app_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMaintenanceWindowForAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **noEmbed** | **bool** | When true, omits embedded resources from the response. Also triggered when the Prefer: no_sensitive_extras&#x3D;true header is present. | 
 **prefer** | **string** | When set to no_sensitive_extras&#x3D;true, omits sensitive fields and embedded resources from the response. | 
 **updateMaintenanceWindowForAppRequest** | [**UpdateMaintenanceWindowForAppRequest**](UpdateMaintenanceWindowForAppRequest.md) |  | 

### Return type

[**MaintenanceWindow**](MaintenanceWindow.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMaintenanceWindowForDatabase

> MaintenanceWindow UpdateMaintenanceWindowForDatabase(ctx, databaseId).NoEmbed(noEmbed).Prefer(prefer).UpdateMaintenanceWindowForAppRequest(updateMaintenanceWindowForAppRequest).Execute()

update maintenance_window



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
	noEmbed := true // bool | When true, omits embedded resources from the response. Also triggered when the Prefer: no_sensitive_extras=true header is present. (optional)
	prefer := "prefer_example" // string | When set to no_sensitive_extras=true, omits sensitive fields and embedded resources from the response. (optional)
	updateMaintenanceWindowForAppRequest := *openapiclient.NewUpdateMaintenanceWindowForAppRequest() // UpdateMaintenanceWindowForAppRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MaintenanceWindowsAPI.UpdateMaintenanceWindowForDatabase(context.Background(), databaseId).NoEmbed(noEmbed).Prefer(prefer).UpdateMaintenanceWindowForAppRequest(updateMaintenanceWindowForAppRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MaintenanceWindowsAPI.UpdateMaintenanceWindowForDatabase``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateMaintenanceWindowForDatabase`: MaintenanceWindow
	fmt.Fprintf(os.Stdout, "Response from `MaintenanceWindowsAPI.UpdateMaintenanceWindowForDatabase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**databaseId** | **int32** | database_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMaintenanceWindowForDatabaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **noEmbed** | **bool** | When true, omits embedded resources from the response. Also triggered when the Prefer: no_sensitive_extras&#x3D;true header is present. | 
 **prefer** | **string** | When set to no_sensitive_extras&#x3D;true, omits sensitive fields and embedded resources from the response. | 
 **updateMaintenanceWindowForAppRequest** | [**UpdateMaintenanceWindowForAppRequest**](UpdateMaintenanceWindowForAppRequest.md) |  | 

### Return type

[**MaintenanceWindow**](MaintenanceWindow.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

