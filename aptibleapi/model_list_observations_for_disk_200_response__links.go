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

// checks if the ListObservationsForDisk200ResponseLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListObservationsForDisk200ResponseLinks{}

// ListObservationsForDisk200ResponseLinks struct for ListObservationsForDisk200ResponseLinks
type ListObservationsForDisk200ResponseLinks struct {
	Disk *ListAccountsForStack200ResponseLinksStack `json:"disk,omitempty"`
	Next *ListAccountsForStack200ResponseLinksStack `json:"next,omitempty"`
	Prev *ListAccountsForStack200ResponseLinksStack `json:"prev,omitempty"`
	Self *ListAccountsForStack200ResponseLinksStack `json:"self,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ListObservationsForDisk200ResponseLinks ListObservationsForDisk200ResponseLinks

// NewListObservationsForDisk200ResponseLinks instantiates a new ListObservationsForDisk200ResponseLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListObservationsForDisk200ResponseLinks() *ListObservationsForDisk200ResponseLinks {
	this := ListObservationsForDisk200ResponseLinks{}
	return &this
}

// NewListObservationsForDisk200ResponseLinksWithDefaults instantiates a new ListObservationsForDisk200ResponseLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListObservationsForDisk200ResponseLinksWithDefaults() *ListObservationsForDisk200ResponseLinks {
	this := ListObservationsForDisk200ResponseLinks{}
	return &this
}

// GetDisk returns the Disk field value if set, zero value otherwise.
func (o *ListObservationsForDisk200ResponseLinks) GetDisk() ListAccountsForStack200ResponseLinksStack {
	if o == nil || IsNil(o.Disk) {
		var ret ListAccountsForStack200ResponseLinksStack
		return ret
	}
	return *o.Disk
}

// GetDiskOk returns a tuple with the Disk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListObservationsForDisk200ResponseLinks) GetDiskOk() (*ListAccountsForStack200ResponseLinksStack, bool) {
	if o == nil || IsNil(o.Disk) {
		return nil, false
	}
	return o.Disk, true
}

// HasDisk returns a boolean if a field has been set.
func (o *ListObservationsForDisk200ResponseLinks) HasDisk() bool {
	if o != nil && !IsNil(o.Disk) {
		return true
	}

	return false
}

// SetDisk gets a reference to the given ListAccountsForStack200ResponseLinksStack and assigns it to the Disk field.
func (o *ListObservationsForDisk200ResponseLinks) SetDisk(v ListAccountsForStack200ResponseLinksStack) {
	o.Disk = &v
}

// GetNext returns the Next field value if set, zero value otherwise.
func (o *ListObservationsForDisk200ResponseLinks) GetNext() ListAccountsForStack200ResponseLinksStack {
	if o == nil || IsNil(o.Next) {
		var ret ListAccountsForStack200ResponseLinksStack
		return ret
	}
	return *o.Next
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListObservationsForDisk200ResponseLinks) GetNextOk() (*ListAccountsForStack200ResponseLinksStack, bool) {
	if o == nil || IsNil(o.Next) {
		return nil, false
	}
	return o.Next, true
}

// HasNext returns a boolean if a field has been set.
func (o *ListObservationsForDisk200ResponseLinks) HasNext() bool {
	if o != nil && !IsNil(o.Next) {
		return true
	}

	return false
}

// SetNext gets a reference to the given ListAccountsForStack200ResponseLinksStack and assigns it to the Next field.
func (o *ListObservationsForDisk200ResponseLinks) SetNext(v ListAccountsForStack200ResponseLinksStack) {
	o.Next = &v
}

// GetPrev returns the Prev field value if set, zero value otherwise.
func (o *ListObservationsForDisk200ResponseLinks) GetPrev() ListAccountsForStack200ResponseLinksStack {
	if o == nil || IsNil(o.Prev) {
		var ret ListAccountsForStack200ResponseLinksStack
		return ret
	}
	return *o.Prev
}

// GetPrevOk returns a tuple with the Prev field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListObservationsForDisk200ResponseLinks) GetPrevOk() (*ListAccountsForStack200ResponseLinksStack, bool) {
	if o == nil || IsNil(o.Prev) {
		return nil, false
	}
	return o.Prev, true
}

// HasPrev returns a boolean if a field has been set.
func (o *ListObservationsForDisk200ResponseLinks) HasPrev() bool {
	if o != nil && !IsNil(o.Prev) {
		return true
	}

	return false
}

// SetPrev gets a reference to the given ListAccountsForStack200ResponseLinksStack and assigns it to the Prev field.
func (o *ListObservationsForDisk200ResponseLinks) SetPrev(v ListAccountsForStack200ResponseLinksStack) {
	o.Prev = &v
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *ListObservationsForDisk200ResponseLinks) GetSelf() ListAccountsForStack200ResponseLinksStack {
	if o == nil || IsNil(o.Self) {
		var ret ListAccountsForStack200ResponseLinksStack
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListObservationsForDisk200ResponseLinks) GetSelfOk() (*ListAccountsForStack200ResponseLinksStack, bool) {
	if o == nil || IsNil(o.Self) {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *ListObservationsForDisk200ResponseLinks) HasSelf() bool {
	if o != nil && !IsNil(o.Self) {
		return true
	}

	return false
}

// SetSelf gets a reference to the given ListAccountsForStack200ResponseLinksStack and assigns it to the Self field.
func (o *ListObservationsForDisk200ResponseLinks) SetSelf(v ListAccountsForStack200ResponseLinksStack) {
	o.Self = &v
}

func (o ListObservationsForDisk200ResponseLinks) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListObservationsForDisk200ResponseLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Disk) {
		toSerialize["disk"] = o.Disk
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

func (o *ListObservationsForDisk200ResponseLinks) UnmarshalJSON(data []byte) (err error) {
	varListObservationsForDisk200ResponseLinks := _ListObservationsForDisk200ResponseLinks{}

	err = json.Unmarshal(data, &varListObservationsForDisk200ResponseLinks)

	if err != nil {
		return err
	}

	*o = ListObservationsForDisk200ResponseLinks(varListObservationsForDisk200ResponseLinks)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "disk")
		delete(additionalProperties, "next")
		delete(additionalProperties, "prev")
		delete(additionalProperties, "self")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListObservationsForDisk200ResponseLinks struct {
	value *ListObservationsForDisk200ResponseLinks
	isSet bool
}

func (v NullableListObservationsForDisk200ResponseLinks) Get() *ListObservationsForDisk200ResponseLinks {
	return v.value
}

func (v *NullableListObservationsForDisk200ResponseLinks) Set(val *ListObservationsForDisk200ResponseLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableListObservationsForDisk200ResponseLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableListObservationsForDisk200ResponseLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListObservationsForDisk200ResponseLinks(val *ListObservationsForDisk200ResponseLinks) *NullableListObservationsForDisk200ResponseLinks {
	return &NullableListObservationsForDisk200ResponseLinks{value: val, isSet: true}
}

func (v NullableListObservationsForDisk200ResponseLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListObservationsForDisk200ResponseLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


