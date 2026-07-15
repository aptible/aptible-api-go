# LlmGatewayConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Database ID of the LLM Gateway Configuration | 
**MetaType** | **string** |  | 
**OrganizationId** | **string** | The organization this configuration belongs to | 
**TeamId** | **string** | The LiteLLM team ID for this organization | 
**Trial** | **bool** | Whether the organization is on a trial AI Gateway plan | 
**SpendLimit** | **string** | Monthly spend limit in USD (decimal, serialized as string) | 
**Models** | **[]string** | List of permitted model identifiers | 
**CreatedAt** | **NullableString** |  | 
**UpdatedAt** | **NullableString** |  | 
**Links** | Pointer to [**LlmGatewayConfigurationLinks**](LlmGatewayConfigurationLinks.md) |  | [optional] 

## Methods

### NewLlmGatewayConfiguration

`func NewLlmGatewayConfiguration(id int32, metaType string, organizationId string, teamId string, trial bool, spendLimit string, models []string, createdAt NullableString, updatedAt NullableString, ) *LlmGatewayConfiguration`

NewLlmGatewayConfiguration instantiates a new LlmGatewayConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLlmGatewayConfigurationWithDefaults

`func NewLlmGatewayConfigurationWithDefaults() *LlmGatewayConfiguration`

NewLlmGatewayConfigurationWithDefaults instantiates a new LlmGatewayConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LlmGatewayConfiguration) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LlmGatewayConfiguration) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LlmGatewayConfiguration) SetId(v int32)`

SetId sets Id field to given value.


### GetMetaType

`func (o *LlmGatewayConfiguration) GetMetaType() string`

GetMetaType returns the MetaType field if non-nil, zero value otherwise.

### GetMetaTypeOk

`func (o *LlmGatewayConfiguration) GetMetaTypeOk() (*string, bool)`

GetMetaTypeOk returns a tuple with the MetaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaType

`func (o *LlmGatewayConfiguration) SetMetaType(v string)`

SetMetaType sets MetaType field to given value.


### GetOrganizationId

`func (o *LlmGatewayConfiguration) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *LlmGatewayConfiguration) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *LlmGatewayConfiguration) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.


### GetTeamId

`func (o *LlmGatewayConfiguration) GetTeamId() string`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *LlmGatewayConfiguration) GetTeamIdOk() (*string, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *LlmGatewayConfiguration) SetTeamId(v string)`

SetTeamId sets TeamId field to given value.


### GetTrial

`func (o *LlmGatewayConfiguration) GetTrial() bool`

GetTrial returns the Trial field if non-nil, zero value otherwise.

### GetTrialOk

`func (o *LlmGatewayConfiguration) GetTrialOk() (*bool, bool)`

GetTrialOk returns a tuple with the Trial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrial

`func (o *LlmGatewayConfiguration) SetTrial(v bool)`

SetTrial sets Trial field to given value.


### GetSpendLimit

`func (o *LlmGatewayConfiguration) GetSpendLimit() string`

GetSpendLimit returns the SpendLimit field if non-nil, zero value otherwise.

### GetSpendLimitOk

`func (o *LlmGatewayConfiguration) GetSpendLimitOk() (*string, bool)`

GetSpendLimitOk returns a tuple with the SpendLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpendLimit

`func (o *LlmGatewayConfiguration) SetSpendLimit(v string)`

SetSpendLimit sets SpendLimit field to given value.


### GetModels

`func (o *LlmGatewayConfiguration) GetModels() []string`

GetModels returns the Models field if non-nil, zero value otherwise.

### GetModelsOk

`func (o *LlmGatewayConfiguration) GetModelsOk() (*[]string, bool)`

GetModelsOk returns a tuple with the Models field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModels

`func (o *LlmGatewayConfiguration) SetModels(v []string)`

SetModels sets Models field to given value.


### GetCreatedAt

`func (o *LlmGatewayConfiguration) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *LlmGatewayConfiguration) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *LlmGatewayConfiguration) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### SetCreatedAtNil

`func (o *LlmGatewayConfiguration) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *LlmGatewayConfiguration) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *LlmGatewayConfiguration) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *LlmGatewayConfiguration) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *LlmGatewayConfiguration) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### SetUpdatedAtNil

`func (o *LlmGatewayConfiguration) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *LlmGatewayConfiguration) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetLinks

`func (o *LlmGatewayConfiguration) GetLinks() LlmGatewayConfigurationLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *LlmGatewayConfiguration) GetLinksOk() (*LlmGatewayConfigurationLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *LlmGatewayConfiguration) SetLinks(v LlmGatewayConfigurationLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *LlmGatewayConfiguration) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


