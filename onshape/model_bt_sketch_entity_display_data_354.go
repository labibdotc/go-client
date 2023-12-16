/*
Onshape REST API

## Welcome to the Onshape REST API Explorer  To use this API explorer, sign in to your [Onshape](https://cad.onshape.com) account in another tab, then click the **Try it out** button below (it toggles to a **Cancel** button when selected).  See the **[API Explorer Guide](https://onshape-public.github.io/docs/api-intro/explorer/)** for help navigating this API Explorer, including **[authentication](https://onshape-public.github.io/docs/api-intro/explorer/#authentication)**.  **Tip:** To ensure the current session isn't used when trying other authentication techniques, make sure to [remove the Onshape cookie](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site) as per the instructions for your browser. Alternatively, you can use a private or incognito window.  ## See Also  * [Onshape API Guide](https://onshape-public.github.io/docs/): Our full suite of developer guides, to be used as an accompaniment to this API Explorer. * [Onshape Developer Portal](https://dev-portal.onshape.com/): The Onshape portal for managing your API keys, OAuth2 credentials, your Onshape applications, and your Onshape App Store entries. * [Authentication Guide](https://onshape-public.github.io/docs/auth/): Our guide to using API keys, request signatures, and OAuth2 in  your Onshape applications.

API version: 1.174.27783-ab3907bf6199
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
	"fmt"
)

// BTSketchEntityDisplayData354 - struct for BTSketchEntityDisplayData354
type BTSketchEntityDisplayData354 struct {
	implBTSketchEntityDisplayData354 interface{}
}

// BTSketchEllipticalArcDisplayData892AsBTSketchEntityDisplayData354 is a convenience function that returns BTSketchEllipticalArcDisplayData892 wrapped in BTSketchEntityDisplayData354
func (o *BTSketchEllipticalArcDisplayData892) AsBTSketchEntityDisplayData354() *BTSketchEntityDisplayData354 {
	return &BTSketchEntityDisplayData354{o}
}

// BTSketchTextDisplayData1707AsBTSketchEntityDisplayData354 is a convenience function that returns BTSketchTextDisplayData1707 wrapped in BTSketchEntityDisplayData354
func (o *BTSketchTextDisplayData1707) AsBTSketchEntityDisplayData354() *BTSketchEntityDisplayData354 {
	return &BTSketchEntityDisplayData354{o}
}

// BTSketchLineDisplayData357AsBTSketchEntityDisplayData354 is a convenience function that returns BTSketchLineDisplayData357 wrapped in BTSketchEntityDisplayData354
func (o *BTSketchLineDisplayData357) AsBTSketchEntityDisplayData354() *BTSketchEntityDisplayData354 {
	return &BTSketchEntityDisplayData354{o}
}

// BTSketchConicDisplayData1085AsBTSketchEntityDisplayData354 is a convenience function that returns BTSketchConicDisplayData1085 wrapped in BTSketchEntityDisplayData354
func (o *BTSketchConicDisplayData1085) AsBTSketchEntityDisplayData354() *BTSketchEntityDisplayData354 {
	return &BTSketchEntityDisplayData354{o}
}

// BTSketchEllipseDisplayData712AsBTSketchEntityDisplayData354 is a convenience function that returns BTSketchEllipseDisplayData712 wrapped in BTSketchEntityDisplayData354
func (o *BTSketchEllipseDisplayData712) AsBTSketchEntityDisplayData354() *BTSketchEntityDisplayData354 {
	return &BTSketchEntityDisplayData354{o}
}

// BTSketchCompositeEntityDisplayData1093AsBTSketchEntityDisplayData354 is a convenience function that returns BTSketchCompositeEntityDisplayData1093 wrapped in BTSketchEntityDisplayData354
func (o *BTSketchCompositeEntityDisplayData1093) AsBTSketchEntityDisplayData354() *BTSketchEntityDisplayData354 {
	return &BTSketchEntityDisplayData354{o}
}

// BTSketchCircleDisplayData350AsBTSketchEntityDisplayData354 is a convenience function that returns BTSketchCircleDisplayData350 wrapped in BTSketchEntityDisplayData354
func (o *BTSketchCircleDisplayData350) AsBTSketchEntityDisplayData354() *BTSketchEntityDisplayData354 {
	return &BTSketchEntityDisplayData354{o}
}

// BTSketchSplineDisplayData359AsBTSketchEntityDisplayData354 is a convenience function that returns BTSketchSplineDisplayData359 wrapped in BTSketchEntityDisplayData354
func (o *BTSketchSplineDisplayData359) AsBTSketchEntityDisplayData354() *BTSketchEntityDisplayData354 {
	return &BTSketchEntityDisplayData354{o}
}

// BTSketchArcDisplayData349AsBTSketchEntityDisplayData354 is a convenience function that returns BTSketchArcDisplayData349 wrapped in BTSketchEntityDisplayData354
func (o *BTSketchArcDisplayData349) AsBTSketchEntityDisplayData354() *BTSketchEntityDisplayData354 {
	return &BTSketchEntityDisplayData354{o}
}

// BTSketchPointDisplayData358AsBTSketchEntityDisplayData354 is a convenience function that returns BTSketchPointDisplayData358 wrapped in BTSketchEntityDisplayData354
func (o *BTSketchPointDisplayData358) AsBTSketchEntityDisplayData354() *BTSketchEntityDisplayData354 {
	return &BTSketchEntityDisplayData354{o}
}

// NewBTSketchEntityDisplayData354 instantiates a new BTSketchEntityDisplayData354 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTSketchEntityDisplayData354() *BTSketchEntityDisplayData354 {
	this := BTSketchEntityDisplayData354{Newbase_BTSketchEntityDisplayData354()}
	return &this
}

// NewBTSketchEntityDisplayData354WithDefaults instantiates a new BTSketchEntityDisplayData354 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTSketchEntityDisplayData354WithDefaults() *BTSketchEntityDisplayData354 {
	this := BTSketchEntityDisplayData354{Newbase_BTSketchEntityDisplayData354WithDefaults()}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTSketchEntityDisplayData354) GetBtType() string {
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
func (o *BTSketchEntityDisplayData354) GetBtTypeOk() (*string, bool) {
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
func (o *BTSketchEntityDisplayData354) HasBtType() bool {
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
func (o *BTSketchEntityDisplayData354) SetBtType(v string) {
	type getResult interface {
		SetBtType(v string)
	}

	o.GetActualInstance().(getResult).SetBtType(v)
}

// GetPoints returns the Points field value if set, zero value otherwise.
func (o *BTSketchEntityDisplayData354) GetPoints() []float64 {
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
func (o *BTSketchEntityDisplayData354) GetPointsOk() ([]float64, bool) {
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
func (o *BTSketchEntityDisplayData354) HasPoints() bool {
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
func (o *BTSketchEntityDisplayData354) SetPoints(v []float64) {
	type getResult interface {
		SetPoints(v []float64)
	}

	o.GetActualInstance().(getResult).SetPoints(v)
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *BTSketchEntityDisplayData354) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'BTSketchArcDisplayData-349'
	if jsonDict["btType"] == "BTSketchArcDisplayData-349" {
		// try to unmarshal JSON data into BTSketchArcDisplayData349
		var qr *BTSketchArcDisplayData349
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTSketchEntityDisplayData354 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTSketchEntityDisplayData354 = nil
			return fmt.Errorf("Failed to unmarshal BTSketchEntityDisplayData354 as BTSketchArcDisplayData349: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTSketchCircleDisplayData-350'
	if jsonDict["btType"] == "BTSketchCircleDisplayData-350" {
		// try to unmarshal JSON data into BTSketchCircleDisplayData350
		var qr *BTSketchCircleDisplayData350
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTSketchEntityDisplayData354 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTSketchEntityDisplayData354 = nil
			return fmt.Errorf("Failed to unmarshal BTSketchEntityDisplayData354 as BTSketchCircleDisplayData350: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTSketchCompositeEntityDisplayData-1093'
	if jsonDict["btType"] == "BTSketchCompositeEntityDisplayData-1093" {
		// try to unmarshal JSON data into BTSketchCompositeEntityDisplayData1093
		var qr *BTSketchCompositeEntityDisplayData1093
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTSketchEntityDisplayData354 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTSketchEntityDisplayData354 = nil
			return fmt.Errorf("Failed to unmarshal BTSketchEntityDisplayData354 as BTSketchCompositeEntityDisplayData1093: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTSketchConicDisplayData-1085'
	if jsonDict["btType"] == "BTSketchConicDisplayData-1085" {
		// try to unmarshal JSON data into BTSketchConicDisplayData1085
		var qr *BTSketchConicDisplayData1085
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTSketchEntityDisplayData354 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTSketchEntityDisplayData354 = nil
			return fmt.Errorf("Failed to unmarshal BTSketchEntityDisplayData354 as BTSketchConicDisplayData1085: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTSketchEllipseDisplayData-712'
	if jsonDict["btType"] == "BTSketchEllipseDisplayData-712" {
		// try to unmarshal JSON data into BTSketchEllipseDisplayData712
		var qr *BTSketchEllipseDisplayData712
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTSketchEntityDisplayData354 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTSketchEntityDisplayData354 = nil
			return fmt.Errorf("Failed to unmarshal BTSketchEntityDisplayData354 as BTSketchEllipseDisplayData712: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTSketchEllipticalArcDisplayData-892'
	if jsonDict["btType"] == "BTSketchEllipticalArcDisplayData-892" {
		// try to unmarshal JSON data into BTSketchEllipticalArcDisplayData892
		var qr *BTSketchEllipticalArcDisplayData892
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTSketchEntityDisplayData354 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTSketchEntityDisplayData354 = nil
			return fmt.Errorf("Failed to unmarshal BTSketchEntityDisplayData354 as BTSketchEllipticalArcDisplayData892: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTSketchLineDisplayData-357'
	if jsonDict["btType"] == "BTSketchLineDisplayData-357" {
		// try to unmarshal JSON data into BTSketchLineDisplayData357
		var qr *BTSketchLineDisplayData357
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTSketchEntityDisplayData354 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTSketchEntityDisplayData354 = nil
			return fmt.Errorf("Failed to unmarshal BTSketchEntityDisplayData354 as BTSketchLineDisplayData357: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTSketchPointDisplayData-358'
	if jsonDict["btType"] == "BTSketchPointDisplayData-358" {
		// try to unmarshal JSON data into BTSketchPointDisplayData358
		var qr *BTSketchPointDisplayData358
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTSketchEntityDisplayData354 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTSketchEntityDisplayData354 = nil
			return fmt.Errorf("Failed to unmarshal BTSketchEntityDisplayData354 as BTSketchPointDisplayData358: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTSketchSplineDisplayData-359'
	if jsonDict["btType"] == "BTSketchSplineDisplayData-359" {
		// try to unmarshal JSON data into BTSketchSplineDisplayData359
		var qr *BTSketchSplineDisplayData359
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTSketchEntityDisplayData354 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTSketchEntityDisplayData354 = nil
			return fmt.Errorf("Failed to unmarshal BTSketchEntityDisplayData354 as BTSketchSplineDisplayData359: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTSketchTextDisplayData-1707'
	if jsonDict["btType"] == "BTSketchTextDisplayData-1707" {
		// try to unmarshal JSON data into BTSketchTextDisplayData1707
		var qr *BTSketchTextDisplayData1707
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTSketchEntityDisplayData354 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTSketchEntityDisplayData354 = nil
			return fmt.Errorf("Failed to unmarshal BTSketchEntityDisplayData354 as BTSketchTextDisplayData1707: %s", err.Error())
		}
	}

	var qtx *base_BTSketchEntityDisplayData354
	err = json.Unmarshal(data, &qtx)
	if err == nil {
		dst.implBTSketchEntityDisplayData354 = qtx
		return nil // data stored in dst.base_BTSketchEntityDisplayData354, return on the first match
	} else {
		dst.implBTSketchEntityDisplayData354 = nil
		return fmt.Errorf("Failed to unmarshal BTSketchEntityDisplayData354 as base_BTSketchEntityDisplayData354: %s", err.Error())
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src BTSketchEntityDisplayData354) MarshalJSON() ([]byte, error) {
	ret := src.GetActualInstance()
	if ret == nil {
		return nil, nil // no data in oneOf schemas
	} else {
		return json.Marshal(&ret)
	}
}

// Get the actual instance
func (obj *BTSketchEntityDisplayData354) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	return obj.implBTSketchEntityDisplayData354
}

type NullableBTSketchEntityDisplayData354 struct {
	value *BTSketchEntityDisplayData354
	isSet bool
}

func (v NullableBTSketchEntityDisplayData354) Get() *BTSketchEntityDisplayData354 {
	return v.value
}

func (v *NullableBTSketchEntityDisplayData354) Set(val *BTSketchEntityDisplayData354) {
	v.value = val
	v.isSet = true
}

func (v NullableBTSketchEntityDisplayData354) IsSet() bool {
	return v.isSet
}

func (v *NullableBTSketchEntityDisplayData354) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTSketchEntityDisplayData354(val *BTSketchEntityDisplayData354) *NullableBTSketchEntityDisplayData354 {
	return &NullableBTSketchEntityDisplayData354{value: val, isSet: true}
}

func (v NullableBTSketchEntityDisplayData354) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTSketchEntityDisplayData354) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

type base_BTSketchEntityDisplayData354 struct {
	BtType *string   `json:"btType,omitempty"`
	Points []float64 `json:"points,omitempty"`
}

// Newbase_BTSketchEntityDisplayData354 instantiates a new base_BTSketchEntityDisplayData354 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func Newbase_BTSketchEntityDisplayData354() *base_BTSketchEntityDisplayData354 {
	this := base_BTSketchEntityDisplayData354{}
	return &this
}

// Newbase_BTSketchEntityDisplayData354WithDefaults instantiates a new base_BTSketchEntityDisplayData354 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func Newbase_BTSketchEntityDisplayData354WithDefaults() *base_BTSketchEntityDisplayData354 {
	this := base_BTSketchEntityDisplayData354{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *base_BTSketchEntityDisplayData354) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTSketchEntityDisplayData354) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *base_BTSketchEntityDisplayData354) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *base_BTSketchEntityDisplayData354) SetBtType(v string) {
	o.BtType = &v
}

// GetPoints returns the Points field value if set, zero value otherwise.
func (o *base_BTSketchEntityDisplayData354) GetPoints() []float64 {
	if o == nil || o.Points == nil {
		var ret []float64
		return ret
	}
	return o.Points
}

// GetPointsOk returns a tuple with the Points field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTSketchEntityDisplayData354) GetPointsOk() ([]float64, bool) {
	if o == nil || o.Points == nil {
		return nil, false
	}
	return o.Points, true
}

// HasPoints returns a boolean if a field has been set.
func (o *base_BTSketchEntityDisplayData354) HasPoints() bool {
	if o != nil && o.Points != nil {
		return true
	}

	return false
}

// SetPoints gets a reference to the given []float64 and assigns it to the Points field.
func (o *base_BTSketchEntityDisplayData354) SetPoints(v []float64) {
	o.Points = v
}

func (o base_BTSketchEntityDisplayData354) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.Points != nil {
		toSerialize["points"] = o.Points
	}
	return json.Marshal(toSerialize)
}
