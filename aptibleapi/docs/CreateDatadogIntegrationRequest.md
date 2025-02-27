# CreateDatadogIntegrationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**OrganizationId** | **string** |  | 
**AwsRoleArn** | Pointer to **string** |  | [optional] 
**ApiKey** | Pointer to **string** |  | [optional] 
**AppKey** | Pointer to **string** |  | [optional] 
**Host** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**Database** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateDatadogIntegrationRequest

`func NewCreateDatadogIntegrationRequest(type_ string, organizationId string, ) *CreateDatadogIntegrationRequest`

NewCreateDatadogIntegrationRequest instantiates a new CreateDatadogIntegrationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDatadogIntegrationRequestWithDefaults

`func NewCreateDatadogIntegrationRequestWithDefaults() *CreateDatadogIntegrationRequest`

NewCreateDatadogIntegrationRequestWithDefaults instantiates a new CreateDatadogIntegrationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CreateDatadogIntegrationRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateDatadogIntegrationRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateDatadogIntegrationRequest) SetType(v string)`

SetType sets Type field to given value.


### GetOrganizationId

`func (o *CreateDatadogIntegrationRequest) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *CreateDatadogIntegrationRequest) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *CreateDatadogIntegrationRequest) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.


### GetAwsRoleArn

`func (o *CreateDatadogIntegrationRequest) GetAwsRoleArn() string`

GetAwsRoleArn returns the AwsRoleArn field if non-nil, zero value otherwise.

### GetAwsRoleArnOk

`func (o *CreateDatadogIntegrationRequest) GetAwsRoleArnOk() (*string, bool)`

GetAwsRoleArnOk returns a tuple with the AwsRoleArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsRoleArn

`func (o *CreateDatadogIntegrationRequest) SetAwsRoleArn(v string)`

SetAwsRoleArn sets AwsRoleArn field to given value.

### HasAwsRoleArn

`func (o *CreateDatadogIntegrationRequest) HasAwsRoleArn() bool`

HasAwsRoleArn returns a boolean if a field has been set.

### GetApiKey

`func (o *CreateDatadogIntegrationRequest) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *CreateDatadogIntegrationRequest) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *CreateDatadogIntegrationRequest) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *CreateDatadogIntegrationRequest) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### GetAppKey

`func (o *CreateDatadogIntegrationRequest) GetAppKey() string`

GetAppKey returns the AppKey field if non-nil, zero value otherwise.

### GetAppKeyOk

`func (o *CreateDatadogIntegrationRequest) GetAppKeyOk() (*string, bool)`

GetAppKeyOk returns a tuple with the AppKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppKey

`func (o *CreateDatadogIntegrationRequest) SetAppKey(v string)`

SetAppKey sets AppKey field to given value.

### HasAppKey

`func (o *CreateDatadogIntegrationRequest) HasAppKey() bool`

HasAppKey returns a boolean if a field has been set.

### GetHost

`func (o *CreateDatadogIntegrationRequest) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *CreateDatadogIntegrationRequest) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *CreateDatadogIntegrationRequest) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *CreateDatadogIntegrationRequest) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetPort

`func (o *CreateDatadogIntegrationRequest) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *CreateDatadogIntegrationRequest) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *CreateDatadogIntegrationRequest) SetPort(v string)`

SetPort sets Port field to given value.

### HasPort

`func (o *CreateDatadogIntegrationRequest) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetUsername

`func (o *CreateDatadogIntegrationRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CreateDatadogIntegrationRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CreateDatadogIntegrationRequest) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *CreateDatadogIntegrationRequest) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *CreateDatadogIntegrationRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CreateDatadogIntegrationRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CreateDatadogIntegrationRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *CreateDatadogIntegrationRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetDatabase

`func (o *CreateDatadogIntegrationRequest) GetDatabase() string`

GetDatabase returns the Database field if non-nil, zero value otherwise.

### GetDatabaseOk

`func (o *CreateDatadogIntegrationRequest) GetDatabaseOk() (*string, bool)`

GetDatabaseOk returns a tuple with the Database field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabase

`func (o *CreateDatadogIntegrationRequest) SetDatabase(v string)`

SetDatabase sets Database field to given value.

### HasDatabase

`func (o *CreateDatadogIntegrationRequest) HasDatabase() bool`

HasDatabase returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


