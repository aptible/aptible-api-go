# UpdateAppExternalAwsRdsConnectionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DatabaseName** | Pointer to **string** |  | [optional] 
**DatabaseUser** | Pointer to **string** |  | [optional] 
**Superuser** | Pointer to **bool** |  | [optional] 

## Methods

### NewUpdateAppExternalAwsRdsConnectionRequest

`func NewUpdateAppExternalAwsRdsConnectionRequest() *UpdateAppExternalAwsRdsConnectionRequest`

NewUpdateAppExternalAwsRdsConnectionRequest instantiates a new UpdateAppExternalAwsRdsConnectionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAppExternalAwsRdsConnectionRequestWithDefaults

`func NewUpdateAppExternalAwsRdsConnectionRequestWithDefaults() *UpdateAppExternalAwsRdsConnectionRequest`

NewUpdateAppExternalAwsRdsConnectionRequestWithDefaults instantiates a new UpdateAppExternalAwsRdsConnectionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatabaseName

`func (o *UpdateAppExternalAwsRdsConnectionRequest) GetDatabaseName() string`

GetDatabaseName returns the DatabaseName field if non-nil, zero value otherwise.

### GetDatabaseNameOk

`func (o *UpdateAppExternalAwsRdsConnectionRequest) GetDatabaseNameOk() (*string, bool)`

GetDatabaseNameOk returns a tuple with the DatabaseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseName

`func (o *UpdateAppExternalAwsRdsConnectionRequest) SetDatabaseName(v string)`

SetDatabaseName sets DatabaseName field to given value.

### HasDatabaseName

`func (o *UpdateAppExternalAwsRdsConnectionRequest) HasDatabaseName() bool`

HasDatabaseName returns a boolean if a field has been set.

### GetDatabaseUser

`func (o *UpdateAppExternalAwsRdsConnectionRequest) GetDatabaseUser() string`

GetDatabaseUser returns the DatabaseUser field if non-nil, zero value otherwise.

### GetDatabaseUserOk

`func (o *UpdateAppExternalAwsRdsConnectionRequest) GetDatabaseUserOk() (*string, bool)`

GetDatabaseUserOk returns a tuple with the DatabaseUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseUser

`func (o *UpdateAppExternalAwsRdsConnectionRequest) SetDatabaseUser(v string)`

SetDatabaseUser sets DatabaseUser field to given value.

### HasDatabaseUser

`func (o *UpdateAppExternalAwsRdsConnectionRequest) HasDatabaseUser() bool`

HasDatabaseUser returns a boolean if a field has been set.

### GetSuperuser

`func (o *UpdateAppExternalAwsRdsConnectionRequest) GetSuperuser() bool`

GetSuperuser returns the Superuser field if non-nil, zero value otherwise.

### GetSuperuserOk

`func (o *UpdateAppExternalAwsRdsConnectionRequest) GetSuperuserOk() (*bool, bool)`

GetSuperuserOk returns a tuple with the Superuser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuperuser

`func (o *UpdateAppExternalAwsRdsConnectionRequest) SetSuperuser(v bool)`

SetSuperuser sets Superuser field to given value.

### HasSuperuser

`func (o *UpdateAppExternalAwsRdsConnectionRequest) HasSuperuser() bool`

HasSuperuser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


