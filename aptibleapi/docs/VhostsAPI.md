# \VhostsAPI

All URIs are relative to *https://api.aptible.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateVhost**](VhostsAPI.md#CreateVhost) | **Post** /services/{service_id}/vhosts | create vhost
[**DeleteVhost**](VhostsAPI.md#DeleteVhost) | **Delete** /vhosts/{id} | delete vhost
[**GetVhost**](VhostsAPI.md#GetVhost) | **Get** /vhosts/{id} | show vhost
[**ListVhostsForAccount**](VhostsAPI.md#ListVhostsForAccount) | **Get** /accounts/{account_id}/vhosts | list vhosts
[**ListVhostsForApp**](VhostsAPI.md#ListVhostsForApp) | **Get** /apps/{app_id}/vhosts | list vhosts
[**ListVhostsForCertificate**](VhostsAPI.md#ListVhostsForCertificate) | **Get** /certificates/{certificate_id}/vhosts | list vhosts
[**ListVhostsForService**](VhostsAPI.md#ListVhostsForService) | **Get** /services/{service_id}/vhosts | list vhosts
[**PatchVhost**](VhostsAPI.md#PatchVhost) | **Patch** /vhosts/{id} | update vhost
[**UpdateVhost**](VhostsAPI.md#UpdateVhost) | **Put** /vhosts/{id} | update vhost



## CreateVhost

> Vhost CreateVhost(ctx, serviceId).CreateVhostRequest(createVhostRequest).Execute()

create vhost

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
	createVhostRequest := *openapiclient.NewCreateVhostRequest("Type_example") // CreateVhostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VhostsAPI.CreateVhost(context.Background(), serviceId).CreateVhostRequest(createVhostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VhostsAPI.CreateVhost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateVhost`: Vhost
	fmt.Fprintf(os.Stdout, "Response from `VhostsAPI.CreateVhost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **int32** | service_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVhostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createVhostRequest** | [**CreateVhostRequest**](CreateVhostRequest.md) |  | 

### Return type

[**Vhost**](Vhost.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteVhost

> DeleteVhost(ctx, id).Execute()

delete vhost

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
	r, err := apiClient.VhostsAPI.DeleteVhost(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VhostsAPI.DeleteVhost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteVhostRequest struct via the builder pattern


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


## GetVhost

> Vhost GetVhost(ctx, id).Execute()

show vhost

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
	resp, r, err := apiClient.VhostsAPI.GetVhost(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VhostsAPI.GetVhost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVhost`: Vhost
	fmt.Fprintf(os.Stdout, "Response from `VhostsAPI.GetVhost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVhostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Vhost**](Vhost.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVhostsForAccount

> ListVhostsForAccount200Response ListVhostsForAccount(ctx, accountId).Page(page).Execute()

list vhosts

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
	resp, r, err := apiClient.VhostsAPI.ListVhostsForAccount(context.Background(), accountId).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VhostsAPI.ListVhostsForAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListVhostsForAccount`: ListVhostsForAccount200Response
	fmt.Fprintf(os.Stdout, "Response from `VhostsAPI.ListVhostsForAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **int32** | account_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiListVhostsForAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | current page of results for pagination | 

### Return type

[**ListVhostsForAccount200Response**](ListVhostsForAccount200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVhostsForApp

> ListVhostsForAccount200Response ListVhostsForApp(ctx, appId).Page(page).Execute()

list vhosts

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
	resp, r, err := apiClient.VhostsAPI.ListVhostsForApp(context.Background(), appId).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VhostsAPI.ListVhostsForApp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListVhostsForApp`: ListVhostsForAccount200Response
	fmt.Fprintf(os.Stdout, "Response from `VhostsAPI.ListVhostsForApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** | app_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiListVhostsForAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | current page of results for pagination | 

### Return type

[**ListVhostsForAccount200Response**](ListVhostsForAccount200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVhostsForCertificate

> ListVhostsForAccount200Response ListVhostsForCertificate(ctx, certificateId).Page(page).Execute()

list vhosts

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
	certificateId := int32(56) // int32 | certificate_id
	page := int32(56) // int32 | current page of results for pagination (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VhostsAPI.ListVhostsForCertificate(context.Background(), certificateId).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VhostsAPI.ListVhostsForCertificate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListVhostsForCertificate`: ListVhostsForAccount200Response
	fmt.Fprintf(os.Stdout, "Response from `VhostsAPI.ListVhostsForCertificate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**certificateId** | **int32** | certificate_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiListVhostsForCertificateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | current page of results for pagination | 

### Return type

[**ListVhostsForAccount200Response**](ListVhostsForAccount200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVhostsForService

> ListVhostsForAccount200Response ListVhostsForService(ctx, serviceId).Page(page).Execute()

list vhosts

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
	resp, r, err := apiClient.VhostsAPI.ListVhostsForService(context.Background(), serviceId).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VhostsAPI.ListVhostsForService``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListVhostsForService`: ListVhostsForAccount200Response
	fmt.Fprintf(os.Stdout, "Response from `VhostsAPI.ListVhostsForService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **int32** | service_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiListVhostsForServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | current page of results for pagination | 

### Return type

[**ListVhostsForAccount200Response**](ListVhostsForAccount200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchVhost

> PatchVhost(ctx, id).UpdateVhostRequest(updateVhostRequest).Execute()

update vhost

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
	updateVhostRequest := *openapiclient.NewUpdateVhostRequest() // UpdateVhostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VhostsAPI.PatchVhost(context.Background(), id).UpdateVhostRequest(updateVhostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VhostsAPI.PatchVhost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPatchVhostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateVhostRequest** | [**UpdateVhostRequest**](UpdateVhostRequest.md) |  | 

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


## UpdateVhost

> UpdateVhost(ctx, id).UpdateVhostRequest(updateVhostRequest).Execute()

update vhost

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
	updateVhostRequest := *openapiclient.NewUpdateVhostRequest() // UpdateVhostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VhostsAPI.UpdateVhost(context.Background(), id).UpdateVhostRequest(updateVhostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VhostsAPI.UpdateVhost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiUpdateVhostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateVhostRequest** | [**UpdateVhostRequest**](UpdateVhostRequest.md) |  | 

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

