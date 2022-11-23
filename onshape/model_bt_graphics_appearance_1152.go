/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.156.7802-6b1a0d70e58f
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
	"fmt"
)

// BTGraphicsAppearance1152 - struct for BTGraphicsAppearance1152
type BTGraphicsAppearance1152 struct {
	implBTGraphicsAppearance1152 interface{}
}

// BTGeneratedGraphicsAppearance4159AsBTGraphicsAppearance1152 is a convenience function that returns BTGeneratedGraphicsAppearance4159 wrapped in BTGraphicsAppearance1152
func (o *BTGeneratedGraphicsAppearance4159) AsBTGraphicsAppearance1152() *BTGraphicsAppearance1152 {
	return &BTGraphicsAppearance1152{o}
}

// NewBTGraphicsAppearance1152 instantiates a new BTGraphicsAppearance1152 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTGraphicsAppearance1152() *BTGraphicsAppearance1152 {
	this := BTGraphicsAppearance1152{Newbase_BTGraphicsAppearance1152()}
	return &this
}

// NewBTGraphicsAppearance1152WithDefaults instantiates a new BTGraphicsAppearance1152 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTGraphicsAppearance1152WithDefaults() *BTGraphicsAppearance1152 {
	this := BTGraphicsAppearance1152{Newbase_BTGraphicsAppearance1152WithDefaults()}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTGraphicsAppearance1152) GetBtType() string {
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
func (o *BTGraphicsAppearance1152) GetBtTypeOk() (*string, bool) {
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
func (o *BTGraphicsAppearance1152) HasBtType() bool {
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
func (o *BTGraphicsAppearance1152) SetBtType(v string) {
	type getResult interface {
		SetBtType(v string)
	}

	o.GetActualInstance().(getResult).SetBtType(v)
}

// GetColor returns the Color field value if set, zero value otherwise.
func (o *BTGraphicsAppearance1152) GetColor() []string {
	type getResult interface {
		GetColor() []string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetColor()
	} else {
		var de []string
		return de
	}
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTGraphicsAppearance1152) GetColorOk() ([]string, bool) {
	type getResult interface {
		GetColorOk() ([]string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetColorOk()
	} else {
		return nil, false
	}
}

// HasColor returns a boolean if a field has been set.
func (o *BTGraphicsAppearance1152) HasColor() bool {
	type getResult interface {
		HasColor() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasColor()
	} else {
		return false
	}
}

// SetColor gets a reference to the given []string and assigns it to the Color field.
func (o *BTGraphicsAppearance1152) SetColor(v []string) {
	type getResult interface {
		SetColor(v []string)
	}

	o.GetActualInstance().(getResult).SetColor(v)
}

// GetNonTrivial returns the NonTrivial field value if set, zero value otherwise.
func (o *BTGraphicsAppearance1152) GetNonTrivial() bool {
	type getResult interface {
		GetNonTrivial() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetNonTrivial()
	} else {
		var de bool
		return de
	}
}

// GetNonTrivialOk returns a tuple with the NonTrivial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTGraphicsAppearance1152) GetNonTrivialOk() (*bool, bool) {
	type getResult interface {
		GetNonTrivialOk() (*bool, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetNonTrivialOk()
	} else {
		return nil, false
	}
}

// HasNonTrivial returns a boolean if a field has been set.
func (o *BTGraphicsAppearance1152) HasNonTrivial() bool {
	type getResult interface {
		HasNonTrivial() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasNonTrivial()
	} else {
		return false
	}
}

// SetNonTrivial gets a reference to the given bool and assigns it to the NonTrivial field.
func (o *BTGraphicsAppearance1152) SetNonTrivial(v bool) {
	type getResult interface {
		SetNonTrivial(v bool)
	}

	o.GetActualInstance().(getResult).SetNonTrivial(v)
}

// GetOpacity returns the Opacity field value if set, zero value otherwise.
func (o *BTGraphicsAppearance1152) GetOpacity() int32 {
	type getResult interface {
		GetOpacity() int32
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetOpacity()
	} else {
		var de int32
		return de
	}
}

// GetOpacityOk returns a tuple with the Opacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTGraphicsAppearance1152) GetOpacityOk() (*int32, bool) {
	type getResult interface {
		GetOpacityOk() (*int32, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetOpacityOk()
	} else {
		return nil, false
	}
}

// HasOpacity returns a boolean if a field has been set.
func (o *BTGraphicsAppearance1152) HasOpacity() bool {
	type getResult interface {
		HasOpacity() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasOpacity()
	} else {
		return false
	}
}

// SetOpacity gets a reference to the given int32 and assigns it to the Opacity field.
func (o *BTGraphicsAppearance1152) SetOpacity(v int32) {
	type getResult interface {
		SetOpacity(v int32)
	}

	o.GetActualInstance().(getResult).SetOpacity(v)
}

// GetReset returns the Reset field value if set, zero value otherwise.
func (o *BTGraphicsAppearance1152) GetReset() bool {
	type getResult interface {
		GetReset() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetReset()
	} else {
		var de bool
		return de
	}
}

// GetResetOk returns a tuple with the Reset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTGraphicsAppearance1152) GetResetOk() (*bool, bool) {
	type getResult interface {
		GetResetOk() (*bool, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetResetOk()
	} else {
		return nil, false
	}
}

// HasReset returns a boolean if a field has been set.
func (o *BTGraphicsAppearance1152) HasReset() bool {
	type getResult interface {
		HasReset() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasReset()
	} else {
		return false
	}
}

// SetReset gets a reference to the given bool and assigns it to the Reset field.
func (o *BTGraphicsAppearance1152) SetReset(v bool) {
	type getResult interface {
		SetReset(v bool)
	}

	o.GetActualInstance().(getResult).SetReset(v)
}

// GetRgbaColor returns the RgbaColor field value if set, zero value otherwise.
func (o *BTGraphicsAppearance1152) GetRgbaColor() []string {
	type getResult interface {
		GetRgbaColor() []string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetRgbaColor()
	} else {
		var de []string
		return de
	}
}

// GetRgbaColorOk returns a tuple with the RgbaColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTGraphicsAppearance1152) GetRgbaColorOk() ([]string, bool) {
	type getResult interface {
		GetRgbaColorOk() ([]string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetRgbaColorOk()
	} else {
		return nil, false
	}
}

// HasRgbaColor returns a boolean if a field has been set.
func (o *BTGraphicsAppearance1152) HasRgbaColor() bool {
	type getResult interface {
		HasRgbaColor() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasRgbaColor()
	} else {
		return false
	}
}

// SetRgbaColor gets a reference to the given []string and assigns it to the RgbaColor field.
func (o *BTGraphicsAppearance1152) SetRgbaColor(v []string) {
	type getResult interface {
		SetRgbaColor(v []string)
	}

	o.GetActualInstance().(getResult).SetRgbaColor(v)
}

// GetUsableAppearance returns the UsableAppearance field value if set, zero value otherwise.
func (o *BTGraphicsAppearance1152) GetUsableAppearance() bool {
	type getResult interface {
		GetUsableAppearance() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetUsableAppearance()
	} else {
		var de bool
		return de
	}
}

// GetUsableAppearanceOk returns a tuple with the UsableAppearance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTGraphicsAppearance1152) GetUsableAppearanceOk() (*bool, bool) {
	type getResult interface {
		GetUsableAppearanceOk() (*bool, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetUsableAppearanceOk()
	} else {
		return nil, false
	}
}

// HasUsableAppearance returns a boolean if a field has been set.
func (o *BTGraphicsAppearance1152) HasUsableAppearance() bool {
	type getResult interface {
		HasUsableAppearance() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasUsableAppearance()
	} else {
		return false
	}
}

// SetUsableAppearance gets a reference to the given bool and assigns it to the UsableAppearance field.
func (o *BTGraphicsAppearance1152) SetUsableAppearance(v bool) {
	type getResult interface {
		SetUsableAppearance(v bool)
	}

	o.GetActualInstance().(getResult).SetUsableAppearance(v)
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *BTGraphicsAppearance1152) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'BTGeneratedGraphicsAppearance-4159'
	if jsonDict["btType"] == "BTGeneratedGraphicsAppearance-4159" {
		// try to unmarshal JSON data into BTGeneratedGraphicsAppearance4159
		var qr *BTGeneratedGraphicsAppearance4159
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTGraphicsAppearance1152 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTGraphicsAppearance1152 = nil
			return fmt.Errorf("Failed to unmarshal BTGraphicsAppearance1152 as BTGeneratedGraphicsAppearance4159: %s", err.Error())
		}
	}

	var qtx *base_BTGraphicsAppearance1152
	err = json.Unmarshal(data, &qtx)
	if err == nil {
		dst.implBTGraphicsAppearance1152 = qtx
		return nil // data stored in dst.base_BTGraphicsAppearance1152, return on the first match
	} else {
		dst.implBTGraphicsAppearance1152 = nil
		return fmt.Errorf("Failed to unmarshal BTGraphicsAppearance1152 as base_BTGraphicsAppearance1152: %s", err.Error())
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src BTGraphicsAppearance1152) MarshalJSON() ([]byte, error) {
	ret := src.GetActualInstance()
	if ret == nil {
		return nil, nil // no data in oneOf schemas
	} else {
		return json.Marshal(&ret)
	}
}

// Get the actual instance
func (obj *BTGraphicsAppearance1152) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	return obj.implBTGraphicsAppearance1152
}

type NullableBTGraphicsAppearance1152 struct {
	value *BTGraphicsAppearance1152
	isSet bool
}

func (v NullableBTGraphicsAppearance1152) Get() *BTGraphicsAppearance1152 {
	return v.value
}

func (v *NullableBTGraphicsAppearance1152) Set(val *BTGraphicsAppearance1152) {
	v.value = val
	v.isSet = true
}

func (v NullableBTGraphicsAppearance1152) IsSet() bool {
	return v.isSet
}

func (v *NullableBTGraphicsAppearance1152) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTGraphicsAppearance1152(val *BTGraphicsAppearance1152) *NullableBTGraphicsAppearance1152 {
	return &NullableBTGraphicsAppearance1152{value: val, isSet: true}
}

func (v NullableBTGraphicsAppearance1152) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTGraphicsAppearance1152) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

type base_BTGraphicsAppearance1152 struct {
	BtType           *string  `json:"btType,omitempty"`
	Color            []string `json:"color,omitempty"`
	NonTrivial       *bool    `json:"nonTrivial,omitempty"`
	Opacity          *int32   `json:"opacity,omitempty"`
	Reset            *bool    `json:"reset,omitempty"`
	RgbaColor        []string `json:"rgbaColor,omitempty"`
	UsableAppearance *bool    `json:"usableAppearance,omitempty"`
}

// Newbase_BTGraphicsAppearance1152 instantiates a new base_BTGraphicsAppearance1152 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func Newbase_BTGraphicsAppearance1152() *base_BTGraphicsAppearance1152 {
	this := base_BTGraphicsAppearance1152{}
	return &this
}

