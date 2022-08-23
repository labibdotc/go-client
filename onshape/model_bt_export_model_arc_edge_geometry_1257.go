/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.152.6159-a7ef074944fe
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTExportModelArcEdgeGeometry1257 struct for BTExportModelArcEdgeGeometry1257
type BTExportModelArcEdgeGeometry1257 struct {
	EndPoint       *BTVector3d389 `json:"endPoint,omitempty"`
	EndVector      *BTVector3d389 `json:"endVector,omitempty"`
	Length         *float64       `json:"length,omitempty"`
	MidPoint       *BTVector3d389 `json:"midPoint,omitempty"`
	QuarterPoint   *BTVector3d389 `json:"quarterPoint,omitempty"`
	StartPoint     *BTVector3d389 `json:"startPoint,omitempty"`
	StartVector    *BTVector3d389 `json:"startVector,omitempty"`
	ArcIsClockwise *bool          `json:"arcIsClockwise,omitempty"`
	ArcSweep       *float64       `json:"arcSweep,omitempty"`
	BtType         *string        `json:"btType,omitempty"`
}

// NewBTExportModelArcEdgeGeometry1257 instantiates a new BTExportModelArcEdgeGeometry1257 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTExportModelArcEdgeGeometry1257() *BTExportModelArcEdgeGeometry1257 {
	this := BTExportModelArcEdgeGeometry1257{}
	return &this
}

// NewBTExportModelArcEdgeGeometry1257WithDefaults instantiates a new BTExportModelArcEdgeGeometry1257 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTExportModelArcEdgeGeometry1257WithDefaults() *BTExportModelArcEdgeGeometry1257 {
	this := BTExportModelArcEdgeGeometry1257{}
	return &this
}

// GetEndPoint returns the EndPoint field value if set, zero value otherwise.
func (o *BTExportModelArcEdgeGeometry1257) GetEndPoint() BTVector3d389 {
	if o == nil || o.EndPoint == nil {
		var ret BTVector3d389
		return ret
	}
	return *o.EndPoint
}

// GetEndPointOk returns a tuple with the EndPoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExportModelArcEdgeGeometry1257) GetEndPointOk() (*BTVector3d389, bool) {
	if o == nil || o.EndPoint == nil {
		return nil, false
	}
	return o.EndPoint, true
}

// HasEndPoint returns a boolean if a field has been set.
func (o *BTExportModelArcEdgeGeometry1257) HasEndPoint() bool {
	if o != nil && o.EndPoint != nil {
		return true
	}

	return false
}

// SetEndPoint gets a reference to the given BTVector3d389 and assigns it to the EndPoint field.
func (o *BTExportModelArcEdgeGeometry1257) SetEndPoint(v BTVector3d389) {
	o.EndPoint = &v
}

// GetEndVector returns the EndVector field value if set, zero value otherwise.
func (o *BTExportModelArcEdgeGeometry1257) GetEndVector() BTVector3d389 {
	if o == nil || o.EndVector == nil {
		var ret BTVector3d389
		return ret
	}
	return *o.EndVector
}

// GetEndVectorOk returns a tuple with the EndVector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExportModelArcEdgeGeometry1257) GetEndVectorOk() (*BTVector3d389, bool) {
	if o == nil || o.EndVector == nil {
		return nil, false
	}
	return o.EndVector, true
}

// HasEndVector returns a boolean if a field has been set.
func (o *BTExportModelArcEdgeGeometry1257) HasEndVector() bool {
	if o != nil && o.EndVector != nil {
		return true
	}

	return false
}

// SetEndVector gets a reference to the given BTVector3d389 and assigns it to the EndVector field.
func (o *BTExportModelArcEdgeGeometry1257) SetEndVector(v BTVector3d389) {
	o.EndVector = &v
}

// GetLength returns the Length field value if set, zero value otherwise.
func (o *BTExportModelArcEdgeGeometry1257) GetLength() float64 {
	if o == nil || o.Length == nil {
		var ret float64
		return ret
	}
	return *o.Length
}

// GetLengthOk returns a tuple with the Length field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExportModelArcEdgeGeometry1257) GetLengthOk() (*float64, bool) {
	if o == nil || o.Length == nil {
		return nil, false
	}
	return o.Length, true
}

// HasLength returns a boolean if a field has been set.
func (o *BTExportModelArcEdgeGeometry1257) HasLength() bool {
	if o != nil && o.Length != nil {
		return true
	}

	return false
}

