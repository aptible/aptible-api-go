# \LLMIntegrationsAPI

All URIs are relative to *https://api.aptible.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateLlmIntegration**](LLMIntegrationsAPI.md#CreateLlmIntegration) | **Post** /llm_integrations | create llm integration
[**DeleteLlmIntegration**](LLMIntegrationsAPI.md#DeleteLlmIntegration) | **Delete** /llm_integrations/{id} | delete llm integration
[**GetLlmIntegration**](LLMIntegrationsAPI.md#GetLlmIntegration) | **Get** /llm_integrations/{id} | show llm integration
[**ListLlmIntegrations**](LLMIntegrationsAPI.md#ListLlmIntegrations) | **Get** /llm_integrations | list llm integrations
[**UpdateLlmIntegration**](LLMIntegrationsAPI.md#UpdateLlmIntegration) | **Put** /llm_integrations/{id} | update llm integration



## CreateLlmIntegration

> LlmIntegration CreateLlmIntegration(ctx).CreateLlmIntegrationRequest(createLlmIntegrationRequest).Execute()

create llm integration

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
	createLlmIntegrationRequest := *openapiclient.NewCreateLlmIntegrationRequest("ProviderType_example", "OrganizationId_example") // CreateLlmIntegrationRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LLMIntegrationsAPI.CreateLlmIntegration(context.Background()).CreateLlmIntegrationRequest(createLlmIntegrationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LLMIntegrationsAPI.CreateLlmIntegration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateLlmIntegration`: LlmIntegration
	fmt.Fprintf(os.Stdout, "Response from `LLMIntegrationsAPI.CreateLlmIntegration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateLlmIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createLlmIntegrationRequest** | [**CreateLlmIntegrationRequest**](CreateLlmIntegrationRequest.md) |  | 

### Return type

[**LlmIntegration**](LlmIntegration.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteLlmIntegration

> DeleteLlmIntegration(ctx, id).Execute()

delete llm integration

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
	r, err := apiClient.LLMIntegrationsAPI.DeleteLlmIntegration(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LLMIntegrationsAPI.DeleteLlmIntegration``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteLlmIntegrationRequest struct via the builder pattern


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


## GetLlmIntegration

> LlmIntegration GetLlmIntegration(ctx, id).Execute()

show llm integration

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
	resp, r, err := apiClient.LLMIntegrationsAPI.GetLlmIntegration(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LLMIntegrationsAPI.GetLlmIntegration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLlmIntegration`: LlmIntegration
	fmt.Fprintf(os.Stdout, "Response from `LLMIntegrationsAPI.GetLlmIntegration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLlmIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LlmIntegration**](LlmIntegration.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLlmIntegrations

> ListLlmIntegrations200Response ListLlmIntegrations(ctx).Page(page).Execute()

list llm integrations

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LLMIntegrationsAPI.ListLlmIntegrations(context.Background()).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LLMIntegrationsAPI.ListLlmIntegrations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListLlmIntegrations`: ListLlmIntegrations200Response
	fmt.Fprintf(os.Stdout, "Response from `LLMIntegrationsAPI.ListLlmIntegrations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListLlmIntegrationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | current page of results for pagination | 

### Return type

[**ListLlmIntegrations200Response**](ListLlmIntegrations200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLlmIntegration

> UpdateLlmIntegration(ctx, id).UpdateLlmIntegrationRequest(updateLlmIntegrationRequest).Execute()

update llm integration

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
	updateLlmIntegrationRequest := *openapiclient.NewUpdateLlmIntegrationRequest() // UpdateLlmIntegrationRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LLMIntegrationsAPI.UpdateLlmIntegration(context.Background(), id).UpdateLlmIntegrationRequest(updateLlmIntegrationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LLMIntegrationsAPI.UpdateLlmIntegration``: %v\n", err)
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

Other parameters are passed through a pointer to a apiUpdateLlmIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateLlmIntegrationRequest** | [**UpdateLlmIntegrationRequest**](UpdateLlmIntegrationRequest.md) |  | 

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