// Newbase_BTGraphicsAppearance1152WithDefaults instantiates a new base_BTGraphicsAppearance1152 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func Newbase_BTGraphicsAppearance1152WithDefaults() *base_BTGraphicsAppearance1152 {
	this := base_BTGraphicsAppearance1152{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *base_BTGraphicsAppearance1152) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTGraphicsAppearance1152) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *base_BTGraphicsAppearance1152) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *base_BTGraphicsAppearance1152) SetBtType(v string) {
	o.BtType = &v
}

// GetColor returns the Color field value if set, zero value otherwise.
func (o *base_BTGraphicsAppearance1152) GetColor() []string {
	if o == nil || o.Color == nil {
		var ret []string
		return ret
	}
	return o.Color
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTGraphicsAppearance1152) GetColorOk() ([]string, bool) {
	if o == nil || o.Color == nil {
		return nil, false
	}
	return o.Color, true
}

// HasColor returns a boolean if a field has been set.
func (o *base_BTGraphicsAppearance1152) HasColor() bool {
	if o != nil && o.Color != nil {
		return true
	}

	return false
}

// SetColor gets a reference to the given []string and assigns it to the Color field.
func (o *base_BTGraphicsAppearance1152) SetColor(v []string) {
	o.Color = v
}

// GetNonTrivial returns the NonTrivial field value if set, zero value otherwise.
func (o *base_BTGraphicsAppearance1152) GetNonTrivial() bool {
	if o == nil || o.NonTrivial == nil {
		var ret bool
		return ret
	}
	return *o.NonTrivial
}

