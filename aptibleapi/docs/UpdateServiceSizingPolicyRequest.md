# UpdateServiceSizingPolicyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ScalingEnabled** | Pointer to **bool** |  | [optional] 
**MetricLookbackSeconds** | Pointer to **int32** |  | [optional] 
**Percentile** | Pointer to **float32** |  | [optional] 
**PostScaleUpCooldownSeconds** | Pointer to **int32** |  | [optional] 
**PostScaleDownCooldownSeconds** | Pointer to **int32** |  | [optional] 
**PostReleaseCooldownSeconds** | Pointer to **int32** |  | [optional] 
**MemCpuRatioRThreshold** | Pointer to **float32** |  | [optional] 
**MemCpuRatioCThreshold** | Pointer to **float32** |  | [optional] 
**MemScaleUpThreshold** | Pointer to **float32** |  | [optional] 
**MemScaleDownThreshold** | Pointer to **float32** |  | [optional] 
**MinimumMemory** | Pointer to **int32** |  | [optional] 
**MaximumMemory** | Pointer to **int32** |  | [optional] 
**Autoscaling** | Pointer to **string** |  | [optional] 
**MinCpuThreshold** | Pointer to **float32** |  | [optional] 
**MaxCpuThreshold** | Pointer to **float32** |  | [optional] 
**MinContainers** | Pointer to **int32** |  | [optional] 
**MaxContainers** | Pointer to **int32** |  | [optional] 
**ScaleUpStep** | Pointer to **int32** |  | [optional] 
**ScaleDownStep** | Pointer to **int32** |  | [optional] 

## Methods

### NewUpdateServiceSizingPolicyRequest

`func NewUpdateServiceSizingPolicyRequest() *UpdateServiceSizingPolicyRequest`

NewUpdateServiceSizingPolicyRequest instantiates a new UpdateServiceSizingPolicyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateServiceSizingPolicyRequestWithDefaults

`func NewUpdateServiceSizingPolicyRequestWithDefaults() *UpdateServiceSizingPolicyRequest`

NewUpdateServiceSizingPolicyRequestWithDefaults instantiates a new UpdateServiceSizingPolicyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScalingEnabled

`func (o *UpdateServiceSizingPolicyRequest) GetScalingEnabled() bool`

GetScalingEnabled returns the ScalingEnabled field if non-nil, zero value otherwise.

### GetScalingEnabledOk

`func (o *UpdateServiceSizingPolicyRequest) GetScalingEnabledOk() (*bool, bool)`

GetScalingEnabledOk returns a tuple with the ScalingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScalingEnabled

`func (o *UpdateServiceSizingPolicyRequest) SetScalingEnabled(v bool)`

SetScalingEnabled sets ScalingEnabled field to given value.

### HasScalingEnabled

`func (o *UpdateServiceSizingPolicyRequest) HasScalingEnabled() bool`

HasScalingEnabled returns a boolean if a field has been set.

### GetMetricLookbackSeconds

`func (o *UpdateServiceSizingPolicyRequest) GetMetricLookbackSeconds() int32`

GetMetricLookbackSeconds returns the MetricLookbackSeconds field if non-nil, zero value otherwise.

### GetMetricLookbackSecondsOk

`func (o *UpdateServiceSizingPolicyRequest) GetMetricLookbackSecondsOk() (*int32, bool)`

GetMetricLookbackSecondsOk returns a tuple with the MetricLookbackSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricLookbackSeconds

`func (o *UpdateServiceSizingPolicyRequest) SetMetricLookbackSeconds(v int32)`

SetMetricLookbackSeconds sets MetricLookbackSeconds field to given value.

### HasMetricLookbackSeconds

`func (o *UpdateServiceSizingPolicyRequest) HasMetricLookbackSeconds() bool`

HasMetricLookbackSeconds returns a boolean if a field has been set.

### GetPercentile

`func (o *UpdateServiceSizingPolicyRequest) GetPercentile() float32`

GetPercentile returns the Percentile field if non-nil, zero value otherwise.

### GetPercentileOk

`func (o *UpdateServiceSizingPolicyRequest) GetPercentileOk() (*float32, bool)`

GetPercentileOk returns a tuple with the Percentile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentile

`func (o *UpdateServiceSizingPolicyRequest) SetPercentile(v float32)`

SetPercentile sets Percentile field to given value.

### HasPercentile

`func (o *UpdateServiceSizingPolicyRequest) HasPercentile() bool`

HasPercentile returns a boolean if a field has been set.

### GetPostScaleUpCooldownSeconds

`func (o *UpdateServiceSizingPolicyRequest) GetPostScaleUpCooldownSeconds() int32`

