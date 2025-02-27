# DatabaseLinks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to [**DatabaseLinksAccount**](DatabaseLinksAccount.md) |  | [optional] 
**DatabaseImage** | Pointer to [**DatabaseLinksDatabaseImage**](DatabaseLinksDatabaseImage.md) |  | [optional] 
**Service** | Pointer to [**DatabaseLinksService**](DatabaseLinksService.md) |  | [optional] 
**Disk** | Pointer to [**DatabaseLinksDisk**](DatabaseLinksDisk.md) |  | [optional] 
**InitializeFrom** | Pointer to [**DatabaseLinksInitializeFrom**](DatabaseLinksInitializeFrom.md) |  | [optional] 
**CurrentConfiguration** | Pointer to [**DatabaseLinksCurrentConfiguration**](DatabaseLinksCurrentConfiguration.md) |  | [optional] 
**Dependents** | Pointer to [**DatabaseLinksDependents**](DatabaseLinksDependents.md) |  | [optional] 
**Operations** | Pointer to [**DatabaseLinksOperations**](DatabaseLinksOperations.md) |  | [optional] 
**Backups** | Pointer to [**DatabaseLinksBackups**](DatabaseLinksBackups.md) |  | [optional] 
**Configurations** | Pointer to [**DatabaseLinksConfigurations**](DatabaseLinksConfigurations.md) |  | [optional] 
**DatabaseCredentials** | Pointer to [**DatabaseLinksDatabaseCredentials**](DatabaseLinksDatabaseCredentials.md) |  | [optional] 
**Self** | Pointer to [**DatabaseLinksSelf**](DatabaseLinksSelf.md) |  | [optional] 

## Methods

### NewDatabaseLinks

`func NewDatabaseLinks() *DatabaseLinks`

NewDatabaseLinks instantiates a new DatabaseLinks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabaseLinksWithDefaults

`func NewDatabaseLinksWithDefaults() *DatabaseLinks`

NewDatabaseLinksWithDefaults instantiates a new DatabaseLinks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *DatabaseLinks) GetAccount() DatabaseLinksAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *DatabaseLinks) GetAccountOk() (*DatabaseLinksAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *DatabaseLinks) SetAccount(v DatabaseLinksAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *DatabaseLinks) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetDatabaseImage

`func (o *DatabaseLinks) GetDatabaseImage() DatabaseLinksDatabaseImage`

GetDatabaseImage returns the DatabaseImage field if non-nil, zero value otherwise.

### GetDatabaseImageOk

`func (o *DatabaseLinks) GetDatabaseImageOk() (*DatabaseLinksDatabaseImage, bool)`

GetDatabaseImageOk returns a tuple with the DatabaseImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseImage

`func (o *DatabaseLinks) SetDatabaseImage(v DatabaseLinksDatabaseImage)`

SetDatabaseImage sets DatabaseImage field to given value.

### HasDatabaseImage

`func (o *DatabaseLinks) HasDatabaseImage() bool`

HasDatabaseImage returns a boolean if a field has been set.

### GetService

