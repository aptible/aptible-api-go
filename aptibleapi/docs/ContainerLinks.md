# ContainerLinks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Release** | Pointer to [**ContainerLinksRelease**](ContainerLinksRelease.md) |  | [optional] 
**Vhost** | Pointer to [**ContainerLinksVhost**](ContainerLinksVhost.md) |  | [optional] 
**LogDrain** | Pointer to [**ContainerLinksLogDrain**](ContainerLinksLogDrain.md) |  | [optional] 
**MetricDrain** | Pointer to [**ContainerLinksMetricDrain**](ContainerLinksMetricDrain.md) |  | [optional] 
**Self** | Pointer to [**ContainerLinksSelf**](ContainerLinksSelf.md) |  | [optional] 

## Methods

### NewContainerLinks

`func NewContainerLinks() *ContainerLinks`

NewContainerLinks instantiates a new ContainerLinks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerLinksWithDefaults

`func NewContainerLinksWithDefaults() *ContainerLinks`

NewContainerLinksWithDefaults instantiates a new ContainerLinks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRelease

`func (o *ContainerLinks) GetRelease() ContainerLinksRelease`

GetRelease returns the Release field if non-nil, zero value otherwise.

### GetReleaseOk

`func (o *ContainerLinks) GetReleaseOk() (*ContainerLinksRelease, bool)`

GetReleaseOk returns a tuple with the Release field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelease

`func (o *ContainerLinks) SetRelease(v ContainerLinksRelease)`

SetRelease sets Release field to given value.

### HasRelease

`func (o *ContainerLinks) HasRelease() bool`

HasRelease returns a boolean if a field has been set.

### GetVhost

`func (o *ContainerLinks) GetVhost() ContainerLinksVhost`

GetVhost returns the Vhost field if non-nil, zero value otherwise.

### GetVhostOk

`func (o *ContainerLinks) GetVhostOk() (*ContainerLinksVhost, bool)`

GetVhostOk returns a tuple with the Vhost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVhost

`func (o *ContainerLinks) SetVhost(v ContainerLinksVhost)`

SetVhost sets Vhost field to given value.

### HasVhost

`func (o *ContainerLinks) HasVhost() bool`

HasVhost returns a boolean if a field has been set.

### GetLogDrain

`func (o *ContainerLinks) GetLogDrain() ContainerLinksLogDrain`

GetLogDrain returns the LogDrain field if non-nil, zero value otherwise.

### GetLogDrainOk

`func (o *ContainerLinks) GetLogDrainOk() (*ContainerLinksLogDrain, bool)`

GetLogDrainOk returns a tuple with the LogDrain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogDrain

`func (o *ContainerLinks) SetLogDrain(v ContainerLinksLogDrain)`

SetLogDrain sets LogDrain field to given value.

### HasLogDrain

`func (o *ContainerLinks) HasLogDrain() bool`

HasLogDrain returns a boolean if a field has been set.

### GetMetricDrain

`func (o *ContainerLinks) GetMetricDrain() ContainerLinksMetricDrain`

GetMetricDrain returns the MetricDrain field if non-nil, zero value otherwise.

### GetMetricDrainOk

`func (o *ContainerLinks) GetMetricDrainOk() (*ContainerLinksMetricDrain, bool)`

GetMetricDrainOk returns a tuple with the MetricDrain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricDrain

`func (o *ContainerLinks) SetMetricDrain(v ContainerLinksMetricDrain)`

SetMetricDrain sets MetricDrain field to given value.

### HasMetricDrain

`func (o *ContainerLinks) HasMetricDrain() bool`

HasMetricDrain returns a boolean if a field has been set.

### GetSelf

`func (o *ContainerLinks) GetSelf() ContainerLinksSelf`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *ContainerLinks) GetSelfOk() (*ContainerLinksSelf, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *ContainerLinks) SetSelf(v ContainerLinksSelf)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *ContainerLinks) HasSelf() bool`

HasSelf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


