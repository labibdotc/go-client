/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.162.14462-13ace71ec1df
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTMParameterLookupTablePath1419 struct for BTMParameterLookupTablePath1419
type BTMParameterLookupTablePath1419 struct {
	ImportMicroversion *string            `json:"importMicroversion,omitempty"`
	NodeId             *string            `json:"nodeId,omitempty"`
	ParameterId        *string            `json:"parameterId,omitempty"`
	BtType             *string            `json:"btType,omitempty"`
	Value              *map[string]string `json:"value,omitempty"`
}

// NewBTMParameterLookupTablePath1419 instantiates a new BTMParameterLookupTablePath1419 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTMParameterLookupTablePath1419() *BTMParameterLookupTablePath1419 {
	this := BTMParameterLookupTablePath1419{}
	return &this
}

// NewBTMParameterLookupTablePath1419WithDefaults instantiates a new BTMParameterLookupTablePath1419 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTMParameterLookupTablePath1419WithDefaults() *BTMParameterLookupTablePath1419 {
	this := BTMParameterLookupTablePath1419{}
	return &this
}

// GetImportMicroversion returns the ImportMicroversion field value if set, zero value otherwise.
func (o *BTMParameterLookupTablePath1419) GetImportMicroversion() string {
	if o == nil || o.ImportMicroversion == nil {
		var ret string
		return ret
	}
	return *o.ImportMicroversion
}

// GetImportMicroversionOk returns a tuple with the ImportMicroversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMParameterLookupTablePath1419) GetImportMicroversionOk() (*string, bool) {
	if o == nil || o.ImportMicroversion == nil {
		return nil, false
	}
	return o.ImportMicroversion, true
}

// HasImportMicroversion returns a boolean if a field has been set.
func (o *BTMParameterLookupTablePath1419) HasImportMicroversion() bool {
	if o != nil && o.ImportMicroversion != nil {
		return true
	}

	return false
}

// SetImportMicroversion gets a reference to the given string and assigns it to the ImportMicroversion field.
func (o *BTMParameterLookupTablePath1419) SetImportMicroversion(v string) {
	o.ImportMicroversion = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *BTMParameterLookupTablePath1419) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMParameterLookupTablePath1419) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *BTMParameterLookupTablePath1419) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *BTMParameterLookupTablePath1419) SetNodeId(v string) {
	o.NodeId = &v
}

// GetParameterId returns the ParameterId field value if set, zero value otherwise.
func (o *BTMParameterLookupTablePath1419) GetParameterId() string {
	if o == nil || o.ParameterId == nil {
		var ret string
		return ret
	}
	return *o.ParameterId
}

// GetParameterIdOk returns a tuple with the ParameterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMParameterLookupTablePath1419) GetParameterIdOk() (*string, bool) {
	if o == nil || o.ParameterId == nil {
		return nil, false
	}
	return o.ParameterId, true
}

// HasParameterId returns a boolean if a field has been set.
func (o *BTMParameterLookupTablePath1419) HasParameterId() bool {
	if o != nil && o.ParameterId != nil {
		return true
	}

	return false
}

// SetParameterId gets a reference to the given string and assigns it to the ParameterId field.
func (o *BTMParameterLookupTablePath1419) SetParameterId(v string) {
	o.ParameterId = &v
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTMParameterLookupTablePath1419) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMParameterLookupTablePath1419) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTMParameterLookupTablePath1419) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTMParameterLookupTablePath1419) SetBtType(v string) {
	o.BtType = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *BTMParameterLookupTablePath1419) GetValue() map[string]string {
	if o == nil || o.Value == nil {
		var ret map[string]string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMParameterLookupTablePath1419) GetValueOk() (*map[string]string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *BTMParameterLookupTablePath1419) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given map[string]string and assigns it to the Value field.
func (o *BTMParameterLookupTablePath1419) SetValue(v map[string]string) {
	o.Value = &v
}

func (o BTMParameterLookupTablePath1419) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ImportMicroversion != nil {
		toSerialize["importMicroversion"] = o.ImportMicroversion
	}
	if o.NodeId != nil {
		toSerialize["nodeId"] = o.NodeId
	}
	if o.ParameterId != nil {
		toSerialize["parameterId"] = o.ParameterId
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableBTMParameterLookupTablePath1419 struct {
	value *BTMParameterLookupTablePath1419
	isSet bool
}

func (v NullableBTMParameterLookupTablePath1419) Get() *BTMParameterLookupTablePath1419 {
	return v.value
}

func (v *NullableBTMParameterLookupTablePath1419) Set(val *BTMParameterLookupTablePath1419) {
	v.value = val
	v.isSet = true
}

func (v NullableBTMParameterLookupTablePath1419) IsSet() bool {
	return v.isSet
}

func (v *NullableBTMParameterLookupTablePath1419) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTMParameterLookupTablePath1419(val *BTMParameterLookupTablePath1419) *NullableBTMParameterLookupTablePath1419 {
	return &NullableBTMParameterLookupTablePath1419{value: val, isSet: true}
}

func (v NullableBTMParameterLookupTablePath1419) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTMParameterLookupTablePath1419) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
