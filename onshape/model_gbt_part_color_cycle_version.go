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

// GBTPartColorCycleVersion the model 'GBTPartColorCycleVersion'
type GBTPartColorCycleVersion string

// List of GBTPartColorCycleVersion
const (
	GBTPartColorCycleVersionColorCycle0 GBTPartColorCycleVersion = "COLOR_CYCLE_0"
	GBTPartColorCycleVersionColorCycle1 GBTPartColorCycleVersion = "COLOR_CYCLE_1"
	GBTPartColorCycleVersionUnknown     GBTPartColorCycleVersion = "UNKNOWN"
)

// All allowed values of GBTPartColorCycleVersion enum
var AllowedGBTPartColorCycleVersionEnumValues = []GBTPartColorCycleVersion{
	"COLOR_CYCLE_0",
	"COLOR_CYCLE_1",
	"UNKNOWN",
}

func (v *GBTPartColorCycleVersion) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GBTPartColorCycleVersion(value)
	for _, existing := range AllowedGBTPartColorCycleVersionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GBTPartColorCycleVersion", value)
}

// NewGBTPartColorCycleVersionFromValue returns a pointer to a valid GBTPartColorCycleVersion
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGBTPartColorCycleVersionFromValue(v string) (*GBTPartColorCycleVersion, error) {
	ev := GBTPartColorCycleVersion(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GBTPartColorCycleVersion: valid values are %v", v, AllowedGBTPartColorCycleVersionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GBTPartColorCycleVersion) IsValid() bool {
	for _, existing := range AllowedGBTPartColorCycleVersionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GBTPartColorCycleVersion value
func (v GBTPartColorCycleVersion) Ptr() *GBTPartColorCycleVersion {
	return &v
}

type NullableGBTPartColorCycleVersion struct {
	value *GBTPartColorCycleVersion
	isSet bool
}

func (v NullableGBTPartColorCycleVersion) Get() *GBTPartColorCycleVersion {
	return v.value
}

func (v *NullableGBTPartColorCycleVersion) Set(val *GBTPartColorCycleVersion) {
	v.value = val
	v.isSet = true
}

func (v NullableGBTPartColorCycleVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableGBTPartColorCycleVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGBTPartColorCycleVersion(val *GBTPartColorCycleVersion) *NullableGBTPartColorCycleVersion {
	return &NullableGBTPartColorCycleVersion{value: val, isSet: true}
}

func (v NullableGBTPartColorCycleVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGBTPartColorCycleVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
