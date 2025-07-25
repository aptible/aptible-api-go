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

// checks if the VhostAcmeConfigurationChallengesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VhostAcmeConfigurationChallengesInner{}

// VhostAcmeConfigurationChallengesInner struct for VhostAcmeConfigurationChallengesInner
type VhostAcmeConfigurationChallengesInner struct {
	Method *string `json:"method,omitempty"`
	From *UpdateDashboardRequest `json:"from,omitempty"`
	To []VhostAcmeConfigurationChallengesInnerToInner `json:"to,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VhostAcmeConfigurationChallengesInner VhostAcmeConfigurationChallengesInner

// NewVhostAcmeConfigurationChallengesInner instantiates a new VhostAcmeConfigurationChallengesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVhostAcmeConfigurationChallengesInner() *VhostAcmeConfigurationChallengesInner {
	this := VhostAcmeConfigurationChallengesInner{}
	return &this
}

// NewVhostAcmeConfigurationChallengesInnerWithDefaults instantiates a new VhostAcmeConfigurationChallengesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVhostAcmeConfigurationChallengesInnerWithDefaults() *VhostAcmeConfigurationChallengesInner {
	this := VhostAcmeConfigurationChallengesInner{}
	return &this
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *VhostAcmeConfigurationChallengesInner) GetMethod() string {
	if o == nil || IsNil(o.Method) {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VhostAcmeConfigurationChallengesInner) GetMethodOk() (*string, bool) {
	if o == nil || IsNil(o.Method) {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *VhostAcmeConfigurationChallengesInner) HasMethod() bool {
	if o != nil && !IsNil(o.Method) {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *VhostAcmeConfigurationChallengesInner) SetMethod(v string) {
	o.Method = &v
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *VhostAcmeConfigurationChallengesInner) GetFrom() UpdateDashboardRequest {
	if o == nil || IsNil(o.From) {
		var ret UpdateDashboardRequest
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VhostAcmeConfigurationChallengesInner) GetFromOk() (*UpdateDashboardRequest, bool) {
	if o == nil || IsNil(o.From) {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *VhostAcmeConfigurationChallengesInner) HasFrom() bool {
	if o != nil && !IsNil(o.From) {
		return true
	}

	return false
}

// SetFrom gets a reference to the given UpdateDashboardRequest and assigns it to the From field.
func (o *VhostAcmeConfigurationChallengesInner) SetFrom(v UpdateDashboardRequest) {
	o.From = &v
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *VhostAcmeConfigurationChallengesInner) GetTo() []VhostAcmeConfigurationChallengesInnerToInner {
	if o == nil || IsNil(o.To) {
		var ret []VhostAcmeConfigurationChallengesInnerToInner
		return ret
	}
	return o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VhostAcmeConfigurationChallengesInner) GetToOk() ([]VhostAcmeConfigurationChallengesInnerToInner, bool) {
	if o == nil || IsNil(o.To) {
		return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *VhostAcmeConfigurationChallengesInner) HasTo() bool {
	if o != nil && !IsNil(o.To) {
		return true
	}

	return false
}

// SetTo gets a reference to the given []VhostAcmeConfigurationChallengesInnerToInner and assigns it to the To field.
func (o *VhostAcmeConfigurationChallengesInner) SetTo(v []VhostAcmeConfigurationChallengesInnerToInner) {
	o.To = v
}

func (o VhostAcmeConfigurationChallengesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VhostAcmeConfigurationChallengesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Method) {
		toSerialize["method"] = o.Method
	}
	if !IsNil(o.From) {
		toSerialize["from"] = o.From
	}
	if !IsNil(o.To) {
		toSerialize["to"] = o.To
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *VhostAcmeConfigurationChallengesInner) UnmarshalJSON(data []byte) (err error) {
	varVhostAcmeConfigurationChallengesInner := _VhostAcmeConfigurationChallengesInner{}

	err = json.Unmarshal(data, &varVhostAcmeConfigurationChallengesInner)

	if err != nil {
		return err
	}

	*o = VhostAcmeConfigurationChallengesInner(varVhostAcmeConfigurationChallengesInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "method")
		delete(additionalProperties, "from")
		delete(additionalProperties, "to")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVhostAcmeConfigurationChallengesInner struct {
	value *VhostAcmeConfigurationChallengesInner
	isSet bool
}

func (v NullableVhostAcmeConfigurationChallengesInner) Get() *VhostAcmeConfigurationChallengesInner {
	return v.value
}

func (v *NullableVhostAcmeConfigurationChallengesInner) Set(val *VhostAcmeConfigurationChallengesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableVhostAcmeConfigurationChallengesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableVhostAcmeConfigurationChallengesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVhostAcmeConfigurationChallengesInner(val *VhostAcmeConfigurationChallengesInner) *NullableVhostAcmeConfigurationChallengesInner {
	return &NullableVhostAcmeConfigurationChallengesInner{value: val, isSet: true}
}

func (v NullableVhostAcmeConfigurationChallengesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVhostAcmeConfigurationChallengesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


