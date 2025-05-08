# LlmIntegration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**MetaType** | **string** |  | 
**CreatedAt** | **string** |  | 
**UpdatedAt** | **string** |  | 
**ProviderType** | **string** |  | 
**OrganizationId** | **string** |  | 
**ApiKey** | **NullableString** |  | 
**BaseUrl** | **NullableString** |  | 
**OpenaiOrganization** | **NullableString** |  | 
**ApiVersion** | **NullableString** |  | 
**AwsAccessKeyId** | **NullableString** |  | 
**AwsSecretAccessKey** | **NullableString** |  | 
**AwsRegion** | **NullableString** |  | 
**Links** | Pointer to [**IntegrationLinks**](IntegrationLinks.md) |  | [optional] 

## Methods

### NewLlmIntegration

`func NewLlmIntegration(id int32, metaType string, createdAt string, updatedAt string, providerType string, organizationId string, apiKey NullableString, baseUrl NullableString, openaiOrganization NullableString, apiVersion NullableString, awsAccessKeyId NullableString, awsSecretAccessKey NullableString, awsRegion NullableString, ) *LlmIntegration`

NewLlmIntegration instantiates a new LlmIntegration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLlmIntegrationWithDefaults

`func NewLlmIntegrationWithDefaults() *LlmIntegration`

NewLlmIntegrationWithDefaults instantiates a new LlmIntegration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LlmIntegration) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LlmIntegration) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LlmIntegration) SetId(v int32)`

SetId sets Id field to given value.


### GetMetaType

`func (o *LlmIntegration) GetMetaType() string`

GetMetaType returns the MetaType field if non-nil, zero value otherwise.

### GetMetaTypeOk

`func (o *LlmIntegration) GetMetaTypeOk() (*string, bool)`

GetMetaTypeOk returns a tuple with the MetaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaType

`func (o *LlmIntegration) SetMetaType(v string)`

SetMetaType sets MetaType field to given value.


### GetCreatedAt

`func (o *LlmIntegration) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *LlmIntegration) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *LlmIntegration) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *LlmIntegration) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *LlmIntegration) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *LlmIntegration) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetProviderType

`func (o *LlmIntegration) GetProviderType() string`

GetProviderType returns the ProviderType field if non-nil, zero value otherwise.

### GetProviderTypeOk

`func (o *LlmIntegration) GetProviderTypeOk() (*string, bool)`

GetProviderTypeOk returns a tuple with the ProviderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderType

`func (o *LlmIntegration) SetProviderType(v string)`

SetProviderType sets ProviderType field to given value.


### GetOrganizationId

`func (o *LlmIntegration) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *LlmIntegration) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *LlmIntegration) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.


### GetApiKey

`func (o *LlmIntegration) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *LlmIntegration) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *LlmIntegration) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.


### SetApiKeyNil

`func (o *LlmIntegration) SetApiKeyNil(b bool)`

 SetApiKeyNil sets the value for ApiKey to be an explicit nil

### UnsetApiKey
`func (o *LlmIntegration) UnsetApiKey()`

UnsetApiKey ensures that no value is present for ApiKey, not even an explicit nil
### GetBaseUrl

`func (o *LlmIntegration) GetBaseUrl() string`

GetBaseUrl returns the BaseUrl field if non-nil, zero value otherwise.

### GetBaseUrlOk

`func (o *LlmIntegration) GetBaseUrlOk() (*string, bool)`

GetBaseUrlOk returns a tuple with the BaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUrl

`func (o *LlmIntegration) SetBaseUrl(v string)`

SetBaseUrl sets BaseUrl field to given value.


### SetBaseUrlNil

`func (o *LlmIntegration) SetBaseUrlNil(b bool)`

 SetBaseUrlNil sets the value for BaseUrl to be an explicit nil

### UnsetBaseUrl
`func (o *LlmIntegration) UnsetBaseUrl()`

UnsetBaseUrl ensures that no value is present for BaseUrl, not even an explicit nil
### GetOpenaiOrganization

`func (o *LlmIntegration) GetOpenaiOrganization() string`

GetOpenaiOrganization returns the OpenaiOrganization field if non-nil, zero value otherwise.

### GetOpenaiOrganizationOk

`func (o *LlmIntegration) GetOpenaiOrganizationOk() (*string, bool)`

