/*
Deploy API v1

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package aptibleapi

import (
	"encoding/json"
)

// checks if the SshPortalConnectionLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SshPortalConnectionLinks{}

// SshPortalConnectionLinks struct for SshPortalConnectionLinks
type SshPortalConnectionLinks struct {
	Operation *ListAccountsForStack200ResponseLinksStack `json:"operation,omitempty"`
	Self *ListAccountsForStack200ResponseLinksStack `json:"self,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SshPortalConnectionLinks SshPortalConnectionLinks

// NewSshPortalConnectionLinks instantiates a new SshPortalConnectionLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSshPortalConnectionLinks() *SshPortalConnectionLinks {
	this := SshPortalConnectionLinks{}
	return &this
}

// NewSshPortalConnectionLinksWithDefaults instantiates a new SshPortalConnectionLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSshPortalConnectionLinksWithDefaults() *SshPortalConnectionLinks {
	this := SshPortalConnectionLinks{}
	return &this
}

// GetOperation returns the Operation field value if set, zero value otherwise.
func (o *SshPortalConnectionLinks) GetOperation() ListAccountsForStack200ResponseLinksStack {
	if o == nil || IsNil(o.Operation) {
		var ret ListAccountsForStack200ResponseLinksStack
		return ret
	}
	return *o.Operation
}

// GetOperationOk returns a tuple with the Operation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SshPortalConnectionLinks) GetOperationOk() (*ListAccountsForStack200ResponseLinksStack, bool) {
	if o == nil || IsNil(o.Operation) {
		return nil, false
	}
	return o.Operation, true
}

// HasOperation returns a boolean if a field has been set.
func (o *SshPortalConnectionLinks) HasOperation() bool {
	if o != nil && !IsNil(o.Operation) {
		return true
	}

	return false
}

// SetOperation gets a reference to the given ListAccountsForStack200ResponseLinksStack and assigns it to the Operation field.
func (o *SshPortalConnectionLinks) SetOperation(v ListAccountsForStack200ResponseLinksStack) {
	o.Operation = &v
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *SshPortalConnectionLinks) GetSelf() ListAccountsForStack200ResponseLinksStack {
	if o == nil || IsNil(o.Self) {
		var ret ListAccountsForStack200ResponseLinksStack
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SshPortalConnectionLinks) GetSelfOk() (*ListAccountsForStack200ResponseLinksStack, bool) {
	if o == nil || IsNil(o.Self) {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *SshPortalConnectionLinks) HasSelf() bool {
	if o != nil && !IsNil(o.Self) {
		return true
	}

	return false
}

// SetSelf gets a reference to the given ListAccountsForStack200ResponseLinksStack and assigns it to the Self field.
func (o *SshPortalConnectionLinks) SetSelf(v ListAccountsForStack200ResponseLinksStack) {
	o.Self = &v
}

func (o SshPortalConnectionLinks) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SshPortalConnectionLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Operation) {
		toSerialize["operation"] = o.Operation
	}
	if !IsNil(o.Self) {
		toSerialize["self"] = o.Self
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SshPortalConnectionLinks) UnmarshalJSON(data []byte) (err error) {
	varSshPortalConnectionLinks := _SshPortalConnectionLinks{}

	err = json.Unmarshal(data, &varSshPortalConnectionLinks)

	if err != nil {
		return err
	}

	*o = SshPortalConnectionLinks(varSshPortalConnectionLinks)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "operation")
		delete(additionalProperties, "self")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSshPortalConnectionLinks struct {
	value *SshPortalConnectionLinks
	isSet bool
}

func (v NullableSshPortalConnectionLinks) Get() *SshPortalConnectionLinks {
	return v.value
}

func (v *NullableSshPortalConnectionLinks) Set(val *SshPortalConnectionLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableSshPortalConnectionLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableSshPortalConnectionLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSshPortalConnectionLinks(val *SshPortalConnectionLinks) *NullableSshPortalConnectionLinks {
	return &NullableSshPortalConnectionLinks{value: val, isSet: true}
}

func (v NullableSshPortalConnectionLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSshPortalConnectionLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


