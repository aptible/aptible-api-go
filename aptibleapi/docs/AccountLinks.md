# AccountLinks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Organization** | Pointer to [**ListAccountsForStack200ResponseLinksStack**](ListAccountsForStack200ResponseLinksStack.md) |  | [optional] 
**Stack** | Pointer to [**ListAccountsForStack200ResponseLinksStack**](ListAccountsForStack200ResponseLinksStack.md) |  | [optional] 
**Apps** | Pointer to [**ListAccountsForStack200ResponseLinksStack**](ListAccountsForStack200ResponseLinksStack.md) |  | [optional] 
**Backups** | Pointer to [**ListAccountsForStack200ResponseLinksStack**](ListAccountsForStack200ResponseLinksStack.md) |  | [optional] 
**Databases** | Pointer to [**ListAccountsForStack200ResponseLinksStack**](ListAccountsForStack200ResponseLinksStack.md) |  | [optional] 
**Disks** | Pointer to [**ListAccountsForStack200ResponseLinksStack**](ListAccountsForStack200ResponseLinksStack.md) |  | [optional] 
**Services** | Pointer to [**ListAccountsForStack200ResponseLinksStack**](ListAccountsForStack200ResponseLinksStack.md) |  | [optional] 
**Operations** | Pointer to [**ListAccountsForStack200ResponseLinksStack**](ListAccountsForStack200ResponseLinksStack.md) |  | [optional] 
**Permissions** | Pointer to [**ListAccountsForStack200ResponseLinksStack**](ListAccountsForStack200ResponseLinksStack.md) |  | [optional] 
**LogDrains** | Pointer to [**ListAccountsForStack200ResponseLinksStack**](ListAccountsForStack200ResponseLinksStack.md) |  | [optional] 
**MetricDrains** | Pointer to [**ListAccountsForStack200ResponseLinksStack**](ListAccountsForStack200ResponseLinksStack.md) |  | [optional] 
**Certificates** | Pointer to [**ListAccountsForStack200ResponseLinksStack**](ListAccountsForStack200ResponseLinksStack.md) |  | [optional] 
**Vhosts** | Pointer to [**ListAccountsForStack200ResponseLinksStack**](ListAccountsForStack200ResponseLinksStack.md) |  | [optional] 
**ActivityReports** | Pointer to [**ListAccountsForStack200ResponseLinksStack**](ListAccountsForStack200ResponseLinksStack.md) |  | [optional] 
**DiskAttachments** | Pointer to [**ListAccountsForStack200ResponseLinksStack**](ListAccountsForStack200ResponseLinksStack.md) |  | [optional] 
**Self** | Pointer to [**ListAccountsForStack200ResponseLinksStack**](ListAccountsForStack200ResponseLinksStack.md) |  | [optional] 

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

`func (o *AccountLinks) GetOrganization() ListAccountsForStack200ResponseLinksStack`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *AccountLinks) GetOrganizationOk() (*ListAccountsForStack200ResponseLinksStack, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *AccountLinks) SetOrganization(v ListAccountsForStack200ResponseLinksStack)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *AccountLinks) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetStack

`func (o *AccountLinks) GetStack() ListAccountsForStack200ResponseLinksStack`

GetStack returns the Stack field if non-nil, zero value otherwise.

### GetStackOk

`func (o *AccountLinks) GetStackOk() (*ListAccountsForStack200ResponseLinksStack, bool)`

GetStackOk returns a tuple with the Stack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStack

`func (o *AccountLinks) SetStack(v ListAccountsForStack200ResponseLinksStack)`

SetStack sets Stack field to given value.

### HasStack

`func (o *AccountLinks) HasStack() bool`

HasStack returns a boolean if a field has been set.

### GetApps

`func (o *AccountLinks) GetApps() ListAccountsForStack200ResponseLinksStack`

GetApps returns the Apps field if non-nil, zero value otherwise.

### GetAppsOk

`func (o *AccountLinks) GetAppsOk() (*ListAccountsForStack200ResponseLinksStack, bool)`

GetAppsOk returns a tuple with the Apps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApps

`func (o *AccountLinks) SetApps(v ListAccountsForStack200ResponseLinksStack)`

