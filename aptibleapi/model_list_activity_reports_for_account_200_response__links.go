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

// checks if the ListActivityReportsForAccount200ResponseLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListActivityReportsForAccount200ResponseLinks{}

// ListActivityReportsForAccount200ResponseLinks struct for ListActivityReportsForAccount200ResponseLinks
type ListActivityReportsForAccount200ResponseLinks struct {
	Account *ListAccountsForStack200ResponseLinksStack `json:"account,omitempty"`
	Next *ListAccountsForStack200ResponseLinksStack `json:"next,omitempty"`
	Prev *ListAccountsForStack200ResponseLinksStack `json:"prev,omitempty"`
	Self *ListAccountsForStack200ResponseLinksStack `json:"self,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ListActivityReportsForAccount200ResponseLinks ListActivityReportsForAccount200ResponseLinks

// NewListActivityReportsForAccount200ResponseLinks instantiates a new ListActivityReportsForAccount200ResponseLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListActivityReportsForAccount200ResponseLinks() *ListActivityReportsForAccount200ResponseLinks {
	this := ListActivityReportsForAccount200ResponseLinks{}
	return &this
}

// NewListActivityReportsForAccount200ResponseLinksWithDefaults instantiates a new ListActivityReportsForAccount200ResponseLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListActivityReportsForAccount200ResponseLinksWithDefaults() *ListActivityReportsForAccount200ResponseLinks {
	this := ListActivityReportsForAccount200ResponseLinks{}
	return &this
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *ListActivityReportsForAccount200ResponseLinks) GetAccount() ListAccountsForStack200ResponseLinksStack {
	if o == nil || IsNil(o.Account) {
		var ret ListAccountsForStack200ResponseLinksStack
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListActivityReportsForAccount200ResponseLinks) GetAccountOk() (*ListAccountsForStack200ResponseLinksStack, bool) {
	if o == nil || IsNil(o.Account) {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *ListActivityReportsForAccount200ResponseLinks) HasAccount() bool {
	if o != nil && !IsNil(o.Account) {
		return true
	}

	return false
}

// SetAccount gets a reference to the given ListAccountsForStack200ResponseLinksStack and assigns it to the Account field.
func (o *ListActivityReportsForAccount200ResponseLinks) SetAccount(v ListAccountsForStack200ResponseLinksStack) {
	o.Account = &v
}

// GetNext returns the Next field value if set, zero value otherwise.
func (o *ListActivityReportsForAccount200ResponseLinks) GetNext() ListAccountsForStack200ResponseLinksStack {
	if o == nil || IsNil(o.Next) {
		var ret ListAccountsForStack200ResponseLinksStack
		return ret
	}
	return *o.Next
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListActivityReportsForAccount200ResponseLinks) GetNextOk() (*ListAccountsForStack200ResponseLinksStack, bool) {
	if o == nil || IsNil(o.Next) {
		return nil, false
	}
	return o.Next, true
}

// HasNext returns a boolean if a field has been set.
func (o *ListActivityReportsForAccount200ResponseLinks) HasNext() bool {
	if o != nil && !IsNil(o.Next) {
		return true
	}

	return false
}

// SetNext gets a reference to the given ListAccountsForStack200ResponseLinksStack and assigns it to the Next field.
func (o *ListActivityReportsForAccount200ResponseLinks) SetNext(v ListAccountsForStack200ResponseLinksStack) {
	o.Next = &v
}

// GetPrev returns the Prev field value if set, zero value otherwise.
func (o *ListActivityReportsForAccount200ResponseLinks) GetPrev() ListAccountsForStack200ResponseLinksStack {
	if o == nil || IsNil(o.Prev) {
		var ret ListAccountsForStack200ResponseLinksStack
		return ret
	}
	return *o.Prev
}

// GetPrevOk returns a tuple with the Prev field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListActivityReportsForAccount200ResponseLinks) GetPrevOk() (*ListAccountsForStack200ResponseLinksStack, bool) {
	if o == nil || IsNil(o.Prev) {
		return nil, false
	}
	return o.Prev, true
}

// HasPrev returns a boolean if a field has been set.
func (o *ListActivityReportsForAccount200ResponseLinks) HasPrev() bool {
	if o != nil && !IsNil(o.Prev) {
		return true
	}

	return false
}

// SetPrev gets a reference to the given ListAccountsForStack200ResponseLinksStack and assigns it to the Prev field.
func (o *ListActivityReportsForAccount200ResponseLinks) SetPrev(v ListAccountsForStack200ResponseLinksStack) {
	o.Prev = &v
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *ListActivityReportsForAccount200ResponseLinks) GetSelf() ListAccountsForStack200ResponseLinksStack {
	if o == nil || IsNil(o.Self) {
		var ret ListAccountsForStack200ResponseLinksStack
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListActivityReportsForAccount200ResponseLinks) GetSelfOk() (*ListAccountsForStack200ResponseLinksStack, bool) {
	if o == nil || IsNil(o.Self) {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *ListActivityReportsForAccount200ResponseLinks) HasSelf() bool {
	if o != nil && !IsNil(o.Self) {
		return true
	}

	return false
}

// SetSelf gets a reference to the given ListAccountsForStack200ResponseLinksStack and assigns it to the Self field.
func (o *ListActivityReportsForAccount200ResponseLinks) SetSelf(v ListAccountsForStack200ResponseLinksStack) {
	o.Self = &v
}

func (o ListActivityReportsForAccount200ResponseLinks) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListActivityReportsForAccount200ResponseLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Account) {
		toSerialize["account"] = o.Account
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

func (o *ListActivityReportsForAccount200ResponseLinks) UnmarshalJSON(data []byte) (err error) {
	varListActivityReportsForAccount200ResponseLinks := _ListActivityReportsForAccount200ResponseLinks{}

	err = json.Unmarshal(data, &varListActivityReportsForAccount200ResponseLinks)

	if err != nil {
		return err
	}

	*o = ListActivityReportsForAccount200ResponseLinks(varListActivityReportsForAccount200ResponseLinks)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "account")
		delete(additionalProperties, "next")
		delete(additionalProperties, "prev")
		delete(additionalProperties, "self")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListActivityReportsForAccount200ResponseLinks struct {
	value *ListActivityReportsForAccount200ResponseLinks
	isSet bool
}

func (v NullableListActivityReportsForAccount200ResponseLinks) Get() *ListActivityReportsForAccount200ResponseLinks {
	return v.value
}

func (v *NullableListActivityReportsForAccount200ResponseLinks) Set(val *ListActivityReportsForAccount200ResponseLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableListActivityReportsForAccount200ResponseLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableListActivityReportsForAccount200ResponseLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListActivityReportsForAccount200ResponseLinks(val *ListActivityReportsForAccount200ResponseLinks) *NullableListActivityReportsForAccount200ResponseLinks {
	return &NullableListActivityReportsForAccount200ResponseLinks{value: val, isSet: true}
}

func (v NullableListActivityReportsForAccount200ResponseLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListActivityReportsForAccount200ResponseLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


