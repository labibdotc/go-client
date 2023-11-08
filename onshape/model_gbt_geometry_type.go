/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.172.25541-b54c5b71ef45
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
	"fmt"
)

// GBTGeometryType the model 'GBTGeometryType'
type GBTGeometryType string

// List of GBTGeometryType
const (
	GBTGeometryTypeLine                 GBTGeometryType = "LINE"
	GBTGeometryTypeCircle               GBTGeometryType = "CIRCLE"
	GBTGeometryTypeArc                  GBTGeometryType = "ARC"
	GBTGeometryTypePlane                GBTGeometryType = "PLANE"
	GBTGeometryTypeCylinder             GBTGeometryType = "CYLINDER"
	GBTGeometryTypeCone                 GBTGeometryType = "CONE"
	GBTGeometryTypeSphere               GBTGeometryType = "SPHERE"
	GBTGeometryTypeTorus                GBTGeometryType = "TORUS"
	GBTGeometryTypeSpline               GBTGeometryType = "SPLINE"
	GBTGeometryTypeEllipse              GBTGeometryType = "ELLIPSE"
	GBTGeometryTypeMesh                 GBTGeometryType = "MESH"
	GBTGeometryTypeConic                GBTGeometryType = "CONIC"
	GBTGeometryTypeRevolved             GBTGeometryType = "REVOLVED"
	GBTGeometryTypeExtruded             GBTGeometryType = "EXTRUDED"
	GBTGeometryTypeAllMesh              GBTGeometryType = "ALL_MESH"
	GBTGeometryTypeMixedMesh            GBTGeometryType = "MIXED_MESH"
	GBTGeometryTypeSplineInternalPoint  GBTGeometryType = "SPLINE_INTERNAL_POINT"
	GBTGeometryTypeSplineControlPolygon GBTGeometryType = "SPLINE_CONTROL_POLYGON"
	GBTGeometryTypeEllipticalArc        GBTGeometryType = "ELLIPTICAL_ARC"
	GBTGeometryTypeUnknown              GBTGeometryType = "UNKNOWN"
)

// All allowed values of GBTGeometryType enum
var AllowedGBTGeometryTypeEnumValues = []GBTGeometryType{
	"LINE",
	"CIRCLE",
	"ARC",
	"PLANE",
	"CYLINDER",
	"CONE",
	"SPHERE",
	"TORUS",
	"SPLINE",
	"ELLIPSE",
	"MESH",
	"CONIC",
	"REVOLVED",
	"EXTRUDED",
	"ALL_MESH",
	"MIXED_MESH",
	"SPLINE_INTERNAL_POINT",
	"SPLINE_CONTROL_POLYGON",
	"ELLIPTICAL_ARC",
	"UNKNOWN",
}

func (v *GBTGeometryType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GBTGeometryType(value)
	for _, existing := range AllowedGBTGeometryTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GBTGeometryType", value)
}

// NewGBTGeometryTypeFromValue returns a pointer to a valid GBTGeometryType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGBTGeometryTypeFromValue(v string) (*GBTGeometryType, error) {
	ev := GBTGeometryType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GBTGeometryType: valid values are %v", v, AllowedGBTGeometryTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GBTGeometryType) IsValid() bool {
	for _, existing := range AllowedGBTGeometryTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GBTGeometryType value
func (v GBTGeometryType) Ptr() *GBTGeometryType {
	return &v
}

type NullableGBTGeometryType struct {
	value *GBTGeometryType
	isSet bool
}

func (v NullableGBTGeometryType) Get() *GBTGeometryType {
	return v.value
}

func (v *NullableGBTGeometryType) Set(val *GBTGeometryType) {
	v.value = val
	v.isSet = true
}

func (v NullableGBTGeometryType) IsSet() bool {
	return v.isSet
}

func (v *NullableGBTGeometryType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGBTGeometryType(val *GBTGeometryType) *NullableGBTGeometryType {
	return &NullableGBTGeometryType{value: val, isSet: true}
}

func (v NullableGBTGeometryType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGBTGeometryType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
