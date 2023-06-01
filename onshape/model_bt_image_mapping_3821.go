/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.164.16643-ef813b2da145
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
	"fmt"
)

// BTImageMapping3821 - struct for BTImageMapping3821
type BTImageMapping3821 struct {
	implBTImageMapping3821 interface{}
}

// BTCylindricalImageMapping1640AsBTImageMapping3821 is a convenience function that returns BTCylindricalImageMapping1640 wrapped in BTImageMapping3821
func (o *BTCylindricalImageMapping1640) AsBTImageMapping3821() *BTImageMapping3821 {
	return &BTImageMapping3821{o}
}

// BTPlanarImageMapping4398AsBTImageMapping3821 is a convenience function that returns BTPlanarImageMapping4398 wrapped in BTImageMapping3821
func (o *BTPlanarImageMapping4398) AsBTImageMapping3821() *BTImageMapping3821 {
	return &BTImageMapping3821{o}
}

// NewBTImageMapping3821 instantiates a new BTImageMapping3821 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTImageMapping3821() *BTImageMapping3821 {
	this := BTImageMapping3821{Newbase_BTImageMapping3821()}
	return &this
}

// NewBTImageMapping3821WithDefaults instantiates a new BTImageMapping3821 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTImageMapping3821WithDefaults() *BTImageMapping3821 {
	this := BTImageMapping3821{Newbase_BTImageMapping3821WithDefaults()}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTImageMapping3821) GetBtType() string {
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
func (o *BTImageMapping3821) GetBtTypeOk() (*string, bool) {
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
func (o *BTImageMapping3821) HasBtType() bool {
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
func (o *BTImageMapping3821) SetBtType(v string) {
	type getResult interface {
		SetBtType(v string)
	}

	o.GetActualInstance().(getResult).SetBtType(v)
}

// GetDeterministicIds returns the DeterministicIds field value if set, zero value otherwise.
func (o *BTImageMapping3821) GetDeterministicIds() []string {
	type getResult interface {
		GetDeterministicIds() []string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetDeterministicIds()
	} else {
		var de []string
		return de
	}
}

// GetDeterministicIdsOk returns a tuple with the DeterministicIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTImageMapping3821) GetDeterministicIdsOk() ([]string, bool) {
	type getResult interface {
		GetDeterministicIdsOk() ([]string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetDeterministicIdsOk()
	} else {
		return nil, false
	}
}

// HasDeterministicIds returns a boolean if a field has been set.
func (o *BTImageMapping3821) HasDeterministicIds() bool {
	type getResult interface {
		HasDeterministicIds() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasDeterministicIds()
	} else {
		return false
	}
}

// SetDeterministicIds gets a reference to the given []string and assigns it to the DeterministicIds field.
func (o *BTImageMapping3821) SetDeterministicIds(v []string) {
	type getResult interface {
		SetDeterministicIds(v []string)
	}

	o.GetActualInstance().(getResult).SetDeterministicIds(v)
}

// GetUvTransform returns the UvTransform field value if set, zero value otherwise.
func (o *BTImageMapping3821) GetUvTransform() BTMatrix3x3340 {
	type getResult interface {
		GetUvTransform() BTMatrix3x3340
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetUvTransform()
	} else {
		var de BTMatrix3x3340
		return de
	}
}

// GetUvTransformOk returns a tuple with the UvTransform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTImageMapping3821) GetUvTransformOk() (*BTMatrix3x3340, bool) {
	type getResult interface {
		GetUvTransformOk() (*BTMatrix3x3340, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetUvTransformOk()
	} else {
		return nil, false
	}
}

// HasUvTransform returns a boolean if a field has been set.
func (o *BTImageMapping3821) HasUvTransform() bool {
	type getResult interface {
		HasUvTransform() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasUvTransform()
	} else {
		return false
	}
}

// SetUvTransform gets a reference to the given BTMatrix3x3340 and assigns it to the UvTransform field.
func (o *BTImageMapping3821) SetUvTransform(v BTMatrix3x3340) {
	type getResult interface {
		SetUvTransform(v BTMatrix3x3340)
	}

	o.GetActualInstance().(getResult).SetUvTransform(v)
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *BTImageMapping3821) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'BTCylindricalImageMapping-1640'
	if jsonDict["btType"] == "BTCylindricalImageMapping-1640" {
		// try to unmarshal JSON data into BTCylindricalImageMapping1640
		var qr *BTCylindricalImageMapping1640
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTImageMapping3821 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTImageMapping3821 = nil
			return fmt.Errorf("Failed to unmarshal BTImageMapping3821 as BTCylindricalImageMapping1640: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTPlanarImageMapping-4398'
	if jsonDict["btType"] == "BTPlanarImageMapping-4398" {
		// try to unmarshal JSON data into BTPlanarImageMapping4398
		var qr *BTPlanarImageMapping4398
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTImageMapping3821 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTImageMapping3821 = nil
			return fmt.Errorf("Failed to unmarshal BTImageMapping3821 as BTPlanarImageMapping4398: %s", err.Error())
		}
	}

	var qtx *base_BTImageMapping3821
	err = json.Unmarshal(data, &qtx)
	if err == nil {
		dst.implBTImageMapping3821 = qtx
		return nil // data stored in dst.base_BTImageMapping3821, return on the first match
	} else {
		dst.implBTImageMapping3821 = nil
		return fmt.Errorf("Failed to unmarshal BTImageMapping3821 as base_BTImageMapping3821: %s", err.Error())
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src BTImageMapping3821) MarshalJSON() ([]byte, error) {
	ret := src.GetActualInstance()
	if ret == nil {
		return nil, nil // no data in oneOf schemas
	} else {
		return json.Marshal(&ret)
	}
}

// Get the actual instance
func (obj *BTImageMapping3821) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	return obj.implBTImageMapping3821
}

type NullableBTImageMapping3821 struct {
	value *BTImageMapping3821
	isSet bool
}

func (v NullableBTImageMapping3821) Get() *BTImageMapping3821 {
	return v.value
}

func (v *NullableBTImageMapping3821) Set(val *BTImageMapping3821) {
	v.value = val
	v.isSet = true
}

func (v NullableBTImageMapping3821) IsSet() bool {
	return v.isSet
}

func (v *NullableBTImageMapping3821) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTImageMapping3821(val *BTImageMapping3821) *NullableBTImageMapping3821 {
	return &NullableBTImageMapping3821{value: val, isSet: true}
}

func (v NullableBTImageMapping3821) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTImageMapping3821) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

type base_BTImageMapping3821 struct {
	BtType           *string         `json:"btType,omitempty"`
	DeterministicIds []string        `json:"deterministicIds,omitempty"`
	UvTransform      *BTMatrix3x3340 `json:"uvTransform,omitempty"`
}

// Newbase_BTImageMapping3821 instantiates a new base_BTImageMapping3821 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func Newbase_BTImageMapping3821() *base_BTImageMapping3821 {
	this := base_BTImageMapping3821{}
	return &this
}

// Newbase_BTImageMapping3821WithDefaults instantiates a new base_BTImageMapping3821 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func Newbase_BTImageMapping3821WithDefaults() *base_BTImageMapping3821 {
	this := base_BTImageMapping3821{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *base_BTImageMapping3821) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTImageMapping3821) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *base_BTImageMapping3821) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *base_BTImageMapping3821) SetBtType(v string) {
	o.BtType = &v
}

