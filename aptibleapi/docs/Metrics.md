# Metrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**MetaType** | **string** |  | 
**CreatedAt** | **string** |  | 
**UpdatedAt** | **string** |  | 
**Name** | **string** |  | 
**Description** | **string** |  | 
**Unit** | **string** |  | 
**Query** | **map[string]interface{}** |  | 
**Links** | Pointer to [**IntegrationLinks**](IntegrationLinks.md) |  | [optional] 

## Methods

### NewMetrics

`func NewMetrics(id int32, metaType string, createdAt string, updatedAt string, name string, description string, unit string, query map[string]interface{}, ) *Metrics`

NewMetrics instantiates a new Metrics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricsWithDefaults

`func NewMetricsWithDefaults() *Metrics`

NewMetricsWithDefaults instantiates a new Metrics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Metrics) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Metrics) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Metrics) SetId(v int32)`

SetId sets Id field to given value.


### GetMetaType

`func (o *Metrics) GetMetaType() string`

GetMetaType returns the MetaType field if non-nil, zero value otherwise.

### GetMetaTypeOk

`func (o *Metrics) GetMetaTypeOk() (*string, bool)`

GetMetaTypeOk returns a tuple with the MetaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaType

`func (o *Metrics) SetMetaType(v string)`

SetMetaType sets MetaType field to given value.


### GetCreatedAt

`func (o *Metrics) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Metrics) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Metrics) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Metrics) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Metrics) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Metrics) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetName

`func (o *Metrics) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Metrics) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Metrics) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Metrics) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Metrics) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Metrics) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetUnit

`func (o *Metrics) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *Metrics) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *Metrics) SetUnit(v string)`

SetUnit sets Unit field to given value.


### GetQuery

`func (o *Metrics) GetQuery() map[string]interface{}`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *Metrics) GetQueryOk() (*map[string]interface{}, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *Metrics) SetQuery(v map[string]interface{})`

SetQuery sets Query field to given value.


### GetLinks

`func (o *Metrics) GetLinks() IntegrationLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Metrics) GetLinksOk() (*IntegrationLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Metrics) SetLinks(v IntegrationLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Metrics) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


