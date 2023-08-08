/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.167.20169-88260985a0b6
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTMArrayParameterItem1843 struct for BTMArrayParameterItem1843
type BTMArrayParameterItem1843 struct {
	BtType             *string `json:"btType,omitempty"`
	ImportMicroversion *string `json:"importMicroversion,omitempty"`
	NodeId             *string `json:"nodeId,omitempty"`
}

// NewBTMArrayParameterItem1843 instantiates a new BTMArrayParameterItem1843 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTMArrayParameterItem1843() *BTMArrayParameterItem1843 {
	this := BTMArrayParameterItem1843{}
	return &this
}

// NewBTMArrayParameterItem1843WithDefaults instantiates a new BTMArrayParameterItem1843 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTMArrayParameterItem1843WithDefaults() *BTMArrayParameterItem1843 {
	this := BTMArrayParameterItem1843{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTMArrayParameterItem1843) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMArrayParameterItem1843) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTMArrayParameterItem1843) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTMArrayParameterItem1843) SetBtType(v string) {
	o.BtType = &v
}

// GetImportMicroversion returns the ImportMicroversion field value if set, zero value otherwise.
func (o *BTMArrayParameterItem1843) GetImportMicroversion() string {
	if o == nil || o.ImportMicroversion == nil {
		var ret string
		return ret
	}
	return *o.ImportMicroversion
}

// GetImportMicroversionOk returns a tuple with the ImportMicroversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMArrayParameterItem1843) GetImportMicroversionOk() (*string, bool) {
	if o == nil || o.ImportMicroversion == nil {
		return nil, false
	}
	return o.ImportMicroversion, true
}

// HasImportMicroversion returns a boolean if a field has been set.
func (o *BTMArrayParameterItem1843) HasImportMicroversion() bool {
	if o != nil && o.ImportMicroversion != nil {
		return true
	}

	return false
}

// SetImportMicroversion gets a reference to the given string and assigns it to the ImportMicroversion field.
func (o *BTMArrayParameterItem1843) SetImportMicroversion(v string) {
	o.ImportMicroversion = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *BTMArrayParameterItem1843) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMArrayParameterItem1843) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *BTMArrayParameterItem1843) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *BTMArrayParameterItem1843) SetNodeId(v string) {
	o.NodeId = &v
}

func (o BTMArrayParameterItem1843) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.ImportMicroversion != nil {
		toSerialize["importMicroversion"] = o.ImportMicroversion
	}
	if o.NodeId != nil {
		toSerialize["nodeId"] = o.NodeId
	}
	return json.Marshal(toSerialize)
}

type NullableBTMArrayParameterItem1843 struct {
	value *BTMArrayParameterItem1843
	isSet bool
}

func (v NullableBTMArrayParameterItem1843) Get() *BTMArrayParameterItem1843 {
	return v.value
}

func (v *NullableBTMArrayParameterItem1843) Set(val *BTMArrayParameterItem1843) {
	v.value = val
	v.isSet = true
}

func (v NullableBTMArrayParameterItem1843) IsSet() bool {
	return v.isSet
}

func (v *NullableBTMArrayParameterItem1843) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTMArrayParameterItem1843(val *BTMArrayParameterItem1843) *NullableBTMArrayParameterItem1843 {
	return &NullableBTMArrayParameterItem1843{value: val, isSet: true}
}

func (v NullableBTMArrayParameterItem1843) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTMArrayParameterItem1843) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
