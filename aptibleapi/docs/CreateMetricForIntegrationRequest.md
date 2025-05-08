# CreateMetricForIntegrationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | **string** |  | 
**Unit** | **string** |  | 
**Query** | **map[string]interface{}** |  | 

## Methods

### NewCreateMetricForIntegrationRequest

`func NewCreateMetricForIntegrationRequest(name string, description string, unit string, query map[string]interface{}, ) *CreateMetricForIntegrationRequest`

NewCreateMetricForIntegrationRequest instantiates a new CreateMetricForIntegrationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateMetricForIntegrationRequestWithDefaults

`func NewCreateMetricForIntegrationRequestWithDefaults() *CreateMetricForIntegrationRequest`

NewCreateMetricForIntegrationRequestWithDefaults instantiates a new CreateMetricForIntegrationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateMetricForIntegrationRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateMetricForIntegrationRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateMetricForIntegrationRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateMetricForIntegrationRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateMetricForIntegrationRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateMetricForIntegrationRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetUnit

`func (o *CreateMetricForIntegrationRequest) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *CreateMetricForIntegrationRequest) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *CreateMetricForIntegrationRequest) SetUnit(v string)`

SetUnit sets Unit field to given value.


### GetQuery

`func (o *CreateMetricForIntegrationRequest) GetQuery() map[string]interface{}`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *CreateMetricForIntegrationRequest) GetQueryOk() (*map[string]interface{}, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *CreateMetricForIntegrationRequest) SetQuery(v map[string]interface{})`

SetQuery sets Query field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


