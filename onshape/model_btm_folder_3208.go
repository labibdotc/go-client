/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.168.20965-e01b1f4d96d1
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTMFolder3208 struct for BTMFolder3208
type BTMFolder3208 struct {
	BtType             *string `json:"btType,omitempty"`
	ImportMicroversion *string `json:"importMicroversion,omitempty"`
	NodeId             *string `json:"nodeId,omitempty"`
	FolderId           *string `json:"folderId,omitempty"`
	IsStartFolder      *bool   `json:"isStartFolder,omitempty"`
	Name               *string `json:"name,omitempty"`
}

// NewBTMFolder3208 instantiates a new BTMFolder3208 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTMFolder3208() *BTMFolder3208 {
	this := BTMFolder3208{}
	return &this
}

// NewBTMFolder3208WithDefaults instantiates a new BTMFolder3208 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTMFolder3208WithDefaults() *BTMFolder3208 {
	this := BTMFolder3208{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTMFolder3208) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMFolder3208) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTMFolder3208) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTMFolder3208) SetBtType(v string) {
	o.BtType = &v
}

// GetImportMicroversion returns the ImportMicroversion field value if set, zero value otherwise.
func (o *BTMFolder3208) GetImportMicroversion() string {
	if o == nil || o.ImportMicroversion == nil {
		var ret string
		return ret
	}
	return *o.ImportMicroversion
}

// GetImportMicroversionOk returns a tuple with the ImportMicroversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMFolder3208) GetImportMicroversionOk() (*string, bool) {
	if o == nil || o.ImportMicroversion == nil {
		return nil, false
	}
	return o.ImportMicroversion, true
}

// HasImportMicroversion returns a boolean if a field has been set.
func (o *BTMFolder3208) HasImportMicroversion() bool {
	if o != nil && o.ImportMicroversion != nil {
		return true
	}

	return false
}

// SetImportMicroversion gets a reference to the given string and assigns it to the ImportMicroversion field.
func (o *BTMFolder3208) SetImportMicroversion(v string) {
	o.ImportMicroversion = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *BTMFolder3208) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMFolder3208) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *BTMFolder3208) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *BTMFolder3208) SetNodeId(v string) {
	o.NodeId = &v
}

// GetFolderId returns the FolderId field value if set, zero value otherwise.
func (o *BTMFolder3208) GetFolderId() string {
	if o == nil || o.FolderId == nil {
		var ret string
		return ret
	}
	return *o.FolderId
}

// GetFolderIdOk returns a tuple with the FolderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMFolder3208) GetFolderIdOk() (*string, bool) {
	if o == nil || o.FolderId == nil {
		return nil, false
	}
	return o.FolderId, true
}

// HasFolderId returns a boolean if a field has been set.
func (o *BTMFolder3208) HasFolderId() bool {
	if o != nil && o.FolderId != nil {
		return true
	}

	return false
}

// SetFolderId gets a reference to the given string and assigns it to the FolderId field.
func (o *BTMFolder3208) SetFolderId(v string) {
	o.FolderId = &v
}

// GetIsStartFolder returns the IsStartFolder field value if set, zero value otherwise.
func (o *BTMFolder3208) GetIsStartFolder() bool {
	if o == nil || o.IsStartFolder == nil {
		var ret bool
		return ret
	}
	return *o.IsStartFolder
}

// GetIsStartFolderOk returns a tuple with the IsStartFolder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMFolder3208) GetIsStartFolderOk() (*bool, bool) {
	if o == nil || o.IsStartFolder == nil {
		return nil, false
	}
	return o.IsStartFolder, true
}

// HasIsStartFolder returns a boolean if a field has been set.
func (o *BTMFolder3208) HasIsStartFolder() bool {
	if o != nil && o.IsStartFolder != nil {
		return true
	}

	return false
}

// SetIsStartFolder gets a reference to the given bool and assigns it to the IsStartFolder field.
func (o *BTMFolder3208) SetIsStartFolder(v bool) {
	o.IsStartFolder = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BTMFolder3208) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMFolder3208) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BTMFolder3208) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BTMFolder3208) SetName(v string) {
	o.Name = &v
}

func (o BTMFolder3208) MarshalJSON() ([]byte, error) {
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
	if o.FolderId != nil {
		toSerialize["folderId"] = o.FolderId
	}
	if o.IsStartFolder != nil {
		toSerialize["isStartFolder"] = o.IsStartFolder
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableBTMFolder3208 struct {
	value *BTMFolder3208
	isSet bool
}

func (v NullableBTMFolder3208) Get() *BTMFolder3208 {
	return v.value
}

func (v *NullableBTMFolder3208) Set(val *BTMFolder3208) {
	v.value = val
	v.isSet = true
}

func (v NullableBTMFolder3208) IsSet() bool {
	return v.isSet
}

func (v *NullableBTMFolder3208) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTMFolder3208(val *BTMFolder3208) *NullableBTMFolder3208 {
	return &NullableBTMFolder3208{value: val, isSet: true}
}

func (v NullableBTMFolder3208) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTMFolder3208) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
