# UpdateExternalAwsResourceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceName** | Pointer to **string** |  | [optional] 
**LockVersion** | Pointer to **int32** |  | [optional] 

## Methods

### NewUpdateExternalAwsResourceRequest

`func NewUpdateExternalAwsResourceRequest() *UpdateExternalAwsResourceRequest`

NewUpdateExternalAwsResourceRequest instantiates a new UpdateExternalAwsResourceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateExternalAwsResourceRequestWithDefaults

`func NewUpdateExternalAwsResourceRequestWithDefaults() *UpdateExternalAwsResourceRequest`

NewUpdateExternalAwsResourceRequestWithDefaults instantiates a new UpdateExternalAwsResourceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceName

`func (o *UpdateExternalAwsResourceRequest) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *UpdateExternalAwsResourceRequest) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *UpdateExternalAwsResourceRequest) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.

### HasResourceName

`func (o *UpdateExternalAwsResourceRequest) HasResourceName() bool`

HasResourceName returns a boolean if a field has been set.

### GetLockVersion

`func (o *UpdateExternalAwsResourceRequest) GetLockVersion() int32`

GetLockVersion returns the LockVersion field if non-nil, zero value otherwise.

### GetLockVersionOk

`func (o *UpdateExternalAwsResourceRequest) GetLockVersionOk() (*int32, bool)`

GetLockVersionOk returns a tuple with the LockVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockVersion

`func (o *UpdateExternalAwsResourceRequest) SetLockVersion(v int32)`

SetLockVersion sets LockVersion field to given value.

### HasLockVersion

`func (o *UpdateExternalAwsResourceRequest) HasLockVersion() bool`

HasLockVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


