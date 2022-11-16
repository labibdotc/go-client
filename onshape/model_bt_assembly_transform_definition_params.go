/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.155.7489-3fe01473c812
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTAssemblyTransformDefinitionParams struct for BTAssemblyTransformDefinitionParams
type BTAssemblyTransformDefinitionParams struct {
	IsRelative  *bool            `json:"isRelative,omitempty"`
	Occurrences []BTOccurrence74 `json:"occurrences,omitempty"`
	Transform   []float64        `json:"transform,omitempty"`
}

// NewBTAssemblyTransformDefinitionParams instantiates a new BTAssemblyTransformDefinitionParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTAssemblyTransformDefinitionParams() *BTAssemblyTransformDefinitionParams {
	this := BTAssemblyTransformDefinitionParams{}
	return &this
}

// NewBTAssemblyTransformDefinitionParamsWithDefaults instantiates a new BTAssemblyTransformDefinitionParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTAssemblyTransformDefinitionParamsWithDefaults() *BTAssemblyTransformDefinitionParams {
	this := BTAssemblyTransformDefinitionParams{}
	return &this
}

// GetIsRelative returns the IsRelative field value if set, zero value otherwise.
func (o *BTAssemblyTransformDefinitionParams) GetIsRelative() bool {
	if o == nil || o.IsRelative == nil {
		var ret bool
		return ret
	}
	return *o.IsRelative
}

// GetIsRelativeOk returns a tuple with the IsRelative field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyTransformDefinitionParams) GetIsRelativeOk() (*bool, bool) {
	if o == nil || o.IsRelative == nil {
		return nil, false
	}
	return o.IsRelative, true
}

// HasIsRelative returns a boolean if a field has been set.
func (o *BTAssemblyTransformDefinitionParams) HasIsRelative() bool {
	if o != nil && o.IsRelative != nil {
		return true
	}

	return false
}

// SetIsRelative gets a reference to the given bool and assigns it to the IsRelative field.
func (o *BTAssemblyTransformDefinitionParams) SetIsRelative(v bool) {
	o.IsRelative = &v
}

// GetOccurrences returns the Occurrences field value if set, zero value otherwise.
func (o *BTAssemblyTransformDefinitionParams) GetOccurrences() []BTOccurrence74 {
	if o == nil || o.Occurrences == nil {
		var ret []BTOccurrence74
		return ret
	}
	return o.Occurrences
}

// GetOccurrencesOk returns a tuple with the Occurrences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyTransformDefinitionParams) GetOccurrencesOk() ([]BTOccurrence74, bool) {
	if o == nil || o.Occurrences == nil {
		return nil, false
	}
	return o.Occurrences, true
}

// HasOccurrences returns a boolean if a field has been set.
func (o *BTAssemblyTransformDefinitionParams) HasOccurrences() bool {
	if o != nil && o.Occurrences != nil {
		return true
	}

	return false
}

// SetOccurrences gets a reference to the given []BTOccurrence74 and assigns it to the Occurrences field.
func (o *BTAssemblyTransformDefinitionParams) SetOccurrences(v []BTOccurrence74) {
	o.Occurrences = v
}

// GetTransform returns the Transform field value if set, zero value otherwise.
func (o *BTAssemblyTransformDefinitionParams) GetTransform() []float64 {
	if o == nil || o.Transform == nil {
		var ret []float64
		return ret
	}
	return o.Transform
}

// GetTransformOk returns a tuple with the Transform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyTransformDefinitionParams) GetTransformOk() ([]float64, bool) {
	if o == nil || o.Transform == nil {
		return nil, false
	}
	return o.Transform, true
}

// HasTransform returns a boolean if a field has been set.
func (o *BTAssemblyTransformDefinitionParams) HasTransform() bool {
	if o != nil && o.Transform != nil {
		return true
	}

	return false
}

// SetTransform gets a reference to the given []float64 and assigns it to the Transform field.
func (o *BTAssemblyTransformDefinitionParams) SetTransform(v []float64) {
	o.Transform = v
}

func (o BTAssemblyTransformDefinitionParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IsRelative != nil {
		toSerialize["isRelative"] = o.IsRelative
	}
	if o.Occurrences != nil {
		toSerialize["occurrences"] = o.Occurrences
	}
	if o.Transform != nil {
		toSerialize["transform"] = o.Transform
	}
	return json.Marshal(toSerialize)
}

type NullableBTAssemblyTransformDefinitionParams struct {
	value *BTAssemblyTransformDefinitionParams
	isSet bool
}

func (v NullableBTAssemblyTransformDefinitionParams) Get() *BTAssemblyTransformDefinitionParams {
	return v.value
}

func (v *NullableBTAssemblyTransformDefinitionParams) Set(val *BTAssemblyTransformDefinitionParams) {
	v.value = val
	v.isSet = true
}

func (v NullableBTAssemblyTransformDefinitionParams) IsSet() bool {
	return v.isSet
}

func (v *NullableBTAssemblyTransformDefinitionParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTAssemblyTransformDefinitionParams(val *BTAssemblyTransformDefinitionParams) *NullableBTAssemblyTransformDefinitionParams {
	return &NullableBTAssemblyTransformDefinitionParams{value: val, isSet: true}
}

func (v NullableBTAssemblyTransformDefinitionParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTAssemblyTransformDefinitionParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
