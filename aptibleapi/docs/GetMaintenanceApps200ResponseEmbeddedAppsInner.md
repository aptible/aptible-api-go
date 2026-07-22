# GetMaintenanceApps200ResponseEmbeddedAppsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Handle** | Pointer to **string** |  | [optional] 
**MetaType** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**MaintenanceDeadline** | Pointer to [**[]time.Time**](time.Time.md) |  | [optional] 
**AccountId** | Pointer to **int32** |  | [optional] 
**Links** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewGetMaintenanceApps200ResponseEmbeddedAppsInner

`func NewGetMaintenanceApps200ResponseEmbeddedAppsInner() *GetMaintenanceApps200ResponseEmbeddedAppsInner`

NewGetMaintenanceApps200ResponseEmbeddedAppsInner instantiates a new GetMaintenanceApps200ResponseEmbeddedAppsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMaintenanceApps200ResponseEmbeddedAppsInnerWithDefaults

`func NewGetMaintenanceApps200ResponseEmbeddedAppsInnerWithDefaults() *GetMaintenanceApps200ResponseEmbeddedAppsInner`

NewGetMaintenanceApps200ResponseEmbeddedAppsInnerWithDefaults instantiates a new GetMaintenanceApps200ResponseEmbeddedAppsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetMaintenanceApps200ResponseEmbeddedAppsInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetMaintenanceApps200ResponseEmbeddedAppsInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetMaintenanceApps200ResponseEmbeddedAppsInner) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GetMaintenanceApps200ResponseEmbeddedAppsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHandle

`func (o *GetMaintenanceApps200ResponseEmbeddedAppsInner) GetHandle() string`

GetHandle returns the Handle field if non-nil, zero value otherwise.

### GetHandleOk

`func (o *GetMaintenanceApps200ResponseEmbeddedAppsInner) GetHandleOk() (*string, bool)`

GetHandleOk returns a tuple with the Handle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandle

`func (o *GetMaintenanceApps200ResponseEmbeddedAppsInner) SetHandle(v string)`

SetHandle sets Handle field to given value.

### HasHandle

`func (o *GetMaintenanceApps200ResponseEmbeddedAppsInner) HasHandle() bool`

HasHandle returns a boolean if a field has been set.

### GetMetaType

`func (o *GetMaintenanceApps200ResponseEmbeddedAppsInner) GetMetaType() string`

GetMetaType returns the MetaType field if non-nil, zero value otherwise.

### GetMetaTypeOk

`func (o *GetMaintenanceApps200ResponseEmbeddedAppsInner) GetMetaTypeOk() (*string, bool)`

GetMetaTypeOk returns a tuple with the MetaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaType

`func (o *GetMaintenanceApps200ResponseEmbeddedAppsInner) SetMetaType(v string)`

SetMetaType sets MetaType field to given value.

### HasMetaType

`func (o *GetMaintenanceApps200ResponseEmbeddedAppsInner) HasMetaType() bool`

HasMetaType returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GetMaintenanceApps200ResponseEmbeddedAppsInner) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetMaintenanceApps200ResponseEmbeddedAppsInner) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetMaintenanceApps200ResponseEmbeddedAppsInner) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GetMaintenanceApps200ResponseEmbeddedAppsInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GetMaintenanceApps200ResponseEmbeddedAppsInner) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GetMaintenanceApps200ResponseEmbeddedAppsInner) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GetMaintenanceApps200ResponseEmbeddedAppsInner) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GetMaintenanceApps200ResponseEmbeddedAppsInner) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetStatus

`func (o *GetMaintenanceApps200ResponseEmbeddedAppsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetMaintenanceApps200ResponseEmbeddedAppsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetMaintenanceApps200ResponseEmbeddedAppsInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetMaintenanceApps200ResponseEmbeddedAppsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMaintenanceDeadline

`func (o *GetMaintenanceApps200ResponseEmbeddedAppsInner) GetMaintenanceDeadline() []time.Time`

GetMaintenanceDeadline returns the MaintenanceDeadline field if non-nil, zero value otherwise.

### GetMaintenanceDeadlineOk

`func (o *GetMaintenanceApps200ResponseEmbeddedAppsInner) GetMaintenanceDeadlineOk() (*[]time.Time, bool)`

GetMaintenanceDeadlineOk returns a tuple with the MaintenanceDeadline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceDeadline

`func (o *GetMaintenanceApps200ResponseEmbeddedAppsInner) SetMaintenanceDeadline(v []time.Time)`

SetMaintenanceDeadline sets MaintenanceDeadline field to given value.

### HasMaintenanceDeadline

`func (o *GetMaintenanceApps200ResponseEmbeddedAppsInner) HasMaintenanceDeadline() bool`

HasMaintenanceDeadline returns a boolean if a field has been set.

### SetMaintenanceDeadlineNil

`func (o *GetMaintenanceApps200ResponseEmbeddedAppsInner) SetMaintenanceDeadlineNil(b bool)`

 SetMaintenanceDeadlineNil sets the value for MaintenanceDeadline to be an explicit nil

### UnsetMaintenanceDeadline
`func (o *GetMaintenanceApps200ResponseEmbeddedAppsInner) UnsetMaintenanceDeadline()`

UnsetMaintenanceDeadline ensures that no value is present for MaintenanceDeadline, not even an explicit nil
### GetAccountId

`func (o *GetMaintenanceApps200ResponseEmbeddedAppsInner) GetAccountId() int32`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *GetMaintenanceApps200ResponseEmbeddedAppsInner) GetAccountIdOk() (*int32, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *GetMaintenanceApps200ResponseEmbeddedAppsInner) SetAccountId(v int32)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *GetMaintenanceApps200ResponseEmbeddedAppsInner) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetLinks

`func (o *GetMaintenanceApps200ResponseEmbeddedAppsInner) GetLinks() map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetMaintenanceApps200ResponseEmbeddedAppsInner) GetLinksOk() (*map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetMaintenanceApps200ResponseEmbeddedAppsInner) SetLinks(v map[string]interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GetMaintenanceApps200ResponseEmbeddedAppsInner) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


