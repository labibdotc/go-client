/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.173.26754-ceeaad064d4a
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTComputedAssemblyPropertyConfig struct for BTComputedAssemblyPropertyConfig
type BTComputedAssemblyPropertyConfig struct {
	AggregatedPropertyId     *string                                        `json:"aggregatedPropertyId,omitempty"`
	AggregationOperator      *BTComputedAssemblyPropertyAggregationOperator `json:"aggregationOperator,omitempty"`
	ErrorValuePolicy         *BTComputedAssemblyPropertyErrorPolicy         `json:"errorValuePolicy,omitempty"`
	FilterPropertyId         *string                                        `json:"filterPropertyId,omitempty"`
	IsFilterPropertyInverted *bool                                          `json:"isFilterPropertyInverted,omitempty"`
	MissingValuePolicy       *BTComputedAssemblyPropertyErrorPolicy         `json:"missingValuePolicy,omitempty"`
	SecondaryPropertyId      *string                                        `json:"secondaryPropertyId,omitempty"`
}

// NewBTComputedAssemblyPropertyConfig instantiates a new BTComputedAssemblyPropertyConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTComputedAssemblyPropertyConfig() *BTComputedAssemblyPropertyConfig {
	this := BTComputedAssemblyPropertyConfig{}
	return &this
}

// NewBTComputedAssemblyPropertyConfigWithDefaults instantiates a new BTComputedAssemblyPropertyConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTComputedAssemblyPropertyConfigWithDefaults() *BTComputedAssemblyPropertyConfig {
	this := BTComputedAssemblyPropertyConfig{}
	return &this
}

// GetAggregatedPropertyId returns the AggregatedPropertyId field value if set, zero value otherwise.
func (o *BTComputedAssemblyPropertyConfig) GetAggregatedPropertyId() string {
	if o == nil || o.AggregatedPropertyId == nil {
		var ret string
		return ret
	}
	return *o.AggregatedPropertyId
}

// GetAggregatedPropertyIdOk returns a tuple with the AggregatedPropertyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTComputedAssemblyPropertyConfig) GetAggregatedPropertyIdOk() (*string, bool) {
	if o == nil || o.AggregatedPropertyId == nil {
		return nil, false
	}
	return o.AggregatedPropertyId, true
}

// HasAggregatedPropertyId returns a boolean if a field has been set.
func (o *BTComputedAssemblyPropertyConfig) HasAggregatedPropertyId() bool {
	if o != nil && o.AggregatedPropertyId != nil {
		return true
	}

	return false
}

// SetAggregatedPropertyId gets a reference to the given string and assigns it to the AggregatedPropertyId field.
func (o *BTComputedAssemblyPropertyConfig) SetAggregatedPropertyId(v string) {
	o.AggregatedPropertyId = &v
}

// GetAggregationOperator returns the AggregationOperator field value if set, zero value otherwise.
func (o *BTComputedAssemblyPropertyConfig) GetAggregationOperator() BTComputedAssemblyPropertyAggregationOperator {
	if o == nil || o.AggregationOperator == nil {
		var ret BTComputedAssemblyPropertyAggregationOperator
		return ret
	}
	return *o.AggregationOperator
}

// GetAggregationOperatorOk returns a tuple with the AggregationOperator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTComputedAssemblyPropertyConfig) GetAggregationOperatorOk() (*BTComputedAssemblyPropertyAggregationOperator, bool) {
	if o == nil || o.AggregationOperator == nil {
		return nil, false
	}
	return o.AggregationOperator, true
}

// HasAggregationOperator returns a boolean if a field has been set.
func (o *BTComputedAssemblyPropertyConfig) HasAggregationOperator() bool {
	if o != nil && o.AggregationOperator != nil {
		return true
	}

	return false
}

// SetAggregationOperator gets a reference to the given BTComputedAssemblyPropertyAggregationOperator and assigns it to the AggregationOperator field.
func (o *BTComputedAssemblyPropertyConfig) SetAggregationOperator(v BTComputedAssemblyPropertyAggregationOperator) {
	o.AggregationOperator = &v
}