GetPostScaleUpCooldownSeconds returns the PostScaleUpCooldownSeconds field if non-nil, zero value otherwise.

### GetPostScaleUpCooldownSecondsOk

`func (o *UpdateServiceSizingPolicyRequest) GetPostScaleUpCooldownSecondsOk() (*int32, bool)`

GetPostScaleUpCooldownSecondsOk returns a tuple with the PostScaleUpCooldownSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostScaleUpCooldownSeconds

`func (o *UpdateServiceSizingPolicyRequest) SetPostScaleUpCooldownSeconds(v int32)`

SetPostScaleUpCooldownSeconds sets PostScaleUpCooldownSeconds field to given value.

### HasPostScaleUpCooldownSeconds

`func (o *UpdateServiceSizingPolicyRequest) HasPostScaleUpCooldownSeconds() bool`

HasPostScaleUpCooldownSeconds returns a boolean if a field has been set.

### GetPostScaleDownCooldownSeconds

`func (o *UpdateServiceSizingPolicyRequest) GetPostScaleDownCooldownSeconds() int32`

GetPostScaleDownCooldownSeconds returns the PostScaleDownCooldownSeconds field if non-nil, zero value otherwise.

### GetPostScaleDownCooldownSecondsOk

`func (o *UpdateServiceSizingPolicyRequest) GetPostScaleDownCooldownSecondsOk() (*int32, bool)`

GetPostScaleDownCooldownSecondsOk returns a tuple with the PostScaleDownCooldownSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostScaleDownCooldownSeconds

`func (o *UpdateServiceSizingPolicyRequest) SetPostScaleDownCooldownSeconds(v int32)`

SetPostScaleDownCooldownSeconds sets PostScaleDownCooldownSeconds field to given value.

### HasPostScaleDownCooldownSeconds

`func (o *UpdateServiceSizingPolicyRequest) HasPostScaleDownCooldownSeconds() bool`

HasPostScaleDownCooldownSeconds returns a boolean if a field has been set.

### GetPostReleaseCooldownSeconds

`func (o *UpdateServiceSizingPolicyRequest) GetPostReleaseCooldownSeconds() int32`

GetPostReleaseCooldownSeconds returns the PostReleaseCooldownSeconds field if non-nil, zero value otherwise.

### GetPostReleaseCooldownSecondsOk

`func (o *UpdateServiceSizingPolicyRequest) GetPostReleaseCooldownSecondsOk() (*int32, bool)`

GetPostReleaseCooldownSecondsOk returns a tuple with the PostReleaseCooldownSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostReleaseCooldownSeconds

`func (o *UpdateServiceSizingPolicyRequest) SetPostReleaseCooldownSeconds(v int32)`

SetPostReleaseCooldownSeconds sets PostReleaseCooldownSeconds field to given value.

### HasPostReleaseCooldownSeconds

`func (o *UpdateServiceSizingPolicyRequest) HasPostReleaseCooldownSeconds() bool`

HasPostReleaseCooldownSeconds returns a boolean if a field has been set.

### GetMemCpuRatioRThreshold

`func (o *UpdateServiceSizingPolicyRequest) GetMemCpuRatioRThreshold() float32`

GetMemCpuRatioRThreshold returns the MemCpuRatioRThreshold field if non-nil, zero value otherwise.

### GetMemCpuRatioRThresholdOk

`func (o *UpdateServiceSizingPolicyRequest) GetMemCpuRatioRThresholdOk() (*float32, bool)`

GetMemCpuRatioRThresholdOk returns a tuple with the MemCpuRatioRThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemCpuRatioRThreshold

`func (o *UpdateServiceSizingPolicyRequest) SetMemCpuRatioRThreshold(v float32)`

SetMemCpuRatioRThreshold sets MemCpuRatioRThreshold field to given value.

### HasMemCpuRatioRThreshold

`func (o *UpdateServiceSizingPolicyRequest) HasMemCpuRatioRThreshold() bool`

HasMemCpuRatioRThreshold returns a boolean if a field has been set.

### GetMemCpuRatioCThreshold

`func (o *UpdateServiceSizingPolicyRequest) GetMemCpuRatioCThreshold() float32`

GetMemCpuRatioCThreshold returns the MemCpuRatioCThreshold field if non-nil, zero value otherwise.

### GetMemCpuRatioCThresholdOk

`func (o *UpdateServiceSizingPolicyRequest) GetMemCpuRatioCThresholdOk() (*float32, bool)`

GetMemCpuRatioCThresholdOk returns a tuple with the MemCpuRatioCThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemCpuRatioCThreshold

`func (o *UpdateServiceSizingPolicyRequest) SetMemCpuRatioCThreshold(v float32)`

SetMemCpuRatioCThreshold sets MemCpuRatioCThreshold field to given value.

### HasMemCpuRatioCThreshold

`func (o *UpdateServiceSizingPolicyRequest) HasMemCpuRatioCThreshold() bool`

HasMemCpuRatioCThreshold returns a boolean if a field has been set.

### GetMemScaleUpThreshold

`func (o *UpdateServiceSizingPolicyRequest) GetMemScaleUpThreshold() float32`

GetMemScaleUpThreshold returns the MemScaleUpThreshold field if non-nil, zero value otherwise.

### GetMemScaleUpThresholdOk

`func (o *UpdateServiceSizingPolicyRequest) GetMemScaleUpThresholdOk() (*float32, bool)`

GetMemScaleUpThresholdOk returns a tuple with the MemScaleUpThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemScaleUpThreshold

`func (o *UpdateServiceSizingPolicyRequest) SetMemScaleUpThreshold(v float32)`

SetMemScaleUpThreshold sets MemScaleUpThreshold field to given value.

### HasMemScaleUpThreshold

`func (o *UpdateServiceSizingPolicyRequest) HasMemScaleUpThreshold() bool`

HasMemScaleUpThreshold returns a boolean if a field has been set.

### GetMemScaleDownThreshold

`func (o *UpdateServiceSizingPolicyRequest) GetMemScaleDownThreshold() float32`

GetMemScaleDownThreshold returns the MemScaleDownThreshold field if non-nil, zero value otherwise.

### GetMemScaleDownThresholdOk

`func (o *UpdateServiceSizingPolicyRequest) GetMemScaleDownThresholdOk() (*float32, bool)`

GetMemScaleDownThresholdOk returns a tuple with the MemScaleDownThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemScaleDownThreshold

`func (o *UpdateServiceSizingPolicyRequest) SetMemScaleDownThreshold(v float32)`

SetMemScaleDownThreshold sets MemScaleDownThreshold field to given value.

### HasMemScaleDownThreshold

`func (o *UpdateServiceSizingPolicyRequest) HasMemScaleDownThreshold() bool`

HasMemScaleDownThreshold returns a boolean if a field has been set.

### GetMinimumMemory

`func (o *UpdateServiceSizingPolicyRequest) GetMinimumMemory() int32`

GetMinimumMemory returns the MinimumMemory field if non-nil, zero value otherwise.

### GetMinimumMemoryOk

`func (o *UpdateServiceSizingPolicyRequest) GetMinimumMemoryOk() (*int32, bool)`

GetMinimumMemoryOk returns a tuple with the MinimumMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumMemory

`func (o *UpdateServiceSizingPolicyRequest) SetMinimumMemory(v int32)`

SetMinimumMemory sets MinimumMemory field to given value.

### HasMinimumMemory

`func (o *UpdateServiceSizingPolicyRequest) HasMinimumMemory() bool`

HasMinimumMemory returns a boolean if a field has been set.

### GetMaximumMemory

`func (o *UpdateServiceSizingPolicyRequest) GetMaximumMemory() int32`

GetMaximumMemory returns the MaximumMemory field if non-nil, zero value otherwise.

### GetMaximumMemoryOk

`func (o *UpdateServiceSizingPolicyRequest) GetMaximumMemoryOk() (*int32, bool)`

GetMaximumMemoryOk returns a tuple with the MaximumMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumMemory

`func (o *UpdateServiceSizingPolicyRequest) SetMaximumMemory(v int32)`

SetMaximumMemory sets MaximumMemory field to given value.

### HasMaximumMemory

`func (o *UpdateServiceSizingPolicyRequest) HasMaximumMemory() bool`

HasMaximumMemory returns a boolean if a field has been set.

### GetAutoscaling

`func (o *UpdateServiceSizingPolicyRequest) GetAutoscaling() string`

GetAutoscaling returns the Autoscaling field if non-nil, zero value otherwise.

### GetAutoscalingOk

`func (o *UpdateServiceSizingPolicyRequest) GetAutoscalingOk() (*string, bool)`

GetAutoscalingOk returns a tuple with the Autoscaling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoscaling

`func (o *UpdateServiceSizingPolicyRequest) SetAutoscaling(v string)`

SetAutoscaling sets Autoscaling field to given value.

### HasAutoscaling

`func (o *UpdateServiceSizingPolicyRequest) HasAutoscaling() bool`

HasAutoscaling returns a boolean if a field has been set.

### GetMinCpuThreshold

`func (o *UpdateServiceSizingPolicyRequest) GetMinCpuThreshold() float32`

