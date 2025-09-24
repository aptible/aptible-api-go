# BackupRetentionPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**MetaType** | **string** |  | 
**CreatedAt** | **string** |  | 
**Daily** | **int32** |  | 
**Monthly** | **int32** |  | 
**Yearly** | **int32** |  | 
**PitrDays** | **int32** |  | 
**MakeCopy** | **bool** |  | 
**KeepFinal** | **bool** |  | 
**Links** | Pointer to [**BackupRetentionPolicyLinks**](BackupRetentionPolicyLinks.md) |  | [optional] 

## Methods

### NewBackupRetentionPolicy

`func NewBackupRetentionPolicy(id int32, metaType string, createdAt string, daily int32, monthly int32, yearly int32, pitrDays int32, makeCopy bool, keepFinal bool, ) *BackupRetentionPolicy`

NewBackupRetentionPolicy instantiates a new BackupRetentionPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupRetentionPolicyWithDefaults

`func NewBackupRetentionPolicyWithDefaults() *BackupRetentionPolicy`

NewBackupRetentionPolicyWithDefaults instantiates a new BackupRetentionPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BackupRetentionPolicy) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BackupRetentionPolicy) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BackupRetentionPolicy) SetId(v int32)`

SetId sets Id field to given value.


### GetMetaType

`func (o *BackupRetentionPolicy) GetMetaType() string`

GetMetaType returns the MetaType field if non-nil, zero value otherwise.

### GetMetaTypeOk

`func (o *BackupRetentionPolicy) GetMetaTypeOk() (*string, bool)`

GetMetaTypeOk returns a tuple with the MetaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaType

`func (o *BackupRetentionPolicy) SetMetaType(v string)`

SetMetaType sets MetaType field to given value.


### GetCreatedAt

`func (o *BackupRetentionPolicy) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BackupRetentionPolicy) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BackupRetentionPolicy) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetDaily

`func (o *BackupRetentionPolicy) GetDaily() int32`

GetDaily returns the Daily field if non-nil, zero value otherwise.

### GetDailyOk

`func (o *BackupRetentionPolicy) GetDailyOk() (*int32, bool)`

GetDailyOk returns a tuple with the Daily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaily

`func (o *BackupRetentionPolicy) SetDaily(v int32)`

SetDaily sets Daily field to given value.


### GetMonthly

`func (o *BackupRetentionPolicy) GetMonthly() int32`

GetMonthly returns the Monthly field if non-nil, zero value otherwise.

### GetMonthlyOk

`func (o *BackupRetentionPolicy) GetMonthlyOk() (*int32, bool)`

GetMonthlyOk returns a tuple with the Monthly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthly

`func (o *BackupRetentionPolicy) SetMonthly(v int32)`

SetMonthly sets Monthly field to given value.


### GetYearly

`func (o *BackupRetentionPolicy) GetYearly() int32`

GetYearly returns the Yearly field if non-nil, zero value otherwise.

### GetYearlyOk

`func (o *BackupRetentionPolicy) GetYearlyOk() (*int32, bool)`

GetYearlyOk returns a tuple with the Yearly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYearly

`func (o *BackupRetentionPolicy) SetYearly(v int32)`

SetYearly sets Yearly field to given value.


### GetPitrDays

`func (o *BackupRetentionPolicy) GetPitrDays() int32`

GetPitrDays returns the PitrDays field if non-nil, zero value otherwise.

### GetPitrDaysOk

`func (o *BackupRetentionPolicy) GetPitrDaysOk() (*int32, bool)`

GetPitrDaysOk returns a tuple with the PitrDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPitrDays

`func (o *BackupRetentionPolicy) SetPitrDays(v int32)`

SetPitrDays sets PitrDays field to given value.


### GetMakeCopy

`func (o *BackupRetentionPolicy) GetMakeCopy() bool`

GetMakeCopy returns the MakeCopy field if non-nil, zero value otherwise.

### GetMakeCopyOk

`func (o *BackupRetentionPolicy) GetMakeCopyOk() (*bool, bool)`

GetMakeCopyOk returns a tuple with the MakeCopy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMakeCopy

`func (o *BackupRetentionPolicy) SetMakeCopy(v bool)`

SetMakeCopy sets MakeCopy field to given value.


### GetKeepFinal

`func (o *BackupRetentionPolicy) GetKeepFinal() bool`

GetKeepFinal returns the KeepFinal field if non-nil, zero value otherwise.

### GetKeepFinalOk

`func (o *BackupRetentionPolicy) GetKeepFinalOk() (*bool, bool)`

GetKeepFinalOk returns a tuple with the KeepFinal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepFinal

`func (o *BackupRetentionPolicy) SetKeepFinal(v bool)`

SetKeepFinal sets KeepFinal field to given value.


### GetLinks

`func (o *BackupRetentionPolicy) GetLinks() BackupRetentionPolicyLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *BackupRetentionPolicy) GetLinksOk() (*BackupRetentionPolicyLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *BackupRetentionPolicy) SetLinks(v BackupRetentionPolicyLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *BackupRetentionPolicy) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


