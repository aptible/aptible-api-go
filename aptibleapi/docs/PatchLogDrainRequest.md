# PatchLogDrainRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DrainApps** | Pointer to **bool** |  | [optional] 
**DrainDatabases** | Pointer to **bool** |  | [optional] 
**DrainEphemeralSessions** | Pointer to **bool** |  | [optional] 
**DrainProxies** | Pointer to **bool** |  | [optional] 

## Methods

### NewPatchLogDrainRequest

`func NewPatchLogDrainRequest() *PatchLogDrainRequest`

NewPatchLogDrainRequest instantiates a new PatchLogDrainRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchLogDrainRequestWithDefaults

`func NewPatchLogDrainRequestWithDefaults() *PatchLogDrainRequest`

NewPatchLogDrainRequestWithDefaults instantiates a new PatchLogDrainRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDrainApps

`func (o *PatchLogDrainRequest) GetDrainApps() bool`

GetDrainApps returns the DrainApps field if non-nil, zero value otherwise.

### GetDrainAppsOk

`func (o *PatchLogDrainRequest) GetDrainAppsOk() (*bool, bool)`

GetDrainAppsOk returns a tuple with the DrainApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrainApps

`func (o *PatchLogDrainRequest) SetDrainApps(v bool)`

SetDrainApps sets DrainApps field to given value.

### HasDrainApps

`func (o *PatchLogDrainRequest) HasDrainApps() bool`

HasDrainApps returns a boolean if a field has been set.

### GetDrainDatabases

`func (o *PatchLogDrainRequest) GetDrainDatabases() bool`

GetDrainDatabases returns the DrainDatabases field if non-nil, zero value otherwise.

### GetDrainDatabasesOk

`func (o *PatchLogDrainRequest) GetDrainDatabasesOk() (*bool, bool)`

GetDrainDatabasesOk returns a tuple with the DrainDatabases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrainDatabases

`func (o *PatchLogDrainRequest) SetDrainDatabases(v bool)`

SetDrainDatabases sets DrainDatabases field to given value.

### HasDrainDatabases

`func (o *PatchLogDrainRequest) HasDrainDatabases() bool`

HasDrainDatabases returns a boolean if a field has been set.

### GetDrainEphemeralSessions

`func (o *PatchLogDrainRequest) GetDrainEphemeralSessions() bool`

GetDrainEphemeralSessions returns the DrainEphemeralSessions field if non-nil, zero value otherwise.

### GetDrainEphemeralSessionsOk

`func (o *PatchLogDrainRequest) GetDrainEphemeralSessionsOk() (*bool, bool)`

GetDrainEphemeralSessionsOk returns a tuple with the DrainEphemeralSessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrainEphemeralSessions

`func (o *PatchLogDrainRequest) SetDrainEphemeralSessions(v bool)`

SetDrainEphemeralSessions sets DrainEphemeralSessions field to given value.

### HasDrainEphemeralSessions

`func (o *PatchLogDrainRequest) HasDrainEphemeralSessions() bool`

HasDrainEphemeralSessions returns a boolean if a field has been set.

### GetDrainProxies

`func (o *PatchLogDrainRequest) GetDrainProxies() bool`

GetDrainProxies returns the DrainProxies field if non-nil, zero value otherwise.

### GetDrainProxiesOk

`func (o *PatchLogDrainRequest) GetDrainProxiesOk() (*bool, bool)`

GetDrainProxiesOk returns a tuple with the DrainProxies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrainProxies

`func (o *PatchLogDrainRequest) SetDrainProxies(v bool)`

SetDrainProxies sets DrainProxies field to given value.

### HasDrainProxies

`func (o *PatchLogDrainRequest) HasDrainProxies() bool`

HasDrainProxies returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


