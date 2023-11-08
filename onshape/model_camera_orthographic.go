/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.172.25541-b54c5b71ef45
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// CameraOrthographic struct for CameraOrthographic
type CameraOrthographic struct {
	Extensions map[string]map[string]interface{} `json:"extensions,omitempty"`
	Extras     *map[string]interface{}           `json:"extras,omitempty"`
	Xmag       *float32                          `json:"xmag,omitempty"`
	Ymag       *float32                          `json:"ymag,omitempty"`
	Zfar       *float32                          `json:"zfar,omitempty"`
	Znear      *float32                          `json:"znear,omitempty"`
}

// NewCameraOrthographic instantiates a new CameraOrthographic object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCameraOrthographic() *CameraOrthographic {
	this := CameraOrthographic{}
	return &this
}

// NewCameraOrthographicWithDefaults instantiates a new CameraOrthographic object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCameraOrthographicWithDefaults() *CameraOrthographic {
	this := CameraOrthographic{}
	return &this
}

// GetExtensions returns the Extensions field value if set, zero value otherwise.
func (o *CameraOrthographic) GetExtensions() map[string]map[string]interface{} {
	if o == nil || o.Extensions == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.Extensions
}

// GetExtensionsOk returns a tuple with the Extensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CameraOrthographic) GetExtensionsOk() (map[string]map[string]interface{}, bool) {
	if o == nil || o.Extensions == nil {
		return nil, false
	}
	return o.Extensions, true
}

// HasExtensions returns a boolean if a field has been set.
func (o *CameraOrthographic) HasExtensions() bool {
	if o != nil && o.Extensions != nil {
		return true
	}

	return false
}

// SetExtensions gets a reference to the given map[string]map[string]interface{} and assigns it to the Extensions field.
func (o *CameraOrthographic) SetExtensions(v map[string]map[string]interface{}) {
	o.Extensions = v
}

// GetExtras returns the Extras field value if set, zero value otherwise.
func (o *CameraOrthographic) GetExtras() map[string]interface{} {
	if o == nil || o.Extras == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Extras
}

// GetExtrasOk returns a tuple with the Extras field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CameraOrthographic) GetExtrasOk() (*map[string]interface{}, bool) {
	if o == nil || o.Extras == nil {
		return nil, false
	}
	return o.Extras, true
}

// HasExtras returns a boolean if a field has been set.
func (o *CameraOrthographic) HasExtras() bool {
	if o != nil && o.Extras != nil {
		return true
	}

	return false
}

// SetExtras gets a reference to the given map[string]interface{} and assigns it to the Extras field.
func (o *CameraOrthographic) SetExtras(v map[string]interface{}) {
	o.Extras = &v
}

// GetXmag returns the Xmag field value if set, zero value otherwise.
func (o *CameraOrthographic) GetXmag() float32 {
	if o == nil || o.Xmag == nil {
		var ret float32
		return ret
	}
	return *o.Xmag
}

// GetXmagOk returns a tuple with the Xmag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CameraOrthographic) GetXmagOk() (*float32, bool) {
	if o == nil || o.Xmag == nil {
		return nil, false
	}
	return o.Xmag, true
}

// HasXmag returns a boolean if a field has been set.
func (o *CameraOrthographic) HasXmag() bool {
	if o != nil && o.Xmag != nil {
		return true
	}

	return false
}

// SetXmag gets a reference to the given float32 and assigns it to the Xmag field.
func (o *CameraOrthographic) SetXmag(v float32) {
	o.Xmag = &v
}

// GetYmag returns the Ymag field value if set, zero value otherwise.
func (o *CameraOrthographic) GetYmag() float32 {
	if o == nil || o.Ymag == nil {
		var ret float32
		return ret
	}
	return *o.Ymag
}

// GetYmagOk returns a tuple with the Ymag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CameraOrthographic) GetYmagOk() (*float32, bool) {
	if o == nil || o.Ymag == nil {
		return nil, false
	}
	return o.Ymag, true
}

// HasYmag returns a boolean if a field has been set.
func (o *CameraOrthographic) HasYmag() bool {
	if o != nil && o.Ymag != nil {
		return true
	}

	return false
}

// SetYmag gets a reference to the given float32 and assigns it to the Ymag field.
func (o *CameraOrthographic) SetYmag(v float32) {
	o.Ymag = &v
}

// GetZfar returns the Zfar field value if set, zero value otherwise.
func (o *CameraOrthographic) GetZfar() float32 {
	if o == nil || o.Zfar == nil {
		var ret float32
		return ret
	}
	return *o.Zfar
}

// GetZfarOk returns a tuple with the Zfar field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CameraOrthographic) GetZfarOk() (*float32, bool) {
	if o == nil || o.Zfar == nil {
		return nil, false
	}
	return o.Zfar, true
}

// HasZfar returns a boolean if a field has been set.
func (o *CameraOrthographic) HasZfar() bool {
	if o != nil && o.Zfar != nil {
		return true
	}

	return false
}

// SetZfar gets a reference to the given float32 and assigns it to the Zfar field.
func (o *CameraOrthographic) SetZfar(v float32) {
	o.Zfar = &v
}

// GetZnear returns the Znear field value if set, zero value otherwise.
func (o *CameraOrthographic) GetZnear() float32 {
	if o == nil || o.Znear == nil {
		var ret float32
		return ret
	}
	return *o.Znear
}

// GetZnearOk returns a tuple with the Znear field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CameraOrthographic) GetZnearOk() (*float32, bool) {
	if o == nil || o.Znear == nil {
		return nil, false
	}
	return o.Znear, true
}

// HasZnear returns a boolean if a field has been set.
func (o *CameraOrthographic) HasZnear() bool {
	if o != nil && o.Znear != nil {
		return true
	}

	return false
}

// SetZnear gets a reference to the given float32 and assigns it to the Znear field.
func (o *CameraOrthographic) SetZnear(v float32) {
	o.Znear = &v
}

func (o CameraOrthographic) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Extensions != nil {
		toSerialize["extensions"] = o.Extensions
	}
	if o.Extras != nil {
		toSerialize["extras"] = o.Extras
	}
	if o.Xmag != nil {
		toSerialize["xmag"] = o.Xmag
	}
	if o.Ymag != nil {
		toSerialize["ymag"] = o.Ymag
	}
	if o.Zfar != nil {
		toSerialize["zfar"] = o.Zfar
	}
	if o.Znear != nil {
		toSerialize["znear"] = o.Znear
	}
	return json.Marshal(toSerialize)
}

type NullableCameraOrthographic struct {
	value *CameraOrthographic
	isSet bool
}

func (v NullableCameraOrthographic) Get() *CameraOrthographic {
	return v.value
}

func (v *NullableCameraOrthographic) Set(val *CameraOrthographic) {
	v.value = val
	v.isSet = true
}

func (v NullableCameraOrthographic) IsSet() bool {
	return v.isSet
}

func (v *NullableCameraOrthographic) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCameraOrthographic(val *CameraOrthographic) *NullableCameraOrthographic {
	return &NullableCameraOrthographic{value: val, isSet: true}
}

func (v NullableCameraOrthographic) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCameraOrthographic) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
