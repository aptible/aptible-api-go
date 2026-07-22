# \LlmPoliciesAPI

All URIs are relative to *https://api.aptible.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrUpdateLlmPolicy**](LlmPoliciesAPI.md#CreateOrUpdateLlmPolicy) | **Post** /accounts/{account_id}/llm_policies/create_or_update | create or update llm_policy
[**GetLlmPolicy**](LlmPoliciesAPI.md#GetLlmPolicy) | **Get** /llm_policies/{id} | show llm_policy
[**GetLlmPolicyUsage**](LlmPoliciesAPI.md#GetLlmPolicyUsage) | **Get** /llm_policies/{id}/usage | get llm policy usage
[**ListLlmPolicies**](LlmPoliciesAPI.md#ListLlmPolicies) | **Get** /llm_policies | list all llm_policies
[**ListLlmPoliciesForAccount**](LlmPoliciesAPI.md#ListLlmPoliciesForAccount) | **Get** /accounts/{account_id}/llm_policies | list llm_policies for account



## CreateOrUpdateLlmPolicy

> LlmPolicy CreateOrUpdateLlmPolicy(ctx, accountId).NoEmbed(noEmbed).Prefer(prefer).CreateOrUpdateLlmPolicyRequest(createOrUpdateLlmPolicyRequest).Execute()

create or update llm_policy



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
	createOrUpdateLlmPolicyRequest := *openapiclient.NewCreateOrUpdateLlmPolicyRequest() // CreateOrUpdateLlmPolicyRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LlmPoliciesAPI.CreateOrUpdateLlmPolicy(context.Background(), accountId).NoEmbed(noEmbed).Prefer(prefer).CreateOrUpdateLlmPolicyRequest(createOrUpdateLlmPolicyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LlmPoliciesAPI.CreateOrUpdateLlmPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrUpdateLlmPolicy`: LlmPolicy
	fmt.Fprintf(os.Stdout, "Response from `LlmPoliciesAPI.CreateOrUpdateLlmPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **int32** | account_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrUpdateLlmPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **noEmbed** | **bool** | When true, omits embedded resources from the response. Also triggered when the Prefer: no_sensitive_extras&#x3D;true header is present. | 
 **prefer** | **string** | When set to no_sensitive_extras&#x3D;true, omits sensitive fields and embedded resources from the response. | 
 **createOrUpdateLlmPolicyRequest** | [**CreateOrUpdateLlmPolicyRequest**](CreateOrUpdateLlmPolicyRequest.md) |  | 

### Return type

[**LlmPolicy**](LlmPolicy.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLlmPolicy

> LlmPolicy GetLlmPolicy(ctx, id).NoEmbed(noEmbed).Prefer(prefer).Execute()

show llm_policy



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
	id := int32(56) // int32 | LLM Policy database ID
	noEmbed := true // bool | When true, omits embedded resources from the response. Also triggered when the Prefer: no_sensitive_extras=true header is present. (optional)
	prefer := "prefer_example" // string | When set to no_sensitive_extras=true, omits sensitive fields and embedded resources from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LlmPoliciesAPI.GetLlmPolicy(context.Background(), id).NoEmbed(noEmbed).Prefer(prefer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LlmPoliciesAPI.GetLlmPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLlmPolicy`: LlmPolicy
	fmt.Fprintf(os.Stdout, "Response from `LlmPoliciesAPI.GetLlmPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | LLM Policy database ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLlmPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **noEmbed** | **bool** | When true, omits embedded resources from the response. Also triggered when the Prefer: no_sensitive_extras&#x3D;true header is present. | 
 **prefer** | **string** | When set to no_sensitive_extras&#x3D;true, omits sensitive fields and embedded resources from the response. | 

### Return type

[**LlmPolicy**](LlmPolicy.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLlmPolicyUsage

> GetLlmKeyUsage200Response GetLlmPolicyUsage(ctx, id).StartDate(startDate).EndDate(endDate).NoEmbed(noEmbed).Prefer(prefer).Execute()

get llm policy usage



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
	id := int32(56) // int32 | LLM Policy database ID
	startDate := "startDate_example" // string | Start date for the usage period (YYYY-MM-DD)
	endDate := "endDate_example" // string | End date for the usage period (YYYY-MM-DD, exclusive)
	noEmbed := true // bool | When true, omits embedded resources from the response. Also triggered when the Prefer: no_sensitive_extras=true header is present. (optional)
	prefer := "prefer_example" // string | When set to no_sensitive_extras=true, omits sensitive fields and embedded resources from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LlmPoliciesAPI.GetLlmPolicyUsage(context.Background(), id).StartDate(startDate).EndDate(endDate).NoEmbed(noEmbed).Prefer(prefer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LlmPoliciesAPI.GetLlmPolicyUsage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLlmPolicyUsage`: GetLlmKeyUsage200Response
	fmt.Fprintf(os.Stdout, "Response from `LlmPoliciesAPI.GetLlmPolicyUsage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | LLM Policy database ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLlmPolicyUsageRequest struct via the builder pattern


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


## ListLlmPolicies

> ListLlmPolicies200Response ListLlmPolicies(ctx).Page(page).PerPage(perPage).NoEmbed(noEmbed).Prefer(prefer).Execute()

list all llm_policies



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
	resp, r, err := apiClient.LlmPoliciesAPI.ListLlmPolicies(context.Background()).Page(page).PerPage(perPage).NoEmbed(noEmbed).Prefer(prefer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LlmPoliciesAPI.ListLlmPolicies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListLlmPolicies`: ListLlmPolicies200Response
	fmt.Fprintf(os.Stdout, "Response from `LlmPoliciesAPI.ListLlmPolicies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListLlmPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Current page of paginated results | 
 **perPage** | **int32** | Number of results to return per page | 
 **noEmbed** | **bool** | When true, omits embedded resources from the response. Also triggered when the Prefer: no_sensitive_extras&#x3D;true header is present. | 
 **prefer** | **string** | When set to no_sensitive_extras&#x3D;true, omits sensitive fields and embedded resources from the response. | 

### Return type

[**ListLlmPolicies200Response**](ListLlmPolicies200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLlmPoliciesForAccount

> ListLlmPoliciesForAccount200Response ListLlmPoliciesForAccount(ctx, accountId).NoEmbed(noEmbed).Prefer(prefer).Execute()

list llm_policies for account



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
	resp, r, err := apiClient.LlmPoliciesAPI.ListLlmPoliciesForAccount(context.Background(), accountId).NoEmbed(noEmbed).Prefer(prefer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LlmPoliciesAPI.ListLlmPoliciesForAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListLlmPoliciesForAccount`: ListLlmPoliciesForAccount200Response
	fmt.Fprintf(os.Stdout, "Response from `LlmPoliciesAPI.ListLlmPoliciesForAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **int32** | account_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiListLlmPoliciesForAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **noEmbed** | **bool** | When true, omits embedded resources from the response. Also triggered when the Prefer: no_sensitive_extras&#x3D;true header is present. | 
 **prefer** | **string** | When set to no_sensitive_extras&#x3D;true, omits sensitive fields and embedded resources from the response. | 

### Return type

[**ListLlmPoliciesForAccount200Response**](ListLlmPoliciesForAccount200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

