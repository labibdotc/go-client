/*
Onshape REST API

## Welcome to the Onshape REST API Explorer  To use this API explorer, sign in to your [Onshape](https://cad.onshape.com) account in another tab, then click the **Try it out** button below (it toggles to a **Cancel** button when selected).  See the **[API Explorer Guide](https://onshape-public.github.io/docs/api-intro/explorer/)** for help navigating this API Explorer, including **[authentication](https://onshape-public.github.io/docs/api-intro/explorer/#authentication)**.  **Tip:** To ensure the current session isn't used when trying other authentication techniques, make sure to [remove the Onshape cookie](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site) as per the instructions for your browser. Alternatively, you can use a private or incognito window.  ## See Also  * [Onshape API Guide](https://onshape-public.github.io/docs/): Our full suite of developer guides, to be used as an accompaniment to this API Explorer. * [Onshape Developer Portal](https://dev-portal.onshape.com/): The Onshape portal for managing your API keys, OAuth2 credentials, your Onshape applications, and your Onshape App Store entries. * [Authentication Guide](https://onshape-public.github.io/docs/auth/): Our guide to using API keys, request signatures, and OAuth2 in  your Onshape applications.

API version: 1.174.28023-41481dfcfdcb
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
	"fmt"
)

// StyleEnum the model 'StyleEnum'
type StyleEnum string

// List of StyleEnum
const (
	StyleEnumForm           StyleEnum = "FORM"
	StyleEnumSpaceDelimited StyleEnum = "SPACE_DELIMITED"
	StyleEnumPipeDelimited  StyleEnum = "PIPE_DELIMITED"
	StyleEnumDeepObject     StyleEnum = "DEEP_OBJECT"
)

// All allowed values of StyleEnum enum
var AllowedStyleEnumEnumValues = []StyleEnum{
	"FORM",
	"SPACE_DELIMITED",
	"PIPE_DELIMITED",
	"DEEP_OBJECT",
}

func (v *StyleEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := StyleEnum(value)
	for _, existing := range AllowedStyleEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid StyleEnum", value)
}

// NewStyleEnumFromValue returns a pointer to a valid StyleEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewStyleEnumFromValue(v string) (*StyleEnum, error) {
	ev := StyleEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for StyleEnum: valid values are %v", v, AllowedStyleEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v StyleEnum) IsValid() bool {
	for _, existing := range AllowedStyleEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to StyleEnum value
func (v StyleEnum) Ptr() *StyleEnum {
	return &v
}

type NullableStyleEnum struct {
	value *StyleEnum
	isSet bool
}

func (v NullableStyleEnum) Get() *StyleEnum {
	return v.value
}

func (v *NullableStyleEnum) Set(val *StyleEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableStyleEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableStyleEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStyleEnum(val *StyleEnum) *NullableStyleEnum {
	return &NullableStyleEnum{value: val, isSet: true}
}

func (v NullableStyleEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStyleEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
