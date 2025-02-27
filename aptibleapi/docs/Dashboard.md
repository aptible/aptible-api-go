# Dashboard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**MetaType** | **string** |  | 
**CreatedAt** | **string** |  | 
**UpdatedAt** | **string** |  | 
**Name** | **string** |  | 
**OrganizationId** | **string** |  | 
**RangeBegin** | **string** |  | 
**RangeEnd** | **string** |  | 
**ObservationTimestamp** | **string** |  | 
**ResourceId** | **int32** |  | 
**ResourceType** | **string** |  | 
**Data** | **map[string]interface{}** |  | 
**Links** | Pointer to [**ConfigurationLinks**](ConfigurationLinks.md) |  | [optional] 

## Methods

### NewDashboard

`func NewDashboard(id string, metaType string, createdAt string, updatedAt string, name string, organizationId string, rangeBegin string, rangeEnd string, observationTimestamp string, resourceId int32, resourceType string, data map[string]interface{}, ) *Dashboard`

NewDashboard instantiates a new Dashboard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDashboardWithDefaults

`func NewDashboardWithDefaults() *Dashboard`

NewDashboardWithDefaults instantiates a new Dashboard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Dashboard) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Dashboard) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Dashboard) SetId(v string)`

SetId sets Id field to given value.


### GetMetaType

`func (o *Dashboard) GetMetaType() string`

GetMetaType returns the MetaType field if non-nil, zero value otherwise.

### GetMetaTypeOk

`func (o *Dashboard) GetMetaTypeOk() (*string, bool)`

GetMetaTypeOk returns a tuple with the MetaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaType

`func (o *Dashboard) SetMetaType(v string)`

SetMetaType sets MetaType field to given value.


### GetCreatedAt

`func (o *Dashboard) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Dashboard) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Dashboard) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Dashboard) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Dashboard) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Dashboard) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetName

`func (o *Dashboard) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Dashboard) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Dashboard) SetName(v string)`

SetName sets Name field to given value.


### GetOrganizationId

`func (o *Dashboard) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *Dashboard) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *Dashboard) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.


### GetRangeBegin

`func (o *Dashboard) GetRangeBegin() string`

GetRangeBegin returns the RangeBegin field if non-nil, zero value otherwise.

### GetRangeBeginOk

`func (o *Dashboard) GetRangeBeginOk() (*string, bool)`

GetRangeBeginOk returns a tuple with the RangeBegin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeBegin

`func (o *Dashboard) SetRangeBegin(v string)`

SetRangeBegin sets RangeBegin field to given value.


### GetRangeEnd

`func (o *Dashboard) GetRangeEnd() string`

GetRangeEnd returns the RangeEnd field if non-nil, zero value otherwise.

### GetRangeEndOk

`func (o *Dashboard) GetRangeEndOk() (*string, bool)`

GetRangeEndOk returns a tuple with the RangeEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeEnd

`func (o *Dashboard) SetRangeEnd(v string)`

SetRangeEnd sets RangeEnd field to given value.


### GetObservationTimestamp

`func (o *Dashboard) GetObservationTimestamp() string`

GetObservationTimestamp returns the ObservationTimestamp field if non-nil, zero value otherwise.

### GetObservationTimestampOk

`func (o *Dashboard) GetObservationTimestampOk() (*string, bool)`

GetObservationTimestampOk returns a tuple with the ObservationTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservationTimestamp

`func (o *Dashboard) SetObservationTimestamp(v string)`

SetObservationTimestamp sets ObservationTimestamp field to given value.


### GetResourceId

`func (o *Dashboard) GetResourceId() int32`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *Dashboard) GetResourceIdOk() (*int32, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *Dashboard) SetResourceId(v int32)`

SetResourceId sets ResourceId field to given value.


### GetResourceType

`func (o *Dashboard) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *Dashboard) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *Dashboard) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.


### GetData

`func (o *Dashboard) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *Dashboard) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *Dashboard) SetData(v map[string]interface{})`

SetData sets Data field to given value.


### GetLinks

`func (o *Dashboard) GetLinks() ConfigurationLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Dashboard) GetLinksOk() (*ConfigurationLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Dashboard) SetLinks(v ConfigurationLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Dashboard) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


