/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.167.20169-88260985a0b6
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTSimplePropertyInfo struct for BTSimplePropertyInfo
type BTSimplePropertyInfo struct {
	DisplayName  *string `json:"displayName,omitempty"`
	Frozen       *bool   `json:"frozen,omitempty"`
	PropertyId   *string `json:"propertyId,omitempty"`
	PublishState *int32  `json:"publishState,omitempty"`
}

// NewBTSimplePropertyInfo instantiates a new BTSimplePropertyInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTSimplePropertyInfo() *BTSimplePropertyInfo {
	this := BTSimplePropertyInfo{}
	return &this
}

// NewBTSimplePropertyInfoWithDefaults instantiates a new BTSimplePropertyInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTSimplePropertyInfoWithDefaults() *BTSimplePropertyInfo {
	this := BTSimplePropertyInfo{}
	return &this
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *BTSimplePropertyInfo) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSimplePropertyInfo) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *BTSimplePropertyInfo) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *BTSimplePropertyInfo) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetFrozen returns the Frozen field value if set, zero value otherwise.
func (o *BTSimplePropertyInfo) GetFrozen() bool {
	if o == nil || o.Frozen == nil {
		var ret bool
		return ret
	}
	return *o.Frozen
}

// GetFrozenOk returns a tuple with the Frozen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSimplePropertyInfo) GetFrozenOk() (*bool, bool) {
	if o == nil || o.Frozen == nil {
		return nil, false
	}
	return o.Frozen, true
}

// HasFrozen returns a boolean if a field has been set.
func (o *BTSimplePropertyInfo) HasFrozen() bool {
	if o != nil && o.Frozen != nil {
		return true
	}

	return false
}

// SetFrozen gets a reference to the given bool and assigns it to the Frozen field.
func (o *BTSimplePropertyInfo) SetFrozen(v bool) {
	o.Frozen = &v
}

// GetPropertyId returns the PropertyId field value if set, zero value otherwise.
func (o *BTSimplePropertyInfo) GetPropertyId() string {
	if o == nil || o.PropertyId == nil {
		var ret string
		return ret
	}
	return *o.PropertyId
}

// GetPropertyIdOk returns a tuple with the PropertyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSimplePropertyInfo) GetPropertyIdOk() (*string, bool) {
	if o == nil || o.PropertyId == nil {
		return nil, false
	}
	return o.PropertyId, true
}

// HasPropertyId returns a boolean if a field has been set.
func (o *BTSimplePropertyInfo) HasPropertyId() bool {
	if o != nil && o.PropertyId != nil {
		return true
	}

	return false
}

// SetPropertyId gets a reference to the given string and assigns it to the PropertyId field.
func (o *BTSimplePropertyInfo) SetPropertyId(v string) {
	o.PropertyId = &v
}

// GetPublishState returns the PublishState field value if set, zero value otherwise.
func (o *BTSimplePropertyInfo) GetPublishState() int32 {
	if o == nil || o.PublishState == nil {
		var ret int32
		return ret
	}
	return *o.PublishState
}

// GetPublishStateOk returns a tuple with the PublishState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSimplePropertyInfo) GetPublishStateOk() (*int32, bool) {
	if o == nil || o.PublishState == nil {
		return nil, false
	}
	return o.PublishState, true
}

// HasPublishState returns a boolean if a field has been set.
func (o *BTSimplePropertyInfo) HasPublishState() bool {
	if o != nil && o.PublishState != nil {
		return true
	}

	return false
}

// SetPublishState gets a reference to the given int32 and assigns it to the PublishState field.
func (o *BTSimplePropertyInfo) SetPublishState(v int32) {
	o.PublishState = &v
}

func (o BTSimplePropertyInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.Frozen != nil {
		toSerialize["frozen"] = o.Frozen
	}
	if o.PropertyId != nil {
		toSerialize["propertyId"] = o.PropertyId
	}
	if o.PublishState != nil {
		toSerialize["publishState"] = o.PublishState
	}
	return json.Marshal(toSerialize)
}

type NullableBTSimplePropertyInfo struct {
	value *BTSimplePropertyInfo
	isSet bool
}

func (v NullableBTSimplePropertyInfo) Get() *BTSimplePropertyInfo {
	return v.value
}

func (v *NullableBTSimplePropertyInfo) Set(val *BTSimplePropertyInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTSimplePropertyInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTSimplePropertyInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTSimplePropertyInfo(val *BTSimplePropertyInfo) *NullableBTSimplePropertyInfo {
	return &NullableBTSimplePropertyInfo{value: val, isSet: true}
}

func (v NullableBTSimplePropertyInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTSimplePropertyInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
