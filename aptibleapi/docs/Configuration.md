# Configuration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**MetaType** | **string** |  | 
**Env** | Pointer to **map[string]interface{}** |  | [optional] 
**Links** | Pointer to [**ConfigurationLinks**](ConfigurationLinks.md) |  | [optional] 

## Methods

### NewConfiguration

`func NewConfiguration(id int32, metaType string, ) *Configuration`

NewConfiguration instantiates a new Configuration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigurationWithDefaults

`func NewConfigurationWithDefaults() *Configuration`

NewConfigurationWithDefaults instantiates a new Configuration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Configuration) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Configuration) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Configuration) SetId(v int32)`

SetId sets Id field to given value.


### GetMetaType

`func (o *Configuration) GetMetaType() string`

GetMetaType returns the MetaType field if non-nil, zero value otherwise.

### GetMetaTypeOk

`func (o *Configuration) GetMetaTypeOk() (*string, bool)`

GetMetaTypeOk returns a tuple with the MetaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaType

`func (o *Configuration) SetMetaType(v string)`

SetMetaType sets MetaType field to given value.


### GetEnv

`func (o *Configuration) GetEnv() map[string]interface{}`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *Configuration) GetEnvOk() (*map[string]interface{}, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *Configuration) SetEnv(v map[string]interface{})`

SetEnv sets Env field to given value.

### HasEnv

`func (o *Configuration) HasEnv() bool`

HasEnv returns a boolean if a field has been set.

### GetLinks

`func (o *Configuration) GetLinks() ConfigurationLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Configuration) GetLinksOk() (*ConfigurationLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Configuration) SetLinks(v ConfigurationLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Configuration) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