// GetNonTrivialOk returns a tuple with the NonTrivial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTGraphicsAppearance1152) GetNonTrivialOk() (*bool, bool) {
	if o == nil || o.NonTrivial == nil {
		return nil, false
	}
	return o.NonTrivial, true
}

// HasNonTrivial returns a boolean if a field has been set.
func (o *base_BTGraphicsAppearance1152) HasNonTrivial() bool {
	if o != nil && o.NonTrivial != nil {
		return true
	}

	return false
}

// SetNonTrivial gets a reference to the given bool and assigns it to the NonTrivial field.
func (o *base_BTGraphicsAppearance1152) SetNonTrivial(v bool) {
	o.NonTrivial = &v
}

// GetOpacity returns the Opacity field value if set, zero value otherwise.
func (o *base_BTGraphicsAppearance1152) GetOpacity() int32 {
	if o == nil || o.Opacity == nil {
		var ret int32
		return ret
	}
	return *o.Opacity
}

// GetOpacityOk returns a tuple with the Opacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTGraphicsAppearance1152) GetOpacityOk() (*int32, bool) {
	if o == nil || o.Opacity == nil {
		return nil, false
	}
	return o.Opacity, true
}

// HasOpacity returns a boolean if a field has been set.
func (o *base_BTGraphicsAppearance1152) HasOpacity() bool {
	if o != nil && o.Opacity != nil {
		return true
	}

	return false
}

