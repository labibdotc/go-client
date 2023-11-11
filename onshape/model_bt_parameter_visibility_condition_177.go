/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.172.25784-25d5580b0721
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
	"fmt"
)

// BTParameterVisibilityCondition177 - struct for BTParameterVisibilityCondition177
type BTParameterVisibilityCondition177 struct {
	implBTParameterVisibilityCondition177 interface{}
}

// BTParameterVisibilityOnMateDOFType2114AsBTParameterVisibilityCondition177 is a convenience function that returns BTParameterVisibilityOnMateDOFType2114 wrapped in BTParameterVisibilityCondition177
func (o *BTParameterVisibilityOnMateDOFType2114) AsBTParameterVisibilityCondition177() *BTParameterVisibilityCondition177 {
	return &BTParameterVisibilityCondition177{o}
}

// BTParameterVisibilityLogical178AsBTParameterVisibilityCondition177 is a convenience function that returns BTParameterVisibilityLogical178 wrapped in BTParameterVisibilityCondition177
func (o *BTParameterVisibilityLogical178) AsBTParameterVisibilityCondition177() *BTParameterVisibilityCondition177 {
	return &BTParameterVisibilityCondition177{o}
}

// BTParameterVisibilityAlwaysHidden176AsBTParameterVisibilityCondition177 is a convenience function that returns BTParameterVisibilityAlwaysHidden176 wrapped in BTParameterVisibilityCondition177
func (o *BTParameterVisibilityAlwaysHidden176) AsBTParameterVisibilityCondition177() *BTParameterVisibilityCondition177 {
	return &BTParameterVisibilityCondition177{o}
}

// BTParameterVisibilityOnEqual180AsBTParameterVisibilityCondition177 is a convenience function that returns BTParameterVisibilityOnEqual180 wrapped in BTParameterVisibilityCondition177
func (o *BTParameterVisibilityOnEqual180) AsBTParameterVisibilityCondition177() *BTParameterVisibilityCondition177 {
	return &BTParameterVisibilityCondition177{o}
}

// NewBTParameterVisibilityCondition177 instantiates a new BTParameterVisibilityCondition177 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTParameterVisibilityCondition177() *BTParameterVisibilityCondition177 {
	this := BTParameterVisibilityCondition177{Newbase_BTParameterVisibilityCondition177()}
	return &this
}

