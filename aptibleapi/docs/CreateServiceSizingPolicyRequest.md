# CreateServiceSizingPolicyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

### NewCreateServiceSizingPolicyRequest

`func NewCreateServiceSizingPolicyRequest() *CreateServiceSizingPolicyRequest`

NewCreateServiceSizingPolicyRequest instantiates a new CreateServiceSizingPolicyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateServiceSizingPolicyRequestWithDefaults

`func NewCreateServiceSizingPolicyRequestWithDefaults() *CreateServiceSizingPolicyRequest`

NewCreateServiceSizingPolicyRequestWithDefaults instantiates a new CreateServiceSizingPolicyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetricLookbackSeconds

`func (o *CreateServiceSizingPolicyRequest) GetMetricLookbackSeconds() int32`

GetMetricLookbackSeconds returns the MetricLookbackSeconds field if non-nil, zero value otherwise.

### GetMetricLookbackSecondsOk

`func (o *CreateServiceSizingPolicyRequest) GetMetricLookbackSecondsOk() (*int32, bool)`

GetMetricLookbackSecondsOk returns a tuple with the MetricLookbackSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricLookbackSeconds

`func (o *CreateServiceSizingPolicyRequest) SetMetricLookbackSeconds(v int32)`

SetMetricLookbackSeconds sets MetricLookbackSeconds field to given value.

### HasMetricLookbackSeconds

`func (o *CreateServiceSizingPolicyRequest) HasMetricLookbackSeconds() bool`

HasMetricLookbackSeconds returns a boolean if a field has been set.

### GetPercentile

`func (o *CreateServiceSizingPolicyRequest) GetPercentile() float32`

GetPercentile returns the Percentile field if non-nil, zero value otherwise.

### GetPercentileOk

`func (o *CreateServiceSizingPolicyRequest) GetPercentileOk() (*float32, bool)`

GetPercentileOk returns a tuple with the Percentile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentile

`func (o *CreateServiceSizingPolicyRequest) SetPercentile(v float32)`

SetPercentile sets Percentile field to given value.

### HasPercentile

`func (o *CreateServiceSizingPolicyRequest) HasPercentile() bool`

HasPercentile returns a boolean if a field has been set.

### GetPostScaleUpCooldownSeconds

`func (o *CreateServiceSizingPolicyRequest) GetPostScaleUpCooldownSeconds() int32`

GetPostScaleUpCooldownSeconds returns the PostScaleUpCooldownSeconds field if non-nil, zero value otherwise.

### GetPostScaleUpCooldownSecondsOk

`func (o *CreateServiceSizingPolicyRequest) GetPostScaleUpCooldownSecondsOk() (*int32, bool)`

GetPostScaleUpCooldownSecondsOk returns a tuple with the PostScaleUpCooldownSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostScaleUpCooldownSeconds

`func (o *CreateServiceSizingPolicyRequest) SetPostScaleUpCooldownSeconds(v int32)`

SetPostScaleUpCooldownSeconds sets PostScaleUpCooldownSeconds field to given value.

### HasPostScaleUpCooldownSeconds

`func (o *CreateServiceSizingPolicyRequest) HasPostScaleUpCooldownSeconds() bool`

HasPostScaleUpCooldownSeconds returns a boolean if a field has been set.

### GetPostScaleDownCooldownSeconds

`func (o *CreateServiceSizingPolicyRequest) GetPostScaleDownCooldownSeconds() int32`

GetPostScaleDownCooldownSeconds returns the PostScaleDownCooldownSeconds field if non-nil, zero value otherwise.

### GetPostScaleDownCooldownSecondsOk

`func (o *CreateServiceSizingPolicyRequest) GetPostScaleDownCooldownSecondsOk() (*int32, bool)`

GetPostScaleDownCooldownSecondsOk returns a tuple with the PostScaleDownCooldownSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostScaleDownCooldownSeconds

`func (o *CreateServiceSizingPolicyRequest) SetPostScaleDownCooldownSeconds(v int32)`

