# \LlmGatewayConfigurationsAPI

All URIs are relative to *https://api.aptible.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateLlmGatewayConfiguration**](LlmGatewayConfigurationsAPI.md#CreateLlmGatewayConfiguration) | **Post** /llm_gateway_configurations | enroll organization in AI Gateway
[**ListLlmGatewayConfigurations**](LlmGatewayConfigurationsAPI.md#ListLlmGatewayConfigurations) | **Get** /llm_gateway_configurations | list llm_gateway_configurations
[**UpdateLlmGatewayConfiguration**](LlmGatewayConfigurationsAPI.md#UpdateLlmGatewayConfiguration) | **Patch** /llm_gateway_configurations/{id} | update llm_gateway_configuration



## CreateLlmGatewayConfiguration

> LlmGatewayConfiguration CreateLlmGatewayConfiguration(ctx).NoEmbed(noEmbed).Prefer(prefer).CreateLlmGatewayConfigurationRequest(createLlmGatewayConfigurationRequest).Execute()

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
	noEmbed := true // bool | When true, omits embedded resources from the response. Also triggered when the Prefer: no_sensitive_extras=true header is present. (optional)
	prefer := "prefer_example" // string | When set to no_sensitive_extras=true, omits sensitive fields and embedded resources from the response. (optional)
	createLlmGatewayConfigurationRequest := *openapiclient.NewCreateLlmGatewayConfigurationRequest("OrganizationId_example") // CreateLlmGatewayConfigurationRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LlmGatewayConfigurationsAPI.CreateLlmGatewayConfiguration(context.Background()).NoEmbed(noEmbed).Prefer(prefer).CreateLlmGatewayConfigurationRequest(createLlmGatewayConfigurationRequest).Execute()
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
 **noEmbed** | **bool** | When true, omits embedded resources from the response. Also triggered when the Prefer: no_sensitive_extras&#x3D;true header is present. | 
 **prefer** | **string** | When set to no_sensitive_extras&#x3D;true, omits sensitive fields and embedded resources from the response. | 
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

> ListLlmGatewayConfigurations200Response ListLlmGatewayConfigurations(ctx).OrganizationId(organizationId).NoEmbed(noEmbed).Prefer(prefer).Execute()

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
	noEmbed := true // bool | When true, omits embedded resources from the response. Also triggered when the Prefer: no_sensitive_extras=true header is present. (optional)
	prefer := "prefer_example" // string | When set to no_sensitive_extras=true, omits sensitive fields and embedded resources from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LlmGatewayConfigurationsAPI.ListLlmGatewayConfigurations(context.Background()).OrganizationId(organizationId).NoEmbed(noEmbed).Prefer(prefer).Execute()
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
 **noEmbed** | **bool** | When true, omits embedded resources from the response. Also triggered when the Prefer: no_sensitive_extras&#x3D;true header is present. | 
 **prefer** | **string** | When set to no_sensitive_extras&#x3D;true, omits sensitive fields and embedded resources from the response. | 

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

> LlmGatewayConfiguration UpdateLlmGatewayConfiguration(ctx, id).NoEmbed(noEmbed).Prefer(prefer).UpdateLlmGatewayConfigurationRequest(updateLlmGatewayConfigurationRequest).Execute()

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
	noEmbed := true // bool | When true, omits embedded resources from the response. Also triggered when the Prefer: no_sensitive_extras=true header is present. (optional)
	prefer := "prefer_example" // string | When set to no_sensitive_extras=true, omits sensitive fields and embedded resources from the response. (optional)
	updateLlmGatewayConfigurationRequest := *openapiclient.NewUpdateLlmGatewayConfigurationRequest() // UpdateLlmGatewayConfigurationRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LlmGatewayConfigurationsAPI.UpdateLlmGatewayConfiguration(context.Background(), id).NoEmbed(noEmbed).Prefer(prefer).UpdateLlmGatewayConfigurationRequest(updateLlmGatewayConfigurationRequest).Execute()
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

 **noEmbed** | **bool** | When true, omits embedded resources from the response. Also triggered when the Prefer: no_sensitive_extras&#x3D;true header is present. | 
 **prefer** | **string** | When set to no_sensitive_extras&#x3D;true, omits sensitive fields and embedded resources from the response. | 
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

