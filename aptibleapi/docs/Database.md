# Database

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Handle** | **string** |  | 
**Type** | **NullableString** |  | 
**Passphrase** | **NullableString** |  | 
**ConnectionUrl** | **NullableString** |  | 
**Provisioned** | **bool** |  | 
**MetaType** | **string** |  | 
**CreatedAt** | **string** |  | 
**UpdatedAt** | **string** |  | 
**Status** | **string** |  | 
**DockerRepo** | **NullableString** |  | 
**PortMapping** | **[][]int32** |  | 
**InitialDiskSize** | **NullableInt32** |  | 
**InitialContainerSize** | **NullableInt32** |  | 
**CurrentKmsArn** | **NullableString** |  | 
**EnableBackups** | **bool** |  | 
**PitrBucket** | **NullableString** |  | 
**OldestPitrTime** | **NullableString** |  | 
**EnablePitr** | **bool** |  | 
**Embedded** | [**DatabaseEmbedded**](DatabaseEmbedded.md) |  | 
**Links** | Pointer to [**DatabaseLinks**](DatabaseLinks.md) |  | [optional] 

## Methods

### NewDatabase

`func NewDatabase(id int32, handle string, type_ NullableString, passphrase NullableString, connectionUrl NullableString, provisioned bool, metaType string, createdAt string, updatedAt string, status string, dockerRepo NullableString, portMapping [][]int32, initialDiskSize NullableInt32, initialContainerSize NullableInt32, currentKmsArn NullableString, enableBackups bool, pitrBucket NullableString, oldestPitrTime NullableString, enablePitr bool, embedded DatabaseEmbedded, ) *Database`

NewDatabase instantiates a new Database object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabaseWithDefaults

`func NewDatabaseWithDefaults() *Database`

NewDatabaseWithDefaults instantiates a new Database object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Database) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Database) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Database) SetId(v int32)`

SetId sets Id field to given value.


### GetHandle

`func (o *Database) GetHandle() string`

GetHandle returns the Handle field if non-nil, zero value otherwise.

### GetHandleOk

`func (o *Database) GetHandleOk() (*string, bool)`

GetHandleOk returns a tuple with the Handle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandle

`func (o *Database) SetHandle(v string)`

SetHandle sets Handle field to given value.


### GetType

`func (o *Database) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Database) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Database) SetType(v string)`

SetType sets Type field to given value.


### SetTypeNil

`func (o *Database) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *Database) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetPassphrase

`func (o *Database) GetPassphrase() string`

GetPassphrase returns the Passphrase field if non-nil, zero value otherwise.

### GetPassphraseOk

`func (o *Database) GetPassphraseOk() (*string, bool)`

GetPassphraseOk returns a tuple with the Passphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassphrase

`func (o *Database) SetPassphrase(v string)`

SetPassphrase sets Passphrase field to given value.


### SetPassphraseNil

`func (o *Database) SetPassphraseNil(b bool)`

 SetPassphraseNil sets the value for Passphrase to be an explicit nil

### UnsetPassphrase
`func (o *Database) UnsetPassphrase()`

UnsetPassphrase ensures that no value is present for Passphrase, not even an explicit nil
### GetConnectionUrl

`func (o *Database) GetConnectionUrl() string`

GetConnectionUrl returns the ConnectionUrl field if non-nil, zero value otherwise.

### GetConnectionUrlOk

`func (o *Database) GetConnectionUrlOk() (*string, bool)`

GetConnectionUrlOk returns a tuple with the ConnectionUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionUrl

`func (o *Database) SetConnectionUrl(v string)`

SetConnectionUrl sets ConnectionUrl field to given value.


### SetConnectionUrlNil

`func (o *Database) SetConnectionUrlNil(b bool)`

 SetConnectionUrlNil sets the value for ConnectionUrl to be an explicit nil

### UnsetConnectionUrl
`func (o *Database) UnsetConnectionUrl()`

UnsetConnectionUrl ensures that no value is present for ConnectionUrl, not even an explicit nil
### GetProvisioned

