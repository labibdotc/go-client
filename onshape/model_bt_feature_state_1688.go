/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.152.6085-0290120322c3
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTFeatureState1688 struct for BTFeatureState1688
type BTFeatureState1688 struct {
	FeatureStatus *string `json:"featureStatus,omitempty"`
}

// NewBTFeatureState1688 instantiates a new BTFeatureState1688 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTFeatureState1688() *BTFeatureState1688 {
	this := BTFeatureState1688{}
	return &this
}

// NewBTFeatureState1688WithDefaults instantiates a new BTFeatureState1688 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTFeatureState1688WithDefaults() *BTFeatureState1688 {
	this := BTFeatureState1688{}
	return &this
}

// GetFeatureStatus returns the FeatureStatus field value if set, zero value otherwise.
func (o *BTFeatureState1688) GetFeatureStatus() string {
	if o == nil || o.FeatureStatus == nil {
		var ret string
		return ret
	}
	return *o.FeatureStatus
}

// GetFeatureStatusOk returns a tuple with the FeatureStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFeatureState1688) GetFeatureStatusOk() (*string, bool) {
	if o == nil || o.FeatureStatus == nil {
		return nil, false
	}
	return o.FeatureStatus, true
}

// HasFeatureStatus returns a boolean if a field has been set.
func (o *BTFeatureState1688) HasFeatureStatus() bool {
	if o != nil && o.FeatureStatus != nil {
		return true
	}

	return false
}

// SetFeatureStatus gets a reference to the given string and assigns it to the FeatureStatus field.
func (o *BTFeatureState1688) SetFeatureStatus(v string) {
	o.FeatureStatus = &v
}

func (o BTFeatureState1688) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.FeatureStatus != nil {
		toSerialize["featureStatus"] = o.FeatureStatus
	}
	return json.Marshal(toSerialize)
}

type NullableBTFeatureState1688 struct {
	value *BTFeatureState1688
	isSet bool
}

func (v NullableBTFeatureState1688) Get() *BTFeatureState1688 {
	return v.value
}

func (v *NullableBTFeatureState1688) Set(val *BTFeatureState1688) {
	v.value = val
	v.isSet = true
}

func (v NullableBTFeatureState1688) IsSet() bool {
	return v.isSet
}

func (v *NullableBTFeatureState1688) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTFeatureState1688(val *BTFeatureState1688) *NullableBTFeatureState1688 {
	return &NullableBTFeatureState1688{value: val, isSet: true}
}

func (v NullableBTFeatureState1688) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTFeatureState1688) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
