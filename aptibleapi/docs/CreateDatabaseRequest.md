# CreateDatabaseRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Handle** | **string** |  | 
**Type** | Pointer to **string** |  | [optional] 
**InitialDiskSize** | Pointer to **int32** |  | [optional] 
**InitialContainerSize** | Pointer to **int32** |  | [optional] 
**InitializeFrom** | Pointer to **string** |  | [optional] 
**DatabaseImage** | Pointer to **int32** | Alternate name for &#x60;database_image_id&#x60; | [optional] 
**DatabaseImageId** | **int32** |  | 
**CurrentKmsArn** | Pointer to **int32** |  | [optional] 

## Methods

### NewCreateDatabaseRequest

`func NewCreateDatabaseRequest(handle string, databaseImageId int32, ) *CreateDatabaseRequest`

NewCreateDatabaseRequest instantiates a new CreateDatabaseRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDatabaseRequestWithDefaults

`func NewCreateDatabaseRequestWithDefaults() *CreateDatabaseRequest`

NewCreateDatabaseRequestWithDefaults instantiates a new CreateDatabaseRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHandle

`func (o *CreateDatabaseRequest) GetHandle() string`

GetHandle returns the Handle field if non-nil, zero value otherwise.

### GetHandleOk

`func (o *CreateDatabaseRequest) GetHandleOk() (*string, bool)`

GetHandleOk returns a tuple with the Handle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandle

`func (o *CreateDatabaseRequest) SetHandle(v string)`

SetHandle sets Handle field to given value.


### GetType

`func (o *CreateDatabaseRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateDatabaseRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateDatabaseRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CreateDatabaseRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetInitialDiskSize

`func (o *CreateDatabaseRequest) GetInitialDiskSize() int32`

GetInitialDiskSize returns the InitialDiskSize field if non-nil, zero value otherwise.

### GetInitialDiskSizeOk

`func (o *CreateDatabaseRequest) GetInitialDiskSizeOk() (*int32, bool)`

GetInitialDiskSizeOk returns a tuple with the InitialDiskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialDiskSize

`func (o *CreateDatabaseRequest) SetInitialDiskSize(v int32)`

SetInitialDiskSize sets InitialDiskSize field to given value.

### HasInitialDiskSize

`func (o *CreateDatabaseRequest) HasInitialDiskSize() bool`

HasInitialDiskSize returns a boolean if a field has been set.

### GetInitialContainerSize

`func (o *CreateDatabaseRequest) GetInitialContainerSize() int32`

GetInitialContainerSize returns the InitialContainerSize field if non-nil, zero value otherwise.

### GetInitialContainerSizeOk

`func (o *CreateDatabaseRequest) GetInitialContainerSizeOk() (*int32, bool)`

GetInitialContainerSizeOk returns a tuple with the InitialContainerSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialContainerSize

`func (o *CreateDatabaseRequest) SetInitialContainerSize(v int32)`

SetInitialContainerSize sets InitialContainerSize field to given value.

### HasInitialContainerSize

`func (o *CreateDatabaseRequest) HasInitialContainerSize() bool`

HasInitialContainerSize returns a boolean if a field has been set.

### GetInitializeFrom

`func (o *CreateDatabaseRequest) GetInitializeFrom() string`

GetInitializeFrom returns the InitializeFrom field if non-nil, zero value otherwise.

### GetInitializeFromOk

`func (o *CreateDatabaseRequest) GetInitializeFromOk() (*string, bool)`

GetInitializeFromOk returns a tuple with the InitializeFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitializeFrom

`func (o *CreateDatabaseRequest) SetInitializeFrom(v string)`

SetInitializeFrom sets InitializeFrom field to given value.

### HasInitializeFrom

`func (o *CreateDatabaseRequest) HasInitializeFrom() bool`

HasInitializeFrom returns a boolean if a field has been set.

### GetDatabaseImage

`func (o *CreateDatabaseRequest) GetDatabaseImage() int32`

GetDatabaseImage returns the DatabaseImage field if non-nil, zero value otherwise.

### GetDatabaseImageOk

`func (o *CreateDatabaseRequest) GetDatabaseImageOk() (*int32, bool)`

GetDatabaseImageOk returns a tuple with the DatabaseImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseImage

`func (o *CreateDatabaseRequest) SetDatabaseImage(v int32)`

SetDatabaseImage sets DatabaseImage field to given value.

### HasDatabaseImage

`func (o *CreateDatabaseRequest) HasDatabaseImage() bool`

HasDatabaseImage returns a boolean if a field has been set.

### GetDatabaseImageId

`func (o *CreateDatabaseRequest) GetDatabaseImageId() int32`

GetDatabaseImageId returns the DatabaseImageId field if non-nil, zero value otherwise.

### GetDatabaseImageIdOk

`func (o *CreateDatabaseRequest) GetDatabaseImageIdOk() (*int32, bool)`

GetDatabaseImageIdOk returns a tuple with the DatabaseImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseImageId

`func (o *CreateDatabaseRequest) SetDatabaseImageId(v int32)`

SetDatabaseImageId sets DatabaseImageId field to given value.


### GetCurrentKmsArn

`func (o *CreateDatabaseRequest) GetCurrentKmsArn() int32`

GetCurrentKmsArn returns the CurrentKmsArn field if non-nil, zero value otherwise.

### GetCurrentKmsArnOk

`func (o *CreateDatabaseRequest) GetCurrentKmsArnOk() (*int32, bool)`

GetCurrentKmsArnOk returns a tuple with the CurrentKmsArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentKmsArn

`func (o *CreateDatabaseRequest) SetCurrentKmsArn(v int32)`

SetCurrentKmsArn sets CurrentKmsArn field to given value.

### HasCurrentKmsArn

`func (o *CreateDatabaseRequest) HasCurrentKmsArn() bool`

HasCurrentKmsArn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


