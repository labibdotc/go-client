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
)

// BTPSOIdentity2741 struct for BTPSOIdentity2741
type BTPSOIdentity2741 struct {
	BtType *string `json:"btType,omitempty"`
	TheId  *string `json:"theId,omitempty"`
}

// NewBTPSOIdentity2741 instantiates a new BTPSOIdentity2741 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTPSOIdentity2741() *BTPSOIdentity2741 {
	this := BTPSOIdentity2741{}
	return &this
}

// NewBTPSOIdentity2741WithDefaults instantiates a new BTPSOIdentity2741 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTPSOIdentity2741WithDefaults() *BTPSOIdentity2741 {
	this := BTPSOIdentity2741{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTPSOIdentity2741) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPSOIdentity2741) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTPSOIdentity2741) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTPSOIdentity2741) SetBtType(v string) {
	o.BtType = &v
}

// GetTheId returns the TheId field value if set, zero value otherwise.
func (o *BTPSOIdentity2741) GetTheId() string {
	if o == nil || o.TheId == nil {
		var ret string
		return ret
	}
	return *o.TheId
}

// GetTheIdOk returns a tuple with the TheId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPSOIdentity2741) GetTheIdOk() (*string, bool) {
	if o == nil || o.TheId == nil {
		return nil, false
	}
	return o.TheId, true
}

// HasTheId returns a boolean if a field has been set.
func (o *BTPSOIdentity2741) HasTheId() bool {
	if o != nil && o.TheId != nil {
		return true
	}

	return false
}

// SetTheId gets a reference to the given string and assigns it to the TheId field.
func (o *BTPSOIdentity2741) SetTheId(v string) {
	o.TheId = &v
}

func (o BTPSOIdentity2741) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.TheId != nil {
		toSerialize["theId"] = o.TheId
	}
	return json.Marshal(toSerialize)
}

type NullableBTPSOIdentity2741 struct {
	value *BTPSOIdentity2741
	isSet bool
}

func (v NullableBTPSOIdentity2741) Get() *BTPSOIdentity2741 {
	return v.value
}

func (v *NullableBTPSOIdentity2741) Set(val *BTPSOIdentity2741) {
	v.value = val
	v.isSet = true
}

func (v NullableBTPSOIdentity2741) IsSet() bool {
	return v.isSet
}

func (v *NullableBTPSOIdentity2741) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTPSOIdentity2741(val *BTPSOIdentity2741) *NullableBTPSOIdentity2741 {
	return &NullableBTPSOIdentity2741{value: val, isSet: true}
}

func (v NullableBTPSOIdentity2741) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTPSOIdentity2741) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
