# CreateVhostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Default** | Pointer to **bool** |  | [optional] 
**Internal** | Pointer to **bool** |  | [optional] 
**Certificate** | Pointer to **int32** |  | [optional] 
**UserDomain** | Pointer to **string** |  | [optional] 
**Acme** | Pointer to **bool** |  | [optional] 
**Platform** | Pointer to **string** |  | [optional] 
**ContainerPort** | Pointer to **int32** |  | [optional] 
**ContainerPorts** | Pointer to **[]int32** |  | [optional] 
**IpWhitelist** | Pointer to **[]string** |  | [optional] 
**ContainerExposedPorts** | Pointer to **[]int32** |  | [optional] 
**Shared** | Pointer to **bool** |  | [optional] 

## Methods

### NewCreateVhostRequest

`func NewCreateVhostRequest(type_ string, ) *CreateVhostRequest`

NewCreateVhostRequest instantiates a new CreateVhostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateVhostRequestWithDefaults

`func NewCreateVhostRequestWithDefaults() *CreateVhostRequest`

NewCreateVhostRequestWithDefaults instantiates a new CreateVhostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CreateVhostRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateVhostRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateVhostRequest) SetType(v string)`

SetType sets Type field to given value.


### GetDefault

`func (o *CreateVhostRequest) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *CreateVhostRequest) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *CreateVhostRequest) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *CreateVhostRequest) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetInternal

`func (o *CreateVhostRequest) GetInternal() bool`

GetInternal returns the Internal field if non-nil, zero value otherwise.

### GetInternalOk

`func (o *CreateVhostRequest) GetInternalOk() (*bool, bool)`

GetInternalOk returns a tuple with the Internal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternal

`func (o *CreateVhostRequest) SetInternal(v bool)`

SetInternal sets Internal field to given value.

### HasInternal

`func (o *CreateVhostRequest) HasInternal() bool`

HasInternal returns a boolean if a field has been set.

### GetCertificate

`func (o *CreateVhostRequest) GetCertificate() int32`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *CreateVhostRequest) GetCertificateOk() (*int32, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *CreateVhostRequest) SetCertificate(v int32)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *CreateVhostRequest) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### GetUserDomain

`func (o *CreateVhostRequest) GetUserDomain() string`

GetUserDomain returns the UserDomain field if non-nil, zero value otherwise.

### GetUserDomainOk

`func (o *CreateVhostRequest) GetUserDomainOk() (*string, bool)`

GetUserDomainOk returns a tuple with the UserDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDomain

`func (o *CreateVhostRequest) SetUserDomain(v string)`

SetUserDomain sets UserDomain field to given value.

### HasUserDomain

`func (o *CreateVhostRequest) HasUserDomain() bool`

HasUserDomain returns a boolean if a field has been set.

### GetAcme

`func (o *CreateVhostRequest) GetAcme() bool`

GetAcme returns the Acme field if non-nil, zero value otherwise.

### GetAcmeOk

`func (o *CreateVhostRequest) GetAcmeOk() (*bool, bool)`

GetAcmeOk returns a tuple with the Acme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcme

`func (o *CreateVhostRequest) SetAcme(v bool)`

SetAcme sets Acme field to given value.

### HasAcme

`func (o *CreateVhostRequest) HasAcme() bool`

HasAcme returns a boolean if a field has been set.

### GetPlatform

`func (o *CreateVhostRequest) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *CreateVhostRequest) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *CreateVhostRequest) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *CreateVhostRequest) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetContainerPort

`func (o *CreateVhostRequest) GetContainerPort() int32`

GetContainerPort returns the ContainerPort field if non-nil, zero value otherwise.

### GetContainerPortOk

`func (o *CreateVhostRequest) GetContainerPortOk() (*int32, bool)`

GetContainerPortOk returns a tuple with the ContainerPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerPort

`func (o *CreateVhostRequest) SetContainerPort(v int32)`

SetContainerPort sets ContainerPort field to given value.

### HasContainerPort

`func (o *CreateVhostRequest) HasContainerPort() bool`

HasContainerPort returns a boolean if a field has been set.

### GetContainerPorts

`func (o *CreateVhostRequest) GetContainerPorts() []int32`

GetContainerPorts returns the ContainerPorts field if non-nil, zero value otherwise.

### GetContainerPortsOk

`func (o *CreateVhostRequest) GetContainerPortsOk() (*[]int32, bool)`

GetContainerPortsOk returns a tuple with the ContainerPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerPorts

`func (o *CreateVhostRequest) SetContainerPorts(v []int32)`

SetContainerPorts sets ContainerPorts field to given value.

### HasContainerPorts

`func (o *CreateVhostRequest) HasContainerPorts() bool`

HasContainerPorts returns a boolean if a field has been set.

### GetIpWhitelist

`func (o *CreateVhostRequest) GetIpWhitelist() []string`

GetIpWhitelist returns the IpWhitelist field if non-nil, zero value otherwise.

### GetIpWhitelistOk

`func (o *CreateVhostRequest) GetIpWhitelistOk() (*[]string, bool)`

GetIpWhitelistOk returns a tuple with the IpWhitelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpWhitelist

`func (o *CreateVhostRequest) SetIpWhitelist(v []string)`

SetIpWhitelist sets IpWhitelist field to given value.

### HasIpWhitelist

`func (o *CreateVhostRequest) HasIpWhitelist() bool`

HasIpWhitelist returns a boolean if a field has been set.

### GetContainerExposedPorts

`func (o *CreateVhostRequest) GetContainerExposedPorts() []int32`

GetContainerExposedPorts returns the ContainerExposedPorts field if non-nil, zero value otherwise.

### GetContainerExposedPortsOk

`func (o *CreateVhostRequest) GetContainerExposedPortsOk() (*[]int32, bool)`

GetContainerExposedPortsOk returns a tuple with the ContainerExposedPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerExposedPorts

`func (o *CreateVhostRequest) SetContainerExposedPorts(v []int32)`

SetContainerExposedPorts sets ContainerExposedPorts field to given value.

### HasContainerExposedPorts

`func (o *CreateVhostRequest) HasContainerExposedPorts() bool`

HasContainerExposedPorts returns a boolean if a field has been set.

### GetShared

`func (o *CreateVhostRequest) GetShared() bool`

GetShared returns the Shared field if non-nil, zero value otherwise.

### GetSharedOk

`func (o *CreateVhostRequest) GetSharedOk() (*bool, bool)`

GetSharedOk returns a tuple with the Shared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShared

`func (o *CreateVhostRequest) SetShared(v bool)`

SetShared sets Shared field to given value.

### HasShared

`func (o *CreateVhostRequest) HasShared() bool`

HasShared returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


