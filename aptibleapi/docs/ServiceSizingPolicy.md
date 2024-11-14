# ServiceSizingPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**MetaType** | **string** |  | 
**CreatedAt** | **string** |  | 
**UpdatedAt** | **string** |  | 
**ScalingEnabled** | **bool** |  | 
**DefaultPolicy** | **bool** |  | 
**MetricLookbackSeconds** | **int32** |  | 
**Percentile** | **float32** |  | 
**PostScaleUpCooldownSeconds** | **int32** |  | 
**PostScaleDownCooldownSeconds** | **int32** |  | 
**PostReleaseCooldownSeconds** | **int32** |  | 
**MemCpuRatioRThreshold** | **float32** |  | 
**MemCpuRatioCThreshold** | **float32** |  | 
**MemScaleUpThreshold** | **float32** |  | 
**MemScaleDownThreshold** | **float32** |  | 
**MinimumMemory** | **int32** |  | 
**AccountId** | **NullableInt32** |  | 
**MaximumMemory** | **NullableInt32** |  | 
**Autoscaling** | **string** |  | 
**MinCpuThreshold** | **NullableFloat32** |  | 
**MaxCpuThreshold** | **NullableFloat32** |  | 
**MinContainers** | **NullableInt32** |  | 
**MaxContainers** | **NullableInt32** |  | 
**ScaleUpStep** | **NullableInt32** |  | 
**ScaleDownStep** | **NullableInt32** |  | 
**Links** | Pointer to [**BackupRetentionPolicyLinks**](BackupRetentionPolicyLinks.md) |  | [optional] 

## Methods

### NewServiceSizingPolicy

`func NewServiceSizingPolicy(id int32, metaType string, createdAt string, updatedAt string, scalingEnabled bool, defaultPolicy bool, metricLookbackSeconds int32, percentile float32, postScaleUpCooldownSeconds int32, postScaleDownCooldownSeconds int32, postReleaseCooldownSeconds int32, memCpuRatioRThreshold float32, memCpuRatioCThreshold float32, memScaleUpThreshold float32, memScaleDownThreshold float32, minimumMemory int32, accountId NullableInt32, maximumMemory NullableInt32, autoscaling string, minCpuThreshold NullableFloat32, maxCpuThreshold NullableFloat32, minContainers NullableInt32, maxContainers NullableInt32, scaleUpStep NullableInt32, scaleDownStep NullableInt32, ) *ServiceSizingPolicy`

NewServiceSizingPolicy instantiates a new ServiceSizingPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceSizingPolicyWithDefaults

`func NewServiceSizingPolicyWithDefaults() *ServiceSizingPolicy`

NewServiceSizingPolicyWithDefaults instantiates a new ServiceSizingPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ServiceSizingPolicy) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServiceSizingPolicy) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServiceSizingPolicy) SetId(v int32)`

SetId sets Id field to given value.


### GetMetaType

`func (o *ServiceSizingPolicy) GetMetaType() string`

GetMetaType returns the MetaType field if non-nil, zero value otherwise.

### GetMetaTypeOk

`func (o *ServiceSizingPolicy) GetMetaTypeOk() (*string, bool)`

GetMetaTypeOk returns a tuple with the MetaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaType

`func (o *ServiceSizingPolicy) SetMetaType(v string)`

SetMetaType sets MetaType field to given value.


### GetCreatedAt

`func (o *ServiceSizingPolicy) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ServiceSizingPolicy) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ServiceSizingPolicy) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *ServiceSizingPolicy) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ServiceSizingPolicy) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ServiceSizingPolicy) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetScalingEnabled

`func (o *ServiceSizingPolicy) GetScalingEnabled() bool`

GetScalingEnabled returns the ScalingEnabled field if non-nil, zero value otherwise.

### GetScalingEnabledOk

`func (o *ServiceSizingPolicy) GetScalingEnabledOk() (*bool, bool)`

GetScalingEnabledOk returns a tuple with the ScalingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScalingEnabled

