/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.153.6415-48a6b2252b8c
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTEllipseDiameterDimensionDisplayData1301 struct for BTEllipseDiameterDimensionDisplayData1301
type BTEllipseDiameterDimensionDisplayData1301 struct {
	CoordinateSystem     *BTMatrix3x3340 `json:"coordinateSystem,omitempty"`
	FeatureId            *string         `json:"featureId,omitempty"`
	HasMaximumLimit_     *bool           `json:"hasMaximumLimit,omitempty"`
	HasMinimumLimit_     *bool           `json:"hasMinimumLimit,omitempty"`
	Id                   *string         `json:"id,omitempty"`
	IsAssociatedWithFlat *bool           `json:"isAssociatedWithFlat,omitempty"`
	IsDriven             *bool           `json:"isDriven,omitempty"`
	IsOverDefined        *bool           `json:"isOverDefined,omitempty"`
	MaximumLimit         *float64        `json:"maximumLimit,omitempty"`
	MinimumLimit         *float64        `json:"minimumLimit,omitempty"`
	ParameterId          *string         `json:"parameterId,omitempty"`
	PlaneMatrix          *BTBSMatrix386  `json:"planeMatrix,omitempty"`
	Value                *float64        `json:"value,omitempty"`
	BtType               *string         `json:"btType,omitempty"`
	PositionX            *float64        `json:"positionX,omitempty"`
	PositionY            *float64        `json:"positionY,omitempty"`
	WitnessEndPoint0X    *float64        `json:"witnessEndPoint0X,omitempty"`
	WitnessEndPoint0Y    *float64        `json:"witnessEndPoint0Y,omitempty"`
	WitnessEndPoint1X    *float64        `json:"witnessEndPoint1X,omitempty"`
	WitnessEndPoint1Y    *float64        `json:"witnessEndPoint1Y,omitempty"`
}

// NewBTEllipseDiameterDimensionDisplayData1301 instantiates a new BTEllipseDiameterDimensionDisplayData1301 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTEllipseDiameterDimensionDisplayData1301() *BTEllipseDiameterDimensionDisplayData1301 {
	this := BTEllipseDiameterDimensionDisplayData1301{}
	return &this
}

// NewBTEllipseDiameterDimensionDisplayData1301WithDefaults instantiates a new BTEllipseDiameterDimensionDisplayData1301 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTEllipseDiameterDimensionDisplayData1301WithDefaults() *BTEllipseDiameterDimensionDisplayData1301 {
	this := BTEllipseDiameterDimensionDisplayData1301{}
	return &this
}

// GetCoordinateSystem returns the CoordinateSystem field value if set, zero value otherwise.
func (o *BTEllipseDiameterDimensionDisplayData1301) GetCoordinateSystem() BTMatrix3x3340 {
	if o == nil || o.CoordinateSystem == nil {
		var ret BTMatrix3x3340
		return ret
	}
	return *o.CoordinateSystem
}

// GetCoordinateSystemOk returns a tuple with the CoordinateSystem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTEllipseDiameterDimensionDisplayData1301) GetCoordinateSystemOk() (*BTMatrix3x3340, bool) {
	if o == nil || o.CoordinateSystem == nil {
		return nil, false
	}
	return o.CoordinateSystem, true
}

// HasCoordinateSystem returns a boolean if a field has been set.
func (o *BTEllipseDiameterDimensionDisplayData1301) HasCoordinateSystem() bool {
	if o != nil && o.CoordinateSystem != nil {
		return true
	}

	return false
}

// SetCoordinateSystem gets a reference to the given BTMatrix3x3340 and assigns it to the CoordinateSystem field.
func (o *BTEllipseDiameterDimensionDisplayData1301) SetCoordinateSystem(v BTMatrix3x3340) {
	o.CoordinateSystem = &v
}

// GetFeatureId returns the FeatureId field value if set, zero value otherwise.
func (o *BTEllipseDiameterDimensionDisplayData1301) GetFeatureId() string {
	if o == nil || o.FeatureId == nil {
		var ret string
		return ret
	}
	return *o.FeatureId
}

// GetFeatureIdOk returns a tuple with the FeatureId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTEllipseDiameterDimensionDisplayData1301) GetFeatureIdOk() (*string, bool) {
	if o == nil || o.FeatureId == nil {
		return nil, false
	}
	return o.FeatureId, true
}

