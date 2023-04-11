/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.161.14306-351f5b17f026
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTJsonMatch2290 struct for BTJsonMatch2290
type BTJsonMatch2290 struct {
	DefiniteJsonPath *string              `json:"definiteJsonPath,omitempty"`
	Node             *BTJsonMatch2290Node `json:"node,omitempty"`
}

// NewBTJsonMatch2290 instantiates a new BTJsonMatch2290 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTJsonMatch2290() *BTJsonMatch2290 {
	this := BTJsonMatch2290{}
	return &this
}

// NewBTJsonMatch2290WithDefaults instantiates a new BTJsonMatch2290 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTJsonMatch2290WithDefaults() *BTJsonMatch2290 {
	this := BTJsonMatch2290{}
	return &this
}

// GetDefiniteJsonPath returns the DefiniteJsonPath field value if set, zero value otherwise.
func (o *BTJsonMatch2290) GetDefiniteJsonPath() string {
	if o == nil || o.DefiniteJsonPath == nil {
		var ret string
		return ret
	}
	return *o.DefiniteJsonPath
}

// GetDefiniteJsonPathOk returns a tuple with the DefiniteJsonPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTJsonMatch2290) GetDefiniteJsonPathOk() (*string, bool) {
	if o == nil || o.DefiniteJsonPath == nil {
		return nil, false
	}
	return o.DefiniteJsonPath, true
}

// HasDefiniteJsonPath returns a boolean if a field has been set.
func (o *BTJsonMatch2290) HasDefiniteJsonPath() bool {
	if o != nil && o.DefiniteJsonPath != nil {
		return true
	}

	return false
}

// SetDefiniteJsonPath gets a reference to the given string and assigns it to the DefiniteJsonPath field.
func (o *BTJsonMatch2290) SetDefiniteJsonPath(v string) {
	o.DefiniteJsonPath = &v
}

// GetNode returns the Node field value if set, zero value otherwise.
func (o *BTJsonMatch2290) GetNode() BTJsonMatch2290Node {
	if o == nil || o.Node == nil {
		var ret BTJsonMatch2290Node
		return ret
	}
	return *o.Node
}

// GetNodeOk returns a tuple with the Node field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTJsonMatch2290) GetNodeOk() (*BTJsonMatch2290Node, bool) {
	if o == nil || o.Node == nil {
		return nil, false
	}
	return o.Node, true
}

// HasNode returns a boolean if a field has been set.
func (o *BTJsonMatch2290) HasNode() bool {
	if o != nil && o.Node != nil {
		return true
	}

	return false
}

// SetNode gets a reference to the given BTJsonMatch2290Node and assigns it to the Node field.
func (o *BTJsonMatch2290) SetNode(v BTJsonMatch2290Node) {
	o.Node = &v
}

func (o BTJsonMatch2290) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DefiniteJsonPath != nil {
		toSerialize["definiteJsonPath"] = o.DefiniteJsonPath
	}
	if o.Node != nil {
		toSerialize["node"] = o.Node
	}
	return json.Marshal(toSerialize)
}

type NullableBTJsonMatch2290 struct {
	value *BTJsonMatch2290
	isSet bool
}

func (v NullableBTJsonMatch2290) Get() *BTJsonMatch2290 {
	return v.value
}

func (v *NullableBTJsonMatch2290) Set(val *BTJsonMatch2290) {
	v.value = val
	v.isSet = true
}

func (v NullableBTJsonMatch2290) IsSet() bool {
	return v.isSet
}

func (v *NullableBTJsonMatch2290) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTJsonMatch2290(val *BTJsonMatch2290) *NullableBTJsonMatch2290 {
	return &NullableBTJsonMatch2290{value: val, isSet: true}
}

func (v NullableBTJsonMatch2290) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTJsonMatch2290) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