GetOpenaiOrganizationOk returns a tuple with the OpenaiOrganization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenaiOrganization

`func (o *LlmIntegration) SetOpenaiOrganization(v string)`

SetOpenaiOrganization sets OpenaiOrganization field to given value.


### SetOpenaiOrganizationNil

`func (o *LlmIntegration) SetOpenaiOrganizationNil(b bool)`

 SetOpenaiOrganizationNil sets the value for OpenaiOrganization to be an explicit nil

### UnsetOpenaiOrganization
`func (o *LlmIntegration) UnsetOpenaiOrganization()`

UnsetOpenaiOrganization ensures that no value is present for OpenaiOrganization, not even an explicit nil
### GetApiVersion

`func (o *LlmIntegration) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *LlmIntegration) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *LlmIntegration) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### SetApiVersionNil

`func (o *LlmIntegration) SetApiVersionNil(b bool)`

 SetApiVersionNil sets the value for ApiVersion to be an explicit nil

### UnsetApiVersion
`func (o *LlmIntegration) UnsetApiVersion()`

UnsetApiVersion ensures that no value is present for ApiVersion, not even an explicit nil
### GetAwsAccessKeyId

`func (o *LlmIntegration) GetAwsAccessKeyId() string`

GetAwsAccessKeyId returns the AwsAccessKeyId field if non-nil, zero value otherwise.

### GetAwsAccessKeyIdOk

`func (o *LlmIntegration) GetAwsAccessKeyIdOk() (*string, bool)`

GetAwsAccessKeyIdOk returns a tuple with the AwsAccessKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsAccessKeyId

`func (o *LlmIntegration) SetAwsAccessKeyId(v string)`

SetAwsAccessKeyId sets AwsAccessKeyId field to given value.


### SetAwsAccessKeyIdNil

`func (o *LlmIntegration) SetAwsAccessKeyIdNil(b bool)`

 SetAwsAccessKeyIdNil sets the value for AwsAccessKeyId to be an explicit nil

### UnsetAwsAccessKeyId
`func (o *LlmIntegration) UnsetAwsAccessKeyId()`

UnsetAwsAccessKeyId ensures that no value is present for AwsAccessKeyId, not even an explicit nil
### GetAwsSecretAccessKey

`func (o *LlmIntegration) GetAwsSecretAccessKey() string`

GetAwsSecretAccessKey returns the AwsSecretAccessKey field if non-nil, zero value otherwise.

### GetAwsSecretAccessKeyOk

`func (o *LlmIntegration) GetAwsSecretAccessKeyOk() (*string, bool)`

GetAwsSecretAccessKeyOk returns a tuple with the AwsSecretAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsSecretAccessKey

`func (o *LlmIntegration) SetAwsSecretAccessKey(v string)`

SetAwsSecretAccessKey sets AwsSecretAccessKey field to given value.


### SetAwsSecretAccessKeyNil

`func (o *LlmIntegration) SetAwsSecretAccessKeyNil(b bool)`

 SetAwsSecretAccessKeyNil sets the value for AwsSecretAccessKey to be an explicit nil

### UnsetAwsSecretAccessKey
`func (o *LlmIntegration) UnsetAwsSecretAccessKey()`

UnsetAwsSecretAccessKey ensures that no value is present for AwsSecretAccessKey, not even an explicit nil
### GetAwsRegion

`func (o *LlmIntegration) GetAwsRegion() string`

GetAwsRegion returns the AwsRegion field if non-nil, zero value otherwise.

### GetAwsRegionOk

`func (o *LlmIntegration) GetAwsRegionOk() (*string, bool)`

GetAwsRegionOk returns a tuple with the AwsRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsRegion

`func (o *LlmIntegration) SetAwsRegion(v string)`

SetAwsRegion sets AwsRegion field to given value.


### SetAwsRegionNil

`func (o *LlmIntegration) SetAwsRegionNil(b bool)`

 SetAwsRegionNil sets the value for AwsRegion to be an explicit nil

### UnsetAwsRegion
`func (o *LlmIntegration) UnsetAwsRegion()`

UnsetAwsRegion ensures that no value is present for AwsRegion, not even an explicit nil
### GetLinks

`func (o *LlmIntegration) GetLinks() IntegrationLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *LlmIntegration) GetLinksOk() (*IntegrationLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *LlmIntegration) SetLinks(v IntegrationLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *LlmIntegration) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


