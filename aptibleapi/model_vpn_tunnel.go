/*
Deploy API v1

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package aptibleapi

import (
	"encoding/json"
	"fmt"
)

// checks if the VpnTunnel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VpnTunnel{}

// VpnTunnel struct for VpnTunnel
type VpnTunnel struct {
	Id int32 `json:"id"`
	MetaType string `json:"_type"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Handle string `json:"handle"`
	Phase1Alg NullableString `json:"phase_1_alg"`
	Phase1DhGroup NullableInt32 `json:"phase_1_dh_group"`
	Phase1Lifetime NullableInt32 `json:"phase_1_lifetime"`
	Phase2Alg NullableString `json:"phase_2_alg"`
	Phase2DhGroup NullableInt32 `json:"phase_2_dh_group"`
	Phase2Lifetime NullableInt32 `json:"phase_2_lifetime"`
	PerfectForwardSecrecy NullableBool `json:"perfect_forward_secrecy"`
	OurGateway NullableString `json:"our_gateway"`
	OurNetworks [][]string `json:"our_networks"`
	PeerGateway NullableString `json:"peer_gateway"`
	PeerNetworks [][]string `json:"peer_networks"`
	Links *VpcPeerLinks `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VpnTunnel VpnTunnel

// NewVpnTunnel instantiates a new VpnTunnel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVpnTunnel(id int32, metaType string, createdAt string, updatedAt string, handle string, phase1Alg NullableString, phase1DhGroup NullableInt32, phase1Lifetime NullableInt32, phase2Alg NullableString, phase2DhGroup NullableInt32, phase2Lifetime NullableInt32, perfectForwardSecrecy NullableBool, ourGateway NullableString, ourNetworks [][]string, peerGateway NullableString, peerNetworks [][]string) *VpnTunnel {
	this := VpnTunnel{}
	this.Id = id
	this.MetaType = metaType
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.Handle = handle
	this.Phase1Alg = phase1Alg
	this.Phase1DhGroup = phase1DhGroup
	this.Phase1Lifetime = phase1Lifetime
	this.Phase2Alg = phase2Alg
	this.Phase2DhGroup = phase2DhGroup
	this.Phase2Lifetime = phase2Lifetime
	this.PerfectForwardSecrecy = perfectForwardSecrecy
	this.OurGateway = ourGateway
	this.OurNetworks = ourNetworks
	this.PeerGateway = peerGateway
	this.PeerNetworks = peerNetworks
	return &this
}

// NewVpnTunnelWithDefaults instantiates a new VpnTunnel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVpnTunnelWithDefaults() *VpnTunnel {
	this := VpnTunnel{}
	return &this
}

// GetId returns the Id field value
func (o *VpnTunnel) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *VpnTunnel) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *VpnTunnel) SetId(v int32) {
	o.Id = v
}

// GetMetaType returns the MetaType field value
func (o *VpnTunnel) GetMetaType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MetaType
}

// GetMetaTypeOk returns a tuple with the MetaType field value
// and a boolean to check if the value has been set.
func (o *VpnTunnel) GetMetaTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetaType, true
}

// SetMetaType sets field value
func (o *VpnTunnel) SetMetaType(v string) {
	o.MetaType = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *VpnTunnel) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *VpnTunnel) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *VpnTunnel) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *VpnTunnel) GetUpdatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *VpnTunnel) GetUpdatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *VpnTunnel) SetUpdatedAt(v string) {
	o.UpdatedAt = v
}

// GetHandle returns the Handle field value
func (o *VpnTunnel) GetHandle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Handle
}

// GetHandleOk returns a tuple with the Handle field value
// and a boolean to check if the value has been set.
func (o *VpnTunnel) GetHandleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Handle, true
}

// SetHandle sets field value
func (o *VpnTunnel) SetHandle(v string) {
	o.Handle = v
}

// GetPhase1Alg returns the Phase1Alg field value
// If the value is explicit nil, the zero value for string will be returned
func (o *VpnTunnel) GetPhase1Alg() string {
	if o == nil || o.Phase1Alg.Get() == nil {
		var ret string
		return ret
	}

	return *o.Phase1Alg.Get()
}

// GetPhase1AlgOk returns a tuple with the Phase1Alg field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VpnTunnel) GetPhase1AlgOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Phase1Alg.Get(), o.Phase1Alg.IsSet()
}

// SetPhase1Alg sets field value
func (o *VpnTunnel) SetPhase1Alg(v string) {
	o.Phase1Alg.Set(&v)
}

// GetPhase1DhGroup returns the Phase1DhGroup field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *VpnTunnel) GetPhase1DhGroup() int32 {
	if o == nil || o.Phase1DhGroup.Get() == nil {
		var ret int32
		return ret
	}

	return *o.Phase1DhGroup.Get()
}

// GetPhase1DhGroupOk returns a tuple with the Phase1DhGroup field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VpnTunnel) GetPhase1DhGroupOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Phase1DhGroup.Get(), o.Phase1DhGroup.IsSet()
}

// SetPhase1DhGroup sets field value
func (o *VpnTunnel) SetPhase1DhGroup(v int32) {
	o.Phase1DhGroup.Set(&v)
}

// GetPhase1Lifetime returns the Phase1Lifetime field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *VpnTunnel) GetPhase1Lifetime() int32 {
	if o == nil || o.Phase1Lifetime.Get() == nil {
		var ret int32
		return ret
	}

	return *o.Phase1Lifetime.Get()
}

// GetPhase1LifetimeOk returns a tuple with the Phase1Lifetime field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VpnTunnel) GetPhase1LifetimeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Phase1Lifetime.Get(), o.Phase1Lifetime.IsSet()
}

// SetPhase1Lifetime sets field value
func (o *VpnTunnel) SetPhase1Lifetime(v int32) {
	o.Phase1Lifetime.Set(&v)
}

// GetPhase2Alg returns the Phase2Alg field value
// If the value is explicit nil, the zero value for string will be returned
func (o *VpnTunnel) GetPhase2Alg() string {
	if o == nil || o.Phase2Alg.Get() == nil {
		var ret string
		return ret
	}

	return *o.Phase2Alg.Get()
}

// GetPhase2AlgOk returns a tuple with the Phase2Alg field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VpnTunnel) GetPhase2AlgOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Phase2Alg.Get(), o.Phase2Alg.IsSet()
}

// SetPhase2Alg sets field value
func (o *VpnTunnel) SetPhase2Alg(v string) {
	o.Phase2Alg.Set(&v)
}

// GetPhase2DhGroup returns the Phase2DhGroup field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *VpnTunnel) GetPhase2DhGroup() int32 {
	if o == nil || o.Phase2DhGroup.Get() == nil {
		var ret int32
		return ret
	}

	return *o.Phase2DhGroup.Get()
}

// GetPhase2DhGroupOk returns a tuple with the Phase2DhGroup field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VpnTunnel) GetPhase2DhGroupOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Phase2DhGroup.Get(), o.Phase2DhGroup.IsSet()
}

// SetPhase2DhGroup sets field value
func (o *VpnTunnel) SetPhase2DhGroup(v int32) {
	o.Phase2DhGroup.Set(&v)
}

// GetPhase2Lifetime returns the Phase2Lifetime field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *VpnTunnel) GetPhase2Lifetime() int32 {
	if o == nil || o.Phase2Lifetime.Get() == nil {
		var ret int32
		return ret
	}

	return *o.Phase2Lifetime.Get()
}

// GetPhase2LifetimeOk returns a tuple with the Phase2Lifetime field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VpnTunnel) GetPhase2LifetimeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Phase2Lifetime.Get(), o.Phase2Lifetime.IsSet()
}

// SetPhase2Lifetime sets field value
func (o *VpnTunnel) SetPhase2Lifetime(v int32) {
	o.Phase2Lifetime.Set(&v)
}

// GetPerfectForwardSecrecy returns the PerfectForwardSecrecy field value
// If the value is explicit nil, the zero value for bool will be returned
func (o *VpnTunnel) GetPerfectForwardSecrecy() bool {
	if o == nil || o.PerfectForwardSecrecy.Get() == nil {
		var ret bool
		return ret
	}

	return *o.PerfectForwardSecrecy.Get()
}

// GetPerfectForwardSecrecyOk returns a tuple with the PerfectForwardSecrecy field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VpnTunnel) GetPerfectForwardSecrecyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.PerfectForwardSecrecy.Get(), o.PerfectForwardSecrecy.IsSet()
}

// SetPerfectForwardSecrecy sets field value
func (o *VpnTunnel) SetPerfectForwardSecrecy(v bool) {
	o.PerfectForwardSecrecy.Set(&v)
}

// GetOurGateway returns the OurGateway field value
// If the value is explicit nil, the zero value for string will be returned
func (o *VpnTunnel) GetOurGateway() string {
	if o == nil || o.OurGateway.Get() == nil {
		var ret string
		return ret
	}

	return *o.OurGateway.Get()
}

// GetOurGatewayOk returns a tuple with the OurGateway field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VpnTunnel) GetOurGatewayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OurGateway.Get(), o.OurGateway.IsSet()
}

// SetOurGateway sets field value
func (o *VpnTunnel) SetOurGateway(v string) {
	o.OurGateway.Set(&v)
}

// GetOurNetworks returns the OurNetworks field value
// If the value is explicit nil, the zero value for [][]string will be returned
func (o *VpnTunnel) GetOurNetworks() [][]string {
	if o == nil {
		var ret [][]string
		return ret
	}

	return o.OurNetworks
}

// GetOurNetworksOk returns a tuple with the OurNetworks field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VpnTunnel) GetOurNetworksOk() ([][]string, bool) {
	if o == nil || IsNil(o.OurNetworks) {
		return nil, false
	}
	return o.OurNetworks, true
}

// SetOurNetworks sets field value
func (o *VpnTunnel) SetOurNetworks(v [][]string) {
	o.OurNetworks = v
}

// GetPeerGateway returns the PeerGateway field value
// If the value is explicit nil, the zero value for string will be returned
func (o *VpnTunnel) GetPeerGateway() string {
	if o == nil || o.PeerGateway.Get() == nil {
		var ret string
		return ret
	}

	return *o.PeerGateway.Get()
}

// GetPeerGatewayOk returns a tuple with the PeerGateway field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VpnTunnel) GetPeerGatewayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PeerGateway.Get(), o.PeerGateway.IsSet()
}

// SetPeerGateway sets field value
func (o *VpnTunnel) SetPeerGateway(v string) {
	o.PeerGateway.Set(&v)
}

// GetPeerNetworks returns the PeerNetworks field value
// If the value is explicit nil, the zero value for [][]string will be returned
func (o *VpnTunnel) GetPeerNetworks() [][]string {
	if o == nil {
		var ret [][]string
		return ret
	}

	return o.PeerNetworks
}

// GetPeerNetworksOk returns a tuple with the PeerNetworks field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VpnTunnel) GetPeerNetworksOk() ([][]string, bool) {
	if o == nil || IsNil(o.PeerNetworks) {
		return nil, false
	}
	return o.PeerNetworks, true
}

// SetPeerNetworks sets field value
func (o *VpnTunnel) SetPeerNetworks(v [][]string) {
	o.PeerNetworks = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *VpnTunnel) GetLinks() VpcPeerLinks {
	if o == nil || IsNil(o.Links) {
		var ret VpcPeerLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VpnTunnel) GetLinksOk() (*VpcPeerLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *VpnTunnel) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given VpcPeerLinks and assigns it to the Links field.
func (o *VpnTunnel) SetLinks(v VpcPeerLinks) {
	o.Links = &v
}

func (o VpnTunnel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VpnTunnel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["_type"] = o.MetaType
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["updated_at"] = o.UpdatedAt
	toSerialize["handle"] = o.Handle
	toSerialize["phase_1_alg"] = o.Phase1Alg.Get()
	toSerialize["phase_1_dh_group"] = o.Phase1DhGroup.Get()
	toSerialize["phase_1_lifetime"] = o.Phase1Lifetime.Get()
	toSerialize["phase_2_alg"] = o.Phase2Alg.Get()
	toSerialize["phase_2_dh_group"] = o.Phase2DhGroup.Get()
	toSerialize["phase_2_lifetime"] = o.Phase2Lifetime.Get()
	toSerialize["perfect_forward_secrecy"] = o.PerfectForwardSecrecy.Get()
	toSerialize["our_gateway"] = o.OurGateway.Get()
	if o.OurNetworks != nil {
		toSerialize["our_networks"] = o.OurNetworks
	}
	toSerialize["peer_gateway"] = o.PeerGateway.Get()
	if o.PeerNetworks != nil {
		toSerialize["peer_networks"] = o.PeerNetworks
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *VpnTunnel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"_type",
		"created_at",
		"updated_at",
		"handle",
		"phase_1_alg",
		"phase_1_dh_group",
		"phase_1_lifetime",
		"phase_2_alg",
		"phase_2_dh_group",
		"phase_2_lifetime",
		"perfect_forward_secrecy",
		"our_gateway",
		"our_networks",
		"peer_gateway",
		"peer_networks",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varVpnTunnel := _VpnTunnel{}

	err = json.Unmarshal(data, &varVpnTunnel)

	if err != nil {
		return err
	}

	*o = VpnTunnel(varVpnTunnel)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "_type")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "handle")
		delete(additionalProperties, "phase_1_alg")
		delete(additionalProperties, "phase_1_dh_group")
		delete(additionalProperties, "phase_1_lifetime")
		delete(additionalProperties, "phase_2_alg")
		delete(additionalProperties, "phase_2_dh_group")
		delete(additionalProperties, "phase_2_lifetime")
		delete(additionalProperties, "perfect_forward_secrecy")
		delete(additionalProperties, "our_gateway")
		delete(additionalProperties, "our_networks")
		delete(additionalProperties, "peer_gateway")
		delete(additionalProperties, "peer_networks")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVpnTunnel struct {
	value *VpnTunnel
	isSet bool
}

func (v NullableVpnTunnel) Get() *VpnTunnel {
	return v.value
}

func (v *NullableVpnTunnel) Set(val *VpnTunnel) {
	v.value = val
	v.isSet = true
}

func (v NullableVpnTunnel) IsSet() bool {
	return v.isSet
}

func (v *NullableVpnTunnel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVpnTunnel(val *VpnTunnel) *NullableVpnTunnel {
	return &NullableVpnTunnel{value: val, isSet: true}
}

func (v NullableVpnTunnel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVpnTunnel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