// HasFeatureId returns a boolean if a field has been set.
func (o *BTEllipseDiameterDimensionDisplayData1301) HasFeatureId() bool {
	if o != nil && o.FeatureId != nil {
		return true
	}

	return false
}

// SetFeatureId gets a reference to the given string and assigns it to the FeatureId field.
func (o *BTEllipseDiameterDimensionDisplayData1301) SetFeatureId(v string) {
	o.FeatureId = &v
}

// GetHasMaximumLimit_ returns the HasMaximumLimit_ field value if set, zero value otherwise.
func (o *BTEllipseDiameterDimensionDisplayData1301) GetHasMaximumLimit_() bool {
	if o == nil || o.HasMaximumLimit_ == nil {
		var ret bool
		return ret
	}
	return *o.HasMaximumLimit_
}

// GetHasMaximumLimit_Ok returns a tuple with the HasMaximumLimit_ field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTEllipseDiameterDimensionDisplayData1301) GetHasMaximumLimit_Ok() (*bool, bool) {
	if o == nil || o.HasMaximumLimit_ == nil {
		return nil, false
	}
	return o.HasMaximumLimit_, true
}

// HasHasMaximumLimit_ returns a boolean if a field has been set.
func (o *BTEllipseDiameterDimensionDisplayData1301) HasHasMaximumLimit_() bool {
	if o != nil && o.HasMaximumLimit_ != nil {
		return true
	}

	return false
}

// SetHasMaximumLimit_ gets a reference to the given bool and assigns it to the HasMaximumLimit_ field.
func (o *BTEllipseDiameterDimensionDisplayData1301) SetHasMaximumLimit_(v bool) {
	o.HasMaximumLimit_ = &v
}

// GetHasMinimumLimit_ returns the HasMinimumLimit_ field value if set, zero value otherwise.
func (o *BTEllipseDiameterDimensionDisplayData1301) GetHasMinimumLimit_() bool {
	if o == nil || o.HasMinimumLimit_ == nil {
		var ret bool
		return ret
	}
	return *o.HasMinimumLimit_
}

// GetHasMinimumLimit_Ok returns a tuple with the HasMinimumLimit_ field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTEllipseDiameterDimensionDisplayData1301) GetHasMinimumLimit_Ok() (*bool, bool) {
	if o == nil || o.HasMinimumLimit_ == nil {
		return nil, false
	}
	return o.HasMinimumLimit_, true
}

// HasHasMinimumLimit_ returns a boolean if a field has been set.
func (o *BTEllipseDiameterDimensionDisplayData1301) HasHasMinimumLimit_() bool {
	if o != nil && o.HasMinimumLimit_ != nil {
		return true
	}

	return false
}

// SetHasMinimumLimit_ gets a reference to the given bool and assigns it to the HasMinimumLimit_ field.
func (o *BTEllipseDiameterDimensionDisplayData1301) SetHasMinimumLimit_(v bool) {
	o.HasMinimumLimit_ = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BTEllipseDiameterDimensionDisplayData1301) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTEllipseDiameterDimensionDisplayData1301) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BTEllipseDiameterDimensionDisplayData1301) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BTEllipseDiameterDimensionDisplayData1301) SetId(v string) {
	o.Id = &v
}

// GetIsAssociatedWithFlat returns the IsAssociatedWithFlat field value if set, zero value otherwise.
func (o *BTEllipseDiameterDimensionDisplayData1301) GetIsAssociatedWithFlat() bool {
	if o == nil || o.IsAssociatedWithFlat == nil {
		var ret bool
		return ret
	}
	return *o.IsAssociatedWithFlat
}

// GetIsAssociatedWithFlatOk returns a tuple with the IsAssociatedWithFlat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTEllipseDiameterDimensionDisplayData1301) GetIsAssociatedWithFlatOk() (*bool, bool) {
	if o == nil || o.IsAssociatedWithFlat == nil {
		return nil, false
	}
	return o.IsAssociatedWithFlat, true
}

// HasIsAssociatedWithFlat returns a boolean if a field has been set.
func (o *BTEllipseDiameterDimensionDisplayData1301) HasIsAssociatedWithFlat() bool {
	if o != nil && o.IsAssociatedWithFlat != nil {
		return true
	}

	return false
}

// SetIsAssociatedWithFlat gets a reference to the given bool and assigns it to the IsAssociatedWithFlat field.
func (o *BTEllipseDiameterDimensionDisplayData1301) SetIsAssociatedWithFlat(v bool) {
	o.IsAssociatedWithFlat = &v
}

