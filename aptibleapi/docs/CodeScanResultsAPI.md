# \CodeScanResultsAPI

All URIs are relative to *https://api.aptible.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCodeScanResult**](CodeScanResultsAPI.md#CreateCodeScanResult) | **Post** /apps/{app_id}/code_scan_results | create code_scan_result
[**DeleteCodeScanResult**](CodeScanResultsAPI.md#DeleteCodeScanResult) | **Delete** /code_scan_results/{id} | delete code_scan_result
[**GetCodeScanResult**](CodeScanResultsAPI.md#GetCodeScanResult) | **Get** /code_scan_results/{id} | show code_scan_result
[**ListCodeScanResultsForApp**](CodeScanResultsAPI.md#ListCodeScanResultsForApp) | **Get** /apps/{app_id}/code_scan_results | list code_scan_results



## CreateCodeScanResult

> CodeScanResult CreateCodeScanResult(ctx, appId).CreateCodeScanResultRequest(createCodeScanResultRequest).Execute()

create code_scan_result

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
	appId := int32(56) // int32 | app_id
	createCodeScanResultRequest := *openapiclient.NewCreateCodeScanResultRequest(int32(123), "GitRef_example", "GitCommit_example") // CreateCodeScanResultRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CodeScanResultsAPI.CreateCodeScanResult(context.Background(), appId).CreateCodeScanResultRequest(createCodeScanResultRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CodeScanResultsAPI.CreateCodeScanResult``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCodeScanResult`: CodeScanResult
	fmt.Fprintf(os.Stdout, "Response from `CodeScanResultsAPI.CreateCodeScanResult`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** | app_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCodeScanResultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createCodeScanResultRequest** | [**CreateCodeScanResultRequest**](CreateCodeScanResultRequest.md) |  | 

### Return type

[**CodeScanResult**](CodeScanResult.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCodeScanResult

> DeleteCodeScanResult(ctx, id).Execute()

delete code_scan_result

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
	r, err := apiClient.CodeScanResultsAPI.DeleteCodeScanResult(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CodeScanResultsAPI.DeleteCodeScanResult``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteCodeScanResultRequest struct via the builder pattern


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


## GetCodeScanResult

> CodeScanResult GetCodeScanResult(ctx, id).Execute()

show code_scan_result

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
	resp, r, err := apiClient.CodeScanResultsAPI.GetCodeScanResult(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CodeScanResultsAPI.GetCodeScanResult``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCodeScanResult`: CodeScanResult
	fmt.Fprintf(os.Stdout, "Response from `CodeScanResultsAPI.GetCodeScanResult`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCodeScanResultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CodeScanResult**](CodeScanResult.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCodeScanResultsForApp

> ListCodeScanResultsForApp200Response ListCodeScanResultsForApp(ctx, appId).Page(page).Execute()

list code_scan_results

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
	appId := int32(56) // int32 | app_id
	page := int32(56) // int32 | current page of results for pagination (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CodeScanResultsAPI.ListCodeScanResultsForApp(context.Background(), appId).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CodeScanResultsAPI.ListCodeScanResultsForApp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCodeScanResultsForApp`: ListCodeScanResultsForApp200Response
	fmt.Fprintf(os.Stdout, "Response from `CodeScanResultsAPI.ListCodeScanResultsForApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** | app_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCodeScanResultsForAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | current page of results for pagination | 

### Return type

[**ListCodeScanResultsForApp200Response**](ListCodeScanResultsForApp200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

