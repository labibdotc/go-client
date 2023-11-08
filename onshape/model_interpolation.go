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

// Interpolation the model 'Interpolation'
type Interpolation string

// List of Interpolation
const (
	InterpolationStep             Interpolation = "STEP"
	InterpolationLinear           Interpolation = "LINEAR"
	InterpolationCatmullromspline Interpolation = "CATMULLROMSPLINE"
	InterpolationCubicspline      Interpolation = "CUBICSPLINE"
)

// All allowed values of Interpolation enum
var AllowedInterpolationEnumValues = []Interpolation{
	"STEP",
	"LINEAR",
	"CATMULLROMSPLINE",
	"CUBICSPLINE",
}

func (v *Interpolation) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Interpolation(value)
	for _, existing := range AllowedInterpolationEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Interpolation", value)
}

// NewInterpolationFromValue returns a pointer to a valid Interpolation
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewInterpolationFromValue(v string) (*Interpolation, error) {
	ev := Interpolation(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for Interpolation: valid values are %v", v, AllowedInterpolationEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Interpolation) IsValid() bool {
	for _, existing := range AllowedInterpolationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Interpolation value
func (v Interpolation) Ptr() *Interpolation {
	return &v
}

type NullableInterpolation struct {
	value *Interpolation
	isSet bool
}

func (v NullableInterpolation) Get() *Interpolation {
	return v.value
}

func (v *NullableInterpolation) Set(val *Interpolation) {
	v.value = val
	v.isSet = true
}

func (v NullableInterpolation) IsSet() bool {
	return v.isSet
}

func (v *NullableInterpolation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInterpolation(val *Interpolation) *NullableInterpolation {
	return &NullableInterpolation{value: val, isSet: true}
}

func (v NullableInterpolation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInterpolation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