`func (o *ServiceSizingPolicy) SetScalingEnabled(v bool)`

SetScalingEnabled sets ScalingEnabled field to given value.


### GetDefaultPolicy

`func (o *ServiceSizingPolicy) GetDefaultPolicy() bool`

GetDefaultPolicy returns the DefaultPolicy field if non-nil, zero value otherwise.

### GetDefaultPolicyOk

`func (o *ServiceSizingPolicy) GetDefaultPolicyOk() (*bool, bool)`

GetDefaultPolicyOk returns a tuple with the DefaultPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPolicy

`func (o *ServiceSizingPolicy) SetDefaultPolicy(v bool)`

SetDefaultPolicy sets DefaultPolicy field to given value.


### GetMetricLookbackSeconds

`func (o *ServiceSizingPolicy) GetMetricLookbackSeconds() int32`

GetMetricLookbackSeconds returns the MetricLookbackSeconds field if non-nil, zero value otherwise.

### GetMetricLookbackSecondsOk

`func (o *ServiceSizingPolicy) GetMetricLookbackSecondsOk() (*int32, bool)`

GetMetricLookbackSecondsOk returns a tuple with the MetricLookbackSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricLookbackSeconds

`func (o *ServiceSizingPolicy) SetMetricLookbackSeconds(v int32)`

SetMetricLookbackSeconds sets MetricLookbackSeconds field to given value.


### GetPercentile

`func (o *ServiceSizingPolicy) GetPercentile() float32`

GetPercentile returns the Percentile field if non-nil, zero value otherwise.

### GetPercentileOk

`func (o *ServiceSizingPolicy) GetPercentileOk() (*float32, bool)`

GetPercentileOk returns a tuple with the Percentile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentile

`func (o *ServiceSizingPolicy) SetPercentile(v float32)`

SetPercentile sets Percentile field to given value.


### GetPostScaleUpCooldownSeconds

`func (o *ServiceSizingPolicy) GetPostScaleUpCooldownSeconds() int32`

GetPostScaleUpCooldownSeconds returns the PostScaleUpCooldownSeconds field if non-nil, zero value otherwise.

### GetPostScaleUpCooldownSecondsOk

`func (o *ServiceSizingPolicy) GetPostScaleUpCooldownSecondsOk() (*int32, bool)`

GetPostScaleUpCooldownSecondsOk returns a tuple with the PostScaleUpCooldownSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostScaleUpCooldownSeconds

`func (o *ServiceSizingPolicy) SetPostScaleUpCooldownSeconds(v int32)`

SetPostScaleUpCooldownSeconds sets PostScaleUpCooldownSeconds field to given value.


### GetPostScaleDownCooldownSeconds

`func (o *ServiceSizingPolicy) GetPostScaleDownCooldownSeconds() int32`

GetPostScaleDownCooldownSeconds returns the PostScaleDownCooldownSeconds field if non-nil, zero value otherwise.

### GetPostScaleDownCooldownSecondsOk

`func (o *ServiceSizingPolicy) GetPostScaleDownCooldownSecondsOk() (*int32, bool)`

GetPostScaleDownCooldownSecondsOk returns a tuple with the PostScaleDownCooldownSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostScaleDownCooldownSeconds

`func (o *ServiceSizingPolicy) SetPostScaleDownCooldownSeconds(v int32)`

SetPostScaleDownCooldownSeconds sets PostScaleDownCooldownSeconds field to given value.


### GetPostReleaseCooldownSeconds

`func (o *ServiceSizingPolicy) GetPostReleaseCooldownSeconds() int32`

GetPostReleaseCooldownSeconds returns the PostReleaseCooldownSeconds field if non-nil, zero value otherwise.

### GetPostReleaseCooldownSecondsOk

`func (o *ServiceSizingPolicy) GetPostReleaseCooldownSecondsOk() (*int32, bool)`

GetPostReleaseCooldownSecondsOk returns a tuple with the PostReleaseCooldownSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostReleaseCooldownSeconds

