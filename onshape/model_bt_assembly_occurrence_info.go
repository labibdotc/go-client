/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.162.14806-89d807e7089c
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTAssemblyOccurrenceInfo struct for BTAssemblyOccurrenceInfo
type BTAssemblyOccurrenceInfo struct {
	Fixed     *bool     `json:"fixed,omitempty"`
	Hidden    *bool     `json:"hidden,omitempty"`
	Path      []string  `json:"path,omitempty"`
	Transform []float64 `json:"transform,omitempty"`
}

// NewBTAssemblyOccurrenceInfo instantiates a new BTAssemblyOccurrenceInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTAssemblyOccurrenceInfo() *BTAssemblyOccurrenceInfo {
	this := BTAssemblyOccurrenceInfo{}
	return &this
}

// NewBTAssemblyOccurrenceInfoWithDefaults instantiates a new BTAssemblyOccurrenceInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTAssemblyOccurrenceInfoWithDefaults() *BTAssemblyOccurrenceInfo {
	this := BTAssemblyOccurrenceInfo{}
	return &this
}

// GetFixed returns the Fixed field value if set, zero value otherwise.
func (o *BTAssemblyOccurrenceInfo) GetFixed() bool {
	if o == nil || o.Fixed == nil {
		var ret bool
		return ret
	}
	return *o.Fixed
}

// GetFixedOk returns a tuple with the Fixed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyOccurrenceInfo) GetFixedOk() (*bool, bool) {
	if o == nil || o.Fixed == nil {
		return nil, false
	}
	return o.Fixed, true
}

// HasFixed returns a boolean if a field has been set.
func (o *BTAssemblyOccurrenceInfo) HasFixed() bool {
	if o != nil && o.Fixed != nil {
		return true
	}

	return false
}

// SetFixed gets a reference to the given bool and assigns it to the Fixed field.
func (o *BTAssemblyOccurrenceInfo) SetFixed(v bool) {
	o.Fixed = &v
}

// GetHidden returns the Hidden field value if set, zero value otherwise.
func (o *BTAssemblyOccurrenceInfo) GetHidden() bool {
	if o == nil || o.Hidden == nil {
		var ret bool
		return ret
	}
	return *o.Hidden
}

// GetHiddenOk returns a tuple with the Hidden field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyOccurrenceInfo) GetHiddenOk() (*bool, bool) {
	if o == nil || o.Hidden == nil {
		return nil, false
	}
	return o.Hidden, true
}

// HasHidden returns a boolean if a field has been set.
func (o *BTAssemblyOccurrenceInfo) HasHidden() bool {
	if o != nil && o.Hidden != nil {
		return true
	}

	return false
}

// SetHidden gets a reference to the given bool and assigns it to the Hidden field.
func (o *BTAssemblyOccurrenceInfo) SetHidden(v bool) {
	o.Hidden = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *BTAssemblyOccurrenceInfo) GetPath() []string {
	if o == nil || o.Path == nil {
		var ret []string
		return ret
	}
	return o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyOccurrenceInfo) GetPathOk() ([]string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *BTAssemblyOccurrenceInfo) HasPath() bool {
	if o != nil && o.Path != nil {
		return true
	}

	return false
}

// SetPath gets a reference to the given []string and assigns it to the Path field.
func (o *BTAssemblyOccurrenceInfo) SetPath(v []string) {
	o.Path = v
}

// GetTransform returns the Transform field value if set, zero value otherwise.
func (o *BTAssemblyOccurrenceInfo) GetTransform() []float64 {
	if o == nil || o.Transform == nil {
		var ret []float64
		return ret
	}
	return o.Transform
}

// GetTransformOk returns a tuple with the Transform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyOccurrenceInfo) GetTransformOk() ([]float64, bool) {
	if o == nil || o.Transform == nil {
		return nil, false
	}
	return o.Transform, true
}

// HasTransform returns a boolean if a field has been set.
func (o *BTAssemblyOccurrenceInfo) HasTransform() bool {
	if o != nil && o.Transform != nil {
		return true
	}

	return false
}

// SetTransform gets a reference to the given []float64 and assigns it to the Transform field.
func (o *BTAssemblyOccurrenceInfo) SetTransform(v []float64) {
	o.Transform = v
}

func (o BTAssemblyOccurrenceInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Fixed != nil {
		toSerialize["fixed"] = o.Fixed
	}
	if o.Hidden != nil {
		toSerialize["hidden"] = o.Hidden
	}
	if o.Path != nil {
		toSerialize["path"] = o.Path
	}
	if o.Transform != nil {
		toSerialize["transform"] = o.Transform
	}
	return json.Marshal(toSerialize)
}

type NullableBTAssemblyOccurrenceInfo struct {
	value *BTAssemblyOccurrenceInfo
	isSet bool
}

func (v NullableBTAssemblyOccurrenceInfo) Get() *BTAssemblyOccurrenceInfo {
	return v.value
}

func (v *NullableBTAssemblyOccurrenceInfo) Set(val *BTAssemblyOccurrenceInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTAssemblyOccurrenceInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTAssemblyOccurrenceInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTAssemblyOccurrenceInfo(val *BTAssemblyOccurrenceInfo) *NullableBTAssemblyOccurrenceInfo {
	return &NullableBTAssemblyOccurrenceInfo{value: val, isSet: true}
}

func (v NullableBTAssemblyOccurrenceInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTAssemblyOccurrenceInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
