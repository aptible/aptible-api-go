# LlmGatewayKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Database ID of the LLM Gateway Key | 
**MetaType** | **string** |  | 
**Handle** | **string** | User-provided handle for the key (unique within the organization) | 
**Note** | **NullableString** | User-provided note | 
**SpendLimit** | **NullableString** | Spend limit in USD for this key (decimal, serialized as string). Null means the key has no limit of its own. Never exceeds the organization&#39;s spend_limit: lowering that lowers any key that was above it. | 
**OrganizationId** | **string** | The organization this key belongs to | 
**CreatedAt** | **NullableString** |  | 
**UpdatedAt** | **NullableString** |  | 
**GatewayUrl** | **NullableString** | Customer-facing LLM Gateway endpoint URL | 
**TokenPrefix** | **NullableString** | The first characters of the key&#39;s token, safe to display and enough to tell keys apart | 
**CreatedBy** | **map[string]interface{}** | User who created the key | 
**RevokedAt** | **NullableString** | Timestamp when the key was revoked | 
**RevokedBy** | **map[string]interface{}** | User who revoked the key | 
**Status** | **string** | Key status (ACTIVE or REVOKED) | 
**Token** | Pointer to **NullableString** | Secret API key value (only present on creation) | [optional] 
**Links** | Pointer to [**LlmGatewayKeyLinks**](LlmGatewayKeyLinks.md) |  | [optional] 

## Methods

### NewLlmGatewayKey

`func NewLlmGatewayKey(id int32, metaType string, handle string, note NullableString, spendLimit NullableString, organizationId string, createdAt NullableString, updatedAt NullableString, gatewayUrl NullableString, tokenPrefix NullableString, createdBy map[string]interface{}, revokedAt NullableString, revokedBy map[string]interface{}, status string, ) *LlmGatewayKey`

NewLlmGatewayKey instantiates a new LlmGatewayKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLlmGatewayKeyWithDefaults

`func NewLlmGatewayKeyWithDefaults() *LlmGatewayKey`

NewLlmGatewayKeyWithDefaults instantiates a new LlmGatewayKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LlmGatewayKey) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LlmGatewayKey) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LlmGatewayKey) SetId(v int32)`

SetId sets Id field to given value.


### GetMetaType

`func (o *LlmGatewayKey) GetMetaType() string`

GetMetaType returns the MetaType field if non-nil, zero value otherwise.

### GetMetaTypeOk

`func (o *LlmGatewayKey) GetMetaTypeOk() (*string, bool)`

GetMetaTypeOk returns a tuple with the MetaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaType

`func (o *LlmGatewayKey) SetMetaType(v string)`

SetMetaType sets MetaType field to given value.


### GetHandle

`func (o *LlmGatewayKey) GetHandle() string`

GetHandle returns the Handle field if non-nil, zero value otherwise.

### GetHandleOk

`func (o *LlmGatewayKey) GetHandleOk() (*string, bool)`

GetHandleOk returns a tuple with the Handle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandle

`func (o *LlmGatewayKey) SetHandle(v string)`

SetHandle sets Handle field to given value.


### GetNote

`func (o *LlmGatewayKey) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *LlmGatewayKey) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *LlmGatewayKey) SetNote(v string)`

SetNote sets Note field to given value.


### SetNoteNil

`func (o *LlmGatewayKey) SetNoteNil(b bool)`

 SetNoteNil sets the value for Note to be an explicit nil

### UnsetNote
`func (o *LlmGatewayKey) UnsetNote()`

UnsetNote ensures that no value is present for Note, not even an explicit nil
### GetSpendLimit

`func (o *LlmGatewayKey) GetSpendLimit() string`

GetSpendLimit returns the SpendLimit field if non-nil, zero value otherwise.

### GetSpendLimitOk

`func (o *LlmGatewayKey) GetSpendLimitOk() (*string, bool)`

GetSpendLimitOk returns a tuple with the SpendLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpendLimit

`func (o *LlmGatewayKey) SetSpendLimit(v string)`

SetSpendLimit sets SpendLimit field to given value.


### SetSpendLimitNil

`func (o *LlmGatewayKey) SetSpendLimitNil(b bool)`

 SetSpendLimitNil sets the value for SpendLimit to be an explicit nil

### UnsetSpendLimit
`func (o *LlmGatewayKey) UnsetSpendLimit()`

UnsetSpendLimit ensures that no value is present for SpendLimit, not even an explicit nil
### GetOrganizationId

`func (o *LlmGatewayKey) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *LlmGatewayKey) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *LlmGatewayKey) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.


### GetCreatedAt

`func (o *LlmGatewayKey) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *LlmGatewayKey) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *LlmGatewayKey) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### SetCreatedAtNil

`func (o *LlmGatewayKey) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *LlmGatewayKey) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *LlmGatewayKey) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *LlmGatewayKey) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *LlmGatewayKey) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### SetUpdatedAtNil

