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

// BTAppElementContentDeltaInfo struct for BTAppElementContentDeltaInfo
type BTAppElementContentDeltaInfo struct {
	Delta *string `json:"delta,omitempty"`
}

// NewBTAppElementContentDeltaInfo instantiates a new BTAppElementContentDeltaInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTAppElementContentDeltaInfo() *BTAppElementContentDeltaInfo {
	this := BTAppElementContentDeltaInfo{}
	return &this
}

// NewBTAppElementContentDeltaInfoWithDefaults instantiates a new BTAppElementContentDeltaInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTAppElementContentDeltaInfoWithDefaults() *BTAppElementContentDeltaInfo {
	this := BTAppElementContentDeltaInfo{}
	return &this
}

// GetDelta returns the Delta field value if set, zero value otherwise.
func (o *BTAppElementContentDeltaInfo) GetDelta() string {
	if o == nil || o.Delta == nil {
		var ret string
		return ret
	}
	return *o.Delta
}

// GetDeltaOk returns a tuple with the Delta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppElementContentDeltaInfo) GetDeltaOk() (*string, bool) {
	if o == nil || o.Delta == nil {
		return nil, false
	}
	return o.Delta, true
}

// HasDelta returns a boolean if a field has been set.
func (o *BTAppElementContentDeltaInfo) HasDelta() bool {
	if o != nil && o.Delta != nil {
		return true
	}

	return false
}

// SetDelta gets a reference to the given string and assigns it to the Delta field.
func (o *BTAppElementContentDeltaInfo) SetDelta(v string) {
	o.Delta = &v
}

func (o BTAppElementContentDeltaInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Delta != nil {
		toSerialize["delta"] = o.Delta
	}
	return json.Marshal(toSerialize)
}

type NullableBTAppElementContentDeltaInfo struct {
	value *BTAppElementContentDeltaInfo
	isSet bool
}

func (v NullableBTAppElementContentDeltaInfo) Get() *BTAppElementContentDeltaInfo {
	return v.value
}

func (v *NullableBTAppElementContentDeltaInfo) Set(val *BTAppElementContentDeltaInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTAppElementContentDeltaInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTAppElementContentDeltaInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTAppElementContentDeltaInfo(val *BTAppElementContentDeltaInfo) *NullableBTAppElementContentDeltaInfo {
	return &NullableBTAppElementContentDeltaInfo{value: val, isSet: true}
}

func (v NullableBTAppElementContentDeltaInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTAppElementContentDeltaInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
