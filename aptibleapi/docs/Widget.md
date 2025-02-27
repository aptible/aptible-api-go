# Widget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**MetaType** | **string** |  | 
**CreatedAt** | **string** |  | 
**UpdatedAt** | **string** |  | 
**Title** | **string** |  | 
**Summary** | **string** |  | 
**Rating** | **int32** |  | 
**WidgetType** | **string** |  | 
**Data** | **map[string]interface{}** |  | 
**RangeBegin** | **string** |  | 
**RangeEnd** | **string** |  | 
**Links** | Pointer to [**WidgetLinks**](WidgetLinks.md) |  | [optional] 

## Methods

### NewWidget

`func NewWidget(id string, metaType string, createdAt string, updatedAt string, title string, summary string, rating int32, widgetType string, data map[string]interface{}, rangeBegin string, rangeEnd string, ) *Widget`

NewWidget instantiates a new Widget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWidgetWithDefaults

`func NewWidgetWithDefaults() *Widget`

NewWidgetWithDefaults instantiates a new Widget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Widget) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Widget) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Widget) SetId(v string)`

SetId sets Id field to given value.


### GetMetaType

`func (o *Widget) GetMetaType() string`

GetMetaType returns the MetaType field if non-nil, zero value otherwise.

### GetMetaTypeOk

`func (o *Widget) GetMetaTypeOk() (*string, bool)`

GetMetaTypeOk returns a tuple with the MetaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaType

`func (o *Widget) SetMetaType(v string)`

SetMetaType sets MetaType field to given value.


### GetCreatedAt

`func (o *Widget) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Widget) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Widget) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Widget) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Widget) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Widget) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetTitle

`func (o *Widget) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Widget) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Widget) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetSummary

`func (o *Widget) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *Widget) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *Widget) SetSummary(v string)`

SetSummary sets Summary field to given value.


### GetRating

`func (o *Widget) GetRating() int32`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *Widget) GetRatingOk() (*int32, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *Widget) SetRating(v int32)`

SetRating sets Rating field to given value.


### GetWidgetType

`func (o *Widget) GetWidgetType() string`

GetWidgetType returns the WidgetType field if non-nil, zero value otherwise.

### GetWidgetTypeOk

`func (o *Widget) GetWidgetTypeOk() (*string, bool)`

GetWidgetTypeOk returns a tuple with the WidgetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidgetType

`func (o *Widget) SetWidgetType(v string)`

SetWidgetType sets WidgetType field to given value.


### GetData

`func (o *Widget) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *Widget) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *Widget) SetData(v map[string]interface{})`

SetData sets Data field to given value.


### GetRangeBegin

`func (o *Widget) GetRangeBegin() string`

GetRangeBegin returns the RangeBegin field if non-nil, zero value otherwise.

### GetRangeBeginOk

`func (o *Widget) GetRangeBeginOk() (*string, bool)`

GetRangeBeginOk returns a tuple with the RangeBegin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeBegin

`func (o *Widget) SetRangeBegin(v string)`

SetRangeBegin sets RangeBegin field to given value.


### GetRangeEnd

`func (o *Widget) GetRangeEnd() string`

GetRangeEnd returns the RangeEnd field if non-nil, zero value otherwise.

### GetRangeEndOk

`func (o *Widget) GetRangeEndOk() (*string, bool)`

GetRangeEndOk returns a tuple with the RangeEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeEnd

`func (o *Widget) SetRangeEnd(v string)`

SetRangeEnd sets RangeEnd field to given value.


### GetLinks

`func (o *Widget) GetLinks() WidgetLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Widget) GetLinksOk() (*WidgetLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Widget) SetLinks(v WidgetLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Widget) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


