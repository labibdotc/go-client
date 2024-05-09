/*
Onshape REST API

## Welcome to the Onshape REST API Explorer  To use this API explorer, sign in to your [Onshape](https://cad.onshape.com) account in another tab, then click the **Try it out** button below (it toggles to a **Cancel** button when selected).  See the **[API Explorer Guide](https://onshape-public.github.io/docs/api-intro/explorer/)** for help navigating this API Explorer, including **[authentication](https://onshape-public.github.io/docs/api-intro/explorer/#authentication)**.  **Tip:** To ensure the current session isn't used when trying other authentication techniques, make sure to [remove the Onshape cookie](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site) as per the instructions for your browser. Alternatively, you can use a private or incognito window.  ## See Also  * [Onshape API Guide](https://onshape-public.github.io/docs/): Our full suite of developer guides, to be used as an accompaniment to this API Explorer. * [Onshape Developer Portal](https://dev-portal.onshape.com/): The Onshape portal for managing your API keys, OAuth2 credentials, your Onshape applications, and your Onshape App Store entries. * [Authentication Guide](https://onshape-public.github.io/docs/auth/): Our guide to using API keys, request signatures, and OAuth2 in  your Onshape applications.

Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTCurveGeometryEllipse1189 struct for BTCurveGeometryEllipse1189
type BTCurveGeometryEllipse1189 struct {
	BTCurveGeometryCircle115
	BtType      *string  `json:"btType,omitempty"`
	Clockwise   *bool    `json:"clockwise,omitempty"`
	Radius      *float64 `json:"radius,omitempty"`
	Xcenter     *float64 `json:"xcenter,omitempty"`
	Xdir        *float64 `json:"xdir,omitempty"`
	Ycenter     *float64 `json:"ycenter,omitempty"`
	Ydir        *float64 `json:"ydir,omitempty"`
	MinorRadius *float64 `json:"minorRadius,omitempty"`
}

// NewBTCurveGeometryEllipse1189 instantiates a new BTCurveGeometryEllipse1189 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTCurveGeometryEllipse1189() *BTCurveGeometryEllipse1189 {
	this := BTCurveGeometryEllipse1189{}
	return &this
}

// NewBTCurveGeometryEllipse1189WithDefaults instantiates a new BTCurveGeometryEllipse1189 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTCurveGeometryEllipse1189WithDefaults() *BTCurveGeometryEllipse1189 {
	this := BTCurveGeometryEllipse1189{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTCurveGeometryEllipse1189) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCurveGeometryEllipse1189) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTCurveGeometryEllipse1189) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTCurveGeometryEllipse1189) SetBtType(v string) {
	o.BtType = &v
}

// GetClockwise returns the Clockwise field value if set, zero value otherwise.
func (o *BTCurveGeometryEllipse1189) GetClockwise() bool {
	if o == nil || o.Clockwise == nil {
		var ret bool
		return ret
	}
	return *o.Clockwise
}

// GetClockwiseOk returns a tuple with the Clockwise field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCurveGeometryEllipse1189) GetClockwiseOk() (*bool, bool) {
	if o == nil || o.Clockwise == nil {
		return nil, false
	}
	return o.Clockwise, true
}

// HasClockwise returns a boolean if a field has been set.
func (o *BTCurveGeometryEllipse1189) HasClockwise() bool {
	if o != nil && o.Clockwise != nil {
		return true
	}

	return false
}

// SetClockwise gets a reference to the given bool and assigns it to the Clockwise field.
func (o *BTCurveGeometryEllipse1189) SetClockwise(v bool) {
	o.Clockwise = &v
}

// GetRadius returns the Radius field value if set, zero value otherwise.
func (o *BTCurveGeometryEllipse1189) GetRadius() float64 {
	if o == nil || o.Radius == nil {
		var ret float64
		return ret
	}
	return *o.Radius
}

// GetRadiusOk returns a tuple with the Radius field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCurveGeometryEllipse1189) GetRadiusOk() (*float64, bool) {
	if o == nil || o.Radius == nil {
		return nil, false
	}
	return o.Radius, true
}

// HasRadius returns a boolean if a field has been set.
func (o *BTCurveGeometryEllipse1189) HasRadius() bool {
	if o != nil && o.Radius != nil {
		return true
	}

	return false
}

// SetRadius gets a reference to the given float64 and assigns it to the Radius field.
func (o *BTCurveGeometryEllipse1189) SetRadius(v float64) {
	o.Radius = &v
}

// GetXcenter returns the Xcenter field value if set, zero value otherwise.
func (o *BTCurveGeometryEllipse1189) GetXcenter() float64 {
	if o == nil || o.Xcenter == nil {
		var ret float64
		return ret
	}
	return *o.Xcenter
}

// GetXcenterOk returns a tuple with the Xcenter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCurveGeometryEllipse1189) GetXcenterOk() (*float64, bool) {
	if o == nil || o.Xcenter == nil {
		return nil, false
	}
	return o.Xcenter, true
}

// HasXcenter returns a boolean if a field has been set.
func (o *BTCurveGeometryEllipse1189) HasXcenter() bool {
	if o != nil && o.Xcenter != nil {
		return true
	}

	return false
}

// SetXcenter gets a reference to the given float64 and assigns it to the Xcenter field.
func (o *BTCurveGeometryEllipse1189) SetXcenter(v float64) {
	o.Xcenter = &v
}

// GetXdir returns the Xdir field value if set, zero value otherwise.
func (o *BTCurveGeometryEllipse1189) GetXdir() float64 {
	if o == nil || o.Xdir == nil {
		var ret float64
		return ret
	}
	return *o.Xdir
}

// GetXdirOk returns a tuple with the Xdir field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCurveGeometryEllipse1189) GetXdirOk() (*float64, bool) {
	if o == nil || o.Xdir == nil {
		return nil, false
	}
	return o.Xdir, true
}

// HasXdir returns a boolean if a field has been set.
func (o *BTCurveGeometryEllipse1189) HasXdir() bool {
	if o != nil && o.Xdir != nil {
		return true
	}

	return false
}

// SetXdir gets a reference to the given float64 and assigns it to the Xdir field.
func (o *BTCurveGeometryEllipse1189) SetXdir(v float64) {
	o.Xdir = &v
}

// GetYcenter returns the Ycenter field value if set, zero value otherwise.
func (o *BTCurveGeometryEllipse1189) GetYcenter() float64 {
	if o == nil || o.Ycenter == nil {
		var ret float64
		return ret
	}
	return *o.Ycenter
}

// GetYcenterOk returns a tuple with the Ycenter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCurveGeometryEllipse1189) GetYcenterOk() (*float64, bool) {
	if o == nil || o.Ycenter == nil {
		return nil, false
	}
	return o.Ycenter, true
}

// HasYcenter returns a boolean if a field has been set.
func (o *BTCurveGeometryEllipse1189) HasYcenter() bool {
	if o != nil && o.Ycenter != nil {
		return true
	}

	return false
}

// SetYcenter gets a reference to the given float64 and assigns it to the Ycenter field.
func (o *BTCurveGeometryEllipse1189) SetYcenter(v float64) {
	o.Ycenter = &v
}

// GetYdir returns the Ydir field value if set, zero value otherwise.
func (o *BTCurveGeometryEllipse1189) GetYdir() float64 {
	if o == nil || o.Ydir == nil {
		var ret float64
		return ret
	}
	return *o.Ydir
}

// GetYdirOk returns a tuple with the Ydir field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCurveGeometryEllipse1189) GetYdirOk() (*float64, bool) {
	if o == nil || o.Ydir == nil {
		return nil, false
	}
	return o.Ydir, true
}

// HasYdir returns a boolean if a field has been set.
func (o *BTCurveGeometryEllipse1189) HasYdir() bool {
	if o != nil && o.Ydir != nil {
		return true
	}

	return false
}

// SetYdir gets a reference to the given float64 and assigns it to the Ydir field.
func (o *BTCurveGeometryEllipse1189) SetYdir(v float64) {
	o.Ydir = &v
}

// GetMinorRadius returns the MinorRadius field value if set, zero value otherwise.
func (o *BTCurveGeometryEllipse1189) GetMinorRadius() float64 {
	if o == nil || o.MinorRadius == nil {
		var ret float64
		return ret
	}
	return *o.MinorRadius
}

// GetMinorRadiusOk returns a tuple with the MinorRadius field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCurveGeometryEllipse1189) GetMinorRadiusOk() (*float64, bool) {
	if o == nil || o.MinorRadius == nil {
		return nil, false
	}
	return o.MinorRadius, true
}

// HasMinorRadius returns a boolean if a field has been set.
func (o *BTCurveGeometryEllipse1189) HasMinorRadius() bool {
	if o != nil && o.MinorRadius != nil {
		return true
	}

	return false
}

// SetMinorRadius gets a reference to the given float64 and assigns it to the MinorRadius field.
func (o *BTCurveGeometryEllipse1189) SetMinorRadius(v float64) {
	o.MinorRadius = &v
}

func (o BTCurveGeometryEllipse1189) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedBTCurveGeometryCircle115, errBTCurveGeometryCircle115 := json.Marshal(o.BTCurveGeometryCircle115)
	if errBTCurveGeometryCircle115 != nil {
		return []byte{}, errBTCurveGeometryCircle115
	}
	errBTCurveGeometryCircle115 = json.Unmarshal([]byte(serializedBTCurveGeometryCircle115), &toSerialize)
	if errBTCurveGeometryCircle115 != nil {
		return []byte{}, errBTCurveGeometryCircle115
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.Clockwise != nil {
		toSerialize["clockwise"] = o.Clockwise
	}
	if o.Radius != nil {
		toSerialize["radius"] = o.Radius
	}
	if o.Xcenter != nil {
		toSerialize["xcenter"] = o.Xcenter
	}
	if o.Xdir != nil {
		toSerialize["xdir"] = o.Xdir
	}
	if o.Ycenter != nil {
		toSerialize["ycenter"] = o.Ycenter
	}
	if o.Ydir != nil {
		toSerialize["ydir"] = o.Ydir
	}
	if o.MinorRadius != nil {
		toSerialize["minorRadius"] = o.MinorRadius
	}
	return json.Marshal(toSerialize)
}

type NullableBTCurveGeometryEllipse1189 struct {
	value *BTCurveGeometryEllipse1189
	isSet bool
}

func (v NullableBTCurveGeometryEllipse1189) Get() *BTCurveGeometryEllipse1189 {
	return v.value
}

func (v *NullableBTCurveGeometryEllipse1189) Set(val *BTCurveGeometryEllipse1189) {
	v.value = val
	v.isSet = true
}

func (v NullableBTCurveGeometryEllipse1189) IsSet() bool {
	return v.isSet
}

func (v *NullableBTCurveGeometryEllipse1189) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTCurveGeometryEllipse1189(val *BTCurveGeometryEllipse1189) *NullableBTCurveGeometryEllipse1189 {
	return &NullableBTCurveGeometryEllipse1189{value: val, isSet: true}
}

func (v NullableBTCurveGeometryEllipse1189) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTCurveGeometryEllipse1189) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
