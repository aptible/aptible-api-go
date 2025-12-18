# CreateExternalAwsDatabaseCredentialRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Default** | Pointer to **bool** |  | [optional] 
**ConnectionUrl** | **string** |  | 

## Methods

### NewCreateExternalAwsDatabaseCredentialRequest

`func NewCreateExternalAwsDatabaseCredentialRequest(type_ string, connectionUrl string, ) *CreateExternalAwsDatabaseCredentialRequest`

NewCreateExternalAwsDatabaseCredentialRequest instantiates a new CreateExternalAwsDatabaseCredentialRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateExternalAwsDatabaseCredentialRequestWithDefaults

`func NewCreateExternalAwsDatabaseCredentialRequestWithDefaults() *CreateExternalAwsDatabaseCredentialRequest`

NewCreateExternalAwsDatabaseCredentialRequestWithDefaults instantiates a new CreateExternalAwsDatabaseCredentialRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CreateExternalAwsDatabaseCredentialRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateExternalAwsDatabaseCredentialRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateExternalAwsDatabaseCredentialRequest) SetType(v string)`

SetType sets Type field to given value.


### GetDefault

`func (o *CreateExternalAwsDatabaseCredentialRequest) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *CreateExternalAwsDatabaseCredentialRequest) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *CreateExternalAwsDatabaseCredentialRequest) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *CreateExternalAwsDatabaseCredentialRequest) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetConnectionUrl

`func (o *CreateExternalAwsDatabaseCredentialRequest) GetConnectionUrl() string`

GetConnectionUrl returns the ConnectionUrl field if non-nil, zero value otherwise.

### GetConnectionUrlOk

`func (o *CreateExternalAwsDatabaseCredentialRequest) GetConnectionUrlOk() (*string, bool)`

GetConnectionUrlOk returns a tuple with the ConnectionUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionUrl

`func (o *CreateExternalAwsDatabaseCredentialRequest) SetConnectionUrl(v string)`

SetConnectionUrl sets ConnectionUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


