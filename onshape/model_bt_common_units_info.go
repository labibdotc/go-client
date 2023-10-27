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

// BTCommonUnitsInfo struct for BTCommonUnitsInfo
type BTCommonUnitsInfo struct {
	QuantityTypeToBaseUnits *map[string]map[string]int32 `json:"quantityTypeToBaseUnits,omitempty"`
	Units                   []BTCommonUnitInfo           `json:"units,omitempty"`
}

// NewBTCommonUnitsInfo instantiates a new BTCommonUnitsInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTCommonUnitsInfo() *BTCommonUnitsInfo {
	this := BTCommonUnitsInfo{}
	return &this
}

// NewBTCommonUnitsInfoWithDefaults instantiates a new BTCommonUnitsInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTCommonUnitsInfoWithDefaults() *BTCommonUnitsInfo {
	this := BTCommonUnitsInfo{}
	return &this
}

// GetQuantityTypeToBaseUnits returns the QuantityTypeToBaseUnits field value if set, zero value otherwise.
func (o *BTCommonUnitsInfo) GetQuantityTypeToBaseUnits() map[string]map[string]int32 {
	if o == nil || o.QuantityTypeToBaseUnits == nil {
		var ret map[string]map[string]int32
		return ret
	}
	return *o.QuantityTypeToBaseUnits
}

// GetQuantityTypeToBaseUnitsOk returns a tuple with the QuantityTypeToBaseUnits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCommonUnitsInfo) GetQuantityTypeToBaseUnitsOk() (*map[string]map[string]int32, bool) {
	if o == nil || o.QuantityTypeToBaseUnits == nil {
		return nil, false
	}
	return o.QuantityTypeToBaseUnits, true
}

// HasQuantityTypeToBaseUnits returns a boolean if a field has been set.
func (o *BTCommonUnitsInfo) HasQuantityTypeToBaseUnits() bool {
	if o != nil && o.QuantityTypeToBaseUnits != nil {
		return true
	}

	return false
}

// SetQuantityTypeToBaseUnits gets a reference to the given map[string]map[string]int32 and assigns it to the QuantityTypeToBaseUnits field.
func (o *BTCommonUnitsInfo) SetQuantityTypeToBaseUnits(v map[string]map[string]int32) {
	o.QuantityTypeToBaseUnits = &v
}

// GetUnits returns the Units field value if set, zero value otherwise.
func (o *BTCommonUnitsInfo) GetUnits() []BTCommonUnitInfo {
	if o == nil || o.Units == nil {
		var ret []BTCommonUnitInfo
		return ret
	}
	return o.Units
}

// GetUnitsOk returns a tuple with the Units field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCommonUnitsInfo) GetUnitsOk() ([]BTCommonUnitInfo, bool) {
	if o == nil || o.Units == nil {
		return nil, false
	}
	return o.Units, true
}

// HasUnits returns a boolean if a field has been set.
func (o *BTCommonUnitsInfo) HasUnits() bool {
	if o != nil && o.Units != nil {
		return true
	}

	return false
}

// SetUnits gets a reference to the given []BTCommonUnitInfo and assigns it to the Units field.
func (o *BTCommonUnitsInfo) SetUnits(v []BTCommonUnitInfo) {
	o.Units = v
}

func (o BTCommonUnitsInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.QuantityTypeToBaseUnits != nil {
		toSerialize["quantityTypeToBaseUnits"] = o.QuantityTypeToBaseUnits
	}
	if o.Units != nil {
		toSerialize["units"] = o.Units
	}
	return json.Marshal(toSerialize)
}

type NullableBTCommonUnitsInfo struct {
	value *BTCommonUnitsInfo
	isSet bool
}

func (v NullableBTCommonUnitsInfo) Get() *BTCommonUnitsInfo {
	return v.value
}

func (v *NullableBTCommonUnitsInfo) Set(val *BTCommonUnitsInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTCommonUnitsInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTCommonUnitsInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTCommonUnitsInfo(val *BTCommonUnitsInfo) *NullableBTCommonUnitsInfo {
	return &NullableBTCommonUnitsInfo{value: val, isSet: true}
}

func (v NullableBTCommonUnitsInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTCommonUnitsInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
