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
	"fmt"
)

// BTSketchCompositeEntityDisplayData1093 - struct for BTSketchCompositeEntityDisplayData1093
type BTSketchCompositeEntityDisplayData1093 struct {
	implBTSketchCompositeEntityDisplayData1093 interface{}
}

// BTSketchTextDisplayData1707AsBTSketchCompositeEntityDisplayData1093 is a convenience function that returns BTSketchTextDisplayData1707 wrapped in BTSketchCompositeEntityDisplayData1093
func (o *BTSketchTextDisplayData1707) AsBTSketchCompositeEntityDisplayData1093() *BTSketchCompositeEntityDisplayData1093 {
	return &BTSketchCompositeEntityDisplayData1093{o}
}

// BTSketchImageDisplayData1379AsBTSketchCompositeEntityDisplayData1093 is a convenience function that returns BTSketchImageDisplayData1379 wrapped in BTSketchCompositeEntityDisplayData1093
func (o *BTSketchImageDisplayData1379) AsBTSketchCompositeEntityDisplayData1093() *BTSketchCompositeEntityDisplayData1093 {
	return &BTSketchCompositeEntityDisplayData1093{o}
}

// NewBTSketchCompositeEntityDisplayData1093 instantiates a new BTSketchCompositeEntityDisplayData1093 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTSketchCompositeEntityDisplayData1093() *BTSketchCompositeEntityDisplayData1093 {
	this := BTSketchCompositeEntityDisplayData1093{Newbase_BTSketchCompositeEntityDisplayData1093()}
	return &this
}

