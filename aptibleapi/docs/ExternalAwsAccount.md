# ExternalAwsAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**MetaType** | **string** |  | 
**OrganizationId** | **string** |  | 
**AwsAccountId** | **string** |  | 
**AccountName** | **string** |  | 
**AwsRegionPrimary** | **string** |  | 
**DiscoveryEnabled** | **bool** |  | 
**DiscoveryFrequency** | **NullableString** |  | 
**LastDiscoveryAt** | **NullableString** |  | 
**Status** | **string** |  | 
**CreatedByUserEmail** | **string** |  | 
**CreatedByUserName** | **string** |  | 
**CreatedByUserId** | **string** |  | 
**CreatedAt** | **string** |  | 
**UpdatedAt** | **string** |  | 
**Links** | Pointer to [**ExternalAwsAccountLinks**](ExternalAwsAccountLinks.md) |  | [optional] 

## Methods

### NewExternalAwsAccount

`func NewExternalAwsAccount(id int32, metaType string, organizationId string, awsAccountId string, accountName string, awsRegionPrimary string, discoveryEnabled bool, discoveryFrequency NullableString, lastDiscoveryAt NullableString, status string, createdByUserEmail string, createdByUserName string, createdByUserId string, createdAt string, updatedAt string, ) *ExternalAwsAccount`

NewExternalAwsAccount instantiates a new ExternalAwsAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalAwsAccountWithDefaults

`func NewExternalAwsAccountWithDefaults() *ExternalAwsAccount`

NewExternalAwsAccountWithDefaults instantiates a new ExternalAwsAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ExternalAwsAccount) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExternalAwsAccount) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExternalAwsAccount) SetId(v int32)`

SetId sets Id field to given value.


### GetMetaType

`func (o *ExternalAwsAccount) GetMetaType() string`

GetMetaType returns the MetaType field if non-nil, zero value otherwise.

### GetMetaTypeOk

`func (o *ExternalAwsAccount) GetMetaTypeOk() (*string, bool)`

GetMetaTypeOk returns a tuple with the MetaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaType

`func (o *ExternalAwsAccount) SetMetaType(v string)`

SetMetaType sets MetaType field to given value.


### GetOrganizationId

`func (o *ExternalAwsAccount) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *ExternalAwsAccount) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *ExternalAwsAccount) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.


### GetAwsAccountId

`func (o *ExternalAwsAccount) GetAwsAccountId() string`

GetAwsAccountId returns the AwsAccountId field if non-nil, zero value otherwise.

### GetAwsAccountIdOk

`func (o *ExternalAwsAccount) GetAwsAccountIdOk() (*string, bool)`

GetAwsAccountIdOk returns a tuple with the AwsAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsAccountId

`func (o *ExternalAwsAccount) SetAwsAccountId(v string)`

SetAwsAccountId sets AwsAccountId field to given value.


### GetAccountName

