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

// checks if the CodeScanResultLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CodeScanResultLinks{}

// CodeScanResultLinks struct for CodeScanResultLinks
type CodeScanResultLinks struct {
	App *ListAccountsForStack200ResponseLinksStack `json:"app,omitempty"`
	Operation *ListAccountsForStack200ResponseLinksStack `json:"operation,omitempty"`
	Self *ListAccountsForStack200ResponseLinksStack `json:"self,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CodeScanResultLinks CodeScanResultLinks

// NewCodeScanResultLinks instantiates a new CodeScanResultLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCodeScanResultLinks() *CodeScanResultLinks {
	this := CodeScanResultLinks{}
	return &this
}

// NewCodeScanResultLinksWithDefaults instantiates a new CodeScanResultLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCodeScanResultLinksWithDefaults() *CodeScanResultLinks {
	this := CodeScanResultLinks{}
	return &this
}

// GetApp returns the App field value if set, zero value otherwise.
func (o *CodeScanResultLinks) GetApp() ListAccountsForStack200ResponseLinksStack {
	if o == nil || IsNil(o.App) {
		var ret ListAccountsForStack200ResponseLinksStack
		return ret
	}
	return *o.App
}

// GetAppOk returns a tuple with the App field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CodeScanResultLinks) GetAppOk() (*ListAccountsForStack200ResponseLinksStack, bool) {
	if o == nil || IsNil(o.App) {
		return nil, false
	}
	return o.App, true
}

// HasApp returns a boolean if a field has been set.
func (o *CodeScanResultLinks) HasApp() bool {
	if o != nil && !IsNil(o.App) {
		return true
	}

	return false
}

// SetApp gets a reference to the given ListAccountsForStack200ResponseLinksStack and assigns it to the App field.
func (o *CodeScanResultLinks) SetApp(v ListAccountsForStack200ResponseLinksStack) {
	o.App = &v
}

// GetOperation returns the Operation field value if set, zero value otherwise.
func (o *CodeScanResultLinks) GetOperation() ListAccountsForStack200ResponseLinksStack {
	if o == nil || IsNil(o.Operation) {
		var ret ListAccountsForStack200ResponseLinksStack
		return ret
	}
	return *o.Operation
}

// GetOperationOk returns a tuple with the Operation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CodeScanResultLinks) GetOperationOk() (*ListAccountsForStack200ResponseLinksStack, bool) {
	if o == nil || IsNil(o.Operation) {
		return nil, false
	}
	return o.Operation, true
}

// HasOperation returns a boolean if a field has been set.
func (o *CodeScanResultLinks) HasOperation() bool {
	if o != nil && !IsNil(o.Operation) {
		return true
	}

	return false
}

// SetOperation gets a reference to the given ListAccountsForStack200ResponseLinksStack and assigns it to the Operation field.
func (o *CodeScanResultLinks) SetOperation(v ListAccountsForStack200ResponseLinksStack) {
	o.Operation = &v
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *CodeScanResultLinks) GetSelf() ListAccountsForStack200ResponseLinksStack {
	if o == nil || IsNil(o.Self) {
		var ret ListAccountsForStack200ResponseLinksStack
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CodeScanResultLinks) GetSelfOk() (*ListAccountsForStack200ResponseLinksStack, bool) {
	if o == nil || IsNil(o.Self) {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *CodeScanResultLinks) HasSelf() bool {
	if o != nil && !IsNil(o.Self) {
		return true
	}

	return false
}

// SetSelf gets a reference to the given ListAccountsForStack200ResponseLinksStack and assigns it to the Self field.
func (o *CodeScanResultLinks) SetSelf(v ListAccountsForStack200ResponseLinksStack) {
	o.Self = &v
}

func (o CodeScanResultLinks) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CodeScanResultLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.App) {
		toSerialize["app"] = o.App
	}
	if !IsNil(o.Operation) {
		toSerialize["operation"] = o.Operation
	}
	if !IsNil(o.Self) {
		toSerialize["self"] = o.Self
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CodeScanResultLinks) UnmarshalJSON(data []byte) (err error) {
	varCodeScanResultLinks := _CodeScanResultLinks{}

	err = json.Unmarshal(data, &varCodeScanResultLinks)

	if err != nil {
		return err
	}

	*o = CodeScanResultLinks(varCodeScanResultLinks)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "app")
		delete(additionalProperties, "operation")
		delete(additionalProperties, "self")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCodeScanResultLinks struct {
	value *CodeScanResultLinks
	isSet bool
}

func (v NullableCodeScanResultLinks) Get() *CodeScanResultLinks {
	return v.value
}

func (v *NullableCodeScanResultLinks) Set(val *CodeScanResultLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableCodeScanResultLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableCodeScanResultLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCodeScanResultLinks(val *CodeScanResultLinks) *NullableCodeScanResultLinks {
	return &NullableCodeScanResultLinks{value: val, isSet: true}
}

func (v NullableCodeScanResultLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCodeScanResultLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

