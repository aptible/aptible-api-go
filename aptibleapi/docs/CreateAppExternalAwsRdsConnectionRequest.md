# CreateAppExternalAwsRdsConnectionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalAwsResourceId** | Pointer to **int32** |  | [optional] 
**DatabaseName** | Pointer to **string** | The name of the database to connect to | [optional] 
**DatabaseUser** | Pointer to **string** | The user to connect to the database as | [optional] 
**Superuser** | Pointer to **bool** |  | [optional] 

## Methods

### NewCreateAppExternalAwsRdsConnectionRequest

`func NewCreateAppExternalAwsRdsConnectionRequest() *CreateAppExternalAwsRdsConnectionRequest`

NewCreateAppExternalAwsRdsConnectionRequest instantiates a new CreateAppExternalAwsRdsConnectionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAppExternalAwsRdsConnectionRequestWithDefaults

`func NewCreateAppExternalAwsRdsConnectionRequestWithDefaults() *CreateAppExternalAwsRdsConnectionRequest`

NewCreateAppExternalAwsRdsConnectionRequestWithDefaults instantiates a new CreateAppExternalAwsRdsConnectionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalAwsResourceId

`func (o *CreateAppExternalAwsRdsConnectionRequest) GetExternalAwsResourceId() int32`

GetExternalAwsResourceId returns the ExternalAwsResourceId field if non-nil, zero value otherwise.

### GetExternalAwsResourceIdOk

`func (o *CreateAppExternalAwsRdsConnectionRequest) GetExternalAwsResourceIdOk() (*int32, bool)`

GetExternalAwsResourceIdOk returns a tuple with the ExternalAwsResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalAwsResourceId

`func (o *CreateAppExternalAwsRdsConnectionRequest) SetExternalAwsResourceId(v int32)`

SetExternalAwsResourceId sets ExternalAwsResourceId field to given value.

### HasExternalAwsResourceId

`func (o *CreateAppExternalAwsRdsConnectionRequest) HasExternalAwsResourceId() bool`

HasExternalAwsResourceId returns a boolean if a field has been set.

### GetDatabaseName

`func (o *CreateAppExternalAwsRdsConnectionRequest) GetDatabaseName() string`

GetDatabaseName returns the DatabaseName field if non-nil, zero value otherwise.

### GetDatabaseNameOk

`func (o *CreateAppExternalAwsRdsConnectionRequest) GetDatabaseNameOk() (*string, bool)`

GetDatabaseNameOk returns a tuple with the DatabaseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseName

`func (o *CreateAppExternalAwsRdsConnectionRequest) SetDatabaseName(v string)`

SetDatabaseName sets DatabaseName field to given value.

### HasDatabaseName

`func (o *CreateAppExternalAwsRdsConnectionRequest) HasDatabaseName() bool`

HasDatabaseName returns a boolean if a field has been set.

### GetDatabaseUser

`func (o *CreateAppExternalAwsRdsConnectionRequest) GetDatabaseUser() string`

GetDatabaseUser returns the DatabaseUser field if non-nil, zero value otherwise.

### GetDatabaseUserOk

`func (o *CreateAppExternalAwsRdsConnectionRequest) GetDatabaseUserOk() (*string, bool)`

GetDatabaseUserOk returns a tuple with the DatabaseUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseUser

`func (o *CreateAppExternalAwsRdsConnectionRequest) SetDatabaseUser(v string)`

SetDatabaseUser sets DatabaseUser field to given value.

### HasDatabaseUser

`func (o *CreateAppExternalAwsRdsConnectionRequest) HasDatabaseUser() bool`

HasDatabaseUser returns a boolean if a field has been set.

### GetSuperuser

`func (o *CreateAppExternalAwsRdsConnectionRequest) GetSuperuser() bool`

GetSuperuser returns the Superuser field if non-nil, zero value otherwise.

### GetSuperuserOk

`func (o *CreateAppExternalAwsRdsConnectionRequest) GetSuperuserOk() (*bool, bool)`

GetSuperuserOk returns a tuple with the Superuser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuperuser

`func (o *CreateAppExternalAwsRdsConnectionRequest) SetSuperuser(v bool)`

SetSuperuser sets Superuser field to given value.

### HasSuperuser

`func (o *CreateAppExternalAwsRdsConnectionRequest) HasSuperuser() bool`

HasSuperuser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


