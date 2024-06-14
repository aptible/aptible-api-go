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

// checks if the ListDatabaseCredentialsForDatabase200ResponseEmbedded type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListDatabaseCredentialsForDatabase200ResponseEmbedded{}

// ListDatabaseCredentialsForDatabase200ResponseEmbedded struct for ListDatabaseCredentialsForDatabase200ResponseEmbedded
type ListDatabaseCredentialsForDatabase200ResponseEmbedded struct {
	DatabaseCredentials []DatabaseCredential `json:"database_credentials,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ListDatabaseCredentialsForDatabase200ResponseEmbedded ListDatabaseCredentialsForDatabase200ResponseEmbedded

// NewListDatabaseCredentialsForDatabase200ResponseEmbedded instantiates a new ListDatabaseCredentialsForDatabase200ResponseEmbedded object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListDatabaseCredentialsForDatabase200ResponseEmbedded() *ListDatabaseCredentialsForDatabase200ResponseEmbedded {
	this := ListDatabaseCredentialsForDatabase200ResponseEmbedded{}
	return &this
}

// NewListDatabaseCredentialsForDatabase200ResponseEmbeddedWithDefaults instantiates a new ListDatabaseCredentialsForDatabase200ResponseEmbedded object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListDatabaseCredentialsForDatabase200ResponseEmbeddedWithDefaults() *ListDatabaseCredentialsForDatabase200ResponseEmbedded {
	this := ListDatabaseCredentialsForDatabase200ResponseEmbedded{}
	return &this
}

// GetDatabaseCredentials returns the DatabaseCredentials field value if set, zero value otherwise.
func (o *ListDatabaseCredentialsForDatabase200ResponseEmbedded) GetDatabaseCredentials() []DatabaseCredential {
	if o == nil || IsNil(o.DatabaseCredentials) {
		var ret []DatabaseCredential
		return ret
	}
	return o.DatabaseCredentials
}

// GetDatabaseCredentialsOk returns a tuple with the DatabaseCredentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListDatabaseCredentialsForDatabase200ResponseEmbedded) GetDatabaseCredentialsOk() ([]DatabaseCredential, bool) {
	if o == nil || IsNil(o.DatabaseCredentials) {
		return nil, false
	}
	return o.DatabaseCredentials, true
}

// HasDatabaseCredentials returns a boolean if a field has been set.
func (o *ListDatabaseCredentialsForDatabase200ResponseEmbedded) HasDatabaseCredentials() bool {
	if o != nil && !IsNil(o.DatabaseCredentials) {
		return true
	}

	return false
}

// SetDatabaseCredentials gets a reference to the given []DatabaseCredential and assigns it to the DatabaseCredentials field.
func (o *ListDatabaseCredentialsForDatabase200ResponseEmbedded) SetDatabaseCredentials(v []DatabaseCredential) {
	o.DatabaseCredentials = v
}

func (o ListDatabaseCredentialsForDatabase200ResponseEmbedded) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListDatabaseCredentialsForDatabase200ResponseEmbedded) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DatabaseCredentials) {
		toSerialize["database_credentials"] = o.DatabaseCredentials
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListDatabaseCredentialsForDatabase200ResponseEmbedded) UnmarshalJSON(data []byte) (err error) {
	varListDatabaseCredentialsForDatabase200ResponseEmbedded := _ListDatabaseCredentialsForDatabase200ResponseEmbedded{}

	err = json.Unmarshal(data, &varListDatabaseCredentialsForDatabase200ResponseEmbedded)

	if err != nil {
		return err
	}

	*o = ListDatabaseCredentialsForDatabase200ResponseEmbedded(varListDatabaseCredentialsForDatabase200ResponseEmbedded)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "database_credentials")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListDatabaseCredentialsForDatabase200ResponseEmbedded struct {
	value *ListDatabaseCredentialsForDatabase200ResponseEmbedded
	isSet bool
}

func (v NullableListDatabaseCredentialsForDatabase200ResponseEmbedded) Get() *ListDatabaseCredentialsForDatabase200ResponseEmbedded {
	return v.value
}

func (v *NullableListDatabaseCredentialsForDatabase200ResponseEmbedded) Set(val *ListDatabaseCredentialsForDatabase200ResponseEmbedded) {
	v.value = val
	v.isSet = true
}

func (v NullableListDatabaseCredentialsForDatabase200ResponseEmbedded) IsSet() bool {
	return v.isSet
}

func (v *NullableListDatabaseCredentialsForDatabase200ResponseEmbedded) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListDatabaseCredentialsForDatabase200ResponseEmbedded(val *ListDatabaseCredentialsForDatabase200ResponseEmbedded) *NullableListDatabaseCredentialsForDatabase200ResponseEmbedded {
	return &NullableListDatabaseCredentialsForDatabase200ResponseEmbedded{value: val, isSet: true}
}

func (v NullableListDatabaseCredentialsForDatabase200ResponseEmbedded) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListDatabaseCredentialsForDatabase200ResponseEmbedded) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


