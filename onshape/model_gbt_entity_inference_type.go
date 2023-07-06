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

// GBTEntityInferenceType the model 'GBTEntityInferenceType'
type GBTEntityInferenceType string

// List of GBTEntityInferenceType
const (
	GBTEntityInferenceTypePartOrigin      GBTEntityInferenceType = "PART_ORIGIN"
	GBTEntityInferenceTypePoint           GBTEntityInferenceType = "POINT"
	GBTEntityInferenceTypeCentroid        GBTEntityInferenceType = "CENTROID"
	GBTEntityInferenceTypeCenter          GBTEntityInferenceType = "CENTER"
	GBTEntityInferenceTypeMidPoint        GBTEntityInferenceType = "MID_POINT"
	GBTEntityInferenceTypeTopAxisPoint    GBTEntityInferenceType = "TOP_AXIS_POINT"
	GBTEntityInferenceTypeMidAxisPoint    GBTEntityInferenceType = "MID_AXIS_POINT"
	GBTEntityInferenceTypeBottomAxisPoint GBTEntityInferenceType = "BOTTOM_AXIS_POINT"
	GBTEntityInferenceTypeOriginX         GBTEntityInferenceType = "ORIGIN_X"
	GBTEntityInferenceTypeOriginY         GBTEntityInferenceType = "ORIGIN_Y"
	GBTEntityInferenceTypeOriginZ         GBTEntityInferenceType = "ORIGIN_Z"
	GBTEntityInferenceTypeLoopCenter      GBTEntityInferenceType = "LOOP_CENTER"
	GBTEntityInferenceTypeUnknown         GBTEntityInferenceType = "UNKNOWN"
)

// All allowed values of GBTEntityInferenceType enum
var AllowedGBTEntityInferenceTypeEnumValues = []GBTEntityInferenceType{
	"PART_ORIGIN",
	"POINT",
	"CENTROID",
	"CENTER",
	"MID_POINT",
	"TOP_AXIS_POINT",
	"MID_AXIS_POINT",
	"BOTTOM_AXIS_POINT",
	"ORIGIN_X",
	"ORIGIN_Y",
	"ORIGIN_Z",
	"LOOP_CENTER",
	"UNKNOWN",
}

func (v *GBTEntityInferenceType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GBTEntityInferenceType(value)
	for _, existing := range AllowedGBTEntityInferenceTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GBTEntityInferenceType", value)
}

// NewGBTEntityInferenceTypeFromValue returns a pointer to a valid GBTEntityInferenceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGBTEntityInferenceTypeFromValue(v string) (*GBTEntityInferenceType, error) {
	ev := GBTEntityInferenceType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GBTEntityInferenceType: valid values are %v", v, AllowedGBTEntityInferenceTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GBTEntityInferenceType) IsValid() bool {
	for _, existing := range AllowedGBTEntityInferenceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GBTEntityInferenceType value
func (v GBTEntityInferenceType) Ptr() *GBTEntityInferenceType {
	return &v
}

type NullableGBTEntityInferenceType struct {
	value *GBTEntityInferenceType
	isSet bool
}

func (v NullableGBTEntityInferenceType) Get() *GBTEntityInferenceType {
	return v.value
}

func (v *NullableGBTEntityInferenceType) Set(val *GBTEntityInferenceType) {
	v.value = val
	v.isSet = true
}

func (v NullableGBTEntityInferenceType) IsSet() bool {
	return v.isSet
}

func (v *NullableGBTEntityInferenceType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGBTEntityInferenceType(val *GBTEntityInferenceType) *NullableGBTEntityInferenceType {
	return &NullableGBTEntityInferenceType{value: val, isSet: true}
}

func (v NullableGBTEntityInferenceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGBTEntityInferenceType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
