# Service

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**MetaType** | **string** |  | 
**Handle** | **string** |  | 
**DockerRepo** | **NullableString** |  | 
**DockerRef** | **NullableString** |  | 
**ProcessType** | **string** |  | 
**Command** | **string** |  | 
**ContainerCount** | **int32** |  | 
**CreatedAt** | **string** |  | 
**UpdatedAt** | **string** |  | 
**ContainerMemoryLimitMb** | **NullableInt32** |  | 
**InstanceClass** | **string** |  | 
**ForceZeroDowntime** | **bool** |  | 
**NaiveHealthCheck** | **bool** |  | 
**StopTimeout** | **NullableInt32** |  | 
**Links** | Pointer to [**ServiceLinks**](ServiceLinks.md) |  | [optional] 

## Methods

### NewService

`func NewService(id int32, metaType string, handle string, dockerRepo NullableString, dockerRef NullableString, processType string, command string, containerCount int32, createdAt string, updatedAt string, containerMemoryLimitMb NullableInt32, instanceClass string, forceZeroDowntime bool, naiveHealthCheck bool, stopTimeout NullableInt32, ) *Service`

NewService instantiates a new Service object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceWithDefaults

`func NewServiceWithDefaults() *Service`

NewServiceWithDefaults instantiates a new Service object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Service) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Service) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Service) SetId(v int32)`

SetId sets Id field to given value.


### GetMetaType

`func (o *Service) GetMetaType() string`

GetMetaType returns the MetaType field if non-nil, zero value otherwise.

### GetMetaTypeOk

`func (o *Service) GetMetaTypeOk() (*string, bool)`

GetMetaTypeOk returns a tuple with the MetaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaType

`func (o *Service) SetMetaType(v string)`

SetMetaType sets MetaType field to given value.


### GetHandle

`func (o *Service) GetHandle() string`

GetHandle returns the Handle field if non-nil, zero value otherwise.

### GetHandleOk

`func (o *Service) GetHandleOk() (*string, bool)`

GetHandleOk returns a tuple with the Handle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandle

`func (o *Service) SetHandle(v string)`

SetHandle sets Handle field to given value.


### GetDockerRepo

`func (o *Service) GetDockerRepo() string`

GetDockerRepo returns the DockerRepo field if non-nil, zero value otherwise.

### GetDockerRepoOk

`func (o *Service) GetDockerRepoOk() (*string, bool)`

GetDockerRepoOk returns a tuple with the DockerRepo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerRepo

`func (o *Service) SetDockerRepo(v string)`

SetDockerRepo sets DockerRepo field to given value.


### SetDockerRepoNil

`func (o *Service) SetDockerRepoNil(b bool)`

 SetDockerRepoNil sets the value for DockerRepo to be an explicit nil

### UnsetDockerRepo
`func (o *Service) UnsetDockerRepo()`

UnsetDockerRepo ensures that no value is present for DockerRepo, not even an explicit nil
### GetDockerRef

`func (o *Service) GetDockerRef() string`

GetDockerRef returns the DockerRef field if non-nil, zero value otherwise.

### GetDockerRefOk

`func (o *Service) GetDockerRefOk() (*string, bool)`

GetDockerRefOk returns a tuple with the DockerRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerRef

`func (o *Service) SetDockerRef(v string)`

SetDockerRef sets DockerRef field to given value.


### SetDockerRefNil

`func (o *Service) SetDockerRefNil(b bool)`

 SetDockerRefNil sets the value for DockerRef to be an explicit nil

### UnsetDockerRef
`func (o *Service) UnsetDockerRef()`

UnsetDockerRef ensures that no value is present for DockerRef, not even an explicit nil
### GetProcessType

`func (o *Service) GetProcessType() string`

GetProcessType returns the ProcessType field if non-nil, zero value otherwise.

### GetProcessTypeOk

`func (o *Service) GetProcessTypeOk() (*string, bool)`

GetProcessTypeOk returns a tuple with the ProcessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessType

`func (o *Service) SetProcessType(v string)`

SetProcessType sets ProcessType field to given value.


### GetCommand

`func (o *Service) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *Service) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *Service) SetCommand(v string)`

SetCommand sets Command field to given value.


### GetContainerCount

`func (o *Service) GetContainerCount() int32`

GetContainerCount returns the ContainerCount field if non-nil, zero value otherwise.

### GetContainerCountOk

`func (o *Service) GetContainerCountOk() (*int32, bool)`

GetContainerCountOk returns a tuple with the ContainerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerCount

`func (o *Service) SetContainerCount(v int32)`

SetContainerCount sets ContainerCount field to given value.


### GetCreatedAt

`func (o *Service) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Service) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Service) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Service) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Service) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Service) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetContainerMemoryLimitMb

