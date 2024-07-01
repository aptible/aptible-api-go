# Disk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Handle** | **string** |  | 
**EbsVolumeId** | **NullableString** |  | 
**EbsVolumeType** | **NullableString** |  | 
**Filesystem** | **string** |  | 
**Passphrase** | Pointer to **string** |  | [optional] 
**KeyBytes** | **int32** |  | 
**Size** | **int32** |  | 
**ProvisionedIops** | **int32** |  | 
**Host** | **NullableString** |  | 
**Device** | **NullableString** |  | 
**Attached** | **bool** |  | 
**MetaType** | **string** |  | 
**CreatedAt** | **string** |  | 
**UpdatedAt** | **string** |  | 
**Ec2InstanceId** | **NullableString** |  | 
**BaselineIops** | **NullableInt32** |  | 
**AvailabilityZone** | **NullableString** |  | 
**CurrentKmsArn** | **NullableString** |  | 
**Links** | Pointer to [**DiskLinks**](DiskLinks.md) |  | [optional] 

## Methods

### NewDisk

`func NewDisk(id int32, handle string, ebsVolumeId NullableString, ebsVolumeType NullableString, filesystem string, keyBytes int32, size int32, provisionedIops int32, host NullableString, device NullableString, attached bool, metaType string, createdAt string, updatedAt string, ec2InstanceId NullableString, baselineIops NullableInt32, availabilityZone NullableString, currentKmsArn NullableString, ) *Disk`

NewDisk instantiates a new Disk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiskWithDefaults

`func NewDiskWithDefaults() *Disk`

NewDiskWithDefaults instantiates a new Disk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Disk) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Disk) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Disk) SetId(v int32)`

SetId sets Id field to given value.


### GetHandle

`func (o *Disk) GetHandle() string`

GetHandle returns the Handle field if non-nil, zero value otherwise.

### GetHandleOk

`func (o *Disk) GetHandleOk() (*string, bool)`

GetHandleOk returns a tuple with the Handle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandle

`func (o *Disk) SetHandle(v string)`

SetHandle sets Handle field to given value.


### GetEbsVolumeId

`func (o *Disk) GetEbsVolumeId() string`

GetEbsVolumeId returns the EbsVolumeId field if non-nil, zero value otherwise.

### GetEbsVolumeIdOk

`func (o *Disk) GetEbsVolumeIdOk() (*string, bool)`

GetEbsVolumeIdOk returns a tuple with the EbsVolumeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEbsVolumeId

`func (o *Disk) SetEbsVolumeId(v string)`

SetEbsVolumeId sets EbsVolumeId field to given value.


### SetEbsVolumeIdNil

`func (o *Disk) SetEbsVolumeIdNil(b bool)`

 SetEbsVolumeIdNil sets the value for EbsVolumeId to be an explicit nil

### UnsetEbsVolumeId
`func (o *Disk) UnsetEbsVolumeId()`

UnsetEbsVolumeId ensures that no value is present for EbsVolumeId, not even an explicit nil
### GetEbsVolumeType

`func (o *Disk) GetEbsVolumeType() string`

GetEbsVolumeType returns the EbsVolumeType field if non-nil, zero value otherwise.

### GetEbsVolumeTypeOk

`func (o *Disk) GetEbsVolumeTypeOk() (*string, bool)`

GetEbsVolumeTypeOk returns a tuple with the EbsVolumeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEbsVolumeType

`func (o *Disk) SetEbsVolumeType(v string)`

SetEbsVolumeType sets EbsVolumeType field to given value.


### SetEbsVolumeTypeNil

`func (o *Disk) SetEbsVolumeTypeNil(b bool)`

 SetEbsVolumeTypeNil sets the value for EbsVolumeType to be an explicit nil

### UnsetEbsVolumeType
`func (o *Disk) UnsetEbsVolumeType()`

UnsetEbsVolumeType ensures that no value is present for EbsVolumeType, not even an explicit nil
### GetFilesystem

`func (o *Disk) GetFilesystem() string`

GetFilesystem returns the Filesystem field if non-nil, zero value otherwise.

### GetFilesystemOk

`func (o *Disk) GetFilesystemOk() (*string, bool)`

GetFilesystemOk returns a tuple with the Filesystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesystem

`func (o *Disk) SetFilesystem(v string)`

SetFilesystem sets Filesystem field to given value.


### GetPassphrase

`func (o *Disk) GetPassphrase() string`

GetPassphrase returns the Passphrase field if non-nil, zero value otherwise.

### GetPassphraseOk

`func (o *Disk) GetPassphraseOk() (*string, bool)`

GetPassphraseOk returns a tuple with the Passphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassphrase

`func (o *Disk) SetPassphrase(v string)`

SetPassphrase sets Passphrase field to given value.

### HasPassphrase

`func (o *Disk) HasPassphrase() bool`

HasPassphrase returns a boolean if a field has been set.

### GetKeyBytes

`func (o *Disk) GetKeyBytes() int32`

GetKeyBytes returns the KeyBytes field if non-nil, zero value otherwise.

### GetKeyBytesOk

`func (o *Disk) GetKeyBytesOk() (*int32, bool)`

GetKeyBytesOk returns a tuple with the KeyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyBytes

`func (o *Disk) SetKeyBytes(v int32)`

SetKeyBytes sets KeyBytes field to given value.


### GetSize

`func (o *Disk) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *Disk) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *Disk) SetSize(v int32)`

SetSize sets Size field to given value.


### GetProvisionedIops

`func (o *Disk) GetProvisionedIops() int32`

GetProvisionedIops returns the ProvisionedIops field if non-nil, zero value otherwise.

### GetProvisionedIopsOk

`func (o *Disk) GetProvisionedIopsOk() (*int32, bool)`

GetProvisionedIopsOk returns a tuple with the ProvisionedIops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionedIops

`func (o *Disk) SetProvisionedIops(v int32)`

SetProvisionedIops sets ProvisionedIops field to given value.


### GetHost

`func (o *Disk) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *Disk) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *Disk) SetHost(v string)`

SetHost sets Host field to given value.


### SetHostNil

`func (o *Disk) SetHostNil(b bool)`

 SetHostNil sets the value for Host to be an explicit nil

### UnsetHost
`func (o *Disk) UnsetHost()`

UnsetHost ensures that no value is present for Host, not even an explicit nil
### GetDevice

`func (o *Disk) GetDevice() string`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *Disk) GetDeviceOk() (*string, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *Disk) SetDevice(v string)`

