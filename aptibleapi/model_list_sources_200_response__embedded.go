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

// checks if the ListSources200ResponseEmbedded type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListSources200ResponseEmbedded{}

// ListSources200ResponseEmbedded struct for ListSources200ResponseEmbedded
type ListSources200ResponseEmbedded struct {
	Sources []Source `json:"sources,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ListSources200ResponseEmbedded ListSources200ResponseEmbedded

// NewListSources200ResponseEmbedded instantiates a new ListSources200ResponseEmbedded object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListSources200ResponseEmbedded() *ListSources200ResponseEmbedded {
	this := ListSources200ResponseEmbedded{}
	return &this
}

// NewListSources200ResponseEmbeddedWithDefaults instantiates a new ListSources200ResponseEmbedded object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListSources200ResponseEmbeddedWithDefaults() *ListSources200ResponseEmbedded {
	this := ListSources200ResponseEmbedded{}
	return &this
}

// GetSources returns the Sources field value if set, zero value otherwise.
func (o *ListSources200ResponseEmbedded) GetSources() []Source {
	if o == nil || IsNil(o.Sources) {
		var ret []Source
		return ret
	}
	return o.Sources
}

// GetSourcesOk returns a tuple with the Sources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSources200ResponseEmbedded) GetSourcesOk() ([]Source, bool) {
	if o == nil || IsNil(o.Sources) {
		return nil, false
	}
	return o.Sources, true
}

// HasSources returns a boolean if a field has been set.
func (o *ListSources200ResponseEmbedded) HasSources() bool {
	if o != nil && !IsNil(o.Sources) {
		return true
	}

	return false
}

// SetSources gets a reference to the given []Source and assigns it to the Sources field.
func (o *ListSources200ResponseEmbedded) SetSources(v []Source) {
	o.Sources = v
}

func (o ListSources200ResponseEmbedded) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListSources200ResponseEmbedded) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Sources) {
		toSerialize["sources"] = o.Sources
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListSources200ResponseEmbedded) UnmarshalJSON(data []byte) (err error) {
	varListSources200ResponseEmbedded := _ListSources200ResponseEmbedded{}

	err = json.Unmarshal(data, &varListSources200ResponseEmbedded)

	if err != nil {
		return err
	}

	*o = ListSources200ResponseEmbedded(varListSources200ResponseEmbedded)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "sources")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListSources200ResponseEmbedded struct {
	value *ListSources200ResponseEmbedded
	isSet bool
}

func (v NullableListSources200ResponseEmbedded) Get() *ListSources200ResponseEmbedded {
	return v.value
}

func (v *NullableListSources200ResponseEmbedded) Set(val *ListSources200ResponseEmbedded) {
	v.value = val
	v.isSet = true
}

func (v NullableListSources200ResponseEmbedded) IsSet() bool {
	return v.isSet
}

func (v *NullableListSources200ResponseEmbedded) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListSources200ResponseEmbedded(val *ListSources200ResponseEmbedded) *NullableListSources200ResponseEmbedded {
	return &NullableListSources200ResponseEmbedded{value: val, isSet: true}
}

func (v NullableListSources200ResponseEmbedded) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListSources200ResponseEmbedded) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


