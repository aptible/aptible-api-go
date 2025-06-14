/*
Deploy API v1

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package aptibleapi

import (
	"encoding/json"
	"fmt"
)

// checks if the CreateWidgetForDashboardRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateWidgetForDashboardRequest{}

// CreateWidgetForDashboardRequest struct for CreateWidgetForDashboardRequest
type CreateWidgetForDashboardRequest struct {
	ResourceId int32 `json:"resource_id"`
	ResourceType string `json:"resource_type"`
	Title string `json:"title"`
	Summary string `json:"summary"`
	Rating int32 `json:"rating"`
	WidgetType string `json:"widget_type"`
	Data map[string]interface{} `json:"data"`
	RangeBegin string `json:"range_begin"`
	RangeEnd string `json:"range_end"`
	AdditionalProperties map[string]interface{}
}

type _CreateWidgetForDashboardRequest CreateWidgetForDashboardRequest

// NewCreateWidgetForDashboardRequest instantiates a new CreateWidgetForDashboardRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateWidgetForDashboardRequest(resourceId int32, resourceType string, title string, summary string, rating int32, widgetType string, data map[string]interface{}, rangeBegin string, rangeEnd string) *CreateWidgetForDashboardRequest {
	this := CreateWidgetForDashboardRequest{}
	this.ResourceId = resourceId
	this.ResourceType = resourceType
	this.Title = title
	this.Summary = summary
	this.Rating = rating
	this.WidgetType = widgetType
	this.Data = data
	this.RangeBegin = rangeBegin
	this.RangeEnd = rangeEnd
	return &this
}

// NewCreateWidgetForDashboardRequestWithDefaults instantiates a new CreateWidgetForDashboardRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateWidgetForDashboardRequestWithDefaults() *CreateWidgetForDashboardRequest {
	this := CreateWidgetForDashboardRequest{}
	return &this
}

// GetResourceId returns the ResourceId field value
func (o *CreateWidgetForDashboardRequest) GetResourceId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value
// and a boolean to check if the value has been set.
func (o *CreateWidgetForDashboardRequest) GetResourceIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceId, true
}

// SetResourceId sets field value
func (o *CreateWidgetForDashboardRequest) SetResourceId(v int32) {
	o.ResourceId = v
}

// GetResourceType returns the ResourceType field value
func (o *CreateWidgetForDashboardRequest) GetResourceType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value
// and a boolean to check if the value has been set.
func (o *CreateWidgetForDashboardRequest) GetResourceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceType, true
}

// SetResourceType sets field value
func (o *CreateWidgetForDashboardRequest) SetResourceType(v string) {
	o.ResourceType = v
}

// GetTitle returns the Title field value
func (o *CreateWidgetForDashboardRequest) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *CreateWidgetForDashboardRequest) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *CreateWidgetForDashboardRequest) SetTitle(v string) {
	o.Title = v
}

// GetSummary returns the Summary field value
func (o *CreateWidgetForDashboardRequest) GetSummary() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value
// and a boolean to check if the value has been set.
func (o *CreateWidgetForDashboardRequest) GetSummaryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Summary, true
}

// SetSummary sets field value
func (o *CreateWidgetForDashboardRequest) SetSummary(v string) {
	o.Summary = v
}

// GetRating returns the Rating field value
func (o *CreateWidgetForDashboardRequest) GetRating() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Rating
}

// GetRatingOk returns a tuple with the Rating field value
// and a boolean to check if the value has been set.
func (o *CreateWidgetForDashboardRequest) GetRatingOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rating, true
}

// SetRating sets field value
func (o *CreateWidgetForDashboardRequest) SetRating(v int32) {
	o.Rating = v
}

// GetWidgetType returns the WidgetType field value
func (o *CreateWidgetForDashboardRequest) GetWidgetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WidgetType
}

// GetWidgetTypeOk returns a tuple with the WidgetType field value
// and a boolean to check if the value has been set.
func (o *CreateWidgetForDashboardRequest) GetWidgetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WidgetType, true
}

// SetWidgetType sets field value
func (o *CreateWidgetForDashboardRequest) SetWidgetType(v string) {
	o.WidgetType = v
}

// GetData returns the Data field value
func (o *CreateWidgetForDashboardRequest) GetData() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CreateWidgetForDashboardRequest) GetDataOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *CreateWidgetForDashboardRequest) SetData(v map[string]interface{}) {
	o.Data = v
}

// GetRangeBegin returns the RangeBegin field value
func (o *CreateWidgetForDashboardRequest) GetRangeBegin() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RangeBegin
}

// GetRangeBeginOk returns a tuple with the RangeBegin field value
// and a boolean to check if the value has been set.
func (o *CreateWidgetForDashboardRequest) GetRangeBeginOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RangeBegin, true
}

// SetRangeBegin sets field value
func (o *CreateWidgetForDashboardRequest) SetRangeBegin(v string) {
	o.RangeBegin = v
}

// GetRangeEnd returns the RangeEnd field value
func (o *CreateWidgetForDashboardRequest) GetRangeEnd() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RangeEnd
}

// GetRangeEndOk returns a tuple with the RangeEnd field value
// and a boolean to check if the value has been set.
func (o *CreateWidgetForDashboardRequest) GetRangeEndOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RangeEnd, true
}

// SetRangeEnd sets field value
func (o *CreateWidgetForDashboardRequest) SetRangeEnd(v string) {
	o.RangeEnd = v
}

func (o CreateWidgetForDashboardRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateWidgetForDashboardRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["resource_id"] = o.ResourceId
	toSerialize["resource_type"] = o.ResourceType
	toSerialize["title"] = o.Title
	toSerialize["summary"] = o.Summary
	toSerialize["rating"] = o.Rating
	toSerialize["widget_type"] = o.WidgetType
	toSerialize["data"] = o.Data
	toSerialize["range_begin"] = o.RangeBegin
	toSerialize["range_end"] = o.RangeEnd

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateWidgetForDashboardRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"resource_id",
		"resource_type",
		"title",
		"summary",
		"rating",
		"widget_type",
		"data",
		"range_begin",
		"range_end",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCreateWidgetForDashboardRequest := _CreateWidgetForDashboardRequest{}

	err = json.Unmarshal(data, &varCreateWidgetForDashboardRequest)

	if err != nil {
		return err
	}

	*o = CreateWidgetForDashboardRequest(varCreateWidgetForDashboardRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "resource_id")
		delete(additionalProperties, "resource_type")
		delete(additionalProperties, "title")
		delete(additionalProperties, "summary")
		delete(additionalProperties, "rating")
		delete(additionalProperties, "widget_type")
		delete(additionalProperties, "data")
		delete(additionalProperties, "range_begin")
		delete(additionalProperties, "range_end")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateWidgetForDashboardRequest struct {
	value *CreateWidgetForDashboardRequest
	isSet bool
}

func (v NullableCreateWidgetForDashboardRequest) Get() *CreateWidgetForDashboardRequest {
	return v.value
}

func (v *NullableCreateWidgetForDashboardRequest) Set(val *CreateWidgetForDashboardRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateWidgetForDashboardRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateWidgetForDashboardRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateWidgetForDashboardRequest(val *CreateWidgetForDashboardRequest) *NullableCreateWidgetForDashboardRequest {
	return &NullableCreateWidgetForDashboardRequest{value: val, isSet: true}
}

func (v NullableCreateWidgetForDashboardRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateWidgetForDashboardRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


