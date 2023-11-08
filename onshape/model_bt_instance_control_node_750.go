/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.172.25541-b54c5b71ef45
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTInstanceControlNode750 struct for BTInstanceControlNode750
type BTInstanceControlNode750 struct {
	BtType                *string                  `json:"btType,omitempty"`
	ImportMicroversion    *string                  `json:"importMicroversion,omitempty"`
	NodeId                *string                  `json:"nodeId,omitempty"`
	Suppressed            *bool                    `json:"suppressed,omitempty"`
	SuppressedFieldIndex  *int32                   `json:"suppressedFieldIndex,omitempty"`
	SuppressionConfigured *bool                    `json:"suppressionConfigured,omitempty"`
	SuppressionState      *BTMSuppressionState1924 `json:"suppressionState,omitempty"`
}

// NewBTInstanceControlNode750 instantiates a new BTInstanceControlNode750 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTInstanceControlNode750() *BTInstanceControlNode750 {
	this := BTInstanceControlNode750{}
	return &this
}

// NewBTInstanceControlNode750WithDefaults instantiates a new BTInstanceControlNode750 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTInstanceControlNode750WithDefaults() *BTInstanceControlNode750 {
	this := BTInstanceControlNode750{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTInstanceControlNode750) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTInstanceControlNode750) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTInstanceControlNode750) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTInstanceControlNode750) SetBtType(v string) {
	o.BtType = &v
}

// GetImportMicroversion returns the ImportMicroversion field value if set, zero value otherwise.
func (o *BTInstanceControlNode750) GetImportMicroversion() string {
	if o == nil || o.ImportMicroversion == nil {
		var ret string
		return ret
	}
	return *o.ImportMicroversion
}

// GetImportMicroversionOk returns a tuple with the ImportMicroversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTInstanceControlNode750) GetImportMicroversionOk() (*string, bool) {
	if o == nil || o.ImportMicroversion == nil {
		return nil, false
	}
	return o.ImportMicroversion, true
}

// HasImportMicroversion returns a boolean if a field has been set.
func (o *BTInstanceControlNode750) HasImportMicroversion() bool {
	if o != nil && o.ImportMicroversion != nil {
		return true
	}

	return false
}

// SetImportMicroversion gets a reference to the given string and assigns it to the ImportMicroversion field.
func (o *BTInstanceControlNode750) SetImportMicroversion(v string) {
	o.ImportMicroversion = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *BTInstanceControlNode750) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTInstanceControlNode750) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *BTInstanceControlNode750) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *BTInstanceControlNode750) SetNodeId(v string) {
	o.NodeId = &v
}

// GetSuppressed returns the Suppressed field value if set, zero value otherwise.
func (o *BTInstanceControlNode750) GetSuppressed() bool {
	if o == nil || o.Suppressed == nil {
		var ret bool
		return ret
	}
	return *o.Suppressed
}

// GetSuppressedOk returns a tuple with the Suppressed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTInstanceControlNode750) GetSuppressedOk() (*bool, bool) {
	if o == nil || o.Suppressed == nil {
		return nil, false
	}
	return o.Suppressed, true
}

// HasSuppressed returns a boolean if a field has been set.
func (o *BTInstanceControlNode750) HasSuppressed() bool {
	if o != nil && o.Suppressed != nil {
		return true
	}

	return false
}

// SetSuppressed gets a reference to the given bool and assigns it to the Suppressed field.
func (o *BTInstanceControlNode750) SetSuppressed(v bool) {
	o.Suppressed = &v
}

// GetSuppressedFieldIndex returns the SuppressedFieldIndex field value if set, zero value otherwise.
func (o *BTInstanceControlNode750) GetSuppressedFieldIndex() int32 {
	if o == nil || o.SuppressedFieldIndex == nil {
		var ret int32
		return ret
	}
	return *o.SuppressedFieldIndex
}

// GetSuppressedFieldIndexOk returns a tuple with the SuppressedFieldIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTInstanceControlNode750) GetSuppressedFieldIndexOk() (*int32, bool) {
	if o == nil || o.SuppressedFieldIndex == nil {
		return nil, false
	}
	return o.SuppressedFieldIndex, true
}

