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

// BTUserState the model 'BTUserState'
type BTUserState string

// List of BTUserState
const (
	BTUserStateDeleted           BTUserState = "DELETED"
	BTUserStateActive            BTUserState = "ACTIVE"
	BTUserStateInactive          BTUserState = "INACTIVE"
	BTUserStateRequested         BTUserState = "REQUESTED"
	BTUserStateApproved          BTUserState = "APPROVED"
	BTUserStateRequestExpired    BTUserState = "REQUEST_EXPIRED"
	BTUserStateAll               BTUserState = "ALL"
	BTUserStateMarkedForDeletion BTUserState = "MARKED_FOR_DELETION"
)

// All allowed values of BTUserState enum
var AllowedBTUserStateEnumValues = []BTUserState{
	"DELETED",
	"ACTIVE",
	"INACTIVE",
	"REQUESTED",
	"APPROVED",
	"REQUEST_EXPIRED",
	"ALL",
	"MARKED_FOR_DELETION",
}

func (v *BTUserState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BTUserState(value)
	for _, existing := range AllowedBTUserStateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BTUserState", value)
}

// NewBTUserStateFromValue returns a pointer to a valid BTUserState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBTUserStateFromValue(v string) (*BTUserState, error) {
	ev := BTUserState(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BTUserState: valid values are %v", v, AllowedBTUserStateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BTUserState) IsValid() bool {
	for _, existing := range AllowedBTUserStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BTUserState value
func (v BTUserState) Ptr() *BTUserState {
	return &v
}

type NullableBTUserState struct {
	value *BTUserState
	isSet bool
}

func (v NullableBTUserState) Get() *BTUserState {
	return v.value
}

func (v *NullableBTUserState) Set(val *BTUserState) {
	v.value = val
	v.isSet = true
}

func (v NullableBTUserState) IsSet() bool {
	return v.isSet
}

func (v *NullableBTUserState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTUserState(val *BTUserState) *NullableBTUserState {
	return &NullableBTUserState{value: val, isSet: true}
}

func (v NullableBTUserState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTUserState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
