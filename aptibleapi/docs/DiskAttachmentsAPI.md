# \DiskAttachmentsAPI

All URIs are relative to *https://api.aptible.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDiskAttachment**](DiskAttachmentsAPI.md#GetDiskAttachment) | **Get** /disk_attachments/{id} | show disk attachment
[**ListDiskAttachments**](DiskAttachmentsAPI.md#ListDiskAttachments) | **Get** /disk_attachments | list disk attachments
[**ListDiskAttachmentsForAccount**](DiskAttachmentsAPI.md#ListDiskAttachmentsForAccount) | **Get** /accounts/{account_id}/disk_attachments | list disk attachments
[**ListDiskAttachmentsForPersistentDisk**](DiskAttachmentsAPI.md#ListDiskAttachmentsForPersistentDisk) | **Get** /persistent_disks/{persistent_disk_id}/disk_attachments | list disk attachments
[**ListDiskAttachmentsForService**](DiskAttachmentsAPI.md#ListDiskAttachmentsForService) | **Get** /services/{service_id}/disk_attachments | list disk attachments



## GetDiskAttachment

> DiskAttachment GetDiskAttachment(ctx, id).WithDeleted(withDeleted).Execute()

show disk attachment

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
	withDeleted := true // bool | include soft-deleted records (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DiskAttachmentsAPI.GetDiskAttachment(context.Background(), id).WithDeleted(withDeleted).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DiskAttachmentsAPI.GetDiskAttachment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDiskAttachment`: DiskAttachment
	fmt.Fprintf(os.Stdout, "Response from `DiskAttachmentsAPI.GetDiskAttachment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDiskAttachmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **withDeleted** | **bool** | include soft-deleted records | 

### Return type

[**DiskAttachment**](DiskAttachment.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDiskAttachments

> ListDiskAttachments200Response ListDiskAttachments(ctx).Page(page).Execute()

list disk attachments

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
	resp, r, err := apiClient.DiskAttachmentsAPI.ListDiskAttachments(context.Background()).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DiskAttachmentsAPI.ListDiskAttachments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDiskAttachments`: ListDiskAttachments200Response
	fmt.Fprintf(os.Stdout, "Response from `DiskAttachmentsAPI.ListDiskAttachments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDiskAttachmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | current page of results for pagination | 

### Return type

[**ListDiskAttachments200Response**](ListDiskAttachments200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDiskAttachmentsForAccount

> ListDiskAttachmentsForAccount200Response ListDiskAttachmentsForAccount(ctx, accountId).Page(page).Execute()

list disk attachments

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
	page := int32(56) // int32 | current page of results for pagination (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DiskAttachmentsAPI.ListDiskAttachmentsForAccount(context.Background(), accountId).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DiskAttachmentsAPI.ListDiskAttachmentsForAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDiskAttachmentsForAccount`: ListDiskAttachmentsForAccount200Response
	fmt.Fprintf(os.Stdout, "Response from `DiskAttachmentsAPI.ListDiskAttachmentsForAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **int32** | account_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiListDiskAttachmentsForAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | current page of results for pagination | 

### Return type

[**ListDiskAttachmentsForAccount200Response**](ListDiskAttachmentsForAccount200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDiskAttachmentsForPersistentDisk

> ListDiskAttachmentsForPersistentDisk200Response ListDiskAttachmentsForPersistentDisk(ctx, persistentDiskId).Page(page).Execute()

list disk attachments

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
	persistentDiskId := int32(56) // int32 | persistent_disk_id
	page := int32(56) // int32 | current page of results for pagination (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DiskAttachmentsAPI.ListDiskAttachmentsForPersistentDisk(context.Background(), persistentDiskId).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DiskAttachmentsAPI.ListDiskAttachmentsForPersistentDisk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDiskAttachmentsForPersistentDisk`: ListDiskAttachmentsForPersistentDisk200Response
	fmt.Fprintf(os.Stdout, "Response from `DiskAttachmentsAPI.ListDiskAttachmentsForPersistentDisk`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**persistentDiskId** | **int32** | persistent_disk_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiListDiskAttachmentsForPersistentDiskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | current page of results for pagination | 

### Return type

[**ListDiskAttachmentsForPersistentDisk200Response**](ListDiskAttachmentsForPersistentDisk200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDiskAttachmentsForService

> ListDiskAttachmentsForService200Response ListDiskAttachmentsForService(ctx, serviceId).Page(page).Execute()

list disk attachments

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
	serviceId := int32(56) // int32 | service_id
	page := int32(56) // int32 | current page of results for pagination (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DiskAttachmentsAPI.ListDiskAttachmentsForService(context.Background(), serviceId).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DiskAttachmentsAPI.ListDiskAttachmentsForService``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDiskAttachmentsForService`: ListDiskAttachmentsForService200Response
	fmt.Fprintf(os.Stdout, "Response from `DiskAttachmentsAPI.ListDiskAttachmentsForService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **int32** | service_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiListDiskAttachmentsForServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | current page of results for pagination | 

### Return type

[**ListDiskAttachmentsForService200Response**](ListDiskAttachmentsForService200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

