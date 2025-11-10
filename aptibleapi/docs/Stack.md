# Stack

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**MetaType** | **string** |  | 
**Name** | **string** |  | 
**Version** | **string** |  | 
**Region** | **string** |  | 
**Default** | **NullableBool** |  | 
**Public** | **bool** |  | 
**CreatedAt** | **string** |  | 
**UpdatedAt** | **string** |  | 
**SshHostDsaPublicKey** | **string** |  | 
**SshHostRsaPublicKey** | **string** |  | 
**SshHostEcdsaPublicKey** | **string** |  | 
**SshPortalHost** | **string** |  | 
**SshPortalPort** | **int32** |  | 
**OutboundIpAddresses** | **[]string** |  | 
**MemoryLimits** | **bool** |  | 
**CpuLimits** | **bool** |  | 
**IntrusionDetection** | **bool** |  | 
**ExposeIntrusionDetectionReports** | **bool** |  | 
**AccountId** | **NullableString** |  | 
**VpcId** | **NullableString** |  | 
**InternalDomain** | **NullableString** |  | 
**DefaultDomain** | **NullableString** |  | 
**BrickwallEnabled** | **bool** |  | 
**VpcCidr** | **NullableString** |  | 
**Links** | Pointer to [**StackLinks**](StackLinks.md) |  | [optional] 

## Methods

### NewStack

`func NewStack(id int32, metaType string, name string, version string, region string, default_ NullableBool, public bool, createdAt string, updatedAt string, sshHostDsaPublicKey string, sshHostRsaPublicKey string, sshHostEcdsaPublicKey string, sshPortalHost string, sshPortalPort int32, outboundIpAddresses []string, memoryLimits bool, cpuLimits bool, intrusionDetection bool, exposeIntrusionDetectionReports bool, accountId NullableString, vpcId NullableString, internalDomain NullableString, defaultDomain NullableString, brickwallEnabled bool, vpcCidr NullableString, ) *Stack`

NewStack instantiates a new Stack object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStackWithDefaults

`func NewStackWithDefaults() *Stack`

NewStackWithDefaults instantiates a new Stack object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Stack) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Stack) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Stack) SetId(v int32)`

SetId sets Id field to given value.


### GetMetaType

`func (o *Stack) GetMetaType() string`

GetMetaType returns the MetaType field if non-nil, zero value otherwise.

### GetMetaTypeOk

`func (o *Stack) GetMetaTypeOk() (*string, bool)`

GetMetaTypeOk returns a tuple with the MetaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaType

`func (o *Stack) SetMetaType(v string)`

SetMetaType sets MetaType field to given value.


### GetName

`func (o *Stack) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Stack) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Stack) SetName(v string)`

SetName sets Name field to given value.


### GetVersion

`func (o *Stack) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Stack) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Stack) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetRegion

`func (o *Stack) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *Stack) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *Stack) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetDefault

`func (o *Stack) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *Stack) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *Stack) SetDefault(v bool)`

SetDefault sets Default field to given value.


### SetDefaultNil

`func (o *Stack) SetDefaultNil(b bool)`

 SetDefaultNil sets the value for Default to be an explicit nil

### UnsetDefault
`func (o *Stack) UnsetDefault()`

UnsetDefault ensures that no value is present for Default, not even an explicit nil
### GetPublic

`func (o *Stack) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *Stack) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *Stack) SetPublic(v bool)`

SetPublic sets Public field to given value.


### GetCreatedAt

`func (o *Stack) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Stack) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Stack) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Stack) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Stack) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Stack) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetSshHostDsaPublicKey

`func (o *Stack) GetSshHostDsaPublicKey() string`

GetSshHostDsaPublicKey returns the SshHostDsaPublicKey field if non-nil, zero value otherwise.

### GetSshHostDsaPublicKeyOk

`func (o *Stack) GetSshHostDsaPublicKeyOk() (*string, bool)`

GetSshHostDsaPublicKeyOk returns a tuple with the SshHostDsaPublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshHostDsaPublicKey

`func (o *Stack) SetSshHostDsaPublicKey(v string)`

SetSshHostDsaPublicKey sets SshHostDsaPublicKey field to given value.


### GetSshHostRsaPublicKey

`func (o *Stack) GetSshHostRsaPublicKey() string`

GetSshHostRsaPublicKey returns the SshHostRsaPublicKey field if non-nil, zero value otherwise.

### GetSshHostRsaPublicKeyOk