// SetOpacity gets a reference to the given int32 and assigns it to the Opacity field.
func (o *base_BTGraphicsAppearance1152) SetOpacity(v int32) {
	o.Opacity = &v
}

// GetReset returns the Reset field value if set, zero value otherwise.
func (o *base_BTGraphicsAppearance1152) GetReset() bool {
	if o == nil || o.Reset == nil {
		var ret bool
		return ret
	}
	return *o.Reset
}

// GetResetOk returns a tuple with the Reset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTGraphicsAppearance1152) GetResetOk() (*bool, bool) {
	if o == nil || o.Reset == nil {
		return nil, false
	}
	return o.Reset, true
}

// HasReset returns a boolean if a field has been set.
func (o *base_BTGraphicsAppearance1152) HasReset() bool {
	if o != nil && o.Reset != nil {
		return true
	}

	return false
}

// SetReset gets a reference to the given bool and assigns it to the Reset field.
func (o *base_BTGraphicsAppearance1152) SetReset(v bool) {
	o.Reset = &v
}

// GetRgbaColor returns the RgbaColor field value if set, zero value otherwise.
func (o *base_BTGraphicsAppearance1152) GetRgbaColor() []string {
	if o == nil || o.RgbaColor == nil {
		var ret []string
		return ret
	}
	return o.RgbaColor
}

// GetRgbaColorOk returns a tuple with the RgbaColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTGraphicsAppearance1152) GetRgbaColorOk() ([]string, bool) {
	if o == nil || o.RgbaColor == nil {
		return nil, false
	}
	return o.RgbaColor, true
}

// HasRgbaColor returns a boolean if a field has been set.
func (o *base_BTGraphicsAppearance1152) HasRgbaColor() bool {
	if o != nil && o.RgbaColor != nil {
		return true
	}

	return false
}

// SetRgbaColor gets a reference to the given []string and assigns it to the RgbaColor field.
func (o *base_BTGraphicsAppearance1152) SetRgbaColor(v []string) {
	o.RgbaColor = v
}

// GetUsableAppearance returns the UsableAppearance field value if set, zero value otherwise.
func (o *base_BTGraphicsAppearance1152) GetUsableAppearance() bool {
	if o == nil || o.UsableAppearance == nil {
		var ret bool
		return ret
	}
	return *o.UsableAppearance
}

// GetUsableAppearanceOk returns a tuple with the UsableAppearance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTGraphicsAppearance1152) GetUsableAppearanceOk() (*bool, bool) {
	if o == nil || o.UsableAppearance == nil {
		return nil, false
	}
	return o.UsableAppearance, true
}

// HasUsableAppearance returns a boolean if a field has been set.
func (o *base_BTGraphicsAppearance1152) HasUsableAppearance() bool {
	if o != nil && o.UsableAppearance != nil {
		return true
	}

	return false
}

// SetUsableAppearance gets a reference to the given bool and assigns it to the UsableAppearance field.
func (o *base_BTGraphicsAppearance1152) SetUsableAppearance(v bool) {
	o.UsableAppearance = &v
}

func (o base_BTGraphicsAppearance1152) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.Color != nil {
		toSerialize["color"] = o.Color
	}
	if o.NonTrivial != nil {
		toSerialize["nonTrivial"] = o.NonTrivial
	}
	if o.Opacity != nil {
		toSerialize["opacity"] = o.Opacity
	}
	if o.Reset != nil {
		toSerialize["reset"] = o.Reset
	}
	if o.RgbaColor != nil {
		toSerialize["rgbaColor"] = o.RgbaColor
	}
	if o.UsableAppearance != nil {
		toSerialize["usableAppearance"] = o.UsableAppearance
	}
	return json.Marshal(toSerialize)
}
