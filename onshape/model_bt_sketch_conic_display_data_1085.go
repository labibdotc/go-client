/*
Onshape REST API

## Welcome to the Onshape REST API Explorer  To use this API explorer, sign in to your [Onshape](https://cad.onshape.com) account in another tab, then click the **Try it out** button below (it toggles to a **Cancel** button when selected).  See the **[API Explorer Guide](https://onshape-public.github.io/docs/api-intro/explorer/)** for help navigating this API Explorer, including **[authentication](https://onshape-public.github.io/docs/api-intro/explorer/#authentication)**.  **Tip:** To ensure the current session isn't used when trying other authentication techniques, make sure to [remove the Onshape cookie](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site) as per the instructions for your browser. Alternatively, you can use a private or incognito window.  ## See Also  * [Onshape API Guide](https://onshape-public.github.io/docs/): Our full suite of developer guides, to be used as an accompaniment to this API Explorer. * [Onshape Developer Portal](https://dev-portal.onshape.com/): The Onshape portal for managing your API keys, OAuth2 credentials, your Onshape applications, and your Onshape App Store entries. * [Authentication Guide](https://onshape-public.github.io/docs/auth/): Our guide to using API keys, request signatures, and OAuth2 in  your Onshape applications.

API version: 1.175.29320-74695940af99
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTSketchConicDisplayData1085 struct for BTSketchConicDisplayData1085
type BTSketchConicDisplayData1085 struct {
	BtType *string   `json:"btType,omitempty"`
	Points []float64 `json:"points,omitempty"`
	Offset *float64  `json:"offset,omitempty"`
	Rho    *float64  `json:"rho,omitempty"`
}

// NewBTSketchConicDisplayData1085 instantiates a new BTSketchConicDisplayData1085 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTSketchConicDisplayData1085() *BTSketchConicDisplayData1085 {
	this := BTSketchConicDisplayData1085{}
	return &this
}

// NewBTSketchConicDisplayData1085WithDefaults instantiates a new BTSketchConicDisplayData1085 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTSketchConicDisplayData1085WithDefaults() *BTSketchConicDisplayData1085 {
	this := BTSketchConicDisplayData1085{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTSketchConicDisplayData1085) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSketchConicDisplayData1085) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTSketchConicDisplayData1085) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTSketchConicDisplayData1085) SetBtType(v string) {
	o.BtType = &v
}

// GetPoints returns the Points field value if set, zero value otherwise.
func (o *BTSketchConicDisplayData1085) GetPoints() []float64 {
	if o == nil || o.Points == nil {
		var ret []float64
		return ret
	}
	return o.Points
}

// GetPointsOk returns a tuple with the Points field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSketchConicDisplayData1085) GetPointsOk() ([]float64, bool) {
	if o == nil || o.Points == nil {
		return nil, false
	}
	return o.Points, true
}

// HasPoints returns a boolean if a field has been set.
func (o *BTSketchConicDisplayData1085) HasPoints() bool {
	if o != nil && o.Points != nil {
		return true
	}

	return false
}

// SetPoints gets a reference to the given []float64 and assigns it to the Points field.
func (o *BTSketchConicDisplayData1085) SetPoints(v []float64) {
	o.Points = v
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *BTSketchConicDisplayData1085) GetOffset() float64 {
	if o == nil || o.Offset == nil {
		var ret float64
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSketchConicDisplayData1085) GetOffsetOk() (*float64, bool) {
	if o == nil || o.Offset == nil {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *BTSketchConicDisplayData1085) HasOffset() bool {
	if o != nil && o.Offset != nil {
		return true
	}

	return false
}

// SetOffset gets a reference to the given float64 and assigns it to the Offset field.
func (o *BTSketchConicDisplayData1085) SetOffset(v float64) {
	o.Offset = &v
}

// GetRho returns the Rho field value if set, zero value otherwise.
func (o *BTSketchConicDisplayData1085) GetRho() float64 {
	if o == nil || o.Rho == nil {
		var ret float64
		return ret
	}
	return *o.Rho
}

// GetRhoOk returns a tuple with the Rho field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSketchConicDisplayData1085) GetRhoOk() (*float64, bool) {
	if o == nil || o.Rho == nil {
		return nil, false
	}
	return o.Rho, true
}

// HasRho returns a boolean if a field has been set.
func (o *BTSketchConicDisplayData1085) HasRho() bool {
	if o != nil && o.Rho != nil {
		return true
	}

	return false
}

// SetRho gets a reference to the given float64 and assigns it to the Rho field.
func (o *BTSketchConicDisplayData1085) SetRho(v float64) {
	o.Rho = &v
}

func (o BTSketchConicDisplayData1085) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.Points != nil {
		toSerialize["points"] = o.Points
	}
	if o.Offset != nil {
		toSerialize["offset"] = o.Offset
	}
	if o.Rho != nil {
		toSerialize["rho"] = o.Rho
	}
	return json.Marshal(toSerialize)
}

type NullableBTSketchConicDisplayData1085 struct {
	value *BTSketchConicDisplayData1085
	isSet bool
}

func (v NullableBTSketchConicDisplayData1085) Get() *BTSketchConicDisplayData1085 {
	return v.value
}

func (v *NullableBTSketchConicDisplayData1085) Set(val *BTSketchConicDisplayData1085) {
	v.value = val
	v.isSet = true
}

func (v NullableBTSketchConicDisplayData1085) IsSet() bool {
	return v.isSet
}

func (v *NullableBTSketchConicDisplayData1085) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTSketchConicDisplayData1085(val *BTSketchConicDisplayData1085) *NullableBTSketchConicDisplayData1085 {
	return &NullableBTSketchConicDisplayData1085{value: val, isSet: true}
}

func (v NullableBTSketchConicDisplayData1085) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTSketchConicDisplayData1085) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
