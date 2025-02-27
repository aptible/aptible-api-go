# AccountLinks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Organization** | Pointer to [**AccountLinksOrganization**](AccountLinksOrganization.md) |  | [optional] 
**Stack** | Pointer to [**AccountLinksStack**](AccountLinksStack.md) |  | [optional] 
**Apps** | Pointer to [**AccountLinksApps**](AccountLinksApps.md) |  | [optional] 
**Backups** | Pointer to [**AccountLinksBackups**](AccountLinksBackups.md) |  | [optional] 
**Databases** | Pointer to [**AccountLinksDatabases**](AccountLinksDatabases.md) |  | [optional] 
**Disks** | Pointer to [**AccountLinksDisks**](AccountLinksDisks.md) |  | [optional] 
**Services** | Pointer to [**AccountLinksServices**](AccountLinksServices.md) |  | [optional] 
**Operations** | Pointer to [**AccountLinksOperations**](AccountLinksOperations.md) |  | [optional] 
**Permissions** | Pointer to [**AccountLinksPermissions**](AccountLinksPermissions.md) |  | [optional] 
**LogDrains** | Pointer to [**AccountLinksLogDrains**](AccountLinksLogDrains.md) |  | [optional] 
**MetricDrains** | Pointer to [**AccountLinksMetricDrains**](AccountLinksMetricDrains.md) |  | [optional] 
**Certificates** | Pointer to [**AccountLinksCertificates**](AccountLinksCertificates.md) |  | [optional] 
**Vhosts** | Pointer to [**AccountLinksVhosts**](AccountLinksVhosts.md) |  | [optional] 
**ActivityReports** | Pointer to [**AccountLinksActivityReports**](AccountLinksActivityReports.md) |  | [optional] 
**DiskAttachments** | Pointer to [**AccountLinksDiskAttachments**](AccountLinksDiskAttachments.md) |  | [optional] 
**Self** | Pointer to [**AccountLinksSelf**](AccountLinksSelf.md) |  | [optional] 

## Methods

### NewAccountLinks

`func NewAccountLinks() *AccountLinks`

NewAccountLinks instantiates a new AccountLinks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountLinksWithDefaults

`func NewAccountLinksWithDefaults() *AccountLinks`

NewAccountLinksWithDefaults instantiates a new AccountLinks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrganization

`func (o *AccountLinks) GetOrganization() AccountLinksOrganization`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *AccountLinks) GetOrganizationOk() (*AccountLinksOrganization, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *AccountLinks) SetOrganization(v AccountLinksOrganization)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *AccountLinks) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetStack

`func (o *AccountLinks) GetStack() AccountLinksStack`

GetStack returns the Stack field if non-nil, zero value otherwise.

### GetStackOk

`func (o *AccountLinks) GetStackOk() (*AccountLinksStack, bool)`

GetStackOk returns a tuple with the Stack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStack

`func (o *AccountLinks) SetStack(v AccountLinksStack)`

SetStack sets Stack field to given value.

### HasStack

`func (o *AccountLinks) HasStack() bool`

HasStack returns a boolean if a field has been set.

### GetApps

`func (o *AccountLinks) GetApps() AccountLinksApps`

GetApps returns the Apps field if non-nil, zero value otherwise.

### GetAppsOk

`func (o *AccountLinks) GetAppsOk() (*AccountLinksApps, bool)`

GetAppsOk returns a tuple with the Apps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApps

`func (o *AccountLinks) SetApps(v AccountLinksApps)`

SetApps sets Apps field to given value.

### HasApps

`func (o *AccountLinks) HasApps() bool`

HasApps returns a boolean if a field has been set.

### GetBackups

`func (o *AccountLinks) GetBackups() AccountLinksBackups`

GetBackups returns the Backups field if non-nil, zero value otherwise.

### GetBackupsOk

`func (o *AccountLinks) GetBackupsOk() (*AccountLinksBackups, bool)`

GetBackupsOk returns a tuple with the Backups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackups

`func (o *AccountLinks) SetBackups(v AccountLinksBackups)`

SetBackups sets Backups field to given value.

### HasBackups

`func (o *AccountLinks) HasBackups() bool`

HasBackups returns a boolean if a field has been set.

### GetDatabases

`func (o *AccountLinks) GetDatabases() AccountLinksDatabases`

GetDatabases returns the Databases field if non-nil, zero value otherwise.

### GetDatabasesOk

`func (o *AccountLinks) GetDatabasesOk() (*AccountLinksDatabases, bool)`

GetDatabasesOk returns a tuple with the Databases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabases

`func (o *AccountLinks) SetDatabases(v AccountLinksDatabases)`

SetDatabases sets Databases field to given value.

### HasDatabases

`func (o *AccountLinks) HasDatabases() bool`

HasDatabases returns a boolean if a field has been set.

### GetDisks

`func (o *AccountLinks) GetDisks() AccountLinksDisks`

GetDisks returns the Disks field if non-nil, zero value otherwise.

### GetDisksOk

`func (o *AccountLinks) GetDisksOk() (*AccountLinksDisks, bool)`

GetDisksOk returns a tuple with the Disks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisks

`func (o *AccountLinks) SetDisks(v AccountLinksDisks)`

SetDisks sets Disks field to given value.

### HasDisks

`func (o *AccountLinks) HasDisks() bool`

HasDisks returns a boolean if a field has been set.

### GetServices

