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
)

// BTAssemblyTransformDefinitionParams struct for BTAssemblyTransformDefinitionParams
type BTAssemblyTransformDefinitionParams struct {
	IsRelative  *bool            `json:"isRelative,omitempty"`
	Occurrences []BTOccurrence74 `json:"occurrences,omitempty"`
	Transform   []float64        `json:"transform,omitempty"`
}

// NewBTAssemblyTransformDefinitionParams instantiates a new BTAssemblyTransformDefinitionParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTAssemblyTransformDefinitionParams() *BTAssemblyTransformDefinitionParams {
	this := BTAssemblyTransformDefinitionParams{}
	return &this
}

// NewBTAssemblyTransformDefinitionParamsWithDefaults instantiates a new BTAssemblyTransformDefinitionParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTAssemblyTransformDefinitionParamsWithDefaults() *BTAssemblyTransformDefinitionParams {
	this := BTAssemblyTransformDefinitionParams{}
	return &this
}

// GetIsRelative returns the IsRelative field value if set, zero value otherwise.
func (o *BTAssemblyTransformDefinitionParams) GetIsRelative() bool {
	if o == nil || o.IsRelative == nil {
		var ret bool
		return ret
	}
	return *o.IsRelative
}

// GetIsRelativeOk returns a tuple with the IsRelative field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyTransformDefinitionParams) GetIsRelativeOk() (*bool, bool) {
	if o == nil || o.IsRelative == nil {
		return nil, false
	}
	return o.IsRelative, true
}

// HasIsRelative returns a boolean if a field has been set.
func (o *BTAssemblyTransformDefinitionParams) HasIsRelative() bool {
	if o != nil && o.IsRelative != nil {
		return true
	}

	return false
}

// SetIsRelative gets a reference to the given bool and assigns it to the IsRelative field.
func (o *BTAssemblyTransformDefinitionParams) SetIsRelative(v bool) {
	o.IsRelative = &v
}

// GetOccurrences returns the Occurrences field value if set, zero value otherwise.
func (o *BTAssemblyTransformDefinitionParams) GetOccurrences() []BTOccurrence74 {
	if o == nil || o.Occurrences == nil {
		var ret []BTOccurrence74
		return ret
	}
	return o.Occurrences
}

// GetOccurrencesOk returns a tuple with the Occurrences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyTransformDefinitionParams) GetOccurrencesOk() ([]BTOccurrence74, bool) {
	if o == nil || o.Occurrences == nil {
		return nil, false
	}
	return o.Occurrences, true
}

// HasOccurrences returns a boolean if a field has been set.
func (o *BTAssemblyTransformDefinitionParams) HasOccurrences() bool {
	if o != nil && o.Occurrences != nil {
		return true
	}

	return false
}

// SetOccurrences gets a reference to the given []BTOccurrence74 and assigns it to the Occurrences field.
func (o *BTAssemblyTransformDefinitionParams) SetOccurrences(v []BTOccurrence74) {
	o.Occurrences = v
}

// GetTransform returns the Transform field value if set, zero value otherwise.
func (o *BTAssemblyTransformDefinitionParams) GetTransform() []float64 {
	if o == nil || o.Transform == nil {
		var ret []float64
		return ret
	}
	return o.Transform
}

// GetTransformOk returns a tuple with the Transform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyTransformDefinitionParams) GetTransformOk() ([]float64, bool) {
	if o == nil || o.Transform == nil {
		return nil, false
	}
	return o.Transform, true
}

// HasTransform returns a boolean if a field has been set.
func (o *BTAssemblyTransformDefinitionParams) HasTransform() bool {
	if o != nil && o.Transform != nil {
		return true
	}

	return false
}

// SetTransform gets a reference to the given []float64 and assigns it to the Transform field.
func (o *BTAssemblyTransformDefinitionParams) SetTransform(v []float64) {
	o.Transform = v
}

func (o BTAssemblyTransformDefinitionParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IsRelative != nil {
		toSerialize["isRelative"] = o.IsRelative
	}
	if o.Occurrences != nil {
		toSerialize["occurrences"] = o.Occurrences
	}
	if o.Transform != nil {
		toSerialize["transform"] = o.Transform
	}
	return json.Marshal(toSerialize)
}

type NullableBTAssemblyTransformDefinitionParams struct {
	value *BTAssemblyTransformDefinitionParams
	isSet bool
}

func (v NullableBTAssemblyTransformDefinitionParams) Get() *BTAssemblyTransformDefinitionParams {
	return v.value
}

func (v *NullableBTAssemblyTransformDefinitionParams) Set(val *BTAssemblyTransformDefinitionParams) {
	v.value = val
	v.isSet = true
}

func (v NullableBTAssemblyTransformDefinitionParams) IsSet() bool {
	return v.isSet
}

func (v *NullableBTAssemblyTransformDefinitionParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTAssemblyTransformDefinitionParams(val *BTAssemblyTransformDefinitionParams) *NullableBTAssemblyTransformDefinitionParams {
	return &NullableBTAssemblyTransformDefinitionParams{value: val, isSet: true}
}

func (v NullableBTAssemblyTransformDefinitionParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTAssemblyTransformDefinitionParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