// GetIsDriven returns the IsDriven field value if set, zero value otherwise.
func (o *BTEllipseDiameterDimensionDisplayData1301) GetIsDriven() bool {
	if o == nil || o.IsDriven == nil {
		var ret bool
		return ret
	}
	return *o.IsDriven
}

// GetIsDrivenOk returns a tuple with the IsDriven field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTEllipseDiameterDimensionDisplayData1301) GetIsDrivenOk() (*bool, bool) {
	if o == nil || o.IsDriven == nil {
		return nil, false
	}
	return o.IsDriven, true
}

// HasIsDriven returns a boolean if a field has been set.
func (o *BTEllipseDiameterDimensionDisplayData1301) HasIsDriven() bool {
	if o != nil && o.IsDriven != nil {
		return true
	}

	return false
}

// SetIsDriven gets a reference to the given bool and assigns it to the IsDriven field.
func (o *BTEllipseDiameterDimensionDisplayData1301) SetIsDriven(v bool) {
	o.IsDriven = &v
}

// GetIsOverDefined returns the IsOverDefined field value if set, zero value otherwise.
func (o *BTEllipseDiameterDimensionDisplayData1301) GetIsOverDefined() bool {
	if o == nil || o.IsOverDefined == nil {
		var ret bool
		return ret
	}
	return *o.IsOverDefined
}

// GetIsOverDefinedOk returns a tuple with the IsOverDefined field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTEllipseDiameterDimensionDisplayData1301) GetIsOverDefinedOk() (*bool, bool) {
	if o == nil || o.IsOverDefined == nil {
		return nil, false
	}
	return o.IsOverDefined, true
}

// HasIsOverDefined returns a boolean if a field has been set.
func (o *BTEllipseDiameterDimensionDisplayData1301) HasIsOverDefined() bool {
	if o != nil && o.IsOverDefined != nil {
		return true
	}

	return false
}

// SetIsOverDefined gets a reference to the given bool and assigns it to the IsOverDefined field.
func (o *BTEllipseDiameterDimensionDisplayData1301) SetIsOverDefined(v bool) {
	o.IsOverDefined = &v
}

// GetMaximumLimit returns the MaximumLimit field value if set, zero value otherwise.
func (o *BTEllipseDiameterDimensionDisplayData1301) GetMaximumLimit() float64 {
	if o == nil || o.MaximumLimit == nil {
		var ret float64
		return ret
	}
	return *o.MaximumLimit
}

// GetMaximumLimitOk returns a tuple with the MaximumLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTEllipseDiameterDimensionDisplayData1301) GetMaximumLimitOk() (*float64, bool) {
	if o == nil || o.MaximumLimit == nil {
		return nil, false
	}
	return o.MaximumLimit, true
}

// HasMaximumLimit returns a boolean if a field has been set.
func (o *BTEllipseDiameterDimensionDisplayData1301) HasMaximumLimit() bool {
	if o != nil && o.MaximumLimit != nil {
		return true
	}

	return false
}

// SetMaximumLimit gets a reference to the given float64 and assigns it to the MaximumLimit field.
func (o *BTEllipseDiameterDimensionDisplayData1301) SetMaximumLimit(v float64) {
	o.MaximumLimit = &v
}

// GetMinimumLimit returns the MinimumLimit field value if set, zero value otherwise.
func (o *BTEllipseDiameterDimensionDisplayData1301) GetMinimumLimit() float64 {
	if o == nil || o.MinimumLimit == nil {
		var ret float64
		return ret
	}
	return *o.MinimumLimit
}

// GetMinimumLimitOk returns a tuple with the MinimumLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTEllipseDiameterDimensionDisplayData1301) GetMinimumLimitOk() (*float64, bool) {
	if o == nil || o.MinimumLimit == nil {
		return nil, false
	}
	return o.MinimumLimit, true
}

// HasMinimumLimit returns a boolean if a field has been set.
func (o *BTEllipseDiameterDimensionDisplayData1301) HasMinimumLimit() bool {
	if o != nil && o.MinimumLimit != nil {
		return true
	}

	return false
}

// SetMinimumLimit gets a reference to the given float64 and assigns it to the MinimumLimit field.
func (o *BTEllipseDiameterDimensionDisplayData1301) SetMinimumLimit(v float64) {
	o.MinimumLimit = &v
}

