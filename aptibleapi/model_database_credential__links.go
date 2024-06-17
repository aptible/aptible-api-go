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

// checks if the DatabaseCredentialLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DatabaseCredentialLinks{}

// DatabaseCredentialLinks struct for DatabaseCredentialLinks
type DatabaseCredentialLinks struct {
	Database *ListAccountsForStack200ResponseLinksStack `json:"database,omitempty"`
	Operations *ListAccountsForStack200ResponseLinksStack `json:"operations,omitempty"`
	Self *ListAccountsForStack200ResponseLinksStack `json:"self,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DatabaseCredentialLinks DatabaseCredentialLinks

// NewDatabaseCredentialLinks instantiates a new DatabaseCredentialLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatabaseCredentialLinks() *DatabaseCredentialLinks {
	this := DatabaseCredentialLinks{}
	return &this
}

// NewDatabaseCredentialLinksWithDefaults instantiates a new DatabaseCredentialLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatabaseCredentialLinksWithDefaults() *DatabaseCredentialLinks {
	this := DatabaseCredentialLinks{}
	return &this
}

// GetDatabase returns the Database field value if set, zero value otherwise.
func (o *DatabaseCredentialLinks) GetDatabase() ListAccountsForStack200ResponseLinksStack {
	if o == nil || IsNil(o.Database) {
		var ret ListAccountsForStack200ResponseLinksStack
		return ret
	}
	return *o.Database
}

// GetDatabaseOk returns a tuple with the Database field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseCredentialLinks) GetDatabaseOk() (*ListAccountsForStack200ResponseLinksStack, bool) {
	if o == nil || IsNil(o.Database) {
		return nil, false
	}
	return o.Database, true
}

// HasDatabase returns a boolean if a field has been set.
func (o *DatabaseCredentialLinks) HasDatabase() bool {
	if o != nil && !IsNil(o.Database) {
		return true
	}

	return false
}

// SetDatabase gets a reference to the given ListAccountsForStack200ResponseLinksStack and assigns it to the Database field.
func (o *DatabaseCredentialLinks) SetDatabase(v ListAccountsForStack200ResponseLinksStack) {
	o.Database = &v
}

// GetOperations returns the Operations field value if set, zero value otherwise.
func (o *DatabaseCredentialLinks) GetOperations() ListAccountsForStack200ResponseLinksStack {
	if o == nil || IsNil(o.Operations) {
		var ret ListAccountsForStack200ResponseLinksStack
		return ret
	}
	return *o.Operations
}

// GetOperationsOk returns a tuple with the Operations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseCredentialLinks) GetOperationsOk() (*ListAccountsForStack200ResponseLinksStack, bool) {
	if o == nil || IsNil(o.Operations) {
		return nil, false
	}
	return o.Operations, true
}

// HasOperations returns a boolean if a field has been set.
func (o *DatabaseCredentialLinks) HasOperations() bool {
	if o != nil && !IsNil(o.Operations) {
		return true
	}

	return false
}

// SetOperations gets a reference to the given ListAccountsForStack200ResponseLinksStack and assigns it to the Operations field.
func (o *DatabaseCredentialLinks) SetOperations(v ListAccountsForStack200ResponseLinksStack) {
	o.Operations = &v
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *DatabaseCredentialLinks) GetSelf() ListAccountsForStack200ResponseLinksStack {
	if o == nil || IsNil(o.Self) {
		var ret ListAccountsForStack200ResponseLinksStack
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseCredentialLinks) GetSelfOk() (*ListAccountsForStack200ResponseLinksStack, bool) {
	if o == nil || IsNil(o.Self) {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *DatabaseCredentialLinks) HasSelf() bool {
	if o != nil && !IsNil(o.Self) {
		return true
	}

	return false
}

// SetSelf gets a reference to the given ListAccountsForStack200ResponseLinksStack and assigns it to the Self field.
func (o *DatabaseCredentialLinks) SetSelf(v ListAccountsForStack200ResponseLinksStack) {
	o.Self = &v
}

func (o DatabaseCredentialLinks) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DatabaseCredentialLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Database) {
		toSerialize["database"] = o.Database
	}
	if !IsNil(o.Operations) {
		toSerialize["operations"] = o.Operations
	}
	if !IsNil(o.Self) {
		toSerialize["self"] = o.Self
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DatabaseCredentialLinks) UnmarshalJSON(data []byte) (err error) {
	varDatabaseCredentialLinks := _DatabaseCredentialLinks{}

	err = json.Unmarshal(data, &varDatabaseCredentialLinks)

	if err != nil {
		return err
	}

	*o = DatabaseCredentialLinks(varDatabaseCredentialLinks)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "database")
		delete(additionalProperties, "operations")
		delete(additionalProperties, "self")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDatabaseCredentialLinks struct {
	value *DatabaseCredentialLinks
	isSet bool
}

func (v NullableDatabaseCredentialLinks) Get() *DatabaseCredentialLinks {
	return v.value
}

func (v *NullableDatabaseCredentialLinks) Set(val *DatabaseCredentialLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableDatabaseCredentialLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableDatabaseCredentialLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatabaseCredentialLinks(val *DatabaseCredentialLinks) *NullableDatabaseCredentialLinks {
	return &NullableDatabaseCredentialLinks{value: val, isSet: true}
}

func (v NullableDatabaseCredentialLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatabaseCredentialLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