`func (o *Database) GetProvisioned() bool`

GetProvisioned returns the Provisioned field if non-nil, zero value otherwise.

### GetProvisionedOk

`func (o *Database) GetProvisionedOk() (*bool, bool)`

GetProvisionedOk returns a tuple with the Provisioned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioned

`func (o *Database) SetProvisioned(v bool)`

SetProvisioned sets Provisioned field to given value.


### GetMetaType

`func (o *Database) GetMetaType() string`

GetMetaType returns the MetaType field if non-nil, zero value otherwise.

### GetMetaTypeOk

`func (o *Database) GetMetaTypeOk() (*string, bool)`

GetMetaTypeOk returns a tuple with the MetaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaType

`func (o *Database) SetMetaType(v string)`

SetMetaType sets MetaType field to given value.


### GetCreatedAt

`func (o *Database) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Database) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Database) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Database) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Database) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Database) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetStatus

`func (o *Database) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Database) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Database) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetDockerRepo

`func (o *Database) GetDockerRepo() string`

GetDockerRepo returns the DockerRepo field if non-nil, zero value otherwise.

### GetDockerRepoOk

`func (o *Database) GetDockerRepoOk() (*string, bool)`

GetDockerRepoOk returns a tuple with the DockerRepo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerRepo

`func (o *Database) SetDockerRepo(v string)`

SetDockerRepo sets DockerRepo field to given value.


### SetDockerRepoNil

`func (o *Database) SetDockerRepoNil(b bool)`

 SetDockerRepoNil sets the value for DockerRepo to be an explicit nil

### UnsetDockerRepo
`func (o *Database) UnsetDockerRepo()`

UnsetDockerRepo ensures that no value is present for DockerRepo, not even an explicit nil
### GetPortMapping

`func (o *Database) GetPortMapping() [][]int32`

GetPortMapping returns the PortMapping field if non-nil, zero value otherwise.

### GetPortMappingOk

`func (o *Database) GetPortMappingOk() (*[][]int32, bool)`

GetPortMappingOk returns a tuple with the PortMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortMapping

`func (o *Database) SetPortMapping(v [][]int32)`

SetPortMapping sets PortMapping field to given value.


### GetInitialDiskSize

`func (o *Database) GetInitialDiskSize() int32`

GetInitialDiskSize returns the InitialDiskSize field if non-nil, zero value otherwise.

### GetInitialDiskSizeOk

`func (o *Database) GetInitialDiskSizeOk() (*int32, bool)`

GetInitialDiskSizeOk returns a tuple with the InitialDiskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialDiskSize

`func (o *Database) SetInitialDiskSize(v int32)`

SetInitialDiskSize sets InitialDiskSize field to given value.


### SetInitialDiskSizeNil

`func (o *Database) SetInitialDiskSizeNil(b bool)`

 SetInitialDiskSizeNil sets the value for InitialDiskSize to be an explicit nil

### UnsetInitialDiskSize
`func (o *Database) UnsetInitialDiskSize()`

UnsetInitialDiskSize ensures that no value is present for InitialDiskSize, not even an explicit nil
### GetInitialContainerSize

`func (o *Database) GetInitialContainerSize() int32`

GetInitialContainerSize returns the InitialContainerSize field if non-nil, zero value otherwise.

### GetInitialContainerSizeOk

`func (o *Database) GetInitialContainerSizeOk() (*int32, bool)`

GetInitialContainerSizeOk returns a tuple with the InitialContainerSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialContainerSize

`func (o *Database) SetInitialContainerSize(v int32)`

SetInitialContainerSize sets InitialContainerSize field to given value.


### SetInitialContainerSizeNil

`func (o *Database) SetInitialContainerSizeNil(b bool)`

 SetInitialContainerSizeNil sets the value for InitialContainerSize to be an explicit nil

### UnsetInitialContainerSize
`func (o *Database) UnsetInitialContainerSize()`

UnsetInitialContainerSize ensures that no value is present for InitialContainerSize, not even an explicit nil
### GetCurrentKmsArn

