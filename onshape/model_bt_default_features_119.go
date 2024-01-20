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
)

// BTDefaultFeatures119 struct for BTDefaultFeatures119
type BTDefaultFeatures119 struct {
	BtType   *string         `json:"btType,omitempty"`
	Children []BTMFeature134 `json:"children,omitempty"`
}

// NewBTDefaultFeatures119 instantiates a new BTDefaultFeatures119 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTDefaultFeatures119() *BTDefaultFeatures119 {
	this := BTDefaultFeatures119{}
	return &this
}

// NewBTDefaultFeatures119WithDefaults instantiates a new BTDefaultFeatures119 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTDefaultFeatures119WithDefaults() *BTDefaultFeatures119 {
	this := BTDefaultFeatures119{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTDefaultFeatures119) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDefaultFeatures119) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTDefaultFeatures119) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTDefaultFeatures119) SetBtType(v string) {
	o.BtType = &v
}

// GetChildren returns the Children field value if set, zero value otherwise.
func (o *BTDefaultFeatures119) GetChildren() []BTMFeature134 {
	if o == nil || o.Children == nil {
		var ret []BTMFeature134
		return ret
	}
	return o.Children
}

// GetChildrenOk returns a tuple with the Children field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDefaultFeatures119) GetChildrenOk() ([]BTMFeature134, bool) {
	if o == nil || o.Children == nil {
		return nil, false
	}
	return o.Children, true
}

// HasChildren returns a boolean if a field has been set.
func (o *BTDefaultFeatures119) HasChildren() bool {
	if o != nil && o.Children != nil {
		return true
	}

	return false
}

// SetChildren gets a reference to the given []BTMFeature134 and assigns it to the Children field.
func (o *BTDefaultFeatures119) SetChildren(v []BTMFeature134) {
	o.Children = v
}

func (o BTDefaultFeatures119) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.Children != nil {
		toSerialize["children"] = o.Children
	}
	return json.Marshal(toSerialize)
}

type NullableBTDefaultFeatures119 struct {
	value *BTDefaultFeatures119
	isSet bool
}

func (v NullableBTDefaultFeatures119) Get() *BTDefaultFeatures119 {
	return v.value
}

func (v *NullableBTDefaultFeatures119) Set(val *BTDefaultFeatures119) {
	v.value = val
	v.isSet = true
}

func (v NullableBTDefaultFeatures119) IsSet() bool {
	return v.isSet
}

func (v *NullableBTDefaultFeatures119) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTDefaultFeatures119(val *BTDefaultFeatures119) *NullableBTDefaultFeatures119 {
	return &NullableBTDefaultFeatures119{value: val, isSet: true}
}

func (v NullableBTDefaultFeatures119) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTDefaultFeatures119) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
