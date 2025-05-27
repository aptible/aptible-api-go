# UpdateVhostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserDomain** | Pointer to **string** |  | [optional] 
**Certificate** | Pointer to **int32** |  | [optional] 
**Platform** | Pointer to **string** |  | [optional] 
**ContainerPort** | Pointer to **int32** |  | [optional] 
**ContainerPorts** | Pointer to **[]int32** |  | [optional] 
**IpWhitelist** | Pointer to **[]string** |  | [optional] 
**SharedFingerprint** | Pointer to **string** |  | [optional] 
**Shared** | Pointer to **bool** |  | [optional] 

## Methods

### NewUpdateVhostRequest

`func NewUpdateVhostRequest() *UpdateVhostRequest`

NewUpdateVhostRequest instantiates a new UpdateVhostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateVhostRequestWithDefaults

`func NewUpdateVhostRequestWithDefaults() *UpdateVhostRequest`

NewUpdateVhostRequestWithDefaults instantiates a new UpdateVhostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserDomain

`func (o *UpdateVhostRequest) GetUserDomain() string`

GetUserDomain returns the UserDomain field if non-nil, zero value otherwise.

### GetUserDomainOk

`func (o *UpdateVhostRequest) GetUserDomainOk() (*string, bool)`

GetUserDomainOk returns a tuple with the UserDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDomain

`func (o *UpdateVhostRequest) SetUserDomain(v string)`

SetUserDomain sets UserDomain field to given value.

### HasUserDomain

`func (o *UpdateVhostRequest) HasUserDomain() bool`

HasUserDomain returns a boolean if a field has been set.

### GetCertificate

`func (o *UpdateVhostRequest) GetCertificate() int32`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *UpdateVhostRequest) GetCertificateOk() (*int32, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *UpdateVhostRequest) SetCertificate(v int32)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *UpdateVhostRequest) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### GetPlatform

`func (o *UpdateVhostRequest) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *UpdateVhostRequest) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *UpdateVhostRequest) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *UpdateVhostRequest) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetContainerPort

`func (o *UpdateVhostRequest) GetContainerPort() int32`

GetContainerPort returns the ContainerPort field if non-nil, zero value otherwise.

### GetContainerPortOk

`func (o *UpdateVhostRequest) GetContainerPortOk() (*int32, bool)`

GetContainerPortOk returns a tuple with the ContainerPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerPort

`func (o *UpdateVhostRequest) SetContainerPort(v int32)`

SetContainerPort sets ContainerPort field to given value.

### HasContainerPort

`func (o *UpdateVhostRequest) HasContainerPort() bool`

HasContainerPort returns a boolean if a field has been set.

### GetContainerPorts

`func (o *UpdateVhostRequest) GetContainerPorts() []int32`

GetContainerPorts returns the ContainerPorts field if non-nil, zero value otherwise.

### GetContainerPortsOk

`func (o *UpdateVhostRequest) GetContainerPortsOk() (*[]int32, bool)`

GetContainerPortsOk returns a tuple with the ContainerPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerPorts

`func (o *UpdateVhostRequest) SetContainerPorts(v []int32)`

SetContainerPorts sets ContainerPorts field to given value.

### HasContainerPorts

`func (o *UpdateVhostRequest) HasContainerPorts() bool`

HasContainerPorts returns a boolean if a field has been set.

### GetIpWhitelist

`func (o *UpdateVhostRequest) GetIpWhitelist() []string`

GetIpWhitelist returns the IpWhitelist field if non-nil, zero value otherwise.

### GetIpWhitelistOk

`func (o *UpdateVhostRequest) GetIpWhitelistOk() (*[]string, bool)`

GetIpWhitelistOk returns a tuple with the IpWhitelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpWhitelist

`func (o *UpdateVhostRequest) SetIpWhitelist(v []string)`

SetIpWhitelist sets IpWhitelist field to given value.

### HasIpWhitelist

`func (o *UpdateVhostRequest) HasIpWhitelist() bool`

HasIpWhitelist returns a boolean if a field has been set.

### GetSharedFingerprint

`func (o *UpdateVhostRequest) GetSharedFingerprint() string`

GetSharedFingerprint returns the SharedFingerprint field if non-nil, zero value otherwise.

### GetSharedFingerprintOk

`func (o *UpdateVhostRequest) GetSharedFingerprintOk() (*string, bool)`

GetSharedFingerprintOk returns a tuple with the SharedFingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedFingerprint

`func (o *UpdateVhostRequest) SetSharedFingerprint(v string)`

SetSharedFingerprint sets SharedFingerprint field to given value.

### HasSharedFingerprint

`func (o *UpdateVhostRequest) HasSharedFingerprint() bool`

HasSharedFingerprint returns a boolean if a field has been set.

### GetShared

`func (o *UpdateVhostRequest) GetShared() bool`

GetShared returns the Shared field if non-nil, zero value otherwise.

### GetSharedOk

`func (o *UpdateVhostRequest) GetSharedOk() (*bool, bool)`

GetSharedOk returns a tuple with the Shared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShared

`func (o *UpdateVhostRequest) SetShared(v bool)`

SetShared sets Shared field to given value.

### HasShared

`func (o *UpdateVhostRequest) HasShared() bool`

HasShared returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


