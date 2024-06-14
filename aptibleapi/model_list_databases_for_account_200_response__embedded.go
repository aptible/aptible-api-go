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

// checks if the ListDatabasesForAccount200ResponseEmbedded type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListDatabasesForAccount200ResponseEmbedded{}

// ListDatabasesForAccount200ResponseEmbedded struct for ListDatabasesForAccount200ResponseEmbedded
type ListDatabasesForAccount200ResponseEmbedded struct {
	Databases []Database `json:"databases,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ListDatabasesForAccount200ResponseEmbedded ListDatabasesForAccount200ResponseEmbedded

// NewListDatabasesForAccount200ResponseEmbedded instantiates a new ListDatabasesForAccount200ResponseEmbedded object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListDatabasesForAccount200ResponseEmbedded() *ListDatabasesForAccount200ResponseEmbedded {
	this := ListDatabasesForAccount200ResponseEmbedded{}
	return &this
}

// NewListDatabasesForAccount200ResponseEmbeddedWithDefaults instantiates a new ListDatabasesForAccount200ResponseEmbedded object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListDatabasesForAccount200ResponseEmbeddedWithDefaults() *ListDatabasesForAccount200ResponseEmbedded {
	this := ListDatabasesForAccount200ResponseEmbedded{}
	return &this
}

// GetDatabases returns the Databases field value if set, zero value otherwise.
func (o *ListDatabasesForAccount200ResponseEmbedded) GetDatabases() []Database {
	if o == nil || IsNil(o.Databases) {
		var ret []Database
		return ret
	}
	return o.Databases
}

// GetDatabasesOk returns a tuple with the Databases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListDatabasesForAccount200ResponseEmbedded) GetDatabasesOk() ([]Database, bool) {
	if o == nil || IsNil(o.Databases) {
		return nil, false
	}
	return o.Databases, true
}

// HasDatabases returns a boolean if a field has been set.
func (o *ListDatabasesForAccount200ResponseEmbedded) HasDatabases() bool {
	if o != nil && !IsNil(o.Databases) {
		return true
	}

	return false
}

// SetDatabases gets a reference to the given []Database and assigns it to the Databases field.
func (o *ListDatabasesForAccount200ResponseEmbedded) SetDatabases(v []Database) {
	o.Databases = v
}

func (o ListDatabasesForAccount200ResponseEmbedded) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListDatabasesForAccount200ResponseEmbedded) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Databases) {
		toSerialize["databases"] = o.Databases
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListDatabasesForAccount200ResponseEmbedded) UnmarshalJSON(data []byte) (err error) {
	varListDatabasesForAccount200ResponseEmbedded := _ListDatabasesForAccount200ResponseEmbedded{}

	err = json.Unmarshal(data, &varListDatabasesForAccount200ResponseEmbedded)

	if err != nil {
		return err
	}

	*o = ListDatabasesForAccount200ResponseEmbedded(varListDatabasesForAccount200ResponseEmbedded)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "databases")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListDatabasesForAccount200ResponseEmbedded struct {
	value *ListDatabasesForAccount200ResponseEmbedded
	isSet bool
}

func (v NullableListDatabasesForAccount200ResponseEmbedded) Get() *ListDatabasesForAccount200ResponseEmbedded {
	return v.value
}

func (v *NullableListDatabasesForAccount200ResponseEmbedded) Set(val *ListDatabasesForAccount200ResponseEmbedded) {
	v.value = val
	v.isSet = true
}

func (v NullableListDatabasesForAccount200ResponseEmbedded) IsSet() bool {
	return v.isSet
}

func (v *NullableListDatabasesForAccount200ResponseEmbedded) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListDatabasesForAccount200ResponseEmbedded(val *ListDatabasesForAccount200ResponseEmbedded) *NullableListDatabasesForAccount200ResponseEmbedded {
	return &NullableListDatabasesForAccount200ResponseEmbedded{value: val, isSet: true}
}

func (v NullableListDatabasesForAccount200ResponseEmbedded) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListDatabasesForAccount200ResponseEmbedded) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