// GetErrorValuePolicy returns the ErrorValuePolicy field value if set, zero value otherwise.
func (o *BTComputedAssemblyPropertyConfig) GetErrorValuePolicy() BTComputedAssemblyPropertyErrorPolicy {
	if o == nil || o.ErrorValuePolicy == nil {
		var ret BTComputedAssemblyPropertyErrorPolicy
		return ret
	}
	return *o.ErrorValuePolicy
}

// GetErrorValuePolicyOk returns a tuple with the ErrorValuePolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTComputedAssemblyPropertyConfig) GetErrorValuePolicyOk() (*BTComputedAssemblyPropertyErrorPolicy, bool) {
	if o == nil || o.ErrorValuePolicy == nil {
		return nil, false
	}
	return o.ErrorValuePolicy, true
}

// HasErrorValuePolicy returns a boolean if a field has been set.
func (o *BTComputedAssemblyPropertyConfig) HasErrorValuePolicy() bool {
	if o != nil && o.ErrorValuePolicy != nil {
		return true
	}

	return false
}

// SetErrorValuePolicy gets a reference to the given BTComputedAssemblyPropertyErrorPolicy and assigns it to the ErrorValuePolicy field.
func (o *BTComputedAssemblyPropertyConfig) SetErrorValuePolicy(v BTComputedAssemblyPropertyErrorPolicy) {
	o.ErrorValuePolicy = &v
}

// GetFilterPropertyId returns the FilterPropertyId field value if set, zero value otherwise.
func (o *BTComputedAssemblyPropertyConfig) GetFilterPropertyId() string {
	if o == nil || o.FilterPropertyId == nil {
		var ret string
		return ret
	}
	return *o.FilterPropertyId
}

// GetFilterPropertyIdOk returns a tuple with the FilterPropertyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTComputedAssemblyPropertyConfig) GetFilterPropertyIdOk() (*string, bool) {
	if o == nil || o.FilterPropertyId == nil {
		return nil, false
	}
	return o.FilterPropertyId, true
}

// HasFilterPropertyId returns a boolean if a field has been set.
func (o *BTComputedAssemblyPropertyConfig) HasFilterPropertyId() bool {
	if o != nil && o.FilterPropertyId != nil {
		return true
	}

	return false
}

// SetFilterPropertyId gets a reference to the given string and assigns it to the FilterPropertyId field.
func (o *BTComputedAssemblyPropertyConfig) SetFilterPropertyId(v string) {
	o.FilterPropertyId = &v
}

// GetIsFilterPropertyInverted returns the IsFilterPropertyInverted field value if set, zero value otherwise.
func (o *BTComputedAssemblyPropertyConfig) GetIsFilterPropertyInverted() bool {
	if o == nil || o.IsFilterPropertyInverted == nil {
		var ret bool
		return ret
	}
	return *o.IsFilterPropertyInverted
}

// GetIsFilterPropertyInvertedOk returns a tuple with the IsFilterPropertyInverted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTComputedAssemblyPropertyConfig) GetIsFilterPropertyInvertedOk() (*bool, bool) {
	if o == nil || o.IsFilterPropertyInverted == nil {
		return nil, false
	}
	return o.IsFilterPropertyInverted, true
}

// HasIsFilterPropertyInverted returns a boolean if a field has been set.
func (o *BTComputedAssemblyPropertyConfig) HasIsFilterPropertyInverted() bool {
	if o != nil && o.IsFilterPropertyInverted != nil {
		return true
	}

	return false
}

// SetIsFilterPropertyInverted gets a reference to the given bool and assigns it to the IsFilterPropertyInverted field.
func (o *BTComputedAssemblyPropertyConfig) SetIsFilterPropertyInverted(v bool) {
	o.IsFilterPropertyInverted = &v
}

// GetMissingValuePolicy returns the MissingValuePolicy field value if set, zero value otherwise.
func (o *BTComputedAssemblyPropertyConfig) GetMissingValuePolicy() BTComputedAssemblyPropertyErrorPolicy {
	if o == nil || o.MissingValuePolicy == nil {
		var ret BTComputedAssemblyPropertyErrorPolicy
		return ret
	}
	return *o.MissingValuePolicy
}

