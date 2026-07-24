# \LlmKeysAPI

All URIs are relative to *https://api.aptible.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLlmKey**](LlmKeysAPI.md#GetLlmKey) | **Get** /llm_keys/{id} | show llm_key
[**GetLlmKeyRequests**](LlmKeysAPI.md#GetLlmKeyRequests) | **Get** /llm_keys/{id}/requests | list llm key requests
[**GetLlmKeyUsage**](LlmKeysAPI.md#GetLlmKeyUsage) | **Get** /llm_keys/{id}/usage | get llm key usage
[**ListLlmKeys**](LlmKeysAPI.md#ListLlmKeys) | **Get** /llm_keys | list all llm_keys
[**ListLlmKeysForAccount**](LlmKeysAPI.md#ListLlmKeysForAccount) | **Get** /accounts/{account_id}/llm_keys | list llm_keys
[**RevokeLlmKey**](LlmKeysAPI.md#RevokeLlmKey) | **Delete** /llm_keys/{id} | revoke llm_key



## GetLlmKey

> LlmKey GetLlmKey(ctx, id).NoEmbed(noEmbed).Prefer(prefer).Execute()

show llm_key



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
	id := int32(56) // int32 | LLM Key database ID
	noEmbed := true // bool | When true, omits embedded resources from the response. Also triggered when the Prefer: no_sensitive_extras=true header is present. (optional)
	prefer := "prefer_example" // string | When set to no_sensitive_extras=true, omits sensitive fields and embedded resources from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LlmKeysAPI.GetLlmKey(context.Background(), id).NoEmbed(noEmbed).Prefer(prefer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LlmKeysAPI.GetLlmKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLlmKey`: LlmKey
	fmt.Fprintf(os.Stdout, "Response from `LlmKeysAPI.GetLlmKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | LLM Key database ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLlmKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **noEmbed** | **bool** | When true, omits embedded resources from the response. Also triggered when the Prefer: no_sensitive_extras&#x3D;true header is present. | 
 **prefer** | **string** | When set to no_sensitive_extras&#x3D;true, omits sensitive fields and embedded resources from the response. | 

### Return type

[**LlmKey**](LlmKey.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLlmKeyRequests

> GetLlmKeyRequests200Response GetLlmKeyRequests(ctx, id).StartTime(startTime).EndTime(endTime).Page(page).PerPage(perPage).RequestId(requestId).Status(status).Model(model).SortBy(sortBy).SortOrder(sortOrder).NoEmbed(noEmbed).Prefer(prefer).Execute()

list llm key requests



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
	id := int32(56) // int32 | LLM Key database ID
	startTime := "startTime_example" // string | Start date/time for filtering requests (ISO8601 format, YYYY-MM-DD or YYYY-MM-DDTHH:MM:SS). Cannot be more than 7 days ago.
	endTime := "endTime_example" // string | End date/time for filtering requests (ISO8601 format, YYYY-MM-DD or YYYY-MM-DDTHH:MM:SS)
	page := int32(56) // int32 | Page number for pagination (optional)
	perPage := int32(56) // int32 | Number of results per page (max 50) (optional)
	requestId := "requestId_example" // string | Filter by specific request ID (optional)
	status := "status_example" // string | Filter by request status ('success' or 'failure') (optional)
	model := "model_example" // string | Filter by model name (optional)
	sortBy := "sortBy_example" // string | Sort field: 'spend', 'total_tokens', 'start_time', or 'end_time'. Default: 'start_time' (optional)
	sortOrder := "sortOrder_example" // string | Sort order: 'asc' or 'desc'. Default: 'desc' (optional)
	noEmbed := true // bool | When true, omits embedded resources from the response. Also triggered when the Prefer: no_sensitive_extras=true header is present. (optional)
	prefer := "prefer_example" // string | When set to no_sensitive_extras=true, omits sensitive fields and embedded resources from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LlmKeysAPI.GetLlmKeyRequests(context.Background(), id).StartTime(startTime).EndTime(endTime).Page(page).PerPage(perPage).RequestId(requestId).Status(status).Model(model).SortBy(sortBy).SortOrder(sortOrder).NoEmbed(noEmbed).Prefer(prefer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LlmKeysAPI.GetLlmKeyRequests``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLlmKeyRequests`: GetLlmKeyRequests200Response
	fmt.Fprintf(os.Stdout, "Response from `LlmKeysAPI.GetLlmKeyRequests`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | LLM Key database ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLlmKeyRequestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startTime** | **string** | Start date/time for filtering requests (ISO8601 format, YYYY-MM-DD or YYYY-MM-DDTHH:MM:SS). Cannot be more than 7 days ago. | 
 **endTime** | **string** | End date/time for filtering requests (ISO8601 format, YYYY-MM-DD or YYYY-MM-DDTHH:MM:SS) | 
 **page** | **int32** | Page number for pagination | 
 **perPage** | **int32** | Number of results per page (max 50) | 
 **requestId** | **string** | Filter by specific request ID | 
 **status** | **string** | Filter by request status (&#39;success&#39; or &#39;failure&#39;) | 
 **model** | **string** | Filter by model name | 
 **sortBy** | **string** | Sort field: &#39;spend&#39;, &#39;total_tokens&#39;, &#39;start_time&#39;, or &#39;end_time&#39;. Default: &#39;start_time&#39; | 
 **sortOrder** | **string** | Sort order: &#39;asc&#39; or &#39;desc&#39;. Default: &#39;desc&#39; | 
 **noEmbed** | **bool** | When true, omits embedded resources from the response. Also triggered when the Prefer: no_sensitive_extras&#x3D;true header is present. | 
 **prefer** | **string** | When set to no_sensitive_extras&#x3D;true, omits sensitive fields and embedded resources from the response. | 

### Return type

[**GetLlmKeyRequests200Response**](GetLlmKeyRequests200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLlmKeyUsage

> GetLlmKeyUsage200Response GetLlmKeyUsage(ctx, id).StartDate(startDate).EndDate(endDate).NoEmbed(noEmbed).Prefer(prefer).Execute()

get llm key usage



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
	id := int32(56) // int32 | LLM Key database ID
	startDate := "startDate_example" // string | Start date for the usage period (YYYY-MM-DD)
	endDate := "endDate_example" // string | End date for the usage period (YYYY-MM-DD, exclusive)
	noEmbed := true // bool | When true, omits embedded resources from the response. Also triggered when the Prefer: no_sensitive_extras=true header is present. (optional)
	prefer := "prefer_example" // string | When set to no_sensitive_extras=true, omits sensitive fields and embedded resources from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LlmKeysAPI.GetLlmKeyUsage(context.Background(), id).StartDate(startDate).EndDate(endDate).NoEmbed(noEmbed).Prefer(prefer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LlmKeysAPI.GetLlmKeyUsage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLlmKeyUsage`: GetLlmKeyUsage200Response
	fmt.Fprintf(os.Stdout, "Response from `LlmKeysAPI.GetLlmKeyUsage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | LLM Key database ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLlmKeyUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startDate** | **string** | Start date for the usage period (YYYY-MM-DD) | 
 **endDate** | **string** | End date for the usage period (YYYY-MM-DD, exclusive) | 
 **noEmbed** | **bool** | When true, omits embedded resources from the response. Also triggered when the Prefer: no_sensitive_extras&#x3D;true header is present. | 
 **prefer** | **string** | When set to no_sensitive_extras&#x3D;true, omits sensitive fields and embedded resources from the response. | 

### Return type

[**GetLlmKeyUsage200Response**](GetLlmKeyUsage200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLlmKeys

> ListLlmKeys200Response ListLlmKeys(ctx).Page(page).PerPage(perPage).NoEmbed(noEmbed).Prefer(prefer).Execute()

list all llm_keys



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
	resp, r, err := apiClient.LlmKeysAPI.ListLlmKeys(context.Background()).Page(page).PerPage(perPage).NoEmbed(noEmbed).Prefer(prefer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LlmKeysAPI.ListLlmKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListLlmKeys`: ListLlmKeys200Response
	fmt.Fprintf(os.Stdout, "Response from `LlmKeysAPI.ListLlmKeys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListLlmKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Current page of paginated results | 
 **perPage** | **int32** | Number of results to return per page | 
 **noEmbed** | **bool** | When true, omits embedded resources from the response. Also triggered when the Prefer: no_sensitive_extras&#x3D;true header is present. | 
 **prefer** | **string** | When set to no_sensitive_extras&#x3D;true, omits sensitive fields and embedded resources from the response. | 

### Return type

[**ListLlmKeys200Response**](ListLlmKeys200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLlmKeysForAccount

> ListLlmKeysForAccount200Response ListLlmKeysForAccount(ctx, accountId).NoEmbed(noEmbed).Prefer(prefer).Execute()

list llm_keys



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LlmKeysAPI.ListLlmKeysForAccount(context.Background(), accountId).NoEmbed(noEmbed).Prefer(prefer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LlmKeysAPI.ListLlmKeysForAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListLlmKeysForAccount`: ListLlmKeysForAccount200Response
	fmt.Fprintf(os.Stdout, "Response from `LlmKeysAPI.ListLlmKeysForAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **int32** | account_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiListLlmKeysForAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **noEmbed** | **bool** | When true, omits embedded resources from the response. Also triggered when the Prefer: no_sensitive_extras&#x3D;true header is present. | 
 **prefer** | **string** | When set to no_sensitive_extras&#x3D;true, omits sensitive fields and embedded resources from the response. | 

### Return type

[**ListLlmKeysForAccount200Response**](ListLlmKeysForAccount200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RevokeLlmKey

> LlmKey RevokeLlmKey(ctx, id).Execute()

revoke llm_key



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
	id := int32(56) // int32 | LLM Key database ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LlmKeysAPI.RevokeLlmKey(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LlmKeysAPI.RevokeLlmKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RevokeLlmKey`: LlmKey
	fmt.Fprintf(os.Stdout, "Response from `LlmKeysAPI.RevokeLlmKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | LLM Key database ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRevokeLlmKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LlmKey**](LlmKey.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

