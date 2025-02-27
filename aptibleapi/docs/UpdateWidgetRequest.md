# UpdateWidgetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** |  | [optional] 
**Summary** | Pointer to **string** |  | [optional] 
**Rating** | Pointer to **int32** |  | [optional] 
**WidgetType** | Pointer to **string** |  | [optional] 
**Data** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewUpdateWidgetRequest

`func NewUpdateWidgetRequest() *UpdateWidgetRequest`

NewUpdateWidgetRequest instantiates a new UpdateWidgetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateWidgetRequestWithDefaults

`func NewUpdateWidgetRequestWithDefaults() *UpdateWidgetRequest`

NewUpdateWidgetRequestWithDefaults instantiates a new UpdateWidgetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *UpdateWidgetRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *UpdateWidgetRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *UpdateWidgetRequest) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *UpdateWidgetRequest) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetSummary

`func (o *UpdateWidgetRequest) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *UpdateWidgetRequest) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *UpdateWidgetRequest) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *UpdateWidgetRequest) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetRating

`func (o *UpdateWidgetRequest) GetRating() int32`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *UpdateWidgetRequest) GetRatingOk() (*int32, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *UpdateWidgetRequest) SetRating(v int32)`

SetRating sets Rating field to given value.

### HasRating

`func (o *UpdateWidgetRequest) HasRating() bool`

HasRating returns a boolean if a field has been set.

### GetWidgetType

`func (o *UpdateWidgetRequest) GetWidgetType() string`

GetWidgetType returns the WidgetType field if non-nil, zero value otherwise.

### GetWidgetTypeOk

`func (o *UpdateWidgetRequest) GetWidgetTypeOk() (*string, bool)`

GetWidgetTypeOk returns a tuple with the WidgetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidgetType

`func (o *UpdateWidgetRequest) SetWidgetType(v string)`

SetWidgetType sets WidgetType field to given value.

### HasWidgetType

`func (o *UpdateWidgetRequest) HasWidgetType() bool`

HasWidgetType returns a boolean if a field has been set.

### GetData

`func (o *UpdateWidgetRequest) GetData() map[string]string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *UpdateWidgetRequest) GetDataOk() (*map[string]string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *UpdateWidgetRequest) SetData(v map[string]string)`

SetData sets Data field to given value.

### HasData

`func (o *UpdateWidgetRequest) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


