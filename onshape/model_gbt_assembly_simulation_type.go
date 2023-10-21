/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.171.24440-f37a82fd6450
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
	"fmt"
)

// GBTAssemblySimulationType the model 'GBTAssemblySimulationType'
type GBTAssemblySimulationType string

// List of GBTAssemblySimulationType
const (
	GBTAssemblySimulationTypeLinearStatic    GBTAssemblySimulationType = "LINEAR_STATIC"
	GBTAssemblySimulationTypeModal           GBTAssemblySimulationType = "MODAL"
	GBTAssemblySimulationTypeContactAnalysis GBTAssemblySimulationType = "CONTACT_ANALYSIS"
	GBTAssemblySimulationTypeUnknown         GBTAssemblySimulationType = "UNKNOWN"
)

// All allowed values of GBTAssemblySimulationType enum
var AllowedGBTAssemblySimulationTypeEnumValues = []GBTAssemblySimulationType{
	"LINEAR_STATIC",
	"MODAL",
	"CONTACT_ANALYSIS",
	"UNKNOWN",
}

func (v *GBTAssemblySimulationType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GBTAssemblySimulationType(value)
	for _, existing := range AllowedGBTAssemblySimulationTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GBTAssemblySimulationType", value)
}

// NewGBTAssemblySimulationTypeFromValue returns a pointer to a valid GBTAssemblySimulationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGBTAssemblySimulationTypeFromValue(v string) (*GBTAssemblySimulationType, error) {
	ev := GBTAssemblySimulationType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GBTAssemblySimulationType: valid values are %v", v, AllowedGBTAssemblySimulationTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GBTAssemblySimulationType) IsValid() bool {
	for _, existing := range AllowedGBTAssemblySimulationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GBTAssemblySimulationType value
func (v GBTAssemblySimulationType) Ptr() *GBTAssemblySimulationType {
	return &v
}

type NullableGBTAssemblySimulationType struct {
	value *GBTAssemblySimulationType
	isSet bool
}

func (v NullableGBTAssemblySimulationType) Get() *GBTAssemblySimulationType {
	return v.value
}

func (v *NullableGBTAssemblySimulationType) Set(val *GBTAssemblySimulationType) {
	v.value = val
	v.isSet = true
}

func (v NullableGBTAssemblySimulationType) IsSet() bool {
	return v.isSet
}

func (v *NullableGBTAssemblySimulationType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGBTAssemblySimulationType(val *GBTAssemblySimulationType) *NullableGBTAssemblySimulationType {
	return &NullableGBTAssemblySimulationType{value: val, isSet: true}
}

func (v NullableGBTAssemblySimulationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGBTAssemblySimulationType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
