/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.156.7250-f7937557e62d
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTToleranceString3274 struct for BTToleranceString3274
type BTToleranceString3274 struct {
	BtType             *string `json:"btType,omitempty"`
	ImportMicroversion *string `json:"importMicroversion,omitempty"`
	Lower              *string `json:"lower,omitempty"`
	NodeId             *string `json:"nodeId,omitempty"`
	Upper              *string `json:"upper,omitempty"`
	Value              *string `json:"value,omitempty"`
}

// NewBTToleranceString3274 instantiates a new BTToleranceString3274 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTToleranceString3274() *BTToleranceString3274 {
	this := BTToleranceString3274{}
	return &this
}

// NewBTToleranceString3274WithDefaults instantiates a new BTToleranceString3274 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTToleranceString3274WithDefaults() *BTToleranceString3274 {
	this := BTToleranceString3274{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTToleranceString3274) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTToleranceString3274) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTToleranceString3274) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTToleranceString3274) SetBtType(v string) {
	o.BtType = &v
}

// GetImportMicroversion returns the ImportMicroversion field value if set, zero value otherwise.
func (o *BTToleranceString3274) GetImportMicroversion() string {
	if o == nil || o.ImportMicroversion == nil {
		var ret string
		return ret
	}
	return *o.ImportMicroversion
}

// GetImportMicroversionOk returns a tuple with the ImportMicroversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTToleranceString3274) GetImportMicroversionOk() (*string, bool) {
	if o == nil || o.ImportMicroversion == nil {
		return nil, false
	}
	return o.ImportMicroversion, true
}

// HasImportMicroversion returns a boolean if a field has been set.
func (o *BTToleranceString3274) HasImportMicroversion() bool {
	if o != nil && o.ImportMicroversion != nil {
		return true
	}

	return false
}

// SetImportMicroversion gets a reference to the given string and assigns it to the ImportMicroversion field.
func (o *BTToleranceString3274) SetImportMicroversion(v string) {
	o.ImportMicroversion = &v
}

// GetLower returns the Lower field value if set, zero value otherwise.
func (o *BTToleranceString3274) GetLower() string {
	if o == nil || o.Lower == nil {
		var ret string
		return ret
	}
	return *o.Lower
}

// GetLowerOk returns a tuple with the Lower field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTToleranceString3274) GetLowerOk() (*string, bool) {
	if o == nil || o.Lower == nil {
		return nil, false
	}
	return o.Lower, true
}

// HasLower returns a boolean if a field has been set.
func (o *BTToleranceString3274) HasLower() bool {
	if o != nil && o.Lower != nil {
		return true
	}

	return false
}

// SetLower gets a reference to the given string and assigns it to the Lower field.
func (o *BTToleranceString3274) SetLower(v string) {
	o.Lower = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *BTToleranceString3274) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTToleranceString3274) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *BTToleranceString3274) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *BTToleranceString3274) SetNodeId(v string) {
	o.NodeId = &v
}

// GetUpper returns the Upper field value if set, zero value otherwise.
func (o *BTToleranceString3274) GetUpper() string {
	if o == nil || o.Upper == nil {
		var ret string
		return ret
	}
	return *o.Upper
}

// GetUpperOk returns a tuple with the Upper field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTToleranceString3274) GetUpperOk() (*string, bool) {
	if o == nil || o.Upper == nil {
		return nil, false
	}
	return o.Upper, true
}

// HasUpper returns a boolean if a field has been set.
func (o *BTToleranceString3274) HasUpper() bool {
	if o != nil && o.Upper != nil {
		return true
	}

	return false
}

// SetUpper gets a reference to the given string and assigns it to the Upper field.
func (o *BTToleranceString3274) SetUpper(v string) {
	o.Upper = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *BTToleranceString3274) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTToleranceString3274) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *BTToleranceString3274) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *BTToleranceString3274) SetValue(v string) {
	o.Value = &v
}

func (o BTToleranceString3274) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.ImportMicroversion != nil {
		toSerialize["importMicroversion"] = o.ImportMicroversion
	}
	if o.Lower != nil {
		toSerialize["lower"] = o.Lower
	}
	if o.NodeId != nil {
		toSerialize["nodeId"] = o.NodeId
	}
	if o.Upper != nil {
		toSerialize["upper"] = o.Upper
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableBTToleranceString3274 struct {
	value *BTToleranceString3274
	isSet bool
}

func (v NullableBTToleranceString3274) Get() *BTToleranceString3274 {
	return v.value
}

func (v *NullableBTToleranceString3274) Set(val *BTToleranceString3274) {
	v.value = val
	v.isSet = true
}

func (v NullableBTToleranceString3274) IsSet() bool {
	return v.isSet
}

func (v *NullableBTToleranceString3274) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTToleranceString3274(val *BTToleranceString3274) *NullableBTToleranceString3274 {
	return &NullableBTToleranceString3274{value: val, isSet: true}
}

func (v NullableBTToleranceString3274) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTToleranceString3274) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
