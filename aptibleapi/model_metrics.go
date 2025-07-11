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

// checks if the Metrics type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Metrics{}

// Metrics struct for Metrics
type Metrics struct {
	Id int32 `json:"id"`
	MetaType string `json:"_type"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Name string `json:"name"`
	Description string `json:"description"`
	Unit string `json:"unit"`
	Query map[string]interface{} `json:"query"`
	Links *IntegrationLinks `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Metrics Metrics

// NewMetrics instantiates a new Metrics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetrics(id int32, metaType string, createdAt string, updatedAt string, name string, description string, unit string, query map[string]interface{}) *Metrics {
	this := Metrics{}
	this.Id = id
	this.MetaType = metaType
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.Name = name
	this.Description = description
	this.Unit = unit
	this.Query = query
	return &this
}

// NewMetricsWithDefaults instantiates a new Metrics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricsWithDefaults() *Metrics {
	this := Metrics{}
	return &this
}

// GetId returns the Id field value
func (o *Metrics) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Metrics) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Metrics) SetId(v int32) {
	o.Id = v
}

// GetMetaType returns the MetaType field value
func (o *Metrics) GetMetaType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MetaType
}

// GetMetaTypeOk returns a tuple with the MetaType field value
// and a boolean to check if the value has been set.
func (o *Metrics) GetMetaTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetaType, true
}

// SetMetaType sets field value
func (o *Metrics) SetMetaType(v string) {
	o.MetaType = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *Metrics) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Metrics) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *Metrics) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *Metrics) GetUpdatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *Metrics) GetUpdatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *Metrics) SetUpdatedAt(v string) {
	o.UpdatedAt = v
}

// GetName returns the Name field value
func (o *Metrics) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Metrics) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Metrics) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *Metrics) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *Metrics) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *Metrics) SetDescription(v string) {
	o.Description = v
}

// GetUnit returns the Unit field value
func (o *Metrics) GetUnit() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Unit
}

// GetUnitOk returns a tuple with the Unit field value
// and a boolean to check if the value has been set.
func (o *Metrics) GetUnitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Unit, true
}

// SetUnit sets field value
func (o *Metrics) SetUnit(v string) {
	o.Unit = v
}

// GetQuery returns the Query field value
func (o *Metrics) GetQuery() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *Metrics) GetQueryOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Query, true
}

// SetQuery sets field value
func (o *Metrics) SetQuery(v map[string]interface{}) {
	o.Query = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *Metrics) GetLinks() IntegrationLinks {
	if o == nil || IsNil(o.Links) {
		var ret IntegrationLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Metrics) GetLinksOk() (*IntegrationLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *Metrics) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given IntegrationLinks and assigns it to the Links field.
func (o *Metrics) SetLinks(v IntegrationLinks) {
	o.Links = &v
}

func (o Metrics) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Metrics) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["_type"] = o.MetaType
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["updated_at"] = o.UpdatedAt
	toSerialize["name"] = o.Name
	toSerialize["description"] = o.Description
	toSerialize["unit"] = o.Unit
	toSerialize["query"] = o.Query
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Metrics) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"_type",
		"created_at",
		"updated_at",
		"name",
		"description",
		"unit",
		"query",
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

	varMetrics := _Metrics{}

	err = json.Unmarshal(data, &varMetrics)

	if err != nil {
		return err
	}

	*o = Metrics(varMetrics)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "_type")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "unit")
		delete(additionalProperties, "query")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMetrics struct {
	value *Metrics
	isSet bool
}

func (v NullableMetrics) Get() *Metrics {
	return v.value
}

func (v *NullableMetrics) Set(val *Metrics) {
	v.value = val
	v.isSet = true
}

func (v NullableMetrics) IsSet() bool {
	return v.isSet
}

func (v *NullableMetrics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetrics(val *Metrics) *NullableMetrics {
	return &NullableMetrics{value: val, isSet: true}
}

func (v NullableMetrics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetrics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


