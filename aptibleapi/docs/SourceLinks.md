# SourceLinks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Organization** | Pointer to [**SourceLinksOrganization**](SourceLinksOrganization.md) |  | [optional] 
**Apps** | Pointer to [**SourceLinksApps**](SourceLinksApps.md) |  | [optional] 
**Deployments** | Pointer to [**SourceLinksDeployments**](SourceLinksDeployments.md) |  | [optional] 
**Self** | Pointer to [**SourceLinksSelf**](SourceLinksSelf.md) |  | [optional] 

## Methods

### NewSourceLinks

`func NewSourceLinks() *SourceLinks`

NewSourceLinks instantiates a new SourceLinks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceLinksWithDefaults

`func NewSourceLinksWithDefaults() *SourceLinks`

NewSourceLinksWithDefaults instantiates a new SourceLinks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrganization

`func (o *SourceLinks) GetOrganization() SourceLinksOrganization`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *SourceLinks) GetOrganizationOk() (*SourceLinksOrganization, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *SourceLinks) SetOrganization(v SourceLinksOrganization)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *SourceLinks) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetApps

`func (o *SourceLinks) GetApps() SourceLinksApps`

GetApps returns the Apps field if non-nil, zero value otherwise.

### GetAppsOk

`func (o *SourceLinks) GetAppsOk() (*SourceLinksApps, bool)`

GetAppsOk returns a tuple with the Apps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApps

`func (o *SourceLinks) SetApps(v SourceLinksApps)`

SetApps sets Apps field to given value.

### HasApps

`func (o *SourceLinks) HasApps() bool`

HasApps returns a boolean if a field has been set.

### GetDeployments

`func (o *SourceLinks) GetDeployments() SourceLinksDeployments`

GetDeployments returns the Deployments field if non-nil, zero value otherwise.

### GetDeploymentsOk

`func (o *SourceLinks) GetDeploymentsOk() (*SourceLinksDeployments, bool)`

GetDeploymentsOk returns a tuple with the Deployments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployments

`func (o *SourceLinks) SetDeployments(v SourceLinksDeployments)`

SetDeployments sets Deployments field to given value.

### HasDeployments

`func (o *SourceLinks) HasDeployments() bool`

HasDeployments returns a boolean if a field has been set.

### GetSelf

`func (o *SourceLinks) GetSelf() SourceLinksSelf`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *SourceLinks) GetSelfOk() (*SourceLinksSelf, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *SourceLinks) SetSelf(v SourceLinksSelf)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *SourceLinks) HasSelf() bool`

HasSelf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


