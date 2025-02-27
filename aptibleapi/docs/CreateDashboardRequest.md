# CreateDashboardRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrganizationId** | **string** |  | 
**ResourceId** | **int32** |  | 
**ResourceType** | **string** |  | 
**Name** | **string** |  | 
**ObservationTimestamp** | Pointer to **string** |  | [optional] 
**RangeBegin** | Pointer to **string** |  | [optional] 
**RangeEnd** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateDashboardRequest

`func NewCreateDashboardRequest(organizationId string, resourceId int32, resourceType string, name string, ) *CreateDashboardRequest`

NewCreateDashboardRequest instantiates a new CreateDashboardRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDashboardRequestWithDefaults

`func NewCreateDashboardRequestWithDefaults() *CreateDashboardRequest`

NewCreateDashboardRequestWithDefaults instantiates a new CreateDashboardRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrganizationId

`func (o *CreateDashboardRequest) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *CreateDashboardRequest) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *CreateDashboardRequest) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.


### GetResourceId

`func (o *CreateDashboardRequest) GetResourceId() int32`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *CreateDashboardRequest) GetResourceIdOk() (*int32, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *CreateDashboardRequest) SetResourceId(v int32)`

SetResourceId sets ResourceId field to given value.


### GetResourceType

`func (o *CreateDashboardRequest) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *CreateDashboardRequest) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *CreateDashboardRequest) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.


### GetName

`func (o *CreateDashboardRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateDashboardRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateDashboardRequest) SetName(v string)`

SetName sets Name field to given value.


### GetObservationTimestamp

`func (o *CreateDashboardRequest) GetObservationTimestamp() string`

GetObservationTimestamp returns the ObservationTimestamp field if non-nil, zero value otherwise.

### GetObservationTimestampOk

`func (o *CreateDashboardRequest) GetObservationTimestampOk() (*string, bool)`

GetObservationTimestampOk returns a tuple with the ObservationTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservationTimestamp

`func (o *CreateDashboardRequest) SetObservationTimestamp(v string)`

SetObservationTimestamp sets ObservationTimestamp field to given value.

### HasObservationTimestamp

`func (o *CreateDashboardRequest) HasObservationTimestamp() bool`

HasObservationTimestamp returns a boolean if a field has been set.

### GetRangeBegin

`func (o *CreateDashboardRequest) GetRangeBegin() string`

GetRangeBegin returns the RangeBegin field if non-nil, zero value otherwise.

### GetRangeBeginOk

`func (o *CreateDashboardRequest) GetRangeBeginOk() (*string, bool)`

GetRangeBeginOk returns a tuple with the RangeBegin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeBegin

`func (o *CreateDashboardRequest) SetRangeBegin(v string)`

SetRangeBegin sets RangeBegin field to given value.

### HasRangeBegin

`func (o *CreateDashboardRequest) HasRangeBegin() bool`

HasRangeBegin returns a boolean if a field has been set.

### GetRangeEnd

`func (o *CreateDashboardRequest) GetRangeEnd() string`

GetRangeEnd returns the RangeEnd field if non-nil, zero value otherwise.

### GetRangeEndOk

`func (o *CreateDashboardRequest) GetRangeEndOk() (*string, bool)`

GetRangeEndOk returns a tuple with the RangeEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeEnd

`func (o *CreateDashboardRequest) SetRangeEnd(v string)`

SetRangeEnd sets RangeEnd field to given value.

### HasRangeEnd

`func (o *CreateDashboardRequest) HasRangeEnd() bool`

HasRangeEnd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


