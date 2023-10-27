/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.171.24804-920f3dc76f2b
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTDeviceLoginSecretInfo struct for BTDeviceLoginSecretInfo
type BTDeviceLoginSecretInfo struct {
	CreatedAt   *JSONTime `json:"createdAt,omitempty"`
	ModifiedAt  *JSONTime `json:"modifiedAt,omitempty"`
	RandomToken *string   `json:"randomToken,omitempty"`
}

// NewBTDeviceLoginSecretInfo instantiates a new BTDeviceLoginSecretInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTDeviceLoginSecretInfo() *BTDeviceLoginSecretInfo {
	this := BTDeviceLoginSecretInfo{}
	return &this
}

// NewBTDeviceLoginSecretInfoWithDefaults instantiates a new BTDeviceLoginSecretInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTDeviceLoginSecretInfoWithDefaults() *BTDeviceLoginSecretInfo {
	this := BTDeviceLoginSecretInfo{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *BTDeviceLoginSecretInfo) GetCreatedAt() JSONTime {
	if o == nil || o.CreatedAt == nil {
		var ret JSONTime
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDeviceLoginSecretInfo) GetCreatedAtOk() (*JSONTime, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *BTDeviceLoginSecretInfo) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given JSONTime and assigns it to the CreatedAt field.
func (o *BTDeviceLoginSecretInfo) SetCreatedAt(v JSONTime) {
	o.CreatedAt = &v
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise.
func (o *BTDeviceLoginSecretInfo) GetModifiedAt() JSONTime {
	if o == nil || o.ModifiedAt == nil {
		var ret JSONTime
		return ret
	}
	return *o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDeviceLoginSecretInfo) GetModifiedAtOk() (*JSONTime, bool) {
	if o == nil || o.ModifiedAt == nil {
		return nil, false
	}
	return o.ModifiedAt, true
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *BTDeviceLoginSecretInfo) HasModifiedAt() bool {
	if o != nil && o.ModifiedAt != nil {
		return true
	}

	return false
}

// SetModifiedAt gets a reference to the given JSONTime and assigns it to the ModifiedAt field.
func (o *BTDeviceLoginSecretInfo) SetModifiedAt(v JSONTime) {
	o.ModifiedAt = &v
}

// GetRandomToken returns the RandomToken field value if set, zero value otherwise.
func (o *BTDeviceLoginSecretInfo) GetRandomToken() string {
	if o == nil || o.RandomToken == nil {
		var ret string
		return ret
	}
	return *o.RandomToken
}

// GetRandomTokenOk returns a tuple with the RandomToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDeviceLoginSecretInfo) GetRandomTokenOk() (*string, bool) {
	if o == nil || o.RandomToken == nil {
		return nil, false
	}
	return o.RandomToken, true
}

// HasRandomToken returns a boolean if a field has been set.
func (o *BTDeviceLoginSecretInfo) HasRandomToken() bool {
	if o != nil && o.RandomToken != nil {
		return true
	}

	return false
}

// SetRandomToken gets a reference to the given string and assigns it to the RandomToken field.
func (o *BTDeviceLoginSecretInfo) SetRandomToken(v string) {
	o.RandomToken = &v
}

func (o BTDeviceLoginSecretInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CreatedAt != nil {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if o.ModifiedAt != nil {
		toSerialize["modifiedAt"] = o.ModifiedAt
	}
	if o.RandomToken != nil {
		toSerialize["randomToken"] = o.RandomToken
	}
	return json.Marshal(toSerialize)
}

type NullableBTDeviceLoginSecretInfo struct {
	value *BTDeviceLoginSecretInfo
	isSet bool
}

func (v NullableBTDeviceLoginSecretInfo) Get() *BTDeviceLoginSecretInfo {
	return v.value
}

func (v *NullableBTDeviceLoginSecretInfo) Set(val *BTDeviceLoginSecretInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTDeviceLoginSecretInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTDeviceLoginSecretInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTDeviceLoginSecretInfo(val *BTDeviceLoginSecretInfo) *NullableBTDeviceLoginSecretInfo {
	return &NullableBTDeviceLoginSecretInfo{value: val, isSet: true}
}

func (v NullableBTDeviceLoginSecretInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTDeviceLoginSecretInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
