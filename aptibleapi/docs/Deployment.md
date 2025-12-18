# Deployment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**MetaType** | **string** |  | 
**Status** | **string** |  | 
**OperationId** | **int32** |  | 
**AppId** | **int32** |  | 
**ConfigurationId** | **NullableInt32** |  | 
**ImageId** | **NullableInt32** |  | 
**SourceId** | **NullableInt32** |  | 
**SettingId** | **NullableInt32** |  | 
**DockerImage** | **NullableString** |  | 
**DockerRepositoryUrl** | **NullableString** |  | 
**GitRepositoryUrl** | **NullableString** |  | 
**GitRef** | **NullableString** |  | 
**GitCommitSha** | **NullableString** |  | 
**GitCommitTimestamp** | **NullableString** |  | 
**GitCommitUrl** | **NullableString** |  | 
**GitCommitMessage** | **NullableString** |  | 
**CreatedAt** | **string** |  | 
**UpdatedAt** | **string** |  | 
**DeletedAt** | **NullableString** |  | 
**Links** | Pointer to [**DeploymentLinks**](DeploymentLinks.md) |  | [optional] 

## Methods

### NewDeployment

`func NewDeployment(id int32, metaType string, status string, operationId int32, appId int32, configurationId NullableInt32, imageId NullableInt32, sourceId NullableInt32, settingId NullableInt32, dockerImage NullableString, dockerRepositoryUrl NullableString, gitRepositoryUrl NullableString, gitRef NullableString, gitCommitSha NullableString, gitCommitTimestamp NullableString, gitCommitUrl NullableString, gitCommitMessage NullableString, createdAt string, updatedAt string, deletedAt NullableString, ) *Deployment`

NewDeployment instantiates a new Deployment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentWithDefaults

`func NewDeploymentWithDefaults() *Deployment`

NewDeploymentWithDefaults instantiates a new Deployment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Deployment) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Deployment) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Deployment) SetId(v int32)`

SetId sets Id field to given value.


### GetMetaType

`func (o *Deployment) GetMetaType() string`

GetMetaType returns the MetaType field if non-nil, zero value otherwise.

### GetMetaTypeOk

`func (o *Deployment) GetMetaTypeOk() (*string, bool)`

GetMetaTypeOk returns a tuple with the MetaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaType

`func (o *Deployment) SetMetaType(v string)`

SetMetaType sets MetaType field to given value.


### GetStatus

`func (o *Deployment) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Deployment) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Deployment) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetOperationId

`func (o *Deployment) GetOperationId() int32`

GetOperationId returns the OperationId field if non-nil, zero value otherwise.

### GetOperationIdOk

`func (o *Deployment) GetOperationIdOk() (*int32, bool)`

GetOperationIdOk returns a tuple with the OperationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationId

`func (o *Deployment) SetOperationId(v int32)`

SetOperationId sets OperationId field to given value.


### GetAppId

`func (o *Deployment) GetAppId() int32`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *Deployment) GetAppIdOk() (*int32, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *Deployment) SetAppId(v int32)`

SetAppId sets AppId field to given value.


### GetConfigurationId

`func (o *Deployment) GetConfigurationId() int32`

GetConfigurationId returns the ConfigurationId field if non-nil, zero value otherwise.

### GetConfigurationIdOk

`func (o *Deployment) GetConfigurationIdOk() (*int32, bool)`

GetConfigurationIdOk returns a tuple with the ConfigurationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationId

`func (o *Deployment) SetConfigurationId(v int32)`

SetConfigurationId sets ConfigurationId field to given value.


### SetConfigurationIdNil

`func (o *Deployment) SetConfigurationIdNil(b bool)`

 SetConfigurationIdNil sets the value for ConfigurationId to be an explicit nil

### UnsetConfigurationId
`func (o *Deployment) UnsetConfigurationId()`

UnsetConfigurationId ensures that no value is present for ConfigurationId, not even an explicit nil
### GetImageId

`func (o *Deployment) GetImageId() int32`

GetImageId returns the ImageId field if non-nil, zero value otherwise.

### GetImageIdOk

`func (o *Deployment) GetImageIdOk() (*int32, bool)`

GetImageIdOk returns a tuple with the ImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageId

`func (o *Deployment) SetImageId(v int32)`

SetImageId sets ImageId field to given value.


### SetImageIdNil

`func (o *Deployment) SetImageIdNil(b bool)`

 SetImageIdNil sets the value for ImageId to be an explicit nil

### UnsetImageId
`func (o *Deployment) UnsetImageId()`

UnsetImageId ensures that no value is present for ImageId, not even an explicit nil
### GetSourceId

`func (o *Deployment) GetSourceId() int32`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *Deployment) GetSourceIdOk() (*int32, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *Deployment) SetSourceId(v int32)`

