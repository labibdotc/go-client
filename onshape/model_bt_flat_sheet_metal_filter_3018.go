/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.167.20169-88260985a0b6
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTFlatSheetMetalFilter3018 struct for BTFlatSheetMetalFilter3018
type BTFlatSheetMetalFilter3018 struct {
	BtType *string                            `json:"btType,omitempty"`
	Allows *GBTFilterFlattenedGeometryOptions `json:"allows,omitempty"`
}

// NewBTFlatSheetMetalFilter3018 instantiates a new BTFlatSheetMetalFilter3018 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTFlatSheetMetalFilter3018() *BTFlatSheetMetalFilter3018 {
	this := BTFlatSheetMetalFilter3018{}
	return &this
}

// NewBTFlatSheetMetalFilter3018WithDefaults instantiates a new BTFlatSheetMetalFilter3018 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTFlatSheetMetalFilter3018WithDefaults() *BTFlatSheetMetalFilter3018 {
	this := BTFlatSheetMetalFilter3018{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTFlatSheetMetalFilter3018) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFlatSheetMetalFilter3018) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTFlatSheetMetalFilter3018) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTFlatSheetMetalFilter3018) SetBtType(v string) {
	o.BtType = &v
}

// GetAllows returns the Allows field value if set, zero value otherwise.
func (o *BTFlatSheetMetalFilter3018) GetAllows() GBTFilterFlattenedGeometryOptions {
	if o == nil || o.Allows == nil {
		var ret GBTFilterFlattenedGeometryOptions
		return ret
	}
	return *o.Allows
}

// GetAllowsOk returns a tuple with the Allows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFlatSheetMetalFilter3018) GetAllowsOk() (*GBTFilterFlattenedGeometryOptions, bool) {
	if o == nil || o.Allows == nil {
		return nil, false
	}
	return o.Allows, true
}

// HasAllows returns a boolean if a field has been set.
func (o *BTFlatSheetMetalFilter3018) HasAllows() bool {
	if o != nil && o.Allows != nil {
		return true
	}

	return false
}

// SetAllows gets a reference to the given GBTFilterFlattenedGeometryOptions and assigns it to the Allows field.
func (o *BTFlatSheetMetalFilter3018) SetAllows(v GBTFilterFlattenedGeometryOptions) {
	o.Allows = &v
}

func (o BTFlatSheetMetalFilter3018) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.Allows != nil {
		toSerialize["allows"] = o.Allows
	}
	return json.Marshal(toSerialize)
}

type NullableBTFlatSheetMetalFilter3018 struct {
	value *BTFlatSheetMetalFilter3018
	isSet bool
}

func (v NullableBTFlatSheetMetalFilter3018) Get() *BTFlatSheetMetalFilter3018 {
	return v.value
}

func (v *NullableBTFlatSheetMetalFilter3018) Set(val *BTFlatSheetMetalFilter3018) {
	v.value = val
	v.isSet = true
}

func (v NullableBTFlatSheetMetalFilter3018) IsSet() bool {
	return v.isSet
}

func (v *NullableBTFlatSheetMetalFilter3018) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTFlatSheetMetalFilter3018(val *BTFlatSheetMetalFilter3018) *NullableBTFlatSheetMetalFilter3018 {
	return &NullableBTFlatSheetMetalFilter3018{value: val, isSet: true}
}

func (v NullableBTFlatSheetMetalFilter3018) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTFlatSheetMetalFilter3018) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
