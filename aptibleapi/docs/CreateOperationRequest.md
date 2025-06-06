# CreateOperationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**GitRef** | Pointer to **string** |  | [optional] 
**DiskSize** | Pointer to **int32** |  | [optional] 
**ContainerCount** | Pointer to **int32** |  | [optional] 
**ContainerSize** | Pointer to **int32** |  | [optional] 
**Command** | Pointer to **string** |  | [optional] 
**Env** | Pointer to **map[string]string** |  | [optional] 
**Handle** | Pointer to **string** |  | [optional] 
**Certificate** | Pointer to **string** |  | [optional] 
**PrivateKey** | Pointer to **string** |  | [optional] 
**DestinationRegion** | Pointer to **string** |  | [optional] 
**Interactive** | Pointer to **bool** |  | [optional] 
**DestinationAccount** | Pointer to **string** |  | [optional] 
**DestinationAccountId** | Pointer to **int32** |  | [optional] 
**DockerRef** | Pointer to **string** |  | [optional] 
**Automated** | Pointer to **bool** |  | [optional] 
**KeyArn** | Pointer to **string** |  | [optional] 
**InstanceProfile** | Pointer to **string** |  | [optional] 
**ProvisionedIops** | Pointer to **int32** |  | [optional] 
**Daily** | Pointer to **int32** |  | [optional] 
**Monthly** | Pointer to **int32** |  | [optional] 
**Yearly** | Pointer to **int32** |  | [optional] 
**PitrDays** | Pointer to **int32** |  | [optional] 
**MakeCopy** | Pointer to **bool** |  | [optional] 
**KeepFinal** | Pointer to **bool** |  | [optional] 

## Methods

### NewCreateOperationRequest

`func NewCreateOperationRequest(type_ string, ) *CreateOperationRequest`

NewCreateOperationRequest instantiates a new CreateOperationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOperationRequestWithDefaults

`func NewCreateOperationRequestWithDefaults() *CreateOperationRequest`

NewCreateOperationRequestWithDefaults instantiates a new CreateOperationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CreateOperationRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateOperationRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateOperationRequest) SetType(v string)`

SetType sets Type field to given value.


### GetGitRef

`func (o *CreateOperationRequest) GetGitRef() string`

GetGitRef returns the GitRef field if non-nil, zero value otherwise.

### GetGitRefOk

`func (o *CreateOperationRequest) GetGitRefOk() (*string, bool)`

GetGitRefOk returns a tuple with the GitRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitRef

`func (o *CreateOperationRequest) SetGitRef(v string)`

SetGitRef sets GitRef field to given value.

### HasGitRef

`func (o *CreateOperationRequest) HasGitRef() bool`

HasGitRef returns a boolean if a field has been set.

### GetDiskSize

`func (o *CreateOperationRequest) GetDiskSize() int32`

GetDiskSize returns the DiskSize field if non-nil, zero value otherwise.

### GetDiskSizeOk

`func (o *CreateOperationRequest) GetDiskSizeOk() (*int32, bool)`

GetDiskSizeOk returns a tuple with the DiskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSize

`func (o *CreateOperationRequest) SetDiskSize(v int32)`

SetDiskSize sets DiskSize field to given value.

### HasDiskSize

`func (o *CreateOperationRequest) HasDiskSize() bool`

HasDiskSize returns a boolean if a field has been set.

### GetContainerCount

`func (o *CreateOperationRequest) GetContainerCount() int32`

GetContainerCount returns the ContainerCount field if non-nil, zero value otherwise.

### GetContainerCountOk

`func (o *CreateOperationRequest) GetContainerCountOk() (*int32, bool)`

GetContainerCountOk returns a tuple with the ContainerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerCount

`func (o *CreateOperationRequest) SetContainerCount(v int32)`

SetContainerCount sets ContainerCount field to given value.

### HasContainerCount

`func (o *CreateOperationRequest) HasContainerCount() bool`

HasContainerCount returns a boolean if a field has been set.

### GetContainerSize

`func (o *CreateOperationRequest) GetContainerSize() int32`

GetContainerSize returns the ContainerSize field if non-nil, zero value otherwise.

### GetContainerSizeOk

`func (o *CreateOperationRequest) GetContainerSizeOk() (*int32, bool)`

GetContainerSizeOk returns a tuple with the ContainerSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerSize

`func (o *CreateOperationRequest) SetContainerSize(v int32)`

SetContainerSize sets ContainerSize field to given value.

### HasContainerSize

`func (o *CreateOperationRequest) HasContainerSize() bool`

HasContainerSize returns a boolean if a field has been set.

### GetCommand

`func (o *CreateOperationRequest) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *CreateOperationRequest) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *CreateOperationRequest) SetCommand(v string)`

SetCommand sets Command field to given value.

### HasCommand

`func (o *CreateOperationRequest) HasCommand() bool`

HasCommand returns a boolean if a field has been set.

### GetEnv

`func (o *CreateOperationRequest) GetEnv() map[string]string`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *CreateOperationRequest) GetEnvOk() (*map[string]string, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *CreateOperationRequest) SetEnv(v map[string]string)`

