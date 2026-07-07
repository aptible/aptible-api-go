# CreateOrUpdateLlmPolicyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Models** | Pointer to **[]string** | List of allowed model identifiers | [optional] 

## Methods

### NewCreateOrUpdateLlmPolicyRequest

`func NewCreateOrUpdateLlmPolicyRequest() *CreateOrUpdateLlmPolicyRequest`

NewCreateOrUpdateLlmPolicyRequest instantiates a new CreateOrUpdateLlmPolicyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrUpdateLlmPolicyRequestWithDefaults

`func NewCreateOrUpdateLlmPolicyRequestWithDefaults() *CreateOrUpdateLlmPolicyRequest`

NewCreateOrUpdateLlmPolicyRequestWithDefaults instantiates a new CreateOrUpdateLlmPolicyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModels

`func (o *CreateOrUpdateLlmPolicyRequest) GetModels() []string`

GetModels returns the Models field if non-nil, zero value otherwise.

### GetModelsOk

`func (o *CreateOrUpdateLlmPolicyRequest) GetModelsOk() (*[]string, bool)`

GetModelsOk returns a tuple with the Models field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModels

`func (o *CreateOrUpdateLlmPolicyRequest) SetModels(v []string)`

SetModels sets Models field to given value.

### HasModels

`func (o *CreateOrUpdateLlmPolicyRequest) HasModels() bool`

HasModels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


