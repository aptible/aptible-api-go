# CreateLlmIntegrationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProviderType** | **string** |  | 
**OrganizationId** | **string** |  | 
**ApiKey** | Pointer to **string** |  | [optional] 
**BaseUrl** | Pointer to **string** |  | [optional] 
**OpenaiOrganization** | Pointer to **string** |  | [optional] 
**ApiVersion** | Pointer to **string** |  | [optional] 
**AwsAccessKeyId** | Pointer to **string** |  | [optional] 
**AwsSecretAccessKey** | Pointer to **string** |  | [optional] 
**AwsRegion** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateLlmIntegrationRequest

`func NewCreateLlmIntegrationRequest(providerType string, organizationId string, ) *CreateLlmIntegrationRequest`

NewCreateLlmIntegrationRequest instantiates a new CreateLlmIntegrationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateLlmIntegrationRequestWithDefaults

`func NewCreateLlmIntegrationRequestWithDefaults() *CreateLlmIntegrationRequest`

NewCreateLlmIntegrationRequestWithDefaults instantiates a new CreateLlmIntegrationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProviderType

`func (o *CreateLlmIntegrationRequest) GetProviderType() string`

GetProviderType returns the ProviderType field if non-nil, zero value otherwise.

### GetProviderTypeOk

`func (o *CreateLlmIntegrationRequest) GetProviderTypeOk() (*string, bool)`

GetProviderTypeOk returns a tuple with the ProviderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderType

`func (o *CreateLlmIntegrationRequest) SetProviderType(v string)`

SetProviderType sets ProviderType field to given value.


### GetOrganizationId

`func (o *CreateLlmIntegrationRequest) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *CreateLlmIntegrationRequest) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *CreateLlmIntegrationRequest) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.


### GetApiKey

`func (o *CreateLlmIntegrationRequest) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *CreateLlmIntegrationRequest) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *CreateLlmIntegrationRequest) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *CreateLlmIntegrationRequest) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### GetBaseUrl

`func (o *CreateLlmIntegrationRequest) GetBaseUrl() string`

GetBaseUrl returns the BaseUrl field if non-nil, zero value otherwise.

### GetBaseUrlOk

`func (o *CreateLlmIntegrationRequest) GetBaseUrlOk() (*string, bool)`

GetBaseUrlOk returns a tuple with the BaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUrl

`func (o *CreateLlmIntegrationRequest) SetBaseUrl(v string)`

SetBaseUrl sets BaseUrl field to given value.

### HasBaseUrl

`func (o *CreateLlmIntegrationRequest) HasBaseUrl() bool`

HasBaseUrl returns a boolean if a field has been set.

### GetOpenaiOrganization

`func (o *CreateLlmIntegrationRequest) GetOpenaiOrganization() string`

GetOpenaiOrganization returns the OpenaiOrganization field if non-nil, zero value otherwise.

### GetOpenaiOrganizationOk

`func (o *CreateLlmIntegrationRequest) GetOpenaiOrganizationOk() (*string, bool)`

GetOpenaiOrganizationOk returns a tuple with the OpenaiOrganization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenaiOrganization

`func (o *CreateLlmIntegrationRequest) SetOpenaiOrganization(v string)`

SetOpenaiOrganization sets OpenaiOrganization field to given value.

### HasOpenaiOrganization

`func (o *CreateLlmIntegrationRequest) HasOpenaiOrganization() bool`

HasOpenaiOrganization returns a boolean if a field has been set.

### GetApiVersion

`func (o *CreateLlmIntegrationRequest) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *CreateLlmIntegrationRequest) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *CreateLlmIntegrationRequest) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *CreateLlmIntegrationRequest) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetAwsAccessKeyId

`func (o *CreateLlmIntegrationRequest) GetAwsAccessKeyId() string`

GetAwsAccessKeyId returns the AwsAccessKeyId field if non-nil, zero value otherwise.

### GetAwsAccessKeyIdOk

`func (o *CreateLlmIntegrationRequest) GetAwsAccessKeyIdOk() (*string, bool)`

GetAwsAccessKeyIdOk returns a tuple with the AwsAccessKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsAccessKeyId

`func (o *CreateLlmIntegrationRequest) SetAwsAccessKeyId(v string)`

SetAwsAccessKeyId sets AwsAccessKeyId field to given value.

### HasAwsAccessKeyId

`func (o *CreateLlmIntegrationRequest) HasAwsAccessKeyId() bool`

HasAwsAccessKeyId returns a boolean if a field has been set.

### GetAwsSecretAccessKey

`func (o *CreateLlmIntegrationRequest) GetAwsSecretAccessKey() string`

GetAwsSecretAccessKey returns the AwsSecretAccessKey field if non-nil, zero value otherwise.

### GetAwsSecretAccessKeyOk

`func (o *CreateLlmIntegrationRequest) GetAwsSecretAccessKeyOk() (*string, bool)`

GetAwsSecretAccessKeyOk returns a tuple with the AwsSecretAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsSecretAccessKey

`func (o *CreateLlmIntegrationRequest) SetAwsSecretAccessKey(v string)`

SetAwsSecretAccessKey sets AwsSecretAccessKey field to given value.

### HasAwsSecretAccessKey

`func (o *CreateLlmIntegrationRequest) HasAwsSecretAccessKey() bool`

HasAwsSecretAccessKey returns a boolean if a field has been set.

### GetAwsRegion

`func (o *CreateLlmIntegrationRequest) GetAwsRegion() string`

GetAwsRegion returns the AwsRegion field if non-nil, zero value otherwise.

### GetAwsRegionOk

`func (o *CreateLlmIntegrationRequest) GetAwsRegionOk() (*string, bool)`

GetAwsRegionOk returns a tuple with the AwsRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsRegion

`func (o *CreateLlmIntegrationRequest) SetAwsRegion(v string)`

SetAwsRegion sets AwsRegion field to given value.

### HasAwsRegion

`func (o *CreateLlmIntegrationRequest) HasAwsRegion() bool`

HasAwsRegion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


