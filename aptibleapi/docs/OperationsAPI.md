# \OperationsAPI

All URIs are relative to *https://api.aptible.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOperationForApp**](OperationsAPI.md#CreateOperationForApp) | **Post** /apps/{app_id}/operations | create operation
[**CreateOperationForBackup**](OperationsAPI.md#CreateOperationForBackup) | **Post** /backups/{backup_id}/operations | create operation
[**CreateOperationForDatabase**](OperationsAPI.md#CreateOperationForDatabase) | **Post** /databases/{database_id}/operations | create operation
[**CreateOperationForDatabaseCredential**](OperationsAPI.md#CreateOperationForDatabaseCredential) | **Post** /database_credentials/{database_credential_id}/operations | create operation
[**CreateOperationForDiskAttachment**](OperationsAPI.md#CreateOperationForDiskAttachment) | **Post** /disk_attachments/{disk_attachment_id}/operations | create operation
[**CreateOperationForEphemeralSession**](OperationsAPI.md#CreateOperationForEphemeralSession) | **Post** /ephemeral_sessions/{ephemeral_session_id}/operations | create operation
[**CreateOperationForImage**](OperationsAPI.md#CreateOperationForImage) | **Post** /images/{image_id}/operations | create operation
[**CreateOperationForLogDrain**](OperationsAPI.md#CreateOperationForLogDrain) | **Post** /log_drains/{log_drain_id}/operations | create operation
[**CreateOperationForMetricDrain**](OperationsAPI.md#CreateOperationForMetricDrain) | **Post** /metric_drains/{metric_drain_id}/operations | create operation
[**CreateOperationForPersistentDisk**](OperationsAPI.md#CreateOperationForPersistentDisk) | **Post** /persistent_disks/{persistent_disk_id}/operations | create operation
[**CreateOperationForService**](OperationsAPI.md#CreateOperationForService) | **Post** /services/{service_id}/operations | create operation
[**CreateOperationForVhost**](OperationsAPI.md#CreateOperationForVhost) | **Post** /vhosts/{vhost_id}/operations | create operation
[**GetOperation**](OperationsAPI.md#GetOperation) | **Get** /operations/{id} | show operation
[**ListOperationsForAccount**](OperationsAPI.md#ListOperationsForAccount) | **Get** /accounts/{account_id}/operations | list operations
[**ListOperationsForApp**](OperationsAPI.md#ListOperationsForApp) | **Get** /apps/{app_id}/operations | list operations
[**ListOperationsForBackup**](OperationsAPI.md#ListOperationsForBackup) | **Get** /backups/{backup_id}/operations | list operations
[**ListOperationsForDatabase**](OperationsAPI.md#ListOperationsForDatabase) | **Get** /databases/{database_id}/operations | list operations
[**ListOperationsForDatabaseCredential**](OperationsAPI.md#ListOperationsForDatabaseCredential) | **Get** /database_credentials/{database_credential_id}/operations | list operations
[**ListOperationsForDiskAttachment**](OperationsAPI.md#ListOperationsForDiskAttachment) | **Get** /disk_attachments/{disk_attachment_id}/operations | list operations
[**ListOperationsForEphemeralSession**](OperationsAPI.md#ListOperationsForEphemeralSession) | **Get** /ephemeral_sessions/{ephemeral_session_id}/operations | list operations
[**ListOperationsForImage**](OperationsAPI.md#ListOperationsForImage) | **Get** /images/{image_id}/operations | list operations
[**ListOperationsForLogDrain**](OperationsAPI.md#ListOperationsForLogDrain) | **Get** /log_drains/{log_drain_id}/operations | list operations
[**ListOperationsForMetricDrain**](OperationsAPI.md#ListOperationsForMetricDrain) | **Get** /metric_drains/{metric_drain_id}/operations | list operations
[**ListOperationsForPersistentDisk**](OperationsAPI.md#ListOperationsForPersistentDisk) | **Get** /persistent_disks/{persistent_disk_id}/operations | list operations
[**ListOperationsForService**](OperationsAPI.md#ListOperationsForService) | **Get** /services/{service_id}/operations | list operations
[**ListOperationsForVhost**](OperationsAPI.md#ListOperationsForVhost) | **Get** /vhosts/{vhost_id}/operations | list operations
[**PatchOperation**](OperationsAPI.md#PatchOperation) | **Patch** /operations/{id} | update operation
[**UpdateOperation**](OperationsAPI.md#UpdateOperation) | **Put** /operations/{id} | update operation



## CreateOperationForApp

> Operation CreateOperationForApp(ctx, appId).CreateOperationRequest(createOperationRequest).Execute()

create operation

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
	createOperationRequest := *openapiclient.NewCreateOperationRequest("Type_example") // CreateOperationRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OperationsAPI.CreateOperationForApp(context.Background(), appId).CreateOperationRequest(createOperationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OperationsAPI.CreateOperationForApp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOperationForApp`: Operation
	fmt.Fprintf(os.Stdout, "Response from `OperationsAPI.CreateOperationForApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** | app_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOperationForAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOperationRequest** | [**CreateOperationRequest**](CreateOperationRequest.md) |  | 

### Return type

[**Operation**](Operation.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOperationForBackup

> Operation CreateOperationForBackup(ctx, backupId).CreateOperationRequest(createOperationRequest).Execute()

create operation

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
	backupId := int32(56) // int32 | backup_id
	createOperationRequest := *openapiclient.NewCreateOperationRequest("Type_example") // CreateOperationRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OperationsAPI.CreateOperationForBackup(context.Background(), backupId).CreateOperationRequest(createOperationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OperationsAPI.CreateOperationForBackup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOperationForBackup`: Operation
	fmt.Fprintf(os.Stdout, "Response from `OperationsAPI.CreateOperationForBackup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**backupId** | **int32** | backup_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOperationForBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOperationRequest** | [**CreateOperationRequest**](CreateOperationRequest.md) |  | 

### Return type

[**Operation**](Operation.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOperationForDatabase

> Operation CreateOperationForDatabase(ctx, databaseId).CreateOperationRequest(createOperationRequest).Execute()

create operation

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
	databaseId := int32(56) // int32 | database_id
	createOperationRequest := *openapiclient.NewCreateOperationRequest("Type_example") // CreateOperationRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OperationsAPI.CreateOperationForDatabase(context.Background(), databaseId).CreateOperationRequest(createOperationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OperationsAPI.CreateOperationForDatabase``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOperationForDatabase`: Operation
	fmt.Fprintf(os.Stdout, "Response from `OperationsAPI.CreateOperationForDatabase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**databaseId** | **int32** | database_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOperationForDatabaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOperationRequest** | [**CreateOperationRequest**](CreateOperationRequest.md) |  | 

### Return type

[**Operation**](Operation.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOperationForDatabaseCredential

> Operation CreateOperationForDatabaseCredential(ctx, databaseCredentialId).CreateOperationRequest(createOperationRequest).Execute()

create operation

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
	databaseCredentialId := int32(56) // int32 | database_credential_id
	createOperationRequest := *openapiclient.NewCreateOperationRequest("Type_example") // CreateOperationRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OperationsAPI.CreateOperationForDatabaseCredential(context.Background(), databaseCredentialId).CreateOperationRequest(createOperationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OperationsAPI.CreateOperationForDatabaseCredential``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOperationForDatabaseCredential`: Operation
	fmt.Fprintf(os.Stdout, "Response from `OperationsAPI.CreateOperationForDatabaseCredential`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**databaseCredentialId** | **int32** | database_credential_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOperationForDatabaseCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOperationRequest** | [**CreateOperationRequest**](CreateOperationRequest.md) |  | 

### Return type

[**Operation**](Operation.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOperationForDiskAttachment

> Operation CreateOperationForDiskAttachment(ctx, diskAttachmentId).CreateOperationRequest(createOperationRequest).Execute()

create operation

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
	diskAttachmentId := int32(56) // int32 | disk_attachment_id
	createOperationRequest := *openapiclient.NewCreateOperationRequest("Type_example") // CreateOperationRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OperationsAPI.CreateOperationForDiskAttachment(context.Background(), diskAttachmentId).CreateOperationRequest(createOperationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OperationsAPI.CreateOperationForDiskAttachment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOperationForDiskAttachment`: Operation
	fmt.Fprintf(os.Stdout, "Response from `OperationsAPI.CreateOperationForDiskAttachment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**diskAttachmentId** | **int32** | disk_attachment_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOperationForDiskAttachmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOperationRequest** | [**CreateOperationRequest**](CreateOperationRequest.md) |  | 

### Return type

[**Operation**](Operation.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOperationForEphemeralSession

> Operation CreateOperationForEphemeralSession(ctx, ephemeralSessionId).CreateOperationRequest(createOperationRequest).Execute()

create operation

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
	ephemeralSessionId := int32(56) // int32 | ephemeral_session_id
	createOperationRequest := *openapiclient.NewCreateOperationRequest("Type_example") // CreateOperationRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OperationsAPI.CreateOperationForEphemeralSession(context.Background(), ephemeralSessionId).CreateOperationRequest(createOperationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OperationsAPI.CreateOperationForEphemeralSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOperationForEphemeralSession`: Operation
	fmt.Fprintf(os.Stdout, "Response from `OperationsAPI.CreateOperationForEphemeralSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ephemeralSessionId** | **int32** | ephemeral_session_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOperationForEphemeralSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOperationRequest** | [**CreateOperationRequest**](CreateOperationRequest.md) |  | 

### Return type

[**Operation**](Operation.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOperationForImage

> Operation CreateOperationForImage(ctx, imageId).CreateOperationRequest(createOperationRequest).Execute()

create operation

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
	imageId := int32(56) // int32 | image_id
	createOperationRequest := *openapiclient.NewCreateOperationRequest("Type_example") // CreateOperationRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OperationsAPI.CreateOperationForImage(context.Background(), imageId).CreateOperationRequest(createOperationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OperationsAPI.CreateOperationForImage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOperationForImage`: Operation
	fmt.Fprintf(os.Stdout, "Response from `OperationsAPI.CreateOperationForImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageId** | **int32** | image_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOperationForImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOperationRequest** | [**CreateOperationRequest**](CreateOperationRequest.md) |  | 

### Return type

[**Operation**](Operation.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOperationForLogDrain

> Operation CreateOperationForLogDrain(ctx, logDrainId).CreateOperationRequest(createOperationRequest).Execute()

create operation

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
	logDrainId := int32(56) // int32 | log_drain_id
	createOperationRequest := *openapiclient.NewCreateOperationRequest("Type_example") // CreateOperationRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OperationsAPI.CreateOperationForLogDrain(context.Background(), logDrainId).CreateOperationRequest(createOperationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OperationsAPI.CreateOperationForLogDrain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOperationForLogDrain`: Operation
	fmt.Fprintf(os.Stdout, "Response from `OperationsAPI.CreateOperationForLogDrain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**logDrainId** | **int32** | log_drain_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOperationForLogDrainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOperationRequest** | [**CreateOperationRequest**](CreateOperationRequest.md) |  | 

### Return type

[**Operation**](Operation.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOperationForMetricDrain

> Operation CreateOperationForMetricDrain(ctx, metricDrainId).CreateOperationRequest(createOperationRequest).Execute()

create operation

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
	metricDrainId := int32(56) // int32 | metric_drain_id
	createOperationRequest := *openapiclient.NewCreateOperationRequest("Type_example") // CreateOperationRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OperationsAPI.CreateOperationForMetricDrain(context.Background(), metricDrainId).CreateOperationRequest(createOperationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OperationsAPI.CreateOperationForMetricDrain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOperationForMetricDrain`: Operation
	fmt.Fprintf(os.Stdout, "Response from `OperationsAPI.CreateOperationForMetricDrain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**metricDrainId** | **int32** | metric_drain_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOperationForMetricDrainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOperationRequest** | [**CreateOperationRequest**](CreateOperationRequest.md) |  | 

### Return type

[**Operation**](Operation.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOperationForPersistentDisk

> Operation CreateOperationForPersistentDisk(ctx, persistentDiskId).CreateOperationRequest(createOperationRequest).Execute()

create operation

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
	createOperationRequest := *openapiclient.NewCreateOperationRequest("Type_example") // CreateOperationRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OperationsAPI.CreateOperationForPersistentDisk(context.Background(), persistentDiskId).CreateOperationRequest(createOperationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OperationsAPI.CreateOperationForPersistentDisk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOperationForPersistentDisk`: Operation
	fmt.Fprintf(os.Stdout, "Response from `OperationsAPI.CreateOperationForPersistentDisk`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**persistentDiskId** | **int32** | persistent_disk_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOperationForPersistentDiskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOperationRequest** | [**CreateOperationRequest**](CreateOperationRequest.md) |  | 

### Return type

[**Operation**](Operation.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOperationForService

> Operation CreateOperationForService(ctx, serviceId).CreateOperationRequest(createOperationRequest).Execute()

create operation

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
	createOperationRequest := *openapiclient.NewCreateOperationRequest("Type_example") // CreateOperationRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OperationsAPI.CreateOperationForService(context.Background(), serviceId).CreateOperationRequest(createOperationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OperationsAPI.CreateOperationForService``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOperationForService`: Operation
	fmt.Fprintf(os.Stdout, "Response from `OperationsAPI.CreateOperationForService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **int32** | service_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOperationForServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOperationRequest** | [**CreateOperationRequest**](CreateOperationRequest.md) |  | 

### Return type

[**Operation**](Operation.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOperationForVhost

> Operation CreateOperationForVhost(ctx, vhostId).CreateOperationRequest(createOperationRequest).Execute()

create operation

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
	vhostId := int32(56) // int32 | vhost_id
	createOperationRequest := *openapiclient.NewCreateOperationRequest("Type_example") // CreateOperationRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OperationsAPI.CreateOperationForVhost(context.Background(), vhostId).CreateOperationRequest(createOperationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OperationsAPI.CreateOperationForVhost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOperationForVhost`: Operation
	fmt.Fprintf(os.Stdout, "Response from `OperationsAPI.CreateOperationForVhost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vhostId** | **int32** | vhost_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOperationForVhostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOperationRequest** | [**CreateOperationRequest**](CreateOperationRequest.md) |  | 

### Return type

[**Operation**](Operation.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOperation

> Operation GetOperation(ctx, id).Execute()

show operation

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
	resp, r, err := apiClient.OperationsAPI.GetOperation(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OperationsAPI.GetOperation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOperation`: Operation
	fmt.Fprintf(os.Stdout, "Response from `OperationsAPI.GetOperation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOperationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Operation**](Operation.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOperationsForAccount

> ListOperationsForAccount200Response ListOperationsForAccount(ctx, accountId).Page(page).Execute()

list operations

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
	resp, r, err := apiClient.OperationsAPI.ListOperationsForAccount(context.Background(), accountId).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OperationsAPI.ListOperationsForAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOperationsForAccount`: ListOperationsForAccount200Response
	fmt.Fprintf(os.Stdout, "Response from `OperationsAPI.ListOperationsForAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **int32** | account_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOperationsForAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | current page of results for pagination | 

### Return type

[**ListOperationsForAccount200Response**](ListOperationsForAccount200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOperationsForApp

> ListOperationsForAccount200Response ListOperationsForApp(ctx, appId).Page(page).Execute()

list operations

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
	resp, r, err := apiClient.OperationsAPI.ListOperationsForApp(context.Background(), appId).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OperationsAPI.ListOperationsForApp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOperationsForApp`: ListOperationsForAccount200Response
	fmt.Fprintf(os.Stdout, "Response from `OperationsAPI.ListOperationsForApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** | app_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOperationsForAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | current page of results for pagination | 

### Return type

[**ListOperationsForAccount200Response**](ListOperationsForAccount200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOperationsForBackup

> ListOperationsForAccount200Response ListOperationsForBackup(ctx, backupId).Page(page).Execute()

list operations

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
	backupId := int32(56) // int32 | backup_id
	page := int32(56) // int32 | current page of results for pagination (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OperationsAPI.ListOperationsForBackup(context.Background(), backupId).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OperationsAPI.ListOperationsForBackup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOperationsForBackup`: ListOperationsForAccount200Response
	fmt.Fprintf(os.Stdout, "Response from `OperationsAPI.ListOperationsForBackup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**backupId** | **int32** | backup_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOperationsForBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | current page of results for pagination | 

### Return type

[**ListOperationsForAccount200Response**](ListOperationsForAccount200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOperationsForDatabase

> ListOperationsForAccount200Response ListOperationsForDatabase(ctx, databaseId).Page(page).Execute()

list operations

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
	databaseId := int32(56) // int32 | database_id
	page := int32(56) // int32 | current page of results for pagination (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OperationsAPI.ListOperationsForDatabase(context.Background(), databaseId).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OperationsAPI.ListOperationsForDatabase``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOperationsForDatabase`: ListOperationsForAccount200Response
	fmt.Fprintf(os.Stdout, "Response from `OperationsAPI.ListOperationsForDatabase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**databaseId** | **int32** | database_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOperationsForDatabaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | current page of results for pagination | 

### Return type

[**ListOperationsForAccount200Response**](ListOperationsForAccount200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOperationsForDatabaseCredential

> ListOperationsForAccount200Response ListOperationsForDatabaseCredential(ctx, databaseCredentialId).Page(page).Execute()

list operations

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
	databaseCredentialId := int32(56) // int32 | database_credential_id
	page := int32(56) // int32 | current page of results for pagination (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OperationsAPI.ListOperationsForDatabaseCredential(context.Background(), databaseCredentialId).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OperationsAPI.ListOperationsForDatabaseCredential``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOperationsForDatabaseCredential`: ListOperationsForAccount200Response
	fmt.Fprintf(os.Stdout, "Response from `OperationsAPI.ListOperationsForDatabaseCredential`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**databaseCredentialId** | **int32** | database_credential_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOperationsForDatabaseCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | current page of results for pagination | 

### Return type

[**ListOperationsForAccount200Response**](ListOperationsForAccount200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOperationsForDiskAttachment

> ListOperationsForAccount200Response ListOperationsForDiskAttachment(ctx, diskAttachmentId).Page(page).Execute()

list operations

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
	diskAttachmentId := int32(56) // int32 | disk_attachment_id
	page := int32(56) // int32 | current page of results for pagination (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OperationsAPI.ListOperationsForDiskAttachment(context.Background(), diskAttachmentId).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OperationsAPI.ListOperationsForDiskAttachment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOperationsForDiskAttachment`: ListOperationsForAccount200Response
	fmt.Fprintf(os.Stdout, "Response from `OperationsAPI.ListOperationsForDiskAttachment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**diskAttachmentId** | **int32** | disk_attachment_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOperationsForDiskAttachmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | current page of results for pagination | 

### Return type

[**ListOperationsForAccount200Response**](ListOperationsForAccount200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOperationsForEphemeralSession

> ListOperationsForAccount200Response ListOperationsForEphemeralSession(ctx, ephemeralSessionId).Page(page).Execute()

list operations

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
	ephemeralSessionId := int32(56) // int32 | ephemeral_session_id
	page := int32(56) // int32 | current page of results for pagination (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OperationsAPI.ListOperationsForEphemeralSession(context.Background(), ephemeralSessionId).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OperationsAPI.ListOperationsForEphemeralSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOperationsForEphemeralSession`: ListOperationsForAccount200Response
	fmt.Fprintf(os.Stdout, "Response from `OperationsAPI.ListOperationsForEphemeralSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ephemeralSessionId** | **int32** | ephemeral_session_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOperationsForEphemeralSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | current page of results for pagination | 

### Return type

[**ListOperationsForAccount200Response**](ListOperationsForAccount200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOperationsForImage

> ListOperationsForAccount200Response ListOperationsForImage(ctx, imageId).Page(page).Execute()

list operations

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
	imageId := int32(56) // int32 | image_id
	page := int32(56) // int32 | current page of results for pagination (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OperationsAPI.ListOperationsForImage(context.Background(), imageId).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OperationsAPI.ListOperationsForImage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOperationsForImage`: ListOperationsForAccount200Response
	fmt.Fprintf(os.Stdout, "Response from `OperationsAPI.ListOperationsForImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageId** | **int32** | image_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOperationsForImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | current page of results for pagination | 

### Return type

[**ListOperationsForAccount200Response**](ListOperationsForAccount200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOperationsForLogDrain

> ListOperationsForAccount200Response ListOperationsForLogDrain(ctx, logDrainId).Page(page).Execute()

list operations

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
	logDrainId := int32(56) // int32 | log_drain_id
	page := int32(56) // int32 | current page of results for pagination (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OperationsAPI.ListOperationsForLogDrain(context.Background(), logDrainId).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OperationsAPI.ListOperationsForLogDrain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOperationsForLogDrain`: ListOperationsForAccount200Response
	fmt.Fprintf(os.Stdout, "Response from `OperationsAPI.ListOperationsForLogDrain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**logDrainId** | **int32** | log_drain_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOperationsForLogDrainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | current page of results for pagination | 

### Return type

[**ListOperationsForAccount200Response**](ListOperationsForAccount200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOperationsForMetricDrain

> ListOperationsForAccount200Response ListOperationsForMetricDrain(ctx, metricDrainId).Page(page).Execute()

list operations

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
	metricDrainId := int32(56) // int32 | metric_drain_id
	page := int32(56) // int32 | current page of results for pagination (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OperationsAPI.ListOperationsForMetricDrain(context.Background(), metricDrainId).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OperationsAPI.ListOperationsForMetricDrain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOperationsForMetricDrain`: ListOperationsForAccount200Response
	fmt.Fprintf(os.Stdout, "Response from `OperationsAPI.ListOperationsForMetricDrain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**metricDrainId** | **int32** | metric_drain_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOperationsForMetricDrainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | current page of results for pagination | 

### Return type

[**ListOperationsForAccount200Response**](ListOperationsForAccount200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOperationsForPersistentDisk

> ListOperationsForAccount200Response ListOperationsForPersistentDisk(ctx, persistentDiskId).Page(page).Execute()

list operations

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
	resp, r, err := apiClient.OperationsAPI.ListOperationsForPersistentDisk(context.Background(), persistentDiskId).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OperationsAPI.ListOperationsForPersistentDisk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOperationsForPersistentDisk`: ListOperationsForAccount200Response
	fmt.Fprintf(os.Stdout, "Response from `OperationsAPI.ListOperationsForPersistentDisk`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**persistentDiskId** | **int32** | persistent_disk_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOperationsForPersistentDiskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | current page of results for pagination | 

### Return type

[**ListOperationsForAccount200Response**](ListOperationsForAccount200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOperationsForService

> ListOperationsForAccount200Response ListOperationsForService(ctx, serviceId).Page(page).Execute()

list operations

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
	resp, r, err := apiClient.OperationsAPI.ListOperationsForService(context.Background(), serviceId).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OperationsAPI.ListOperationsForService``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOperationsForService`: ListOperationsForAccount200Response
	fmt.Fprintf(os.Stdout, "Response from `OperationsAPI.ListOperationsForService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **int32** | service_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOperationsForServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | current page of results for pagination | 

### Return type

[**ListOperationsForAccount200Response**](ListOperationsForAccount200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOperationsForVhost

> ListOperationsForAccount200Response ListOperationsForVhost(ctx, vhostId).Page(page).Execute()

list operations

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
	vhostId := int32(56) // int32 | vhost_id
	page := int32(56) // int32 | current page of results for pagination (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OperationsAPI.ListOperationsForVhost(context.Background(), vhostId).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OperationsAPI.ListOperationsForVhost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOperationsForVhost`: ListOperationsForAccount200Response
	fmt.Fprintf(os.Stdout, "Response from `OperationsAPI.ListOperationsForVhost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vhostId** | **int32** | vhost_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOperationsForVhostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | current page of results for pagination | 

### Return type

[**ListOperationsForAccount200Response**](ListOperationsForAccount200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchOperation

> PatchOperation(ctx, id).UpdateOperationRequest(updateOperationRequest).Execute()

update operation

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
	updateOperationRequest := *openapiclient.NewUpdateOperationRequest() // UpdateOperationRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OperationsAPI.PatchOperation(context.Background(), id).UpdateOperationRequest(updateOperationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OperationsAPI.PatchOperation``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPatchOperationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateOperationRequest** | [**UpdateOperationRequest**](UpdateOperationRequest.md) |  | 

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


## UpdateOperation

> UpdateOperation(ctx, id).UpdateOperationRequest(updateOperationRequest).Execute()

update operation

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
	updateOperationRequest := *openapiclient.NewUpdateOperationRequest() // UpdateOperationRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OperationsAPI.UpdateOperation(context.Background(), id).UpdateOperationRequest(updateOperationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OperationsAPI.UpdateOperation``: %v\n", err)
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

Other parameters are passed through a pointer to a apiUpdateOperationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateOperationRequest** | [**UpdateOperationRequest**](UpdateOperationRequest.md) |  | 

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

