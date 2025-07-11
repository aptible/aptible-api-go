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

// checks if the CreateLlmIntegrationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateLlmIntegrationRequest{}

// CreateLlmIntegrationRequest struct for CreateLlmIntegrationRequest
type CreateLlmIntegrationRequest struct {
	ProviderType string `json:"provider_type"`
	OrganizationId string `json:"organization_id"`
	ApiKey *string `json:"api_key,omitempty"`
	BaseUrl *string `json:"base_url,omitempty"`
	OpenaiOrganization *string `json:"openai_organization,omitempty"`
	ApiVersion *string `json:"api_version,omitempty"`
	AwsAccessKeyId *string `json:"aws_access_key_id,omitempty"`
	AwsSecretAccessKey *string `json:"aws_secret_access_key,omitempty"`
	AwsRegion *string `json:"aws_region,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateLlmIntegrationRequest CreateLlmIntegrationRequest

// NewCreateLlmIntegrationRequest instantiates a new CreateLlmIntegrationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateLlmIntegrationRequest(providerType string, organizationId string) *CreateLlmIntegrationRequest {
	this := CreateLlmIntegrationRequest{}
	this.ProviderType = providerType
	this.OrganizationId = organizationId
	return &this
}

// NewCreateLlmIntegrationRequestWithDefaults instantiates a new CreateLlmIntegrationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateLlmIntegrationRequestWithDefaults() *CreateLlmIntegrationRequest {
	this := CreateLlmIntegrationRequest{}
	return &this
}

// GetProviderType returns the ProviderType field value
func (o *CreateLlmIntegrationRequest) GetProviderType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProviderType
}

// GetProviderTypeOk returns a tuple with the ProviderType field value
// and a boolean to check if the value has been set.
func (o *CreateLlmIntegrationRequest) GetProviderTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProviderType, true
}

// SetProviderType sets field value
func (o *CreateLlmIntegrationRequest) SetProviderType(v string) {
	o.ProviderType = v
}

// GetOrganizationId returns the OrganizationId field value
func (o *CreateLlmIntegrationRequest) GetOrganizationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value
// and a boolean to check if the value has been set.
func (o *CreateLlmIntegrationRequest) GetOrganizationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrganizationId, true
}

// SetOrganizationId sets field value
func (o *CreateLlmIntegrationRequest) SetOrganizationId(v string) {
	o.OrganizationId = v
}

// GetApiKey returns the ApiKey field value if set, zero value otherwise.
func (o *CreateLlmIntegrationRequest) GetApiKey() string {
	if o == nil || IsNil(o.ApiKey) {
		var ret string
		return ret
	}
	return *o.ApiKey
}

// GetApiKeyOk returns a tuple with the ApiKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLlmIntegrationRequest) GetApiKeyOk() (*string, bool) {
	if o == nil || IsNil(o.ApiKey) {
		return nil, false
	}
	return o.ApiKey, true
}

// HasApiKey returns a boolean if a field has been set.
func (o *CreateLlmIntegrationRequest) HasApiKey() bool {
	if o != nil && !IsNil(o.ApiKey) {
		return true
	}

	return false
}

// SetApiKey gets a reference to the given string and assigns it to the ApiKey field.
func (o *CreateLlmIntegrationRequest) SetApiKey(v string) {
	o.ApiKey = &v
}

// GetBaseUrl returns the BaseUrl field value if set, zero value otherwise.
func (o *CreateLlmIntegrationRequest) GetBaseUrl() string {
	if o == nil || IsNil(o.BaseUrl) {
		var ret string
		return ret
	}
	return *o.BaseUrl
}

// GetBaseUrlOk returns a tuple with the BaseUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLlmIntegrationRequest) GetBaseUrlOk() (*string, bool) {
	if o == nil || IsNil(o.BaseUrl) {
		return nil, false
	}
	return o.BaseUrl, true
}

// HasBaseUrl returns a boolean if a field has been set.
func (o *CreateLlmIntegrationRequest) HasBaseUrl() bool {
	if o != nil && !IsNil(o.BaseUrl) {
		return true
	}

	return false
}

// SetBaseUrl gets a reference to the given string and assigns it to the BaseUrl field.
func (o *CreateLlmIntegrationRequest) SetBaseUrl(v string) {
	o.BaseUrl = &v
}

// GetOpenaiOrganization returns the OpenaiOrganization field value if set, zero value otherwise.
func (o *CreateLlmIntegrationRequest) GetOpenaiOrganization() string {
	if o == nil || IsNil(o.OpenaiOrganization) {
		var ret string
		return ret
	}
	return *o.OpenaiOrganization
}

// GetOpenaiOrganizationOk returns a tuple with the OpenaiOrganization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLlmIntegrationRequest) GetOpenaiOrganizationOk() (*string, bool) {
	if o == nil || IsNil(o.OpenaiOrganization) {
		return nil, false
	}
	return o.OpenaiOrganization, true
}

// HasOpenaiOrganization returns a boolean if a field has been set.
func (o *CreateLlmIntegrationRequest) HasOpenaiOrganization() bool {
	if o != nil && !IsNil(o.OpenaiOrganization) {
		return true
	}

	return false
}

// SetOpenaiOrganization gets a reference to the given string and assigns it to the OpenaiOrganization field.
func (o *CreateLlmIntegrationRequest) SetOpenaiOrganization(v string) {
	o.OpenaiOrganization = &v
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *CreateLlmIntegrationRequest) GetApiVersion() string {
	if o == nil || IsNil(o.ApiVersion) {
		var ret string
		return ret
	}
	return *o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLlmIntegrationRequest) GetApiVersionOk() (*string, bool) {
	if o == nil || IsNil(o.ApiVersion) {
		return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *CreateLlmIntegrationRequest) HasApiVersion() bool {
	if o != nil && !IsNil(o.ApiVersion) {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given string and assigns it to the ApiVersion field.
func (o *CreateLlmIntegrationRequest) SetApiVersion(v string) {
	o.ApiVersion = &v
}

// GetAwsAccessKeyId returns the AwsAccessKeyId field value if set, zero value otherwise.
func (o *CreateLlmIntegrationRequest) GetAwsAccessKeyId() string {
	if o == nil || IsNil(o.AwsAccessKeyId) {
		var ret string
		return ret
	}
	return *o.AwsAccessKeyId
}

// GetAwsAccessKeyIdOk returns a tuple with the AwsAccessKeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLlmIntegrationRequest) GetAwsAccessKeyIdOk() (*string, bool) {
	if o == nil || IsNil(o.AwsAccessKeyId) {
		return nil, false
	}
	return o.AwsAccessKeyId, true
}

// HasAwsAccessKeyId returns a boolean if a field has been set.
func (o *CreateLlmIntegrationRequest) HasAwsAccessKeyId() bool {
	if o != nil && !IsNil(o.AwsAccessKeyId) {
		return true
	}

	return false
}

// SetAwsAccessKeyId gets a reference to the given string and assigns it to the AwsAccessKeyId field.
func (o *CreateLlmIntegrationRequest) SetAwsAccessKeyId(v string) {
	o.AwsAccessKeyId = &v
}

// GetAwsSecretAccessKey returns the AwsSecretAccessKey field value if set, zero value otherwise.
func (o *CreateLlmIntegrationRequest) GetAwsSecretAccessKey() string {
	if o == nil || IsNil(o.AwsSecretAccessKey) {
		var ret string
		return ret
	}
	return *o.AwsSecretAccessKey
}

// GetAwsSecretAccessKeyOk returns a tuple with the AwsSecretAccessKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLlmIntegrationRequest) GetAwsSecretAccessKeyOk() (*string, bool) {
	if o == nil || IsNil(o.AwsSecretAccessKey) {
		return nil, false
	}
	return o.AwsSecretAccessKey, true
}

// HasAwsSecretAccessKey returns a boolean if a field has been set.
func (o *CreateLlmIntegrationRequest) HasAwsSecretAccessKey() bool {
	if o != nil && !IsNil(o.AwsSecretAccessKey) {
		return true
	}

	return false
}

// SetAwsSecretAccessKey gets a reference to the given string and assigns it to the AwsSecretAccessKey field.
func (o *CreateLlmIntegrationRequest) SetAwsSecretAccessKey(v string) {
	o.AwsSecretAccessKey = &v
}

// GetAwsRegion returns the AwsRegion field value if set, zero value otherwise.
func (o *CreateLlmIntegrationRequest) GetAwsRegion() string {
	if o == nil || IsNil(o.AwsRegion) {
		var ret string
		return ret
	}
	return *o.AwsRegion
}

// GetAwsRegionOk returns a tuple with the AwsRegion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLlmIntegrationRequest) GetAwsRegionOk() (*string, bool) {
	if o == nil || IsNil(o.AwsRegion) {
		return nil, false
	}
	return o.AwsRegion, true
}

// HasAwsRegion returns a boolean if a field has been set.
func (o *CreateLlmIntegrationRequest) HasAwsRegion() bool {
	if o != nil && !IsNil(o.AwsRegion) {
		return true
	}

	return false
}

// SetAwsRegion gets a reference to the given string and assigns it to the AwsRegion field.
func (o *CreateLlmIntegrationRequest) SetAwsRegion(v string) {
	o.AwsRegion = &v
}

func (o CreateLlmIntegrationRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateLlmIntegrationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["provider_type"] = o.ProviderType
	toSerialize["organization_id"] = o.OrganizationId
	if !IsNil(o.ApiKey) {
		toSerialize["api_key"] = o.ApiKey
	}
	if !IsNil(o.BaseUrl) {
		toSerialize["base_url"] = o.BaseUrl
	}
	if !IsNil(o.OpenaiOrganization) {
		toSerialize["openai_organization"] = o.OpenaiOrganization
	}
	if !IsNil(o.ApiVersion) {
		toSerialize["api_version"] = o.ApiVersion
	}
	if !IsNil(o.AwsAccessKeyId) {
		toSerialize["aws_access_key_id"] = o.AwsAccessKeyId
	}
	if !IsNil(o.AwsSecretAccessKey) {
		toSerialize["aws_secret_access_key"] = o.AwsSecretAccessKey
	}
	if !IsNil(o.AwsRegion) {
		toSerialize["aws_region"] = o.AwsRegion
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateLlmIntegrationRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"provider_type",
		"organization_id",
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

	varCreateLlmIntegrationRequest := _CreateLlmIntegrationRequest{}

	err = json.Unmarshal(data, &varCreateLlmIntegrationRequest)

	if err != nil {
		return err
	}

	*o = CreateLlmIntegrationRequest(varCreateLlmIntegrationRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "provider_type")
		delete(additionalProperties, "organization_id")
		delete(additionalProperties, "api_key")
		delete(additionalProperties, "base_url")
		delete(additionalProperties, "openai_organization")
		delete(additionalProperties, "api_version")
		delete(additionalProperties, "aws_access_key_id")
		delete(additionalProperties, "aws_secret_access_key")
		delete(additionalProperties, "aws_region")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateLlmIntegrationRequest struct {
	value *CreateLlmIntegrationRequest
	isSet bool
}

func (v NullableCreateLlmIntegrationRequest) Get() *CreateLlmIntegrationRequest {
	return v.value
}

func (v *NullableCreateLlmIntegrationRequest) Set(val *CreateLlmIntegrationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateLlmIntegrationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateLlmIntegrationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateLlmIntegrationRequest(val *CreateLlmIntegrationRequest) *NullableCreateLlmIntegrationRequest {
	return &NullableCreateLlmIntegrationRequest{value: val, isSet: true}
}

func (v NullableCreateLlmIntegrationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateLlmIntegrationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