// GetDeterministicIds returns the DeterministicIds field value if set, zero value otherwise.
func (o *base_BTImageMapping3821) GetDeterministicIds() []string {
	if o == nil || o.DeterministicIds == nil {
		var ret []string
		return ret
	}
	return o.DeterministicIds
}

// GetDeterministicIdsOk returns a tuple with the DeterministicIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTImageMapping3821) GetDeterministicIdsOk() ([]string, bool) {
	if o == nil || o.DeterministicIds == nil {
		return nil, false
	}
	return o.DeterministicIds, true
}

// HasDeterministicIds returns a boolean if a field has been set.
func (o *base_BTImageMapping3821) HasDeterministicIds() bool {
	if o != nil && o.DeterministicIds != nil {
		return true
	}

	return false
}

// SetDeterministicIds gets a reference to the given []string and assigns it to the DeterministicIds field.
func (o *base_BTImageMapping3821) SetDeterministicIds(v []string) {
	o.DeterministicIds = v
}

// GetUvTransform returns the UvTransform field value if set, zero value otherwise.
func (o *base_BTImageMapping3821) GetUvTransform() BTMatrix3x3340 {
	if o == nil || o.UvTransform == nil {
		var ret BTMatrix3x3340
		return ret
	}
	return *o.UvTransform
}

// GetUvTransformOk returns a tuple with the UvTransform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTImageMapping3821) GetUvTransformOk() (*BTMatrix3x3340, bool) {
	if o == nil || o.UvTransform == nil {
		return nil, false
	}
	return o.UvTransform, true
}

// HasUvTransform returns a boolean if a field has been set.
func (o *base_BTImageMapping3821) HasUvTransform() bool {
	if o != nil && o.UvTransform != nil {
		return true
	}

	return false
}

// SetUvTransform gets a reference to the given BTMatrix3x3340 and assigns it to the UvTransform field.
func (o *base_BTImageMapping3821) SetUvTransform(v BTMatrix3x3340) {
	o.UvTransform = &v
}

func (o base_BTImageMapping3821) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.DeterministicIds != nil {
		toSerialize["deterministicIds"] = o.DeterministicIds
	}
	if o.UvTransform != nil {
		toSerialize["uvTransform"] = o.UvTransform
	}
	return json.Marshal(toSerialize)
}