SetPostScaleDownCooldownSeconds sets PostScaleDownCooldownSeconds field to given value.

### HasPostScaleDownCooldownSeconds

`func (o *CreateServiceSizingPolicyRequest) HasPostScaleDownCooldownSeconds() bool`

HasPostScaleDownCooldownSeconds returns a boolean if a field has been set.

### GetPostReleaseCooldownSeconds

`func (o *CreateServiceSizingPolicyRequest) GetPostReleaseCooldownSeconds() int32`

GetPostReleaseCooldownSeconds returns the PostReleaseCooldownSeconds field if non-nil, zero value otherwise.

### GetPostReleaseCooldownSecondsOk

`func (o *CreateServiceSizingPolicyRequest) GetPostReleaseCooldownSecondsOk() (*int32, bool)`

GetPostReleaseCooldownSecondsOk returns a tuple with the PostReleaseCooldownSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostReleaseCooldownSeconds

`func (o *CreateServiceSizingPolicyRequest) SetPostReleaseCooldownSeconds(v int32)`

SetPostReleaseCooldownSeconds sets PostReleaseCooldownSeconds field to given value.

### HasPostReleaseCooldownSeconds

`func (o *CreateServiceSizingPolicyRequest) HasPostReleaseCooldownSeconds() bool`

HasPostReleaseCooldownSeconds returns a boolean if a field has been set.

### GetMemCpuRatioRThreshold

`func (o *CreateServiceSizingPolicyRequest) GetMemCpuRatioRThreshold() float32`

GetMemCpuRatioRThreshold returns the MemCpuRatioRThreshold field if non-nil, zero value otherwise.

### GetMemCpuRatioRThresholdOk

`func (o *CreateServiceSizingPolicyRequest) GetMemCpuRatioRThresholdOk() (*float32, bool)`

GetMemCpuRatioRThresholdOk returns a tuple with the MemCpuRatioRThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemCpuRatioRThreshold

`func (o *CreateServiceSizingPolicyRequest) SetMemCpuRatioRThreshold(v float32)`

SetMemCpuRatioRThreshold sets MemCpuRatioRThreshold field to given value.

### HasMemCpuRatioRThreshold

`func (o *CreateServiceSizingPolicyRequest) HasMemCpuRatioRThreshold() bool`

HasMemCpuRatioRThreshold returns a boolean if a field has been set.

### GetMemCpuRatioCThreshold

`func (o *CreateServiceSizingPolicyRequest) GetMemCpuRatioCThreshold() float32`

GetMemCpuRatioCThreshold returns the MemCpuRatioCThreshold field if non-nil, zero value otherwise.

### GetMemCpuRatioCThresholdOk

`func (o *CreateServiceSizingPolicyRequest) GetMemCpuRatioCThresholdOk() (*float32, bool)`

GetMemCpuRatioCThresholdOk returns a tuple with the MemCpuRatioCThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemCpuRatioCThreshold

`func (o *CreateServiceSizingPolicyRequest) SetMemCpuRatioCThreshold(v float32)`

SetMemCpuRatioCThreshold sets MemCpuRatioCThreshold field to given value.

### HasMemCpuRatioCThreshold

`func (o *CreateServiceSizingPolicyRequest) HasMemCpuRatioCThreshold() bool`

HasMemCpuRatioCThreshold returns a boolean if a field has been set.

### GetMemScaleUpThreshold

`func (o *CreateServiceSizingPolicyRequest) GetMemScaleUpThreshold() float32`

GetMemScaleUpThreshold returns the MemScaleUpThreshold field if non-nil, zero value otherwise.

### GetMemScaleUpThresholdOk

`func (o *CreateServiceSizingPolicyRequest) GetMemScaleUpThresholdOk() (*float32, bool)`

GetMemScaleUpThresholdOk returns a tuple with the MemScaleUpThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemScaleUpThreshold

`func (o *CreateServiceSizingPolicyRequest) SetMemScaleUpThreshold(v float32)`

SetMemScaleUpThreshold sets MemScaleUpThreshold field to given value.

