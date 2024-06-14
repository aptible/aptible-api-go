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

// checks if the Image type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Image{}

// Image struct for Image
type Image struct {
	Id int32 `json:"id"`
	GitRepo NullableString `json:"git_repo"`
	GitRef NullableString `json:"git_ref"`
	DockerRepo NullableString `json:"docker_repo"`
	DockerRef NullableString `json:"docker_ref"`
	DualstackHint NullableInt32 `json:"dualstack_hint"`
	Platform NullableString `json:"platform"`
	Release NullableString `json:"release"`
	Scan NullableString `json:"scan"`
	ExposedPorts []int32 `json:"exposed_ports"`
	MetaType string `json:"_type"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Links *ImageLinks `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Image Image

// NewImage instantiates a new Image object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImage(id int32, gitRepo NullableString, gitRef NullableString, dockerRepo NullableString, dockerRef NullableString, dualstackHint NullableInt32, platform NullableString, release NullableString, scan NullableString, exposedPorts []int32, metaType string, createdAt string, updatedAt string) *Image {
	this := Image{}
	this.Id = id
	this.GitRepo = gitRepo
	this.GitRef = gitRef
	this.DockerRepo = dockerRepo
	this.DockerRef = dockerRef
	this.DualstackHint = dualstackHint
	this.Platform = platform
	this.Release = release
	this.Scan = scan
	this.ExposedPorts = exposedPorts
	this.MetaType = metaType
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	return &this
}

// NewImageWithDefaults instantiates a new Image object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImageWithDefaults() *Image {
	this := Image{}
	return &this
}

// GetId returns the Id field value
func (o *Image) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Image) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Image) SetId(v int32) {
	o.Id = v
}

// GetGitRepo returns the GitRepo field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Image) GetGitRepo() string {
	if o == nil || o.GitRepo.Get() == nil {
		var ret string
		return ret
	}

	return *o.GitRepo.Get()
}

// GetGitRepoOk returns a tuple with the GitRepo field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Image) GetGitRepoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.GitRepo.Get(), o.GitRepo.IsSet()
}

// SetGitRepo sets field value
func (o *Image) SetGitRepo(v string) {
	o.GitRepo.Set(&v)
}

// GetGitRef returns the GitRef field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Image) GetGitRef() string {
	if o == nil || o.GitRef.Get() == nil {
		var ret string
		return ret
	}

	return *o.GitRef.Get()
}

// GetGitRefOk returns a tuple with the GitRef field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Image) GetGitRefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.GitRef.Get(), o.GitRef.IsSet()
}

// SetGitRef sets field value
func (o *Image) SetGitRef(v string) {
	o.GitRef.Set(&v)
}

// GetDockerRepo returns the DockerRepo field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Image) GetDockerRepo() string {
	if o == nil || o.DockerRepo.Get() == nil {
		var ret string
		return ret
	}

	return *o.DockerRepo.Get()
}

// GetDockerRepoOk returns a tuple with the DockerRepo field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Image) GetDockerRepoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DockerRepo.Get(), o.DockerRepo.IsSet()
}

// SetDockerRepo sets field value
func (o *Image) SetDockerRepo(v string) {
	o.DockerRepo.Set(&v)
}

// GetDockerRef returns the DockerRef field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Image) GetDockerRef() string {
	if o == nil || o.DockerRef.Get() == nil {
		var ret string
		return ret
	}

	return *o.DockerRef.Get()
}

// GetDockerRefOk returns a tuple with the DockerRef field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Image) GetDockerRefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DockerRef.Get(), o.DockerRef.IsSet()
}

// SetDockerRef sets field value
func (o *Image) SetDockerRef(v string) {
	o.DockerRef.Set(&v)
}

// GetDualstackHint returns the DualstackHint field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *Image) GetDualstackHint() int32 {
	if o == nil || o.DualstackHint.Get() == nil {
		var ret int32
		return ret
	}

	return *o.DualstackHint.Get()
}

// GetDualstackHintOk returns a tuple with the DualstackHint field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Image) GetDualstackHintOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.DualstackHint.Get(), o.DualstackHint.IsSet()
}

// SetDualstackHint sets field value
func (o *Image) SetDualstackHint(v int32) {
	o.DualstackHint.Set(&v)
}

