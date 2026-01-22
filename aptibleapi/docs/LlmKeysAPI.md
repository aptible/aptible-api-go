# \LlmKeysAPI

All URIs are relative to *https://api.aptible.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateLlmKey**](LlmKeysAPI.md#CreateLlmKey) | **Post** /accounts/{account_id}/llm_keys | create llm_key
[**GetLlmKey**](LlmKeysAPI.md#GetLlmKey) | **Get** /llm_keys/{id} | show llm_key
[**ListLlmKeys**](LlmKeysAPI.md#ListLlmKeys) | **Get** /llm_keys | list all llm_keys
[**ListLlmKeysForAccount**](LlmKeysAPI.md#ListLlmKeysForAccount) | **Get** /accounts/{account_id}/llm_keys | list llm_keys
[**RevokeLlmKey**](LlmKeysAPI.md#RevokeLlmKey) | **Delete** /llm_keys/{id} | revoke llm_key



## CreateLlmKey

> LlmKey CreateLlmKey(ctx, accountId).CreateLlmKeyRequest(createLlmKeyRequest).Execute()

create llm_key

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
	createLlmKeyRequest := *openapiclient.NewCreateLlmKeyRequest("Handle_example") // CreateLlmKeyRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LlmKeysAPI.CreateLlmKey(context.Background(), accountId).CreateLlmKeyRequest(createLlmKeyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LlmKeysAPI.CreateLlmKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateLlmKey`: LlmKey
	fmt.Fprintf(os.Stdout, "Response from `LlmKeysAPI.CreateLlmKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **int32** | account_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateLlmKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createLlmKeyRequest** | [**CreateLlmKeyRequest**](CreateLlmKeyRequest.md) |  | 

### Return type

[**LlmKey**](LlmKey.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLlmKey

> LlmKey GetLlmKey(ctx, id).Execute()

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LlmKeysAPI.GetLlmKey(context.Background(), id).Execute()
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


## ListLlmKeys

> ListLlmKeys200Response ListLlmKeys(ctx).Page(page).PerPage(perPage).Execute()

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
	page := int32(56) // int32 | current page of results for pagination (optional)
	perPage := int32(56) // int32 | number of results per page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LlmKeysAPI.ListLlmKeys(context.Background()).Page(page).PerPage(perPage).Execute()
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
 **page** | **int32** | current page of results for pagination | 
 **perPage** | **int32** | number of results per page | 

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

> ListLlmKeysForAccount200Response ListLlmKeysForAccount(ctx, accountId).Execute()

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LlmKeysAPI.ListLlmKeysForAccount(context.Background(), accountId).Execute()
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

