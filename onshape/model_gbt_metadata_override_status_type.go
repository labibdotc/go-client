/*
Onshape REST API

## Welcome to the Onshape REST API Explorer  To use this API explorer, sign in to your [Onshape](https://cad.onshape.com) account in another tab, then click the **Try it out** button below (it toggles to a **Cancel** button when selected).  See the **[API Explorer Guide](https://onshape-public.github.io/docs/api-intro/explorer/)** for help navigating this API Explorer, including **[authentication](https://onshape-public.github.io/docs/api-intro/explorer/#authentication)**.  **Tip:** To ensure the current session isn't used when trying other authentication techniques, make sure to [remove the Onshape cookie](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site) as per the instructions for your browser. Alternatively, you can use a private or incognito window.  ## See Also  * [Onshape API Guide](https://onshape-public.github.io/docs/): Our full suite of developer guides, to be used as an accompaniment to this API Explorer. * [Onshape Developer Portal](https://dev-portal.onshape.com/): The Onshape portal for managing your API keys, OAuth2 credentials, your Onshape applications, and your Onshape App Store entries. * [Authentication Guide](https://onshape-public.github.io/docs/auth/): Our guide to using API keys, request signatures, and OAuth2 in  your Onshape applications.

API version: 1.175.29320-74695940af99
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
	"fmt"
)

// GBTMetadataOverrideStatusType the model 'GBTMetadataOverrideStatusType'
type GBTMetadataOverrideStatusType string

// List of GBTMetadataOverrideStatusType
const (
	GBTMetadataOverrideStatusTypeUnknown                          GBTMetadataOverrideStatusType = "UNKNOWN"
	GBTMetadataOverrideStatusTypePropertyNotComputed              GBTMetadataOverrideStatusType = "PROPERTY_NOT_COMPUTED"
	GBTMetadataOverrideStatusTypeComputedWithNoOverrides          GBTMetadataOverrideStatusType = "COMPUTED_WITH_NO_OVERRIDES"
	GBTMetadataOverrideStatusTypeComputedWithOverrides            GBTMetadataOverrideStatusType = "COMPUTED_WITH_OVERRIDES"
	GBTMetadataOverrideStatusTypeComputedWithSubassemblyOverrides GBTMetadataOverrideStatusType = "COMPUTED_WITH_SUBASSEMBLY_OVERRIDES"
	GBTMetadataOverrideStatusTypeOverridden                       GBTMetadataOverrideStatusType = "OVERRIDDEN"
)

// All allowed values of GBTMetadataOverrideStatusType enum
var AllowedGBTMetadataOverrideStatusTypeEnumValues = []GBTMetadataOverrideStatusType{
	"UNKNOWN",
	"PROPERTY_NOT_COMPUTED",
	"COMPUTED_WITH_NO_OVERRIDES",
	"COMPUTED_WITH_OVERRIDES",
	"COMPUTED_WITH_SUBASSEMBLY_OVERRIDES",
	"OVERRIDDEN",
}

func (v *GBTMetadataOverrideStatusType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GBTMetadataOverrideStatusType(value)
	for _, existing := range AllowedGBTMetadataOverrideStatusTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GBTMetadataOverrideStatusType", value)
}

// NewGBTMetadataOverrideStatusTypeFromValue returns a pointer to a valid GBTMetadataOverrideStatusType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGBTMetadataOverrideStatusTypeFromValue(v string) (*GBTMetadataOverrideStatusType, error) {
	ev := GBTMetadataOverrideStatusType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GBTMetadataOverrideStatusType: valid values are %v", v, AllowedGBTMetadataOverrideStatusTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GBTMetadataOverrideStatusType) IsValid() bool {
	for _, existing := range AllowedGBTMetadataOverrideStatusTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GBTMetadataOverrideStatusType value
func (v GBTMetadataOverrideStatusType) Ptr() *GBTMetadataOverrideStatusType {
	return &v
}

type NullableGBTMetadataOverrideStatusType struct {
	value *GBTMetadataOverrideStatusType
	isSet bool
}

func (v NullableGBTMetadataOverrideStatusType) Get() *GBTMetadataOverrideStatusType {
	return v.value
}

func (v *NullableGBTMetadataOverrideStatusType) Set(val *GBTMetadataOverrideStatusType) {
	v.value = val
	v.isSet = true
}

func (v NullableGBTMetadataOverrideStatusType) IsSet() bool {
	return v.isSet
}

func (v *NullableGBTMetadataOverrideStatusType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGBTMetadataOverrideStatusType(val *GBTMetadataOverrideStatusType) *NullableGBTMetadataOverrideStatusType {
	return &NullableGBTMetadataOverrideStatusType{value: val, isSet: true}
}

func (v NullableGBTMetadataOverrideStatusType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGBTMetadataOverrideStatusType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
