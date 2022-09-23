/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.154.6590-f8226b4e1789
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
	"fmt"
)

// BTBasePartColorCycle2614 - struct for BTBasePartColorCycle2614
type BTBasePartColorCycle2614 struct {
	implBTBasePartColorCycle2614 interface{}
}

// BTSystemPartColorCycle1580AsBTBasePartColorCycle2614 is a convenience function that returns BTSystemPartColorCycle1580 wrapped in BTBasePartColorCycle2614
func (o *BTSystemPartColorCycle1580) AsBTBasePartColorCycle2614() *BTBasePartColorCycle2614 {
	return &BTBasePartColorCycle2614{o}
}

// NewBTBasePartColorCycle2614 instantiates a new BTBasePartColorCycle2614 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTBasePartColorCycle2614() *BTBasePartColorCycle2614 {
	this := BTBasePartColorCycle2614{Newbase_BTBasePartColorCycle2614()}
	return &this
}

// NewBTBasePartColorCycle2614WithDefaults instantiates a new BTBasePartColorCycle2614 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTBasePartColorCycle2614WithDefaults() *BTBasePartColorCycle2614 {
	this := BTBasePartColorCycle2614{Newbase_BTBasePartColorCycle2614WithDefaults()}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTBasePartColorCycle2614) GetBtType() string {
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
func (o *BTBasePartColorCycle2614) GetBtTypeOk() (*string, bool) {
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
func (o *BTBasePartColorCycle2614) HasBtType() bool {
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
func (o *BTBasePartColorCycle2614) SetBtType(v string) {
	type getResult interface {
		SetBtType(v string)
	}

	o.GetActualInstance().(getResult).SetBtType(v)
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *BTBasePartColorCycle2614) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'BTSystemPartColorCycle-1580'
	if jsonDict["btType"] == "BTSystemPartColorCycle-1580" {
		// try to unmarshal JSON data into BTSystemPartColorCycle1580
		var qr *BTSystemPartColorCycle1580
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTBasePartColorCycle2614 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTBasePartColorCycle2614 = nil
			return fmt.Errorf("Failed to unmarshal BTBasePartColorCycle2614 as BTSystemPartColorCycle1580: %s", err.Error())
		}
	}

	var qtx *base_BTBasePartColorCycle2614
	err = json.Unmarshal(data, &qtx)
	if err == nil {
		dst.implBTBasePartColorCycle2614 = qtx
		return nil // data stored in dst.base_BTBasePartColorCycle2614, return on the first match
	} else {
		dst.implBTBasePartColorCycle2614 = nil
		return fmt.Errorf("Failed to unmarshal BTBasePartColorCycle2614 as base_BTBasePartColorCycle2614: %s", err.Error())
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src BTBasePartColorCycle2614) MarshalJSON() ([]byte, error) {
	ret := src.GetActualInstance()
	if ret == nil {
		return nil, nil // no data in oneOf schemas
	} else {
		return json.Marshal(&ret)
	}
}

// Get the actual instance
func (obj *BTBasePartColorCycle2614) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	return obj.implBTBasePartColorCycle2614
}

type NullableBTBasePartColorCycle2614 struct {
	value *BTBasePartColorCycle2614
	isSet bool
}

func (v NullableBTBasePartColorCycle2614) Get() *BTBasePartColorCycle2614 {
	return v.value
}

func (v *NullableBTBasePartColorCycle2614) Set(val *BTBasePartColorCycle2614) {
	v.value = val
	v.isSet = true
}

func (v NullableBTBasePartColorCycle2614) IsSet() bool {
	return v.isSet
}

func (v *NullableBTBasePartColorCycle2614) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTBasePartColorCycle2614(val *BTBasePartColorCycle2614) *NullableBTBasePartColorCycle2614 {
	return &NullableBTBasePartColorCycle2614{value: val, isSet: true}
}

func (v NullableBTBasePartColorCycle2614) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTBasePartColorCycle2614) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

type base_BTBasePartColorCycle2614 struct {
	BtType *string `json:"btType,omitempty"`
}

// Newbase_BTBasePartColorCycle2614 instantiates a new base_BTBasePartColorCycle2614 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func Newbase_BTBasePartColorCycle2614() *base_BTBasePartColorCycle2614 {
	this := base_BTBasePartColorCycle2614{}
	return &this
}

// Newbase_BTBasePartColorCycle2614WithDefaults instantiates a new base_BTBasePartColorCycle2614 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func Newbase_BTBasePartColorCycle2614WithDefaults() *base_BTBasePartColorCycle2614 {
	this := base_BTBasePartColorCycle2614{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *base_BTBasePartColorCycle2614) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTBasePartColorCycle2614) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *base_BTBasePartColorCycle2614) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *base_BTBasePartColorCycle2614) SetBtType(v string) {
	o.BtType = &v
}

func (o base_BTBasePartColorCycle2614) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	return json.Marshal(toSerialize)
}