`func (o *Service) GetContainerMemoryLimitMb() int32`

GetContainerMemoryLimitMb returns the ContainerMemoryLimitMb field if non-nil, zero value otherwise.

### GetContainerMemoryLimitMbOk

`func (o *Service) GetContainerMemoryLimitMbOk() (*int32, bool)`

GetContainerMemoryLimitMbOk returns a tuple with the ContainerMemoryLimitMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerMemoryLimitMb

`func (o *Service) SetContainerMemoryLimitMb(v int32)`

SetContainerMemoryLimitMb sets ContainerMemoryLimitMb field to given value.


### SetContainerMemoryLimitMbNil

`func (o *Service) SetContainerMemoryLimitMbNil(b bool)`

 SetContainerMemoryLimitMbNil sets the value for ContainerMemoryLimitMb to be an explicit nil

### UnsetContainerMemoryLimitMb
`func (o *Service) UnsetContainerMemoryLimitMb()`

UnsetContainerMemoryLimitMb ensures that no value is present for ContainerMemoryLimitMb, not even an explicit nil
### GetInstanceClass

`func (o *Service) GetInstanceClass() string`

GetInstanceClass returns the InstanceClass field if non-nil, zero value otherwise.

### GetInstanceClassOk

`func (o *Service) GetInstanceClassOk() (*string, bool)`

GetInstanceClassOk returns a tuple with the InstanceClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceClass

`func (o *Service) SetInstanceClass(v string)`

SetInstanceClass sets InstanceClass field to given value.


### GetForceZeroDowntime

`func (o *Service) GetForceZeroDowntime() bool`

GetForceZeroDowntime returns the ForceZeroDowntime field if non-nil, zero value otherwise.

### GetForceZeroDowntimeOk

`func (o *Service) GetForceZeroDowntimeOk() (*bool, bool)`

GetForceZeroDowntimeOk returns a tuple with the ForceZeroDowntime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceZeroDowntime

`func (o *Service) SetForceZeroDowntime(v bool)`

SetForceZeroDowntime sets ForceZeroDowntime field to given value.


### GetNaiveHealthCheck

`func (o *Service) GetNaiveHealthCheck() bool`

GetNaiveHealthCheck returns the NaiveHealthCheck field if non-nil, zero value otherwise.

### GetNaiveHealthCheckOk

`func (o *Service) GetNaiveHealthCheckOk() (*bool, bool)`

GetNaiveHealthCheckOk returns a tuple with the NaiveHealthCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNaiveHealthCheck

`func (o *Service) SetNaiveHealthCheck(v bool)`

SetNaiveHealthCheck sets NaiveHealthCheck field to given value.


### GetStopTimeout

`func (o *Service) GetStopTimeout() int32`

GetStopTimeout returns the StopTimeout field if non-nil, zero value otherwise.

### GetStopTimeoutOk

`func (o *Service) GetStopTimeoutOk() (*int32, bool)`

GetStopTimeoutOk returns a tuple with the StopTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopTimeout

`func (o *Service) SetStopTimeout(v int32)`

SetStopTimeout sets StopTimeout field to given value.


### SetStopTimeoutNil

`func (o *Service) SetStopTimeoutNil(b bool)`

 SetStopTimeoutNil sets the value for StopTimeout to be an explicit nil

### UnsetStopTimeout
`func (o *Service) UnsetStopTimeout()`

UnsetStopTimeout ensures that no value is present for StopTimeout, not even an explicit nil
### GetLinks

`func (o *Service) GetLinks() ServiceLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Service) GetLinksOk() (*ServiceLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Service) SetLinks(v ServiceLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Service) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


