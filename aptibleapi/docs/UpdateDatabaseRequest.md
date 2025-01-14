# UpdateDatabaseRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentConfiguration** | Pointer to **int32** |  | [optional] 
**Handle** | Pointer to **string** |  | [optional] 
**InitialDiskSize** | Pointer to **int32** |  | [optional] 
**InitialContinerSize** | Pointer to **int32** |  | [optional] 
**DatabaseImageId** | Pointer to **int32** |  | [optional] 
**CurrentKmsArn** | Pointer to **int32** |  | [optional] 
**EnableBackups** | Pointer to **bool** |  | [optional] 

## Methods

### NewUpdateDatabaseRequest

`func NewUpdateDatabaseRequest() *UpdateDatabaseRequest`

NewUpdateDatabaseRequest instantiates a new UpdateDatabaseRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDatabaseRequestWithDefaults

`func NewUpdateDatabaseRequestWithDefaults() *UpdateDatabaseRequest`

NewUpdateDatabaseRequestWithDefaults instantiates a new UpdateDatabaseRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentConfiguration

`func (o *UpdateDatabaseRequest) GetCurrentConfiguration() int32`

GetCurrentConfiguration returns the CurrentConfiguration field if non-nil, zero value otherwise.

### GetCurrentConfigurationOk

`func (o *UpdateDatabaseRequest) GetCurrentConfigurationOk() (*int32, bool)`

GetCurrentConfigurationOk returns a tuple with the CurrentConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentConfiguration

`func (o *UpdateDatabaseRequest) SetCurrentConfiguration(v int32)`

SetCurrentConfiguration sets CurrentConfiguration field to given value.

### HasCurrentConfiguration

`func (o *UpdateDatabaseRequest) HasCurrentConfiguration() bool`

HasCurrentConfiguration returns a boolean if a field has been set.

### GetHandle

`func (o *UpdateDatabaseRequest) GetHandle() string`

GetHandle returns the Handle field if non-nil, zero value otherwise.

### GetHandleOk

`func (o *UpdateDatabaseRequest) GetHandleOk() (*string, bool)`

GetHandleOk returns a tuple with the Handle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandle

`func (o *UpdateDatabaseRequest) SetHandle(v string)`

SetHandle sets Handle field to given value.

### HasHandle

`func (o *UpdateDatabaseRequest) HasHandle() bool`

HasHandle returns a boolean if a field has been set.

### GetInitialDiskSize

`func (o *UpdateDatabaseRequest) GetInitialDiskSize() int32`

GetInitialDiskSize returns the InitialDiskSize field if non-nil, zero value otherwise.

### GetInitialDiskSizeOk

`func (o *UpdateDatabaseRequest) GetInitialDiskSizeOk() (*int32, bool)`

GetInitialDiskSizeOk returns a tuple with the InitialDiskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialDiskSize

`func (o *UpdateDatabaseRequest) SetInitialDiskSize(v int32)`

SetInitialDiskSize sets InitialDiskSize field to given value.

### HasInitialDiskSize

`func (o *UpdateDatabaseRequest) HasInitialDiskSize() bool`

HasInitialDiskSize returns a boolean if a field has been set.

### GetInitialContinerSize

`func (o *UpdateDatabaseRequest) GetInitialContinerSize() int32`

GetInitialContinerSize returns the InitialContinerSize field if non-nil, zero value otherwise.

### GetInitialContinerSizeOk

`func (o *UpdateDatabaseRequest) GetInitialContinerSizeOk() (*int32, bool)`

GetInitialContinerSizeOk returns a tuple with the InitialContinerSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialContinerSize

`func (o *UpdateDatabaseRequest) SetInitialContinerSize(v int32)`

SetInitialContinerSize sets InitialContinerSize field to given value.

### HasInitialContinerSize

`func (o *UpdateDatabaseRequest) HasInitialContinerSize() bool`

HasInitialContinerSize returns a boolean if a field has been set.

### GetDatabaseImageId

`func (o *UpdateDatabaseRequest) GetDatabaseImageId() int32`

GetDatabaseImageId returns the DatabaseImageId field if non-nil, zero value otherwise.

### GetDatabaseImageIdOk

`func (o *UpdateDatabaseRequest) GetDatabaseImageIdOk() (*int32, bool)`

GetDatabaseImageIdOk returns a tuple with the DatabaseImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseImageId

`func (o *UpdateDatabaseRequest) SetDatabaseImageId(v int32)`

SetDatabaseImageId sets DatabaseImageId field to given value.

### HasDatabaseImageId

`func (o *UpdateDatabaseRequest) HasDatabaseImageId() bool`

HasDatabaseImageId returns a boolean if a field has been set.

### GetCurrentKmsArn

`func (o *UpdateDatabaseRequest) GetCurrentKmsArn() int32`

GetCurrentKmsArn returns the CurrentKmsArn field if non-nil, zero value otherwise.

### GetCurrentKmsArnOk

`func (o *UpdateDatabaseRequest) GetCurrentKmsArnOk() (*int32, bool)`

GetCurrentKmsArnOk returns a tuple with the CurrentKmsArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentKmsArn

`func (o *UpdateDatabaseRequest) SetCurrentKmsArn(v int32)`

SetCurrentKmsArn sets CurrentKmsArn field to given value.

### HasCurrentKmsArn

`func (o *UpdateDatabaseRequest) HasCurrentKmsArn() bool`

HasCurrentKmsArn returns a boolean if a field has been set.

### GetEnableBackups

`func (o *UpdateDatabaseRequest) GetEnableBackups() bool`

GetEnableBackups returns the EnableBackups field if non-nil, zero value otherwise.

### GetEnableBackupsOk

`func (o *UpdateDatabaseRequest) GetEnableBackupsOk() (*bool, bool)`

GetEnableBackupsOk returns a tuple with the EnableBackups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableBackups

`func (o *UpdateDatabaseRequest) SetEnableBackups(v bool)`

SetEnableBackups sets EnableBackups field to given value.

### HasEnableBackups

`func (o *UpdateDatabaseRequest) HasEnableBackups() bool`

HasEnableBackups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