// NewBTParameterVisibilityCondition177WithDefaults instantiates a new BTParameterVisibilityCondition177 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTParameterVisibilityCondition177WithDefaults() *BTParameterVisibilityCondition177 {
	this := BTParameterVisibilityCondition177{Newbase_BTParameterVisibilityCondition177WithDefaults()}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTParameterVisibilityCondition177) GetBtType() string {
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
func (o *BTParameterVisibilityCondition177) GetBtTypeOk() (*string, bool) {
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
func (o *BTParameterVisibilityCondition177) HasBtType() bool {
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
func (o *BTParameterVisibilityCondition177) SetBtType(v string) {
	type getResult interface {
		SetBtType(v string)
	}

	o.GetActualInstance().(getResult).SetBtType(v)
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *BTParameterVisibilityCondition177) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'BTParameterVisibilityAlwaysHidden-176'
	if jsonDict["btType"] == "BTParameterVisibilityAlwaysHidden-176" {
		// try to unmarshal JSON data into BTParameterVisibilityAlwaysHidden176
		var qr *BTParameterVisibilityAlwaysHidden176
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTParameterVisibilityCondition177 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTParameterVisibilityCondition177 = nil
			return fmt.Errorf("Failed to unmarshal BTParameterVisibilityCondition177 as BTParameterVisibilityAlwaysHidden176: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTParameterVisibilityLogical-178'
	if jsonDict["btType"] == "BTParameterVisibilityLogical-178" {
		// try to unmarshal JSON data into BTParameterVisibilityLogical178
		var qr *BTParameterVisibilityLogical178
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTParameterVisibilityCondition177 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTParameterVisibilityCondition177 = nil
			return fmt.Errorf("Failed to unmarshal BTParameterVisibilityCondition177 as BTParameterVisibilityLogical178: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTParameterVisibilityOnEqual-180'
	if jsonDict["btType"] == "BTParameterVisibilityOnEqual-180" {
		// try to unmarshal JSON data into BTParameterVisibilityOnEqual180
		var qr *BTParameterVisibilityOnEqual180
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTParameterVisibilityCondition177 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTParameterVisibilityCondition177 = nil
			return fmt.Errorf("Failed to unmarshal BTParameterVisibilityCondition177 as BTParameterVisibilityOnEqual180: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTParameterVisibilityOnMateDOFType-2114'
	if jsonDict["btType"] == "BTParameterVisibilityOnMateDOFType-2114" {
		// try to unmarshal JSON data into BTParameterVisibilityOnMateDOFType2114
		var qr *BTParameterVisibilityOnMateDOFType2114
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTParameterVisibilityCondition177 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTParameterVisibilityCondition177 = nil
			return fmt.Errorf("Failed to unmarshal BTParameterVisibilityCondition177 as BTParameterVisibilityOnMateDOFType2114: %s", err.Error())
		}
	}

	var qtx *base_BTParameterVisibilityCondition177
	err = json.Unmarshal(data, &qtx)
	if err == nil {
		dst.implBTParameterVisibilityCondition177 = qtx
		return nil // data stored in dst.base_BTParameterVisibilityCondition177, return on the first match
	} else {
		dst.implBTParameterVisibilityCondition177 = nil
		return fmt.Errorf("Failed to unmarshal BTParameterVisibilityCondition177 as base_BTParameterVisibilityCondition177: %s", err.Error())
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src BTParameterVisibilityCondition177) MarshalJSON() ([]byte, error) {
	ret := src.GetActualInstance()
	if ret == nil {
		return nil, nil // no data in oneOf schemas
	} else {
		return json.Marshal(&ret)
	}
}

// Get the actual instance
func (obj *BTParameterVisibilityCondition177) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	return obj.implBTParameterVisibilityCondition177
}

type NullableBTParameterVisibilityCondition177 struct {
	value *BTParameterVisibilityCondition177
	isSet bool
}

func (v NullableBTParameterVisibilityCondition177) Get() *BTParameterVisibilityCondition177 {
	return v.value
}

func (v *NullableBTParameterVisibilityCondition177) Set(val *BTParameterVisibilityCondition177) {
	v.value = val
	v.isSet = true
}

func (v NullableBTParameterVisibilityCondition177) IsSet() bool {
	return v.isSet
}

func (v *NullableBTParameterVisibilityCondition177) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTParameterVisibilityCondition177(val *BTParameterVisibilityCondition177) *NullableBTParameterVisibilityCondition177 {
	return &NullableBTParameterVisibilityCondition177{value: val, isSet: true}
}

func (v NullableBTParameterVisibilityCondition177) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTParameterVisibilityCondition177) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

type base_BTParameterVisibilityCondition177 struct {
	BtType *string `json:"btType,omitempty"`
}

// Newbase_BTParameterVisibilityCondition177 instantiates a new base_BTParameterVisibilityCondition177 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func Newbase_BTParameterVisibilityCondition177() *base_BTParameterVisibilityCondition177 {
	this := base_BTParameterVisibilityCondition177{}
	return &this
}

// Newbase_BTParameterVisibilityCondition177WithDefaults instantiates a new base_BTParameterVisibilityCondition177 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func Newbase_BTParameterVisibilityCondition177WithDefaults() *base_BTParameterVisibilityCondition177 {
	this := base_BTParameterVisibilityCondition177{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *base_BTParameterVisibilityCondition177) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTParameterVisibilityCondition177) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *base_BTParameterVisibilityCondition177) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *base_BTParameterVisibilityCondition177) SetBtType(v string) {
	o.BtType = &v
}

func (o base_BTParameterVisibilityCondition177) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	return json.Marshal(toSerialize)
}
