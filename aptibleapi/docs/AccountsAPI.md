# \AccountsAPI

All URIs are relative to *https://api.aptible.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccount**](AccountsAPI.md#CreateAccount) | **Post** /accounts | create account
[**DeleteAccount**](AccountsAPI.md#DeleteAccount) | **Delete** /accounts/{id} | delete account
[**GetAccount**](AccountsAPI.md#GetAccount) | **Get** /accounts/{id} | show account
[**ListAccounts**](AccountsAPI.md#ListAccounts) | **Get** /accounts | list accounts
[**ListAccountsForStack**](AccountsAPI.md#ListAccountsForStack) | **Get** /stacks/{stack_id}/accounts | list accounts
[**PatchAccount**](AccountsAPI.md#PatchAccount) | **Patch** /accounts/{id} | update account
[**UpdateAccount**](AccountsAPI.md#UpdateAccount) | **Put** /accounts/{id} | update account



## CreateAccount

> Account CreateAccount(ctx).CreateAccountRequest(createAccountRequest).Execute()

create account

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
	createAccountRequest := *openapiclient.NewCreateAccountRequest("Type_example", "Handle_example", "OrganizationId_example") // CreateAccountRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsAPI.CreateAccount(context.Background()).CreateAccountRequest(createAccountRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.CreateAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAccount`: Account
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.CreateAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createAccountRequest** | [**CreateAccountRequest**](CreateAccountRequest.md) |  | 

### Return type

[**Account**](Account.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAccount

> DeleteAccount(ctx, id).Execute()

delete account

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
	r, err := apiClient.AccountsAPI.DeleteAccount(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.DeleteAccount``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteAccountRequest struct via the builder pattern


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


## GetAccount

> Account GetAccount(ctx, id).Execute()

show account

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
	resp, r, err := apiClient.AccountsAPI.GetAccount(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.GetAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccount`: Account
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.GetAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Account**](Account.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAccounts

> ListAccountsForStack200Response ListAccounts(ctx).Page(page).Execute()

list accounts

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
	resp, r, err := apiClient.AccountsAPI.ListAccounts(context.Background()).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.ListAccounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAccounts`: ListAccountsForStack200Response
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.ListAccounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | current page of results for pagination | 

### Return type

[**ListAccountsForStack200Response**](ListAccountsForStack200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAccountsForStack

> ListAccountsForStack200Response ListAccountsForStack(ctx, stackId).Page(page).Execute()

list accounts

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
	stackId := int32(56) // int32 | stack_id
	page := int32(56) // int32 | current page of results for pagination (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsAPI.ListAccountsForStack(context.Background(), stackId).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.ListAccountsForStack``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAccountsForStack`: ListAccountsForStack200Response
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.ListAccountsForStack`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **int32** | stack_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAccountsForStackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | current page of results for pagination | 

### Return type

[**ListAccountsForStack200Response**](ListAccountsForStack200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchAccount

> PatchAccount(ctx, id).PatchAccountRequest(patchAccountRequest).Execute()

update account

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
	patchAccountRequest := *openapiclient.NewPatchAccountRequest() // PatchAccountRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AccountsAPI.PatchAccount(context.Background(), id).PatchAccountRequest(patchAccountRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.PatchAccount``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPatchAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchAccountRequest** | [**PatchAccountRequest**](PatchAccountRequest.md) |  | 

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


## UpdateAccount

> UpdateAccount(ctx, id).UpdateAccountRequest(updateAccountRequest).Execute()

update account

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
	updateAccountRequest := *openapiclient.NewUpdateAccountRequest() // UpdateAccountRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AccountsAPI.UpdateAccount(context.Background(), id).UpdateAccountRequest(updateAccountRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.UpdateAccount``: %v\n", err)
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

Other parameters are passed through a pointer to a apiUpdateAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateAccountRequest** | [**UpdateAccountRequest**](UpdateAccountRequest.md) |  | 

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

