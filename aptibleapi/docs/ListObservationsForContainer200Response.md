# ListObservationsForContainer200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Embedded** | [**ListObservationsForApp200ResponseEmbedded**](ListObservationsForApp200ResponseEmbedded.md) |  | 
**TotalCount** | **int32** |  | 
**PerPage** | **int32** |  | 
**CurrentPage** | **int32** |  | 
**Links** | [**ListObservationsForContainer200ResponseLinks**](ListObservationsForContainer200ResponseLinks.md) |  | 

## Methods

### NewListObservationsForContainer200Response

`func NewListObservationsForContainer200Response(embedded ListObservationsForApp200ResponseEmbedded, totalCount int32, perPage int32, currentPage int32, links ListObservationsForContainer200ResponseLinks, ) *ListObservationsForContainer200Response`

NewListObservationsForContainer200Response instantiates a new ListObservationsForContainer200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListObservationsForContainer200ResponseWithDefaults

`func NewListObservationsForContainer200ResponseWithDefaults() *ListObservationsForContainer200Response`

NewListObservationsForContainer200ResponseWithDefaults instantiates a new ListObservationsForContainer200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmbedded

`func (o *ListObservationsForContainer200Response) GetEmbedded() ListObservationsForApp200ResponseEmbedded`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *ListObservationsForContainer200Response) GetEmbeddedOk() (*ListObservationsForApp200ResponseEmbedded, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *ListObservationsForContainer200Response) SetEmbedded(v ListObservationsForApp200ResponseEmbedded)`

SetEmbedded sets Embedded field to given value.


### GetTotalCount

`func (o *ListObservationsForContainer200Response) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ListObservationsForContainer200Response) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ListObservationsForContainer200Response) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.


### GetPerPage

`func (o *ListObservationsForContainer200Response) GetPerPage() int32`

GetPerPage returns the PerPage field if non-nil, zero value otherwise.

### GetPerPageOk

`func (o *ListObservationsForContainer200Response) GetPerPageOk() (*int32, bool)`

GetPerPageOk returns a tuple with the PerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerPage

`func (o *ListObservationsForContainer200Response) SetPerPage(v int32)`

SetPerPage sets PerPage field to given value.


### GetCurrentPage

`func (o *ListObservationsForContainer200Response) GetCurrentPage() int32`

GetCurrentPage returns the CurrentPage field if non-nil, zero value otherwise.

### GetCurrentPageOk

`func (o *ListObservationsForContainer200Response) GetCurrentPageOk() (*int32, bool)`

GetCurrentPageOk returns a tuple with the CurrentPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPage

`func (o *ListObservationsForContainer200Response) SetCurrentPage(v int32)`

SetCurrentPage sets CurrentPage field to given value.


### GetLinks

`func (o *ListObservationsForContainer200Response) GetLinks() ListObservationsForContainer200ResponseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ListObservationsForContainer200Response) GetLinksOk() (*ListObservationsForContainer200ResponseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ListObservationsForContainer200Response) SetLinks(v ListObservationsForContainer200ResponseLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


