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

// checks if the ListDiskAttachmentsForPersistentDisk200ResponseLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListDiskAttachmentsForPersistentDisk200ResponseLinks{}

// ListDiskAttachmentsForPersistentDisk200ResponseLinks struct for ListDiskAttachmentsForPersistentDisk200ResponseLinks
type ListDiskAttachmentsForPersistentDisk200ResponseLinks struct {
	PersistentDisk *ListAccountsForStack200ResponseLinksStack `json:"persistent_disk,omitempty"`
	Next *ListAccountsForStack200ResponseLinksStack `json:"next,omitempty"`
	Prev *ListAccountsForStack200ResponseLinksStack `json:"prev,omitempty"`
	Self *ListAccountsForStack200ResponseLinksStack `json:"self,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ListDiskAttachmentsForPersistentDisk200ResponseLinks ListDiskAttachmentsForPersistentDisk200ResponseLinks

// NewListDiskAttachmentsForPersistentDisk200ResponseLinks instantiates a new ListDiskAttachmentsForPersistentDisk200ResponseLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListDiskAttachmentsForPersistentDisk200ResponseLinks() *ListDiskAttachmentsForPersistentDisk200ResponseLinks {
	this := ListDiskAttachmentsForPersistentDisk200ResponseLinks{}
	return &this
}

// NewListDiskAttachmentsForPersistentDisk200ResponseLinksWithDefaults instantiates a new ListDiskAttachmentsForPersistentDisk200ResponseLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListDiskAttachmentsForPersistentDisk200ResponseLinksWithDefaults() *ListDiskAttachmentsForPersistentDisk200ResponseLinks {
	this := ListDiskAttachmentsForPersistentDisk200ResponseLinks{}
	return &this
}

// GetPersistentDisk returns the PersistentDisk field value if set, zero value otherwise.
func (o *ListDiskAttachmentsForPersistentDisk200ResponseLinks) GetPersistentDisk() ListAccountsForStack200ResponseLinksStack {
	if o == nil || IsNil(o.PersistentDisk) {
		var ret ListAccountsForStack200ResponseLinksStack
		return ret
	}
	return *o.PersistentDisk
}

// GetPersistentDiskOk returns a tuple with the PersistentDisk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListDiskAttachmentsForPersistentDisk200ResponseLinks) GetPersistentDiskOk() (*ListAccountsForStack200ResponseLinksStack, bool) {
	if o == nil || IsNil(o.PersistentDisk) {
		return nil, false
	}
	return o.PersistentDisk, true
}

// HasPersistentDisk returns a boolean if a field has been set.
func (o *ListDiskAttachmentsForPersistentDisk200ResponseLinks) HasPersistentDisk() bool {
	if o != nil && !IsNil(o.PersistentDisk) {
		return true
	}

	return false
}

// SetPersistentDisk gets a reference to the given ListAccountsForStack200ResponseLinksStack and assigns it to the PersistentDisk field.
func (o *ListDiskAttachmentsForPersistentDisk200ResponseLinks) SetPersistentDisk(v ListAccountsForStack200ResponseLinksStack) {
	o.PersistentDisk = &v
}

// GetNext returns the Next field value if set, zero value otherwise.
func (o *ListDiskAttachmentsForPersistentDisk200ResponseLinks) GetNext() ListAccountsForStack200ResponseLinksStack {
	if o == nil || IsNil(o.Next) {
		var ret ListAccountsForStack200ResponseLinksStack
		return ret
	}
	return *o.Next
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListDiskAttachmentsForPersistentDisk200ResponseLinks) GetNextOk() (*ListAccountsForStack200ResponseLinksStack, bool) {
	if o == nil || IsNil(o.Next) {
		return nil, false
	}
	return o.Next, true
}

// HasNext returns a boolean if a field has been set.
func (o *ListDiskAttachmentsForPersistentDisk200ResponseLinks) HasNext() bool {
	if o != nil && !IsNil(o.Next) {
		return true
	}

	return false
}

// SetNext gets a reference to the given ListAccountsForStack200ResponseLinksStack and assigns it to the Next field.
func (o *ListDiskAttachmentsForPersistentDisk200ResponseLinks) SetNext(v ListAccountsForStack200ResponseLinksStack) {
	o.Next = &v
}

// GetPrev returns the Prev field value if set, zero value otherwise.
func (o *ListDiskAttachmentsForPersistentDisk200ResponseLinks) GetPrev() ListAccountsForStack200ResponseLinksStack {
	if o == nil || IsNil(o.Prev) {
		var ret ListAccountsForStack200ResponseLinksStack
		return ret
	}
	return *o.Prev
}

// GetPrevOk returns a tuple with the Prev field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListDiskAttachmentsForPersistentDisk200ResponseLinks) GetPrevOk() (*ListAccountsForStack200ResponseLinksStack, bool) {
	if o == nil || IsNil(o.Prev) {
		return nil, false
	}
	return o.Prev, true
}

// HasPrev returns a boolean if a field has been set.
func (o *ListDiskAttachmentsForPersistentDisk200ResponseLinks) HasPrev() bool {
	if o != nil && !IsNil(o.Prev) {
		return true
	}

	return false
}

// SetPrev gets a reference to the given ListAccountsForStack200ResponseLinksStack and assigns it to the Prev field.
func (o *ListDiskAttachmentsForPersistentDisk200ResponseLinks) SetPrev(v ListAccountsForStack200ResponseLinksStack) {
	o.Prev = &v
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *ListDiskAttachmentsForPersistentDisk200ResponseLinks) GetSelf() ListAccountsForStack200ResponseLinksStack {
	if o == nil || IsNil(o.Self) {
		var ret ListAccountsForStack200ResponseLinksStack
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListDiskAttachmentsForPersistentDisk200ResponseLinks) GetSelfOk() (*ListAccountsForStack200ResponseLinksStack, bool) {
	if o == nil || IsNil(o.Self) {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *ListDiskAttachmentsForPersistentDisk200ResponseLinks) HasSelf() bool {
	if o != nil && !IsNil(o.Self) {
		return true
	}

	return false
}

// SetSelf gets a reference to the given ListAccountsForStack200ResponseLinksStack and assigns it to the Self field.
func (o *ListDiskAttachmentsForPersistentDisk200ResponseLinks) SetSelf(v ListAccountsForStack200ResponseLinksStack) {
	o.Self = &v
}

func (o ListDiskAttachmentsForPersistentDisk200ResponseLinks) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListDiskAttachmentsForPersistentDisk200ResponseLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PersistentDisk) {
		toSerialize["persistent_disk"] = o.PersistentDisk
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

func (o *ListDiskAttachmentsForPersistentDisk200ResponseLinks) UnmarshalJSON(data []byte) (err error) {
	varListDiskAttachmentsForPersistentDisk200ResponseLinks := _ListDiskAttachmentsForPersistentDisk200ResponseLinks{}

	err = json.Unmarshal(data, &varListDiskAttachmentsForPersistentDisk200ResponseLinks)

	if err != nil {
		return err
	}

	*o = ListDiskAttachmentsForPersistentDisk200ResponseLinks(varListDiskAttachmentsForPersistentDisk200ResponseLinks)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "persistent_disk")
		delete(additionalProperties, "next")
		delete(additionalProperties, "prev")
		delete(additionalProperties, "self")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListDiskAttachmentsForPersistentDisk200ResponseLinks struct {
	value *ListDiskAttachmentsForPersistentDisk200ResponseLinks
	isSet bool
}

func (v NullableListDiskAttachmentsForPersistentDisk200ResponseLinks) Get() *ListDiskAttachmentsForPersistentDisk200ResponseLinks {
	return v.value
}

func (v *NullableListDiskAttachmentsForPersistentDisk200ResponseLinks) Set(val *ListDiskAttachmentsForPersistentDisk200ResponseLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableListDiskAttachmentsForPersistentDisk200ResponseLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableListDiskAttachmentsForPersistentDisk200ResponseLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListDiskAttachmentsForPersistentDisk200ResponseLinks(val *ListDiskAttachmentsForPersistentDisk200ResponseLinks) *NullableListDiskAttachmentsForPersistentDisk200ResponseLinks {
	return &NullableListDiskAttachmentsForPersistentDisk200ResponseLinks{value: val, isSet: true}
}

func (v NullableListDiskAttachmentsForPersistentDisk200ResponseLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListDiskAttachmentsForPersistentDisk200ResponseLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


