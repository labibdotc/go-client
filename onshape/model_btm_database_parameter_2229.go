/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.152.6159-a7ef074944fe
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTMDatabaseParameter2229 struct for BTMDatabaseParameter2229
type BTMDatabaseParameter2229 struct {
	ImportMicroversion *string `json:"importMicroversion,omitempty"`
	NodeId             *string `json:"nodeId,omitempty"`
	ParameterId        *string `json:"parameterId,omitempty"`
	BtType             *string `json:"btType,omitempty"`
	DatabaseId         *string `json:"databaseId,omitempty"`
}

// NewBTMDatabaseParameter2229 instantiates a new BTMDatabaseParameter2229 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTMDatabaseParameter2229() *BTMDatabaseParameter2229 {
	this := BTMDatabaseParameter2229{}
	return &this
}

// NewBTMDatabaseParameter2229WithDefaults instantiates a new BTMDatabaseParameter2229 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTMDatabaseParameter2229WithDefaults() *BTMDatabaseParameter2229 {
	this := BTMDatabaseParameter2229{}
	return &this
}

// GetImportMicroversion returns the ImportMicroversion field value if set, zero value otherwise.
func (o *BTMDatabaseParameter2229) GetImportMicroversion() string {
	if o == nil || o.ImportMicroversion == nil {
		var ret string
		return ret
	}
	return *o.ImportMicroversion
}

// GetImportMicroversionOk returns a tuple with the ImportMicroversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMDatabaseParameter2229) GetImportMicroversionOk() (*string, bool) {
	if o == nil || o.ImportMicroversion == nil {
		return nil, false
	}
	return o.ImportMicroversion, true
}

// HasImportMicroversion returns a boolean if a field has been set.
func (o *BTMDatabaseParameter2229) HasImportMicroversion() bool {
	if o != nil && o.ImportMicroversion != nil {
		return true
	}

	return false
}

// SetImportMicroversion gets a reference to the given string and assigns it to the ImportMicroversion field.
func (o *BTMDatabaseParameter2229) SetImportMicroversion(v string) {
	o.ImportMicroversion = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *BTMDatabaseParameter2229) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMDatabaseParameter2229) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *BTMDatabaseParameter2229) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *BTMDatabaseParameter2229) SetNodeId(v string) {
	o.NodeId = &v
}

// GetParameterId returns the ParameterId field value if set, zero value otherwise.
func (o *BTMDatabaseParameter2229) GetParameterId() string {
	if o == nil || o.ParameterId == nil {
		var ret string
		return ret
	}
	return *o.ParameterId
}

// GetParameterIdOk returns a tuple with the ParameterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMDatabaseParameter2229) GetParameterIdOk() (*string, bool) {
	if o == nil || o.ParameterId == nil {
		return nil, false
	}
	return o.ParameterId, true
}

// HasParameterId returns a boolean if a field has been set.
func (o *BTMDatabaseParameter2229) HasParameterId() bool {
	if o != nil && o.ParameterId != nil {
		return true
	}

	return false
}

// SetParameterId gets a reference to the given string and assigns it to the ParameterId field.
func (o *BTMDatabaseParameter2229) SetParameterId(v string) {
	o.ParameterId = &v
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTMDatabaseParameter2229) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMDatabaseParameter2229) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTMDatabaseParameter2229) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTMDatabaseParameter2229) SetBtType(v string) {
	o.BtType = &v
}

// GetDatabaseId returns the DatabaseId field value if set, zero value otherwise.
func (o *BTMDatabaseParameter2229) GetDatabaseId() string {
	if o == nil || o.DatabaseId == nil {
		var ret string
		return ret
	}
	return *o.DatabaseId
}

// GetDatabaseIdOk returns a tuple with the DatabaseId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMDatabaseParameter2229) GetDatabaseIdOk() (*string, bool) {
	if o == nil || o.DatabaseId == nil {
		return nil, false
	}
	return o.DatabaseId, true
}

// HasDatabaseId returns a boolean if a field has been set.
func (o *BTMDatabaseParameter2229) HasDatabaseId() bool {
	if o != nil && o.DatabaseId != nil {
		return true
	}

	return false
}

// SetDatabaseId gets a reference to the given string and assigns it to the DatabaseId field.
func (o *BTMDatabaseParameter2229) SetDatabaseId(v string) {
	o.DatabaseId = &v
}

func (o BTMDatabaseParameter2229) MarshalJSON() ([]byte, error) {
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
	if o.DatabaseId != nil {
		toSerialize["databaseId"] = o.DatabaseId
	}
	return json.Marshal(toSerialize)
}

type NullableBTMDatabaseParameter2229 struct {
	value *BTMDatabaseParameter2229
	isSet bool
}

func (v NullableBTMDatabaseParameter2229) Get() *BTMDatabaseParameter2229 {
	return v.value
}

func (v *NullableBTMDatabaseParameter2229) Set(val *BTMDatabaseParameter2229) {
	v.value = val
	v.isSet = true
}

func (v NullableBTMDatabaseParameter2229) IsSet() bool {
	return v.isSet
}

func (v *NullableBTMDatabaseParameter2229) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTMDatabaseParameter2229(val *BTMDatabaseParameter2229) *NullableBTMDatabaseParameter2229 {
	return &NullableBTMDatabaseParameter2229{value: val, isSet: true}
}

func (v NullableBTMDatabaseParameter2229) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTMDatabaseParameter2229) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
