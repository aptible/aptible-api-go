# LlmKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Database ID of the LLM Key | 
**MetaType** | **string** |  | 
**Handle** | **string** | User-provided handle for the token (unique within account) | 
**Note** | **NullableString** | User-provided note (URL-safe base64 encoded) | 
**CreatedAt** | **NullableString** |  | 
**UpdatedAt** | **NullableString** |  | 
**GatewayUrl** | **NullableString** | Customer-facing LLM Gateway endpoint URL | 
**CreatedBy** | **map[string]interface{}** | User who created the token | 
**RevokedAt** | **NullableString** | Timestamp when token was revoked | 
**RevokedBy** | **map[string]interface{}** | User who revoked the token | 
**Status** | **string** | Token status (ACTIVE or REVOKED) | 
**Token** | Pointer to **NullableString** | Secret API key value (only present on creation) | [optional] 
**Links** | Pointer to [**BackupRetentionPolicyLinks**](BackupRetentionPolicyLinks.md) |  | [optional] 

## Methods

### NewLlmKey

`func NewLlmKey(id int32, metaType string, handle string, note NullableString, createdAt NullableString, updatedAt NullableString, gatewayUrl NullableString, createdBy map[string]interface{}, revokedAt NullableString, revokedBy map[string]interface{}, status string, ) *LlmKey`

NewLlmKey instantiates a new LlmKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLlmKeyWithDefaults

`func NewLlmKeyWithDefaults() *LlmKey`

NewLlmKeyWithDefaults instantiates a new LlmKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LlmKey) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LlmKey) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LlmKey) SetId(v int32)`

SetId sets Id field to given value.


### GetMetaType

`func (o *LlmKey) GetMetaType() string`

GetMetaType returns the MetaType field if non-nil, zero value otherwise.

### GetMetaTypeOk

`func (o *LlmKey) GetMetaTypeOk() (*string, bool)`

GetMetaTypeOk returns a tuple with the MetaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaType

`func (o *LlmKey) SetMetaType(v string)`

SetMetaType sets MetaType field to given value.


### GetHandle

`func (o *LlmKey) GetHandle() string`

GetHandle returns the Handle field if non-nil, zero value otherwise.

### GetHandleOk

`func (o *LlmKey) GetHandleOk() (*string, bool)`

GetHandleOk returns a tuple with the Handle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandle

`func (o *LlmKey) SetHandle(v string)`

SetHandle sets Handle field to given value.


### GetNote

`func (o *LlmKey) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *LlmKey) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *LlmKey) SetNote(v string)`

SetNote sets Note field to given value.


### SetNoteNil

`func (o *LlmKey) SetNoteNil(b bool)`

 SetNoteNil sets the value for Note to be an explicit nil

### UnsetNote
`func (o *LlmKey) UnsetNote()`

UnsetNote ensures that no value is present for Note, not even an explicit nil
### GetCreatedAt

`func (o *LlmKey) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *LlmKey) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *LlmKey) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### SetCreatedAtNil

`func (o *LlmKey) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *LlmKey) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *LlmKey) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *LlmKey) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *LlmKey) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### SetUpdatedAtNil

`func (o *LlmKey) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *LlmKey) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetGatewayUrl

`func (o *LlmKey) GetGatewayUrl() string`

GetGatewayUrl returns the GatewayUrl field if non-nil, zero value otherwise.

### GetGatewayUrlOk

`func (o *LlmKey) GetGatewayUrlOk() (*string, bool)`

GetGatewayUrlOk returns a tuple with the GatewayUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayUrl

`func (o *LlmKey) SetGatewayUrl(v string)`

SetGatewayUrl sets GatewayUrl field to given value.


### SetGatewayUrlNil

`func (o *LlmKey) SetGatewayUrlNil(b bool)`

 SetGatewayUrlNil sets the value for GatewayUrl to be an explicit nil

### UnsetGatewayUrl
`func (o *LlmKey) UnsetGatewayUrl()`

UnsetGatewayUrl ensures that no value is present for GatewayUrl, not even an explicit nil
### GetCreatedBy

`func (o *LlmKey) GetCreatedBy() map[string]interface{}`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *LlmKey) GetCreatedByOk() (*map[string]interface{}, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *LlmKey) SetCreatedBy(v map[string]interface{})`

SetCreatedBy sets CreatedBy field to given value.


### SetCreatedByNil

`func (o *LlmKey) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *LlmKey) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetRevokedAt

`func (o *LlmKey) GetRevokedAt() string`

GetRevokedAt returns the RevokedAt field if non-nil, zero value otherwise.

### GetRevokedAtOk

`func (o *LlmKey) GetRevokedAtOk() (*string, bool)`

GetRevokedAtOk returns a tuple with the RevokedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevokedAt

`func (o *LlmKey) SetRevokedAt(v string)`

SetRevokedAt sets RevokedAt field to given value.


### SetRevokedAtNil

`func (o *LlmKey) SetRevokedAtNil(b bool)`

 SetRevokedAtNil sets the value for RevokedAt to be an explicit nil

### UnsetRevokedAt
`func (o *LlmKey) UnsetRevokedAt()`

UnsetRevokedAt ensures that no value is present for RevokedAt, not even an explicit nil
### GetRevokedBy

`func (o *LlmKey) GetRevokedBy() map[string]interface{}`

GetRevokedBy returns the RevokedBy field if non-nil, zero value otherwise.

### GetRevokedByOk

`func (o *LlmKey) GetRevokedByOk() (*map[string]interface{}, bool)`

GetRevokedByOk returns a tuple with the RevokedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevokedBy

`func (o *LlmKey) SetRevokedBy(v map[string]interface{})`

SetRevokedBy sets RevokedBy field to given value.


### SetRevokedByNil

`func (o *LlmKey) SetRevokedByNil(b bool)`

 SetRevokedByNil sets the value for RevokedBy to be an explicit nil

### UnsetRevokedBy
`func (o *LlmKey) UnsetRevokedBy()`

UnsetRevokedBy ensures that no value is present for RevokedBy, not even an explicit nil
### GetStatus

`func (o *LlmKey) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *LlmKey) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *LlmKey) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetToken

`func (o *LlmKey) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *LlmKey) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *LlmKey) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *LlmKey) HasToken() bool`

HasToken returns a boolean if a field has been set.

### SetTokenNil

`func (o *LlmKey) SetTokenNil(b bool)`

 SetTokenNil sets the value for Token to be an explicit nil

### UnsetToken
`func (o *LlmKey) UnsetToken()`

UnsetToken ensures that no value is present for Token, not even an explicit nil
### GetLinks

`func (o *LlmKey) GetLinks() BackupRetentionPolicyLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *LlmKey) GetLinksOk() (*BackupRetentionPolicyLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *LlmKey) SetLinks(v BackupRetentionPolicyLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *LlmKey) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


