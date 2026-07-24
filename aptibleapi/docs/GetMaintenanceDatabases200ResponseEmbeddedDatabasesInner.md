# GetMaintenanceDatabases200ResponseEmbeddedDatabasesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Handle** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **NullableString** |  | [optional] 
**Provisioned** | Pointer to **bool** |  | [optional] 
**MetaType** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**DeletedAt** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**MaintenanceDeadline** | Pointer to [**[]time.Time**](time.Time.md) |  | [optional] 
**Links** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewGetMaintenanceDatabases200ResponseEmbeddedDatabasesInner

`func NewGetMaintenanceDatabases200ResponseEmbeddedDatabasesInner() *GetMaintenanceDatabases200ResponseEmbeddedDatabasesInner`

NewGetMaintenanceDatabases200ResponseEmbeddedDatabasesInner instantiates a new GetMaintenanceDatabases200ResponseEmbeddedDatabasesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMaintenanceDatabases200ResponseEmbeddedDatabasesInnerWithDefaults

`func NewGetMaintenanceDatabases200ResponseEmbeddedDatabasesInnerWithDefaults() *GetMaintenanceDatabases200ResponseEmbeddedDatabasesInner`

NewGetMaintenanceDatabases200ResponseEmbeddedDatabasesInnerWithDefaults instantiates a new GetMaintenanceDatabases200ResponseEmbeddedDatabasesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetMaintenanceDatabases200ResponseEmbeddedDatabasesInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetMaintenanceDatabases200ResponseEmbeddedDatabasesInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetMaintenanceDatabases200ResponseEmbeddedDatabasesInner) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GetMaintenanceDatabases200ResponseEmbeddedDatabasesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHandle

`func (o *GetMaintenanceDatabases200ResponseEmbeddedDatabasesInner) GetHandle() string`

GetHandle returns the Handle field if non-nil, zero value otherwise.

### GetHandleOk

`func (o *GetMaintenanceDatabases200ResponseEmbeddedDatabasesInner) GetHandleOk() (*string, bool)`

GetHandleOk returns a tuple with the Handle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandle

`func (o *GetMaintenanceDatabases200ResponseEmbeddedDatabasesInner) SetHandle(v string)`

SetHandle sets Handle field to given value.

### HasHandle

`func (o *GetMaintenanceDatabases200ResponseEmbeddedDatabasesInner) HasHandle() bool`

HasHandle returns a boolean if a field has been set.

### GetType

