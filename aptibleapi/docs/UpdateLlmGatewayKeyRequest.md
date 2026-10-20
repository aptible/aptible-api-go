# UpdateLlmGatewayKeyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Handle** | Pointer to **string** | New handle for the key, unique within the organization | [optional] 
**SpendLimit** | Pointer to **NullableFloat32** | New spend limit in USD for this key, or null to remove it. Must not exceed the organization&#39;s spend_limit. | [optional] 

## Methods

### NewUpdateLlmGatewayKeyRequest

`func NewUpdateLlmGatewayKeyRequest() *UpdateLlmGatewayKeyRequest`

NewUpdateLlmGatewayKeyRequest instantiates a new UpdateLlmGatewayKeyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateLlmGatewayKeyRequestWithDefaults

`func NewUpdateLlmGatewayKeyRequestWithDefaults() *UpdateLlmGatewayKeyRequest`

NewUpdateLlmGatewayKeyRequestWithDefaults instantiates a new UpdateLlmGatewayKeyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHandle

`func (o *UpdateLlmGatewayKeyRequest) GetHandle() string`

GetHandle returns the Handle field if non-nil, zero value otherwise.

### GetHandleOk

`func (o *UpdateLlmGatewayKeyRequest) GetHandleOk() (*string, bool)`

GetHandleOk returns a tuple with the Handle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandle

`func (o *UpdateLlmGatewayKeyRequest) SetHandle(v string)`

SetHandle sets Handle field to given value.

### HasHandle

`func (o *UpdateLlmGatewayKeyRequest) HasHandle() bool`

HasHandle returns a boolean if a field has been set.

### GetSpendLimit

`func (o *UpdateLlmGatewayKeyRequest) GetSpendLimit() float32`

GetSpendLimit returns the SpendLimit field if non-nil, zero value otherwise.

### GetSpendLimitOk

`func (o *UpdateLlmGatewayKeyRequest) GetSpendLimitOk() (*float32, bool)`

GetSpendLimitOk returns a tuple with the SpendLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpendLimit

`func (o *UpdateLlmGatewayKeyRequest) SetSpendLimit(v float32)`

SetSpendLimit sets SpendLimit field to given value.

### HasSpendLimit

`func (o *UpdateLlmGatewayKeyRequest) HasSpendLimit() bool`

HasSpendLimit returns a boolean if a field has been set.

### SetSpendLimitNil

`func (o *UpdateLlmGatewayKeyRequest) SetSpendLimitNil(b bool)`

 SetSpendLimitNil sets the value for SpendLimit to be an explicit nil

### UnsetSpendLimit
`func (o *UpdateLlmGatewayKeyRequest) UnsetSpendLimit()`

UnsetSpendLimit ensures that no value is present for SpendLimit, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


