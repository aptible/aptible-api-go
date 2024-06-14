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

// checks if the ListDiskAttachmentsForAccount200ResponseEmbedded type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListDiskAttachmentsForAccount200ResponseEmbedded{}

// ListDiskAttachmentsForAccount200ResponseEmbedded struct for ListDiskAttachmentsForAccount200ResponseEmbedded
type ListDiskAttachmentsForAccount200ResponseEmbedded struct {
	DiskAttachments []DiskAttachment `json:"disk_attachments,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ListDiskAttachmentsForAccount200ResponseEmbedded ListDiskAttachmentsForAccount200ResponseEmbedded

// NewListDiskAttachmentsForAccount200ResponseEmbedded instantiates a new ListDiskAttachmentsForAccount200ResponseEmbedded object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListDiskAttachmentsForAccount200ResponseEmbedded() *ListDiskAttachmentsForAccount200ResponseEmbedded {
	this := ListDiskAttachmentsForAccount200ResponseEmbedded{}
	return &this
}

// NewListDiskAttachmentsForAccount200ResponseEmbeddedWithDefaults instantiates a new ListDiskAttachmentsForAccount200ResponseEmbedded object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListDiskAttachmentsForAccount200ResponseEmbeddedWithDefaults() *ListDiskAttachmentsForAccount200ResponseEmbedded {
	this := ListDiskAttachmentsForAccount200ResponseEmbedded{}
	return &this
}

// GetDiskAttachments returns the DiskAttachments field value if set, zero value otherwise.
func (o *ListDiskAttachmentsForAccount200ResponseEmbedded) GetDiskAttachments() []DiskAttachment {
	if o == nil || IsNil(o.DiskAttachments) {
		var ret []DiskAttachment
		return ret
	}
	return o.DiskAttachments
}

// GetDiskAttachmentsOk returns a tuple with the DiskAttachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListDiskAttachmentsForAccount200ResponseEmbedded) GetDiskAttachmentsOk() ([]DiskAttachment, bool) {
	if o == nil || IsNil(o.DiskAttachments) {
		return nil, false
	}
	return o.DiskAttachments, true
}

// HasDiskAttachments returns a boolean if a field has been set.
func (o *ListDiskAttachmentsForAccount200ResponseEmbedded) HasDiskAttachments() bool {
	if o != nil && !IsNil(o.DiskAttachments) {
		return true
	}

	return false
}

// SetDiskAttachments gets a reference to the given []DiskAttachment and assigns it to the DiskAttachments field.
func (o *ListDiskAttachmentsForAccount200ResponseEmbedded) SetDiskAttachments(v []DiskAttachment) {
	o.DiskAttachments = v
}

func (o ListDiskAttachmentsForAccount200ResponseEmbedded) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListDiskAttachmentsForAccount200ResponseEmbedded) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DiskAttachments) {
		toSerialize["disk_attachments"] = o.DiskAttachments
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListDiskAttachmentsForAccount200ResponseEmbedded) UnmarshalJSON(data []byte) (err error) {
	varListDiskAttachmentsForAccount200ResponseEmbedded := _ListDiskAttachmentsForAccount200ResponseEmbedded{}

	err = json.Unmarshal(data, &varListDiskAttachmentsForAccount200ResponseEmbedded)

	if err != nil {
		return err
	}

	*o = ListDiskAttachmentsForAccount200ResponseEmbedded(varListDiskAttachmentsForAccount200ResponseEmbedded)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "disk_attachments")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListDiskAttachmentsForAccount200ResponseEmbedded struct {
	value *ListDiskAttachmentsForAccount200ResponseEmbedded
	isSet bool
}

func (v NullableListDiskAttachmentsForAccount200ResponseEmbedded) Get() *ListDiskAttachmentsForAccount200ResponseEmbedded {
	return v.value
}

func (v *NullableListDiskAttachmentsForAccount200ResponseEmbedded) Set(val *ListDiskAttachmentsForAccount200ResponseEmbedded) {
	v.value = val
	v.isSet = true
}

func (v NullableListDiskAttachmentsForAccount200ResponseEmbedded) IsSet() bool {
	return v.isSet
}

func (v *NullableListDiskAttachmentsForAccount200ResponseEmbedded) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListDiskAttachmentsForAccount200ResponseEmbedded(val *ListDiskAttachmentsForAccount200ResponseEmbedded) *NullableListDiskAttachmentsForAccount200ResponseEmbedded {
	return &NullableListDiskAttachmentsForAccount200ResponseEmbedded{value: val, isSet: true}
}

func (v NullableListDiskAttachmentsForAccount200ResponseEmbedded) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListDiskAttachmentsForAccount200ResponseEmbedded) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


