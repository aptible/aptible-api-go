# CreateOperationRequest7

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

## Methods

### NewCreateOperationRequest7

`func NewCreateOperationRequest7(type_ string, ) *CreateOperationRequest7`

NewCreateOperationRequest7 instantiates a new CreateOperationRequest7 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOperationRequest7WithDefaults

`func NewCreateOperationRequest7WithDefaults() *CreateOperationRequest7`

NewCreateOperationRequest7WithDefaults instantiates a new CreateOperationRequest7 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CreateOperationRequest7) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateOperationRequest7) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateOperationRequest7) SetType(v string)`

SetType sets Type field to given value.


### GetGitRef

`func (o *CreateOperationRequest7) GetGitRef() string`

GetGitRef returns the GitRef field if non-nil, zero value otherwise.

### GetGitRefOk

`func (o *CreateOperationRequest7) GetGitRefOk() (*string, bool)`

GetGitRefOk returns a tuple with the GitRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitRef

`func (o *CreateOperationRequest7) SetGitRef(v string)`

SetGitRef sets GitRef field to given value.

### HasGitRef

`func (o *CreateOperationRequest7) HasGitRef() bool`

HasGitRef returns a boolean if a field has been set.

### GetDiskSize

`func (o *CreateOperationRequest7) GetDiskSize() int32`

GetDiskSize returns the DiskSize field if non-nil, zero value otherwise.

### GetDiskSizeOk

`func (o *CreateOperationRequest7) GetDiskSizeOk() (*int32, bool)`

GetDiskSizeOk returns a tuple with the DiskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSize

`func (o *CreateOperationRequest7) SetDiskSize(v int32)`

SetDiskSize sets DiskSize field to given value.

### HasDiskSize

`func (o *CreateOperationRequest7) HasDiskSize() bool`

HasDiskSize returns a boolean if a field has been set.

### GetContainerCount

`func (o *CreateOperationRequest7) GetContainerCount() int32`

GetContainerCount returns the ContainerCount field if non-nil, zero value otherwise.

### GetContainerCountOk

`func (o *CreateOperationRequest7) GetContainerCountOk() (*int32, bool)`

GetContainerCountOk returns a tuple with the ContainerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerCount

`func (o *CreateOperationRequest7) SetContainerCount(v int32)`

SetContainerCount sets ContainerCount field to given value.

### HasContainerCount

`func (o *CreateOperationRequest7) HasContainerCount() bool`

HasContainerCount returns a boolean if a field has been set.

### GetContainerSize

`func (o *CreateOperationRequest7) GetContainerSize() int32`

GetContainerSize returns the ContainerSize field if non-nil, zero value otherwise.

### GetContainerSizeOk

`func (o *CreateOperationRequest7) GetContainerSizeOk() (*int32, bool)`

GetContainerSizeOk returns a tuple with the ContainerSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerSize

`func (o *CreateOperationRequest7) SetContainerSize(v int32)`

SetContainerSize sets ContainerSize field to given value.

### HasContainerSize

`func (o *CreateOperationRequest7) HasContainerSize() bool`

HasContainerSize returns a boolean if a field has been set.

### GetCommand

`func (o *CreateOperationRequest7) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *CreateOperationRequest7) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *CreateOperationRequest7) SetCommand(v string)`

SetCommand sets Command field to given value.

### HasCommand

`func (o *CreateOperationRequest7) HasCommand() bool`

HasCommand returns a boolean if a field has been set.

### GetEnv

`func (o *CreateOperationRequest7) GetEnv() map[string]string`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *CreateOperationRequest7) GetEnvOk() (*map[string]string, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *CreateOperationRequest7) SetEnv(v map[string]string)`

SetEnv sets Env field to given value.

### HasEnv

`func (o *CreateOperationRequest7) HasEnv() bool`

HasEnv returns a boolean if a field has been set.

### GetHandle

`func (o *CreateOperationRequest7) GetHandle() string`

GetHandle returns the Handle field if non-nil, zero value otherwise.

### GetHandleOk

`func (o *CreateOperationRequest7) GetHandleOk() (*string, bool)`

GetHandleOk returns a tuple with the Handle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandle

`func (o *CreateOperationRequest7) SetHandle(v string)`

SetHandle sets Handle field to given value.

### HasHandle

`func (o *CreateOperationRequest7) HasHandle() bool`

HasHandle returns a boolean if a field has been set.

### GetCertificate

`func (o *CreateOperationRequest7) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *CreateOperationRequest7) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *CreateOperationRequest7) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *CreateOperationRequest7) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### GetPrivateKey

`func (o *CreateOperationRequest7) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *CreateOperationRequest7) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *CreateOperationRequest7) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.

### HasPrivateKey

