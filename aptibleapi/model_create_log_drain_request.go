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

// checks if the CreateLogDrainRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateLogDrainRequest{}

// CreateLogDrainRequest struct for CreateLogDrainRequest
type CreateLogDrainRequest struct {
	Handle string `json:"handle"`
	DrainType string `json:"drain_type"`
	// For syslog-tls-tcp drains
	DrainHost *string `json:"drain_host,omitempty"`
	// For syslog-tls-tcp drains
	DrainPort *int32 `json:"drain_port,omitempty"`
	// DEPRECATED: For legacy Elasticsearch drains
	DrainUsername *string `json:"drain_username,omitempty"`
	// For Tail drains
	DrainPassword *string `json:"drain_password,omitempty"`
	// For Deploy-hosted Elasticsearch drains
	Database *string `json:"database,omitempty"`
	// For Deploy-hosted Elasticsearch drains
	DatabaseId *int32 `json:"database_id,omitempty"`
	// For HTTPs Post drains
	Url *string `json:"url,omitempty"`
	LoggingToken *string `json:"logging_token,omitempty"`
	DrainApps *bool `json:"drain_apps,omitempty"`
	DrainDatabases *bool `json:"drain_databases,omitempty"`
	DrainEphemeralSessions *bool `json:"drain_ephemeral_sessions,omitempty"`
	DrainProxies *bool `json:"drain_proxies,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateLogDrainRequest CreateLogDrainRequest

// NewCreateLogDrainRequest instantiates a new CreateLogDrainRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateLogDrainRequest(handle string, drainType string) *CreateLogDrainRequest {
	this := CreateLogDrainRequest{}
	this.Handle = handle
	this.DrainType = drainType
	return &this
}

// NewCreateLogDrainRequestWithDefaults instantiates a new CreateLogDrainRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateLogDrainRequestWithDefaults() *CreateLogDrainRequest {
	this := CreateLogDrainRequest{}
	return &this
}

// GetHandle returns the Handle field value
func (o *CreateLogDrainRequest) GetHandle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Handle
}

// GetHandleOk returns a tuple with the Handle field value
// and a boolean to check if the value has been set.
func (o *CreateLogDrainRequest) GetHandleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Handle, true
}

// SetHandle sets field value
func (o *CreateLogDrainRequest) SetHandle(v string) {
	o.Handle = v
}

// GetDrainType returns the DrainType field value
func (o *CreateLogDrainRequest) GetDrainType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DrainType
}

// GetDrainTypeOk returns a tuple with the DrainType field value
// and a boolean to check if the value has been set.
func (o *CreateLogDrainRequest) GetDrainTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DrainType, true
}

// SetDrainType sets field value
func (o *CreateLogDrainRequest) SetDrainType(v string) {
	o.DrainType = v
}

// GetDrainHost returns the DrainHost field value if set, zero value otherwise.
func (o *CreateLogDrainRequest) GetDrainHost() string {
	if o == nil || IsNil(o.DrainHost) {
		var ret string
		return ret
	}
	return *o.DrainHost
}

// GetDrainHostOk returns a tuple with the DrainHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLogDrainRequest) GetDrainHostOk() (*string, bool) {
	if o == nil || IsNil(o.DrainHost) {
		return nil, false
	}
	return o.DrainHost, true
}

// HasDrainHost returns a boolean if a field has been set.
func (o *CreateLogDrainRequest) HasDrainHost() bool {
	if o != nil && !IsNil(o.DrainHost) {
		return true
	}

	return false
}

// SetDrainHost gets a reference to the given string and assigns it to the DrainHost field.
func (o *CreateLogDrainRequest) SetDrainHost(v string) {
	o.DrainHost = &v
}

// GetDrainPort returns the DrainPort field value if set, zero value otherwise.
func (o *CreateLogDrainRequest) GetDrainPort() int32 {
	if o == nil || IsNil(o.DrainPort) {
		var ret int32
		return ret
	}
	return *o.DrainPort
}

// GetDrainPortOk returns a tuple with the DrainPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLogDrainRequest) GetDrainPortOk() (*int32, bool) {
	if o == nil || IsNil(o.DrainPort) {
		return nil, false
	}
	return o.DrainPort, true
}

// HasDrainPort returns a boolean if a field has been set.
func (o *CreateLogDrainRequest) HasDrainPort() bool {
	if o != nil && !IsNil(o.DrainPort) {
		return true
	}

	return false
}

// SetDrainPort gets a reference to the given int32 and assigns it to the DrainPort field.
func (o *CreateLogDrainRequest) SetDrainPort(v int32) {
	o.DrainPort = &v
}

// GetDrainUsername returns the DrainUsername field value if set, zero value otherwise.
func (o *CreateLogDrainRequest) GetDrainUsername() string {
	if o == nil || IsNil(o.DrainUsername) {
		var ret string
		return ret
	}
	return *o.DrainUsername
}

// GetDrainUsernameOk returns a tuple with the DrainUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLogDrainRequest) GetDrainUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.DrainUsername) {
		return nil, false
	}
	return o.DrainUsername, true
}

// HasDrainUsername returns a boolean if a field has been set.
func (o *CreateLogDrainRequest) HasDrainUsername() bool {
	if o != nil && !IsNil(o.DrainUsername) {
		return true
	}

	return false
}

// SetDrainUsername gets a reference to the given string and assigns it to the DrainUsername field.
func (o *CreateLogDrainRequest) SetDrainUsername(v string) {
	o.DrainUsername = &v
}

// GetDrainPassword returns the DrainPassword field value if set, zero value otherwise.
func (o *CreateLogDrainRequest) GetDrainPassword() string {
	if o == nil || IsNil(o.DrainPassword) {
		var ret string
		return ret
	}
	return *o.DrainPassword
}

// GetDrainPasswordOk returns a tuple with the DrainPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLogDrainRequest) GetDrainPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.DrainPassword) {
		return nil, false
	}
	return o.DrainPassword, true
}

// HasDrainPassword returns a boolean if a field has been set.
func (o *CreateLogDrainRequest) HasDrainPassword() bool {
	if o != nil && !IsNil(o.DrainPassword) {
		return true
	}

	return false
}

// SetDrainPassword gets a reference to the given string and assigns it to the DrainPassword field.
func (o *CreateLogDrainRequest) SetDrainPassword(v string) {
	o.DrainPassword = &v
}

// GetDatabase returns the Database field value if set, zero value otherwise.
func (o *CreateLogDrainRequest) GetDatabase() string {
	if o == nil || IsNil(o.Database) {
		var ret string
		return ret
	}
	return *o.Database
}

// GetDatabaseOk returns a tuple with the Database field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLogDrainRequest) GetDatabaseOk() (*string, bool) {
	if o == nil || IsNil(o.Database) {
		return nil, false
	}
	return o.Database, true
}

// HasDatabase returns a boolean if a field has been set.
func (o *CreateLogDrainRequest) HasDatabase() bool {
	if o != nil && !IsNil(o.Database) {
		return true
	}

	return false
}

// SetDatabase gets a reference to the given string and assigns it to the Database field.
func (o *CreateLogDrainRequest) SetDatabase(v string) {
	o.Database = &v
}

// GetDatabaseId returns the DatabaseId field value if set, zero value otherwise.
func (o *CreateLogDrainRequest) GetDatabaseId() int32 {
	if o == nil || IsNil(o.DatabaseId) {
		var ret int32
		return ret
	}
	return *o.DatabaseId
}

// GetDatabaseIdOk returns a tuple with the DatabaseId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLogDrainRequest) GetDatabaseIdOk() (*int32, bool) {
	if o == nil || IsNil(o.DatabaseId) {
		return nil, false
	}
	return o.DatabaseId, true
}

// HasDatabaseId returns a boolean if a field has been set.
func (o *CreateLogDrainRequest) HasDatabaseId() bool {
	if o != nil && !IsNil(o.DatabaseId) {
		return true
	}

	return false
}

// SetDatabaseId gets a reference to the given int32 and assigns it to the DatabaseId field.
func (o *CreateLogDrainRequest) SetDatabaseId(v int32) {
	o.DatabaseId = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *CreateLogDrainRequest) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLogDrainRequest) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *CreateLogDrainRequest) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *CreateLogDrainRequest) SetUrl(v string) {
	o.Url = &v
}

// GetLoggingToken returns the LoggingToken field value if set, zero value otherwise.
func (o *CreateLogDrainRequest) GetLoggingToken() string {
	if o == nil || IsNil(o.LoggingToken) {
		var ret string
		return ret
	}
	return *o.LoggingToken
}

// GetLoggingTokenOk returns a tuple with the LoggingToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLogDrainRequest) GetLoggingTokenOk() (*string, bool) {
	if o == nil || IsNil(o.LoggingToken) {
		return nil, false
	}
	return o.LoggingToken, true
}

// HasLoggingToken returns a boolean if a field has been set.
func (o *CreateLogDrainRequest) HasLoggingToken() bool {
	if o != nil && !IsNil(o.LoggingToken) {
		return true
	}

	return false
}

// SetLoggingToken gets a reference to the given string and assigns it to the LoggingToken field.
func (o *CreateLogDrainRequest) SetLoggingToken(v string) {
	o.LoggingToken = &v
}

// GetDrainApps returns the DrainApps field value if set, zero value otherwise.
func (o *CreateLogDrainRequest) GetDrainApps() bool {
	if o == nil || IsNil(o.DrainApps) {
		var ret bool
		return ret
	}
	return *o.DrainApps
}

// GetDrainAppsOk returns a tuple with the DrainApps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLogDrainRequest) GetDrainAppsOk() (*bool, bool) {
	if o == nil || IsNil(o.DrainApps) {
		return nil, false
	}
	return o.DrainApps, true
}

// HasDrainApps returns a boolean if a field has been set.
func (o *CreateLogDrainRequest) HasDrainApps() bool {
	if o != nil && !IsNil(o.DrainApps) {
		return true
	}

	return false
}

// SetDrainApps gets a reference to the given bool and assigns it to the DrainApps field.
func (o *CreateLogDrainRequest) SetDrainApps(v bool) {
	o.DrainApps = &v
}

// GetDrainDatabases returns the DrainDatabases field value if set, zero value otherwise.
func (o *CreateLogDrainRequest) GetDrainDatabases() bool {
	if o == nil || IsNil(o.DrainDatabases) {
		var ret bool
		return ret
	}
	return *o.DrainDatabases
}

// GetDrainDatabasesOk returns a tuple with the DrainDatabases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLogDrainRequest) GetDrainDatabasesOk() (*bool, bool) {
	if o == nil || IsNil(o.DrainDatabases) {
		return nil, false
	}
	return o.DrainDatabases, true
}

// HasDrainDatabases returns a boolean if a field has been set.
func (o *CreateLogDrainRequest) HasDrainDatabases() bool {
	if o != nil && !IsNil(o.DrainDatabases) {
		return true
	}

	return false
}

// SetDrainDatabases gets a reference to the given bool and assigns it to the DrainDatabases field.
func (o *CreateLogDrainRequest) SetDrainDatabases(v bool) {
	o.DrainDatabases = &v
}

// GetDrainEphemeralSessions returns the DrainEphemeralSessions field value if set, zero value otherwise.
func (o *CreateLogDrainRequest) GetDrainEphemeralSessions() bool {
	if o == nil || IsNil(o.DrainEphemeralSessions) {
		var ret bool
		return ret
	}
	return *o.DrainEphemeralSessions
}

// GetDrainEphemeralSessionsOk returns a tuple with the DrainEphemeralSessions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLogDrainRequest) GetDrainEphemeralSessionsOk() (*bool, bool) {
	if o == nil || IsNil(o.DrainEphemeralSessions) {
		return nil, false
	}
	return o.DrainEphemeralSessions, true
}

// HasDrainEphemeralSessions returns a boolean if a field has been set.
func (o *CreateLogDrainRequest) HasDrainEphemeralSessions() bool {
	if o != nil && !IsNil(o.DrainEphemeralSessions) {
		return true
	}

	return false
}

// SetDrainEphemeralSessions gets a reference to the given bool and assigns it to the DrainEphemeralSessions field.
func (o *CreateLogDrainRequest) SetDrainEphemeralSessions(v bool) {
	o.DrainEphemeralSessions = &v
}

// GetDrainProxies returns the DrainProxies field value if set, zero value otherwise.
func (o *CreateLogDrainRequest) GetDrainProxies() bool {
	if o == nil || IsNil(o.DrainProxies) {
		var ret bool
		return ret
	}
	return *o.DrainProxies
}

// GetDrainProxiesOk returns a tuple with the DrainProxies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLogDrainRequest) GetDrainProxiesOk() (*bool, bool) {
	if o == nil || IsNil(o.DrainProxies) {
		return nil, false
	}
	return o.DrainProxies, true
}

// HasDrainProxies returns a boolean if a field has been set.
func (o *CreateLogDrainRequest) HasDrainProxies() bool {
	if o != nil && !IsNil(o.DrainProxies) {
		return true
	}

	return false
}

// SetDrainProxies gets a reference to the given bool and assigns it to the DrainProxies field.
func (o *CreateLogDrainRequest) SetDrainProxies(v bool) {
	o.DrainProxies = &v
}

func (o CreateLogDrainRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateLogDrainRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["handle"] = o.Handle
	toSerialize["drain_type"] = o.DrainType
	if !IsNil(o.DrainHost) {
		toSerialize["drain_host"] = o.DrainHost
	}
	if !IsNil(o.DrainPort) {
		toSerialize["drain_port"] = o.DrainPort
	}
	if !IsNil(o.DrainUsername) {
		toSerialize["drain_username"] = o.DrainUsername
	}
	if !IsNil(o.DrainPassword) {
		toSerialize["drain_password"] = o.DrainPassword
	}
	if !IsNil(o.Database) {
		toSerialize["database"] = o.Database
	}
	if !IsNil(o.DatabaseId) {
		toSerialize["database_id"] = o.DatabaseId
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.LoggingToken) {
		toSerialize["logging_token"] = o.LoggingToken
	}
	if !IsNil(o.DrainApps) {
		toSerialize["drain_apps"] = o.DrainApps
	}
	if !IsNil(o.DrainDatabases) {
		toSerialize["drain_databases"] = o.DrainDatabases
	}
	if !IsNil(o.DrainEphemeralSessions) {
		toSerialize["drain_ephemeral_sessions"] = o.DrainEphemeralSessions
	}
	if !IsNil(o.DrainProxies) {
		toSerialize["drain_proxies"] = o.DrainProxies
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateLogDrainRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"handle",
		"drain_type",
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

	varCreateLogDrainRequest := _CreateLogDrainRequest{}

	err = json.Unmarshal(data, &varCreateLogDrainRequest)

	if err != nil {
		return err
	}

	*o = CreateLogDrainRequest(varCreateLogDrainRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "handle")
		delete(additionalProperties, "drain_type")
		delete(additionalProperties, "drain_host")
		delete(additionalProperties, "drain_port")
		delete(additionalProperties, "drain_username")
		delete(additionalProperties, "drain_password")
		delete(additionalProperties, "database")
		delete(additionalProperties, "database_id")
		delete(additionalProperties, "url")
		delete(additionalProperties, "logging_token")
		delete(additionalProperties, "drain_apps")
		delete(additionalProperties, "drain_databases")
		delete(additionalProperties, "drain_ephemeral_sessions")
		delete(additionalProperties, "drain_proxies")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateLogDrainRequest struct {
	value *CreateLogDrainRequest
	isSet bool
}

func (v NullableCreateLogDrainRequest) Get() *CreateLogDrainRequest {
	return v.value
}

func (v *NullableCreateLogDrainRequest) Set(val *CreateLogDrainRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateLogDrainRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateLogDrainRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateLogDrainRequest(val *CreateLogDrainRequest) *NullableCreateLogDrainRequest {
	return &NullableCreateLogDrainRequest{value: val, isSet: true}
}

func (v NullableCreateLogDrainRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateLogDrainRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


