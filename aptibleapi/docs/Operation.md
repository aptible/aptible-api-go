# Operation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**MetaType** | **string** |  | 
**Type** | **string** |  | 
**Status** | **string** |  | 
**Cancelled** | **bool** |  | 
**Aborted** | **bool** |  | 
**GitRef** | **NullableString** |  | 
**DockerRef** | **NullableString** |  | 
**Env** | **map[string]interface{}** |  | 
**ContainerSize** | **NullableInt32** |  | 
**ContainerCount** | **NullableInt32** |  | 
**DiskSize** | **int32** |  | 
**Command** | **NullableString** |  | 
**Handle** | **NullableString** |  | 
**CreatedAt** | **string** |  | 
**UpdatedAt** | **string** |  | 
**Certificate** | **NullableString** |  | 
**PrivateKey** | **NullableString** |  | 
**UserName** | **string** |  | 
**UserEmail** | **string** |  | 
**DestinationRegion** | **NullableString** |  | 
**Interactive** | **NullableBool** |  | 
**InstanceProfile** | **NullableString** |  | 
**MountPoint** | **NullableString** |  | 
**Daily** | **NullableInt32** |  | 
**Monthly** | **NullableInt32** |  | 
**Yearly** | **NullableInt32** |  | 
**PitrDays** | **NullableInt32** |  | 
**MakeCopy** | **NullableBool** |  | 
**KeepFinal** | **NullableBool** |  | 
**Links** | Pointer to [**OperationLinks**](OperationLinks.md) |  | [optional] 

## Methods

### NewOperation

`func NewOperation(id int32, metaType string, type_ string, status string, cancelled bool, aborted bool, gitRef NullableString, dockerRef NullableString, env map[string]interface{}, containerSize NullableInt32, containerCount NullableInt32, diskSize int32, command NullableString, handle NullableString, createdAt string, updatedAt string, certificate NullableString, privateKey NullableString, userName string, userEmail string, destinationRegion NullableString, interactive NullableBool, instanceProfile NullableString, mountPoint NullableString, daily NullableInt32, monthly NullableInt32, yearly NullableInt32, pitrDays NullableInt32, makeCopy NullableBool, keepFinal NullableBool, ) *Operation`

NewOperation instantiates a new Operation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOperationWithDefaults

`func NewOperationWithDefaults() *Operation`

NewOperationWithDefaults instantiates a new Operation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Operation) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Operation) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Operation) SetId(v int32)`

SetId sets Id field to given value.


### GetMetaType

`func (o *Operation) GetMetaType() string`

GetMetaType returns the MetaType field if non-nil, zero value otherwise.

### GetMetaTypeOk

`func (o *Operation) GetMetaTypeOk() (*string, bool)`

GetMetaTypeOk returns a tuple with the MetaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaType

`func (o *Operation) SetMetaType(v string)`

SetMetaType sets MetaType field to given value.


### GetType

`func (o *Operation) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Operation) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Operation) SetType(v string)`

SetType sets Type field to given value.


### GetStatus

`func (o *Operation) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Operation) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Operation) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetCancelled

`func (o *Operation) GetCancelled() bool`

GetCancelled returns the Cancelled field if non-nil, zero value otherwise.

### GetCancelledOk

`func (o *Operation) GetCancelledOk() (*bool, bool)`

GetCancelledOk returns a tuple with the Cancelled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelled

`func (o *Operation) SetCancelled(v bool)`

SetCancelled sets Cancelled field to given value.


### GetAborted

`func (o *Operation) GetAborted() bool`

GetAborted returns the Aborted field if non-nil, zero value otherwise.

### GetAbortedOk

`func (o *Operation) GetAbortedOk() (*bool, bool)`

GetAbortedOk returns a tuple with the Aborted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAborted

`func (o *Operation) SetAborted(v bool)`

SetAborted sets Aborted field to given value.


### GetGitRef

`func (o *Operation) GetGitRef() string`

GetGitRef returns the GitRef field if non-nil, zero value otherwise.

### GetGitRefOk

`func (o *Operation) GetGitRefOk() (*string, bool)`

GetGitRefOk returns a tuple with the GitRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitRef

`func (o *Operation) SetGitRef(v string)`

SetGitRef sets GitRef field to given value.


### SetGitRefNil

`func (o *Operation) SetGitRefNil(b bool)`

 SetGitRefNil sets the value for GitRef to be an explicit nil

### UnsetGitRef
`func (o *Operation) UnsetGitRef()`

UnsetGitRef ensures that no value is present for GitRef, not even an explicit nil
### GetDockerRef

