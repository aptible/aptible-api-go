# VpnTunnel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**MetaType** | **string** |  | 
**CreatedAt** | **string** |  | 
**UpdatedAt** | **string** |  | 
**Handle** | **string** |  | 
**Phase1Alg** | **NullableString** |  | 
**Phase1DhGroup** | **NullableInt32** |  | 
**Phase1Lifetime** | **NullableInt32** |  | 
**Phase2Alg** | **NullableString** |  | 
**Phase2DhGroup** | **NullableInt32** |  | 
**Phase2Lifetime** | **NullableInt32** |  | 
**PerfectForwardSecrecy** | **NullableBool** |  | 
**OurGateway** | **NullableString** |  | 
**OurNetworks** | **[][]string** |  | 
**PeerGateway** | **NullableString** |  | 
**PeerNetworks** | **[][]string** |  | 
**Links** | Pointer to [**VpnTunnelLinks**](VpnTunnelLinks.md) |  | [optional] 

## Methods

### NewVpnTunnel

`func NewVpnTunnel(id int32, metaType string, createdAt string, updatedAt string, handle string, phase1Alg NullableString, phase1DhGroup NullableInt32, phase1Lifetime NullableInt32, phase2Alg NullableString, phase2DhGroup NullableInt32, phase2Lifetime NullableInt32, perfectForwardSecrecy NullableBool, ourGateway NullableString, ourNetworks [][]string, peerGateway NullableString, peerNetworks [][]string, ) *VpnTunnel`

NewVpnTunnel instantiates a new VpnTunnel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpnTunnelWithDefaults

`func NewVpnTunnelWithDefaults() *VpnTunnel`

NewVpnTunnelWithDefaults instantiates a new VpnTunnel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VpnTunnel) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VpnTunnel) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VpnTunnel) SetId(v int32)`

SetId sets Id field to given value.


### GetMetaType

`func (o *VpnTunnel) GetMetaType() string`

GetMetaType returns the MetaType field if non-nil, zero value otherwise.

### GetMetaTypeOk

`func (o *VpnTunnel) GetMetaTypeOk() (*string, bool)`

GetMetaTypeOk returns a tuple with the MetaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaType

`func (o *VpnTunnel) SetMetaType(v string)`

SetMetaType sets MetaType field to given value.


### GetCreatedAt

`func (o *VpnTunnel) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *VpnTunnel) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *VpnTunnel) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *VpnTunnel) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *VpnTunnel) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *VpnTunnel) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetHandle

`func (o *VpnTunnel) GetHandle() string`

GetHandle returns the Handle field if non-nil, zero value otherwise.

### GetHandleOk

`func (o *VpnTunnel) GetHandleOk() (*string, bool)`

GetHandleOk returns a tuple with the Handle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandle

`func (o *VpnTunnel) SetHandle(v string)`

SetHandle sets Handle field to given value.


### GetPhase1Alg

`func (o *VpnTunnel) GetPhase1Alg() string`

GetPhase1Alg returns the Phase1Alg field if non-nil, zero value otherwise.

### GetPhase1AlgOk

`func (o *VpnTunnel) GetPhase1AlgOk() (*string, bool)`

GetPhase1AlgOk returns a tuple with the Phase1Alg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase1Alg

`func (o *VpnTunnel) SetPhase1Alg(v string)`

SetPhase1Alg sets Phase1Alg field to given value.


### SetPhase1AlgNil

`func (o *VpnTunnel) SetPhase1AlgNil(b bool)`

 SetPhase1AlgNil sets the value for Phase1Alg to be an explicit nil

### UnsetPhase1Alg
`func (o *VpnTunnel) UnsetPhase1Alg()`

UnsetPhase1Alg ensures that no value is present for Phase1Alg, not even an explicit nil
### GetPhase1DhGroup

`func (o *VpnTunnel) GetPhase1DhGroup() int32`

GetPhase1DhGroup returns the Phase1DhGroup field if non-nil, zero value otherwise.

### GetPhase1DhGroupOk

`func (o *VpnTunnel) GetPhase1DhGroupOk() (*int32, bool)`

GetPhase1DhGroupOk returns a tuple with the Phase1DhGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase1DhGroup

`func (o *VpnTunnel) SetPhase1DhGroup(v int32)`

SetPhase1DhGroup sets Phase1DhGroup field to given value.


### SetPhase1DhGroupNil

`func (o *VpnTunnel) SetPhase1DhGroupNil(b bool)`

 SetPhase1DhGroupNil sets the value for Phase1DhGroup to be an explicit nil

