/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.154.6700-f2963ce28904
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTActiveSheetMetalFilter2944 struct for BTActiveSheetMetalFilter2944
type BTActiveSheetMetalFilter2944 struct {
	BtType                 *string `json:"btType,omitempty"`
	IsFromActiveSheetMetal *bool   `json:"isFromActiveSheetMetal,omitempty"`
}

// NewBTActiveSheetMetalFilter2944 instantiates a new BTActiveSheetMetalFilter2944 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTActiveSheetMetalFilter2944() *BTActiveSheetMetalFilter2944 {
	this := BTActiveSheetMetalFilter2944{}
	return &this
}

// NewBTActiveSheetMetalFilter2944WithDefaults instantiates a new BTActiveSheetMetalFilter2944 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTActiveSheetMetalFilter2944WithDefaults() *BTActiveSheetMetalFilter2944 {
	this := BTActiveSheetMetalFilter2944{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTActiveSheetMetalFilter2944) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTActiveSheetMetalFilter2944) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTActiveSheetMetalFilter2944) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTActiveSheetMetalFilter2944) SetBtType(v string) {
	o.BtType = &v
}

// GetIsFromActiveSheetMetal returns the IsFromActiveSheetMetal field value if set, zero value otherwise.
func (o *BTActiveSheetMetalFilter2944) GetIsFromActiveSheetMetal() bool {
	if o == nil || o.IsFromActiveSheetMetal == nil {
		var ret bool
		return ret
	}
	return *o.IsFromActiveSheetMetal
}

// GetIsFromActiveSheetMetalOk returns a tuple with the IsFromActiveSheetMetal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTActiveSheetMetalFilter2944) GetIsFromActiveSheetMetalOk() (*bool, bool) {
	if o == nil || o.IsFromActiveSheetMetal == nil {
		return nil, false
	}
	return o.IsFromActiveSheetMetal, true
}

// HasIsFromActiveSheetMetal returns a boolean if a field has been set.
func (o *BTActiveSheetMetalFilter2944) HasIsFromActiveSheetMetal() bool {
	if o != nil && o.IsFromActiveSheetMetal != nil {
		return true
	}

	return false
}

// SetIsFromActiveSheetMetal gets a reference to the given bool and assigns it to the IsFromActiveSheetMetal field.
func (o *BTActiveSheetMetalFilter2944) SetIsFromActiveSheetMetal(v bool) {
	o.IsFromActiveSheetMetal = &v
}

func (o BTActiveSheetMetalFilter2944) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.IsFromActiveSheetMetal != nil {
		toSerialize["isFromActiveSheetMetal"] = o.IsFromActiveSheetMetal
	}
	return json.Marshal(toSerialize)
}

type NullableBTActiveSheetMetalFilter2944 struct {
	value *BTActiveSheetMetalFilter2944
	isSet bool
}

func (v NullableBTActiveSheetMetalFilter2944) Get() *BTActiveSheetMetalFilter2944 {
	return v.value
}

func (v *NullableBTActiveSheetMetalFilter2944) Set(val *BTActiveSheetMetalFilter2944) {
	v.value = val
	v.isSet = true
}

func (v NullableBTActiveSheetMetalFilter2944) IsSet() bool {
	return v.isSet
}

func (v *NullableBTActiveSheetMetalFilter2944) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTActiveSheetMetalFilter2944(val *BTActiveSheetMetalFilter2944) *NullableBTActiveSheetMetalFilter2944 {
	return &NullableBTActiveSheetMetalFilter2944{value: val, isSet: true}
}

func (v NullableBTActiveSheetMetalFilter2944) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTActiveSheetMetalFilter2944) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
