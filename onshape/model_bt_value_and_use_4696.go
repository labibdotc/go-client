/*
Onshape REST API

## Welcome to the Onshape REST API Explorer  To use this API explorer, sign in to your [Onshape](https://cad.onshape.com) account in another tab, then click the **Try it out** button below (it toggles to a **Cancel** button when selected).  See the **[API Explorer Guide](https://onshape-public.github.io/docs/api-intro/explorer/)** for help navigating this API Explorer, including **[authentication](https://onshape-public.github.io/docs/api-intro/explorer/#authentication)**.  **Tip:** To ensure the current session isn't used when trying other authentication techniques, make sure to [remove the Onshape cookie](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site) as per the instructions for your browser. Alternatively, you can use a private or incognito window.  ## See Also  * [Onshape API Guide](https://onshape-public.github.io/docs/): Our full suite of developer guides, to be used as an accompaniment to this API Explorer. * [Onshape Developer Portal](https://dev-portal.onshape.com/): The Onshape portal for managing your API keys, OAuth2 credentials, your Onshape applications, and your Onshape App Store entries. * [Authentication Guide](https://onshape-public.github.io/docs/auth/): Our guide to using API keys, request signatures, and OAuth2 in  your Onshape applications.

Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTValueAndUse4696 struct for BTValueAndUse4696
type BTValueAndUse4696 struct {
	// Type of JSON object.
	BtType *string        `json:"btType,omitempty"`
	Use    *GBTValueUse   `json:"use,omitempty"`
	Value  *BTFSValue1888 `json:"value,omitempty"`
}

// NewBTValueAndUse4696 instantiates a new BTValueAndUse4696 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTValueAndUse4696() *BTValueAndUse4696 {
	this := BTValueAndUse4696{}
	return &this
}

// NewBTValueAndUse4696WithDefaults instantiates a new BTValueAndUse4696 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTValueAndUse4696WithDefaults() *BTValueAndUse4696 {
	this := BTValueAndUse4696{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTValueAndUse4696) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTValueAndUse4696) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTValueAndUse4696) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTValueAndUse4696) SetBtType(v string) {
	o.BtType = &v
}

// GetUse returns the Use field value if set, zero value otherwise.
func (o *BTValueAndUse4696) GetUse() GBTValueUse {
	if o == nil || o.Use == nil {
		var ret GBTValueUse
		return ret
	}
	return *o.Use
}

// GetUseOk returns a tuple with the Use field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTValueAndUse4696) GetUseOk() (*GBTValueUse, bool) {
	if o == nil || o.Use == nil {
		return nil, false
	}
	return o.Use, true
}

// HasUse returns a boolean if a field has been set.
func (o *BTValueAndUse4696) HasUse() bool {
	if o != nil && o.Use != nil {
		return true
	}

	return false
}

// SetUse gets a reference to the given GBTValueUse and assigns it to the Use field.
func (o *BTValueAndUse4696) SetUse(v GBTValueUse) {
	o.Use = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *BTValueAndUse4696) GetValue() BTFSValue1888 {
	if o == nil || o.Value == nil {
		var ret BTFSValue1888
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTValueAndUse4696) GetValueOk() (*BTFSValue1888, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *BTValueAndUse4696) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given BTFSValue1888 and assigns it to the Value field.
func (o *BTValueAndUse4696) SetValue(v BTFSValue1888) {
	o.Value = &v
}

func (o BTValueAndUse4696) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.Use != nil {
		toSerialize["use"] = o.Use
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableBTValueAndUse4696 struct {
	value *BTValueAndUse4696
	isSet bool
}

func (v NullableBTValueAndUse4696) Get() *BTValueAndUse4696 {
	return v.value
}

func (v *NullableBTValueAndUse4696) Set(val *BTValueAndUse4696) {
	v.value = val
	v.isSet = true
}

func (v NullableBTValueAndUse4696) IsSet() bool {
	return v.isSet
}

func (v *NullableBTValueAndUse4696) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTValueAndUse4696(val *BTValueAndUse4696) *NullableBTValueAndUse4696 {
	return &NullableBTValueAndUse4696{value: val, isSet: true}
}

func (v NullableBTValueAndUse4696) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTValueAndUse4696) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
