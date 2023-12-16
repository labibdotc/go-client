/*
Onshape REST API

## Welcome to the Onshape REST API Explorer  To use this API explorer, sign in to your [Onshape](https://cad.onshape.com) account in another tab, then click the **Try it out** button below (it toggles to a **Cancel** button when selected).  See the **[API Explorer Guide](https://onshape-public.github.io/docs/api-intro/explorer/)** for help navigating this API Explorer, including **[authentication](https://onshape-public.github.io/docs/api-intro/explorer/#authentication)**.  **Tip:** To ensure the current session isn't used when trying other authentication techniques, make sure to [remove the Onshape cookie](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site) as per the instructions for your browser. Alternatively, you can use a private or incognito window.  ## See Also  * [Onshape API Guide](https://onshape-public.github.io/docs/): Our full suite of developer guides, to be used as an accompaniment to this API Explorer. * [Onshape Developer Portal](https://dev-portal.onshape.com/): The Onshape portal for managing your API keys, OAuth2 credentials, your Onshape applications, and your Onshape App Store entries. * [Authentication Guide](https://onshape-public.github.io/docs/auth/): Our guide to using API keys, request signatures, and OAuth2 in  your Onshape applications.

API version: 1.174.27783-ab3907bf6199
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTMassPropertiesBulkInfo struct for BTMassPropertiesBulkInfo
type BTMassPropertiesBulkInfo struct {
	Bodies         *map[string]BTMassPropertiesInfo `json:"bodies,omitempty"`
	MicroversionId *string                          `json:"microversionId,omitempty"`
}

// NewBTMassPropertiesBulkInfo instantiates a new BTMassPropertiesBulkInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTMassPropertiesBulkInfo() *BTMassPropertiesBulkInfo {
	this := BTMassPropertiesBulkInfo{}
	return &this
}

// NewBTMassPropertiesBulkInfoWithDefaults instantiates a new BTMassPropertiesBulkInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTMassPropertiesBulkInfoWithDefaults() *BTMassPropertiesBulkInfo {
	this := BTMassPropertiesBulkInfo{}
	return &this
}

// GetBodies returns the Bodies field value if set, zero value otherwise.
func (o *BTMassPropertiesBulkInfo) GetBodies() map[string]BTMassPropertiesInfo {
	if o == nil || o.Bodies == nil {
		var ret map[string]BTMassPropertiesInfo
		return ret
	}
	return *o.Bodies
}

// GetBodiesOk returns a tuple with the Bodies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMassPropertiesBulkInfo) GetBodiesOk() (*map[string]BTMassPropertiesInfo, bool) {
	if o == nil || o.Bodies == nil {
		return nil, false
	}
	return o.Bodies, true
}

// HasBodies returns a boolean if a field has been set.
func (o *BTMassPropertiesBulkInfo) HasBodies() bool {
	if o != nil && o.Bodies != nil {
		return true
	}

	return false
}

// SetBodies gets a reference to the given map[string]BTMassPropertiesInfo and assigns it to the Bodies field.
func (o *BTMassPropertiesBulkInfo) SetBodies(v map[string]BTMassPropertiesInfo) {
	o.Bodies = &v
}

// GetMicroversionId returns the MicroversionId field value if set, zero value otherwise.
func (o *BTMassPropertiesBulkInfo) GetMicroversionId() string {
	if o == nil || o.MicroversionId == nil {
		var ret string
		return ret
	}
	return *o.MicroversionId
}

// GetMicroversionIdOk returns a tuple with the MicroversionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMassPropertiesBulkInfo) GetMicroversionIdOk() (*string, bool) {
	if o == nil || o.MicroversionId == nil {
		return nil, false
	}
	return o.MicroversionId, true
}

// HasMicroversionId returns a boolean if a field has been set.
func (o *BTMassPropertiesBulkInfo) HasMicroversionId() bool {
	if o != nil && o.MicroversionId != nil {
		return true
	}

	return false
}

// SetMicroversionId gets a reference to the given string and assigns it to the MicroversionId field.
func (o *BTMassPropertiesBulkInfo) SetMicroversionId(v string) {
	o.MicroversionId = &v
}

func (o BTMassPropertiesBulkInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Bodies != nil {
		toSerialize["bodies"] = o.Bodies
	}
	if o.MicroversionId != nil {
		toSerialize["microversionId"] = o.MicroversionId
	}
	return json.Marshal(toSerialize)
}

type NullableBTMassPropertiesBulkInfo struct {
	value *BTMassPropertiesBulkInfo
	isSet bool
}

func (v NullableBTMassPropertiesBulkInfo) Get() *BTMassPropertiesBulkInfo {
	return v.value
}

func (v *NullableBTMassPropertiesBulkInfo) Set(val *BTMassPropertiesBulkInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTMassPropertiesBulkInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTMassPropertiesBulkInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTMassPropertiesBulkInfo(val *BTMassPropertiesBulkInfo) *NullableBTMassPropertiesBulkInfo {
	return &NullableBTMassPropertiesBulkInfo{value: val, isSet: true}
}

func (v NullableBTMassPropertiesBulkInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTMassPropertiesBulkInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
