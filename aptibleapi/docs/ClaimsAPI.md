# \ClaimsAPI

All URIs are relative to *https://api.aptible.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClaimsTypePost**](ClaimsAPI.md#ClaimsTypePost) | **Post** /claims/{type} | create claim
[**CreateClaim**](ClaimsAPI.md#CreateClaim) | **Post** /claims | create claim
[**CreateClaimForAccount**](ClaimsAPI.md#CreateClaimForAccount) | **Post** /accounts/{account_id}/claims/{type} | create claim



## ClaimsTypePost

> ClaimsTypePost(ctx, type_).ClaimsTypePostRequest(claimsTypePostRequest).Execute()

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
	claimsTypePostRequest := *openapiclient.NewClaimsTypePostRequest() // ClaimsTypePostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ClaimsAPI.ClaimsTypePost(context.Background(), type_).ClaimsTypePostRequest(claimsTypePostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClaimsAPI.ClaimsTypePost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiClaimsTypePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **claimsTypePostRequest** | [**ClaimsTypePostRequest**](ClaimsTypePostRequest.md) |  | 

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


## CreateClaim

> CreateClaim(ctx).CreateClaimRequest(createClaimRequest).Execute()

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
	createClaimRequest := *openapiclient.NewCreateClaimRequest("Type_example") // CreateClaimRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ClaimsAPI.CreateClaim(context.Background()).CreateClaimRequest(createClaimRequest).Execute()
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

> CreateClaimForAccount(ctx, accountId, type_).CreateAppRequest(createAppRequest).Execute()

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
	createAppRequest := *openapiclient.NewCreateAppRequest("Handle_example") // CreateAppRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ClaimsAPI.CreateClaimForAccount(context.Background(), accountId, type_).CreateAppRequest(createAppRequest).Execute()
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

