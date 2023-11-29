/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.173.26754-ceeaad064d4a
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
	"fmt"
)

// GBTCurveTypeEnum the model 'GBTCurveTypeEnum'
type GBTCurveTypeEnum string

// List of GBTCurveTypeEnum
const (
	GBTCurveTypeEnumOther   GBTCurveTypeEnum = "OTHER"
	GBTCurveTypeEnumLine    GBTCurveTypeEnum = "LINE"
	GBTCurveTypeEnumCircle  GBTCurveTypeEnum = "CIRCLE"
	GBTCurveTypeEnumEllipse GBTCurveTypeEnum = "ELLIPSE"
	GBTCurveTypeEnumBcurve  GBTCurveTypeEnum = "BCURVE"
	GBTCurveTypeEnumIcurve  GBTCurveTypeEnum = "ICURVE"
	GBTCurveTypeEnumUnknown GBTCurveTypeEnum = "UNKNOWN"
)

// All allowed values of GBTCurveTypeEnum enum
var AllowedGBTCurveTypeEnumEnumValues = []GBTCurveTypeEnum{
	"OTHER",
	"LINE",
	"CIRCLE",
	"ELLIPSE",
	"BCURVE",
	"ICURVE",
	"UNKNOWN",
}

func (v *GBTCurveTypeEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GBTCurveTypeEnum(value)
	for _, existing := range AllowedGBTCurveTypeEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GBTCurveTypeEnum", value)
}

// NewGBTCurveTypeEnumFromValue returns a pointer to a valid GBTCurveTypeEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGBTCurveTypeEnumFromValue(v string) (*GBTCurveTypeEnum, error) {
	ev := GBTCurveTypeEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GBTCurveTypeEnum: valid values are %v", v, AllowedGBTCurveTypeEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GBTCurveTypeEnum) IsValid() bool {
	for _, existing := range AllowedGBTCurveTypeEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GBTCurveTypeEnum value
func (v GBTCurveTypeEnum) Ptr() *GBTCurveTypeEnum {
	return &v
}

type NullableGBTCurveTypeEnum struct {
	value *GBTCurveTypeEnum
	isSet bool
}

func (v NullableGBTCurveTypeEnum) Get() *GBTCurveTypeEnum {
	return v.value
}

func (v *NullableGBTCurveTypeEnum) Set(val *GBTCurveTypeEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableGBTCurveTypeEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableGBTCurveTypeEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGBTCurveTypeEnum(val *GBTCurveTypeEnum) *NullableGBTCurveTypeEnum {
	return &NullableGBTCurveTypeEnum{value: val, isSet: true}
}

func (v NullableGBTCurveTypeEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGBTCurveTypeEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
