# ListExternalAwsResources200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Embedded** | [**ListExternalAwsResourcesForExternalAwsAccount200ResponseEmbedded**](ListExternalAwsResourcesForExternalAwsAccount200ResponseEmbedded.md) |  | 
**TotalCount** | **int32** |  | 
**PerPage** | **int32** |  | 
**CurrentPage** | **int32** |  | 
**Links** | [**ListDatabaseImages200ResponseLinks**](ListDatabaseImages200ResponseLinks.md) |  | 

## Methods

### NewListExternalAwsResources200Response

`func NewListExternalAwsResources200Response(embedded ListExternalAwsResourcesForExternalAwsAccount200ResponseEmbedded, totalCount int32, perPage int32, currentPage int32, links ListDatabaseImages200ResponseLinks, ) *ListExternalAwsResources200Response`

NewListExternalAwsResources200Response instantiates a new ListExternalAwsResources200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListExternalAwsResources200ResponseWithDefaults

`func NewListExternalAwsResources200ResponseWithDefaults() *ListExternalAwsResources200Response`

NewListExternalAwsResources200ResponseWithDefaults instantiates a new ListExternalAwsResources200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmbedded

`func (o *ListExternalAwsResources200Response) GetEmbedded() ListExternalAwsResourcesForExternalAwsAccount200ResponseEmbedded`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *ListExternalAwsResources200Response) GetEmbeddedOk() (*ListExternalAwsResourcesForExternalAwsAccount200ResponseEmbedded, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *ListExternalAwsResources200Response) SetEmbedded(v ListExternalAwsResourcesForExternalAwsAccount200ResponseEmbedded)`

SetEmbedded sets Embedded field to given value.


### GetTotalCount

`func (o *ListExternalAwsResources200Response) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ListExternalAwsResources200Response) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ListExternalAwsResources200Response) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.


### GetPerPage

`func (o *ListExternalAwsResources200Response) GetPerPage() int32`

GetPerPage returns the PerPage field if non-nil, zero value otherwise.

### GetPerPageOk

`func (o *ListExternalAwsResources200Response) GetPerPageOk() (*int32, bool)`

GetPerPageOk returns a tuple with the PerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerPage

`func (o *ListExternalAwsResources200Response) SetPerPage(v int32)`

SetPerPage sets PerPage field to given value.


### GetCurrentPage

`func (o *ListExternalAwsResources200Response) GetCurrentPage() int32`

GetCurrentPage returns the CurrentPage field if non-nil, zero value otherwise.

### GetCurrentPageOk

`func (o *ListExternalAwsResources200Response) GetCurrentPageOk() (*int32, bool)`

GetCurrentPageOk returns a tuple with the CurrentPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPage

`func (o *ListExternalAwsResources200Response) SetCurrentPage(v int32)`

SetCurrentPage sets CurrentPage field to given value.


### GetLinks

`func (o *ListExternalAwsResources200Response) GetLinks() ListDatabaseImages200ResponseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ListExternalAwsResources200Response) GetLinksOk() (*ListDatabaseImages200ResponseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ListExternalAwsResources200Response) SetLinks(v ListDatabaseImages200ResponseLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


