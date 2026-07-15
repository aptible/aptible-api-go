# GetLlmKeyUsage200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Spend** | Pointer to **float32** | Total spend in USD | [optional] 
**InputTokens** | Pointer to **int32** | Total input (prompt) tokens | [optional] 
**OutputTokens** | Pointer to **int32** | Total output (completion) tokens | [optional] 
**SuccessfulRequests** | Pointer to **int32** | Total successful API requests | [optional] 
**FailedRequests** | Pointer to **int32** | Total failed API requests | [optional] 

## Methods

### NewGetLlmKeyUsage200Response

`func NewGetLlmKeyUsage200Response() *GetLlmKeyUsage200Response`

NewGetLlmKeyUsage200Response instantiates a new GetLlmKeyUsage200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetLlmKeyUsage200ResponseWithDefaults

`func NewGetLlmKeyUsage200ResponseWithDefaults() *GetLlmKeyUsage200Response`

NewGetLlmKeyUsage200ResponseWithDefaults instantiates a new GetLlmKeyUsage200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpend

`func (o *GetLlmKeyUsage200Response) GetSpend() float32`

GetSpend returns the Spend field if non-nil, zero value otherwise.

### GetSpendOk

`func (o *GetLlmKeyUsage200Response) GetSpendOk() (*float32, bool)`

GetSpendOk returns a tuple with the Spend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpend

`func (o *GetLlmKeyUsage200Response) SetSpend(v float32)`

SetSpend sets Spend field to given value.

### HasSpend

`func (o *GetLlmKeyUsage200Response) HasSpend() bool`

HasSpend returns a boolean if a field has been set.

### GetInputTokens

`func (o *GetLlmKeyUsage200Response) GetInputTokens() int32`

GetInputTokens returns the InputTokens field if non-nil, zero value otherwise.

### GetInputTokensOk

`func (o *GetLlmKeyUsage200Response) GetInputTokensOk() (*int32, bool)`

GetInputTokensOk returns a tuple with the InputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputTokens

`func (o *GetLlmKeyUsage200Response) SetInputTokens(v int32)`

SetInputTokens sets InputTokens field to given value.

### HasInputTokens

`func (o *GetLlmKeyUsage200Response) HasInputTokens() bool`

HasInputTokens returns a boolean if a field has been set.

### GetOutputTokens

`func (o *GetLlmKeyUsage200Response) GetOutputTokens() int32`

GetOutputTokens returns the OutputTokens field if non-nil, zero value otherwise.

### GetOutputTokensOk

`func (o *GetLlmKeyUsage200Response) GetOutputTokensOk() (*int32, bool)`

GetOutputTokensOk returns a tuple with the OutputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputTokens

`func (o *GetLlmKeyUsage200Response) SetOutputTokens(v int32)`

SetOutputTokens sets OutputTokens field to given value.

### HasOutputTokens

`func (o *GetLlmKeyUsage200Response) HasOutputTokens() bool`

HasOutputTokens returns a boolean if a field has been set.

### GetSuccessfulRequests

`func (o *GetLlmKeyUsage200Response) GetSuccessfulRequests() int32`

GetSuccessfulRequests returns the SuccessfulRequests field if non-nil, zero value otherwise.

### GetSuccessfulRequestsOk

`func (o *GetLlmKeyUsage200Response) GetSuccessfulRequestsOk() (*int32, bool)`

GetSuccessfulRequestsOk returns a tuple with the SuccessfulRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessfulRequests

`func (o *GetLlmKeyUsage200Response) SetSuccessfulRequests(v int32)`

SetSuccessfulRequests sets SuccessfulRequests field to given value.

### HasSuccessfulRequests

`func (o *GetLlmKeyUsage200Response) HasSuccessfulRequests() bool`

HasSuccessfulRequests returns a boolean if a field has been set.

### GetFailedRequests

`func (o *GetLlmKeyUsage200Response) GetFailedRequests() int32`

GetFailedRequests returns the FailedRequests field if non-nil, zero value otherwise.

### GetFailedRequestsOk

`func (o *GetLlmKeyUsage200Response) GetFailedRequestsOk() (*int32, bool)`

GetFailedRequestsOk returns a tuple with the FailedRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedRequests

`func (o *GetLlmKeyUsage200Response) SetFailedRequests(v int32)`

SetFailedRequests sets FailedRequests field to given value.

### HasFailedRequests

`func (o *GetLlmKeyUsage200Response) HasFailedRequests() bool`

HasFailedRequests returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


