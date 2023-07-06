/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.166.18417-1bd990c6fbaa
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
	"fmt"
)

// GBTDataItemFormat the model 'GBTDataItemFormat'
type GBTDataItemFormat string

// List of GBTDataItemFormat
const (
	GBTDataItemFormatXT        GBTDataItemFormat = "X_T"
	GBTDataItemFormatXmm       GBTDataItemFormat = "XMM"
	GBTDataItemFormatStl       GBTDataItemFormat = "STL"
	GBTDataItemFormatOnshape   GBTDataItemFormat = "ONSHAPE"
	GBTDataItemFormatZip       GBTDataItemFormat = "ZIP"
	GBTDataItemFormatXTXmmZip  GBTDataItemFormat = "X_T_XMM_ZIP"
	GBTDataItemFormatObjMtlZip GBTDataItemFormat = "OBJ_MTL_ZIP"
	GBTDataItemFormatXB        GBTDataItemFormat = "X_B"
	GBTDataItemFormatUnknown   GBTDataItemFormat = "UNKNOWN"
)

// All allowed values of GBTDataItemFormat enum
var AllowedGBTDataItemFormatEnumValues = []GBTDataItemFormat{
	"X_T",
	"XMM",
	"STL",
	"ONSHAPE",
	"ZIP",
	"X_T_XMM_ZIP",
	"OBJ_MTL_ZIP",
	"X_B",
	"UNKNOWN",
}

func (v *GBTDataItemFormat) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GBTDataItemFormat(value)
	for _, existing := range AllowedGBTDataItemFormatEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GBTDataItemFormat", value)
}

// NewGBTDataItemFormatFromValue returns a pointer to a valid GBTDataItemFormat
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGBTDataItemFormatFromValue(v string) (*GBTDataItemFormat, error) {
	ev := GBTDataItemFormat(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GBTDataItemFormat: valid values are %v", v, AllowedGBTDataItemFormatEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GBTDataItemFormat) IsValid() bool {
	for _, existing := range AllowedGBTDataItemFormatEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GBTDataItemFormat value
func (v GBTDataItemFormat) Ptr() *GBTDataItemFormat {
	return &v
}

type NullableGBTDataItemFormat struct {
	value *GBTDataItemFormat
	isSet bool
}

func (v NullableGBTDataItemFormat) Get() *GBTDataItemFormat {
	return v.value
}

func (v *NullableGBTDataItemFormat) Set(val *GBTDataItemFormat) {
	v.value = val
	v.isSet = true
}

func (v NullableGBTDataItemFormat) IsSet() bool {
	return v.isSet
}

func (v *NullableGBTDataItemFormat) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGBTDataItemFormat(val *GBTDataItemFormat) *NullableGBTDataItemFormat {
	return &NullableGBTDataItemFormat{value: val, isSet: true}
}

func (v NullableGBTDataItemFormat) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGBTDataItemFormat) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
