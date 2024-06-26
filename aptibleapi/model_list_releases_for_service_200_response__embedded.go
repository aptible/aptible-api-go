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

// checks if the ListReleasesForService200ResponseEmbedded type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListReleasesForService200ResponseEmbedded{}

// ListReleasesForService200ResponseEmbedded struct for ListReleasesForService200ResponseEmbedded
type ListReleasesForService200ResponseEmbedded struct {
	Releases []Release `json:"releases,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ListReleasesForService200ResponseEmbedded ListReleasesForService200ResponseEmbedded

// NewListReleasesForService200ResponseEmbedded instantiates a new ListReleasesForService200ResponseEmbedded object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListReleasesForService200ResponseEmbedded() *ListReleasesForService200ResponseEmbedded {
	this := ListReleasesForService200ResponseEmbedded{}
	return &this
}

// NewListReleasesForService200ResponseEmbeddedWithDefaults instantiates a new ListReleasesForService200ResponseEmbedded object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListReleasesForService200ResponseEmbeddedWithDefaults() *ListReleasesForService200ResponseEmbedded {
	this := ListReleasesForService200ResponseEmbedded{}
	return &this
}

// GetReleases returns the Releases field value if set, zero value otherwise.
func (o *ListReleasesForService200ResponseEmbedded) GetReleases() []Release {
	if o == nil || IsNil(o.Releases) {
		var ret []Release
		return ret
	}
	return o.Releases
}

// GetReleasesOk returns a tuple with the Releases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListReleasesForService200ResponseEmbedded) GetReleasesOk() ([]Release, bool) {
	if o == nil || IsNil(o.Releases) {
		return nil, false
	}
	return o.Releases, true
}

// HasReleases returns a boolean if a field has been set.
func (o *ListReleasesForService200ResponseEmbedded) HasReleases() bool {
	if o != nil && !IsNil(o.Releases) {
		return true
	}

	return false
}

// SetReleases gets a reference to the given []Release and assigns it to the Releases field.
func (o *ListReleasesForService200ResponseEmbedded) SetReleases(v []Release) {
	o.Releases = v
}

func (o ListReleasesForService200ResponseEmbedded) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListReleasesForService200ResponseEmbedded) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Releases) {
		toSerialize["releases"] = o.Releases
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListReleasesForService200ResponseEmbedded) UnmarshalJSON(data []byte) (err error) {
	varListReleasesForService200ResponseEmbedded := _ListReleasesForService200ResponseEmbedded{}

	err = json.Unmarshal(data, &varListReleasesForService200ResponseEmbedded)

	if err != nil {
		return err
	}

	*o = ListReleasesForService200ResponseEmbedded(varListReleasesForService200ResponseEmbedded)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "releases")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListReleasesForService200ResponseEmbedded struct {
	value *ListReleasesForService200ResponseEmbedded
	isSet bool
}

func (v NullableListReleasesForService200ResponseEmbedded) Get() *ListReleasesForService200ResponseEmbedded {
	return v.value
}

func (v *NullableListReleasesForService200ResponseEmbedded) Set(val *ListReleasesForService200ResponseEmbedded) {
	v.value = val
	v.isSet = true
}

func (v NullableListReleasesForService200ResponseEmbedded) IsSet() bool {
	return v.isSet
}

func (v *NullableListReleasesForService200ResponseEmbedded) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListReleasesForService200ResponseEmbedded(val *ListReleasesForService200ResponseEmbedded) *NullableListReleasesForService200ResponseEmbedded {
	return &NullableListReleasesForService200ResponseEmbedded{value: val, isSet: true}
}

func (v NullableListReleasesForService200ResponseEmbedded) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListReleasesForService200ResponseEmbedded) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


