# MaintenanceWindow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**AccountId** | Pointer to **int32** |  | [optional] 
**DayOfWeek** | Pointer to **NullableInt32** |  | [optional] 
**StartTimeUtc** | Pointer to **string** | Start of the window, in UTC, as HH:MM (e.g. \&quot;22:00\&quot;) | [optional] 
**DurationHours** | Pointer to **int32** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**Links** | Pointer to [**MaintenanceWindowLinks**](MaintenanceWindowLinks.md) |  | [optional] 

## Methods

### NewMaintenanceWindow

`func NewMaintenanceWindow(id int32, ) *MaintenanceWindow`

NewMaintenanceWindow instantiates a new MaintenanceWindow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMaintenanceWindowWithDefaults

`func NewMaintenanceWindowWithDefaults() *MaintenanceWindow`

NewMaintenanceWindowWithDefaults instantiates a new MaintenanceWindow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MaintenanceWindow) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MaintenanceWindow) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MaintenanceWindow) SetId(v int32)`

SetId sets Id field to given value.


### GetAccountId

`func (o *MaintenanceWindow) GetAccountId() int32`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *MaintenanceWindow) GetAccountIdOk() (*int32, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *MaintenanceWindow) SetAccountId(v int32)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *MaintenanceWindow) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetDayOfWeek

`func (o *MaintenanceWindow) GetDayOfWeek() int32`

GetDayOfWeek returns the DayOfWeek field if non-nil, zero value otherwise.

### GetDayOfWeekOk

`func (o *MaintenanceWindow) GetDayOfWeekOk() (*int32, bool)`

GetDayOfWeekOk returns a tuple with the DayOfWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfWeek

`func (o *MaintenanceWindow) SetDayOfWeek(v int32)`

SetDayOfWeek sets DayOfWeek field to given value.

### HasDayOfWeek

`func (o *MaintenanceWindow) HasDayOfWeek() bool`

HasDayOfWeek returns a boolean if a field has been set.

### SetDayOfWeekNil

`func (o *MaintenanceWindow) SetDayOfWeekNil(b bool)`

 SetDayOfWeekNil sets the value for DayOfWeek to be an explicit nil

### UnsetDayOfWeek
`func (o *MaintenanceWindow) UnsetDayOfWeek()`

UnsetDayOfWeek ensures that no value is present for DayOfWeek, not even an explicit nil
### GetStartTimeUtc

`func (o *MaintenanceWindow) GetStartTimeUtc() string`

GetStartTimeUtc returns the StartTimeUtc field if non-nil, zero value otherwise.

### GetStartTimeUtcOk

`func (o *MaintenanceWindow) GetStartTimeUtcOk() (*string, bool)`

GetStartTimeUtcOk returns a tuple with the StartTimeUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimeUtc

`func (o *MaintenanceWindow) SetStartTimeUtc(v string)`

SetStartTimeUtc sets StartTimeUtc field to given value.

### HasStartTimeUtc

`func (o *MaintenanceWindow) HasStartTimeUtc() bool`

HasStartTimeUtc returns a boolean if a field has been set.

### GetDurationHours

`func (o *MaintenanceWindow) GetDurationHours() int32`

GetDurationHours returns the DurationHours field if non-nil, zero value otherwise.

### GetDurationHoursOk

`func (o *MaintenanceWindow) GetDurationHoursOk() (*int32, bool)`

GetDurationHoursOk returns a tuple with the DurationHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationHours

`func (o *MaintenanceWindow) SetDurationHours(v int32)`

SetDurationHours sets DurationHours field to given value.

### HasDurationHours

`func (o *MaintenanceWindow) HasDurationHours() bool`

HasDurationHours returns a boolean if a field has been set.

### GetEnabled

`func (o *MaintenanceWindow) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *MaintenanceWindow) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *MaintenanceWindow) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *MaintenanceWindow) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetCreatedAt

`func (o *MaintenanceWindow) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MaintenanceWindow) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MaintenanceWindow) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *MaintenanceWindow) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *MaintenanceWindow) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *MaintenanceWindow) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *MaintenanceWindow) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *MaintenanceWindow) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetLinks

`func (o *MaintenanceWindow) GetLinks() MaintenanceWindowLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MaintenanceWindow) GetLinksOk() (*MaintenanceWindowLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MaintenanceWindow) SetLinks(v MaintenanceWindowLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *MaintenanceWindow) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


