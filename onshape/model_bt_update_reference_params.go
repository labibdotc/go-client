/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.155.7328-9aa102855752
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTUpdateReferenceParams struct for BTUpdateReferenceParams
type BTUpdateReferenceParams struct {
	ConnectionId     *string        `json:"connectionId,omitempty"`
	EditDescription  *string        `json:"editDescription,omitempty"`
	ReferenceUpdates []UpdateParams `json:"referenceUpdates,omitempty"`
}

// NewBTUpdateReferenceParams instantiates a new BTUpdateReferenceParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTUpdateReferenceParams() *BTUpdateReferenceParams {
	this := BTUpdateReferenceParams{}
	return &this
}

// NewBTUpdateReferenceParamsWithDefaults instantiates a new BTUpdateReferenceParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTUpdateReferenceParamsWithDefaults() *BTUpdateReferenceParams {
	this := BTUpdateReferenceParams{}
	return &this
}

// GetConnectionId returns the ConnectionId field value if set, zero value otherwise.
func (o *BTUpdateReferenceParams) GetConnectionId() string {
	if o == nil || o.ConnectionId == nil {
		var ret string
		return ret
	}
	return *o.ConnectionId
}

// GetConnectionIdOk returns a tuple with the ConnectionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUpdateReferenceParams) GetConnectionIdOk() (*string, bool) {
	if o == nil || o.ConnectionId == nil {
		return nil, false
	}
	return o.ConnectionId, true
}

// HasConnectionId returns a boolean if a field has been set.
func (o *BTUpdateReferenceParams) HasConnectionId() bool {
	if o != nil && o.ConnectionId != nil {
		return true
	}

	return false
}

// SetConnectionId gets a reference to the given string and assigns it to the ConnectionId field.
func (o *BTUpdateReferenceParams) SetConnectionId(v string) {
	o.ConnectionId = &v
}

// GetEditDescription returns the EditDescription field value if set, zero value otherwise.
func (o *BTUpdateReferenceParams) GetEditDescription() string {
	if o == nil || o.EditDescription == nil {
		var ret string
		return ret
	}
	return *o.EditDescription
}

// GetEditDescriptionOk returns a tuple with the EditDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUpdateReferenceParams) GetEditDescriptionOk() (*string, bool) {
	if o == nil || o.EditDescription == nil {
		return nil, false
	}
	return o.EditDescription, true
}

// HasEditDescription returns a boolean if a field has been set.
func (o *BTUpdateReferenceParams) HasEditDescription() bool {
	if o != nil && o.EditDescription != nil {
		return true
	}

	return false
}

// SetEditDescription gets a reference to the given string and assigns it to the EditDescription field.
func (o *BTUpdateReferenceParams) SetEditDescription(v string) {
	o.EditDescription = &v
}

// GetReferenceUpdates returns the ReferenceUpdates field value if set, zero value otherwise.
func (o *BTUpdateReferenceParams) GetReferenceUpdates() []UpdateParams {
	if o == nil || o.ReferenceUpdates == nil {
		var ret []UpdateParams
		return ret
	}
	return o.ReferenceUpdates
}

// GetReferenceUpdatesOk returns a tuple with the ReferenceUpdates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUpdateReferenceParams) GetReferenceUpdatesOk() ([]UpdateParams, bool) {
	if o == nil || o.ReferenceUpdates == nil {
		return nil, false
	}
	return o.ReferenceUpdates, true
}

// HasReferenceUpdates returns a boolean if a field has been set.
func (o *BTUpdateReferenceParams) HasReferenceUpdates() bool {
	if o != nil && o.ReferenceUpdates != nil {
		return true
	}

	return false
}

// SetReferenceUpdates gets a reference to the given []UpdateParams and assigns it to the ReferenceUpdates field.
func (o *BTUpdateReferenceParams) SetReferenceUpdates(v []UpdateParams) {
	o.ReferenceUpdates = v
}

func (o BTUpdateReferenceParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ConnectionId != nil {
		toSerialize["connectionId"] = o.ConnectionId
	}
	if o.EditDescription != nil {
		toSerialize["editDescription"] = o.EditDescription
	}
	if o.ReferenceUpdates != nil {
		toSerialize["referenceUpdates"] = o.ReferenceUpdates
	}
	return json.Marshal(toSerialize)
}

type NullableBTUpdateReferenceParams struct {
	value *BTUpdateReferenceParams
	isSet bool
}

func (v NullableBTUpdateReferenceParams) Get() *BTUpdateReferenceParams {
	return v.value
}

func (v *NullableBTUpdateReferenceParams) Set(val *BTUpdateReferenceParams) {
	v.value = val
	v.isSet = true
}

func (v NullableBTUpdateReferenceParams) IsSet() bool {
	return v.isSet
}

func (v *NullableBTUpdateReferenceParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTUpdateReferenceParams(val *BTUpdateReferenceParams) *NullableBTUpdateReferenceParams {
	return &NullableBTUpdateReferenceParams{value: val, isSet: true}
}

func (v NullableBTUpdateReferenceParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTUpdateReferenceParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
