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

// GBTNodeChange the model 'GBTNodeChange'
type GBTNodeChange string

// List of GBTNodeChange
const (
	GBTNodeChangeNone             GBTNodeChange = "NONE"
	GBTNodeChangeMoved            GBTNodeChange = "MOVED"
	GBTNodeChangeModified         GBTNodeChange = "MODIFIED"
	GBTNodeChangeMovedAndModified GBTNodeChange = "MOVED_AND_MODIFIED"
	GBTNodeChangeAdded            GBTNodeChange = "ADDED"
	GBTNodeChangeDeleted          GBTNodeChange = "DELETED"
	GBTNodeChangeUnknown          GBTNodeChange = "UNKNOWN"
)

// All allowed values of GBTNodeChange enum
var AllowedGBTNodeChangeEnumValues = []GBTNodeChange{
	"NONE",
	"MOVED",
	"MODIFIED",
	"MOVED_AND_MODIFIED",
	"ADDED",
	"DELETED",
	"UNKNOWN",
}

func (v *GBTNodeChange) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GBTNodeChange(value)
	for _, existing := range AllowedGBTNodeChangeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GBTNodeChange", value)
}

// NewGBTNodeChangeFromValue returns a pointer to a valid GBTNodeChange
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGBTNodeChangeFromValue(v string) (*GBTNodeChange, error) {
	ev := GBTNodeChange(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GBTNodeChange: valid values are %v", v, AllowedGBTNodeChangeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GBTNodeChange) IsValid() bool {
	for _, existing := range AllowedGBTNodeChangeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GBTNodeChange value
func (v GBTNodeChange) Ptr() *GBTNodeChange {
	return &v
}

type NullableGBTNodeChange struct {
	value *GBTNodeChange
	isSet bool
}

func (v NullableGBTNodeChange) Get() *GBTNodeChange {
	return v.value
}

func (v *NullableGBTNodeChange) Set(val *GBTNodeChange) {
	v.value = val
	v.isSet = true
}

func (v NullableGBTNodeChange) IsSet() bool {
	return v.isSet
}

func (v *NullableGBTNodeChange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGBTNodeChange(val *GBTNodeChange) *NullableGBTNodeChange {
	return &NullableGBTNodeChange{value: val, isSet: true}
}

func (v NullableGBTNodeChange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGBTNodeChange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
