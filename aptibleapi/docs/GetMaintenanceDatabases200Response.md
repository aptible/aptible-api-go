# GetMaintenanceDatabases200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Embedded** | Pointer to [**GetMaintenanceDatabases200ResponseEmbedded**](GetMaintenanceDatabases200ResponseEmbedded.md) |  | [optional] 
**TotalCount** | Pointer to **int32** |  | [optional] 
**PerPage** | Pointer to **int32** |  | [optional] 
**CurrentPage** | Pointer to **int32** |  | [optional] 
**Links** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewGetMaintenanceDatabases200Response

`func NewGetMaintenanceDatabases200Response() *GetMaintenanceDatabases200Response`

NewGetMaintenanceDatabases200Response instantiates a new GetMaintenanceDatabases200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMaintenanceDatabases200ResponseWithDefaults

`func NewGetMaintenanceDatabases200ResponseWithDefaults() *GetMaintenanceDatabases200Response`

NewGetMaintenanceDatabases200ResponseWithDefaults instantiates a new GetMaintenanceDatabases200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmbedded

`func (o *GetMaintenanceDatabases200Response) GetEmbedded() GetMaintenanceDatabases200ResponseEmbedded`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *GetMaintenanceDatabases200Response) GetEmbeddedOk() (*GetMaintenanceDatabases200ResponseEmbedded, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *GetMaintenanceDatabases200Response) SetEmbedded(v GetMaintenanceDatabases200ResponseEmbedded)`

SetEmbedded sets Embedded field to given value.

### HasEmbedded

`func (o *GetMaintenanceDatabases200Response) HasEmbedded() bool`

HasEmbedded returns a boolean if a field has been set.

### GetTotalCount

`func (o *GetMaintenanceDatabases200Response) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *GetMaintenanceDatabases200Response) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *GetMaintenanceDatabases200Response) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *GetMaintenanceDatabases200Response) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetPerPage

`func (o *GetMaintenanceDatabases200Response) GetPerPage() int32`

GetPerPage returns the PerPage field if non-nil, zero value otherwise.

### GetPerPageOk

`func (o *GetMaintenanceDatabases200Response) GetPerPageOk() (*int32, bool)`

GetPerPageOk returns a tuple with the PerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerPage

`func (o *GetMaintenanceDatabases200Response) SetPerPage(v int32)`

SetPerPage sets PerPage field to given value.

### HasPerPage

`func (o *GetMaintenanceDatabases200Response) HasPerPage() bool`

HasPerPage returns a boolean if a field has been set.

### GetCurrentPage

`func (o *GetMaintenanceDatabases200Response) GetCurrentPage() int32`

GetCurrentPage returns the CurrentPage field if non-nil, zero value otherwise.

### GetCurrentPageOk

`func (o *GetMaintenanceDatabases200Response) GetCurrentPageOk() (*int32, bool)`

GetCurrentPageOk returns a tuple with the CurrentPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPage

`func (o *GetMaintenanceDatabases200Response) SetCurrentPage(v int32)`

SetCurrentPage sets CurrentPage field to given value.

### HasCurrentPage

`func (o *GetMaintenanceDatabases200Response) HasCurrentPage() bool`

HasCurrentPage returns a boolean if a field has been set.

### GetLinks

`func (o *GetMaintenanceDatabases200Response) GetLinks() map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetMaintenanceDatabases200Response) GetLinksOk() (*map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetMaintenanceDatabases200Response) SetLinks(v map[string]interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GetMaintenanceDatabases200Response) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


