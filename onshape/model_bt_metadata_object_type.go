/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.167.19303-3cbf47a47fe4
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
	"fmt"
)

// BTMetadataObjectType the model 'BTMetadataObjectType'
type BTMetadataObjectType string

// List of BTMetadataObjectType
const (
	BTMetadataObjectTypeGlobal         BTMetadataObjectType = "GLOBAL"
	BTMetadataObjectTypeDocument       BTMetadataObjectType = "DOCUMENT"
	BTMetadataObjectTypePart           BTMetadataObjectType = "PART"
	BTMetadataObjectTypeAssembly       BTMetadataObjectType = "ASSEMBLY"
	BTMetadataObjectTypeDrawing        BTMetadataObjectType = "DRAWING"
	BTMetadataObjectTypePartStudio     BTMetadataObjectType = "PART_STUDIO"
	BTMetadataObjectTypeBlobElement    BTMetadataObjectType = "BLOB_ELEMENT"
	BTMetadataObjectTypeAppElement     BTMetadataObjectType = "APP_ELEMENT"
	BTMetadataObjectTypeVersion        BTMetadataObjectType = "VERSION"
	BTMetadataObjectTypeWorkspace      BTMetadataObjectType = "WORKSPACE"
	BTMetadataObjectTypeProject        BTMetadataObjectType = "PROJECT"
	BTMetadataObjectTypeItem           BTMetadataObjectType = "ITEM"
	BTMetadataObjectTypeFeatureStudio  BTMetadataObjectType = "FEATURE_STUDIO"
	BTMetadataObjectTypeChangeRequest  BTMetadataObjectType = "CHANGE_REQUEST"
	BTMetadataObjectTypeTask           BTMetadataObjectType = "TASK"
	BTMetadataObjectTypeChangeOrder    BTMetadataObjectType = "CHANGE_ORDER"
	BTMetadataObjectTypeChangeTask     BTMetadataObjectType = "CHANGE_TASK"
	BTMetadataObjectTypeVariableStudio BTMetadataObjectType = "VARIABLE_STUDIO"
)

// All allowed values of BTMetadataObjectType enum
var AllowedBTMetadataObjectTypeEnumValues = []BTMetadataObjectType{
	"GLOBAL",
	"DOCUMENT",
	"PART",
	"ASSEMBLY",
	"DRAWING",
	"PART_STUDIO",
	"BLOB_ELEMENT",
	"APP_ELEMENT",
	"VERSION",
	"WORKSPACE",
	"PROJECT",
	"ITEM",
	"FEATURE_STUDIO",
	"CHANGE_REQUEST",
	"TASK",
	"CHANGE_ORDER",
	"CHANGE_TASK",
	"VARIABLE_STUDIO",
}

func (v *BTMetadataObjectType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BTMetadataObjectType(value)
	for _, existing := range AllowedBTMetadataObjectTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BTMetadataObjectType", value)
}

// NewBTMetadataObjectTypeFromValue returns a pointer to a valid BTMetadataObjectType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBTMetadataObjectTypeFromValue(v string) (*BTMetadataObjectType, error) {
	ev := BTMetadataObjectType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BTMetadataObjectType: valid values are %v", v, AllowedBTMetadataObjectTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BTMetadataObjectType) IsValid() bool {
	for _, existing := range AllowedBTMetadataObjectTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BTMetadataObjectType value
func (v BTMetadataObjectType) Ptr() *BTMetadataObjectType {
	return &v
}

type NullableBTMetadataObjectType struct {
	value *BTMetadataObjectType
	isSet bool
}

func (v NullableBTMetadataObjectType) Get() *BTMetadataObjectType {
	return v.value
}

func (v *NullableBTMetadataObjectType) Set(val *BTMetadataObjectType) {
	v.value = val
	v.isSet = true
}

func (v NullableBTMetadataObjectType) IsSet() bool {
	return v.isSet
}

func (v *NullableBTMetadataObjectType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTMetadataObjectType(val *BTMetadataObjectType) *NullableBTMetadataObjectType {
	return &NullableBTMetadataObjectType{value: val, isSet: true}
}

func (v NullableBTMetadataObjectType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTMetadataObjectType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}