# DiskAttachment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**MetaType** | **string** |  | 
**MountPoint** | **string** |  | 
**Status** | **string** |  | 
**CreatedAt** | **string** |  | 
**UpdatedAt** | **string** |  | 
**DeletedAt** | **NullableString** |  | 
**Embedded** | [**DiskAttachmentEmbedded**](DiskAttachmentEmbedded.md) |  | 
**Links** | Pointer to [**DiskAttachmentLinks**](DiskAttachmentLinks.md) |  | [optional] 

## Methods

### NewDiskAttachment

`func NewDiskAttachment(id int32, metaType string, mountPoint string, status string, createdAt string, updatedAt string, deletedAt NullableString, embedded DiskAttachmentEmbedded, ) *DiskAttachment`

NewDiskAttachment instantiates a new DiskAttachment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiskAttachmentWithDefaults

`func NewDiskAttachmentWithDefaults() *DiskAttachment`

NewDiskAttachmentWithDefaults instantiates a new DiskAttachment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DiskAttachment) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DiskAttachment) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DiskAttachment) SetId(v int32)`

SetId sets Id field to given value.


### GetMetaType

`func (o *DiskAttachment) GetMetaType() string`

GetMetaType returns the MetaType field if non-nil, zero value otherwise.

### GetMetaTypeOk

`func (o *DiskAttachment) GetMetaTypeOk() (*string, bool)`

GetMetaTypeOk returns a tuple with the MetaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaType

`func (o *DiskAttachment) SetMetaType(v string)`

SetMetaType sets MetaType field to given value.


### GetMountPoint

`func (o *DiskAttachment) GetMountPoint() string`

GetMountPoint returns the MountPoint field if non-nil, zero value otherwise.

### GetMountPointOk

`func (o *DiskAttachment) GetMountPointOk() (*string, bool)`

GetMountPointOk returns a tuple with the MountPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountPoint

`func (o *DiskAttachment) SetMountPoint(v string)`

SetMountPoint sets MountPoint field to given value.


### GetStatus

`func (o *DiskAttachment) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DiskAttachment) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DiskAttachment) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetCreatedAt

`func (o *DiskAttachment) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DiskAttachment) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DiskAttachment) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *DiskAttachment) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DiskAttachment) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DiskAttachment) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetDeletedAt

`func (o *DiskAttachment) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *DiskAttachment) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *DiskAttachment) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.


### SetDeletedAtNil

`func (o *DiskAttachment) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *DiskAttachment) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
### GetEmbedded

`func (o *DiskAttachment) GetEmbedded() DiskAttachmentEmbedded`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *DiskAttachment) GetEmbeddedOk() (*DiskAttachmentEmbedded, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *DiskAttachment) SetEmbedded(v DiskAttachmentEmbedded)`

SetEmbedded sets Embedded field to given value.


### GetLinks

`func (o *DiskAttachment) GetLinks() DiskAttachmentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DiskAttachment) GetLinksOk() (*DiskAttachmentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DiskAttachment) SetLinks(v DiskAttachmentLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *DiskAttachment) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


