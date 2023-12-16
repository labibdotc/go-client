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

// BTParameterExpressionErrorInfo2037 struct for BTParameterExpressionErrorInfo2037
type BTParameterExpressionErrorInfo2037 struct {
	BtType                 *string             `json:"btType,omitempty"`
	ErrorMessageIdentifier *GBTErrorStringEnum `json:"errorMessageIdentifier,omitempty"`
	MessageArguments       []BTValueAndUse4696 `json:"messageArguments,omitempty"`
}

// NewBTParameterExpressionErrorInfo2037 instantiates a new BTParameterExpressionErrorInfo2037 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTParameterExpressionErrorInfo2037() *BTParameterExpressionErrorInfo2037 {
	this := BTParameterExpressionErrorInfo2037{}
	return &this
}

// NewBTParameterExpressionErrorInfo2037WithDefaults instantiates a new BTParameterExpressionErrorInfo2037 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTParameterExpressionErrorInfo2037WithDefaults() *BTParameterExpressionErrorInfo2037 {
	this := BTParameterExpressionErrorInfo2037{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTParameterExpressionErrorInfo2037) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterExpressionErrorInfo2037) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTParameterExpressionErrorInfo2037) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTParameterExpressionErrorInfo2037) SetBtType(v string) {
	o.BtType = &v
}

// GetErrorMessageIdentifier returns the ErrorMessageIdentifier field value if set, zero value otherwise.
func (o *BTParameterExpressionErrorInfo2037) GetErrorMessageIdentifier() GBTErrorStringEnum {
	if o == nil || o.ErrorMessageIdentifier == nil {
		var ret GBTErrorStringEnum
		return ret
	}
	return *o.ErrorMessageIdentifier
}

// GetErrorMessageIdentifierOk returns a tuple with the ErrorMessageIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterExpressionErrorInfo2037) GetErrorMessageIdentifierOk() (*GBTErrorStringEnum, bool) {
	if o == nil || o.ErrorMessageIdentifier == nil {
		return nil, false
	}
	return o.ErrorMessageIdentifier, true
}

// HasErrorMessageIdentifier returns a boolean if a field has been set.
func (o *BTParameterExpressionErrorInfo2037) HasErrorMessageIdentifier() bool {
	if o != nil && o.ErrorMessageIdentifier != nil {
		return true
	}

	return false
}

// SetErrorMessageIdentifier gets a reference to the given GBTErrorStringEnum and assigns it to the ErrorMessageIdentifier field.
func (o *BTParameterExpressionErrorInfo2037) SetErrorMessageIdentifier(v GBTErrorStringEnum) {
	o.ErrorMessageIdentifier = &v
}

// GetMessageArguments returns the MessageArguments field value if set, zero value otherwise.
func (o *BTParameterExpressionErrorInfo2037) GetMessageArguments() []BTValueAndUse4696 {
	if o == nil || o.MessageArguments == nil {
		var ret []BTValueAndUse4696
		return ret
	}
	return o.MessageArguments
}

// GetMessageArgumentsOk returns a tuple with the MessageArguments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterExpressionErrorInfo2037) GetMessageArgumentsOk() ([]BTValueAndUse4696, bool) {
	if o == nil || o.MessageArguments == nil {
		return nil, false
	}
	return o.MessageArguments, true
}

// HasMessageArguments returns a boolean if a field has been set.
func (o *BTParameterExpressionErrorInfo2037) HasMessageArguments() bool {
	if o != nil && o.MessageArguments != nil {
		return true
	}

	return false
}

// SetMessageArguments gets a reference to the given []BTValueAndUse4696 and assigns it to the MessageArguments field.
func (o *BTParameterExpressionErrorInfo2037) SetMessageArguments(v []BTValueAndUse4696) {
	o.MessageArguments = v
}

func (o BTParameterExpressionErrorInfo2037) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.ErrorMessageIdentifier != nil {
		toSerialize["errorMessageIdentifier"] = o.ErrorMessageIdentifier
	}
	if o.MessageArguments != nil {
		toSerialize["messageArguments"] = o.MessageArguments
	}
	return json.Marshal(toSerialize)
}

type NullableBTParameterExpressionErrorInfo2037 struct {
	value *BTParameterExpressionErrorInfo2037
	isSet bool
}

func (v NullableBTParameterExpressionErrorInfo2037) Get() *BTParameterExpressionErrorInfo2037 {
	return v.value
}

func (v *NullableBTParameterExpressionErrorInfo2037) Set(val *BTParameterExpressionErrorInfo2037) {
	v.value = val
	v.isSet = true
}

func (v NullableBTParameterExpressionErrorInfo2037) IsSet() bool {
	return v.isSet
}

func (v *NullableBTParameterExpressionErrorInfo2037) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTParameterExpressionErrorInfo2037(val *BTParameterExpressionErrorInfo2037) *NullableBTParameterExpressionErrorInfo2037 {
	return &NullableBTParameterExpressionErrorInfo2037{value: val, isSet: true}
}

func (v NullableBTParameterExpressionErrorInfo2037) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTParameterExpressionErrorInfo2037) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
