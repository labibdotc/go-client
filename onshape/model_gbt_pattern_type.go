/*
Onshape REST API

## Welcome to the Onshape REST API Explorer  To use this API explorer, sign in to your [Onshape](https://cad.onshape.com) account in another tab, then click the **Try it out** button below (it toggles to a **Cancel** button when selected).  See the **[API Explorer Guide](https://onshape-public.github.io/docs/api-intro/explorer/)** for help navigating this API Explorer, including **[authentication](https://onshape-public.github.io/docs/api-intro/explorer/#authentication)**.  **Tip:** To ensure the current session isn't used when trying other authentication techniques, make sure to [remove the Onshape cookie](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site) as per the instructions for your browser. Alternatively, you can use a private or incognito window.  ## See Also  * [Onshape API Guide](https://onshape-public.github.io/docs/): Our full suite of developer guides, to be used as an accompaniment to this API Explorer. * [Onshape Developer Portal](https://dev-portal.onshape.com/): The Onshape portal for managing your API keys, OAuth2 credentials, your Onshape applications, and your Onshape App Store entries. * [Authentication Guide](https://onshape-public.github.io/docs/auth/): Our guide to using API keys, request signatures, and OAuth2 in  your Onshape applications.

API version: 1.174.27783-ab3907bf6199
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
	"fmt"
)

// GBTPatternType the model 'GBTPatternType'
type GBTPatternType string

// List of GBTPatternType
const (
	GBTPatternTypeLinear   GBTPatternType = "LINEAR"
	GBTPatternTypeCircular GBTPatternType = "CIRCULAR"
	GBTPatternTypeUnknown  GBTPatternType = "UNKNOWN"
)

// All allowed values of GBTPatternType enum
var AllowedGBTPatternTypeEnumValues = []GBTPatternType{
	"LINEAR",
	"CIRCULAR",
	"UNKNOWN",
}

func (v *GBTPatternType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GBTPatternType(value)
	for _, existing := range AllowedGBTPatternTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GBTPatternType", value)
}

// NewGBTPatternTypeFromValue returns a pointer to a valid GBTPatternType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGBTPatternTypeFromValue(v string) (*GBTPatternType, error) {
	ev := GBTPatternType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GBTPatternType: valid values are %v", v, AllowedGBTPatternTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GBTPatternType) IsValid() bool {
	for _, existing := range AllowedGBTPatternTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GBTPatternType value
func (v GBTPatternType) Ptr() *GBTPatternType {
	return &v
}

type NullableGBTPatternType struct {
	value *GBTPatternType
	isSet bool
}

func (v NullableGBTPatternType) Get() *GBTPatternType {
	return v.value
}

func (v *NullableGBTPatternType) Set(val *GBTPatternType) {
	v.value = val
	v.isSet = true
}

func (v NullableGBTPatternType) IsSet() bool {
	return v.isSet
}

func (v *NullableGBTPatternType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGBTPatternType(val *GBTPatternType) *NullableGBTPatternType {
	return &NullableGBTPatternType{value: val, isSet: true}
}

func (v NullableGBTPatternType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGBTPatternType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
