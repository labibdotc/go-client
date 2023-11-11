/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.172.25784-25d5580b0721
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTDrawingModificationParams struct for BTDrawingModificationParams
type BTDrawingModificationParams struct {
	EditDescription *string                 `json:"editDescription,omitempty"`
	JsonRequests    *map[string]interface{} `json:"jsonRequests,omitempty"`
}

// NewBTDrawingModificationParams instantiates a new BTDrawingModificationParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTDrawingModificationParams() *BTDrawingModificationParams {
	this := BTDrawingModificationParams{}
	return &this
}

// NewBTDrawingModificationParamsWithDefaults instantiates a new BTDrawingModificationParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTDrawingModificationParamsWithDefaults() *BTDrawingModificationParams {
	this := BTDrawingModificationParams{}
	return &this
}

// GetEditDescription returns the EditDescription field value if set, zero value otherwise.
func (o *BTDrawingModificationParams) GetEditDescription() string {
	if o == nil || o.EditDescription == nil {
		var ret string
		return ret
	}
	return *o.EditDescription
}

// GetEditDescriptionOk returns a tuple with the EditDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDrawingModificationParams) GetEditDescriptionOk() (*string, bool) {
	if o == nil || o.EditDescription == nil {
		return nil, false
	}
	return o.EditDescription, true
}

// HasEditDescription returns a boolean if a field has been set.
func (o *BTDrawingModificationParams) HasEditDescription() bool {
	if o != nil && o.EditDescription != nil {
		return true
	}

	return false
}

// SetEditDescription gets a reference to the given string and assigns it to the EditDescription field.
func (o *BTDrawingModificationParams) SetEditDescription(v string) {
	o.EditDescription = &v
}

// GetJsonRequests returns the JsonRequests field value if set, zero value otherwise.
func (o *BTDrawingModificationParams) GetJsonRequests() map[string]interface{} {
	if o == nil || o.JsonRequests == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.JsonRequests
}

// GetJsonRequestsOk returns a tuple with the JsonRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDrawingModificationParams) GetJsonRequestsOk() (*map[string]interface{}, bool) {
	if o == nil || o.JsonRequests == nil {
		return nil, false
	}
	return o.JsonRequests, true
}

// HasJsonRequests returns a boolean if a field has been set.
func (o *BTDrawingModificationParams) HasJsonRequests() bool {
	if o != nil && o.JsonRequests != nil {
		return true
	}

	return false
}

// SetJsonRequests gets a reference to the given map[string]interface{} and assigns it to the JsonRequests field.
func (o *BTDrawingModificationParams) SetJsonRequests(v map[string]interface{}) {
	o.JsonRequests = &v
}

func (o BTDrawingModificationParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EditDescription != nil {
		toSerialize["editDescription"] = o.EditDescription
	}
	if o.JsonRequests != nil {
		toSerialize["jsonRequests"] = o.JsonRequests
	}
	return json.Marshal(toSerialize)
}

type NullableBTDrawingModificationParams struct {
	value *BTDrawingModificationParams
	isSet bool
}

func (v NullableBTDrawingModificationParams) Get() *BTDrawingModificationParams {
	return v.value
}

func (v *NullableBTDrawingModificationParams) Set(val *BTDrawingModificationParams) {
	v.value = val
	v.isSet = true
}

func (v NullableBTDrawingModificationParams) IsSet() bool {
	return v.isSet
}

func (v *NullableBTDrawingModificationParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTDrawingModificationParams(val *BTDrawingModificationParams) *NullableBTDrawingModificationParams {
	return &NullableBTDrawingModificationParams{value: val, isSet: true}
}

func (v NullableBTDrawingModificationParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTDrawingModificationParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