`func (o *AccountLinks) GetServices() AccountLinksServices`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *AccountLinks) GetServicesOk() (*AccountLinksServices, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *AccountLinks) SetServices(v AccountLinksServices)`

SetServices sets Services field to given value.

### HasServices

`func (o *AccountLinks) HasServices() bool`

HasServices returns a boolean if a field has been set.

### GetOperations

`func (o *AccountLinks) GetOperations() AccountLinksOperations`

GetOperations returns the Operations field if non-nil, zero value otherwise.

### GetOperationsOk

`func (o *AccountLinks) GetOperationsOk() (*AccountLinksOperations, bool)`

GetOperationsOk returns a tuple with the Operations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperations

`func (o *AccountLinks) SetOperations(v AccountLinksOperations)`

SetOperations sets Operations field to given value.

### HasOperations

`func (o *AccountLinks) HasOperations() bool`

HasOperations returns a boolean if a field has been set.

### GetPermissions

`func (o *AccountLinks) GetPermissions() AccountLinksPermissions`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *AccountLinks) GetPermissionsOk() (*AccountLinksPermissions, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *AccountLinks) SetPermissions(v AccountLinksPermissions)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *AccountLinks) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetLogDrains

`func (o *AccountLinks) GetLogDrains() AccountLinksLogDrains`

GetLogDrains returns the LogDrains field if non-nil, zero value otherwise.

### GetLogDrainsOk

`func (o *AccountLinks) GetLogDrainsOk() (*AccountLinksLogDrains, bool)`

GetLogDrainsOk returns a tuple with the LogDrains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogDrains

`func (o *AccountLinks) SetLogDrains(v AccountLinksLogDrains)`

SetLogDrains sets LogDrains field to given value.

### HasLogDrains

`func (o *AccountLinks) HasLogDrains() bool`

HasLogDrains returns a boolean if a field has been set.

### GetMetricDrains

`func (o *AccountLinks) GetMetricDrains() AccountLinksMetricDrains`

GetMetricDrains returns the MetricDrains field if non-nil, zero value otherwise.

### GetMetricDrainsOk

`func (o *AccountLinks) GetMetricDrainsOk() (*AccountLinksMetricDrains, bool)`

GetMetricDrainsOk returns a tuple with the MetricDrains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricDrains

`func (o *AccountLinks) SetMetricDrains(v AccountLinksMetricDrains)`

SetMetricDrains sets MetricDrains field to given value.

### HasMetricDrains

`func (o *AccountLinks) HasMetricDrains() bool`

HasMetricDrains returns a boolean if a field has been set.

### GetCertificates

`func (o *AccountLinks) GetCertificates() AccountLinksCertificates`

GetCertificates returns the Certificates field if non-nil, zero value otherwise.

### GetCertificatesOk

`func (o *AccountLinks) GetCertificatesOk() (*AccountLinksCertificates, bool)`

GetCertificatesOk returns a tuple with the Certificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificates

`func (o *AccountLinks) SetCertificates(v AccountLinksCertificates)`

SetCertificates sets Certificates field to given value.

### HasCertificates

`func (o *AccountLinks) HasCertificates() bool`

HasCertificates returns a boolean if a field has been set.

### GetVhosts

`func (o *AccountLinks) GetVhosts() AccountLinksVhosts`

GetVhosts returns the Vhosts field if non-nil, zero value otherwise.

### GetVhostsOk

`func (o *AccountLinks) GetVhostsOk() (*AccountLinksVhosts, bool)`

GetVhostsOk returns a tuple with the Vhosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVhosts

`func (o *AccountLinks) SetVhosts(v AccountLinksVhosts)`

SetVhosts sets Vhosts field to given value.

### HasVhosts

`func (o *AccountLinks) HasVhosts() bool`

HasVhosts returns a boolean if a field has been set.

### GetActivityReports

`func (o *AccountLinks) GetActivityReports() AccountLinksActivityReports`

GetActivityReports returns the ActivityReports field if non-nil, zero value otherwise.

### GetActivityReportsOk

`func (o *AccountLinks) GetActivityReportsOk() (*AccountLinksActivityReports, bool)`

GetActivityReportsOk returns a tuple with the ActivityReports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityReports

`func (o *AccountLinks) SetActivityReports(v AccountLinksActivityReports)`

SetActivityReports sets ActivityReports field to given value.

### HasActivityReports

`func (o *AccountLinks) HasActivityReports() bool`

HasActivityReports returns a boolean if a field has been set.

### GetDiskAttachments

`func (o *AccountLinks) GetDiskAttachments() AccountLinksDiskAttachments`

GetDiskAttachments returns the DiskAttachments field if non-nil, zero value otherwise.

### GetDiskAttachmentsOk

`func (o *AccountLinks) GetDiskAttachmentsOk() (*AccountLinksDiskAttachments, bool)`

GetDiskAttachmentsOk returns a tuple with the DiskAttachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskAttachments

`func (o *AccountLinks) SetDiskAttachments(v AccountLinksDiskAttachments)`

SetDiskAttachments sets DiskAttachments field to given value.

### HasDiskAttachments

`func (o *AccountLinks) HasDiskAttachments() bool`

HasDiskAttachments returns a boolean if a field has been set.

### GetSelf

`func (o *AccountLinks) GetSelf() AccountLinksSelf`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *AccountLinks) GetSelfOk() (*AccountLinksSelf, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *AccountLinks) SetSelf(v AccountLinksSelf)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *AccountLinks) HasSelf() bool`

HasSelf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


