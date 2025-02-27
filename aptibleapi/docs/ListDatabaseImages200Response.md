# ListDatabaseImages200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Embedded** | [**ListDatabaseImages200ResponseEmbedded**](ListDatabaseImages200ResponseEmbedded.md) |  | 
**TotalCount** | **int32** |  | 
**PerPage** | **int32** |  | 
**CurrentPage** | **int32** |  | 
**Links** | [**ListDashboards200ResponseLinks**](ListDashboards200ResponseLinks.md) |  | 

## Methods

### NewListDatabaseImages200Response

`func NewListDatabaseImages200Response(embedded ListDatabaseImages200ResponseEmbedded, totalCount int32, perPage int32, currentPage int32, links ListDashboards200ResponseLinks, ) *ListDatabaseImages200Response`

NewListDatabaseImages200Response instantiates a new ListDatabaseImages200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListDatabaseImages200ResponseWithDefaults

`func NewListDatabaseImages200ResponseWithDefaults() *ListDatabaseImages200Response`

NewListDatabaseImages200ResponseWithDefaults instantiates a new ListDatabaseImages200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmbedded

`func (o *ListDatabaseImages200Response) GetEmbedded() ListDatabaseImages200ResponseEmbedded`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *ListDatabaseImages200Response) GetEmbeddedOk() (*ListDatabaseImages200ResponseEmbedded, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *ListDatabaseImages200Response) SetEmbedded(v ListDatabaseImages200ResponseEmbedded)`

SetEmbedded sets Embedded field to given value.


### GetTotalCount

`func (o *ListDatabaseImages200Response) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ListDatabaseImages200Response) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ListDatabaseImages200Response) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.


### GetPerPage

`func (o *ListDatabaseImages200Response) GetPerPage() int32`

GetPerPage returns the PerPage field if non-nil, zero value otherwise.

### GetPerPageOk

`func (o *ListDatabaseImages200Response) GetPerPageOk() (*int32, bool)`

GetPerPageOk returns a tuple with the PerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerPage

`func (o *ListDatabaseImages200Response) SetPerPage(v int32)`

SetPerPage sets PerPage field to given value.


### GetCurrentPage

`func (o *ListDatabaseImages200Response) GetCurrentPage() int32`

GetCurrentPage returns the CurrentPage field if non-nil, zero value otherwise.

### GetCurrentPageOk

`func (o *ListDatabaseImages200Response) GetCurrentPageOk() (*int32, bool)`

GetCurrentPageOk returns a tuple with the CurrentPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPage

`func (o *ListDatabaseImages200Response) SetCurrentPage(v int32)`

SetCurrentPage sets CurrentPage field to given value.


### GetLinks

`func (o *ListDatabaseImages200Response) GetLinks() ListDashboards200ResponseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ListDatabaseImages200Response) GetLinksOk() (*ListDashboards200ResponseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ListDatabaseImages200Response) SetLinks(v ListDashboards200ResponseLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


