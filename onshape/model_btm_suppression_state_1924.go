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
	"fmt"
)

// BTMSuppressionState1924 - struct for BTMSuppressionState1924
type BTMSuppressionState1924 struct {
	implBTMSuppressionState1924 interface{}
}

// BTMSuppressionStateConfigured2598AsBTMSuppressionState1924 is a convenience function that returns BTMSuppressionStateConfigured2598 wrapped in BTMSuppressionState1924
func (o *BTMSuppressionStateConfigured2598) AsBTMSuppressionState1924() *BTMSuppressionState1924 {
	return &BTMSuppressionState1924{o}
}

// NewBTMSuppressionState1924 instantiates a new BTMSuppressionState1924 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTMSuppressionState1924() *BTMSuppressionState1924 {
	this := BTMSuppressionState1924{Newbase_BTMSuppressionState1924()}
	return &this
}

// NewBTMSuppressionState1924WithDefaults instantiates a new BTMSuppressionState1924 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTMSuppressionState1924WithDefaults() *BTMSuppressionState1924 {
	this := BTMSuppressionState1924{Newbase_BTMSuppressionState1924WithDefaults()}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTMSuppressionState1924) GetBtType() string {
	type getResult interface {
		GetBtType() string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetBtType()
	} else {
		var de string
		return de
	}
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMSuppressionState1924) GetBtTypeOk() (*string, bool) {
	type getResult interface {
		GetBtTypeOk() (*string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetBtTypeOk()
	} else {
		return nil, false
	}
}

// HasBtType returns a boolean if a field has been set.
func (o *BTMSuppressionState1924) HasBtType() bool {
	type getResult interface {
		HasBtType() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasBtType()
	} else {
		return false
	}
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTMSuppressionState1924) SetBtType(v string) {
	type getResult interface {
		SetBtType(v string)
	}

	o.GetActualInstance().(getResult).SetBtType(v)
}

// GetImportMicroversion returns the ImportMicroversion field value if set, zero value otherwise.
func (o *BTMSuppressionState1924) GetImportMicroversion() string {
	type getResult interface {
		GetImportMicroversion() string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetImportMicroversion()
	} else {
		var de string
		return de
	}
}

// GetImportMicroversionOk returns a tuple with the ImportMicroversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMSuppressionState1924) GetImportMicroversionOk() (*string, bool) {
	type getResult interface {
		GetImportMicroversionOk() (*string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetImportMicroversionOk()
	} else {
		return nil, false
	}
}

// HasImportMicroversion returns a boolean if a field has been set.
func (o *BTMSuppressionState1924) HasImportMicroversion() bool {
	type getResult interface {
		HasImportMicroversion() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasImportMicroversion()
	} else {
		return false
	}
}

// SetImportMicroversion gets a reference to the given string and assigns it to the ImportMicroversion field.
func (o *BTMSuppressionState1924) SetImportMicroversion(v string) {
	type getResult interface {
		SetImportMicroversion(v string)
	}

	o.GetActualInstance().(getResult).SetImportMicroversion(v)
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *BTMSuppressionState1924) GetNodeId() string {
	type getResult interface {
		GetNodeId() string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetNodeId()
	} else {
		var de string
		return de
	}
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMSuppressionState1924) GetNodeIdOk() (*string, bool) {
	type getResult interface {
		GetNodeIdOk() (*string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetNodeIdOk()
	} else {
		return nil, false
	}
}

// HasNodeId returns a boolean if a field has been set.
func (o *BTMSuppressionState1924) HasNodeId() bool {
	type getResult interface {
		HasNodeId() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasNodeId()
	} else {
		return false
	}
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *BTMSuppressionState1924) SetNodeId(v string) {
	type getResult interface {
		SetNodeId(v string)
	}

	o.GetActualInstance().(getResult).SetNodeId(v)
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *BTMSuppressionState1924) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'BTMSuppressionStateConfigured-2598'
	if jsonDict["btType"] == "BTMSuppressionStateConfigured-2598" {
		// try to unmarshal JSON data into BTMSuppressionStateConfigured2598
		var qr *BTMSuppressionStateConfigured2598
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTMSuppressionState1924 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTMSuppressionState1924 = nil
			return fmt.Errorf("Failed to unmarshal BTMSuppressionState1924 as BTMSuppressionStateConfigured2598: %s", err.Error())
		}
	}

	var qtx *base_BTMSuppressionState1924
	err = json.Unmarshal(data, &qtx)
	if err == nil {
		dst.implBTMSuppressionState1924 = qtx
		return nil // data stored in dst.base_BTMSuppressionState1924, return on the first match
	} else {
		dst.implBTMSuppressionState1924 = nil
		return fmt.Errorf("Failed to unmarshal BTMSuppressionState1924 as base_BTMSuppressionState1924: %s", err.Error())
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src BTMSuppressionState1924) MarshalJSON() ([]byte, error) {
	ret := src.GetActualInstance()
	if ret == nil {
		return nil, nil // no data in oneOf schemas
	} else {
		return json.Marshal(&ret)
	}
}

// Get the actual instance
func (obj *BTMSuppressionState1924) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	return obj.implBTMSuppressionState1924
}

