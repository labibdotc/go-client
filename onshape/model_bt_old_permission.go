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

// BTOldPermission the model 'BTOldPermission'
type BTOldPermission string

// List of BTOldPermission
const (
	BTOldPermissionNoaccess        BTOldPermission = "NOACCESS"
	BTOldPermissionAnonymousAccess BTOldPermission = "ANONYMOUS_ACCESS"
	BTOldPermissionRead            BTOldPermission = "READ"
	BTOldPermissionReadCopyExport  BTOldPermission = "READ_COPY_EXPORT"
	BTOldPermissionComment         BTOldPermission = "COMMENT"
	BTOldPermissionWrite           BTOldPermission = "WRITE"
	BTOldPermissionReshare         BTOldPermission = "RESHARE"
	BTOldPermissionFull            BTOldPermission = "FULL"
	BTOldPermissionOwner           BTOldPermission = "OWNER"
)

// All allowed values of BTOldPermission enum
var AllowedBTOldPermissionEnumValues = []BTOldPermission{
	"NOACCESS",
	"ANONYMOUS_ACCESS",
	"READ",
	"READ_COPY_EXPORT",
	"COMMENT",
	"WRITE",
	"RESHARE",
	"FULL",
	"OWNER",
}

func (v *BTOldPermission) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BTOldPermission(value)
	for _, existing := range AllowedBTOldPermissionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BTOldPermission", value)
}

// NewBTOldPermissionFromValue returns a pointer to a valid BTOldPermission
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBTOldPermissionFromValue(v string) (*BTOldPermission, error) {
	ev := BTOldPermission(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BTOldPermission: valid values are %v", v, AllowedBTOldPermissionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BTOldPermission) IsValid() bool {
	for _, existing := range AllowedBTOldPermissionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BTOldPermission value
func (v BTOldPermission) Ptr() *BTOldPermission {
	return &v
}

type NullableBTOldPermission struct {
	value *BTOldPermission
	isSet bool
}

func (v NullableBTOldPermission) Get() *BTOldPermission {
	return v.value
}

func (v *NullableBTOldPermission) Set(val *BTOldPermission) {
	v.value = val
	v.isSet = true
}

func (v NullableBTOldPermission) IsSet() bool {
	return v.isSet
}

func (v *NullableBTOldPermission) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTOldPermission(val *BTOldPermission) *NullableBTOldPermission {
	return &NullableBTOldPermission{value: val, isSet: true}
}

func (v NullableBTOldPermission) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTOldPermission) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
