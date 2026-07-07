# LlmPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Database ID of the LLM Policy | 
**MetaType** | **string** |  | 
**SpendLimit** | **NullableInt32** | Monthly spend limit in USD cents (nil means unlimited) | 
**Models** | **[]string** | List of allowed model identifiers | 
**CreatedAt** | **NullableString** |  | 
**UpdatedAt** | **NullableString** |  | 
**CurrentUsage** | Pointer to [**NullableLlmPolicyCurrentUsage**](LlmPolicyCurrentUsage.md) |  | [optional] 
**Links** | Pointer to [**BackupRetentionPolicyLinks**](BackupRetentionPolicyLinks.md) |  | [optional] 

## Methods

### NewLlmPolicy

`func NewLlmPolicy(id int32, metaType string, spendLimit NullableInt32, models []string, createdAt NullableString, updatedAt NullableString, ) *LlmPolicy`

NewLlmPolicy instantiates a new LlmPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLlmPolicyWithDefaults

`func NewLlmPolicyWithDefaults() *LlmPolicy`

NewLlmPolicyWithDefaults instantiates a new LlmPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LlmPolicy) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LlmPolicy) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LlmPolicy) SetId(v int32)`

SetId sets Id field to given value.


### GetMetaType

`func (o *LlmPolicy) GetMetaType() string`

GetMetaType returns the MetaType field if non-nil, zero value otherwise.

### GetMetaTypeOk

`func (o *LlmPolicy) GetMetaTypeOk() (*string, bool)`

GetMetaTypeOk returns a tuple with the MetaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaType

`func (o *LlmPolicy) SetMetaType(v string)`

SetMetaType sets MetaType field to given value.


### GetSpendLimit

`func (o *LlmPolicy) GetSpendLimit() int32`

GetSpendLimit returns the SpendLimit field if non-nil, zero value otherwise.

### GetSpendLimitOk

`func (o *LlmPolicy) GetSpendLimitOk() (*int32, bool)`

GetSpendLimitOk returns a tuple with the SpendLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpendLimit

`func (o *LlmPolicy) SetSpendLimit(v int32)`

SetSpendLimit sets SpendLimit field to given value.


### SetSpendLimitNil

`func (o *LlmPolicy) SetSpendLimitNil(b bool)`

 SetSpendLimitNil sets the value for SpendLimit to be an explicit nil

### UnsetSpendLimit
`func (o *LlmPolicy) UnsetSpendLimit()`

UnsetSpendLimit ensures that no value is present for SpendLimit, not even an explicit nil
### GetModels

`func (o *LlmPolicy) GetModels() []string`

GetModels returns the Models field if non-nil, zero value otherwise.

### GetModelsOk

`func (o *LlmPolicy) GetModelsOk() (*[]string, bool)`

GetModelsOk returns a tuple with the Models field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModels

`func (o *LlmPolicy) SetModels(v []string)`

SetModels sets Models field to given value.


### GetCreatedAt

`func (o *LlmPolicy) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *LlmPolicy) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *LlmPolicy) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### SetCreatedAtNil

`func (o *LlmPolicy) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *LlmPolicy) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *LlmPolicy) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *LlmPolicy) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *LlmPolicy) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### SetUpdatedAtNil

`func (o *LlmPolicy) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *LlmPolicy) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetCurrentUsage

`func (o *LlmPolicy) GetCurrentUsage() LlmPolicyCurrentUsage`

GetCurrentUsage returns the CurrentUsage field if non-nil, zero value otherwise.

### GetCurrentUsageOk

`func (o *LlmPolicy) GetCurrentUsageOk() (*LlmPolicyCurrentUsage, bool)`

GetCurrentUsageOk returns a tuple with the CurrentUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentUsage

`func (o *LlmPolicy) SetCurrentUsage(v LlmPolicyCurrentUsage)`

SetCurrentUsage sets CurrentUsage field to given value.

### HasCurrentUsage

`func (o *LlmPolicy) HasCurrentUsage() bool`

HasCurrentUsage returns a boolean if a field has been set.

### SetCurrentUsageNil

`func (o *LlmPolicy) SetCurrentUsageNil(b bool)`

 SetCurrentUsageNil sets the value for CurrentUsage to be an explicit nil

### UnsetCurrentUsage
`func (o *LlmPolicy) UnsetCurrentUsage()`

UnsetCurrentUsage ensures that no value is present for CurrentUsage, not even an explicit nil
### GetLinks

`func (o *LlmPolicy) GetLinks() BackupRetentionPolicyLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *LlmPolicy) GetLinksOk() (*BackupRetentionPolicyLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *LlmPolicy) SetLinks(v BackupRetentionPolicyLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *LlmPolicy) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


