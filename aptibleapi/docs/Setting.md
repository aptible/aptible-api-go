# Setting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**MetaType** | **string** |  | 
**Settings** | **map[string]interface{}** |  | 
**SensitiveSettings** | **map[string]interface{}** |  | 
**Links** | Pointer to [**ConfigurationLinks**](ConfigurationLinks.md) |  | [optional] 

## Methods

### NewSetting

`func NewSetting(id int32, metaType string, settings map[string]interface{}, sensitiveSettings map[string]interface{}, ) *Setting`

NewSetting instantiates a new Setting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSettingWithDefaults

`func NewSettingWithDefaults() *Setting`

NewSettingWithDefaults instantiates a new Setting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Setting) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Setting) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Setting) SetId(v int32)`

SetId sets Id field to given value.


### GetMetaType

`func (o *Setting) GetMetaType() string`

GetMetaType returns the MetaType field if non-nil, zero value otherwise.

### GetMetaTypeOk

`func (o *Setting) GetMetaTypeOk() (*string, bool)`

GetMetaTypeOk returns a tuple with the MetaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaType

`func (o *Setting) SetMetaType(v string)`

SetMetaType sets MetaType field to given value.


### GetSettings

`func (o *Setting) GetSettings() map[string]interface{}`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *Setting) GetSettingsOk() (*map[string]interface{}, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *Setting) SetSettings(v map[string]interface{})`

SetSettings sets Settings field to given value.


### GetSensitiveSettings

`func (o *Setting) GetSensitiveSettings() map[string]interface{}`

GetSensitiveSettings returns the SensitiveSettings field if non-nil, zero value otherwise.

### GetSensitiveSettingsOk

`func (o *Setting) GetSensitiveSettingsOk() (*map[string]interface{}, bool)`

GetSensitiveSettingsOk returns a tuple with the SensitiveSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSensitiveSettings

`func (o *Setting) SetSensitiveSettings(v map[string]interface{})`

SetSensitiveSettings sets SensitiveSettings field to given value.


### GetLinks

`func (o *Setting) GetLinks() ConfigurationLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Setting) GetLinksOk() (*ConfigurationLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Setting) SetLinks(v ConfigurationLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Setting) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


