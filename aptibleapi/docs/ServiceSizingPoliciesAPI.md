# \ServiceSizingPoliciesAPI

All URIs are relative to *https://api.aptible.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateServiceSizingPolicy**](ServiceSizingPoliciesAPI.md#CreateServiceSizingPolicy) | **Post** /services/{service_id}/service_sizing_policies | create service_sizing_policy
[**DeleteServiceSizingPolicy**](ServiceSizingPoliciesAPI.md#DeleteServiceSizingPolicy) | **Delete** /services/{service_id}/service_sizing_policy | delete service_sizing_policy
[**GetServiceSizingPolicy**](ServiceSizingPoliciesAPI.md#GetServiceSizingPolicy) | **Get** /service_sizing_policies/{id} | show service_sizing_policy
[**ListServiceSizingPoliciesForAccount**](ServiceSizingPoliciesAPI.md#ListServiceSizingPoliciesForAccount) | **Get** /accounts/{account_id}/service_sizing_policies | list service_sizing_policies
[**ListServiceSizingPoliciesForService**](ServiceSizingPoliciesAPI.md#ListServiceSizingPoliciesForService) | **Get** /services/{service_id}/service_sizing_policies | list service_sizing_policies
[**UpdateServiceSizingPolicy**](ServiceSizingPoliciesAPI.md#UpdateServiceSizingPolicy) | **Put** /services/{service_id}/service_sizing_policies | update service_sizing_policy



## CreateServiceSizingPolicy

> CreateServiceSizingPolicy(ctx, serviceId).NoEmbed(noEmbed).Prefer(prefer).UpdateServiceSizingPolicyRequest(updateServiceSizingPolicyRequest).Execute()

create service_sizing_policy



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
	serviceId := int32(56) // int32 | service_id
	noEmbed := true // bool | When true, omits embedded resources from the response. Also triggered when the Prefer: no_sensitive_extras=true header is present. (optional)
	prefer := "prefer_example" // string | When set to no_sensitive_extras=true, omits sensitive fields and embedded resources from the response. (optional)
	updateServiceSizingPolicyRequest := *openapiclient.NewUpdateServiceSizingPolicyRequest() // UpdateServiceSizingPolicyRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ServiceSizingPoliciesAPI.CreateServiceSizingPolicy(context.Background(), serviceId).NoEmbed(noEmbed).Prefer(prefer).UpdateServiceSizingPolicyRequest(updateServiceSizingPolicyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceSizingPoliciesAPI.CreateServiceSizingPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **int32** | service_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateServiceSizingPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **noEmbed** | **bool** | When true, omits embedded resources from the response. Also triggered when the Prefer: no_sensitive_extras&#x3D;true header is present. | 
 **prefer** | **string** | When set to no_sensitive_extras&#x3D;true, omits sensitive fields and embedded resources from the response. | 
 **updateServiceSizingPolicyRequest** | [**UpdateServiceSizingPolicyRequest**](UpdateServiceSizingPolicyRequest.md) |  | 

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


## DeleteServiceSizingPolicy

> DeleteServiceSizingPolicy(ctx, serviceId).Execute()

delete service_sizing_policy



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
	serviceId := int32(56) // int32 | service_id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ServiceSizingPoliciesAPI.DeleteServiceSizingPolicy(context.Background(), serviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceSizingPoliciesAPI.DeleteServiceSizingPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **int32** | service_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteServiceSizingPolicyRequest struct via the builder pattern


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


## GetServiceSizingPolicy

> ServiceSizingPolicy GetServiceSizingPolicy(ctx, id).NoEmbed(noEmbed).Prefer(prefer).Execute()

show service_sizing_policy



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
	resp, r, err := apiClient.ServiceSizingPoliciesAPI.GetServiceSizingPolicy(context.Background(), id).NoEmbed(noEmbed).Prefer(prefer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceSizingPoliciesAPI.GetServiceSizingPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServiceSizingPolicy`: ServiceSizingPolicy
	fmt.Fprintf(os.Stdout, "Response from `ServiceSizingPoliciesAPI.GetServiceSizingPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceSizingPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **noEmbed** | **bool** | When true, omits embedded resources from the response. Also triggered when the Prefer: no_sensitive_extras&#x3D;true header is present. | 
 **prefer** | **string** | When set to no_sensitive_extras&#x3D;true, omits sensitive fields and embedded resources from the response. | 

### Return type

[**ServiceSizingPolicy**](ServiceSizingPolicy.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListServiceSizingPoliciesForAccount

> ListServiceSizingPoliciesForAccount200Response ListServiceSizingPoliciesForAccount(ctx, accountId).Page(page).PerPage(perPage).NoEmbed(noEmbed).Prefer(prefer).Execute()

list service_sizing_policies



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
	resp, r, err := apiClient.ServiceSizingPoliciesAPI.ListServiceSizingPoliciesForAccount(context.Background(), accountId).Page(page).PerPage(perPage).NoEmbed(noEmbed).Prefer(prefer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceSizingPoliciesAPI.ListServiceSizingPoliciesForAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListServiceSizingPoliciesForAccount`: ListServiceSizingPoliciesForAccount200Response
	fmt.Fprintf(os.Stdout, "Response from `ServiceSizingPoliciesAPI.ListServiceSizingPoliciesForAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **int32** | account_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiListServiceSizingPoliciesForAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Current page of paginated results | 
 **perPage** | **int32** | Number of results to return per page | 
 **noEmbed** | **bool** | When true, omits embedded resources from the response. Also triggered when the Prefer: no_sensitive_extras&#x3D;true header is present. | 
 **prefer** | **string** | When set to no_sensitive_extras&#x3D;true, omits sensitive fields and embedded resources from the response. | 

### Return type

[**ListServiceSizingPoliciesForAccount200Response**](ListServiceSizingPoliciesForAccount200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListServiceSizingPoliciesForService

> ListServiceSizingPoliciesForAccount200Response ListServiceSizingPoliciesForService(ctx, serviceId).Page(page).PerPage(perPage).NoEmbed(noEmbed).Prefer(prefer).Execute()

list service_sizing_policies



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
	serviceId := int32(56) // int32 | service_id
	page := int32(56) // int32 | Current page of paginated results (optional)
	perPage := int32(56) // int32 | Number of results to return per page (optional)
	noEmbed := true // bool | When true, omits embedded resources from the response. Also triggered when the Prefer: no_sensitive_extras=true header is present. (optional)
	prefer := "prefer_example" // string | When set to no_sensitive_extras=true, omits sensitive fields and embedded resources from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceSizingPoliciesAPI.ListServiceSizingPoliciesForService(context.Background(), serviceId).Page(page).PerPage(perPage).NoEmbed(noEmbed).Prefer(prefer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceSizingPoliciesAPI.ListServiceSizingPoliciesForService``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListServiceSizingPoliciesForService`: ListServiceSizingPoliciesForAccount200Response
	fmt.Fprintf(os.Stdout, "Response from `ServiceSizingPoliciesAPI.ListServiceSizingPoliciesForService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **int32** | service_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiListServiceSizingPoliciesForServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Current page of paginated results | 
 **perPage** | **int32** | Number of results to return per page | 
 **noEmbed** | **bool** | When true, omits embedded resources from the response. Also triggered when the Prefer: no_sensitive_extras&#x3D;true header is present. | 
 **prefer** | **string** | When set to no_sensitive_extras&#x3D;true, omits sensitive fields and embedded resources from the response. | 

### Return type

[**ListServiceSizingPoliciesForAccount200Response**](ListServiceSizingPoliciesForAccount200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateServiceSizingPolicy

> UpdateServiceSizingPolicy(ctx, serviceId).NoEmbed(noEmbed).Prefer(prefer).UpdateServiceSizingPolicyRequest(updateServiceSizingPolicyRequest).Execute()

update service_sizing_policy



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
	serviceId := int32(56) // int32 | service_id
	noEmbed := true // bool | When true, omits embedded resources from the response. Also triggered when the Prefer: no_sensitive_extras=true header is present. (optional)
	prefer := "prefer_example" // string | When set to no_sensitive_extras=true, omits sensitive fields and embedded resources from the response. (optional)
	updateServiceSizingPolicyRequest := *openapiclient.NewUpdateServiceSizingPolicyRequest() // UpdateServiceSizingPolicyRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ServiceSizingPoliciesAPI.UpdateServiceSizingPolicy(context.Background(), serviceId).NoEmbed(noEmbed).Prefer(prefer).UpdateServiceSizingPolicyRequest(updateServiceSizingPolicyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceSizingPoliciesAPI.UpdateServiceSizingPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **int32** | service_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateServiceSizingPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **noEmbed** | **bool** | When true, omits embedded resources from the response. Also triggered when the Prefer: no_sensitive_extras&#x3D;true header is present. | 
 **prefer** | **string** | When set to no_sensitive_extras&#x3D;true, omits sensitive fields and embedded resources from the response. | 
 **updateServiceSizingPolicyRequest** | [**UpdateServiceSizingPolicyRequest**](UpdateServiceSizingPolicyRequest.md) |  | 

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

