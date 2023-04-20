/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.162.14806-89d807e7089c
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTStringFormatMatchPattern2446 struct for BTStringFormatMatchPattern2446
type BTStringFormatMatchPattern2446 struct {
	BtType                        *string `json:"btType,omitempty"`
	ErrorMessage                  *string `json:"errorMessage,omitempty"`
	ShouldResetValueWhenConfirmed *bool   `json:"shouldResetValueWhenConfirmed,omitempty"`
	RegExpToMatch                 *string `json:"regExpToMatch,omitempty"`
}

// NewBTStringFormatMatchPattern2446 instantiates a new BTStringFormatMatchPattern2446 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTStringFormatMatchPattern2446() *BTStringFormatMatchPattern2446 {
	this := BTStringFormatMatchPattern2446{}
	return &this
}

// NewBTStringFormatMatchPattern2446WithDefaults instantiates a new BTStringFormatMatchPattern2446 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTStringFormatMatchPattern2446WithDefaults() *BTStringFormatMatchPattern2446 {
	this := BTStringFormatMatchPattern2446{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTStringFormatMatchPattern2446) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTStringFormatMatchPattern2446) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTStringFormatMatchPattern2446) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTStringFormatMatchPattern2446) SetBtType(v string) {
	o.BtType = &v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *BTStringFormatMatchPattern2446) GetErrorMessage() string {
	if o == nil || o.ErrorMessage == nil {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTStringFormatMatchPattern2446) GetErrorMessageOk() (*string, bool) {
	if o == nil || o.ErrorMessage == nil {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *BTStringFormatMatchPattern2446) HasErrorMessage() bool {
	if o != nil && o.ErrorMessage != nil {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *BTStringFormatMatchPattern2446) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

// GetShouldResetValueWhenConfirmed returns the ShouldResetValueWhenConfirmed field value if set, zero value otherwise.
func (o *BTStringFormatMatchPattern2446) GetShouldResetValueWhenConfirmed() bool {
	if o == nil || o.ShouldResetValueWhenConfirmed == nil {
		var ret bool
		return ret
	}
	return *o.ShouldResetValueWhenConfirmed
}

// GetShouldResetValueWhenConfirmedOk returns a tuple with the ShouldResetValueWhenConfirmed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTStringFormatMatchPattern2446) GetShouldResetValueWhenConfirmedOk() (*bool, bool) {
	if o == nil || o.ShouldResetValueWhenConfirmed == nil {
		return nil, false
	}
	return o.ShouldResetValueWhenConfirmed, true
}

// HasShouldResetValueWhenConfirmed returns a boolean if a field has been set.
func (o *BTStringFormatMatchPattern2446) HasShouldResetValueWhenConfirmed() bool {
	if o != nil && o.ShouldResetValueWhenConfirmed != nil {
		return true
	}

	return false
}

// SetShouldResetValueWhenConfirmed gets a reference to the given bool and assigns it to the ShouldResetValueWhenConfirmed field.
func (o *BTStringFormatMatchPattern2446) SetShouldResetValueWhenConfirmed(v bool) {
	o.ShouldResetValueWhenConfirmed = &v
}

// GetRegExpToMatch returns the RegExpToMatch field value if set, zero value otherwise.
func (o *BTStringFormatMatchPattern2446) GetRegExpToMatch() string {
	if o == nil || o.RegExpToMatch == nil {
		var ret string
		return ret
	}
	return *o.RegExpToMatch
}

// GetRegExpToMatchOk returns a tuple with the RegExpToMatch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTStringFormatMatchPattern2446) GetRegExpToMatchOk() (*string, bool) {
	if o == nil || o.RegExpToMatch == nil {
		return nil, false
	}
	return o.RegExpToMatch, true
}

// HasRegExpToMatch returns a boolean if a field has been set.
func (o *BTStringFormatMatchPattern2446) HasRegExpToMatch() bool {
	if o != nil && o.RegExpToMatch != nil {
		return true
	}

	return false
}

// SetRegExpToMatch gets a reference to the given string and assigns it to the RegExpToMatch field.
func (o *BTStringFormatMatchPattern2446) SetRegExpToMatch(v string) {
	o.RegExpToMatch = &v
}

func (o BTStringFormatMatchPattern2446) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.ErrorMessage != nil {
		toSerialize["errorMessage"] = o.ErrorMessage
	}
	if o.ShouldResetValueWhenConfirmed != nil {
		toSerialize["shouldResetValueWhenConfirmed"] = o.ShouldResetValueWhenConfirmed
	}
	if o.RegExpToMatch != nil {
		toSerialize["regExpToMatch"] = o.RegExpToMatch
	}
	return json.Marshal(toSerialize)
}

type NullableBTStringFormatMatchPattern2446 struct {
	value *BTStringFormatMatchPattern2446
	isSet bool
}

func (v NullableBTStringFormatMatchPattern2446) Get() *BTStringFormatMatchPattern2446 {
	return v.value
}

func (v *NullableBTStringFormatMatchPattern2446) Set(val *BTStringFormatMatchPattern2446) {
	v.value = val
	v.isSet = true
}

func (v NullableBTStringFormatMatchPattern2446) IsSet() bool {
	return v.isSet
}

func (v *NullableBTStringFormatMatchPattern2446) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTStringFormatMatchPattern2446(val *BTStringFormatMatchPattern2446) *NullableBTStringFormatMatchPattern2446 {
	return &NullableBTStringFormatMatchPattern2446{value: val, isSet: true}
}

func (v NullableBTStringFormatMatchPattern2446) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTStringFormatMatchPattern2446) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