`func (o *GetMaintenanceDatabases200ResponseEmbeddedDatabasesInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetMaintenanceDatabases200ResponseEmbeddedDatabasesInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetMaintenanceDatabases200ResponseEmbeddedDatabasesInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetMaintenanceDatabases200ResponseEmbeddedDatabasesInner) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *GetMaintenanceDatabases200ResponseEmbeddedDatabasesInner) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *GetMaintenanceDatabases200ResponseEmbeddedDatabasesInner) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetProvisioned

`func (o *GetMaintenanceDatabases200ResponseEmbeddedDatabasesInner) GetProvisioned() bool`

GetProvisioned returns the Provisioned field if non-nil, zero value otherwise.

### GetProvisionedOk

`func (o *GetMaintenanceDatabases200ResponseEmbeddedDatabasesInner) GetProvisionedOk() (*bool, bool)`

GetProvisionedOk returns a tuple with the Provisioned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioned

`func (o *GetMaintenanceDatabases200ResponseEmbeddedDatabasesInner) SetProvisioned(v bool)`

SetProvisioned sets Provisioned field to given value.

### HasProvisioned

`func (o *GetMaintenanceDatabases200ResponseEmbeddedDatabasesInner) HasProvisioned() bool`

HasProvisioned returns a boolean if a field has been set.

### GetMetaType

`func (o *GetMaintenanceDatabases200ResponseEmbeddedDatabasesInner) GetMetaType() string`

GetMetaType returns the MetaType field if non-nil, zero value otherwise.

### GetMetaTypeOk

`func (o *GetMaintenanceDatabases200ResponseEmbeddedDatabasesInner) GetMetaTypeOk() (*string, bool)`

GetMetaTypeOk returns a tuple with the MetaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaType

`func (o *GetMaintenanceDatabases200ResponseEmbeddedDatabasesInner) SetMetaType(v string)`

SetMetaType sets MetaType field to given value.

### HasMetaType

`func (o *GetMaintenanceDatabases200ResponseEmbeddedDatabasesInner) HasMetaType() bool`

HasMetaType returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GetMaintenanceDatabases200ResponseEmbeddedDatabasesInner) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetMaintenanceDatabases200ResponseEmbeddedDatabasesInner) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetMaintenanceDatabases200ResponseEmbeddedDatabasesInner) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GetMaintenanceDatabases200ResponseEmbeddedDatabasesInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GetMaintenanceDatabases200ResponseEmbeddedDatabasesInner) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GetMaintenanceDatabases200ResponseEmbeddedDatabasesInner) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GetMaintenanceDatabases200ResponseEmbeddedDatabasesInner) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GetMaintenanceDatabases200ResponseEmbeddedDatabasesInner) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *GetMaintenanceDatabases200ResponseEmbeddedDatabasesInner) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *GetMaintenanceDatabases200ResponseEmbeddedDatabasesInner) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *GetMaintenanceDatabases200ResponseEmbeddedDatabasesInner) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *GetMaintenanceDatabases200ResponseEmbeddedDatabasesInner) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### SetDeletedAtNil

`func (o *GetMaintenanceDatabases200ResponseEmbeddedDatabasesInner) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *GetMaintenanceDatabases200ResponseEmbeddedDatabasesInner) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
### GetStatus

`func (o *GetMaintenanceDatabases200ResponseEmbeddedDatabasesInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetMaintenanceDatabases200ResponseEmbeddedDatabasesInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetMaintenanceDatabases200ResponseEmbeddedDatabasesInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetMaintenanceDatabases200ResponseEmbeddedDatabasesInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMaintenanceDeadline

`func (o *GetMaintenanceDatabases200ResponseEmbeddedDatabasesInner) GetMaintenanceDeadline() []time.Time`

GetMaintenanceDeadline returns the MaintenanceDeadline field if non-nil, zero value otherwise.

### GetMaintenanceDeadlineOk

`func (o *GetMaintenanceDatabases200ResponseEmbeddedDatabasesInner) GetMaintenanceDeadlineOk() (*[]time.Time, bool)`

GetMaintenanceDeadlineOk returns a tuple with the MaintenanceDeadline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceDeadline

`func (o *GetMaintenanceDatabases200ResponseEmbeddedDatabasesInner) SetMaintenanceDeadline(v []time.Time)`

SetMaintenanceDeadline sets MaintenanceDeadline field to given value.

### HasMaintenanceDeadline

`func (o *GetMaintenanceDatabases200ResponseEmbeddedDatabasesInner) HasMaintenanceDeadline() bool`

HasMaintenanceDeadline returns a boolean if a field has been set.

### SetMaintenanceDeadlineNil

`func (o *GetMaintenanceDatabases200ResponseEmbeddedDatabasesInner) SetMaintenanceDeadlineNil(b bool)`

 SetMaintenanceDeadlineNil sets the value for MaintenanceDeadline to be an explicit nil

### UnsetMaintenanceDeadline
`func (o *GetMaintenanceDatabases200ResponseEmbeddedDatabasesInner) UnsetMaintenanceDeadline()`

UnsetMaintenanceDeadline ensures that no value is present for MaintenanceDeadline, not even an explicit nil
### GetLinks

`func (o *GetMaintenanceDatabases200ResponseEmbeddedDatabasesInner) GetLinks() map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetMaintenanceDatabases200ResponseEmbeddedDatabasesInner) GetLinksOk() (*map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetMaintenanceDatabases200ResponseEmbeddedDatabasesInner) SetLinks(v map[string]interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GetMaintenanceDatabases200ResponseEmbeddedDatabasesInner) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


