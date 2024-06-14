/*
Deploy API v1

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package aptibleapi

import (
	"encoding/json"
	"fmt"
)

// checks if the Source type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Source{}

// Source struct for Source
type Source struct {
	Id int32 `json:"id"`
	MetaType string `json:"_type"`
	QualifiedName string `json:"qualified_name"`
	DisplayName string `json:"display_name"`
	Url string `json:"url"`
	Links *SourceLinks `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Source Source

// NewSource instantiates a new Source object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSource(id int32, metaType string, qualifiedName string, displayName string, url string) *Source {
	this := Source{}
	this.Id = id
	this.MetaType = metaType
	this.QualifiedName = qualifiedName
	this.DisplayName = displayName
	this.Url = url
	return &this
}

// NewSourceWithDefaults instantiates a new Source object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSourceWithDefaults() *Source {
	this := Source{}
	return &this
}

// GetId returns the Id field value
func (o *Source) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Source) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Source) SetId(v int32) {
	o.Id = v
}

// GetMetaType returns the MetaType field value
func (o *Source) GetMetaType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MetaType
}

// GetMetaTypeOk returns a tuple with the MetaType field value
// and a boolean to check if the value has been set.
func (o *Source) GetMetaTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetaType, true
}

// SetMetaType sets field value
func (o *Source) SetMetaType(v string) {
	o.MetaType = v
}

// GetQualifiedName returns the QualifiedName field value
func (o *Source) GetQualifiedName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.QualifiedName
}

// GetQualifiedNameOk returns a tuple with the QualifiedName field value
// and a boolean to check if the value has been set.
func (o *Source) GetQualifiedNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QualifiedName, true
}

// SetQualifiedName sets field value
func (o *Source) SetQualifiedName(v string) {
	o.QualifiedName = v
}

// GetDisplayName returns the DisplayName field value
func (o *Source) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *Source) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *Source) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetUrl returns the Url field value
func (o *Source) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *Source) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *Source) SetUrl(v string) {
	o.Url = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *Source) GetLinks() SourceLinks {
	if o == nil || IsNil(o.Links) {
		var ret SourceLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Source) GetLinksOk() (*SourceLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *Source) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given SourceLinks and assigns it to the Links field.
func (o *Source) SetLinks(v SourceLinks) {
	o.Links = &v
}

func (o Source) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Source) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["_type"] = o.MetaType
	toSerialize["qualified_name"] = o.QualifiedName
	toSerialize["display_name"] = o.DisplayName
	toSerialize["url"] = o.Url
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Source) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"_type",
		"qualified_name",
		"display_name",
		"url",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varSource := _Source{}

	err = json.Unmarshal(data, &varSource)

	if err != nil {
		return err
	}

	*o = Source(varSource)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "_type")
		delete(additionalProperties, "qualified_name")
		delete(additionalProperties, "display_name")
		delete(additionalProperties, "url")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSource struct {
	value *Source
	isSet bool
}

func (v NullableSource) Get() *Source {
	return v.value
}

func (v *NullableSource) Set(val *Source) {
	v.value = val
	v.isSet = true
}

func (v NullableSource) IsSet() bool {
	return v.isSet
}

func (v *NullableSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSource(val *Source) *NullableSource {
	return &NullableSource{value: val, isSet: true}
}

func (v NullableSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


