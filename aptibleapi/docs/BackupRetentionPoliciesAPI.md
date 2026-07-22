# \BackupRetentionPoliciesAPI

All URIs are relative to *https://api.aptible.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBackupRetentionPolicy**](BackupRetentionPoliciesAPI.md#CreateBackupRetentionPolicy) | **Post** /accounts/{account_id}/backup_retention_policies | create backup retention policy
[**GetBackupRetentionPolicy**](BackupRetentionPoliciesAPI.md#GetBackupRetentionPolicy) | **Get** /backup_retention_policies/{id} | show backup_retention_policy
[**ListBackupRetentionPoliciesForAccount**](BackupRetentionPoliciesAPI.md#ListBackupRetentionPoliciesForAccount) | **Get** /accounts/{account_id}/backup_retention_policies | list backup retention policies



## CreateBackupRetentionPolicy

> CreateBackupRetentionPolicy(ctx, accountId).NoEmbed(noEmbed).Prefer(prefer).CreateBackupRetentionPolicyRequest(createBackupRetentionPolicyRequest).Execute()

create backup retention policy



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
	createBackupRetentionPolicyRequest := *openapiclient.NewCreateBackupRetentionPolicyRequest() // CreateBackupRetentionPolicyRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BackupRetentionPoliciesAPI.CreateBackupRetentionPolicy(context.Background(), accountId).NoEmbed(noEmbed).Prefer(prefer).CreateBackupRetentionPolicyRequest(createBackupRetentionPolicyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupRetentionPoliciesAPI.CreateBackupRetentionPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **int32** | account_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBackupRetentionPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **noEmbed** | **bool** | When true, omits embedded resources from the response. Also triggered when the Prefer: no_sensitive_extras&#x3D;true header is present. | 
 **prefer** | **string** | When set to no_sensitive_extras&#x3D;true, omits sensitive fields and embedded resources from the response. | 
 **createBackupRetentionPolicyRequest** | [**CreateBackupRetentionPolicyRequest**](CreateBackupRetentionPolicyRequest.md) |  | 

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


## GetBackupRetentionPolicy

> BackupRetentionPolicy GetBackupRetentionPolicy(ctx, id).NoEmbed(noEmbed).Prefer(prefer).Execute()

show backup_retention_policy



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
	resp, r, err := apiClient.BackupRetentionPoliciesAPI.GetBackupRetentionPolicy(context.Background(), id).NoEmbed(noEmbed).Prefer(prefer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupRetentionPoliciesAPI.GetBackupRetentionPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBackupRetentionPolicy`: BackupRetentionPolicy
	fmt.Fprintf(os.Stdout, "Response from `BackupRetentionPoliciesAPI.GetBackupRetentionPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBackupRetentionPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **noEmbed** | **bool** | When true, omits embedded resources from the response. Also triggered when the Prefer: no_sensitive_extras&#x3D;true header is present. | 
 **prefer** | **string** | When set to no_sensitive_extras&#x3D;true, omits sensitive fields and embedded resources from the response. | 

### Return type

[**BackupRetentionPolicy**](BackupRetentionPolicy.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBackupRetentionPoliciesForAccount

> ListBackupRetentionPoliciesForAccount200Response ListBackupRetentionPoliciesForAccount(ctx, accountId).Page(page).PerPage(perPage).NoEmbed(noEmbed).Prefer(prefer).Execute()

list backup retention policies



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
	resp, r, err := apiClient.BackupRetentionPoliciesAPI.ListBackupRetentionPoliciesForAccount(context.Background(), accountId).Page(page).PerPage(perPage).NoEmbed(noEmbed).Prefer(prefer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupRetentionPoliciesAPI.ListBackupRetentionPoliciesForAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListBackupRetentionPoliciesForAccount`: ListBackupRetentionPoliciesForAccount200Response
	fmt.Fprintf(os.Stdout, "Response from `BackupRetentionPoliciesAPI.ListBackupRetentionPoliciesForAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **int32** | account_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiListBackupRetentionPoliciesForAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Current page of paginated results | 
 **perPage** | **int32** | Number of results to return per page | 
 **noEmbed** | **bool** | When true, omits embedded resources from the response. Also triggered when the Prefer: no_sensitive_extras&#x3D;true header is present. | 
 **prefer** | **string** | When set to no_sensitive_extras&#x3D;true, omits sensitive fields and embedded resources from the response. | 

### Return type

[**ListBackupRetentionPoliciesForAccount200Response**](ListBackupRetentionPoliciesForAccount200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

