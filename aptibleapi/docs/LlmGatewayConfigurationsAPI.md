# \LlmGatewayConfigurationsAPI

All URIs are relative to *https://api.aptible.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateLlmGatewayConfiguration**](LlmGatewayConfigurationsAPI.md#CreateLlmGatewayConfiguration) | **Post** /llm_gateway_configurations | enroll organization in AI Gateway
[**ListLlmGatewayConfigurations**](LlmGatewayConfigurationsAPI.md#ListLlmGatewayConfigurations) | **Get** /llm_gateway_configurations | list llm_gateway_configurations
[**UpdateLlmGatewayConfiguration**](LlmGatewayConfigurationsAPI.md#UpdateLlmGatewayConfiguration) | **Patch** /llm_gateway_configurations/{id} | update llm_gateway_configuration



## CreateLlmGatewayConfiguration

> LlmGatewayConfiguration CreateLlmGatewayConfiguration(ctx).CreateLlmGatewayConfigurationRequest(createLlmGatewayConfigurationRequest).Execute()

enroll organization in AI Gateway

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
	createLlmGatewayConfigurationRequest := *openapiclient.NewCreateLlmGatewayConfigurationRequest("OrganizationId_example") // CreateLlmGatewayConfigurationRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LlmGatewayConfigurationsAPI.CreateLlmGatewayConfiguration(context.Background()).CreateLlmGatewayConfigurationRequest(createLlmGatewayConfigurationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LlmGatewayConfigurationsAPI.CreateLlmGatewayConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateLlmGatewayConfiguration`: LlmGatewayConfiguration
	fmt.Fprintf(os.Stdout, "Response from `LlmGatewayConfigurationsAPI.CreateLlmGatewayConfiguration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateLlmGatewayConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createLlmGatewayConfigurationRequest** | [**CreateLlmGatewayConfigurationRequest**](CreateLlmGatewayConfigurationRequest.md) |  | 

### Return type

[**LlmGatewayConfiguration**](LlmGatewayConfiguration.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLlmGatewayConfigurations

> ListLlmGatewayConfigurations200Response ListLlmGatewayConfigurations(ctx).OrganizationId(organizationId).Execute()

list llm_gateway_configurations

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LlmGatewayConfigurationsAPI.ListLlmGatewayConfigurations(context.Background()).OrganizationId(organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LlmGatewayConfigurationsAPI.ListLlmGatewayConfigurations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListLlmGatewayConfigurations`: ListLlmGatewayConfigurations200Response
	fmt.Fprintf(os.Stdout, "Response from `LlmGatewayConfigurationsAPI.ListLlmGatewayConfigurations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListLlmGatewayConfigurationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organizationId** | **string** | Filter by organization ID | 

### Return type

[**ListLlmGatewayConfigurations200Response**](ListLlmGatewayConfigurations200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLlmGatewayConfiguration

> LlmGatewayConfiguration UpdateLlmGatewayConfiguration(ctx, id).UpdateLlmGatewayConfigurationRequest(updateLlmGatewayConfigurationRequest).Execute()

update llm_gateway_configuration

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
	id := int32(56) // int32 | LLM Gateway Configuration database ID
	updateLlmGatewayConfigurationRequest := *openapiclient.NewUpdateLlmGatewayConfigurationRequest() // UpdateLlmGatewayConfigurationRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LlmGatewayConfigurationsAPI.UpdateLlmGatewayConfiguration(context.Background(), id).UpdateLlmGatewayConfigurationRequest(updateLlmGatewayConfigurationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LlmGatewayConfigurationsAPI.UpdateLlmGatewayConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateLlmGatewayConfiguration`: LlmGatewayConfiguration
	fmt.Fprintf(os.Stdout, "Response from `LlmGatewayConfigurationsAPI.UpdateLlmGatewayConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | LLM Gateway Configuration database ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLlmGatewayConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateLlmGatewayConfigurationRequest** | [**UpdateLlmGatewayConfigurationRequest**](UpdateLlmGatewayConfigurationRequest.md) |  | 

### Return type

[**LlmGatewayConfiguration**](LlmGatewayConfiguration.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

