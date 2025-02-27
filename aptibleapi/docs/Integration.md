# Integration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**MetaType** | **string** |  | 
**CreatedAt** | **string** |  | 
**UpdatedAt** | **string** |  | 
**Type** | **string** |  | 
**OrganizationId** | **string** |  | 
**AwsRoleArn** | **NullableString** |  | 
**ApiKey** | **NullableString** |  | 
**AppKey** | **NullableString** |  | 
**Host** | **NullableString** |  | 
**Port** | **NullableString** |  | 
**Username** | **NullableString** |  | 
**Password** | **NullableString** |  | 
**Database** | **NullableString** |  | 
**Links** | Pointer to [**IntegrationLinks**](IntegrationLinks.md) |  | [optional] 

## Methods

### NewIntegration

`func NewIntegration(id int32, metaType string, createdAt string, updatedAt string, type_ string, organizationId string, awsRoleArn NullableString, apiKey NullableString, appKey NullableString, host NullableString, port NullableString, username NullableString, password NullableString, database NullableString, ) *Integration`

NewIntegration instantiates a new Integration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationWithDefaults

`func NewIntegrationWithDefaults() *Integration`

NewIntegrationWithDefaults instantiates a new Integration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Integration) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Integration) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Integration) SetId(v int32)`

SetId sets Id field to given value.


### GetMetaType

`func (o *Integration) GetMetaType() string`

GetMetaType returns the MetaType field if non-nil, zero value otherwise.

### GetMetaTypeOk

`func (o *Integration) GetMetaTypeOk() (*string, bool)`

GetMetaTypeOk returns a tuple with the MetaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaType

`func (o *Integration) SetMetaType(v string)`

SetMetaType sets MetaType field to given value.


### GetCreatedAt

`func (o *Integration) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Integration) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Integration) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Integration) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Integration) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Integration) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetType

`func (o *Integration) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Integration) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Integration) SetType(v string)`

SetType sets Type field to given value.


### GetOrganizationId

`func (o *Integration) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *Integration) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *Integration) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.


### GetAwsRoleArn

`func (o *Integration) GetAwsRoleArn() string`

GetAwsRoleArn returns the AwsRoleArn field if non-nil, zero value otherwise.

### GetAwsRoleArnOk

`func (o *Integration) GetAwsRoleArnOk() (*string, bool)`

GetAwsRoleArnOk returns a tuple with the AwsRoleArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsRoleArn

`func (o *Integration) SetAwsRoleArn(v string)`

SetAwsRoleArn sets AwsRoleArn field to given value.


### SetAwsRoleArnNil

`func (o *Integration) SetAwsRoleArnNil(b bool)`

 SetAwsRoleArnNil sets the value for AwsRoleArn to be an explicit nil

### UnsetAwsRoleArn
`func (o *Integration) UnsetAwsRoleArn()`

UnsetAwsRoleArn ensures that no value is present for AwsRoleArn, not even an explicit nil
### GetApiKey

`func (o *Integration) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *Integration) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *Integration) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.


### SetApiKeyNil

`func (o *Integration) SetApiKeyNil(b bool)`

 SetApiKeyNil sets the value for ApiKey to be an explicit nil

### UnsetApiKey
`func (o *Integration) UnsetApiKey()`

UnsetApiKey ensures that no value is present for ApiKey, not even an explicit nil
### GetAppKey

`func (o *Integration) GetAppKey() string`

GetAppKey returns the AppKey field if non-nil, zero value otherwise.

### GetAppKeyOk

`func (o *Integration) GetAppKeyOk() (*string, bool)`

GetAppKeyOk returns a tuple with the AppKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppKey

`func (o *Integration) SetAppKey(v string)`

SetAppKey sets AppKey field to given value.


### SetAppKeyNil

`func (o *Integration) SetAppKeyNil(b bool)`

 SetAppKeyNil sets the value for AppKey to be an explicit nil

### UnsetAppKey
`func (o *Integration) UnsetAppKey()`

UnsetAppKey ensures that no value is present for AppKey, not even an explicit nil
### GetHost

`func (o *Integration) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *Integration) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *Integration) SetHost(v string)`

SetHost sets Host field to given value.


### SetHostNil

`func (o *Integration) SetHostNil(b bool)`

 SetHostNil sets the value for Host to be an explicit nil

### UnsetHost
`func (o *Integration) UnsetHost()`

UnsetHost ensures that no value is present for Host, not even an explicit nil
### GetPort

`func (o *Integration) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *Integration) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *Integration) SetPort(v string)`

SetPort sets Port field to given value.


### SetPortNil

`func (o *Integration) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *Integration) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil
### GetUsername

`func (o *Integration) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *Integration) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *Integration) SetUsername(v string)`

SetUsername sets Username field to given value.


### SetUsernameNil

`func (o *Integration) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *Integration) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetPassword

`func (o *Integration) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *Integration) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *Integration) SetPassword(v string)`

SetPassword sets Password field to given value.


### SetPasswordNil

`func (o *Integration) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *Integration) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetDatabase

`func (o *Integration) GetDatabase() string`

GetDatabase returns the Database field if non-nil, zero value otherwise.

### GetDatabaseOk

`func (o *Integration) GetDatabaseOk() (*string, bool)`

GetDatabaseOk returns a tuple with the Database field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabase

`func (o *Integration) SetDatabase(v string)`

SetDatabase sets Database field to given value.


### SetDatabaseNil

`func (o *Integration) SetDatabaseNil(b bool)`

 SetDatabaseNil sets the value for Database to be an explicit nil

### UnsetDatabase
`func (o *Integration) UnsetDatabase()`

UnsetDatabase ensures that no value is present for Database, not even an explicit nil
### GetLinks

`func (o *Integration) GetLinks() IntegrationLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Integration) GetLinksOk() (*IntegrationLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Integration) SetLinks(v IntegrationLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Integration) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


