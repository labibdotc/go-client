/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.163.15808-38acf80dff96
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTSystemPartColorCycle1580 struct for BTSystemPartColorCycle1580
type BTSystemPartColorCycle1580 struct {
	BtType  *string                   `json:"btType,omitempty"`
	Version *GBTPartColorCycleVersion `json:"version,omitempty"`
}

// NewBTSystemPartColorCycle1580 instantiates a new BTSystemPartColorCycle1580 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTSystemPartColorCycle1580() *BTSystemPartColorCycle1580 {
	this := BTSystemPartColorCycle1580{}
	return &this
}

// NewBTSystemPartColorCycle1580WithDefaults instantiates a new BTSystemPartColorCycle1580 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTSystemPartColorCycle1580WithDefaults() *BTSystemPartColorCycle1580 {
	this := BTSystemPartColorCycle1580{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTSystemPartColorCycle1580) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSystemPartColorCycle1580) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTSystemPartColorCycle1580) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTSystemPartColorCycle1580) SetBtType(v string) {
	o.BtType = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *BTSystemPartColorCycle1580) GetVersion() GBTPartColorCycleVersion {
	if o == nil || o.Version == nil {
		var ret GBTPartColorCycleVersion
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSystemPartColorCycle1580) GetVersionOk() (*GBTPartColorCycleVersion, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *BTSystemPartColorCycle1580) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given GBTPartColorCycleVersion and assigns it to the Version field.
func (o *BTSystemPartColorCycle1580) SetVersion(v GBTPartColorCycleVersion) {
	o.Version = &v
}

func (o BTSystemPartColorCycle1580) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableBTSystemPartColorCycle1580 struct {
	value *BTSystemPartColorCycle1580
	isSet bool
}

func (v NullableBTSystemPartColorCycle1580) Get() *BTSystemPartColorCycle1580 {
	return v.value
}

func (v *NullableBTSystemPartColorCycle1580) Set(val *BTSystemPartColorCycle1580) {
	v.value = val
	v.isSet = true
}

func (v NullableBTSystemPartColorCycle1580) IsSet() bool {
	return v.isSet
}

func (v *NullableBTSystemPartColorCycle1580) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTSystemPartColorCycle1580(val *BTSystemPartColorCycle1580) *NullableBTSystemPartColorCycle1580 {
	return &NullableBTSystemPartColorCycle1580{value: val, isSet: true}
}

func (v NullableBTSystemPartColorCycle1580) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTSystemPartColorCycle1580) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
