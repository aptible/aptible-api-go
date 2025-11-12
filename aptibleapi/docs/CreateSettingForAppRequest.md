# CreateSettingForAppRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Keys** | Pointer to **map[string]interface{}** | Mapping of non-sensitive key-value pairs | [optional] 
**SensitiveKeys** | Pointer to **map[string]interface{}** | Mapping of sensitive key-value pairs | [optional] 

## Methods

### NewCreateSettingForAppRequest

`func NewCreateSettingForAppRequest() *CreateSettingForAppRequest`

NewCreateSettingForAppRequest instantiates a new CreateSettingForAppRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSettingForAppRequestWithDefaults

`func NewCreateSettingForAppRequestWithDefaults() *CreateSettingForAppRequest`

NewCreateSettingForAppRequestWithDefaults instantiates a new CreateSettingForAppRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeys

`func (o *CreateSettingForAppRequest) GetKeys() map[string]interface{}`

GetKeys returns the Keys field if non-nil, zero value otherwise.

### GetKeysOk

`func (o *CreateSettingForAppRequest) GetKeysOk() (*map[string]interface{}, bool)`

GetKeysOk returns a tuple with the Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeys

`func (o *CreateSettingForAppRequest) SetKeys(v map[string]interface{})`

SetKeys sets Keys field to given value.

### HasKeys

`func (o *CreateSettingForAppRequest) HasKeys() bool`

HasKeys returns a boolean if a field has been set.

### GetSensitiveKeys

`func (o *CreateSettingForAppRequest) GetSensitiveKeys() map[string]interface{}`

GetSensitiveKeys returns the SensitiveKeys field if non-nil, zero value otherwise.

### GetSensitiveKeysOk

`func (o *CreateSettingForAppRequest) GetSensitiveKeysOk() (*map[string]interface{}, bool)`

GetSensitiveKeysOk returns a tuple with the SensitiveKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSensitiveKeys

`func (o *CreateSettingForAppRequest) SetSensitiveKeys(v map[string]interface{})`

SetSensitiveKeys sets SensitiveKeys field to given value.

### HasSensitiveKeys

`func (o *CreateSettingForAppRequest) HasSensitiveKeys() bool`

HasSensitiveKeys returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


