# UpdateLlmGatewayConfigurationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SpendLimit** | Pointer to **float32** | New monthly spend limit in USD. Lowering it also lowers the spend_limit of any of the organization&#39;s keys that are above the new limit, to match. | [optional] 
**Models** | Pointer to **[]string** | List of permitted model identifiers | [optional] 

## Methods

### NewUpdateLlmGatewayConfigurationRequest

`func NewUpdateLlmGatewayConfigurationRequest() *UpdateLlmGatewayConfigurationRequest`

NewUpdateLlmGatewayConfigurationRequest instantiates a new UpdateLlmGatewayConfigurationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateLlmGatewayConfigurationRequestWithDefaults

`func NewUpdateLlmGatewayConfigurationRequestWithDefaults() *UpdateLlmGatewayConfigurationRequest`

NewUpdateLlmGatewayConfigurationRequestWithDefaults instantiates a new UpdateLlmGatewayConfigurationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpendLimit

`func (o *UpdateLlmGatewayConfigurationRequest) GetSpendLimit() float32`

GetSpendLimit returns the SpendLimit field if non-nil, zero value otherwise.

### GetSpendLimitOk

`func (o *UpdateLlmGatewayConfigurationRequest) GetSpendLimitOk() (*float32, bool)`

GetSpendLimitOk returns a tuple with the SpendLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpendLimit

`func (o *UpdateLlmGatewayConfigurationRequest) SetSpendLimit(v float32)`

SetSpendLimit sets SpendLimit field to given value.

### HasSpendLimit

`func (o *UpdateLlmGatewayConfigurationRequest) HasSpendLimit() bool`

HasSpendLimit returns a boolean if a field has been set.

### GetModels

`func (o *UpdateLlmGatewayConfigurationRequest) GetModels() []string`

GetModels returns the Models field if non-nil, zero value otherwise.

### GetModelsOk

`func (o *UpdateLlmGatewayConfigurationRequest) GetModelsOk() (*[]string, bool)`

GetModelsOk returns a tuple with the Models field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModels

`func (o *UpdateLlmGatewayConfigurationRequest) SetModels(v []string)`

SetModels sets Models field to given value.

### HasModels

`func (o *UpdateLlmGatewayConfigurationRequest) HasModels() bool`

HasModels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


