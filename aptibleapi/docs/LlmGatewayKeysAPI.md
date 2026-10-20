# \LlmGatewayKeysAPI

All URIs are relative to *https://api.aptible.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateLlmGatewayKey**](LlmGatewayKeysAPI.md#CreateLlmGatewayKey) | **Post** /llm_gateway_keys | create llm_gateway_key
[**GetLlmGatewayKey**](LlmGatewayKeysAPI.md#GetLlmGatewayKey) | **Get** /llm_gateway_keys/{id} | show llm_gateway_key
[**ListLlmGatewayKeys**](LlmGatewayKeysAPI.md#ListLlmGatewayKeys) | **Get** /llm_gateway_keys | list llm_gateway_keys
[**RevokeLlmGatewayKey**](LlmGatewayKeysAPI.md#RevokeLlmGatewayKey) | **Delete** /llm_gateway_keys/{id} | revoke llm_gateway_key
[**UpdateLlmGatewayKey**](LlmGatewayKeysAPI.md#UpdateLlmGatewayKey) | **Patch** /llm_gateway_keys/{id} | update llm_gateway_key



## CreateLlmGatewayKey

> LlmGatewayKey CreateLlmGatewayKey(ctx).NoEmbed(noEmbed).Prefer(prefer).CreateLlmGatewayKeyRequest(createLlmGatewayKeyRequest).Execute()

create llm_gateway_key



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
	noEmbed := true // bool | When true, omits embedded resources from the response. Also triggered when the Prefer: no_sensitive_extras=true header is present. (optional)
	prefer := "prefer_example" // string | When set to no_sensitive_extras=true, omits sensitive fields and embedded resources from the response. (optional)
	createLlmGatewayKeyRequest := *openapiclient.NewCreateLlmGatewayKeyRequest("Handle_example") // CreateLlmGatewayKeyRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LlmGatewayKeysAPI.CreateLlmGatewayKey(context.Background()).NoEmbed(noEmbed).Prefer(prefer).CreateLlmGatewayKeyRequest(createLlmGatewayKeyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LlmGatewayKeysAPI.CreateLlmGatewayKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateLlmGatewayKey`: LlmGatewayKey
	fmt.Fprintf(os.Stdout, "Response from `LlmGatewayKeysAPI.CreateLlmGatewayKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateLlmGatewayKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **noEmbed** | **bool** | When true, omits embedded resources from the response. Also triggered when the Prefer: no_sensitive_extras&#x3D;true header is present. | 
 **prefer** | **string** | When set to no_sensitive_extras&#x3D;true, omits sensitive fields and embedded resources from the response. | 
 **createLlmGatewayKeyRequest** | [**CreateLlmGatewayKeyRequest**](CreateLlmGatewayKeyRequest.md) |  | 

### Return type

[**LlmGatewayKey**](LlmGatewayKey.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLlmGatewayKey

> LlmGatewayKey GetLlmGatewayKey(ctx, id).NoEmbed(noEmbed).Prefer(prefer).Execute()

show llm_gateway_key



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
	id := int32(56) // int32 | LLM Gateway Key database ID
	noEmbed := true // bool | When true, omits embedded resources from the response. Also triggered when the Prefer: no_sensitive_extras=true header is present. (optional)
	prefer := "prefer_example" // string | When set to no_sensitive_extras=true, omits sensitive fields and embedded resources from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LlmGatewayKeysAPI.GetLlmGatewayKey(context.Background(), id).NoEmbed(noEmbed).Prefer(prefer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LlmGatewayKeysAPI.GetLlmGatewayKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLlmGatewayKey`: LlmGatewayKey
	fmt.Fprintf(os.Stdout, "Response from `LlmGatewayKeysAPI.GetLlmGatewayKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | LLM Gateway Key database ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLlmGatewayKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **noEmbed** | **bool** | When true, omits embedded resources from the response. Also triggered when the Prefer: no_sensitive_extras&#x3D;true header is present. | 
 **prefer** | **string** | When set to no_sensitive_extras&#x3D;true, omits sensitive fields and embedded resources from the response. | 

### Return type

[**LlmGatewayKey**](LlmGatewayKey.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLlmGatewayKeys

> ListLlmGatewayKeys200Response ListLlmGatewayKeys(ctx).OrganizationId(organizationId).Page(page).PerPage(perPage).NoEmbed(noEmbed).Prefer(prefer).Execute()

list llm_gateway_keys



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
	organizationId := "organizationId_example" // string | Filter by organization ID (optional)
	page := int32(56) // int32 | Current page of paginated results (optional)
	perPage := int32(56) // int32 | Number of results to return per page (optional)
	noEmbed := true // bool | When true, omits embedded resources from the response. Also triggered when the Prefer: no_sensitive_extras=true header is present. (optional)
	prefer := "prefer_example" // string | When set to no_sensitive_extras=true, omits sensitive fields and embedded resources from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LlmGatewayKeysAPI.ListLlmGatewayKeys(context.Background()).OrganizationId(organizationId).Page(page).PerPage(perPage).NoEmbed(noEmbed).Prefer(prefer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LlmGatewayKeysAPI.ListLlmGatewayKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListLlmGatewayKeys`: ListLlmGatewayKeys200Response
	fmt.Fprintf(os.Stdout, "Response from `LlmGatewayKeysAPI.ListLlmGatewayKeys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListLlmGatewayKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organizationId** | **string** | Filter by organization ID | 
 **page** | **int32** | Current page of paginated results | 
 **perPage** | **int32** | Number of results to return per page | 
 **noEmbed** | **bool** | When true, omits embedded resources from the response. Also triggered when the Prefer: no_sensitive_extras&#x3D;true header is present. | 
 **prefer** | **string** | When set to no_sensitive_extras&#x3D;true, omits sensitive fields and embedded resources from the response. | 

### Return type

[**ListLlmGatewayKeys200Response**](ListLlmGatewayKeys200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RevokeLlmGatewayKey

> LlmGatewayKey RevokeLlmGatewayKey(ctx, id).NoEmbed(noEmbed).Prefer(prefer).Execute()

revoke llm_gateway_key



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
	id := int32(56) // int32 | LLM Gateway Key database ID
	noEmbed := true // bool | When true, omits embedded resources from the response. Also triggered when the Prefer: no_sensitive_extras=true header is present. (optional)
	prefer := "prefer_example" // string | When set to no_sensitive_extras=true, omits sensitive fields and embedded resources from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LlmGatewayKeysAPI.RevokeLlmGatewayKey(context.Background(), id).NoEmbed(noEmbed).Prefer(prefer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LlmGatewayKeysAPI.RevokeLlmGatewayKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RevokeLlmGatewayKey`: LlmGatewayKey
	fmt.Fprintf(os.Stdout, "Response from `LlmGatewayKeysAPI.RevokeLlmGatewayKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | LLM Gateway Key database ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRevokeLlmGatewayKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **noEmbed** | **bool** | When true, omits embedded resources from the response. Also triggered when the Prefer: no_sensitive_extras&#x3D;true header is present. | 
 **prefer** | **string** | When set to no_sensitive_extras&#x3D;true, omits sensitive fields and embedded resources from the response. | 

### Return type

[**LlmGatewayKey**](LlmGatewayKey.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLlmGatewayKey

> LlmGatewayKey UpdateLlmGatewayKey(ctx, id).NoEmbed(noEmbed).Prefer(prefer).UpdateLlmGatewayKeyRequest(updateLlmGatewayKeyRequest).Execute()

update llm_gateway_key



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
	id := int32(56) // int32 | LLM Gateway Key database ID
	noEmbed := true // bool | When true, omits embedded resources from the response. Also triggered when the Prefer: no_sensitive_extras=true header is present. (optional)
	prefer := "prefer_example" // string | When set to no_sensitive_extras=true, omits sensitive fields and embedded resources from the response. (optional)
	updateLlmGatewayKeyRequest := *openapiclient.NewUpdateLlmGatewayKeyRequest() // UpdateLlmGatewayKeyRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LlmGatewayKeysAPI.UpdateLlmGatewayKey(context.Background(), id).NoEmbed(noEmbed).Prefer(prefer).UpdateLlmGatewayKeyRequest(updateLlmGatewayKeyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LlmGatewayKeysAPI.UpdateLlmGatewayKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateLlmGatewayKey`: LlmGatewayKey
	fmt.Fprintf(os.Stdout, "Response from `LlmGatewayKeysAPI.UpdateLlmGatewayKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | LLM Gateway Key database ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLlmGatewayKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **noEmbed** | **bool** | When true, omits embedded resources from the response. Also triggered when the Prefer: no_sensitive_extras&#x3D;true header is present. | 
 **prefer** | **string** | When set to no_sensitive_extras&#x3D;true, omits sensitive fields and embedded resources from the response. | 
 **updateLlmGatewayKeyRequest** | [**UpdateLlmGatewayKeyRequest**](UpdateLlmGatewayKeyRequest.md) |  | 

### Return type

[**LlmGatewayKey**](LlmGatewayKey.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

