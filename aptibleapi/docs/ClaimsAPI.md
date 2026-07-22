# \ClaimsAPI

All URIs are relative to *https://api.aptible.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateClaim**](ClaimsAPI.md#CreateClaim) | **Post** /claims | create claim
[**CreateClaimForAccount**](ClaimsAPI.md#CreateClaimForAccount) | **Post** /accounts/{account_id}/claims/{type} | create claim
[**CreateClaimForType**](ClaimsAPI.md#CreateClaimForType) | **Post** /claims/{type} | create claim



## CreateClaim

> CreateClaim(ctx).NoEmbed(noEmbed).Prefer(prefer).CreateClaimRequest(createClaimRequest).Execute()

create claim



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
	createClaimRequest := *openapiclient.NewCreateClaimRequest("Type_example") // CreateClaimRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ClaimsAPI.CreateClaim(context.Background()).NoEmbed(noEmbed).Prefer(prefer).CreateClaimRequest(createClaimRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClaimsAPI.CreateClaim``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateClaimRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **noEmbed** | **bool** | When true, omits embedded resources from the response. Also triggered when the Prefer: no_sensitive_extras&#x3D;true header is present. | 
 **prefer** | **string** | When set to no_sensitive_extras&#x3D;true, omits sensitive fields and embedded resources from the response. | 
 **createClaimRequest** | [**CreateClaimRequest**](CreateClaimRequest.md) |  | 

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


## CreateClaimForAccount

> CreateClaimForAccount(ctx, accountId, type_).NoEmbed(noEmbed).Prefer(prefer).CreateAppRequest(createAppRequest).Execute()

create claim



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
	type_ := "type__example" // string | type
	noEmbed := true // bool | When true, omits embedded resources from the response. Also triggered when the Prefer: no_sensitive_extras=true header is present. (optional)
	prefer := "prefer_example" // string | When set to no_sensitive_extras=true, omits sensitive fields and embedded resources from the response. (optional)
	createAppRequest := *openapiclient.NewCreateAppRequest("Handle_example") // CreateAppRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ClaimsAPI.CreateClaimForAccount(context.Background(), accountId, type_).NoEmbed(noEmbed).Prefer(prefer).CreateAppRequest(createAppRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClaimsAPI.CreateClaimForAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **int32** | account_id | 
**type_** | **string** | type | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateClaimForAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **noEmbed** | **bool** | When true, omits embedded resources from the response. Also triggered when the Prefer: no_sensitive_extras&#x3D;true header is present. | 
 **prefer** | **string** | When set to no_sensitive_extras&#x3D;true, omits sensitive fields and embedded resources from the response. | 
 **createAppRequest** | [**CreateAppRequest**](CreateAppRequest.md) |  | 

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


## CreateClaimForType

> CreateClaimForType(ctx, type_).NoEmbed(noEmbed).Prefer(prefer).CreateClaimForTypeRequest(createClaimForTypeRequest).Execute()

create claim



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
	type_ := "type__example" // string | type
	noEmbed := true // bool | When true, omits embedded resources from the response. Also triggered when the Prefer: no_sensitive_extras=true header is present. (optional)
	prefer := "prefer_example" // string | When set to no_sensitive_extras=true, omits sensitive fields and embedded resources from the response. (optional)
	createClaimForTypeRequest := *openapiclient.NewCreateClaimForTypeRequest() // CreateClaimForTypeRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ClaimsAPI.CreateClaimForType(context.Background(), type_).NoEmbed(noEmbed).Prefer(prefer).CreateClaimForTypeRequest(createClaimForTypeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClaimsAPI.CreateClaimForType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | **string** | type | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateClaimForTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **noEmbed** | **bool** | When true, omits embedded resources from the response. Also triggered when the Prefer: no_sensitive_extras&#x3D;true header is present. | 
 **prefer** | **string** | When set to no_sensitive_extras&#x3D;true, omits sensitive fields and embedded resources from the response. | 
 **createClaimForTypeRequest** | [**CreateClaimForTypeRequest**](CreateClaimForTypeRequest.md) |  | 

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

