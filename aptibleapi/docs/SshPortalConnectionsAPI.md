# \SshPortalConnectionsAPI

All URIs are relative to *https://api.aptible.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSshPortalConnection**](SshPortalConnectionsAPI.md#CreateSshPortalConnection) | **Post** /operations/{operation_id}/ssh_portal_connections | create ssh_portal_connection
[**GetSshPortalConnection**](SshPortalConnectionsAPI.md#GetSshPortalConnection) | **Get** /ssh_portal_connections/{id} | show ssh_portal_connection
[**ListSshPortalConnectionsForOperation**](SshPortalConnectionsAPI.md#ListSshPortalConnectionsForOperation) | **Get** /operations/{operation_id}/ssh_portal_connections | list ssh_portal_connections



## CreateSshPortalConnection

> SshPortalConnection CreateSshPortalConnection(ctx, operationId).NoEmbed(noEmbed).Prefer(prefer).CreateSshPortalConnectionRequest(createSshPortalConnectionRequest).Execute()

create ssh_portal_connection



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
	operationId := int32(56) // int32 | operation_id
	noEmbed := true // bool | When true, omits embedded resources from the response. Also triggered when the Prefer: no_sensitive_extras=true header is present. (optional)
	prefer := "prefer_example" // string | When set to no_sensitive_extras=true, omits sensitive fields and embedded resources from the response. (optional)
	createSshPortalConnectionRequest := *openapiclient.NewCreateSshPortalConnectionRequest("SshPublicKey_example") // CreateSshPortalConnectionRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SshPortalConnectionsAPI.CreateSshPortalConnection(context.Background(), operationId).NoEmbed(noEmbed).Prefer(prefer).CreateSshPortalConnectionRequest(createSshPortalConnectionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SshPortalConnectionsAPI.CreateSshPortalConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSshPortalConnection`: SshPortalConnection
	fmt.Fprintf(os.Stdout, "Response from `SshPortalConnectionsAPI.CreateSshPortalConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operationId** | **int32** | operation_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSshPortalConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **noEmbed** | **bool** | When true, omits embedded resources from the response. Also triggered when the Prefer: no_sensitive_extras&#x3D;true header is present. | 
 **prefer** | **string** | When set to no_sensitive_extras&#x3D;true, omits sensitive fields and embedded resources from the response. | 
 **createSshPortalConnectionRequest** | [**CreateSshPortalConnectionRequest**](CreateSshPortalConnectionRequest.md) |  | 

### Return type

[**SshPortalConnection**](SshPortalConnection.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSshPortalConnection

> SshPortalConnection GetSshPortalConnection(ctx, id).NoEmbed(noEmbed).Prefer(prefer).Execute()

show ssh_portal_connection



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
	resp, r, err := apiClient.SshPortalConnectionsAPI.GetSshPortalConnection(context.Background(), id).NoEmbed(noEmbed).Prefer(prefer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SshPortalConnectionsAPI.GetSshPortalConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSshPortalConnection`: SshPortalConnection
	fmt.Fprintf(os.Stdout, "Response from `SshPortalConnectionsAPI.GetSshPortalConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSshPortalConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **noEmbed** | **bool** | When true, omits embedded resources from the response. Also triggered when the Prefer: no_sensitive_extras&#x3D;true header is present. | 
 **prefer** | **string** | When set to no_sensitive_extras&#x3D;true, omits sensitive fields and embedded resources from the response. | 

### Return type

[**SshPortalConnection**](SshPortalConnection.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSshPortalConnectionsForOperation

> ListSshPortalConnectionsForOperation200Response ListSshPortalConnectionsForOperation(ctx, operationId).Page(page).PerPage(perPage).NoEmbed(noEmbed).Prefer(prefer).Execute()

list ssh_portal_connections



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
	operationId := int32(56) // int32 | operation_id
	page := int32(56) // int32 | Current page of paginated results (optional)
	perPage := int32(56) // int32 | Number of results to return per page (optional)
	noEmbed := true // bool | When true, omits embedded resources from the response. Also triggered when the Prefer: no_sensitive_extras=true header is present. (optional)
	prefer := "prefer_example" // string | When set to no_sensitive_extras=true, omits sensitive fields and embedded resources from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SshPortalConnectionsAPI.ListSshPortalConnectionsForOperation(context.Background(), operationId).Page(page).PerPage(perPage).NoEmbed(noEmbed).Prefer(prefer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SshPortalConnectionsAPI.ListSshPortalConnectionsForOperation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSshPortalConnectionsForOperation`: ListSshPortalConnectionsForOperation200Response
	fmt.Fprintf(os.Stdout, "Response from `SshPortalConnectionsAPI.ListSshPortalConnectionsForOperation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operationId** | **int32** | operation_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSshPortalConnectionsForOperationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Current page of paginated results | 
 **perPage** | **int32** | Number of results to return per page | 
 **noEmbed** | **bool** | When true, omits embedded resources from the response. Also triggered when the Prefer: no_sensitive_extras&#x3D;true header is present. | 
 **prefer** | **string** | When set to no_sensitive_extras&#x3D;true, omits sensitive fields and embedded resources from the response. | 

### Return type

[**ListSshPortalConnectionsForOperation200Response**](ListSshPortalConnectionsForOperation200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

