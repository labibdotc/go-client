/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.166.18417-1bd990c6fbaa
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTNodeReference21 struct for BTNodeReference21
type BTNodeReference21 struct {
	BtType    *string     `json:"btType,omitempty"`
	NodeId    *string     `json:"nodeId,omitempty"`
	NodeIdRaw *BTObjectId `json:"nodeIdRaw,omitempty"`
}

// NewBTNodeReference21 instantiates a new BTNodeReference21 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTNodeReference21() *BTNodeReference21 {
	this := BTNodeReference21{}
	return &this
}

// NewBTNodeReference21WithDefaults instantiates a new BTNodeReference21 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTNodeReference21WithDefaults() *BTNodeReference21 {
	this := BTNodeReference21{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTNodeReference21) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTNodeReference21) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTNodeReference21) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTNodeReference21) SetBtType(v string) {
	o.BtType = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *BTNodeReference21) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTNodeReference21) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *BTNodeReference21) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *BTNodeReference21) SetNodeId(v string) {
	o.NodeId = &v
}

// GetNodeIdRaw returns the NodeIdRaw field value if set, zero value otherwise.
func (o *BTNodeReference21) GetNodeIdRaw() BTObjectId {
	if o == nil || o.NodeIdRaw == nil {
		var ret BTObjectId
		return ret
	}
	return *o.NodeIdRaw
}

// GetNodeIdRawOk returns a tuple with the NodeIdRaw field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTNodeReference21) GetNodeIdRawOk() (*BTObjectId, bool) {
	if o == nil || o.NodeIdRaw == nil {
		return nil, false
	}
	return o.NodeIdRaw, true
}

// HasNodeIdRaw returns a boolean if a field has been set.
func (o *BTNodeReference21) HasNodeIdRaw() bool {
	if o != nil && o.NodeIdRaw != nil {
		return true
	}

	return false
}

// SetNodeIdRaw gets a reference to the given BTObjectId and assigns it to the NodeIdRaw field.
func (o *BTNodeReference21) SetNodeIdRaw(v BTObjectId) {
	o.NodeIdRaw = &v
}

func (o BTNodeReference21) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.NodeId != nil {
		toSerialize["nodeId"] = o.NodeId
	}
	if o.NodeIdRaw != nil {
		toSerialize["nodeIdRaw"] = o.NodeIdRaw
	}
	return json.Marshal(toSerialize)
}

type NullableBTNodeReference21 struct {
	value *BTNodeReference21
	isSet bool
}

func (v NullableBTNodeReference21) Get() *BTNodeReference21 {
	return v.value
}

func (v *NullableBTNodeReference21) Set(val *BTNodeReference21) {
	v.value = val
	v.isSet = true
}

func (v NullableBTNodeReference21) IsSet() bool {
	return v.isSet
}

func (v *NullableBTNodeReference21) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTNodeReference21(val *BTNodeReference21) *NullableBTNodeReference21 {
	return &NullableBTNodeReference21{value: val, isSet: true}
}

func (v NullableBTNodeReference21) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTNodeReference21) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
