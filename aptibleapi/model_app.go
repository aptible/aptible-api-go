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

// checks if the App type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &App{}

// App struct for App
type App struct {
	Id int32 `json:"id"`
	Handle string `json:"handle"`
	GitRepo string `json:"git_repo"`
	MetaType string `json:"_type"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Status string `json:"status"`
	DeploymentMethod NullableString `json:"deployment_method"`
	Embedded AppEmbedded `json:"_embedded"`
	Links *AppLinks `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _App App

// NewApp instantiates a new App object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApp(id int32, handle string, gitRepo string, metaType string, createdAt string, updatedAt string, status string, deploymentMethod NullableString, embedded AppEmbedded) *App {
	this := App{}
	this.Id = id
	this.Handle = handle
	this.GitRepo = gitRepo
	this.MetaType = metaType
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.Status = status
	this.DeploymentMethod = deploymentMethod
	this.Embedded = embedded
	return &this
}

// NewAppWithDefaults instantiates a new App object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppWithDefaults() *App {
	this := App{}
	return &this
}

// GetId returns the Id field value
func (o *App) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *App) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *App) SetId(v int32) {
	o.Id = v
}

// GetHandle returns the Handle field value
func (o *App) GetHandle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Handle
}

// GetHandleOk returns a tuple with the Handle field value
// and a boolean to check if the value has been set.
func (o *App) GetHandleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Handle, true
}

// SetHandle sets field value
func (o *App) SetHandle(v string) {
	o.Handle = v
}

// GetGitRepo returns the GitRepo field value
func (o *App) GetGitRepo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GitRepo
}

// GetGitRepoOk returns a tuple with the GitRepo field value
// and a boolean to check if the value has been set.
func (o *App) GetGitRepoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GitRepo, true
}

// SetGitRepo sets field value
func (o *App) SetGitRepo(v string) {
	o.GitRepo = v
}

// GetMetaType returns the MetaType field value
func (o *App) GetMetaType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MetaType
}

// GetMetaTypeOk returns a tuple with the MetaType field value
// and a boolean to check if the value has been set.
func (o *App) GetMetaTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetaType, true
}

// SetMetaType sets field value
func (o *App) SetMetaType(v string) {
	o.MetaType = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *App) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *App) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *App) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *App) GetUpdatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *App) GetUpdatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *App) SetUpdatedAt(v string) {
	o.UpdatedAt = v
}

// GetStatus returns the Status field value
func (o *App) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *App) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *App) SetStatus(v string) {
	o.Status = v
}

// GetDeploymentMethod returns the DeploymentMethod field value
// If the value is explicit nil, the zero value for string will be returned
func (o *App) GetDeploymentMethod() string {
	if o == nil || o.DeploymentMethod.Get() == nil {
		var ret string
		return ret
	}

	return *o.DeploymentMethod.Get()
}

// GetDeploymentMethodOk returns a tuple with the DeploymentMethod field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *App) GetDeploymentMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeploymentMethod.Get(), o.DeploymentMethod.IsSet()
}

// SetDeploymentMethod sets field value
func (o *App) SetDeploymentMethod(v string) {
	o.DeploymentMethod.Set(&v)
}

// GetEmbedded returns the Embedded field value
func (o *App) GetEmbedded() AppEmbedded {
	if o == nil {
		var ret AppEmbedded
		return ret
	}

	return o.Embedded
}

// GetEmbeddedOk returns a tuple with the Embedded field value
// and a boolean to check if the value has been set.
func (o *App) GetEmbeddedOk() (*AppEmbedded, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Embedded, true
}

// SetEmbedded sets field value
func (o *App) SetEmbedded(v AppEmbedded) {
	o.Embedded = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *App) GetLinks() AppLinks {
	if o == nil || IsNil(o.Links) {
		var ret AppLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *App) GetLinksOk() (*AppLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *App) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AppLinks and assigns it to the Links field.
func (o *App) SetLinks(v AppLinks) {
	o.Links = &v
}

func (o App) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o App) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["handle"] = o.Handle
	toSerialize["git_repo"] = o.GitRepo
	toSerialize["_type"] = o.MetaType
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["updated_at"] = o.UpdatedAt
	toSerialize["status"] = o.Status
	toSerialize["deployment_method"] = o.DeploymentMethod.Get()
	toSerialize["_embedded"] = o.Embedded
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *App) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"handle",
		"git_repo",
		"_type",
		"created_at",
		"updated_at",
		"status",
		"deployment_method",
		"_embedded",
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

	varApp := _App{}

	err = json.Unmarshal(data, &varApp)

	if err != nil {
		return err
	}

	*o = App(varApp)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "handle")
		delete(additionalProperties, "git_repo")
		delete(additionalProperties, "_type")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "status")
		delete(additionalProperties, "deployment_method")
		delete(additionalProperties, "_embedded")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableApp struct {
	value *App
	isSet bool
}

func (v NullableApp) Get() *App {
	return v.value
}

func (v *NullableApp) Set(val *App) {
	v.value = val
	v.isSet = true
}

func (v NullableApp) IsSet() bool {
	return v.isSet
}

func (v *NullableApp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApp(val *App) *NullableApp {
	return &NullableApp{value: val, isSet: true}
}

func (v NullableApp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

