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

// checks if the EphemeralSession type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EphemeralSession{}

// EphemeralSession struct for EphemeralSession
type EphemeralSession struct {
	Id int32 `json:"id"`
	AwsInstanceId string `json:"aws_instance_id"`
	Host string `json:"host"`
	Name string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	MetaType string `json:"_type"`
	Links *EphemeralSessionLinks `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EphemeralSession EphemeralSession

// NewEphemeralSession instantiates a new EphemeralSession object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEphemeralSession(id int32, awsInstanceId string, host string, name string, createdAt string, updatedAt string, metaType string) *EphemeralSession {
	this := EphemeralSession{}
	this.Id = id
	this.AwsInstanceId = awsInstanceId
	this.Host = host
	this.Name = name
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.MetaType = metaType
	return &this
}

// NewEphemeralSessionWithDefaults instantiates a new EphemeralSession object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEphemeralSessionWithDefaults() *EphemeralSession {
	this := EphemeralSession{}
	return &this
}

// GetId returns the Id field value
func (o *EphemeralSession) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *EphemeralSession) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *EphemeralSession) SetId(v int32) {
	o.Id = v
}

// GetAwsInstanceId returns the AwsInstanceId field value
func (o *EphemeralSession) GetAwsInstanceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AwsInstanceId
}

// GetAwsInstanceIdOk returns a tuple with the AwsInstanceId field value
// and a boolean to check if the value has been set.
func (o *EphemeralSession) GetAwsInstanceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AwsInstanceId, true
}

// SetAwsInstanceId sets field value
func (o *EphemeralSession) SetAwsInstanceId(v string) {
	o.AwsInstanceId = v
}

// GetHost returns the Host field value
func (o *EphemeralSession) GetHost() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Host
}

// GetHostOk returns a tuple with the Host field value
// and a boolean to check if the value has been set.
func (o *EphemeralSession) GetHostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Host, true
}

// SetHost sets field value
func (o *EphemeralSession) SetHost(v string) {
	o.Host = v
}

// GetName returns the Name field value
func (o *EphemeralSession) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EphemeralSession) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *EphemeralSession) SetName(v string) {
	o.Name = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *EphemeralSession) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *EphemeralSession) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *EphemeralSession) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *EphemeralSession) GetUpdatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *EphemeralSession) GetUpdatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *EphemeralSession) SetUpdatedAt(v string) {
	o.UpdatedAt = v
}

// GetMetaType returns the MetaType field value
func (o *EphemeralSession) GetMetaType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MetaType
}

// GetMetaTypeOk returns a tuple with the MetaType field value
// and a boolean to check if the value has been set.
func (o *EphemeralSession) GetMetaTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetaType, true
}

// SetMetaType sets field value
func (o *EphemeralSession) SetMetaType(v string) {
	o.MetaType = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *EphemeralSession) GetLinks() EphemeralSessionLinks {
	if o == nil || IsNil(o.Links) {
		var ret EphemeralSessionLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EphemeralSession) GetLinksOk() (*EphemeralSessionLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *EphemeralSession) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given EphemeralSessionLinks and assigns it to the Links field.
func (o *EphemeralSession) SetLinks(v EphemeralSessionLinks) {
	o.Links = &v
}

func (o EphemeralSession) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EphemeralSession) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["aws_instance_id"] = o.AwsInstanceId
	toSerialize["host"] = o.Host
	toSerialize["name"] = o.Name
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["updated_at"] = o.UpdatedAt
	toSerialize["_type"] = o.MetaType
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EphemeralSession) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"aws_instance_id",
		"host",
		"name",
		"created_at",
		"updated_at",
		"_type",
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

	varEphemeralSession := _EphemeralSession{}

	err = json.Unmarshal(data, &varEphemeralSession)

	if err != nil {
		return err
	}

	*o = EphemeralSession(varEphemeralSession)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "aws_instance_id")
		delete(additionalProperties, "host")
		delete(additionalProperties, "name")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "_type")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEphemeralSession struct {
	value *EphemeralSession
	isSet bool
}

func (v NullableEphemeralSession) Get() *EphemeralSession {
	return v.value
}

func (v *NullableEphemeralSession) Set(val *EphemeralSession) {
	v.value = val
	v.isSet = true
}

func (v NullableEphemeralSession) IsSet() bool {
	return v.isSet
}

func (v *NullableEphemeralSession) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEphemeralSession(val *EphemeralSession) *NullableEphemeralSession {
	return &NullableEphemeralSession{value: val, isSet: true}
}

func (v NullableEphemeralSession) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEphemeralSession) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


