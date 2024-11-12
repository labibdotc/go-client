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

// BTJPathKey3221 Identifies a subtree of the json structure by field name.
type BTJPathKey3221 struct {
	BTJPathElement2297
	BtType *string `json:"btType,omitempty"`
	Key    *string `json:"key,omitempty"`
}

// NewBTJPathKey3221 instantiates a new BTJPathKey3221 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTJPathKey3221() *BTJPathKey3221 {
	this := BTJPathKey3221{}
	return &this
}

// NewBTJPathKey3221WithDefaults instantiates a new BTJPathKey3221 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTJPathKey3221WithDefaults() *BTJPathKey3221 {
	this := BTJPathKey3221{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTJPathKey3221) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTJPathKey3221) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTJPathKey3221) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTJPathKey3221) SetBtType(v string) {
	o.BtType = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *BTJPathKey3221) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTJPathKey3221) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *BTJPathKey3221) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *BTJPathKey3221) SetKey(v string) {
	o.Key = &v
}

func (o BTJPathKey3221) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedBTJPathElement2297, errBTJPathElement2297 := json.Marshal(o.BTJPathElement2297)
	if errBTJPathElement2297 != nil {
		return []byte{}, errBTJPathElement2297
	}
	errBTJPathElement2297 = json.Unmarshal([]byte(serializedBTJPathElement2297), &toSerialize)
	if errBTJPathElement2297 != nil {
		return []byte{}, errBTJPathElement2297
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	return json.Marshal(toSerialize)
}

type NullableBTJPathKey3221 struct {
	value *BTJPathKey3221
	isSet bool
}

func (v NullableBTJPathKey3221) Get() *BTJPathKey3221 {
	return v.value
}

func (v *NullableBTJPathKey3221) Set(val *BTJPathKey3221) {
	v.value = val
	v.isSet = true
}

func (v NullableBTJPathKey3221) IsSet() bool {
	return v.isSet
}

func (v *NullableBTJPathKey3221) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTJPathKey3221(val *BTJPathKey3221) *NullableBTJPathKey3221 {
	return &NullableBTJPathKey3221{value: val, isSet: true}
}

func (v NullableBTJPathKey3221) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTJPathKey3221) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
