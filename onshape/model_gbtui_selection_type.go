/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.165.18120-f464f720d215
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
	"fmt"
)

// GBTUiSelectionType the model 'GBTUiSelectionType'
type GBTUiSelectionType string

// List of GBTUiSelectionType
const (
	GBTUiSelectionTypeEntity                GBTUiSelectionType = "ENTITY"
	GBTUiSelectionTypeFeature               GBTUiSelectionType = "FEATURE"
	GBTUiSelectionTypeBody                  GBTUiSelectionType = "BODY"
	GBTUiSelectionTypeOccurrence            GBTUiSelectionType = "OCCURRENCE"
	GBTUiSelectionTypeUsercode              GBTUiSelectionType = "USERCODE"
	GBTUiSelectionTypeRollbackbar           GBTUiSelectionType = "ROLLBACKBAR"
	GBTUiSelectionTypeElement               GBTUiSelectionType = "ELEMENT"
	GBTUiSelectionTypeMate                  GBTUiSelectionType = "MATE"
	GBTUiSelectionTypeMateConnector         GBTUiSelectionType = "MATE_CONNECTOR"
	GBTUiSelectionTypeEdgePoint             GBTUiSelectionType = "EDGE_POINT"
	GBTUiSelectionTypeMeshPoint             GBTUiSelectionType = "MESH_POINT"
	GBTUiSelectionTypeTableItem             GBTUiSelectionType = "TABLE_ITEM"
	GBTUiSelectionTypeSketchGroup           GBTUiSelectionType = "SKETCH_GROUP"
	GBTUiSelectionTypeFolder                GBTUiSelectionType = "FOLDER"
	GBTUiSelectionTypeNonGeometricItem      GBTUiSelectionType = "NON_GEOMETRIC_ITEM"
	GBTUiSelectionTypeTemporaryGeometry     GBTUiSelectionType = "TEMPORARY_GEOMETRY"
	GBTUiSelectionTypeProperty              GBTUiSelectionType = "PROPERTY"
	GBTUiSelectionTypeSimulationLoad        GBTUiSelectionType = "SIMULATION_LOAD"
	GBTUiSelectionTypePersistentQueryString GBTUiSelectionType = "PERSISTENT_QUERY_STRING"
	GBTUiSelectionTypeUnknown               GBTUiSelectionType = "UNKNOWN"
)

// All allowed values of GBTUiSelectionType enum
var AllowedGBTUiSelectionTypeEnumValues = []GBTUiSelectionType{
	"ENTITY",
	"FEATURE",
	"BODY",
	"OCCURRENCE",
	"USERCODE",
	"ROLLBACKBAR",
	"ELEMENT",
	"MATE",
	"MATE_CONNECTOR",
	"EDGE_POINT",
	"MESH_POINT",
	"TABLE_ITEM",
	"SKETCH_GROUP",
	"FOLDER",
	"NON_GEOMETRIC_ITEM",
	"TEMPORARY_GEOMETRY",
	"PROPERTY",
	"SIMULATION_LOAD",
	"PERSISTENT_QUERY_STRING",
	"UNKNOWN",
}

func (v *GBTUiSelectionType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GBTUiSelectionType(value)
	for _, existing := range AllowedGBTUiSelectionTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GBTUiSelectionType", value)
}

// NewGBTUiSelectionTypeFromValue returns a pointer to a valid GBTUiSelectionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGBTUiSelectionTypeFromValue(v string) (*GBTUiSelectionType, error) {
	ev := GBTUiSelectionType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GBTUiSelectionType: valid values are %v", v, AllowedGBTUiSelectionTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GBTUiSelectionType) IsValid() bool {
	for _, existing := range AllowedGBTUiSelectionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GBTUiSelectionType value
func (v GBTUiSelectionType) Ptr() *GBTUiSelectionType {
	return &v
}

type NullableGBTUiSelectionType struct {
	value *GBTUiSelectionType
	isSet bool
}

func (v NullableGBTUiSelectionType) Get() *GBTUiSelectionType {
	return v.value
}

func (v *NullableGBTUiSelectionType) Set(val *GBTUiSelectionType) {
	v.value = val
	v.isSet = true
}

func (v NullableGBTUiSelectionType) IsSet() bool {
	return v.isSet
}

func (v *NullableGBTUiSelectionType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGBTUiSelectionType(val *GBTUiSelectionType) *NullableGBTUiSelectionType {
	return &NullableGBTUiSelectionType{value: val, isSet: true}
}

func (v NullableGBTUiSelectionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGBTUiSelectionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
