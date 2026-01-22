# CreateExternalAwsResourceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceType** | **string** |  | 
**ResourceArn** | **string** |  | 
**ResourceId** | **string** |  | 
**ResourceName** | Pointer to **string** |  | [optional] 
**Region** | **string** |  | 
**DiscoveredAt** | Pointer to **string** |  | [optional] 
**LastSyncedAt** | Pointer to **string** |  | [optional] 
**SyncStatus** | **string** |  | 
**Metadata** | **map[string]interface{}** |  | 
**Tags** | **map[string]interface{}** |  | 

## Methods

### NewCreateExternalAwsResourceRequest

`func NewCreateExternalAwsResourceRequest(resourceType string, resourceArn string, resourceId string, region string, syncStatus string, metadata map[string]interface{}, tags map[string]interface{}, ) *CreateExternalAwsResourceRequest`

NewCreateExternalAwsResourceRequest instantiates a new CreateExternalAwsResourceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateExternalAwsResourceRequestWithDefaults

`func NewCreateExternalAwsResourceRequestWithDefaults() *CreateExternalAwsResourceRequest`

NewCreateExternalAwsResourceRequestWithDefaults instantiates a new CreateExternalAwsResourceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceType

`func (o *CreateExternalAwsResourceRequest) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *CreateExternalAwsResourceRequest) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *CreateExternalAwsResourceRequest) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.


### GetResourceArn

`func (o *CreateExternalAwsResourceRequest) GetResourceArn() string`

GetResourceArn returns the ResourceArn field if non-nil, zero value otherwise.

### GetResourceArnOk

`func (o *CreateExternalAwsResourceRequest) GetResourceArnOk() (*string, bool)`

GetResourceArnOk returns a tuple with the ResourceArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceArn

`func (o *CreateExternalAwsResourceRequest) SetResourceArn(v string)`

SetResourceArn sets ResourceArn field to given value.


### GetResourceId

`func (o *CreateExternalAwsResourceRequest) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *CreateExternalAwsResourceRequest) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *CreateExternalAwsResourceRequest) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.


### GetResourceName

`func (o *CreateExternalAwsResourceRequest) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *CreateExternalAwsResourceRequest) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *CreateExternalAwsResourceRequest) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.

### HasResourceName

`func (o *CreateExternalAwsResourceRequest) HasResourceName() bool`

HasResourceName returns a boolean if a field has been set.

### GetRegion

`func (o *CreateExternalAwsResourceRequest) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *CreateExternalAwsResourceRequest) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *CreateExternalAwsResourceRequest) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetDiscoveredAt

`func (o *CreateExternalAwsResourceRequest) GetDiscoveredAt() string`

GetDiscoveredAt returns the DiscoveredAt field if non-nil, zero value otherwise.

### GetDiscoveredAtOk

`func (o *CreateExternalAwsResourceRequest) GetDiscoveredAtOk() (*string, bool)`

GetDiscoveredAtOk returns a tuple with the DiscoveredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveredAt

`func (o *CreateExternalAwsResourceRequest) SetDiscoveredAt(v string)`

SetDiscoveredAt sets DiscoveredAt field to given value.

### HasDiscoveredAt

`func (o *CreateExternalAwsResourceRequest) HasDiscoveredAt() bool`

HasDiscoveredAt returns a boolean if a field has been set.

### GetLastSyncedAt

`func (o *CreateExternalAwsResourceRequest) GetLastSyncedAt() string`

GetLastSyncedAt returns the LastSyncedAt field if non-nil, zero value otherwise.

### GetLastSyncedAtOk

`func (o *CreateExternalAwsResourceRequest) GetLastSyncedAtOk() (*string, bool)`

GetLastSyncedAtOk returns a tuple with the LastSyncedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSyncedAt

`func (o *CreateExternalAwsResourceRequest) SetLastSyncedAt(v string)`

SetLastSyncedAt sets LastSyncedAt field to given value.

### HasLastSyncedAt

`func (o *CreateExternalAwsResourceRequest) HasLastSyncedAt() bool`

HasLastSyncedAt returns a boolean if a field has been set.

### GetSyncStatus

`func (o *CreateExternalAwsResourceRequest) GetSyncStatus() string`

GetSyncStatus returns the SyncStatus field if non-nil, zero value otherwise.

### GetSyncStatusOk

`func (o *CreateExternalAwsResourceRequest) GetSyncStatusOk() (*string, bool)`

GetSyncStatusOk returns a tuple with the SyncStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncStatus

`func (o *CreateExternalAwsResourceRequest) SetSyncStatus(v string)`

SetSyncStatus sets SyncStatus field to given value.


### GetMetadata

`func (o *CreateExternalAwsResourceRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreateExternalAwsResourceRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreateExternalAwsResourceRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.


### GetTags

`func (o *CreateExternalAwsResourceRequest) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateExternalAwsResourceRequest) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateExternalAwsResourceRequest) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


