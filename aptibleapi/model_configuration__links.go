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

// checks if the ConfigurationLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfigurationLinks{}

// ConfigurationLinks struct for ConfigurationLinks
type ConfigurationLinks struct {
	Resource *ListAccountsForStack200ResponseLinksStack `json:"resource,omitempty"`
	Self *ListAccountsForStack200ResponseLinksStack `json:"self,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConfigurationLinks ConfigurationLinks

// NewConfigurationLinks instantiates a new ConfigurationLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigurationLinks() *ConfigurationLinks {
	this := ConfigurationLinks{}
	return &this
}

// NewConfigurationLinksWithDefaults instantiates a new ConfigurationLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigurationLinksWithDefaults() *ConfigurationLinks {
	this := ConfigurationLinks{}
	return &this
}

// GetResource returns the Resource field value if set, zero value otherwise.
func (o *ConfigurationLinks) GetResource() ListAccountsForStack200ResponseLinksStack {
	if o == nil || IsNil(o.Resource) {
		var ret ListAccountsForStack200ResponseLinksStack
		return ret
	}
	return *o.Resource
}

// GetResourceOk returns a tuple with the Resource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationLinks) GetResourceOk() (*ListAccountsForStack200ResponseLinksStack, bool) {
	if o == nil || IsNil(o.Resource) {
		return nil, false
	}
	return o.Resource, true
}

// HasResource returns a boolean if a field has been set.
func (o *ConfigurationLinks) HasResource() bool {
	if o != nil && !IsNil(o.Resource) {
		return true
	}

	return false
}

// SetResource gets a reference to the given ListAccountsForStack200ResponseLinksStack and assigns it to the Resource field.
func (o *ConfigurationLinks) SetResource(v ListAccountsForStack200ResponseLinksStack) {
	o.Resource = &v
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *ConfigurationLinks) GetSelf() ListAccountsForStack200ResponseLinksStack {
	if o == nil || IsNil(o.Self) {
		var ret ListAccountsForStack200ResponseLinksStack
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationLinks) GetSelfOk() (*ListAccountsForStack200ResponseLinksStack, bool) {
	if o == nil || IsNil(o.Self) {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *ConfigurationLinks) HasSelf() bool {
	if o != nil && !IsNil(o.Self) {
		return true
	}

	return false
}

// SetSelf gets a reference to the given ListAccountsForStack200ResponseLinksStack and assigns it to the Self field.
func (o *ConfigurationLinks) SetSelf(v ListAccountsForStack200ResponseLinksStack) {
	o.Self = &v
}

func (o ConfigurationLinks) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConfigurationLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Resource) {
		toSerialize["resource"] = o.Resource
	}
	if !IsNil(o.Self) {
		toSerialize["self"] = o.Self
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ConfigurationLinks) UnmarshalJSON(data []byte) (err error) {
	varConfigurationLinks := _ConfigurationLinks{}

	err = json.Unmarshal(data, &varConfigurationLinks)

	if err != nil {
		return err
	}

	*o = ConfigurationLinks(varConfigurationLinks)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "resource")
		delete(additionalProperties, "self")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableConfigurationLinks struct {
	value *ConfigurationLinks
	isSet bool
}

func (v NullableConfigurationLinks) Get() *ConfigurationLinks {
	return v.value
}

func (v *NullableConfigurationLinks) Set(val *ConfigurationLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigurationLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigurationLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigurationLinks(val *ConfigurationLinks) *NullableConfigurationLinks {
	return &NullableConfigurationLinks{value: val, isSet: true}
}

func (v NullableConfigurationLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigurationLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