`func (o *ServiceSizingPolicy) SetPostReleaseCooldownSeconds(v int32)`

SetPostReleaseCooldownSeconds sets PostReleaseCooldownSeconds field to given value.


### GetMemCpuRatioRThreshold

`func (o *ServiceSizingPolicy) GetMemCpuRatioRThreshold() float32`

GetMemCpuRatioRThreshold returns the MemCpuRatioRThreshold field if non-nil, zero value otherwise.

### GetMemCpuRatioRThresholdOk

`func (o *ServiceSizingPolicy) GetMemCpuRatioRThresholdOk() (*float32, bool)`

GetMemCpuRatioRThresholdOk returns a tuple with the MemCpuRatioRThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemCpuRatioRThreshold

`func (o *ServiceSizingPolicy) SetMemCpuRatioRThreshold(v float32)`

SetMemCpuRatioRThreshold sets MemCpuRatioRThreshold field to given value.


### GetMemCpuRatioCThreshold

`func (o *ServiceSizingPolicy) GetMemCpuRatioCThreshold() float32`

GetMemCpuRatioCThreshold returns the MemCpuRatioCThreshold field if non-nil, zero value otherwise.

### GetMemCpuRatioCThresholdOk

`func (o *ServiceSizingPolicy) GetMemCpuRatioCThresholdOk() (*float32, bool)`

GetMemCpuRatioCThresholdOk returns a tuple with the MemCpuRatioCThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemCpuRatioCThreshold

`func (o *ServiceSizingPolicy) SetMemCpuRatioCThreshold(v float32)`

SetMemCpuRatioCThreshold sets MemCpuRatioCThreshold field to given value.


### GetMemScaleUpThreshold

`func (o *ServiceSizingPolicy) GetMemScaleUpThreshold() float32`

GetMemScaleUpThreshold returns the MemScaleUpThreshold field if non-nil, zero value otherwise.

### GetMemScaleUpThresholdOk

`func (o *ServiceSizingPolicy) GetMemScaleUpThresholdOk() (*float32, bool)`

GetMemScaleUpThresholdOk returns a tuple with the MemScaleUpThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemScaleUpThreshold

`func (o *ServiceSizingPolicy) SetMemScaleUpThreshold(v float32)`

SetMemScaleUpThreshold sets MemScaleUpThreshold field to given value.


### GetMemScaleDownThreshold

`func (o *ServiceSizingPolicy) GetMemScaleDownThreshold() float32`

GetMemScaleDownThreshold returns the MemScaleDownThreshold field if non-nil, zero value otherwise.

### GetMemScaleDownThresholdOk

`func (o *ServiceSizingPolicy) GetMemScaleDownThresholdOk() (*float32, bool)`

GetMemScaleDownThresholdOk returns a tuple with the MemScaleDownThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemScaleDownThreshold

`func (o *ServiceSizingPolicy) SetMemScaleDownThreshold(v float32)`

SetMemScaleDownThreshold sets MemScaleDownThreshold field to given value.


### GetMinimumMemory

`func (o *ServiceSizingPolicy) GetMinimumMemory() int32`

GetMinimumMemory returns the MinimumMemory field if non-nil, zero value otherwise.

### GetMinimumMemoryOk

`func (o *ServiceSizingPolicy) GetMinimumMemoryOk() (*int32, bool)`

GetMinimumMemoryOk returns a tuple with the MinimumMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumMemory

`func (o *ServiceSizingPolicy) SetMinimumMemory(v int32)`

SetMinimumMemory sets MinimumMemory field to given value.


### GetAccountId

