# OperationLinks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | Pointer to [**OperationLinksUser**](OperationLinksUser.md) |  | [optional] 
**Resource** | Pointer to [**OperationLinksResource**](OperationLinksResource.md) |  | [optional] 
**Account** | Pointer to [**OperationLinksAccount**](OperationLinksAccount.md) |  | [optional] 
**DestinationAccount** | Pointer to [**OperationLinksDestinationAccount**](OperationLinksDestinationAccount.md) |  | [optional] 
**SshPortalConnections** | Pointer to [**OperationLinksSshPortalConnections**](OperationLinksSshPortalConnections.md) |  | [optional] 
**EphemeralSessions** | Pointer to [**OperationLinksEphemeralSessions**](OperationLinksEphemeralSessions.md) |  | [optional] 
**CodeScanResult** | Pointer to [**OperationLinksCodeScanResult**](OperationLinksCodeScanResult.md) |  | [optional] 
**Service** | Pointer to [**OperationLinksService**](OperationLinksService.md) |  | [optional] 
**Self** | Pointer to [**OperationLinksSelf**](OperationLinksSelf.md) |  | [optional] 

## Methods

### NewOperationLinks

`func NewOperationLinks() *OperationLinks`

NewOperationLinks instantiates a new OperationLinks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOperationLinksWithDefaults

`func NewOperationLinksWithDefaults() *OperationLinks`

NewOperationLinksWithDefaults instantiates a new OperationLinks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUser

`func (o *OperationLinks) GetUser() OperationLinksUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *OperationLinks) GetUserOk() (*OperationLinksUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *OperationLinks) SetUser(v OperationLinksUser)`

SetUser sets User field to given value.

### HasUser

`func (o *OperationLinks) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetResource

`func (o *OperationLinks) GetResource() OperationLinksResource`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *OperationLinks) GetResourceOk() (*OperationLinksResource, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *OperationLinks) SetResource(v OperationLinksResource)`

SetResource sets Resource field to given value.

### HasResource

`func (o *OperationLinks) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetAccount

`func (o *OperationLinks) GetAccount() OperationLinksAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *OperationLinks) GetAccountOk() (*OperationLinksAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *OperationLinks) SetAccount(v OperationLinksAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *OperationLinks) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetDestinationAccount

`func (o *OperationLinks) GetDestinationAccount() OperationLinksDestinationAccount`

GetDestinationAccount returns the DestinationAccount field if non-nil, zero value otherwise.

### GetDestinationAccountOk

`func (o *OperationLinks) GetDestinationAccountOk() (*OperationLinksDestinationAccount, bool)`

GetDestinationAccountOk returns a tuple with the DestinationAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationAccount

`func (o *OperationLinks) SetDestinationAccount(v OperationLinksDestinationAccount)`

SetDestinationAccount sets DestinationAccount field to given value.

### HasDestinationAccount

`func (o *OperationLinks) HasDestinationAccount() bool`

HasDestinationAccount returns a boolean if a field has been set.

### GetSshPortalConnections

`func (o *OperationLinks) GetSshPortalConnections() OperationLinksSshPortalConnections`

GetSshPortalConnections returns the SshPortalConnections field if non-nil, zero value otherwise.

### GetSshPortalConnectionsOk

`func (o *OperationLinks) GetSshPortalConnectionsOk() (*OperationLinksSshPortalConnections, bool)`

GetSshPortalConnectionsOk returns a tuple with the SshPortalConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPortalConnections

`func (o *OperationLinks) SetSshPortalConnections(v OperationLinksSshPortalConnections)`

SetSshPortalConnections sets SshPortalConnections field to given value.

### HasSshPortalConnections

`func (o *OperationLinks) HasSshPortalConnections() bool`

HasSshPortalConnections returns a boolean if a field has been set.

### GetEphemeralSessions

`func (o *OperationLinks) GetEphemeralSessions() OperationLinksEphemeralSessions`

GetEphemeralSessions returns the EphemeralSessions field if non-nil, zero value otherwise.

### GetEphemeralSessionsOk

`func (o *OperationLinks) GetEphemeralSessionsOk() (*OperationLinksEphemeralSessions, bool)`

GetEphemeralSessionsOk returns a tuple with the EphemeralSessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEphemeralSessions

`func (o *OperationLinks) SetEphemeralSessions(v OperationLinksEphemeralSessions)`

SetEphemeralSessions sets EphemeralSessions field to given value.

### HasEphemeralSessions

`func (o *OperationLinks) HasEphemeralSessions() bool`

HasEphemeralSessions returns a boolean if a field has been set.

### GetCodeScanResult

`func (o *OperationLinks) GetCodeScanResult() OperationLinksCodeScanResult`

GetCodeScanResult returns the CodeScanResult field if non-nil, zero value otherwise.

### GetCodeScanResultOk

`func (o *OperationLinks) GetCodeScanResultOk() (*OperationLinksCodeScanResult, bool)`

GetCodeScanResultOk returns a tuple with the CodeScanResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeScanResult

`func (o *OperationLinks) SetCodeScanResult(v OperationLinksCodeScanResult)`

SetCodeScanResult sets CodeScanResult field to given value.

### HasCodeScanResult

`func (o *OperationLinks) HasCodeScanResult() bool`

HasCodeScanResult returns a boolean if a field has been set.

### GetService

`func (o *OperationLinks) GetService() OperationLinksService`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *OperationLinks) GetServiceOk() (*OperationLinksService, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *OperationLinks) SetService(v OperationLinksService)`

SetService sets Service field to given value.

### HasService

`func (o *OperationLinks) HasService() bool`

HasService returns a boolean if a field has been set.

### GetSelf

`func (o *OperationLinks) GetSelf() OperationLinksSelf`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *OperationLinks) GetSelfOk() (*OperationLinksSelf, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *OperationLinks) SetSelf(v OperationLinksSelf)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *OperationLinks) HasSelf() bool`

HasSelf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