// GetPlatform returns the Platform field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Image) GetPlatform() string {
	if o == nil || o.Platform.Get() == nil {
		var ret string
		return ret
	}

	return *o.Platform.Get()
}

// GetPlatformOk returns a tuple with the Platform field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Image) GetPlatformOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Platform.Get(), o.Platform.IsSet()
}

// SetPlatform sets field value
func (o *Image) SetPlatform(v string) {
	o.Platform.Set(&v)
}

// GetRelease returns the Release field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Image) GetRelease() string {
	if o == nil || o.Release.Get() == nil {
		var ret string
		return ret
	}

	return *o.Release.Get()
}

// GetReleaseOk returns a tuple with the Release field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Image) GetReleaseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Release.Get(), o.Release.IsSet()
}

// SetRelease sets field value
func (o *Image) SetRelease(v string) {
	o.Release.Set(&v)
}

// GetScan returns the Scan field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Image) GetScan() string {
	if o == nil || o.Scan.Get() == nil {
		var ret string
		return ret
	}

	return *o.Scan.Get()
}

// GetScanOk returns a tuple with the Scan field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Image) GetScanOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Scan.Get(), o.Scan.IsSet()
}

// SetScan sets field value
func (o *Image) SetScan(v string) {
	o.Scan.Set(&v)
}

// GetExposedPorts returns the ExposedPorts field value
func (o *Image) GetExposedPorts() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.ExposedPorts
}

// GetExposedPortsOk returns a tuple with the ExposedPorts field value
// and a boolean to check if the value has been set.
func (o *Image) GetExposedPortsOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExposedPorts, true
}

// SetExposedPorts sets field value
func (o *Image) SetExposedPorts(v []int32) {
	o.ExposedPorts = v
}

// GetMetaType returns the MetaType field value
func (o *Image) GetMetaType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MetaType
}

// GetMetaTypeOk returns a tuple with the MetaType field value
// and a boolean to check if the value has been set.
func (o *Image) GetMetaTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetaType, true
}

// SetMetaType sets field value
func (o *Image) SetMetaType(v string) {
	o.MetaType = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *Image) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Image) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *Image) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *Image) GetUpdatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *Image) GetUpdatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *Image) SetUpdatedAt(v string) {
	o.UpdatedAt = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *Image) GetLinks() ImageLinks {
	if o == nil || IsNil(o.Links) {
		var ret ImageLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Image) GetLinksOk() (*ImageLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *Image) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given ImageLinks and assigns it to the Links field.
func (o *Image) SetLinks(v ImageLinks) {
	o.Links = &v
}

func (o Image) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Image) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["git_repo"] = o.GitRepo.Get()
	toSerialize["git_ref"] = o.GitRef.Get()
	toSerialize["docker_repo"] = o.DockerRepo.Get()
	toSerialize["docker_ref"] = o.DockerRef.Get()
	toSerialize["dualstack_hint"] = o.DualstackHint.Get()
	toSerialize["platform"] = o.Platform.Get()
	toSerialize["release"] = o.Release.Get()
	toSerialize["scan"] = o.Scan.Get()
	toSerialize["exposed_ports"] = o.ExposedPorts
	toSerialize["_type"] = o.MetaType
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

func (o *Image) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"git_repo",
		"git_ref",
		"docker_repo",
		"docker_ref",
		"dualstack_hint",
		"platform",
		"release",
		"scan",
		"exposed_ports",
		"_type",
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

	varImage := _Image{}

	err = json.Unmarshal(data, &varImage)

	if err != nil {
		return err
	}

	*o = Image(varImage)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "git_repo")
		delete(additionalProperties, "git_ref")
		delete(additionalProperties, "docker_repo")
		delete(additionalProperties, "docker_ref")
		delete(additionalProperties, "dualstack_hint")
		delete(additionalProperties, "platform")
		delete(additionalProperties, "release")
		delete(additionalProperties, "scan")
		delete(additionalProperties, "exposed_ports")
		delete(additionalProperties, "_type")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableImage struct {
	value *Image
	isSet bool
}

func (v NullableImage) Get() *Image {
	return v.value
}

func (v *NullableImage) Set(val *Image) {
	v.value = val
	v.isSet = true
}

func (v NullableImage) IsSet() bool {
	return v.isSet
}

func (v *NullableImage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImage(val *Image) *NullableImage {
	return &NullableImage{value: val, isSet: true}
}

func (v NullableImage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


