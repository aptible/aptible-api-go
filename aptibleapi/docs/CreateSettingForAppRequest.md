# CreateSettingForAppRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Settings** | Pointer to **map[string]interface{}** | Mapping of non-sensitive key-value pairs | [optional] 
**SensitiveSettings** | Pointer to **map[string]interface{}** | Mapping of sensitive key-value pairs | [optional] 

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

### GetSettings

`func (o *CreateSettingForAppRequest) GetSettings() map[string]interface{}`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *CreateSettingForAppRequest) GetSettingsOk() (*map[string]interface{}, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *CreateSettingForAppRequest) SetSettings(v map[string]interface{})`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *CreateSettingForAppRequest) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetSensitiveSettings

`func (o *CreateSettingForAppRequest) GetSensitiveSettings() map[string]interface{}`

GetSensitiveSettings returns the SensitiveSettings field if non-nil, zero value otherwise.

### GetSensitiveSettingsOk

`func (o *CreateSettingForAppRequest) GetSensitiveSettingsOk() (*map[string]interface{}, bool)`

GetSensitiveSettingsOk returns a tuple with the SensitiveSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSensitiveSettings

`func (o *CreateSettingForAppRequest) SetSensitiveSettings(v map[string]interface{})`

SetSensitiveSettings sets SensitiveSettings field to given value.

### HasSensitiveSettings

`func (o *CreateSettingForAppRequest) HasSensitiveSettings() bool`

HasSensitiveSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


