# SshPortalConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**MetaType** | **string** |  | 
**SshUser** | **string** |  | 
**SshCertificateBody** | **string** |  | 
**SshPortForwardSocket** | **NullableString** |  | 
**SshPty** | **bool** |  | 
**CreatedAt** | **string** |  | 
**UpdatedAt** | **string** |  | 
**Links** | Pointer to [**SshPortalConnectionLinks**](SshPortalConnectionLinks.md) |  | [optional] 

## Methods

### NewSshPortalConnection

`func NewSshPortalConnection(id string, metaType string, sshUser string, sshCertificateBody string, sshPortForwardSocket NullableString, sshPty bool, createdAt string, updatedAt string, ) *SshPortalConnection`

NewSshPortalConnection instantiates a new SshPortalConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSshPortalConnectionWithDefaults

`func NewSshPortalConnectionWithDefaults() *SshPortalConnection`

NewSshPortalConnectionWithDefaults instantiates a new SshPortalConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SshPortalConnection) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SshPortalConnection) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SshPortalConnection) SetId(v string)`

SetId sets Id field to given value.


### GetMetaType

`func (o *SshPortalConnection) GetMetaType() string`

GetMetaType returns the MetaType field if non-nil, zero value otherwise.

### GetMetaTypeOk

`func (o *SshPortalConnection) GetMetaTypeOk() (*string, bool)`

GetMetaTypeOk returns a tuple with the MetaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaType

`func (o *SshPortalConnection) SetMetaType(v string)`

SetMetaType sets MetaType field to given value.


### GetSshUser

`func (o *SshPortalConnection) GetSshUser() string`

GetSshUser returns the SshUser field if non-nil, zero value otherwise.

### GetSshUserOk

`func (o *SshPortalConnection) GetSshUserOk() (*string, bool)`

GetSshUserOk returns a tuple with the SshUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshUser

`func (o *SshPortalConnection) SetSshUser(v string)`

SetSshUser sets SshUser field to given value.


### GetSshCertificateBody

`func (o *SshPortalConnection) GetSshCertificateBody() string`

GetSshCertificateBody returns the SshCertificateBody field if non-nil, zero value otherwise.

### GetSshCertificateBodyOk

`func (o *SshPortalConnection) GetSshCertificateBodyOk() (*string, bool)`

GetSshCertificateBodyOk returns a tuple with the SshCertificateBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshCertificateBody

`func (o *SshPortalConnection) SetSshCertificateBody(v string)`

SetSshCertificateBody sets SshCertificateBody field to given value.


### GetSshPortForwardSocket

`func (o *SshPortalConnection) GetSshPortForwardSocket() string`

GetSshPortForwardSocket returns the SshPortForwardSocket field if non-nil, zero value otherwise.

### GetSshPortForwardSocketOk

`func (o *SshPortalConnection) GetSshPortForwardSocketOk() (*string, bool)`

GetSshPortForwardSocketOk returns a tuple with the SshPortForwardSocket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPortForwardSocket

`func (o *SshPortalConnection) SetSshPortForwardSocket(v string)`

SetSshPortForwardSocket sets SshPortForwardSocket field to given value.


### SetSshPortForwardSocketNil

`func (o *SshPortalConnection) SetSshPortForwardSocketNil(b bool)`

 SetSshPortForwardSocketNil sets the value for SshPortForwardSocket to be an explicit nil

### UnsetSshPortForwardSocket
`func (o *SshPortalConnection) UnsetSshPortForwardSocket()`

UnsetSshPortForwardSocket ensures that no value is present for SshPortForwardSocket, not even an explicit nil
### GetSshPty

`func (o *SshPortalConnection) GetSshPty() bool`

GetSshPty returns the SshPty field if non-nil, zero value otherwise.

### GetSshPtyOk

`func (o *SshPortalConnection) GetSshPtyOk() (*bool, bool)`

GetSshPtyOk returns a tuple with the SshPty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPty

`func (o *SshPortalConnection) SetSshPty(v bool)`

SetSshPty sets SshPty field to given value.


### GetCreatedAt

`func (o *SshPortalConnection) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SshPortalConnection) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SshPortalConnection) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *SshPortalConnection) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SshPortalConnection) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SshPortalConnection) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetLinks

`func (o *SshPortalConnection) GetLinks() SshPortalConnectionLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *SshPortalConnection) GetLinksOk() (*SshPortalConnectionLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *SshPortalConnection) SetLinks(v SshPortalConnectionLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *SshPortalConnection) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