SetSourceId sets SourceId field to given value.


### SetSourceIdNil

`func (o *Deployment) SetSourceIdNil(b bool)`

 SetSourceIdNil sets the value for SourceId to be an explicit nil

### UnsetSourceId
`func (o *Deployment) UnsetSourceId()`

UnsetSourceId ensures that no value is present for SourceId, not even an explicit nil
### GetSettingId

`func (o *Deployment) GetSettingId() int32`

GetSettingId returns the SettingId field if non-nil, zero value otherwise.

### GetSettingIdOk

`func (o *Deployment) GetSettingIdOk() (*int32, bool)`

GetSettingIdOk returns a tuple with the SettingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingId

`func (o *Deployment) SetSettingId(v int32)`

SetSettingId sets SettingId field to given value.


### SetSettingIdNil

`func (o *Deployment) SetSettingIdNil(b bool)`

 SetSettingIdNil sets the value for SettingId to be an explicit nil

### UnsetSettingId
`func (o *Deployment) UnsetSettingId()`

UnsetSettingId ensures that no value is present for SettingId, not even an explicit nil
### GetDockerImage

`func (o *Deployment) GetDockerImage() string`

GetDockerImage returns the DockerImage field if non-nil, zero value otherwise.

### GetDockerImageOk

`func (o *Deployment) GetDockerImageOk() (*string, bool)`

GetDockerImageOk returns a tuple with the DockerImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerImage

`func (o *Deployment) SetDockerImage(v string)`

SetDockerImage sets DockerImage field to given value.


### SetDockerImageNil

`func (o *Deployment) SetDockerImageNil(b bool)`

 SetDockerImageNil sets the value for DockerImage to be an explicit nil

### UnsetDockerImage
`func (o *Deployment) UnsetDockerImage()`

UnsetDockerImage ensures that no value is present for DockerImage, not even an explicit nil
### GetDockerRepositoryUrl

`func (o *Deployment) GetDockerRepositoryUrl() string`

GetDockerRepositoryUrl returns the DockerRepositoryUrl field if non-nil, zero value otherwise.

### GetDockerRepositoryUrlOk

`func (o *Deployment) GetDockerRepositoryUrlOk() (*string, bool)`

GetDockerRepositoryUrlOk returns a tuple with the DockerRepositoryUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerRepositoryUrl

`func (o *Deployment) SetDockerRepositoryUrl(v string)`

SetDockerRepositoryUrl sets DockerRepositoryUrl field to given value.


### SetDockerRepositoryUrlNil

`func (o *Deployment) SetDockerRepositoryUrlNil(b bool)`

 SetDockerRepositoryUrlNil sets the value for DockerRepositoryUrl to be an explicit nil

### UnsetDockerRepositoryUrl
`func (o *Deployment) UnsetDockerRepositoryUrl()`

UnsetDockerRepositoryUrl ensures that no value is present for DockerRepositoryUrl, not even an explicit nil
### GetGitRepositoryUrl

`func (o *Deployment) GetGitRepositoryUrl() string`

GetGitRepositoryUrl returns the GitRepositoryUrl field if non-nil, zero value otherwise.

### GetGitRepositoryUrlOk

`func (o *Deployment) GetGitRepositoryUrlOk() (*string, bool)`

GetGitRepositoryUrlOk returns a tuple with the GitRepositoryUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitRepositoryUrl

`func (o *Deployment) SetGitRepositoryUrl(v string)`

SetGitRepositoryUrl sets GitRepositoryUrl field to given value.


### SetGitRepositoryUrlNil

`func (o *Deployment) SetGitRepositoryUrlNil(b bool)`

 SetGitRepositoryUrlNil sets the value for GitRepositoryUrl to be an explicit nil

### UnsetGitRepositoryUrl
`func (o *Deployment) UnsetGitRepositoryUrl()`

UnsetGitRepositoryUrl ensures that no value is present for GitRepositoryUrl, not even an explicit nil
### GetGitRef

`func (o *Deployment) GetGitRef() string`

GetGitRef returns the GitRef field if non-nil, zero value otherwise.

### GetGitRefOk

`func (o *Deployment) GetGitRefOk() (*string, bool)`

GetGitRefOk returns a tuple with the GitRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitRef

`func (o *Deployment) SetGitRef(v string)`

