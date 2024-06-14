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

// checks if the ImageLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImageLinks{}

// ImageLinks struct for ImageLinks
type ImageLinks struct {
	App *ListAccountsForStack200ResponseLinksStack `json:"app,omitempty"`
	Operations *ListAccountsForStack200ResponseLinksStack `json:"operations,omitempty"`
	Self *ListAccountsForStack200ResponseLinksStack `json:"self,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ImageLinks ImageLinks

// NewImageLinks instantiates a new ImageLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImageLinks() *ImageLinks {
	this := ImageLinks{}
	return &this
}

// NewImageLinksWithDefaults instantiates a new ImageLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImageLinksWithDefaults() *ImageLinks {
	this := ImageLinks{}
	return &this
}

// GetApp returns the App field value if set, zero value otherwise.
func (o *ImageLinks) GetApp() ListAccountsForStack200ResponseLinksStack {
	if o == nil || IsNil(o.App) {
		var ret ListAccountsForStack200ResponseLinksStack
		return ret
	}
	return *o.App
}

// GetAppOk returns a tuple with the App field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageLinks) GetAppOk() (*ListAccountsForStack200ResponseLinksStack, bool) {
	if o == nil || IsNil(o.App) {
		return nil, false
	}
	return o.App, true
}

// HasApp returns a boolean if a field has been set.
func (o *ImageLinks) HasApp() bool {
	if o != nil && !IsNil(o.App) {
		return true
	}

	return false
}

// SetApp gets a reference to the given ListAccountsForStack200ResponseLinksStack and assigns it to the App field.
func (o *ImageLinks) SetApp(v ListAccountsForStack200ResponseLinksStack) {
	o.App = &v
}

// GetOperations returns the Operations field value if set, zero value otherwise.
func (o *ImageLinks) GetOperations() ListAccountsForStack200ResponseLinksStack {
	if o == nil || IsNil(o.Operations) {
		var ret ListAccountsForStack200ResponseLinksStack
		return ret
	}
	return *o.Operations
}

// GetOperationsOk returns a tuple with the Operations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageLinks) GetOperationsOk() (*ListAccountsForStack200ResponseLinksStack, bool) {
	if o == nil || IsNil(o.Operations) {
		return nil, false
	}
	return o.Operations, true
}

// HasOperations returns a boolean if a field has been set.
func (o *ImageLinks) HasOperations() bool {
	if o != nil && !IsNil(o.Operations) {
		return true
	}

	return false
}

// SetOperations gets a reference to the given ListAccountsForStack200ResponseLinksStack and assigns it to the Operations field.
func (o *ImageLinks) SetOperations(v ListAccountsForStack200ResponseLinksStack) {
	o.Operations = &v
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *ImageLinks) GetSelf() ListAccountsForStack200ResponseLinksStack {
	if o == nil || IsNil(o.Self) {
		var ret ListAccountsForStack200ResponseLinksStack
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageLinks) GetSelfOk() (*ListAccountsForStack200ResponseLinksStack, bool) {
	if o == nil || IsNil(o.Self) {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *ImageLinks) HasSelf() bool {
	if o != nil && !IsNil(o.Self) {
		return true
	}

	return false
}

// SetSelf gets a reference to the given ListAccountsForStack200ResponseLinksStack and assigns it to the Self field.
func (o *ImageLinks) SetSelf(v ListAccountsForStack200ResponseLinksStack) {
	o.Self = &v
}

func (o ImageLinks) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ImageLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.App) {
		toSerialize["app"] = o.App
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

func (o *ImageLinks) UnmarshalJSON(data []byte) (err error) {
	varImageLinks := _ImageLinks{}

	err = json.Unmarshal(data, &varImageLinks)

	if err != nil {
		return err
	}

	*o = ImageLinks(varImageLinks)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "app")
		delete(additionalProperties, "operations")
		delete(additionalProperties, "self")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableImageLinks struct {
	value *ImageLinks
	isSet bool
}

func (v NullableImageLinks) Get() *ImageLinks {
	return v.value
}

func (v *NullableImageLinks) Set(val *ImageLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableImageLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableImageLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImageLinks(val *ImageLinks) *NullableImageLinks {
	return &NullableImageLinks{value: val, isSet: true}
}

func (v NullableImageLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImageLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