// HasSuppressedFieldIndex returns a boolean if a field has been set.
func (o *BTInstanceControlNode750) HasSuppressedFieldIndex() bool {
	if o != nil && o.SuppressedFieldIndex != nil {
		return true
	}

	return false
}

// SetSuppressedFieldIndex gets a reference to the given int32 and assigns it to the SuppressedFieldIndex field.
func (o *BTInstanceControlNode750) SetSuppressedFieldIndex(v int32) {
	o.SuppressedFieldIndex = &v
}

// GetSuppressionConfigured returns the SuppressionConfigured field value if set, zero value otherwise.
func (o *BTInstanceControlNode750) GetSuppressionConfigured() bool {
	if o == nil || o.SuppressionConfigured == nil {
		var ret bool
		return ret
	}
	return *o.SuppressionConfigured
}

// GetSuppressionConfiguredOk returns a tuple with the SuppressionConfigured field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTInstanceControlNode750) GetSuppressionConfiguredOk() (*bool, bool) {
	if o == nil || o.SuppressionConfigured == nil {
		return nil, false
	}
	return o.SuppressionConfigured, true
}

// HasSuppressionConfigured returns a boolean if a field has been set.
func (o *BTInstanceControlNode750) HasSuppressionConfigured() bool {
	if o != nil && o.SuppressionConfigured != nil {
		return true
	}

	return false
}

// SetSuppressionConfigured gets a reference to the given bool and assigns it to the SuppressionConfigured field.
func (o *BTInstanceControlNode750) SetSuppressionConfigured(v bool) {
	o.SuppressionConfigured = &v
}

// GetSuppressionState returns the SuppressionState field value if set, zero value otherwise.
func (o *BTInstanceControlNode750) GetSuppressionState() BTMSuppressionState1924 {
	if o == nil || o.SuppressionState == nil {
		var ret BTMSuppressionState1924
		return ret
	}
	return *o.SuppressionState
}

// GetSuppressionStateOk returns a tuple with the SuppressionState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTInstanceControlNode750) GetSuppressionStateOk() (*BTMSuppressionState1924, bool) {
	if o == nil || o.SuppressionState == nil {
		return nil, false
	}
	return o.SuppressionState, true
}

// HasSuppressionState returns a boolean if a field has been set.
func (o *BTInstanceControlNode750) HasSuppressionState() bool {
	if o != nil && o.SuppressionState != nil {
		return true
	}

	return false
}

// SetSuppressionState gets a reference to the given BTMSuppressionState1924 and assigns it to the SuppressionState field.
func (o *BTInstanceControlNode750) SetSuppressionState(v BTMSuppressionState1924) {
	o.SuppressionState = &v
}

func (o BTInstanceControlNode750) MarshalJSON() ([]byte, error) {
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
	if o.Suppressed != nil {
		toSerialize["suppressed"] = o.Suppressed
	}
	if o.SuppressedFieldIndex != nil {
		toSerialize["suppressedFieldIndex"] = o.SuppressedFieldIndex
	}
	if o.SuppressionConfigured != nil {
		toSerialize["suppressionConfigured"] = o.SuppressionConfigured
	}
	if o.SuppressionState != nil {
		toSerialize["suppressionState"] = o.SuppressionState
	}
	return json.Marshal(toSerialize)
}

type NullableBTInstanceControlNode750 struct {
	value *BTInstanceControlNode750
	isSet bool
}

func (v NullableBTInstanceControlNode750) Get() *BTInstanceControlNode750 {
	return v.value
}

func (v *NullableBTInstanceControlNode750) Set(val *BTInstanceControlNode750) {
	v.value = val
	v.isSet = true
}

func (v NullableBTInstanceControlNode750) IsSet() bool {
	return v.isSet
}

func (v *NullableBTInstanceControlNode750) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTInstanceControlNode750(val *BTInstanceControlNode750) *NullableBTInstanceControlNode750 {
	return &NullableBTInstanceControlNode750{value: val, isSet: true}
}

func (v NullableBTInstanceControlNode750) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTInstanceControlNode750) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
