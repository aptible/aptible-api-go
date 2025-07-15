# Vhost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**MetaType** | **string** |  | 
**VirtualDomain** | **string** |  | 
**Type** | **string** |  | 
**ElasticLoadBalancerName** | **NullableString** |  | 
**ApplicationLoadBalancerArn** | **NullableString** |  | 
**SecurityGroupId** | **NullableString** |  | 
**ExternalHost** | **NullableString** |  | 
**ExternalHttpPort** | **NullableInt32** |  | 
**ExternalHttpsPort** | **NullableInt32** |  | 
**InternalHost** | **NullableString** |  | 
**InternalHttpPort** | **NullableInt32** |  | 
**InternalHttpsPort** | **NullableInt32** |  | 
**InternalHealthPort** | **NullableInt32** |  | 
**DockerName** | **NullableString** |  | 
**CreatedAt** | **string** |  | 
**UpdatedAt** | **string** |  | 
**Status** | **string** |  | 
**Platform** | **string** |  | 
**Default** | **bool** |  | 
**Internal** | **bool** |  | 
**ContainerExposedPorts** | **[]int32** |  | 
**HostMappedPorts** | **[]int32** |  | 
**IpWhitelist** | **[]string** |  | 
**UserDomain** | **NullableString** |  | 
**Acme** | **bool** |  | 
**AcmeStatus** | **NullableString** |  | 
**AcmeDnsChallengeHost** | **NullableString** |  | 
**ContainerPort** | **NullableInt32** |  | 
**ContainerPorts** | **[]int32** |  | 
**AcmeConfiguration** | [**NullableVhostAcmeConfiguration**](VhostAcmeConfiguration.md) |  | 
**Shared** | **NullableBool** |  | 
**SharedFingerprint** | **NullableString** |  | 
**LoadBalancingAlgorithmType** | **NullableString** |  | 
**Links** | Pointer to [**VhostLinks**](VhostLinks.md) |  | [optional] 

## Methods

### NewVhost

`func NewVhost(id int32, metaType string, virtualDomain string, type_ string, elasticLoadBalancerName NullableString, applicationLoadBalancerArn NullableString, securityGroupId NullableString, externalHost NullableString, externalHttpPort NullableInt32, externalHttpsPort NullableInt32, internalHost NullableString, internalHttpPort NullableInt32, internalHttpsPort NullableInt32, internalHealthPort NullableInt32, dockerName NullableString, createdAt string, updatedAt string, status string, platform string, default_ bool, internal bool, containerExposedPorts []int32, hostMappedPorts []int32, ipWhitelist []string, userDomain NullableString, acme bool, acmeStatus NullableString, acmeDnsChallengeHost NullableString, containerPort NullableInt32, containerPorts []int32, acmeConfiguration NullableVhostAcmeConfiguration, shared NullableBool, sharedFingerprint NullableString, loadBalancingAlgorithmType NullableString, ) *Vhost`

NewVhost instantiates a new Vhost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVhostWithDefaults

`func NewVhostWithDefaults() *Vhost`

