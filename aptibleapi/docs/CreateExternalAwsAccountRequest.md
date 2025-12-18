# CreateExternalAwsAccountRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrganizationId** | **string** |  | 
**AwsAccountId** | **string** |  | 
**AccountName** | **string** |  | 
**AwsRegionPrimary** | **string** |  | 
**DiscoveryEnabled** | Pointer to **bool** |  | [optional] 
**DiscoveryFrequency** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateExternalAwsAccountRequest

`func NewCreateExternalAwsAccountRequest(organizationId string, awsAccountId string, accountName string, awsRegionPrimary string, ) *CreateExternalAwsAccountRequest`

NewCreateExternalAwsAccountRequest instantiates a new CreateExternalAwsAccountRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateExternalAwsAccountRequestWithDefaults

`func NewCreateExternalAwsAccountRequestWithDefaults() *CreateExternalAwsAccountRequest`

NewCreateExternalAwsAccountRequestWithDefaults instantiates a new CreateExternalAwsAccountRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrganizationId

`func (o *CreateExternalAwsAccountRequest) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *CreateExternalAwsAccountRequest) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *CreateExternalAwsAccountRequest) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.


### GetAwsAccountId

`func (o *CreateExternalAwsAccountRequest) GetAwsAccountId() string`

GetAwsAccountId returns the AwsAccountId field if non-nil, zero value otherwise.

### GetAwsAccountIdOk

`func (o *CreateExternalAwsAccountRequest) GetAwsAccountIdOk() (*string, bool)`

GetAwsAccountIdOk returns a tuple with the AwsAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsAccountId

`func (o *CreateExternalAwsAccountRequest) SetAwsAccountId(v string)`

SetAwsAccountId sets AwsAccountId field to given value.


### GetAccountName

`func (o *CreateExternalAwsAccountRequest) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *CreateExternalAwsAccountRequest) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *CreateExternalAwsAccountRequest) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.


### GetAwsRegionPrimary

`func (o *CreateExternalAwsAccountRequest) GetAwsRegionPrimary() string`

GetAwsRegionPrimary returns the AwsRegionPrimary field if non-nil, zero value otherwise.

### GetAwsRegionPrimaryOk

`func (o *CreateExternalAwsAccountRequest) GetAwsRegionPrimaryOk() (*string, bool)`

GetAwsRegionPrimaryOk returns a tuple with the AwsRegionPrimary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsRegionPrimary

`func (o *CreateExternalAwsAccountRequest) SetAwsRegionPrimary(v string)`

SetAwsRegionPrimary sets AwsRegionPrimary field to given value.


### GetDiscoveryEnabled

`func (o *CreateExternalAwsAccountRequest) GetDiscoveryEnabled() bool`

GetDiscoveryEnabled returns the DiscoveryEnabled field if non-nil, zero value otherwise.

### GetDiscoveryEnabledOk

`func (o *CreateExternalAwsAccountRequest) GetDiscoveryEnabledOk() (*bool, bool)`

GetDiscoveryEnabledOk returns a tuple with the DiscoveryEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveryEnabled

`func (o *CreateExternalAwsAccountRequest) SetDiscoveryEnabled(v bool)`

SetDiscoveryEnabled sets DiscoveryEnabled field to given value.

### HasDiscoveryEnabled

`func (o *CreateExternalAwsAccountRequest) HasDiscoveryEnabled() bool`

HasDiscoveryEnabled returns a boolean if a field has been set.

### GetDiscoveryFrequency

`func (o *CreateExternalAwsAccountRequest) GetDiscoveryFrequency() string`

GetDiscoveryFrequency returns the DiscoveryFrequency field if non-nil, zero value otherwise.

### GetDiscoveryFrequencyOk

`func (o *CreateExternalAwsAccountRequest) GetDiscoveryFrequencyOk() (*string, bool)`

GetDiscoveryFrequencyOk returns a tuple with the DiscoveryFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveryFrequency

`func (o *CreateExternalAwsAccountRequest) SetDiscoveryFrequency(v string)`

SetDiscoveryFrequency sets DiscoveryFrequency field to given value.

### HasDiscoveryFrequency

`func (o *CreateExternalAwsAccountRequest) HasDiscoveryFrequency() bool`

HasDiscoveryFrequency returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