`func (o *DatabaseLinks) GetService() DatabaseLinksService`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *DatabaseLinks) GetServiceOk() (*DatabaseLinksService, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *DatabaseLinks) SetService(v DatabaseLinksService)`

SetService sets Service field to given value.

### HasService

`func (o *DatabaseLinks) HasService() bool`

HasService returns a boolean if a field has been set.

### GetDisk

`func (o *DatabaseLinks) GetDisk() DatabaseLinksDisk`

GetDisk returns the Disk field if non-nil, zero value otherwise.

### GetDiskOk

`func (o *DatabaseLinks) GetDiskOk() (*DatabaseLinksDisk, bool)`

GetDiskOk returns a tuple with the Disk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisk

`func (o *DatabaseLinks) SetDisk(v DatabaseLinksDisk)`

SetDisk sets Disk field to given value.

### HasDisk

`func (o *DatabaseLinks) HasDisk() bool`

HasDisk returns a boolean if a field has been set.

### GetInitializeFrom

`func (o *DatabaseLinks) GetInitializeFrom() DatabaseLinksInitializeFrom`

GetInitializeFrom returns the InitializeFrom field if non-nil, zero value otherwise.

### GetInitializeFromOk

`func (o *DatabaseLinks) GetInitializeFromOk() (*DatabaseLinksInitializeFrom, bool)`

GetInitializeFromOk returns a tuple with the InitializeFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitializeFrom

`func (o *DatabaseLinks) SetInitializeFrom(v DatabaseLinksInitializeFrom)`

SetInitializeFrom sets InitializeFrom field to given value.

### HasInitializeFrom

`func (o *DatabaseLinks) HasInitializeFrom() bool`

HasInitializeFrom returns a boolean if a field has been set.

### GetCurrentConfiguration

`func (o *DatabaseLinks) GetCurrentConfiguration() DatabaseLinksCurrentConfiguration`

GetCurrentConfiguration returns the CurrentConfiguration field if non-nil, zero value otherwise.

### GetCurrentConfigurationOk

`func (o *DatabaseLinks) GetCurrentConfigurationOk() (*DatabaseLinksCurrentConfiguration, bool)`

GetCurrentConfigurationOk returns a tuple with the CurrentConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentConfiguration

`func (o *DatabaseLinks) SetCurrentConfiguration(v DatabaseLinksCurrentConfiguration)`

SetCurrentConfiguration sets CurrentConfiguration field to given value.

### HasCurrentConfiguration

`func (o *DatabaseLinks) HasCurrentConfiguration() bool`

HasCurrentConfiguration returns a boolean if a field has been set.

### GetDependents

`func (o *DatabaseLinks) GetDependents() DatabaseLinksDependents`

GetDependents returns the Dependents field if non-nil, zero value otherwise.

### GetDependentsOk

`func (o *DatabaseLinks) GetDependentsOk() (*DatabaseLinksDependents, bool)`

GetDependentsOk returns a tuple with the Dependents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependents

`func (o *DatabaseLinks) SetDependents(v DatabaseLinksDependents)`

SetDependents sets Dependents field to given value.

### HasDependents

`func (o *DatabaseLinks) HasDependents() bool`

HasDependents returns a boolean if a field has been set.

### GetOperations

`func (o *DatabaseLinks) GetOperations() DatabaseLinksOperations`

GetOperations returns the Operations field if non-nil, zero value otherwise.

### GetOperationsOk

`func (o *DatabaseLinks) GetOperationsOk() (*DatabaseLinksOperations, bool)`

GetOperationsOk returns a tuple with the Operations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperations

`func (o *DatabaseLinks) SetOperations(v DatabaseLinksOperations)`

SetOperations sets Operations field to given value.

### HasOperations

`func (o *DatabaseLinks) HasOperations() bool`

HasOperations returns a boolean if a field has been set.

### GetBackups

`func (o *DatabaseLinks) GetBackups() DatabaseLinksBackups`

GetBackups returns the Backups field if non-nil, zero value otherwise.

### GetBackupsOk

`func (o *DatabaseLinks) GetBackupsOk() (*DatabaseLinksBackups, bool)`

GetBackupsOk returns a tuple with the Backups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackups

`func (o *DatabaseLinks) SetBackups(v DatabaseLinksBackups)`

SetBackups sets Backups field to given value.

### HasBackups

`func (o *DatabaseLinks) HasBackups() bool`

HasBackups returns a boolean if a field has been set.

### GetConfigurations

`func (o *DatabaseLinks) GetConfigurations() DatabaseLinksConfigurations`

GetConfigurations returns the Configurations field if non-nil, zero value otherwise.

### GetConfigurationsOk

`func (o *DatabaseLinks) GetConfigurationsOk() (*DatabaseLinksConfigurations, bool)`

GetConfigurationsOk returns a tuple with the Configurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurations

`func (o *DatabaseLinks) SetConfigurations(v DatabaseLinksConfigurations)`

SetConfigurations sets Configurations field to given value.

### HasConfigurations

`func (o *DatabaseLinks) HasConfigurations() bool`

HasConfigurations returns a boolean if a field has been set.

### GetDatabaseCredentials

`func (o *DatabaseLinks) GetDatabaseCredentials() DatabaseLinksDatabaseCredentials`

GetDatabaseCredentials returns the DatabaseCredentials field if non-nil, zero value otherwise.

### GetDatabaseCredentialsOk

`func (o *DatabaseLinks) GetDatabaseCredentialsOk() (*DatabaseLinksDatabaseCredentials, bool)`

GetDatabaseCredentialsOk returns a tuple with the DatabaseCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseCredentials

`func (o *DatabaseLinks) SetDatabaseCredentials(v DatabaseLinksDatabaseCredentials)`

SetDatabaseCredentials sets DatabaseCredentials field to given value.

### HasDatabaseCredentials

`func (o *DatabaseLinks) HasDatabaseCredentials() bool`

HasDatabaseCredentials returns a boolean if a field has been set.

### GetSelf

`func (o *DatabaseLinks) GetSelf() DatabaseLinksSelf`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *DatabaseLinks) GetSelfOk() (*DatabaseLinksSelf, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *DatabaseLinks) SetSelf(v DatabaseLinksSelf)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *DatabaseLinks) HasSelf() bool`

HasSelf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


