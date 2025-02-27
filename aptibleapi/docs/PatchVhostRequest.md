# PatchVhostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserDomain** | Pointer to **string** |  | [optional] 
**Certificate** | Pointer to **int32** |  | [optional] 
**Platform** | Pointer to **string** |  | [optional] 
**ContainerPort** | Pointer to **int32** |  | [optional] 
**ContainerPorts** | Pointer to **[]int32** |  | [optional] 
**IpWhitelist** | Pointer to **[]string** |  | [optional] 

## Methods

### NewPatchVhostRequest

`func NewPatchVhostRequest() *PatchVhostRequest`

NewPatchVhostRequest instantiates a new PatchVhostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchVhostRequestWithDefaults

`func NewPatchVhostRequestWithDefaults() *PatchVhostRequest`

NewPatchVhostRequestWithDefaults instantiates a new PatchVhostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserDomain

`func (o *PatchVhostRequest) GetUserDomain() string`

GetUserDomain returns the UserDomain field if non-nil, zero value otherwise.

### GetUserDomainOk

`func (o *PatchVhostRequest) GetUserDomainOk() (*string, bool)`

GetUserDomainOk returns a tuple with the UserDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDomain

`func (o *PatchVhostRequest) SetUserDomain(v string)`

SetUserDomain sets UserDomain field to given value.

### HasUserDomain

`func (o *PatchVhostRequest) HasUserDomain() bool`

HasUserDomain returns a boolean if a field has been set.

### GetCertificate

`func (o *PatchVhostRequest) GetCertificate() int32`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *PatchVhostRequest) GetCertificateOk() (*int32, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *PatchVhostRequest) SetCertificate(v int32)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *PatchVhostRequest) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### GetPlatform

`func (o *PatchVhostRequest) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *PatchVhostRequest) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *PatchVhostRequest) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *PatchVhostRequest) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetContainerPort

`func (o *PatchVhostRequest) GetContainerPort() int32`

GetContainerPort returns the ContainerPort field if non-nil, zero value otherwise.

### GetContainerPortOk

`func (o *PatchVhostRequest) GetContainerPortOk() (*int32, bool)`

GetContainerPortOk returns a tuple with the ContainerPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerPort

`func (o *PatchVhostRequest) SetContainerPort(v int32)`

SetContainerPort sets ContainerPort field to given value.

### HasContainerPort

`func (o *PatchVhostRequest) HasContainerPort() bool`

HasContainerPort returns a boolean if a field has been set.

### GetContainerPorts

`func (o *PatchVhostRequest) GetContainerPorts() []int32`

GetContainerPorts returns the ContainerPorts field if non-nil, zero value otherwise.

### GetContainerPortsOk

`func (o *PatchVhostRequest) GetContainerPortsOk() (*[]int32, bool)`

GetContainerPortsOk returns a tuple with the ContainerPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerPorts

`func (o *PatchVhostRequest) SetContainerPorts(v []int32)`

SetContainerPorts sets ContainerPorts field to given value.

### HasContainerPorts

`func (o *PatchVhostRequest) HasContainerPorts() bool`

HasContainerPorts returns a boolean if a field has been set.

### GetIpWhitelist

`func (o *PatchVhostRequest) GetIpWhitelist() []string`

GetIpWhitelist returns the IpWhitelist field if non-nil, zero value otherwise.

### GetIpWhitelistOk

`func (o *PatchVhostRequest) GetIpWhitelistOk() (*[]string, bool)`

GetIpWhitelistOk returns a tuple with the IpWhitelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpWhitelist

`func (o *PatchVhostRequest) SetIpWhitelist(v []string)`

SetIpWhitelist sets IpWhitelist field to given value.

### HasIpWhitelist

`func (o *PatchVhostRequest) HasIpWhitelist() bool`

HasIpWhitelist returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