// GetParameterId returns the ParameterId field value if set, zero value otherwise.
func (o *BTEllipseDiameterDimensionDisplayData1301) GetParameterId() string {
	if o == nil || o.ParameterId == nil {
		var ret string
		return ret
	}
	return *o.ParameterId
}

// GetParameterIdOk returns a tuple with the ParameterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTEllipseDiameterDimensionDisplayData1301) GetParameterIdOk() (*string, bool) {
	if o == nil || o.ParameterId == nil {
		return nil, false
	}
	return o.ParameterId, true
}

// HasParameterId returns a boolean if a field has been set.
func (o *BTEllipseDiameterDimensionDisplayData1301) HasParameterId() bool {
	if o != nil && o.ParameterId != nil {
		return true
	}

	return false
}

// SetParameterId gets a reference to the given string and assigns it to the ParameterId field.
func (o *BTEllipseDiameterDimensionDisplayData1301) SetParameterId(v string) {
	o.ParameterId = &v
}

// GetPlaneMatrix returns the PlaneMatrix field value if set, zero value otherwise.
func (o *BTEllipseDiameterDimensionDisplayData1301) GetPlaneMatrix() BTBSMatrix386 {
	if o == nil || o.PlaneMatrix == nil {
		var ret BTBSMatrix386
		return ret
	}
	return *o.PlaneMatrix
}

// GetPlaneMatrixOk returns a tuple with the PlaneMatrix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTEllipseDiameterDimensionDisplayData1301) GetPlaneMatrixOk() (*BTBSMatrix386, bool) {
	if o == nil || o.PlaneMatrix == nil {
		return nil, false
	}
	return o.PlaneMatrix, true
}

// HasPlaneMatrix returns a boolean if a field has been set.
func (o *BTEllipseDiameterDimensionDisplayData1301) HasPlaneMatrix() bool {
	if o != nil && o.PlaneMatrix != nil {
		return true
	}

	return false
}

// SetPlaneMatrix gets a reference to the given BTBSMatrix386 and assigns it to the PlaneMatrix field.
func (o *BTEllipseDiameterDimensionDisplayData1301) SetPlaneMatrix(v BTBSMatrix386) {
	o.PlaneMatrix = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *BTEllipseDiameterDimensionDisplayData1301) GetValue() float64 {
	if o == nil || o.Value == nil {
		var ret float64
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTEllipseDiameterDimensionDisplayData1301) GetValueOk() (*float64, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *BTEllipseDiameterDimensionDisplayData1301) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given float64 and assigns it to the Value field.
func (o *BTEllipseDiameterDimensionDisplayData1301) SetValue(v float64) {
	o.Value = &v
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTEllipseDiameterDimensionDisplayData1301) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTEllipseDiameterDimensionDisplayData1301) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTEllipseDiameterDimensionDisplayData1301) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTEllipseDiameterDimensionDisplayData1301) SetBtType(v string) {
	o.BtType = &v
}

// GetPositionX returns the PositionX field value if set, zero value otherwise.
func (o *BTEllipseDiameterDimensionDisplayData1301) GetPositionX() float64 {
	if o == nil || o.PositionX == nil {
		var ret float64
		return ret
	}
	return *o.PositionX
}

// GetPositionXOk returns a tuple with the PositionX field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTEllipseDiameterDimensionDisplayData1301) GetPositionXOk() (*float64, bool) {
	if o == nil || o.PositionX == nil {
		return nil, false
	}
	return o.PositionX, true
}

// HasPositionX returns a boolean if a field has been set.
func (o *BTEllipseDiameterDimensionDisplayData1301) HasPositionX() bool {
	if o != nil && o.PositionX != nil {
		return true
	}

	return false
}

// SetPositionX gets a reference to the given float64 and assigns it to the PositionX field.
func (o *BTEllipseDiameterDimensionDisplayData1301) SetPositionX(v float64) {
	o.PositionX = &v
}

// GetPositionY returns the PositionY field value if set, zero value otherwise.
func (o *BTEllipseDiameterDimensionDisplayData1301) GetPositionY() float64 {
	if o == nil || o.PositionY == nil {
		var ret float64
		return ret
	}
	return *o.PositionY
}

// GetPositionYOk returns a tuple with the PositionY field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTEllipseDiameterDimensionDisplayData1301) GetPositionYOk() (*float64, bool) {
	if o == nil || o.PositionY == nil {
		return nil, false
	}
	return o.PositionY, true
}

