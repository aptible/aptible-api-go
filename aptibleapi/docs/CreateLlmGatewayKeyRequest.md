# CreateLlmGatewayKeyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Handle** | **string** | Handle for the key, unique within the organization | 
**Note** | Pointer to **string** | Note about the key, up to 256 characters | [optional] 
**SpendLimit** | Pointer to **NullableFloat32** | Spend limit in USD for this key. Must not exceed the organization&#39;s spend_limit. Omit it, or send null, for a key with no limit of its own. | [optional] 
**OrganizationId** | Pointer to **string** | The organization to create the key in. Required unless the caller belongs to exactly one organization, in which case it is inferred from the token. | [optional] 

## Methods

### NewCreateLlmGatewayKeyRequest

`func NewCreateLlmGatewayKeyRequest(handle string, ) *CreateLlmGatewayKeyRequest`

NewCreateLlmGatewayKeyRequest instantiates a new CreateLlmGatewayKeyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateLlmGatewayKeyRequestWithDefaults

`func NewCreateLlmGatewayKeyRequestWithDefaults() *CreateLlmGatewayKeyRequest`

NewCreateLlmGatewayKeyRequestWithDefaults instantiates a new CreateLlmGatewayKeyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHandle

`func (o *CreateLlmGatewayKeyRequest) GetHandle() string`

GetHandle returns the Handle field if non-nil, zero value otherwise.

### GetHandleOk

`func (o *CreateLlmGatewayKeyRequest) GetHandleOk() (*string, bool)`

GetHandleOk returns a tuple with the Handle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandle

`func (o *CreateLlmGatewayKeyRequest) SetHandle(v string)`

SetHandle sets Handle field to given value.


### GetNote

`func (o *CreateLlmGatewayKeyRequest) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *CreateLlmGatewayKeyRequest) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *CreateLlmGatewayKeyRequest) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *CreateLlmGatewayKeyRequest) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetSpendLimit

`func (o *CreateLlmGatewayKeyRequest) GetSpendLimit() float32`

GetSpendLimit returns the SpendLimit field if non-nil, zero value otherwise.

### GetSpendLimitOk

`func (o *CreateLlmGatewayKeyRequest) GetSpendLimitOk() (*float32, bool)`

GetSpendLimitOk returns a tuple with the SpendLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpendLimit

`func (o *CreateLlmGatewayKeyRequest) SetSpendLimit(v float32)`

SetSpendLimit sets SpendLimit field to given value.

### HasSpendLimit

`func (o *CreateLlmGatewayKeyRequest) HasSpendLimit() bool`

HasSpendLimit returns a boolean if a field has been set.

### SetSpendLimitNil

`func (o *CreateLlmGatewayKeyRequest) SetSpendLimitNil(b bool)`

 SetSpendLimitNil sets the value for SpendLimit to be an explicit nil

### UnsetSpendLimit
`func (o *CreateLlmGatewayKeyRequest) UnsetSpendLimit()`

UnsetSpendLimit ensures that no value is present for SpendLimit, not even an explicit nil
### GetOrganizationId

`func (o *CreateLlmGatewayKeyRequest) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *CreateLlmGatewayKeyRequest) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *CreateLlmGatewayKeyRequest) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *CreateLlmGatewayKeyRequest) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