`func (o *Database) GetCurrentKmsArn() string`

GetCurrentKmsArn returns the CurrentKmsArn field if non-nil, zero value otherwise.

### GetCurrentKmsArnOk

`func (o *Database) GetCurrentKmsArnOk() (*string, bool)`

GetCurrentKmsArnOk returns a tuple with the CurrentKmsArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentKmsArn

`func (o *Database) SetCurrentKmsArn(v string)`

SetCurrentKmsArn sets CurrentKmsArn field to given value.


### SetCurrentKmsArnNil

`func (o *Database) SetCurrentKmsArnNil(b bool)`

 SetCurrentKmsArnNil sets the value for CurrentKmsArn to be an explicit nil

### UnsetCurrentKmsArn
`func (o *Database) UnsetCurrentKmsArn()`

UnsetCurrentKmsArn ensures that no value is present for CurrentKmsArn, not even an explicit nil
### GetEnableBackups

`func (o *Database) GetEnableBackups() bool`

GetEnableBackups returns the EnableBackups field if non-nil, zero value otherwise.

### GetEnableBackupsOk

`func (o *Database) GetEnableBackupsOk() (*bool, bool)`

GetEnableBackupsOk returns a tuple with the EnableBackups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableBackups

`func (o *Database) SetEnableBackups(v bool)`

SetEnableBackups sets EnableBackups field to given value.


### GetPitrBucket

`func (o *Database) GetPitrBucket() string`

GetPitrBucket returns the PitrBucket field if non-nil, zero value otherwise.

### GetPitrBucketOk

`func (o *Database) GetPitrBucketOk() (*string, bool)`

GetPitrBucketOk returns a tuple with the PitrBucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPitrBucket

`func (o *Database) SetPitrBucket(v string)`

SetPitrBucket sets PitrBucket field to given value.


### SetPitrBucketNil

`func (o *Database) SetPitrBucketNil(b bool)`

 SetPitrBucketNil sets the value for PitrBucket to be an explicit nil

### UnsetPitrBucket
`func (o *Database) UnsetPitrBucket()`

UnsetPitrBucket ensures that no value is present for PitrBucket, not even an explicit nil
### GetOldestPitrTime

`func (o *Database) GetOldestPitrTime() string`

GetOldestPitrTime returns the OldestPitrTime field if non-nil, zero value otherwise.

### GetOldestPitrTimeOk

`func (o *Database) GetOldestPitrTimeOk() (*string, bool)`

GetOldestPitrTimeOk returns a tuple with the OldestPitrTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldestPitrTime

`func (o *Database) SetOldestPitrTime(v string)`

SetOldestPitrTime sets OldestPitrTime field to given value.


### SetOldestPitrTimeNil

`func (o *Database) SetOldestPitrTimeNil(b bool)`

 SetOldestPitrTimeNil sets the value for OldestPitrTime to be an explicit nil

### UnsetOldestPitrTime
`func (o *Database) UnsetOldestPitrTime()`

UnsetOldestPitrTime ensures that no value is present for OldestPitrTime, not even an explicit nil
### GetEnablePitr

`func (o *Database) GetEnablePitr() bool`

GetEnablePitr returns the EnablePitr field if non-nil, zero value otherwise.

### GetEnablePitrOk

`func (o *Database) GetEnablePitrOk() (*bool, bool)`

GetEnablePitrOk returns a tuple with the EnablePitr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnablePitr

`func (o *Database) SetEnablePitr(v bool)`

SetEnablePitr sets EnablePitr field to given value.


### GetEmbedded

`func (o *Database) GetEmbedded() DatabaseEmbedded`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *Database) GetEmbeddedOk() (*DatabaseEmbedded, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *Database) SetEmbedded(v DatabaseEmbedded)`

SetEmbedded sets Embedded field to given value.


### GetLinks

`func (o *Database) GetLinks() DatabaseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Database) GetLinksOk() (*DatabaseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Database) SetLinks(v DatabaseLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Database) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


