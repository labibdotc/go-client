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

// BTMetadataValueType the model 'BTMetadataValueType'
type BTMetadataValueType string

// List of BTMetadataValueType
const (
	BTMetadataValueTypeString         BTMetadataValueType = "STRING"
	BTMetadataValueTypeBool           BTMetadataValueType = "BOOL"
	BTMetadataValueTypeInt            BTMetadataValueType = "INT"
	BTMetadataValueTypeDouble         BTMetadataValueType = "DOUBLE"
	BTMetadataValueTypeDate           BTMetadataValueType = "DATE"
	BTMetadataValueTypeEnum           BTMetadataValueType = "ENUM"
	BTMetadataValueTypeObject         BTMetadataValueType = "OBJECT"
	BTMetadataValueTypeBlob           BTMetadataValueType = "BLOB"
	BTMetadataValueTypeUser           BTMetadataValueType = "USER"
	BTMetadataValueTypeList           BTMetadataValueType = "LIST"
	BTMetadataValueTypeForeign        BTMetadataValueType = "FOREIGN"
	BTMetadataValueTypeCategory       BTMetadataValueType = "CATEGORY"
	BTMetadataValueTypeComputed       BTMetadataValueType = "COMPUTED"
	BTMetadataValueTypeValueWithUnits BTMetadataValueType = "VALUE_WITH_UNITS"
)

// All allowed values of BTMetadataValueType enum
var AllowedBTMetadataValueTypeEnumValues = []BTMetadataValueType{
	"STRING",
	"BOOL",
	"INT",
	"DOUBLE",
	"DATE",
	"ENUM",
	"OBJECT",
	"BLOB",
	"USER",
	"LIST",
	"FOREIGN",
	"CATEGORY",
	"COMPUTED",
	"VALUE_WITH_UNITS",
}

func (v *BTMetadataValueType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BTMetadataValueType(value)
	for _, existing := range AllowedBTMetadataValueTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BTMetadataValueType", value)
}

// NewBTMetadataValueTypeFromValue returns a pointer to a valid BTMetadataValueType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBTMetadataValueTypeFromValue(v string) (*BTMetadataValueType, error) {
	ev := BTMetadataValueType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BTMetadataValueType: valid values are %v", v, AllowedBTMetadataValueTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BTMetadataValueType) IsValid() bool {
	for _, existing := range AllowedBTMetadataValueTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BTMetadataValueType value
func (v BTMetadataValueType) Ptr() *BTMetadataValueType {
	return &v
}

type NullableBTMetadataValueType struct {
	value *BTMetadataValueType
	isSet bool
}

func (v NullableBTMetadataValueType) Get() *BTMetadataValueType {
	return v.value
}

func (v *NullableBTMetadataValueType) Set(val *BTMetadataValueType) {
	v.value = val
	v.isSet = true
}

func (v NullableBTMetadataValueType) IsSet() bool {
	return v.isSet
}

func (v *NullableBTMetadataValueType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTMetadataValueType(val *BTMetadataValueType) *NullableBTMetadataValueType {
	return &NullableBTMetadataValueType{value: val, isSet: true}
}

func (v NullableBTMetadataValueType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTMetadataValueType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