SetApps sets Apps field to given value.

### HasApps

`func (o *AccountLinks) HasApps() bool`

HasApps returns a boolean if a field has been set.

### GetBackups

`func (o *AccountLinks) GetBackups() ListAccountsForStack200ResponseLinksStack`

GetBackups returns the Backups field if non-nil, zero value otherwise.

### GetBackupsOk

`func (o *AccountLinks) GetBackupsOk() (*ListAccountsForStack200ResponseLinksStack, bool)`

GetBackupsOk returns a tuple with the Backups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackups

`func (o *AccountLinks) SetBackups(v ListAccountsForStack200ResponseLinksStack)`

SetBackups sets Backups field to given value.

### HasBackups

`func (o *AccountLinks) HasBackups() bool`

HasBackups returns a boolean if a field has been set.

### GetDatabases

`func (o *AccountLinks) GetDatabases() ListAccountsForStack200ResponseLinksStack`

GetDatabases returns the Databases field if non-nil, zero value otherwise.

### GetDatabasesOk

`func (o *AccountLinks) GetDatabasesOk() (*ListAccountsForStack200ResponseLinksStack, bool)`

GetDatabasesOk returns a tuple with the Databases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabases

`func (o *AccountLinks) SetDatabases(v ListAccountsForStack200ResponseLinksStack)`

SetDatabases sets Databases field to given value.

### HasDatabases

`func (o *AccountLinks) HasDatabases() bool`

HasDatabases returns a boolean if a field has been set.

### GetDisks

`func (o *AccountLinks) GetDisks() ListAccountsForStack200ResponseLinksStack`

GetDisks returns the Disks field if non-nil, zero value otherwise.

### GetDisksOk

`func (o *AccountLinks) GetDisksOk() (*ListAccountsForStack200ResponseLinksStack, bool)`

GetDisksOk returns a tuple with the Disks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisks

`func (o *AccountLinks) SetDisks(v ListAccountsForStack200ResponseLinksStack)`

SetDisks sets Disks field to given value.

### HasDisks

`func (o *AccountLinks) HasDisks() bool`

HasDisks returns a boolean if a field has been set.

### GetServices

