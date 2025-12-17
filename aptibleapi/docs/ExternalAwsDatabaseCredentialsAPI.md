# \ExternalAwsDatabaseCredentialsAPI

All URIs are relative to *https://api.aptible.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateExternalAwsDatabaseCredential**](ExternalAwsDatabaseCredentialsAPI.md#CreateExternalAwsDatabaseCredential) | **Post** /external_aws_resources/{external_aws_resource_id}/external_aws_database_credentials | create external aws database credential
[**GetExternalAwsDatabaseCredential**](ExternalAwsDatabaseCredentialsAPI.md#GetExternalAwsDatabaseCredential) | **Get** /external_aws_database_credentials/{id} | show external aws database credential



## CreateExternalAwsDatabaseCredential

> ExternalAwsDatabaseCredential CreateExternalAwsDatabaseCredential(ctx, externalAwsResourceId).CreateExternalAwsDatabaseCredentialRequest(createExternalAwsDatabaseCredentialRequest).Execute()

create external aws database credential

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
	externalAwsResourceId := int32(56) // int32 | external_aws_resource_id
	createExternalAwsDatabaseCredentialRequest := *openapiclient.NewCreateExternalAwsDatabaseCredentialRequest("Type_example", "ConnectionUrl_example") // CreateExternalAwsDatabaseCredentialRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExternalAwsDatabaseCredentialsAPI.CreateExternalAwsDatabaseCredential(context.Background(), externalAwsResourceId).CreateExternalAwsDatabaseCredentialRequest(createExternalAwsDatabaseCredentialRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExternalAwsDatabaseCredentialsAPI.CreateExternalAwsDatabaseCredential``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateExternalAwsDatabaseCredential`: ExternalAwsDatabaseCredential
	fmt.Fprintf(os.Stdout, "Response from `ExternalAwsDatabaseCredentialsAPI.CreateExternalAwsDatabaseCredential`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalAwsResourceId** | **int32** | external_aws_resource_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateExternalAwsDatabaseCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createExternalAwsDatabaseCredentialRequest** | [**CreateExternalAwsDatabaseCredentialRequest**](CreateExternalAwsDatabaseCredentialRequest.md) |  | 

### Return type

[**ExternalAwsDatabaseCredential**](ExternalAwsDatabaseCredential.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExternalAwsDatabaseCredential

> ExternalAwsDatabaseCredential GetExternalAwsDatabaseCredential(ctx, id).Execute()

show external aws database credential

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
	resp, r, err := apiClient.ExternalAwsDatabaseCredentialsAPI.GetExternalAwsDatabaseCredential(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExternalAwsDatabaseCredentialsAPI.GetExternalAwsDatabaseCredential``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExternalAwsDatabaseCredential`: ExternalAwsDatabaseCredential
	fmt.Fprintf(os.Stdout, "Response from `ExternalAwsDatabaseCredentialsAPI.GetExternalAwsDatabaseCredential`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExternalAwsDatabaseCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ExternalAwsDatabaseCredential**](ExternalAwsDatabaseCredential.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