GetMinCpuThreshold returns the MinCpuThreshold field if non-nil, zero value otherwise.

### GetMinCpuThresholdOk

`func (o *UpdateServiceSizingPolicyRequest) GetMinCpuThresholdOk() (*float32, bool)`

GetMinCpuThresholdOk returns a tuple with the MinCpuThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinCpuThreshold

`func (o *UpdateServiceSizingPolicyRequest) SetMinCpuThreshold(v float32)`

SetMinCpuThreshold sets MinCpuThreshold field to given value.

### HasMinCpuThreshold

`func (o *UpdateServiceSizingPolicyRequest) HasMinCpuThreshold() bool`

HasMinCpuThreshold returns a boolean if a field has been set.

### GetMaxCpuThreshold

`func (o *UpdateServiceSizingPolicyRequest) GetMaxCpuThreshold() float32`

GetMaxCpuThreshold returns the MaxCpuThreshold field if non-nil, zero value otherwise.

### GetMaxCpuThresholdOk

`func (o *UpdateServiceSizingPolicyRequest) GetMaxCpuThresholdOk() (*float32, bool)`

GetMaxCpuThresholdOk returns a tuple with the MaxCpuThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCpuThreshold

`func (o *UpdateServiceSizingPolicyRequest) SetMaxCpuThreshold(v float32)`

SetMaxCpuThreshold sets MaxCpuThreshold field to given value.

### HasMaxCpuThreshold

`func (o *UpdateServiceSizingPolicyRequest) HasMaxCpuThreshold() bool`

HasMaxCpuThreshold returns a boolean if a field has been set.

### GetMinContainers

`func (o *UpdateServiceSizingPolicyRequest) GetMinContainers() int32`

GetMinContainers returns the MinContainers field if non-nil, zero value otherwise.

### GetMinContainersOk

`func (o *UpdateServiceSizingPolicyRequest) GetMinContainersOk() (*int32, bool)`

GetMinContainersOk returns a tuple with the MinContainers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinContainers

`func (o *UpdateServiceSizingPolicyRequest) SetMinContainers(v int32)`

SetMinContainers sets MinContainers field to given value.

### HasMinContainers

`func (o *UpdateServiceSizingPolicyRequest) HasMinContainers() bool`

HasMinContainers returns a boolean if a field has been set.

### GetMaxContainers

`func (o *UpdateServiceSizingPolicyRequest) GetMaxContainers() int32`

GetMaxContainers returns the MaxContainers field if non-nil, zero value otherwise.

### GetMaxContainersOk

`func (o *UpdateServiceSizingPolicyRequest) GetMaxContainersOk() (*int32, bool)`

GetMaxContainersOk returns a tuple with the MaxContainers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxContainers

`func (o *UpdateServiceSizingPolicyRequest) SetMaxContainers(v int32)`

SetMaxContainers sets MaxContainers field to given value.

### HasMaxContainers

`func (o *UpdateServiceSizingPolicyRequest) HasMaxContainers() bool`

HasMaxContainers returns a boolean if a field has been set.

### GetScaleUpStep

`func (o *UpdateServiceSizingPolicyRequest) GetScaleUpStep() int32`

GetScaleUpStep returns the ScaleUpStep field if non-nil, zero value otherwise.

### GetScaleUpStepOk

`func (o *UpdateServiceSizingPolicyRequest) GetScaleUpStepOk() (*int32, bool)`

GetScaleUpStepOk returns a tuple with the ScaleUpStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScaleUpStep

`func (o *UpdateServiceSizingPolicyRequest) SetScaleUpStep(v int32)`

SetScaleUpStep sets ScaleUpStep field to given value.

### HasScaleUpStep

`func (o *UpdateServiceSizingPolicyRequest) HasScaleUpStep() bool`

HasScaleUpStep returns a boolean if a field has been set.

### GetScaleDownStep

`func (o *UpdateServiceSizingPolicyRequest) GetScaleDownStep() int32`

GetScaleDownStep returns the ScaleDownStep field if non-nil, zero value otherwise.

### GetScaleDownStepOk

`func (o *UpdateServiceSizingPolicyRequest) GetScaleDownStepOk() (*int32, bool)`

GetScaleDownStepOk returns a tuple with the ScaleDownStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScaleDownStep

`func (o *UpdateServiceSizingPolicyRequest) SetScaleDownStep(v int32)`

SetScaleDownStep sets ScaleDownStep field to given value.

### HasScaleDownStep

`func (o *UpdateServiceSizingPolicyRequest) HasScaleDownStep() bool`

HasScaleDownStep returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