// HasPositionY returns a boolean if a field has been set.
func (o *BTEllipseDiameterDimensionDisplayData1301) HasPositionY() bool {
	if o != nil && o.PositionY != nil {
		return true
	}

	return false
}

// SetPositionY gets a reference to the given float64 and assigns it to the PositionY field.
func (o *BTEllipseDiameterDimensionDisplayData1301) SetPositionY(v float64) {
	o.PositionY = &v
}

// GetWitnessEndPoint0X returns the WitnessEndPoint0X field value if set, zero value otherwise.
func (o *BTEllipseDiameterDimensionDisplayData1301) GetWitnessEndPoint0X() float64 {
	if o == nil || o.WitnessEndPoint0X == nil {
		var ret float64
		return ret
	}
	return *o.WitnessEndPoint0X
}

// GetWitnessEndPoint0XOk returns a tuple with the WitnessEndPoint0X field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTEllipseDiameterDimensionDisplayData1301) GetWitnessEndPoint0XOk() (*float64, bool) {
	if o == nil || o.WitnessEndPoint0X == nil {
		return nil, false
	}
	return o.WitnessEndPoint0X, true
}

// HasWitnessEndPoint0X returns a boolean if a field has been set.
func (o *BTEllipseDiameterDimensionDisplayData1301) HasWitnessEndPoint0X() bool {
	if o != nil && o.WitnessEndPoint0X != nil {
		return true
	}

	return false
}

// SetWitnessEndPoint0X gets a reference to the given float64 and assigns it to the WitnessEndPoint0X field.
func (o *BTEllipseDiameterDimensionDisplayData1301) SetWitnessEndPoint0X(v float64) {
	o.WitnessEndPoint0X = &v
}

// GetWitnessEndPoint0Y returns the WitnessEndPoint0Y field value if set, zero value otherwise.
func (o *BTEllipseDiameterDimensionDisplayData1301) GetWitnessEndPoint0Y() float64 {
	if o == nil || o.WitnessEndPoint0Y == nil {
		var ret float64
		return ret
	}
	return *o.WitnessEndPoint0Y
}

// GetWitnessEndPoint0YOk returns a tuple with the WitnessEndPoint0Y field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTEllipseDiameterDimensionDisplayData1301) GetWitnessEndPoint0YOk() (*float64, bool) {
	if o == nil || o.WitnessEndPoint0Y == nil {
		return nil, false
	}
	return o.WitnessEndPoint0Y, true
}

// HasWitnessEndPoint0Y returns a boolean if a field has been set.
func (o *BTEllipseDiameterDimensionDisplayData1301) HasWitnessEndPoint0Y() bool {
	if o != nil && o.WitnessEndPoint0Y != nil {
		return true
	}

	return false
}

// SetWitnessEndPoint0Y gets a reference to the given float64 and assigns it to the WitnessEndPoint0Y field.
func (o *BTEllipseDiameterDimensionDisplayData1301) SetWitnessEndPoint0Y(v float64) {
	o.WitnessEndPoint0Y = &v
}

// GetWitnessEndPoint1X returns the WitnessEndPoint1X field value if set, zero value otherwise.
func (o *BTEllipseDiameterDimensionDisplayData1301) GetWitnessEndPoint1X() float64 {
	if o == nil || o.WitnessEndPoint1X == nil {
		var ret float64
		return ret
	}
	return *o.WitnessEndPoint1X
}

// GetWitnessEndPoint1XOk returns a tuple with the WitnessEndPoint1X field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTEllipseDiameterDimensionDisplayData1301) GetWitnessEndPoint1XOk() (*float64, bool) {
	if o == nil || o.WitnessEndPoint1X == nil {
		return nil, false
	}
	return o.WitnessEndPoint1X, true
}

// HasWitnessEndPoint1X returns a boolean if a field has been set.
func (o *BTEllipseDiameterDimensionDisplayData1301) HasWitnessEndPoint1X() bool {
	if o != nil && o.WitnessEndPoint1X != nil {
		return true
	}

	return false
}

// SetWitnessEndPoint1X gets a reference to the given float64 and assigns it to the WitnessEndPoint1X field.
func (o *BTEllipseDiameterDimensionDisplayData1301) SetWitnessEndPoint1X(v float64) {
	o.WitnessEndPoint1X = &v
}

