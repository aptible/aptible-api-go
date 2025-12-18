# AppExternalAwsRdsConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**MetaType** | **string** |  | 
**DatabaseName** | **string** |  | 
**DatabaseUser** | **string** |  | 
**Superuser** | **bool** |  | 
**CreatedAt** | **string** |  | 
**UpdatedAt** | **string** |  | 
**Links** | Pointer to [**AppExternalAwsRdsConnectionLinks**](AppExternalAwsRdsConnectionLinks.md) |  | [optional] 

## Methods

### NewAppExternalAwsRdsConnection

`func NewAppExternalAwsRdsConnection(id int32, metaType string, databaseName string, databaseUser string, superuser bool, createdAt string, updatedAt string, ) *AppExternalAwsRdsConnection`

NewAppExternalAwsRdsConnection instantiates a new AppExternalAwsRdsConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppExternalAwsRdsConnectionWithDefaults

`func NewAppExternalAwsRdsConnectionWithDefaults() *AppExternalAwsRdsConnection`

NewAppExternalAwsRdsConnectionWithDefaults instantiates a new AppExternalAwsRdsConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AppExternalAwsRdsConnection) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AppExternalAwsRdsConnection) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AppExternalAwsRdsConnection) SetId(v int32)`

SetId sets Id field to given value.


### GetMetaType

`func (o *AppExternalAwsRdsConnection) GetMetaType() string`

GetMetaType returns the MetaType field if non-nil, zero value otherwise.

### GetMetaTypeOk

`func (o *AppExternalAwsRdsConnection) GetMetaTypeOk() (*string, bool)`

GetMetaTypeOk returns a tuple with the MetaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaType

`func (o *AppExternalAwsRdsConnection) SetMetaType(v string)`

SetMetaType sets MetaType field to given value.


### GetDatabaseName

`func (o *AppExternalAwsRdsConnection) GetDatabaseName() string`

GetDatabaseName returns the DatabaseName field if non-nil, zero value otherwise.

### GetDatabaseNameOk

`func (o *AppExternalAwsRdsConnection) GetDatabaseNameOk() (*string, bool)`

GetDatabaseNameOk returns a tuple with the DatabaseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseName

`func (o *AppExternalAwsRdsConnection) SetDatabaseName(v string)`

SetDatabaseName sets DatabaseName field to given value.


### GetDatabaseUser

`func (o *AppExternalAwsRdsConnection) GetDatabaseUser() string`

GetDatabaseUser returns the DatabaseUser field if non-nil, zero value otherwise.

### GetDatabaseUserOk

`func (o *AppExternalAwsRdsConnection) GetDatabaseUserOk() (*string, bool)`

GetDatabaseUserOk returns a tuple with the DatabaseUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseUser

`func (o *AppExternalAwsRdsConnection) SetDatabaseUser(v string)`

SetDatabaseUser sets DatabaseUser field to given value.


### GetSuperuser

`func (o *AppExternalAwsRdsConnection) GetSuperuser() bool`

GetSuperuser returns the Superuser field if non-nil, zero value otherwise.

### GetSuperuserOk

`func (o *AppExternalAwsRdsConnection) GetSuperuserOk() (*bool, bool)`

GetSuperuserOk returns a tuple with the Superuser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuperuser

`func (o *AppExternalAwsRdsConnection) SetSuperuser(v bool)`

SetSuperuser sets Superuser field to given value.


### GetCreatedAt

`func (o *AppExternalAwsRdsConnection) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AppExternalAwsRdsConnection) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AppExternalAwsRdsConnection) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *AppExternalAwsRdsConnection) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AppExternalAwsRdsConnection) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AppExternalAwsRdsConnection) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetLinks

`func (o *AppExternalAwsRdsConnection) GetLinks() AppExternalAwsRdsConnectionLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AppExternalAwsRdsConnection) GetLinksOk() (*AppExternalAwsRdsConnectionLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AppExternalAwsRdsConnection) SetLinks(v AppExternalAwsRdsConnectionLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AppExternalAwsRdsConnection) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


