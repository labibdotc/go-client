/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.152.6268-501b86d3d653
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTRbacPermissionSchemeInfo struct for BTRbacPermissionSchemeInfo
type BTRbacPermissionSchemeInfo struct {
	Active                     *bool   `json:"active,omitempty"`
	Description                *string `json:"description,omitempty"`
	Entries                    []Entry `json:"entries,omitempty"`
	Href                       *string `json:"href,omitempty"`
	Id                         *string `json:"id,omitempty"`
	Name                       *string `json:"name,omitempty"`
	PredefinedPermissionScheme *int32  `json:"predefinedPermissionScheme,omitempty"`
	ViewRef                    *string `json:"viewRef,omitempty"`
}

// NewBTRbacPermissionSchemeInfo instantiates a new BTRbacPermissionSchemeInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTRbacPermissionSchemeInfo() *BTRbacPermissionSchemeInfo {
	this := BTRbacPermissionSchemeInfo{}
	return &this
}

// NewBTRbacPermissionSchemeInfoWithDefaults instantiates a new BTRbacPermissionSchemeInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTRbacPermissionSchemeInfoWithDefaults() *BTRbacPermissionSchemeInfo {
	this := BTRbacPermissionSchemeInfo{}
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *BTRbacPermissionSchemeInfo) GetActive() bool {
	if o == nil || o.Active == nil {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRbacPermissionSchemeInfo) GetActiveOk() (*bool, bool) {
	if o == nil || o.Active == nil {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *BTRbacPermissionSchemeInfo) HasActive() bool {
	if o != nil && o.Active != nil {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *BTRbacPermissionSchemeInfo) SetActive(v bool) {
	o.Active = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BTRbacPermissionSchemeInfo) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRbacPermissionSchemeInfo) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BTRbacPermissionSchemeInfo) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BTRbacPermissionSchemeInfo) SetDescription(v string) {
	o.Description = &v
}

// GetEntries returns the Entries field value if set, zero value otherwise.
func (o *BTRbacPermissionSchemeInfo) GetEntries() []Entry {
	if o == nil || o.Entries == nil {
		var ret []Entry
		return ret
	}
	return o.Entries
}

// GetEntriesOk returns a tuple with the Entries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRbacPermissionSchemeInfo) GetEntriesOk() ([]Entry, bool) {
	if o == nil || o.Entries == nil {
		return nil, false
	}
	return o.Entries, true
}

// HasEntries returns a boolean if a field has been set.
func (o *BTRbacPermissionSchemeInfo) HasEntries() bool {
	if o != nil && o.Entries != nil {
		return true
	}

	return false
}

// SetEntries gets a reference to the given []Entry and assigns it to the Entries field.
func (o *BTRbacPermissionSchemeInfo) SetEntries(v []Entry) {
	o.Entries = v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *BTRbacPermissionSchemeInfo) GetHref() string {
	if o == nil || o.Href == nil {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRbacPermissionSchemeInfo) GetHrefOk() (*string, bool) {
	if o == nil || o.Href == nil {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *BTRbacPermissionSchemeInfo) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *BTRbacPermissionSchemeInfo) SetHref(v string) {
	o.Href = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BTRbacPermissionSchemeInfo) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRbacPermissionSchemeInfo) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BTRbacPermissionSchemeInfo) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BTRbacPermissionSchemeInfo) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BTRbacPermissionSchemeInfo) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRbacPermissionSchemeInfo) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BTRbacPermissionSchemeInfo) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BTRbacPermissionSchemeInfo) SetName(v string) {
	o.Name = &v
}

// GetPredefinedPermissionScheme returns the PredefinedPermissionScheme field value if set, zero value otherwise.
func (o *BTRbacPermissionSchemeInfo) GetPredefinedPermissionScheme() int32 {
	if o == nil || o.PredefinedPermissionScheme == nil {
		var ret int32
		return ret
	}
	return *o.PredefinedPermissionScheme
}

// GetPredefinedPermissionSchemeOk returns a tuple with the PredefinedPermissionScheme field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRbacPermissionSchemeInfo) GetPredefinedPermissionSchemeOk() (*int32, bool) {
	if o == nil || o.PredefinedPermissionScheme == nil {
		return nil, false
	}
	return o.PredefinedPermissionScheme, true
}

// HasPredefinedPermissionScheme returns a boolean if a field has been set.
func (o *BTRbacPermissionSchemeInfo) HasPredefinedPermissionScheme() bool {
	if o != nil && o.PredefinedPermissionScheme != nil {
		return true
	}

	return false
}

// SetPredefinedPermissionScheme gets a reference to the given int32 and assigns it to the PredefinedPermissionScheme field.
func (o *BTRbacPermissionSchemeInfo) SetPredefinedPermissionScheme(v int32) {
	o.PredefinedPermissionScheme = &v
}

// GetViewRef returns the ViewRef field value if set, zero value otherwise.
func (o *BTRbacPermissionSchemeInfo) GetViewRef() string {
	if o == nil || o.ViewRef == nil {
		var ret string
		return ret
	}
	return *o.ViewRef
}

// GetViewRefOk returns a tuple with the ViewRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRbacPermissionSchemeInfo) GetViewRefOk() (*string, bool) {
	if o == nil || o.ViewRef == nil {
		return nil, false
	}
	return o.ViewRef, true
}

// HasViewRef returns a boolean if a field has been set.
func (o *BTRbacPermissionSchemeInfo) HasViewRef() bool {
	if o != nil && o.ViewRef != nil {
		return true
	}

	return false
}

// SetViewRef gets a reference to the given string and assigns it to the ViewRef field.
func (o *BTRbacPermissionSchemeInfo) SetViewRef(v string) {
	o.ViewRef = &v
}

func (o BTRbacPermissionSchemeInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Active != nil {
		toSerialize["active"] = o.Active
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Entries != nil {
		toSerialize["entries"] = o.Entries
	}
	if o.Href != nil {
		toSerialize["href"] = o.Href
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.PredefinedPermissionScheme != nil {
		toSerialize["predefinedPermissionScheme"] = o.PredefinedPermissionScheme
	}
	if o.ViewRef != nil {
		toSerialize["viewRef"] = o.ViewRef
	}
	return json.Marshal(toSerialize)
}

type NullableBTRbacPermissionSchemeInfo struct {
	value *BTRbacPermissionSchemeInfo
	isSet bool
}

func (v NullableBTRbacPermissionSchemeInfo) Get() *BTRbacPermissionSchemeInfo {
	return v.value
}

func (v *NullableBTRbacPermissionSchemeInfo) Set(val *BTRbacPermissionSchemeInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTRbacPermissionSchemeInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTRbacPermissionSchemeInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTRbacPermissionSchemeInfo(val *BTRbacPermissionSchemeInfo) *NullableBTRbacPermissionSchemeInfo {
	return &NullableBTRbacPermissionSchemeInfo{value: val, isSet: true}
}

func (v NullableBTRbacPermissionSchemeInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTRbacPermissionSchemeInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
