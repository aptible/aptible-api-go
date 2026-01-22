# ListLlmKeys200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Embedded** | Pointer to [**ListLlmKeysForAccount200ResponseEmbedded**](ListLlmKeysForAccount200ResponseEmbedded.md) |  | [optional] 
**Links** | Pointer to **map[string]interface{}** |  | [optional] 
**TotalCount** | Pointer to **int32** |  | [optional] 
**PerPage** | Pointer to **int32** |  | [optional] 
**CurrentPage** | Pointer to **int32** |  | [optional] 

## Methods

### NewListLlmKeys200Response

`func NewListLlmKeys200Response() *ListLlmKeys200Response`

NewListLlmKeys200Response instantiates a new ListLlmKeys200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListLlmKeys200ResponseWithDefaults

`func NewListLlmKeys200ResponseWithDefaults() *ListLlmKeys200Response`

NewListLlmKeys200ResponseWithDefaults instantiates a new ListLlmKeys200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmbedded

`func (o *ListLlmKeys200Response) GetEmbedded() ListLlmKeysForAccount200ResponseEmbedded`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *ListLlmKeys200Response) GetEmbeddedOk() (*ListLlmKeysForAccount200ResponseEmbedded, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *ListLlmKeys200Response) SetEmbedded(v ListLlmKeysForAccount200ResponseEmbedded)`

SetEmbedded sets Embedded field to given value.

### HasEmbedded

`func (o *ListLlmKeys200Response) HasEmbedded() bool`

HasEmbedded returns a boolean if a field has been set.

### GetLinks

`func (o *ListLlmKeys200Response) GetLinks() map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ListLlmKeys200Response) GetLinksOk() (*map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ListLlmKeys200Response) SetLinks(v map[string]interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ListLlmKeys200Response) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetTotalCount

`func (o *ListLlmKeys200Response) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ListLlmKeys200Response) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ListLlmKeys200Response) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *ListLlmKeys200Response) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetPerPage

`func (o *ListLlmKeys200Response) GetPerPage() int32`

GetPerPage returns the PerPage field if non-nil, zero value otherwise.

### GetPerPageOk

`func (o *ListLlmKeys200Response) GetPerPageOk() (*int32, bool)`

GetPerPageOk returns a tuple with the PerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerPage

`func (o *ListLlmKeys200Response) SetPerPage(v int32)`

SetPerPage sets PerPage field to given value.

### HasPerPage

`func (o *ListLlmKeys200Response) HasPerPage() bool`

HasPerPage returns a boolean if a field has been set.

### GetCurrentPage

`func (o *ListLlmKeys200Response) GetCurrentPage() int32`

GetCurrentPage returns the CurrentPage field if non-nil, zero value otherwise.

### GetCurrentPageOk

`func (o *ListLlmKeys200Response) GetCurrentPageOk() (*int32, bool)`

GetCurrentPageOk returns a tuple with the CurrentPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPage

`func (o *ListLlmKeys200Response) SetCurrentPage(v int32)`

SetCurrentPage sets CurrentPage field to given value.

### HasCurrentPage

`func (o *ListLlmKeys200Response) HasCurrentPage() bool`

HasCurrentPage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


