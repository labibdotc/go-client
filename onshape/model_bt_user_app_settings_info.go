/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.171.24804-920f3dc76f2b
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTUserAppSettingsInfo struct for BTUserAppSettingsInfo
type BTUserAppSettingsInfo struct {
	Settings []BTSettingInfo `json:"settings,omitempty"`
}

// NewBTUserAppSettingsInfo instantiates a new BTUserAppSettingsInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTUserAppSettingsInfo() *BTUserAppSettingsInfo {
	this := BTUserAppSettingsInfo{}
	return &this
}

// NewBTUserAppSettingsInfoWithDefaults instantiates a new BTUserAppSettingsInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTUserAppSettingsInfoWithDefaults() *BTUserAppSettingsInfo {
	this := BTUserAppSettingsInfo{}
	return &this
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *BTUserAppSettingsInfo) GetSettings() []BTSettingInfo {
	if o == nil || o.Settings == nil {
		var ret []BTSettingInfo
		return ret
	}
	return o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserAppSettingsInfo) GetSettingsOk() ([]BTSettingInfo, bool) {
	if o == nil || o.Settings == nil {
		return nil, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *BTUserAppSettingsInfo) HasSettings() bool {
	if o != nil && o.Settings != nil {
		return true
	}

	return false
}

// SetSettings gets a reference to the given []BTSettingInfo and assigns it to the Settings field.
func (o *BTUserAppSettingsInfo) SetSettings(v []BTSettingInfo) {
	o.Settings = v
}

func (o BTUserAppSettingsInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Settings != nil {
		toSerialize["settings"] = o.Settings
	}
	return json.Marshal(toSerialize)
}

type NullableBTUserAppSettingsInfo struct {
	value *BTUserAppSettingsInfo
	isSet bool
}

func (v NullableBTUserAppSettingsInfo) Get() *BTUserAppSettingsInfo {
	return v.value
}

func (v *NullableBTUserAppSettingsInfo) Set(val *BTUserAppSettingsInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTUserAppSettingsInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTUserAppSettingsInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTUserAppSettingsInfo(val *BTUserAppSettingsInfo) *NullableBTUserAppSettingsInfo {
	return &NullableBTUserAppSettingsInfo{value: val, isSet: true}
}

func (v NullableBTUserAppSettingsInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTUserAppSettingsInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
