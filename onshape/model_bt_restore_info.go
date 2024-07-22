/*
Onshape REST API

## Welcome to the Onshape REST API Explorer  To use this API explorer, sign in to your [Onshape](https://cad.onshape.com) account in another tab, then click the **Try it out** button below (it toggles to a **Cancel** button when selected).  See the **[API Explorer Guide](https://onshape-public.github.io/docs/api-intro/explorer/)** for help navigating this API Explorer, including **[authentication](https://onshape-public.github.io/docs/api-intro/explorer/#authentication)**.  **Tip:** To ensure the current session isn't used when trying other authentication techniques, make sure to [remove the Onshape cookie](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site) as per the instructions for your browser. Alternatively, you can use a private or incognito window.  ## See Also  * [Onshape API Guide](https://onshape-public.github.io/docs/): Our full suite of developer guides, to be used as an accompaniment to this API Explorer. * [Onshape Developer Portal](https://dev-portal.onshape.com/): The Onshape portal for managing your API keys, OAuth2 credentials, your Onshape applications, and your Onshape App Store entries. * [Authentication Guide](https://onshape-public.github.io/docs/auth/): Our guide to using API keys, request signatures, and OAuth2 in  your Onshape applications.

Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTRestoreInfo struct for BTRestoreInfo
type BTRestoreInfo struct {
	DefaultRestoreStrategy      *BTRestoreStrategy            `json:"defaultRestoreStrategy,omitempty"`
	ElementIdToStrategyOverride *map[string]BTRestoreStrategy `json:"elementIdToStrategyOverride,omitempty"`
}

// NewBTRestoreInfo instantiates a new BTRestoreInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTRestoreInfo() *BTRestoreInfo {
	this := BTRestoreInfo{}
	return &this
}

// NewBTRestoreInfoWithDefaults instantiates a new BTRestoreInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTRestoreInfoWithDefaults() *BTRestoreInfo {
	this := BTRestoreInfo{}
	return &this
}

// GetDefaultRestoreStrategy returns the DefaultRestoreStrategy field value if set, zero value otherwise.
func (o *BTRestoreInfo) GetDefaultRestoreStrategy() BTRestoreStrategy {
	if o == nil || o.DefaultRestoreStrategy == nil {
		var ret BTRestoreStrategy
		return ret
	}
	return *o.DefaultRestoreStrategy
}

// GetDefaultRestoreStrategyOk returns a tuple with the DefaultRestoreStrategy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRestoreInfo) GetDefaultRestoreStrategyOk() (*BTRestoreStrategy, bool) {
	if o == nil || o.DefaultRestoreStrategy == nil {
		return nil, false
	}
	return o.DefaultRestoreStrategy, true
}

// HasDefaultRestoreStrategy returns a boolean if a field has been set.
func (o *BTRestoreInfo) HasDefaultRestoreStrategy() bool {
	if o != nil && o.DefaultRestoreStrategy != nil {
		return true
	}

	return false
}

// SetDefaultRestoreStrategy gets a reference to the given BTRestoreStrategy and assigns it to the DefaultRestoreStrategy field.
func (o *BTRestoreInfo) SetDefaultRestoreStrategy(v BTRestoreStrategy) {
	o.DefaultRestoreStrategy = &v
}

// GetElementIdToStrategyOverride returns the ElementIdToStrategyOverride field value if set, zero value otherwise.
func (o *BTRestoreInfo) GetElementIdToStrategyOverride() map[string]BTRestoreStrategy {
	if o == nil || o.ElementIdToStrategyOverride == nil {
		var ret map[string]BTRestoreStrategy
		return ret
	}
	return *o.ElementIdToStrategyOverride
}

// GetElementIdToStrategyOverrideOk returns a tuple with the ElementIdToStrategyOverride field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRestoreInfo) GetElementIdToStrategyOverrideOk() (*map[string]BTRestoreStrategy, bool) {
	if o == nil || o.ElementIdToStrategyOverride == nil {
		return nil, false
	}
	return o.ElementIdToStrategyOverride, true
}

// HasElementIdToStrategyOverride returns a boolean if a field has been set.
func (o *BTRestoreInfo) HasElementIdToStrategyOverride() bool {
	if o != nil && o.ElementIdToStrategyOverride != nil {
		return true
	}

	return false
}

// SetElementIdToStrategyOverride gets a reference to the given map[string]BTRestoreStrategy and assigns it to the ElementIdToStrategyOverride field.
func (o *BTRestoreInfo) SetElementIdToStrategyOverride(v map[string]BTRestoreStrategy) {
	o.ElementIdToStrategyOverride = &v
}

func (o BTRestoreInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DefaultRestoreStrategy != nil {
		toSerialize["defaultRestoreStrategy"] = o.DefaultRestoreStrategy
	}
	if o.ElementIdToStrategyOverride != nil {
		toSerialize["elementIdToStrategyOverride"] = o.ElementIdToStrategyOverride
	}
	return json.Marshal(toSerialize)
}

type NullableBTRestoreInfo struct {
	value *BTRestoreInfo
	isSet bool
}

func (v NullableBTRestoreInfo) Get() *BTRestoreInfo {
	return v.value
}

func (v *NullableBTRestoreInfo) Set(val *BTRestoreInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTRestoreInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTRestoreInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTRestoreInfo(val *BTRestoreInfo) *NullableBTRestoreInfo {
	return &NullableBTRestoreInfo{value: val, isSet: true}
}

func (v NullableBTRestoreInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTRestoreInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}