`func (o *ExternalAwsAccount) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *ExternalAwsAccount) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *ExternalAwsAccount) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.


### GetAwsRegionPrimary

`func (o *ExternalAwsAccount) GetAwsRegionPrimary() string`

GetAwsRegionPrimary returns the AwsRegionPrimary field if non-nil, zero value otherwise.

### GetAwsRegionPrimaryOk

`func (o *ExternalAwsAccount) GetAwsRegionPrimaryOk() (*string, bool)`

GetAwsRegionPrimaryOk returns a tuple with the AwsRegionPrimary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsRegionPrimary

`func (o *ExternalAwsAccount) SetAwsRegionPrimary(v string)`

SetAwsRegionPrimary sets AwsRegionPrimary field to given value.


### GetDiscoveryEnabled

`func (o *ExternalAwsAccount) GetDiscoveryEnabled() bool`

GetDiscoveryEnabled returns the DiscoveryEnabled field if non-nil, zero value otherwise.

### GetDiscoveryEnabledOk

`func (o *ExternalAwsAccount) GetDiscoveryEnabledOk() (*bool, bool)`

GetDiscoveryEnabledOk returns a tuple with the DiscoveryEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveryEnabled

`func (o *ExternalAwsAccount) SetDiscoveryEnabled(v bool)`

SetDiscoveryEnabled sets DiscoveryEnabled field to given value.


### GetDiscoveryFrequency

`func (o *ExternalAwsAccount) GetDiscoveryFrequency() string`

GetDiscoveryFrequency returns the DiscoveryFrequency field if non-nil, zero value otherwise.

### GetDiscoveryFrequencyOk

`func (o *ExternalAwsAccount) GetDiscoveryFrequencyOk() (*string, bool)`

GetDiscoveryFrequencyOk returns a tuple with the DiscoveryFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveryFrequency

`func (o *ExternalAwsAccount) SetDiscoveryFrequency(v string)`

SetDiscoveryFrequency sets DiscoveryFrequency field to given value.


### SetDiscoveryFrequencyNil

`func (o *ExternalAwsAccount) SetDiscoveryFrequencyNil(b bool)`

 SetDiscoveryFrequencyNil sets the value for DiscoveryFrequency to be an explicit nil

### UnsetDiscoveryFrequency
`func (o *ExternalAwsAccount) UnsetDiscoveryFrequency()`

UnsetDiscoveryFrequency ensures that no value is present for DiscoveryFrequency, not even an explicit nil
### GetLastDiscoveryAt

`func (o *ExternalAwsAccount) GetLastDiscoveryAt() string`

GetLastDiscoveryAt returns the LastDiscoveryAt field if non-nil, zero value otherwise.

### GetLastDiscoveryAtOk

`func (o *ExternalAwsAccount) GetLastDiscoveryAtOk() (*string, bool)`

GetLastDiscoveryAtOk returns a tuple with the LastDiscoveryAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDiscoveryAt

`func (o *ExternalAwsAccount) SetLastDiscoveryAt(v string)`

SetLastDiscoveryAt sets LastDiscoveryAt field to given value.


### SetLastDiscoveryAtNil

`func (o *ExternalAwsAccount) SetLastDiscoveryAtNil(b bool)`

 SetLastDiscoveryAtNil sets the value for LastDiscoveryAt to be an explicit nil

### UnsetLastDiscoveryAt
`func (o *ExternalAwsAccount) UnsetLastDiscoveryAt()`

UnsetLastDiscoveryAt ensures that no value is present for LastDiscoveryAt, not even an explicit nil
### GetStatus

`func (o *ExternalAwsAccount) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ExternalAwsAccount) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ExternalAwsAccount) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetCreatedByUserEmail

`func (o *ExternalAwsAccount) GetCreatedByUserEmail() string`

GetCreatedByUserEmail returns the CreatedByUserEmail field if non-nil, zero value otherwise.

### GetCreatedByUserEmailOk

`func (o *ExternalAwsAccount) GetCreatedByUserEmailOk() (*string, bool)`

GetCreatedByUserEmailOk returns a tuple with the CreatedByUserEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUserEmail

`func (o *ExternalAwsAccount) SetCreatedByUserEmail(v string)`

SetCreatedByUserEmail sets CreatedByUserEmail field to given value.


### GetCreatedByUserName

`func (o *ExternalAwsAccount) GetCreatedByUserName() string`

GetCreatedByUserName returns the CreatedByUserName field if non-nil, zero value otherwise.

### GetCreatedByUserNameOk

`func (o *ExternalAwsAccount) GetCreatedByUserNameOk() (*string, bool)`

GetCreatedByUserNameOk returns a tuple with the CreatedByUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUserName

`func (o *ExternalAwsAccount) SetCreatedByUserName(v string)`

SetCreatedByUserName sets CreatedByUserName field to given value.


### GetCreatedByUserId

`func (o *ExternalAwsAccount) GetCreatedByUserId() string`

GetCreatedByUserId returns the CreatedByUserId field if non-nil, zero value otherwise.

### GetCreatedByUserIdOk

`func (o *ExternalAwsAccount) GetCreatedByUserIdOk() (*string, bool)`

GetCreatedByUserIdOk returns a tuple with the CreatedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUserId

`func (o *ExternalAwsAccount) SetCreatedByUserId(v string)`

SetCreatedByUserId sets CreatedByUserId field to given value.


### GetCreatedAt

`func (o *ExternalAwsAccount) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ExternalAwsAccount) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ExternalAwsAccount) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *ExternalAwsAccount) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ExternalAwsAccount) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ExternalAwsAccount) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetLinks

`func (o *ExternalAwsAccount) GetLinks() ExternalAwsAccountLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ExternalAwsAccount) GetLinksOk() (*ExternalAwsAccountLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ExternalAwsAccount) SetLinks(v ExternalAwsAccountLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ExternalAwsAccount) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


