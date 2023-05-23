/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.164.16301-d273853a12e7
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
	"fmt"
)

// GBTInsertableType the model 'GBTInsertableType'
type GBTInsertableType string

// List of GBTInsertableType
const (
	GBTInsertableTypeParts          GBTInsertableType = "PARTS"
	GBTInsertableTypeSketches       GBTInsertableType = "SKETCHES"
	GBTInsertableTypeSurfaces       GBTInsertableType = "SURFACES"
	GBTInsertableTypeFlattenedParts GBTInsertableType = "FLATTENED_PARTS"
	GBTInsertableTypeCompositeParts GBTInsertableType = "COMPOSITE_PARTS"
	GBTInsertableTypePartStudios    GBTInsertableType = "PART_STUDIOS"
	GBTInsertableTypeWires          GBTInsertableType = "WIRES"
	GBTInsertableTypeUnknown        GBTInsertableType = "UNKNOWN"
)

// All allowed values of GBTInsertableType enum
var AllowedGBTInsertableTypeEnumValues = []GBTInsertableType{
	"PARTS",
	"SKETCHES",
	"SURFACES",
	"FLATTENED_PARTS",
	"COMPOSITE_PARTS",
	"PART_STUDIOS",
	"WIRES",
	"UNKNOWN",
}

func (v *GBTInsertableType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GBTInsertableType(value)
	for _, existing := range AllowedGBTInsertableTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GBTInsertableType", value)
}

// NewGBTInsertableTypeFromValue returns a pointer to a valid GBTInsertableType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGBTInsertableTypeFromValue(v string) (*GBTInsertableType, error) {
	ev := GBTInsertableType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GBTInsertableType: valid values are %v", v, AllowedGBTInsertableTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GBTInsertableType) IsValid() bool {
	for _, existing := range AllowedGBTInsertableTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GBTInsertableType value
func (v GBTInsertableType) Ptr() *GBTInsertableType {
	return &v
}

type NullableGBTInsertableType struct {
	value *GBTInsertableType
	isSet bool
}

func (v NullableGBTInsertableType) Get() *GBTInsertableType {
	return v.value
}

func (v *NullableGBTInsertableType) Set(val *GBTInsertableType) {
	v.value = val
	v.isSet = true
}

func (v NullableGBTInsertableType) IsSet() bool {
	return v.isSet
}

func (v *NullableGBTInsertableType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGBTInsertableType(val *GBTInsertableType) *NullableGBTInsertableType {
	return &NullableGBTInsertableType{value: val, isSet: true}
}

func (v NullableGBTInsertableType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGBTInsertableType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
