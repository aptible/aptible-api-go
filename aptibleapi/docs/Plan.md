# Plan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**MetaType** | **string** |  | 
**ContainerMemoryLimitMb** | **int32** |  | 
**CpuAllowedProfiles** | **[]string** |  | 
**DiskLimitGb** | **int32** |  | 
**EphemeralSessionLimit** | **int32** |  | 
**EnvironmentLimit** | **int32** |  | 
**VhostLimit** | **int32** |  | 
**CostCents** | **int32** |  | 
**IncludedContainerMb** | **int32** |  | 
**IncludedDiskGb** | **int32** |  | 
**IncludedVhosts** | **int32** |  | 
**ComplianceDashboardAccess** | **bool** |  | 
**CreatedAt** | **string** |  | 
**UpdatedAt** | **string** |  | 
**Links** | Pointer to [**IntegrationLinks**](IntegrationLinks.md) |  | [optional] 

## Methods

### NewPlan

`func NewPlan(id int32, metaType string, containerMemoryLimitMb int32, cpuAllowedProfiles []string, diskLimitGb int32, ephemeralSessionLimit int32, environmentLimit int32, vhostLimit int32, costCents int32, includedContainerMb int32, includedDiskGb int32, includedVhosts int32, complianceDashboardAccess bool, createdAt string, updatedAt string, ) *Plan`

NewPlan instantiates a new Plan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlanWithDefaults

`func NewPlanWithDefaults() *Plan`

