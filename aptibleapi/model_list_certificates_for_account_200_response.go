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

// checks if the ListCertificatesForAccount200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListCertificatesForAccount200Response{}

// ListCertificatesForAccount200Response struct for ListCertificatesForAccount200Response
type ListCertificatesForAccount200Response struct {
	Embedded ListCertificatesForAccount200ResponseEmbedded `json:"_embedded"`
	TotalCount int32 `json:"total_count"`
	PerPage int32 `json:"per_page"`
	CurrentPage int32 `json:"current_page"`
	Links ListActivityReportsForAccount200ResponseLinks `json:"_links"`
	AdditionalProperties map[string]interface{}
}

type _ListCertificatesForAccount200Response ListCertificatesForAccount200Response

// NewListCertificatesForAccount200Response instantiates a new ListCertificatesForAccount200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListCertificatesForAccount200Response(embedded ListCertificatesForAccount200ResponseEmbedded, totalCount int32, perPage int32, currentPage int32, links ListActivityReportsForAccount200ResponseLinks) *ListCertificatesForAccount200Response {
	this := ListCertificatesForAccount200Response{}
	this.Embedded = embedded
	this.TotalCount = totalCount
	this.PerPage = perPage
	this.CurrentPage = currentPage
	this.Links = links
	return &this
}

// NewListCertificatesForAccount200ResponseWithDefaults instantiates a new ListCertificatesForAccount200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListCertificatesForAccount200ResponseWithDefaults() *ListCertificatesForAccount200Response {
	this := ListCertificatesForAccount200Response{}
	return &this
}

// GetEmbedded returns the Embedded field value
func (o *ListCertificatesForAccount200Response) GetEmbedded() ListCertificatesForAccount200ResponseEmbedded {
	if o == nil {
		var ret ListCertificatesForAccount200ResponseEmbedded
		return ret
	}

	return o.Embedded
}

// GetEmbeddedOk returns a tuple with the Embedded field value
// and a boolean to check if the value has been set.
func (o *ListCertificatesForAccount200Response) GetEmbeddedOk() (*ListCertificatesForAccount200ResponseEmbedded, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Embedded, true
}

// SetEmbedded sets field value
func (o *ListCertificatesForAccount200Response) SetEmbedded(v ListCertificatesForAccount200ResponseEmbedded) {
	o.Embedded = v
}

// GetTotalCount returns the TotalCount field value
func (o *ListCertificatesForAccount200Response) GetTotalCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value
// and a boolean to check if the value has been set.
func (o *ListCertificatesForAccount200Response) GetTotalCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalCount, true
}

// SetTotalCount sets field value
func (o *ListCertificatesForAccount200Response) SetTotalCount(v int32) {
	o.TotalCount = v
}

// GetPerPage returns the PerPage field value
func (o *ListCertificatesForAccount200Response) GetPerPage() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PerPage
}

// GetPerPageOk returns a tuple with the PerPage field value
// and a boolean to check if the value has been set.
func (o *ListCertificatesForAccount200Response) GetPerPageOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PerPage, true
}

// SetPerPage sets field value
func (o *ListCertificatesForAccount200Response) SetPerPage(v int32) {
	o.PerPage = v
}

// GetCurrentPage returns the CurrentPage field value
func (o *ListCertificatesForAccount200Response) GetCurrentPage() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CurrentPage
}

// GetCurrentPageOk returns a tuple with the CurrentPage field value
// and a boolean to check if the value has been set.
func (o *ListCertificatesForAccount200Response) GetCurrentPageOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrentPage, true
}

// SetCurrentPage sets field value
func (o *ListCertificatesForAccount200Response) SetCurrentPage(v int32) {
	o.CurrentPage = v
}

// GetLinks returns the Links field value
func (o *ListCertificatesForAccount200Response) GetLinks() ListActivityReportsForAccount200ResponseLinks {
	if o == nil {
		var ret ListActivityReportsForAccount200ResponseLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *ListCertificatesForAccount200Response) GetLinksOk() (*ListActivityReportsForAccount200ResponseLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *ListCertificatesForAccount200Response) SetLinks(v ListActivityReportsForAccount200ResponseLinks) {
	o.Links = v
}

func (o ListCertificatesForAccount200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListCertificatesForAccount200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["_embedded"] = o.Embedded
	toSerialize["total_count"] = o.TotalCount
	toSerialize["per_page"] = o.PerPage
	toSerialize["current_page"] = o.CurrentPage
	toSerialize["_links"] = o.Links

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListCertificatesForAccount200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"_embedded",
		"total_count",
		"per_page",
		"current_page",
		"_links",
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

	varListCertificatesForAccount200Response := _ListCertificatesForAccount200Response{}

	err = json.Unmarshal(data, &varListCertificatesForAccount200Response)

	if err != nil {
		return err
	}

	*o = ListCertificatesForAccount200Response(varListCertificatesForAccount200Response)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "_embedded")
		delete(additionalProperties, "total_count")
		delete(additionalProperties, "per_page")
		delete(additionalProperties, "current_page")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListCertificatesForAccount200Response struct {
	value *ListCertificatesForAccount200Response
	isSet bool
}

func (v NullableListCertificatesForAccount200Response) Get() *ListCertificatesForAccount200Response {
	return v.value
}

func (v *NullableListCertificatesForAccount200Response) Set(val *ListCertificatesForAccount200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableListCertificatesForAccount200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableListCertificatesForAccount200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListCertificatesForAccount200Response(val *ListCertificatesForAccount200Response) *NullableListCertificatesForAccount200Response {
	return &NullableListCertificatesForAccount200Response{value: val, isSet: true}
}

func (v NullableListCertificatesForAccount200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListCertificatesForAccount200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