// NewBTSketchCompositeEntityDisplayData1093WithDefaults instantiates a new BTSketchCompositeEntityDisplayData1093 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTSketchCompositeEntityDisplayData1093WithDefaults() *BTSketchCompositeEntityDisplayData1093 {
	this := BTSketchCompositeEntityDisplayData1093{Newbase_BTSketchCompositeEntityDisplayData1093WithDefaults()}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTSketchCompositeEntityDisplayData1093) GetBtType() string {
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
func (o *BTSketchCompositeEntityDisplayData1093) GetBtTypeOk() (*string, bool) {
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
func (o *BTSketchCompositeEntityDisplayData1093) HasBtType() bool {
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
func (o *BTSketchCompositeEntityDisplayData1093) SetBtType(v string) {
	type getResult interface {
		SetBtType(v string)
	}

	o.GetActualInstance().(getResult).SetBtType(v)
}

// GetPoints returns the Points field value if set, zero value otherwise.
func (o *BTSketchCompositeEntityDisplayData1093) GetPoints() []float64 {
	type getResult interface {
		GetPoints() []float64
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetPoints()
	} else {
		var de []float64
		return de
	}
}

// GetPointsOk returns a tuple with the Points field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSketchCompositeEntityDisplayData1093) GetPointsOk() ([]float64, bool) {
	type getResult interface {
		GetPointsOk() ([]float64, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetPointsOk()
	} else {
		return nil, false
	}
}

// HasPoints returns a boolean if a field has been set.
func (o *BTSketchCompositeEntityDisplayData1093) HasPoints() bool {
	type getResult interface {
		HasPoints() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasPoints()
	} else {
		return false
	}
}

// SetPoints gets a reference to the given []float64 and assigns it to the Points field.
func (o *BTSketchCompositeEntityDisplayData1093) SetPoints(v []float64) {
	type getResult interface {
		SetPoints(v []float64)
	}

	o.GetActualInstance().(getResult).SetPoints(v)
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *BTSketchCompositeEntityDisplayData1093) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'BTSketchImageDisplayData-1379'
	if jsonDict["btType"] == "BTSketchImageDisplayData-1379" {
		// try to unmarshal JSON data into BTSketchImageDisplayData1379
		var qr *BTSketchImageDisplayData1379
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTSketchCompositeEntityDisplayData1093 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTSketchCompositeEntityDisplayData1093 = nil
			return fmt.Errorf("Failed to unmarshal BTSketchCompositeEntityDisplayData1093 as BTSketchImageDisplayData1379: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTSketchTextDisplayData-1707'
	if jsonDict["btType"] == "BTSketchTextDisplayData-1707" {
		// try to unmarshal JSON data into BTSketchTextDisplayData1707
		var qr *BTSketchTextDisplayData1707
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTSketchCompositeEntityDisplayData1093 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTSketchCompositeEntityDisplayData1093 = nil
			return fmt.Errorf("Failed to unmarshal BTSketchCompositeEntityDisplayData1093 as BTSketchTextDisplayData1707: %s", err.Error())
		}
	}

	var qtx *base_BTSketchCompositeEntityDisplayData1093
	err = json.Unmarshal(data, &qtx)
	if err == nil {
		dst.implBTSketchCompositeEntityDisplayData1093 = qtx
		return nil // data stored in dst.base_BTSketchCompositeEntityDisplayData1093, return on the first match
	} else {
		dst.implBTSketchCompositeEntityDisplayData1093 = nil
		return fmt.Errorf("Failed to unmarshal BTSketchCompositeEntityDisplayData1093 as base_BTSketchCompositeEntityDisplayData1093: %s", err.Error())
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src BTSketchCompositeEntityDisplayData1093) MarshalJSON() ([]byte, error) {
	ret := src.GetActualInstance()
	if ret == nil {
		return nil, nil // no data in oneOf schemas
	} else {
		return json.Marshal(&ret)
	}
}

// Get the actual instance
func (obj *BTSketchCompositeEntityDisplayData1093) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	return obj.implBTSketchCompositeEntityDisplayData1093
}

type NullableBTSketchCompositeEntityDisplayData1093 struct {
	value *BTSketchCompositeEntityDisplayData1093
	isSet bool
}

func (v NullableBTSketchCompositeEntityDisplayData1093) Get() *BTSketchCompositeEntityDisplayData1093 {
	return v.value
}

func (v *NullableBTSketchCompositeEntityDisplayData1093) Set(val *BTSketchCompositeEntityDisplayData1093) {
	v.value = val
	v.isSet = true
}

func (v NullableBTSketchCompositeEntityDisplayData1093) IsSet() bool {
	return v.isSet
}

func (v *NullableBTSketchCompositeEntityDisplayData1093) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTSketchCompositeEntityDisplayData1093(val *BTSketchCompositeEntityDisplayData1093) *NullableBTSketchCompositeEntityDisplayData1093 {
	return &NullableBTSketchCompositeEntityDisplayData1093{value: val, isSet: true}
}

func (v NullableBTSketchCompositeEntityDisplayData1093) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTSketchCompositeEntityDisplayData1093) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

type base_BTSketchCompositeEntityDisplayData1093 struct {
	BtType *string   `json:"btType,omitempty"`
	Points []float64 `json:"points,omitempty"`
}

// Newbase_BTSketchCompositeEntityDisplayData1093 instantiates a new base_BTSketchCompositeEntityDisplayData1093 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func Newbase_BTSketchCompositeEntityDisplayData1093() *base_BTSketchCompositeEntityDisplayData1093 {
	this := base_BTSketchCompositeEntityDisplayData1093{}
	return &this
}

// Newbase_BTSketchCompositeEntityDisplayData1093WithDefaults instantiates a new base_BTSketchCompositeEntityDisplayData1093 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func Newbase_BTSketchCompositeEntityDisplayData1093WithDefaults() *base_BTSketchCompositeEntityDisplayData1093 {
	this := base_BTSketchCompositeEntityDisplayData1093{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *base_BTSketchCompositeEntityDisplayData1093) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTSketchCompositeEntityDisplayData1093) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *base_BTSketchCompositeEntityDisplayData1093) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *base_BTSketchCompositeEntityDisplayData1093) SetBtType(v string) {
	o.BtType = &v
}

// GetPoints returns the Points field value if set, zero value otherwise.
func (o *base_BTSketchCompositeEntityDisplayData1093) GetPoints() []float64 {
	if o == nil || o.Points == nil {
		var ret []float64
		return ret
	}
	return o.Points
}

// GetPointsOk returns a tuple with the Points field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTSketchCompositeEntityDisplayData1093) GetPointsOk() ([]float64, bool) {
	if o == nil || o.Points == nil {
		return nil, false
	}
	return o.Points, true
}

// HasPoints returns a boolean if a field has been set.
func (o *base_BTSketchCompositeEntityDisplayData1093) HasPoints() bool {
	if o != nil && o.Points != nil {
		return true
	}

	return false
}

// SetPoints gets a reference to the given []float64 and assigns it to the Points field.
func (o *base_BTSketchCompositeEntityDisplayData1093) SetPoints(v []float64) {
	o.Points = v
}

func (o base_BTSketchCompositeEntityDisplayData1093) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.Points != nil {
		toSerialize["points"] = o.Points
	}
	return json.Marshal(toSerialize)
}