SetGitRef sets GitRef field to given value.


### SetGitRefNil

`func (o *Deployment) SetGitRefNil(b bool)`

 SetGitRefNil sets the value for GitRef to be an explicit nil

### UnsetGitRef
`func (o *Deployment) UnsetGitRef()`

UnsetGitRef ensures that no value is present for GitRef, not even an explicit nil
### GetGitCommitSha

`func (o *Deployment) GetGitCommitSha() string`

GetGitCommitSha returns the GitCommitSha field if non-nil, zero value otherwise.

### GetGitCommitShaOk

`func (o *Deployment) GetGitCommitShaOk() (*string, bool)`

GetGitCommitShaOk returns a tuple with the GitCommitSha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitCommitSha

`func (o *Deployment) SetGitCommitSha(v string)`

SetGitCommitSha sets GitCommitSha field to given value.


### SetGitCommitShaNil

`func (o *Deployment) SetGitCommitShaNil(b bool)`

 SetGitCommitShaNil sets the value for GitCommitSha to be an explicit nil

### UnsetGitCommitSha
`func (o *Deployment) UnsetGitCommitSha()`

UnsetGitCommitSha ensures that no value is present for GitCommitSha, not even an explicit nil
### GetGitCommitTimestamp

`func (o *Deployment) GetGitCommitTimestamp() string`

GetGitCommitTimestamp returns the GitCommitTimestamp field if non-nil, zero value otherwise.

### GetGitCommitTimestampOk

`func (o *Deployment) GetGitCommitTimestampOk() (*string, bool)`

GetGitCommitTimestampOk returns a tuple with the GitCommitTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitCommitTimestamp

`func (o *Deployment) SetGitCommitTimestamp(v string)`

SetGitCommitTimestamp sets GitCommitTimestamp field to given value.


### SetGitCommitTimestampNil

`func (o *Deployment) SetGitCommitTimestampNil(b bool)`

 SetGitCommitTimestampNil sets the value for GitCommitTimestamp to be an explicit nil

### UnsetGitCommitTimestamp
`func (o *Deployment) UnsetGitCommitTimestamp()`

UnsetGitCommitTimestamp ensures that no value is present for GitCommitTimestamp, not even an explicit nil
### GetGitCommitUrl

`func (o *Deployment) GetGitCommitUrl() string`

GetGitCommitUrl returns the GitCommitUrl field if non-nil, zero value otherwise.

### GetGitCommitUrlOk

`func (o *Deployment) GetGitCommitUrlOk() (*string, bool)`

GetGitCommitUrlOk returns a tuple with the GitCommitUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitCommitUrl

`func (o *Deployment) SetGitCommitUrl(v string)`

SetGitCommitUrl sets GitCommitUrl field to given value.


### SetGitCommitUrlNil

`func (o *Deployment) SetGitCommitUrlNil(b bool)`

 SetGitCommitUrlNil sets the value for GitCommitUrl to be an explicit nil

### UnsetGitCommitUrl
`func (o *Deployment) UnsetGitCommitUrl()`

UnsetGitCommitUrl ensures that no value is present for GitCommitUrl, not even an explicit nil
### GetGitCommitMessage

`func (o *Deployment) GetGitCommitMessage() string`

GetGitCommitMessage returns the GitCommitMessage field if non-nil, zero value otherwise.

### GetGitCommitMessageOk

`func (o *Deployment) GetGitCommitMessageOk() (*string, bool)`

GetGitCommitMessageOk returns a tuple with the GitCommitMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitCommitMessage

`func (o *Deployment) SetGitCommitMessage(v string)`

SetGitCommitMessage sets GitCommitMessage field to given value.


### SetGitCommitMessageNil

`func (o *Deployment) SetGitCommitMessageNil(b bool)`

 SetGitCommitMessageNil sets the value for GitCommitMessage to be an explicit nil

### UnsetGitCommitMessage
`func (o *Deployment) UnsetGitCommitMessage()`

UnsetGitCommitMessage ensures that no value is present for GitCommitMessage, not even an explicit nil
### GetCreatedAt

`func (o *Deployment) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Deployment) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Deployment) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Deployment) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Deployment) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Deployment) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetDeletedAt

`func (o *Deployment) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *Deployment) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *Deployment) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.


### SetDeletedAtNil

`func (o *Deployment) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *Deployment) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
### GetLinks

`func (o *Deployment) GetLinks() DeploymentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Deployment) GetLinksOk() (*DeploymentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Deployment) SetLinks(v DeploymentLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Deployment) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


