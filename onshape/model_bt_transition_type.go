/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.169.22266-e2d421ffb3ea
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
	"fmt"
)

// BTTransitionType Transition types(SUBMIT, APPROVE, REJECT)
type BTTransitionType string

// List of BTTransitionType
const (
	BTTransitionTypeDefault       BTTransitionType = "DEFAULT"
	BTTransitionTypeSubmit        BTTransitionType = "SUBMIT"
	BTTransitionTypeApprove       BTTransitionType = "APPROVE"
	BTTransitionTypeApproveDirect BTTransitionType = "APPROVE_DIRECT"
	BTTransitionTypeReject        BTTransitionType = "REJECT"
	BTTransitionTypeDelete        BTTransitionType = "DELETE"
	BTTransitionTypeComment       BTTransitionType = "COMMENT"
	BTTransitionTypeReassignTask  BTTransitionType = "REASSIGN_TASK"
)

// All allowed values of BTTransitionType enum
var AllowedBTTransitionTypeEnumValues = []BTTransitionType{
	"DEFAULT",
	"SUBMIT",
	"APPROVE",
	"APPROVE_DIRECT",
	"REJECT",
	"DELETE",
	"COMMENT",
	"REASSIGN_TASK",
}

func (v *BTTransitionType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BTTransitionType(value)
	for _, existing := range AllowedBTTransitionTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BTTransitionType", value)
}

// NewBTTransitionTypeFromValue returns a pointer to a valid BTTransitionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBTTransitionTypeFromValue(v string) (*BTTransitionType, error) {
	ev := BTTransitionType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BTTransitionType: valid values are %v", v, AllowedBTTransitionTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BTTransitionType) IsValid() bool {
	for _, existing := range AllowedBTTransitionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BTTransitionType value
func (v BTTransitionType) Ptr() *BTTransitionType {
	return &v
}

type NullableBTTransitionType struct {
	value *BTTransitionType
	isSet bool
}

func (v NullableBTTransitionType) Get() *BTTransitionType {
	return v.value
}

func (v *NullableBTTransitionType) Set(val *BTTransitionType) {
	v.value = val
	v.isSet = true
}

func (v NullableBTTransitionType) IsSet() bool {
	return v.isSet
}

func (v *NullableBTTransitionType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTTransitionType(val *BTTransitionType) *NullableBTTransitionType {
	return &NullableBTTransitionType{value: val, isSet: true}
}

func (v NullableBTTransitionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTTransitionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