// SetLength gets a reference to the given float64 and assigns it to the Length field.
func (o *BTExportModelArcEdgeGeometry1257) SetLength(v float64) {
	o.Length = &v
}

// GetMidPoint returns the MidPoint field value if set, zero value otherwise.
func (o *BTExportModelArcEdgeGeometry1257) GetMidPoint() BTVector3d389 {
	if o == nil || o.MidPoint == nil {
		var ret BTVector3d389
		return ret
	}
	return *o.MidPoint
}

// GetMidPointOk returns a tuple with the MidPoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExportModelArcEdgeGeometry1257) GetMidPointOk() (*BTVector3d389, bool) {
	if o == nil || o.MidPoint == nil {
		return nil, false
	}
	return o.MidPoint, true
}

// HasMidPoint returns a boolean if a field has been set.
func (o *BTExportModelArcEdgeGeometry1257) HasMidPoint() bool {
	if o != nil && o.MidPoint != nil {
		return true
	}

	return false
}

// SetMidPoint gets a reference to the given BTVector3d389 and assigns it to the MidPoint field.
func (o *BTExportModelArcEdgeGeometry1257) SetMidPoint(v BTVector3d389) {
	o.MidPoint = &v
}

// GetQuarterPoint returns the QuarterPoint field value if set, zero value otherwise.
func (o *BTExportModelArcEdgeGeometry1257) GetQuarterPoint() BTVector3d389 {
	if o == nil || o.QuarterPoint == nil {
		var ret BTVector3d389
		return ret
	}
	return *o.QuarterPoint
}

// GetQuarterPointOk returns a tuple with the QuarterPoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExportModelArcEdgeGeometry1257) GetQuarterPointOk() (*BTVector3d389, bool) {
	if o == nil || o.QuarterPoint == nil {
		return nil, false
	}
	return o.QuarterPoint, true
}

// HasQuarterPoint returns a boolean if a field has been set.
func (o *BTExportModelArcEdgeGeometry1257) HasQuarterPoint() bool {
	if o != nil && o.QuarterPoint != nil {
		return true
	}

	return false
}

// SetQuarterPoint gets a reference to the given BTVector3d389 and assigns it to the QuarterPoint field.
func (o *BTExportModelArcEdgeGeometry1257) SetQuarterPoint(v BTVector3d389) {
	o.QuarterPoint = &v
}

// GetStartPoint returns the StartPoint field value if set, zero value otherwise.
func (o *BTExportModelArcEdgeGeometry1257) GetStartPoint() BTVector3d389 {
	if o == nil || o.StartPoint == nil {
		var ret BTVector3d389
		return ret
	}
	return *o.StartPoint
}

// GetStartPointOk returns a tuple with the StartPoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExportModelArcEdgeGeometry1257) GetStartPointOk() (*BTVector3d389, bool) {
	if o == nil || o.StartPoint == nil {
		return nil, false
	}
	return o.StartPoint, true
}

// HasStartPoint returns a boolean if a field has been set.
func (o *BTExportModelArcEdgeGeometry1257) HasStartPoint() bool {
	if o != nil && o.StartPoint != nil {
		return true
	}

	return false
}

// SetStartPoint gets a reference to the given BTVector3d389 and assigns it to the StartPoint field.
func (o *BTExportModelArcEdgeGeometry1257) SetStartPoint(v BTVector3d389) {
	o.StartPoint = &v
}

// GetStartVector returns the StartVector field value if set, zero value otherwise.
func (o *BTExportModelArcEdgeGeometry1257) GetStartVector() BTVector3d389 {
	if o == nil || o.StartVector == nil {
		var ret BTVector3d389
		return ret
	}
	return *o.StartVector
}

// GetStartVectorOk returns a tuple with the StartVector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExportModelArcEdgeGeometry1257) GetStartVectorOk() (*BTVector3d389, bool) {
	if o == nil || o.StartVector == nil {
		return nil, false
	}
	return o.StartVector, true
}

// HasStartVector returns a boolean if a field has been set.
func (o *BTExportModelArcEdgeGeometry1257) HasStartVector() bool {
	if o != nil && o.StartVector != nil {
		return true
	}

	return false
}

// SetStartVector gets a reference to the given BTVector3d389 and assigns it to the StartVector field.
func (o *BTExportModelArcEdgeGeometry1257) SetStartVector(v BTVector3d389) {
	o.StartVector = &v
}