`func (o *ServiceSizingPolicy) GetAccountId() int32`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *ServiceSizingPolicy) GetAccountIdOk() (*int32, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *ServiceSizingPolicy) SetAccountId(v int32)`

SetAccountId sets AccountId field to given value.


### SetAccountIdNil

`func (o *ServiceSizingPolicy) SetAccountIdNil(b bool)`

 SetAccountIdNil sets the value for AccountId to be an explicit nil

### UnsetAccountId
`func (o *ServiceSizingPolicy) UnsetAccountId()`

UnsetAccountId ensures that no value is present for AccountId, not even an explicit nil
### GetMaximumMemory

`func (o *ServiceSizingPolicy) GetMaximumMemory() int32`

GetMaximumMemory returns the MaximumMemory field if non-nil, zero value otherwise.

### GetMaximumMemoryOk

`func (o *ServiceSizingPolicy) GetMaximumMemoryOk() (*int32, bool)`

GetMaximumMemoryOk returns a tuple with the MaximumMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumMemory

`func (o *ServiceSizingPolicy) SetMaximumMemory(v int32)`

SetMaximumMemory sets MaximumMemory field to given value.


### SetMaximumMemoryNil

`func (o *ServiceSizingPolicy) SetMaximumMemoryNil(b bool)`

 SetMaximumMemoryNil sets the value for MaximumMemory to be an explicit nil

### UnsetMaximumMemory
`func (o *ServiceSizingPolicy) UnsetMaximumMemory()`

UnsetMaximumMemory ensures that no value is present for MaximumMemory, not even an explicit nil
### GetAutoscaling

`func (o *ServiceSizingPolicy) GetAutoscaling() string`

GetAutoscaling returns the Autoscaling field if non-nil, zero value otherwise.

### GetAutoscalingOk

`func (o *ServiceSizingPolicy) GetAutoscalingOk() (*string, bool)`

GetAutoscalingOk returns a tuple with the Autoscaling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoscaling

`func (o *ServiceSizingPolicy) SetAutoscaling(v string)`

SetAutoscaling sets Autoscaling field to given value.


### GetMinCpuThreshold

`func (o *ServiceSizingPolicy) GetMinCpuThreshold() float32`

GetMinCpuThreshold returns the MinCpuThreshold field if non-nil, zero value otherwise.

### GetMinCpuThresholdOk

`func (o *ServiceSizingPolicy) GetMinCpuThresholdOk() (*float32, bool)`

GetMinCpuThresholdOk returns a tuple with the MinCpuThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinCpuThreshold

`func (o *ServiceSizingPolicy) SetMinCpuThreshold(v float32)`

SetMinCpuThreshold sets MinCpuThreshold field to given value.


### SetMinCpuThresholdNil

`func (o *ServiceSizingPolicy) SetMinCpuThresholdNil(b bool)`

 SetMinCpuThresholdNil sets the value for MinCpuThreshold to be an explicit nil

### UnsetMinCpuThreshold
`func (o *ServiceSizingPolicy) UnsetMinCpuThreshold()`

UnsetMinCpuThreshold ensures that no value is present for MinCpuThreshold, not even an explicit nil
### GetMaxCpuThreshold

`func (o *ServiceSizingPolicy) GetMaxCpuThreshold() float32`

GetMaxCpuThreshold returns the MaxCpuThreshold field if non-nil, zero value otherwise.

### GetMaxCpuThresholdOk

`func (o *ServiceSizingPolicy) GetMaxCpuThresholdOk() (*float32, bool)`

GetMaxCpuThresholdOk returns a tuple with the MaxCpuThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCpuThreshold

`func (o *ServiceSizingPolicy) SetMaxCpuThreshold(v float32)`

SetMaxCpuThreshold sets MaxCpuThreshold field to given value.


### SetMaxCpuThresholdNil

`func (o *ServiceSizingPolicy) SetMaxCpuThresholdNil(b bool)`

 SetMaxCpuThresholdNil sets the value for MaxCpuThreshold to be an explicit nil

### UnsetMaxCpuThreshold
`func (o *ServiceSizingPolicy) UnsetMaxCpuThreshold()`

UnsetMaxCpuThreshold ensures that no value is present for MaxCpuThreshold, not even an explicit nil
### GetMinContainers

`func (o *ServiceSizingPolicy) GetMinContainers() int32`

GetMinContainers returns the MinContainers field if non-nil, zero value otherwise.

### GetMinContainersOk

`func (o *ServiceSizingPolicy) GetMinContainersOk() (*int32, bool)`

GetMinContainersOk returns a tuple with the MinContainers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinContainers

`func (o *ServiceSizingPolicy) SetMinContainers(v int32)`

SetMinContainers sets MinContainers field to given value.


### SetMinContainersNil

`func (o *ServiceSizingPolicy) SetMinContainersNil(b bool)`

 SetMinContainersNil sets the value for MinContainers to be an explicit nil

### UnsetMinContainers
`func (o *ServiceSizingPolicy) UnsetMinContainers()`

UnsetMinContainers ensures that no value is present for MinContainers, not even an explicit nil
### GetMaxContainers

`func (o *ServiceSizingPolicy) GetMaxContainers() int32`

GetMaxContainers returns the MaxContainers field if non-nil, zero value otherwise.

### GetMaxContainersOk

`func (o *ServiceSizingPolicy) GetMaxContainersOk() (*int32, bool)`

GetMaxContainersOk returns a tuple with the MaxContainers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxContainers

`func (o *ServiceSizingPolicy) SetMaxContainers(v int32)`

SetMaxContainers sets MaxContainers field to given value.


### SetMaxContainersNil

`func (o *ServiceSizingPolicy) SetMaxContainersNil(b bool)`

 SetMaxContainersNil sets the value for MaxContainers to be an explicit nil

### UnsetMaxContainers
`func (o *ServiceSizingPolicy) UnsetMaxContainers()`

UnsetMaxContainers ensures that no value is present for MaxContainers, not even an explicit nil
### GetScaleUpStep

`func (o *ServiceSizingPolicy) GetScaleUpStep() int32`

GetScaleUpStep returns the ScaleUpStep field if non-nil, zero value otherwise.

### GetScaleUpStepOk

`func (o *ServiceSizingPolicy) GetScaleUpStepOk() (*int32, bool)`

GetScaleUpStepOk returns a tuple with the ScaleUpStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScaleUpStep

`func (o *ServiceSizingPolicy) SetScaleUpStep(v int32)`

SetScaleUpStep sets ScaleUpStep field to given value.


### SetScaleUpStepNil

`func (o *ServiceSizingPolicy) SetScaleUpStepNil(b bool)`

 SetScaleUpStepNil sets the value for ScaleUpStep to be an explicit nil

### UnsetScaleUpStep
`func (o *ServiceSizingPolicy) UnsetScaleUpStep()`

UnsetScaleUpStep ensures that no value is present for ScaleUpStep, not even an explicit nil
### GetScaleDownStep

`func (o *ServiceSizingPolicy) GetScaleDownStep() int32`

GetScaleDownStep returns the ScaleDownStep field if non-nil, zero value otherwise.

### GetScaleDownStepOk

`func (o *ServiceSizingPolicy) GetScaleDownStepOk() (*int32, bool)`

GetScaleDownStepOk returns a tuple with the ScaleDownStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScaleDownStep

`func (o *ServiceSizingPolicy) SetScaleDownStep(v int32)`

SetScaleDownStep sets ScaleDownStep field to given value.


### SetScaleDownStepNil

`func (o *ServiceSizingPolicy) SetScaleDownStepNil(b bool)`

 SetScaleDownStepNil sets the value for ScaleDownStep to be an explicit nil

### UnsetScaleDownStep
`func (o *ServiceSizingPolicy) UnsetScaleDownStep()`

UnsetScaleDownStep ensures that no value is present for ScaleDownStep, not even an explicit nil
### GetLinks

`func (o *ServiceSizingPolicy) GetLinks() BackupRetentionPolicyLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ServiceSizingPolicy) GetLinksOk() (*BackupRetentionPolicyLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ServiceSizingPolicy) SetLinks(v BackupRetentionPolicyLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ServiceSizingPolicy) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


