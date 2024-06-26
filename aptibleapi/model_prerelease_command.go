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

// checks if the PrereleaseCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PrereleaseCommand{}

// PrereleaseCommand struct for PrereleaseCommand
type PrereleaseCommand struct {
	Id int32 `json:"id"`
	MetaType string `json:"_type"`
	Command string `json:"command"`
	Index int32 `json:"index"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Links *PrereleaseCommandLinks `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PrereleaseCommand PrereleaseCommand

// NewPrereleaseCommand instantiates a new PrereleaseCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrereleaseCommand(id int32, metaType string, command string, index int32, createdAt string, updatedAt string) *PrereleaseCommand {
	this := PrereleaseCommand{}
	this.Id = id
	this.MetaType = metaType
	this.Command = command
	this.Index = index
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	return &this
}

// NewPrereleaseCommandWithDefaults instantiates a new PrereleaseCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrereleaseCommandWithDefaults() *PrereleaseCommand {
	this := PrereleaseCommand{}
	return &this
}

// GetId returns the Id field value
func (o *PrereleaseCommand) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PrereleaseCommand) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PrereleaseCommand) SetId(v int32) {
	o.Id = v
}

// GetMetaType returns the MetaType field value
func (o *PrereleaseCommand) GetMetaType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MetaType
}

// GetMetaTypeOk returns a tuple with the MetaType field value
// and a boolean to check if the value has been set.
func (o *PrereleaseCommand) GetMetaTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetaType, true
}

// SetMetaType sets field value
func (o *PrereleaseCommand) SetMetaType(v string) {
	o.MetaType = v
}

// GetCommand returns the Command field value
func (o *PrereleaseCommand) GetCommand() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Command
}

// GetCommandOk returns a tuple with the Command field value
// and a boolean to check if the value has been set.
func (o *PrereleaseCommand) GetCommandOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Command, true
}

// SetCommand sets field value
func (o *PrereleaseCommand) SetCommand(v string) {
	o.Command = v
}

// GetIndex returns the Index field value
func (o *PrereleaseCommand) GetIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *PrereleaseCommand) GetIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *PrereleaseCommand) SetIndex(v int32) {
	o.Index = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *PrereleaseCommand) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *PrereleaseCommand) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *PrereleaseCommand) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *PrereleaseCommand) GetUpdatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *PrereleaseCommand) GetUpdatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *PrereleaseCommand) SetUpdatedAt(v string) {
	o.UpdatedAt = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *PrereleaseCommand) GetLinks() PrereleaseCommandLinks {
	if o == nil || IsNil(o.Links) {
		var ret PrereleaseCommandLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrereleaseCommand) GetLinksOk() (*PrereleaseCommandLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *PrereleaseCommand) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given PrereleaseCommandLinks and assigns it to the Links field.
func (o *PrereleaseCommand) SetLinks(v PrereleaseCommandLinks) {
	o.Links = &v
}

func (o PrereleaseCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PrereleaseCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["_type"] = o.MetaType
	toSerialize["command"] = o.Command
	toSerialize["index"] = o.Index
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["updated_at"] = o.UpdatedAt
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PrereleaseCommand) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"_type",
		"command",
		"index",
		"created_at",
		"updated_at",
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

	varPrereleaseCommand := _PrereleaseCommand{}

	err = json.Unmarshal(data, &varPrereleaseCommand)

	if err != nil {
		return err
	}

	*o = PrereleaseCommand(varPrereleaseCommand)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "_type")
		delete(additionalProperties, "command")
		delete(additionalProperties, "index")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePrereleaseCommand struct {
	value *PrereleaseCommand
	isSet bool
}

func (v NullablePrereleaseCommand) Get() *PrereleaseCommand {
	return v.value
}

func (v *NullablePrereleaseCommand) Set(val *PrereleaseCommand) {
	v.value = val
	v.isSet = true
}

func (v NullablePrereleaseCommand) IsSet() bool {
	return v.isSet
}

func (v *NullablePrereleaseCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrereleaseCommand(val *PrereleaseCommand) *NullablePrereleaseCommand {
	return &NullablePrereleaseCommand{value: val, isSet: true}
}

func (v NullablePrereleaseCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrereleaseCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


