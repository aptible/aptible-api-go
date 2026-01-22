# CreateLlmKeyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Handle** | **string** | Unique handle for the token within the account | 
**Note** | Pointer to **string** | URL-safe base64 encoded note (max 256 chars decoded) | [optional] 

## Methods

### NewCreateLlmKeyRequest

`func NewCreateLlmKeyRequest(handle string, ) *CreateLlmKeyRequest`

NewCreateLlmKeyRequest instantiates a new CreateLlmKeyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateLlmKeyRequestWithDefaults

`func NewCreateLlmKeyRequestWithDefaults() *CreateLlmKeyRequest`

NewCreateLlmKeyRequestWithDefaults instantiates a new CreateLlmKeyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHandle

`func (o *CreateLlmKeyRequest) GetHandle() string`

GetHandle returns the Handle field if non-nil, zero value otherwise.

### GetHandleOk

`func (o *CreateLlmKeyRequest) GetHandleOk() (*string, bool)`

GetHandleOk returns a tuple with the Handle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandle

`func (o *CreateLlmKeyRequest) SetHandle(v string)`

SetHandle sets Handle field to given value.


### GetNote

`func (o *CreateLlmKeyRequest) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *CreateLlmKeyRequest) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *CreateLlmKeyRequest) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *CreateLlmKeyRequest) HasNote() bool`

HasNote returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


