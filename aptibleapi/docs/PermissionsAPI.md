# \PermissionsAPI

All URIs are relative to *https://api.aptible.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePermission**](PermissionsAPI.md#CreatePermission) | **Post** /accounts/{account_id}/permissions | create permission
[**DeletePermission**](PermissionsAPI.md#DeletePermission) | **Delete** /permissions/{id} | delete permission
[**GetPermission**](PermissionsAPI.md#GetPermission) | **Get** /permissions/{id} | show permission
[**ListPermissions**](PermissionsAPI.md#ListPermissions) | **Get** /permissions | list permissions
[**ListPermissionsForAccount**](PermissionsAPI.md#ListPermissionsForAccount) | **Get** /accounts/{account_id}/permissions | list permissions



## CreatePermission

> Permission CreatePermission(ctx, accountId).NoEmbed(noEmbed).Prefer(prefer).CreatePermissionRequest(createPermissionRequest).Execute()

create permission



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
	noEmbed := true // bool | When true, omits embedded resources from the response. Also triggered when the Prefer: no_sensitive_extras=true header is present. (optional)
	prefer := "prefer_example" // string | When set to no_sensitive_extras=true, omits sensitive fields and embedded resources from the response. (optional)
	createPermissionRequest := *openapiclient.NewCreatePermissionRequest("Role_example", "Scope_example") // CreatePermissionRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PermissionsAPI.CreatePermission(context.Background(), accountId).NoEmbed(noEmbed).Prefer(prefer).CreatePermissionRequest(createPermissionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PermissionsAPI.CreatePermission``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePermission`: Permission
	fmt.Fprintf(os.Stdout, "Response from `PermissionsAPI.CreatePermission`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **int32** | account_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePermissionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **noEmbed** | **bool** | When true, omits embedded resources from the response. Also triggered when the Prefer: no_sensitive_extras&#x3D;true header is present. | 
 **prefer** | **string** | When set to no_sensitive_extras&#x3D;true, omits sensitive fields and embedded resources from the response. | 
 **createPermissionRequest** | [**CreatePermissionRequest**](CreatePermissionRequest.md) |  | 

### Return type

[**Permission**](Permission.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePermission

> DeletePermission(ctx, id).Execute()

delete permission



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
	r, err := apiClient.PermissionsAPI.DeletePermission(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PermissionsAPI.DeletePermission``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeletePermissionRequest struct via the builder pattern


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


## GetPermission

> Permission GetPermission(ctx, id).NoEmbed(noEmbed).Prefer(prefer).Execute()

show permission



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
	resp, r, err := apiClient.PermissionsAPI.GetPermission(context.Background(), id).NoEmbed(noEmbed).Prefer(prefer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PermissionsAPI.GetPermission``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPermission`: Permission
	fmt.Fprintf(os.Stdout, "Response from `PermissionsAPI.GetPermission`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPermissionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **noEmbed** | **bool** | When true, omits embedded resources from the response. Also triggered when the Prefer: no_sensitive_extras&#x3D;true header is present. | 
 **prefer** | **string** | When set to no_sensitive_extras&#x3D;true, omits sensitive fields and embedded resources from the response. | 

### Return type

[**Permission**](Permission.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPermissions

> ListPermissionsForAccount200Response ListPermissions(ctx).Page(page).PerPage(perPage).NoEmbed(noEmbed).Prefer(prefer).Execute()

list permissions



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
	noEmbed := true // bool | When true, omits embedded resources from the response. Also triggered when the Prefer: no_sensitive_extras=true header is present. (optional)
	prefer := "prefer_example" // string | When set to no_sensitive_extras=true, omits sensitive fields and embedded resources from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PermissionsAPI.ListPermissions(context.Background()).Page(page).PerPage(perPage).NoEmbed(noEmbed).Prefer(prefer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PermissionsAPI.ListPermissions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPermissions`: ListPermissionsForAccount200Response
	fmt.Fprintf(os.Stdout, "Response from `PermissionsAPI.ListPermissions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Current page of paginated results | 
 **perPage** | **int32** | Number of results to return per page | 
 **noEmbed** | **bool** | When true, omits embedded resources from the response. Also triggered when the Prefer: no_sensitive_extras&#x3D;true header is present. | 
 **prefer** | **string** | When set to no_sensitive_extras&#x3D;true, omits sensitive fields and embedded resources from the response. | 

### Return type

[**ListPermissionsForAccount200Response**](ListPermissionsForAccount200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPermissionsForAccount

> ListPermissionsForAccount200Response ListPermissionsForAccount(ctx, accountId).Page(page).PerPage(perPage).NoEmbed(noEmbed).Prefer(prefer).Execute()

list permissions



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
	page := int32(56) // int32 | Current page of paginated results (optional)
	perPage := int32(56) // int32 | Number of results to return per page (optional)
	noEmbed := true // bool | When true, omits embedded resources from the response. Also triggered when the Prefer: no_sensitive_extras=true header is present. (optional)
	prefer := "prefer_example" // string | When set to no_sensitive_extras=true, omits sensitive fields and embedded resources from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PermissionsAPI.ListPermissionsForAccount(context.Background(), accountId).Page(page).PerPage(perPage).NoEmbed(noEmbed).Prefer(prefer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PermissionsAPI.ListPermissionsForAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPermissionsForAccount`: ListPermissionsForAccount200Response
	fmt.Fprintf(os.Stdout, "Response from `PermissionsAPI.ListPermissionsForAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **int32** | account_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiListPermissionsForAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Current page of paginated results | 
 **perPage** | **int32** | Number of results to return per page | 
 **noEmbed** | **bool** | When true, omits embedded resources from the response. Also triggered when the Prefer: no_sensitive_extras&#x3D;true header is present. | 
 **prefer** | **string** | When set to no_sensitive_extras&#x3D;true, omits sensitive fields and embedded resources from the response. | 

### Return type

[**ListPermissionsForAccount200Response**](ListPermissionsForAccount200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

