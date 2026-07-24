# DatabaseCredential

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**MetaType** | **string** |  | 
**CreatedAt** | **string** |  | 
**UpdatedAt** | **string** |  | 
**ConnectionUrl** | Pointer to **string** |  | [optional] 
**Type** | **string** |  | 
**Default** | **bool** |  | 
**Links** | Pointer to [**DatabaseCredentialLinks**](DatabaseCredentialLinks.md) |  | [optional] 

## Methods

### NewDatabaseCredential

`func NewDatabaseCredential(id int32, metaType string, createdAt string, updatedAt string, type_ string, default_ bool, ) *DatabaseCredential`

NewDatabaseCredential instantiates a new DatabaseCredential object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabaseCredentialWithDefaults

`func NewDatabaseCredentialWithDefaults() *DatabaseCredential`

NewDatabaseCredentialWithDefaults instantiates a new DatabaseCredential object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DatabaseCredential) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DatabaseCredential) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DatabaseCredential) SetId(v int32)`

SetId sets Id field to given value.


### GetMetaType

`func (o *DatabaseCredential) GetMetaType() string`

GetMetaType returns the MetaType field if non-nil, zero value otherwise.

### GetMetaTypeOk

`func (o *DatabaseCredential) GetMetaTypeOk() (*string, bool)`

GetMetaTypeOk returns a tuple with the MetaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaType

`func (o *DatabaseCredential) SetMetaType(v string)`

SetMetaType sets MetaType field to given value.


### GetCreatedAt

`func (o *DatabaseCredential) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DatabaseCredential) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DatabaseCredential) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *DatabaseCredential) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DatabaseCredential) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DatabaseCredential) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetConnectionUrl

`func (o *DatabaseCredential) GetConnectionUrl() string`

GetConnectionUrl returns the ConnectionUrl field if non-nil, zero value otherwise.

### GetConnectionUrlOk

`func (o *DatabaseCredential) GetConnectionUrlOk() (*string, bool)`

GetConnectionUrlOk returns a tuple with the ConnectionUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionUrl

`func (o *DatabaseCredential) SetConnectionUrl(v string)`

SetConnectionUrl sets ConnectionUrl field to given value.

### HasConnectionUrl

`func (o *DatabaseCredential) HasConnectionUrl() bool`

HasConnectionUrl returns a boolean if a field has been set.

### GetType

`func (o *DatabaseCredential) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DatabaseCredential) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DatabaseCredential) SetType(v string)`

SetType sets Type field to given value.


### GetDefault

`func (o *DatabaseCredential) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *DatabaseCredential) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *DatabaseCredential) SetDefault(v bool)`

SetDefault sets Default field to given value.


### GetLinks

`func (o *DatabaseCredential) GetLinks() DatabaseCredentialLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DatabaseCredential) GetLinksOk() (*DatabaseCredentialLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DatabaseCredential) SetLinks(v DatabaseCredentialLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *DatabaseCredential) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


