# MetricDrain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**MetaType** | **string** |  | 
**Handle** | **string** |  | 
**CreatedAt** | **string** |  | 
**UpdatedAt** | **string** |  | 
**Status** | **string** |  | 
**DrainType** | **string** |  | 
**DrainConfiguration** | Pointer to [**MetricDrainDrainConfiguration**](MetricDrainDrainConfiguration.md) |  | [optional] 
**AggregatorCaCertificate** | **NullableString** |  | 
**AggregatorCaPrivateKeyBlob** | **NullableString** |  | 
**AggregatorHost** | **NullableString** |  | 
**AggregatorPortMapping** | **[][]int32** |  | 
**AggregatorInstanceId** | **NullableString** |  | 
**AggregatorDockerName** | **NullableString** |  | 
**AggregatorAllocation** | **[]string** |  | 
**Links** | Pointer to [**MetricDrainLinks**](MetricDrainLinks.md) |  | [optional] 

## Methods

### NewMetricDrain

`func NewMetricDrain(id int32, metaType string, handle string, createdAt string, updatedAt string, status string, drainType string, aggregatorCaCertificate NullableString, aggregatorCaPrivateKeyBlob NullableString, aggregatorHost NullableString, aggregatorPortMapping [][]int32, aggregatorInstanceId NullableString, aggregatorDockerName NullableString, aggregatorAllocation []string, ) *MetricDrain`

NewMetricDrain instantiates a new MetricDrain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricDrainWithDefaults

`func NewMetricDrainWithDefaults() *MetricDrain`

NewMetricDrainWithDefaults instantiates a new MetricDrain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MetricDrain) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MetricDrain) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MetricDrain) SetId(v int32)`

SetId sets Id field to given value.


### GetMetaType

`func (o *MetricDrain) GetMetaType() string`

GetMetaType returns the MetaType field if non-nil, zero value otherwise.

### GetMetaTypeOk

`func (o *MetricDrain) GetMetaTypeOk() (*string, bool)`

GetMetaTypeOk returns a tuple with the MetaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaType

`func (o *MetricDrain) SetMetaType(v string)`

SetMetaType sets MetaType field to given value.


### GetHandle

`func (o *MetricDrain) GetHandle() string`

GetHandle returns the Handle field if non-nil, zero value otherwise.

### GetHandleOk

`func (o *MetricDrain) GetHandleOk() (*string, bool)`

GetHandleOk returns a tuple with the Handle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandle

`func (o *MetricDrain) SetHandle(v string)`

SetHandle sets Handle field to given value.


### GetCreatedAt

`func (o *MetricDrain) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MetricDrain) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MetricDrain) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *MetricDrain) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *MetricDrain) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *MetricDrain) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetStatus

`func (o *MetricDrain) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MetricDrain) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MetricDrain) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetDrainType

`func (o *MetricDrain) GetDrainType() string`

GetDrainType returns the DrainType field if non-nil, zero value otherwise.

### GetDrainTypeOk

`func (o *MetricDrain) GetDrainTypeOk() (*string, bool)`

GetDrainTypeOk returns a tuple with the DrainType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrainType

`func (o *MetricDrain) SetDrainType(v string)`

SetDrainType sets DrainType field to given value.


### GetDrainConfiguration

`func (o *MetricDrain) GetDrainConfiguration() MetricDrainDrainConfiguration`

GetDrainConfiguration returns the DrainConfiguration field if non-nil, zero value otherwise.

### GetDrainConfigurationOk

`func (o *MetricDrain) GetDrainConfigurationOk() (*MetricDrainDrainConfiguration, bool)`

GetDrainConfigurationOk returns a tuple with the DrainConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrainConfiguration

`func (o *MetricDrain) SetDrainConfiguration(v MetricDrainDrainConfiguration)`

SetDrainConfiguration sets DrainConfiguration field to given value.

### HasDrainConfiguration

`func (o *MetricDrain) HasDrainConfiguration() bool`

HasDrainConfiguration returns a boolean if a field has been set.

### GetAggregatorCaCertificate

`func (o *MetricDrain) GetAggregatorCaCertificate() string`

GetAggregatorCaCertificate returns the AggregatorCaCertificate field if non-nil, zero value otherwise.

### GetAggregatorCaCertificateOk

`func (o *MetricDrain) GetAggregatorCaCertificateOk() (*string, bool)`

GetAggregatorCaCertificateOk returns a tuple with the AggregatorCaCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregatorCaCertificate

`func (o *MetricDrain) SetAggregatorCaCertificate(v string)`

SetAggregatorCaCertificate sets AggregatorCaCertificate field to given value.


### SetAggregatorCaCertificateNil

`func (o *MetricDrain) SetAggregatorCaCertificateNil(b bool)`

 SetAggregatorCaCertificateNil sets the value for AggregatorCaCertificate to be an explicit nil

### UnsetAggregatorCaCertificate
`func (o *MetricDrain) UnsetAggregatorCaCertificate()`

UnsetAggregatorCaCertificate ensures that no value is present for AggregatorCaCertificate, not even an explicit nil
### GetAggregatorCaPrivateKeyBlob

`func (o *MetricDrain) GetAggregatorCaPrivateKeyBlob() string`

GetAggregatorCaPrivateKeyBlob returns the AggregatorCaPrivateKeyBlob field if non-nil, zero value otherwise.

### GetAggregatorCaPrivateKeyBlobOk

`func (o *MetricDrain) GetAggregatorCaPrivateKeyBlobOk() (*string, bool)`

GetAggregatorCaPrivateKeyBlobOk returns a tuple with the AggregatorCaPrivateKeyBlob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregatorCaPrivateKeyBlob

`func (o *MetricDrain) SetAggregatorCaPrivateKeyBlob(v string)`

SetAggregatorCaPrivateKeyBlob sets AggregatorCaPrivateKeyBlob field to given value.


### SetAggregatorCaPrivateKeyBlobNil

`func (o *MetricDrain) SetAggregatorCaPrivateKeyBlobNil(b bool)`

 SetAggregatorCaPrivateKeyBlobNil sets the value for AggregatorCaPrivateKeyBlob to be an explicit nil

### UnsetAggregatorCaPrivateKeyBlob
`func (o *MetricDrain) UnsetAggregatorCaPrivateKeyBlob()`

UnsetAggregatorCaPrivateKeyBlob ensures that no value is present for AggregatorCaPrivateKeyBlob, not even an explicit nil
### GetAggregatorHost

`func (o *MetricDrain) GetAggregatorHost() string`

GetAggregatorHost returns the AggregatorHost field if non-nil, zero value otherwise.

### GetAggregatorHostOk

`func (o *MetricDrain) GetAggregatorHostOk() (*string, bool)`

GetAggregatorHostOk returns a tuple with the AggregatorHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregatorHost

`func (o *MetricDrain) SetAggregatorHost(v string)`

SetAggregatorHost sets AggregatorHost field to given value.


### SetAggregatorHostNil

`func (o *MetricDrain) SetAggregatorHostNil(b bool)`

 SetAggregatorHostNil sets the value for AggregatorHost to be an explicit nil

### UnsetAggregatorHost
`func (o *MetricDrain) UnsetAggregatorHost()`

UnsetAggregatorHost ensures that no value is present for AggregatorHost, not even an explicit nil
### GetAggregatorPortMapping

`func (o *MetricDrain) GetAggregatorPortMapping() [][]int32`

GetAggregatorPortMapping returns the AggregatorPortMapping field if non-nil, zero value otherwise.

### GetAggregatorPortMappingOk

`func (o *MetricDrain) GetAggregatorPortMappingOk() (*[][]int32, bool)`

GetAggregatorPortMappingOk returns a tuple with the AggregatorPortMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregatorPortMapping

`func (o *MetricDrain) SetAggregatorPortMapping(v [][]int32)`

SetAggregatorPortMapping sets AggregatorPortMapping field to given value.


### GetAggregatorInstanceId

`func (o *MetricDrain) GetAggregatorInstanceId() string`

GetAggregatorInstanceId returns the AggregatorInstanceId field if non-nil, zero value otherwise.

### GetAggregatorInstanceIdOk

`func (o *MetricDrain) GetAggregatorInstanceIdOk() (*string, bool)`

GetAggregatorInstanceIdOk returns a tuple with the AggregatorInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregatorInstanceId

`func (o *MetricDrain) SetAggregatorInstanceId(v string)`

SetAggregatorInstanceId sets AggregatorInstanceId field to given value.


### SetAggregatorInstanceIdNil

`func (o *MetricDrain) SetAggregatorInstanceIdNil(b bool)`

 SetAggregatorInstanceIdNil sets the value for AggregatorInstanceId to be an explicit nil

### UnsetAggregatorInstanceId
`func (o *MetricDrain) UnsetAggregatorInstanceId()`

UnsetAggregatorInstanceId ensures that no value is present for AggregatorInstanceId, not even an explicit nil
### GetAggregatorDockerName

`func (o *MetricDrain) GetAggregatorDockerName() string`

GetAggregatorDockerName returns the AggregatorDockerName field if non-nil, zero value otherwise.

### GetAggregatorDockerNameOk

`func (o *MetricDrain) GetAggregatorDockerNameOk() (*string, bool)`

GetAggregatorDockerNameOk returns a tuple with the AggregatorDockerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregatorDockerName

`func (o *MetricDrain) SetAggregatorDockerName(v string)`

SetAggregatorDockerName sets AggregatorDockerName field to given value.


### SetAggregatorDockerNameNil

`func (o *MetricDrain) SetAggregatorDockerNameNil(b bool)`

 SetAggregatorDockerNameNil sets the value for AggregatorDockerName to be an explicit nil

### UnsetAggregatorDockerName
`func (o *MetricDrain) UnsetAggregatorDockerName()`

UnsetAggregatorDockerName ensures that no value is present for AggregatorDockerName, not even an explicit nil
### GetAggregatorAllocation

`func (o *MetricDrain) GetAggregatorAllocation() []string`

GetAggregatorAllocation returns the AggregatorAllocation field if non-nil, zero value otherwise.

### GetAggregatorAllocationOk

`func (o *MetricDrain) GetAggregatorAllocationOk() (*[]string, bool)`

GetAggregatorAllocationOk returns a tuple with the AggregatorAllocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregatorAllocation

`func (o *MetricDrain) SetAggregatorAllocation(v []string)`

SetAggregatorAllocation sets AggregatorAllocation field to given value.


### GetLinks

`func (o *MetricDrain) GetLinks() MetricDrainLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MetricDrain) GetLinksOk() (*MetricDrainLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MetricDrain) SetLinks(v MetricDrainLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *MetricDrain) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