`func (o *Operation) GetDockerRef() string`

GetDockerRef returns the DockerRef field if non-nil, zero value otherwise.

### GetDockerRefOk

`func (o *Operation) GetDockerRefOk() (*string, bool)`

GetDockerRefOk returns a tuple with the DockerRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerRef

`func (o *Operation) SetDockerRef(v string)`

SetDockerRef sets DockerRef field to given value.


### SetDockerRefNil

`func (o *Operation) SetDockerRefNil(b bool)`

 SetDockerRefNil sets the value for DockerRef to be an explicit nil

### UnsetDockerRef
`func (o *Operation) UnsetDockerRef()`

UnsetDockerRef ensures that no value is present for DockerRef, not even an explicit nil
### GetEnv

`func (o *Operation) GetEnv() map[string]interface{}`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *Operation) GetEnvOk() (*map[string]interface{}, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *Operation) SetEnv(v map[string]interface{})`

SetEnv sets Env field to given value.


### SetEnvNil

`func (o *Operation) SetEnvNil(b bool)`

 SetEnvNil sets the value for Env to be an explicit nil

### UnsetEnv
`func (o *Operation) UnsetEnv()`

UnsetEnv ensures that no value is present for Env, not even an explicit nil
### GetContainerSize

`func (o *Operation) GetContainerSize() int32`

GetContainerSize returns the ContainerSize field if non-nil, zero value otherwise.

### GetContainerSizeOk

`func (o *Operation) GetContainerSizeOk() (*int32, bool)`

GetContainerSizeOk returns a tuple with the ContainerSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerSize

`func (o *Operation) SetContainerSize(v int32)`

SetContainerSize sets ContainerSize field to given value.


### SetContainerSizeNil

`func (o *Operation) SetContainerSizeNil(b bool)`

 SetContainerSizeNil sets the value for ContainerSize to be an explicit nil

### UnsetContainerSize
`func (o *Operation) UnsetContainerSize()`

UnsetContainerSize ensures that no value is present for ContainerSize, not even an explicit nil
### GetContainerCount

`func (o *Operation) GetContainerCount() int32`

GetContainerCount returns the ContainerCount field if non-nil, zero value otherwise.

### GetContainerCountOk

`func (o *Operation) GetContainerCountOk() (*int32, bool)`

GetContainerCountOk returns a tuple with the ContainerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerCount

`func (o *Operation) SetContainerCount(v int32)`

SetContainerCount sets ContainerCount field to given value.


### SetContainerCountNil

`func (o *Operation) SetContainerCountNil(b bool)`

 SetContainerCountNil sets the value for ContainerCount to be an explicit nil

### UnsetContainerCount
`func (o *Operation) UnsetContainerCount()`

UnsetContainerCount ensures that no value is present for ContainerCount, not even an explicit nil
### GetDiskSize

`func (o *Operation) GetDiskSize() int32`

GetDiskSize returns the DiskSize field if non-nil, zero value otherwise.

### GetDiskSizeOk

`func (o *Operation) GetDiskSizeOk() (*int32, bool)`

GetDiskSizeOk returns a tuple with the DiskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSize

`func (o *Operation) SetDiskSize(v int32)`

SetDiskSize sets DiskSize field to given value.


### GetCommand

`func (o *Operation) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *Operation) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *Operation) SetCommand(v string)`

SetCommand sets Command field to given value.


### SetCommandNil

`func (o *Operation) SetCommandNil(b bool)`

 SetCommandNil sets the value for Command to be an explicit nil

### UnsetCommand
`func (o *Operation) UnsetCommand()`

UnsetCommand ensures that no value is present for Command, not even an explicit nil
### GetHandle

`func (o *Operation) GetHandle() string`

GetHandle returns the Handle field if non-nil, zero value otherwise.

### GetHandleOk

`func (o *Operation) GetHandleOk() (*string, bool)`

GetHandleOk returns a tuple with the Handle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandle

`func (o *Operation) SetHandle(v string)`

SetHandle sets Handle field to given value.


### SetHandleNil

`func (o *Operation) SetHandleNil(b bool)`

 SetHandleNil sets the value for Handle to be an explicit nil

### UnsetHandle
`func (o *Operation) UnsetHandle()`

UnsetHandle ensures that no value is present for Handle, not even an explicit nil
### GetCreatedAt

`func (o *Operation) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Operation) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Operation) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Operation) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Operation) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Operation) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetCertificate

`func (o *Operation) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *Operation) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *Operation) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.


### SetCertificateNil

`func (o *Operation) SetCertificateNil(b bool)`

 SetCertificateNil sets the value for Certificate to be an explicit nil

### UnsetCertificate
`func (o *Operation) UnsetCertificate()`

UnsetCertificate ensures that no value is present for Certificate, not even an explicit nil
### GetPrivateKey

`func (o *Operation) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *Operation) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *Operation) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.


### SetPrivateKeyNil

`func (o *Operation) SetPrivateKeyNil(b bool)`

 SetPrivateKeyNil sets the value for PrivateKey to be an explicit nil

### UnsetPrivateKey
`func (o *Operation) UnsetPrivateKey()`

UnsetPrivateKey ensures that no value is present for PrivateKey, not even an explicit nil
### GetUserName

`func (o *Operation) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *Operation) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *Operation) SetUserName(v string)`

SetUserName sets UserName field to given value.


### GetUserEmail

`func (o *Operation) GetUserEmail() string`

GetUserEmail returns the UserEmail field if non-nil, zero value otherwise.

### GetUserEmailOk

`func (o *Operation) GetUserEmailOk() (*string, bool)`

GetUserEmailOk returns a tuple with the UserEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserEmail

`func (o *Operation) SetUserEmail(v string)`

SetUserEmail sets UserEmail field to given value.


### GetDestinationRegion

`func (o *Operation) GetDestinationRegion() string`

GetDestinationRegion returns the DestinationRegion field if non-nil, zero value otherwise.

### GetDestinationRegionOk

`func (o *Operation) GetDestinationRegionOk() (*string, bool)`

GetDestinationRegionOk returns a tuple with the DestinationRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationRegion

`func (o *Operation) SetDestinationRegion(v string)`

SetDestinationRegion sets DestinationRegion field to given value.


### SetDestinationRegionNil

`func (o *Operation) SetDestinationRegionNil(b bool)`

 SetDestinationRegionNil sets the value for DestinationRegion to be an explicit nil

### UnsetDestinationRegion
`func (o *Operation) UnsetDestinationRegion()`

UnsetDestinationRegion ensures that no value is present for DestinationRegion, not even an explicit nil
### GetInteractive

`func (o *Operation) GetInteractive() bool`

GetInteractive returns the Interactive field if non-nil, zero value otherwise.

### GetInteractiveOk

`func (o *Operation) GetInteractiveOk() (*bool, bool)`

GetInteractiveOk returns a tuple with the Interactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInteractive

`func (o *Operation) SetInteractive(v bool)`

SetInteractive sets Interactive field to given value.


### SetInteractiveNil

`func (o *Operation) SetInteractiveNil(b bool)`

 SetInteractiveNil sets the value for Interactive to be an explicit nil

### UnsetInteractive
`func (o *Operation) UnsetInteractive()`

UnsetInteractive ensures that no value is present for Interactive, not even an explicit nil
### GetInstanceProfile

`func (o *Operation) GetInstanceProfile() string`

GetInstanceProfile returns the InstanceProfile field if non-nil, zero value otherwise.

### GetInstanceProfileOk

`func (o *Operation) GetInstanceProfileOk() (*string, bool)`

GetInstanceProfileOk returns a tuple with the InstanceProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceProfile

`func (o *Operation) SetInstanceProfile(v string)`

SetInstanceProfile sets InstanceProfile field to given value.


### SetInstanceProfileNil

`func (o *Operation) SetInstanceProfileNil(b bool)`

 SetInstanceProfileNil sets the value for InstanceProfile to be an explicit nil

### UnsetInstanceProfile
`func (o *Operation) UnsetInstanceProfile()`

UnsetInstanceProfile ensures that no value is present for InstanceProfile, not even an explicit nil
### GetMountPoint

`func (o *Operation) GetMountPoint() string`

GetMountPoint returns the MountPoint field if non-nil, zero value otherwise.

### GetMountPointOk

`func (o *Operation) GetMountPointOk() (*string, bool)`

GetMountPointOk returns a tuple with the MountPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountPoint

`func (o *Operation) SetMountPoint(v string)`

SetMountPoint sets MountPoint field to given value.


### SetMountPointNil

`func (o *Operation) SetMountPointNil(b bool)`

 SetMountPointNil sets the value for MountPoint to be an explicit nil

### UnsetMountPoint
`func (o *Operation) UnsetMountPoint()`

UnsetMountPoint ensures that no value is present for MountPoint, not even an explicit nil
### GetDaily

`func (o *Operation) GetDaily() int32`

GetDaily returns the Daily field if non-nil, zero value otherwise.