### HasMemScaleUpThreshold

`func (o *CreateServiceSizingPolicyRequest) HasMemScaleUpThreshold() bool`

HasMemScaleUpThreshold returns a boolean if a field has been set.

### GetMemScaleDownThreshold

`func (o *CreateServiceSizingPolicyRequest) GetMemScaleDownThreshold() float32`

GetMemScaleDownThreshold returns the MemScaleDownThreshold field if non-nil, zero value otherwise.

### GetMemScaleDownThresholdOk

`func (o *CreateServiceSizingPolicyRequest) GetMemScaleDownThresholdOk() (*float32, bool)`

GetMemScaleDownThresholdOk returns a tuple with the MemScaleDownThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemScaleDownThreshold

`func (o *CreateServiceSizingPolicyRequest) SetMemScaleDownThreshold(v float32)`

SetMemScaleDownThreshold sets MemScaleDownThreshold field to given value.

### HasMemScaleDownThreshold

`func (o *CreateServiceSizingPolicyRequest) HasMemScaleDownThreshold() bool`

HasMemScaleDownThreshold returns a boolean if a field has been set.

### GetMinimumMemory

`func (o *CreateServiceSizingPolicyRequest) GetMinimumMemory() int32`

GetMinimumMemory returns the MinimumMemory field if non-nil, zero value otherwise.

### GetMinimumMemoryOk

`func (o *CreateServiceSizingPolicyRequest) GetMinimumMemoryOk() (*int32, bool)`

GetMinimumMemoryOk returns a tuple with the MinimumMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumMemory

`func (o *CreateServiceSizingPolicyRequest) SetMinimumMemory(v int32)`

SetMinimumMemory sets MinimumMemory field to given value.

### HasMinimumMemory

`func (o *CreateServiceSizingPolicyRequest) HasMinimumMemory() bool`

HasMinimumMemory returns a boolean if a field has been set.

### GetMaximumMemory

`func (o *CreateServiceSizingPolicyRequest) GetMaximumMemory() int32`

GetMaximumMemory returns the MaximumMemory field if non-nil, zero value otherwise.

### GetMaximumMemoryOk

`func (o *CreateServiceSizingPolicyRequest) GetMaximumMemoryOk() (*int32, bool)`

GetMaximumMemoryOk returns a tuple with the MaximumMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumMemory

`func (o *CreateServiceSizingPolicyRequest) SetMaximumMemory(v int32)`

SetMaximumMemory sets MaximumMemory field to given value.

### HasMaximumMemory

`func (o *CreateServiceSizingPolicyRequest) HasMaximumMemory() bool`

HasMaximumMemory returns a boolean if a field has been set.

### GetAutoscaling

`func (o *CreateServiceSizingPolicyRequest) GetAutoscaling() string`

GetAutoscaling returns the Autoscaling field if non-nil, zero value otherwise.

### GetAutoscalingOk

`func (o *CreateServiceSizingPolicyRequest) GetAutoscalingOk() (*string, bool)`

GetAutoscalingOk returns a tuple with the Autoscaling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoscaling

`func (o *CreateServiceSizingPolicyRequest) SetAutoscaling(v string)`

SetAutoscaling sets Autoscaling field to given value.

### HasAutoscaling

`func (o *CreateServiceSizingPolicyRequest) HasAutoscaling() bool`

HasAutoscaling returns a boolean if a field has been set.

### GetMinCpuThreshold

`func (o *CreateServiceSizingPolicyRequest) GetMinCpuThreshold() float32`

GetMinCpuThreshold returns the MinCpuThreshold field if non-nil, zero value otherwise.

### GetMinCpuThresholdOk

`func (o *CreateServiceSizingPolicyRequest) GetMinCpuThresholdOk() (*float32, bool)`

GetMinCpuThresholdOk returns a tuple with the MinCpuThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinCpuThreshold

`func (o *CreateServiceSizingPolicyRequest) SetMinCpuThreshold(v float32)`

SetMinCpuThreshold sets MinCpuThreshold field to given value.

