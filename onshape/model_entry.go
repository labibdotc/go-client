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

// Entry struct for Entry
type Entry struct {
	PermissionSet []string        `json:"permissionSet,omitempty"`
	Role          *BTRbacRoleInfo `json:"role,omitempty"`
}

// NewEntry instantiates a new Entry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntry() *Entry {
	this := Entry{}
	return &this
}

// NewEntryWithDefaults instantiates a new Entry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntryWithDefaults() *Entry {
	this := Entry{}
	return &this
}

// GetPermissionSet returns the PermissionSet field value if set, zero value otherwise.
func (o *Entry) GetPermissionSet() []string {
	if o == nil || o.PermissionSet == nil {
		var ret []string
		return ret
	}
	return o.PermissionSet
}

// GetPermissionSetOk returns a tuple with the PermissionSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Entry) GetPermissionSetOk() ([]string, bool) {
	if o == nil || o.PermissionSet == nil {
		return nil, false
	}
	return o.PermissionSet, true
}

// HasPermissionSet returns a boolean if a field has been set.
func (o *Entry) HasPermissionSet() bool {
	if o != nil && o.PermissionSet != nil {
		return true
	}

	return false
}

// SetPermissionSet gets a reference to the given []string and assigns it to the PermissionSet field.
func (o *Entry) SetPermissionSet(v []string) {
	o.PermissionSet = v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *Entry) GetRole() BTRbacRoleInfo {
	if o == nil || o.Role == nil {
		var ret BTRbacRoleInfo
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Entry) GetRoleOk() (*BTRbacRoleInfo, bool) {
	if o == nil || o.Role == nil {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *Entry) HasRole() bool {
	if o != nil && o.Role != nil {
		return true
	}

	return false
}

// SetRole gets a reference to the given BTRbacRoleInfo and assigns it to the Role field.
func (o *Entry) SetRole(v BTRbacRoleInfo) {
	o.Role = &v
}

func (o Entry) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PermissionSet != nil {
		toSerialize["permissionSet"] = o.PermissionSet
	}
	if o.Role != nil {
		toSerialize["role"] = o.Role
	}
	return json.Marshal(toSerialize)
}

type NullableEntry struct {
	value *Entry
	isSet bool
}

func (v NullableEntry) Get() *Entry {
	return v.value
}

func (v *NullableEntry) Set(val *Entry) {
	v.value = val
	v.isSet = true
}

func (v NullableEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntry(val *Entry) *NullableEntry {
	return &NullableEntry{value: val, isSet: true}
}

func (v NullableEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
