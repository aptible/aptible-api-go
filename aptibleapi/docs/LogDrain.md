# LogDrain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**MetaType** | **string** |  | 
**Handle** | **string** |  | 
**DrainType** | **string** |  | 
**DrainHost** | **string** |  | 
**DrainPort** | **int32** |  | 
**CreatedAt** | **string** |  | 
**UpdatedAt** | **string** |  | 
**Status** | **string** |  | 
**DrainUsername** | Pointer to **NullableString** |  | [optional] 
**DrainPassword** | Pointer to **NullableString** |  | [optional] 
**DrainApps** | **bool** |  | 
**DrainDatabases** | **bool** |  | 
**DrainEphemeralSessions** | **bool** |  | 
**DrainProxies** | **bool** |  | 
**GentlemanjerryCertificate** | **NullableString** |  | 
**GentlemanjerryDockerName** | **NullableString** |  | 
**GentlemanjerryInstanceId** | **NullableString** |  | 
**GentlemanjerryHost** | **NullableString** |  | 
**GentlemanjerryPortMapping** | **[][]int32** |  | 
**GentlemanjerryAllocation** | **[]string** |  | 
**Url** | Pointer to **NullableString** |  | [optional] 
**LoggingToken** | Pointer to **NullableString** |  | [optional] 
**Links** | Pointer to [**LogDrainLinks**](LogDrainLinks.md) |  | [optional] 

## Methods

### NewLogDrain

`func NewLogDrain(id int32, metaType string, handle string, drainType string, drainHost string, drainPort int32, createdAt string, updatedAt string, status string, drainApps bool, drainDatabases bool, drainEphemeralSessions bool, drainProxies bool, gentlemanjerryCertificate NullableString, gentlemanjerryDockerName NullableString, gentlemanjerryInstanceId NullableString, gentlemanjerryHost NullableString, gentlemanjerryPortMapping [][]int32, gentlemanjerryAllocation []string, ) *LogDrain`

NewLogDrain instantiates a new LogDrain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogDrainWithDefaults

`func NewLogDrainWithDefaults() *LogDrain`

NewLogDrainWithDefaults instantiates a new LogDrain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LogDrain) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LogDrain) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LogDrain) SetId(v int32)`

SetId sets Id field to given value.


### GetMetaType

`func (o *LogDrain) GetMetaType() string`

GetMetaType returns the MetaType field if non-nil, zero value otherwise.

### GetMetaTypeOk

`func (o *LogDrain) GetMetaTypeOk() (*string, bool)`

GetMetaTypeOk returns a tuple with the MetaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaType

`func (o *LogDrain) SetMetaType(v string)`

SetMetaType sets MetaType field to given value.


### GetHandle

`func (o *LogDrain) GetHandle() string`

GetHandle returns the Handle field if non-nil, zero value otherwise.

### GetHandleOk

`func (o *LogDrain) GetHandleOk() (*string, bool)`

GetHandleOk returns a tuple with the Handle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandle

`func (o *LogDrain) SetHandle(v string)`

SetHandle sets Handle field to given value.


### GetDrainType

`func (o *LogDrain) GetDrainType() string`

GetDrainType returns the DrainType field if non-nil, zero value otherwise.

### GetDrainTypeOk

`func (o *LogDrain) GetDrainTypeOk() (*string, bool)`

GetDrainTypeOk returns a tuple with the DrainType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrainType

`func (o *LogDrain) SetDrainType(v string)`

SetDrainType sets DrainType field to given value.


### GetDrainHost

`func (o *LogDrain) GetDrainHost() string`

GetDrainHost returns the DrainHost field if non-nil, zero value otherwise.

### GetDrainHostOk

`func (o *LogDrain) GetDrainHostOk() (*string, bool)`

GetDrainHostOk returns a tuple with the DrainHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrainHost

`func (o *LogDrain) SetDrainHost(v string)`

SetDrainHost sets DrainHost field to given value.


### GetDrainPort

`func (o *LogDrain) GetDrainPort() int32`

GetDrainPort returns the DrainPort field if non-nil, zero value otherwise.

### GetDrainPortOk

`func (o *LogDrain) GetDrainPortOk() (*int32, bool)`

GetDrainPortOk returns a tuple with the DrainPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrainPort

`func (o *LogDrain) SetDrainPort(v int32)`

SetDrainPort sets DrainPort field to given value.


### GetCreatedAt

`func (o *LogDrain) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *LogDrain) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *LogDrain) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *LogDrain) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *LogDrain) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *LogDrain) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetStatus

`func (o *LogDrain) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *LogDrain) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *LogDrain) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetDrainUsername

`func (o *LogDrain) GetDrainUsername() string`

GetDrainUsername returns the DrainUsername field if non-nil, zero value otherwise.

### GetDrainUsernameOk

`func (o *LogDrain) GetDrainUsernameOk() (*string, bool)`

GetDrainUsernameOk returns a tuple with the DrainUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrainUsername

`func (o *LogDrain) SetDrainUsername(v string)`

SetDrainUsername sets DrainUsername field to given value.

### HasDrainUsername

`func (o *LogDrain) HasDrainUsername() bool`

HasDrainUsername returns a boolean if a field has been set.

### SetDrainUsernameNil

`func (o *LogDrain) SetDrainUsernameNil(b bool)`

 SetDrainUsernameNil sets the value for DrainUsername to be an explicit nil

### UnsetDrainUsername
`func (o *LogDrain) UnsetDrainUsername()`

UnsetDrainUsername ensures that no value is present for DrainUsername, not even an explicit nil
### GetDrainPassword

`func (o *LogDrain) GetDrainPassword() string`

GetDrainPassword returns the DrainPassword field if non-nil, zero value otherwise.

### GetDrainPasswordOk

`func (o *LogDrain) GetDrainPasswordOk() (*string, bool)`

GetDrainPasswordOk returns a tuple with the DrainPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrainPassword

`func (o *LogDrain) SetDrainPassword(v string)`

SetDrainPassword sets DrainPassword field to given value.

### HasDrainPassword

`func (o *LogDrain) HasDrainPassword() bool`

HasDrainPassword returns a boolean if a field has been set.

### SetDrainPasswordNil

`func (o *LogDrain) SetDrainPasswordNil(b bool)`

 SetDrainPasswordNil sets the value for DrainPassword to be an explicit nil

### UnsetDrainPassword
`func (o *LogDrain) UnsetDrainPassword()`

UnsetDrainPassword ensures that no value is present for DrainPassword, not even an explicit nil
### GetDrainApps

`func (o *LogDrain) GetDrainApps() bool`

GetDrainApps returns the DrainApps field if non-nil, zero value otherwise.

### GetDrainAppsOk

`func (o *LogDrain) GetDrainAppsOk() (*bool, bool)`

GetDrainAppsOk returns a tuple with the DrainApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrainApps

`func (o *LogDrain) SetDrainApps(v bool)`

SetDrainApps sets DrainApps field to given value.


### GetDrainDatabases

`func (o *LogDrain) GetDrainDatabases() bool`

GetDrainDatabases returns the DrainDatabases field if non-nil, zero value otherwise.

### GetDrainDatabasesOk

`func (o *LogDrain) GetDrainDatabasesOk() (*bool, bool)`

GetDrainDatabasesOk returns a tuple with the DrainDatabases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrainDatabases

`func (o *LogDrain) SetDrainDatabases(v bool)`

SetDrainDatabases sets DrainDatabases field to given value.


### GetDrainEphemeralSessions

`func (o *LogDrain) GetDrainEphemeralSessions() bool`

GetDrainEphemeralSessions returns the DrainEphemeralSessions field if non-nil, zero value otherwise.

### GetDrainEphemeralSessionsOk

`func (o *LogDrain) GetDrainEphemeralSessionsOk() (*bool, bool)`

GetDrainEphemeralSessionsOk returns a tuple with the DrainEphemeralSessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrainEphemeralSessions

`func (o *LogDrain) SetDrainEphemeralSessions(v bool)`

SetDrainEphemeralSessions sets DrainEphemeralSessions field to given value.


### GetDrainProxies

`func (o *LogDrain) GetDrainProxies() bool`

GetDrainProxies returns the DrainProxies field if non-nil, zero value otherwise.

### GetDrainProxiesOk

`func (o *LogDrain) GetDrainProxiesOk() (*bool, bool)`

GetDrainProxiesOk returns a tuple with the DrainProxies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrainProxies

`func (o *LogDrain) SetDrainProxies(v bool)`

SetDrainProxies sets DrainProxies field to given value.


### GetGentlemanjerryCertificate

`func (o *LogDrain) GetGentlemanjerryCertificate() string`

GetGentlemanjerryCertificate returns the GentlemanjerryCertificate field if non-nil, zero value otherwise.

### GetGentlemanjerryCertificateOk

`func (o *LogDrain) GetGentlemanjerryCertificateOk() (*string, bool)`

GetGentlemanjerryCertificateOk returns a tuple with the GentlemanjerryCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGentlemanjerryCertificate

`func (o *LogDrain) SetGentlemanjerryCertificate(v string)`

SetGentlemanjerryCertificate sets GentlemanjerryCertificate field to given value.


### SetGentlemanjerryCertificateNil

`func (o *LogDrain) SetGentlemanjerryCertificateNil(b bool)`

 SetGentlemanjerryCertificateNil sets the value for GentlemanjerryCertificate to be an explicit nil

### UnsetGentlemanjerryCertificate
`func (o *LogDrain) UnsetGentlemanjerryCertificate()`

UnsetGentlemanjerryCertificate ensures that no value is present for GentlemanjerryCertificate, not even an explicit nil
### GetGentlemanjerryDockerName

`func (o *LogDrain) GetGentlemanjerryDockerName() string`

GetGentlemanjerryDockerName returns the GentlemanjerryDockerName field if non-nil, zero value otherwise.

### GetGentlemanjerryDockerNameOk

`func (o *LogDrain) GetGentlemanjerryDockerNameOk() (*string, bool)`

GetGentlemanjerryDockerNameOk returns a tuple with the GentlemanjerryDockerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGentlemanjerryDockerName

`func (o *LogDrain) SetGentlemanjerryDockerName(v string)`

SetGentlemanjerryDockerName sets GentlemanjerryDockerName field to given value.


### SetGentlemanjerryDockerNameNil

`func (o *LogDrain) SetGentlemanjerryDockerNameNil(b bool)`

 SetGentlemanjerryDockerNameNil sets the value for GentlemanjerryDockerName to be an explicit nil

### UnsetGentlemanjerryDockerName
`func (o *LogDrain) UnsetGentlemanjerryDockerName()`

UnsetGentlemanjerryDockerName ensures that no value is present for GentlemanjerryDockerName, not even an explicit nil
### GetGentlemanjerryInstanceId

`func (o *LogDrain) GetGentlemanjerryInstanceId() string`

GetGentlemanjerryInstanceId returns the GentlemanjerryInstanceId field if non-nil, zero value otherwise.

### GetGentlemanjerryInstanceIdOk

`func (o *LogDrain) GetGentlemanjerryInstanceIdOk() (*string, bool)`

GetGentlemanjerryInstanceIdOk returns a tuple with the GentlemanjerryInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGentlemanjerryInstanceId

`func (o *LogDrain) SetGentlemanjerryInstanceId(v string)`

SetGentlemanjerryInstanceId sets GentlemanjerryInstanceId field to given value.


### SetGentlemanjerryInstanceIdNil

`func (o *LogDrain) SetGentlemanjerryInstanceIdNil(b bool)`

 SetGentlemanjerryInstanceIdNil sets the value for GentlemanjerryInstanceId to be an explicit nil

### UnsetGentlemanjerryInstanceId
`func (o *LogDrain) UnsetGentlemanjerryInstanceId()`

UnsetGentlemanjerryInstanceId ensures that no value is present for GentlemanjerryInstanceId, not even an explicit nil
### GetGentlemanjerryHost

`func (o *LogDrain) GetGentlemanjerryHost() string`

GetGentlemanjerryHost returns the GentlemanjerryHost field if non-nil, zero value otherwise.

### GetGentlemanjerryHostOk

`func (o *LogDrain) GetGentlemanjerryHostOk() (*string, bool)`

GetGentlemanjerryHostOk returns a tuple with the GentlemanjerryHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGentlemanjerryHost

`func (o *LogDrain) SetGentlemanjerryHost(v string)`

SetGentlemanjerryHost sets GentlemanjerryHost field to given value.


### SetGentlemanjerryHostNil

`func (o *LogDrain) SetGentlemanjerryHostNil(b bool)`

 SetGentlemanjerryHostNil sets the value for GentlemanjerryHost to be an explicit nil

### UnsetGentlemanjerryHost
`func (o *LogDrain) UnsetGentlemanjerryHost()`

UnsetGentlemanjerryHost ensures that no value is present for GentlemanjerryHost, not even an explicit nil
### GetGentlemanjerryPortMapping

`func (o *LogDrain) GetGentlemanjerryPortMapping() [][]int32`

GetGentlemanjerryPortMapping returns the GentlemanjerryPortMapping field if non-nil, zero value otherwise.

### GetGentlemanjerryPortMappingOk

`func (o *LogDrain) GetGentlemanjerryPortMappingOk() (*[][]int32, bool)`

GetGentlemanjerryPortMappingOk returns a tuple with the GentlemanjerryPortMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGentlemanjerryPortMapping

`func (o *LogDrain) SetGentlemanjerryPortMapping(v [][]int32)`

SetGentlemanjerryPortMapping sets GentlemanjerryPortMapping field to given value.


### GetGentlemanjerryAllocation

`func (o *LogDrain) GetGentlemanjerryAllocation() []string`

GetGentlemanjerryAllocation returns the GentlemanjerryAllocation field if non-nil, zero value otherwise.

### GetGentlemanjerryAllocationOk

`func (o *LogDrain) GetGentlemanjerryAllocationOk() (*[]string, bool)`

GetGentlemanjerryAllocationOk returns a tuple with the GentlemanjerryAllocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGentlemanjerryAllocation

`func (o *LogDrain) SetGentlemanjerryAllocation(v []string)`

SetGentlemanjerryAllocation sets GentlemanjerryAllocation field to given value.


### GetUrl

`func (o *LogDrain) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *LogDrain) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *LogDrain) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *LogDrain) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *LogDrain) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *LogDrain) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetLoggingToken

`func (o *LogDrain) GetLoggingToken() string`

GetLoggingToken returns the LoggingToken field if non-nil, zero value otherwise.

### GetLoggingTokenOk

`func (o *LogDrain) GetLoggingTokenOk() (*string, bool)`

GetLoggingTokenOk returns a tuple with the LoggingToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoggingToken

`func (o *LogDrain) SetLoggingToken(v string)`

SetLoggingToken sets LoggingToken field to given value.

### HasLoggingToken

`func (o *LogDrain) HasLoggingToken() bool`

HasLoggingToken returns a boolean if a field has been set.

### SetLoggingTokenNil

`func (o *LogDrain) SetLoggingTokenNil(b bool)`

 SetLoggingTokenNil sets the value for LoggingToken to be an explicit nil

### UnsetLoggingToken
`func (o *LogDrain) UnsetLoggingToken()`

UnsetLoggingToken ensures that no value is present for LoggingToken, not even an explicit nil
### GetLinks

`func (o *LogDrain) GetLinks() LogDrainLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *LogDrain) GetLinksOk() (*LogDrainLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *LogDrain) SetLinks(v LogDrainLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *LogDrain) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