`func (o *LlmGatewayKey) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *LlmGatewayKey) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetGatewayUrl

`func (o *LlmGatewayKey) GetGatewayUrl() string`

GetGatewayUrl returns the GatewayUrl field if non-nil, zero value otherwise.

### GetGatewayUrlOk

`func (o *LlmGatewayKey) GetGatewayUrlOk() (*string, bool)`

GetGatewayUrlOk returns a tuple with the GatewayUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayUrl

`func (o *LlmGatewayKey) SetGatewayUrl(v string)`

SetGatewayUrl sets GatewayUrl field to given value.


### SetGatewayUrlNil

`func (o *LlmGatewayKey) SetGatewayUrlNil(b bool)`

 SetGatewayUrlNil sets the value for GatewayUrl to be an explicit nil

### UnsetGatewayUrl
`func (o *LlmGatewayKey) UnsetGatewayUrl()`

UnsetGatewayUrl ensures that no value is present for GatewayUrl, not even an explicit nil
### GetTokenPrefix

`func (o *LlmGatewayKey) GetTokenPrefix() string`

GetTokenPrefix returns the TokenPrefix field if non-nil, zero value otherwise.

### GetTokenPrefixOk

`func (o *LlmGatewayKey) GetTokenPrefixOk() (*string, bool)`

GetTokenPrefixOk returns a tuple with the TokenPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenPrefix

`func (o *LlmGatewayKey) SetTokenPrefix(v string)`

SetTokenPrefix sets TokenPrefix field to given value.


### SetTokenPrefixNil

`func (o *LlmGatewayKey) SetTokenPrefixNil(b bool)`

 SetTokenPrefixNil sets the value for TokenPrefix to be an explicit nil

### UnsetTokenPrefix
`func (o *LlmGatewayKey) UnsetTokenPrefix()`

UnsetTokenPrefix ensures that no value is present for TokenPrefix, not even an explicit nil
### GetCreatedBy

`func (o *LlmGatewayKey) GetCreatedBy() map[string]interface{}`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *LlmGatewayKey) GetCreatedByOk() (*map[string]interface{}, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *LlmGatewayKey) SetCreatedBy(v map[string]interface{})`

SetCreatedBy sets CreatedBy field to given value.


### SetCreatedByNil

`func (o *LlmGatewayKey) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *LlmGatewayKey) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetRevokedAt

`func (o *LlmGatewayKey) GetRevokedAt() string`

GetRevokedAt returns the RevokedAt field if non-nil, zero value otherwise.

### GetRevokedAtOk

`func (o *LlmGatewayKey) GetRevokedAtOk() (*string, bool)`

GetRevokedAtOk returns a tuple with the RevokedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevokedAt

`func (o *LlmGatewayKey) SetRevokedAt(v string)`

SetRevokedAt sets RevokedAt field to given value.


### SetRevokedAtNil

`func (o *LlmGatewayKey) SetRevokedAtNil(b bool)`

 SetRevokedAtNil sets the value for RevokedAt to be an explicit nil

### UnsetRevokedAt
`func (o *LlmGatewayKey) UnsetRevokedAt()`

UnsetRevokedAt ensures that no value is present for RevokedAt, not even an explicit nil
### GetRevokedBy

`func (o *LlmGatewayKey) GetRevokedBy() map[string]interface{}`

GetRevokedBy returns the RevokedBy field if non-nil, zero value otherwise.

### GetRevokedByOk

`func (o *LlmGatewayKey) GetRevokedByOk() (*map[string]interface{}, bool)`

GetRevokedByOk returns a tuple with the RevokedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevokedBy

`func (o *LlmGatewayKey) SetRevokedBy(v map[string]interface{})`

SetRevokedBy sets RevokedBy field to given value.


### SetRevokedByNil

`func (o *LlmGatewayKey) SetRevokedByNil(b bool)`

 SetRevokedByNil sets the value for RevokedBy to be an explicit nil

### UnsetRevokedBy
`func (o *LlmGatewayKey) UnsetRevokedBy()`

UnsetRevokedBy ensures that no value is present for RevokedBy, not even an explicit nil
### GetStatus

`func (o *LlmGatewayKey) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *LlmGatewayKey) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *LlmGatewayKey) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetToken

`func (o *LlmGatewayKey) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *LlmGatewayKey) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *LlmGatewayKey) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *LlmGatewayKey) HasToken() bool`

HasToken returns a boolean if a field has been set.

### SetTokenNil

`func (o *LlmGatewayKey) SetTokenNil(b bool)`

 SetTokenNil sets the value for Token to be an explicit nil

### UnsetToken
`func (o *LlmGatewayKey) UnsetToken()`

UnsetToken ensures that no value is present for Token, not even an explicit nil
### GetLinks

`func (o *LlmGatewayKey) GetLinks() LlmGatewayKeyLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *LlmGatewayKey) GetLinksOk() (*LlmGatewayKeyLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *LlmGatewayKey) SetLinks(v LlmGatewayKeyLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *LlmGatewayKey) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