// GetArcIsClockwise returns the ArcIsClockwise field value if set, zero value otherwise.
func (o *BTExportModelArcEdgeGeometry1257) GetArcIsClockwise() bool {
	if o == nil || o.ArcIsClockwise == nil {
		var ret bool
		return ret
	}
	return *o.ArcIsClockwise
}

// GetArcIsClockwiseOk returns a tuple with the ArcIsClockwise field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExportModelArcEdgeGeometry1257) GetArcIsClockwiseOk() (*bool, bool) {
	if o == nil || o.ArcIsClockwise == nil {
		return nil, false
	}
	return o.ArcIsClockwise, true
}

// HasArcIsClockwise returns a boolean if a field has been set.
func (o *BTExportModelArcEdgeGeometry1257) HasArcIsClockwise() bool {
	if o != nil && o.ArcIsClockwise != nil {
		return true
	}

	return false
}

// SetArcIsClockwise gets a reference to the given bool and assigns it to the ArcIsClockwise field.
func (o *BTExportModelArcEdgeGeometry1257) SetArcIsClockwise(v bool) {
	o.ArcIsClockwise = &v
}

// GetArcSweep returns the ArcSweep field value if set, zero value otherwise.
func (o *BTExportModelArcEdgeGeometry1257) GetArcSweep() float64 {
	if o == nil || o.ArcSweep == nil {
		var ret float64
		return ret
	}
	return *o.ArcSweep
}

// GetArcSweepOk returns a tuple with the ArcSweep field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExportModelArcEdgeGeometry1257) GetArcSweepOk() (*float64, bool) {
	if o == nil || o.ArcSweep == nil {
		return nil, false
	}
	return o.ArcSweep, true
}

// HasArcSweep returns a boolean if a field has been set.
func (o *BTExportModelArcEdgeGeometry1257) HasArcSweep() bool {
	if o != nil && o.ArcSweep != nil {
		return true
	}

	return false
}

// SetArcSweep gets a reference to the given float64 and assigns it to the ArcSweep field.
func (o *BTExportModelArcEdgeGeometry1257) SetArcSweep(v float64) {
	o.ArcSweep = &v
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTExportModelArcEdgeGeometry1257) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExportModelArcEdgeGeometry1257) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTExportModelArcEdgeGeometry1257) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTExportModelArcEdgeGeometry1257) SetBtType(v string) {
	o.BtType = &v
}

func (o BTExportModelArcEdgeGeometry1257) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EndPoint != nil {
		toSerialize["endPoint"] = o.EndPoint
	}
	if o.EndVector != nil {
		toSerialize["endVector"] = o.EndVector
	}
	if o.Length != nil {
		toSerialize["length"] = o.Length
	}
	if o.MidPoint != nil {
		toSerialize["midPoint"] = o.MidPoint
	}
	if o.QuarterPoint != nil {
		toSerialize["quarterPoint"] = o.QuarterPoint
	}
	if o.StartPoint != nil {
		toSerialize["startPoint"] = o.StartPoint
	}
	if o.StartVector != nil {
		toSerialize["startVector"] = o.StartVector
	}
	if o.ArcIsClockwise != nil {
		toSerialize["arcIsClockwise"] = o.ArcIsClockwise
	}
	if o.ArcSweep != nil {
		toSerialize["arcSweep"] = o.ArcSweep
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	return json.Marshal(toSerialize)
}

type NullableBTExportModelArcEdgeGeometry1257 struct {
	value *BTExportModelArcEdgeGeometry1257
	isSet bool
}

func (v NullableBTExportModelArcEdgeGeometry1257) Get() *BTExportModelArcEdgeGeometry1257 {
	return v.value
}

func (v *NullableBTExportModelArcEdgeGeometry1257) Set(val *BTExportModelArcEdgeGeometry1257) {
	v.value = val
	v.isSet = true
}

func (v NullableBTExportModelArcEdgeGeometry1257) IsSet() bool {
	return v.isSet
}

func (v *NullableBTExportModelArcEdgeGeometry1257) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTExportModelArcEdgeGeometry1257(val *BTExportModelArcEdgeGeometry1257) *NullableBTExportModelArcEdgeGeometry1257 {
	return &NullableBTExportModelArcEdgeGeometry1257{value: val, isSet: true}
}

func (v NullableBTExportModelArcEdgeGeometry1257) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTExportModelArcEdgeGeometry1257) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
