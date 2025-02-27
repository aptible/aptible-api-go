# ListConfigurationsForDatabase200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Embedded** | [**ListConfigurationsForDatabase200ResponseEmbedded**](ListConfigurationsForDatabase200ResponseEmbedded.md) |  | 
**TotalCount** | **int32** |  | 
**PerPage** | **int32** |  | 
**CurrentPage** | **int32** |  | 
**Links** | [**ListConfigurationsForDatabase200ResponseLinks**](ListConfigurationsForDatabase200ResponseLinks.md) |  | 

## Methods

### NewListConfigurationsForDatabase200Response

`func NewListConfigurationsForDatabase200Response(embedded ListConfigurationsForDatabase200ResponseEmbedded, totalCount int32, perPage int32, currentPage int32, links ListConfigurationsForDatabase200ResponseLinks, ) *ListConfigurationsForDatabase200Response`

NewListConfigurationsForDatabase200Response instantiates a new ListConfigurationsForDatabase200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListConfigurationsForDatabase200ResponseWithDefaults

`func NewListConfigurationsForDatabase200ResponseWithDefaults() *ListConfigurationsForDatabase200Response`

NewListConfigurationsForDatabase200ResponseWithDefaults instantiates a new ListConfigurationsForDatabase200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmbedded

`func (o *ListConfigurationsForDatabase200Response) GetEmbedded() ListConfigurationsForDatabase200ResponseEmbedded`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *ListConfigurationsForDatabase200Response) GetEmbeddedOk() (*ListConfigurationsForDatabase200ResponseEmbedded, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *ListConfigurationsForDatabase200Response) SetEmbedded(v ListConfigurationsForDatabase200ResponseEmbedded)`

SetEmbedded sets Embedded field to given value.


### GetTotalCount

`func (o *ListConfigurationsForDatabase200Response) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ListConfigurationsForDatabase200Response) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ListConfigurationsForDatabase200Response) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.


### GetPerPage

`func (o *ListConfigurationsForDatabase200Response) GetPerPage() int32`

GetPerPage returns the PerPage field if non-nil, zero value otherwise.

### GetPerPageOk

`func (o *ListConfigurationsForDatabase200Response) GetPerPageOk() (*int32, bool)`

GetPerPageOk returns a tuple with the PerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerPage

`func (o *ListConfigurationsForDatabase200Response) SetPerPage(v int32)`

SetPerPage sets PerPage field to given value.


### GetCurrentPage

`func (o *ListConfigurationsForDatabase200Response) GetCurrentPage() int32`

GetCurrentPage returns the CurrentPage field if non-nil, zero value otherwise.

### GetCurrentPageOk

`func (o *ListConfigurationsForDatabase200Response) GetCurrentPageOk() (*int32, bool)`

GetCurrentPageOk returns a tuple with the CurrentPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPage

`func (o *ListConfigurationsForDatabase200Response) SetCurrentPage(v int32)`

SetCurrentPage sets CurrentPage field to given value.


### GetLinks

`func (o *ListConfigurationsForDatabase200Response) GetLinks() ListConfigurationsForDatabase200ResponseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ListConfigurationsForDatabase200Response) GetLinksOk() (*ListConfigurationsForDatabase200ResponseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ListConfigurationsForDatabase200Response) SetLinks(v ListConfigurationsForDatabase200ResponseLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


