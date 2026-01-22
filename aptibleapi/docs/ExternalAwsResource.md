# ExternalAwsResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**MetaType** | **string** |  | 
**ExternalAwsAccountId** | **int32** |  | 
**ResourceType** | **string** |  | 
**ResourceArn** | **string** |  | 
**ResourceId** | **string** |  | 
**ResourceName** | **NullableString** |  | 
**Region** | **string** |  | 
**Metadata** | **map[string]interface{}** |  | 
**Tags** | **map[string]interface{}** |  | 
**DiscoveredAt** | **NullableString** |  | 
**LastSyncedAt** | **NullableString** |  | 
**SyncStatus** | **string** |  | 
**CreatedAt** | **string** |  | 
**UpdatedAt** | **string** |  | 
**Links** | Pointer to [**ExternalAwsResourceLinks**](ExternalAwsResourceLinks.md) |  | [optional] 

## Methods

### NewExternalAwsResource

`func NewExternalAwsResource(id int32, metaType string, externalAwsAccountId int32, resourceType string, resourceArn string, resourceId string, resourceName NullableString, region string, metadata map[string]interface{}, tags map[string]interface{}, discoveredAt NullableString, lastSyncedAt NullableString, syncStatus string, createdAt string, updatedAt string, ) *ExternalAwsResource`

NewExternalAwsResource instantiates a new ExternalAwsResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalAwsResourceWithDefaults

`func NewExternalAwsResourceWithDefaults() *ExternalAwsResource`

NewExternalAwsResourceWithDefaults instantiates a new ExternalAwsResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ExternalAwsResource) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExternalAwsResource) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExternalAwsResource) SetId(v int32)`

SetId sets Id field to given value.


### GetMetaType

`func (o *ExternalAwsResource) GetMetaType() string`

GetMetaType returns the MetaType field if non-nil, zero value otherwise.

### GetMetaTypeOk

`func (o *ExternalAwsResource) GetMetaTypeOk() (*string, bool)`

GetMetaTypeOk returns a tuple with the MetaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaType

`func (o *ExternalAwsResource) SetMetaType(v string)`

SetMetaType sets MetaType field to given value.


### GetExternalAwsAccountId

`func (o *ExternalAwsResource) GetExternalAwsAccountId() int32`

GetExternalAwsAccountId returns the ExternalAwsAccountId field if non-nil, zero value otherwise.

### GetExternalAwsAccountIdOk

`func (o *ExternalAwsResource) GetExternalAwsAccountIdOk() (*int32, bool)`

GetExternalAwsAccountIdOk returns a tuple with the ExternalAwsAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalAwsAccountId

`func (o *ExternalAwsResource) SetExternalAwsAccountId(v int32)`

SetExternalAwsAccountId sets ExternalAwsAccountId field to given value.


### GetResourceType

`func (o *ExternalAwsResource) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *ExternalAwsResource) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *ExternalAwsResource) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.


### GetResourceArn

`func (o *ExternalAwsResource) GetResourceArn() string`

GetResourceArn returns the ResourceArn field if non-nil, zero value otherwise.

### GetResourceArnOk

`func (o *ExternalAwsResource) GetResourceArnOk() (*string, bool)`

GetResourceArnOk returns a tuple with the ResourceArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceArn

`func (o *ExternalAwsResource) SetResourceArn(v string)`

SetResourceArn sets ResourceArn field to given value.


### GetResourceId

`func (o *ExternalAwsResource) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *ExternalAwsResource) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *ExternalAwsResource) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.


### GetResourceName

`func (o *ExternalAwsResource) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *ExternalAwsResource) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *ExternalAwsResource) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.


### SetResourceNameNil

`func (o *ExternalAwsResource) SetResourceNameNil(b bool)`

 SetResourceNameNil sets the value for ResourceName to be an explicit nil

### UnsetResourceName
`func (o *ExternalAwsResource) UnsetResourceName()`

UnsetResourceName ensures that no value is present for ResourceName, not even an explicit nil
### GetRegion

`func (o *ExternalAwsResource) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *ExternalAwsResource) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *ExternalAwsResource) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetMetadata

`func (o *ExternalAwsResource) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ExternalAwsResource) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ExternalAwsResource) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.


### GetTags

`func (o *ExternalAwsResource) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ExternalAwsResource) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ExternalAwsResource) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.


### GetDiscoveredAt

`func (o *ExternalAwsResource) GetDiscoveredAt() string`

GetDiscoveredAt returns the DiscoveredAt field if non-nil, zero value otherwise.

### GetDiscoveredAtOk

`func (o *ExternalAwsResource) GetDiscoveredAtOk() (*string, bool)`

GetDiscoveredAtOk returns a tuple with the DiscoveredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveredAt

`func (o *ExternalAwsResource) SetDiscoveredAt(v string)`

SetDiscoveredAt sets DiscoveredAt field to given value.


### SetDiscoveredAtNil

`func (o *ExternalAwsResource) SetDiscoveredAtNil(b bool)`

 SetDiscoveredAtNil sets the value for DiscoveredAt to be an explicit nil

### UnsetDiscoveredAt
`func (o *ExternalAwsResource) UnsetDiscoveredAt()`

UnsetDiscoveredAt ensures that no value is present for DiscoveredAt, not even an explicit nil
### GetLastSyncedAt

`func (o *ExternalAwsResource) GetLastSyncedAt() string`

GetLastSyncedAt returns the LastSyncedAt field if non-nil, zero value otherwise.

### GetLastSyncedAtOk

`func (o *ExternalAwsResource) GetLastSyncedAtOk() (*string, bool)`

GetLastSyncedAtOk returns a tuple with the LastSyncedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSyncedAt

`func (o *ExternalAwsResource) SetLastSyncedAt(v string)`

SetLastSyncedAt sets LastSyncedAt field to given value.


### SetLastSyncedAtNil

`func (o *ExternalAwsResource) SetLastSyncedAtNil(b bool)`

 SetLastSyncedAtNil sets the value for LastSyncedAt to be an explicit nil

### UnsetLastSyncedAt
`func (o *ExternalAwsResource) UnsetLastSyncedAt()`

UnsetLastSyncedAt ensures that no value is present for LastSyncedAt, not even an explicit nil
### GetSyncStatus

`func (o *ExternalAwsResource) GetSyncStatus() string`

GetSyncStatus returns the SyncStatus field if non-nil, zero value otherwise.

### GetSyncStatusOk

`func (o *ExternalAwsResource) GetSyncStatusOk() (*string, bool)`

GetSyncStatusOk returns a tuple with the SyncStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncStatus

`func (o *ExternalAwsResource) SetSyncStatus(v string)`

SetSyncStatus sets SyncStatus field to given value.


### GetCreatedAt

`func (o *ExternalAwsResource) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ExternalAwsResource) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ExternalAwsResource) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *ExternalAwsResource) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ExternalAwsResource) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ExternalAwsResource) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetLinks

`func (o *ExternalAwsResource) GetLinks() ExternalAwsResourceLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ExternalAwsResource) GetLinksOk() (*ExternalAwsResourceLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ExternalAwsResource) SetLinks(v ExternalAwsResourceLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ExternalAwsResource) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