`func (o *Stack) GetSshHostRsaPublicKeyOk() (*string, bool)`

GetSshHostRsaPublicKeyOk returns a tuple with the SshHostRsaPublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshHostRsaPublicKey

`func (o *Stack) SetSshHostRsaPublicKey(v string)`

SetSshHostRsaPublicKey sets SshHostRsaPublicKey field to given value.


### GetSshHostEcdsaPublicKey

`func (o *Stack) GetSshHostEcdsaPublicKey() string`

GetSshHostEcdsaPublicKey returns the SshHostEcdsaPublicKey field if non-nil, zero value otherwise.

### GetSshHostEcdsaPublicKeyOk

`func (o *Stack) GetSshHostEcdsaPublicKeyOk() (*string, bool)`

GetSshHostEcdsaPublicKeyOk returns a tuple with the SshHostEcdsaPublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshHostEcdsaPublicKey

`func (o *Stack) SetSshHostEcdsaPublicKey(v string)`

SetSshHostEcdsaPublicKey sets SshHostEcdsaPublicKey field to given value.


### GetSshPortalHost

`func (o *Stack) GetSshPortalHost() string`

GetSshPortalHost returns the SshPortalHost field if non-nil, zero value otherwise.

### GetSshPortalHostOk

`func (o *Stack) GetSshPortalHostOk() (*string, bool)`

GetSshPortalHostOk returns a tuple with the SshPortalHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPortalHost

`func (o *Stack) SetSshPortalHost(v string)`

SetSshPortalHost sets SshPortalHost field to given value.


### GetSshPortalPort

`func (o *Stack) GetSshPortalPort() int32`

GetSshPortalPort returns the SshPortalPort field if non-nil, zero value otherwise.

### GetSshPortalPortOk

`func (o *Stack) GetSshPortalPortOk() (*int32, bool)`

GetSshPortalPortOk returns a tuple with the SshPortalPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPortalPort

`func (o *Stack) SetSshPortalPort(v int32)`

SetSshPortalPort sets SshPortalPort field to given value.


### GetOutboundIpAddresses

`func (o *Stack) GetOutboundIpAddresses() []string`

GetOutboundIpAddresses returns the OutboundIpAddresses field if non-nil, zero value otherwise.

### GetOutboundIpAddressesOk

`func (o *Stack) GetOutboundIpAddressesOk() (*[]string, bool)`

GetOutboundIpAddressesOk returns a tuple with the OutboundIpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutboundIpAddresses

`func (o *Stack) SetOutboundIpAddresses(v []string)`

SetOutboundIpAddresses sets OutboundIpAddresses field to given value.


### GetMemoryLimits

`func (o *Stack) GetMemoryLimits() bool`

GetMemoryLimits returns the MemoryLimits field if non-nil, zero value otherwise.

### GetMemoryLimitsOk

`func (o *Stack) GetMemoryLimitsOk() (*bool, bool)`

GetMemoryLimitsOk returns a tuple with the MemoryLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryLimits

`func (o *Stack) SetMemoryLimits(v bool)`

SetMemoryLimits sets MemoryLimits field to given value.


### GetCpuLimits

`func (o *Stack) GetCpuLimits() bool`

GetCpuLimits returns the CpuLimits field if non-nil, zero value otherwise.

### GetCpuLimitsOk

`func (o *Stack) GetCpuLimitsOk() (*bool, bool)`

GetCpuLimitsOk returns a tuple with the CpuLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuLimits

`func (o *Stack) SetCpuLimits(v bool)`

SetCpuLimits sets CpuLimits field to given value.


### GetIntrusionDetection

`func (o *Stack) GetIntrusionDetection() bool`

GetIntrusionDetection returns the IntrusionDetection field if non-nil, zero value otherwise.

### GetIntrusionDetectionOk

`func (o *Stack) GetIntrusionDetectionOk() (*bool, bool)`

GetIntrusionDetectionOk returns a tuple with the IntrusionDetection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntrusionDetection

`func (o *Stack) SetIntrusionDetection(v bool)`

SetIntrusionDetection sets IntrusionDetection field to given value.


### GetExposeIntrusionDetectionReports

`func (o *Stack) GetExposeIntrusionDetectionReports() bool`

GetExposeIntrusionDetectionReports returns the ExposeIntrusionDetectionReports field if non-nil, zero value otherwise.

### GetExposeIntrusionDetectionReportsOk

