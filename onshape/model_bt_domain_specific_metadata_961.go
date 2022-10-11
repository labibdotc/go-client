/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.154.6859-c8a2bd7d8338
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
	"fmt"
)

// BTDomainSpecificMetadata961 - struct for BTDomainSpecificMetadata961
type BTDomainSpecificMetadata961 struct {
	implBTDomainSpecificMetadata961 interface{}
}

// BTSMSpecificMetadata1315AsBTDomainSpecificMetadata961 is a convenience function that returns BTSMSpecificMetadata1315 wrapped in BTDomainSpecificMetadata961
func (o *BTSMSpecificMetadata1315) AsBTDomainSpecificMetadata961() *BTDomainSpecificMetadata961 {
	return &BTDomainSpecificMetadata961{o}
}

// NewBTDomainSpecificMetadata961 instantiates a new BTDomainSpecificMetadata961 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTDomainSpecificMetadata961() *BTDomainSpecificMetadata961 {
	this := BTDomainSpecificMetadata961{Newbase_BTDomainSpecificMetadata961()}
	return &this
}

// NewBTDomainSpecificMetadata961WithDefaults instantiates a new BTDomainSpecificMetadata961 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTDomainSpecificMetadata961WithDefaults() *BTDomainSpecificMetadata961 {
	this := BTDomainSpecificMetadata961{Newbase_BTDomainSpecificMetadata961WithDefaults()}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTDomainSpecificMetadata961) GetBtType() string {
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
func (o *BTDomainSpecificMetadata961) GetBtTypeOk() (*string, bool) {
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
func (o *BTDomainSpecificMetadata961) HasBtType() bool {
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
func (o *BTDomainSpecificMetadata961) SetBtType(v string) {
	type getResult interface {
		SetBtType(v string)
	}

	o.GetActualInstance().(getResult).SetBtType(v)
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *BTDomainSpecificMetadata961) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'BTSMSpecificMetadata-1315'
	if jsonDict["btType"] == "BTSMSpecificMetadata-1315" {
		// try to unmarshal JSON data into BTSMSpecificMetadata1315
		var qr *BTSMSpecificMetadata1315
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTDomainSpecificMetadata961 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTDomainSpecificMetadata961 = nil
			return fmt.Errorf("Failed to unmarshal BTDomainSpecificMetadata961 as BTSMSpecificMetadata1315: %s", err.Error())
		}
	}

	var qtx *base_BTDomainSpecificMetadata961
	err = json.Unmarshal(data, &qtx)
	if err == nil {
		dst.implBTDomainSpecificMetadata961 = qtx
		return nil // data stored in dst.base_BTDomainSpecificMetadata961, return on the first match
	} else {
		dst.implBTDomainSpecificMetadata961 = nil
		return fmt.Errorf("Failed to unmarshal BTDomainSpecificMetadata961 as base_BTDomainSpecificMetadata961: %s", err.Error())
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src BTDomainSpecificMetadata961) MarshalJSON() ([]byte, error) {
	ret := src.GetActualInstance()
	if ret == nil {
		return nil, nil // no data in oneOf schemas
	} else {
		return json.Marshal(&ret)
	}
}

// Get the actual instance
func (obj *BTDomainSpecificMetadata961) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	return obj.implBTDomainSpecificMetadata961
}

type NullableBTDomainSpecificMetadata961 struct {
	value *BTDomainSpecificMetadata961
	isSet bool
}

func (v NullableBTDomainSpecificMetadata961) Get() *BTDomainSpecificMetadata961 {
	return v.value
}

func (v *NullableBTDomainSpecificMetadata961) Set(val *BTDomainSpecificMetadata961) {
	v.value = val
	v.isSet = true
}

func (v NullableBTDomainSpecificMetadata961) IsSet() bool {
	return v.isSet
}

func (v *NullableBTDomainSpecificMetadata961) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTDomainSpecificMetadata961(val *BTDomainSpecificMetadata961) *NullableBTDomainSpecificMetadata961 {
	return &NullableBTDomainSpecificMetadata961{value: val, isSet: true}
}

func (v NullableBTDomainSpecificMetadata961) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTDomainSpecificMetadata961) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

type base_BTDomainSpecificMetadata961 struct {
	BtType *string `json:"btType,omitempty"`
}

// Newbase_BTDomainSpecificMetadata961 instantiates a new base_BTDomainSpecificMetadata961 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func Newbase_BTDomainSpecificMetadata961() *base_BTDomainSpecificMetadata961 {
	this := base_BTDomainSpecificMetadata961{}
	return &this
}

// Newbase_BTDomainSpecificMetadata961WithDefaults instantiates a new base_BTDomainSpecificMetadata961 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func Newbase_BTDomainSpecificMetadata961WithDefaults() *base_BTDomainSpecificMetadata961 {
	this := base_BTDomainSpecificMetadata961{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *base_BTDomainSpecificMetadata961) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTDomainSpecificMetadata961) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *base_BTDomainSpecificMetadata961) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *base_BTDomainSpecificMetadata961) SetBtType(v string) {
	o.BtType = &v
}

func (o base_BTDomainSpecificMetadata961) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	return json.Marshal(toSerialize)
}
