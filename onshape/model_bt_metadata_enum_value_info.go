/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.172.25784-25d5580b0721
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTMetadataEnumValueInfo struct for BTMetadataEnumValueInfo
type BTMetadataEnumValueInfo struct {
	Label *string `json:"label,omitempty"`
	State *int32  `json:"state,omitempty"`
	Value *string `json:"value,omitempty"`
}

// NewBTMetadataEnumValueInfo instantiates a new BTMetadataEnumValueInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTMetadataEnumValueInfo() *BTMetadataEnumValueInfo {
	this := BTMetadataEnumValueInfo{}
	return &this
}

// NewBTMetadataEnumValueInfoWithDefaults instantiates a new BTMetadataEnumValueInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTMetadataEnumValueInfoWithDefaults() *BTMetadataEnumValueInfo {
	this := BTMetadataEnumValueInfo{}
	return &this
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *BTMetadataEnumValueInfo) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMetadataEnumValueInfo) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *BTMetadataEnumValueInfo) HasLabel() bool {
	if o != nil && o.Label != nil {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *BTMetadataEnumValueInfo) SetLabel(v string) {
	o.Label = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *BTMetadataEnumValueInfo) GetState() int32 {
	if o == nil || o.State == nil {
		var ret int32
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMetadataEnumValueInfo) GetStateOk() (*int32, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *BTMetadataEnumValueInfo) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given int32 and assigns it to the State field.
func (o *BTMetadataEnumValueInfo) SetState(v int32) {
	o.State = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *BTMetadataEnumValueInfo) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMetadataEnumValueInfo) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *BTMetadataEnumValueInfo) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *BTMetadataEnumValueInfo) SetValue(v string) {
	o.Value = &v
}

func (o BTMetadataEnumValueInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Label != nil {
		toSerialize["label"] = o.Label
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableBTMetadataEnumValueInfo struct {
	value *BTMetadataEnumValueInfo
	isSet bool
}

func (v NullableBTMetadataEnumValueInfo) Get() *BTMetadataEnumValueInfo {
	return v.value
}

func (v *NullableBTMetadataEnumValueInfo) Set(val *BTMetadataEnumValueInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTMetadataEnumValueInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTMetadataEnumValueInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTMetadataEnumValueInfo(val *BTMetadataEnumValueInfo) *NullableBTMetadataEnumValueInfo {
	return &NullableBTMetadataEnumValueInfo{value: val, isSet: true}
}

func (v NullableBTMetadataEnumValueInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTMetadataEnumValueInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
