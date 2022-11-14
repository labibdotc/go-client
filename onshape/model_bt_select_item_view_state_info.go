/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.155.7328-9aa102855752
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTSelectItemViewStateInfo struct for BTSelectItemViewStateInfo
type BTSelectItemViewStateInfo struct {
	ActiveSelectorId  *string                  `json:"activeSelectorId,omitempty"`
	DocumentSelectors []BTDocumentSelectorInfo `json:"documentSelectors,omitempty"`
	Purpose           *string                  `json:"purpose,omitempty"`
}

// NewBTSelectItemViewStateInfo instantiates a new BTSelectItemViewStateInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTSelectItemViewStateInfo() *BTSelectItemViewStateInfo {
	this := BTSelectItemViewStateInfo{}
	return &this
}

// NewBTSelectItemViewStateInfoWithDefaults instantiates a new BTSelectItemViewStateInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTSelectItemViewStateInfoWithDefaults() *BTSelectItemViewStateInfo {
	this := BTSelectItemViewStateInfo{}
	return &this
}

// GetActiveSelectorId returns the ActiveSelectorId field value if set, zero value otherwise.
func (o *BTSelectItemViewStateInfo) GetActiveSelectorId() string {
	if o == nil || o.ActiveSelectorId == nil {
		var ret string
		return ret
	}
	return *o.ActiveSelectorId
}

// GetActiveSelectorIdOk returns a tuple with the ActiveSelectorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSelectItemViewStateInfo) GetActiveSelectorIdOk() (*string, bool) {
	if o == nil || o.ActiveSelectorId == nil {
		return nil, false
	}
	return o.ActiveSelectorId, true
}

// HasActiveSelectorId returns a boolean if a field has been set.
func (o *BTSelectItemViewStateInfo) HasActiveSelectorId() bool {
	if o != nil && o.ActiveSelectorId != nil {
		return true
	}

	return false
}

// SetActiveSelectorId gets a reference to the given string and assigns it to the ActiveSelectorId field.
func (o *BTSelectItemViewStateInfo) SetActiveSelectorId(v string) {
	o.ActiveSelectorId = &v
}

// GetDocumentSelectors returns the DocumentSelectors field value if set, zero value otherwise.
func (o *BTSelectItemViewStateInfo) GetDocumentSelectors() []BTDocumentSelectorInfo {
	if o == nil || o.DocumentSelectors == nil {
		var ret []BTDocumentSelectorInfo
		return ret
	}
	return o.DocumentSelectors
}

// GetDocumentSelectorsOk returns a tuple with the DocumentSelectors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSelectItemViewStateInfo) GetDocumentSelectorsOk() ([]BTDocumentSelectorInfo, bool) {
	if o == nil || o.DocumentSelectors == nil {
		return nil, false
	}
	return o.DocumentSelectors, true
}

// HasDocumentSelectors returns a boolean if a field has been set.
func (o *BTSelectItemViewStateInfo) HasDocumentSelectors() bool {
	if o != nil && o.DocumentSelectors != nil {
		return true
	}

	return false
}

// SetDocumentSelectors gets a reference to the given []BTDocumentSelectorInfo and assigns it to the DocumentSelectors field.
func (o *BTSelectItemViewStateInfo) SetDocumentSelectors(v []BTDocumentSelectorInfo) {
	o.DocumentSelectors = v
}

// GetPurpose returns the Purpose field value if set, zero value otherwise.
func (o *BTSelectItemViewStateInfo) GetPurpose() string {
	if o == nil || o.Purpose == nil {
		var ret string
		return ret
	}
	return *o.Purpose
}

// GetPurposeOk returns a tuple with the Purpose field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSelectItemViewStateInfo) GetPurposeOk() (*string, bool) {
	if o == nil || o.Purpose == nil {
		return nil, false
	}
	return o.Purpose, true
}

// HasPurpose returns a boolean if a field has been set.
func (o *BTSelectItemViewStateInfo) HasPurpose() bool {
	if o != nil && o.Purpose != nil {
		return true
	}

	return false
}

// SetPurpose gets a reference to the given string and assigns it to the Purpose field.
func (o *BTSelectItemViewStateInfo) SetPurpose(v string) {
	o.Purpose = &v
}

func (o BTSelectItemViewStateInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ActiveSelectorId != nil {
		toSerialize["activeSelectorId"] = o.ActiveSelectorId
	}
	if o.DocumentSelectors != nil {
		toSerialize["documentSelectors"] = o.DocumentSelectors
	}
	if o.Purpose != nil {
		toSerialize["purpose"] = o.Purpose
	}
	return json.Marshal(toSerialize)
}

type NullableBTSelectItemViewStateInfo struct {
	value *BTSelectItemViewStateInfo
	isSet bool
}

func (v NullableBTSelectItemViewStateInfo) Get() *BTSelectItemViewStateInfo {
	return v.value
}

func (v *NullableBTSelectItemViewStateInfo) Set(val *BTSelectItemViewStateInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTSelectItemViewStateInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTSelectItemViewStateInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTSelectItemViewStateInfo(val *BTSelectItemViewStateInfo) *NullableBTSelectItemViewStateInfo {
	return &NullableBTSelectItemViewStateInfo{value: val, isSet: true}
}

func (v NullableBTSelectItemViewStateInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTSelectItemViewStateInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
