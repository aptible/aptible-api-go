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

// checks if the BackupEmbedded type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BackupEmbedded{}

// BackupEmbedded struct for BackupEmbedded
type BackupEmbedded struct {
	CopiedFrom map[string]interface{} `json:"copied_from,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BackupEmbedded BackupEmbedded

// NewBackupEmbedded instantiates a new BackupEmbedded object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBackupEmbedded() *BackupEmbedded {
	this := BackupEmbedded{}
	return &this
}

// NewBackupEmbeddedWithDefaults instantiates a new BackupEmbedded object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBackupEmbeddedWithDefaults() *BackupEmbedded {
	this := BackupEmbedded{}
	return &this
}

// GetCopiedFrom returns the CopiedFrom field value if set, zero value otherwise.
func (o *BackupEmbedded) GetCopiedFrom() map[string]interface{} {
	if o == nil || IsNil(o.CopiedFrom) {
		var ret map[string]interface{}
		return ret
	}
	return o.CopiedFrom
}

// GetCopiedFromOk returns a tuple with the CopiedFrom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupEmbedded) GetCopiedFromOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CopiedFrom) {
		return map[string]interface{}{}, false
	}
	return o.CopiedFrom, true
}

// HasCopiedFrom returns a boolean if a field has been set.
func (o *BackupEmbedded) HasCopiedFrom() bool {
	if o != nil && !IsNil(o.CopiedFrom) {
		return true
	}

	return false
}

// SetCopiedFrom gets a reference to the given map[string]interface{} and assigns it to the CopiedFrom field.
func (o *BackupEmbedded) SetCopiedFrom(v map[string]interface{}) {
	o.CopiedFrom = v
}

func (o BackupEmbedded) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BackupEmbedded) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CopiedFrom) {
		toSerialize["copied_from"] = o.CopiedFrom
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BackupEmbedded) UnmarshalJSON(data []byte) (err error) {
	varBackupEmbedded := _BackupEmbedded{}

	err = json.Unmarshal(data, &varBackupEmbedded)

	if err != nil {
		return err
	}

	*o = BackupEmbedded(varBackupEmbedded)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "copied_from")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBackupEmbedded struct {
	value *BackupEmbedded
	isSet bool
}

func (v NullableBackupEmbedded) Get() *BackupEmbedded {
	return v.value
}

func (v *NullableBackupEmbedded) Set(val *BackupEmbedded) {
	v.value = val
	v.isSet = true
}

func (v NullableBackupEmbedded) IsSet() bool {
	return v.isSet
}

func (v *NullableBackupEmbedded) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackupEmbedded(val *BackupEmbedded) *NullableBackupEmbedded {
	return &NullableBackupEmbedded{value: val, isSet: true}
}

func (v NullableBackupEmbedded) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackupEmbedded) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

