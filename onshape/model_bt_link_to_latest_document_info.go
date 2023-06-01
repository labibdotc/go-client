/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.164.16643-ef813b2da145
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTLinkToLatestDocumentInfo struct for BTLinkToLatestDocumentInfo
type BTLinkToLatestDocumentInfo struct {
	ChangedElements []string `json:"changedElements,omitempty"`
}

// NewBTLinkToLatestDocumentInfo instantiates a new BTLinkToLatestDocumentInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTLinkToLatestDocumentInfo() *BTLinkToLatestDocumentInfo {
	this := BTLinkToLatestDocumentInfo{}
	return &this
}

// NewBTLinkToLatestDocumentInfoWithDefaults instantiates a new BTLinkToLatestDocumentInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTLinkToLatestDocumentInfoWithDefaults() *BTLinkToLatestDocumentInfo {
	this := BTLinkToLatestDocumentInfo{}
	return &this
}

// GetChangedElements returns the ChangedElements field value if set, zero value otherwise.
func (o *BTLinkToLatestDocumentInfo) GetChangedElements() []string {
	if o == nil || o.ChangedElements == nil {
		var ret []string
		return ret
	}
	return o.ChangedElements
}

// GetChangedElementsOk returns a tuple with the ChangedElements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTLinkToLatestDocumentInfo) GetChangedElementsOk() ([]string, bool) {
	if o == nil || o.ChangedElements == nil {
		return nil, false
	}
	return o.ChangedElements, true
}

// HasChangedElements returns a boolean if a field has been set.
func (o *BTLinkToLatestDocumentInfo) HasChangedElements() bool {
	if o != nil && o.ChangedElements != nil {
		return true
	}

	return false
}

// SetChangedElements gets a reference to the given []string and assigns it to the ChangedElements field.
func (o *BTLinkToLatestDocumentInfo) SetChangedElements(v []string) {
	o.ChangedElements = v
}

func (o BTLinkToLatestDocumentInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ChangedElements != nil {
		toSerialize["changedElements"] = o.ChangedElements
	}
	return json.Marshal(toSerialize)
}

type NullableBTLinkToLatestDocumentInfo struct {
	value *BTLinkToLatestDocumentInfo
	isSet bool
}

func (v NullableBTLinkToLatestDocumentInfo) Get() *BTLinkToLatestDocumentInfo {
	return v.value
}

func (v *NullableBTLinkToLatestDocumentInfo) Set(val *BTLinkToLatestDocumentInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTLinkToLatestDocumentInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTLinkToLatestDocumentInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTLinkToLatestDocumentInfo(val *BTLinkToLatestDocumentInfo) *NullableBTLinkToLatestDocumentInfo {
	return &NullableBTLinkToLatestDocumentInfo{value: val, isSet: true}
}

func (v NullableBTLinkToLatestDocumentInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTLinkToLatestDocumentInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
