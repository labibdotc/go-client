/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.171.24804-920f3dc76f2b
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
	"fmt"
)

// GBTMeshState the model 'GBTMeshState'
type GBTMeshState string

// List of GBTMeshState
const (
	GBTMeshStateNoMesh  GBTMeshState = "NO_MESH"
	GBTMeshStateMixed   GBTMeshState = "MIXED"
	GBTMeshStateAllMesh GBTMeshState = "ALL_MESH"
	GBTMeshStateUnknown GBTMeshState = "UNKNOWN"
)

// All allowed values of GBTMeshState enum
var AllowedGBTMeshStateEnumValues = []GBTMeshState{
	"NO_MESH",
	"MIXED",
	"ALL_MESH",
	"UNKNOWN",
}

func (v *GBTMeshState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GBTMeshState(value)
	for _, existing := range AllowedGBTMeshStateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GBTMeshState", value)
}

// NewGBTMeshStateFromValue returns a pointer to a valid GBTMeshState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGBTMeshStateFromValue(v string) (*GBTMeshState, error) {
	ev := GBTMeshState(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GBTMeshState: valid values are %v", v, AllowedGBTMeshStateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GBTMeshState) IsValid() bool {
	for _, existing := range AllowedGBTMeshStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GBTMeshState value
func (v GBTMeshState) Ptr() *GBTMeshState {
	return &v
}

type NullableGBTMeshState struct {
	value *GBTMeshState
	isSet bool
}

func (v NullableGBTMeshState) Get() *GBTMeshState {
	return v.value
}

func (v *NullableGBTMeshState) Set(val *GBTMeshState) {
	v.value = val
	v.isSet = true
}

func (v NullableGBTMeshState) IsSet() bool {
	return v.isSet
}

func (v *NullableGBTMeshState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGBTMeshState(val *GBTMeshState) *NullableGBTMeshState {
	return &NullableGBTMeshState{value: val, isSet: true}
}

func (v NullableGBTMeshState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGBTMeshState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
