# UpdateServiceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NaiveHealthCheck** | Pointer to **bool** |  | [optional] 
**ForceZeroDowntime** | Pointer to **bool** |  | [optional] 

## Methods

### NewUpdateServiceRequest

`func NewUpdateServiceRequest() *UpdateServiceRequest`

NewUpdateServiceRequest instantiates a new UpdateServiceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateServiceRequestWithDefaults

`func NewUpdateServiceRequestWithDefaults() *UpdateServiceRequest`

NewUpdateServiceRequestWithDefaults instantiates a new UpdateServiceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNaiveHealthCheck

`func (o *UpdateServiceRequest) GetNaiveHealthCheck() bool`

GetNaiveHealthCheck returns the NaiveHealthCheck field if non-nil, zero value otherwise.

### GetNaiveHealthCheckOk

`func (o *UpdateServiceRequest) GetNaiveHealthCheckOk() (*bool, bool)`

GetNaiveHealthCheckOk returns a tuple with the NaiveHealthCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNaiveHealthCheck

`func (o *UpdateServiceRequest) SetNaiveHealthCheck(v bool)`

SetNaiveHealthCheck sets NaiveHealthCheck field to given value.

### HasNaiveHealthCheck

`func (o *UpdateServiceRequest) HasNaiveHealthCheck() bool`

HasNaiveHealthCheck returns a boolean if a field has been set.

### GetForceZeroDowntime

`func (o *UpdateServiceRequest) GetForceZeroDowntime() bool`

GetForceZeroDowntime returns the ForceZeroDowntime field if non-nil, zero value otherwise.

### GetForceZeroDowntimeOk

`func (o *UpdateServiceRequest) GetForceZeroDowntimeOk() (*bool, bool)`

GetForceZeroDowntimeOk returns a tuple with the ForceZeroDowntime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceZeroDowntime

`func (o *UpdateServiceRequest) SetForceZeroDowntime(v bool)`

SetForceZeroDowntime sets ForceZeroDowntime field to given value.

### HasForceZeroDowntime

`func (o *UpdateServiceRequest) HasForceZeroDowntime() bool`

HasForceZeroDowntime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