`func (o *Stack) GetExposeIntrusionDetectionReportsOk() (*bool, bool)`

GetExposeIntrusionDetectionReportsOk returns a tuple with the ExposeIntrusionDetectionReports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExposeIntrusionDetectionReports

`func (o *Stack) SetExposeIntrusionDetectionReports(v bool)`

SetExposeIntrusionDetectionReports sets ExposeIntrusionDetectionReports field to given value.


### GetAccountId

`func (o *Stack) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *Stack) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *Stack) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### SetAccountIdNil

`func (o *Stack) SetAccountIdNil(b bool)`

 SetAccountIdNil sets the value for AccountId to be an explicit nil

### UnsetAccountId
`func (o *Stack) UnsetAccountId()`

UnsetAccountId ensures that no value is present for AccountId, not even an explicit nil
### GetVpcId

`func (o *Stack) GetVpcId() string`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *Stack) GetVpcIdOk() (*string, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *Stack) SetVpcId(v string)`

SetVpcId sets VpcId field to given value.


### SetVpcIdNil

`func (o *Stack) SetVpcIdNil(b bool)`

 SetVpcIdNil sets the value for VpcId to be an explicit nil

### UnsetVpcId
`func (o *Stack) UnsetVpcId()`

UnsetVpcId ensures that no value is present for VpcId, not even an explicit nil
### GetInternalDomain

`func (o *Stack) GetInternalDomain() string`

GetInternalDomain returns the InternalDomain field if non-nil, zero value otherwise.

### GetInternalDomainOk

`func (o *Stack) GetInternalDomainOk() (*string, bool)`

GetInternalDomainOk returns a tuple with the InternalDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalDomain

`func (o *Stack) SetInternalDomain(v string)`

SetInternalDomain sets InternalDomain field to given value.


### SetInternalDomainNil

`func (o *Stack) SetInternalDomainNil(b bool)`

 SetInternalDomainNil sets the value for InternalDomain to be an explicit nil

### UnsetInternalDomain
`func (o *Stack) UnsetInternalDomain()`

UnsetInternalDomain ensures that no value is present for InternalDomain, not even an explicit nil
### GetDefaultDomain

`func (o *Stack) GetDefaultDomain() string`

GetDefaultDomain returns the DefaultDomain field if non-nil, zero value otherwise.

### GetDefaultDomainOk

`func (o *Stack) GetDefaultDomainOk() (*string, bool)`

GetDefaultDomainOk returns a tuple with the DefaultDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultDomain

`func (o *Stack) SetDefaultDomain(v string)`

SetDefaultDomain sets DefaultDomain field to given value.


### SetDefaultDomainNil

`func (o *Stack) SetDefaultDomainNil(b bool)`

 SetDefaultDomainNil sets the value for DefaultDomain to be an explicit nil

### UnsetDefaultDomain
`func (o *Stack) UnsetDefaultDomain()`

UnsetDefaultDomain ensures that no value is present for DefaultDomain, not even an explicit nil
### GetBrickwallEnabled

`func (o *Stack) GetBrickwallEnabled() bool`

GetBrickwallEnabled returns the BrickwallEnabled field if non-nil, zero value otherwise.

### GetBrickwallEnabledOk

`func (o *Stack) GetBrickwallEnabledOk() (*bool, bool)`

GetBrickwallEnabledOk returns a tuple with the BrickwallEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrickwallEnabled

`func (o *Stack) SetBrickwallEnabled(v bool)`

SetBrickwallEnabled sets BrickwallEnabled field to given value.


### GetVpcCidr

`func (o *Stack) GetVpcCidr() string`

GetVpcCidr returns the VpcCidr field if non-nil, zero value otherwise.

### GetVpcCidrOk

`func (o *Stack) GetVpcCidrOk() (*string, bool)`

GetVpcCidrOk returns a tuple with the VpcCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcCidr

`func (o *Stack) SetVpcCidr(v string)`

SetVpcCidr sets VpcCidr field to given value.


### SetVpcCidrNil

`func (o *Stack) SetVpcCidrNil(b bool)`

 SetVpcCidrNil sets the value for VpcCidr to be an explicit nil

### UnsetVpcCidr
`func (o *Stack) UnsetVpcCidr()`

UnsetVpcCidr ensures that no value is present for VpcCidr, not even an explicit nil
### GetLinks

`func (o *Stack) GetLinks() StackLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Stack) GetLinksOk() (*StackLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Stack) SetLinks(v StackLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Stack) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


