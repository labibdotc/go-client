/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.154.6590-f8226b4e1789
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTUpdateMeshUnitsParams struct for BTUpdateMeshUnitsParams
type BTUpdateMeshUnitsParams struct {
	Units *string `json:"units,omitempty"`
}

// NewBTUpdateMeshUnitsParams instantiates a new BTUpdateMeshUnitsParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTUpdateMeshUnitsParams() *BTUpdateMeshUnitsParams {
	this := BTUpdateMeshUnitsParams{}
	return &this
}

// NewBTUpdateMeshUnitsParamsWithDefaults instantiates a new BTUpdateMeshUnitsParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTUpdateMeshUnitsParamsWithDefaults() *BTUpdateMeshUnitsParams {
	this := BTUpdateMeshUnitsParams{}
	return &this
}

// GetUnits returns the Units field value if set, zero value otherwise.
func (o *BTUpdateMeshUnitsParams) GetUnits() string {
	if o == nil || o.Units == nil {
		var ret string
		return ret
	}
	return *o.Units
}

// GetUnitsOk returns a tuple with the Units field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUpdateMeshUnitsParams) GetUnitsOk() (*string, bool) {
	if o == nil || o.Units == nil {
		return nil, false
	}
	return o.Units, true
}

// HasUnits returns a boolean if a field has been set.
func (o *BTUpdateMeshUnitsParams) HasUnits() bool {
	if o != nil && o.Units != nil {
		return true
	}

	return false
}

// SetUnits gets a reference to the given string and assigns it to the Units field.
func (o *BTUpdateMeshUnitsParams) SetUnits(v string) {
	o.Units = &v
}

func (o BTUpdateMeshUnitsParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Units != nil {
		toSerialize["units"] = o.Units
	}
	return json.Marshal(toSerialize)
}

type NullableBTUpdateMeshUnitsParams struct {
	value *BTUpdateMeshUnitsParams
	isSet bool
}

func (v NullableBTUpdateMeshUnitsParams) Get() *BTUpdateMeshUnitsParams {
	return v.value
}

func (v *NullableBTUpdateMeshUnitsParams) Set(val *BTUpdateMeshUnitsParams) {
	v.value = val
	v.isSet = true
}

func (v NullableBTUpdateMeshUnitsParams) IsSet() bool {
	return v.isSet
}

func (v *NullableBTUpdateMeshUnitsParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTUpdateMeshUnitsParams(val *BTUpdateMeshUnitsParams) *NullableBTUpdateMeshUnitsParams {
	return &NullableBTUpdateMeshUnitsParams{value: val, isSet: true}
}

func (v NullableBTUpdateMeshUnitsParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTUpdateMeshUnitsParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
