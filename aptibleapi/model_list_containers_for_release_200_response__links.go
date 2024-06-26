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

// checks if the ListContainersForRelease200ResponseLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListContainersForRelease200ResponseLinks{}

// ListContainersForRelease200ResponseLinks struct for ListContainersForRelease200ResponseLinks
type ListContainersForRelease200ResponseLinks struct {
	Release *ListAccountsForStack200ResponseLinksStack `json:"release,omitempty"`
	LogDrain *ListAccountsForStack200ResponseLinksStack `json:"log_drain,omitempty"`
	Next *ListAccountsForStack200ResponseLinksStack `json:"next,omitempty"`
	Prev *ListAccountsForStack200ResponseLinksStack `json:"prev,omitempty"`
	Self *ListAccountsForStack200ResponseLinksStack `json:"self,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ListContainersForRelease200ResponseLinks ListContainersForRelease200ResponseLinks

// NewListContainersForRelease200ResponseLinks instantiates a new ListContainersForRelease200ResponseLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListContainersForRelease200ResponseLinks() *ListContainersForRelease200ResponseLinks {
	this := ListContainersForRelease200ResponseLinks{}
	return &this
}

// NewListContainersForRelease200ResponseLinksWithDefaults instantiates a new ListContainersForRelease200ResponseLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListContainersForRelease200ResponseLinksWithDefaults() *ListContainersForRelease200ResponseLinks {
	this := ListContainersForRelease200ResponseLinks{}
	return &this
}

// GetRelease returns the Release field value if set, zero value otherwise.
func (o *ListContainersForRelease200ResponseLinks) GetRelease() ListAccountsForStack200ResponseLinksStack {
	if o == nil || IsNil(o.Release) {
		var ret ListAccountsForStack200ResponseLinksStack
		return ret
	}
	return *o.Release
}

// GetReleaseOk returns a tuple with the Release field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListContainersForRelease200ResponseLinks) GetReleaseOk() (*ListAccountsForStack200ResponseLinksStack, bool) {
	if o == nil || IsNil(o.Release) {
		return nil, false
	}
	return o.Release, true
}

// HasRelease returns a boolean if a field has been set.
func (o *ListContainersForRelease200ResponseLinks) HasRelease() bool {
	if o != nil && !IsNil(o.Release) {
		return true
	}

	return false
}

// SetRelease gets a reference to the given ListAccountsForStack200ResponseLinksStack and assigns it to the Release field.
func (o *ListContainersForRelease200ResponseLinks) SetRelease(v ListAccountsForStack200ResponseLinksStack) {
	o.Release = &v
}

// GetLogDrain returns the LogDrain field value if set, zero value otherwise.
func (o *ListContainersForRelease200ResponseLinks) GetLogDrain() ListAccountsForStack200ResponseLinksStack {
	if o == nil || IsNil(o.LogDrain) {
		var ret ListAccountsForStack200ResponseLinksStack
		return ret
	}
	return *o.LogDrain
}

// GetLogDrainOk returns a tuple with the LogDrain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListContainersForRelease200ResponseLinks) GetLogDrainOk() (*ListAccountsForStack200ResponseLinksStack, bool) {
	if o == nil || IsNil(o.LogDrain) {
		return nil, false
	}
	return o.LogDrain, true
}

// HasLogDrain returns a boolean if a field has been set.
func (o *ListContainersForRelease200ResponseLinks) HasLogDrain() bool {
	if o != nil && !IsNil(o.LogDrain) {
		return true
	}

	return false
}

// SetLogDrain gets a reference to the given ListAccountsForStack200ResponseLinksStack and assigns it to the LogDrain field.
func (o *ListContainersForRelease200ResponseLinks) SetLogDrain(v ListAccountsForStack200ResponseLinksStack) {
	o.LogDrain = &v
}

// GetNext returns the Next field value if set, zero value otherwise.
func (o *ListContainersForRelease200ResponseLinks) GetNext() ListAccountsForStack200ResponseLinksStack {
	if o == nil || IsNil(o.Next) {
		var ret ListAccountsForStack200ResponseLinksStack
		return ret
	}
	return *o.Next
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListContainersForRelease200ResponseLinks) GetNextOk() (*ListAccountsForStack200ResponseLinksStack, bool) {
	if o == nil || IsNil(o.Next) {
		return nil, false
	}
	return o.Next, true
}

// HasNext returns a boolean if a field has been set.
func (o *ListContainersForRelease200ResponseLinks) HasNext() bool {
	if o != nil && !IsNil(o.Next) {
		return true
	}

	return false
}

// SetNext gets a reference to the given ListAccountsForStack200ResponseLinksStack and assigns it to the Next field.
func (o *ListContainersForRelease200ResponseLinks) SetNext(v ListAccountsForStack200ResponseLinksStack) {
	o.Next = &v
}

// GetPrev returns the Prev field value if set, zero value otherwise.
func (o *ListContainersForRelease200ResponseLinks) GetPrev() ListAccountsForStack200ResponseLinksStack {
	if o == nil || IsNil(o.Prev) {
		var ret ListAccountsForStack200ResponseLinksStack
		return ret
	}
	return *o.Prev
}

// GetPrevOk returns a tuple with the Prev field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListContainersForRelease200ResponseLinks) GetPrevOk() (*ListAccountsForStack200ResponseLinksStack, bool) {
	if o == nil || IsNil(o.Prev) {
		return nil, false
	}
	return o.Prev, true
}

// HasPrev returns a boolean if a field has been set.
func (o *ListContainersForRelease200ResponseLinks) HasPrev() bool {
	if o != nil && !IsNil(o.Prev) {
		return true
	}

	return false
}

// SetPrev gets a reference to the given ListAccountsForStack200ResponseLinksStack and assigns it to the Prev field.
func (o *ListContainersForRelease200ResponseLinks) SetPrev(v ListAccountsForStack200ResponseLinksStack) {
	o.Prev = &v
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *ListContainersForRelease200ResponseLinks) GetSelf() ListAccountsForStack200ResponseLinksStack {
	if o == nil || IsNil(o.Self) {
		var ret ListAccountsForStack200ResponseLinksStack
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListContainersForRelease200ResponseLinks) GetSelfOk() (*ListAccountsForStack200ResponseLinksStack, bool) {
	if o == nil || IsNil(o.Self) {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *ListContainersForRelease200ResponseLinks) HasSelf() bool {
	if o != nil && !IsNil(o.Self) {
		return true
	}

	return false
}

// SetSelf gets a reference to the given ListAccountsForStack200ResponseLinksStack and assigns it to the Self field.
func (o *ListContainersForRelease200ResponseLinks) SetSelf(v ListAccountsForStack200ResponseLinksStack) {
	o.Self = &v
}

func (o ListContainersForRelease200ResponseLinks) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListContainersForRelease200ResponseLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Release) {
		toSerialize["release"] = o.Release
	}
	if !IsNil(o.LogDrain) {
		toSerialize["log_drain"] = o.LogDrain
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

func (o *ListContainersForRelease200ResponseLinks) UnmarshalJSON(data []byte) (err error) {
	varListContainersForRelease200ResponseLinks := _ListContainersForRelease200ResponseLinks{}

	err = json.Unmarshal(data, &varListContainersForRelease200ResponseLinks)

	if err != nil {
		return err
	}

	*o = ListContainersForRelease200ResponseLinks(varListContainersForRelease200ResponseLinks)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "release")
		delete(additionalProperties, "log_drain")
		delete(additionalProperties, "next")
		delete(additionalProperties, "prev")
		delete(additionalProperties, "self")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListContainersForRelease200ResponseLinks struct {
	value *ListContainersForRelease200ResponseLinks
	isSet bool
}

func (v NullableListContainersForRelease200ResponseLinks) Get() *ListContainersForRelease200ResponseLinks {
	return v.value
}

func (v *NullableListContainersForRelease200ResponseLinks) Set(val *ListContainersForRelease200ResponseLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableListContainersForRelease200ResponseLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableListContainersForRelease200ResponseLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListContainersForRelease200ResponseLinks(val *ListContainersForRelease200ResponseLinks) *NullableListContainersForRelease200ResponseLinks {
	return &NullableListContainersForRelease200ResponseLinks{value: val, isSet: true}
}

func (v NullableListContainersForRelease200ResponseLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListContainersForRelease200ResponseLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