SetEnv sets Env field to given value.

### HasEnv

`func (o *CreateOperationRequest) HasEnv() bool`

HasEnv returns a boolean if a field has been set.

### GetHandle

`func (o *CreateOperationRequest) GetHandle() string`

GetHandle returns the Handle field if non-nil, zero value otherwise.

### GetHandleOk

`func (o *CreateOperationRequest) GetHandleOk() (*string, bool)`

GetHandleOk returns a tuple with the Handle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandle

`func (o *CreateOperationRequest) SetHandle(v string)`

SetHandle sets Handle field to given value.

### HasHandle

`func (o *CreateOperationRequest) HasHandle() bool`

HasHandle returns a boolean if a field has been set.

### GetCertificate

`func (o *CreateOperationRequest) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *CreateOperationRequest) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *CreateOperationRequest) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *CreateOperationRequest) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### GetPrivateKey

`func (o *CreateOperationRequest) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *CreateOperationRequest) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *CreateOperationRequest) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.

### HasPrivateKey

`func (o *CreateOperationRequest) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.

### GetDestinationRegion

`func (o *CreateOperationRequest) GetDestinationRegion() string`

GetDestinationRegion returns the DestinationRegion field if non-nil, zero value otherwise.

### GetDestinationRegionOk

`func (o *CreateOperationRequest) GetDestinationRegionOk() (*string, bool)`

GetDestinationRegionOk returns a tuple with the DestinationRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationRegion

`func (o *CreateOperationRequest) SetDestinationRegion(v string)`

SetDestinationRegion sets DestinationRegion field to given value.

### HasDestinationRegion

`func (o *CreateOperationRequest) HasDestinationRegion() bool`

HasDestinationRegion returns a boolean if a field has been set.

### GetInteractive

`func (o *CreateOperationRequest) GetInteractive() bool`

GetInteractive returns the Interactive field if non-nil, zero value otherwise.

### GetInteractiveOk

`func (o *CreateOperationRequest) GetInteractiveOk() (*bool, bool)`

GetInteractiveOk returns a tuple with the Interactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInteractive

`func (o *CreateOperationRequest) SetInteractive(v bool)`

SetInteractive sets Interactive field to given value.

### HasInteractive

`func (o *CreateOperationRequest) HasInteractive() bool`

HasInteractive returns a boolean if a field has been set.

### GetDestinationAccount

`func (o *CreateOperationRequest) GetDestinationAccount() string`

GetDestinationAccount returns the DestinationAccount field if non-nil, zero value otherwise.

### GetDestinationAccountOk

`func (o *CreateOperationRequest) GetDestinationAccountOk() (*string, bool)`

GetDestinationAccountOk returns a tuple with the DestinationAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationAccount

`func (o *CreateOperationRequest) SetDestinationAccount(v string)`

SetDestinationAccount sets DestinationAccount field to given value.

### HasDestinationAccount

`func (o *CreateOperationRequest) HasDestinationAccount() bool`

HasDestinationAccount returns a boolean if a field has been set.

### GetDestinationAccountId

`func (o *CreateOperationRequest) GetDestinationAccountId() int32`

GetDestinationAccountId returns the DestinationAccountId field if non-nil, zero value otherwise.

### GetDestinationAccountIdOk

`func (o *CreateOperationRequest) GetDestinationAccountIdOk() (*int32, bool)`

GetDestinationAccountIdOk returns a tuple with the DestinationAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationAccountId

`func (o *CreateOperationRequest) SetDestinationAccountId(v int32)`

SetDestinationAccountId sets DestinationAccountId field to given value.

### HasDestinationAccountId

`func (o *CreateOperationRequest) HasDestinationAccountId() bool`

HasDestinationAccountId returns a boolean if a field has been set.

### GetDockerRef

`func (o *CreateOperationRequest) GetDockerRef() string`

GetDockerRef returns the DockerRef field if non-nil, zero value otherwise.

### GetDockerRefOk

`func (o *CreateOperationRequest) GetDockerRefOk() (*string, bool)`

GetDockerRefOk returns a tuple with the DockerRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerRef

`func (o *CreateOperationRequest) SetDockerRef(v string)`

SetDockerRef sets DockerRef field to given value.

### HasDockerRef

`func (o *CreateOperationRequest) HasDockerRef() bool`

HasDockerRef returns a boolean if a field has been set.

### GetAutomated

`func (o *CreateOperationRequest) GetAutomated() bool`

GetAutomated returns the Automated field if non-nil, zero value otherwise.

### GetAutomatedOk

`func (o *CreateOperationRequest) GetAutomatedOk() (*bool, bool)`

GetAutomatedOk returns a tuple with the Automated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomated

`func (o *CreateOperationRequest) SetAutomated(v bool)`

SetAutomated sets Automated field to given value.

### HasAutomated

`func (o *CreateOperationRequest) HasAutomated() bool`

HasAutomated returns a boolean if a field has been set.

### GetKeyArn

`func (o *CreateOperationRequest) GetKeyArn() string`

GetKeyArn returns the KeyArn field if non-nil, zero value otherwise.

### GetKeyArnOk

`func (o *CreateOperationRequest) GetKeyArnOk() (*string, bool)`

GetKeyArnOk returns a tuple with the KeyArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyArn

`func (o *CreateOperationRequest) SetKeyArn(v string)`

SetKeyArn sets KeyArn field to given value.

### HasKeyArn

`func (o *CreateOperationRequest) HasKeyArn() bool`

HasKeyArn returns a boolean if a field has been set.

### GetInstanceProfile

`func (o *CreateOperationRequest) GetInstanceProfile() string`

GetInstanceProfile returns the InstanceProfile field if non-nil, zero value otherwise.

### GetInstanceProfileOk

`func (o *CreateOperationRequest) GetInstanceProfileOk() (*string, bool)`

GetInstanceProfileOk returns a tuple with the InstanceProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceProfile

`func (o *CreateOperationRequest) SetInstanceProfile(v string)`

SetInstanceProfile sets InstanceProfile field to given value.

### HasInstanceProfile

`func (o *CreateOperationRequest) HasInstanceProfile() bool`

HasInstanceProfile returns a boolean if a field has been set.

### GetProvisionedIops

`func (o *CreateOperationRequest) GetProvisionedIops() int32`

GetProvisionedIops returns the ProvisionedIops field if non-nil, zero value otherwise.

### GetProvisionedIopsOk

`func (o *CreateOperationRequest) GetProvisionedIopsOk() (*int32, bool)`

GetProvisionedIopsOk returns a tuple with the ProvisionedIops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionedIops

`func (o *CreateOperationRequest) SetProvisionedIops(v int32)`

SetProvisionedIops sets ProvisionedIops field to given value.

### HasProvisionedIops

`func (o *CreateOperationRequest) HasProvisionedIops() bool`

HasProvisionedIops returns a boolean if a field has been set.

### GetDaily

`func (o *CreateOperationRequest) GetDaily() int32`

GetDaily returns the Daily field if non-nil, zero value otherwise.

### GetDailyOk

`func (o *CreateOperationRequest) GetDailyOk() (*int32, bool)`

GetDailyOk returns a tuple with the Daily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaily

`func (o *CreateOperationRequest) SetDaily(v int32)`

SetDaily sets Daily field to given value.

### HasDaily

`func (o *CreateOperationRequest) HasDaily() bool`

HasDaily returns a boolean if a field has been set.

### GetMonthly

`func (o *CreateOperationRequest) GetMonthly() int32`

GetMonthly returns the Monthly field if non-nil, zero value otherwise.

### GetMonthlyOk

`func (o *CreateOperationRequest) GetMonthlyOk() (*int32, bool)`

GetMonthlyOk returns a tuple with the Monthly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthly

`func (o *CreateOperationRequest) SetMonthly(v int32)`

SetMonthly sets Monthly field to given value.

### HasMonthly

`func (o *CreateOperationRequest) HasMonthly() bool`

HasMonthly returns a boolean if a field has been set.

### GetYearly

`func (o *CreateOperationRequest) GetYearly() int32`

GetYearly returns the Yearly field if non-nil, zero value otherwise.

### GetYearlyOk

`func (o *CreateOperationRequest) GetYearlyOk() (*int32, bool)`

GetYearlyOk returns a tuple with the Yearly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYearly

`func (o *CreateOperationRequest) SetYearly(v int32)`

SetYearly sets Yearly field to given value.

### HasYearly

`func (o *CreateOperationRequest) HasYearly() bool`

HasYearly returns a boolean if a field has been set.

### GetPitrDays

`func (o *CreateOperationRequest) GetPitrDays() int32`

GetPitrDays returns the PitrDays field if non-nil, zero value otherwise.

### GetPitrDaysOk

`func (o *CreateOperationRequest) GetPitrDaysOk() (*int32, bool)`

GetPitrDaysOk returns a tuple with the PitrDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPitrDays

`func (o *CreateOperationRequest) SetPitrDays(v int32)`

SetPitrDays sets PitrDays field to given value.

### HasPitrDays

`func (o *CreateOperationRequest) HasPitrDays() bool`

HasPitrDays returns a boolean if a field has been set.

### GetMakeCopy

`func (o *CreateOperationRequest) GetMakeCopy() bool`

GetMakeCopy returns the MakeCopy field if non-nil, zero value otherwise.

### GetMakeCopyOk

`func (o *CreateOperationRequest) GetMakeCopyOk() (*bool, bool)`

GetMakeCopyOk returns a tuple with the MakeCopy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMakeCopy

`func (o *CreateOperationRequest) SetMakeCopy(v bool)`

SetMakeCopy sets MakeCopy field to given value.

### HasMakeCopy

`func (o *CreateOperationRequest) HasMakeCopy() bool`

HasMakeCopy returns a boolean if a field has been set.

### GetKeepFinal

`func (o *CreateOperationRequest) GetKeepFinal() bool`

GetKeepFinal returns the KeepFinal field if non-nil, zero value otherwise.

### GetKeepFinalOk

`func (o *CreateOperationRequest) GetKeepFinalOk() (*bool, bool)`

GetKeepFinalOk returns a tuple with the KeepFinal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepFinal

`func (o *CreateOperationRequest) SetKeepFinal(v bool)`

SetKeepFinal sets KeepFinal field to given value.

### HasKeepFinal

`func (o *CreateOperationRequest) HasKeepFinal() bool`

HasKeepFinal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


