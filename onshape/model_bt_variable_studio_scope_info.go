/*
Onshape REST API

## Welcome to the Onshape REST API Explorer  To use this API explorer, sign in to your [Onshape](https://cad.onshape.com) account in another tab, then click the **Try it out** button below (it toggles to a **Cancel** button when selected).  See the **[API Explorer Guide](https://onshape-public.github.io/docs/api-intro/explorer/)** for help navigating this API Explorer, including **[authentication](https://onshape-public.github.io/docs/api-intro/explorer/#authentication)**.  **Tip:** To ensure the current session isn't used when trying other authentication techniques, make sure to [remove the Onshape cookie](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site) as per the instructions for your browser. Alternatively, you can use a private or incognito window.  ## See Also  * [Onshape API Guide](https://onshape-public.github.io/docs/): Our full suite of developer guides, to be used as an accompaniment to this API Explorer. * [Onshape Developer Portal](https://dev-portal.onshape.com/): The Onshape portal for managing your API keys, OAuth2 credentials, your Onshape applications, and your Onshape App Store entries. * [Authentication Guide](https://onshape-public.github.io/docs/auth/): Our guide to using API keys, request signatures, and OAuth2 in  your Onshape applications.

Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// BTVariableStudioScopeInfo struct for BTVariableStudioScopeInfo
type BTVariableStudioScopeInfo struct {
	// Whether variable studio is automatically inserted into part studios and assemblies in the workspace
	IsAutomaticallyInserted bool `json:"isAutomaticallyInserted"`
}

// NewBTVariableStudioScopeInfo instantiates a new BTVariableStudioScopeInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTVariableStudioScopeInfo(isAutomaticallyInserted bool) *BTVariableStudioScopeInfo {
	this := BTVariableStudioScopeInfo{}
	this.IsAutomaticallyInserted = isAutomaticallyInserted
	return &this
}

// NewBTVariableStudioScopeInfoWithDefaults instantiates a new BTVariableStudioScopeInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTVariableStudioScopeInfoWithDefaults() *BTVariableStudioScopeInfo {
	this := BTVariableStudioScopeInfo{}
	return &this
}

// GetIsAutomaticallyInserted returns the IsAutomaticallyInserted field value
func (o *BTVariableStudioScopeInfo) GetIsAutomaticallyInserted() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsAutomaticallyInserted
}

// GetIsAutomaticallyInsertedOk returns a tuple with the IsAutomaticallyInserted field value
// and a boolean to check if the value has been set.
func (o *BTVariableStudioScopeInfo) GetIsAutomaticallyInsertedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsAutomaticallyInserted, true
}

// SetIsAutomaticallyInserted sets field value
func (o *BTVariableStudioScopeInfo) SetIsAutomaticallyInserted(v bool) {
	o.IsAutomaticallyInserted = v
}

func (o BTVariableStudioScopeInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["isAutomaticallyInserted"] = o.IsAutomaticallyInserted
	}
	return json.Marshal(toSerialize)
}

type NullableBTVariableStudioScopeInfo struct {
	value *BTVariableStudioScopeInfo
	isSet bool
}

func (v NullableBTVariableStudioScopeInfo) Get() *BTVariableStudioScopeInfo {
	return v.value
}

func (v *NullableBTVariableStudioScopeInfo) Set(val *BTVariableStudioScopeInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTVariableStudioScopeInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTVariableStudioScopeInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTVariableStudioScopeInfo(val *BTVariableStudioScopeInfo) *NullableBTVariableStudioScopeInfo {
	return &NullableBTVariableStudioScopeInfo{value: val, isSet: true}
}

func (v NullableBTVariableStudioScopeInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTVariableStudioScopeInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