### UnsetPhase1DhGroup
`func (o *VpnTunnel) UnsetPhase1DhGroup()`

UnsetPhase1DhGroup ensures that no value is present for Phase1DhGroup, not even an explicit nil
### GetPhase1Lifetime

`func (o *VpnTunnel) GetPhase1Lifetime() int32`

GetPhase1Lifetime returns the Phase1Lifetime field if non-nil, zero value otherwise.

### GetPhase1LifetimeOk

`func (o *VpnTunnel) GetPhase1LifetimeOk() (*int32, bool)`

GetPhase1LifetimeOk returns a tuple with the Phase1Lifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase1Lifetime

`func (o *VpnTunnel) SetPhase1Lifetime(v int32)`

SetPhase1Lifetime sets Phase1Lifetime field to given value.


### SetPhase1LifetimeNil

`func (o *VpnTunnel) SetPhase1LifetimeNil(b bool)`

 SetPhase1LifetimeNil sets the value for Phase1Lifetime to be an explicit nil

### UnsetPhase1Lifetime
`func (o *VpnTunnel) UnsetPhase1Lifetime()`

UnsetPhase1Lifetime ensures that no value is present for Phase1Lifetime, not even an explicit nil
### GetPhase2Alg

`func (o *VpnTunnel) GetPhase2Alg() string`

GetPhase2Alg returns the Phase2Alg field if non-nil, zero value otherwise.

### GetPhase2AlgOk

`func (o *VpnTunnel) GetPhase2AlgOk() (*string, bool)`

GetPhase2AlgOk returns a tuple with the Phase2Alg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase2Alg

`func (o *VpnTunnel) SetPhase2Alg(v string)`

SetPhase2Alg sets Phase2Alg field to given value.


### SetPhase2AlgNil

`func (o *VpnTunnel) SetPhase2AlgNil(b bool)`

 SetPhase2AlgNil sets the value for Phase2Alg to be an explicit nil

### UnsetPhase2Alg
`func (o *VpnTunnel) UnsetPhase2Alg()`

UnsetPhase2Alg ensures that no value is present for Phase2Alg, not even an explicit nil
### GetPhase2DhGroup

`func (o *VpnTunnel) GetPhase2DhGroup() int32`

GetPhase2DhGroup returns the Phase2DhGroup field if non-nil, zero value otherwise.

### GetPhase2DhGroupOk

`func (o *VpnTunnel) GetPhase2DhGroupOk() (*int32, bool)`

GetPhase2DhGroupOk returns a tuple with the Phase2DhGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase2DhGroup

`func (o *VpnTunnel) SetPhase2DhGroup(v int32)`

SetPhase2DhGroup sets Phase2DhGroup field to given value.


### SetPhase2DhGroupNil

`func (o *VpnTunnel) SetPhase2DhGroupNil(b bool)`

 SetPhase2DhGroupNil sets the value for Phase2DhGroup to be an explicit nil

### UnsetPhase2DhGroup
`func (o *VpnTunnel) UnsetPhase2DhGroup()`

UnsetPhase2DhGroup ensures that no value is present for Phase2DhGroup, not even an explicit nil
### GetPhase2Lifetime

`func (o *VpnTunnel) GetPhase2Lifetime() int32`

GetPhase2Lifetime returns the Phase2Lifetime field if non-nil, zero value otherwise.

### GetPhase2LifetimeOk

`func (o *VpnTunnel) GetPhase2LifetimeOk() (*int32, bool)`

GetPhase2LifetimeOk returns a tuple with the Phase2Lifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase2Lifetime

`func (o *VpnTunnel) SetPhase2Lifetime(v int32)`

SetPhase2Lifetime sets Phase2Lifetime field to given value.


### SetPhase2LifetimeNil

`func (o *VpnTunnel) SetPhase2LifetimeNil(b bool)`

 SetPhase2LifetimeNil sets the value for Phase2Lifetime to be an explicit nil

### UnsetPhase2Lifetime
`func (o *VpnTunnel) UnsetPhase2Lifetime()`

UnsetPhase2Lifetime ensures that no value is present for Phase2Lifetime, not even an explicit nil
### GetPerfectForwardSecrecy

`func (o *VpnTunnel) GetPerfectForwardSecrecy() bool`

GetPerfectForwardSecrecy returns the PerfectForwardSecrecy field if non-nil, zero value otherwise.

### GetPerfectForwardSecrecyOk

`func (o *VpnTunnel) GetPerfectForwardSecrecyOk() (*bool, bool)`