NewVhostWithDefaults instantiates a new Vhost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Vhost) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Vhost) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Vhost) SetId(v int32)`

SetId sets Id field to given value.


### GetMetaType

`func (o *Vhost) GetMetaType() string`

GetMetaType returns the MetaType field if non-nil, zero value otherwise.

### GetMetaTypeOk

`func (o *Vhost) GetMetaTypeOk() (*string, bool)`

GetMetaTypeOk returns a tuple with the MetaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaType

`func (o *Vhost) SetMetaType(v string)`

SetMetaType sets MetaType field to given value.


### GetVirtualDomain

`func (o *Vhost) GetVirtualDomain() string`

GetVirtualDomain returns the VirtualDomain field if non-nil, zero value otherwise.

### GetVirtualDomainOk

`func (o *Vhost) GetVirtualDomainOk() (*string, bool)`

GetVirtualDomainOk returns a tuple with the VirtualDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualDomain

`func (o *Vhost) SetVirtualDomain(v string)`

SetVirtualDomain sets VirtualDomain field to given value.


### GetType

`func (o *Vhost) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Vhost) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Vhost) SetType(v string)`

SetType sets Type field to given value.


### GetElasticLoadBalancerName

`func (o *Vhost) GetElasticLoadBalancerName() string`

GetElasticLoadBalancerName returns the ElasticLoadBalancerName field if non-nil, zero value otherwise.

### GetElasticLoadBalancerNameOk

`func (o *Vhost) GetElasticLoadBalancerNameOk() (*string, bool)`

GetElasticLoadBalancerNameOk returns a tuple with the ElasticLoadBalancerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElasticLoadBalancerName

`func (o *Vhost) SetElasticLoadBalancerName(v string)`

SetElasticLoadBalancerName sets ElasticLoadBalancerName field to given value.


### SetElasticLoadBalancerNameNil

`func (o *Vhost) SetElasticLoadBalancerNameNil(b bool)`

 SetElasticLoadBalancerNameNil sets the value for ElasticLoadBalancerName to be an explicit nil

### UnsetElasticLoadBalancerName
`func (o *Vhost) UnsetElasticLoadBalancerName()`

UnsetElasticLoadBalancerName ensures that no value is present for ElasticLoadBalancerName, not even an explicit nil
### GetApplicationLoadBalancerArn

`func (o *Vhost) GetApplicationLoadBalancerArn() string`

GetApplicationLoadBalancerArn returns the ApplicationLoadBalancerArn field if non-nil, zero value otherwise.

### GetApplicationLoadBalancerArnOk

`func (o *Vhost) GetApplicationLoadBalancerArnOk() (*string, bool)`

GetApplicationLoadBalancerArnOk returns a tuple with the ApplicationLoadBalancerArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationLoadBalancerArn

`func (o *Vhost) SetApplicationLoadBalancerArn(v string)`

SetApplicationLoadBalancerArn sets ApplicationLoadBalancerArn field to given value.


### SetApplicationLoadBalancerArnNil

`func (o *Vhost) SetApplicationLoadBalancerArnNil(b bool)`

 SetApplicationLoadBalancerArnNil sets the value for ApplicationLoadBalancerArn to be an explicit nil

### UnsetApplicationLoadBalancerArn
`func (o *Vhost) UnsetApplicationLoadBalancerArn()`

UnsetApplicationLoadBalancerArn ensures that no value is present for ApplicationLoadBalancerArn, not even an explicit nil
### GetSecurityGroupId

`func (o *Vhost) GetSecurityGroupId() string`

GetSecurityGroupId returns the SecurityGroupId field if non-nil, zero value otherwise.

### GetSecurityGroupIdOk

`func (o *Vhost) GetSecurityGroupIdOk() (*string, bool)`

GetSecurityGroupIdOk returns a tuple with the SecurityGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroupId

`func (o *Vhost) SetSecurityGroupId(v string)`

SetSecurityGroupId sets SecurityGroupId field to given value.


### SetSecurityGroupIdNil

`func (o *Vhost) SetSecurityGroupIdNil(b bool)`

 SetSecurityGroupIdNil sets the value for SecurityGroupId to be an explicit nil

### UnsetSecurityGroupId
`func (o *Vhost) UnsetSecurityGroupId()`

UnsetSecurityGroupId ensures that no value is present for SecurityGroupId, not even an explicit nil
### GetExternalHost

`func (o *Vhost) GetExternalHost() string`

GetExternalHost returns the ExternalHost field if non-nil, zero value otherwise.

### GetExternalHostOk

`func (o *Vhost) GetExternalHostOk() (*string, bool)`

GetExternalHostOk returns a tuple with the ExternalHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalHost

`func (o *Vhost) SetExternalHost(v string)`

SetExternalHost sets ExternalHost field to given value.


### SetExternalHostNil

`func (o *Vhost) SetExternalHostNil(b bool)`

 SetExternalHostNil sets the value for ExternalHost to be an explicit nil

### UnsetExternalHost
`func (o *Vhost) UnsetExternalHost()`

UnsetExternalHost ensures that no value is present for ExternalHost, not even an explicit nil
### GetExternalHttpPort

`func (o *Vhost) GetExternalHttpPort() int32`

GetExternalHttpPort returns the ExternalHttpPort field if non-nil, zero value otherwise.

### GetExternalHttpPortOk

`func (o *Vhost) GetExternalHttpPortOk() (*int32, bool)`

GetExternalHttpPortOk returns a tuple with the ExternalHttpPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalHttpPort

`func (o *Vhost) SetExternalHttpPort(v int32)`

SetExternalHttpPort sets ExternalHttpPort field to given value.


### SetExternalHttpPortNil

`func (o *Vhost) SetExternalHttpPortNil(b bool)`

 SetExternalHttpPortNil sets the value for ExternalHttpPort to be an explicit nil

### UnsetExternalHttpPort
`func (o *Vhost) UnsetExternalHttpPort()`

UnsetExternalHttpPort ensures that no value is present for ExternalHttpPort, not even an explicit nil
### GetExternalHttpsPort

`func (o *Vhost) GetExternalHttpsPort() int32`

GetExternalHttpsPort returns the ExternalHttpsPort field if non-nil, zero value otherwise.

### GetExternalHttpsPortOk

`func (o *Vhost) GetExternalHttpsPortOk() (*int32, bool)`

GetExternalHttpsPortOk returns a tuple with the ExternalHttpsPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalHttpsPort

`func (o *Vhost) SetExternalHttpsPort(v int32)`

SetExternalHttpsPort sets ExternalHttpsPort field to given value.


### SetExternalHttpsPortNil

`func (o *Vhost) SetExternalHttpsPortNil(b bool)`

 SetExternalHttpsPortNil sets the value for ExternalHttpsPort to be an explicit nil

### UnsetExternalHttpsPort
`func (o *Vhost) UnsetExternalHttpsPort()`

UnsetExternalHttpsPort ensures that no value is present for ExternalHttpsPort, not even an explicit nil
### GetInternalHost

`func (o *Vhost) GetInternalHost() string`

GetInternalHost returns the InternalHost field if non-nil, zero value otherwise.

### GetInternalHostOk

`func (o *Vhost) GetInternalHostOk() (*string, bool)`

GetInternalHostOk returns a tuple with the InternalHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalHost

`func (o *Vhost) SetInternalHost(v string)`

SetInternalHost sets InternalHost field to given value.


### SetInternalHostNil

`func (o *Vhost) SetInternalHostNil(b bool)`

 SetInternalHostNil sets the value for InternalHost to be an explicit nil

### UnsetInternalHost
`func (o *Vhost) UnsetInternalHost()`

UnsetInternalHost ensures that no value is present for InternalHost, not even an explicit nil
### GetInternalHttpPort

`func (o *Vhost) GetInternalHttpPort() int32`

GetInternalHttpPort returns the InternalHttpPort field if non-nil, zero value otherwise.

### GetInternalHttpPortOk

`func (o *Vhost) GetInternalHttpPortOk() (*int32, bool)`

GetInternalHttpPortOk returns a tuple with the InternalHttpPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalHttpPort

`func (o *Vhost) SetInternalHttpPort(v int32)`

SetInternalHttpPort sets InternalHttpPort field to given value.


### SetInternalHttpPortNil

`func (o *Vhost) SetInternalHttpPortNil(b bool)`

 SetInternalHttpPortNil sets the value for InternalHttpPort to be an explicit nil

### UnsetInternalHttpPort
`func (o *Vhost) UnsetInternalHttpPort()`

UnsetInternalHttpPort ensures that no value is present for InternalHttpPort, not even an explicit nil
### GetInternalHttpsPort

`func (o *Vhost) GetInternalHttpsPort() int32`

GetInternalHttpsPort returns the InternalHttpsPort field if non-nil, zero value otherwise.

### GetInternalHttpsPortOk

`func (o *Vhost) GetInternalHttpsPortOk() (*int32, bool)`

GetInternalHttpsPortOk returns a tuple with the InternalHttpsPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalHttpsPort

`func (o *Vhost) SetInternalHttpsPort(v int32)`

SetInternalHttpsPort sets InternalHttpsPort field to given value.


### SetInternalHttpsPortNil

`func (o *Vhost) SetInternalHttpsPortNil(b bool)`

 SetInternalHttpsPortNil sets the value for InternalHttpsPort to be an explicit nil

### UnsetInternalHttpsPort
`func (o *Vhost) UnsetInternalHttpsPort()`

UnsetInternalHttpsPort ensures that no value is present for InternalHttpsPort, not even an explicit nil
### GetInternalHealthPort

`func (o *Vhost) GetInternalHealthPort() int32`

GetInternalHealthPort returns the InternalHealthPort field if non-nil, zero value otherwise.

### GetInternalHealthPortOk

`func (o *Vhost) GetInternalHealthPortOk() (*int32, bool)`

GetInternalHealthPortOk returns a tuple with the InternalHealthPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalHealthPort

`func (o *Vhost) SetInternalHealthPort(v int32)`

SetInternalHealthPort sets InternalHealthPort field to given value.


### SetInternalHealthPortNil

`func (o *Vhost) SetInternalHealthPortNil(b bool)`

 SetInternalHealthPortNil sets the value for InternalHealthPort to be an explicit nil

### UnsetInternalHealthPort
`func (o *Vhost) UnsetInternalHealthPort()`

UnsetInternalHealthPort ensures that no value is present for InternalHealthPort, not even an explicit nil
### GetDockerName

`func (o *Vhost) GetDockerName() string`

GetDockerName returns the DockerName field if non-nil, zero value otherwise.

### GetDockerNameOk

`func (o *Vhost) GetDockerNameOk() (*string, bool)`

GetDockerNameOk returns a tuple with the DockerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerName

`func (o *Vhost) SetDockerName(v string)`

SetDockerName sets DockerName field to given value.


### SetDockerNameNil

`func (o *Vhost) SetDockerNameNil(b bool)`

 SetDockerNameNil sets the value for DockerName to be an explicit nil

### UnsetDockerName
`func (o *Vhost) UnsetDockerName()`

UnsetDockerName ensures that no value is present for DockerName, not even an explicit nil
### GetCreatedAt

`func (o *Vhost) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Vhost) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Vhost) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Vhost) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Vhost) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Vhost) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetStatus

`func (o *Vhost) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Vhost) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Vhost) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetPlatform

`func (o *Vhost) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *Vhost) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *Vhost) SetPlatform(v string)`

SetPlatform sets Platform field to given value.


### GetDefault

`func (o *Vhost) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *Vhost) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *Vhost) SetDefault(v bool)`

SetDefault sets Default field to given value.


### GetInternal

`func (o *Vhost) GetInternal() bool`

GetInternal returns the Internal field if non-nil, zero value otherwise.

### GetInternalOk

`func (o *Vhost) GetInternalOk() (*bool, bool)`

GetInternalOk returns a tuple with the Internal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternal

`func (o *Vhost) SetInternal(v bool)`

SetInternal sets Internal field to given value.


### GetContainerExposedPorts

`func (o *Vhost) GetContainerExposedPorts() []int32`

GetContainerExposedPorts returns the ContainerExposedPorts field if non-nil, zero value otherwise.

### GetContainerExposedPortsOk

`func (o *Vhost) GetContainerExposedPortsOk() (*[]int32, bool)`

GetContainerExposedPortsOk returns a tuple with the ContainerExposedPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerExposedPorts

`func (o *Vhost) SetContainerExposedPorts(v []int32)`

SetContainerExposedPorts sets ContainerExposedPorts field to given value.


### SetContainerExposedPortsNil

`func (o *Vhost) SetContainerExposedPortsNil(b bool)`

 SetContainerExposedPortsNil sets the value for ContainerExposedPorts to be an explicit nil

### UnsetContainerExposedPorts
`func (o *Vhost) UnsetContainerExposedPorts()`

UnsetContainerExposedPorts ensures that no value is present for ContainerExposedPorts, not even an explicit nil
### GetHostMappedPorts

`func (o *Vhost) GetHostMappedPorts() []int32`

GetHostMappedPorts returns the HostMappedPorts field if non-nil, zero value otherwise.

### GetHostMappedPortsOk

`func (o *Vhost) GetHostMappedPortsOk() (*[]int32, bool)`

GetHostMappedPortsOk returns a tuple with the HostMappedPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostMappedPorts

`func (o *Vhost) SetHostMappedPorts(v []int32)`

SetHostMappedPorts sets HostMappedPorts field to given value.


### SetHostMappedPortsNil

`func (o *Vhost) SetHostMappedPortsNil(b bool)`

 SetHostMappedPortsNil sets the value for HostMappedPorts to be an explicit nil

### UnsetHostMappedPorts
`func (o *Vhost) UnsetHostMappedPorts()`

UnsetHostMappedPorts ensures that no value is present for HostMappedPorts, not even an explicit nil
### GetIpWhitelist

`func (o *Vhost) GetIpWhitelist() []string`

GetIpWhitelist returns the IpWhitelist field if non-nil, zero value otherwise.

### GetIpWhitelistOk

`func (o *Vhost) GetIpWhitelistOk() (*[]string, bool)`

GetIpWhitelistOk returns a tuple with the IpWhitelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpWhitelist

`func (o *Vhost) SetIpWhitelist(v []string)`

SetIpWhitelist sets IpWhitelist field to given value.


### GetUserDomain

`func (o *Vhost) GetUserDomain() string`

GetUserDomain returns the UserDomain field if non-nil, zero value otherwise.

### GetUserDomainOk

`func (o *Vhost) GetUserDomainOk() (*string, bool)`

GetUserDomainOk returns a tuple with the UserDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDomain

`func (o *Vhost) SetUserDomain(v string)`

SetUserDomain sets UserDomain field to given value.


### SetUserDomainNil

`func (o *Vhost) SetUserDomainNil(b bool)`

 SetUserDomainNil sets the value for UserDomain to be an explicit nil

### UnsetUserDomain
`func (o *Vhost) UnsetUserDomain()`

UnsetUserDomain ensures that no value is present for UserDomain, not even an explicit nil
### GetAcme

`func (o *Vhost) GetAcme() bool`

GetAcme returns the Acme field if non-nil, zero value otherwise.

### GetAcmeOk

`func (o *Vhost) GetAcmeOk() (*bool, bool)`

GetAcmeOk returns a tuple with the Acme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcme

`func (o *Vhost) SetAcme(v bool)`

SetAcme sets Acme field to given value.


### GetAcmeStatus

`func (o *Vhost) GetAcmeStatus() string`

GetAcmeStatus returns the AcmeStatus field if non-nil, zero value otherwise.

### GetAcmeStatusOk

`func (o *Vhost) GetAcmeStatusOk() (*string, bool)`

GetAcmeStatusOk returns a tuple with the AcmeStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcmeStatus

`func (o *Vhost) SetAcmeStatus(v string)`

SetAcmeStatus sets AcmeStatus field to given value.


### SetAcmeStatusNil

`func (o *Vhost) SetAcmeStatusNil(b bool)`

 SetAcmeStatusNil sets the value for AcmeStatus to be an explicit nil

### UnsetAcmeStatus
`func (o *Vhost) UnsetAcmeStatus()`

UnsetAcmeStatus ensures that no value is present for AcmeStatus, not even an explicit nil
### GetAcmeDnsChallengeHost

`func (o *Vhost) GetAcmeDnsChallengeHost() string`

GetAcmeDnsChallengeHost returns the AcmeDnsChallengeHost field if non-nil, zero value otherwise.

### GetAcmeDnsChallengeHostOk

`func (o *Vhost) GetAcmeDnsChallengeHostOk() (*string, bool)`

GetAcmeDnsChallengeHostOk returns a tuple with the AcmeDnsChallengeHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcmeDnsChallengeHost

`func (o *Vhost) SetAcmeDnsChallengeHost(v string)`

SetAcmeDnsChallengeHost sets AcmeDnsChallengeHost field to given value.


### SetAcmeDnsChallengeHostNil

`func (o *Vhost) SetAcmeDnsChallengeHostNil(b bool)`

 SetAcmeDnsChallengeHostNil sets the value for AcmeDnsChallengeHost to be an explicit nil

### UnsetAcmeDnsChallengeHost
`func (o *Vhost) UnsetAcmeDnsChallengeHost()`

UnsetAcmeDnsChallengeHost ensures that no value is present for AcmeDnsChallengeHost, not even an explicit nil
### GetContainerPort

`func (o *Vhost) GetContainerPort() int32`

GetContainerPort returns the ContainerPort field if non-nil, zero value otherwise.

### GetContainerPortOk

`func (o *Vhost) GetContainerPortOk() (*int32, bool)`

GetContainerPortOk returns a tuple with the ContainerPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerPort

`func (o *Vhost) SetContainerPort(v int32)`

SetContainerPort sets ContainerPort field to given value.


### SetContainerPortNil

`func (o *Vhost) SetContainerPortNil(b bool)`

 SetContainerPortNil sets the value for ContainerPort to be an explicit nil

### UnsetContainerPort
`func (o *Vhost) UnsetContainerPort()`

UnsetContainerPort ensures that no value is present for ContainerPort, not even an explicit nil
### GetContainerPorts

`func (o *Vhost) GetContainerPorts() []int32`

GetContainerPorts returns the ContainerPorts field if non-nil, zero value otherwise.

### GetContainerPortsOk

`func (o *Vhost) GetContainerPortsOk() (*[]int32, bool)`

GetContainerPortsOk returns a tuple with the ContainerPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerPorts

`func (o *Vhost) SetContainerPorts(v []int32)`

SetContainerPorts sets ContainerPorts field to given value.


### GetAcmeConfiguration

`func (o *Vhost) GetAcmeConfiguration() VhostAcmeConfiguration`

GetAcmeConfiguration returns the AcmeConfiguration field if non-nil, zero value otherwise.

### GetAcmeConfigurationOk

`func (o *Vhost) GetAcmeConfigurationOk() (*VhostAcmeConfiguration, bool)`

GetAcmeConfigurationOk returns a tuple with the AcmeConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcmeConfiguration

`func (o *Vhost) SetAcmeConfiguration(v VhostAcmeConfiguration)`

SetAcmeConfiguration sets AcmeConfiguration field to given value.


### SetAcmeConfigurationNil

`func (o *Vhost) SetAcmeConfigurationNil(b bool)`

 SetAcmeConfigurationNil sets the value for AcmeConfiguration to be an explicit nil

### UnsetAcmeConfiguration
`func (o *Vhost) UnsetAcmeConfiguration()`

UnsetAcmeConfiguration ensures that no value is present for AcmeConfiguration, not even an explicit nil
### GetShared

`func (o *Vhost) GetShared() bool`

GetShared returns the Shared field if non-nil, zero value otherwise.

### GetSharedOk

`func (o *Vhost) GetSharedOk() (*bool, bool)`

GetSharedOk returns a tuple with the Shared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShared

`func (o *Vhost) SetShared(v bool)`

SetShared sets Shared field to given value.


### SetSharedNil

`func (o *Vhost) SetSharedNil(b bool)`

 SetSharedNil sets the value for Shared to be an explicit nil

### UnsetShared
`func (o *Vhost) UnsetShared()`

UnsetShared ensures that no value is present for Shared, not even an explicit nil
### GetSharedFingerprint

`func (o *Vhost) GetSharedFingerprint() string`

GetSharedFingerprint returns the SharedFingerprint field if non-nil, zero value otherwise.

### GetSharedFingerprintOk

`func (o *Vhost) GetSharedFingerprintOk() (*string, bool)`

GetSharedFingerprintOk returns a tuple with the SharedFingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedFingerprint

`func (o *Vhost) SetSharedFingerprint(v string)`

SetSharedFingerprint sets SharedFingerprint field to given value.


### SetSharedFingerprintNil

`func (o *Vhost) SetSharedFingerprintNil(b bool)`

 SetSharedFingerprintNil sets the value for SharedFingerprint to be an explicit nil

### UnsetSharedFingerprint
`func (o *Vhost) UnsetSharedFingerprint()`

UnsetSharedFingerprint ensures that no value is present for SharedFingerprint, not even an explicit nil
### GetLoadBalancingAlgorithmType

`func (o *Vhost) GetLoadBalancingAlgorithmType() string`

GetLoadBalancingAlgorithmType returns the LoadBalancingAlgorithmType field if non-nil, zero value otherwise.

### GetLoadBalancingAlgorithmTypeOk

`func (o *Vhost) GetLoadBalancingAlgorithmTypeOk() (*string, bool)`

GetLoadBalancingAlgorithmTypeOk returns a tuple with the LoadBalancingAlgorithmType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancingAlgorithmType

`func (o *Vhost) SetLoadBalancingAlgorithmType(v string)`

SetLoadBalancingAlgorithmType sets LoadBalancingAlgorithmType field to given value.


### SetLoadBalancingAlgorithmTypeNil

`func (o *Vhost) SetLoadBalancingAlgorithmTypeNil(b bool)`

 SetLoadBalancingAlgorithmTypeNil sets the value for LoadBalancingAlgorithmType to be an explicit nil

### UnsetLoadBalancingAlgorithmType
`func (o *Vhost) UnsetLoadBalancingAlgorithmType()`

UnsetLoadBalancingAlgorithmType ensures that no value is present for LoadBalancingAlgorithmType, not even an explicit nil
### GetLinks

`func (o *Vhost) GetLinks() VhostLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Vhost) GetLinksOk() (*VhostLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Vhost) SetLinks(v VhostLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Vhost) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