`func (o *AccountLinks) GetServices() ListAccountsForStack200ResponseLinksStack`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *AccountLinks) GetServicesOk() (*ListAccountsForStack200ResponseLinksStack, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *AccountLinks) SetServices(v ListAccountsForStack200ResponseLinksStack)`

SetServices sets Services field to given value.

### HasServices

`func (o *AccountLinks) HasServices() bool`

HasServices returns a boolean if a field has been set.

### GetOperations

`func (o *AccountLinks) GetOperations() ListAccountsForStack200ResponseLinksStack`

GetOperations returns the Operations field if non-nil, zero value otherwise.

### GetOperationsOk

`func (o *AccountLinks) GetOperationsOk() (*ListAccountsForStack200ResponseLinksStack, bool)`

GetOperationsOk returns a tuple with the Operations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperations

`func (o *AccountLinks) SetOperations(v ListAccountsForStack200ResponseLinksStack)`

SetOperations sets Operations field to given value.

### HasOperations

`func (o *AccountLinks) HasOperations() bool`

HasOperations returns a boolean if a field has been set.

### GetPermissions

`func (o *AccountLinks) GetPermissions() ListAccountsForStack200ResponseLinksStack`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *AccountLinks) GetPermissionsOk() (*ListAccountsForStack200ResponseLinksStack, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *AccountLinks) SetPermissions(v ListAccountsForStack200ResponseLinksStack)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *AccountLinks) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetLogDrains

`func (o *AccountLinks) GetLogDrains() ListAccountsForStack200ResponseLinksStack`

GetLogDrains returns the LogDrains field if non-nil, zero value otherwise.

### GetLogDrainsOk

`func (o *AccountLinks) GetLogDrainsOk() (*ListAccountsForStack200ResponseLinksStack, bool)`

GetLogDrainsOk returns a tuple with the LogDrains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogDrains

`func (o *AccountLinks) SetLogDrains(v ListAccountsForStack200ResponseLinksStack)`

SetLogDrains sets LogDrains field to given value.

### HasLogDrains

`func (o *AccountLinks) HasLogDrains() bool`

HasLogDrains returns a boolean if a field has been set.

### GetMetricDrains

`func (o *AccountLinks) GetMetricDrains() ListAccountsForStack200ResponseLinksStack`

GetMetricDrains returns the MetricDrains field if non-nil, zero value otherwise.

### GetMetricDrainsOk

`func (o *AccountLinks) GetMetricDrainsOk() (*ListAccountsForStack200ResponseLinksStack, bool)`

GetMetricDrainsOk returns a tuple with the MetricDrains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricDrains

`func (o *AccountLinks) SetMetricDrains(v ListAccountsForStack200ResponseLinksStack)`

SetMetricDrains sets MetricDrains field to given value.

### HasMetricDrains

`func (o *AccountLinks) HasMetricDrains() bool`

HasMetricDrains returns a boolean if a field has been set.

### GetCertificates

`func (o *AccountLinks) GetCertificates() ListAccountsForStack200ResponseLinksStack`

GetCertificates returns the Certificates field if non-nil, zero value otherwise.

### GetCertificatesOk

`func (o *AccountLinks) GetCertificatesOk() (*ListAccountsForStack200ResponseLinksStack, bool)`

GetCertificatesOk returns a tuple with the Certificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificates

`func (o *AccountLinks) SetCertificates(v ListAccountsForStack200ResponseLinksStack)`

SetCertificates sets Certificates field to given value.

### HasCertificates

`func (o *AccountLinks) HasCertificates() bool`

HasCertificates returns a boolean if a field has been set.

### GetVhosts

`func (o *AccountLinks) GetVhosts() ListAccountsForStack200ResponseLinksStack`

GetVhosts returns the Vhosts field if non-nil, zero value otherwise.

### GetVhostsOk

`func (o *AccountLinks) GetVhostsOk() (*ListAccountsForStack200ResponseLinksStack, bool)`

GetVhostsOk returns a tuple with the Vhosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVhosts

`func (o *AccountLinks) SetVhosts(v ListAccountsForStack200ResponseLinksStack)`

SetVhosts sets Vhosts field to given value.

### HasVhosts

`func (o *AccountLinks) HasVhosts() bool`

HasVhosts returns a boolean if a field has been set.

### GetActivityReports

`func (o *AccountLinks) GetActivityReports() ListAccountsForStack200ResponseLinksStack`

GetActivityReports returns the ActivityReports field if non-nil, zero value otherwise.

### GetActivityReportsOk

`func (o *AccountLinks) GetActivityReportsOk() (*ListAccountsForStack200ResponseLinksStack, bool)`

GetActivityReportsOk returns a tuple with the ActivityReports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityReports

`func (o *AccountLinks) SetActivityReports(v ListAccountsForStack200ResponseLinksStack)`

SetActivityReports sets ActivityReports field to given value.

### HasActivityReports

`func (o *AccountLinks) HasActivityReports() bool`

HasActivityReports returns a boolean if a field has been set.

### GetDiskAttachments

`func (o *AccountLinks) GetDiskAttachments() ListAccountsForStack200ResponseLinksStack`

GetDiskAttachments returns the DiskAttachments field if non-nil, zero value otherwise.

### GetDiskAttachmentsOk

`func (o *AccountLinks) GetDiskAttachmentsOk() (*ListAccountsForStack200ResponseLinksStack, bool)`

GetDiskAttachmentsOk returns a tuple with the DiskAttachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskAttachments

`func (o *AccountLinks) SetDiskAttachments(v ListAccountsForStack200ResponseLinksStack)`

SetDiskAttachments sets DiskAttachments field to given value.

### HasDiskAttachments

`func (o *AccountLinks) HasDiskAttachments() bool`

HasDiskAttachments returns a boolean if a field has been set.

### GetSelf

`func (o *AccountLinks) GetSelf() ListAccountsForStack200ResponseLinksStack`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *AccountLinks) GetSelfOk() (*ListAccountsForStack200ResponseLinksStack, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *AccountLinks) SetSelf(v ListAccountsForStack200ResponseLinksStack)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *AccountLinks) HasSelf() bool`

HasSelf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


