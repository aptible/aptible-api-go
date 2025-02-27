# Observation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**MetaType** | **string** |  | 
**CreatedAt** | **string** |  | 
**UpdatedAt** | **string** |  | 
**RangeBegin** | **string** |  | 
**RangeEnd** | **string** |  | 
**Data** | **map[string]interface{}** |  | 
**Links** | Pointer to [**ConfigurationLinks**](ConfigurationLinks.md) |  | [optional] 

## Methods

### NewObservation

`func NewObservation(id int32, metaType string, createdAt string, updatedAt string, rangeBegin string, rangeEnd string, data map[string]interface{}, ) *Observation`

NewObservation instantiates a new Observation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObservationWithDefaults

`func NewObservationWithDefaults() *Observation`

NewObservationWithDefaults instantiates a new Observation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Observation) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Observation) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Observation) SetId(v int32)`

SetId sets Id field to given value.


### GetMetaType

`func (o *Observation) GetMetaType() string`

GetMetaType returns the MetaType field if non-nil, zero value otherwise.

### GetMetaTypeOk

`func (o *Observation) GetMetaTypeOk() (*string, bool)`

GetMetaTypeOk returns a tuple with the MetaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaType

`func (o *Observation) SetMetaType(v string)`

SetMetaType sets MetaType field to given value.


### GetCreatedAt

`func (o *Observation) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Observation) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Observation) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Observation) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Observation) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Observation) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetRangeBegin

`func (o *Observation) GetRangeBegin() string`

GetRangeBegin returns the RangeBegin field if non-nil, zero value otherwise.

### GetRangeBeginOk

`func (o *Observation) GetRangeBeginOk() (*string, bool)`

GetRangeBeginOk returns a tuple with the RangeBegin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeBegin

`func (o *Observation) SetRangeBegin(v string)`

SetRangeBegin sets RangeBegin field to given value.


### GetRangeEnd

`func (o *Observation) GetRangeEnd() string`

GetRangeEnd returns the RangeEnd field if non-nil, zero value otherwise.

### GetRangeEndOk

`func (o *Observation) GetRangeEndOk() (*string, bool)`

GetRangeEndOk returns a tuple with the RangeEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeEnd

`func (o *Observation) SetRangeEnd(v string)`

SetRangeEnd sets RangeEnd field to given value.


### GetData

`func (o *Observation) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *Observation) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *Observation) SetData(v map[string]interface{})`

SetData sets Data field to given value.


### GetLinks

`func (o *Observation) GetLinks() ConfigurationLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Observation) GetLinksOk() (*ConfigurationLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Observation) SetLinks(v ConfigurationLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Observation) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