// GetWitnessEndPoint1Y returns the WitnessEndPoint1Y field value if set, zero value otherwise.
func (o *BTEllipseDiameterDimensionDisplayData1301) GetWitnessEndPoint1Y() float64 {
	if o == nil || o.WitnessEndPoint1Y == nil {
		var ret float64
		return ret
	}
	return *o.WitnessEndPoint1Y
}

// GetWitnessEndPoint1YOk returns a tuple with the WitnessEndPoint1Y field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTEllipseDiameterDimensionDisplayData1301) GetWitnessEndPoint1YOk() (*float64, bool) {
	if o == nil || o.WitnessEndPoint1Y == nil {
		return nil, false
	}
	return o.WitnessEndPoint1Y, true
}

// HasWitnessEndPoint1Y returns a boolean if a field has been set.
func (o *BTEllipseDiameterDimensionDisplayData1301) HasWitnessEndPoint1Y() bool {
	if o != nil && o.WitnessEndPoint1Y != nil {
		return true
	}

	return false
}

// SetWitnessEndPoint1Y gets a reference to the given float64 and assigns it to the WitnessEndPoint1Y field.
func (o *BTEllipseDiameterDimensionDisplayData1301) SetWitnessEndPoint1Y(v float64) {
	o.WitnessEndPoint1Y = &v
}

func (o BTEllipseDiameterDimensionDisplayData1301) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CoordinateSystem != nil {
		toSerialize["coordinateSystem"] = o.CoordinateSystem
	}
	if o.FeatureId != nil {
		toSerialize["featureId"] = o.FeatureId
	}
	if o.HasMaximumLimit_ != nil {
		toSerialize["hasMaximumLimit"] = o.HasMaximumLimit_
	}
	if o.HasMinimumLimit_ != nil {
		toSerialize["hasMinimumLimit"] = o.HasMinimumLimit_
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.IsAssociatedWithFlat != nil {
		toSerialize["isAssociatedWithFlat"] = o.IsAssociatedWithFlat
	}
	if o.IsDriven != nil {
		toSerialize["isDriven"] = o.IsDriven
	}
	if o.IsOverDefined != nil {
		toSerialize["isOverDefined"] = o.IsOverDefined
	}
	if o.MaximumLimit != nil {
		toSerialize["maximumLimit"] = o.MaximumLimit
	}
	if o.MinimumLimit != nil {
		toSerialize["minimumLimit"] = o.MinimumLimit
	}
	if o.ParameterId != nil {
		toSerialize["parameterId"] = o.ParameterId
	}
	if o.PlaneMatrix != nil {
		toSerialize["planeMatrix"] = o.PlaneMatrix
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.PositionX != nil {
		toSerialize["positionX"] = o.PositionX
	}
	if o.PositionY != nil {
		toSerialize["positionY"] = o.PositionY
	}
	if o.WitnessEndPoint0X != nil {
		toSerialize["witnessEndPoint0X"] = o.WitnessEndPoint0X
	}
	if o.WitnessEndPoint0Y != nil {
		toSerialize["witnessEndPoint0Y"] = o.WitnessEndPoint0Y
	}
	if o.WitnessEndPoint1X != nil {
		toSerialize["witnessEndPoint1X"] = o.WitnessEndPoint1X
	}
	if o.WitnessEndPoint1Y != nil {
		toSerialize["witnessEndPoint1Y"] = o.WitnessEndPoint1Y
	}
	return json.Marshal(toSerialize)
}

type NullableBTEllipseDiameterDimensionDisplayData1301 struct {
	value *BTEllipseDiameterDimensionDisplayData1301
	isSet bool
}

func (v NullableBTEllipseDiameterDimensionDisplayData1301) Get() *BTEllipseDiameterDimensionDisplayData1301 {
	return v.value
}

func (v *NullableBTEllipseDiameterDimensionDisplayData1301) Set(val *BTEllipseDiameterDimensionDisplayData1301) {
	v.value = val
	v.isSet = true
}

func (v NullableBTEllipseDiameterDimensionDisplayData1301) IsSet() bool {
	return v.isSet
}

func (v *NullableBTEllipseDiameterDimensionDisplayData1301) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTEllipseDiameterDimensionDisplayData1301(val *BTEllipseDiameterDimensionDisplayData1301) *NullableBTEllipseDiameterDimensionDisplayData1301 {
	return &NullableBTEllipseDiameterDimensionDisplayData1301{value: val, isSet: true}
}

func (v NullableBTEllipseDiameterDimensionDisplayData1301) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTEllipseDiameterDimensionDisplayData1301) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}