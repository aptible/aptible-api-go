# Setting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**MetaType** | **string** |  | 
**Keys** | **map[string]interface{}** |  | 
**SensitiveKeys** | **map[string]interface{}** |  | 
**Links** | Pointer to [**ConfigurationLinks**](ConfigurationLinks.md) |  | [optional] 

## Methods

### NewSetting

`func NewSetting(id int32, metaType string, keys map[string]interface{}, sensitiveKeys map[string]interface{}, ) *Setting`

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


### GetKeys

`func (o *Setting) GetKeys() map[string]interface{}`

GetKeys returns the Keys field if non-nil, zero value otherwise.

### GetKeysOk

`func (o *Setting) GetKeysOk() (*map[string]interface{}, bool)`

GetKeysOk returns a tuple with the Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeys

`func (o *Setting) SetKeys(v map[string]interface{})`

SetKeys sets Keys field to given value.


### GetSensitiveKeys

`func (o *Setting) GetSensitiveKeys() map[string]interface{}`

GetSensitiveKeys returns the SensitiveKeys field if non-nil, zero value otherwise.

### GetSensitiveKeysOk

`func (o *Setting) GetSensitiveKeysOk() (*map[string]interface{}, bool)`

GetSensitiveKeysOk returns a tuple with the SensitiveKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSensitiveKeys

`func (o *Setting) SetSensitiveKeys(v map[string]interface{})`

SetSensitiveKeys sets SensitiveKeys field to given value.


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