type NullableBTMSuppressionState1924 struct {
	value *BTMSuppressionState1924
	isSet bool
}

func (v NullableBTMSuppressionState1924) Get() *BTMSuppressionState1924 {
	return v.value
}

func (v *NullableBTMSuppressionState1924) Set(val *BTMSuppressionState1924) {
	v.value = val
	v.isSet = true
}

func (v NullableBTMSuppressionState1924) IsSet() bool {
	return v.isSet
}

func (v *NullableBTMSuppressionState1924) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTMSuppressionState1924(val *BTMSuppressionState1924) *NullableBTMSuppressionState1924 {
	return &NullableBTMSuppressionState1924{value: val, isSet: true}
}

func (v NullableBTMSuppressionState1924) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTMSuppressionState1924) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

type base_BTMSuppressionState1924 struct {
	BtType             *string `json:"btType,omitempty"`
	ImportMicroversion *string `json:"importMicroversion,omitempty"`
	NodeId             *string `json:"nodeId,omitempty"`
}

// Newbase_BTMSuppressionState1924 instantiates a new base_BTMSuppressionState1924 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func Newbase_BTMSuppressionState1924() *base_BTMSuppressionState1924 {
	this := base_BTMSuppressionState1924{}
	return &this
}

// Newbase_BTMSuppressionState1924WithDefaults instantiates a new base_BTMSuppressionState1924 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func Newbase_BTMSuppressionState1924WithDefaults() *base_BTMSuppressionState1924 {
	this := base_BTMSuppressionState1924{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *base_BTMSuppressionState1924) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTMSuppressionState1924) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *base_BTMSuppressionState1924) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *base_BTMSuppressionState1924) SetBtType(v string) {
	o.BtType = &v
}

// GetImportMicroversion returns the ImportMicroversion field value if set, zero value otherwise.
func (o *base_BTMSuppressionState1924) GetImportMicroversion() string {
	if o == nil || o.ImportMicroversion == nil {
		var ret string
		return ret
	}
	return *o.ImportMicroversion
}

// GetImportMicroversionOk returns a tuple with the ImportMicroversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTMSuppressionState1924) GetImportMicroversionOk() (*string, bool) {
	if o == nil || o.ImportMicroversion == nil {
		return nil, false
	}
	return o.ImportMicroversion, true
}

// HasImportMicroversion returns a boolean if a field has been set.
func (o *base_BTMSuppressionState1924) HasImportMicroversion() bool {
	if o != nil && o.ImportMicroversion != nil {
		return true
	}

	return false
}

// SetImportMicroversion gets a reference to the given string and assigns it to the ImportMicroversion field.
func (o *base_BTMSuppressionState1924) SetImportMicroversion(v string) {
	o.ImportMicroversion = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *base_BTMSuppressionState1924) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTMSuppressionState1924) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *base_BTMSuppressionState1924) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *base_BTMSuppressionState1924) SetNodeId(v string) {
	o.NodeId = &v
}

func (o base_BTMSuppressionState1924) MarshalJSON() ([]byte, error) {
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
