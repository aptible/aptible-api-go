# UpdateMaintenanceWindowForAppRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DayOfWeek** | Pointer to **NullableInt32** |  | [optional] 
**StartTimeUtc** | Pointer to **string** |  | [optional] 
**DurationHours** | Pointer to **int32** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewUpdateMaintenanceWindowForAppRequest

`func NewUpdateMaintenanceWindowForAppRequest() *UpdateMaintenanceWindowForAppRequest`

NewUpdateMaintenanceWindowForAppRequest instantiates a new UpdateMaintenanceWindowForAppRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateMaintenanceWindowForAppRequestWithDefaults

`func NewUpdateMaintenanceWindowForAppRequestWithDefaults() *UpdateMaintenanceWindowForAppRequest`

NewUpdateMaintenanceWindowForAppRequestWithDefaults instantiates a new UpdateMaintenanceWindowForAppRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDayOfWeek

`func (o *UpdateMaintenanceWindowForAppRequest) GetDayOfWeek() int32`

GetDayOfWeek returns the DayOfWeek field if non-nil, zero value otherwise.

### GetDayOfWeekOk

`func (o *UpdateMaintenanceWindowForAppRequest) GetDayOfWeekOk() (*int32, bool)`

GetDayOfWeekOk returns a tuple with the DayOfWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfWeek

`func (o *UpdateMaintenanceWindowForAppRequest) SetDayOfWeek(v int32)`

SetDayOfWeek sets DayOfWeek field to given value.

### HasDayOfWeek

`func (o *UpdateMaintenanceWindowForAppRequest) HasDayOfWeek() bool`

HasDayOfWeek returns a boolean if a field has been set.

### SetDayOfWeekNil

`func (o *UpdateMaintenanceWindowForAppRequest) SetDayOfWeekNil(b bool)`

 SetDayOfWeekNil sets the value for DayOfWeek to be an explicit nil

### UnsetDayOfWeek
`func (o *UpdateMaintenanceWindowForAppRequest) UnsetDayOfWeek()`

UnsetDayOfWeek ensures that no value is present for DayOfWeek, not even an explicit nil
### GetStartTimeUtc

`func (o *UpdateMaintenanceWindowForAppRequest) GetStartTimeUtc() string`

GetStartTimeUtc returns the StartTimeUtc field if non-nil, zero value otherwise.

### GetStartTimeUtcOk

`func (o *UpdateMaintenanceWindowForAppRequest) GetStartTimeUtcOk() (*string, bool)`

GetStartTimeUtcOk returns a tuple with the StartTimeUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimeUtc

`func (o *UpdateMaintenanceWindowForAppRequest) SetStartTimeUtc(v string)`

SetStartTimeUtc sets StartTimeUtc field to given value.

### HasStartTimeUtc

`func (o *UpdateMaintenanceWindowForAppRequest) HasStartTimeUtc() bool`

HasStartTimeUtc returns a boolean if a field has been set.

### GetDurationHours

`func (o *UpdateMaintenanceWindowForAppRequest) GetDurationHours() int32`

GetDurationHours returns the DurationHours field if non-nil, zero value otherwise.

### GetDurationHoursOk

`func (o *UpdateMaintenanceWindowForAppRequest) GetDurationHoursOk() (*int32, bool)`

GetDurationHoursOk returns a tuple with the DurationHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationHours

`func (o *UpdateMaintenanceWindowForAppRequest) SetDurationHours(v int32)`

SetDurationHours sets DurationHours field to given value.

### HasDurationHours

`func (o *UpdateMaintenanceWindowForAppRequest) HasDurationHours() bool`

HasDurationHours returns a boolean if a field has been set.

### GetEnabled

`func (o *UpdateMaintenanceWindowForAppRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateMaintenanceWindowForAppRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateMaintenanceWindowForAppRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdateMaintenanceWindowForAppRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