GetPerfectForwardSecrecyOk returns a tuple with the PerfectForwardSecrecy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerfectForwardSecrecy

`func (o *VpnTunnel) SetPerfectForwardSecrecy(v bool)`

SetPerfectForwardSecrecy sets PerfectForwardSecrecy field to given value.


### SetPerfectForwardSecrecyNil

`func (o *VpnTunnel) SetPerfectForwardSecrecyNil(b bool)`

 SetPerfectForwardSecrecyNil sets the value for PerfectForwardSecrecy to be an explicit nil

### UnsetPerfectForwardSecrecy
`func (o *VpnTunnel) UnsetPerfectForwardSecrecy()`

UnsetPerfectForwardSecrecy ensures that no value is present for PerfectForwardSecrecy, not even an explicit nil
### GetOurGateway

`func (o *VpnTunnel) GetOurGateway() string`

GetOurGateway returns the OurGateway field if non-nil, zero value otherwise.

### GetOurGatewayOk

`func (o *VpnTunnel) GetOurGatewayOk() (*string, bool)`

GetOurGatewayOk returns a tuple with the OurGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOurGateway

`func (o *VpnTunnel) SetOurGateway(v string)`

SetOurGateway sets OurGateway field to given value.


### SetOurGatewayNil

`func (o *VpnTunnel) SetOurGatewayNil(b bool)`

 SetOurGatewayNil sets the value for OurGateway to be an explicit nil

### UnsetOurGateway
`func (o *VpnTunnel) UnsetOurGateway()`

UnsetOurGateway ensures that no value is present for OurGateway, not even an explicit nil
### GetOurNetworks

`func (o *VpnTunnel) GetOurNetworks() [][]string`

GetOurNetworks returns the OurNetworks field if non-nil, zero value otherwise.

### GetOurNetworksOk

`func (o *VpnTunnel) GetOurNetworksOk() (*[][]string, bool)`

GetOurNetworksOk returns a tuple with the OurNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOurNetworks

`func (o *VpnTunnel) SetOurNetworks(v [][]string)`

SetOurNetworks sets OurNetworks field to given value.


### SetOurNetworksNil

`func (o *VpnTunnel) SetOurNetworksNil(b bool)`

 SetOurNetworksNil sets the value for OurNetworks to be an explicit nil

### UnsetOurNetworks
`func (o *VpnTunnel) UnsetOurNetworks()`

UnsetOurNetworks ensures that no value is present for OurNetworks, not even an explicit nil
### GetPeerGateway

`func (o *VpnTunnel) GetPeerGateway() string`

GetPeerGateway returns the PeerGateway field if non-nil, zero value otherwise.

### GetPeerGatewayOk

`func (o *VpnTunnel) GetPeerGatewayOk() (*string, bool)`

GetPeerGatewayOk returns a tuple with the PeerGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerGateway

`func (o *VpnTunnel) SetPeerGateway(v string)`

SetPeerGateway sets PeerGateway field to given value.


### SetPeerGatewayNil

`func (o *VpnTunnel) SetPeerGatewayNil(b bool)`

 SetPeerGatewayNil sets the value for PeerGateway to be an explicit nil

### UnsetPeerGateway
`func (o *VpnTunnel) UnsetPeerGateway()`

UnsetPeerGateway ensures that no value is present for PeerGateway, not even an explicit nil
### GetPeerNetworks

`func (o *VpnTunnel) GetPeerNetworks() [][]string`

GetPeerNetworks returns the PeerNetworks field if non-nil, zero value otherwise.

### GetPeerNetworksOk

`func (o *VpnTunnel) GetPeerNetworksOk() (*[][]string, bool)`

GetPeerNetworksOk returns a tuple with the PeerNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerNetworks

`func (o *VpnTunnel) SetPeerNetworks(v [][]string)`

SetPeerNetworks sets PeerNetworks field to given value.


### SetPeerNetworksNil

`func (o *VpnTunnel) SetPeerNetworksNil(b bool)`

 SetPeerNetworksNil sets the value for PeerNetworks to be an explicit nil

### UnsetPeerNetworks
`func (o *VpnTunnel) UnsetPeerNetworks()`

UnsetPeerNetworks ensures that no value is present for PeerNetworks, not even an explicit nil
### GetLinks

`func (o *VpnTunnel) GetLinks() VpnTunnelLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *VpnTunnel) GetLinksOk() (*VpnTunnelLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *VpnTunnel) SetLinks(v VpnTunnelLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *VpnTunnel) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


