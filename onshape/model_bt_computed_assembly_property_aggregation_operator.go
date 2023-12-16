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

// BTComputedAssemblyPropertyAggregationOperator the model 'BTComputedAssemblyPropertyAggregationOperator'
type BTComputedAssemblyPropertyAggregationOperator string

// List of BTComputedAssemblyPropertyAggregationOperator
const (
	BTComputedAssemblyPropertyAggregationOperatorSum             BTComputedAssemblyPropertyAggregationOperator = "SUM"
	BTComputedAssemblyPropertyAggregationOperatorMinimum         BTComputedAssemblyPropertyAggregationOperator = "MINIMUM"
	BTComputedAssemblyPropertyAggregationOperatorMaximum         BTComputedAssemblyPropertyAggregationOperator = "MAXIMUM"
	BTComputedAssemblyPropertyAggregationOperatorAverage         BTComputedAssemblyPropertyAggregationOperator = "AVERAGE"
	BTComputedAssemblyPropertyAggregationOperatorWeightedSum     BTComputedAssemblyPropertyAggregationOperator = "WEIGHTED_SUM"
	BTComputedAssemblyPropertyAggregationOperatorWeightedAverage BTComputedAssemblyPropertyAggregationOperator = "WEIGHTED_AVERAGE"
	BTComputedAssemblyPropertyAggregationOperatorAny             BTComputedAssemblyPropertyAggregationOperator = "ANY"
	BTComputedAssemblyPropertyAggregationOperatorAll             BTComputedAssemblyPropertyAggregationOperator = "ALL"
	BTComputedAssemblyPropertyAggregationOperatorNotAny          BTComputedAssemblyPropertyAggregationOperator = "NOT_ANY"
	BTComputedAssemblyPropertyAggregationOperatorNotAll          BTComputedAssemblyPropertyAggregationOperator = "NOT_ALL"
)

// All allowed values of BTComputedAssemblyPropertyAggregationOperator enum
var AllowedBTComputedAssemblyPropertyAggregationOperatorEnumValues = []BTComputedAssemblyPropertyAggregationOperator{
	"SUM",
	"MINIMUM",
	"MAXIMUM",
	"AVERAGE",
	"WEIGHTED_SUM",
	"WEIGHTED_AVERAGE",
	"ANY",
	"ALL",
	"NOT_ANY",
	"NOT_ALL",
}

func (v *BTComputedAssemblyPropertyAggregationOperator) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BTComputedAssemblyPropertyAggregationOperator(value)
	for _, existing := range AllowedBTComputedAssemblyPropertyAggregationOperatorEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BTComputedAssemblyPropertyAggregationOperator", value)
}

// NewBTComputedAssemblyPropertyAggregationOperatorFromValue returns a pointer to a valid BTComputedAssemblyPropertyAggregationOperator
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBTComputedAssemblyPropertyAggregationOperatorFromValue(v string) (*BTComputedAssemblyPropertyAggregationOperator, error) {
	ev := BTComputedAssemblyPropertyAggregationOperator(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BTComputedAssemblyPropertyAggregationOperator: valid values are %v", v, AllowedBTComputedAssemblyPropertyAggregationOperatorEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BTComputedAssemblyPropertyAggregationOperator) IsValid() bool {
	for _, existing := range AllowedBTComputedAssemblyPropertyAggregationOperatorEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BTComputedAssemblyPropertyAggregationOperator value
func (v BTComputedAssemblyPropertyAggregationOperator) Ptr() *BTComputedAssemblyPropertyAggregationOperator {
	return &v
}

type NullableBTComputedAssemblyPropertyAggregationOperator struct {
	value *BTComputedAssemblyPropertyAggregationOperator
	isSet bool
}

func (v NullableBTComputedAssemblyPropertyAggregationOperator) Get() *BTComputedAssemblyPropertyAggregationOperator {
	return v.value
}

func (v *NullableBTComputedAssemblyPropertyAggregationOperator) Set(val *BTComputedAssemblyPropertyAggregationOperator) {
	v.value = val
	v.isSet = true
}

func (v NullableBTComputedAssemblyPropertyAggregationOperator) IsSet() bool {
	return v.isSet
}

func (v *NullableBTComputedAssemblyPropertyAggregationOperator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTComputedAssemblyPropertyAggregationOperator(val *BTComputedAssemblyPropertyAggregationOperator) *NullableBTComputedAssemblyPropertyAggregationOperator {
	return &NullableBTComputedAssemblyPropertyAggregationOperator{value: val, isSet: true}
}

func (v NullableBTComputedAssemblyPropertyAggregationOperator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTComputedAssemblyPropertyAggregationOperator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