NewPlanWithDefaults instantiates a new Plan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Plan) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Plan) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Plan) SetId(v int32)`

SetId sets Id field to given value.


### GetMetaType

`func (o *Plan) GetMetaType() string`

GetMetaType returns the MetaType field if non-nil, zero value otherwise.

### GetMetaTypeOk

`func (o *Plan) GetMetaTypeOk() (*string, bool)`

GetMetaTypeOk returns a tuple with the MetaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaType

`func (o *Plan) SetMetaType(v string)`

SetMetaType sets MetaType field to given value.


### GetContainerMemoryLimitMb

`func (o *Plan) GetContainerMemoryLimitMb() int32`

GetContainerMemoryLimitMb returns the ContainerMemoryLimitMb field if non-nil, zero value otherwise.

### GetContainerMemoryLimitMbOk

`func (o *Plan) GetContainerMemoryLimitMbOk() (*int32, bool)`

GetContainerMemoryLimitMbOk returns a tuple with the ContainerMemoryLimitMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerMemoryLimitMb

`func (o *Plan) SetContainerMemoryLimitMb(v int32)`

SetContainerMemoryLimitMb sets ContainerMemoryLimitMb field to given value.


### GetCpuAllowedProfiles

`func (o *Plan) GetCpuAllowedProfiles() []string`

GetCpuAllowedProfiles returns the CpuAllowedProfiles field if non-nil, zero value otherwise.

### GetCpuAllowedProfilesOk

`func (o *Plan) GetCpuAllowedProfilesOk() (*[]string, bool)`

GetCpuAllowedProfilesOk returns a tuple with the CpuAllowedProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuAllowedProfiles

`func (o *Plan) SetCpuAllowedProfiles(v []string)`

SetCpuAllowedProfiles sets CpuAllowedProfiles field to given value.


### GetDiskLimitGb

`func (o *Plan) GetDiskLimitGb() int32`

GetDiskLimitGb returns the DiskLimitGb field if non-nil, zero value otherwise.

### GetDiskLimitGbOk

`func (o *Plan) GetDiskLimitGbOk() (*int32, bool)`

GetDiskLimitGbOk returns a tuple with the DiskLimitGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskLimitGb

`func (o *Plan) SetDiskLimitGb(v int32)`

SetDiskLimitGb sets DiskLimitGb field to given value.


### GetEphemeralSessionLimit

`func (o *Plan) GetEphemeralSessionLimit() int32`

GetEphemeralSessionLimit returns the EphemeralSessionLimit field if non-nil, zero value otherwise.

### GetEphemeralSessionLimitOk

`func (o *Plan) GetEphemeralSessionLimitOk() (*int32, bool)`

GetEphemeralSessionLimitOk returns a tuple with the EphemeralSessionLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEphemeralSessionLimit

`func (o *Plan) SetEphemeralSessionLimit(v int32)`

SetEphemeralSessionLimit sets EphemeralSessionLimit field to given value.


### GetEnvironmentLimit

`func (o *Plan) GetEnvironmentLimit() int32`

GetEnvironmentLimit returns the EnvironmentLimit field if non-nil, zero value otherwise.

### GetEnvironmentLimitOk

`func (o *Plan) GetEnvironmentLimitOk() (*int32, bool)`

GetEnvironmentLimitOk returns a tuple with the EnvironmentLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentLimit

`func (o *Plan) SetEnvironmentLimit(v int32)`

SetEnvironmentLimit sets EnvironmentLimit field to given value.


### GetVhostLimit

`func (o *Plan) GetVhostLimit() int32`

GetVhostLimit returns the VhostLimit field if non-nil, zero value otherwise.

### GetVhostLimitOk

`func (o *Plan) GetVhostLimitOk() (*int32, bool)`

GetVhostLimitOk returns a tuple with the VhostLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVhostLimit

`func (o *Plan) SetVhostLimit(v int32)`

SetVhostLimit sets VhostLimit field to given value.


### GetCostCents

`func (o *Plan) GetCostCents() int32`

GetCostCents returns the CostCents field if non-nil, zero value otherwise.

### GetCostCentsOk

`func (o *Plan) GetCostCentsOk() (*int32, bool)`

GetCostCentsOk returns a tuple with the CostCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostCents

`func (o *Plan) SetCostCents(v int32)`

SetCostCents sets CostCents field to given value.


### GetIncludedContainerMb

`func (o *Plan) GetIncludedContainerMb() int32`

GetIncludedContainerMb returns the IncludedContainerMb field if non-nil, zero value otherwise.

### GetIncludedContainerMbOk

`func (o *Plan) GetIncludedContainerMbOk() (*int32, bool)`

GetIncludedContainerMbOk returns a tuple with the IncludedContainerMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedContainerMb

`func (o *Plan) SetIncludedContainerMb(v int32)`

SetIncludedContainerMb sets IncludedContainerMb field to given value.


### GetIncludedDiskGb

`func (o *Plan) GetIncludedDiskGb() int32`

GetIncludedDiskGb returns the IncludedDiskGb field if non-nil, zero value otherwise.

### GetIncludedDiskGbOk

`func (o *Plan) GetIncludedDiskGbOk() (*int32, bool)`

GetIncludedDiskGbOk returns a tuple with the IncludedDiskGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedDiskGb

`func (o *Plan) SetIncludedDiskGb(v int32)`

SetIncludedDiskGb sets IncludedDiskGb field to given value.


### GetIncludedVhosts

`func (o *Plan) GetIncludedVhosts() int32`

GetIncludedVhosts returns the IncludedVhosts field if non-nil, zero value otherwise.

### GetIncludedVhostsOk

`func (o *Plan) GetIncludedVhostsOk() (*int32, bool)`

GetIncludedVhostsOk returns a tuple with the IncludedVhosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedVhosts

`func (o *Plan) SetIncludedVhosts(v int32)`

SetIncludedVhosts sets IncludedVhosts field to given value.


### GetComplianceDashboardAccess

`func (o *Plan) GetComplianceDashboardAccess() bool`

GetComplianceDashboardAccess returns the ComplianceDashboardAccess field if non-nil, zero value otherwise.

### GetComplianceDashboardAccessOk

`func (o *Plan) GetComplianceDashboardAccessOk() (*bool, bool)`

GetComplianceDashboardAccessOk returns a tuple with the ComplianceDashboardAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplianceDashboardAccess

`func (o *Plan) SetComplianceDashboardAccess(v bool)`

SetComplianceDashboardAccess sets ComplianceDashboardAccess field to given value.


### GetCreatedAt

`func (o *Plan) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Plan) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Plan) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Plan) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Plan) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Plan) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetLinks

`func (o *Plan) GetLinks() IntegrationLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Plan) GetLinksOk() (*IntegrationLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Plan) SetLinks(v IntegrationLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Plan) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


