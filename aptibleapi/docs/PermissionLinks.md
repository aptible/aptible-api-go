# PermissionLinks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Role** | Pointer to [**PermissionLinksRole**](PermissionLinksRole.md) |  | [optional] 
**Account** | Pointer to [**PermissionLinksAccount**](PermissionLinksAccount.md) |  | [optional] 
**Self** | Pointer to [**PermissionLinksSelf**](PermissionLinksSelf.md) |  | [optional] 

## Methods

### NewPermissionLinks

`func NewPermissionLinks() *PermissionLinks`

NewPermissionLinks instantiates a new PermissionLinks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPermissionLinksWithDefaults

`func NewPermissionLinksWithDefaults() *PermissionLinks`

NewPermissionLinksWithDefaults instantiates a new PermissionLinks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRole

`func (o *PermissionLinks) GetRole() PermissionLinksRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *PermissionLinks) GetRoleOk() (*PermissionLinksRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *PermissionLinks) SetRole(v PermissionLinksRole)`

SetRole sets Role field to given value.

### HasRole

`func (o *PermissionLinks) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetAccount

`func (o *PermissionLinks) GetAccount() PermissionLinksAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *PermissionLinks) GetAccountOk() (*PermissionLinksAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *PermissionLinks) SetAccount(v PermissionLinksAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *PermissionLinks) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetSelf

`func (o *PermissionLinks) GetSelf() PermissionLinksSelf`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *PermissionLinks) GetSelfOk() (*PermissionLinksSelf, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *PermissionLinks) SetSelf(v PermissionLinksSelf)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *PermissionLinks) HasSelf() bool`

HasSelf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