// GetMissingValuePolicyOk returns a tuple with the MissingValuePolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTComputedAssemblyPropertyConfig) GetMissingValuePolicyOk() (*BTComputedAssemblyPropertyErrorPolicy, bool) {
	if o == nil || o.MissingValuePolicy == nil {
		return nil, false
	}
	return o.MissingValuePolicy, true
}

// HasMissingValuePolicy returns a boolean if a field has been set.
func (o *BTComputedAssemblyPropertyConfig) HasMissingValuePolicy() bool {
	if o != nil && o.MissingValuePolicy != nil {
		return true
	}

	return false
}

// SetMissingValuePolicy gets a reference to the given BTComputedAssemblyPropertyErrorPolicy and assigns it to the MissingValuePolicy field.
func (o *BTComputedAssemblyPropertyConfig) SetMissingValuePolicy(v BTComputedAssemblyPropertyErrorPolicy) {
	o.MissingValuePolicy = &v
}

// GetSecondaryPropertyId returns the SecondaryPropertyId field value if set, zero value otherwise.
func (o *BTComputedAssemblyPropertyConfig) GetSecondaryPropertyId() string {
	if o == nil || o.SecondaryPropertyId == nil {
		var ret string
		return ret
	}
	return *o.SecondaryPropertyId
}

// GetSecondaryPropertyIdOk returns a tuple with the SecondaryPropertyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTComputedAssemblyPropertyConfig) GetSecondaryPropertyIdOk() (*string, bool) {
	if o == nil || o.SecondaryPropertyId == nil {
		return nil, false
	}
	return o.SecondaryPropertyId, true
}

// HasSecondaryPropertyId returns a boolean if a field has been set.
func (o *BTComputedAssemblyPropertyConfig) HasSecondaryPropertyId() bool {
	if o != nil && o.SecondaryPropertyId != nil {
		return true
	}

	return false
}

// SetSecondaryPropertyId gets a reference to the given string and assigns it to the SecondaryPropertyId field.
func (o *BTComputedAssemblyPropertyConfig) SetSecondaryPropertyId(v string) {
	o.SecondaryPropertyId = &v
}

func (o BTComputedAssemblyPropertyConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AggregatedPropertyId != nil {
		toSerialize["aggregatedPropertyId"] = o.AggregatedPropertyId
	}
	if o.AggregationOperator != nil {
		toSerialize["aggregationOperator"] = o.AggregationOperator
	}
	if o.ErrorValuePolicy != nil {
		toSerialize["errorValuePolicy"] = o.ErrorValuePolicy
	}
	if o.FilterPropertyId != nil {
		toSerialize["filterPropertyId"] = o.FilterPropertyId
	}
	if o.IsFilterPropertyInverted != nil {
		toSerialize["isFilterPropertyInverted"] = o.IsFilterPropertyInverted
	}
	if o.MissingValuePolicy != nil {
		toSerialize["missingValuePolicy"] = o.MissingValuePolicy
	}
	if o.SecondaryPropertyId != nil {
		toSerialize["secondaryPropertyId"] = o.SecondaryPropertyId
	}
	return json.Marshal(toSerialize)
}

type NullableBTComputedAssemblyPropertyConfig struct {
	value *BTComputedAssemblyPropertyConfig
	isSet bool
}

func (v NullableBTComputedAssemblyPropertyConfig) Get() *BTComputedAssemblyPropertyConfig {
	return v.value
}

func (v *NullableBTComputedAssemblyPropertyConfig) Set(val *BTComputedAssemblyPropertyConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableBTComputedAssemblyPropertyConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableBTComputedAssemblyPropertyConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTComputedAssemblyPropertyConfig(val *BTComputedAssemblyPropertyConfig) *NullableBTComputedAssemblyPropertyConfig {
	return &NullableBTComputedAssemblyPropertyConfig{value: val, isSet: true}
}

func (v NullableBTComputedAssemblyPropertyConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTComputedAssemblyPropertyConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