### HasMinCpuThreshold

`func (o *CreateServiceSizingPolicyRequest) HasMinCpuThreshold() bool`

HasMinCpuThreshold returns a boolean if a field has been set.

### GetMaxCpuThreshold

`func (o *CreateServiceSizingPolicyRequest) GetMaxCpuThreshold() float32`

GetMaxCpuThreshold returns the MaxCpuThreshold field if non-nil, zero value otherwise.

### GetMaxCpuThresholdOk

`func (o *CreateServiceSizingPolicyRequest) GetMaxCpuThresholdOk() (*float32, bool)`

GetMaxCpuThresholdOk returns a tuple with the MaxCpuThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCpuThreshold

`func (o *CreateServiceSizingPolicyRequest) SetMaxCpuThreshold(v float32)`

SetMaxCpuThreshold sets MaxCpuThreshold field to given value.

### HasMaxCpuThreshold

`func (o *CreateServiceSizingPolicyRequest) HasMaxCpuThreshold() bool`

HasMaxCpuThreshold returns a boolean if a field has been set.

### GetMinContainers

`func (o *CreateServiceSizingPolicyRequest) GetMinContainers() int32`

GetMinContainers returns the MinContainers field if non-nil, zero value otherwise.

### GetMinContainersOk

`func (o *CreateServiceSizingPolicyRequest) GetMinContainersOk() (*int32, bool)`

GetMinContainersOk returns a tuple with the MinContainers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinContainers

`func (o *CreateServiceSizingPolicyRequest) SetMinContainers(v int32)`

SetMinContainers sets MinContainers field to given value.

### HasMinContainers

`func (o *CreateServiceSizingPolicyRequest) HasMinContainers() bool`

HasMinContainers returns a boolean if a field has been set.

### GetMaxContainers

`func (o *CreateServiceSizingPolicyRequest) GetMaxContainers() int32`

GetMaxContainers returns the MaxContainers field if non-nil, zero value otherwise.

### GetMaxContainersOk

`func (o *CreateServiceSizingPolicyRequest) GetMaxContainersOk() (*int32, bool)`

GetMaxContainersOk returns a tuple with the MaxContainers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxContainers

`func (o *CreateServiceSizingPolicyRequest) SetMaxContainers(v int32)`

SetMaxContainers sets MaxContainers field to given value.

### HasMaxContainers

`func (o *CreateServiceSizingPolicyRequest) HasMaxContainers() bool`

HasMaxContainers returns a boolean if a field has been set.

### GetScaleUpStep

`func (o *CreateServiceSizingPolicyRequest) GetScaleUpStep() int32`

GetScaleUpStep returns the ScaleUpStep field if non-nil, zero value otherwise.

### GetScaleUpStepOk

`func (o *CreateServiceSizingPolicyRequest) GetScaleUpStepOk() (*int32, bool)`

GetScaleUpStepOk returns a tuple with the ScaleUpStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScaleUpStep

`func (o *CreateServiceSizingPolicyRequest) SetScaleUpStep(v int32)`

SetScaleUpStep sets ScaleUpStep field to given value.

### HasScaleUpStep

`func (o *CreateServiceSizingPolicyRequest) HasScaleUpStep() bool`

HasScaleUpStep returns a boolean if a field has been set.

### GetScaleDownStep

`func (o *CreateServiceSizingPolicyRequest) GetScaleDownStep() int32`

GetScaleDownStep returns the ScaleDownStep field if non-nil, zero value otherwise.

### GetScaleDownStepOk

`func (o *CreateServiceSizingPolicyRequest) GetScaleDownStepOk() (*int32, bool)`

GetScaleDownStepOk returns a tuple with the ScaleDownStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScaleDownStep

`func (o *CreateServiceSizingPolicyRequest) SetScaleDownStep(v int32)`

SetScaleDownStep sets ScaleDownStep field to given value.

### HasScaleDownStep

`func (o *CreateServiceSizingPolicyRequest) HasScaleDownStep() bool`

HasScaleDownStep returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


