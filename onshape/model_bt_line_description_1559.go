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
)

// BTLineDescription1559 struct for BTLineDescription1559
type BTLineDescription1559 struct {
	BtType                    *string           `json:"btType,omitempty"`
	Direction                 *BTVector3d389    `json:"direction,omitempty"`
	DirectionOrientedWithFace *BTVector3d389    `json:"directionOrientedWithFace,omitempty"`
	Origin                    *BTVector3d389    `json:"origin,omitempty"`
	Type                      *GBTCurveTypeEnum `json:"type,omitempty"`
}

// NewBTLineDescription1559 instantiates a new BTLineDescription1559 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTLineDescription1559() *BTLineDescription1559 {
	this := BTLineDescription1559{}
	return &this
}

// NewBTLineDescription1559WithDefaults instantiates a new BTLineDescription1559 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTLineDescription1559WithDefaults() *BTLineDescription1559 {
	this := BTLineDescription1559{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTLineDescription1559) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTLineDescription1559) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTLineDescription1559) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTLineDescription1559) SetBtType(v string) {
	o.BtType = &v
}

// GetDirection returns the Direction field value if set, zero value otherwise.
func (o *BTLineDescription1559) GetDirection() BTVector3d389 {
	if o == nil || o.Direction == nil {
		var ret BTVector3d389
		return ret
	}
	return *o.Direction
}

// GetDirectionOk returns a tuple with the Direction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTLineDescription1559) GetDirectionOk() (*BTVector3d389, bool) {
	if o == nil || o.Direction == nil {
		return nil, false
	}
	return o.Direction, true
}

// HasDirection returns a boolean if a field has been set.
func (o *BTLineDescription1559) HasDirection() bool {
	if o != nil && o.Direction != nil {
		return true
	}

	return false
}

// SetDirection gets a reference to the given BTVector3d389 and assigns it to the Direction field.
func (o *BTLineDescription1559) SetDirection(v BTVector3d389) {
	o.Direction = &v
}

// GetDirectionOrientedWithFace returns the DirectionOrientedWithFace field value if set, zero value otherwise.
func (o *BTLineDescription1559) GetDirectionOrientedWithFace() BTVector3d389 {
	if o == nil || o.DirectionOrientedWithFace == nil {
		var ret BTVector3d389
		return ret
	}
	return *o.DirectionOrientedWithFace
}

// GetDirectionOrientedWithFaceOk returns a tuple with the DirectionOrientedWithFace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTLineDescription1559) GetDirectionOrientedWithFaceOk() (*BTVector3d389, bool) {
	if o == nil || o.DirectionOrientedWithFace == nil {
		return nil, false
	}
	return o.DirectionOrientedWithFace, true
}

// HasDirectionOrientedWithFace returns a boolean if a field has been set.
func (o *BTLineDescription1559) HasDirectionOrientedWithFace() bool {
	if o != nil && o.DirectionOrientedWithFace != nil {
		return true
	}

	return false
}

// SetDirectionOrientedWithFace gets a reference to the given BTVector3d389 and assigns it to the DirectionOrientedWithFace field.
func (o *BTLineDescription1559) SetDirectionOrientedWithFace(v BTVector3d389) {
	o.DirectionOrientedWithFace = &v
}

// GetOrigin returns the Origin field value if set, zero value otherwise.
func (o *BTLineDescription1559) GetOrigin() BTVector3d389 {
	if o == nil || o.Origin == nil {
		var ret BTVector3d389
		return ret
	}
	return *o.Origin
}

// GetOriginOk returns a tuple with the Origin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTLineDescription1559) GetOriginOk() (*BTVector3d389, bool) {
	if o == nil || o.Origin == nil {
		return nil, false
	}
	return o.Origin, true
}

// HasOrigin returns a boolean if a field has been set.
func (o *BTLineDescription1559) HasOrigin() bool {
	if o != nil && o.Origin != nil {
		return true
	}

	return false
}

// SetOrigin gets a reference to the given BTVector3d389 and assigns it to the Origin field.
func (o *BTLineDescription1559) SetOrigin(v BTVector3d389) {
	o.Origin = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *BTLineDescription1559) GetType() GBTCurveTypeEnum {
	if o == nil || o.Type == nil {
		var ret GBTCurveTypeEnum
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTLineDescription1559) GetTypeOk() (*GBTCurveTypeEnum, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *BTLineDescription1559) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given GBTCurveTypeEnum and assigns it to the Type field.
func (o *BTLineDescription1559) SetType(v GBTCurveTypeEnum) {
	o.Type = &v
}

func (o BTLineDescription1559) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.Direction != nil {
		toSerialize["direction"] = o.Direction
	}
	if o.DirectionOrientedWithFace != nil {
		toSerialize["directionOrientedWithFace"] = o.DirectionOrientedWithFace
	}
	if o.Origin != nil {
		toSerialize["origin"] = o.Origin
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableBTLineDescription1559 struct {
	value *BTLineDescription1559
	isSet bool
}

func (v NullableBTLineDescription1559) Get() *BTLineDescription1559 {
	return v.value
}

func (v *NullableBTLineDescription1559) Set(val *BTLineDescription1559) {
	v.value = val
	v.isSet = true
}

func (v NullableBTLineDescription1559) IsSet() bool {
	return v.isSet
}

func (v *NullableBTLineDescription1559) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTLineDescription1559(val *BTLineDescription1559) *NullableBTLineDescription1559 {
	return &NullableBTLineDescription1559{value: val, isSet: true}
}

func (v NullableBTLineDescription1559) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTLineDescription1559) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