`func (o *CreateOperationRequest7) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.

### GetDestinationRegion

`func (o *CreateOperationRequest7) GetDestinationRegion() string`

GetDestinationRegion returns the DestinationRegion field if non-nil, zero value otherwise.

### GetDestinationRegionOk

`func (o *CreateOperationRequest7) GetDestinationRegionOk() (*string, bool)`

GetDestinationRegionOk returns a tuple with the DestinationRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationRegion

`func (o *CreateOperationRequest7) SetDestinationRegion(v string)`

SetDestinationRegion sets DestinationRegion field to given value.

### HasDestinationRegion

`func (o *CreateOperationRequest7) HasDestinationRegion() bool`

HasDestinationRegion returns a boolean if a field has been set.

### GetInteractive

`func (o *CreateOperationRequest7) GetInteractive() bool`

GetInteractive returns the Interactive field if non-nil, zero value otherwise.

### GetInteractiveOk

`func (o *CreateOperationRequest7) GetInteractiveOk() (*bool, bool)`

GetInteractiveOk returns a tuple with the Interactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInteractive

`func (o *CreateOperationRequest7) SetInteractive(v bool)`

SetInteractive sets Interactive field to given value.

### HasInteractive

`func (o *CreateOperationRequest7) HasInteractive() bool`

HasInteractive returns a boolean if a field has been set.

### GetDestinationAccount

`func (o *CreateOperationRequest7) GetDestinationAccount() string`

GetDestinationAccount returns the DestinationAccount field if non-nil, zero value otherwise.

### GetDestinationAccountOk

`func (o *CreateOperationRequest7) GetDestinationAccountOk() (*string, bool)`

GetDestinationAccountOk returns a tuple with the DestinationAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationAccount

`func (o *CreateOperationRequest7) SetDestinationAccount(v string)`

SetDestinationAccount sets DestinationAccount field to given value.

### HasDestinationAccount

`func (o *CreateOperationRequest7) HasDestinationAccount() bool`

HasDestinationAccount returns a boolean if a field has been set.

### GetDestinationAccountId

`func (o *CreateOperationRequest7) GetDestinationAccountId() int32`

GetDestinationAccountId returns the DestinationAccountId field if non-nil, zero value otherwise.

### GetDestinationAccountIdOk

`func (o *CreateOperationRequest7) GetDestinationAccountIdOk() (*int32, bool)`

GetDestinationAccountIdOk returns a tuple with the DestinationAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationAccountId

`func (o *CreateOperationRequest7) SetDestinationAccountId(v int32)`

SetDestinationAccountId sets DestinationAccountId field to given value.

### HasDestinationAccountId

`func (o *CreateOperationRequest7) HasDestinationAccountId() bool`

HasDestinationAccountId returns a boolean if a field has been set.

### GetDockerRef

`func (o *CreateOperationRequest7) GetDockerRef() string`

GetDockerRef returns the DockerRef field if non-nil, zero value otherwise.

### GetDockerRefOk

`func (o *CreateOperationRequest7) GetDockerRefOk() (*string, bool)`

GetDockerRefOk returns a tuple with the DockerRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerRef

`func (o *CreateOperationRequest7) SetDockerRef(v string)`

SetDockerRef sets DockerRef field to given value.

### HasDockerRef

`func (o *CreateOperationRequest7) HasDockerRef() bool`

HasDockerRef returns a boolean if a field has been set.

### GetAutomated

`func (o *CreateOperationRequest7) GetAutomated() bool`

GetAutomated returns the Automated field if non-nil, zero value otherwise.

### GetAutomatedOk

`func (o *CreateOperationRequest7) GetAutomatedOk() (*bool, bool)`

GetAutomatedOk returns a tuple with the Automated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomated

`func (o *CreateOperationRequest7) SetAutomated(v bool)`

SetAutomated sets Automated field to given value.

### HasAutomated

`func (o *CreateOperationRequest7) HasAutomated() bool`

HasAutomated returns a boolean if a field has been set.

### GetKeyArn

`func (o *CreateOperationRequest7) GetKeyArn() string`

GetKeyArn returns the KeyArn field if non-nil, zero value otherwise.

### GetKeyArnOk

`func (o *CreateOperationRequest7) GetKeyArnOk() (*string, bool)`

GetKeyArnOk returns a tuple with the KeyArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyArn

`func (o *CreateOperationRequest7) SetKeyArn(v string)`

SetKeyArn sets KeyArn field to given value.

### HasKeyArn

`func (o *CreateOperationRequest7) HasKeyArn() bool`

HasKeyArn returns a boolean if a field has been set.

### GetInstanceProfile

`func (o *CreateOperationRequest7) GetInstanceProfile() string`

GetInstanceProfile returns the InstanceProfile field if non-nil, zero value otherwise.

### GetInstanceProfileOk

`func (o *CreateOperationRequest7) GetInstanceProfileOk() (*string, bool)`

GetInstanceProfileOk returns a tuple with the InstanceProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceProfile

`func (o *CreateOperationRequest7) SetInstanceProfile(v string)`

SetInstanceProfile sets InstanceProfile field to given value.

### HasInstanceProfile

`func (o *CreateOperationRequest7) HasInstanceProfile() bool`

HasInstanceProfile returns a boolean if a field has been set.

### GetProvisionedIops

`func (o *CreateOperationRequest7) GetProvisionedIops() int32`

GetProvisionedIops returns the ProvisionedIops field if non-nil, zero value otherwise.

### GetProvisionedIopsOk

`func (o *CreateOperationRequest7) GetProvisionedIopsOk() (*int32, bool)`

GetProvisionedIopsOk returns a tuple with the ProvisionedIops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionedIops

`func (o *CreateOperationRequest7) SetProvisionedIops(v int32)`

SetProvisionedIops sets ProvisionedIops field to given value.

### HasProvisionedIops

`func (o *CreateOperationRequest7) HasProvisionedIops() bool`

HasProvisionedIops returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


