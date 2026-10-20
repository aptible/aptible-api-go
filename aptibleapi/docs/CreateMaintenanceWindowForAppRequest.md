# CreateMaintenanceWindowForAppRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DayOfWeek** | Pointer to **NullableInt32** | Day of week the window applies to (0 &#x3D; Sunday .. 6 &#x3D; Saturday). Omit to apply every day. | [optional] 
**StartTimeUtc** | **string** | Start of the window, in UTC (e.g. \&quot;22:00\&quot;) | 
**DurationHours** | **int32** | Length of the window in hours (2-24) | 
**Enabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewCreateMaintenanceWindowForAppRequest

`func NewCreateMaintenanceWindowForAppRequest(startTimeUtc string, durationHours int32, ) *CreateMaintenanceWindowForAppRequest`

NewCreateMaintenanceWindowForAppRequest instantiates a new CreateMaintenanceWindowForAppRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateMaintenanceWindowForAppRequestWithDefaults

`func NewCreateMaintenanceWindowForAppRequestWithDefaults() *CreateMaintenanceWindowForAppRequest`

NewCreateMaintenanceWindowForAppRequestWithDefaults instantiates a new CreateMaintenanceWindowForAppRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDayOfWeek

`func (o *CreateMaintenanceWindowForAppRequest) GetDayOfWeek() int32`

GetDayOfWeek returns the DayOfWeek field if non-nil, zero value otherwise.

### GetDayOfWeekOk

`func (o *CreateMaintenanceWindowForAppRequest) GetDayOfWeekOk() (*int32, bool)`

GetDayOfWeekOk returns a tuple with the DayOfWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfWeek

`func (o *CreateMaintenanceWindowForAppRequest) SetDayOfWeek(v int32)`

SetDayOfWeek sets DayOfWeek field to given value.

### HasDayOfWeek

`func (o *CreateMaintenanceWindowForAppRequest) HasDayOfWeek() bool`

HasDayOfWeek returns a boolean if a field has been set.

### SetDayOfWeekNil

`func (o *CreateMaintenanceWindowForAppRequest) SetDayOfWeekNil(b bool)`

 SetDayOfWeekNil sets the value for DayOfWeek to be an explicit nil

### UnsetDayOfWeek
`func (o *CreateMaintenanceWindowForAppRequest) UnsetDayOfWeek()`

UnsetDayOfWeek ensures that no value is present for DayOfWeek, not even an explicit nil
### GetStartTimeUtc

`func (o *CreateMaintenanceWindowForAppRequest) GetStartTimeUtc() string`

GetStartTimeUtc returns the StartTimeUtc field if non-nil, zero value otherwise.

### GetStartTimeUtcOk

`func (o *CreateMaintenanceWindowForAppRequest) GetStartTimeUtcOk() (*string, bool)`

GetStartTimeUtcOk returns a tuple with the StartTimeUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimeUtc

`func (o *CreateMaintenanceWindowForAppRequest) SetStartTimeUtc(v string)`

SetStartTimeUtc sets StartTimeUtc field to given value.


### GetDurationHours

`func (o *CreateMaintenanceWindowForAppRequest) GetDurationHours() int32`

GetDurationHours returns the DurationHours field if non-nil, zero value otherwise.

### GetDurationHoursOk

`func (o *CreateMaintenanceWindowForAppRequest) GetDurationHoursOk() (*int32, bool)`

GetDurationHoursOk returns a tuple with the DurationHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationHours

`func (o *CreateMaintenanceWindowForAppRequest) SetDurationHours(v int32)`

SetDurationHours sets DurationHours field to given value.


### GetEnabled

`func (o *CreateMaintenanceWindowForAppRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CreateMaintenanceWindowForAppRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CreateMaintenanceWindowForAppRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CreateMaintenanceWindowForAppRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


