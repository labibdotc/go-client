/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.164.16643-ef813b2da145
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
	"fmt"
)

// GBTSketchCurveType the model 'GBTSketchCurveType'
type GBTSketchCurveType string

// List of GBTSketchCurveType
const (
	GBTSketchCurveTypeUnset              GBTSketchCurveType = "UNSET"
	GBTSketchCurveTypeInterpolatedSpline GBTSketchCurveType = "INTERPOLATED_SPLINE"
	GBTSketchCurveTypeBezierCurve        GBTSketchCurveType = "BEZIER_CURVE"
	GBTSketchCurveTypeUnknown            GBTSketchCurveType = "UNKNOWN"
)

// All allowed values of GBTSketchCurveType enum
var AllowedGBTSketchCurveTypeEnumValues = []GBTSketchCurveType{
	"UNSET",
	"INTERPOLATED_SPLINE",
	"BEZIER_CURVE",
	"UNKNOWN",
}

func (v *GBTSketchCurveType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GBTSketchCurveType(value)
	for _, existing := range AllowedGBTSketchCurveTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GBTSketchCurveType", value)
}

// NewGBTSketchCurveTypeFromValue returns a pointer to a valid GBTSketchCurveType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGBTSketchCurveTypeFromValue(v string) (*GBTSketchCurveType, error) {
	ev := GBTSketchCurveType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GBTSketchCurveType: valid values are %v", v, AllowedGBTSketchCurveTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GBTSketchCurveType) IsValid() bool {
	for _, existing := range AllowedGBTSketchCurveTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GBTSketchCurveType value
func (v GBTSketchCurveType) Ptr() *GBTSketchCurveType {
	return &v
}

type NullableGBTSketchCurveType struct {
	value *GBTSketchCurveType
	isSet bool
}

func (v NullableGBTSketchCurveType) Get() *GBTSketchCurveType {
	return v.value
}

func (v *NullableGBTSketchCurveType) Set(val *GBTSketchCurveType) {
	v.value = val
	v.isSet = true
}

func (v NullableGBTSketchCurveType) IsSet() bool {
	return v.isSet
}

func (v *NullableGBTSketchCurveType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGBTSketchCurveType(val *GBTSketchCurveType) *NullableGBTSketchCurveType {
	return &NullableGBTSketchCurveType{value: val, isSet: true}
}

func (v NullableGBTSketchCurveType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGBTSketchCurveType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