SetDevice sets Device field to given value.


### SetDeviceNil

`func (o *Disk) SetDeviceNil(b bool)`

 SetDeviceNil sets the value for Device to be an explicit nil

### UnsetDevice
`func (o *Disk) UnsetDevice()`

UnsetDevice ensures that no value is present for Device, not even an explicit nil
### GetAttached

`func (o *Disk) GetAttached() bool`

GetAttached returns the Attached field if non-nil, zero value otherwise.

### GetAttachedOk

`func (o *Disk) GetAttachedOk() (*bool, bool)`

GetAttachedOk returns a tuple with the Attached field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttached

`func (o *Disk) SetAttached(v bool)`

SetAttached sets Attached field to given value.


### GetMetaType

`func (o *Disk) GetMetaType() string`

GetMetaType returns the MetaType field if non-nil, zero value otherwise.

### GetMetaTypeOk

`func (o *Disk) GetMetaTypeOk() (*string, bool)`

GetMetaTypeOk returns a tuple with the MetaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaType

`func (o *Disk) SetMetaType(v string)`

SetMetaType sets MetaType field to given value.


### GetCreatedAt

`func (o *Disk) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Disk) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Disk) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Disk) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Disk) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Disk) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetEc2InstanceId

`func (o *Disk) GetEc2InstanceId() string`

GetEc2InstanceId returns the Ec2InstanceId field if non-nil, zero value otherwise.

### GetEc2InstanceIdOk

`func (o *Disk) GetEc2InstanceIdOk() (*string, bool)`

GetEc2InstanceIdOk returns a tuple with the Ec2InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEc2InstanceId

`func (o *Disk) SetEc2InstanceId(v string)`

SetEc2InstanceId sets Ec2InstanceId field to given value.


### SetEc2InstanceIdNil

`func (o *Disk) SetEc2InstanceIdNil(b bool)`

 SetEc2InstanceIdNil sets the value for Ec2InstanceId to be an explicit nil

### UnsetEc2InstanceId
`func (o *Disk) UnsetEc2InstanceId()`

UnsetEc2InstanceId ensures that no value is present for Ec2InstanceId, not even an explicit nil
### GetBaselineIops

`func (o *Disk) GetBaselineIops() int32`

GetBaselineIops returns the BaselineIops field if non-nil, zero value otherwise.

### GetBaselineIopsOk

`func (o *Disk) GetBaselineIopsOk() (*int32, bool)`

GetBaselineIopsOk returns a tuple with the BaselineIops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaselineIops

`func (o *Disk) SetBaselineIops(v int32)`

SetBaselineIops sets BaselineIops field to given value.


### SetBaselineIopsNil

`func (o *Disk) SetBaselineIopsNil(b bool)`

 SetBaselineIopsNil sets the value for BaselineIops to be an explicit nil

### UnsetBaselineIops
`func (o *Disk) UnsetBaselineIops()`

UnsetBaselineIops ensures that no value is present for BaselineIops, not even an explicit nil
### GetAvailabilityZone

`func (o *Disk) GetAvailabilityZone() string`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *Disk) GetAvailabilityZoneOk() (*string, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *Disk) SetAvailabilityZone(v string)`

SetAvailabilityZone sets AvailabilityZone field to given value.


### SetAvailabilityZoneNil

`func (o *Disk) SetAvailabilityZoneNil(b bool)`

 SetAvailabilityZoneNil sets the value for AvailabilityZone to be an explicit nil

### UnsetAvailabilityZone
`func (o *Disk) UnsetAvailabilityZone()`

UnsetAvailabilityZone ensures that no value is present for AvailabilityZone, not even an explicit nil
### GetCurrentKmsArn

`func (o *Disk) GetCurrentKmsArn() string`

GetCurrentKmsArn returns the CurrentKmsArn field if non-nil, zero value otherwise.

### GetCurrentKmsArnOk

`func (o *Disk) GetCurrentKmsArnOk() (*string, bool)`

GetCurrentKmsArnOk returns a tuple with the CurrentKmsArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentKmsArn

`func (o *Disk) SetCurrentKmsArn(v string)`

SetCurrentKmsArn sets CurrentKmsArn field to given value.


### SetCurrentKmsArnNil

`func (o *Disk) SetCurrentKmsArnNil(b bool)`

 SetCurrentKmsArnNil sets the value for CurrentKmsArn to be an explicit nil

### UnsetCurrentKmsArn
`func (o *Disk) UnsetCurrentKmsArn()`

UnsetCurrentKmsArn ensures that no value is present for CurrentKmsArn, not even an explicit nil
### GetLinks

`func (o *Disk) GetLinks() DiskLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Disk) GetLinksOk() (*DiskLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Disk) SetLinks(v DiskLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Disk) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


