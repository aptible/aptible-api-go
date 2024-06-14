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

// checks if the ListDatabaseCredentialsForDatabase200ResponseLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListDatabaseCredentialsForDatabase200ResponseLinks{}

// ListDatabaseCredentialsForDatabase200ResponseLinks struct for ListDatabaseCredentialsForDatabase200ResponseLinks
type ListDatabaseCredentialsForDatabase200ResponseLinks struct {
	Database *ListAccountsForStack200ResponseLinksStack `json:"database,omitempty"`
	Next *ListAccountsForStack200ResponseLinksStack `json:"next,omitempty"`
	Prev *ListAccountsForStack200ResponseLinksStack `json:"prev,omitempty"`
	Self *ListAccountsForStack200ResponseLinksStack `json:"self,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ListDatabaseCredentialsForDatabase200ResponseLinks ListDatabaseCredentialsForDatabase200ResponseLinks

// NewListDatabaseCredentialsForDatabase200ResponseLinks instantiates a new ListDatabaseCredentialsForDatabase200ResponseLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListDatabaseCredentialsForDatabase200ResponseLinks() *ListDatabaseCredentialsForDatabase200ResponseLinks {
	this := ListDatabaseCredentialsForDatabase200ResponseLinks{}
	return &this
}

// NewListDatabaseCredentialsForDatabase200ResponseLinksWithDefaults instantiates a new ListDatabaseCredentialsForDatabase200ResponseLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListDatabaseCredentialsForDatabase200ResponseLinksWithDefaults() *ListDatabaseCredentialsForDatabase200ResponseLinks {
	this := ListDatabaseCredentialsForDatabase200ResponseLinks{}
	return &this
}

// GetDatabase returns the Database field value if set, zero value otherwise.
func (o *ListDatabaseCredentialsForDatabase200ResponseLinks) GetDatabase() ListAccountsForStack200ResponseLinksStack {
	if o == nil || IsNil(o.Database) {
		var ret ListAccountsForStack200ResponseLinksStack
		return ret
	}
	return *o.Database
}

// GetDatabaseOk returns a tuple with the Database field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListDatabaseCredentialsForDatabase200ResponseLinks) GetDatabaseOk() (*ListAccountsForStack200ResponseLinksStack, bool) {
	if o == nil || IsNil(o.Database) {
		return nil, false
	}
	return o.Database, true
}

// HasDatabase returns a boolean if a field has been set.
func (o *ListDatabaseCredentialsForDatabase200ResponseLinks) HasDatabase() bool {
	if o != nil && !IsNil(o.Database) {
		return true
	}

	return false
}

// SetDatabase gets a reference to the given ListAccountsForStack200ResponseLinksStack and assigns it to the Database field.
func (o *ListDatabaseCredentialsForDatabase200ResponseLinks) SetDatabase(v ListAccountsForStack200ResponseLinksStack) {
	o.Database = &v
}

// GetNext returns the Next field value if set, zero value otherwise.
func (o *ListDatabaseCredentialsForDatabase200ResponseLinks) GetNext() ListAccountsForStack200ResponseLinksStack {
	if o == nil || IsNil(o.Next) {
		var ret ListAccountsForStack200ResponseLinksStack
		return ret
	}
	return *o.Next
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListDatabaseCredentialsForDatabase200ResponseLinks) GetNextOk() (*ListAccountsForStack200ResponseLinksStack, bool) {
	if o == nil || IsNil(o.Next) {
		return nil, false
	}
	return o.Next, true
}

// HasNext returns a boolean if a field has been set.
func (o *ListDatabaseCredentialsForDatabase200ResponseLinks) HasNext() bool {
	if o != nil && !IsNil(o.Next) {
		return true
	}

	return false
}

// SetNext gets a reference to the given ListAccountsForStack200ResponseLinksStack and assigns it to the Next field.
func (o *ListDatabaseCredentialsForDatabase200ResponseLinks) SetNext(v ListAccountsForStack200ResponseLinksStack) {
	o.Next = &v
}

// GetPrev returns the Prev field value if set, zero value otherwise.
func (o *ListDatabaseCredentialsForDatabase200ResponseLinks) GetPrev() ListAccountsForStack200ResponseLinksStack {
	if o == nil || IsNil(o.Prev) {
		var ret ListAccountsForStack200ResponseLinksStack
		return ret
	}
	return *o.Prev
}

// GetPrevOk returns a tuple with the Prev field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListDatabaseCredentialsForDatabase200ResponseLinks) GetPrevOk() (*ListAccountsForStack200ResponseLinksStack, bool) {
	if o == nil || IsNil(o.Prev) {
		return nil, false
	}
	return o.Prev, true
}

// HasPrev returns a boolean if a field has been set.
func (o *ListDatabaseCredentialsForDatabase200ResponseLinks) HasPrev() bool {
	if o != nil && !IsNil(o.Prev) {
		return true
	}

	return false
}

// SetPrev gets a reference to the given ListAccountsForStack200ResponseLinksStack and assigns it to the Prev field.
func (o *ListDatabaseCredentialsForDatabase200ResponseLinks) SetPrev(v ListAccountsForStack200ResponseLinksStack) {
	o.Prev = &v
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *ListDatabaseCredentialsForDatabase200ResponseLinks) GetSelf() ListAccountsForStack200ResponseLinksStack {
	if o == nil || IsNil(o.Self) {
		var ret ListAccountsForStack200ResponseLinksStack
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListDatabaseCredentialsForDatabase200ResponseLinks) GetSelfOk() (*ListAccountsForStack200ResponseLinksStack, bool) {
	if o == nil || IsNil(o.Self) {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *ListDatabaseCredentialsForDatabase200ResponseLinks) HasSelf() bool {
	if o != nil && !IsNil(o.Self) {
		return true
	}

	return false
}

// SetSelf gets a reference to the given ListAccountsForStack200ResponseLinksStack and assigns it to the Self field.
func (o *ListDatabaseCredentialsForDatabase200ResponseLinks) SetSelf(v ListAccountsForStack200ResponseLinksStack) {
	o.Self = &v
}

func (o ListDatabaseCredentialsForDatabase200ResponseLinks) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListDatabaseCredentialsForDatabase200ResponseLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Database) {
		toSerialize["database"] = o.Database
	}
	if !IsNil(o.Next) {
		toSerialize["next"] = o.Next
	}
	if !IsNil(o.Prev) {
		toSerialize["prev"] = o.Prev
	}
	if !IsNil(o.Self) {
		toSerialize["self"] = o.Self
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListDatabaseCredentialsForDatabase200ResponseLinks) UnmarshalJSON(data []byte) (err error) {
	varListDatabaseCredentialsForDatabase200ResponseLinks := _ListDatabaseCredentialsForDatabase200ResponseLinks{}

	err = json.Unmarshal(data, &varListDatabaseCredentialsForDatabase200ResponseLinks)

	if err != nil {
		return err
	}

	*o = ListDatabaseCredentialsForDatabase200ResponseLinks(varListDatabaseCredentialsForDatabase200ResponseLinks)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "database")
		delete(additionalProperties, "next")
		delete(additionalProperties, "prev")
		delete(additionalProperties, "self")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListDatabaseCredentialsForDatabase200ResponseLinks struct {
	value *ListDatabaseCredentialsForDatabase200ResponseLinks
	isSet bool
}

func (v NullableListDatabaseCredentialsForDatabase200ResponseLinks) Get() *ListDatabaseCredentialsForDatabase200ResponseLinks {
	return v.value
}

func (v *NullableListDatabaseCredentialsForDatabase200ResponseLinks) Set(val *ListDatabaseCredentialsForDatabase200ResponseLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableListDatabaseCredentialsForDatabase200ResponseLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableListDatabaseCredentialsForDatabase200ResponseLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListDatabaseCredentialsForDatabase200ResponseLinks(val *ListDatabaseCredentialsForDatabase200ResponseLinks) *NullableListDatabaseCredentialsForDatabase200ResponseLinks {
	return &NullableListDatabaseCredentialsForDatabase200ResponseLinks{value: val, isSet: true}
}

func (v NullableListDatabaseCredentialsForDatabase200ResponseLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListDatabaseCredentialsForDatabase200ResponseLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