### GetDailyOk

`func (o *Operation) GetDailyOk() (*int32, bool)`

GetDailyOk returns a tuple with the Daily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaily

`func (o *Operation) SetDaily(v int32)`

SetDaily sets Daily field to given value.


### SetDailyNil

`func (o *Operation) SetDailyNil(b bool)`

 SetDailyNil sets the value for Daily to be an explicit nil

### UnsetDaily
`func (o *Operation) UnsetDaily()`

UnsetDaily ensures that no value is present for Daily, not even an explicit nil
### GetMonthly

`func (o *Operation) GetMonthly() int32`

GetMonthly returns the Monthly field if non-nil, zero value otherwise.

### GetMonthlyOk

`func (o *Operation) GetMonthlyOk() (*int32, bool)`

GetMonthlyOk returns a tuple with the Monthly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthly

`func (o *Operation) SetMonthly(v int32)`

SetMonthly sets Monthly field to given value.


### SetMonthlyNil

`func (o *Operation) SetMonthlyNil(b bool)`

 SetMonthlyNil sets the value for Monthly to be an explicit nil

### UnsetMonthly
`func (o *Operation) UnsetMonthly()`

UnsetMonthly ensures that no value is present for Monthly, not even an explicit nil
### GetYearly

`func (o *Operation) GetYearly() int32`

GetYearly returns the Yearly field if non-nil, zero value otherwise.

### GetYearlyOk

`func (o *Operation) GetYearlyOk() (*int32, bool)`

GetYearlyOk returns a tuple with the Yearly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYearly

`func (o *Operation) SetYearly(v int32)`

SetYearly sets Yearly field to given value.


### SetYearlyNil

`func (o *Operation) SetYearlyNil(b bool)`

 SetYearlyNil sets the value for Yearly to be an explicit nil

### UnsetYearly
`func (o *Operation) UnsetYearly()`

UnsetYearly ensures that no value is present for Yearly, not even an explicit nil
### GetPitrDays

`func (o *Operation) GetPitrDays() int32`

GetPitrDays returns the PitrDays field if non-nil, zero value otherwise.

### GetPitrDaysOk

`func (o *Operation) GetPitrDaysOk() (*int32, bool)`

GetPitrDaysOk returns a tuple with the PitrDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPitrDays

`func (o *Operation) SetPitrDays(v int32)`

SetPitrDays sets PitrDays field to given value.


### SetPitrDaysNil

`func (o *Operation) SetPitrDaysNil(b bool)`

 SetPitrDaysNil sets the value for PitrDays to be an explicit nil

### UnsetPitrDays
`func (o *Operation) UnsetPitrDays()`

UnsetPitrDays ensures that no value is present for PitrDays, not even an explicit nil
### GetMakeCopy

`func (o *Operation) GetMakeCopy() bool`

GetMakeCopy returns the MakeCopy field if non-nil, zero value otherwise.

### GetMakeCopyOk

`func (o *Operation) GetMakeCopyOk() (*bool, bool)`

GetMakeCopyOk returns a tuple with the MakeCopy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMakeCopy

`func (o *Operation) SetMakeCopy(v bool)`

SetMakeCopy sets MakeCopy field to given value.


### SetMakeCopyNil

`func (o *Operation) SetMakeCopyNil(b bool)`

 SetMakeCopyNil sets the value for MakeCopy to be an explicit nil

### UnsetMakeCopy
`func (o *Operation) UnsetMakeCopy()`

UnsetMakeCopy ensures that no value is present for MakeCopy, not even an explicit nil
### GetKeepFinal

`func (o *Operation) GetKeepFinal() bool`

GetKeepFinal returns the KeepFinal field if non-nil, zero value otherwise.

### GetKeepFinalOk

`func (o *Operation) GetKeepFinalOk() (*bool, bool)`

GetKeepFinalOk returns a tuple with the KeepFinal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepFinal

`func (o *Operation) SetKeepFinal(v bool)`

SetKeepFinal sets KeepFinal field to given value.


### SetKeepFinalNil

`func (o *Operation) SetKeepFinalNil(b bool)`

 SetKeepFinalNil sets the value for KeepFinal to be an explicit nil

### UnsetKeepFinal
`func (o *Operation) UnsetKeepFinal()`

UnsetKeepFinal ensures that no value is present for KeepFinal, not even an explicit nil
### GetLinks

`func (o *Operation) GetLinks() OperationLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Operation) GetLinksOk() (*OperationLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Operation) SetLinks(v OperationLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Operation) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


