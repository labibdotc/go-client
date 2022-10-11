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
)

// BTClosedConstituentPartData2911 struct for BTClosedConstituentPartData2911
type BTClosedConstituentPartData2911 struct {
	BodyType           *string `json:"bodyType,omitempty"`
	IsActiveSheetMetal *bool   `json:"isActiveSheetMetal,omitempty"`
	IsMesh             *bool   `json:"isMesh,omitempty"`
	MeshState          *string `json:"meshState,omitempty"`
}

// NewBTClosedConstituentPartData2911 instantiates a new BTClosedConstituentPartData2911 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTClosedConstituentPartData2911() *BTClosedConstituentPartData2911 {
	this := BTClosedConstituentPartData2911{}
	return &this
}

// NewBTClosedConstituentPartData2911WithDefaults instantiates a new BTClosedConstituentPartData2911 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTClosedConstituentPartData2911WithDefaults() *BTClosedConstituentPartData2911 {
	this := BTClosedConstituentPartData2911{}
	return &this
}

// GetBodyType returns the BodyType field value if set, zero value otherwise.
func (o *BTClosedConstituentPartData2911) GetBodyType() string {
	if o == nil || o.BodyType == nil {
		var ret string
		return ret
	}
	return *o.BodyType
}

// GetBodyTypeOk returns a tuple with the BodyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTClosedConstituentPartData2911) GetBodyTypeOk() (*string, bool) {
	if o == nil || o.BodyType == nil {
		return nil, false
	}
	return o.BodyType, true
}

// HasBodyType returns a boolean if a field has been set.
func (o *BTClosedConstituentPartData2911) HasBodyType() bool {
	if o != nil && o.BodyType != nil {
		return true
	}

	return false
}

// SetBodyType gets a reference to the given string and assigns it to the BodyType field.
func (o *BTClosedConstituentPartData2911) SetBodyType(v string) {
	o.BodyType = &v
}

// GetIsActiveSheetMetal returns the IsActiveSheetMetal field value if set, zero value otherwise.
func (o *BTClosedConstituentPartData2911) GetIsActiveSheetMetal() bool {
	if o == nil || o.IsActiveSheetMetal == nil {
		var ret bool
		return ret
	}
	return *o.IsActiveSheetMetal
}

// GetIsActiveSheetMetalOk returns a tuple with the IsActiveSheetMetal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTClosedConstituentPartData2911) GetIsActiveSheetMetalOk() (*bool, bool) {
	if o == nil || o.IsActiveSheetMetal == nil {
		return nil, false
	}
	return o.IsActiveSheetMetal, true
}

// HasIsActiveSheetMetal returns a boolean if a field has been set.
func (o *BTClosedConstituentPartData2911) HasIsActiveSheetMetal() bool {
	if o != nil && o.IsActiveSheetMetal != nil {
		return true
	}

	return false
}

// SetIsActiveSheetMetal gets a reference to the given bool and assigns it to the IsActiveSheetMetal field.
func (o *BTClosedConstituentPartData2911) SetIsActiveSheetMetal(v bool) {
	o.IsActiveSheetMetal = &v
}

// GetIsMesh returns the IsMesh field value if set, zero value otherwise.
func (o *BTClosedConstituentPartData2911) GetIsMesh() bool {
	if o == nil || o.IsMesh == nil {
		var ret bool
		return ret
	}
	return *o.IsMesh
}

// GetIsMeshOk returns a tuple with the IsMesh field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTClosedConstituentPartData2911) GetIsMeshOk() (*bool, bool) {
	if o == nil || o.IsMesh == nil {
		return nil, false
	}
	return o.IsMesh, true
}

// HasIsMesh returns a boolean if a field has been set.
func (o *BTClosedConstituentPartData2911) HasIsMesh() bool {
	if o != nil && o.IsMesh != nil {
		return true
	}

	return false
}

// SetIsMesh gets a reference to the given bool and assigns it to the IsMesh field.
func (o *BTClosedConstituentPartData2911) SetIsMesh(v bool) {
	o.IsMesh = &v
}

// GetMeshState returns the MeshState field value if set, zero value otherwise.
func (o *BTClosedConstituentPartData2911) GetMeshState() string {
	if o == nil || o.MeshState == nil {
		var ret string
		return ret
	}
	return *o.MeshState
}

// GetMeshStateOk returns a tuple with the MeshState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTClosedConstituentPartData2911) GetMeshStateOk() (*string, bool) {
	if o == nil || o.MeshState == nil {
		return nil, false
	}
	return o.MeshState, true
}

// HasMeshState returns a boolean if a field has been set.
func (o *BTClosedConstituentPartData2911) HasMeshState() bool {
	if o != nil && o.MeshState != nil {
		return true
	}

	return false
}

// SetMeshState gets a reference to the given string and assigns it to the MeshState field.
func (o *BTClosedConstituentPartData2911) SetMeshState(v string) {
	o.MeshState = &v
}

func (o BTClosedConstituentPartData2911) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BodyType != nil {
		toSerialize["bodyType"] = o.BodyType
	}
	if o.IsActiveSheetMetal != nil {
		toSerialize["isActiveSheetMetal"] = o.IsActiveSheetMetal
	}
	if o.IsMesh != nil {
		toSerialize["isMesh"] = o.IsMesh
	}
	if o.MeshState != nil {
		toSerialize["meshState"] = o.MeshState
	}
	return json.Marshal(toSerialize)
}

type NullableBTClosedConstituentPartData2911 struct {
	value *BTClosedConstituentPartData2911
	isSet bool
}

func (v NullableBTClosedConstituentPartData2911) Get() *BTClosedConstituentPartData2911 {
	return v.value
}

func (v *NullableBTClosedConstituentPartData2911) Set(val *BTClosedConstituentPartData2911) {
	v.value = val
	v.isSet = true
}

func (v NullableBTClosedConstituentPartData2911) IsSet() bool {
	return v.isSet
}

func (v *NullableBTClosedConstituentPartData2911) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTClosedConstituentPartData2911(val *BTClosedConstituentPartData2911) *NullableBTClosedConstituentPartData2911 {
	return &NullableBTClosedConstituentPartData2911{value: val, isSet: true}
}

func (v NullableBTClosedConstituentPartData2911) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTClosedConstituentPartData2911) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
